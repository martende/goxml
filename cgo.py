from pyclibrary import CParser
import os.path
import os
import sys
import re
FUNC_DESCS = (
	('xmlCreatePushParserCtxt','void *','user_data'),('SKIP',)
)
getHandler =  lambda n,t:"\tc_%s := %s.handler" % (n,n)
getConverter = lambda n,t:'\tc_%s := C.%s(%s)' % (n,t,n)
def getNullOrHandler(n,t): 
	a = """\tvar c_%(n)s C.%(t)s=nil ;if %(n)s !=nil { c_%(n)s = %(n)s.handler }""" % {'n':n,'t':t}
	return a
	
TYPEINFO = {
	'xmlNodePtr' : {
		'goArgType' : '*XmlNode',
		'go2cConverter' : getNullOrHandler,
		'c2goConverter' : retNullOrObject,
	},
	'xmlDocPtr' : {
		'goArgType' : '*XmlDoc',
		'go2cConverter' : getNullOrHandler
	},
	'int' : {
		'goArgType' : 'int',
		'go2cConverter' : getConverter
	},
}

INCLUDES = (
	"/usr/include/libxml2/libxml/tree.h",
	"/usr/include/libxml2/libxml/parser.h",
	"/usr/include/libxml2/libxml/xmlreader.h",
	"/usr/include/libxml2/libxml/xmlmemory.h",
	"/usr/include/libxml2/libxml/xmlstring.h",
	"/usr/include/libxml2/libxml/HTMLparser.h"
);

IMPORTS = (
	'xmlDocCopyNode',
	'UTF8ToHtml',
	'xmlMemBlocks',
	'xmlStrlen',
	'xmlFreeTextReader',
	'xmlTextReaderConstName',
	'xmlTextReaderConstValue',
	'xmlTextReaderDepth',
	'xmlTextReaderNodeType',
	'xmlTextReaderRead',
	'xmlTextReaderIsEmptyElement',
	'xmlTextReaderHasValue',
	'xmlReaderForFile','xmlParseChunk','xmlReadFile','xmlReadMemory','xmlFreeDoc','xmlAddChild','xmlCleanupParser','xmlMemoryDump','xmlNewParserCtxt','xmlFreeParserCtxt','xmlCtxtReadFile','xmlCreatePushParserCtxt');

TMP = "/tmp/tmp1"

GO_TPL = """package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/%(filename)s>
*/
import "C"
%(unsafe_import)s

%(consts_list)s
%(funcs_list)s

"""

def lookInDb(fname,ptype,pname):
	for i, v in enumerate(FUNC_DESCS):
		if v[0] == (fname,ptype,pname):
			return v[1]
	return None


def createFuncArgString(fname,sig):
	#print fname,sig
	goFname = fname[0].upper() + fname[1:]
	cReturnType = "".join(sig[0]) 
	args = []
	errs = []
	goReturnType = None
	try:
		goReturnType = TYPEINFO[cReturnType]['goReturnType']
	except:
		try:
			goReturnType = TYPEINFO[cReturnType]['goArgType']
		except:
			#goArgType = ptype
			goReturnType = cReturnType
			errs.append('Warn: ReturnType %s Not defined' % (goReturnType))
	
	for (pname,ptype,_) in sig[1]:
		
		
		ptype = "".join(ptype)
		
		dbData = lookInDb(fname,ptype,pname)
		goArgType = None
		if dbData:
			raise ('Not implemented')
		else:
			try:
				goArgType = TYPEINFO[ptype]['goArgType']
			except:
				goArgType = ptype
				errs.append('Warn: %s %s Not defined' % (pname,ptype))
				
		args.append("%s %s" % (pname,goArgType))
	
	arg_string ="func " + goFname + "(" + ",".join(args) +") "  + goReturnType
	
	return arg_string,errs

def createInputInits(fname,sig):
	errs = []
	outs = []
	for (pname,ptype,_) in sig[1]:
		ptype = "".join(ptype)
		dbData = lookInDb(fname,ptype,pname)
		go2cConvert = None
		if dbData:
			raise ('Not implemented')
		try:
			go2cConvert = TYPEINFO[ptype]['go2cConverter']
		except:
			go2cConvert = lambda a,b:pname
			errs.append('Warn: %s %s No converter to C' % (pname,ptype))
		
		outs.append(go2cConvert(pname,ptype) )
	
	return "\n".join(outs),errs

def createCallLine(fName,sig):
	cReturnType = "".join(sig[0])
	errs = []
	outs = []
	callargs = []
	for (pname,ptype,_) in sig[1]:
		callargs.append("c_" + pname)
	callLine = 'C.' + fName + "(" + ",".join(callargs)+")"
	if ( cReturnType != 'void') :
		callLine = "\tc_ret := " + callLine 
		return "\n".join((callLine,)),errs

def processFuncsList(functionSignatures):
	for fName in functionSignatures:
		errs = []
		sig = functionSignatures[fName]
		funcComment = """/* 
   Function: %s
   ReturnType: %s
   Args: %s
*/
"""	% ( fName , "".join(sig[0]) , sig[1]) 
		#print fName,functionSignatures[fName]
		funcCode = funcComment
		funArgString,err = createFuncArgString(fName,sig)
		inputInits,err = createInputInits(fName,sig)
		callLine,err = createCallLine(fName,sig)
		errs+=err
		
		if fName not in IMPORTS  or errs:
			funcCode+="/*\n"
		if errs:
			funcCode+= "\n\t" + "\n\t".join(errs) + "\n\n"
		
		funcCode+=funArgString +"\n{\n"
		
		funcCode+=inputInits + "\n"
		
		funcCode+=callLine + "\n";
		
		funcCode+="}\n"
		
		if fName not in IMPORTS or errs:
			funcCode+="\n*/\n"
		
		return funcCode
def processFile(include):
	filename = include.split("/")[-1] 
	p = CParser(include,cache=TMP + "/" + filename + ".cache",macros={'XMLCALL': ''})
	gofilename = re.sub(r'\.h$',".go",filename)
	#p.printAll()
	varsdict = {
		'filename' : filename,
		'unsafe_import' : '',
		'consts_list' : '',
		'funcs_list' : '',
	}
	varsdict['funcs_list'] = processFuncsList(p.defs['functions']) 
	print  GO_TPL % varsdict
	
	sys.exit(0)
	
if not os.path.exists(TMP):
	os.mkdir(TMP) 


for include in INCLUDES:
	processFile(include)
	
#p.printAll()

from pyclibrary import CParser
import os.path
import os
import sys
import re
import optparse

FUNC_DESCS = (
	('f','xmlCreatePushParserCtxt','void *','user_data'),('SKIP',),
	('s','xmlDocPtr',None,'_private'),('PRIVATE'),
	('s','xmlDocPtr',None,'ids'),('PRIVATE'),
	('s','xmlDocPtr',None,'refs'),('PRIVATE'),
	('s','xmlDocPtr',None,'psvi'),('PRIVATE'),
	('s','xmlDtdPtr','void*',None),('PRIVATE'),
	('s',None,'void*','userData'),('PRIVATE'),
	('s',None,None,'type'),('PRIVATE'),
	('s',None,None,'_private'),('PRIVATE'),
	
)
getHandler =  lambda n,t:"\tc_%s := %s.handler" % (n,n)
getConverter = lambda n,t:'\tc_%s := C.%s(%s)' % (n,t,n)

def toCharConverter(n,t):
	return """\tc_%s:= C.CString(%s)""" % (n,n)

def getNullOrHandler(n,t): 
	return """\tvar c_%(n)s C.%(t)s=nil ;if %(n)s !=nil { c_%(n)s = %(n)s.handler }""" % {'n':n,'t':t}
def retNullOrObject(t,goRetType):
	return 'if c_ret == nil {return nil}\n\treturn &%s{handler:c_ret}' % goRetType  
def retObject(t,goRetType):
	return "return %s(c_ret)" % t

TYPEALIAS = {
	'struct _xmlNode*' : 'xmlNodePtr',
	'struct _xmlDoc*' : 'xmlDocPtr',
	'struct _xmlDtd*' : 'xmlDtdPtr',
	'struct _xmlDict*' : 'xmlDictPtr',
	'struct _xmlNs*' : 'xmlNsPtr',
	'struct _xmlSAXHandler*' : 'xmlSAXHandlerPtr',
}
TYPEINFO = {
	'xmlDtdPtr' : {
		'goArgType' : '*XmlDtd',
		'exportStruct' : '_xmlDtd',
	},
	'xmlDictPtr' : {
		'goArgType' : '*XmlDict',
		'exportStruct' : '_xmlDict',
		'go2cConverter' : getNullOrHandler,
		'c2goConverter' : retNullOrObject,
	},
	'xmlNsPtr' : {
		'goArgType' : '*XmlNs',
		'exportStruct' : '_xmlNs',
	},
	'xmlElementType': {
		'goArgType':'int',
	},
	'xmlChar*' : {
		'goArgType':'string',
	},
	
	'char*' : {
		'goArgType':'string',
		'go2cConverter':toCharConverter,
	},
	'xmlParserInputPtr' : {
		'goArgType' : '*XmlParserInput',
		'go2cConverter' : getNullOrHandler,
		'c2goConverter' : retNullOrObject,
		'exportStruct' : '_xmlParserInput'
	},
	'xmlParserCtxtPtr': {
		'goArgType' : '*XmlParserCtxt',
		'go2cConverter' : getNullOrHandler,
		'c2goConverter' : retNullOrObject,
		'exportStruct' : '_xmlParserCtxt'
	},
	'xmlNodePtr' : {
		'goArgType' : '*XmlNode',
		'go2cConverter' : getNullOrHandler,
		'c2goConverter' : retNullOrObject,
	},
	'xmlDocPtr' : {
		'goArgType' : '*XmlDoc',
		'go2cConverter' : getNullOrHandler,
		'c2goConverter' : retNullOrObject,
		'exportStruct' : '_xmlDoc'
	},
	'xmlSAXHandlerPtr' : {
		'goArgType' : '*XmlSAXHandler',
		'go2cConverter' : getNullOrHandler,
		'c2goConverter' : retNullOrObject,
		'exportStruct' : '_xmlSAXHandler'
	},
	'int' : {
		'goArgType' : 'int',
		'go2cConverter' : getConverter,
		'c2goConverter' : retObject,
	},
	'void' : {
		'goReturnType':''
	}
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
%(structs_list)s
%(funcs_list)s

"""

CONF_PRINT_SKIPPED = False

def lookInDb(stype,fname,ptype,pname):
	for i in range(0,len(FUNC_DESCS),2):
		v=FUNC_DESCS[i]
		t = [stype,fname,ptype,pname]
		for j,tv in enumerate(v):
			if tv is None:
				t[j]=None
		if v == tuple(t):
			return FUNC_DESCS[i+1]
	return None

class FileConverter():
	def __init__(self,include):
		self.filename = include.split("/")[-1]
		self.include = include
		self.gofilename = re.sub(r'\.h$',".go",self.filename)
		self.unsafe = False
	def createFuncArgString(self,fname,sig):
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
			if pname is not None:
				ptype = "".join(ptype)
				dbData = lookInDb('f',fname,ptype,pname)
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
		
		arg_string ="func " + goFname + "(" + ",".join(args) +")"
		if goReturnType:
			arg_string+= " "+goReturnType
		
		return arg_string,errs
	
	def createInputInits(self,fname,sig):
		errs = []
		outs = []
		for (pname,ptype,_) in sig[1]:
			if pname is not None:
				ptype = "".join(ptype)
				dbData = lookInDb('f',fname,ptype,pname)
				go2cConvert = None
				if dbData:
					raise ('Not implemented')
				try:
					go2cConvert = TYPEINFO[ptype]['go2cConverter']
				except:
					go2cConvert = lambda a,b:pname
					errs.append('Warn: %s %s No converter to C(go2cConverter)' % (pname,ptype))
				
				outs.append(str(go2cConvert(pname,ptype)) )
			
		return "\n".join(outs),errs
	
	def createCallLine(self,fName,sig):
		
		cReturnType = "".join(sig[0])
		errs = []
		outs = []
		callargs = []
		for (pname,ptype,_) in sig[1]:
			if pname is not None:
				callargs.append("c_" + pname)
		callLine = 'C.' + fName + "(" + ",".join(callargs)+")"
		if ( cReturnType != 'void') :
			callLine = "\tc_ret := " + callLine
			c2goConverter = None
			try:
				c2goConverter = TYPEINFO[cReturnType]['c2goConverter']
			except:
				c2goConverter = lambda t,p:t
				errs.append('Warn: type (%s) has no Return converter to Go' % (cReturnType))
			
			goReturnType = None
			try:
				goReturnType = TYPEINFO[cReturnType]['goReturnType']
			except:
				try:
					goReturnType = re.sub(r'^\*','',TYPEINFO[cReturnType]['goArgType'])
				except:
					#goArgType = ptype
					goReturnType = cReturnType
					errs.append('Warn[createCallLine]: ReturnType %s Not defined' % (goReturnType))
			
			callLine+="\n\t" + c2goConverter(cReturnType,goReturnType) 
		else:
			callLine = "\t" + callLine
		return callLine,errs
	
	def processFuncsList(self,functionSignatures):
		flist = []
		for fName in functionSignatures:
			inImports = fName in IMPORTS
			if inImports or CONF_PRINT_SKIPPED: 
				errs = []
				sig = functionSignatures[fName]
				funcComment = """/* 
	   Function: %s
	   ReturnType: %s
	   Args: %s
*/
"""	% ( fName , "".join(sig[0]) , sig[1]) 
				funcCode = funcComment
				funArgString,err = self.createFuncArgString(fName,sig)
				errs+=err
				inputInits,err = self.createInputInits(fName,sig)
				errs+=err
				callLine,err = self.createCallLine(fName,sig)
				errs+=err
				if not inImports or errs:
					funcCode+="/*\n"
				if errs:
					funcCode+= "\n\t" + "\n\t".join(errs) + "\n\n"
				
				funcCode+=funArgString +" {\n"
				
				funcCode+=inputInits + "\n"
				
				funcCode+=callLine + "\n";
				
				funcCode+="}\n"
				
				if not inImports or errs:
					funcCode+="\n*/\n"
			
				flist.append(funcCode)
		return flist
	def processConstsList(self,enumSignatures):
		r = []
		for enum in enumSignatures:
			for name in enumSignatures[enum]:
				r.append(name)
				#r.append("const %s = C.%s\n" % (name,name))
		
		return r
	def prepareStruct(self,sig,cStruct,innerStruct):
		fields = [('handler',"C."+innerStruct)]
		errs = []
		getters = []
		goStructName = TYPEINFO[innerStruct]['goArgType']
		goStructName = re.sub(r'^\*','',goStructName)
		
		for el in sig:
			errored = False
			elName = el[0]
			fieldName=elName[0].upper()+elName[1:] 
			elType = "".join(el[1])
			
			info= lookInDb('s',innerStruct,elType,elName)
			
			if info:
				if 'PRIVATE' in info:
					fields.append(('// ' + elName , elType + " // Private" ) )
					continue
			if elType not in TYPEINFO:
				if elType in TYPEALIAS:
					elType =  TYPEALIAS[elType]
			if elType not in TYPEINFO:
				errored = True
				errs.append("Element %s has not registered type %s " % (elName,elType))
			goType = elType
			try:
				goType = TYPEINFO[elType]['goArgType']
			except:
				pass
			inner = None
			
			if goType == 'int':
				inner = "func (this *%(goStructName)s) Get%(fieldName)s() %(goType)s {\n" % {'goType':goType,'fieldName':fieldName,'goStructName':goStructName}
				inner += "\treturn int(this.handler."+elName+")\n}\n"
			elif goType[0] == '*':
				self.unsafe=True
				# this._next.handler = (C.xmlNsPtr)(unsafe.Pointer(this.handler.next))
				constr = re.sub('^\*','&',goType)
				fields.append(("_"+elName,goType))
				inner = "func (this *%(goStructName)s) Get%(fieldName)s() %(goType)s {\n" % {'goType':goType,'fieldName':fieldName,'goStructName':goStructName}
				inner += "\tif this.handler.%s == nil {\n\t\treturn nil\n\t}\n" % elName
				inner += "\tif this._%s == nil {\n\t\tthis._%s = %s{}\n\t}\n" % (elName,elName,constr)
				inner += "\tthis._%s.handler = (C.%s)(unsafe.Pointer(this.handler.%s))\n" % (elName,elType,elName) 
				inner += "\treturn this._"+elName+"\n}\n"
			elif goType== 'string':
				inner = "func (this *%(goStructName)s) Get%(fieldName)s() %(goType)s {\n" % {'goType':goType,'fieldName':fieldName,'goStructName':goStructName}
				inner += "\treturn C.GoString((*C.char)(unsafe.Pointer(this.handler."+elName+")))\n}\n"
				self.unsafe=True
			else:
				errored = True
				errs.append("Element %s not recognized getter for type %s " % (elName,elType))
				inner = "func (this *%(goStructName)s) Get%(fieldName)s() %(goType)s {\n" % {'goType':goType,'fieldName':fieldName,'goStructName':goStructName}
				inner += "\treturn int(this.handler."+elName+")\n}\n"
			
			if inner:
				if errored:
					inner = "/*\n" + inner + "*/\n"
				getters.append(inner)
			
		
		struct_desc =""
		if (errs ):
			struct_desc+="/*"
			struct_desc+= "\n\t" + "\n\t".join(errs) + "\n\n"
			struct_desc+="*/\n"
		
		struct_desc+= "type " + goStructName +" struct {\n\t" 
		struct_desc+= "\n\t".join([" ".join(f) for f in fields]) +"\n}\n"
		
		struct_desc+= "".join(getters)
		#if (errs ):
		#	struct_desc+="*/\n"
		
		return struct_desc 
	def processStructs(self,structs):
		ss = []
		for s in structs:
			for t in TYPEINFO:
				if  TYPEINFO[t].get('exportStruct','')==s and structs[s]['members']:
					ss.append(self.prepareStruct(structs[s]['members'],s,t))
					break
		return ss
	def processFile(self):
		if True:
			f=open(self.include,"r")
			s = f.read()
			s = re.sub("\s+(\d+)<<(\d+)(,?)",lambda x:str(int(x.group(1))<<int(x.group(2)))+ x.group(3),s)  
			#s = re.sub("\s+(\d+)<<(\d+),",lambda x: x.group(1) + "    <->     " + x.group(2),s)
			f.close()
			f = open(TMP + "/" + self.filename + ".h","w")
			f.write(s)
			f.close()
			(mode, ino, dev, nlink, uid, gid, size, atime, mtime, ctime)  = os.stat(self.include)
			os.utime(TMP + "/" + self.filename + ".h", (atime,mtime))


			
		p = CParser(TMP + "/" + self.filename + ".h",cache=TMP + "/" + self.filename + ".cache",macros={'XMLCALL': ''})
		#mk = {'XMLCALL': '',
		#'LIBXML_PUSH_ENABLED':1
		#}
		
		#p = CParser(["/tmp/tmp1/parser.h.h"],cache=TMP + "/" + self.filename + ".cache",macros=mk)
		
		
		varsdict = {
			'filename' : self.filename,
			'unsafe_import' : '',
			'consts_list' : '',
			'structs_list' : '',
			'funcs_list' : '',
		}
		
		#varsdict['consts_list'] = "".join(self.processConstsList(p.defs['enums']))
		
		self.consts = self.processConstsList(p.defs['enums'])
		
		varsdict['funcs_list'] = "".join(self.processFuncsList(p.defs['functions']))
		varsdict['structs_list'] = "".join(self.processStructs(p.defs['structs']))
		if self.unsafe:
			varsdict['unsafe_import'] = "import \"unsafe\""
		open(self.gofilename,"w").write(GO_TPL % varsdict)
		
		
		
		#sys.exit(0)

parser = optparse.OptionParser()

parser.add_option('-f', '--functions',
    action="append", dest="IMPORTS",
    help="Force change import functions", default=None)

parser.add_option('-i', '--includes',
    action="append", dest="INCLUDES",
    help="Force change import functions", default=None)


options, args = parser.parse_args()

if options.IMPORTS is not None:
	IMPORTS = options.IMPORTS
if options.INCLUDES is not None:
	INCLUDES = options.INCLUDES


if not os.path.exists(TMP):
	os.mkdir(TMP) 

consts = []
includes = []
for include in INCLUDES:
	p = FileConverter(include)
	p.processFile()
	includes.append(p.filename) 
	consts += p.consts

consts = dict(map(lambda x : (x ,1 ),consts)).keys()
incls = "\n".join(["#include <libxml/%s>" % fn for fn in includes])
varsdict = {
	'filename' : incls,
	'consts_list' : "\n".join(["const %s = C.%s" % (c,c) for c in consts],)
}
open("const.go","w").write("""package goxml
/*
#cgo pkg-config: libxml-2.0
%(filename)s
*/
import "C"

%(consts_list)s
""" % varsdict)

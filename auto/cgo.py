from pyclibrary import CParser
import os.path
import os
import sys
import re
import optparse

def calc_len(param,mul=1):
	def w(n,t):
		if t[-1]=='*':
			s = "c0_%(n)s:=C.int(len(%(p)s)*%(mul)s)\nc_%(n)s:=&c0_%(n)s" % {"n":n,"p":param,"t":t,'mul':mul}
		else:
			s = "c_%(n)s:=C.int(len(%(p)s)*%(mul)s)" % {"n":n,"p":param,"t":t,'mul':mul}
		return s
	return w
def create_buffer_as(param,mul=1):
	def w(n,t):
		#if t[-1]=='*':
		s = "\tc_%(n)s:= (%(ct)s)(C.calloc(  (C.size_t)( len(%(p)s)*%(mul)s+ 1 )  ,1))\n\tdefer C.free(unsafe.Pointer(c_%(n)s))" % {"n":n,"p":param,"t":t,'mul':mul,"ct":c2goc(t)}
		return s
	return w

# p1 - element for return as value 
# p2 - element for error calculation
# ec - error condition
def return_mapper(p1,p2,ec="%s != 0"):
	class t:
		def arg(self):
			return p1
		def errorCondition(self,vName):
			return ec % vName
		def errArg(self):
			return p2
		def getReturnTuple(self,argCType,cReturnType):
			return (("g_"+p1,TYPEINFO[argCType]['goArgType']),("err","error"))
		def getReturnType(self,argCType,cReturnType):
			return "(g_"+p1+" " + TYPEINFO[argCType]['goArgType'] + ",err error)"
			
	def w(n,t):
		#if t[-1]=='*':
		s = "\tc_%(n)s:= (%(ct)s)(C.calloc(  (C.size_t)( len(%(p)s)*%(mul)s+ 1 )  ,1))\n\tdefer C.free(unsafe.Pointer(c_%(n)s))" % {"n":n,"p":param,"t":t,'mul':mul,"ct":c2goc(t)}
		return s
	
	return t()
		
FUNC_DESCS = (
	
	
	('f','UTF8ToHtml',None,'inlen'),('CALC',calc_len('in')),
	('f','UTF8ToHtml',None,'outlen'),('CALC',calc_len('in',3)),
	('f','UTF8ToHtml',None,'out'),('CALC',create_buffer_as('in',3)),
	('r','UTF8ToHtml',None,None),('CALC',return_mapper('out','ret')),
	
	('f','htmlEncodeEntities',None,'inlen'),('CALC',calc_len('in')),
	('f','htmlEncodeEntities',None,'outlen'),('CALC',calc_len('in')),
	('f','htmlEncodeEntities',None,'out'),('CALC',create_buffer_as('in',3)),
	('r','htmlEncodeEntities',None,None),('CALC',return_mapper('out','ret')),
	
	('f','htmlCreateMemoryParserCtxt',None,'size'),('CALC',calc_len('buffer')),
	('r','htmlCreateMemoryParserCtxt',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	
	('f','htmlCreatePushParserCtxt','void*','user_data'),('SKIP',),
	('f','htmlCreatePushParserCtxt',None,'size'),('CALC',calc_len('chunk')),
	('r','htmlCreatePushParserCtxt',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	
	('r','htmlCtxtReadDoc',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	
	('r','htmlReadFile',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	('r','htmlCtxtReadFile',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	('r','htmlCtxtReadFd',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	('r','htmlCtxtReadIO',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	
	('f','htmlCtxtReadMemory',None,'size'),('CALC',calc_len('buffer')),
	('r','htmlCtxtReadMemory',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	
	('f','htmlSAXParseFile','void*','userData'),('SKIP',),
	
	('f','xmlReadMemory',None,'size'),('CALC',calc_len('buffer')),
	('r','xmlReadMemory',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	
	('f','xmlCreatePushParserCtxt','void*','user_data'),('SKIP',),
	('f','xmlCreatePushParserCtxt',None,'size'),('CALC',calc_len('chunk')),
	('r','xmlCreatePushParserCtxt',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	
	('r','xmlReadFile',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	('r','xmlCtxtReadFile',None,None),('CALC',return_mapper('ret','ret','%s == nil')),

	('r','xmlReaderForFile',None,None),('CALC',return_mapper('ret','ret','%s == nil')),

	
	#('f','','char*','filename'),
	('s','xmlDocPtr',None,'_private'),('PRIVATE'),
	('s','xmlDocPtr',None,'ids'),('PRIVATE'),
	('s','xmlDocPtr',None,'refs'),('PRIVATE'),
	('s','xmlDocPtr',None,'psvi'),('PRIVATE'),
	('s','xmlDtdPtr','void*',None),('PRIVATE'),
	('s','htmlElemDescPtr','char',None),('RETYPE','__bool'),
	('s',None,'void*','userData'),('PRIVATE'),
	('s',None,None,'type'),('PRIVATE'),
	('s',None,None,'_private'),('PRIVATE'),
	('s',None,'void*',None),('PRIVATE'),
	('s',None,'int*',None),('PRIVATE'),
	
)
getHandler =  lambda n,t:"c_%s := %s.handler" % (n,n)
getConverter = lambda n,t:'c_%s := %s(%s)' % (n,c2goc(t),n)

# The standard C numeric types are available under the names C.char, C.schar (signed char), 
# C.uchar (unsigned char), C.short, C.ushort (unsigned short), C.int, C.uint (unsigned int), C.long, C.ulong (unsigned long), C.longlong (long long), C.ulonglong (unsigned long long), C.float, C.double. 
# The C type void* is represented by Go's unsafe.Pointer.
def c2goc(t):
	return {
		'unsigned char*' : "*C.uchar",
		'unsigned int'	: "C.uint",
		'int'	: "C.int",
		'char*' : "*C.char",
		'xmlChar*' : "*C.xmlChar"
	}.get(t,"C."+t)
	
def toCharConverter(n,t):
	return """c_%(n)s:= (%(t)s)(unsafe.Pointer(C.CString(%(n)s)))
defer C.free(unsafe.Pointer(c_%(n)s))""" % {'n':n,'t':c2goc(t)}

def getNullOrHandler(n,t): 
	return """var c_%(n)s C.%(t)s=nil
if %(n)s !=nil { c_%(n)s = (C.%(t)s)(%(n)s.handler) }""" % {'n':n,'t':t}

#
# returnConverter
#
def retNullOrObject(t,goRetType):
	return 'if c_ret == nil {return nil}\nreturn &%s{handler:(C.%s)(c_ret)}' % (goRetType,t)  
def retObject(t,goRetType):
	return "return %s(c_ret)" % goRetType
def retString(t,goRetType):
	return "if c_ret == nil {return \"\"}\ng_ret:=C.GoString((*C.char)(unsafe.Pointer(c_ret)))\nreturn g_ret" 

# Converters

def c_cchar2string(vName,cName):
	return """if %(i)s == nil {
	%(a)s=""
} else {
	%(a)s = C.GoString((*C.char)(unsafe.Pointer(%(i)s)))
}""" % {"a":vName,"i":cName}

def c2GoConverter1(g_name,g_type,c_name,c_type):
	return "%s =  &%s{handler:(C.%s)(%s)}" % (g_name,g_type.replace('*',''),c_type,c_name)

TYPEALIAS = {
	'struct _xmlNode*' : 'xmlNodePtr',
	'struct _xmlDoc*' : 'xmlDocPtr',
	'struct _xmlDtd*' : 'xmlDtdPtr',
	'struct _xmlDict*' : 'xmlDictPtr',
	'struct _xmlNs*' : 'xmlNsPtr',
	'struct _xmlSAXHandler*' : 'xmlSAXHandlerPtr',
	'struct _xmlTextReader*' : 'xmlTextReaderPtr',
	'htmlElemDesc*' : 'htmlElemDescPtr',
	'htmlEntityDesc*' : 'htmlEntityDescPtr',
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
		'returnConverter' : retNullOrObject,
	},
	'xmlNsPtr' : {
		'goArgType' : '*XmlNs',
		'exportStruct' : '_xmlNs',
	},
	
	'htmlElemDescPtr' : {
		'goArgType' : '*HtmlElemDesc',
		'go2cConverter' : getNullOrHandler,
		'exportStruct' : '_htmlElemDesc',
		'returnConverter' : retNullOrObject,
	},
	'xmlElementType': {
		'goArgType':'int',
	},
	'xmlChar*' : {
		'goArgType':'string',
		'go2cConverter':toCharConverter,
		'returnConverter' : retString,
	},
	'int*' : {
		'goArgType':'*int',
		'go2cConverter' : lambda n,t:"\tc0_%(n)s:=C.int(*%(n)s)\n\tc_%(n)s:=&c0_%(n)s" % {"n":n,"t":t},
		'postProcessor' : lambda n: "*%(n)s = int(c0_%(n)s)" % {"n":n}
	},
	'unsigned char*' : ('alias','char*'),
	'char*' : {
		'goArgType':'string',
		'go2cConverter':toCharConverter,
		"to-string" : c_cchar2string,
	},
	'__string_ucharptr' : {
		'goArgType':'*string',
		'go2cConverter':lambda n,t: "\tc_%(n)s:= (*C.uchar)(unsafe.Pointer((C.CString(*%(n)s))))" % {"n":n,"t":t},
		'postProcessor' : lambda n: "*%(n)s = C.GoString((*C.char)(unsafe.Pointer(c_%(n)s)))" % {"n":n}
	},
	'xmlParserInputPtr' : {
		'goArgType' : '*XmlParserInput',
		'go2cConverter' : getNullOrHandler,
		'returnConverter' : retNullOrObject,
		'exportStruct' : '_xmlParserInput'
	},
	'xmlParserCtxtPtr': {
		'goArgType' : '*XmlParserCtxt',
		'go2cConverter' : getNullOrHandler,
		'returnConverter' : retNullOrObject,
		'c2GoConverter'	: c2GoConverter1, 
		'exportStruct' : '_xmlParserCtxt'
	},
	'xmlTextReaderPtr': {
		'goArgType' : '*XmlTextReader',
		'go2cConverter' : getNullOrHandler,
		'returnConverter' : retNullOrObject,
		'exportStruct' : '_xmlTextReader',
		'c2GoConverter'	: c2GoConverter1, 
	},
	
	'xmlNodePtr' : {
		'goArgType' : '*XmlNode',
		'go2cConverter' : getNullOrHandler,
		'returnConverter' : retNullOrObject,
		'exportStruct' : '_xmlNode',
	},
	
	'htmlDocPtr': ('alias','xmlDocPtr'),
	'htmlNodePtr':  ('alias','xmlNodePtr'),
	'htmlParserCtxtPtr':  ('alias','xmlParserCtxtPtr'),
	'htmlSAXHandlerPtr':  ('alias','xmlSAXHandlerPtr'),
	
	'xmlDocPtr' : {
		'goArgType' : '*XmlDoc',
		'go2cConverter' : getNullOrHandler,
		'returnConverter' : retNullOrObject,
		'exportStruct' : '_xmlDoc',
		'c2GoConverter'	: c2GoConverter1,
	},
	'htmlEntityDescPtr' : {
		'goArgType' : '*HtmlEntityDesc',
		'go2cConverter' : getNullOrHandler,
		'returnConverter' : retNullOrObject,
		'exportStruct' : '_htmlEntityDesc'
	},
	'xmlSAXHandlerPtr' : {
		'goArgType' : '*XmlSAXHandler',
		'go2cConverter' : getNullOrHandler,
		'returnConverter' : retNullOrObject,
		'exportStruct' : '_xmlSAXHandler'
	},
	'htmlStatus' : {
		'goArgType' : 'int',
		'go2cConverter' : getConverter,
		'returnConverter' : retObject,
	},
	'xmlCharEncoding' : ('alias','int'),
	'unsigned int' : {
		'goArgType' : 'uint',
		'go2cConverter' : getConverter,
		'returnConverter' : retObject,
	},
	'int' : {
		'goArgType' : 'int',
		'go2cConverter' : getConverter,
		'returnConverter' : retObject,
	},
	'void' : {
		'goReturnType':''
	},
	'__bool' : {
		'goArgType' : 'byte'
	}
}

INCLUDES = (
	"/usr/include/libxml2/libxml/tree.h",
	"/usr/include/libxml2/libxml/parser.h",
	"/usr/include/libxml2/libxml/xmlreader.h",
	"/usr/include/libxml2/libxml/xmlmemory.h",
	"/usr/include/libxml2/libxml/xmlstring.h",
	"/usr/include/libxml2/libxml/HTMLparser.h",
	"/usr/include/libxml2/libxml/encoding.h",
);

HTMLparser_IMPORTS = (
	'UTF8ToHtml',
	'htmlAttrAllowed',
	'htmlAutoCloseTag',
	'htmlCreateMemoryParserCtxt',
	'htmlCreatePushParserCtxt',
	'htmlCtxtReadDoc',
	'htmlReadFile',
	'htmlCtxtReadFd',
	'htmlCtxtReadFile',
	'#htmlCtxtReadIO',
	'htmlCtxtReadMemory',
	'#htmlReadIO',
	'htmlCtxtReset',
	'htmlCtxtUseOptions',
	'htmlElementAllowedHere',
	'htmlElementStatusHere',
	'htmlEncodeEntities',
	'htmlEntityLookup',
	'htmlEntityValueLookup',
	'htmlFreeParserCtxt',
	'htmlHandleOmittedElem',
	'htmlIsAutoClosed',
	'htmlIsScriptAttribute',
	'htmlNewParserCtxt',
	'htmlNodeStatus',
	'htmlParseCharRef',
	'htmlParseChunk',
	'htmlParseDoc',
	'htmlParseDocument',
	'htmlParseElement',
	'#htmlParseEntityRef',
	'htmlParseFile',
	'htmlReadDoc',
	'htmlReadFd',
	'htmlReadFile',
	'#htmlReadIO',
	'htmlReadMemory',
	'htmlSAXParseDoc',
	'htmlSAXParseFile',
	'htmlTagLookup',
	
)

parser_IMPORTS = (
	'xmlReadFile',
	'xmlFreeDoc',
	'xmlCleanupParser',
	'xmlFreeParserCtxt',
	'xmlNewParserCtxt',
	'xmlCtxtReadFile',
	'xmlReadMemory',
	'xmlCreatePushParserCtxt',
	'xmlParseChunk',
)

memory_IMPORTS = (
	'xmlMemoryDump',
)

reader_IMPORTS = (
	'xmlReaderForFile',
	'xmlTextReaderRead',
	'xmlFreeTextReader',
	'xmlTextReaderConstName',
	'xmlTextReaderConstValue',
	'xmlTextReaderDepth',
	'xmlTextReaderNodeType',
	'xmlTextReaderRead',
	'xmlTextReaderIsEmptyElement',
	'xmlTextReaderHasValue',
)

tree_IMPORTS = (
	'xmlDocGetRootElement',
)
IMPORTS = HTMLparser_IMPORTS + parser_IMPORTS + memory_IMPORTS + reader_IMPORTS + tree_IMPORTS



ALLI = (
	'xmlDocCopyNode',
	'xmlMemBlocks',
	'xmlStrlen',
	'xmlFreeTextReader',
	
	'xmlReaderForFile','xmlParseChunk','xmlReadFile','xmlReadMemory','xmlFreeDoc','xmlAddChild','xmlCleanupParser','xmlMemoryDump','xmlNewParserCtxt','xmlFreeParserCtxt','xmlCtxtReadFile','xmlCreatePushParserCtxt');

TMP = "/tmp/tmp1"

GO_TPL = """package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/%(filename)s>
*/
import "C"
%(imports)s

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
		self.consts = []
	def calcArgType(self,arg,sig):
		if arg == 'ret':
			return "".join(sig[0])
		for s in sig[1]:
			if arg==s[0]:
				return "".join(s[1])
		raise Exception("calcArgType not found Name:\"%s\" Signature:%s" %(arg,sig))
	def createFuncArgString(self,fname,sig):
		goFname = fname[0].upper() + fname[1:]
		cReturnType = "".join(sig[0]) 
		args = []
		errs = []
		goReturnType = None
		
		dbData = lookInDb('r',fname,cReturnType,None)
		if dbData:
			if (dbData[0] == 'CALC'):
				goReturnType=dbData[1].getReturnType(self.calcArgType(dbData[1].arg(),sig),cReturnType)
		if cReturnType in TYPEALIAS:
			cReturnType = TYPEALIAS[cReturnType]
		if goReturnType is None:
			try:
				goReturnType = TYPEINFO[cReturnType]['goReturnType']
			except:
				try:
					goReturnType = TYPEINFO[cReturnType]['goArgType']
				except:
					#goArgType = ptype
					goReturnType = cReturnType
					errs.append('Warn[createFuncArgString]: ReturnType %s Not defined' % (goReturnType))
		
		for (pname,ptype,_) in sig[1]:
			if pname is not None:
				ptype = "".join(ptype)
				dbData = lookInDb('f',fname,ptype,pname)
				goArgType = None
				if dbData:
					if dbData[0] in ('SKIP','CALC'):
						continue
					elif dbData[0] == 'RETYPE':
						ptype = dbData[1]
					else:
						raise (Exception('Not implemented %s' % str(dbData)))
				
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
		outVarsBlock = []
		recalcsBlock = []
		for (pname,ptype,_) in sig[1]:
			if pname is not None:
				ptype = "".join(ptype)
				dbData = lookInDb('f',fname,ptype,pname)
				go2cConvert = None
				if dbData:
					if 'SKIP' in dbData:
						continue
					elif dbData[0] == 'RETYPE':
						ptype = dbData[1]
					elif dbData[0] == 'CALC':
						recalcsBlock.append(dbData[1](pname,ptype))
						continue
					else:
						raise (Exception('Not implemented %s' % str(dbData)))
				try:
					go2cConvert = TYPEINFO[ptype]['go2cConverter']
				except:
					go2cConvert = lambda a,b:pname
					errs.append('Warn: %s %s No converter to C(go2cConverter)' % (pname,ptype))
				
				outVarsBlock.append(str(go2cConvert(pname,ptype)) )
		
		outStr = self.funcJoin(outVarsBlock)
		recalcStr = self.funcJoin(recalcsBlock)
		
		return "\n".join((outStr , recalcStr)),errs
	
	def createCallLine(self,fName,sig):
		
		cReturnType = "".join(sig[0])
		errs = []
		outs = []
		callargs = []
		for (pname,ptype,_) in sig[1]:
			ptype = "".join(ptype)
			dbData = lookInDb('f',fName,ptype,pname)
			if dbData:
				if 'SKIP' in dbData:
					callargs.append("nil")
					continue
				elif dbData[0] == 'RETYPE':
					ptype = dbData[1]
				elif dbData[0] == 'CALC':
					pass
				else:
					raise (Exception('Not implemented %s' % str(dbData)))
			if pname is not None:
				callargs.append("c_" + pname)
			
		callLine = 'C.' + fName + "(" + ",".join(callargs)+")"
		
		if ( cReturnType != 'void') :
			callLine = "\tc_ret := " + callLine + "\n"
		else:
			callLine = "\t" + callLine + "\n"
		
		returnBlock = []
		errorProcessBlock = []
		dbData = lookInDb('r',fName,cReturnType,None)
		if dbData and dbData[0] == 'CALC':
			mapper = dbData[1]
			errArg = mapper.errArg()
			if errArg == 'ret':
				errorCondition = mapper.errorCondition('c_ret')
				errorProcessBlock.append("""if %(errorCondition)s {
	err = fmt.Errorf("%(fName)s errno %%d" ,c_ret)
}"""% {"errorCondition":errorCondition,"fName":fName})
				returnBlock.append("return")
			else:
				raise Exception("Not Implemented")
			#goReturnType=dbData[1].getReturnType(self.calcArgType(dbData[1].arg(),sig),cReturnType)
		elif cReturnType != 'void' :
			if cReturnType in TYPEALIAS:
				cReturnType = TYPEALIAS[cReturnType]
			if  'realType' in TYPEINFO[cReturnType]:
				cReturnType = TYPEINFO[cReturnType]['realType']
			returnConverter = None
			try:
				returnConverter = TYPEINFO[cReturnType]['returnConverter']
			except:
				returnConverter = lambda t,p:t
				errs.append('Warn[createCallLine]: type (%s) has no Return converter to Go' % (cReturnType))
			 
			goReturnType = None
			try:
				goReturnType = TYPEINFO[cReturnType]['goReturnType']
			except:
				try:
					goReturnType = re.sub(r'^\*','',TYPEINFO[cReturnType]['goArgType'])
				except:
					#goArgType = ptype
					goReturnType = cReturnType
					errs.append('Warn[createCallLine]: ReturnType %s found but not defined' % (goReturnType))
			
			returnBlock.append(returnConverter(cReturnType,goReturnType)) 
		# Post Process Parameters
		postProcessBlock = []
		for (pname,ptype,_) in sig[1]:
			if pname is not None:
				ptype = "".join(ptype)
				dbData = lookInDb('f',fName,ptype,pname)
				if dbData:
					if dbData[0] == 'RETYPE':
						ptype = dbData[1]
					elif dbData[0] == 'CALC':
						continue
				postProcessor = TYPEINFO.get(ptype,{}).get('postProcessor')
				if postProcessor:
					postProcessBlock.append("\t" + postProcessor(pname))
		
		dbData = lookInDb('r',fName,cReturnType,None)
		if dbData and dbData[0] == 'CALC':
			mapper = dbData[1]
			arg1 = mapper.arg()
			if arg1!="ret":
				#print arg1, "\t",self.calcArgType(arg1,sig),"\t",mapper.getReturnTuple(self.calcArgType(arg1,sig),cReturnType)
				v = mapper.getReturnTuple(self.calcArgType(arg1,sig),cReturnType)
				vName = v[0][0]
				vType = v[0][1]
				cName = arg1
				cType = self.calcArgType(arg1,sig)
				cInfo = TYPEINFO[cType]
				if 'to-' + vType in cInfo:
					converter = cInfo['to-' + vType]
					postProcessBlock.append(converter(vName,"c_"+cName))
				else:
					raise Exception("Not implemented converstion to %s to from Type %s(%s)" % (vType,cType,cInfo))
			else:
				cType = self.calcArgType(arg1,sig)
				cInfo = TYPEINFO[cType]
				# avoid aliases problems
				if 'realType' in cInfo:
					cType = cInfo['realType']
				
				postProcessBlock.append(cInfo['c2GoConverter']('g_ret',cInfo['goArgType'],'c_ret',cType))
			
		returnStr = self.funcJoin(returnBlock)
		postProcessStr = self.funcJoin(postProcessBlock)
		errorProcessStr= self.funcJoin(errorProcessBlock)
		if errorProcessStr and postProcessStr:
			errorProcessStr += " else {"
			postProcessStr = self.funcJoin((postProcessStr,"}"))
		return "\n".join((callLine,errorProcessStr,postProcessStr,returnStr)),errs
	def funcJoin(self,l):
		r = []
		for i in l:
			r.append("\n".join(["\t" + k for k in i.split("\n")]))
		return "\n".join(r)
	def log(self,s):
		if VERBOSE:
			print self.filename+": " +s
	
	def processFuncsList(self,functionSignatures):
		flist = []
		#print functionSignatures
		for fName in functionSignatures:
			inImports = fName in IMPORTS
			if not inImports:
				self.log("Skip %s"%fName)
			if inImports or CONF_PRINT_SKIPPED: 
				errs = []
				sig = functionSignatures[fName]
				funcComment = """/* 
	   Function: %s
	   ReturnType: %s
	   Args: %s
*/
"""	% ( fName , "".join(sig[0]) , sig[1])
				sig = list(sig)
				sig[1] = list(sig[1])
				funcCode = funcComment
				self.preprocessArgsSig(sig)
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
				if 'unsafe.Pointer' in funcCode:
					self.unsafe = True
				flist.append(funcCode)
				
				if errs:
					self.log("Errored %s"%fName)
				else:
					self.log("Process %s"%fName)
			
		return flist
	def preprocessArgsSig(self,sig):
		c = 1
		for i,(n,t,_) in enumerate(sig[1]):
			if n is None and "".join(t) != 'void':
				sig[1][i]=("arg%i" % c , t, _)
				c+=1
		for i,(n,t,_) in enumerate(sig[1]):
			elType = "".join(t)
			if elType in TYPEALIAS:
				sig[1][i]=(n , TYPEALIAS[elType], _)
				
		
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
			retyped = False
			elName = el[0]
			fieldName=elName[0].upper()+elName[1:] 
			elType = "".join(el[1])
			
			info= lookInDb('s',innerStruct,elType,elName)
			
			if info:
				if 'PRIVATE' in info:
					fields.append(('// ' + elName , elType + " // Private" ) )
					continue
				elif info[0] == 'RETYPE':
					elType = info[1] 
					retyped = True
			if not retyped and elType not in TYPEINFO:
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
			elif goType == 'byte':
				inner = "func (this *%(goStructName)s) Get%(fieldName)s() %(goType)s {\n" % {'goType':goType,'fieldName':fieldName,'goStructName':goStructName}
				inner += "\treturn byte(this.handler."+elName+")\n}\n"
			elif goType== '*string':
				inner = "func (this *%(goStructName)s) Get%(fieldName)s() %(goType)s {\n" % {'goType':goType,'fieldName':fieldName,'goStructName':goStructName}
				inner += "if this.handler."+elName+" == nil { return nil }\n"
				inner += "\ts:=C.GoString((*C.char)(unsafe.Pointer(this.handler."+elName+")))\n\treturn &s\n}\n"
				self.unsafe=True
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
				errs.append("Element %s not recognized getter for elType:%s goType:%s" % (elName,elType,goType))
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
		replace = {}
		if True:
			# enum a = 1 << 213 bug
			f=open(self.include,"r")
			s = f.read()
			s = re.findall("(\s+(\d+)<<(\d+)(,?))",s)
			for r in s:
				replace[r[0]] = str(int(r[1]) << int(r[2]) ) +  r[3] 
		
		# Read defines from version ( defines only - not funcs and etc )
		
		vp = CParser(
			"/".join(self.include.split("/")[:-1] + list(("xmlversion.h",))),
			cache=TMP + "/" + "xmlversion.h" + ".cache",
		)
		#print vp.defs['macros']
		#return
		macros = {
			'XMLCALL': '',
			#'LIBXML_PUSH_ENABLED':'',
			#'LIBXML_READER_ENABLED':''
		}
		macros.update(vp.defs['macros'])
		p = CParser(
			# TMP + "/" + self.filename + ".h",
			self.include,
			cache=TMP + "/" + self.filename + ".cache",
			#copyFrom=vp,
			replace  = replace,	
			#replace = {"\s+(\d+)<<(\d+)(,?)" : lambda x:str(int(x.group(1))<<int(x.group(2)))+ x.group(3)},
			macros=macros
		)
		if options.PRINTALL:
			p.printAll()
			return
		#mk = {'XMLCALL': '',
		#'LIBXML_PUSH_ENABLED':1
		#}
		
		#p = CParser(["/tmp/tmp1/parser.h.h"],cache=TMP + "/" + self.filename + ".cache",macros=mk)
		
				
		varsdict = {
			'filename' : self.filename,
			"imports" : '',
			'consts_list' : '',
			'structs_list' : '',
			'funcs_list' : '',
		}
		
		#varsdict['consts_list'] = "".join(self.processConstsList(p.defs['enums']))
		
		self.consts = self.processConstsList(p.defs['enums'])
		
		varsdict['funcs_list'] = "".join(self.processFuncsList(p.defs['functions']))
		varsdict['structs_list'] = "".join(self.processStructs(p.defs['structs']))
		imports = []
		
		if "errors.New" in varsdict['funcs_list'] or "errors.New" in varsdict['structs_list']  :
			imports.append("import \"errors\"\n")
		if "unsafe.Pointer" in varsdict['funcs_list'] or "unsafe.Pointer" in varsdict['structs_list']  :
			imports.append("import \"unsafe\"\n")
		if "fmt.Errorf" in varsdict['funcs_list'] or "fmt.Errorf" in varsdict['structs_list']  :
			imports.append("import \"fmt\"\n")
		
		varsdict['imports'] =  	"".join(imports)
		
		open(self.gofilename,"w").write(GO_TPL % varsdict)
		
		
		
		

parser = optparse.OptionParser()

parser.add_option('-f', '--functions',
    action="append", dest="IMPORTS",
    help="Force change import functions", default=None)

parser.add_option('-i', '--includes',
    action="append", dest="INCLUDES",
    help="Force change import functions", default=None)

parser.add_option('-p', '--print-all',
    action="store_true", dest="PRINTALL",
    help="Print parser content", default=None)

parser.add_option('-v', '--verbose',
    action="count", dest="VERBOSE",
    help="Verbose", default=0)


options, args = parser.parse_args()

if options.IMPORTS is not None:
	IMPORTS = options.IMPORTS
if options.INCLUDES is not None:
	INCLUDES = options.INCLUDES
if options.VERBOSE is not None:
	VERBOSE = options.VERBOSE


def convertAliases():
	for t in TYPEINFO:
		if isinstance(TYPEINFO[t],tuple):
			if TYPEINFO[t][0] == 'alias':
				parent = TYPEINFO[t][1]
				TYPEINFO[t] = {'realType' : parent}
				TYPEINFO[t].update(TYPEINFO[parent])
				if 'exportStruct' in TYPEINFO[t]:
					del TYPEINFO[t]['exportStruct']
	
if not os.path.exists(TMP):
	os.mkdir(TMP) 

convertAliases()
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

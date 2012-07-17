"""
	goxml description
"""
# FILES TO INCLUDE
INCLUDES = (
	"/usr/include/libxml2/libxml/tree.h",
	"/usr/include/libxml2/libxml/parser.h",
	"/usr/include/libxml2/libxml/xmlreader.h",
	"/usr/include/libxml2/libxml/xmlmemory.h",
	"/usr/include/libxml2/libxml/xmlstring.h",
	"/usr/include/libxml2/libxml/HTMLparser.h",
	"/usr/include/libxml2/libxml/encoding.h",
	"/usr/include/libxml2/libxml/xpath.h",
);


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
		
FUNC_DESCS = [
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
	('r','xmlParseFile',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	('r','xmlCtxtReadFile',None,None),('CALC',return_mapper('ret','ret','%s == nil')),

	('r','xmlReaderForFile',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	
	('f','xmlTextReaderPreservePattern',None,'namespaces'),('SKIP',),
	
	('r','xmlNewProp',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	('r','xmlNewChild',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	('r','xmlNewDoc',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	('r','xmlNewNode',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	('r','xmlDocSetRootElement',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	('r','xmlAddChild',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	('r','xmlNewText',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	('r','xmlCreateIntSubset',None,None),('CALC',return_mapper('ret','ret','%s == nil')),
	
	('s','xmlNodePtr',None,'type'),('RETYPE','__XmlNode_reserved_type'),
	('s','htmlElemDescPtr','char',None),('RETYPE','__bool'),
	
	#('f','','char*','filename'),
	('s','xmlDocPtr',None,'_private'),('PRIVATE'),
	('s','xmlDocPtr',None,'ids'),('PRIVATE'),
	('s','xmlDocPtr',None,'refs'),('PRIVATE'),
	('s','xmlDocPtr',None,'psvi'),('PRIVATE'),
	('s','xmlDtdPtr','void*',None),('PRIVATE'),
	
	('s',None,'void*','userData'),('PRIVATE'),
	
	('s',None,None,'type'),('PRIVATE'),
	('s',None,None,'_private'),('PRIVATE'),
	('s',None,'void*',None),('PRIVATE'),
	('s',None,'int*',None),('PRIVATE'),
	
]
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
	'xmlNs*' :'xmlNsPtr',
	'struct _xmlSAXHandler*' : 'xmlSAXHandlerPtr',
	'struct _xmlTextReader*' : 'xmlTextReaderPtr',
	'htmlElemDesc*' : 'htmlElemDescPtr',
	'htmlEntityDesc*' : 'htmlEntityDescPtr',
}
TYPEINFO = {                                                
	'xmlDtdPtr' : {
		'goArgType' : '*XmlDtd',
		'exportStruct' : '_xmlDtd',
		'c2GoConverter'	: c2GoConverter1, 
		'go2cConverter' : getNullOrHandler,
		'returnConverter' : retNullOrObject,
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
		'go2cConverter' : getNullOrHandler,
	},
	'xmlAttrPtr' : {
		'goArgType' : '*XmlAttr',
		'exportStruct' : '_xmlAttr',
		'c2GoConverter'	: c2GoConverter1, 
		'go2cConverter' : getNullOrHandler,
		'returnConverter' : retNullOrObject,
	},
	'xmlXPathContextPtr' : {
		'goArgType' : '*XmlXPathContext',
		'exportStruct' : '_xmlXPathContext',
		'c2GoConverter'	: c2GoConverter1, 
		'go2cConverter' : getNullOrHandler,
		'returnConverter' : retNullOrObject,
	},
	'xmlXPathObjectPtr' : {
		'goArgType' : '*XmlXPathObject',
		'exportStruct' : '_xmlXPathObject',
		'c2GoConverter'	: c2GoConverter1, 
		'go2cConverter' : getNullOrHandler,
		'returnConverter' : retNullOrObject,
	},
	'htmlElemDescPtr' : {
		'goArgType' : '*HtmlElemDesc',
		'go2cConverter' : getNullOrHandler,
		'exportStruct' : '_htmlElemDesc',
		'returnConverter' : retNullOrObject,
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
		'c2GoConverter'	: c2GoConverter1, 
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
	'xmlElementType' : {
		'goArgType' : 'int',
		'go2cConverter' : getConverter,
		'returnConverter' : retObject,
	},
	'__XmlNode_reserved_type' : {
		#'goArgType' : 'int',
		#'go2cConverter' : lambda t:"return int(C.FETCH_TYPE)"
		#'returnConverter' : lambda t:"return int(C.FETCH_TYPE)"
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
	'FILE*' : {
		'goArgType' : '*os.File',
		'go2cConverter' : lambda n,t:"""var c_%(n)s *C.FILE
{
	tp:= (*C.char)(unsafe.Pointer(C.CString("w")));
	defer C.free(unsafe.Pointer(tp));
	c_%(n)s = C.fdopen((C.int)(%(n)s.Fd()),tp)
}
""" % {'n':n}
	},
	'void' : {
		'goReturnType':''
	},
	'__bool' : {
		'goArgType' : 'byte'
	}
}

CPOSTPROCESSORS = {
	'tree.h' : """int XmlNode_fetch_type(xmlNodePtr h) {
	return (int)h->type;
}
"""
}

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
	'xmlParseFile',
	'xmlInitParser',
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
	'xmlTextReaderIsValid',
	'xmlTextReaderCurrentDoc',
	'xmlTextReaderPreservePattern',
	'xmlReaderNewFile',
)

tree_IMPORTS = (
	'xmlDocGetRootElement',
	'xmlDocDump',
	'xmlNewChild',
	'xmlNewProp',
	'xmlSaveFormatFileEnc',
	'xmlNewDoc',
	'xmlNewNode',
	'xmlDocSetRootElement',
	'xmlAddChild',
	'xmlNewText',
	'xmlCreateIntSubset'
)

xpath_IMPORTS = (
	'@xmlXPathNewContext',
	'@xmlXPathEvalExpression',
	'xmlXPathFreeContext',
	'xmlXPathFreeObject'
)

IMPORTS = list(HTMLparser_IMPORTS + parser_IMPORTS + memory_IMPORTS + reader_IMPORTS + tree_IMPORTS + xpath_IMPORTS) 

GO_TPL = """package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/%(filename)s>
%(cpostprocessors)s
*/
import "C"
%(imports)s

%(consts_list)s
%(structs_list)s
%(funcs_list)s

"""




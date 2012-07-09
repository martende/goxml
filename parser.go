package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/parser.h>
*/
import "C"

import "unsafe"
 
const XML_PARSER_ATTRIBUTE_VALUE=C.XML_PARSER_ATTRIBUTE_VALUE
const XML_PARSER_CDATA_SECTION=C.XML_PARSER_CDATA_SECTION
const XML_PARSER_COMMENT=C.XML_PARSER_COMMENT
const XML_PARSER_CONTENT=C.XML_PARSER_CONTENT
const XML_PARSER_DTD=C.XML_PARSER_DTD
const XML_PARSER_END_TAG=C.XML_PARSER_END_TAG
const XML_PARSER_ENTITY_DECL=C.XML_PARSER_ENTITY_DECL
const XML_PARSER_ENTITY_VALUE=C.XML_PARSER_ENTITY_VALUE
const XML_PARSER_EOF=C.XML_PARSER_EOF
const XML_PARSER_EPILOG=C.XML_PARSER_EPILOG
const XML_PARSER_IGNORE=C.XML_PARSER_IGNORE
const XML_PARSER_MISC=C.XML_PARSER_MISC
const XML_PARSER_PI=C.XML_PARSER_PI
const XML_PARSER_PROLOG=C.XML_PARSER_PROLOG
const XML_PARSER_PUBLIC_LITERAL=C.XML_PARSER_PUBLIC_LITERAL
const XML_PARSER_START=C.XML_PARSER_START
const XML_PARSER_START_TAG=C.XML_PARSER_START_TAG
const XML_PARSER_SYSTEM_LITERAL=C.XML_PARSER_SYSTEM_LITERAL
const XML_PARSE_COMPACT=C.XML_PARSE_COMPACT
const XML_PARSE_DOM=C.XML_PARSE_DOM
const XML_PARSE_DTDATTR=C.XML_PARSE_DTDATTR
const XML_PARSE_DTDLOAD=C.XML_PARSE_DTDLOAD
const XML_PARSE_DTDVALID=C.XML_PARSE_DTDVALID
const XML_PARSE_HUGE=C.XML_PARSE_HUGE
const XML_PARSE_NOBASEFIX=C.XML_PARSE_NOBASEFIX
const XML_PARSE_NOBLANKS=C.XML_PARSE_NOBLANKS
const XML_PARSE_NOCDATA=C.XML_PARSE_NOCDATA
const XML_PARSE_NODICT=C.XML_PARSE_NODICT
const XML_PARSE_NOENT=C.XML_PARSE_NOENT
const XML_PARSE_NOERROR=C.XML_PARSE_NOERROR
const XML_PARSE_NONET=C.XML_PARSE_NONET
const XML_PARSE_NOWARNING=C.XML_PARSE_NOWARNING
const XML_PARSE_NOXINCNODE=C.XML_PARSE_NOXINCNODE
const XML_PARSE_NSCLEAN=C.XML_PARSE_NSCLEAN
const XML_PARSE_OLD10=C.XML_PARSE_OLD10
const XML_PARSE_OLDSAX=C.XML_PARSE_OLDSAX
const XML_PARSE_PEDANTIC=C.XML_PARSE_PEDANTIC
const XML_PARSE_PUSH_DOM=C.XML_PARSE_PUSH_DOM
const XML_PARSE_PUSH_SAX=C.XML_PARSE_PUSH_SAX
const XML_PARSE_READER=C.XML_PARSE_READER
const XML_PARSE_RECOVER=C.XML_PARSE_RECOVER
const XML_PARSE_SAX=C.XML_PARSE_SAX
const XML_PARSE_SAX1=C.XML_PARSE_SAX1
const XML_PARSE_UNKNOWN=C.XML_PARSE_UNKNOWN
const XML_PARSE_XINCLUDE=C.XML_PARSE_XINCLUDE
const XML_WITH_AUTOMATA=C.XML_WITH_AUTOMATA
const XML_WITH_C14N=C.XML_WITH_C14N
const XML_WITH_CATALOG=C.XML_WITH_CATALOG
const XML_WITH_DEBUG=C.XML_WITH_DEBUG
const XML_WITH_DEBUG_MEM=C.XML_WITH_DEBUG_MEM
const XML_WITH_DEBUG_RUN=C.XML_WITH_DEBUG_RUN
const XML_WITH_EXPR=C.XML_WITH_EXPR
const XML_WITH_FTP=C.XML_WITH_FTP
const XML_WITH_HTML=C.XML_WITH_HTML
const XML_WITH_HTTP=C.XML_WITH_HTTP
const XML_WITH_ICONV=C.XML_WITH_ICONV
const XML_WITH_ICU=C.XML_WITH_ICU
const XML_WITH_ISO8859X=C.XML_WITH_ISO8859X
const XML_WITH_LEGACY=C.XML_WITH_LEGACY
const XML_WITH_MODULES=C.XML_WITH_MODULES
const XML_WITH_NONE=C.XML_WITH_NONE
const XML_WITH_OUTPUT=C.XML_WITH_OUTPUT
const XML_WITH_PATTERN=C.XML_WITH_PATTERN
const XML_WITH_PUSH=C.XML_WITH_PUSH
const XML_WITH_READER=C.XML_WITH_READER
const XML_WITH_REGEXP=C.XML_WITH_REGEXP
const XML_WITH_SAX1=C.XML_WITH_SAX1
const XML_WITH_SCHEMAS=C.XML_WITH_SCHEMAS
const XML_WITH_SCHEMATRON=C.XML_WITH_SCHEMATRON
const XML_WITH_THREAD=C.XML_WITH_THREAD
const XML_WITH_TREE=C.XML_WITH_TREE
const XML_WITH_UNICODE=C.XML_WITH_UNICODE
const XML_WITH_VALID=C.XML_WITH_VALID
const XML_WITH_WRITER=C.XML_WITH_WRITER
const XML_WITH_XINCLUDE=C.XML_WITH_XINCLUDE
const XML_WITH_XPATH=C.XML_WITH_XPATH
const XML_WITH_XPTR=C.XML_WITH_XPTR
const XML_WITH_ZLIB=C.XML_WITH_ZLIB
func XmlCleanupParser() {
	C.xmlCleanupParser()
}

func XmlCreatePushParserCtxt(sax *XmlSAXHandler,chunk string,size int,filename string) *XmlParserCtxt {
	var c_ret C.xmlParserCtxtPtr
	g_ret := &XmlParserCtxt{}
	var c_sax C.xmlSAXHandlerPtr=nil
	if sax!=nil { c_sax = sax.handler}
	c_user_data := unsafe.Pointer(nil)
	c_chunk := C.CString(chunk)
	c_size := C.int(size)
	c_filename := C.CString(filename)
	c_ret = C.xmlCreatePushParserCtxt(c_sax,c_user_data,c_chunk,c_size,c_filename)
	g_ret.handler = c_ret
	return g_ret
}

func XmlCtxtReadFile(ctxt *XmlParserCtxt,filename string,encoding string,options int) *XmlDoc {
	var c_ret C.xmlDocPtr
	g_ret := &XmlDoc{}
	c_ctxt := ctxt.handler
	c_filename := C.CString(filename)
	c_encoding := C.CString(encoding)
	c_options := C.int(options)
	c_ret = C.xmlCtxtReadFile(c_ctxt,c_filename,c_encoding,c_options)
	g_ret.handler = c_ret
	return g_ret
}

func XmlFreeParserCtxt(ctxt *XmlParserCtxt) {
	c_ctxt := ctxt.handler
	C.xmlFreeParserCtxt(c_ctxt)
}

func XmlNewParserCtxt() *XmlParserCtxt {
	var c_ret C.xmlParserCtxtPtr
	g_ret := &XmlParserCtxt{}
	c_ret = C.xmlNewParserCtxt()
	g_ret.handler = c_ret
	return g_ret
}

func XmlParseChunk(ctxt *XmlParserCtxt,chunk string,size int,terminate int) int {
	var c_ret C.int
	var g_ret int
	c_ctxt := ctxt.handler
	c_chunk := C.CString(chunk)
	c_size := C.int(size)
	c_terminate := C.int(terminate)
	c_ret = C.xmlParseChunk(c_ctxt,c_chunk,c_size,c_terminate)
	g_ret = int(c_ret)
	return g_ret
}

func XmlReadFile(URL string,encoding string,options int) *XmlDoc {
	var c_ret C.xmlDocPtr
	g_ret := &XmlDoc{}
	c_URL := C.CString(URL)
	c_encoding := C.CString(encoding)
	c_options := C.int(options)
	c_ret = C.xmlReadFile(c_URL,c_encoding,c_options)
	g_ret.handler = c_ret
	return g_ret
}

func XmlReadMemory(buffer string,size int,URL string,encoding string,options int) *XmlDoc {
	var c_ret C.xmlDocPtr
	g_ret := &XmlDoc{}
	c_buffer := C.CString(buffer)
	c_size := C.int(size)
	c_URL := C.CString(URL)
	c_encoding := C.CString(encoding)
	c_options := C.int(options)
	c_ret = C.xmlReadMemory(c_buffer,c_size,c_URL,c_encoding,c_options)
	g_ret.handler = c_ret
	return g_ret
}

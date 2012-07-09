package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/xmlreader.h>
*/
import "C"

import "unsafe"
 
const XML_PARSER_DEFAULTATTRS=C.XML_PARSER_DEFAULTATTRS
const XML_PARSER_LOADDTD=C.XML_PARSER_LOADDTD
const XML_PARSER_SEVERITY_ERROR=C.XML_PARSER_SEVERITY_ERROR
const XML_PARSER_SEVERITY_VALIDITY_ERROR=C.XML_PARSER_SEVERITY_VALIDITY_ERROR
const XML_PARSER_SEVERITY_VALIDITY_WARNING=C.XML_PARSER_SEVERITY_VALIDITY_WARNING
const XML_PARSER_SEVERITY_WARNING=C.XML_PARSER_SEVERITY_WARNING
const XML_PARSER_SUBST_ENTITIES=C.XML_PARSER_SUBST_ENTITIES
const XML_PARSER_VALIDATE=C.XML_PARSER_VALIDATE
const XML_READER_TYPE_ATTRIBUTE=C.XML_READER_TYPE_ATTRIBUTE
const XML_READER_TYPE_CDATA=C.XML_READER_TYPE_CDATA
const XML_READER_TYPE_COMMENT=C.XML_READER_TYPE_COMMENT
const XML_READER_TYPE_DOCUMENT=C.XML_READER_TYPE_DOCUMENT
const XML_READER_TYPE_DOCUMENT_FRAGMENT=C.XML_READER_TYPE_DOCUMENT_FRAGMENT
const XML_READER_TYPE_DOCUMENT_TYPE=C.XML_READER_TYPE_DOCUMENT_TYPE
const XML_READER_TYPE_ELEMENT=C.XML_READER_TYPE_ELEMENT
const XML_READER_TYPE_END_ELEMENT=C.XML_READER_TYPE_END_ELEMENT
const XML_READER_TYPE_END_ENTITY=C.XML_READER_TYPE_END_ENTITY
const XML_READER_TYPE_ENTITY=C.XML_READER_TYPE_ENTITY
const XML_READER_TYPE_ENTITY_REFERENCE=C.XML_READER_TYPE_ENTITY_REFERENCE
const XML_READER_TYPE_NONE=C.XML_READER_TYPE_NONE
const XML_READER_TYPE_NOTATION=C.XML_READER_TYPE_NOTATION
const XML_READER_TYPE_PROCESSING_INSTRUCTION=C.XML_READER_TYPE_PROCESSING_INSTRUCTION
const XML_READER_TYPE_SIGNIFICANT_WHITESPACE=C.XML_READER_TYPE_SIGNIFICANT_WHITESPACE
const XML_READER_TYPE_TEXT=C.XML_READER_TYPE_TEXT
const XML_READER_TYPE_WHITESPACE=C.XML_READER_TYPE_WHITESPACE
const XML_READER_TYPE_XML_DECLARATION=C.XML_READER_TYPE_XML_DECLARATION
const XML_TEXTREADER_MODE_CLOSED=C.XML_TEXTREADER_MODE_CLOSED
const XML_TEXTREADER_MODE_EOF=C.XML_TEXTREADER_MODE_EOF
const XML_TEXTREADER_MODE_ERROR=C.XML_TEXTREADER_MODE_ERROR
const XML_TEXTREADER_MODE_INITIAL=C.XML_TEXTREADER_MODE_INITIAL
const XML_TEXTREADER_MODE_INTERACTIVE=C.XML_TEXTREADER_MODE_INTERACTIVE
const XML_TEXTREADER_MODE_READING=C.XML_TEXTREADER_MODE_READING
func XmlFreeTextReader(reader *XmlTextReader) {
	c_reader := reader.handler
	C.xmlFreeTextReader(c_reader)
}

func XmlReaderForFile(filename string,encoding string,options int) *XmlTextReader {
	var c_ret C.xmlTextReaderPtr
	g_ret := &XmlTextReader{}
	c_filename := C.CString(filename)
	c_encoding := C.CString(encoding)
	c_options := C.int(options)
	c_ret = C.xmlReaderForFile(c_filename,c_encoding,c_options)
	g_ret.handler = c_ret
	return g_ret
}

func XmlTextReaderConstName(reader *XmlTextReader) string {
	var c_ret *C.char
	var g_ret string
	c_reader := reader.handler
	c_ret = (*C.char)(unsafe.Pointer(C.xmlTextReaderConstName(c_reader)))
	g_ret = C.GoString(c_ret)
	return g_ret
}

func XmlTextReaderConstValue(reader *XmlTextReader) string {
	var c_ret *C.char
	var g_ret string
	c_reader := reader.handler
	c_ret = (*C.char)(unsafe.Pointer(C.xmlTextReaderConstValue(c_reader)))
	g_ret = C.GoString(c_ret)
	return g_ret
}

func XmlTextReaderDepth(reader *XmlTextReader) int {
	var c_ret C.int
	var g_ret int
	c_reader := reader.handler
	c_ret = C.xmlTextReaderDepth(c_reader)
	g_ret = int(c_ret)
	return g_ret
}

func XmlTextReaderHasValue(reader *XmlTextReader) int {
	var c_ret C.int
	var g_ret int
	c_reader := reader.handler
	c_ret = C.xmlTextReaderHasValue(c_reader)
	g_ret = int(c_ret)
	return g_ret
}

func XmlTextReaderIsEmptyElement(reader *XmlTextReader) int {
	var c_ret C.int
	var g_ret int
	c_reader := reader.handler
	c_ret = C.xmlTextReaderIsEmptyElement(c_reader)
	g_ret = int(c_ret)
	return g_ret
}

func XmlTextReaderNodeType(reader *XmlTextReader) int {
	var c_ret C.int
	var g_ret int
	c_reader := reader.handler
	c_ret = C.xmlTextReaderNodeType(c_reader)
	g_ret = int(c_ret)
	return g_ret
}

func XmlTextReaderRead(reader *XmlTextReader) int {
	var c_ret C.int
	var g_ret int
	c_reader := reader.handler
	c_ret = C.xmlTextReaderRead(c_reader)
	g_ret = int(c_ret)
	return g_ret
}

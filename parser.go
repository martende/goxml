package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/parser.h>
*/
import "C"
func XmlCleanupParser() {
	C.xmlCleanupParser()
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

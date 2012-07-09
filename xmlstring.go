package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/xmlstring.h>
*/
import "C"

import "unsafe"
 
func XmlStrlen(str string) int {
	var c_ret C.int
	var g_ret int
	c_str := (*C.xmlChar)(unsafe.Pointer(C.CString(str)))
	c_ret = C.xmlStrlen(c_str)
	g_ret = int(c_ret)
	return g_ret
}

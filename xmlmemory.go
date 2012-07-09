package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/xmlmemory.h>
*/
import "C"


func XmlMemBlocks() int {
	var c_ret C.int
	var g_ret int
	c_ret = C.xmlMemBlocks()
	g_ret = int(c_ret)
	return g_ret
}

func XmlMemoryDump() {
	C.xmlMemoryDump()
}

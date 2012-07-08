package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/xmlmemory.h>
*/
import "C"

import "unsafe"

func XmlMemoryDump() {
	C.xmlMemoryDump()
}

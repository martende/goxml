package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/xmlmemory.h>

*/
import "C"




/* 
	   Function: xmlMemoryDump
	   ReturnType: void
	   Args: ((None, ['void'], None),)
*/
func XmlMemoryDump() {


	C.xmlMemoryDump()




}



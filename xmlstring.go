package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/xmlstring.h>
*/
import "C"




/* 
	   Function: xmlStrlen
	   ReturnType: int
	   Args: (('str', ['xmlChar', '*'], None),)
*/
/*

	Warn: str xmlChar* No converter to C(go2cConverter)

func XmlStrlen(str string) int {
str
	c_ret := C.xmlStrlen(c_str)
	return int(c_ret)
}

*/



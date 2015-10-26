package goxml

/*
#cgo pkg-config: libxml-2.0
#include <libxml/xpathInternals.h>

*/
import "C"
import "unsafe"

/*
   Function: xmlXPathRegisterNs
   ReturnType: int
   Args: (('ctxt', ['xmlXPathContextPtr'], None), ('prefix', ['xmlChar', '*'], None), ('ns_uri', ['xmlChar', '*'], None))
*/
func XmlXPathRegisterNs(ctxt *XmlXPathContext, prefix string, ns_uri string) int {
	var c_ctxt C.xmlXPathContextPtr = nil
	if ctxt != nil {
		c_ctxt = (C.xmlXPathContextPtr)(ctxt.handler)
	}
	c_prefix := (*C.xmlChar)(unsafe.Pointer(C.CString(prefix)))
	defer C.free(unsafe.Pointer(c_prefix))
	c_ns_uri := (*C.xmlChar)(unsafe.Pointer(C.CString(ns_uri)))
	defer C.free(unsafe.Pointer(c_ns_uri))

	c_ret := C.xmlXPathRegisterNs(c_ctxt, c_prefix, c_ns_uri)

	return int(c_ret)
}

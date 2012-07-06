package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/parser.h>
*/
import "C"
type XmlParserCtxt struct {
	handler C.xmlParserCtxtPtr
}
/*
func (this *XmlParserCtxt) GetSax() struct _xmlSAXHandler* {
}
*/

package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/tree.h>
*/
import "C"
type XmlNode struct {
	handler C.xmlNodePtr
}
type XmlDoc struct {
	handler C.xmlDocPtr
}
/*
type XmlParserCtxt struct {
	handler C.xmlParserCtxtPtr
}

func (this *XmlParserCtxt) GetValid() int{
	return int(this.handler.valid)
}
*/

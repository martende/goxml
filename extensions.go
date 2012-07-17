package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/xpath.h>

*/
import "C"
import "unsafe"
//import "fmt"
//import "reflect"

func (this *XmlXPathObject) GetNodesetval() []*XmlNode {
	if this.handler.nodesetval == nil || this.handler.nodesetval.nodeNr == 0 {
		return nil
	}
	l:=int(this.handler.nodesetval.nodeNr)
	ret:=make([]*XmlNode,l)
	
	off:=(uintptr)(unsafe.Pointer(this.handler.nodesetval.nodeTab))
	
	for i:=0;i<l;i++ {
		t:=*(*C.xmlNodePtr)(unsafe.Pointer(off))
		ret[i] = &XmlNode{handler:t}
		off+=unsafe.Sizeof(t)
	}
	return ret
}
/**
	static xmlNodePtr xmlXPathNodeSetDupNs(xmlNodePtr node, xmlNsPtr ns) hack
**/
func (this *XmlNode) GetNodesetval() *XmlNs {
	return nil
}

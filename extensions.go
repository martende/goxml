package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/xpath.h>

*/
import "C"
import "unsafe"
import "fmt"
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
	GetChildren extension - return all ChildNodes as list
*/
func (this *XmlNode ) GetAllChildren() []*XmlNode {
	l:=0
	
	for cur_node:= this.handler.children; cur_node!=nil; cur_node = cur_node.next {
		fmt.Printf("GetAllChildren %v\n",cur_node);
	
		l++
	}
	ret:=make([]*XmlNode,l)
	i:=0
	for cur_node:= this.handler.children; cur_node!=nil; cur_node = cur_node.next {
		ret[i] =  &XmlNode{handler:(C.xmlNodePtr)(unsafe.Pointer(this.handler))}
		i++;
	}
	return ret
}

/**
	GetProperties extension - return all attributes as list
*/
func (this *XmlNode ) GetAllProperties() []*XmlAttr {
	fmt.Printf("GetAllProperties\n");
	l:=0
	for cur_node:= this.handler.properties; cur_node!=nil; cur_node = cur_node.next {
		l++
	}
	fmt.Printf("GetAllProperties l=%d\n",l);
	ret:=make([]*XmlAttr,l)
	i:=0
	for cur_node:= this.handler.children; cur_node!=nil; cur_node = cur_node.next {
		ret[i] =  &XmlAttr{handler:(C.xmlAttrPtr)(unsafe.Pointer(this.handler))}
		i++;
	}
	return ret
}

/**
	static xmlNodePtr xmlXPathNodeSetDupNs(xmlNodePtr node, xmlNsPtr ns) hack
	node - is ns 
	next - is node
**/
func (this *XmlNode) ConverttoNs() *XmlNs {
	return &XmlNs{handler:(C.xmlNsPtr)(unsafe.Pointer(this.handler))}
}

func (this *XmlNs) ConverttoNode() *XmlNode {
	return &XmlNode{handler:(C.xmlNodePtr)(unsafe.Pointer(this.handler))}
}


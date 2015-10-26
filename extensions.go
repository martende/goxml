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
	l := int(this.handler.nodesetval.nodeNr)
	ret := make([]*XmlNode, l)

	off := (uintptr)(unsafe.Pointer(this.handler.nodesetval.nodeTab))

	for i := 0; i < l; i++ {
		t := *(*C.xmlNodePtr)(unsafe.Pointer(off))
		ret[i] = &XmlNode{handler: t}
		off += unsafe.Sizeof(t)
	}
	return ret
}
func (this *XmlNode) GetHandler() uint {
	return uint((uintptr)(unsafe.Pointer(this.handler)))
}

/**
GetChildren extension - return all ChildNodes as list
*/
func (this *XmlNode) GetAllChildren() []*XmlNode {
	l := 0
	for cur_node := this.handler.children; cur_node != nil; cur_node = cur_node.next {
		l++
	}
	ret := make([]*XmlNode, l)
	i := 0
	for cur_node := this.handler.children; cur_node != nil; cur_node = cur_node.next {
		ret[i] = &XmlNode{handler: (C.xmlNodePtr)(unsafe.Pointer(cur_node))}
		i++
	}
	return ret
}

/**
GetProperties extension - return all attributes as list
*/
func (this *XmlNode) GetAllProperties() []*XmlAttr {
	l := 0
	for cur_node := this.handler.properties; cur_node != nil; cur_node = cur_node.next {
		l++
	}
	ret := make([]*XmlAttr, l)
	i := 0
	for cur_node := this.handler.properties; cur_node != nil; cur_node = cur_node.next {
		ret[i] = &XmlAttr{handler: (C.xmlAttrPtr)(unsafe.Pointer(cur_node))}
		i++
	}
	return ret
}

/**
GetProperties extension - return all attributes as map
*/
func (this *XmlNode) GetMapProperties() map[string]string {
	var ret = map[string]string{}
	for cur_node := this.handler.properties; cur_node != nil; cur_node = cur_node.next {

		if cur_node.name != nil {
			name := C.GoString((*C.char)(unsafe.Pointer(cur_node.name)))
			value := ""
			if cur_node.children != nil {
				content := C.xmlNodeGetContent((*C.xmlNode)(unsafe.Pointer(cur_node.children)))
				if content != nil {
					value = C.GoString((*C.char)(unsafe.Pointer(content)))
				}
			}
			ret[name] = value
		}
	}
	return ret
}

/**
	static xmlNodePtr xmlXPathNodeSetDupNs(xmlNodePtr node, xmlNsPtr ns) hack
	node - is ns
	next - is node
**/
func (this *XmlNode) ConverttoNs() *XmlNs {
	return &XmlNs{handler: (C.xmlNsPtr)(unsafe.Pointer(this.handler))}
}

func (this *XmlNs) ConverttoNode() *XmlNode {
	return &XmlNode{handler: (C.xmlNodePtr)(unsafe.Pointer(this.handler))}
}

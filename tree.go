package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/tree.h>
*/
import "C"
func XmlAddChild(parent *XmlNode,cur *XmlNode) *XmlNode {
	var c_ret C.xmlNodePtr
	g_ret := &XmlNode{}
	c_parent := parent.handler
	c_cur := cur.handler
	c_ret = C.xmlAddChild(c_parent,c_cur)
	g_ret.handler = c_ret
	return g_ret
}

func XmlFreeDoc(cur *XmlDoc) {
	c_cur := cur.handler
	C.xmlFreeDoc(c_cur)
}

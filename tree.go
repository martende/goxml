package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/tree.h>
*/
import "C"
import "unsafe"



/*
	Element ns has not registered type xmlNs* 
	Element ns not recognized getter for elType:xmlNs* goType:xmlNs*
	Element properties has not registered type struct _xmlAttr* 
	Element properties not recognized getter for elType:struct _xmlAttr* goType:struct _xmlAttr*
	Element nsDef has not registered type xmlNs* 
	Element nsDef not recognized getter for elType:xmlNs* goType:xmlNs*
	Element line has not registered type unsigned short 
	Element line not recognized getter for elType:unsigned short goType:unsigned short
	Element extra has not registered type unsigned short 
	Element extra not recognized getter for elType:unsigned short goType:unsigned short

*/
type XmlNode struct {
	handler C.xmlNodePtr
	// _private void* // Private
	// type xmlElementType // Private
	_children *XmlNode
	_last *XmlNode
	_parent *XmlNode
	_next *XmlNode
	_prev *XmlNode
	_doc *XmlDoc
	// psvi void* // Private
}
func (this *XmlNode) GetName() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.name)))
}
func (this *XmlNode) GetChildren() *XmlNode {
	if this.handler.children == nil {
		return nil
	}
	if this._children == nil {
		this._children = &XmlNode{}
	}
	this._children.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.children))
	return this._children
}
func (this *XmlNode) GetLast() *XmlNode {
	if this.handler.last == nil {
		return nil
	}
	if this._last == nil {
		this._last = &XmlNode{}
	}
	this._last.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.last))
	return this._last
}
func (this *XmlNode) GetParent() *XmlNode {
	if this.handler.parent == nil {
		return nil
	}
	if this._parent == nil {
		this._parent = &XmlNode{}
	}
	this._parent.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.parent))
	return this._parent
}
func (this *XmlNode) GetNext() *XmlNode {
	if this.handler.next == nil {
		return nil
	}
	if this._next == nil {
		this._next = &XmlNode{}
	}
	this._next.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.next))
	return this._next
}
func (this *XmlNode) GetPrev() *XmlNode {
	if this.handler.prev == nil {
		return nil
	}
	if this._prev == nil {
		this._prev = &XmlNode{}
	}
	this._prev.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.prev))
	return this._prev
}
func (this *XmlNode) GetDoc() *XmlDoc {
	if this.handler.doc == nil {
		return nil
	}
	if this._doc == nil {
		this._doc = &XmlDoc{}
	}
	this._doc.handler = (C.xmlDocPtr)(unsafe.Pointer(this.handler.doc))
	return this._doc
}
/*
func (this *XmlNode) GetNs() xmlNs* {
	return int(this.handler.ns)
}
*/
func (this *XmlNode) GetContent() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.content)))
}
/*
func (this *XmlNode) GetProperties() struct _xmlAttr* {
	return int(this.handler.properties)
}
*/
/*
func (this *XmlNode) GetNsDef() xmlNs* {
	return int(this.handler.nsDef)
}
*/
/*
func (this *XmlNode) GetLine() unsigned short {
	return int(this.handler.line)
}
*/
/*
func (this *XmlNode) GetExtra() unsigned short {
	return int(this.handler.extra)
}
*/
type XmlNs struct {
	handler C.xmlNsPtr
	_next *XmlNs
	// type xmlNsType // Private
	// _private void* // Private
	_context *XmlDoc
}
func (this *XmlNs) GetNext() *XmlNs {
	if this.handler.next == nil {
		return nil
	}
	if this._next == nil {
		this._next = &XmlNs{}
	}
	this._next.handler = (C.xmlNsPtr)(unsafe.Pointer(this.handler.next))
	return this._next
}
func (this *XmlNs) GetHref() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.href)))
}
func (this *XmlNs) GetPrefix() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.prefix)))
}
func (this *XmlNs) GetContext() *XmlDoc {
	if this.handler.context == nil {
		return nil
	}
	if this._context == nil {
		this._context = &XmlDoc{}
	}
	this._context.handler = (C.xmlDocPtr)(unsafe.Pointer(this.handler.context))
	return this._context
}
type XmlDoc struct {
	handler C.xmlDocPtr
	// _private void* // Private
	// type xmlElementType // Private
	_children *XmlNode
	_last *XmlNode
	_parent *XmlNode
	_next *XmlNode
	_prev *XmlNode
	_doc *XmlDoc
	_intSubset *XmlDtd
	_extSubset *XmlDtd
	_oldNs *XmlNs
	// ids void* // Private
	// refs void* // Private
	_dict *XmlDict
	// psvi void* // Private
}
func (this *XmlDoc) GetName() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.name)))
}
func (this *XmlDoc) GetChildren() *XmlNode {
	if this.handler.children == nil {
		return nil
	}
	if this._children == nil {
		this._children = &XmlNode{}
	}
	this._children.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.children))
	return this._children
}
func (this *XmlDoc) GetLast() *XmlNode {
	if this.handler.last == nil {
		return nil
	}
	if this._last == nil {
		this._last = &XmlNode{}
	}
	this._last.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.last))
	return this._last
}
func (this *XmlDoc) GetParent() *XmlNode {
	if this.handler.parent == nil {
		return nil
	}
	if this._parent == nil {
		this._parent = &XmlNode{}
	}
	this._parent.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.parent))
	return this._parent
}
func (this *XmlDoc) GetNext() *XmlNode {
	if this.handler.next == nil {
		return nil
	}
	if this._next == nil {
		this._next = &XmlNode{}
	}
	this._next.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.next))
	return this._next
}
func (this *XmlDoc) GetPrev() *XmlNode {
	if this.handler.prev == nil {
		return nil
	}
	if this._prev == nil {
		this._prev = &XmlNode{}
	}
	this._prev.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.prev))
	return this._prev
}
func (this *XmlDoc) GetDoc() *XmlDoc {
	if this.handler.doc == nil {
		return nil
	}
	if this._doc == nil {
		this._doc = &XmlDoc{}
	}
	this._doc.handler = (C.xmlDocPtr)(unsafe.Pointer(this.handler.doc))
	return this._doc
}
func (this *XmlDoc) GetCompression() int {
	return int(this.handler.compression)
}
func (this *XmlDoc) GetStandalone() int {
	return int(this.handler.standalone)
}
func (this *XmlDoc) GetIntSubset() *XmlDtd {
	if this.handler.intSubset == nil {
		return nil
	}
	if this._intSubset == nil {
		this._intSubset = &XmlDtd{}
	}
	this._intSubset.handler = (C.xmlDtdPtr)(unsafe.Pointer(this.handler.intSubset))
	return this._intSubset
}
func (this *XmlDoc) GetExtSubset() *XmlDtd {
	if this.handler.extSubset == nil {
		return nil
	}
	if this._extSubset == nil {
		this._extSubset = &XmlDtd{}
	}
	this._extSubset.handler = (C.xmlDtdPtr)(unsafe.Pointer(this.handler.extSubset))
	return this._extSubset
}
func (this *XmlDoc) GetOldNs() *XmlNs {
	if this.handler.oldNs == nil {
		return nil
	}
	if this._oldNs == nil {
		this._oldNs = &XmlNs{}
	}
	this._oldNs.handler = (C.xmlNsPtr)(unsafe.Pointer(this.handler.oldNs))
	return this._oldNs
}
func (this *XmlDoc) GetVersion() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.version)))
}
func (this *XmlDoc) GetEncoding() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.encoding)))
}
func (this *XmlDoc) GetURL() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.URL)))
}
func (this *XmlDoc) GetCharset() int {
	return int(this.handler.charset)
}
func (this *XmlDoc) GetDict() *XmlDict {
	if this.handler.dict == nil {
		return nil
	}
	if this._dict == nil {
		this._dict = &XmlDict{}
	}
	this._dict.handler = (C.xmlDictPtr)(unsafe.Pointer(this.handler.dict))
	return this._dict
}
func (this *XmlDoc) GetParseFlags() int {
	return int(this.handler.parseFlags)
}
func (this *XmlDoc) GetProperties() int {
	return int(this.handler.properties)
}
type XmlDtd struct {
	handler C.xmlDtdPtr
	// _private void* // Private
	// type xmlElementType // Private
	_children *XmlNode
	_last *XmlNode
	_parent *XmlDoc
	_next *XmlNode
	_prev *XmlNode
	_doc *XmlDoc
	// notations void* // Private
	// elements void* // Private
	// attributes void* // Private
	// entities void* // Private
	// pentities void* // Private
}
func (this *XmlDtd) GetName() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.name)))
}
func (this *XmlDtd) GetChildren() *XmlNode {
	if this.handler.children == nil {
		return nil
	}
	if this._children == nil {
		this._children = &XmlNode{}
	}
	this._children.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.children))
	return this._children
}
func (this *XmlDtd) GetLast() *XmlNode {
	if this.handler.last == nil {
		return nil
	}
	if this._last == nil {
		this._last = &XmlNode{}
	}
	this._last.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.last))
	return this._last
}
func (this *XmlDtd) GetParent() *XmlDoc {
	if this.handler.parent == nil {
		return nil
	}
	if this._parent == nil {
		this._parent = &XmlDoc{}
	}
	this._parent.handler = (C.xmlDocPtr)(unsafe.Pointer(this.handler.parent))
	return this._parent
}
func (this *XmlDtd) GetNext() *XmlNode {
	if this.handler.next == nil {
		return nil
	}
	if this._next == nil {
		this._next = &XmlNode{}
	}
	this._next.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.next))
	return this._next
}
func (this *XmlDtd) GetPrev() *XmlNode {
	if this.handler.prev == nil {
		return nil
	}
	if this._prev == nil {
		this._prev = &XmlNode{}
	}
	this._prev.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.prev))
	return this._prev
}
func (this *XmlDtd) GetDoc() *XmlDoc {
	if this.handler.doc == nil {
		return nil
	}
	if this._doc == nil {
		this._doc = &XmlDoc{}
	}
	this._doc.handler = (C.xmlDocPtr)(unsafe.Pointer(this.handler.doc))
	return this._doc
}
func (this *XmlDtd) GetExternalID() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.ExternalID)))
}
func (this *XmlDtd) GetSystemID() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.SystemID)))
}

/* 
	   Function: xmlDocGetRootElement
	   ReturnType: xmlNodePtr
	   Args: (('doc', ['xmlDocPtr'], None),)
*/
func XmlDocGetRootElement(doc *XmlDoc) *XmlNode {
	var c_doc C.xmlDocPtr=nil
	if doc !=nil { c_doc = (C.xmlDocPtr)(doc.handler) }

	c_ret := C.xmlDocGetRootElement(c_doc)



	if c_ret == nil {return nil}
	return &XmlNode{handler:(C.xmlNodePtr)(c_ret)}
}
/* 
	   Function: xmlFreeDoc
	   ReturnType: void
	   Args: (('cur', ['xmlDocPtr'], None),)
*/
func XmlFreeDoc(cur *XmlDoc) {
	var c_cur C.xmlDocPtr=nil
	if cur !=nil { c_cur = (C.xmlDocPtr)(cur.handler) }

	C.xmlFreeDoc(c_cur)




}



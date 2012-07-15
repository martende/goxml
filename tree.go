package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/tree.h>
int XmlNode_fetch_type(xmlNodePtr h) {
	return (int)h->type;
}

*/
import "C"
import "unsafe"
import "fmt"
import "os"



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
	_children *XmlNode
	_last *XmlNode
	_parent *XmlNode
	_next *XmlNode
	_prev *XmlNode
	_doc *XmlDoc
	// psvi void* // Private
}
func (this *XmlNode) GetType() int {
	return int(C.XmlNode_fetch_type(this.handler))
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
/*
	Element next has not registered type struct _xmlAttr* 
	Element next not recognized getter for elType:struct _xmlAttr* goType:struct _xmlAttr*
	Element prev has not registered type struct _xmlAttr* 
	Element prev not recognized getter for elType:struct _xmlAttr* goType:struct _xmlAttr*
	Element ns has not registered type xmlNs* 
	Element ns not recognized getter for elType:xmlNs* goType:xmlNs*
	Element atype has not registered type xmlAttributeType 
	Element atype not recognized getter for elType:xmlAttributeType goType:xmlAttributeType

*/
type XmlAttr struct {
	handler C.xmlAttrPtr
	// _private void* // Private
	// type xmlElementType // Private
	_children *XmlNode
	_last *XmlNode
	_parent *XmlNode
	_doc *XmlDoc
	// psvi void* // Private
}
func (this *XmlAttr) GetName() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.name)))
}
func (this *XmlAttr) GetChildren() *XmlNode {
	if this.handler.children == nil {
		return nil
	}
	if this._children == nil {
		this._children = &XmlNode{}
	}
	this._children.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.children))
	return this._children
}
func (this *XmlAttr) GetLast() *XmlNode {
	if this.handler.last == nil {
		return nil
	}
	if this._last == nil {
		this._last = &XmlNode{}
	}
	this._last.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.last))
	return this._last
}
func (this *XmlAttr) GetParent() *XmlNode {
	if this.handler.parent == nil {
		return nil
	}
	if this._parent == nil {
		this._parent = &XmlNode{}
	}
	this._parent.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.parent))
	return this._parent
}
/*
func (this *XmlAttr) GetNext() struct _xmlAttr* {
	return int(this.handler.next)
}
*/
/*
func (this *XmlAttr) GetPrev() struct _xmlAttr* {
	return int(this.handler.prev)
}
*/
func (this *XmlAttr) GetDoc() *XmlDoc {
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
func (this *XmlAttr) GetNs() xmlNs* {
	return int(this.handler.ns)
}
*/
/*
func (this *XmlAttr) GetAtype() xmlAttributeType {
	return int(this.handler.atype)
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
	   Function: xmlSaveFormatFileEnc
	   ReturnType: int
	   Args: (('filename', ['char', '*'], None), ('cur', ['xmlDocPtr'], None), ('encoding', ['char', '*'], None), ('format', ['int'], None))
*/
func XmlSaveFormatFileEnc(filename string,cur *XmlDoc,encoding string,format int) int {
	c_filename:= (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(c_filename))
	var c_cur C.xmlDocPtr=nil
	if cur !=nil { c_cur = (C.xmlDocPtr)(cur.handler) }
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	c_format := C.int(format)

	c_ret := C.xmlSaveFormatFileEnc(c_filename,c_cur,c_encoding,c_format)



	return int(c_ret)
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
	   Function: xmlNewNode
	   ReturnType: xmlNodePtr
	   Args: (('ns', ['xmlNsPtr'], None), ('name', ['xmlChar', '*'], None))
*/
func XmlNewNode(ns *XmlNs,name string) (g_ret *XmlNode,err error) {
	var c_ns C.xmlNsPtr=nil
	if ns !=nil { c_ns = (C.xmlNsPtr)(ns.handler) }
	c_name:= (*C.xmlChar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(c_name))

	c_ret := C.xmlNewNode(c_ns,c_name)

	if c_ret == nil {
		err = fmt.Errorf("xmlNewNode errno %d" ,c_ret)
	} else {
		g_ret =  &XmlNode{handler:(C.xmlNodePtr)(c_ret)}
	}
	return
}
/* 
	   Function: xmlAddChild
	   ReturnType: xmlNodePtr
	   Args: (('parent', ['xmlNodePtr'], None), ('cur', ['xmlNodePtr'], None))
*/
func XmlAddChild(parent *XmlNode,cur *XmlNode) (g_ret *XmlNode,err error) {
	var c_parent C.xmlNodePtr=nil
	if parent !=nil { c_parent = (C.xmlNodePtr)(parent.handler) }
	var c_cur C.xmlNodePtr=nil
	if cur !=nil { c_cur = (C.xmlNodePtr)(cur.handler) }

	c_ret := C.xmlAddChild(c_parent,c_cur)

	if c_ret == nil {
		err = fmt.Errorf("xmlAddChild errno %d" ,c_ret)
	} else {
		g_ret =  &XmlNode{handler:(C.xmlNodePtr)(c_ret)}
	}
	return
}
/* 
	   Function: xmlDocDump
	   ReturnType: int
	   Args: (('f', ['FILE', '*'], None), ('cur', ['xmlDocPtr'], None))
*/
func XmlDocDump(f *os.File,cur *XmlDoc) int {
	var c_f *C.FILE
	{
		tp:= (*C.char)(unsafe.Pointer(C.CString("w")));
		defer C.free(unsafe.Pointer(tp));
		c_f = C.fdopen((C.int)(f.Fd()),tp)
	}
	
	var c_cur C.xmlDocPtr=nil
	if cur !=nil { c_cur = (C.xmlDocPtr)(cur.handler) }

	c_ret := C.xmlDocDump(c_f,c_cur)



	return int(c_ret)
}
/* 
	   Function: xmlNewProp
	   ReturnType: xmlAttrPtr
	   Args: (('node', ['xmlNodePtr'], None), ('name', ['xmlChar', '*'], None), ('value', ['xmlChar', '*'], None))
*/
func XmlNewProp(node *XmlNode,name string,value string) (g_ret *XmlAttr,err error) {
	var c_node C.xmlNodePtr=nil
	if node !=nil { c_node = (C.xmlNodePtr)(node.handler) }
	c_name:= (*C.xmlChar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(c_name))
	c_value:= (*C.xmlChar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(c_value))

	c_ret := C.xmlNewProp(c_node,c_name,c_value)

	if c_ret == nil {
		err = fmt.Errorf("xmlNewProp errno %d" ,c_ret)
	} else {
		g_ret =  &XmlAttr{handler:(C.xmlAttrPtr)(c_ret)}
	}
	return
}
/* 
	   Function: xmlDocSetRootElement
	   ReturnType: xmlNodePtr
	   Args: (('doc', ['xmlDocPtr'], None), ('root', ['xmlNodePtr'], None))
*/
func XmlDocSetRootElement(doc *XmlDoc,root *XmlNode) (g_ret *XmlNode,err error) {
	var c_doc C.xmlDocPtr=nil
	if doc !=nil { c_doc = (C.xmlDocPtr)(doc.handler) }
	var c_root C.xmlNodePtr=nil
	if root !=nil { c_root = (C.xmlNodePtr)(root.handler) }

	c_ret := C.xmlDocSetRootElement(c_doc,c_root)

	if c_ret == nil {
		err = fmt.Errorf("xmlDocSetRootElement errno %d" ,c_ret)
	} else {
		g_ret =  &XmlNode{handler:(C.xmlNodePtr)(c_ret)}
	}
	return
}
/* 
	   Function: xmlNewChild
	   ReturnType: xmlNodePtr
	   Args: (('parent', ['xmlNodePtr'], None), ('ns', ['xmlNsPtr'], None), ('name', ['xmlChar', '*'], None), ('content', ['xmlChar', '*'], None))
*/
func XmlNewChild(parent *XmlNode,ns *XmlNs,name string,content string) (g_ret *XmlNode,err error) {
	var c_parent C.xmlNodePtr=nil
	if parent !=nil { c_parent = (C.xmlNodePtr)(parent.handler) }
	var c_ns C.xmlNsPtr=nil
	if ns !=nil { c_ns = (C.xmlNsPtr)(ns.handler) }
	c_name:= (*C.xmlChar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(c_name))
	c_content:= (*C.xmlChar)(unsafe.Pointer(C.CString(content)))
	defer C.free(unsafe.Pointer(c_content))

	c_ret := C.xmlNewChild(c_parent,c_ns,c_name,c_content)

	if c_ret == nil {
		err = fmt.Errorf("xmlNewChild errno %d" ,c_ret)
	} else {
		g_ret =  &XmlNode{handler:(C.xmlNodePtr)(c_ret)}
	}
	return
}
/* 
	   Function: xmlCreateIntSubset
	   ReturnType: xmlDtdPtr
	   Args: (('doc', ['xmlDocPtr'], None), ('name', ['xmlChar', '*'], None), ('ExternalID', ['xmlChar', '*'], None), ('SystemID', ['xmlChar', '*'], None))
*/
func XmlCreateIntSubset(doc *XmlDoc,name string,ExternalID string,SystemID string) (g_ret *XmlDtd,err error) {
	var c_doc C.xmlDocPtr=nil
	if doc !=nil { c_doc = (C.xmlDocPtr)(doc.handler) }
	c_name:= (*C.xmlChar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(c_name))
	c_ExternalID:= (*C.xmlChar)(unsafe.Pointer(C.CString(ExternalID)))
	defer C.free(unsafe.Pointer(c_ExternalID))
	c_SystemID:= (*C.xmlChar)(unsafe.Pointer(C.CString(SystemID)))
	defer C.free(unsafe.Pointer(c_SystemID))

	c_ret := C.xmlCreateIntSubset(c_doc,c_name,c_ExternalID,c_SystemID)

	if c_ret == nil {
		err = fmt.Errorf("xmlCreateIntSubset errno %d" ,c_ret)
	} else {
		g_ret =  &XmlDtd{handler:(C.xmlDtdPtr)(c_ret)}
	}
	return
}
/* 
	   Function: xmlNewDoc
	   ReturnType: xmlDocPtr
	   Args: (('version', ['xmlChar', '*'], None),)
*/
func XmlNewDoc(version string) (g_ret *XmlDoc,err error) {
	c_version:= (*C.xmlChar)(unsafe.Pointer(C.CString(version)))
	defer C.free(unsafe.Pointer(c_version))

	c_ret := C.xmlNewDoc(c_version)

	if c_ret == nil {
		err = fmt.Errorf("xmlNewDoc errno %d" ,c_ret)
	} else {
		g_ret =  &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
	}
	return
}
/* 
	   Function: xmlNewText
	   ReturnType: xmlNodePtr
	   Args: (('content', ['xmlChar', '*'], None),)
*/
func XmlNewText(content string) (g_ret *XmlNode,err error) {
	c_content:= (*C.xmlChar)(unsafe.Pointer(C.CString(content)))
	defer C.free(unsafe.Pointer(c_content))

	c_ret := C.xmlNewText(c_content)

	if c_ret == nil {
		err = fmt.Errorf("xmlNewText errno %d" ,c_ret)
	} else {
		g_ret =  &XmlNode{handler:(C.xmlNodePtr)(c_ret)}
	}
	return
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



package goxml

/*
#cgo pkg-config: libxml-2.0
#include <libxml/xpath.h>

*/
import "C"
import "unsafe"
import "fmt"

/*
	Element nodesetval has not registered type xmlNodeSetPtr
	Element nodesetval not recognized getter for elType:xmlNodeSetPtr goType:xmlNodeSetPtr
	Element floatval has not registered type double
	Element floatval not recognized getter for elType:double goType:double

*/
type XmlXPathObject struct {
	handler C.xmlXPathObjectPtr
	// type xmlXPathObjectType // Private
	// user void* // Private
	// user2 void* // Private
}

/*
func (this *XmlXPathObject) GetNodesetval() xmlNodeSetPtr {
	return int(this.handler.nodesetval)
}
*/
func (this *XmlXPathObject) GetBoolval() int {
	return int(this.handler.boolval)
}

/*
func (this *XmlXPathObject) GetFloatval() double {
	return int(this.handler.floatval)
}
*/
func (this *XmlXPathObject) GetStringval() string {
	if this.handler.stringval == nil {
		return ""
	}
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.stringval)))
}
func (this *XmlXPathObject) GetIndex() int {
	return int(this.handler.index)
}
func (this *XmlXPathObject) GetIndex2() int {
	return int(this.handler.index2)
}

/*
	Element varHash has not registered type xmlHashTablePtr
	Element varHash not recognized getter for elType:xmlHashTablePtr goType:xmlHashTablePtr
	Element types has not registered type xmlXPathTypePtr
	Element types not recognized getter for elType:xmlXPathTypePtr goType:xmlXPathTypePtr
	Element funcHash has not registered type xmlHashTablePtr
	Element funcHash not recognized getter for elType:xmlHashTablePtr goType:xmlHashTablePtr
	Element axis has not registered type xmlXPathAxisPtr
	Element axis not recognized getter for elType:xmlXPathAxisPtr goType:xmlXPathAxisPtr
	Element namespaces has not registered type xmlNsPtr*
	Element namespaces not recognized getter for elType:xmlNsPtr* goType:xmlNsPtr*
	Element nsHash has not registered type xmlHashTablePtr
	Element nsHash not recognized getter for elType:xmlHashTablePtr goType:xmlHashTablePtr
	Element varLookupFunc has not registered type xmlXPathVariableLookupFunc
	Element varLookupFunc not recognized getter for elType:xmlXPathVariableLookupFunc goType:xmlXPathVariableLookupFunc
	Element funcLookupFunc has not registered type xmlXPathFuncLookupFunc
	Element funcLookupFunc not recognized getter for elType:xmlXPathFuncLookupFunc goType:xmlXPathFuncLookupFunc
	Element tmpNsList has not registered type xmlNsPtr*
	Element tmpNsList not recognized getter for elType:xmlNsPtr* goType:xmlNsPtr*
	Element error has not registered type xmlStructuredErrorFunc
	Element error not recognized getter for elType:xmlStructuredErrorFunc goType:xmlStructuredErrorFunc
	Element lastError has not registered type xmlError
	Element lastError not recognized getter for elType:xmlError goType:xmlError

*/
type XmlXPathContext struct {
	handler C.xmlXPathContextPtr
	_doc    *XmlDoc
	_node   *XmlNode
	// user void* // Private
	_here   *XmlNode
	_origin *XmlNode
	// varLookupData void* // Private
	// extra void* // Private
	// funcLookupData void* // Private
	// userData void* // Private
	_debugNode *XmlNode
	_dict      *XmlDict
	// cache void* // Private
}

func (this *XmlXPathContext) GetDoc() *XmlDoc {
	if this.handler.doc == nil {
		return nil
	}
	if this._doc == nil {
		this._doc = &XmlDoc{}
	}
	this._doc.handler = (C.xmlDocPtr)(unsafe.Pointer(this.handler.doc))
	return this._doc
}
func (this *XmlXPathContext) GetNode() *XmlNode {
	if this.handler.node == nil {
		return nil
	}
	if this._node == nil {
		this._node = &XmlNode{}
	}
	this._node.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.node))
	return this._node
}
func (this *XmlXPathContext) GetNb_variables_unused() int {
	return int(this.handler.nb_variables_unused)
}
func (this *XmlXPathContext) GetMax_variables_unused() int {
	return int(this.handler.max_variables_unused)
}

/*
func (this *XmlXPathContext) GetVarHash() xmlHashTablePtr {
	return int(this.handler.varHash)
}
*/
func (this *XmlXPathContext) GetNb_types() int {
	return int(this.handler.nb_types)
}
func (this *XmlXPathContext) GetMax_types() int {
	return int(this.handler.max_types)
}

/*
func (this *XmlXPathContext) GetTypes() xmlXPathTypePtr {
	return int(this.handler.types)
}
*/
func (this *XmlXPathContext) GetNb_funcs_unused() int {
	return int(this.handler.nb_funcs_unused)
}
func (this *XmlXPathContext) GetMax_funcs_unused() int {
	return int(this.handler.max_funcs_unused)
}

/*
func (this *XmlXPathContext) GetFuncHash() xmlHashTablePtr {
	return int(this.handler.funcHash)
}
*/
func (this *XmlXPathContext) GetNb_axis() int {
	return int(this.handler.nb_axis)
}
func (this *XmlXPathContext) GetMax_axis() int {
	return int(this.handler.max_axis)
}

/*
func (this *XmlXPathContext) GetAxis() xmlXPathAxisPtr {
	return int(this.handler.axis)
}
*/
/*
func (this *XmlXPathContext) GetNamespaces() xmlNsPtr* {
	return int(this.handler.namespaces)
}
*/
func (this *XmlXPathContext) GetNsNr() int {
	return int(this.handler.nsNr)
}
func (this *XmlXPathContext) GetContextSize() int {
	return int(this.handler.contextSize)
}
func (this *XmlXPathContext) GetProximityPosition() int {
	return int(this.handler.proximityPosition)
}
func (this *XmlXPathContext) GetXptr() int {
	return int(this.handler.xptr)
}
func (this *XmlXPathContext) GetHere() *XmlNode {
	if this.handler.here == nil {
		return nil
	}
	if this._here == nil {
		this._here = &XmlNode{}
	}
	this._here.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.here))
	return this._here
}
func (this *XmlXPathContext) GetOrigin() *XmlNode {
	if this.handler.origin == nil {
		return nil
	}
	if this._origin == nil {
		this._origin = &XmlNode{}
	}
	this._origin.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.origin))
	return this._origin
}

/*
func (this *XmlXPathContext) GetNsHash() xmlHashTablePtr {
	return int(this.handler.nsHash)
}
*/
/*
func (this *XmlXPathContext) GetVarLookupFunc() xmlXPathVariableLookupFunc {
	return int(this.handler.varLookupFunc)
}
*/
func (this *XmlXPathContext) GetFunction() string {
	if this.handler.function == nil {
		return ""
	}
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.function)))
}
func (this *XmlXPathContext) GetFunctionURI() string {
	if this.handler.functionURI == nil {
		return ""
	}
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.functionURI)))
}

/*
func (this *XmlXPathContext) GetFuncLookupFunc() xmlXPathFuncLookupFunc {
	return int(this.handler.funcLookupFunc)
}
*/
/*
func (this *XmlXPathContext) GetTmpNsList() xmlNsPtr* {
	return int(this.handler.tmpNsList)
}
*/
func (this *XmlXPathContext) GetTmpNsNr() int {
	return int(this.handler.tmpNsNr)
}

/*
func (this *XmlXPathContext) GetError() xmlStructuredErrorFunc {
	return int(this.handler.error)
}
*/
/*
func (this *XmlXPathContext) GetLastError() xmlError {
	return int(this.handler.lastError)
}
*/
func (this *XmlXPathContext) GetDebugNode() *XmlNode {
	if this.handler.debugNode == nil {
		return nil
	}
	if this._debugNode == nil {
		this._debugNode = &XmlNode{}
	}
	this._debugNode.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.debugNode))
	return this._debugNode
}
func (this *XmlXPathContext) GetDict() *XmlDict {
	if this.handler.dict == nil {
		return nil
	}
	if this._dict == nil {
		this._dict = &XmlDict{}
	}
	this._dict.handler = (C.xmlDictPtr)(unsafe.Pointer(this.handler.dict))
	return this._dict
}
func (this *XmlXPathContext) GetFlags() int {
	return int(this.handler.flags)
}

/*
   Function: xmlXPathFreeObject
   ReturnType: void
   Args: (('obj', ['xmlXPathObjectPtr'], None),)
*/
func XmlXPathFreeObject(obj *XmlXPathObject) {
	var c_obj C.xmlXPathObjectPtr = nil
	if obj != nil {
		c_obj = (C.xmlXPathObjectPtr)(obj.handler)
	}

	C.xmlXPathFreeObject(c_obj)

}

/*
   Function: xmlXPathFreeContext
   ReturnType: void
   Args: (('ctxt', ['xmlXPathContextPtr'], None),)
*/
func XmlXPathFreeContext(ctxt *XmlXPathContext) {
	var c_ctxt C.xmlXPathContextPtr = nil
	if ctxt != nil {
		c_ctxt = (C.xmlXPathContextPtr)(ctxt.handler)
	}

	C.xmlXPathFreeContext(c_ctxt)

}

/*
   Function: xmlXPathEvalExpression
   ReturnType: xmlXPathObjectPtr
   Args: (('str', ['xmlChar', '*'], None), ('ctxt', ['xmlXPathContextPtr'], None))
*/
func XmlXPathEvalExpression(str string, ctxt *XmlXPathContext) (g_ret *XmlXPathObject, err error) {
	c_str := (*C.xmlChar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(c_str))
	var c_ctxt C.xmlXPathContextPtr = nil
	if ctxt != nil {
		c_ctxt = (C.xmlXPathContextPtr)(ctxt.handler)
	}

	c_ret := C.xmlXPathEvalExpression(c_str, c_ctxt)

	if c_ret == nil {
		err = fmt.Errorf("xmlXPathEvalExpression errno %d", c_ret)
	} else {
		g_ret = &XmlXPathObject{handler: (C.xmlXPathObjectPtr)(c_ret)}
	}
	return
}

/*
   Function: xmlXPathNewContext
   ReturnType: xmlXPathContextPtr
   Args: (('doc', ['xmlDocPtr'], None),)
*/
func XmlXPathNewContext(doc *XmlDoc) (g_ret *XmlXPathContext, err error) {
	var c_doc C.xmlDocPtr = nil
	if doc != nil {
		c_doc = (C.xmlDocPtr)(doc.handler)
	}

	c_ret := C.xmlXPathNewContext(c_doc)

	if c_ret == nil {
		err = fmt.Errorf("xmlXPathNewContext errno %d", c_ret)
	} else {
		g_ret = &XmlXPathContext{handler: (C.xmlXPathContextPtr)(c_ret)}
	}
	return
}

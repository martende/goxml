package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/HTMLparser.h>
*/
import "C"
import "unsafe"



/*
	Element subelts has not registered type char** 
	Element subelts not recognized getter for type char** 
	Element attrs_opt has not registered type char** 
	Element attrs_opt not recognized getter for type char** 
	Element attrs_depr has not registered type char** 
	Element attrs_depr not recognized getter for type char** 
	Element attrs_req has not registered type char** 
	Element attrs_req not recognized getter for type char** 

*/
type HtmlElemDesc struct {
	handler C.htmlElemDescPtr
	// startTag char // Private
	// endTag char // Private
	// saveEndTag char // Private
	// empty char // Private
	// depr char // Private
	// dtd char // Private
	// isinline char // Private
}
func (this *HtmlElemDesc) GetName() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.name)))
}
func (this *HtmlElemDesc) GetDesc() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.desc)))
}
/*
func (this *HtmlElemDesc) GetSubelts() char** {
	return int(this.handler.subelts)
}
*/
func (this *HtmlElemDesc) GetDefaultsubelt() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.defaultsubelt)))
}
/*
func (this *HtmlElemDesc) GetAttrs_opt() char** {
	return int(this.handler.attrs_opt)
}
*/
/*
func (this *HtmlElemDesc) GetAttrs_depr() char** {
	return int(this.handler.attrs_depr)
}
*/
/*
func (this *HtmlElemDesc) GetAttrs_req() char** {
	return int(this.handler.attrs_req)
}
*/

/* 
	   Function: htmlAttrAllowed
	   ReturnType: htmlStatus
	   Args: ((None, ['htmlElemDesc', '*'], None), (None, ['xmlChar', '*'], None), (None, ['int'], None))
*/
func HtmlAttrAllowed(arg1 *HtmlElemDesc,arg2 string,arg3 int) int {
	var c_arg1 C.htmlElemDescPtr=nil
	if arg1 !=nil { c_arg1 = arg1.handler }
	c_arg2:= (*C.xmlChar)(unsafe.Pointer(C.CString(arg2)))
	defer C.free(unsafe.Pointer(c_arg2))
	c_arg3 := C.int(arg3)

	c_ret := C.htmlAttrAllowed(c_arg1,c_arg2,c_arg3)



	return int(c_ret)
}



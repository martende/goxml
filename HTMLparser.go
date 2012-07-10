package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/HTMLparser.h>
*/
import "C"
import "unsafe"
import "fmt"



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
}
func (this *HtmlElemDesc) GetName() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.name)))
}
func (this *HtmlElemDesc) GetStartTag() byte {
	return byte(this.handler.startTag)
}
func (this *HtmlElemDesc) GetEndTag() byte {
	return byte(this.handler.endTag)
}
func (this *HtmlElemDesc) GetSaveEndTag() byte {
	return byte(this.handler.saveEndTag)
}
func (this *HtmlElemDesc) GetEmpty() byte {
	return byte(this.handler.empty)
}
func (this *HtmlElemDesc) GetDepr() byte {
	return byte(this.handler.depr)
}
func (this *HtmlElemDesc) GetDtd() byte {
	return byte(this.handler.dtd)
}
func (this *HtmlElemDesc) GetIsinline() byte {
	return byte(this.handler.isinline)
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
	   Function: htmlAutoCloseTag
	   ReturnType: int
	   Args: (('doc', ['htmlDocPtr'], None), ('name', ['xmlChar', '*'], None), ('elem', ['htmlNodePtr'], None))
*/
func HtmlAutoCloseTag(doc *XmlDoc,name string,elem *XmlNode) int {
	var c_doc C.htmlDocPtr=nil
	if doc !=nil { c_doc = doc.handler }
	c_name:= (*C.xmlChar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(c_name))
	var c_elem C.htmlNodePtr=nil
	if elem !=nil { c_elem = elem.handler }

	c_ret := C.htmlAutoCloseTag(c_doc,c_name,c_elem)



	return int(c_ret)
}
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
/* 
	   Function: UTF8ToHtml
	   ReturnType: int
	   Args: (('out', ['unsigned char', '*'], None), ('outlen', ['int', '*'], None), ('in', ['unsigned char', '*'], None), ('inlen', ['int', '*'], None))
*/
func UTF8ToHtml(in string) (g_out string,err error) {
	c_in:= (*C.uchar)(unsafe.Pointer(C.CString(in)))
	defer C.free(unsafe.Pointer(c_in))
		c_out:= (*C.uchar)(C.calloc(  (C.size_t)( len(in)*3+ 1 )  ,1))
		defer C.free(unsafe.Pointer(c_out))
		c0_outlen:=C.int(len(in)*3+1)
		c_outlen:=&c0_outlen
		c0_inlen:=C.int(len(in)*1+1)
		c_inlen:=&c0_inlen
	c_ret := C.UTF8ToHtml(c_out,c_outlen,c_in,c_inlen)

	if c_ret != 0 {
		err = fmt.Errorf("UTF8ToHtml errno %d" ,c_ret)
	} else {
		if c_out == nil {
			g_out=""
		} else {
			g_out = C.GoString((*C.char)(unsafe.Pointer(c_out)))
		}
	}
	return
}



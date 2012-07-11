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
	   Function: htmlCtxtReadDoc
	   ReturnType: htmlDocPtr
	   Args: (('ctxt', ['xmlParserCtxtPtr'], None), ('cur', ['xmlChar', '*'], None), ('URL', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
func HtmlCtxtReadDoc(ctxt *XmlParserCtxt,cur string,URL string,encoding string,options int) (g_ret *XmlDoc,err error) {
	var c_ctxt C.xmlParserCtxtPtr=nil
	if ctxt !=nil { c_ctxt = (C.xmlParserCtxtPtr)(ctxt.handler) }
	c_cur:= (*C.xmlChar)(unsafe.Pointer(C.CString(cur)))
	defer C.free(unsafe.Pointer(c_cur))
	c_URL:= (*C.char)(unsafe.Pointer(C.CString(URL)))
	defer C.free(unsafe.Pointer(c_URL))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	c_options := C.int(options)

	c_ret := C.htmlCtxtReadDoc(c_ctxt,c_cur,c_URL,c_encoding,c_options)

	if c_ret == nil {
		err = fmt.Errorf("htmlCtxtReadDoc errno %d" ,c_ret)
	} else {
		g_ret =  &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
	}
	return
}
/* 
	   Function: htmlAutoCloseTag
	   ReturnType: int
	   Args: (('doc', ['htmlDocPtr'], None), ('name', ['xmlChar', '*'], None), ('elem', ['htmlNodePtr'], None))
*/
func HtmlAutoCloseTag(doc *XmlDoc,name string,elem *XmlNode) int {
	var c_doc C.htmlDocPtr=nil
	if doc !=nil { c_doc = (C.htmlDocPtr)(doc.handler) }
	c_name:= (*C.xmlChar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(c_name))
	var c_elem C.htmlNodePtr=nil
	if elem !=nil { c_elem = (C.htmlNodePtr)(elem.handler) }

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
	if arg1 !=nil { c_arg1 = (C.htmlElemDescPtr)(arg1.handler) }
	c_arg2:= (*C.xmlChar)(unsafe.Pointer(C.CString(arg2)))
	defer C.free(unsafe.Pointer(c_arg2))
	c_arg3 := C.int(arg3)

	c_ret := C.htmlAttrAllowed(c_arg1,c_arg2,c_arg3)



	return int(c_ret)
}
/* 
	   Function: htmlCreateMemoryParserCtxt
	   ReturnType: htmlParserCtxtPtr
	   Args: (('buffer', ['char', '*'], None), ('size', ['int'], None))
*/
func HtmlCreateMemoryParserCtxt(buffer string) (g_ret *XmlParserCtxt,err error) {
	c_buffer:= (*C.char)(unsafe.Pointer(C.CString(buffer)))
	defer C.free(unsafe.Pointer(c_buffer))
	c_size:=C.int(len(buffer)*1+1)
	c_ret := C.htmlCreateMemoryParserCtxt(c_buffer,c_size)

	if c_ret == nil {
		err = fmt.Errorf("htmlCreateMemoryParserCtxt errno %d" ,c_ret)
	} else {
		g_ret =  &XmlParserCtxt{handler:(C.xmlParserCtxtPtr)(c_ret)}
	}
	return
}
/* 
	   Function: htmlCtxtReadIO
	   ReturnType: htmlDocPtr
	   Args: (('ctxt', ['xmlParserCtxtPtr'], None), ('ioread', ['xmlInputReadCallback'], None), ('ioclose', ['xmlInputCloseCallback'], None), ('ioctx', ['void', '*'], None), ('URL', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
/*

	Warn: ioread xmlInputReadCallback Not defined
	Warn: ioclose xmlInputCloseCallback Not defined
	Warn: ioctx void* Not defined
	Warn: ioread xmlInputReadCallback No converter to C(go2cConverter)
	Warn: ioclose xmlInputCloseCallback No converter to C(go2cConverter)
	Warn: ioctx void* No converter to C(go2cConverter)

func HtmlCtxtReadIO(ctxt *XmlParserCtxt,ioread xmlInputReadCallback,ioclose xmlInputCloseCallback,ioctx void*,URL string,encoding string,options int) (g_ret *XmlDoc,err error) {
	var c_ctxt C.xmlParserCtxtPtr=nil
	if ctxt !=nil { c_ctxt = (C.xmlParserCtxtPtr)(ctxt.handler) }
	ioread
	ioclose
	ioctx
	c_URL:= (*C.char)(unsafe.Pointer(C.CString(URL)))
	defer C.free(unsafe.Pointer(c_URL))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	c_options := C.int(options)

	c_ret := C.htmlCtxtReadIO(c_ctxt,c_ioread,c_ioclose,c_ioctx,c_URL,c_encoding,c_options)

	if c_ret == nil {
		err = fmt.Errorf("htmlCtxtReadIO errno %d" ,c_ret)
	} else {
		g_ret =  &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
	}
	return
}

*/
/* 
	   Function: htmlCreatePushParserCtxt
	   ReturnType: htmlParserCtxtPtr
	   Args: (('sax', ['htmlSAXHandlerPtr'], None), ('user_data', ['void', '*'], None), ('chunk', ['char', '*'], None), ('size', ['int'], None), ('filename', ['char', '*'], None), ('enc', ['xmlCharEncoding'], None))
*/
func HtmlCreatePushParserCtxt(sax *XmlSAXHandler,chunk string,filename string,enc int) (g_ret *XmlParserCtxt,err error) {
	var c_sax C.htmlSAXHandlerPtr=nil
	if sax !=nil { c_sax = (C.htmlSAXHandlerPtr)(sax.handler) }
	c_chunk:= (*C.char)(unsafe.Pointer(C.CString(chunk)))
	defer C.free(unsafe.Pointer(c_chunk))
	c_filename:= (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(c_filename))
	c_enc := C.xmlCharEncoding(enc)
	c_size:=C.int(len(chunk)*1+1)
	c_ret := C.htmlCreatePushParserCtxt(c_sax,nil,c_chunk,c_size,c_filename,c_enc)

	if c_ret == nil {
		err = fmt.Errorf("htmlCreatePushParserCtxt errno %d" ,c_ret)
	} else {
		g_ret =  &XmlParserCtxt{handler:(C.xmlParserCtxtPtr)(c_ret)}
	}
	return
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
/* 
	   Function: htmlCtxtReadMemory
	   ReturnType: htmlDocPtr
	   Args: (('ctxt', ['xmlParserCtxtPtr'], None), ('buffer', ['char', '*'], None), ('size', ['int'], None), ('URL', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
func HtmlCtxtReadMemory(ctxt *XmlParserCtxt,buffer string,URL string,encoding string,options int) (g_ret *XmlDoc,err error) {
	var c_ctxt C.xmlParserCtxtPtr=nil
	if ctxt !=nil { c_ctxt = (C.xmlParserCtxtPtr)(ctxt.handler) }
	c_buffer:= (*C.char)(unsafe.Pointer(C.CString(buffer)))
	defer C.free(unsafe.Pointer(c_buffer))
	c_URL:= (*C.char)(unsafe.Pointer(C.CString(URL)))
	defer C.free(unsafe.Pointer(c_URL))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	c_options := C.int(options)
	c_size:=C.int(len(buffer)*1+1)
	c_ret := C.htmlCtxtReadMemory(c_ctxt,c_buffer,c_size,c_URL,c_encoding,c_options)

	if c_ret == nil {
		err = fmt.Errorf("htmlCtxtReadMemory errno %d" ,c_ret)
	} else {
		g_ret =  &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
	}
	return
}
/* 
	   Function: htmlReadFile
	   ReturnType: htmlDocPtr
	   Args: (('URL', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
func HtmlReadFile(URL string,encoding string,options int) (g_ret *XmlDoc,err error) {
	c_URL:= (*C.char)(unsafe.Pointer(C.CString(URL)))
	defer C.free(unsafe.Pointer(c_URL))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	c_options := C.int(options)

	c_ret := C.htmlReadFile(c_URL,c_encoding,c_options)

	if c_ret == nil {
		err = fmt.Errorf("htmlReadFile errno %d" ,c_ret)
	} else {
		g_ret =  &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
	}
	return
}
/* 
	   Function: htmlCtxtReadFile
	   ReturnType: htmlDocPtr
	   Args: (('ctxt', ['xmlParserCtxtPtr'], None), ('filename', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
func HtmlCtxtReadFile(ctxt *XmlParserCtxt,filename string,encoding string,options int) (g_ret *XmlDoc,err error) {
	var c_ctxt C.xmlParserCtxtPtr=nil
	if ctxt !=nil { c_ctxt = (C.xmlParserCtxtPtr)(ctxt.handler) }
	c_filename:= (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(c_filename))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	c_options := C.int(options)

	c_ret := C.htmlCtxtReadFile(c_ctxt,c_filename,c_encoding,c_options)

	if c_ret == nil {
		err = fmt.Errorf("htmlCtxtReadFile errno %d" ,c_ret)
	} else {
		g_ret =  &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
	}
	return
}
/* 
	   Function: htmlCtxtReadFd
	   ReturnType: htmlDocPtr
	   Args: (('ctxt', ['xmlParserCtxtPtr'], None), ('fd', ['int'], None), ('URL', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
func HtmlCtxtReadFd(ctxt *XmlParserCtxt,fd int,URL string,encoding string,options int) (g_ret *XmlDoc,err error) {
	var c_ctxt C.xmlParserCtxtPtr=nil
	if ctxt !=nil { c_ctxt = (C.xmlParserCtxtPtr)(ctxt.handler) }
	c_fd := C.int(fd)
	c_URL:= (*C.char)(unsafe.Pointer(C.CString(URL)))
	defer C.free(unsafe.Pointer(c_URL))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	c_options := C.int(options)

	c_ret := C.htmlCtxtReadFd(c_ctxt,c_fd,c_URL,c_encoding,c_options)

	if c_ret == nil {
		err = fmt.Errorf("htmlCtxtReadFd errno %d" ,c_ret)
	} else {
		g_ret =  &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
	}
	return
}



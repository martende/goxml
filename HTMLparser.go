package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/HTMLparser.h>
*/
import "C"
import "unsafe"
import "fmt"



/*
	Element value not recognized getter for elType:unsigned int goType:uint

*/
type HtmlEntityDesc struct {
	handler C.htmlEntityDescPtr
}
/*
func (this *HtmlEntityDesc) GetValue() uint {
	return int(this.handler.value)
}
*/
func (this *HtmlEntityDesc) GetName() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.name)))
}
func (this *HtmlEntityDesc) GetDesc() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.desc)))
}
/*
	Element subelts has not registered type char** 
	Element subelts not recognized getter for elType:char** goType:char**
	Element attrs_opt has not registered type char** 
	Element attrs_opt not recognized getter for elType:char** goType:char**
	Element attrs_depr has not registered type char** 
	Element attrs_depr not recognized getter for elType:char** goType:char**
	Element attrs_req has not registered type char** 
	Element attrs_req not recognized getter for elType:char** goType:char**

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
	   Function: htmlIsScriptAttribute
	   ReturnType: int
	   Args: (('name', ['xmlChar', '*'], None),)
*/
func HtmlIsScriptAttribute(name string) int {
	c_name:= (*C.xmlChar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(c_name))

	c_ret := C.htmlIsScriptAttribute(c_name)



	return int(c_ret)
}
/* 
	   Function: htmlHandleOmittedElem
	   ReturnType: int
	   Args: (('val', ['int'], None),)
*/
func HtmlHandleOmittedElem(val int) int {
	c_val := C.int(val)

	c_ret := C.htmlHandleOmittedElem(c_val)



	return int(c_ret)
}
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
	   Function: htmlParseFile
	   ReturnType: htmlDocPtr
	   Args: (('filename', ['char', '*'], None), ('encoding', ['char', '*'], None))
*/
func HtmlParseFile(filename string,encoding string) *XmlDoc {
	c_filename:= (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(c_filename))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))

	c_ret := C.htmlParseFile(c_filename,c_encoding)



	if c_ret == nil {return nil}
	return &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
}
/* 
	   Function: htmlReadFd
	   ReturnType: htmlDocPtr
	   Args: (('fd', ['int'], None), ('URL', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
func HtmlReadFd(fd int,URL string,encoding string,options int) *XmlDoc {
	c_fd := C.int(fd)
	c_URL:= (*C.char)(unsafe.Pointer(C.CString(URL)))
	defer C.free(unsafe.Pointer(c_URL))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	c_options := C.int(options)

	c_ret := C.htmlReadFd(c_fd,c_URL,c_encoding,c_options)



	if c_ret == nil {return nil}
	return &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
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
	   Function: htmlTagLookup
	   ReturnType: htmlElemDesc*
	   Args: (('tag', ['xmlChar', '*'], None),)
*/
func HtmlTagLookup(tag string) *HtmlElemDesc {
	c_tag:= (*C.xmlChar)(unsafe.Pointer(C.CString(tag)))
	defer C.free(unsafe.Pointer(c_tag))

	c_ret := C.htmlTagLookup(c_tag)



	if c_ret == nil {return nil}
	return &HtmlElemDesc{handler:(C.htmlElemDescPtr)(c_ret)}
}
/* 
	   Function: htmlCreateMemoryParserCtxt
	   ReturnType: htmlParserCtxtPtr
	   Args: (('buffer', ['char', '*'], None), ('size', ['int'], None))
*/
func HtmlCreateMemoryParserCtxt(buffer string) (g_ret *XmlParserCtxt,err error) {
	c_buffer:= (*C.char)(unsafe.Pointer(C.CString(buffer)))
	defer C.free(unsafe.Pointer(c_buffer))
	c_size:=C.int(len(buffer)*1)
	c_ret := C.htmlCreateMemoryParserCtxt(c_buffer,c_size)

	if c_ret == nil {
		err = fmt.Errorf("htmlCreateMemoryParserCtxt errno %d" ,c_ret)
	} else {
		g_ret =  &XmlParserCtxt{handler:(C.xmlParserCtxtPtr)(c_ret)}
	}
	return
}
/* 
	   Function: htmlCtxtReset
	   ReturnType: void
	   Args: (('ctxt', ['htmlParserCtxtPtr'], None),)
*/
func HtmlCtxtReset(ctxt *XmlParserCtxt) {
	var c_ctxt C.htmlParserCtxtPtr=nil
	if ctxt !=nil { c_ctxt = (C.htmlParserCtxtPtr)(ctxt.handler) }

	C.htmlCtxtReset(c_ctxt)




}
/* 
	   Function: htmlElementAllowedHere
	   ReturnType: int
	   Args: ((None, ['htmlElemDesc', '*'], None), (None, ['xmlChar', '*'], None))
*/
func HtmlElementAllowedHere(arg1 *HtmlElemDesc,arg2 string) int {
	var c_arg1 C.htmlElemDescPtr=nil
	if arg1 !=nil { c_arg1 = (C.htmlElemDescPtr)(arg1.handler) }
	c_arg2:= (*C.xmlChar)(unsafe.Pointer(C.CString(arg2)))
	defer C.free(unsafe.Pointer(c_arg2))

	c_ret := C.htmlElementAllowedHere(c_arg1,c_arg2)



	return int(c_ret)
}
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
	c_size:=C.int(len(chunk)*1)
	c_ret := C.htmlCreatePushParserCtxt(c_sax,nil,c_chunk,c_size,c_filename,c_enc)

	if c_ret == nil {
		err = fmt.Errorf("htmlCreatePushParserCtxt errno %d" ,c_ret)
	} else {
		g_ret =  &XmlParserCtxt{handler:(C.xmlParserCtxtPtr)(c_ret)}
	}
	return
}
/* 
	   Function: htmlReadMemory
	   ReturnType: htmlDocPtr
	   Args: (('buffer', ['char', '*'], None), ('size', ['int'], None), ('URL', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
func HtmlReadMemory(buffer string,size int,URL string,encoding string,options int) *XmlDoc {
	c_buffer:= (*C.char)(unsafe.Pointer(C.CString(buffer)))
	defer C.free(unsafe.Pointer(c_buffer))
	c_size := C.int(size)
	c_URL:= (*C.char)(unsafe.Pointer(C.CString(URL)))
	defer C.free(unsafe.Pointer(c_URL))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	c_options := C.int(options)

	c_ret := C.htmlReadMemory(c_buffer,c_size,c_URL,c_encoding,c_options)



	if c_ret == nil {return nil}
	return &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
}
/* 
	   Function: htmlIsAutoClosed
	   ReturnType: int
	   Args: (('doc', ['htmlDocPtr'], None), ('elem', ['htmlNodePtr'], None))
*/
func HtmlIsAutoClosed(doc *XmlDoc,elem *XmlNode) int {
	var c_doc C.htmlDocPtr=nil
	if doc !=nil { c_doc = (C.htmlDocPtr)(doc.handler) }
	var c_elem C.htmlNodePtr=nil
	if elem !=nil { c_elem = (C.htmlNodePtr)(elem.handler) }

	c_ret := C.htmlIsAutoClosed(c_doc,c_elem)



	return int(c_ret)
}
/* 
	   Function: htmlParseCharRef
	   ReturnType: int
	   Args: (('ctxt', ['htmlParserCtxtPtr'], None),)
*/
func HtmlParseCharRef(ctxt *XmlParserCtxt) int {
	var c_ctxt C.htmlParserCtxtPtr=nil
	if ctxt !=nil { c_ctxt = (C.htmlParserCtxtPtr)(ctxt.handler) }

	c_ret := C.htmlParseCharRef(c_ctxt)



	return int(c_ret)
}
/* 
	   Function: htmlReadDoc
	   ReturnType: htmlDocPtr
	   Args: (('cur', ['xmlChar', '*'], None), ('URL', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
func HtmlReadDoc(cur string,URL string,encoding string,options int) *XmlDoc {
	c_cur:= (*C.xmlChar)(unsafe.Pointer(C.CString(cur)))
	defer C.free(unsafe.Pointer(c_cur))
	c_URL:= (*C.char)(unsafe.Pointer(C.CString(URL)))
	defer C.free(unsafe.Pointer(c_URL))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	c_options := C.int(options)

	c_ret := C.htmlReadDoc(c_cur,c_URL,c_encoding,c_options)



	if c_ret == nil {return nil}
	return &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
}
/* 
	   Function: htmlParseDocument
	   ReturnType: int
	   Args: (('ctxt', ['htmlParserCtxtPtr'], None),)
*/
func HtmlParseDocument(ctxt *XmlParserCtxt) int {
	var c_ctxt C.htmlParserCtxtPtr=nil
	if ctxt !=nil { c_ctxt = (C.htmlParserCtxtPtr)(ctxt.handler) }

	c_ret := C.htmlParseDocument(c_ctxt)



	return int(c_ret)
}
/* 
	   Function: htmlNodeStatus
	   ReturnType: htmlStatus
	   Args: ((None, ['htmlNodePtr'], None), (None, ['int'], None))
*/
func HtmlNodeStatus(arg1 *XmlNode,arg2 int) int {
	var c_arg1 C.htmlNodePtr=nil
	if arg1 !=nil { c_arg1 = (C.htmlNodePtr)(arg1.handler) }
	c_arg2 := C.int(arg2)

	c_ret := C.htmlNodeStatus(c_arg1,c_arg2)



	return int(c_ret)
}
/* 
	   Function: htmlParseChunk
	   ReturnType: int
	   Args: (('ctxt', ['htmlParserCtxtPtr'], None), ('chunk', ['char', '*'], None), ('size', ['int'], None), ('terminate', ['int'], None))
*/
func HtmlParseChunk(ctxt *XmlParserCtxt,chunk string,size int,terminate int) int {
	var c_ctxt C.htmlParserCtxtPtr=nil
	if ctxt !=nil { c_ctxt = (C.htmlParserCtxtPtr)(ctxt.handler) }
	c_chunk:= (*C.char)(unsafe.Pointer(C.CString(chunk)))
	defer C.free(unsafe.Pointer(c_chunk))
	c_size := C.int(size)
	c_terminate := C.int(terminate)

	c_ret := C.htmlParseChunk(c_ctxt,c_chunk,c_size,c_terminate)



	return int(c_ret)
}
/* 
	   Function: htmlSAXParseFile
	   ReturnType: htmlDocPtr
	   Args: (('filename', ['char', '*'], None), ('encoding', ['char', '*'], None), ('sax', ['htmlSAXHandlerPtr'], None), ('userData', ['void', '*'], None))
*/
func HtmlSAXParseFile(filename string,encoding string,sax *XmlSAXHandler) *XmlDoc {
	c_filename:= (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(c_filename))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	var c_sax C.htmlSAXHandlerPtr=nil
	if sax !=nil { c_sax = (C.htmlSAXHandlerPtr)(sax.handler) }

	c_ret := C.htmlSAXParseFile(c_filename,c_encoding,c_sax,nil)



	if c_ret == nil {return nil}
	return &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
}
/* 
	   Function: htmlElementStatusHere
	   ReturnType: htmlStatus
	   Args: ((None, ['htmlElemDesc', '*'], None), (None, ['htmlElemDesc', '*'], None))
*/
func HtmlElementStatusHere(arg1 *HtmlElemDesc,arg2 *HtmlElemDesc) int {
	var c_arg1 C.htmlElemDescPtr=nil
	if arg1 !=nil { c_arg1 = (C.htmlElemDescPtr)(arg1.handler) }
	var c_arg2 C.htmlElemDescPtr=nil
	if arg2 !=nil { c_arg2 = (C.htmlElemDescPtr)(arg2.handler) }

	c_ret := C.htmlElementStatusHere(c_arg1,c_arg2)



	return int(c_ret)
}
/* 
	   Function: htmlEntityValueLookup
	   ReturnType: htmlEntityDesc*
	   Args: (('value', ['unsigned int'], None),)
*/
func HtmlEntityValueLookup(value uint) *HtmlEntityDesc {
	c_value := C.uint(value)

	c_ret := C.htmlEntityValueLookup(c_value)



	if c_ret == nil {return nil}
	return &HtmlEntityDesc{handler:(C.htmlEntityDescPtr)(c_ret)}
}
/* 
	   Function: htmlParseElement
	   ReturnType: void
	   Args: (('ctxt', ['htmlParserCtxtPtr'], None),)
*/
func HtmlParseElement(ctxt *XmlParserCtxt) {
	var c_ctxt C.htmlParserCtxtPtr=nil
	if ctxt !=nil { c_ctxt = (C.htmlParserCtxtPtr)(ctxt.handler) }

	C.htmlParseElement(c_ctxt)




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
	c0_outlen:=C.int(len(in)*3)
	c_outlen:=&c0_outlen
	c0_inlen:=C.int(len(in)*1)
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
	   Function: htmlEntityLookup
	   ReturnType: htmlEntityDesc*
	   Args: (('name', ['xmlChar', '*'], None),)
*/
func HtmlEntityLookup(name string) *HtmlEntityDesc {
	c_name:= (*C.xmlChar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(c_name))

	c_ret := C.htmlEntityLookup(c_name)



	if c_ret == nil {return nil}
	return &HtmlEntityDesc{handler:(C.htmlEntityDescPtr)(c_ret)}
}
/* 
	   Function: htmlFreeParserCtxt
	   ReturnType: void
	   Args: (('ctxt', ['htmlParserCtxtPtr'], None),)
*/
func HtmlFreeParserCtxt(ctxt *XmlParserCtxt) {
	var c_ctxt C.htmlParserCtxtPtr=nil
	if ctxt !=nil { c_ctxt = (C.htmlParserCtxtPtr)(ctxt.handler) }

	C.htmlFreeParserCtxt(c_ctxt)




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
	c_size:=C.int(len(buffer)*1)
	c_ret := C.htmlCtxtReadMemory(c_ctxt,c_buffer,c_size,c_URL,c_encoding,c_options)

	if c_ret == nil {
		err = fmt.Errorf("htmlCtxtReadMemory errno %d" ,c_ret)
	} else {
		g_ret =  &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
	}
	return
}
/* 
	   Function: htmlParseDoc
	   ReturnType: htmlDocPtr
	   Args: (('cur', ['xmlChar', '*'], None), ('encoding', ['char', '*'], None))
*/
func HtmlParseDoc(cur string,encoding string) *XmlDoc {
	c_cur:= (*C.xmlChar)(unsafe.Pointer(C.CString(cur)))
	defer C.free(unsafe.Pointer(c_cur))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))

	c_ret := C.htmlParseDoc(c_cur,c_encoding)



	if c_ret == nil {return nil}
	return &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
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
	   Function: htmlEncodeEntities
	   ReturnType: int
	   Args: (('out', ['unsigned char', '*'], None), ('outlen', ['int', '*'], None), ('in', ['unsigned char', '*'], None), ('inlen', ['int', '*'], None), ('quoteChar', ['int'], None))
*/
func HtmlEncodeEntities(in string,quoteChar int) (g_out string,err error) {
	c_in:= (*C.uchar)(unsafe.Pointer(C.CString(in)))
	defer C.free(unsafe.Pointer(c_in))
	c_quoteChar := C.int(quoteChar)
		c_out:= (*C.uchar)(C.calloc(  (C.size_t)( len(in)*3+ 1 )  ,1))
		defer C.free(unsafe.Pointer(c_out))
	c0_outlen:=C.int(len(in)*1)
	c_outlen:=&c0_outlen
	c0_inlen:=C.int(len(in)*1)
	c_inlen:=&c0_inlen
	c_ret := C.htmlEncodeEntities(c_out,c_outlen,c_in,c_inlen,c_quoteChar)

	if c_ret != 0 {
		err = fmt.Errorf("htmlEncodeEntities errno %d" ,c_ret)
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
	   Function: htmlNewParserCtxt
	   ReturnType: htmlParserCtxtPtr
	   Args: ((None, ['void'], None),)
*/
func HtmlNewParserCtxt() *XmlParserCtxt {


	c_ret := C.htmlNewParserCtxt()



	if c_ret == nil {return nil}
	return &XmlParserCtxt{handler:(C.xmlParserCtxtPtr)(c_ret)}
}
/* 
	   Function: htmlSAXParseDoc
	   ReturnType: htmlDocPtr
	   Args: (('cur', ['xmlChar', '*'], None), ('encoding', ['char', '*'], None), ('sax', ['htmlSAXHandlerPtr'], None), ('userData', ['void', '*'], None))
*/
/*

	Warn: userData void* Not defined
	Warn: userData void* No converter to C(go2cConverter)

func HtmlSAXParseDoc(cur string,encoding string,sax *XmlSAXHandler,userData void*) *XmlDoc {
	c_cur:= (*C.xmlChar)(unsafe.Pointer(C.CString(cur)))
	defer C.free(unsafe.Pointer(c_cur))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	var c_sax C.htmlSAXHandlerPtr=nil
	if sax !=nil { c_sax = (C.htmlSAXHandlerPtr)(sax.handler) }
	userData

	c_ret := C.htmlSAXParseDoc(c_cur,c_encoding,c_sax,c_userData)



	if c_ret == nil {return nil}
	return &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
}

*/
/* 
	   Function: htmlCtxtUseOptions
	   ReturnType: int
	   Args: (('ctxt', ['htmlParserCtxtPtr'], None), ('options', ['int'], None))
*/
func HtmlCtxtUseOptions(ctxt *XmlParserCtxt,options int) int {
	var c_ctxt C.htmlParserCtxtPtr=nil
	if ctxt !=nil { c_ctxt = (C.htmlParserCtxtPtr)(ctxt.handler) }
	c_options := C.int(options)

	c_ret := C.htmlCtxtUseOptions(c_ctxt,c_options)



	return int(c_ret)
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



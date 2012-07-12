package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/xmlreader.h>
*/
import "C"
import "unsafe"
import "fmt"




/* 
	   Function: xmlTextReaderHasValue
	   ReturnType: int
	   Args: (('reader', ['xmlTextReaderPtr'], None),)
*/
func XmlTextReaderHasValue(reader *XmlTextReader) int {
	var c_reader C.xmlTextReaderPtr=nil
	if reader !=nil { c_reader = (C.xmlTextReaderPtr)(reader.handler) }

	c_ret := C.xmlTextReaderHasValue(c_reader)



	return int(c_ret)
}
/* 
	   Function: xmlTextReaderDepth
	   ReturnType: int
	   Args: (('reader', ['xmlTextReaderPtr'], None),)
*/
func XmlTextReaderDepth(reader *XmlTextReader) int {
	var c_reader C.xmlTextReaderPtr=nil
	if reader !=nil { c_reader = (C.xmlTextReaderPtr)(reader.handler) }

	c_ret := C.xmlTextReaderDepth(c_reader)



	return int(c_ret)
}
/* 
	   Function: xmlReaderForFile
	   ReturnType: xmlTextReaderPtr
	   Args: (('filename', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
func XmlReaderForFile(filename string,encoding string,options int) (g_ret *XmlTextReader,err error) {
	c_filename:= (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(c_filename))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	c_options := C.int(options)

	c_ret := C.xmlReaderForFile(c_filename,c_encoding,c_options)

	if c_ret == nil {
		err = fmt.Errorf("xmlReaderForFile errno %d" ,c_ret)
	} else {
		g_ret =  &XmlTextReader{handler:(C.xmlTextReaderPtr)(c_ret)}
	}
	return
}
/* 
	   Function: xmlTextReaderNodeType
	   ReturnType: int
	   Args: (('reader', ['xmlTextReaderPtr'], None),)
*/
func XmlTextReaderNodeType(reader *XmlTextReader) int {
	var c_reader C.xmlTextReaderPtr=nil
	if reader !=nil { c_reader = (C.xmlTextReaderPtr)(reader.handler) }

	c_ret := C.xmlTextReaderNodeType(c_reader)



	return int(c_ret)
}
/* 
	   Function: xmlFreeTextReader
	   ReturnType: void
	   Args: (('reader', ['xmlTextReaderPtr'], None),)
*/
func XmlFreeTextReader(reader *XmlTextReader) {
	var c_reader C.xmlTextReaderPtr=nil
	if reader !=nil { c_reader = (C.xmlTextReaderPtr)(reader.handler) }

	C.xmlFreeTextReader(c_reader)




}
/* 
	   Function: xmlTextReaderRead
	   ReturnType: int
	   Args: (('reader', ['xmlTextReaderPtr'], None),)
*/
func XmlTextReaderRead(reader *XmlTextReader) int {
	var c_reader C.xmlTextReaderPtr=nil
	if reader !=nil { c_reader = (C.xmlTextReaderPtr)(reader.handler) }

	c_ret := C.xmlTextReaderRead(c_reader)



	return int(c_ret)
}
/* 
	   Function: xmlTextReaderIsEmptyElement
	   ReturnType: int
	   Args: (('reader', ['xmlTextReaderPtr'], None),)
*/
func XmlTextReaderIsEmptyElement(reader *XmlTextReader) int {
	var c_reader C.xmlTextReaderPtr=nil
	if reader !=nil { c_reader = (C.xmlTextReaderPtr)(reader.handler) }

	c_ret := C.xmlTextReaderIsEmptyElement(c_reader)



	return int(c_ret)
}
/* 
	   Function: xmlTextReaderConstValue
	   ReturnType: xmlChar*
	   Args: (('reader', ['xmlTextReaderPtr'], None),)
*/
func XmlTextReaderConstValue(reader *XmlTextReader) string {
	var c_reader C.xmlTextReaderPtr=nil
	if reader !=nil { c_reader = (C.xmlTextReaderPtr)(reader.handler) }

	c_ret := C.xmlTextReaderConstValue(c_reader)



	if c_ret == nil {return ""}
	g_ret:=C.GoString((*C.char)(unsafe.Pointer(c_ret)))
	return g_ret
}
/* 
	   Function: xmlTextReaderConstName
	   ReturnType: xmlChar*
	   Args: (('reader', ['xmlTextReaderPtr'], None),)
*/
func XmlTextReaderConstName(reader *XmlTextReader) string {
	var c_reader C.xmlTextReaderPtr=nil
	if reader !=nil { c_reader = (C.xmlTextReaderPtr)(reader.handler) }

	c_ret := C.xmlTextReaderConstName(c_reader)



	if c_ret == nil {return ""}
	g_ret:=C.GoString((*C.char)(unsafe.Pointer(c_ret)))
	return g_ret
}



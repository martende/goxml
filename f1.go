package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/parser.h>
#include <libxml/tree.h>
#define SOSKA 1
*/
import "C"

type xmlCore struct {

}

type XmlNode struct {
	xmlCore
}

type XmlDoc struct {
	XmlNode
}

func test1() int {
	return 1
}


func parseFile(s string) (doc *XmlDoc, err error) {
	/*
	var docPtr C.xmlDocPtr
	var filename *C.char
    var encoding *C.char
    var options int 
    options = C.XML_PARSE_DTDATTR
    filename = C.CString(filename)
    encoding = C.CString(filename)
    
	docPtr = C.xmlReadFile(filename)
    //if ret is None:raise parserError('xmlParseFile() failed')
    */
    return &XmlDoc{},nil

}




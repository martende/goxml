package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/tree.h>
#include <libxml/xmlreader.h>
*/
import "C"

type XmlNode struct {
	handler C.xmlNodePtr
}


type XmlTextReader struct {
	handler C.xmlTextReaderPtr
}


type XmlDict struct {
	handler C.xmlDictPtr
}


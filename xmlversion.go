package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/xmlversion.h>
*/
import "C"


const LIBXML_DOTTED_VERSION = C.LIBXML_DOTTED_VERSION
const LIBXML_VERSION = C.LIBXML_VERSION
const LIBXML_VERSION_STRING =C.LIBXML_VERSION_STRING
const LIBXML_VERSION_EXTRA = C.LIBXML_VERSION_EXTRA

func XmlCheckVersion() {
        C.xmlCheckVersion(LIBXML_VERSION)
}


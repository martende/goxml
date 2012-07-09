package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/HTMLparser.h>
*/
import "C"


const HTML_DEPRECATED=C.HTML_DEPRECATED
const HTML_INVALID=C.HTML_INVALID
const HTML_NA=C.HTML_NA
const HTML_PARSE_COMPACT=C.HTML_PARSE_COMPACT
const HTML_PARSE_NOBLANKS=C.HTML_PARSE_NOBLANKS
const HTML_PARSE_NODEFDTD=C.HTML_PARSE_NODEFDTD
const HTML_PARSE_NOERROR=C.HTML_PARSE_NOERROR
const HTML_PARSE_NOIMPLIED=C.HTML_PARSE_NOIMPLIED
const HTML_PARSE_NONET=C.HTML_PARSE_NONET
const HTML_PARSE_NOWARNING=C.HTML_PARSE_NOWARNING
const HTML_PARSE_PEDANTIC=C.HTML_PARSE_PEDANTIC
const HTML_PARSE_RECOVER=C.HTML_PARSE_RECOVER
const HTML_REQUIRED=C.HTML_REQUIRED
const HTML_VALID=C.HTML_VALID
func UTF8ToHtml(out string,outlen int,in string,inlen int) int {
	var c_ret C.int
	var g_ret int
	c_out := 
	c_outlen := 
	c_in := 
	c_inlen := 
	c_ret = C.UTF8ToHtml(c_out,c_outlen,c_in,c_inlen)
	g_ret = int(c_ret)
	return g_ret
}

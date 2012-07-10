package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/HTMLparser.h>
*/
import "C"
import "unsafe"



/* 
	   Function: UTF8ToHtml
	   ReturnType: int
	   Args: (('out', ['unsigned char', '*'], None), ('outlen', ['int', '*'], None), ('in', ['unsigned char', '*'], None), ('inlen', ['int', '*'], None))
*/
func UTF8ToHtml(out *string,outlen *int,in string) int {
	c_out:= (*C.uchar)(unsafe.Pointer((C.CString(*out))))
	c0_outlen:=C.int(*outlen)
	c_outlen:=&c0_outlen
	c_in:= (*C.uchar)(unsafe.Pointer(C.CString(in)))
	c0_inlen:=C.int(len(in)+1)
	c_inlen:=&c0_inlen
	c_ret := C.UTF8ToHtml2(c_out,c_outlen,c_in,c_inlen)
	*out = C.GoString((*C.char)(unsafe.Pointer(c_out)))
	*outlen = int(c0_outlen)
	return int(c_ret)
}



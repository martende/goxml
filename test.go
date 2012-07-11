package goxml
/*
#cgo pkg-config: libxml-2.0
#include "test.h"
#include <libxml/parser.h>
*/
import "C"
import "unsafe"
import "fmt"

//export go_callback_xmlInputReadCallback
func go_callback_xmlInputReadCallback(pfoo unsafe.Pointer, p1 C.int) {
	foo := *(*func(C.int))(pfoo)
    foo(p1)
}

type Udata interface {
	GetData() string
	SetData(t string)
}

func SetUserdata(s Udata) {
	fmt.Printf("SetUserdata %v\n",s)
	C.set_userdata( unsafe.Pointer(&s) )
}

func GetUserdata() Udata {
	var t unsafe.Pointer
	t = C.get_userdata()
	fmt.Printf("GetUserdata %v\n",t)
	k:=(*Udata)(t)
	m:=*k
	fmt.Printf("GetUserdata %v\n",m)
	return m
	//C.set_userdata( unsafe.Pointer(&s) )
}

func Callbacker(t1 func(a1,a2 int) int,t2 func(a1,a2 int) int,a1 int) int{
	tcb1:= (C.callbackInfoPtr)(C.calloc(C.CBSIZE,1))
	
	
	ctxt := C.xmlNewParserCtxt();
	C.xmlCtxtReadIO(ctxt,(C.xmlInputReadCallback)(unsafe.Pointer(&tcb1)),(C.xmlInputCloseCallback)(unsafe.Pointer(&t2)),nil,C.CString("test.xml"),nil,0)
	return 1
	//return t(a1,1)
}


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

func Struct_expr() {
	k:=C.createStructArray()
	off:=(uintptr)(unsafe.Pointer(k.data))
	//off2:=unsafe.Pointer((uintptr)(off)+8)
	t:=*(**C.struc1)(unsafe.Pointer(off))
	fmt.Printf("off=%p %p\n",t,t.a1)
	off+=unsafe.Sizeof(t)
	t=*(**C.struc1)(unsafe.Pointer(off))
	fmt.Printf("off=%p %p\n",t,t.a1)
	//t2:=*(**C.struc1)(off2)
	//fmt.Printf("%p\n",t2.a1)
	
	//fmt.Printf("%p\n",t[0])
	/*m:=t[0]
	fmt.Printf("%d\n",m.a1)
	*/
}

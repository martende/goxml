package main

import (
	"github.com/martende/goxml"
	//"os"
	"fmt"
	"runtime"
)

type M struct {
	s string;
}
func (this *M) GetData() string {
	return this.s
}
func (this *M) SetData(t string) {
	this.s = t
}

func koko(a *int) {
	*a = 5;
}

func t2(ss string) {
	k:=M{s:ss}
	goxml.SetUserdata(&k)
	runtime.GC()
}

func t(ss string) {
	k:=M{s:ss}
	goxml.SetUserdata(&k)
	t2(ss+"KO")
	runtime.GC()
}
func TT(a1,a2 int) int {
	fmt.Printf("TT Called %d %d\n",a1,a2)
	return a1+a2
}
func main() {
	
	fmt.Printf("CBRET %d\n",goxml.Callbacker(TT,TT,2))
	
	runtime.GC()
	fmt.Printf("")
	//fmt.Printf("K=%v\n",goxml.GetUserdata().GetData())
}

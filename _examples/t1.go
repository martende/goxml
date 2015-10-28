package main

import (
	. "github.com/martende/goxml"
	//"os"
	"fmt"
)

type M struct {
	s string;
}
func koko(a *int) {
	*a = 5;
}
func main() {
	fmt.Printf("121\n")
	k,_:=UTF8ToHtml("t1")
	fmt.Printf("i=%s\n",k)
        k,_=UTF8ToHtml("t1ß")
        fmt.Printf("i=%s\n",k)
        k,_=UTF8ToHtml("")
        fmt.Printf("i=%s\n",k)
}

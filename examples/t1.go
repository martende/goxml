package main

import (
	. "goxml"
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
	s := ""
	i := 12
	UTF8ToHtml(&s,&i,"s")
}

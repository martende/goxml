package main

import (
	//. "goxml"
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
	//s := M{"soski"}
	var a int;
	koko(&a)
	fmt.Printf("a = %d\n",a)
}

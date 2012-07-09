package main

import (
	. "goxml"
	//"os"
	"fmt"
)

type M struct {
	s string;
}

func main() {
	fmt.Printf("121\n")
	s := M{"soski"}
	SetUserdata(s)
}

package main

import (
	. "github.com/martende/goxml"
	"os"
	"fmt"
)
/*

Analog http://www.xmlsoft.org/examples/parse3.c

Usage:

./parse3

*/

var document = "<doc/>";

func example3Func(content string) {
    doc,err := XmlReadMemory(content, "noname.xml", "", 0)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to parse document\n")
    }
    XmlFreeDoc(doc);
}

func main() {
	XmlCheckVersion()
    example3Func(document)
	XmlCleanupParser()
	XmlMemoryDump()
}

package main

import (
	. "goxml"
	"os"
	"fmt"
)
/*

Analog http://www.xmlsoft.org/examples/parse3.c

Usage:

./parse3

*/

var document = "<doc/>";

func example3Func(content string,length int) {
    doc := XmlReadMemory(content, length, "noname.xml", "", 0)
    if doc == nil {
        fmt.Fprintf(os.Stderr, "Failed to parse document\n")
    }
	
    XmlFreeDoc(doc);
}

func main() {
	XmlCheckVersion()
	
    example3Func(document, 6)
    
	XmlCleanupParser()
	XmlMemoryDump()
}

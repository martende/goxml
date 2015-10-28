package main

import (
	. "github.com/martende/goxml"
	"os"
	"fmt"
)
/*

Analog http://www.xmlsoft.org/examples/parse1.c

Usage:

./parse1 test1.xml

*/
func example1Func(filename string) {
	doc,err:=XmlReadFile(filename, "", 0)
	if err != nil {
        fmt.Fprint(os.Stderr,"Failed to parse %s\n", filename)
    }
    XmlFreeDoc(doc);
}

func main() {
	XmlCheckVersion()
	
	if len(os.Args) != 2 {
		fmt.Println("usage: ./example filename")
        return
    }
    
    example1Func(os.Args[1])
    
	XmlCleanupParser()
	XmlMemoryDump()
}

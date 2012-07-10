package main

import (
	. "goxml"
	"os"
	"fmt"
)
/*

Analog http://www.xmlsoft.org/examples/parse1.c

Usage:

./parse1 test1.xml

*/
func example1Func(filename string) string {
	doc:=XmlReadFile(filename, "", 0)
	if doc == nil {
        fmt.Fprint(os.Stderr,"Failed to parse %s\n", filename);
	return string(nil);
    }
    XmlFreeDoc(doc);
    return "123";
}

func main() {
	XmlCheckVersion()
	
	if len(os.Args) != 2 {
        return
    }
    
    example1Func(os.Args[1])
    
	XmlCleanupParser()
	XmlMemoryDump()
}

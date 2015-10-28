package main

import (
	. "github.com/martende/goxml"
	"os"
	"fmt"
)

/*

Analog http://www.xmlsoft.org/examples/reader3.c

Usage:

./reader3 test3.xml

*/

func extractFile(filename,pattern string)  (doc *XmlDoc,err error) {

    reader,err:=XmlReaderForFile(filename,"",0) 

    if ( err == nil) {
    	defer XmlFreeTextReader(reader)
    	
    	XmlTextReaderPreservePattern(reader,pattern)
    	
    	ret := XmlTextReaderRead(reader)
    	
        for ;ret == 1;ret = XmlTextReaderRead(reader) {
            //processNode(reader)
        }
        if ret != 0 {
        	err=fmt.Errorf("%s : failed to parse", filename)
        } else {
        	doc = XmlTextReaderCurrentDoc(reader)
        }
    } else {
    	fmt.Fprintf(os.Stderr,"Unable to open %s\n", filename)
    }
    return
}

func main() {
	filename:="test3.xml"
	pattern:="preserved"
	
	XmlCheckVersion()
	if len(os.Args) == 3 {
		filename = os.Args[1]
		pattern = os.Args[2]
    }
    doc,err:= extractFile(filename,pattern)
    if err==nil {
    	XmlDocDump(os.Stdout, doc)
    	XmlFreeDoc(doc)
    }
	XmlCleanupParser()
	XmlMemoryDump()
}

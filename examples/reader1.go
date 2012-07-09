package main

import (
	. "goxml"
	"os"
	"fmt"
)
/*

Analog http://www.xmlsoft.org/examples/reader1.c

Usage:

./reader1 test1.xml

*/


func processNode(reader *XmlTextReader) {

    name := XmlTextReaderConstName(reader)
    if name == "" {
    	name = "--"
    }

	value := XmlTextReaderConstValue(reader)

    fmt.Printf("%d %d %s %d %d", 
	    XmlTextReaderDepth(reader),
	    XmlTextReaderNodeType(reader),
	    name,
	    XmlTextReaderIsEmptyElement(reader),
	    XmlTextReaderHasValue(reader));
    if (value == "") {
    	fmt.Printf("\n");
    } else {
        if (len(value) > 40) {
            fmt.Printf(" %.40s...\n", value);
        } else {
        	fmt.Printf(" %s\n", value);
        }
    }
}


func streamFile(filename string) {
    reader:=XmlReaderForFile(filename,"",0)
    if ( reader != nil) {
    	ret := XmlTextReaderRead(reader)
        for ;ret == 1;ret = XmlTextReaderRead(reader) {
            processNode(reader)
        }
        XmlFreeTextReader(reader);
        if ret != 0 {
            fmt.Fprintf(os.Stderr, "%s : failed to parse\n", filename);
        }
    } else {
    	fmt.Fprintf(os.Stderr,"Unable to open %s\n", filename)
    }
}

func main() {
	XmlCheckVersion()
	
	if len(os.Args) != 2 {
        return
    }
    
    streamFile(os.Args[1])
    
	XmlCleanupParser()
	XmlMemoryDump()
}

package main

import (
	. "goxml"
	"os"
	"fmt"
	"io"
)
/*

Analog http://www.xmlsoft.org/examples/reader1.c

Usage:

./reader1 test1.xml

*/

func processNode(XmlTextReader *reader) {

    name := XmlTextReaderConstName(reader)
    if (name == nil) {
    	name = "--"
    }

	value: = XmlTextReaderConstValue(reader);

    Fmt.Printf("%d %d %s %d %d", 
	    XmlTextReaderDepth(reader),
	    XmlTextReaderNodeType(reader),
	    name,
	    XmlTextReaderIsEmptyElement(reader),
	    XmlTextReaderHasValue(reader));
    if (value == nil) {
    	Fmt.Printf("\n");
    } else {
        if (XmlStrlen(value) > 40) {
            Fmt.Printf(" %.40s...\n", value);
        } else {
        	Fmt.Printf(" %s\n", value);
        }
    }
}


func streamFile(filename string) {
    reader:=XmlReaderForFile(filename,"",0)
    if ( reader != nil) {
        for (ret := XmlTextReaderRead(reader);ret == 1;ret = XmlTextReaderRead(reader);) {
            processNode(reader);
        }
        XmlFreeTextReader(reader);
        if (ret != 0) {
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
    var desc *os.File
    var err error

    streamFile(os.Args[1])
    
	XmlCleanupParser()
	XmlMemoryDump()
}

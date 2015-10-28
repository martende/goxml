package main

import (
	. "github.com/martende/goxml"
	"os"
	"fmt"
)

/*

Analog http://www.xmlsoft.org/examples/reader2.c

Usage:

./reader2 test1.xml

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
	/* XML_PARSE_DTDATTR default DTD attributes */
	/* XML_PARSE_NOENT substitute entities */
	/* XML_PARSE_DTDVALID validate with the DTD */
    reader,err:=XmlReaderForFile(filename,"",XML_PARSE_DTDATTR | XML_PARSE_NOENT | XML_PARSE_DTDVALID) 

    if ( err == nil) {
    	ret := XmlTextReaderRead(reader)
        for ;ret == 1;ret = XmlTextReaderRead(reader) {
            processNode(reader)
        }
        /*
        	* Once the document has been fully parsed check the validation results
	 	*/
	 	if (XmlTextReaderIsValid(reader) != 1) {
		    fmt.Fprintf(os.Stderr, "Document %s does not validate\n", filename);
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

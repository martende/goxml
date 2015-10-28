package main

import (
	. "github.com/martende/goxml"
	"os"
	"fmt"
)
/*

Analog http://www.xmlsoft.org/examples/parse2.c

Usage:

./parse2 test1.xml

*/
func example2Func(filename string) {
	ctxt := XmlNewParserCtxt()
	if ctxt == nil {
		fmt.Fprintf(os.Stderr,"Failed to allocate parser context\n");
        return;
	}
	doc,err := XmlCtxtReadFile(ctxt, filename, "", XML_PARSE_DTDVALID);

    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to parse %s\n", filename);
    } else {
    	/* check if validation suceeded */
        if ctxt.GetValid() == 0 {
        	fmt.Fprintf(os.Stderr,"Failed to validate %s\n", filename)
	    }
	    /* free up the resulting document */
	    XmlFreeDoc(doc);
    }
    /* free up the parser context */
    XmlFreeParserCtxt(ctxt);
}

func main() {
	XmlCheckVersion()
	
	if len(os.Args) != 2 {
        return
    }
    
    example2Func(os.Args[1])
    
	XmlCleanupParser()
	XmlMemoryDump()
}

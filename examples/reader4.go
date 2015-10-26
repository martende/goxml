package main

import (
	. "goxml"
	"os"
	"fmt"
)

/*

Analog http://www.xmlsoft.org/examples/reader4.c

Usage:

./reader4 test2.xml test3.xml

*/

func processDoc(reader *XmlTextReader) {

    ret := XmlTextReaderRead(reader)
    	
    for ;ret == 1;ret = XmlTextReaderRead(reader) {
    	//processNode(reader)
    }
    doc := XmlTextReaderCurrentDoc(reader)
    if (doc == nil) {
    	fmt.Fprintf(os.Stderr,"failed to obtain document\n" )
    	return
    }
    URL:=doc.GetURL()
    
    if ret != 0 {
       	fmt.Fprintf(os.Stderr,"%s : failed to parse\n", URL)
       	return
    }
    
    fmt.Printf("%s: Processed ok\n", URL)
    
    return
}

func main() {
	if len(os.Args) < 2 {
		return
    }
    XmlCheckVersion()
    reader,err:=XmlReaderForFile(os.Args[1],"",0)
    if err!=nil {
    	fmt.Fprintf(os.Stderr,"%s: failed to create reader\n", os.Args[1])
    	return
    }
    defer XmlFreeTextReader(reader)
    processDoc(reader)
    for i:=2;i<len(os.Args);i++ {
    	if ret:=XmlReaderNewFile(reader,os.Args[i],"",0); ret == -1 {
    		fmt.Fprintf(os.Stderr,"%s: failed to recreate reader\n", os.Args[i])
    		return
    	}
    	processDoc(reader)    	
    }
    
	/*
     * Since we've called xmlTextReaderCurrentDoc, we now have to
     * clean up after ourselves.  We only have to do this the last
     * time, because xmlReaderNewFile calls xmlCtxtReset which takes
     * care of it.
    */
    doc := XmlTextReaderCurrentDoc(reader);
    if (doc != nil) {
      XmlFreeDoc(doc)
    }
	XmlCleanupParser()
	XmlMemoryDump()
}

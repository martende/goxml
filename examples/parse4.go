package main

import (
	. "goxml"
	"os"
	"fmt"
	"io"
)
/*

Analog http://www.xmlsoft.org/examples/parse4.c

Usage:

./parse4 test1.xml

*/


func example4Func(filename string,desc *os.File) {
    chars := make([]byte, 4)
    
    //XmlDoc doc; /* the resulting document tree */
    
    if cnt,err:=desc.Read(chars);cnt < len(chars) && err != nil {
    	fmt.Fprintf(os.Stderr,"Failed to parse %s\n", filename)
        return
    }
    
    ctxt,err := XmlCreatePushParserCtxt(nil,string(chars), filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create parser context !\n")
        return
    }

    for {
    	cnt,err:=desc.Read(chars);
    	if err!=nil {
	    	if err == io.EOF {
	    		XmlParseChunk(ctxt, string(chars), cnt, 0);
                break
            }
    		fmt.Fprintf(os.Stderr,"Failed to goparse %s %s\n", filename,err)
    		return
    	}
    	XmlParseChunk(ctxt, string(chars), cnt, 0);
    }
    doc := ctxt.GetMyDoc()
  	res := ctxt.GetWellFormed()  
    XmlFreeDoc(doc)
    
    if (res==0) {
        fmt.Fprintf(os.Stderr, "Failed to parse %s\n", filename);
    }
    
}

func main() {
	XmlCheckVersion()
	
	if len(os.Args) != 2 {
        return
    }
    var desc *os.File
    var err error
    
    if desc, err= os.Open(os.Args[1]); err != nil {
        fmt.Println("Failed to open: %s Error %s\n", os.Args[1] , err)
        return
    }

    defer desc.Close()
    
    example4Func(os.Args[1],desc)
    
	XmlCleanupParser()
	XmlMemoryDump()
}

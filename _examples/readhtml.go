package main

import (
	. "github.com/martende/goxml"
	"os"
	"fmt"
)

func GetHtml(doclocation string,encoding string) (doc *XmlDoc,err error) {    
    doc,err= HtmlReadFile(doclocation, encoding, HTML_PARSE_NOBLANKS | HTML_PARSE_NOERROR | HTML_PARSE_NOWARNING | HTML_PARSE_NONET)
    if err!=nil {
    	fmt.Printf( "Document Parsing Failed %s\n",err)
    }    	
    return
}

func GetRoot(doc *XmlDoc) {
    
    cur := XmlDocGetRootElement(doc)

    if (cur == nil) {
        fmt.Printf( "empty document\n")
    } else {
    	fmt.Printf("RootElement %s\n", cur.GetName())
    }

}


func main() {
	website:="http://xmlsoft.org/"
	encoding:="UTF-8"
	
	XmlCheckVersion()
	
	if len(os.Args) >= 2 {
		website = os.Args[1]
        
    }
    
    doc,err := GetHtml(website, encoding);
    if err != nil {
    	panic(err)
    }
    
    GetRoot(doc)
    XmlFreeDoc(doc)
	XmlCleanupParser()
	XmlMemoryDump()
}


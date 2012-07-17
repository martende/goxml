package main

import (
	. "goxml"
	"os"
	"fmt"
)

/*

Analog http://www.xmlsoft.org/examples/xpath1.c

Usage:
 
 ./xpath1 <xml-file> <xpath-expr> [<known-ns-list>]
 ./xpath1 test3.xml '//child2' > xpath1.tmp ; diff xpath1.tmp xpath1.res ; rm xpath1.tmp

*/

func execute_xpath_expression(filename string, xpathExpr string, nsList string) error {
    
	doc,err := XmlParseFile(filename)
	if err != nil {
        fmt.Fprintf(os.Stderr,"Failed to parse %s\n", filename)
        return fmt.Errorf("error")
    }
    defer XmlFreeDoc(doc)
    /* Create xpath evaluation context */
    xpathCtx,err := XmlXPathNewContext(doc)
    if err != nil {
        fmt.Fprintf(os.Stderr,"Error: unable to create new XPath context\n")
        return fmt.Errorf("error %d",1)
    }
    defer XmlXPathFreeContext(xpathCtx)
    
    xpathObj,err := XmlXPathEvalExpression(xpathExpr,xpathCtx)
    if err != nil {
        fmt.Fprintf(os.Stderr,"Error: unable to evaluate xpath expression  \"%s\"\n" , xpathExpr)
        return fmt.Errorf("error %d",1)
    }
    
    defer XmlXPathFreeObject(xpathObj)
    
    print_xpath_nodes(xpathObj.GetNodesetval())
    
    return nil
}

func print_xpath_nodes(nodes []XmlNode) {

}
func usage(name string) {
    fmt.Fprintf(os.Stderr,  "Usage: %s <xml-file> <xpath-expr> [<known-ns-list>]\n", name);
    fmt.Fprintf(os.Stderr,  "where <known-ns-list> is a list of known namespaces\n");
    fmt.Fprintf(os.Stderr,  "in \"<prefix1>=<href1> <prefix2>=href2> ...\" format\n");
}


func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Fprintf(os.Stderr, "Error: wrong number of arguments.\n")
		usage(os.Args[0])
		return
    }
    XmlCheckVersion()
    XmlInitParser()
    
     /* Do the main job */
    var err error;
    if len(os.Args) == 3 {
    	err = execute_xpath_expression(os.Args[1], os.Args[2],"")
    } else {
    	err = execute_xpath_expression(os.Args[1], os.Args[2],os.Args[3])
    }
    
    if err != nil {
		usage(os.Args[0])
		return
    }
    
    XmlCleanupParser()
	XmlMemoryDump()
}

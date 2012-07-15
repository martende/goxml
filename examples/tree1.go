package main

import (
	. "goxml"
	"os"
	"fmt"
)

/*

Analog http://www.xmlsoft.org/examples/tree1.c

Usage:

./tree1 test2.xml

*/

func print_element_names(a_node *XmlNode) {

    for cur_node:=a_node;cur_node!=nil;cur_node = cur_node.GetNext() {
    	if (cur_node.GetType() ==XML_ELEMENT_NODE) {
    		fmt.Printf("node type: Element, name: %s\n", cur_node.GetName())
    	}
    	print_element_names(cur_node.GetChildren())
    }
    
    
    return
}

func main() {
	if len(os.Args) != 2 {
		return
    }
    defer XmlCleanupParser()
    
    XmlCheckVersion()
    
	doc,err:=XmlReadFile(os.Args[1], "", 0)
	if err != nil {
        fmt.Fprint(os.Stderr,"Failed to parse %s\n", os.Args[1])
        return
    }
    
    defer XmlFreeDoc(doc)
    
    root_element:=XmlDocGetRootElement(doc)
    print_element_names(root_element)

}

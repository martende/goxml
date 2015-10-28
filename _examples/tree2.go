package main

import (
	. "github.com/martende/goxml"
	"os"
	"fmt"
)

/*

Analog http://www.xmlsoft.org/examples/tree2.c

Usage:

./tree2 result.html

*/

func main() {
	
    XmlCheckVersion()
    
    defer XmlMemoryDump()
    defer XmlCleanupParser()
    
    doc,err := XmlNewDoc("1.0")
    if (err != nil) {
    	fmt.Fprint(os.Stderr,"Failed to create new document %s\n",err)
    	return
    }
    defer XmlFreeDoc(doc)
    root_node,err := XmlNewNode(nil, "root")
    if (err != nil) {
    	fmt.Fprint(os.Stderr,"Failed to create new root node %s\n",err)
    	return
    }
    XmlDocSetRootElement(doc, root_node)
    /*
     * Creates a DTD declaration. Isn't mandatory. 
     */
     _,err = XmlCreateIntSubset(doc, "root", "", "tree2.dtd")
     if (err != nil) {
    	fmt.Fprint(os.Stderr,"Failed to create new DTD declaration %s\n",err)
    	return
     }
     XmlNewChild(root_node, nil, "node1", "content of node 1")
     XmlNewChild(root_node, nil, "node2", "")
     node,err := XmlNewChild(root_node, nil, "node3", "this node has attributes")
     if (err != nil) {
    	fmt.Fprint(os.Stderr,"Failed to create new node %s\n",err)
    	return
     }
     XmlNewProp(node, "attribute", "yes")
     XmlNewProp(node, "foo", "bar")
     node,err = XmlNewNode(nil ,"node4")
     if (err != nil) {
    	fmt.Fprint(os.Stderr,"Failed to create new node %s\n",err)
    	return
     }
     node1,err := XmlNewText("other way to create content (which is also a node)")
     if (err != nil) {
    	fmt.Fprint(os.Stderr,"Failed to create new content %s\n",err)
    	return
     }
     XmlAddChild(node, node1)
     XmlAddChild(root_node, node)
     for i:= 5; i < 7; i++ {
     	 buff := fmt.Sprintf("node%d", i)
     	 _,err= XmlNewChild(root_node, nil, buff, "")
     	 if (err != nil) {
     	 	 fmt.Fprint(os.Stderr,"Failed to create new node %s\n",err)
     	 	 return
    	 }
     }
     if len(os.Args) > 1 {
     	 XmlSaveFormatFileEnc(os.Args[1], doc, "UTF-8", 1)
     } else {
     	 XmlSaveFormatFileEnc("-", doc, "UTF-8", 1)
     }
     

}

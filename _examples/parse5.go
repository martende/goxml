package main

import (
	"github.com/martende/goxml"
	"os"
	"fmt"
)

var xmlData = `<?xml version="1.0"?>
<catalog>
<book id="bk101" available="true">
<author>Gambardella, Matthew</author>
<title>XML Developer's Guide</title>
<genre>Computer</genre>
<price>44.95</price>
<publish_date>2000-10-01</publish_date>
<description>An in-depth look at creating applications 
with XML.</description>
</book>
<book id="bk102" available="false">
<author>Ralls, Kim</author>
<title>Midnight Rain</title>
<genre>Fantasy</genre>
<price>5.95</price>
<publish_date>2000-12-16</publish_date>
<description>A former architect battles corporate zombies, 
an evil sorceress, and her own childhood to become queen 
of the world.</description>
</book>
</catalog>`
	
func main() {
	
	doc,_:=goxml.XmlParseDoc(xmlData)
	defer goxml.XmlFreeDoc(doc)
	
	// Dump Document
	// goxml.XmlDocDump(os.Stderr,doc)
	
	root:= goxml.XmlDocGetRootElement(doc)
	
	// Dump Node 
	// goxml.XmlElemDump(os.Stderr,doc,root)
	// or so 
	// goxml.XmlElemDump(os.Stderr,nil,root)
	
	// Get Children libxml2 api
	fmt.Printf("Print Root children via libxml2 api root=%08x\n",root.GetHandler())
	for cur_node:=root.GetChildren();cur_node!=nil;cur_node = cur_node.GetNext() {
		switch cur_node.GetType() {
			case goxml.XML_ELEMENT_NODE:
				fmt.Printf("node:%08x type: Element, name: %s\n",cur_node.GetHandler(), cur_node.GetName())
			case goxml.XML_TEXT_NODE:
				fmt.Printf("node:%08x type: XML_TEXT_NODE, name: %s Content:[%s]\n",cur_node.GetHandler(), cur_node.GetName(),goxml.XmlNodeGetContent(cur_node))
			default:
				fmt.Printf("node:%08x type: %d, name: %s\n",cur_node.GetHandler(), cur_node.GetType(),cur_node.GetName())
		}
    }
	
	// Get Children as list
	fmt.Printf("Print Root children via root.GetAllChildren root=%08x\n",root.GetHandler())
	childNodes:= root.GetAllChildren()
	fmt.Fprintf(os.Stdout,"Root have %d children\n" , len(childNodes) )
	for i,v := range childNodes {
		fmt.Fprintf(os.Stdout,"%d. Node[%s] Type[%v]\n",i,v.GetName(),v.GetType())
	}
	
	// Get Attrs as list
	attrs:= childNodes[1].GetAllProperties()
	for i,v := range attrs {
		fmt.Fprintf(os.Stdout,"%d. Node[%s] Attr %s has value=\"%s\"\n",i,childNodes[1].GetName() , v.GetName(), goxml.XmlNodeGetContent(v.GetChildren()) )
	}
	
	// Get Attrs as Map
	
	mapattrs:= childNodes[1].GetMapProperties()
	
	fmt.Fprintf(os.Stdout,"mapattrs = %v\n",mapattrs)
	
}

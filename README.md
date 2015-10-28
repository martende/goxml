# goxml - libxml2 go binding
##Info

goxml is pure libxml2 binding to go with at most same syntax as native c library.

### Folowing rules were used for reformating 

- all fields are replaced as getters
f.valid replaced as f.GetValid()
- all functions replaced with uppercased analog
- all "creating something" functions return object,err
- all data are strings

### Supported libxml interfaces
- encoding.h implemented 0 from 19 functions
- xpath.h implemented 4 from 38 functions
- xmlreader.h implemented 13 from 85 functions
- tree.h implemented 16 from 157 functions
- xpathInternals.h implemented 1 from 117 functions
- xmlmemory.h implemented 1 from 17 functions
- HTMLparser.h implemented 34 from 38 functions
- parser.h implemented 11 from 70 functions
- xmlstring.h implemented 0 from 29 functions


### Supported libxml structures
- XmlNode
- XmlAttr
- XmlNs
- XmlDoc
- XmlDtd
- HtmlEntityDesc
- HtmlElemDesc
- XmlTextReader
- XmlDict
- XmlXPathObject
- XmlXPathContext
- XmlParserCtxt
- XmlParserInput

htmlNode , htmlDoc and etc.. are replaced with XmlNode, XmlDoc
### Examples
- parse5.go


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

 look _examples/* and wiki
##TODO:

Change Array Implementation

  	func (this *XmlXPathObject) GetNodesetval() []*XmlNode {
			if this.handler.nodesetval == nil || this.handler.nodesetval.nodeNr == 0 {
				return nil
			}
			l:=int(this.handler.nodesetval.nodeNr)
			ret:=make([]*XmlNode,l)
			
			off:=(uintptr)(unsafe.Pointer(this.handler.nodesetval.nodeTab))
			
			for i:=0;i<l;i++ {
				t:=*(*C.xmlNodePtr)(unsafe.Pointer(off))
				ret[i] = &XmlNode{handler:t}
				off+=unsafe.Sizeof(t)
			}
			return ret
		}

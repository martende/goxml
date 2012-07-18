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
- tree.h
- parser.h
- xmlreader.h
- xmlmemory.h
- xmlstring.h
- HTMLparser.h
- encoding.h
- xpath.h
- xpathInternals.h

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
- parse3.go


        package main
        
        import (
          . "goxml"
        	"os"
        	"fmt"
        )
        /*
        Analog http://www.xmlsoft.org/examples/parse3.c
        Usage:
        ./parse3
        */
        
        var document = "<doc/>";
        func example3Func(content string) {
            doc,err := XmlReadMemory(content, "noname.xml", "", 0)
            if err != nil {
                fmt.Fprintf(os.Stderr, "Failed to parse document\n")
            }
            XmlFreeDoc(doc);
        }
        
        func main() {
        	XmlCheckVersion()
            example3Func(document)
        	XmlCleanupParser()
        	XmlMemoryDump()
        }

 look examples/* and wiki
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

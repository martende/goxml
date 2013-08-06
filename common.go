package goxml
/*
// #cgo pkg-config: libxml-2.0
// #include <libxml/tree.h>


import "C"
import "unsafe"
import "fmt"
import "os"
*/

/*
	High Level Interface functions
*/

func HtmlExtractXpathNodes(buffer []byte,xpathExpr string,baseUrl string)  (ret []*XmlNode,err error)  {
	data := string(buffer)
	ctxt,err:=HtmlCreateMemoryParserCtxt(data)
	if err != nil {
		return
	}
	defer HtmlFreeParserCtxt(ctxt)

	doc,err:=HtmlCtxtReadDoc(ctxt,data,baseUrl,"",HTML_PARSE_RECOVER | HTML_PARSE_NOERROR | HTML_PARSE_NOWARNING)

	if err != nil {
		return
	}

	defer XmlFreeDoc(doc)
	
	xpathCtx,err := XmlXPathNewContext(doc)

	if err != nil {
		return
	}

	defer XmlXPathFreeContext(xpathCtx)

    xpathObj,err := XmlXPathEvalExpression(xpathExpr,xpathCtx)
    if err != nil {
        //fmt.Fprintf(os.Stderr,"Error: unable to evaluate xpath expression  \"%s\"\n" , xpathExpr)
        //return fmt.Errorf("error %d",1)
        return
    }
    
    defer XmlXPathFreeObject(xpathObj)

	ret = xpathObj.GetNodesetval()

	return
}
package goxml_test

import (
	"fmt"
	"github.com/martende/goxml"
	. "launchpad.net/gocheck"
)

const gen_nb_unsigned_char_ptr = 1
const gen_nb_int_ptr = 2
const gen_nb_const_unsigned_char_ptr = 1

type HTTPparser_suite struct{}

var _ = Suite(&HTTPparser_suite{})

func gen_int_ptr(no int, nr int) *int {
	/*if no == 0 {
		return inttab
	}*/
	return nil
}

func (s *HTTPparser_suite) SetUpSuite(c *C) {
	fmt.Printf("")
	//inttab := make([]int,1024)
}

func (s *HTTPparser_suite) TestUTF8ToHtml(c *C) {
	i1 := ""
	r1, _ := goxml.UTF8ToHtml(i1)
	c.Check(r1, Equals, "")
	i2 := "alle"
	r2, _ := goxml.UTF8ToHtml(i2)
	c.Check(r2, Equals, "alle")
	i3 := "alleß"
	r3, _ := goxml.UTF8ToHtml(i3)
	c.Check(r3, Equals, "alle&szlig;")
}

func (s *HTTPparser_suite) TesthtmlAttrAllowed(c *C) {
	r := goxml.HtmlAttrAllowed(nil, "attr1", 1)
	c.Check(r, Equals, goxml.HTML_INVALID)
}

func (s *S) TesthtmlAutoCloseTag(c *C) {
	r := goxml.HtmlAutoCloseTag(nil, "TAG", nil)
	c.Check(r, Equals, 1)
	//fmt.Printf("R=%d",r)
}

func (s *S) TesthtmlCreateMemoryParserCtxt(c *C) {
	r1, err1 := goxml.HtmlCreateMemoryParserCtxt("<html></html>")
	c.Check(err1, Equals, nil)
	c.Check(r1, Not(Equals), nil)
	r2, err2 := goxml.HtmlCreateMemoryParserCtxt("ichwisenicht")
	c.Check(err2, Equals, nil)
	c.Check(r2, Not(Equals), nil)
}

func (s *S) TestHtmlCreatePushParserCtxt(c *C) {
	r, err := goxml.HtmlCreatePushParserCtxt(nil, "<htm", "filename.txt", goxml.XML_CHAR_ENCODING_UTF8)
	c.Check(err, Equals, nil)
	c.Check(r, Not(Equals), nil)
}

func (s *S) TestHtmlCtxtReadDoc(c *C) {
	r, err := goxml.HtmlCtxtReadDoc(nil, "<document></document>", "http://baseurl/", "", 0)
	c.Check(err, Not(Equals), nil)
	c.Check(r, Equals, (*goxml.XmlDoc)(nil))
	content := "<html></html>"
	ctxt, err := goxml.HtmlCreateMemoryParserCtxt(content)
	c.Log("AA",ctxt,err,ctxt == nil)
	c.Check(ctxt, Not(IsNil))
	c.Check(err, Equals, nil)
	r, err = goxml.HtmlCtxtReadDoc(ctxt, content, "http://baseurl/", "", 0)
	c.Check(err, Equals, nil)
	c.Check(r, Not(Equals), (*goxml.XmlDoc)(nil))
	c.Check(ctxt.GetValid(), Equals, 1)
}

func (s *S) TestHtmlCtxtReadFd(c *C) {
	r, err := goxml.HtmlCtxtReadFd(nil, 10, "http://baseurl/", "UTF-8", 0)
	c.Check(err, Not(Equals), nil)
	c.Check(r, Equals, (*goxml.XmlDoc)(nil))
}

func (s *S) TestHtmlCtxtReadFile(c *C) {
	r, err := goxml.HtmlCtxtReadFile(nil, "file://tmp/noexists.html", "UTF-8", goxml.HTML_PARSE_NOBLANKS|goxml.HTML_PARSE_NOERROR|goxml.HTML_PARSE_NOWARNING|goxml.HTML_PARSE_NONET)
	c.Check(err, Not(Equals), nil)
	c.Check(r, Equals, (*goxml.XmlDoc)(nil))

	ctxt := goxml.HtmlNewParserCtxt()
	r, err = goxml.HtmlCtxtReadFile(ctxt, "examples/test.html", "UTF-8", goxml.HTML_PARSE_NOBLANKS|goxml.HTML_PARSE_NOERROR|goxml.HTML_PARSE_NOWARNING|goxml.HTML_PARSE_NONET)

	c.Check(err, Equals, nil)
	c.Check(r, Not(Equals), (*goxml.XmlDoc)(nil))
	c.Check(ctxt.GetValid(), Equals, 1)

}

func (s *S) TestHtmlCtxtReadMemory(c *C) {
	ctxt, _ := goxml.HtmlCreateMemoryParserCtxt("<html></html>")
	c.Check(ctxt, Not(Equals), nil)
	r, err := goxml.HtmlCtxtReadMemory(ctxt, "<html></html>", "http://baseurl/", "UTF-8", 0)
	c.Check(err, Equals, nil)
	c.Check(r, Not(Equals), (*goxml.XmlDoc)(nil))
	c.Check(ctxt.GetValid(), Equals, 1)
}

func (s *S) TestHtmlCtxtReset(c *C) {
}

func (s *S) TestHtmlCtxtUseOptions(c *C) {
}

func (s *S) TestHtmlElementAllowedHere(c *C) {
}

func (s *S) TestHtmlElementStatusHere(c *C) {
}

func (s *S) TestHtmlEncodeEntities(c *C) {
}

func (s *S) TestHtmlEntityLookup(c *C) {
	// what is it i dont know
	r := goxml.HtmlEntityLookup("body")
	c.Check(r, Equals, (*goxml.HtmlEntityDesc)(nil))
}

func (s *S) TestHtmlEntityValueLookup(c *C) {
}

func (s *S) TestHtmlHandleOmittedElem(c *C) {
}

func (s *S) TestHtmlIsAutoClosed(c *C) {
}

func (s *S) TestHtmlIsScriptAttribute(c *C) {
}

func (s *S) test_htmlNewParserCtxt(c *C) {
}

func (s *S) test_htmlNodeStatus(c *C) {
}

func (s *S) test_htmlParseCharRef(c *C) {
}

func (s *S) test_htmlParseChunk(c *C) {
}

func (s *S) test_htmlParseDoc(c *C) {
}

func (s *S) test_htmlParseDocument(c *C) {
}

func (s *S) test_htmlParseElement(c *C) {
}

func (s *S) test_htmlParseEntityRef(c *C) {
}

func (s *S) test_htmlParseFile(c *C) {
}

func (s *S) test_htmlReadDoc(c *C) {
}

func (s *S) TestHtmlReadFile(c *C) {
	doc, err := goxml.HtmlReadFile("file://tmp/noexists.html", "UTF-8", goxml.HTML_PARSE_NOBLANKS|goxml.HTML_PARSE_NOERROR|goxml.HTML_PARSE_NOWARNING|goxml.HTML_PARSE_NONET)
	c.Check(doc, Equals, (*goxml.XmlDoc)(nil))
	c.Check(err, Not(Equals), nil)

	doc, err = goxml.HtmlReadFile("./examples/test.html", "UTF-8", goxml.HTML_PARSE_NOBLANKS|goxml.HTML_PARSE_NOERROR|goxml.HTML_PARSE_NOWARNING|goxml.HTML_PARSE_NONET)
	c.Check(doc, Not(Equals), (*goxml.XmlDoc)(nil))
	c.Check(err, Equals, nil)

	doc, err = goxml.HtmlReadFile("./examples/test_broken.html", "UTF-8", goxml.HTML_PARSE_NOBLANKS|goxml.HTML_PARSE_NOERROR|goxml.HTML_PARSE_NOWARNING|goxml.HTML_PARSE_NONET)
	c.Check(doc, Not(Equals), (*goxml.XmlDoc)(nil))
	c.Check(err, Equals, nil)
}

func (s *S) test_htmlReadMemory(c *C) {
}

func (s *S) test_htmlSAXParseDoc(c *C) {
}

func (s *S) test_htmlSAXParseFile(c *C) {
}

func (s *S) test_htmlTagLookup(c *C) {
}

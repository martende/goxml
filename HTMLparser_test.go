package goxml_test

import (
    . "launchpad.net/gocheck"
	"fmt"
	"goxml"    
)

const gen_nb_unsigned_char_ptr=1
const gen_nb_int_ptr=2
const gen_nb_const_unsigned_char_ptr=1

type HTTPparser_suite struct{}

var _ = Suite(&HTTPparser_suite{})


func gen_int_ptr(no int,nr int) *int{
	/*if no == 0 {
		return inttab
	}*/
	return nil
}

func (s *HTTPparser_suite) SetUpSuite(c *C) {
	//inttab := make([]int,1024)
}
func (s *HTTPparser_suite) TestUTF8ToHtml(c *C) {
	var out,in string
	var outlen,inlen int
	in = "test1"
	inlen = len(in);
	outlen = 1000;
	goxml.UTF8ToHtml(&out,&outlen,in,&inlen)
	fmt.Printf("out = %s in=%s\n" , out,in);
}

func (s *HTTPparser_suite) TesthtmlAttrAllowed(c *C) {
	fmt.Printf("TesthtmlAttrAllowed\n")
}

func (s *S) TesthtmlAutoCloseTag(c *C) {
	fmt.Printf("test_htmlAutoCloseTag\n")
}

func (s *S) test_htmlCreateMemoryParserCtxt(c *C) {
}

func (s *S) test_htmlCreatePushParserCtxt(c *C) {
}

func (s *S) test_htmlCtxtReadDoc(c *C) {
}

func (s *S) test_htmlCtxtReadFile(c *C) {
}

func (s *S) test_htmlCtxtReadMemory(c *C) {
}

func (s *S) test_htmlCtxtReset(c *C) {
}

func (s *S) test_htmlCtxtUseOptions(c *C) {
}

func (s *S) test_htmlElementAllowedHere(c *C) {
}

func (s *S) test_htmlElementStatusHere(c *C) {
}

func (s *S) test_htmlEncodeEntities(c *C) {
}

func (s *S) test_htmlEntityLookup(c *C) {
}

func (s *S) test_htmlEntityValueLookup(c *C) {
}

func (s *S) test_htmlHandleOmittedElem(c *C) {
}

func (s *S) test_htmlIsAutoClosed(c *C) {
}

func (s *S) test_htmlIsScriptAttribute(c *C) {
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

func (s *S) test_htmlReadFile(c *C) {
}

func (s *S) test_htmlReadMemory(c *C) {
}

func (s *S) test_htmlSAXParseDoc(c *C) {
}

func (s *S) test_htmlSAXParseFile(c *C) {
}

func (s *S) test_htmlTagLookup(c *C) {
}


package goxml_test

import (
    . "launchpad.net/gocheck"
	"fmt"
	"goxml"    
)

const gen_nb_unsigned_char_ptr=1
const gen_nb_int_ptr=2
const gen_nb_const_unsigned_char_ptr=1



gen_int_ptr(no int,nr int) *int{
	/*if no == 0 {
		return inttab
	}*/
	return nil
}

func (s *S) SetUpSuite(c *C) {
	inttab =make([]int,1024)
}
func (s *S) TestUTF8ToHtml(c *C) {
	//fmt.Printf("AAAAAAAAAAAAAAAAAA\n")
	
	for n_out:= 0;n_out < gen_nb_unsigned_char_ptr;n_out++ {
    	for n_outlen := 0;n_outlen < gen_nb_int_ptr;n_outlen++ {
    		for n_in := 0;n_in < gen_nb_const_unsigned_char_ptr;n_in++ {
    			for n_inlen := 0;n_inlen < gen_nb_int_ptr;n_inlen++ {
    				mem_base:=goxml.XmlMemBlocks()
    				out:=nil
    				outlen:=gen_int_ptr(n_outlen, 1)
    				fmt.Printf("%i \n",mem_base)
    			}
    		}
    	}
    }
}

func (s *S) test_htmlAttrAllowed(c *C) {
}

func (s *S) test_htmlAutoCloseTag(c *C) {
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


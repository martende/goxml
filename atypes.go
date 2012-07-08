package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/tree.h>
#include <libxml/parser.h>
#include <libxml/xmlreader.h>
#include <libxml/xmlmemory.h>
*/
import "C"
type XmlParserCtxt struct {
	handler C.xmlParserCtxtPtr
	_myDoc *XmlDoc
}
/*
func (this *XmlParserCtxt) GetSax() struct _xmlSAXHandler* {
}
*/
/*
func (this *XmlParserCtxt) GetUserData() void* {
}
*/
func (this *XmlParserCtxt) GetMyDoc() *XmlDoc {
	if this.handler.myDoc == nil {
		return nil
	}
	if this._myDoc == nil {
		this._myDoc = &XmlDoc{}
	}
	this._myDoc.handler = this.handler.myDoc
	return this._myDoc
}
func (this *XmlParserCtxt) GetWellFormed() int {
	return int(this.handler.wellFormed)
}
func (this *XmlParserCtxt) GetReplaceEntities() int {
	return int(this.handler.replaceEntities)
}
/*
func (this *XmlParserCtxt) GetVersion() xmlChar* {
}
*/
/*
func (this *XmlParserCtxt) GetEncoding() xmlChar* {
}
*/
func (this *XmlParserCtxt) GetStandalone() int {
	return int(this.handler.standalone)
}
func (this *XmlParserCtxt) GetHtml() int {
	return int(this.handler.html)
}
/*
func (this *XmlParserCtxt) GetInput() xmlParserInput* {
}
*/
func (this *XmlParserCtxt) GetInputNr() int {
	return int(this.handler.inputNr)
}
func (this *XmlParserCtxt) GetInputMax() int {
	return int(this.handler.inputMax)
}
/*
func (this *XmlParserCtxt) GetInputTab() xmlParserInput** {
}
*/
/*
func (this *XmlParserCtxt) GetNode() xmlNode* {
}
*/
func (this *XmlParserCtxt) GetNodeNr() int {
	return int(this.handler.nodeNr)
}
func (this *XmlParserCtxt) GetNodeMax() int {
	return int(this.handler.nodeMax)
}
/*
func (this *XmlParserCtxt) GetNodeTab() xmlNode** {
}
*/
func (this *XmlParserCtxt) GetRecord_info() int {
	return int(this.handler.record_info)
}
/*
func (this *XmlParserCtxt) GetNode_seq() xmlParserNodeInfoSeq {
}
*/
func (this *XmlParserCtxt) GetErrNo() int {
	return int(this.handler.errNo)
}
func (this *XmlParserCtxt) GetHasExternalSubset() int {
	return int(this.handler.hasExternalSubset)
}
func (this *XmlParserCtxt) GetHasPErefs() int {
	return int(this.handler.hasPErefs)
}
func (this *XmlParserCtxt) GetExternal() int {
	return int(this.handler.external)
}
func (this *XmlParserCtxt) GetValid() int {
	return int(this.handler.valid)
}
func (this *XmlParserCtxt) GetValidate() int {
	return int(this.handler.validate)
}
/*
func (this *XmlParserCtxt) GetVctxt() xmlValidCtxt {
}
*/
/*
func (this *XmlParserCtxt) GetInstate() xmlParserInputState {
}
*/
func (this *XmlParserCtxt) GetToken() int {
	return int(this.handler.token)
}
/*
func (this *XmlParserCtxt) GetDirectory() char* {
}
*/
/*
func (this *XmlParserCtxt) GetName() xmlChar* {
}
*/
func (this *XmlParserCtxt) GetNameNr() int {
	return int(this.handler.nameNr)
}
func (this *XmlParserCtxt) GetNameMax() int {
	return int(this.handler.nameMax)
}
/*
func (this *XmlParserCtxt) GetNameTab() xmlChar** {
}
*/
/*
func (this *XmlParserCtxt) GetNbChars() long int {
}
*/
/*
func (this *XmlParserCtxt) GetCheckIndex() long int {
}
*/
func (this *XmlParserCtxt) GetKeepBlanks() int {
	return int(this.handler.keepBlanks)
}
func (this *XmlParserCtxt) GetDisableSAX() int {
	return int(this.handler.disableSAX)
}
func (this *XmlParserCtxt) GetInSubset() int {
	return int(this.handler.inSubset)
}
/*
func (this *XmlParserCtxt) GetIntSubName() xmlChar* {
}
*/
/*
func (this *XmlParserCtxt) GetExtSubURI() xmlChar* {
}
*/
/*
func (this *XmlParserCtxt) GetExtSubSystem() xmlChar* {
}
*/
/*
func (this *XmlParserCtxt) GetSpace() int* {
}
*/
func (this *XmlParserCtxt) GetSpaceNr() int {
	return int(this.handler.spaceNr)
}
func (this *XmlParserCtxt) GetSpaceMax() int {
	return int(this.handler.spaceMax)
}
/*
func (this *XmlParserCtxt) GetSpaceTab() int* {
}
*/
func (this *XmlParserCtxt) GetDepth() int {
	return int(this.handler.depth)
}
/*
func (this *XmlParserCtxt) GetEntity() xmlParserInput* {
}
*/
func (this *XmlParserCtxt) GetCharset() int {
	return int(this.handler.charset)
}
func (this *XmlParserCtxt) GetNodelen() int {
	return int(this.handler.nodelen)
}
func (this *XmlParserCtxt) GetNodemem() int {
	return int(this.handler.nodemem)
}
func (this *XmlParserCtxt) GetPedantic() int {
	return int(this.handler.pedantic)
}
/*
func (this *XmlParserCtxt) Get_private() void* {
}
*/
func (this *XmlParserCtxt) GetLoadsubset() int {
	return int(this.handler.loadsubset)
}
func (this *XmlParserCtxt) GetLinenumbers() int {
	return int(this.handler.linenumbers)
}
/*
func (this *XmlParserCtxt) GetCatalogs() void* {
}
*/
func (this *XmlParserCtxt) GetRecovery() int {
	return int(this.handler.recovery)
}
func (this *XmlParserCtxt) GetProgressive() int {
	return int(this.handler.progressive)
}
/*
func (this *XmlParserCtxt) GetDict() xmlDict* {
}
*/
/*
func (this *XmlParserCtxt) GetAtts() xmlChar** {
}
*/
func (this *XmlParserCtxt) GetMaxatts() int {
	return int(this.handler.maxatts)
}
func (this *XmlParserCtxt) GetDocdict() int {
	return int(this.handler.docdict)
}
/*
func (this *XmlParserCtxt) GetStr_xml() xmlChar* {
}
*/
/*
func (this *XmlParserCtxt) GetStr_xmlns() xmlChar* {
}
*/
/*
func (this *XmlParserCtxt) GetStr_xml_ns() xmlChar* {
}
*/
func (this *XmlParserCtxt) GetSax2() int {
	return int(this.handler.sax2)
}
func (this *XmlParserCtxt) GetNsNr() int {
	return int(this.handler.nsNr)
}
func (this *XmlParserCtxt) GetNsMax() int {
	return int(this.handler.nsMax)
}
/*
func (this *XmlParserCtxt) GetNsTab() xmlChar** {
}
*/
/*
func (this *XmlParserCtxt) GetAttallocs() int* {
}
*/
/*
func (this *XmlParserCtxt) GetPushTab() void** {
}
*/
/*
func (this *XmlParserCtxt) GetAttsDefault() xmlHashTable* {
}
*/
/*
func (this *XmlParserCtxt) GetAttsSpecial() xmlHashTable* {
}
*/
func (this *XmlParserCtxt) GetNsWellFormed() int {
	return int(this.handler.nsWellFormed)
}
func (this *XmlParserCtxt) GetOptions() int {
	return int(this.handler.options)
}
func (this *XmlParserCtxt) GetDictNames() int {
	return int(this.handler.dictNames)
}
func (this *XmlParserCtxt) GetFreeElemsNr() int {
	return int(this.handler.freeElemsNr)
}
/*
func (this *XmlParserCtxt) GetFreeElems() xmlNode* {
}
*/
func (this *XmlParserCtxt) GetFreeAttrsNr() int {
	return int(this.handler.freeAttrsNr)
}
/*
func (this *XmlParserCtxt) GetFreeAttrs() xmlAttr* {
}
*/
/*
func (this *XmlParserCtxt) GetLastError() xmlError {
}
*/
/*
func (this *XmlParserCtxt) GetParseMode() xmlParserMode {
}
*/
/*
func (this *XmlParserCtxt) GetNbentities() long unsigned int {
}
*/
/*
func (this *XmlParserCtxt) GetSizeentities() long unsigned int {
}
*/
/*
func (this *XmlParserCtxt) GetNodeInfo() xmlParserNodeInfo* {
}
*/
func (this *XmlParserCtxt) GetNodeInfoNr() int {
	return int(this.handler.nodeInfoNr)
}
func (this *XmlParserCtxt) GetNodeInfoMax() int {
	return int(this.handler.nodeInfoMax)
}
/*
func (this *XmlParserCtxt) GetNodeInfoTab() xmlParserNodeInfo* {
}
*/

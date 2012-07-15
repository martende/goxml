package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/parser.h>

*/
import "C"
import "unsafe"
import "fmt"



/*
	Element inputTab has not registered type xmlParserInputPtr* 
	Element inputTab not recognized getter for elType:xmlParserInputPtr* goType:xmlParserInputPtr*
	Element nodeTab has not registered type xmlNodePtr* 
	Element nodeTab not recognized getter for elType:xmlNodePtr* goType:xmlNodePtr*
	Element node_seq has not registered type xmlParserNodeInfoSeq 
	Element node_seq not recognized getter for elType:xmlParserNodeInfoSeq goType:xmlParserNodeInfoSeq
	Element vctxt has not registered type xmlValidCtxt 
	Element vctxt not recognized getter for elType:xmlValidCtxt goType:xmlValidCtxt
	Element instate has not registered type xmlParserInputState 
	Element instate not recognized getter for elType:xmlParserInputState goType:xmlParserInputState
	Element nameTab has not registered type xmlChar** 
	Element nameTab not recognized getter for elType:xmlChar** goType:xmlChar**
	Element nbChars has not registered type long 
	Element nbChars not recognized getter for elType:long goType:long
	Element checkIndex has not registered type long 
	Element checkIndex not recognized getter for elType:long goType:long
	Element atts has not registered type xmlChar** 
	Element atts not recognized getter for elType:xmlChar** goType:xmlChar**
	Element nsTab has not registered type xmlChar** 
	Element nsTab not recognized getter for elType:xmlChar** goType:xmlChar**
	Element pushTab has not registered type void** 
	Element pushTab not recognized getter for elType:void** goType:void**
	Element attsDefault has not registered type xmlHashTablePtr 
	Element attsDefault not recognized getter for elType:xmlHashTablePtr goType:xmlHashTablePtr
	Element attsSpecial has not registered type xmlHashTablePtr 
	Element attsSpecial not recognized getter for elType:xmlHashTablePtr goType:xmlHashTablePtr
	Element lastError has not registered type xmlError 
	Element lastError not recognized getter for elType:xmlError goType:xmlError
	Element parseMode has not registered type xmlParserMode 
	Element parseMode not recognized getter for elType:xmlParserMode goType:xmlParserMode
	Element nbentities has not registered type unsigned long 
	Element nbentities not recognized getter for elType:unsigned long goType:unsigned long
	Element sizeentities has not registered type unsigned long 
	Element sizeentities not recognized getter for elType:unsigned long goType:unsigned long
	Element nodeInfo has not registered type xmlParserNodeInfo* 
	Element nodeInfo not recognized getter for elType:xmlParserNodeInfo* goType:xmlParserNodeInfo*
	Element nodeInfoTab has not registered type xmlParserNodeInfo* 
	Element nodeInfoTab not recognized getter for elType:xmlParserNodeInfo* goType:xmlParserNodeInfo*

*/
type XmlParserCtxt struct {
	handler C.xmlParserCtxtPtr
	_sax *XmlSAXHandler
	// userData void* // Private
	_myDoc *XmlDoc
	_input *XmlParserInput
	_node *XmlNode
	// space int* // Private
	// spaceTab int* // Private
	_entity *XmlParserInput
	// _private void* // Private
	// catalogs void* // Private
	_dict *XmlDict
	// attallocs int* // Private
	_freeElems *XmlNode
	_freeAttrs *XmlAttr
}
func (this *XmlParserCtxt) GetSax() *XmlSAXHandler {
	if this.handler.sax == nil {
		return nil
	}
	if this._sax == nil {
		this._sax = &XmlSAXHandler{}
	}
	this._sax.handler = (C.xmlSAXHandlerPtr)(unsafe.Pointer(this.handler.sax))
	return this._sax
}
func (this *XmlParserCtxt) GetMyDoc() *XmlDoc {
	if this.handler.myDoc == nil {
		return nil
	}
	if this._myDoc == nil {
		this._myDoc = &XmlDoc{}
	}
	this._myDoc.handler = (C.xmlDocPtr)(unsafe.Pointer(this.handler.myDoc))
	return this._myDoc
}
func (this *XmlParserCtxt) GetWellFormed() int {
	return int(this.handler.wellFormed)
}
func (this *XmlParserCtxt) GetReplaceEntities() int {
	return int(this.handler.replaceEntities)
}
func (this *XmlParserCtxt) GetVersion() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.version)))
}
func (this *XmlParserCtxt) GetEncoding() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.encoding)))
}
func (this *XmlParserCtxt) GetStandalone() int {
	return int(this.handler.standalone)
}
func (this *XmlParserCtxt) GetHtml() int {
	return int(this.handler.html)
}
func (this *XmlParserCtxt) GetInput() *XmlParserInput {
	if this.handler.input == nil {
		return nil
	}
	if this._input == nil {
		this._input = &XmlParserInput{}
	}
	this._input.handler = (C.xmlParserInputPtr)(unsafe.Pointer(this.handler.input))
	return this._input
}
func (this *XmlParserCtxt) GetInputNr() int {
	return int(this.handler.inputNr)
}
func (this *XmlParserCtxt) GetInputMax() int {
	return int(this.handler.inputMax)
}
/*
func (this *XmlParserCtxt) GetInputTab() xmlParserInputPtr* {
	return int(this.handler.inputTab)
}
*/
func (this *XmlParserCtxt) GetNode() *XmlNode {
	if this.handler.node == nil {
		return nil
	}
	if this._node == nil {
		this._node = &XmlNode{}
	}
	this._node.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.node))
	return this._node
}
func (this *XmlParserCtxt) GetNodeNr() int {
	return int(this.handler.nodeNr)
}
func (this *XmlParserCtxt) GetNodeMax() int {
	return int(this.handler.nodeMax)
}
/*
func (this *XmlParserCtxt) GetNodeTab() xmlNodePtr* {
	return int(this.handler.nodeTab)
}
*/
func (this *XmlParserCtxt) GetRecord_info() int {
	return int(this.handler.record_info)
}
/*
func (this *XmlParserCtxt) GetNode_seq() xmlParserNodeInfoSeq {
	return int(this.handler.node_seq)
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
	return int(this.handler.vctxt)
}
*/
/*
func (this *XmlParserCtxt) GetInstate() xmlParserInputState {
	return int(this.handler.instate)
}
*/
func (this *XmlParserCtxt) GetToken() int {
	return int(this.handler.token)
}
func (this *XmlParserCtxt) GetDirectory() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.directory)))
}
func (this *XmlParserCtxt) GetName() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.name)))
}
func (this *XmlParserCtxt) GetNameNr() int {
	return int(this.handler.nameNr)
}
func (this *XmlParserCtxt) GetNameMax() int {
	return int(this.handler.nameMax)
}
/*
func (this *XmlParserCtxt) GetNameTab() xmlChar** {
	return int(this.handler.nameTab)
}
*/
/*
func (this *XmlParserCtxt) GetNbChars() long {
	return int(this.handler.nbChars)
}
*/
/*
func (this *XmlParserCtxt) GetCheckIndex() long {
	return int(this.handler.checkIndex)
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
func (this *XmlParserCtxt) GetIntSubName() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.intSubName)))
}
func (this *XmlParserCtxt) GetExtSubURI() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.extSubURI)))
}
func (this *XmlParserCtxt) GetExtSubSystem() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.extSubSystem)))
}
func (this *XmlParserCtxt) GetSpaceNr() int {
	return int(this.handler.spaceNr)
}
func (this *XmlParserCtxt) GetSpaceMax() int {
	return int(this.handler.spaceMax)
}
func (this *XmlParserCtxt) GetDepth() int {
	return int(this.handler.depth)
}
func (this *XmlParserCtxt) GetEntity() *XmlParserInput {
	if this.handler.entity == nil {
		return nil
	}
	if this._entity == nil {
		this._entity = &XmlParserInput{}
	}
	this._entity.handler = (C.xmlParserInputPtr)(unsafe.Pointer(this.handler.entity))
	return this._entity
}
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
func (this *XmlParserCtxt) GetLoadsubset() int {
	return int(this.handler.loadsubset)
}
func (this *XmlParserCtxt) GetLinenumbers() int {
	return int(this.handler.linenumbers)
}
func (this *XmlParserCtxt) GetRecovery() int {
	return int(this.handler.recovery)
}
func (this *XmlParserCtxt) GetProgressive() int {
	return int(this.handler.progressive)
}
func (this *XmlParserCtxt) GetDict() *XmlDict {
	if this.handler.dict == nil {
		return nil
	}
	if this._dict == nil {
		this._dict = &XmlDict{}
	}
	this._dict.handler = (C.xmlDictPtr)(unsafe.Pointer(this.handler.dict))
	return this._dict
}
/*
func (this *XmlParserCtxt) GetAtts() xmlChar** {
	return int(this.handler.atts)
}
*/
func (this *XmlParserCtxt) GetMaxatts() int {
	return int(this.handler.maxatts)
}
func (this *XmlParserCtxt) GetDocdict() int {
	return int(this.handler.docdict)
}
func (this *XmlParserCtxt) GetStr_xml() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.str_xml)))
}
func (this *XmlParserCtxt) GetStr_xmlns() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.str_xmlns)))
}
func (this *XmlParserCtxt) GetStr_xml_ns() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.str_xml_ns)))
}
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
	return int(this.handler.nsTab)
}
*/
/*
func (this *XmlParserCtxt) GetPushTab() void** {
	return int(this.handler.pushTab)
}
*/
/*
func (this *XmlParserCtxt) GetAttsDefault() xmlHashTablePtr {
	return int(this.handler.attsDefault)
}
*/
/*
func (this *XmlParserCtxt) GetAttsSpecial() xmlHashTablePtr {
	return int(this.handler.attsSpecial)
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
func (this *XmlParserCtxt) GetFreeElems() *XmlNode {
	if this.handler.freeElems == nil {
		return nil
	}
	if this._freeElems == nil {
		this._freeElems = &XmlNode{}
	}
	this._freeElems.handler = (C.xmlNodePtr)(unsafe.Pointer(this.handler.freeElems))
	return this._freeElems
}
func (this *XmlParserCtxt) GetFreeAttrsNr() int {
	return int(this.handler.freeAttrsNr)
}
func (this *XmlParserCtxt) GetFreeAttrs() *XmlAttr {
	if this.handler.freeAttrs == nil {
		return nil
	}
	if this._freeAttrs == nil {
		this._freeAttrs = &XmlAttr{}
	}
	this._freeAttrs.handler = (C.xmlAttrPtr)(unsafe.Pointer(this.handler.freeAttrs))
	return this._freeAttrs
}
/*
func (this *XmlParserCtxt) GetLastError() xmlError {
	return int(this.handler.lastError)
}
*/
/*
func (this *XmlParserCtxt) GetParseMode() xmlParserMode {
	return int(this.handler.parseMode)
}
*/
/*
func (this *XmlParserCtxt) GetNbentities() unsigned long {
	return int(this.handler.nbentities)
}
*/
/*
func (this *XmlParserCtxt) GetSizeentities() unsigned long {
	return int(this.handler.sizeentities)
}
*/
/*
func (this *XmlParserCtxt) GetNodeInfo() xmlParserNodeInfo* {
	return int(this.handler.nodeInfo)
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
	return int(this.handler.nodeInfoTab)
}
*/
/*
	Element buf has not registered type xmlParserInputBufferPtr 
	Element buf not recognized getter for elType:xmlParserInputBufferPtr goType:xmlParserInputBufferPtr
	Element consumed has not registered type unsigned long 
	Element consumed not recognized getter for elType:unsigned long goType:unsigned long
	Element free has not registered type xmlParserInputDeallocate 
	Element free not recognized getter for elType:xmlParserInputDeallocate goType:xmlParserInputDeallocate

*/
type XmlParserInput struct {
	handler C.xmlParserInputPtr
}
/*
func (this *XmlParserInput) GetBuf() xmlParserInputBufferPtr {
	return int(this.handler.buf)
}
*/
func (this *XmlParserInput) GetFilename() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.filename)))
}
func (this *XmlParserInput) GetDirectory() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.directory)))
}
func (this *XmlParserInput) GetBase() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.base)))
}
func (this *XmlParserInput) GetCur() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.cur)))
}
func (this *XmlParserInput) GetEnd() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.end)))
}
func (this *XmlParserInput) GetLength() int {
	return int(this.handler.length)
}
func (this *XmlParserInput) GetLine() int {
	return int(this.handler.line)
}
func (this *XmlParserInput) GetCol() int {
	return int(this.handler.col)
}
/*
func (this *XmlParserInput) GetConsumed() unsigned long {
	return int(this.handler.consumed)
}
*/
/*
func (this *XmlParserInput) GetFree() xmlParserInputDeallocate {
	return int(this.handler.free)
}
*/
func (this *XmlParserInput) GetEncoding() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.encoding)))
}
func (this *XmlParserInput) GetVersion() string {
	return C.GoString((*C.char)(unsafe.Pointer(this.handler.version)))
}
func (this *XmlParserInput) GetStandalone() int {
	return int(this.handler.standalone)
}
func (this *XmlParserInput) GetId() int {
	return int(this.handler.id)
}
/*
	Element internalSubset has not registered type internalSubsetSAXFunc 
	Element internalSubset not recognized getter for elType:internalSubsetSAXFunc goType:internalSubsetSAXFunc
	Element isStandalone has not registered type isStandaloneSAXFunc 
	Element isStandalone not recognized getter for elType:isStandaloneSAXFunc goType:isStandaloneSAXFunc
	Element hasInternalSubset has not registered type hasInternalSubsetSAXFunc 
	Element hasInternalSubset not recognized getter for elType:hasInternalSubsetSAXFunc goType:hasInternalSubsetSAXFunc
	Element hasExternalSubset has not registered type hasExternalSubsetSAXFunc 
	Element hasExternalSubset not recognized getter for elType:hasExternalSubsetSAXFunc goType:hasExternalSubsetSAXFunc
	Element resolveEntity has not registered type resolveEntitySAXFunc 
	Element resolveEntity not recognized getter for elType:resolveEntitySAXFunc goType:resolveEntitySAXFunc
	Element getEntity has not registered type getEntitySAXFunc 
	Element getEntity not recognized getter for elType:getEntitySAXFunc goType:getEntitySAXFunc
	Element entityDecl has not registered type entityDeclSAXFunc 
	Element entityDecl not recognized getter for elType:entityDeclSAXFunc goType:entityDeclSAXFunc
	Element notationDecl has not registered type notationDeclSAXFunc 
	Element notationDecl not recognized getter for elType:notationDeclSAXFunc goType:notationDeclSAXFunc
	Element attributeDecl has not registered type attributeDeclSAXFunc 
	Element attributeDecl not recognized getter for elType:attributeDeclSAXFunc goType:attributeDeclSAXFunc
	Element elementDecl has not registered type elementDeclSAXFunc 
	Element elementDecl not recognized getter for elType:elementDeclSAXFunc goType:elementDeclSAXFunc
	Element unparsedEntityDecl has not registered type unparsedEntityDeclSAXFunc 
	Element unparsedEntityDecl not recognized getter for elType:unparsedEntityDeclSAXFunc goType:unparsedEntityDeclSAXFunc
	Element setDocumentLocator has not registered type setDocumentLocatorSAXFunc 
	Element setDocumentLocator not recognized getter for elType:setDocumentLocatorSAXFunc goType:setDocumentLocatorSAXFunc
	Element startDocument has not registered type startDocumentSAXFunc 
	Element startDocument not recognized getter for elType:startDocumentSAXFunc goType:startDocumentSAXFunc
	Element endDocument has not registered type endDocumentSAXFunc 
	Element endDocument not recognized getter for elType:endDocumentSAXFunc goType:endDocumentSAXFunc
	Element startElement has not registered type startElementSAXFunc 
	Element startElement not recognized getter for elType:startElementSAXFunc goType:startElementSAXFunc
	Element endElement has not registered type endElementSAXFunc 
	Element endElement not recognized getter for elType:endElementSAXFunc goType:endElementSAXFunc
	Element reference has not registered type referenceSAXFunc 
	Element reference not recognized getter for elType:referenceSAXFunc goType:referenceSAXFunc
	Element characters has not registered type charactersSAXFunc 
	Element characters not recognized getter for elType:charactersSAXFunc goType:charactersSAXFunc
	Element ignorableWhitespace has not registered type ignorableWhitespaceSAXFunc 
	Element ignorableWhitespace not recognized getter for elType:ignorableWhitespaceSAXFunc goType:ignorableWhitespaceSAXFunc
	Element processingInstruction has not registered type processingInstructionSAXFunc 
	Element processingInstruction not recognized getter for elType:processingInstructionSAXFunc goType:processingInstructionSAXFunc
	Element comment has not registered type commentSAXFunc 
	Element comment not recognized getter for elType:commentSAXFunc goType:commentSAXFunc
	Element warning has not registered type warningSAXFunc 
	Element warning not recognized getter for elType:warningSAXFunc goType:warningSAXFunc
	Element error has not registered type errorSAXFunc 
	Element error not recognized getter for elType:errorSAXFunc goType:errorSAXFunc
	Element fatalError has not registered type fatalErrorSAXFunc 
	Element fatalError not recognized getter for elType:fatalErrorSAXFunc goType:fatalErrorSAXFunc
	Element getParameterEntity has not registered type getParameterEntitySAXFunc 
	Element getParameterEntity not recognized getter for elType:getParameterEntitySAXFunc goType:getParameterEntitySAXFunc
	Element cdataBlock has not registered type cdataBlockSAXFunc 
	Element cdataBlock not recognized getter for elType:cdataBlockSAXFunc goType:cdataBlockSAXFunc
	Element externalSubset has not registered type externalSubsetSAXFunc 
	Element externalSubset not recognized getter for elType:externalSubsetSAXFunc goType:externalSubsetSAXFunc
	Element initialized not recognized getter for elType:unsigned int goType:uint
	Element startElementNs has not registered type startElementNsSAX2Func 
	Element startElementNs not recognized getter for elType:startElementNsSAX2Func goType:startElementNsSAX2Func
	Element endElementNs has not registered type endElementNsSAX2Func 
	Element endElementNs not recognized getter for elType:endElementNsSAX2Func goType:endElementNsSAX2Func
	Element serror has not registered type xmlStructuredErrorFunc 
	Element serror not recognized getter for elType:xmlStructuredErrorFunc goType:xmlStructuredErrorFunc

*/
type XmlSAXHandler struct {
	handler C.xmlSAXHandlerPtr
	// _private void* // Private
}
/*
func (this *XmlSAXHandler) GetInternalSubset() internalSubsetSAXFunc {
	return int(this.handler.internalSubset)
}
*/
/*
func (this *XmlSAXHandler) GetIsStandalone() isStandaloneSAXFunc {
	return int(this.handler.isStandalone)
}
*/
/*
func (this *XmlSAXHandler) GetHasInternalSubset() hasInternalSubsetSAXFunc {
	return int(this.handler.hasInternalSubset)
}
*/
/*
func (this *XmlSAXHandler) GetHasExternalSubset() hasExternalSubsetSAXFunc {
	return int(this.handler.hasExternalSubset)
}
*/
/*
func (this *XmlSAXHandler) GetResolveEntity() resolveEntitySAXFunc {
	return int(this.handler.resolveEntity)
}
*/
/*
func (this *XmlSAXHandler) GetGetEntity() getEntitySAXFunc {
	return int(this.handler.getEntity)
}
*/
/*
func (this *XmlSAXHandler) GetEntityDecl() entityDeclSAXFunc {
	return int(this.handler.entityDecl)
}
*/
/*
func (this *XmlSAXHandler) GetNotationDecl() notationDeclSAXFunc {
	return int(this.handler.notationDecl)
}
*/
/*
func (this *XmlSAXHandler) GetAttributeDecl() attributeDeclSAXFunc {
	return int(this.handler.attributeDecl)
}
*/
/*
func (this *XmlSAXHandler) GetElementDecl() elementDeclSAXFunc {
	return int(this.handler.elementDecl)
}
*/
/*
func (this *XmlSAXHandler) GetUnparsedEntityDecl() unparsedEntityDeclSAXFunc {
	return int(this.handler.unparsedEntityDecl)
}
*/
/*
func (this *XmlSAXHandler) GetSetDocumentLocator() setDocumentLocatorSAXFunc {
	return int(this.handler.setDocumentLocator)
}
*/
/*
func (this *XmlSAXHandler) GetStartDocument() startDocumentSAXFunc {
	return int(this.handler.startDocument)
}
*/
/*
func (this *XmlSAXHandler) GetEndDocument() endDocumentSAXFunc {
	return int(this.handler.endDocument)
}
*/
/*
func (this *XmlSAXHandler) GetStartElement() startElementSAXFunc {
	return int(this.handler.startElement)
}
*/
/*
func (this *XmlSAXHandler) GetEndElement() endElementSAXFunc {
	return int(this.handler.endElement)
}
*/
/*
func (this *XmlSAXHandler) GetReference() referenceSAXFunc {
	return int(this.handler.reference)
}
*/
/*
func (this *XmlSAXHandler) GetCharacters() charactersSAXFunc {
	return int(this.handler.characters)
}
*/
/*
func (this *XmlSAXHandler) GetIgnorableWhitespace() ignorableWhitespaceSAXFunc {
	return int(this.handler.ignorableWhitespace)
}
*/
/*
func (this *XmlSAXHandler) GetProcessingInstruction() processingInstructionSAXFunc {
	return int(this.handler.processingInstruction)
}
*/
/*
func (this *XmlSAXHandler) GetComment() commentSAXFunc {
	return int(this.handler.comment)
}
*/
/*
func (this *XmlSAXHandler) GetWarning() warningSAXFunc {
	return int(this.handler.warning)
}
*/
/*
func (this *XmlSAXHandler) GetError() errorSAXFunc {
	return int(this.handler.error)
}
*/
/*
func (this *XmlSAXHandler) GetFatalError() fatalErrorSAXFunc {
	return int(this.handler.fatalError)
}
*/
/*
func (this *XmlSAXHandler) GetGetParameterEntity() getParameterEntitySAXFunc {
	return int(this.handler.getParameterEntity)
}
*/
/*
func (this *XmlSAXHandler) GetCdataBlock() cdataBlockSAXFunc {
	return int(this.handler.cdataBlock)
}
*/
/*
func (this *XmlSAXHandler) GetExternalSubset() externalSubsetSAXFunc {
	return int(this.handler.externalSubset)
}
*/
/*
func (this *XmlSAXHandler) GetInitialized() uint {
	return int(this.handler.initialized)
}
*/
/*
func (this *XmlSAXHandler) GetStartElementNs() startElementNsSAX2Func {
	return int(this.handler.startElementNs)
}
*/
/*
func (this *XmlSAXHandler) GetEndElementNs() endElementNsSAX2Func {
	return int(this.handler.endElementNs)
}
*/
/*
func (this *XmlSAXHandler) GetSerror() xmlStructuredErrorFunc {
	return int(this.handler.serror)
}
*/

/* 
	   Function: xmlParseChunk
	   ReturnType: int
	   Args: (('ctxt', ['xmlParserCtxtPtr'], None), ('chunk', ['char', '*'], None), ('size', ['int'], None), ('terminate', ['int'], None))
*/
func XmlParseChunk(ctxt *XmlParserCtxt,chunk string,size int,terminate int) int {
	var c_ctxt C.xmlParserCtxtPtr=nil
	if ctxt !=nil { c_ctxt = (C.xmlParserCtxtPtr)(ctxt.handler) }
	c_chunk:= (*C.char)(unsafe.Pointer(C.CString(chunk)))
	defer C.free(unsafe.Pointer(c_chunk))
	c_size := C.int(size)
	c_terminate := C.int(terminate)

	c_ret := C.xmlParseChunk(c_ctxt,c_chunk,c_size,c_terminate)



	return int(c_ret)
}
/* 
	   Function: xmlCleanupParser
	   ReturnType: void
	   Args: ((None, ['void'], None),)
*/
func XmlCleanupParser() {


	C.xmlCleanupParser()




}
/* 
	   Function: xmlCreatePushParserCtxt
	   ReturnType: xmlParserCtxtPtr
	   Args: (('sax', ['xmlSAXHandlerPtr'], None), ('user_data', ['void', '*'], None), ('chunk', ['char', '*'], None), ('size', ['int'], None), ('filename', ['char', '*'], None))
*/
func XmlCreatePushParserCtxt(sax *XmlSAXHandler,chunk string,filename string) (g_ret *XmlParserCtxt,err error) {
	var c_sax C.xmlSAXHandlerPtr=nil
	if sax !=nil { c_sax = (C.xmlSAXHandlerPtr)(sax.handler) }
	c_chunk:= (*C.char)(unsafe.Pointer(C.CString(chunk)))
	defer C.free(unsafe.Pointer(c_chunk))
	c_filename:= (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(c_filename))
	c_size:=C.int(len(chunk)*1)
	c_ret := C.xmlCreatePushParserCtxt(c_sax,nil,c_chunk,c_size,c_filename)

	if c_ret == nil {
		err = fmt.Errorf("xmlCreatePushParserCtxt errno %d" ,c_ret)
	} else {
		g_ret =  &XmlParserCtxt{handler:(C.xmlParserCtxtPtr)(c_ret)}
	}
	return
}
/* 
	   Function: xmlCtxtReadFile
	   ReturnType: xmlDocPtr
	   Args: (('ctxt', ['xmlParserCtxtPtr'], None), ('filename', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
func XmlCtxtReadFile(ctxt *XmlParserCtxt,filename string,encoding string,options int) (g_ret *XmlDoc,err error) {
	var c_ctxt C.xmlParserCtxtPtr=nil
	if ctxt !=nil { c_ctxt = (C.xmlParserCtxtPtr)(ctxt.handler) }
	c_filename:= (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(c_filename))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	c_options := C.int(options)

	c_ret := C.xmlCtxtReadFile(c_ctxt,c_filename,c_encoding,c_options)

	if c_ret == nil {
		err = fmt.Errorf("xmlCtxtReadFile errno %d" ,c_ret)
	} else {
		g_ret =  &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
	}
	return
}
/* 
	   Function: xmlNewParserCtxt
	   ReturnType: xmlParserCtxtPtr
	   Args: ((None, ['void'], None),)
*/
func XmlNewParserCtxt() *XmlParserCtxt {


	c_ret := C.xmlNewParserCtxt()



	if c_ret == nil {return nil}
	return &XmlParserCtxt{handler:(C.xmlParserCtxtPtr)(c_ret)}
}
/* 
	   Function: xmlFreeParserCtxt
	   ReturnType: void
	   Args: (('ctxt', ['xmlParserCtxtPtr'], None),)
*/
func XmlFreeParserCtxt(ctxt *XmlParserCtxt) {
	var c_ctxt C.xmlParserCtxtPtr=nil
	if ctxt !=nil { c_ctxt = (C.xmlParserCtxtPtr)(ctxt.handler) }

	C.xmlFreeParserCtxt(c_ctxt)




}
/* 
	   Function: xmlReadMemory
	   ReturnType: xmlDocPtr
	   Args: (('buffer', ['char', '*'], None), ('size', ['int'], None), ('URL', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
func XmlReadMemory(buffer string,URL string,encoding string,options int) (g_ret *XmlDoc,err error) {
	c_buffer:= (*C.char)(unsafe.Pointer(C.CString(buffer)))
	defer C.free(unsafe.Pointer(c_buffer))
	c_URL:= (*C.char)(unsafe.Pointer(C.CString(URL)))
	defer C.free(unsafe.Pointer(c_URL))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	c_options := C.int(options)
	c_size:=C.int(len(buffer)*1)
	c_ret := C.xmlReadMemory(c_buffer,c_size,c_URL,c_encoding,c_options)

	if c_ret == nil {
		err = fmt.Errorf("xmlReadMemory errno %d" ,c_ret)
	} else {
		g_ret =  &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
	}
	return
}
/* 
	   Function: xmlReadFile
	   ReturnType: xmlDocPtr
	   Args: (('URL', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
func XmlReadFile(URL string,encoding string,options int) (g_ret *XmlDoc,err error) {
	c_URL:= (*C.char)(unsafe.Pointer(C.CString(URL)))
	defer C.free(unsafe.Pointer(c_URL))
	c_encoding:= (*C.char)(unsafe.Pointer(C.CString(encoding)))
	defer C.free(unsafe.Pointer(c_encoding))
	c_options := C.int(options)

	c_ret := C.xmlReadFile(c_URL,c_encoding,c_options)

	if c_ret == nil {
		err = fmt.Errorf("xmlReadFile errno %d" ,c_ret)
	} else {
		g_ret =  &XmlDoc{handler:(C.xmlDocPtr)(c_ret)}
	}
	return
}



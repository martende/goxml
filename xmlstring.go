package goxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/xmlstring.h>
*/
import "C"
import "unsafe"


/*
	Element inputTab has not registered type xmlParserInputPtr* 
	Element inputTab not recognized getter for type xmlParserInputPtr* 
	Element nodeTab has not registered type xmlNodePtr* 
	Element nodeTab not recognized getter for type xmlNodePtr* 
	Element node_seq has not registered type xmlParserNodeInfoSeq 
	Element node_seq not recognized getter for type xmlParserNodeInfoSeq 
	Element vctxt has not registered type xmlValidCtxt 
	Element vctxt not recognized getter for type xmlValidCtxt 
	Element instate has not registered type xmlParserInputState 
	Element instate not recognized getter for type xmlParserInputState 
	Element nameTab has not registered type xmlChar** 
	Element nameTab not recognized getter for type xmlChar** 
	Element nbChars has not registered type long 
	Element nbChars not recognized getter for type long 
	Element checkIndex has not registered type long 
	Element checkIndex not recognized getter for type long 
	Element space has not registered type int* 
	Element space not recognized getter for type int* 
	Element spaceTab has not registered type int* 
	Element spaceTab not recognized getter for type int* 
	Element catalogs has not registered type void* 
	Element catalogs not recognized getter for type void* 
	Element atts has not registered type xmlChar** 
	Element atts not recognized getter for type xmlChar** 
	Element nsTab has not registered type xmlChar** 
	Element nsTab not recognized getter for type xmlChar** 
	Element attallocs has not registered type int* 
	Element attallocs not recognized getter for type int* 
	Element pushTab has not registered type void** 
	Element pushTab not recognized getter for type void** 
	Element attsDefault has not registered type xmlHashTablePtr 
	Element attsDefault not recognized getter for type xmlHashTablePtr 
	Element attsSpecial has not registered type xmlHashTablePtr 
	Element attsSpecial not recognized getter for type xmlHashTablePtr 
	Element freeAttrs has not registered type xmlAttrPtr 
	Element freeAttrs not recognized getter for type xmlAttrPtr 
	Element lastError has not registered type xmlError 
	Element lastError not recognized getter for type xmlError 
	Element parseMode has not registered type xmlParserMode 
	Element parseMode not recognized getter for type xmlParserMode 
	Element nbentities has not registered type unsigned long 
	Element nbentities not recognized getter for type unsigned long 
	Element sizeentities has not registered type unsigned long 
	Element sizeentities not recognized getter for type unsigned long 
	Element nodeInfo has not registered type xmlParserNodeInfo* 
	Element nodeInfo not recognized getter for type xmlParserNodeInfo* 
	Element nodeInfoTab has not registered type xmlParserNodeInfo* 
	Element nodeInfoTab not recognized getter for type xmlParserNodeInfo* 

*/
type XmlParserCtxt struct {
	handler C.xmlParserCtxtPtr
	_sax *XmlSAXHandler
	// userData void* // Private
	_myDoc *XmlDoc
	_input *XmlParserInput
	_node *XmlNode
	_entity *XmlParserInput
	// _private void* // Private
	_dict *XmlDict
	_freeElems *XmlNode
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
/*
func (this *XmlParserCtxt) GetSpace() int* {
	return int(this.handler.space)
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
	return int(this.handler.spaceTab)
}
*/
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
/*
func (this *XmlParserCtxt) GetCatalogs() void* {
	return int(this.handler.catalogs)
}
*/
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
func (this *XmlParserCtxt) GetAttallocs() int* {
	return int(this.handler.attallocs)
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
/*
func (this *XmlParserCtxt) GetFreeAttrs() xmlAttrPtr {
	return int(this.handler.freeAttrs)
}
*/
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
	Element buf not recognized getter for type xmlParserInputBufferPtr 
	Element consumed has not registered type unsigned long 
	Element consumed not recognized getter for type unsigned long 
	Element free has not registered type xmlParserInputDeallocate 
	Element free not recognized getter for type xmlParserInputDeallocate 

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
	Element internalSubset not recognized getter for type internalSubsetSAXFunc 
	Element isStandalone has not registered type isStandaloneSAXFunc 
	Element isStandalone not recognized getter for type isStandaloneSAXFunc 
	Element hasInternalSubset has not registered type hasInternalSubsetSAXFunc 
	Element hasInternalSubset not recognized getter for type hasInternalSubsetSAXFunc 
	Element hasExternalSubset has not registered type hasExternalSubsetSAXFunc 
	Element hasExternalSubset not recognized getter for type hasExternalSubsetSAXFunc 
	Element resolveEntity has not registered type resolveEntitySAXFunc 
	Element resolveEntity not recognized getter for type resolveEntitySAXFunc 
	Element getEntity has not registered type getEntitySAXFunc 
	Element getEntity not recognized getter for type getEntitySAXFunc 
	Element entityDecl has not registered type entityDeclSAXFunc 
	Element entityDecl not recognized getter for type entityDeclSAXFunc 
	Element notationDecl has not registered type notationDeclSAXFunc 
	Element notationDecl not recognized getter for type notationDeclSAXFunc 
	Element attributeDecl has not registered type attributeDeclSAXFunc 
	Element attributeDecl not recognized getter for type attributeDeclSAXFunc 
	Element elementDecl has not registered type elementDeclSAXFunc 
	Element elementDecl not recognized getter for type elementDeclSAXFunc 
	Element unparsedEntityDecl has not registered type unparsedEntityDeclSAXFunc 
	Element unparsedEntityDecl not recognized getter for type unparsedEntityDeclSAXFunc 
	Element setDocumentLocator has not registered type setDocumentLocatorSAXFunc 
	Element setDocumentLocator not recognized getter for type setDocumentLocatorSAXFunc 
	Element startDocument has not registered type startDocumentSAXFunc 
	Element startDocument not recognized getter for type startDocumentSAXFunc 
	Element endDocument has not registered type endDocumentSAXFunc 
	Element endDocument not recognized getter for type endDocumentSAXFunc 
	Element startElement has not registered type startElementSAXFunc 
	Element startElement not recognized getter for type startElementSAXFunc 
	Element endElement has not registered type endElementSAXFunc 
	Element endElement not recognized getter for type endElementSAXFunc 
	Element reference has not registered type referenceSAXFunc 
	Element reference not recognized getter for type referenceSAXFunc 
	Element characters has not registered type charactersSAXFunc 
	Element characters not recognized getter for type charactersSAXFunc 
	Element ignorableWhitespace has not registered type ignorableWhitespaceSAXFunc 
	Element ignorableWhitespace not recognized getter for type ignorableWhitespaceSAXFunc 
	Element processingInstruction has not registered type processingInstructionSAXFunc 
	Element processingInstruction not recognized getter for type processingInstructionSAXFunc 
	Element comment has not registered type commentSAXFunc 
	Element comment not recognized getter for type commentSAXFunc 
	Element warning has not registered type warningSAXFunc 
	Element warning not recognized getter for type warningSAXFunc 
	Element error has not registered type errorSAXFunc 
	Element error not recognized getter for type errorSAXFunc 
	Element fatalError has not registered type fatalErrorSAXFunc 
	Element fatalError not recognized getter for type fatalErrorSAXFunc 
	Element getParameterEntity has not registered type getParameterEntitySAXFunc 
	Element getParameterEntity not recognized getter for type getParameterEntitySAXFunc 
	Element cdataBlock has not registered type cdataBlockSAXFunc 
	Element cdataBlock not recognized getter for type cdataBlockSAXFunc 
	Element externalSubset has not registered type externalSubsetSAXFunc 
	Element externalSubset not recognized getter for type externalSubsetSAXFunc 
	Element initialized has not registered type unsigned int 
	Element initialized not recognized getter for type unsigned int 
	Element startElementNs has not registered type startElementNsSAX2Func 
	Element startElementNs not recognized getter for type startElementNsSAX2Func 
	Element endElementNs has not registered type endElementNsSAX2Func 
	Element endElementNs not recognized getter for type endElementNsSAX2Func 
	Element serror has not registered type xmlStructuredErrorFunc 
	Element serror not recognized getter for type xmlStructuredErrorFunc 

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
func (this *XmlSAXHandler) GetInitialized() unsigned int {
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
	   Function: xmlCtxtReadFile
	   ReturnType: xmlDocPtr
	   Args: (('ctxt', ['xmlParserCtxtPtr'], None), ('filename', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
func XmlCtxtReadFile(ctxt *XmlParserCtxt,filename string,encoding string,options int) *XmlDoc {
	var c_ctxt C.xmlParserCtxtPtr=nil ;if ctxt !=nil { c_ctxt = ctxt.handler }
	c_filename:= C.CString(filename)
	c_encoding:= C.CString(encoding)
	c_options := C.int(options)
	c_ret := C.xmlCtxtReadFile(c_ctxt,c_filename,c_encoding,c_options)
	if c_ret == nil {return nil}
	return &XmlDoc{handler:c_ret}
}
/* 
	   Function: xmlNewParserCtxt
	   ReturnType: xmlParserCtxtPtr
	   Args: ((None, ['void'], None),)
*/
func XmlNewParserCtxt() *XmlParserCtxt {

	c_ret := C.xmlNewParserCtxt()
	if c_ret == nil {return nil}
	return &XmlParserCtxt{handler:c_ret}
}
/* 
	   Function: xmlReadFile
	   ReturnType: xmlDocPtr
	   Args: (('URL', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
func XmlReadFile(URL string,encoding string,options int) *XmlDoc {
	c_URL:= C.CString(URL)
	c_encoding:= C.CString(encoding)
	c_options := C.int(options)
	c_ret := C.xmlReadFile(c_URL,c_encoding,c_options)
	if c_ret == nil {return nil}
	return &XmlDoc{handler:c_ret}
}
/* 
	   Function: xmlFreeParserCtxt
	   ReturnType: void
	   Args: (('ctxt', ['xmlParserCtxtPtr'], None),)
*/
func XmlFreeParserCtxt(ctxt *XmlParserCtxt) {
	var c_ctxt C.xmlParserCtxtPtr=nil ;if ctxt !=nil { c_ctxt = ctxt.handler }
	C.xmlFreeParserCtxt(c_ctxt)
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
	   Function: xmlReadMemory
	   ReturnType: xmlDocPtr
	   Args: (('buffer', ['char', '*'], None), ('size', ['int'], None), ('URL', ['char', '*'], None), ('encoding', ['char', '*'], None), ('options', ['int'], None))
*/
func XmlReadMemory(buffer string,size int,URL string,encoding string,options int) *XmlDoc {
	c_buffer:= C.CString(buffer)
	c_size := C.int(size)
	c_URL:= C.CString(URL)
	c_encoding:= C.CString(encoding)
	c_options := C.int(options)
	c_ret := C.xmlReadMemory(c_buffer,c_size,c_URL,c_encoding,c_options)
	if c_ret == nil {return nil}
	return &XmlDoc{handler:c_ret}
}



# goxml - libxml2 go binding

TODO:

Change Array Implementation

		func (this *XmlXPathObject) GetNodesetval() []XmlNode {
		if this.handler.nodesetval == nil || this.handler.nodesetval.nodeNr == 0 {
			return nil
		}
		ret:=make([]XmlNode,this.handler.nodesetval.nodeNr)
	
		nodeTab := *(*[]C.xmlNodePtr)(unsafe.Pointer(this.handler.nodesetval.nodeTab))
		for i:=0;i<len(ret);i++ {
			k:= nodeTab[i]
			ret[i] = XmlNode{handler:k}
		}
		return ret
		}

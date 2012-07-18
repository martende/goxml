# goxml - libxml2 go binding

TODO:

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

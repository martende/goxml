from pyclibrary import CParser
import os.path
import os
import sys
import re
import optparse


TMP = "/tmp/tmp1"

CONF_PRINT_SKIPPED = False

def lookInDb(stype,fname,ptype,pname):
	for i in range(0,len(FUNC_DESCS),2):
		v=FUNC_DESCS[i]
		t = [stype,fname,ptype,pname]
		for j,tv in enumerate(v):
			if tv is None:
				t[j]=None
		if v == tuple(t):
			return FUNC_DESCS[i+1]
	return None

class FileConverter():
	def __init__(self,include):
		self.filename = include.split("/")[-1]
		self.include = include
		self.gofilename = re.sub(r'\.h$',".go",self.filename)
		self.unsafe = False
		self.consts = []
	def calcArgType(self,arg,sig):
		if arg == 'ret':
			return "".join(sig[0])
		for s in sig[1]:
			if arg==s[0]:
				return "".join(s[1])
		raise Exception("calcArgType not found Name:\"%s\" Signature:%s" %(arg,sig))
	def createFuncArgString(self,fname,sig):
		goFname = fname[0].upper() + fname[1:]
		cReturnType = "".join(sig[0]) 
		args = []
		errs = []
		goReturnType = None
		
		dbData = lookInDb('r',fname,cReturnType,None)
		if dbData:
			if (dbData[0] == 'CALC'):
				goReturnType=dbData[1].getReturnType(self.calcArgType(dbData[1].arg(),sig),cReturnType)
		if cReturnType in TYPEALIAS:
			cReturnType = TYPEALIAS[cReturnType]
		if goReturnType is None:
			try:
				goReturnType = TYPEINFO[cReturnType]['goReturnType']
			except:
				try:
					goReturnType = TYPEINFO[cReturnType]['goArgType']
				except:
					#goArgType = ptype
					goReturnType = cReturnType
					errs.append('Warn[createFuncArgString]: ReturnType %s Not defined' % (goReturnType))
		
		for (pname,ptype,_) in sig[1]:
			if pname is not None:
				ptype = "".join(ptype)
				dbData = lookInDb('f',fname,ptype,pname)
				goArgType = None
				if dbData:
					if dbData[0] in ('SKIP','CALC'):
						continue
					elif dbData[0] == 'RETYPE':
						ptype = dbData[1]
					else:
						raise (Exception('Not implemented %s' % str(dbData)))
				
				try:
					goArgType = TYPEINFO[ptype]['goArgType']
				except:
					goArgType = ptype
					errs.append('Warn: %s %s Not defined' % (pname,ptype))
					
				args.append("%s %s" % (pname,goArgType))
		
		arg_string ="func " + goFname + "(" + ",".join(args) +")"
		if goReturnType:
			arg_string+= " "+goReturnType
		
		return arg_string,errs
	
	def createInputInits(self,fname,sig):
		errs = []
		outVarsBlock = []
		recalcsBlock = []
		for (pname,ptype,_) in sig[1]:
			if pname is not None:
				ptype = "".join(ptype)
				dbData = lookInDb('f',fname,ptype,pname)
				go2cConvert = None
				if dbData:
					if 'SKIP' in dbData:
						continue
					elif dbData[0] == 'RETYPE':
						ptype = dbData[1]
					elif dbData[0] == 'CALC':
						recalcsBlock.append(dbData[1](pname,ptype))
						continue
					else:
						raise (Exception('Not implemented %s' % str(dbData)))
				try:
					go2cConvert = TYPEINFO[ptype]['go2cConverter']
				except:
					go2cConvert = lambda a,b:pname
					errs.append('Warn: %s %s No converter to C(go2cConverter)' % (pname,ptype))
				
				outVarsBlock.append(str(go2cConvert(pname,ptype)) )
		
		outStr = self.funcJoin(outVarsBlock)
		recalcStr = self.funcJoin(recalcsBlock)
		
		return "\n".join((outStr , recalcStr)),errs
	
	def createCallLine(self,fName,sig):
		
		cReturnType = "".join(sig[0])
		errs = []
		outs = []
		callargs = []
		for (pname,ptype,_) in sig[1]:
			ptype = "".join(ptype)
			dbData = lookInDb('f',fName,ptype,pname)
			if dbData:
				if 'SKIP' in dbData:
					callargs.append("nil")
					continue
				elif dbData[0] == 'RETYPE':
					ptype = dbData[1]
				elif dbData[0] == 'CALC':
					pass
				else:
					raise (Exception('Not implemented %s' % str(dbData)))
			if pname is not None:
				callargs.append("c_" + pname)
			
		callLine = 'C.' + fName + "(" + ",".join(callargs)+")"
		
		if ( cReturnType != 'void') :
			callLine = "\tc_ret := " + callLine + "\n"
		else:
			callLine = "\t" + callLine + "\n"
		
		returnBlock = []
		errorProcessBlock = []
		dbData = lookInDb('r',fName,cReturnType,None)
		if dbData and dbData[0] == 'CALC':
			mapper = dbData[1]
			errArg = mapper.errArg()
			if errArg == 'ret':
				errorCondition = mapper.errorCondition('c_ret')
				errorProcessBlock.append("""if %(errorCondition)s {
	err = fmt.Errorf("%(fName)s errno %%d" ,c_ret)
}"""% {"errorCondition":errorCondition,"fName":fName})
				returnBlock.append("return")
			else:
				raise Exception("Not Implemented")
			#goReturnType=dbData[1].getReturnType(self.calcArgType(dbData[1].arg(),sig),cReturnType)
		elif cReturnType != 'void' :
			if cReturnType in TYPEALIAS:
				cReturnType = TYPEALIAS[cReturnType]
			if cReturnType in TYPEINFO:
				if  'realType' in TYPEINFO[cReturnType]:
					cReturnType = TYPEINFO[cReturnType]['realType']
				returnConverter = None
				try:
					returnConverter = TYPEINFO[cReturnType]['returnConverter']
				except:
					returnConverter = lambda t,p:t
					errs.append('Warn[createCallLine]: type (%s) has no Return converter to Go' % (cReturnType))
				 
				goReturnType = None
				try:
					goReturnType = TYPEINFO[cReturnType]['goReturnType']
				except:
					try:
						goReturnType = re.sub(r'^\*','',TYPEINFO[cReturnType]['goArgType'])
					except:
						#goArgType = ptype
						goReturnType = cReturnType
						errs.append('Warn[createCallLine]: ReturnType %s found but not defined' % (goReturnType))
				
				returnBlock.append(returnConverter(cReturnType,goReturnType))
			else:
				errs.append('Warn[createCallLine]: type (%s) has no Return converter(no info in TYPEINFO) to Go' % (cReturnType))
		# Post Process Parameters
		postProcessBlock = []
		for (pname,ptype,_) in sig[1]:
			if pname is not None:
				ptype = "".join(ptype)
				dbData = lookInDb('f',fName,ptype,pname)
				if dbData:
					if dbData[0] == 'RETYPE':
						ptype = dbData[1]
					elif dbData[0] == 'CALC':
						continue
				postProcessor = TYPEINFO.get(ptype,{}).get('postProcessor')
				if postProcessor:
					postProcessBlock.append("\t" + postProcessor(pname))
		
		dbData = lookInDb('r',fName,cReturnType,None)
		if dbData and dbData[0] == 'CALC':
			mapper = dbData[1]
			arg1 = mapper.arg()
			if arg1!="ret":
				#print arg1, "\t",self.calcArgType(arg1,sig),"\t",mapper.getReturnTuple(self.calcArgType(arg1,sig),cReturnType)
				v = mapper.getReturnTuple(self.calcArgType(arg1,sig),cReturnType)
				vName = v[0][0]
				vType = v[0][1]
				cName = arg1
				cType = self.calcArgType(arg1,sig)
				cInfo = TYPEINFO[cType]
				if 'to-' + vType in cInfo:
					converter = cInfo['to-' + vType]
					postProcessBlock.append(converter(vName,"c_"+cName))
				else:
					raise Exception("Not implemented converstion to %s to from Type %s(%s)" % (vType,cType,cInfo))
			else:
				cType = self.calcArgType(arg1,sig)
				cInfo = TYPEINFO[cType]
				# avoid aliases problems
				if 'realType' in cInfo:
					cType = cInfo['realType']
				if 'c2GoConverter' in cInfo:
					postProcessBlock.append(cInfo['c2GoConverter']('g_ret',cInfo['goArgType'],'c_ret',cType))
				else:
					raise Exception("need c2GoConverter for type %s" % (cType))
			
		returnStr = self.funcJoin(returnBlock)
		postProcessStr = self.funcJoin(postProcessBlock)
		errorProcessStr= self.funcJoin(errorProcessBlock)
		if errorProcessStr and postProcessStr:
			errorProcessStr += " else {"
			postProcessStr = self.funcJoin((postProcessStr,"}"))
		return "\n".join((callLine,errorProcessStr,postProcessStr,returnStr)),errs
	def funcJoin(self,l):
		r = []
		for i in l:
			r.append("\n".join(["\t" + k for k in i.split("\n")]))
		return "\n".join(r)
	def log(self,s):
		if VERBOSE:
			print self.filename+": " +s
	
	def processFuncsList(self,functionSignatures):
		flist = []
		#print functionSignatures
		for fName in functionSignatures:
			inImports = fName in IMPORTS
			if not inImports:
				self.log("Skip %s"%fName)
			if inImports or CONF_PRINT_SKIPPED: 
				errs = []
				sig = functionSignatures[fName]
				funcComment = """/* 
	   Function: %s
	   ReturnType: %s
	   Args: %s
*/
"""	% ( fName , "".join(sig[0]) , sig[1])
				sig = list(sig)
				sig[1] = list(sig[1])
				funcCode = funcComment
				self.preprocessArgsSig(sig)
				funArgString,err = self.createFuncArgString(fName,sig)
				errs+=err
				inputInits,err = self.createInputInits(fName,sig)
				errs+=err
				callLine,err = self.createCallLine(fName,sig)
				errs+=err
				if not inImports or errs:
					funcCode+="/*\n"
				if errs:
					funcCode+= "\n\t" + "\n\t".join(errs) + "\n\n"
				
				funcCode+=funArgString +" {\n"
				
				funcCode+=inputInits + "\n"
				
				funcCode+=callLine + "\n";
				
				funcCode+="}\n"
				
				if not inImports or errs:
					funcCode+="\n*/\n"
				if 'unsafe.Pointer' in funcCode:
					self.unsafe = True
				flist.append(funcCode)
				
				if errs:
					self.log("Errored %s"%fName)
				else:
					self.log("Process %s"%fName)
			
		return flist
	def preprocessArgsSig(self,sig):
		c = 1
		for i,(n,t,_) in enumerate(sig[1]):
			if n is None and "".join(t) != 'void':
				sig[1][i]=("arg%i" % c , t, _)
				c+=1
		for i,(n,t,_) in enumerate(sig[1]):
			elType = "".join(t)
			if elType in TYPEALIAS:
				sig[1][i]=(n , TYPEALIAS[elType], _)
				
		
	def processConstsList(self,enumSignatures):
		r = []
		for enum in enumSignatures:
			for name in enumSignatures[enum]:
				r.append(name)
				#r.append("const %s = C.%s\n" % (name,name))
		
		return r
	def prepareStruct(self,sig,cStruct,innerStruct):
		fields = [('handler',"C."+innerStruct)]
		errs = []
		getters = []
		goStructName = TYPEINFO[innerStruct]['goArgType']
		goStructName = re.sub(r'^\*','',goStructName)
		
		for el in sig:
			errored = False
			retyped = False
			elName = el[0]
			fieldName=elName[0].upper()+elName[1:] 
			elType = "".join(el[1])
			
			info= lookInDb('s',innerStruct,elType,elName)
			self.log("Struct[" + cStruct + "] Field[n=%s,Ct=%s,db=%r]" % (elName,elType,info ))
			if info:
				if 'PRIVATE' in info:
					fields.append(('// ' + elName , elType + " // Private" ) )
					continue
				elif info[0] == 'RETYPE':
					elType = info[1] 
					retyped = True
				elif info[0] == 'PASS':
					# JUST NOP to skip wildmarks
					pass
			if not retyped and elType not in TYPEINFO:
				if elType in TYPEALIAS:
					elType =  TYPEALIAS[elType]
			
			if elType not in TYPEINFO:
				errored = True
				errs.append("Element %s has not registered type %s " % (elName,elType))

			goType = elType			
			try:
				goType = TYPEINFO[elType]['goArgType']
			except:
				pass
			inner = None
			
			if goType == 'int':
				inner = "func (this *%(goStructName)s) Get%(fieldName)s() %(goType)s {\n" % {'goType':goType,'fieldName':fieldName,'goStructName':goStructName}
				inner += "\treturn int(this.handler."+elName+")\n}\n"
			elif goType == '__XmlNode_reserved_type':
				inner = "func (this *%(goStructName)s) Get%(fieldName)s() %(goType)s {\n" % {'goType':'int','fieldName':fieldName,'goStructName':goStructName}
				inner += "\treturn int(C.XmlNode_fetch_type(this.handler))\n}\n"
			elif goType == 'byte':
				inner = "func (this *%(goStructName)s) Get%(fieldName)s() %(goType)s {\n" % {'goType':goType,'fieldName':fieldName,'goStructName':goStructName}
				inner += "\treturn byte(this.handler."+elName+")\n}\n"
			elif goType== '*string':
				inner = "func (this *%(goStructName)s) Get%(fieldName)s() %(goType)s {\n" % {'goType':goType,'fieldName':fieldName,'goStructName':goStructName}
				inner += "if this.handler."+elName+" == nil { return nil }\n"
				inner += "\ts:=C.GoString((*C.char)(unsafe.Pointer(this.handler."+elName+")))\n\treturn &s\n}\n"
				self.unsafe=True
			elif goType[0] == '*':
				self.unsafe=True
				# this._next.handler = (C.xmlNsPtr)(unsafe.Pointer(this.handler.next))
				constr = re.sub('^\*','&',goType)
				fields.append(("_"+elName,goType))
				inner = "func (this *%(goStructName)s) Get%(fieldName)s() %(goType)s {\n" % {'goType':goType,'fieldName':fieldName,'goStructName':goStructName}
				inner += "\tif this.handler.%s == nil {\n\t\treturn nil\n\t}\n" % elName
				inner += "\tif this._%s == nil {\n\t\tthis._%s = %s{}\n\t}\n" % (elName,elName,constr)
				inner += "\tthis._%s.handler = (C.%s)(unsafe.Pointer(this.handler.%s))\n" % (elName,elType,elName) 
				inner += "\treturn this._"+elName+"\n}\n"
			elif goType== 'string':
				inner = "func (this *%(goStructName)s) Get%(fieldName)s() %(goType)s {\n" % {'goType':goType,'fieldName':fieldName,'goStructName':goStructName}
				inner += "\treturn C.GoString((*C.char)(unsafe.Pointer(this.handler."+elName+")))\n}\n"
				self.unsafe=True
			else:
				errored = True
				errs.append("Element %s not recognized getter for elType:%s goType:%s" % (elName,elType,goType))
				inner = "func (this *%(goStructName)s) Get%(fieldName)s() %(goType)s {\n" % {'goType':goType,'fieldName':fieldName,'goStructName':goStructName}
				inner += "\treturn int(this.handler."+elName+")\n}\n"
			
			if inner:
				if errored:
					inner = "/*\n" + inner + "*/\n"
				getters.append(inner)
			
		
		struct_desc =""
		if (errs ):
			struct_desc+="/*"
			struct_desc+= "\n\t" + "\n\t".join(errs) + "\n\n"
			struct_desc+="*/\n"
		
		struct_desc+= "type " + goStructName +" struct {\n\t" 
		struct_desc+= "\n\t".join([" ".join(f) for f in fields]) +"\n}\n"
		
		struct_desc+= "".join(getters)
		#if (errs ):
		#	struct_desc+="*/\n"
		
		return struct_desc 
	def processStructs(self,structs):
		ss = []
		for s in structs:
			for t in TYPEINFO:
				if  TYPEINFO[t].get('exportStruct','')==s and structs[s]['members']:
					self.log("Parse Struct %s"%t)
					ss.append(self.prepareStruct(structs[s]['members'],s,t))
					break
		return ss
	def processFile(self):
		replace = {}
		if True:
			# enum a = 1 << 213 bug
			f=open(self.include,"r")
			s = f.read()
			s = re.findall("(\s+(\d+)<<(\d+)(,?))",s)
			for r in s:
				replace[r[0]] = str(int(r[1]) << int(r[2]) ) +  r[3] 
		
		# Read defines from version ( defines only - not funcs and etc )
		
		vp = CParser(
			"/".join(self.include.split("/")[:-1] + list(("xmlversion.h",))),
			cache=TMP + "/" + "xmlversion.h" + ".cache",
		)
		#print vp.defs['macros']
		#return
		macros = {
			'XMLCALL': '',
			#'LIBXML_PUSH_ENABLED':'',
			#'LIBXML_READER_ENABLED':''
		}
		macros.update(vp.defs['macros'])
		p = CParser(
			# TMP + "/" + self.filename + ".h",
			self.include,
			cache=TMP + "/" + self.filename + ".cache",
			#copyFrom=vp,
			replace  = replace,	
			#replace = {"\s+(\d+)<<(\d+)(,?)" : lambda x:str(int(x.group(1))<<int(x.group(2)))+ x.group(3)},
			macros=macros
		)
		if options.PRINTALL:
			p.printAll()
			return
		#mk = {'XMLCALL': '',
		#'LIBXML_PUSH_ENABLED':1
		#}
		
		#p = CParser(["/tmp/tmp1/parser.h.h"],cache=TMP + "/" + self.filename + ".cache",macros=mk)
		
				
		varsdict = {
			'filename' : self.filename,
			"imports" : '',
			'consts_list' : '',
			'structs_list' : '',
			'funcs_list' : '',
			'cpostprocessors':''
		}
		if self.filename in CPOSTPROCESSORS:
			varsdict['cpostprocessors'] = CPOSTPROCESSORS[self.filename] 
		#varsdict['consts_list'] = "".join(self.processConstsList(p.defs['enums']))
		
		self.consts = self.processConstsList(p.defs['enums'])
		if PROCESS_FUNCTIONS:
			varsdict['funcs_list'] = "".join(self.processFuncsList(p.defs['functions']))
		if PROCESS_STRUCTS:
			varsdict['structs_list'] = "".join(self.processStructs(p.defs['structs']))
		imports = []
		
		if "errors.New" in varsdict['funcs_list'] or "errors.New" in varsdict['structs_list']  :
			imports.append("import \"errors\"\n")
		if "unsafe.Pointer" in varsdict['funcs_list'] or "unsafe.Pointer" in varsdict['structs_list']  :
			imports.append("import \"unsafe\"\n")
		if "fmt.Errorf" in varsdict['funcs_list'] or "fmt.Errorf" in varsdict['structs_list']  :
			imports.append("import \"fmt\"\n")
		if "*os.File" in varsdict['funcs_list'] or "*File" in varsdict['structs_list']  :
			imports.append("import \"os\"\n")
		
		
		varsdict['imports'] =  	"".join(imports)
		
		open(self.gofilename,"w").write(GO_TPL % varsdict)
		
		
		
		
PROCESS_STRUCTS = False
PROCESS_FUNCTIONS = False
parser = optparse.OptionParser()

parser.add_option('--pf',
    action="store_true", dest="PROCESS_FUNCTIONS",
    help="Process Functions (default ALL)", default=None)

parser.add_option('--ps',
    action="store_true", dest="PROCESS_STRUCTS",
    help="Process Structures (default ALL)", default=None)

parser.add_option('-f', '--functions',
    action="append", dest="IMPORTS",
    help="Force change import functions", default=None)

parser.add_option('-i', '--includes',
    action="append", dest="INCLUDES",
    help="Force change import functions", default=None)

parser.add_option('-p', '--print-all',
    action="store_true", dest="PRINTALL",
    help="Print parser content", default=None)

parser.add_option('-v', '--verbose',
    action="count", dest="VERBOSE",
    help="Verbose", default=0)


options, args = parser.parse_args()

from goxml import *

if options.IMPORTS is not None:
	IMPORTS = options.IMPORTS
if options.INCLUDES is not None:
	INCLUDES = options.INCLUDES
if options.VERBOSE is not None:
	VERBOSE = options.VERBOSE
if options.PROCESS_STRUCTS is not None:
	PROCESS_STRUCTS = options.PROCESS_STRUCTS
if options.PROCESS_FUNCTIONS is not None:
	PROCESS_FUNCTIONS = options.PROCESS_FUNCTIONS

if not PROCESS_FUNCTIONS and not PROCESS_STRUCTS:
	# Process All
	PROCESS_STRUCTS = True
	PROCESS_FUNCTIONS = True

def convertAliases():
	for t in TYPEINFO:
		if isinstance(TYPEINFO[t],tuple):
			if TYPEINFO[t][0] == 'alias':
				parent = TYPEINFO[t][1]
				TYPEINFO[t] = {'realType' : parent}
				TYPEINFO[t].update(TYPEINFO[parent])
				if 'exportStruct' in TYPEINFO[t]:
					del TYPEINFO[t]['exportStruct']

def processImports():
	"""
		Convert simple rules in IMPORTS to dbRules
		@ - return error on null
	"""
	for i,func in enumerate(IMPORTS):
		if func.startswith('@'):
			IMPORTS[i] = IMPORTS[i][1:]
			FUNC_DESCS.insert(0,('CALC',return_mapper('ret','ret','%s == nil')))
			FUNC_DESCS.insert(0,('r',IMPORTS[i],None,None))
			
			
if not os.path.exists(TMP):
	os.mkdir(TMP) 

convertAliases()
processImports()
consts = []
includes = []
for include in INCLUDES:
	if VERBOSE:
		print "Parse: " + include 
	p = FileConverter(include)
	p.processFile()
	includes.append(p.filename)
	consts += p.consts

consts = dict(map(lambda x : (x ,1 ),consts)).keys()
incls = "\n".join(["#include <libxml/%s>" % fn for fn in includes])

varsdict = {
	'filename' : incls,
	'consts_list' : "\n".join(["const %s = C.%s" % (c,c) for c in consts],)
}

open("const.go","w").write("""package goxml
/*
#cgo pkg-config: libxml-2.0
%(filename)s
*/
import "C"

%(consts_list)s
""" % varsdict)

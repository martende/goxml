import sys
import glob
import string

debug=0
#debugsym='ignorableWhitespaceSAXFunc'
debugsym=None

ignored_words = {
}

class identifier:
    def __init__(self, name, header=None, module=None, type=None, lineno = 0,
                 info=None, extra=None, conditionals = None):
        self.name = name
	self.header = header
	self.module = module
	self.type = type
	self.info = info
	self.extra = extra
	self.lineno = lineno
	self.static = 0
	if conditionals == None or len(conditionals) == 0:
	    self.conditionals = None
	else:
	    self.conditionals = conditionals[:]
	if self.name == debugsym:
	    print "=> define %s : %s" % (debugsym, (module, type, info,
	                                 extra, conditionals))

    def __repr__(self):
        r = "%s %s:" % (self.type, self.name)
	if self.static:
	    r = r + " static"
	if self.module != None:
	    r = r + " from %s" % (self.module)
	if self.info != None:
	    r = r + " " +  `self.info`
	if self.extra != None:
	    r = r + " " + `self.extra`
	if self.conditionals != None:
	    r = r + " " + `self.conditionals`
	return r


    def set_header(self, header):
        self.header = header
    def set_module(self, module):
        self.module = module
    def set_type(self, type):
        self.type = type
    def set_info(self, info):
        self.info = info
    def set_extra(self, extra):
        self.extra = extra
    def set_lineno(self, lineno):
        self.lineno = lineno
    def set_static(self, static):
        self.static = static
    def set_conditionals(self, conditionals):
	if conditionals == None or len(conditionals) == 0:
	    self.conditionals = None
	else:
	    self.conditionals = conditionals[:]

    def get_name(self):
        return self.name
    def get_header(self):
        return self.module
    def get_module(self):
        return self.module
    def get_type(self):
        return self.type
    def get_info(self):
        return self.info
    def get_lineno(self):
        return self.lineno
    def get_extra(self):
        return self.extra
    def get_static(self):
        return self.static
    def get_conditionals(self):
        return self.conditionals

    def update(self, header, module, type = None, info = None, extra=None,
               conditionals=None):
	if self.name == debugsym:
	    print "=> update %s : %s" % (debugsym, (module, type, info,
	                                 extra, conditionals))
        if header != None and self.header == None:
	    self.set_header(module)
        if module != None and (self.module == None or self.header == self.module):
	    self.set_module(module)
        if type != None and self.type == None:
	    self.set_type(type)
        if info != None:
	    self.set_info(info)
        if extra != None:
	    self.set_extra(extra)
        if conditionals != None:
	    self.set_conditionals(conditionals)


class index:
    def __init__(self, name = "noname"):
        self.name = name
        self.identifiers = {}
        self.functions = {}
	self.variables = {}
	self.includes = {}
	self.structs = {}
	self.enums = {}
	self.typedefs = {}
	self.macros = {}
	self.references = {}
	self.info = {}

    def add_ref(self, name, header, module, static, type, lineno, info=None, extra=None, conditionals = None):
        if name[0:2] == '__':
	    return None
        d = None
        try:
	   d = self.identifiers[name]
	   d.update(header, module, type, lineno, info, extra, conditionals)
	except:
	   d = identifier(name, header, module, type, lineno, info, extra, conditionals)
	   self.identifiers[name] = d

	if d != None and static == 1:
	    d.set_static(1)

	if d != None and name != None and type != None:
	    self.references[name] = d

	if name == debugsym:
	    print "New ref: %s" % (d)

	return d

    def add(self, name, header, module, static, type, lineno, info=None, extra=None, conditionals = None):
        if name[0:2] == '__':
	    return None
        d = None
        try:
	   d = self.identifiers[name]
	   d.update(header, module, type, lineno, info, extra, conditionals)
	except:
	   d = identifier(name, header, module, type, lineno, info, extra, conditionals)
	   self.identifiers[name] = d

	if d != None and static == 1:
	    d.set_static(1)

	if d != None and name != None and type != None:
	    if type == "function":
	        self.functions[name] = d
	    elif type == "functype":
	        self.functions[name] = d
	    elif type == "variable":
	        self.variables[name] = d
	    elif type == "include":
	        self.includes[name] = d
	    elif type == "struct":
	        self.structs[name] = d
	    elif type == "enum":
	        self.enums[name] = d
	    elif type == "typedef":
	        self.typedefs[name] = d
	    elif type == "macro":
	        self.macros[name] = d
	    else:
	        print "Unable to register type ", type

	if name == debugsym:
	    print "New symbol: %s" % (d)

	return d

    def merge(self, idx):
        for id in idx.functions.keys():
              #
              # macro might be used to override functions or variables
              # definitions
              #
	     if self.macros.has_key(id):
	         del self.macros[id]
	     if self.functions.has_key(id):
	         print "function %s from %s redeclared in %s" % (
		    id, self.functions[id].header, idx.functions[id].header)
	     else:
	         self.functions[id] = idx.functions[id]
		 self.identifiers[id] = idx.functions[id]
        for id in idx.variables.keys():
              #
              # macro might be used to override functions or variables
              # definitions
              #
	     if self.macros.has_key(id):
	         del self.macros[id]
	     if self.variables.has_key(id):
	         print "variable %s from %s redeclared in %s" % (
		    id, self.variables[id].header, idx.variables[id].header)
	     else:
	         self.variables[id] = idx.variables[id]
		 self.identifiers[id] = idx.variables[id]
        for id in idx.structs.keys():
	     if self.structs.has_key(id):
	         print "struct %s from %s redeclared in %s" % (
		    id, self.structs[id].header, idx.structs[id].header)
	     else:
	         self.structs[id] = idx.structs[id]
		 self.identifiers[id] = idx.structs[id]
        for id in idx.typedefs.keys():
	     if self.typedefs.has_key(id):
	         print "typedef %s from %s redeclared in %s" % (
		    id, self.typedefs[id].header, idx.typedefs[id].header)
	     else:
	         self.typedefs[id] = idx.typedefs[id]
		 self.identifiers[id] = idx.typedefs[id]
        for id in idx.macros.keys():
              #
              # macro might be used to override functions or variables
              # definitions
              #
             if self.variables.has_key(id):
                 continue
             if self.functions.has_key(id):
                 continue
             if self.enums.has_key(id):
                 continue
	     if self.macros.has_key(id):
	         print "macro %s from %s redeclared in %s" % (
		    id, self.macros[id].header, idx.macros[id].header)
	     else:
	         self.macros[id] = idx.macros[id]
		 self.identifiers[id] = idx.macros[id]
        for id in idx.enums.keys():
	     if self.enums.has_key(id):
	         print "enum %s from %s redeclared in %s" % (
		    id, self.enums[id].header, idx.enums[id].header)
	     else:
	         self.enums[id] = idx.enums[id]
		 self.identifiers[id] = idx.enums[id]

    def merge_public(self, idx):
        for id in idx.functions.keys():
	     if self.functions.has_key(id):
	         # check that function condition agrees with header
	         if idx.functions[id].conditionals != \
		    self.functions[id].conditionals:
		     print "Header condition differs from Function for %s:" \
		        % id
		     print "  H: %s" % self.functions[id].conditionals
		     print "  C: %s" % idx.functions[id].conditionals
	         up = idx.functions[id]
	         self.functions[id].update(None, up.module, up.type, up.info, up.extra)
	 #     else:
	 #         print "Function %s from %s is not declared in headers" % (
	 #	        id, idx.functions[id].module)
	 # TODO: do the same for variables.

    def analyze_dict(self, type, dict):
        count = 0
	public = 0
        for name in dict.keys():
	    id = dict[name]
	    count = count + 1
	    if id.static == 0:
	        public = public + 1
        if count != public:
	    print "  %d %s , %d public" % (count, type, public)
	elif count != 0:
	    print "  %d public %s" % (count, type)


    def analyze(self):
	self.analyze_dict("functions", self.functions)
	self.analyze_dict("variables", self.variables)
	self.analyze_dict("structs", self.structs)
	self.analyze_dict("typedefs", self.typedefs)
	self.analyze_dict("macros", self.macros)

class CLexer:
    """A lexer for the C language, tokenize the input by reading and
       analyzing it line by line"""
    def __init__(self, input):
        self.input = input
	self.tokens = []
	self.line = ""
	self.lineno = 0

    def getline(self):
        line = ''
	while line == '':
	    line = self.input.readline()
	    if not line:
		return None
	    self.lineno = self.lineno + 1
	    line = string.lstrip(line)
	    line = string.rstrip(line)
	    if line == '':
	        continue
	    while line[-1] == '\\':
	        line = line[:-1]
		n = self.input.readline()
		self.lineno = self.lineno + 1
		n = string.lstrip(n)
		n = string.rstrip(n)
		if not n:
		    break
		else:
		    line = line + n
        return line

    def getlineno(self):
        return self.lineno

    def push(self, token):
        self.tokens.insert(0, token);

    def debug(self):
        print "Last token: ", self.last
	print "Token queue: ", self.tokens
	print "Line %d end: " % (self.lineno), self.line

    def token(self):
        while self.tokens == []:
	    if self.line == "":
		line = self.getline()
	    else:
	        line = self.line
		self.line = ""
	    if line == None:
	        return None

	    if line[0] == '#':
	        self.tokens = map((lambda x: ('preproc', x)),
		                  string.split(line))
		break;
	    l = len(line)
	    if line[0] == '"' or line[0] == "'":
	        end = line[0]
	        line = line[1:]
		found = 0
		tok = ""
		while found == 0:
		    i = 0
		    l = len(line)
		    while i < l:
			if line[i] == end:
			    self.line = line[i+1:]
			    line = line[:i]
			    l = i
			    found = 1
			    break
			if line[i] == '\\':
			    i = i + 1
			i = i + 1
		    tok = tok + line
		    if found == 0:
		        line = self.getline()
			if line == None:
			    return None
		self.last = ('string', tok)
		return self.last

	    if l >= 2 and line[0] == '/' and line[1] == '*':
	        line = line[2:]
		found = 0
		tok = ""
		while found == 0:
		    i = 0
		    l = len(line)
		    while i < l:
			if line[i] == '*' and i+1 < l and line[i+1] == '/':
			    self.line = line[i+2:]
			    line = line[:i-1]
			    l = i
			    found = 1
			    break
			i = i + 1
	            if tok != "":
		        tok = tok + "\n"
		    tok = tok + line
		    if found == 0:
		        line = self.getline()
			if line == None:
			    return None
		self.last = ('comment', tok)
		return self.last
	    if l >= 2 and line[0] == '/' and line[1] == '/':
	        line = line[2:]
		self.last = ('comment', line)
		return self.last
	    i = 0
	    while i < l:
	        if line[i] == '/' and i+1 < l and line[i+1] == '/':
		    self.line = line[i:]
		    line = line[:i]
		    break
	        if line[i] == '/' and i+1 < l and line[i+1] == '*':
		    self.line = line[i:]
		    line = line[:i]
		    break
		if line[i] == '"' or line[i] == "'":
		    self.line = line[i:]
		    line = line[:i]
		    break
		i = i + 1
	    l = len(line)
	    i = 0
	    while i < l:
	        if line[i] == ' ' or line[i] == '\t':
		    i = i + 1
		    continue
		o = ord(line[i])
		if (o >= 97 and o <= 122) or (o >= 65 and o <= 90) or \
		   (o >= 48 and o <= 57):
		    s = i
		    while i < l:
			o = ord(line[i])
			if (o >= 97 and o <= 122) or (o >= 65 and o <= 90) or \
			   (o >= 48 and o <= 57) or string.find(
			       " \t(){}:;,+-*/%&!|[]=><", line[i]) == -1:
			    i = i + 1
			else:
			    break
		    self.tokens.append(('name', line[s:i]))
		    continue
		if string.find("(){}:;,[]", line[i]) != -1:
#                 if line[i] == '(' or line[i] == ')' or line[i] == '{' or \
#		    line[i] == '}' or line[i] == ':' or line[i] == ';' or \
#		    line[i] == ',' or line[i] == '[' or line[i] == ']':
		    self.tokens.append(('sep', line[i]))
		    i = i + 1
		    continue
		if string.find("+-*><=/%&!|.", line[i]) != -1:
#                 if line[i] == '+' or line[i] == '-' or line[i] == '*' or \
#		    line[i] == '>' or line[i] == '<' or line[i] == '=' or \
#		    line[i] == '/' or line[i] == '%' or line[i] == '&' or \
#		    line[i] == '!' or line[i] == '|' or line[i] == '.':
		    if line[i] == '.' and  i + 2 < l and \
		       line[i+1] == '.' and line[i+2] == '.':
			self.tokens.append(('name', '...'))
			i = i + 3
			continue

		    j = i + 1
		    if j < l and (
		       string.find("+-*><=/%&!|", line[j]) != -1):
#		        line[j] == '+' or line[j] == '-' or line[j] == '*' or \
#			line[j] == '>' or line[j] == '<' or line[j] == '=' or \
#			line[j] == '/' or line[j] == '%' or line[j] == '&' or \
#			line[j] == '!' or line[j] == '|'):
			self.tokens.append(('op', line[i:j+1]))
			i = j + 1
		    else:
			self.tokens.append(('op', line[i]))
			i = i + 1
		    continue
		s = i
		while i < l:
		    o = ord(line[i])
		    if (o >= 97 and o <= 122) or (o >= 65 and o <= 90) or \
		       (o >= 48 and o <= 57) or (
		        string.find(" \t(){}:;,+-*/%&!|[]=><", line[i]) == -1):
#		         line[i] != ' ' and line[i] != '\t' and
#			 line[i] != '(' and line[i] != ')' and
#			 line[i] != '{'  and line[i] != '}' and
#			 line[i] != ':' and line[i] != ';' and
#			 line[i] != ',' and line[i] != '+' and
#			 line[i] != '-' and line[i] != '*' and
#			 line[i] != '/' and line[i] != '%' and
#			 line[i] != '&' and line[i] != '!' and
#			 line[i] != '|' and line[i] != '[' and
#			 line[i] != ']' and line[i] != '=' and
#			 line[i] != '*' and line[i] != '>' and
#			 line[i] != '<'):
			i = i + 1
		    else:
		        break
		self.tokens.append(('name', line[s:i]))

	tok = self.tokens[0]
	self.tokens = self.tokens[1:]
	self.last = tok
	return tok

class CParser:
    """The C module parser"""
    def __init__(self, filename, idx = None):
    	self.filename = filename
        
	
	if len(filename) > 2 and filename[-2:] == '.h':
	    self.is_header = 1
	else:
	    self.is_header = 0
        self.input = open(filename)
	self.lexer = CLexer(self.input)
	if idx == None:
	    self.index = index()
	else:
	    self.index = idx
	self.top_comment = ""
	self.last_comment = ""
	self.comment = None
	self.collect_ref = 0
	self.no_error = 0
	self.conditionals = []
	self.defines = []
	
	
    def collect_references(self):
        self.collect_ref = 1

    def stop_error(self):
        self.no_error = 1

    def start_error(self):
        self.no_error = 0

    def lineno(self):
        return self.lexer.getlineno()

    def index_add(self, name, module, static, type, info=None, extra = None):
	if self.is_header == 1:
	    self.index.add(name, module, module, static, type, self.lineno(),
			   info, extra, self.conditionals)
	else:
	    self.index.add(name, None, module, static, type, self.lineno(),
			   info, extra, self.conditionals)

    def index_add_ref(self, name, module, static, type, info=None,
                      extra = None):
	if self.is_header == 1:
	    self.index.add_ref(name, module, module, static, type,
	                       self.lineno(), info, extra, self.conditionals)
	else:
	    self.index.add_ref(name, None, module, static, type, self.lineno(),
			       info, extra, self.conditionals)

    def warning(self, msg):
        if self.no_error:
	    return
	print msg

    def error(self, msg, token=-1):
        if self.no_error:
	    return

        print "Parse Error: " + msg
	if token != -1:
	    print "Got token ", token
	self.lexer.debug()
	sys.exit(1)

    def debug(self, msg, token=-1):
        print "Debug: " + msg
	if token != -1:
	    print "Got token ", token
	self.lexer.debug()

    def parseTopComment(self, comment):
	res = {}
	lines = string.split(comment, "\n")
	item = None
	for line in lines:
	    while line != "" and (line[0] == ' ' or line[0] == '\t'):
		line = line[1:]
	    while line != "" and line[0] == '*':
		line = line[1:]
	    while line != "" and (line[0] == ' ' or line[0] == '\t'):
		line = line[1:]
	    try:
		(it, line) = string.split(line, ":", 1)
		item = it
		while line != "" and (line[0] == ' ' or line[0] == '\t'):
		    line = line[1:]
		if res.has_key(item):
		    res[item] = res[item] + " " + line
		else:
		    res[item] = line
	    except:
		if item != None:
		    if res.has_key(item):
			res[item] = res[item] + " " + line
		    else:
			res[item] = line
	self.index.info = res

    def parseComment(self, token):
        if self.top_comment == "":
	    self.top_comment = token[1]
	if self.comment == None or token[1][0] == '*':
	    self.comment = token[1];
	else:
	    self.comment = self.comment + token[1]
	token = self.lexer.token()

        if string.find(self.comment, "DOC_DISABLE") != -1:
	    self.stop_error()

        if string.find(self.comment, "DOC_ENABLE") != -1:
	    self.start_error()

	return token

    #
    # Parse a comment block associate to a typedef
    #
    def parseTypeComment(self, name, quiet = 0):
        if name[0:2] == '__':
	    quiet = 1

        args = []
	desc = ""

        if self.comment == None:
	    if not quiet:
		self.warning("Missing comment for type %s" % (name))
	    return((args, desc))
        if self.comment[0] != '*':
	    if not quiet:
		self.warning("Missing * in type comment for %s" % (name))
	    return((args, desc))
	lines = string.split(self.comment, '\n')
	if lines[0] == '*':
	    del lines[0]
	if lines[0] != "* %s:" % (name):
	    if not quiet:
		self.warning("Misformatted type comment for %s" % (name))
		self.warning("  Expecting '* %s:' got '%s'" % (name, lines[0]))
	    return((args, desc))
	del lines[0]
	while len(lines) > 0 and lines[0] == '*':
	    del lines[0]
	desc = ""
	while len(lines) > 0:
	    l = lines[0]
	    while len(l) > 0 and l[0] == '*':
	        l = l[1:]
	    l = string.strip(l)
	    desc = desc + " " + l
	    del lines[0]

	desc = string.strip(desc)

	if quiet == 0:
	    if desc == "":
	        self.warning("Type comment for %s lack description of the macro" % (name))

	return(desc)
    #
    # Parse a comment block associate to a macro
    #
    def parseMacroComment(self, name, quiet = 0):
        if name[0:2] == '__':
	    quiet = 1

        args = []
	desc = ""

        if self.comment == None:
	    if not quiet:
		self.warning("Missing comment for macro %s" % (name))
	    return((args, desc))
        if self.comment[0] != '*':
	    if not quiet:
		self.warning("Missing * in macro comment for %s" % (name))
	    return((args, desc))
	lines = string.split(self.comment, '\n')
	if lines[0] == '*':
	    del lines[0]
	if lines[0] != "* %s:" % (name):
	    if not quiet:
		self.warning("Misformatted macro comment for %s" % (name))
		self.warning("  Expecting '* %s:' got '%s'" % (name, lines[0]))
	    return((args, desc))
	del lines[0]
	while lines[0] == '*':
	    del lines[0]
	while len(lines) > 0 and lines[0][0:3] == '* @':
	    l = lines[0][3:]
	    try:
	        (arg, desc) = string.split(l, ':', 1)
		desc=string.strip(desc)
		arg=string.strip(arg)
            except:
		if not quiet:
		    self.warning("Misformatted macro comment for %s" % (name))
		    self.warning("  problem with '%s'" % (lines[0]))
		del lines[0]
		continue
	    del lines[0]
	    l = string.strip(lines[0])
	    while len(l) > 2 and l[0:3] != '* @':
	        while l[0] == '*':
		    l = l[1:]
		desc = desc + ' ' + string.strip(l)
		del lines[0]
		if len(lines) == 0:
		    break
		l = lines[0]
            args.append((arg, desc))
	while len(lines) > 0 and lines[0] == '*':
	    del lines[0]
	desc = ""
	while len(lines) > 0:
	    l = lines[0]
	    while len(l) > 0 and l[0] == '*':
	        l = l[1:]
	    l = string.strip(l)
	    desc = desc + " " + l
	    del lines[0]

	desc = string.strip(desc)

	if quiet == 0:
	    if desc == "":
	        self.warning("Macro comment for %s lack description of the macro" % (name))

	return((args, desc))

     #
     # Parse a comment block and merge the informations found in the
     # parameters descriptions, finally returns a block as complete
     # as possible
     #
    def mergeFunctionComment(self, name, description, quiet = 0):
        if name == 'main':
	    quiet = 1
        if name[0:2] == '__':
	    quiet = 1

	(ret, args) = description
	desc = ""
	retdesc = ""

        if self.comment == None:
	    if not quiet:
		self.warning("Missing comment for function %s" % (name))
	    return(((ret[0], retdesc), args, desc))
        if self.comment[0] != '*':
	    if not quiet:
		self.warning("Missing * in function comment for %s" % (name))
	    return(((ret[0], retdesc), args, desc))
	lines = string.split(self.comment, '\n')
	if lines[0] == '*':
	    del lines[0]
	if lines[0] != "* %s:" % (name):
	    if not quiet:
		self.warning("Misformatted function comment for %s" % (name))
		self.warning("  Expecting '* %s:' got '%s'" % (name, lines[0]))
	    return(((ret[0], retdesc), args, desc))
	del lines[0]
	while lines[0] == '*':
	    del lines[0]
	nbargs = len(args)
	while len(lines) > 0 and lines[0][0:3] == '* @':
	    l = lines[0][3:]
	    try:
	        (arg, desc) = string.split(l, ':', 1)
		desc=string.strip(desc)
		arg=string.strip(arg)
            except:
		if not quiet:
		    self.warning("Misformatted function comment for %s" % (name))
		    self.warning("  problem with '%s'" % (lines[0]))
		del lines[0]
		continue
	    del lines[0]
	    l = string.strip(lines[0])
	    while len(l) > 2 and l[0:3] != '* @':
	        while l[0] == '*':
		    l = l[1:]
		desc = desc + ' ' + string.strip(l)
		del lines[0]
		if len(lines) == 0:
		    break
		l = lines[0]
	    i = 0
	    while i < nbargs:
	        if args[i][1] == arg:
		    args[i] = (args[i][0], arg, desc)
		    break;
		i = i + 1
	    if i >= nbargs:
		if not quiet:
		    self.warning("Unable to find arg %s from function comment for %s" % (
		       arg, name))
	while len(lines) > 0 and lines[0] == '*':
	    del lines[0]
	desc = ""
	while len(lines) > 0:
	    l = lines[0]
	    while len(l) > 0 and l[0] == '*':
	        l = l[1:]
	    l = string.strip(l)
	    if len(l) >= 6 and  l[0:6] == "return" or l[0:6] == "Return":
	        try:
		    l = string.split(l, ' ', 1)[1]
		except:
		    l = ""
		retdesc = string.strip(l)
		del lines[0]
		while len(lines) > 0:
		    l = lines[0]
		    while len(l) > 0 and l[0] == '*':
			l = l[1:]
		    l = string.strip(l)
		    retdesc = retdesc + " " + l
		    del lines[0]
	    else:
	        desc = desc + " " + l
		del lines[0]

	retdesc = string.strip(retdesc)
	desc = string.strip(desc)

	if quiet == 0:
	     #
	     # report missing comments
	     #
	    i = 0
	    while i < nbargs:
	        if args[i][2] == None and args[i][0] != "void" and \
		   ((args[i][1] != None) or (args[i][1] == '')):
		    self.warning("Function comment for %s lacks description of arg %s" % (name, args[i][1]))
		i = i + 1
	    if retdesc == "" and ret[0] != "void":
		self.warning("Function comment for %s lacks description of return value" % (name))
	    if desc == "":
	        self.warning("Function comment for %s lacks description of the function" % (name))

	return(((ret[0], retdesc), args, desc))

    def parsePreproc(self, token):
	if debug:
	    print "=> preproc ", token, self.lexer.tokens
        name = token[1]
	if name == "#include":
	    token = self.lexer.token()
	    if token == None:
	        return None
	    if token[0] == 'preproc':
		self.index_add(token[1], self.filename, not self.is_header,
		                "include")
		return self.lexer.token()
	    return token
	if name == "#define":
	    token = self.lexer.token()
	    if token == None:
	        return None
	    if token[0] == 'preproc':
	         # TODO macros with arguments
		name = token[1]
	        lst = []
		token = self.lexer.token()
		while token != None and token[0] == 'preproc' and \
		      token[1][0] != '#':
		    lst.append(token[1])
		    token = self.lexer.token()
                try:
		    name = string.split(name, '(') [0]
                except:
                    pass
                info = self.parseMacroComment(name, not self.is_header)
		self.index_add(name, self.filename, not self.is_header,
		                "macro", info)
		return token

	#
	# Processing of conditionals modified by Bill 1/1/05
	#
	# We process conditionals (i.e. tokens from #ifdef, #ifndef,
	# #if, #else and #endif) for headers and mainline code,
	# store the ones from the header in libxml2-api.xml, and later
	# (in the routine merge_public) verify that the two (header and
	# mainline code) agree.
	#
	# There is a small problem with processing the headers. Some of
	# the variables are not concerned with enabling / disabling of
	# library functions (e.g. '__XML_PARSER_H__'), and we don't want
	# them to be included in libxml2-api.xml, or involved in
	# the check between the header and the mainline code.  To
	# accomplish this, we ignore any conditional which doesn't include
	# the string 'ENABLED'
	#
	if name == "#ifdef":
	    apstr = self.lexer.tokens[0][1]
	    try:
	        self.defines.append(apstr)
		if string.find(apstr, 'ENABLED') != -1:
		    self.conditionals.append("defined(%s)" % apstr)
	    except:
	        pass
	elif name == "#ifndef":
	    apstr = self.lexer.tokens[0][1]
	    try:
	        self.defines.append(apstr)
		if string.find(apstr, 'ENABLED') != -1:
		    self.conditionals.append("!defined(%s)" % apstr)
	    except:
	        pass
	elif name == "#if":
	    apstr = ""
	    for tok in self.lexer.tokens:
	        if apstr != "":
		    apstr = apstr + " "
	        apstr = apstr + tok[1]
	    try:
	        self.defines.append(apstr)
		if string.find(apstr, 'ENABLED') != -1:
		    self.conditionals.append(apstr)
	    except:
	        pass
	elif name == "#else":
	    if self.conditionals != [] and \
	       string.find(self.defines[-1], 'ENABLED') != -1:
	        self.conditionals[-1] = "!(%s)" % self.conditionals[-1]
	elif name == "#endif":
	    if self.conditionals != [] and \
	       string.find(self.defines[-1], 'ENABLED') != -1:
	        self.conditionals = self.conditionals[:-1]
	    self.defines = self.defines[:-1]
	token = self.lexer.token()
	while token != None and token[0] == 'preproc' and \
	    token[1][0] != '#':
	    token = self.lexer.token()
	return token

     #
     # token acquisition on top of the lexer, it handle internally
     # preprocessor and comments since they are logically not part of
     # the program structure.
     #
    def token(self):
        global ignored_words

        token = self.lexer.token()
	while token != None:
	    if token[0] == 'comment':
		token = self.parseComment(token)
		continue
	    elif token[0] == 'preproc':
		token = self.parsePreproc(token)
		continue
	    elif token[0] == "name" and token[1] == "__const":
	        token = ("name", "const")
		return token
	    elif token[0] == "name" and token[1] == "__attribute":
		token = self.lexer.token()
		while token != None and token[1] != ";":
		    token = self.lexer.token()
		return token
	    elif token[0] == "name" and ignored_words.has_key(token[1]):
	        (n, info) = ignored_words[token[1]]
		i = 0
		while i < n:
		    token = self.lexer.token()
		    i = i + 1
		token = self.lexer.token()
		continue
	    else:
	        if debug:
		    print "=> ", token
	        return token
	return None

     #
     # Parse a typedef, it records the type and its name.
     #
    def parseTypedef(self, token):
        if token == None:
	    return None
	token = self.parseType(token)
	if token == None:
	    self.error("parsing typedef")
	    return None
	base_type = self.type
	type = base_type
	 #self.debug("end typedef type", token)
	while token != None:
	    if token[0] == "name":
		name = token[1]
		signature = self.signature
		if signature != None:
		    type = string.split(type, '(')[0]
		    d = self.mergeFunctionComment(name,
			    ((type, None), signature), 1)
		    self.index_add(name, self.filename, not self.is_header,
				    "functype", d)
		else:
		    if base_type == "struct":
			self.index_add(name, self.filename, not self.is_header,
					"struct", type)
			base_type = "struct " + name
	            else:
			# TODO report missing or misformatted comments
			info = self.parseTypeComment(name, 1)
			self.index_add(name, self.filename, not self.is_header,
		                    "typedef", type, info)
		token = self.token()
	    else:
		self.error("parsing typedef: expecting a name")
		return token
	     #self.debug("end typedef", token)
	    if token != None and token[0] == 'sep' and token[1] == ',':
	        type = base_type
	        token = self.token()
		while token != None and token[0] == "op":
		    type = type + token[1]
		    token = self.token()
	    elif token != None and token[0] == 'sep' and token[1] == ';':
	        break;
	    elif token != None and token[0] == 'name':
	        type = base_type
	        continue;
	    else:
		self.error("parsing typedef: expecting ';'", token)
		return token
	token = self.token()
	return token

     #
     # Parse a C code block, used for functions it parse till
     # the balancing } included
     #
    def parseBlock(self, token):
        while token != None:
	    if token[0] == "sep" and token[1] == "{":
	        token = self.token()
		token = self.parseBlock(token)
	    elif token[0] == "sep" and token[1] == "}":
	        self.comment = None
	        token = self.token()
		return token
	    else:
	        if self.collect_ref == 1:
		    oldtok = token
		    token = self.token()
		    if oldtok[0] == "name" and oldtok[1][0:3] == "xml":
		        if token[0] == "sep" and token[1] == "(":
			    self.index_add_ref(oldtok[1], self.filename,
			                        0, "function")
			    token = self.token()
			elif token[0] == "name":
			    token = self.token()
			    if token[0] == "sep" and (token[1] == ";" or
			       token[1] == "," or token[1] == "="):
				self.index_add_ref(oldtok[1], self.filename,
						    0, "type")
		    elif oldtok[0] == "name" and oldtok[1][0:4] == "XML_":
			self.index_add_ref(oldtok[1], self.filename,
					    0, "typedef")
		    elif oldtok[0] == "name" and oldtok[1][0:7] == "LIBXML_":
			self.index_add_ref(oldtok[1], self.filename,
					    0, "typedef")

		else:
		    token = self.token()
	return token

     #
     # Parse a C struct definition till the balancing }
     #
    def parseStruct(self, token):
        fields = []
	 #self.debug("start parseStruct", token)
        while token != None:
	    if token[0] == "sep" and token[1] == "{":
	        token = self.token()
		token = self.parseTypeBlock(token)
	    elif token[0] == "sep" and token[1] == "}":
		self.struct_fields = fields
		 #self.debug("end parseStruct", token)
		 #print fields
	        token = self.token()
		return token
	    else:
	        base_type = self.type
		 #self.debug("before parseType", token)
		token = self.parseType(token)
		 #self.debug("after parseType", token)
		if token != None and token[0] == "name":
		    fname = token[1]
		    token = self.token()
		    if token[0] == "sep" and token[1] == ";":
		        self.comment = None
		        token = self.token()
			fields.append((self.type, fname, self.comment))
			self.comment = None
		    else:
		        self.error("parseStruct: expecting ;", token)
		elif token != None and token[0] == "sep" and token[1] == "{":
		    token = self.token()
		    token = self.parseTypeBlock(token)
		    if token != None and token[0] == "name":
			token = self.token()
		    if token != None and token[0] == "sep" and token[1] == ";":
			token = self.token()
		    else:
		        self.error("parseStruct: expecting ;", token)
		else:
		    self.error("parseStruct: name", token)
		    token = self.token()
		self.type = base_type;
        self.struct_fields = fields
	 #self.debug("end parseStruct", token)
	 #print fields
	return token

     #
     # Parse a C enum block, parse till the balancing }
     #
    def parseEnumBlock(self, token):
        self.enums = []
	name = None
	self.comment = None
	comment = ""
	value = "0"
        while token != None:
	    if token[0] == "sep" and token[1] == "{":
	        token = self.token()
		token = self.parseTypeBlock(token)
	    elif token[0] == "sep" and token[1] == "}":
		if name != None:
		    if self.comment != None:
			comment = self.comment
			self.comment = None
		    self.enums.append((name, value, comment))
	        token = self.token()
		return token
	    elif token[0] == "name":
		    if name != None:
			if self.comment != None:
			    comment = string.strip(self.comment)
			    self.comment = None
			self.enums.append((name, value, comment))
		    name = token[1]
		    comment = ""
		    token = self.token()
		    if token[0] == "op" and token[1][0] == "=":
		        value = ""
		        if len(token[1]) > 1:
			    value = token[1][1:]
		        token = self.token()
		        while token[0] != "sep" or (token[1] != ',' and
			      token[1] != '}'):
			    value = value + token[1]
			    token = self.token()
		    else:
		        try:
			    value = "%d" % (int(value) + 1)
			except:
			    self.warning("Failed to compute value of enum %s" % (name))
			    value=""
		    if token[0] == "sep" and token[1] == ",":
			token = self.token()
	    else:
	        token = self.token()
	return token

     #
     # Parse a C definition block, used for structs it parse till
     # the balancing }
     #
    def parseTypeBlock(self, token):
        while token != None:
	    if token[0] == "sep" and token[1] == "{":
	        token = self.token()
		token = self.parseTypeBlock(token)
	    elif token[0] == "sep" and token[1] == "}":
	        token = self.token()
		return token
	    else:
	        token = self.token()
	return token

     #
     # Parse a type: the fact that the type name can either occur after
     #    the definition or within the definition makes it a little harder
     #    if inside, the name token is pushed back before returning
     #
    def parseType(self, token):
        self.type = ""
	self.struct_fields = []
        self.signature = None
	if token == None:
	    return token

	while token[0] == "name" and (
	      token[1] == "const" or \
	      token[1] == "unsigned" or \
	      token[1] == "signed"):
	    if self.type == "":
	        self.type = token[1]
	    else:
	        self.type = self.type + " " + token[1]
	    token = self.token()

        if token[0] == "name" and (token[1] == "long" or token[1] == "short"):
	    if self.type == "":
	        self.type = token[1]
	    else:
	        self.type = self.type + " " + token[1]
	    if token[0] == "name" and token[1] == "int":
		if self.type == "":
		    self.type = tmp[1]
		else:
		    self.type = self.type + " " + tmp[1]

        elif token[0] == "name" and token[1] == "struct":
	    if self.type == "":
	        self.type = token[1]
	    else:
	        self.type = self.type + " " + token[1]
	    token = self.token()
	    nametok = None
	    if token[0] == "name":
	        nametok = token
		token = self.token()
	    if token != None and token[0] == "sep" and token[1] == "{":
		token = self.token()
		token = self.parseStruct(token)
	    elif token != None and token[0] == "op" and token[1] == "*":
	        self.type = self.type + " " + nametok[1] + " *"
		token = self.token()
		while token != None and token[0] == "op" and token[1] == "*":
		    self.type = self.type + " *"
		    token = self.token()
		if token[0] == "name":
		    nametok = token
		    token = self.token()
		else:
		    self.error("struct : expecting name", token)
		    return token
	    elif token != None and token[0] == "name" and nametok != None:
	        self.type = self.type + " " + nametok[1]
		return token

	    if nametok != None:
		self.lexer.push(token)
		token = nametok
	    return token

        elif token[0] == "name" and token[1] == "enum":
	    if self.type == "":
	        self.type = token[1]
	    else:
	        self.type = self.type + " " + token[1]
	    self.enums = []
	    token = self.token()
	    if token != None and token[0] == "sep" and token[1] == "{":
		token = self.token()
		token = self.parseEnumBlock(token)
	    else:
		self.error("parsing enum: expecting '{'", token)
	    enum_type = None
	    if token != None and token[0] != "name":
	        self.lexer.push(token)
	        token = ("name", "enum")
	    else:
	        enum_type = token[1]
	    for enum in self.enums:
		self.index_add(enum[0], self.filename,
			       not self.is_header, "enum",
			       (enum[1], enum[2], enum_type))
	    return token

	elif token[0] == "name":
	    if self.type == "":
	        self.type = token[1]
	    else:
	        self.type = self.type + " " + token[1]
	else:
	    self.error("parsing type %s: expecting a name" % (self.type),
	               token)
	    return token
	token = self.token()
        while token != None and (token[0] == "op" or
	      token[0] == "name" and token[1] == "const"):
	    self.type = self.type + " " + token[1]
	    token = self.token()

	 #
	 # if there is a parenthesis here, this means a function type
	 #
	if token != None and token[0] == "sep" and token[1] == '(':
	    self.type = self.type + token[1]
	    token = self.token()
	    while token != None and token[0] == "op" and token[1] == '*':
	        self.type = self.type + token[1]
		token = self.token()
	    if token == None or token[0] != "name" :
		self.error("parsing function type, name expected", token);
	        return token
	    self.type = self.type + token[1]
	    nametok = token
	    token = self.token()
	    if token != None and token[0] == "sep" and token[1] == ')':
		self.type = self.type + token[1]
		token = self.token()
		if token != None and token[0] == "sep" and token[1] == '(':
		    token = self.token()
		    type = self.type;
		    token = self.parseSignature(token);
		    self.type = type;
		else:
		    self.error("parsing function type, '(' expected", token);
		    return token
	    else:
	        self.error("parsing function type, ')' expected", token);
		return token
	    self.lexer.push(token)
	    token = nametok
	    return token

         #
	 # do some lookahead for arrays
	 #
	if token != None and token[0] == "name":
	    nametok = token
	    token = self.token()
	    if token != None and token[0] == "sep" and token[1] == '[':
	        self.type = self.type + nametok[1]
		while token != None and token[0] == "sep" and token[1] == '[':
		    self.type = self.type + token[1]
		    token = self.token()
		    while token != None and token[0] != 'sep' and \
		          token[1] != ']' and token[1] != ';':
			self.type = self.type + token[1]
			token = self.token()
		if token != None and token[0] == 'sep' and token[1] == ']':
		    self.type = self.type + token[1]
		    token = self.token()
		else:
		    self.error("parsing array type, ']' expected", token);
		    return token
	    elif token != None and token[0] == "sep" and token[1] == ':':
	         # remove :12 in case it's a limited int size
		token = self.token()
		token = self.token()
	    self.lexer.push(token)
	    token = nametok

	return token

     #
     # Parse a signature: '(' has been parsed and we scan the type definition
     #    up to the ')' included
    def parseSignature(self, token):
        signature = []
	if token != None and token[0] == "sep" and token[1] == ')':
	    self.signature = []
	    token = self.token()
	    return token
	while token != None:
	    token = self.parseType(token)
	    if token != None and token[0] == "name":
	        signature.append((self.type, token[1], None))
		token = self.token()
	    elif token != None and token[0] == "sep" and token[1] == ',':
		token = self.token()
		continue
	    elif token != None and token[0] == "sep" and token[1] == ')':
	         # only the type was provided
		if self.type == "...":
		    signature.append((self.type, "...", None))
		else:
		    signature.append((self.type, None, None))
	    if token != None and token[0] == "sep":
	        if token[1] == ',':
		    token = self.token()
		    continue
		elif token[1] == ')':
		    token = self.token()
		    break
	self.signature = signature
	return token

     #
     # Parse a global definition, be it a type, variable or function
     # the extern "C" blocks are a bit nasty and require it to recurse.
     #
    def parseGlobal(self, token):
        static = 0
        if token[1] == 'extern':
	    token = self.token()
	    if token == None:
	        return token
	    if token[0] == 'string':
	        if token[1] == 'C':
		    token = self.token()
		    if token == None:
			return token
		    if token[0] == 'sep' and token[1] == "{":
		        token = self.token()
#			 print 'Entering extern "C line ', self.lineno()
			while token != None and (token[0] != 'sep' or
			      token[1] != "}"):
			    if token[0] == 'name':
				token = self.parseGlobal(token)
			    else:
				self.error(
				 "token %s %s unexpected at the top level" % (
					token[0], token[1]))
				token = self.parseGlobal(token)
#			 print 'Exiting extern "C" line', self.lineno()
			token = self.token()
			return token
		else:
		    return token
	elif token[1] == 'static':
	    static = 1
	    token = self.token()
	    if token == None or  token[0] != 'name':
	        return token

	if token[1] == 'typedef':
	    token = self.token()
	    return self.parseTypedef(token)
	else:
	    token = self.parseType(token)
	    type_orig = self.type
	if token == None or token[0] != "name":
	    return token
	type = type_orig
	self.name = token[1]
	token = self.token()
	while token != None and (token[0] == "sep" or token[0] == "op"):
	    if token[0] == "sep":
		if token[1] == "[":
		    type = type + token[1]
		    token = self.token()
		    while token != None and (token[0] != "sep" or \
		          token[1] != ";"):
			type = type + token[1]
			token = self.token()

	    if token != None and token[0] == "op" and token[1] == "=":
		 #
		 # Skip the initialization of the variable
		 #
		token = self.token()
		if token[0] == 'sep' and token[1] == '{':
		    token = self.token()
		    token = self.parseBlock(token)
		else:
		    self.comment = None
		    while token != None and (token[0] != "sep" or \
			  (token[1] != ';' and token[1] != ',')):
			    token = self.token()
		self.comment = None
		if token == None or token[0] != "sep" or (token[1] != ';' and
		   token[1] != ','):
		    self.error("missing ';' or ',' after value")

	    if token != None and token[0] == "sep":
		if token[1] == ";":
		    self.comment = None
		    token = self.token()
		    if type == "struct":
		        self.index_add(self.name, self.filename,
			     not self.is_header, "struct", self.struct_fields)
		    else:
			self.index_add(self.name, self.filename,
			     not self.is_header, "variable", type)
		    break
		elif token[1] == "(":
		    token = self.token()
		    token = self.parseSignature(token)
		    if token == None:
			return None
		    if token[0] == "sep" and token[1] == ";":
		        d = self.mergeFunctionComment(self.name,
				((type, None), self.signature), 1)
			self.index_add(self.name, self.filename, static,
			                "function", d)
			token = self.token()
		    elif token[0] == "sep" and token[1] == "{":
		        d = self.mergeFunctionComment(self.name,
				((type, None), self.signature), static)
			self.index_add(self.name, self.filename, static,
			                "function", d)
			token = self.token()
			token = self.parseBlock(token);
		elif token[1] == ',':
		    self.comment = None
		    self.index_add(self.name, self.filename, static,
		                    "variable", type)
		    type = type_orig
		    token = self.token()
		    while token != None and token[0] == "sep":
		        type = type + token[1]
			token = self.token()
		    if token != None and token[0] == "name":
		        self.name = token[1]
			token = self.token()
		else:
		    break

	return token

    def parse(self):
        self.warning("Parsing %s" % (self.filename))
        token = self.token()
	while token != None:
            if token[0] == 'name':
	        token = self.parseGlobal(token)
            else:
	        self.error("token %s %s unexpected at the top level" % (
		       token[0], token[1]))
		token = self.parseGlobal(token)
		return
	self.parseTopComment(self.top_comment)
        return self.index



def parse(filename):
    parser = CParser(filename)
    idx = parser.parse()
    return idx

def rebuild():
	pass

if __name__ == "__main__":
    if len(sys.argv) > 1:
        debug = 1
        parse(sys.argv[1])
    else:
    	rebuild()

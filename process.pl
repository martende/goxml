#!/usr/bin/perl


@IMPORTS = (
	'xmlTextReaderConstName',
	'xmlTextReaderConstValue',
	'xmlTextReaderDepth',
	'xmlTextReaderNodeType',
	'xmlTextReaderRead',
	'xmlTextReaderIsEmptyElement',
	'xmlTextReaderHasValue',
	'xmlReaderForFile','xmlParseChunk','xmlReadFile','xmlReadMemory','xmlFreeDoc','xmlAddChild','xmlCleanupParser','xmlMemoryDump','xmlNewParserCtxt','xmlFreeParserCtxt','xmlCtxtReadFile','xmlCreatePushParserCtxt');
%TYPE_ALIASES = (
	'xmlDoc*' => 'xmlDocPtr'
);
%TYPE_CONVERSIONS = (
	'xmlNodePtr' => {
		'goType'=>'*XmlNode',
		'cConverter' => '%s.handler'
		#'cType'
	},
	'xmlDocPtr' => {
		'goType'=>'*XmlDoc',
		'cConverter' => '%s.handler'
		#'cType'
	},
	'xmlParserCtxtPtr' => {
		'goType'=>'*XmlParserCtxt',
		'cConverter' => '%s.handler',
		'export_as' => 'xmlParserCtxt'
		#'cType'
	},
	'xmlSAXHandlerPtr' => {
		'goType'=>'*XmlSAXHandler',
		'cConverter' => 'NULL_OR_HANDLER'
	},
	'xmlTextReaderPtr' => {
		'goType'=>'*XmlTextReader',
		'cConverter' => '%s.handler'
	},
	'const char *' => {
		'goType' => 'string',
		'cType'=>'C.CString',
		'cConverter'=> 'C.CString(%s)' ,
	},
	'int' => {
		'goType' => 'int',
		#'cType'=>'C.int',
		'cConverter'=> 'C.int(%s)' ,
	},
	'const xmlChar*' => {
		'goType' => 'string',
		'cType'=>'*C.char',
		'cConverter'=> 'C.CString(%s)' 
	},
	#'void *' => {
	#	'goType' => 'interface[]',
	#	'cType' => 'unsafe.Pointer',
	#	#'cConverter'=> 'C.GoBytes(%s,%i)' ,
	#}
	'void' => {
		
	}	
);

require 'funcdescs.pl';

%EXPORTED_TYPES = map {$TYPE_CONVERSIONS{$_}->{export_as}=>$_}  
					grep {exists $TYPE_CONVERSIONS{$_}->{export_as}} 
						keys %TYPE_CONVERSIONS;

use Data::Dumper;
$Data::Dumper::Maxdepth = 3;
$Data::Dumper::Indent = 2;
$Data::Dumper::Seen = ['chan'];


$CFLAGS = `pkg-config --cflags libxml-2.0`;
chomp($CFLAGS);
$r = time;
$r=1;
$TMP = "/tmp/tmp".$r;
mkdir $TMP;



sub remove_includes {
	my $content = shift;
	my $f = shift;
	my $status = 'go';
	my @lines = split "\n", $content;
	###  Remove all from other libs
	my @postprocess = ();
	my $sfile = "";
	my $opcount = 0;
	for (my $i=0;$i<scalar @lines;$i++) {
		# Processops
		my %flags = ();
		my $file = ();
		if ( $lines[$i]=~/^\# (\d+) \"([^"]+)\" ([\d ]+)$/) {
			$file = $2;
			# 1 This indicates the start of a new file. 
			# 2 This indicates returning to a file (after having included another file). 
			# 3 This indicates that the following text comes from a system header file, so certain warnings should be suppressed. 
			# 4 This indicates that the following text should be treated as being wrapped in an implicit extern "C" block.
			%flags = map {$_=>1;} split(" ",$3);
		}
		if ( $status  eq 'lookend' ) {
			if ( $file eq $f and $flags{2} ) {
				$status = 'go';
				#push @postprocess,"--- REMOVE " . $lines[$i];
			} else {
				#push @postprocess,"--- REMOVE " . $lines[$i];
			}
		} elsif ( $status eq 'go' ) {
			if ( $flags{1} ) {
				# look for include end
				$status = 'lookend';
				#$sfile = $file;
				#$opcount+=1;
				#push @postprocess,$lines[$i];
			} else {
				push @postprocess,$lines[$i];
			}
		}
	}
	return join "\n",@postprocess;
}


sub process_enums {
	my $content = shift;	
	open(FH,">$TMP/t.h");
	print FH $content;
	close(FH);
	unlink "$TMP/tags";
	my $sys = "anjuta-tags --fields=iKSTs -o $TMP/tags  $TMP/t.h";
	#print $sys."\n";
	system($sys);
	
	open(FH,"$TMP/tags");
	@d=<FH>;
	close(FH);
	
	my @vals;
	@functions = ();
	
	foreach (@d) {
		chomp;
		my ($name,$file,$re,$type,$sig,undef)  = split /\t/;
		if ($type eq 'enumerator') {
			push @vals,$name;
		}
		elsif ($type eq 'struct' and exists $EXPORTED_TYPES{$name}) {	
			foreach my $ssearch(@d) {
				chomp $ssearch;
				my ($name2,undef,$re,$type,$sig,undef)  = split /\t/,$ssearch;
				if ($type eq 'member' and $sig eq "struct:$name") {
					#print $ssearch ."\n";
				}
			}
		}
	}
	return \@vals;
}


sub process_functions {
	my $content = shift;	
	open(FH,">$TMP/t.h");
	print FH $content;
	close(FH);
	unlink "$TMP/tags";
	my $sys = "anjuta-tags --fields=iKST --c-kinds=p -o $TMP/tags  $TMP/t.h";
	#print $sys."\n";
	system($sys);
	
	open(FH,"$TMP/tags");
	@d=<FH>;
	close(FH);
	
	my @functions = ();
	
	foreach (@d) {
		chomp;
		my ($name,$file,$re,$type,$sig,$returntype)  = split /\t/;
		if ($type eq 'prototype') {
			my $funcName = $name;
			$sig=~s/(^signature:\(|\)$)//g;
			$returntype=~s/^returntype://g;
			my @params = ();
			if ($sig ne 'void') {
				foreach (split /,/,$sig) {
					my ($k,$v) = m/^(.*\s\**)([a-zA-Z]\w+)$/;
					$k=~s/(^ | $)//g;
					push @params , [$k,$v];
				}
			}
			$returntype=~s/^const([^ ])/const $1/;
			push @functions,{
				name => $funcName,
				retType => $returntype,
				params => \@params
			}
			#print "$returntype $funcName($sig)",Dumper(\%params);
		}
	}
	return \@functions;
}

sub look_in_db {
	my ($f,$t,$n) = @_;
	my $s = "$f/$t/$n";
	for (my $i=0;$i<scalar @FUNC_DESCS;$i+=2) {
		if ($s eq $FUNC_DESCS[$i]) {
			return $FUNC_DESCS[$i+1];
		}
	}
	return undef
}

sub create_func_args_string {
	my ($retType,$funcName,$params) = @_;
	my @args = ();
	foreach (@{$params}) {
		my ($type,$name) = @{$_};
		my $dbres = &look_in_db($funcName,$type,$name);
		if ($dbres) {
			if ( $dbres =~ /^SKIP/) {
				next;
			} else {
				die "undefined DBRES $dbres ";
			}
		}
		my $goType = $TYPE_CONVERSIONS{$type}->{goType};
		die "$type not found for func $f->{name}" if (! $goType ) ;
		push @args,"$name $goType";
	}
	
	die "RetType $retType not found for func $funcName" if (! exists $TYPE_CONVERSIONS{$retType} ) ;
	
	my $func_str = "func \u$funcName(" . join(",",@args) . ")";
	$func_str.=" " . $TYPE_CONVERSIONS{$retType}->{goType} if ($retType ne "void");
	
	return $func_str; 
}

sub create_func_input_inits {
	my ($funcName,$params) = @_;
	my @content;
	foreach (@{$params}) {
		my $cConverter ;
		my ($type,$name) = @{$_};
		my $dbres = &look_in_db($funcName,$type,$name);
		if ($dbres) {
			if ( $dbres =~ /^SKIP/) {
				$cConverter  =$';
				if ($sConverter ) {
					$cConverter =~s/^\///;
				} else {
					$cConverter ='unsafe.Pointer(nil)';
				}
			} else {
				die "undefined DBRES $dbres ";
			}
		} else {
			my $cFormat =  $TYPE_CONVERSIONS{$type}->{cConverter};
			if ($cFormat eq 'NULL_OR_HANDLER') {
				push @content,"\tvar c_$name C.$type=nil";
				push @content,"\tif $name!=nil { c_$name = $name.handler}";
				next;
				#$cConverter = sprintf $TYPE_CONVERSIONS{$type}->{cConverter} , $name;
			} else {
				$cConverter = sprintf $TYPE_CONVERSIONS{$type}->{cConverter} , $name;
			}
		}
		die "no Sconverter" unless defined $cConverter; 
		my $line = "\tc_$name := $cConverter";
		push @content,$line;
	}
	return @content;
}

sub compile_function {
	my $f = shift;
	my $funcName = $f->{name};
	my $retType = $f->{retType};
	my $params = $f->{params};
	
	my $func_str = &create_func_args_string($retType,$funcName,$params);
	
	my @content;
	
	my $goRetType = $TYPE_CONVERSIONS{$retType}->{goType};
	
	if (  $retType ne 'void' ) {
		if ( exists $TYPE_CONVERSIONS{$retType}->{cType}) {
			push @content,"\tvar c_ret $TYPE_CONVERSIONS{$retType}->{cType}";
		} else {
			push @content,"\tvar c_ret C.$retType";
		}	 
		if ( $goRetType =~/^\*\s*(\w+)$/ ) {
			my $rt = $1;
			push @content,"\tg_ret := &$rt\{\}";
		} else {
			push @content,"\tvar g_ret $goRetType";
		}
	}
	my @inputs_inits = &create_func_input_inits($funcName,$params);
	
	push @content , @inputs_inits; 
	
	my @callargs = ();
	foreach (@{$f->{params}}) {
		my ($type,$name) = @{$_};
		push @callargs,"c_$name";
	}
	my $callline = "C." .$f->{name} . "(" . join(",",@callargs) . ")";
	
	if ($f->{retType} ne "void") {
		$callline = "c_ret = " . $callline;
	}
	push @content,"\t".$callline;
	
	if ($retType ne "void") {
		if ($goRetType eq 'string') {
			push @content, "\tg_ret = C.GoString(c_ret)";
		} elsif ( $goRetType=~/^\*\s*(\w+)$/ ) {
			#my $rt = $1;
			my $f = sprintf $TYPE_CONVERSIONS{$f->{retType}}->{cConverter} , 'g_ret';
			push @content, "\t$f = c_ret";
		} else {
			push @content, "\tg_ret = $goRetType(c_ret)";
		}
		push @content,"\treturn g_ret";
	}
	
	$func_str.= " {\n" . join("\n",@content) . "\n}\n";
	
	return $func_str; 
}

sub process {
	my $f = shift;
	my ($include_path) = $f=~/(libxml\/\w+\.h)$/;
	my ($file_path) = $f=~/(\w+\.h)$/;
	$file_path =~s/\.h$/.go/;
	my $content = `gcc -E $CFLAGS $f`;
	
	$content = &remove_includes($content,$f);
	
	$functions = &process_functions($content);
	my $enums = &process_enums($content);
	
	my @funcs = ();
	my $max = scalar @$functions;
	#$max = 2;
	
	%funcfilter = map{$_=>1} @IMPORTS;
	
	if ($ARGV[0]) {
		%funcfilter = map{$_=>1} @ARGV;
	}
	
	for (my $i=0;$i<$max;$i++) {
		next unless (exists $funcfilter{$functions->[$i]->{name}});
		#print Dumper($functions->[$i]);
		push @funcs ,&compile_function($functions->[$i]);
	}
	my $out_data = 
'package goxml
/*
#cgo pkg-config: libxml-2.0
#include <'.$include_path.'>
*/
import "C"

import "unsafe"
';
	if (scalar @{$enums}) {
		$out_data.= "\n";
		$out_data.= join "\n",map {"const $_=C.$_" } @{$enums};
	}
	$out_data.= "\n";	
	$out_data.= join "\n",@funcs;
	open(FH,">$file_path");
	print FH $out_data;
	close(FH);
	
	#print $out_data;
	#print @d;
}

sub test_type_processor() {
	
	#foreach (@$f) {
	#	print FH 
	#}
	my $v = &process_types('
int kokoshnichek=2;

struct m1 {
	int a1;
	int a2;
};
typedef struct m1 m1type;

struct _xmlParserCtxt {
	int kotenka1;
	int *kotenka2;
	int **kotenka3;
	int ***kotenka4;
	void *void1;
	unsigned long* uint;
	m1type *sref;
	struct m1 sref2; 
};

typedef struct _xmlParserCtxt xmlParserCtxt;
xmlParserCtxt val_xmlParserCtxt;
main(){};
',['xmlParserCtxt']);
	print Dumper($v); 
}


sub process_xml_types {
	my $includes = shift;
	my $types = shift;
	my $content = "";
	my $is = "";
	foreach my $v(@{$includes}) {
		$content .= "#include \"$v\"\n";
		my ($include_path) = $v=~/(libxml\/\w+\.h)$/;
		$is.= "#include <$include_path>\n" if ($include_path);
	}
	$content .= "main() {\n";
	foreach (@$types) {
		$content .= "$_ val_$_;\n";
	}
	$content .= "}\n";
	my $v = &process_types($content,$types);
	#print Dumper($v);
	#print &create_C_struct("xmlParserCtxt",$v->{xmlParserCtxt});
	$content = 'package goxml
/*
#cgo pkg-config: libxml-2.0
'.$is.'*/
import "C"
';
	foreach (@$types) {
		$content.= &create_go_struct($_,$v->{$_});
	}
	open(FH,">atypes.go");
	print FH $content;
	close(FH);
}

sub create_C_struct {
	my $t = shift;
	my $s = shift;
	my $content = "struct _$t \{\n";
	$content .= join("\n" , (map {"\t$_->[1]\t$_->[0];";} @{$s} ) );
	$content .= "\n";
	$content .="\} $t;\n";
	
	#print Dumper($t);
	return $content;
}

sub create_go_struct {
	my $t = shift;
	my $s = shift;
	my $content ="";
	# Create getters
	my $gotype = $TYPE_CONVERSIONS{$EXPORTED_TYPES{$t}}->{goType};
	$gotype =~s/^\*//;
	my $ctype = $EXPORTED_TYPES{$t};
	my $declaration = "";
	
	foreach (@$s) {
		my ($name,$type) = @{$_};
		my $goftype = $type;
		if (exists $TYPE_ALIASES{$type}) {
			$goftype = $TYPE_ALIASES{$type};
		}
		if (exists $TYPE_CONVERSIONS{$goftype}) {
			$goftype = $TYPE_CONVERSIONS{$goftype}->{goType};
		}
		
		my $inner = "func (this *$gotype) Get\u$name() $goftype {\n";
		my $found = 0;
		if ( $goftype eq 'int') {
			$inner .= "\treturn int(this.handler.$name)\n";
			$found = 1;
		} elsif ($goftype=~/^\*(\w+)$/) {
			my $rt = $1;
			$found = 1;
			$inner.= "\tif this.handler.$name == nil {\n";
			$inner.= "\t\treturn nil\n";
			$inner.= "\t}\n";
			$inner.= "\tif this._$name == nil {\n";
			$inner.= "\t\tthis._$name = &$rt\{\}\n";
			$inner.= "\t}\n";
			$inner.= "\tthis._$name.handler = this.handler.$name\n";
			$inner .= "\treturn this._$name\n";
			push @addon_types , "\t_$name $goftype"
		}
		$inner .= "}\n";
		if ( $found ) {
			$content .= $inner; 
		} else {
			$content .= "/*\n$inner*/\n";
		}
	}
	$declaration .= "type $gotype struct {\n";
	$declaration .= "\thandler C.$ctype\n";
	if ( scalar @addon_types) {
		$declaration .= join "\n"  , @addon_types;
		$declaration .="\n";
	}
	$declaration .= "}";
	$content = "$declaration\n$content";
	return $content;
}
sub process_types {
	my $content = shift;
	my $types = shift;
	
	open FH,">$TMP/t.c";
	print FH $content;
	close(FH);
	unlink("$TMP/t.c.001t.tu");
	my $cmd = "cd $TMP/;gcc -fdump-translation-unit  $CFLAGS t.c";
	#print $cmd ;
	system($cmd);
	
	my @tu=();
	open(FH,"$TMP/t.c.001t.tu");
	my $dump;
	while(my $line=<FH>) {
		chomp $line;
		if($line =~ /^\@(\d+)/) {
			push @tu,&parse_tu_node($dump) if ($dump);
			$dump = $line;
		} else {
			$dump .= $line;
		}
	}
	close(FH);
	
	push @tu,&parse_tu_node($dump) if ($dump);
	
	foreach my $node (@tu) {
		foreach (grep { $_!~/^__/ } keys %{$node}) {
			if ($node->{$_}=~/^\@(.*)$/) {
				$node->{$_}=$tu[$1-1];
			}
		}
    }
    
    #print Dumper(@tu);
    my  %ret;
    foreach (@$types) {
    	$ret{$_}=&restore_type(\@tu,"val_$_");
    }
    return \%ret;
}

sub restore_type {
	my $tu = shift;
	my $varname = shift;
	
	my @fields ; 
	foreach my $el (@$tu) {
		if ( $el->{TYPE} eq 'var_decl' and $el->{name}->{strg} eq $varname) {
			# $el->{type}->{name}->{name}->{strg} eq $type
			my $el_name = $el->{name}->{strg};
			#if ( $el->{type}->{TYPE} eq 'record_type'
			
			die "NO type_decl" unless ($el->{type}->{name}->{TYPE} eq 'type_decl');
			die "NO record_type" unless ($el->{type}->{TYPE} eq 'record_type');
			
			my $field = $el->{type}->{flds};
			while ( $field) {
				die "NO field_decl" unless ( $field->{TYPE} eq 'field_decl') ;
				
				my $name = $field->{name}->{strg};
				my $type = &resrore_type_desc($field->{type});
				#print Dumper($field->{type});
				$field = $field->{chan};
				
				push @fields , [$name , $type];
			}
			last;
		}
	}
	return \@fields;
}

sub resrore_type_desc {
	#my \@tu = shift;
	my $el = shift;
	
	if ($el->{TYPE} eq 'pointer_type') {
		return &resrore_type_desc($el->{ptd}) . "*"; 
	} elsif ( $el->{tag} eq 'struct') {
		if ( $el->{name}->{TYPE} eq 'identifier_node') { 
			return 'struct ' . $el->{name}->{strg}; 
		} elsif ($el->{name}->{TYPE} eq 'type_decl' ){
			return $el->{name}->{name}->{strg};
		}	else  {
			die "resrore_type_desc1";
		}

	}
	return $el->{name}->{name}->{strg};
}
sub parse_tu_node {
	my $dump = shift;
	my $node = {};
	my $td = $dump;
	
	unless($dump =~ s/^\@(\d+)\s+(\w+)(?=\s)//) {
		die "Unknown node format:\n$dump";
    }
    my $index = $1;
    my $type = $2;
    $node->{TYPE} = $type;
    $node->{__index} = $index;
    $node->{__str} = $td;
    if($dump =~ s/\s+strg:\s(.*)\slngt:\s(\d+)//s) {
    	# identifier_node and string_cst come here, at least
    	my($string, $length) = ($1, $2);
    	# string_cst's lngt includes the NUL character, which fprintf()
    	# doesn't print, obviously. Make sure to factor that in...
    	$length-- if $type eq 'string_cst';
    	$node->{'strg'} = substr($string, 0, $length);
    	$node->{'lngt'} = $length;
    }
    $node->{'source'} = $1 if $dump =~ s/\ssrcp:\s(.*?:\d+)(?=\s)//;
    if($dump =~ s/\squal:\s(.{3})\s//) {
    	my $qual = $1;
    	$node->{'const'}    = 1 if $qual =~ /c/;
    	$node->{'volatile'} = 1 if $qual =~ /v/;
    	$node->{'restrict'} = 1 if $qual =~ /r/;
    }
    
    while($dump =~ s/\s(\w+)\s*:\s(\S+)//) {
    	my($key, $value) = ($1, $2);
    	$node->{$key} = $value;
    }
    # All that should remain is flags
    while($dump =~ s/(\w+)//) {
    	#print "TRUE $1\n";
		$node->{$1} = 1;
    }
    
    if($dump =~ /\S/) {
    	$dump =~ s/\s+/ /g;
    	die "Unparsed data: $dump\nFrom: $_[1] " . Dumper($node);
    }

	return $node;
}

@f= (
	"/usr/include/libxml2/libxml/tree.h",
	"/usr/include/libxml2/libxml/parser.h",
	"/usr/include/libxml2/libxml/xmlreader.h",
	"/usr/include/libxml2/libxml/xmlmemory.h",
);
foreach (@f) {
	&process($_);
}
&process_xml_types(\@f,[keys %EXPORTED_TYPES])
#&test_type_processor(\@f);

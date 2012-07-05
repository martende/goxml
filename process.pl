#!/usr/bin/perl

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
		'export_as' => '_xmlParserCtxt'
		#'cType'
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
	'void' => {
		
	}	
);
use GCC::TranslationUnit;
use Data::Dumper;
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


sub compile_function {
	my $f = shift;
	
	@args = ();
	foreach (@{$f->{params}}) {
		my ($type,$name) = @{$_};
		my $goType = $TYPE_CONVERSIONS{$type}->{goType};
		die "$type not found for func $f->{name}" if (! $goType ) ;
		push @args,"$name $goType";
	}
	
	die "RetType $f->{retType} not found for func $f->{name}" if (! exists $TYPE_CONVERSIONS{$f->{retType}} ) ;
	
	my $func_str = "func \u$f->{name}(" . join(",",@args) . ")";
	$func_str.=" " . $TYPE_CONVERSIONS{$f->{retType}}->{goType} if ($f->{retType} ne "void");
	my @content;
	
	#foreach (@{$f->{params}}) {
	#	my ($type,$name) = @{$_};
	#	my $ctype =  $TYPE_CONVERSIONS{$type}->{cType} || "C.".$type;
	#	my $line = "\tc_$name :=";
	#	push @content,$line;
	#}
	
	if (  $f->{retType} ne 'void' ) {
		if ( $f->{retType}=~/\*$/ ) {
			die "notimpl1!";
		} else {
			push @content,"\tvar c_ret C.$f->{retType}";
		}
		if ( $TYPE_CONVERSIONS{$f->{retType}}->{goType}=~/^*\s*(\w+)$/ ) {
			my $rt = $1;
			push @content,"\tg_ret := &$rt\{\}";
		} else {
			die "notimpl2!";
		}
	}
	
	foreach (@{$f->{params}}) {
		my ($type,$name) = @{$_};
		$cConverter = sprintf $TYPE_CONVERSIONS{$type}->{cConverter} , $name;
		my $line = "\tc_$name := $cConverter";
		push @content,$line;
	}
	
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
	
	if ($f->{retType} ne "void") {
		if ( $TYPE_CONVERSIONS{$f->{retType}}->{goType}=~/^*\s*(\w+)$/ ) {
			#my $rt = $1;
			my $f = sprintf $TYPE_CONVERSIONS{$f->{retType}}->{cConverter} , 'g_ret';
			push @content, "\t$f = c_ret";
			push @content,"\treturn g_ret";
		}
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
	%funcfilter = map{$_=>1} ('xmlReadFile','xmlReadMemory','xmlFreeDoc','xmlAddChild','xmlCleanupParser','xmlMemoryDump','xmlNewParserCtxt','xmlFreeParserCtxt','xmlCtxtReadFile');
	
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

sub process_types {
	my $f = shift;
	open FH,">$TMP/t.c";
	#foreach (@$f) {
	#	print FH "#include \"$_\"\n";
	#}
	print FH '
int kokoshnichek=2;
struct _xmlParserCtxt {
int kotenka; 
};
typedef struct _xmlParserCtxt xmlParserCtxt;
xmlParserCtxt mimimishka;
main(){};
';
# typedef struct _xmlParserCtxt xmlParserCtxt;
#;
	close(FH);
	unlink("$TMP/t.c.001t.tu");
	my $cmd = "cd $TMP/;gcc -fdump-translation-unit  $CFLAGS t.c";
	#print $cmd ;
	system($cmd);
	
	my $node = GCC::TranslationUnit::Parser->parsefile("$TMP/t.c.001t.tu")->root;
	
	while($node) {
		#print $node->name->identifier ."\t" . $node->source . "\n";
		next if ($node->source eq '<built-in>:0' or ! $node->source);
		print "NODE  " . ref $node ;
		print "\n";
		if( $node->isa('GCC::Node::function_decl') and $node->name) {
			printf "%s declared in %s\n",
			$node->name->identifier, $node->source;
		}
		if( $node->isa('GCC::Node::var_decl') and $node->name) {
			if ($node->name->identifier eq 'mimimishka') {
				#if ( $node->name->identifier
				#delete $node->{init};
				delete $node->{chan};
				#delete $node->{type}->{name}->{chan};
				print Dumper($node);
			}
			
			printf "%s declared in %s %s %s\n",
			$node->name->identifier, $node->source, $node->type->name->name->identifier , $node->type->name->name->{string};
			
			#print Dumper($node);
		}
		if( $node->isa('GCC::Node::type_decl') and $node->name) {
		  printf "%s declared in %s\n",
		  	
			$node->name->identifier, $node->source;
		}
	  } continue {
		$node = $node->chain;
	  }

}
%EXPORTED_TYPES = map {$TYPE_CONVERSIONS{$_}->{export_as}=>$_}  grep {exists $TYPE_CONVERSIONS{$_}->{export_as}} keys %TYPE_CONVERSIONS;

@f= (
#	"/usr/include/libxml2/libxml/tree.h",
	"/usr/include/libxml2/libxml/parser.h",
#	"/usr/include/libxml2/libxml/xmlmemory.h",
);
foreach (@f) {
	&process($_);
}

&process_types(\@f);

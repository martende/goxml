#!/usr/bin/perl
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

sub process_functions {
	my $content = shift;	
	open(FH,">$TMP/t.h");
	print FH $content;
	close(FH);
	my $sys = "anjuta-tags --fields=iKST --c-kinds=p -o $TMP/tags  $TMP/t.h";
	print $sys."\n";
	system($sys);
	
	open(FH,"$TMP/tags");
	@d=<FH>;
	close(FH);
	
	@functions = ();
	
	foreach (@d) {
		chomp;
		my ($name,$file,$re,$type,$sig,$returntype)  = split /\t/;
		if ($type eq 'prototype') {
			my $funcName = $name;
			$sig=~s/(^signature:\(|\)$)//g;
			$returntype=~s/^returntype://g;
			my @params = ();
			foreach (split /,/,$sig) {
				my ($k,$v) = m/^(.*\s\**)([a-zA-Z]\w+)$/;
				$k=~s/(^ | $)//g;
				push @params , [$k,$v];
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
%TYPE_CONVERSIONS = (
	'xmlNodePtr' => {
		'goType'=>'*XmlNode',
		#'cType'
	}
);

sub compile_function {
	my $f = shift;
	
	@args = ();
	foreach (@{$f->{params}}) {
		my ($type,$name) = @{$_};
		my $goType = $TYPE_CONVERSIONS{$type}->{goType};
		die "$type not found for func $f->{name}" if (! $goType ) ;
		push @args,"$name $goType";
	}
	my $retType = $TYPE_CONVERSIONS{$f->{retType}}->{goType};
	
	die "RetType $f->{retType} not found for func $f->{name}" if (! $retType ) ;
	
	my $func_str = "func $f->{name}(" . join(",",@args) . ")";
	$func_str.=" $retType" if ($retType ne "void");
	my @content;
	
	foreach (@{$f->{params}}) {
		my ($type,$name) = @{$_};
		my $ctype =  $TYPE_CONVERSIONS{$type}->{cType};
		my $line = "var c_$name C.$type";
		push @content,$line;
	}
	
	#foreach (@{$f->{params}}) {
	#	my ($type,$name) = @{$_};
	#	my $ctype =  $TYPE_CONVERSIONS{$type}->{cType};
	#	my $line = "var c_$name C.$type";
	#	push @content,$line;
	#}
	
	if (  $retType ne 'void' ) {
		if ( $retType=~/\*$/ ) {
			
		} else {
			push @content,"var c_ret C.$retType";
		}
	}
	
	$func_str.= " {\n" . join("\n",@content) . "\n}\n";
	print $func_str."";
}
sub process {
	my $f = shift;
	
	my $content = `gcc -E $CFLAGS $f`;
	
	$content = &remove_includes($content,$f);
	
	$functions = &process_functions($content);
	
	
	for (my $i=0;$i<scalar @$functions;$i++) {
		&compile_function($functions->[$i]);
		last;
	}
	
	#print @d;
}

@f= ("/usr/include/libxml2/libxml/tree.h");
foreach (@f) {
	&process($_);
}

#include <libxml/parser.h>
#include <libxml/parserInternals.h>
#include <stdio.h>
#include <string.h>
int xmlInputReadCallback_Function (void * context, char * buffer, int length ) {
	sprintf(buffer,"<br>");
	printf("xmlInputReadCallback_Function %p %p %i\n",context,buffer,length);
	return strlen(buffer);
}
int xmlInputCloseCallback_Function(void * context) {
	printf("xmlInputCloseCallback_Function %p\n",context);
	return 1;
}

int main()
{
	xmlParserCtxtPtr ctxt; /* the parser context */

	ctxt = xmlNewParserCtxt();

	printf("Start %p\n",ctxt);
   xmlCtxtReadIO(ctxt,
       &xmlInputReadCallback_Function,
       &xmlInputCloseCallback_Function,
       NULL,
       "test.xml",
       NULL,
       0
   );

   return 0;
}

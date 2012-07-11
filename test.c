#include <stdio.h>
#include <libxml/parser.h>

void * userdata ;
void set_userdata(void*d) {
	printf("C-set_userdata1: 0x%08x\n",(unsigned int)d);
	printf("C-set_userdata2: 0x%08x\n",*(unsigned int*)d);
	printf("C-set_userdata3: 0x%08x\n",**(unsigned int**)d);
	userdata = d;
	
}
void * get_userdata(void) {
	printf("get_userdata: 0x%08x\n",(unsigned int)userdata);
	return userdata;
}
int UTF8ToHtml2(unsigned char* a,int*b,unsigned char*c,int*d)  {
	printf("Svisni v huy pridutrok d=%i \n",*d);
	*b = 666;
}


extern void go_callback_xmlInputReadCallback(void* foo, int p1);

void callbackcaller(void* pfoo) {
	go_callback_xmlInputReadCallback(pfoo, 5);
}


#include <stdio.h>
#include <libxml/parser.h>
#include "test.h"

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

struc2* createStructArray() {
	struc1 *s1,*s2;
	struc2 *s;
	
	s1 = calloc(sizeof(struc1),1);
	s2 = calloc(sizeof(struc1),1);
	s = calloc(sizeof(struc2),1);
	s1->a1 = 0x0a;
	s1->a2 = 0x0b;
	s2->a1 = 0x0c;
	s2->a2 = 0x0d;
	s->len=2;
	s->data = calloc(sizeof(struc1*),2);
	s->data[0]=s1;
	s->data[1]=s2;
	printf("C.createStructArray s=%p s->data = %p s->data[0] %p\n" , s,s->data,s->data[0]);
	return s;
}

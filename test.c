#include <stdio.h>
void * userdata ;
int set_userdata(void*d) {
	printf("set_userdata: 0x%08x\n",(unsigned int)d);
	userdata=d;
}
void * get_userdata(void) {
	printf("get_userdata: 0x%08x\n",(unsigned int)userdata);
	return userdata;
}
int UTF8ToHtml2(char* a,int*b,unsigned char*c,int*d)  {
	printf("Svisni v huy pridutrok\n");
}

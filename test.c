#include <stdio.h>
void * userdata ;
set_userdata(void*d) {
	printf("set_userdata: 0x%08x\n",(unsigned int)d);
	userdata=d;
}
void * get_userdata() {
	printf("get_userdata: 0x%08x\n",(unsigned int)userdata);
	return userdata;
}

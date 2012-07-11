
struct _callbackInfo {
	void *fakedCallback;
	void *realCallback;
};

typedef struct _callbackInfo callbackInfo;
typedef callbackInfo *callbackInfoPtr;

#define CBSIZE sizeof(callbackInfo) 

void set_userdata(void*d);
void* get_userdata(void);
int UTF8ToHtml2(unsigned char* a,int*b,unsigned char*c,int*d);
void callbackcaller(void* pfoo);

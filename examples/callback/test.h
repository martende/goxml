struct _struc1 {
	int a1,a2 ;
} ;

typedef struct _struc1 struc1;

struct _struc2 {
	int len;
	struc1 **data;
};
typedef struct _struc2 struc2;

struc2* createStructArray() ;

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

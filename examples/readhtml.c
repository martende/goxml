#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include <libxml/HTMLparser.h>

htmlDocPtr gethtml(char *doclocation,char *encoding) {
    htmlDocPtr doc;

    doc = htmlReadFile(doclocation, encoding, HTML_PARSE_NOBLANKS | HTML_PARSE_NOERROR | HTML_PARSE_NOWARNING | HTML_PARSE_NONET);

    if (doc == NULL) {
        fprintf(stderr, "Document not parsed successfully.\n");
        return;
    }

    return doc;
}

void getroot(htmlDocPtr doc) {
    xmlNode *cur = NULL;

    cur = xmlDocGetRootElement(doc);

    if (cur == NULL) {
        fprintf(stderr, "empty document\n");
        xmlFreeDoc(doc);
        return;
    }

    printf("%s\n", cur->name);

}

int main(void) {
    char *website = "http://www.google.com/index.html";
    char *encoding = "UTF-8";
    htmlDocPtr doc;

    doc = gethtml(website, encoding);

    getroot(doc);

    return 0;
}

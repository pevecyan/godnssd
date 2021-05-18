#include "dns_sd.h"

int test();

TXTRecordRef CreateTXTRecordsList();
int SetStringTXTRecordValue(TXTRecordRef* ref, char * key, int length, char* value);
int UpdateTXTRecords(DNSServiceRef serviceRef, TXTRecordRef* txtRef);


DNSServiceRef  RegisterService(long *status, char* name, char* protocol, uint16_t  port, TXTRecordRef* txtRef);

void DeregisterService(DNSServiceRef  serviceRef, TXTRecordRef *txtRef);
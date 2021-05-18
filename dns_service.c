#include "dns_sd.h"
#include <arpa/inet.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int test(){
	return 0;
}

TXTRecordRef CreateTXTRecordsList() {
	TXTRecordRef recordRef;
	TXTRecordCreate(&recordRef, 0, NULL);
	return recordRef;
}

int SetStringTXTRecordValue(TXTRecordRef* ref, char * key, int length, char* value) {
	
	return TXTRecordSetValue(ref, key, length, value);
}

int UpdateTXTRecords(DNSServiceRef serviceRef, TXTRecordRef* txtRef) {
	return DNSServiceUpdateRecord(serviceRef, NULL, kDNSServiceFlagsMoreComing, TXTRecordGetLength(txtRef), TXTRecordGetBytesPtr(txtRef), 0);
}

int SetIntTXTRecordValue(TXTRecordRef* ref, char * key, int length, char* value) {
	
	return TXTRecordSetValue(ref, key, length, value);
}

DNSServiceRef RegisterService(long *status, char* name, char* protocol, uint16_t  port, TXTRecordRef* txtRef)
{
		DNSServiceRef serviceRef;
	
		*status = DNSServiceRegister(&serviceRef,
			0, 
			0, 
			name, 
			protocol, 
			NULL, 
			NULL, 
			htons(port), 
			TXTRecordGetLength(txtRef), 
			TXTRecordGetBytesPtr(txtRef), 
			NULL, 
			NULL
		);
		return serviceRef;
}

void DeregisterService(DNSServiceRef serviceRef, TXTRecordRef *txtRef)
{
	//free(serviceRef);
	DNSServiceRefDeallocate(serviceRef);
	TXTRecordDeallocate(txtRef);
}
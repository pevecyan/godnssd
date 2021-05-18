package godnssd

// #cgo LDFLAGS: -L/mnt/c/Users/pevec/Documents/Development/godnssd -ldns_sd
// #cgo CFLAGS: -I/mnt/c/Users/pevec/Documents/Development/godnssd
// #include "dns_service.h"
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

func (service *DNSService) Register() int {
	var status *int = new(int)

	txtRecordRef := C.CreateTXTRecordsList()
	service.TXTRecordRef = txtRecordRef

	for key := range service.Config.Records {
		value := service.Config.Records[key]
		keyC := C.CString(key)
		valueC := C.CString(value)
		C.SetStringTXTRecordValue(&txtRecordRef, keyC, C.int(len([]byte(value))), valueC)

		//Free strings
		C.free(unsafe.Pointer(keyC))
		C.free(unsafe.Pointer(valueC))
	}

	nameC := C.CString(service.Config.Name)
	protocolC := C.CString(service.Config.Protocol)
	ref := C.RegisterService((*C.long)(unsafe.Pointer(status)), nameC, protocolC, C.ushort(service.Config.Port), &txtRecordRef)
	service.ServiceRef = ref

	//Free strings
	C.free(unsafe.Pointer(nameC))
	C.free(unsafe.Pointer(protocolC))
	return *status
}

func (service *DNSService) Unregister() {
	if service.ServiceRef != nil {
		C.DeregisterService(service.ServiceRef, &service.TXTRecordRef)
		service.ServiceRef = nil
	}
}

func (service *DNSService) UpdateRecord(updateKey, updateValue string) int {
	if service.ServiceRef != nil {
		service.Config.Records[updateKey] = updateValue

		for key := range service.Config.Records {
			value := service.Config.Records[key]
			keyC := C.CString(key)
			valueC := C.CString(value)
			C.SetStringTXTRecordValue(&service.TXTRecordRef, keyC, C.int(len([]byte(value))), valueC)

			//Free strings
			C.free(unsafe.Pointer(keyC))
			C.free(unsafe.Pointer(valueC))
		}
		status := C.UpdateTXTRecords(service.ServiceRef, &service.TXTRecordRef)
		return int(status)

	}
	return KDNSServiceErr_NotInitialized
}

func NewService(config DNSServiceConfig) *DNSService {
	return &DNSService{
		Config: config,
	}
}

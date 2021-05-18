package godnssd

// #cgo LDFLAGS: -L/mnt/c/Users/pevec/Documents/Development/godnssd -ldns_sd
// #cgo CFLAGS: -I/mnt/c/Users/pevec/Documents/Development/godnssd
// #include "dns_service.h"
import "C"

const (
	KDNSServiceErr_NoError                   = 0
	KDNSServiceErr_Unknown                   = -65537 /* 0xFFFE FFFF */
	KDNSServiceErr_NoSuchName                = -65538
	KDNSServiceErr_NoMemory                  = -65539
	KDNSServiceErr_BadParam                  = -65540
	KDNSServiceErr_BadReference              = -65541
	KDNSServiceErr_BadState                  = -65542
	KDNSServiceErr_BadFlags                  = -65543
	KDNSServiceErr_Unsupported               = -65544
	KDNSServiceErr_NotInitialized            = -65545
	KDNSServiceErr_AlreadyRegistered         = -65547
	KDNSServiceErr_NameConflict              = -65548
	KDNSServiceErr_Invalid                   = -65549
	KDNSServiceErr_Firewall                  = -65550
	KDNSServiceErr_Incompatible              = -65551 /* client library incompatible with daemon */
	KDNSServiceErr_BadInterfaceIndex         = -65552
	KDNSServiceErr_Refused                   = -65553
	KDNSServiceErr_NoSuchRecord              = -65554
	KDNSServiceErr_NoAuth                    = -65555
	KDNSServiceErr_NoSuchKey                 = -65556
	KDNSServiceErr_NATTraversal              = -65557
	KDNSServiceErr_DoubleNAT                 = -65558
	KDNSServiceErr_BadTime                   = -65559 /* Codes up to here existed in Tiger */
	KDNSServiceErr_BadSig                    = -65560
	KDNSServiceErr_BadKey                    = -65561
	KDNSServiceErr_Transient                 = -65562
	KDNSServiceErr_ServiceNotRunning         = -65563 /* Background daemon not running */
	KDNSServiceErr_NATPortMappingUnsupported = -65564 /* NAT doesn't support PCP, NAT-PMP or UPnP */
	KDNSServiceErr_NATPortMappingDisabled    = -65565 /* NAT supports PCP, NAT-PMP or UPnP, but it's disabled by the administrator */
	KDNSServiceErr_NoRouter                  = -65566 /* No router currently configured (probably no network connectivity) */
	KDNSServiceErr_PollingMode               = -65567
	KDNSServiceErr_Timeout                   = -65568
)

type DNSService struct {
	Config       DNSServiceConfig
	ServiceRef   C.DNSServiceRef
	TXTRecordRef C.TXTRecordRef
}

type DNSServiceConfig struct {
	InterfceIndex int
	Name          string
	Protocol      string
	Port          uint16
	Records       map[string]string
}

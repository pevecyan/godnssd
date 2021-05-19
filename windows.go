// +build windows
package godnssd

import "fmt"

func NewService(config DNSServiceConfig) *DNSService {
	return &DNSService{
		Config: config,
	}
}

func (service *DNSService) Register() int {
	fmt.Println("[GODNSSD] Service registered")
	return 0
}

func (service *DNSService) UpdateRecord(key, value string) {
	service.Config.Records[key] = value
	fmt.Println("[GODNSSD] Record updated")
}

func (service *DNSService) Unregister() {
	fmt.Println("[GODNSSD] Service unregistered")
}

type DNSService struct {
	Config DNSServiceConfig
}

type DNSServiceConfig struct {
	InterfceIndex int
	Name          string
	Protocol      string
	Port          uint16
	Records       map[string]string
}

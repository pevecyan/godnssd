package main

import (
	"fmt"
	"time"

	"github.com/pevecyan/godnssd"
)

func main() {
	dnsService := godnssd.NewService(godnssd.DNSServiceConfig{
		InterfceIndex: 0,
		Name:          "Webserver test",
		Protocol:      "_http._tcp",
		Port:          1337,
		Records: map[string]string{
			"a": "1",
			"b": "a",
			"c": "žč",
			"d": "1.1",
			"e": "qwew8=?",
			"f": "1238213978",
		},
	})

	dnsService.Register()
	fmt.Println("REGISTER")

	<-time.After(time.Second * 5)

	dnsService.UpdateRecord("c", "abc")
	fmt.Println("UPDATE")

	<-time.After(time.Second * 10)

	dnsService.Unregister()
	fmt.Println("DEREGISTER")

	<-time.After(time.Second * 10)
	//fmt.Println(test)

	fmt.Println("END")
}

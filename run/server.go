package main

import (
	"etcdRegistry/discovery"
	"fmt"
	"time"
)

func main()  {
	serviceName := "s-test"
	serviceInfo := discovery.ServiceInfo{IP: "127.0.0.1"}

	s, err := discovery.NewService(serviceName, serviceInfo, []string{
		"http://127.0.0.1:2379",
		//"http://127.0.0.1:2479",
		//"http://127.0.0.1:2579",
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("name:%s, ip:%s\n", s.Name, s.Info.IP)

	go func() {
		time.Sleep(time.Second * 20)
		s.Stop()
	}()

	s.Start()
}

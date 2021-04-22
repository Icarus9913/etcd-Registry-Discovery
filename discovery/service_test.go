package discovery

import (
	"fmt"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	t.Run("service_test", func(t *testing.T) {
		serviceName := "s-test"
		serviceInfo := ServiceInfo{IP: "127.0.0.1"}

		s, err := NewService(serviceName, serviceInfo, []string{
			"http://127.0.0.1:2379",
			//"http://127.0.0.1:2479",
			//"http://127.0.0.1:2579",
		})

		if err != nil {
			panic(err)
		}

		fmt.Printf("name:%s, ip:%s\n", s.Name, s.Info.IP)

		go func() {
			time.Sleep(time.Second * 10)
			s.Stop()
		}()

		s.Start()
	})
}

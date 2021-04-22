package discovery

import (
	"fmt"
	"testing"
	"time"
)

func TestMaster(t *testing.T)  {
	m, err := NewMaster([]string{
		"http://127.0.0.1:2379",
		//"http://127.0.0.1:2479",
		//"http://127.0.0.1:2579",
	}, "services/")

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		for k, v := range  m.Nodes {
			fmt.Printf("node:%s, ip=%s\n", k, v.Info.IP)
		}
		fmt.Printf("nodes num = %d\n",len(m.Nodes))
		time.Sleep(time.Second * 5)
	}

}

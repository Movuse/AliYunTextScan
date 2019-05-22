package AliYunTextScan

import (
	"fmt"
	"log"
	"testing"
)

const accessKeyId = "<your access key id>"
const accessKeySecret = "<your access key secret>"

func TestAliTextScan(t *testing.T) {
	clientInfo := ClientInfo{IP: "127.0.0.1"}
	c, err := NewTextScan(accessKeyId, accessKeySecret, "/green/text/scan", &clientInfo, []string{"这是一句骂人的话: 傻逼"})
	if err != nil {
		log.Println(err)
	}

	resp, err := c.AliTextScan()
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("resp: %+v", resp)
}

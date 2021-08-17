package head_first

import (
	"fmt"
	"sync"
)

var single sync.Once

type Client struct {
	hostName string
	headers  map[string]string
}

var client *Client
func initializeClient(){
	fmt.Println("init client")
	client = new(Client)
	client.hostName = "https://www.google.com"
	client.headers = map[string]string{
		"auth":"dummy",
	}
}

func GetClient() *Client {
	single.Do(initializeClient)
	return client
}
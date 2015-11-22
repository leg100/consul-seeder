package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	//"gopkg.in/yaml.v2"
)

func main() {
	// Get a new client
	config := api.DefaultConfig()
	config.Address = "docker:8500"

	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// Lookup the pair
	pair, _, err := kv.Get("foo", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v", string(pair.Value))
}

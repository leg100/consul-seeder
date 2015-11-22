package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "consul-seeder"
	app.Usage = "consul-seeder"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host, c",
			Value:  "localhost",
			Usage:  "Consul host to connect to",
			EnvVar: "CONSUL_HOST",
		},
		cli.StringFlag{
			Name:   "port, p",
			Value:  "8500",
			Usage:  "Consul HTTP API port to connect to",
			EnvVar: "CONSUL_PORT",
		},
		cli.StringFlag{
			Name:   "yaml, y",
			Usage:  "URL of yaml file",
			EnvVar: "YAML_URL",
		},
	}

	app.Action = func(c *cli.Context) {
		// Get a new client
		config := api.DefaultConfig()
		config.Address = c.String("host") + ":" + c.String("port")

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

		m := make(map[interface{}]interface{})

		// assume local file path for now
		data, err := ioutil.ReadFile(c.String("yaml"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
		}

		err = yaml.Unmarshal(data, &m)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
		}
		fmt.Printf("--- m:\n%v\n\n", m)
	}

	app.Run(os.Args)
}

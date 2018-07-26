package main

import (
	"github.com/hashicorp/terraform/plugin"
	c "github.com/spothero/terraform-provider-kafka-connect/connect"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: c.Provider})
}

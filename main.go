package main

import (
	"context"

	"github.com/cbsinteractive/terraform-provider-azurermfutures/azurermfutures"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func main() {
	tfsdk.Serve(context.Background(), azurermfutures.NewProvider, tfsdk.ServeOpts{
		Name: "azurermfutures",
	})
}

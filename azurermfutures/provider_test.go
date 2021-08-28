package azurermfutures

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var testAccProviderFactories map[string]func() (tfprotov6.ProviderServer, error)

func init() {
	testAccProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"azurermfutures": func() (tfprotov6.ProviderServer, error) {
			return tfsdk.NewProtocol6Server(NewProvider()), nil
		},
	}
}

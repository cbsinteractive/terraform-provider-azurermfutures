package azurermfutures

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

type myProvider struct{}

// GetSchema returns the schema for this provider's configuration. If
// this provider has no configuration, return an empty schema.Schema.
func (p myProvider) GetSchema(context.Context) (schema.Schema, []*tfprotov6.Diagnostic) {
	log.Println("[DEBUG] Inside myProvider.GetSchema")
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"arm_subscription_id": {
				Type:      types.StringType,
				Optional:  true,
				Sensitive: true,
			},
			"arm_tenant_id": {
				Type:      types.StringType,
				Optional:  true,
				Sensitive: true,
			},
			"arm_client_id": {
				Type:      types.StringType,
				Optional:  true,
				Sensitive: true,
			},
			"arm_client_secret": {
				Type:      types.StringType,
				Optional:  true,
				Sensitive: true,
			},
		},
	}, nil
}

// Configure is called at the beginning of the provider lifecycle, when
// Terraform sends to the provider the values the user specified in the
// provider configuration block. These are supplied in the
// ConfigureProviderRequest argument.
// Values from provider configuration are often used to initialise an
// API client, which should be stored on the struct implementing the
// Provider interface.
func (p myProvider) Configure(context.Context, tfsdk.ConfigureProviderRequest, *tfsdk.ConfigureProviderResponse) {
	log.Println("[DEBUG] Inside myProvider.Configure")
}

// GetResources returns a map of the resource types this provider
// supports.
func (p myProvider) GetResources(context.Context) (map[string]tfsdk.ResourceType, []*tfprotov6.Diagnostic) {
	log.Println("[DEBUG] Inside myProvider.GetResources")
	return map[string]tfsdk.ResourceType{
		"azurermfutures_streamanalytics_job": streamAnalyticsJobType{},
	}, nil
}

// GetDataSources returns a map of the data source types this provider
// supports.
func (p myProvider) GetDataSources(context.Context) (map[string]tfsdk.DataSourceType, []*tfprotov6.Diagnostic) {
	log.Println("[DEBUG] Inside myProvider.GetDataSources")
	return nil, nil
}

func newProvider() tfsdk.Provider {
	log.Println("[DEBUG] Inside newProvider")
	return myProvider{}
}

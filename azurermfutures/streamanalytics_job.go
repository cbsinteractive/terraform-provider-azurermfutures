package azurermfutures

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// NewResource instantiates a new Resource of this ResourceType.
func (t resourcesStreamAnalyticsJob) NewResource(ctx context.Context, p tfsdk.Provider) (tfsdk.Resource, []*tfprotov6.Diagnostic) {
	log.Println("[DEBUG] Inside streamAnalyticsJob.NewResource")
	return resourcesStreamAnalyticsJob{
		p: *(p.(*provider)),
	}, nil
}

type resourcesStreamAnalyticsJob struct {
	p provider
}

// GetSchema returns the schema for this resource.
func (t resourcesStreamAnalyticsJob) GetSchema(context.Context) (schema.Schema, []*tfprotov6.Diagnostic) {
	log.Println("[DEBUG] Inside streamAnalyticsJob.GetSchema")
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
		},
	}, nil
}

// Create is called when the provider must create a new resource. Config
// and planned state values should be read from the
// CreateResourceRequest and new state values set on the
// CreateResourceResponse.
func (r resourcesStreamAnalyticsJob) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	log.Println("[DEBUG] Inside streamAnalyticsJob.Create")
	if !r.p.configured {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Provider not configured",
			Detail:   "Terraform attempted to create a resource before the provider was properly configured.",
		})
		return
	}
	var plan StreamAnalyticsJob
	err := req.Plan.Get(ctx, &plan)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error reading plan",
			Detail:   "An unexpected error was encountered while reading the plan: " + err.Error(),
		})
		return
	}
	result := StreamAnalyticsJob{
		ID: types.String{Value: "some-id"},
	}
	err = resp.State.Set(ctx, result)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error setting state",
			Detail:   "Could not set state, unexpected error: " + err.Error(),
		})
		return
	}
}

// Read is called when the provider must read resource values in order
// to update state. Planned state values should be read from the
// ReadResourceRequest and new state values set on the
// ReadResourceResponse.
func (t resourcesStreamAnalyticsJob) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	log.Println("[DEBUG] Inside streamAnalyticsJob.Read")
	var state StreamAnalyticsJob
	err := req.State.Get(ctx, &state)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error reading state",
			Detail:   "An unexpected error was encountered while reading the state: " + err.Error(),
		})
	}
	log.Printf("[DEBUG] state: %+v\n", state)
}

// Update is called to update the state of the resource. Config, planned
// state, and prior state values should be read from the
// UpdateResourceRequest and new state values set on the
// UpdateResourceResponse.
func (t resourcesStreamAnalyticsJob) Update(context.Context, tfsdk.UpdateResourceRequest, *tfsdk.UpdateResourceResponse) {
	log.Println("[DEBUG] Inside streamAnalyticsJob.Update")
}

// Delete is called when the provider must delete the resource. Config
// values may be read from the DeleteResourceRequest.
func (t resourcesStreamAnalyticsJob) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	log.Println("[DEBUG] Inside streamAnalyticsJob.Delete")
	var state StreamAnalyticsJob
	err := req.State.Get(ctx, &state)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error reading configuration",
			Detail:   "An unexpected error was encountered while reading the configuration: " + err.Error(),
		})
		return
	}
	log.Printf("[DEBUG] state: %+v\n", state)
	// Remove resource from state
	resp.State.RemoveResource(ctx)
}

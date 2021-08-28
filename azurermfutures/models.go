package azurermfutures

import "github.com/hashicorp/terraform-plugin-framework/types"

type StreamAnalyticsJob struct {
	ID types.String `tfsdk:"id"`
}

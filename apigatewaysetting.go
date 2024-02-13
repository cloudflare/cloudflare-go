// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// APIGatewaySettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAPIGatewaySettingService] method
// instead.
type APIGatewaySettingService struct {
	Options          []option.RequestOption
	SchemaValidation *APIGatewaySettingSchemaValidationService
}

// NewAPIGatewaySettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIGatewaySettingService(opts ...option.RequestOption) (r *APIGatewaySettingService) {
	r = &APIGatewaySettingService{}
	r.Options = opts
	r.SchemaValidation = NewAPIGatewaySettingSchemaValidationService(opts...)
	return
}

// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// APIGatewayService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAPIGatewayService] method instead.
type APIGatewayService struct {
	Options          []option.RequestOption
	Configurations   *APIGatewayConfigurationService
	Discoveries      *APIGatewayDiscoveryService
	Operations       *APIGatewayOperationService
	Schemas          *APIGatewaySchemaService
	Settings         *APIGatewaySettingService
	UserSchemas      *APIGatewayUserSchemaService
	Discovery        *APIGatewayDiscoveryService
	SchemaValidation *APIGatewaySchemaValidationService
}

// NewAPIGatewayService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIGatewayService(opts ...option.RequestOption) (r *APIGatewayService) {
	r = &APIGatewayService{}
	r.Options = opts
	r.Configurations = NewAPIGatewayConfigurationService(opts...)
	r.Discoveries = NewAPIGatewayDiscoveryService(opts...)
	r.Operations = NewAPIGatewayOperationService(opts...)
	r.Schemas = NewAPIGatewaySchemaService(opts...)
	r.Settings = NewAPIGatewaySettingService(opts...)
	r.UserSchemas = NewAPIGatewayUserSchemaService(opts...)
	r.Discovery = NewAPIGatewayDiscoveryService(opts...)
	r.SchemaValidation = NewAPIGatewaySchemaValidationService(opts...)
	return
}

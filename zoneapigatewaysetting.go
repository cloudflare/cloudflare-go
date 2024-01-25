// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneAPIGatewaySettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneAPIGatewaySettingService]
// method instead.
type ZoneAPIGatewaySettingService struct {
	Options           []option.RequestOption
	SchemaValidations *ZoneAPIGatewaySettingSchemaValidationService
}

// NewZoneAPIGatewaySettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAPIGatewaySettingService(opts ...option.RequestOption) (r *ZoneAPIGatewaySettingService) {
	r = &ZoneAPIGatewaySettingService{}
	r.Options = opts
	r.SchemaValidations = NewZoneAPIGatewaySettingSchemaValidationService(opts...)
	return
}

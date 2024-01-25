// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneAPIGatewayService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneAPIGatewayService] method
// instead.
type ZoneAPIGatewayService struct {
	Options        []option.RequestOption
	Configurations *ZoneAPIGatewayConfigurationService
	Discovery      *ZoneAPIGatewayDiscoveryService
	Schemas        *ZoneAPIGatewaySchemaService
	Settings       *ZoneAPIGatewaySettingService
	UserSchemas    *ZoneAPIGatewayUserSchemaService
}

// NewZoneAPIGatewayService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneAPIGatewayService(opts ...option.RequestOption) (r *ZoneAPIGatewayService) {
	r = &ZoneAPIGatewayService{}
	r.Options = opts
	r.Configurations = NewZoneAPIGatewayConfigurationService(opts...)
	r.Discovery = NewZoneAPIGatewayDiscoveryService(opts...)
	r.Schemas = NewZoneAPIGatewaySchemaService(opts...)
	r.Settings = NewZoneAPIGatewaySettingService(opts...)
	r.UserSchemas = NewZoneAPIGatewayUserSchemaService(opts...)
	return
}

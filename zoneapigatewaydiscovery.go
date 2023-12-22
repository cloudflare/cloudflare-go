// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneAPIGatewayDiscoveryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneAPIGatewayDiscoveryService] method instead.
type ZoneAPIGatewayDiscoveryService struct {
	Options []option.RequestOption
}

// NewZoneAPIGatewayDiscoveryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAPIGatewayDiscoveryService(opts ...option.RequestOption) (r *ZoneAPIGatewayDiscoveryService) {
	r = &ZoneAPIGatewayDiscoveryService{}
	r.Options = opts
	return
}

// Retrieve the most up to date view of API Discovery on a zone.
func (r *ZoneAPIGatewayDiscoveryService) APIShieldEndpointManagementGetAPIDiscoveryResultsForAZone(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SchemaResponseDiscovery, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/discovery", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SchemaResponseDiscovery struct {
	Errors   []SchemaResponseDiscoveryError   `json:"errors"`
	Messages []SchemaResponseDiscoveryMessage `json:"messages"`
	Result   SchemaResponseDiscoveryResult    `json:"result"`
	// Whether the API call was successful
	Success SchemaResponseDiscoverySuccess `json:"success"`
	JSON    schemaResponseDiscoveryJSON    `json:"-"`
}

// schemaResponseDiscoveryJSON contains the JSON metadata for the struct
// [SchemaResponseDiscovery]
type schemaResponseDiscoveryJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaResponseDiscovery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemaResponseDiscoveryError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    schemaResponseDiscoveryErrorJSON `json:"-"`
}

// schemaResponseDiscoveryErrorJSON contains the JSON metadata for the struct
// [SchemaResponseDiscoveryError]
type schemaResponseDiscoveryErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaResponseDiscoveryError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemaResponseDiscoveryMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    schemaResponseDiscoveryMessageJSON `json:"-"`
}

// schemaResponseDiscoveryMessageJSON contains the JSON metadata for the struct
// [SchemaResponseDiscoveryMessage]
type schemaResponseDiscoveryMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaResponseDiscoveryMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemaResponseDiscoveryResult struct {
	Schemas   []interface{}                     `json:"schemas"`
	Timestamp string                            `json:"timestamp"`
	JSON      schemaResponseDiscoveryResultJSON `json:"-"`
}

// schemaResponseDiscoveryResultJSON contains the JSON metadata for the struct
// [SchemaResponseDiscoveryResult]
type schemaResponseDiscoveryResultJSON struct {
	Schemas     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaResponseDiscoveryResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemaResponseDiscoverySuccess bool

const (
	SchemaResponseDiscoverySuccessTrue SchemaResponseDiscoverySuccess = true
)

// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
	Options    []option.RequestOption
	Operations *ZoneAPIGatewayDiscoveryOperationService
}

// NewZoneAPIGatewayDiscoveryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAPIGatewayDiscoveryService(opts ...option.RequestOption) (r *ZoneAPIGatewayDiscoveryService) {
	r = &ZoneAPIGatewayDiscoveryService{}
	r.Options = opts
	r.Operations = NewZoneAPIGatewayDiscoveryOperationService(opts...)
	return
}

// Retrieve the most up to date view of discovered operations, rendered as OpenAPI
// schemas
func (r *ZoneAPIGatewayDiscoveryService) APIShieldEndpointManagementGetAPIDiscoveryResultsForAZone(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/discovery", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponse struct {
	Errors   []ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseError   `json:"errors"`
	Messages []ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseMessage `json:"messages"`
	Result   ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                                                                                         `json:"success"`
	JSON    zoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponse]
type zoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseError struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    zoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseErrorJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseError]
type zoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseMessage struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    zoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseMessageJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseMessage]
type zoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseResult struct {
	Schemas   []interface{}                                                                                      `json:"schemas"`
	Timestamp time.Time                                                                                          `json:"timestamp" format:"date-time"`
	JSON      zoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseResultJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseResult]
type zoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseResultJSON struct {
	Schemas     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

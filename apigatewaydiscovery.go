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

// APIGatewayDiscoveryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAPIGatewayDiscoveryService]
// method instead.
type APIGatewayDiscoveryService struct {
	Options []option.RequestOption
}

// NewAPIGatewayDiscoveryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIGatewayDiscoveryService(opts ...option.RequestOption) (r *APIGatewayDiscoveryService) {
	r = &APIGatewayDiscoveryService{}
	r.Options = opts
	return
}

// Retrieve the most up to date view of discovered operations, rendered as OpenAPI
// schemas
func (r *APIGatewayDiscoveryService) APIShieldEndpointManagementGetAPIDiscoveryResultsForAZone(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *APIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env APIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseEnvelope
	path := fmt.Sprintf("zones/%s/api_gateway/discovery", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type APIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponse struct {
	Schemas   []interface{}                                                                            `json:"schemas"`
	Timestamp time.Time                                                                                `json:"timestamp" format:"date-time"`
	JSON      apiGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseJSON `json:"-"`
}

// apiGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseJSON
// contains the JSON metadata for the struct
// [APIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponse]
type apiGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseJSON struct {
	Schemas     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseEnvelope struct {
	Errors   APIShieldMessages                                                                    `json:"errors,required"`
	Messages APIShieldMessages                                                                    `json:"messages,required"`
	Result   APIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                                                                             `json:"success,required"`
	JSON    apiGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseEnvelopeJSON `json:"-"`
}

// apiGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [APIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseEnvelope]
type apiGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayDiscoveryAPIShieldEndpointManagementGetAPIDiscoveryResultsForAZoneResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// APIGatewayDiscoveryOperationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAPIGatewayDiscoveryOperationService] method instead.
type APIGatewayDiscoveryOperationService struct {
	Options []option.RequestOption
}

// NewAPIGatewayDiscoveryOperationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAPIGatewayDiscoveryOperationService(opts ...option.RequestOption) (r *APIGatewayDiscoveryOperationService) {
	r = &APIGatewayDiscoveryOperationService{}
	r.Options = opts
	return
}

// Update the `state` on one or more discovered operations
func (r *APIGatewayDiscoveryOperationService) Update(ctx context.Context, zoneID string, body APIGatewayDiscoveryOperationUpdateParams, opts ...option.RequestOption) (res *APIGatewayDiscoveryOperationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env APIGatewayDiscoveryOperationUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/api_gateway/discovery/operations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type APIGatewayDiscoveryOperationUpdateResponse map[string]APIGatewayDiscoveryOperationUpdateResponse

type APIGatewayDiscoveryOperationUpdateParams struct {
	Body param.Field[map[string]APIGatewayDiscoveryOperationUpdateParamsBody] `json:"body,required"`
}

func (r APIGatewayDiscoveryOperationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Operation ID to patch state mappings
type APIGatewayDiscoveryOperationUpdateParamsBody struct {
	// Mark state of operation in API Discovery
	//
	// - `review` - Mark operation as for review
	// - `ignored` - Mark operation as ignored
	State param.Field[APIGatewayDiscoveryOperationUpdateParamsBodyState] `json:"state"`
}

func (r APIGatewayDiscoveryOperationUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mark state of operation in API Discovery
//
// - `review` - Mark operation as for review
// - `ignored` - Mark operation as ignored
type APIGatewayDiscoveryOperationUpdateParamsBodyState string

const (
	APIGatewayDiscoveryOperationUpdateParamsBodyStateReview  APIGatewayDiscoveryOperationUpdateParamsBodyState = "review"
	APIGatewayDiscoveryOperationUpdateParamsBodyStateIgnored APIGatewayDiscoveryOperationUpdateParamsBodyState = "ignored"
)

type APIGatewayDiscoveryOperationUpdateResponseEnvelope struct {
	Errors   APIShieldMessages                          `json:"errors,required"`
	Messages APIShieldMessages                          `json:"messages,required"`
	Result   APIGatewayDiscoveryOperationUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                                   `json:"success,required"`
	JSON    apiGatewayDiscoveryOperationUpdateResponseEnvelopeJSON `json:"-"`
}

// apiGatewayDiscoveryOperationUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [APIGatewayDiscoveryOperationUpdateResponseEnvelope]
type apiGatewayDiscoveryOperationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIGatewayDiscoveryOperationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

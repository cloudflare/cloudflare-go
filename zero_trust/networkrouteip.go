// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// NetworkRouteIPService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewNetworkRouteIPService] method
// instead.
type NetworkRouteIPService struct {
	Options []option.RequestOption
}

// NewNetworkRouteIPService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkRouteIPService(opts ...option.RequestOption) (r *NetworkRouteIPService) {
	r = &NetworkRouteIPService{}
	r.Options = opts
	return
}

// Fetches routes that contain the given IP address.
func (r *NetworkRouteIPService) Get(ctx context.Context, ip string, params NetworkRouteIPGetParams, opts ...option.RequestOption) (res *TunnelTeamnet, err error) {
	opts = append(r.Options[:], opts...)
	var env NetworkRouteIPGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/routes/ip/%s", params.AccountID, ip)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NetworkRouteIPGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `query:"virtual_network_id"`
}

// URLQuery serializes [NetworkRouteIPGetParams]'s query parameters as
// `url.Values`.
func (r NetworkRouteIPGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NetworkRouteIPGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   TunnelTeamnet                `json:"result,required"`
	// Whether the API call was successful
	Success NetworkRouteIPGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkRouteIPGetResponseEnvelopeJSON    `json:"-"`
}

// networkRouteIPGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetworkRouteIPGetResponseEnvelope]
type networkRouteIPGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteIPGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteIPGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkRouteIPGetResponseEnvelopeSuccess bool

const (
	NetworkRouteIPGetResponseEnvelopeSuccessTrue NetworkRouteIPGetResponseEnvelopeSuccess = true
)

func (r NetworkRouteIPGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkRouteIPGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

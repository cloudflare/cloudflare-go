// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// NetworkSubnetInitialResolvedIPService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkSubnetInitialResolvedIPService] method instead.
type NetworkSubnetInitialResolvedIPService struct {
	Options []option.RequestOption
}

// NewNetworkSubnetInitialResolvedIPService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewNetworkSubnetInitialResolvedIPService(opts ...option.RequestOption) (r *NetworkSubnetInitialResolvedIPService) {
	r = &NetworkSubnetInitialResolvedIPService{}
	r.Options = opts
	return
}

// Updates the CIDR for the account's default Initial Resolved IP Subnet of the
// given address family. The new CIDR must not conflict with existing private
// routes in the account.
func (r *NetworkSubnetInitialResolvedIPService) Update(ctx context.Context, addressFamily NetworkSubnetInitialResolvedIPUpdateParamsAddressFamily, params NetworkSubnetInitialResolvedIPUpdateParams, opts ...option.RequestOption) (res *Subnet, err error) {
	var env NetworkSubnetInitialResolvedIPUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/subnets/initial_resolved_ip/%v", params.AccountID, addressFamily)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the account's default Initial Resolved IP Subnet for the given address
// family.
func (r *NetworkSubnetInitialResolvedIPService) Get(ctx context.Context, addressFamily NetworkSubnetInitialResolvedIPGetParamsAddressFamily, query NetworkSubnetInitialResolvedIPGetParams, opts ...option.RequestOption) (res *Subnet, err error) {
	var env NetworkSubnetInitialResolvedIPGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/subnets/initial_resolved_ip/%v", query.AccountID, addressFamily)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type NetworkSubnetInitialResolvedIPUpdateParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// An optional description of the subnet.
	Comment param.Field[string] `json:"comment"`
	// A user-friendly name for the subnet.
	Name param.Field[string] `json:"name"`
	// The private IPv4 or IPv6 range defining the subnet, in CIDR notation.
	Network param.Field[string] `json:"network"`
}

func (r NetworkSubnetInitialResolvedIPUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// IP address family, either `v4` (IPv4) or `v6` (IPv6)
type NetworkSubnetInitialResolvedIPUpdateParamsAddressFamily string

const (
	NetworkSubnetInitialResolvedIPUpdateParamsAddressFamilyV4 NetworkSubnetInitialResolvedIPUpdateParamsAddressFamily = "v4"
	NetworkSubnetInitialResolvedIPUpdateParamsAddressFamilyV6 NetworkSubnetInitialResolvedIPUpdateParamsAddressFamily = "v6"
)

func (r NetworkSubnetInitialResolvedIPUpdateParamsAddressFamily) IsKnown() bool {
	switch r {
	case NetworkSubnetInitialResolvedIPUpdateParamsAddressFamilyV4, NetworkSubnetInitialResolvedIPUpdateParamsAddressFamilyV6:
		return true
	}
	return false
}

type NetworkSubnetInitialResolvedIPUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   Subnet                `json:"result" api:"required"`
	// Whether the API call was successful
	Success NetworkSubnetInitialResolvedIPUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    networkSubnetInitialResolvedIPUpdateResponseEnvelopeJSON    `json:"-"`
}

// networkSubnetInitialResolvedIPUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [NetworkSubnetInitialResolvedIPUpdateResponseEnvelope]
type networkSubnetInitialResolvedIPUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkSubnetInitialResolvedIPUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkSubnetInitialResolvedIPUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkSubnetInitialResolvedIPUpdateResponseEnvelopeSuccess bool

const (
	NetworkSubnetInitialResolvedIPUpdateResponseEnvelopeSuccessTrue NetworkSubnetInitialResolvedIPUpdateResponseEnvelopeSuccess = true
)

func (r NetworkSubnetInitialResolvedIPUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkSubnetInitialResolvedIPUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NetworkSubnetInitialResolvedIPGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

// IP address family, either `v4` (IPv4) or `v6` (IPv6)
type NetworkSubnetInitialResolvedIPGetParamsAddressFamily string

const (
	NetworkSubnetInitialResolvedIPGetParamsAddressFamilyV4 NetworkSubnetInitialResolvedIPGetParamsAddressFamily = "v4"
	NetworkSubnetInitialResolvedIPGetParamsAddressFamilyV6 NetworkSubnetInitialResolvedIPGetParamsAddressFamily = "v6"
)

func (r NetworkSubnetInitialResolvedIPGetParamsAddressFamily) IsKnown() bool {
	switch r {
	case NetworkSubnetInitialResolvedIPGetParamsAddressFamilyV4, NetworkSubnetInitialResolvedIPGetParamsAddressFamilyV6:
		return true
	}
	return false
}

type NetworkSubnetInitialResolvedIPGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   Subnet                `json:"result" api:"required"`
	// Whether the API call was successful
	Success NetworkSubnetInitialResolvedIPGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    networkSubnetInitialResolvedIPGetResponseEnvelopeJSON    `json:"-"`
}

// networkSubnetInitialResolvedIPGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [NetworkSubnetInitialResolvedIPGetResponseEnvelope]
type networkSubnetInitialResolvedIPGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkSubnetInitialResolvedIPGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkSubnetInitialResolvedIPGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkSubnetInitialResolvedIPGetResponseEnvelopeSuccess bool

const (
	NetworkSubnetInitialResolvedIPGetResponseEnvelopeSuccessTrue NetworkSubnetInitialResolvedIPGetResponseEnvelopeSuccess = true
)

func (r NetworkSubnetInitialResolvedIPGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkSubnetInitialResolvedIPGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

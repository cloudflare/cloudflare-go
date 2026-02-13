// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// NetworkSubnetWARPService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkSubnetWARPService] method instead.
type NetworkSubnetWARPService struct {
	Options []option.RequestOption
}

// NewNetworkSubnetWARPService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNetworkSubnetWARPService(opts ...option.RequestOption) (r *NetworkSubnetWARPService) {
	r = &NetworkSubnetWARPService{}
	r.Options = opts
	return
}

// Create a WARP IP assignment subnet. Currently, only IPv4 subnets can be created.
//
// **Network constraints:**
//
// - The network must be within one of the following private IP ranges:
//   - `10.0.0.0/8` (RFC 1918)
//   - `172.16.0.0/12` (RFC 1918)
//   - `192.168.0.0/16` (RFC 1918)
//   - `100.64.0.0/10` (RFC 6598 - CGNAT)
//   - The subnet must have a prefix length of `/24` or larger (e.g., `/16`, `/20`,
//     `/24` are valid; `/25`, `/28` are not)
func (r *NetworkSubnetWARPService) New(ctx context.Context, params NetworkSubnetWARPNewParams, opts ...option.RequestOption) (res *NetworkSubnetWARPNewResponse, err error) {
	var env NetworkSubnetWARPNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/subnets/warp", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a WARP IP assignment subnet. This operation is idempotent - deleting an
// already-deleted or non-existent subnet will return success with a null result.
func (r *NetworkSubnetWARPService) Delete(ctx context.Context, subnetID string, body NetworkSubnetWARPDeleteParams, opts ...option.RequestOption) (res *NetworkSubnetWARPDeleteResponse, err error) {
	var env NetworkSubnetWARPDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if subnetID == "" {
		err = errors.New("missing required subnet_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/subnets/warp/%s", body.AccountID, subnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a WARP IP assignment subnet.
//
// **Update constraints:**
//
//   - The `network` field cannot be modified for WARP subnets. Only `name`,
//     `comment`, and `is_default_network` can be updated.
//   - IPv6 subnets cannot be updated
func (r *NetworkSubnetWARPService) Edit(ctx context.Context, subnetID string, params NetworkSubnetWARPEditParams, opts ...option.RequestOption) (res *NetworkSubnetWARPEditResponse, err error) {
	var env NetworkSubnetWARPEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if subnetID == "" {
		err = errors.New("missing required subnet_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/subnets/warp/%s", params.AccountID, subnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a WARP IP assignment subnet.
func (r *NetworkSubnetWARPService) Get(ctx context.Context, subnetID string, query NetworkSubnetWARPGetParams, opts ...option.RequestOption) (res *NetworkSubnetWARPGetResponse, err error) {
	var env NetworkSubnetWARPGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if subnetID == "" {
		err = errors.New("missing required subnet_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/subnets/warp/%s", query.AccountID, subnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NetworkSubnetWARPNewResponse struct {
	// The UUID of the subnet.
	ID string `json:"id" format:"uuid"`
	// An optional description of the subnet.
	Comment string `json:"comment"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// If `true`, this is the default subnet for the account. There can only be one
	// default subnet per account.
	IsDefaultNetwork bool `json:"is_default_network"`
	// A user-friendly name for the subnet.
	Name string `json:"name"`
	// The private IPv4 or IPv6 range defining the subnet, in CIDR notation.
	Network string `json:"network"`
	// The type of subnet.
	SubnetType NetworkSubnetWARPNewResponseSubnetType `json:"subnet_type"`
	JSON       networkSubnetWARPNewResponseJSON       `json:"-"`
}

// networkSubnetWARPNewResponseJSON contains the JSON metadata for the struct
// [NetworkSubnetWARPNewResponse]
type networkSubnetWARPNewResponseJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	DeletedAt        apijson.Field
	IsDefaultNetwork apijson.Field
	Name             apijson.Field
	Network          apijson.Field
	SubnetType       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NetworkSubnetWARPNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkSubnetWARPNewResponseJSON) RawJSON() string {
	return r.raw
}

// The type of subnet.
type NetworkSubnetWARPNewResponseSubnetType string

const (
	NetworkSubnetWARPNewResponseSubnetTypeCloudflareSource NetworkSubnetWARPNewResponseSubnetType = "cloudflare_source"
	NetworkSubnetWARPNewResponseSubnetTypeWARP             NetworkSubnetWARPNewResponseSubnetType = "warp"
)

func (r NetworkSubnetWARPNewResponseSubnetType) IsKnown() bool {
	switch r {
	case NetworkSubnetWARPNewResponseSubnetTypeCloudflareSource, NetworkSubnetWARPNewResponseSubnetTypeWARP:
		return true
	}
	return false
}

type NetworkSubnetWARPDeleteResponse struct {
	// The UUID of the subnet.
	ID string `json:"id" format:"uuid"`
	// An optional description of the subnet.
	Comment string `json:"comment"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// If `true`, this is the default subnet for the account. There can only be one
	// default subnet per account.
	IsDefaultNetwork bool `json:"is_default_network"`
	// A user-friendly name for the subnet.
	Name string `json:"name"`
	// The private IPv4 or IPv6 range defining the subnet, in CIDR notation.
	Network string `json:"network"`
	// The type of subnet.
	SubnetType NetworkSubnetWARPDeleteResponseSubnetType `json:"subnet_type"`
	JSON       networkSubnetWARPDeleteResponseJSON       `json:"-"`
}

// networkSubnetWARPDeleteResponseJSON contains the JSON metadata for the struct
// [NetworkSubnetWARPDeleteResponse]
type networkSubnetWARPDeleteResponseJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	DeletedAt        apijson.Field
	IsDefaultNetwork apijson.Field
	Name             apijson.Field
	Network          apijson.Field
	SubnetType       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NetworkSubnetWARPDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkSubnetWARPDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The type of subnet.
type NetworkSubnetWARPDeleteResponseSubnetType string

const (
	NetworkSubnetWARPDeleteResponseSubnetTypeCloudflareSource NetworkSubnetWARPDeleteResponseSubnetType = "cloudflare_source"
	NetworkSubnetWARPDeleteResponseSubnetTypeWARP             NetworkSubnetWARPDeleteResponseSubnetType = "warp"
)

func (r NetworkSubnetWARPDeleteResponseSubnetType) IsKnown() bool {
	switch r {
	case NetworkSubnetWARPDeleteResponseSubnetTypeCloudflareSource, NetworkSubnetWARPDeleteResponseSubnetTypeWARP:
		return true
	}
	return false
}

type NetworkSubnetWARPEditResponse struct {
	// The UUID of the subnet.
	ID string `json:"id" format:"uuid"`
	// An optional description of the subnet.
	Comment string `json:"comment"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// If `true`, this is the default subnet for the account. There can only be one
	// default subnet per account.
	IsDefaultNetwork bool `json:"is_default_network"`
	// A user-friendly name for the subnet.
	Name string `json:"name"`
	// The private IPv4 or IPv6 range defining the subnet, in CIDR notation.
	Network string `json:"network"`
	// The type of subnet.
	SubnetType NetworkSubnetWARPEditResponseSubnetType `json:"subnet_type"`
	JSON       networkSubnetWARPEditResponseJSON       `json:"-"`
}

// networkSubnetWARPEditResponseJSON contains the JSON metadata for the struct
// [NetworkSubnetWARPEditResponse]
type networkSubnetWARPEditResponseJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	DeletedAt        apijson.Field
	IsDefaultNetwork apijson.Field
	Name             apijson.Field
	Network          apijson.Field
	SubnetType       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NetworkSubnetWARPEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkSubnetWARPEditResponseJSON) RawJSON() string {
	return r.raw
}

// The type of subnet.
type NetworkSubnetWARPEditResponseSubnetType string

const (
	NetworkSubnetWARPEditResponseSubnetTypeCloudflareSource NetworkSubnetWARPEditResponseSubnetType = "cloudflare_source"
	NetworkSubnetWARPEditResponseSubnetTypeWARP             NetworkSubnetWARPEditResponseSubnetType = "warp"
)

func (r NetworkSubnetWARPEditResponseSubnetType) IsKnown() bool {
	switch r {
	case NetworkSubnetWARPEditResponseSubnetTypeCloudflareSource, NetworkSubnetWARPEditResponseSubnetTypeWARP:
		return true
	}
	return false
}

type NetworkSubnetWARPGetResponse struct {
	// The UUID of the subnet.
	ID string `json:"id" format:"uuid"`
	// An optional description of the subnet.
	Comment string `json:"comment"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// If `true`, this is the default subnet for the account. There can only be one
	// default subnet per account.
	IsDefaultNetwork bool `json:"is_default_network"`
	// A user-friendly name for the subnet.
	Name string `json:"name"`
	// The private IPv4 or IPv6 range defining the subnet, in CIDR notation.
	Network string `json:"network"`
	// The type of subnet.
	SubnetType NetworkSubnetWARPGetResponseSubnetType `json:"subnet_type"`
	JSON       networkSubnetWARPGetResponseJSON       `json:"-"`
}

// networkSubnetWARPGetResponseJSON contains the JSON metadata for the struct
// [NetworkSubnetWARPGetResponse]
type networkSubnetWARPGetResponseJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	DeletedAt        apijson.Field
	IsDefaultNetwork apijson.Field
	Name             apijson.Field
	Network          apijson.Field
	SubnetType       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NetworkSubnetWARPGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkSubnetWARPGetResponseJSON) RawJSON() string {
	return r.raw
}

// The type of subnet.
type NetworkSubnetWARPGetResponseSubnetType string

const (
	NetworkSubnetWARPGetResponseSubnetTypeCloudflareSource NetworkSubnetWARPGetResponseSubnetType = "cloudflare_source"
	NetworkSubnetWARPGetResponseSubnetTypeWARP             NetworkSubnetWARPGetResponseSubnetType = "warp"
)

func (r NetworkSubnetWARPGetResponseSubnetType) IsKnown() bool {
	switch r {
	case NetworkSubnetWARPGetResponseSubnetTypeCloudflareSource, NetworkSubnetWARPGetResponseSubnetTypeWARP:
		return true
	}
	return false
}

type NetworkSubnetWARPNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// A user-friendly name for the subnet.
	Name param.Field[string] `json:"name,required"`
	// The private IPv4 or IPv6 range defining the subnet, in CIDR notation.
	Network param.Field[string] `json:"network,required"`
	// An optional description of the subnet.
	Comment param.Field[string] `json:"comment"`
	// If `true`, this is the default subnet for the account. There can only be one
	// default subnet per account.
	IsDefaultNetwork param.Field[bool] `json:"is_default_network"`
}

func (r NetworkSubnetWARPNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkSubnetWARPNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo        `json:"errors,required"`
	Messages []shared.ResponseInfo        `json:"messages,required"`
	Result   NetworkSubnetWARPNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success NetworkSubnetWARPNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkSubnetWARPNewResponseEnvelopeJSON    `json:"-"`
}

// networkSubnetWARPNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkSubnetWARPNewResponseEnvelope]
type networkSubnetWARPNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkSubnetWARPNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkSubnetWARPNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkSubnetWARPNewResponseEnvelopeSuccess bool

const (
	NetworkSubnetWARPNewResponseEnvelopeSuccessTrue NetworkSubnetWARPNewResponseEnvelopeSuccess = true
)

func (r NetworkSubnetWARPNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkSubnetWARPNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NetworkSubnetWARPDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type NetworkSubnetWARPDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   NetworkSubnetWARPDeleteResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success NetworkSubnetWARPDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkSubnetWARPDeleteResponseEnvelopeJSON    `json:"-"`
}

// networkSubnetWARPDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkSubnetWARPDeleteResponseEnvelope]
type networkSubnetWARPDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkSubnetWARPDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkSubnetWARPDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkSubnetWARPDeleteResponseEnvelopeSuccess bool

const (
	NetworkSubnetWARPDeleteResponseEnvelopeSuccessTrue NetworkSubnetWARPDeleteResponseEnvelopeSuccess = true
)

func (r NetworkSubnetWARPDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkSubnetWARPDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NetworkSubnetWARPEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// An optional description of the subnet.
	Comment param.Field[string] `json:"comment"`
	// If `true`, this is the default subnet for the account. There can only be one
	// default subnet per account.
	IsDefaultNetwork param.Field[bool] `json:"is_default_network"`
	// A user-friendly name for the subnet.
	Name param.Field[string] `json:"name"`
	// The private IPv4 or IPv6 range defining the subnet, in CIDR notation.
	Network param.Field[string] `json:"network"`
}

func (r NetworkSubnetWARPEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkSubnetWARPEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []shared.ResponseInfo         `json:"messages,required"`
	Result   NetworkSubnetWARPEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success NetworkSubnetWARPEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkSubnetWARPEditResponseEnvelopeJSON    `json:"-"`
}

// networkSubnetWARPEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkSubnetWARPEditResponseEnvelope]
type networkSubnetWARPEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkSubnetWARPEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkSubnetWARPEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkSubnetWARPEditResponseEnvelopeSuccess bool

const (
	NetworkSubnetWARPEditResponseEnvelopeSuccessTrue NetworkSubnetWARPEditResponseEnvelopeSuccess = true
)

func (r NetworkSubnetWARPEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkSubnetWARPEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NetworkSubnetWARPGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type NetworkSubnetWARPGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo        `json:"errors,required"`
	Messages []shared.ResponseInfo        `json:"messages,required"`
	Result   NetworkSubnetWARPGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success NetworkSubnetWARPGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkSubnetWARPGetResponseEnvelopeJSON    `json:"-"`
}

// networkSubnetWARPGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkSubnetWARPGetResponseEnvelope]
type networkSubnetWARPGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkSubnetWARPGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkSubnetWARPGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkSubnetWARPGetResponseEnvelopeSuccess bool

const (
	NetworkSubnetWARPGetResponseEnvelopeSuccessTrue NetworkSubnetWARPGetResponseEnvelopeSuccess = true
)

func (r NetworkSubnetWARPGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkSubnetWARPGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

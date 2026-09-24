// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// BGPFilterProfileService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPFilterProfileService] method instead.
type BGPFilterProfileService struct {
	Options []option.RequestOption
}

// NewBGPFilterProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBGPFilterProfileService(opts ...option.RequestOption) (r *BGPFilterProfileService) {
	r = &BGPFilterProfileService{}
	r.Options = opts
	return
}

// Creates a new BGP filter profile for an account.
func (r *BGPFilterProfileService) New(ctx context.Context, params BGPFilterProfileNewParams, opts ...option.RequestOption) (res *BGPFilterProfileNewResponse, err error) {
	var env BGPFilterProfileNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/bgp/filter_profiles", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates a BGP filter profile. Omitted properties are left unchanged. To clear an
// existing description send `description: ""`.
func (r *BGPFilterProfileService) Update(ctx context.Context, profileID string, params BGPFilterProfileUpdateParams, opts ...option.RequestOption) (res *BGPFilterProfileUpdateResponse, err error) {
	var env BGPFilterProfileUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/bgp/filter_profiles/%s", params.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all BGP filter profiles for an account.
func (r *BGPFilterProfileService) List(ctx context.Context, query BGPFilterProfileListParams, opts ...option.RequestOption) (res *pagination.SinglePage[BGPFilterProfileListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/bgp/filter_profiles", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all BGP filter profiles for an account.
func (r *BGPFilterProfileService) ListAutoPaging(ctx context.Context, query BGPFilterProfileListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[BGPFilterProfileListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a BGP filter profile.
func (r *BGPFilterProfileService) Delete(ctx context.Context, profileID string, body BGPFilterProfileDeleteParams, opts ...option.RequestOption) (res *BGPFilterProfileDeleteResponse, err error) {
	var env BGPFilterProfileDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/bgp/filter_profiles/%s", body.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Gets a specific BGP filter profile for an account.
func (r *BGPFilterProfileService) Get(ctx context.Context, profileID string, query BGPFilterProfileGetParams, opts ...option.RequestOption) (res *BGPFilterProfileGetResponse, err error) {
	var env BGPFilterProfileGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/bgp/filter_profiles/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type BGPFilterProfileNewResponse struct {
	// Identifier
	ID string `json:"id" api:"required"`
	// Description of the filter profile
	Description string `json:"description" api:"required"`
	// Action to take when a route matches one of the targets in this profile
	MatchAction BGPFilterProfileNewResponseMatchAction `json:"match_action" api:"required"`
	// Friendly name for the filter profile
	Name string `json:"name" api:"required"`
	// List of CIDR prefixes. Each entry may carry an optional suffix that specifies
	// which prefix lengths to match relative to the prefix length N: '{X,Y}' matches
	// prefix lengths in the inclusive range [X, Y] where N <= X <= Y <= max (max is 32
	// for IPv4, 128 for IPv6), '{X}' matches exactly length X (equivalent to {X,X}),
	// '+' is shorthand for {N, max} (the prefix and all more-specific subnets,
	// including at length N itself; valid even when N is the maximum length). Omit the
	// suffix to match the prefix exactly at length N.
	Targets    []string                        `json:"targets" api:"required"`
	CreatedOn  time.Time                       `json:"created_on" format:"date-time"`
	ModifiedOn time.Time                       `json:"modified_on" format:"date-time"`
	JSON       bgpFilterProfileNewResponseJSON `json:"-"`
}

// bgpFilterProfileNewResponseJSON contains the JSON metadata for the struct
// [BGPFilterProfileNewResponse]
type bgpFilterProfileNewResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	MatchAction apijson.Field
	Name        apijson.Field
	Targets     apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPFilterProfileNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpFilterProfileNewResponseJSON) RawJSON() string {
	return r.raw
}

// Action to take when a route matches one of the targets in this profile
type BGPFilterProfileNewResponseMatchAction string

const (
	BGPFilterProfileNewResponseMatchActionAllow BGPFilterProfileNewResponseMatchAction = "allow"
	BGPFilterProfileNewResponseMatchActionDeny  BGPFilterProfileNewResponseMatchAction = "deny"
)

func (r BGPFilterProfileNewResponseMatchAction) IsKnown() bool {
	switch r {
	case BGPFilterProfileNewResponseMatchActionAllow, BGPFilterProfileNewResponseMatchActionDeny:
		return true
	}
	return false
}

type BGPFilterProfileUpdateResponse struct {
	// Identifier
	ID string `json:"id" api:"required"`
	// Description of the filter profile
	Description string `json:"description" api:"required"`
	// Action to take when a route matches one of the targets in this profile
	MatchAction BGPFilterProfileUpdateResponseMatchAction `json:"match_action" api:"required"`
	// Friendly name for the filter profile
	Name string `json:"name" api:"required"`
	// List of CIDR prefixes. Each entry may carry an optional suffix that specifies
	// which prefix lengths to match relative to the prefix length N: '{X,Y}' matches
	// prefix lengths in the inclusive range [X, Y] where N <= X <= Y <= max (max is 32
	// for IPv4, 128 for IPv6), '{X}' matches exactly length X (equivalent to {X,X}),
	// '+' is shorthand for {N, max} (the prefix and all more-specific subnets,
	// including at length N itself; valid even when N is the maximum length). Omit the
	// suffix to match the prefix exactly at length N.
	Targets    []string                           `json:"targets" api:"required"`
	CreatedOn  time.Time                          `json:"created_on" format:"date-time"`
	ModifiedOn time.Time                          `json:"modified_on" format:"date-time"`
	JSON       bgpFilterProfileUpdateResponseJSON `json:"-"`
}

// bgpFilterProfileUpdateResponseJSON contains the JSON metadata for the struct
// [BGPFilterProfileUpdateResponse]
type bgpFilterProfileUpdateResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	MatchAction apijson.Field
	Name        apijson.Field
	Targets     apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPFilterProfileUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpFilterProfileUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Action to take when a route matches one of the targets in this profile
type BGPFilterProfileUpdateResponseMatchAction string

const (
	BGPFilterProfileUpdateResponseMatchActionAllow BGPFilterProfileUpdateResponseMatchAction = "allow"
	BGPFilterProfileUpdateResponseMatchActionDeny  BGPFilterProfileUpdateResponseMatchAction = "deny"
)

func (r BGPFilterProfileUpdateResponseMatchAction) IsKnown() bool {
	switch r {
	case BGPFilterProfileUpdateResponseMatchActionAllow, BGPFilterProfileUpdateResponseMatchActionDeny:
		return true
	}
	return false
}

type BGPFilterProfileListResponse struct {
	// Identifier
	ID string `json:"id" api:"required"`
	// Description of the filter profile
	Description string `json:"description" api:"required"`
	// Action to take when a route matches one of the targets in this profile
	MatchAction BGPFilterProfileListResponseMatchAction `json:"match_action" api:"required"`
	// Friendly name for the filter profile
	Name string `json:"name" api:"required"`
	// List of CIDR prefixes. Each entry may carry an optional suffix that specifies
	// which prefix lengths to match relative to the prefix length N: '{X,Y}' matches
	// prefix lengths in the inclusive range [X, Y] where N <= X <= Y <= max (max is 32
	// for IPv4, 128 for IPv6), '{X}' matches exactly length X (equivalent to {X,X}),
	// '+' is shorthand for {N, max} (the prefix and all more-specific subnets,
	// including at length N itself; valid even when N is the maximum length). Omit the
	// suffix to match the prefix exactly at length N.
	Targets    []string                         `json:"targets" api:"required"`
	CreatedOn  time.Time                        `json:"created_on" format:"date-time"`
	ModifiedOn time.Time                        `json:"modified_on" format:"date-time"`
	JSON       bgpFilterProfileListResponseJSON `json:"-"`
}

// bgpFilterProfileListResponseJSON contains the JSON metadata for the struct
// [BGPFilterProfileListResponse]
type bgpFilterProfileListResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	MatchAction apijson.Field
	Name        apijson.Field
	Targets     apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPFilterProfileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpFilterProfileListResponseJSON) RawJSON() string {
	return r.raw
}

// Action to take when a route matches one of the targets in this profile
type BGPFilterProfileListResponseMatchAction string

const (
	BGPFilterProfileListResponseMatchActionAllow BGPFilterProfileListResponseMatchAction = "allow"
	BGPFilterProfileListResponseMatchActionDeny  BGPFilterProfileListResponseMatchAction = "deny"
)

func (r BGPFilterProfileListResponseMatchAction) IsKnown() bool {
	switch r {
	case BGPFilterProfileListResponseMatchActionAllow, BGPFilterProfileListResponseMatchActionDeny:
		return true
	}
	return false
}

type BGPFilterProfileDeleteResponse struct {
	// Identifier
	ID string `json:"id" api:"required"`
	// Description of the filter profile
	Description string `json:"description" api:"required"`
	// Action to take when a route matches one of the targets in this profile
	MatchAction BGPFilterProfileDeleteResponseMatchAction `json:"match_action" api:"required"`
	// Friendly name for the filter profile
	Name string `json:"name" api:"required"`
	// List of CIDR prefixes. Each entry may carry an optional suffix that specifies
	// which prefix lengths to match relative to the prefix length N: '{X,Y}' matches
	// prefix lengths in the inclusive range [X, Y] where N <= X <= Y <= max (max is 32
	// for IPv4, 128 for IPv6), '{X}' matches exactly length X (equivalent to {X,X}),
	// '+' is shorthand for {N, max} (the prefix and all more-specific subnets,
	// including at length N itself; valid even when N is the maximum length). Omit the
	// suffix to match the prefix exactly at length N.
	Targets    []string                           `json:"targets" api:"required"`
	CreatedOn  time.Time                          `json:"created_on" format:"date-time"`
	ModifiedOn time.Time                          `json:"modified_on" format:"date-time"`
	JSON       bgpFilterProfileDeleteResponseJSON `json:"-"`
}

// bgpFilterProfileDeleteResponseJSON contains the JSON metadata for the struct
// [BGPFilterProfileDeleteResponse]
type bgpFilterProfileDeleteResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	MatchAction apijson.Field
	Name        apijson.Field
	Targets     apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPFilterProfileDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpFilterProfileDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Action to take when a route matches one of the targets in this profile
type BGPFilterProfileDeleteResponseMatchAction string

const (
	BGPFilterProfileDeleteResponseMatchActionAllow BGPFilterProfileDeleteResponseMatchAction = "allow"
	BGPFilterProfileDeleteResponseMatchActionDeny  BGPFilterProfileDeleteResponseMatchAction = "deny"
)

func (r BGPFilterProfileDeleteResponseMatchAction) IsKnown() bool {
	switch r {
	case BGPFilterProfileDeleteResponseMatchActionAllow, BGPFilterProfileDeleteResponseMatchActionDeny:
		return true
	}
	return false
}

type BGPFilterProfileGetResponse struct {
	// Identifier
	ID string `json:"id" api:"required"`
	// Description of the filter profile
	Description string `json:"description" api:"required"`
	// Action to take when a route matches one of the targets in this profile
	MatchAction BGPFilterProfileGetResponseMatchAction `json:"match_action" api:"required"`
	// Friendly name for the filter profile
	Name string `json:"name" api:"required"`
	// List of CIDR prefixes. Each entry may carry an optional suffix that specifies
	// which prefix lengths to match relative to the prefix length N: '{X,Y}' matches
	// prefix lengths in the inclusive range [X, Y] where N <= X <= Y <= max (max is 32
	// for IPv4, 128 for IPv6), '{X}' matches exactly length X (equivalent to {X,X}),
	// '+' is shorthand for {N, max} (the prefix and all more-specific subnets,
	// including at length N itself; valid even when N is the maximum length). Omit the
	// suffix to match the prefix exactly at length N.
	Targets    []string                        `json:"targets" api:"required"`
	CreatedOn  time.Time                       `json:"created_on" format:"date-time"`
	ModifiedOn time.Time                       `json:"modified_on" format:"date-time"`
	JSON       bgpFilterProfileGetResponseJSON `json:"-"`
}

// bgpFilterProfileGetResponseJSON contains the JSON metadata for the struct
// [BGPFilterProfileGetResponse]
type bgpFilterProfileGetResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	MatchAction apijson.Field
	Name        apijson.Field
	Targets     apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPFilterProfileGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpFilterProfileGetResponseJSON) RawJSON() string {
	return r.raw
}

// Action to take when a route matches one of the targets in this profile
type BGPFilterProfileGetResponseMatchAction string

const (
	BGPFilterProfileGetResponseMatchActionAllow BGPFilterProfileGetResponseMatchAction = "allow"
	BGPFilterProfileGetResponseMatchActionDeny  BGPFilterProfileGetResponseMatchAction = "deny"
)

func (r BGPFilterProfileGetResponseMatchAction) IsKnown() bool {
	switch r {
	case BGPFilterProfileGetResponseMatchActionAllow, BGPFilterProfileGetResponseMatchActionDeny:
		return true
	}
	return false
}

type BGPFilterProfileNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Action to take when a route matches one of the targets in this profile
	MatchAction param.Field[BGPFilterProfileNewParamsMatchAction] `json:"match_action" api:"required"`
	// Friendly name for the filter profile
	Name param.Field[string] `json:"name" api:"required"`
	// List of CIDR prefixes. Each entry may carry an optional suffix that specifies
	// which prefix lengths to match relative to the prefix length N: '{X,Y}' matches
	// prefix lengths in the inclusive range [X, Y] where N <= X <= Y <= max (max is 32
	// for IPv4, 128 for IPv6), '{X}' matches exactly length X (equivalent to {X,X}),
	// '+' is shorthand for {N, max} (the prefix and all more-specific subnets,
	// including at length N itself; valid even when N is the maximum length). Omit the
	// suffix to match the prefix exactly at length N.
	Targets param.Field[[]string] `json:"targets" api:"required"`
	// Description of the filter profile
	Description param.Field[string] `json:"description"`
}

func (r BGPFilterProfileNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action to take when a route matches one of the targets in this profile
type BGPFilterProfileNewParamsMatchAction string

const (
	BGPFilterProfileNewParamsMatchActionAllow BGPFilterProfileNewParamsMatchAction = "allow"
	BGPFilterProfileNewParamsMatchActionDeny  BGPFilterProfileNewParamsMatchAction = "deny"
)

func (r BGPFilterProfileNewParamsMatchAction) IsKnown() bool {
	switch r {
	case BGPFilterProfileNewParamsMatchActionAllow, BGPFilterProfileNewParamsMatchActionDeny:
		return true
	}
	return false
}

type BGPFilterProfileNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors" api:"required"`
	Messages []shared.ResponseInfo       `json:"messages" api:"required"`
	Result   BGPFilterProfileNewResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success BGPFilterProfileNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    bgpFilterProfileNewResponseEnvelopeJSON    `json:"-"`
}

// bgpFilterProfileNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [BGPFilterProfileNewResponseEnvelope]
type bgpFilterProfileNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPFilterProfileNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpFilterProfileNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BGPFilterProfileNewResponseEnvelopeSuccess bool

const (
	BGPFilterProfileNewResponseEnvelopeSuccessTrue BGPFilterProfileNewResponseEnvelopeSuccess = true
)

func (r BGPFilterProfileNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BGPFilterProfileNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BGPFilterProfileUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Description of the filter profile
	Description param.Field[string] `json:"description"`
	// Action to take when a route matches one of the targets in this profile
	MatchAction param.Field[BGPFilterProfileUpdateParamsMatchAction] `json:"match_action"`
	// Friendly name for the filter profile
	Name param.Field[string] `json:"name"`
	// List of CIDR prefixes. Each entry may carry an optional suffix that specifies
	// which prefix lengths to match relative to the prefix length N: '{X,Y}' matches
	// prefix lengths in the inclusive range [X, Y] where N <= X <= Y <= max (max is 32
	// for IPv4, 128 for IPv6), '{X}' matches exactly length X (equivalent to {X,X}),
	// '+' is shorthand for {N, max} (the prefix and all more-specific subnets,
	// including at length N itself; valid even when N is the maximum length). Omit the
	// suffix to match the prefix exactly at length N.
	Targets param.Field[[]string] `json:"targets"`
}

func (r BGPFilterProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action to take when a route matches one of the targets in this profile
type BGPFilterProfileUpdateParamsMatchAction string

const (
	BGPFilterProfileUpdateParamsMatchActionAllow BGPFilterProfileUpdateParamsMatchAction = "allow"
	BGPFilterProfileUpdateParamsMatchActionDeny  BGPFilterProfileUpdateParamsMatchAction = "deny"
)

func (r BGPFilterProfileUpdateParamsMatchAction) IsKnown() bool {
	switch r {
	case BGPFilterProfileUpdateParamsMatchActionAllow, BGPFilterProfileUpdateParamsMatchActionDeny:
		return true
	}
	return false
}

type BGPFilterProfileUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors" api:"required"`
	Messages []shared.ResponseInfo          `json:"messages" api:"required"`
	Result   BGPFilterProfileUpdateResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success BGPFilterProfileUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    bgpFilterProfileUpdateResponseEnvelopeJSON    `json:"-"`
}

// bgpFilterProfileUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [BGPFilterProfileUpdateResponseEnvelope]
type bgpFilterProfileUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPFilterProfileUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpFilterProfileUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BGPFilterProfileUpdateResponseEnvelopeSuccess bool

const (
	BGPFilterProfileUpdateResponseEnvelopeSuccessTrue BGPFilterProfileUpdateResponseEnvelopeSuccess = true
)

func (r BGPFilterProfileUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BGPFilterProfileUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BGPFilterProfileListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BGPFilterProfileDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BGPFilterProfileDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors" api:"required"`
	Messages []shared.ResponseInfo          `json:"messages" api:"required"`
	Result   BGPFilterProfileDeleteResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success BGPFilterProfileDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    bgpFilterProfileDeleteResponseEnvelopeJSON    `json:"-"`
}

// bgpFilterProfileDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [BGPFilterProfileDeleteResponseEnvelope]
type bgpFilterProfileDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPFilterProfileDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpFilterProfileDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BGPFilterProfileDeleteResponseEnvelopeSuccess bool

const (
	BGPFilterProfileDeleteResponseEnvelopeSuccessTrue BGPFilterProfileDeleteResponseEnvelopeSuccess = true
)

func (r BGPFilterProfileDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BGPFilterProfileDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BGPFilterProfileGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BGPFilterProfileGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors" api:"required"`
	Messages []shared.ResponseInfo       `json:"messages" api:"required"`
	Result   BGPFilterProfileGetResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success BGPFilterProfileGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    bgpFilterProfileGetResponseEnvelopeJSON    `json:"-"`
}

// bgpFilterProfileGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [BGPFilterProfileGetResponseEnvelope]
type bgpFilterProfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPFilterProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpFilterProfileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BGPFilterProfileGetResponseEnvelopeSuccess bool

const (
	BGPFilterProfileGetResponseEnvelopeSuccessTrue BGPFilterProfileGetResponseEnvelopeSuccess = true
)

func (r BGPFilterProfileGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BGPFilterProfileGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dls

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// RegionalServicePrefixBindingService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegionalServicePrefixBindingService] method instead.
type RegionalServicePrefixBindingService struct {
	Options []option.RequestOption
}

// NewRegionalServicePrefixBindingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRegionalServicePrefixBindingService(opts ...option.RequestOption) (r *RegionalServicePrefixBindingService) {
	r = &RegionalServicePrefixBindingService{}
	r.Options = opts
	return
}

// Create a DLS prefix binding
func (r *RegionalServicePrefixBindingService) New(ctx context.Context, params RegionalServicePrefixBindingNewParams, opts ...option.RequestOption) (res *RegionalServicePrefixBindingNewResponse, err error) {
	var env RegionalServicePrefixBindingNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dls/regional_services/prefix_bindings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List DLS prefix bindings for an account
func (r *RegionalServicePrefixBindingService) List(ctx context.Context, params RegionalServicePrefixBindingListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[RegionalServicePrefixBindingListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dls/regional_services/prefix_bindings", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List DLS prefix bindings for an account
func (r *RegionalServicePrefixBindingService) ListAutoPaging(ctx context.Context, params RegionalServicePrefixBindingListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[RegionalServicePrefixBindingListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Delete a DLS prefix binding
func (r *RegionalServicePrefixBindingService) Delete(ctx context.Context, bindingID string, body RegionalServicePrefixBindingDeleteParams, opts ...option.RequestOption) (res *RegionalServicePrefixBindingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if bindingID == "" {
		err = errors.New("missing required binding_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dls/regional_services/prefix_bindings/%s", body.AccountID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Update a DLS prefix binding
func (r *RegionalServicePrefixBindingService) Edit(ctx context.Context, bindingID string, params RegionalServicePrefixBindingEditParams, opts ...option.RequestOption) (res *RegionalServicePrefixBindingEditResponse, err error) {
	var env RegionalServicePrefixBindingEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if bindingID == "" {
		err = errors.New("missing required binding_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dls/regional_services/prefix_bindings/%s", params.AccountID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get a DLS prefix binding
func (r *RegionalServicePrefixBindingService) Get(ctx context.Context, bindingID string, query RegionalServicePrefixBindingGetParams, opts ...option.RequestOption) (res *RegionalServicePrefixBindingGetResponse, err error) {
	var env RegionalServicePrefixBindingGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if bindingID == "" {
		err = errors.New("missing required binding_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dls/regional_services/prefix_bindings/%s", query.AccountID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type RegionalServicePrefixBindingNewResponse struct {
	// The ID of the binding.
	ID string `json:"id" api:"required"`
	// The CIDR that is bound.
	CIDR string `json:"cidr" api:"required"`
	// The ID of the parent prefix.
	PrefixID string `json:"prefix_id" api:"required"`
	// The region key used for the binding.
	RegionKey string                                      `json:"region_key" api:"required"`
	JSON      regionalServicePrefixBindingNewResponseJSON `json:"-"`
}

// regionalServicePrefixBindingNewResponseJSON contains the JSON metadata for the
// struct [RegionalServicePrefixBindingNewResponse]
type regionalServicePrefixBindingNewResponseJSON struct {
	ID          apijson.Field
	CIDR        apijson.Field
	PrefixID    apijson.Field
	RegionKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingNewResponseJSON) RawJSON() string {
	return r.raw
}

type RegionalServicePrefixBindingListResponse struct {
	// The ID of the binding.
	ID string `json:"id" api:"required"`
	// The CIDR that is bound.
	CIDR string `json:"cidr" api:"required"`
	// The ID of the parent prefix.
	PrefixID string `json:"prefix_id" api:"required"`
	// The region key used for the binding.
	RegionKey string                                       `json:"region_key" api:"required"`
	JSON      regionalServicePrefixBindingListResponseJSON `json:"-"`
}

// regionalServicePrefixBindingListResponseJSON contains the JSON metadata for the
// struct [RegionalServicePrefixBindingListResponse]
type regionalServicePrefixBindingListResponseJSON struct {
	ID          apijson.Field
	CIDR        apijson.Field
	PrefixID    apijson.Field
	RegionKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingListResponseJSON) RawJSON() string {
	return r.raw
}

type RegionalServicePrefixBindingDeleteResponse struct {
	Messages []shared.ResponseInfo                          `json:"messages" api:"required"`
	Success  bool                                           `json:"success" api:"required"`
	Errors   []shared.ResponseInfo                          `json:"errors"`
	JSON     regionalServicePrefixBindingDeleteResponseJSON `json:"-"`
}

// regionalServicePrefixBindingDeleteResponseJSON contains the JSON metadata for
// the struct [RegionalServicePrefixBindingDeleteResponse]
type regionalServicePrefixBindingDeleteResponseJSON struct {
	Messages    apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type RegionalServicePrefixBindingEditResponse struct {
	// The ID of the binding.
	ID string `json:"id" api:"required"`
	// The CIDR that is bound.
	CIDR string `json:"cidr" api:"required"`
	// The ID of the parent prefix.
	PrefixID string `json:"prefix_id" api:"required"`
	// The region key used for the binding.
	RegionKey string                                       `json:"region_key" api:"required"`
	JSON      regionalServicePrefixBindingEditResponseJSON `json:"-"`
}

// regionalServicePrefixBindingEditResponseJSON contains the JSON metadata for the
// struct [RegionalServicePrefixBindingEditResponse]
type regionalServicePrefixBindingEditResponseJSON struct {
	ID          apijson.Field
	CIDR        apijson.Field
	PrefixID    apijson.Field
	RegionKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingEditResponseJSON) RawJSON() string {
	return r.raw
}

type RegionalServicePrefixBindingGetResponse struct {
	// The ID of the binding.
	ID string `json:"id" api:"required"`
	// The CIDR that is bound.
	CIDR string `json:"cidr" api:"required"`
	// The ID of the parent prefix.
	PrefixID string `json:"prefix_id" api:"required"`
	// The region key used for the binding.
	RegionKey string                                      `json:"region_key" api:"required"`
	JSON      regionalServicePrefixBindingGetResponseJSON `json:"-"`
}

// regionalServicePrefixBindingGetResponseJSON contains the JSON metadata for the
// struct [RegionalServicePrefixBindingGetResponse]
type regionalServicePrefixBindingGetResponseJSON struct {
	ID          apijson.Field
	CIDR        apijson.Field
	PrefixID    apijson.Field
	RegionKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingGetResponseJSON) RawJSON() string {
	return r.raw
}

type RegionalServicePrefixBindingNewParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// IP prefix in CIDR notation to bind.
	CIDR param.Field[string] `json:"cidr" api:"required"`
	// The ID of the parent IP prefix that contains the CIDR.
	PrefixID param.Field[string] `json:"prefix_id" api:"required"`
	// Region key from managed regions (e.g., "us", "eu").
	RegionKey param.Field[string] `json:"region_key" api:"required"`
}

func (r RegionalServicePrefixBindingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RegionalServicePrefixBindingNewResponseEnvelope struct {
	Messages []shared.ResponseInfo                               `json:"messages" api:"required"`
	Result   RegionalServicePrefixBindingNewResponse             `json:"result" api:"required"`
	Success  bool                                                `json:"success" api:"required"`
	Errors   []shared.ResponseInfo                               `json:"errors"`
	JSON     regionalServicePrefixBindingNewResponseEnvelopeJSON `json:"-"`
}

// regionalServicePrefixBindingNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [RegionalServicePrefixBindingNewResponseEnvelope]
type regionalServicePrefixBindingNewResponseEnvelopeJSON struct {
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RegionalServicePrefixBindingListParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Opaque token for cursor-based pagination. Omit for the first page. Pass the
	// value from a previous response to fetch the next page.
	Cursor  param.Field[string] `query:"cursor"`
	PerPage param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [RegionalServicePrefixBindingListParams]'s query parameters
// as `url.Values`.
func (r RegionalServicePrefixBindingListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RegionalServicePrefixBindingDeleteParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RegionalServicePrefixBindingEditParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// New region key to assign (e.g., "us", "eu", "cfcanary").
	RegionKey param.Field[string] `json:"region_key" api:"required"`
}

func (r RegionalServicePrefixBindingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RegionalServicePrefixBindingEditResponseEnvelope struct {
	Messages []shared.ResponseInfo                                `json:"messages" api:"required"`
	Result   RegionalServicePrefixBindingEditResponse             `json:"result" api:"required"`
	Success  bool                                                 `json:"success" api:"required"`
	Errors   []shared.ResponseInfo                                `json:"errors"`
	JSON     regionalServicePrefixBindingEditResponseEnvelopeJSON `json:"-"`
}

// regionalServicePrefixBindingEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [RegionalServicePrefixBindingEditResponseEnvelope]
type regionalServicePrefixBindingEditResponseEnvelopeJSON struct {
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RegionalServicePrefixBindingGetParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RegionalServicePrefixBindingGetResponseEnvelope struct {
	Messages []shared.ResponseInfo                               `json:"messages" api:"required"`
	Result   RegionalServicePrefixBindingGetResponse             `json:"result" api:"required"`
	Success  bool                                                `json:"success" api:"required"`
	Errors   []shared.ResponseInfo                               `json:"errors"`
	JSON     regionalServicePrefixBindingGetResponseEnvelopeJSON `json:"-"`
}

// regionalServicePrefixBindingGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RegionalServicePrefixBindingGetResponseEnvelope]
type regionalServicePrefixBindingGetResponseEnvelopeJSON struct {
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

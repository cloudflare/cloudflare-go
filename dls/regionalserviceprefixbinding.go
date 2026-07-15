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

// Bind a CIDR from a BYOIP prefix to a region.
//
// This requires the **IP Prefixes Write** permission in addition to **DLS Write**,
// because the binding is created against a BYOIP prefix in Addressing.
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

// List the BYOIP prefix bindings configured for an account.
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

// List the BYOIP prefix bindings configured for an account.
func (r *RegionalServicePrefixBindingService) ListAutoPaging(ctx context.Context, params RegionalServicePrefixBindingListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[RegionalServicePrefixBindingListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Delete a BYOIP prefix binding.
//
// Like creating a binding, this requires **IP Prefixes Write** in addition to
// **DLS Write**.
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

// Update the region of an existing BYOIP prefix binding.
//
// Like creating a binding, this requires **IP Prefixes Write** in addition to
// **DLS Write**.
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

// Retrieve a single BYOIP prefix binding by ID.
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
	Messages []RegionalServicePrefixBindingDeleteResponseMessage `json:"messages" api:"required"`
	Success  bool                                                `json:"success" api:"required"`
	Errors   []RegionalServicePrefixBindingDeleteResponseError   `json:"errors"`
	JSON     regionalServicePrefixBindingDeleteResponseJSON      `json:"-"`
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

type RegionalServicePrefixBindingDeleteResponseMessage struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Optional upstream error context for APIv4 errors that wrap downstream service
	// failures.
	ErrorChain []interface{}                                         `json:"error_chain"`
	JSON       regionalServicePrefixBindingDeleteResponseMessageJSON `json:"-"`
}

// regionalServicePrefixBindingDeleteResponseMessageJSON contains the JSON metadata
// for the struct [RegionalServicePrefixBindingDeleteResponseMessage]
type regionalServicePrefixBindingDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	ErrorChain  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type RegionalServicePrefixBindingDeleteResponseError struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Optional upstream error context for APIv4 errors that wrap downstream service
	// failures.
	ErrorChain []interface{}                                       `json:"error_chain"`
	JSON       regionalServicePrefixBindingDeleteResponseErrorJSON `json:"-"`
}

// regionalServicePrefixBindingDeleteResponseErrorJSON contains the JSON metadata
// for the struct [RegionalServicePrefixBindingDeleteResponseError]
type regionalServicePrefixBindingDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	ErrorChain  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingDeleteResponseErrorJSON) RawJSON() string {
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
	Messages []RegionalServicePrefixBindingNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   RegionalServicePrefixBindingNewResponse                   `json:"result" api:"required"`
	Success  bool                                                      `json:"success" api:"required"`
	Errors   []RegionalServicePrefixBindingNewResponseEnvelopeErrors   `json:"errors"`
	JSON     regionalServicePrefixBindingNewResponseEnvelopeJSON       `json:"-"`
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

type RegionalServicePrefixBindingNewResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Optional upstream error context for APIv4 errors that wrap downstream service
	// failures.
	ErrorChain []interface{}                                               `json:"error_chain"`
	JSON       regionalServicePrefixBindingNewResponseEnvelopeMessagesJSON `json:"-"`
}

// regionalServicePrefixBindingNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [RegionalServicePrefixBindingNewResponseEnvelopeMessages]
type regionalServicePrefixBindingNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	ErrorChain  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RegionalServicePrefixBindingNewResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Optional upstream error context for APIv4 errors that wrap downstream service
	// failures.
	ErrorChain []interface{}                                             `json:"error_chain"`
	JSON       regionalServicePrefixBindingNewResponseEnvelopeErrorsJSON `json:"-"`
}

// regionalServicePrefixBindingNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [RegionalServicePrefixBindingNewResponseEnvelopeErrors]
type regionalServicePrefixBindingNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	ErrorChain  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingNewResponseEnvelopeErrorsJSON) RawJSON() string {
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
	Messages []RegionalServicePrefixBindingEditResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   RegionalServicePrefixBindingEditResponse                   `json:"result" api:"required"`
	Success  bool                                                       `json:"success" api:"required"`
	Errors   []RegionalServicePrefixBindingEditResponseEnvelopeErrors   `json:"errors"`
	JSON     regionalServicePrefixBindingEditResponseEnvelopeJSON       `json:"-"`
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

type RegionalServicePrefixBindingEditResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Optional upstream error context for APIv4 errors that wrap downstream service
	// failures.
	ErrorChain []interface{}                                                `json:"error_chain"`
	JSON       regionalServicePrefixBindingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// regionalServicePrefixBindingEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [RegionalServicePrefixBindingEditResponseEnvelopeMessages]
type regionalServicePrefixBindingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	ErrorChain  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RegionalServicePrefixBindingEditResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Optional upstream error context for APIv4 errors that wrap downstream service
	// failures.
	ErrorChain []interface{}                                              `json:"error_chain"`
	JSON       regionalServicePrefixBindingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// regionalServicePrefixBindingEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [RegionalServicePrefixBindingEditResponseEnvelopeErrors]
type regionalServicePrefixBindingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	ErrorChain  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RegionalServicePrefixBindingGetParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RegionalServicePrefixBindingGetResponseEnvelope struct {
	Messages []RegionalServicePrefixBindingGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   RegionalServicePrefixBindingGetResponse                   `json:"result" api:"required"`
	Success  bool                                                      `json:"success" api:"required"`
	Errors   []RegionalServicePrefixBindingGetResponseEnvelopeErrors   `json:"errors"`
	JSON     regionalServicePrefixBindingGetResponseEnvelopeJSON       `json:"-"`
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

type RegionalServicePrefixBindingGetResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Optional upstream error context for APIv4 errors that wrap downstream service
	// failures.
	ErrorChain []interface{}                                               `json:"error_chain"`
	JSON       regionalServicePrefixBindingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// regionalServicePrefixBindingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [RegionalServicePrefixBindingGetResponseEnvelopeMessages]
type regionalServicePrefixBindingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	ErrorChain  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RegionalServicePrefixBindingGetResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Optional upstream error context for APIv4 errors that wrap downstream service
	// failures.
	ErrorChain []interface{}                                             `json:"error_chain"`
	JSON       regionalServicePrefixBindingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// regionalServicePrefixBindingGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [RegionalServicePrefixBindingGetResponseEnvelopeErrors]
type regionalServicePrefixBindingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	ErrorChain  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalServicePrefixBindingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalServicePrefixBindingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

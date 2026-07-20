// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

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
)

// ThreatEventTagCategoryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventTagCategoryService] method instead.
type ThreatEventTagCategoryService struct {
	Options []option.RequestOption
}

// NewThreatEventTagCategoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventTagCategoryService(opts ...option.RequestOption) (r *ThreatEventTagCategoryService) {
	r = &ThreatEventTagCategoryService{}
	r.Options = opts
	return
}

// Creates a new Source-of-Truth tag category for an account.
func (r *ThreatEventTagCategoryService) New(ctx context.Context, params ThreatEventTagCategoryNewParams, opts ...option.RequestOption) (res *ThreatEventTagCategoryNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/tags/categories/create", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns all Source-of-Truth tag categories for an account.
func (r *ThreatEventTagCategoryService) List(ctx context.Context, params ThreatEventTagCategoryListParams, opts ...option.RequestOption) (res *ThreatEventTagCategoryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/tags/categories", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Deletes a Source-of-Truth tag category by UUID.
func (r *ThreatEventTagCategoryService) Delete(ctx context.Context, categoryUUID string, body ThreatEventTagCategoryDeleteParams, opts ...option.RequestOption) (res *ThreatEventTagCategoryDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if categoryUUID == "" {
		err = errors.New("missing required category_uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/tags/categories/%s", body.AccountID, categoryUUID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Updates a Source-of-Truth tag category by UUID.
func (r *ThreatEventTagCategoryService) Edit(ctx context.Context, categoryUUID string, params ThreatEventTagCategoryEditParams, opts ...option.RequestOption) (res *ThreatEventTagCategoryEditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if categoryUUID == "" {
		err = errors.New("missing required category_uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/tags/categories/%s", params.AccountID, categoryUUID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

type ThreatEventTagCategoryNewResponse struct {
	Name        string                                `json:"name" api:"required"`
	UUID        string                                `json:"uuid" api:"required"`
	CreatedAt   string                                `json:"createdAt"`
	Description string                                `json:"description"`
	UpdatedAt   string                                `json:"updatedAt"`
	JSON        threatEventTagCategoryNewResponseJSON `json:"-"`
}

// threatEventTagCategoryNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagCategoryNewResponse]
type threatEventTagCategoryNewResponseJSON struct {
	Name        apijson.Field
	UUID        apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryListResponse struct {
	Categories []ThreatEventTagCategoryListResponseCategory `json:"categories" api:"required"`
	JSON       threatEventTagCategoryListResponseJSON       `json:"-"`
}

// threatEventTagCategoryListResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagCategoryListResponse]
type threatEventTagCategoryListResponseJSON struct {
	Categories  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryListResponseCategory struct {
	Name        string                                         `json:"name" api:"required"`
	UUID        string                                         `json:"uuid" api:"required"`
	CreatedAt   string                                         `json:"createdAt"`
	Description string                                         `json:"description"`
	UpdatedAt   string                                         `json:"updatedAt"`
	JSON        threatEventTagCategoryListResponseCategoryJSON `json:"-"`
}

// threatEventTagCategoryListResponseCategoryJSON contains the JSON metadata for
// the struct [ThreatEventTagCategoryListResponseCategory]
type threatEventTagCategoryListResponseCategoryJSON struct {
	Name        apijson.Field
	UUID        apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryListResponseCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryListResponseCategoryJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryDeleteResponse struct {
	UUID string                                   `json:"uuid" api:"required"`
	JSON threatEventTagCategoryDeleteResponseJSON `json:"-"`
}

// threatEventTagCategoryDeleteResponseJSON contains the JSON metadata for the
// struct [ThreatEventTagCategoryDeleteResponse]
type threatEventTagCategoryDeleteResponseJSON struct {
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryEditResponse struct {
	Name        string                                 `json:"name" api:"required"`
	UUID        string                                 `json:"uuid" api:"required"`
	CreatedAt   string                                 `json:"createdAt"`
	Description string                                 `json:"description"`
	UpdatedAt   string                                 `json:"updatedAt"`
	JSON        threatEventTagCategoryEditResponseJSON `json:"-"`
}

// threatEventTagCategoryEditResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagCategoryEditResponse]
type threatEventTagCategoryEditResponseJSON struct {
	Name        apijson.Field
	UUID        apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryEditResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryNewParams struct {
	// Account ID.
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Name        param.Field[string] `json:"name" api:"required"`
	Description param.Field[string] `json:"description"`
}

func (r ThreatEventTagCategoryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagCategoryListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Search    param.Field[string] `query:"search"`
}

// URLQuery serializes [ThreatEventTagCategoryListParams]'s query parameters as
// `url.Values`.
func (r ThreatEventTagCategoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatEventTagCategoryDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ThreatEventTagCategoryEditParams struct {
	// Account ID.
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
}

func (r ThreatEventTagCategoryEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

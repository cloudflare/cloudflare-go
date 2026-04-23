// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_pages

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// AssetService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetService] method instead.
type AssetService struct {
	Options []option.RequestOption
}

// NewAssetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAssetService(opts ...option.RequestOption) (r *AssetService) {
	r = &AssetService{}
	r.Options = opts
	return
}

// Creates a new custom asset.
func (r *AssetService) New(ctx context.Context, params AssetNewParams, opts ...option.RequestOption) (res *AssetNewResponse, err error) {
	var env AssetNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/custom_pages/assets", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates the configuration of an existing custom asset.
func (r *AssetService) Update(ctx context.Context, assetName string, params AssetUpdateParams, opts ...option.RequestOption) (res *AssetUpdateResponse, err error) {
	var env AssetUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if assetName == "" {
		err = errors.New("missing required asset_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/custom_pages/assets/%s", accountOrZone, accountOrZoneID, assetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetches all the custom assets.
func (r *AssetService) List(ctx context.Context, params AssetListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AssetListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/custom_pages/assets", accountOrZone, accountOrZoneID)
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

// Fetches all the custom assets.
func (r *AssetService) ListAutoPaging(ctx context.Context, params AssetListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AssetListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an existing custom asset.
func (r *AssetService) Delete(ctx context.Context, assetName string, body AssetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	if assetName == "" {
		err = errors.New("missing required asset_name parameter")
		return err
	}
	path := fmt.Sprintf("%s/%s/custom_pages/assets/%s", accountOrZone, accountOrZoneID, assetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Fetches the details of a custom asset.
func (r *AssetService) Get(ctx context.Context, assetName string, query AssetGetParams, opts ...option.RequestOption) (res *AssetGetResponse, err error) {
	var env AssetGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	if assetName == "" {
		err = errors.New("missing required asset_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/custom_pages/assets/%s", accountOrZone, accountOrZoneID, assetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AssetNewResponse struct {
	// A short description of the custom asset.
	Description string    `json:"description"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The unique name of the custom asset. Can only contain letters (A-Z, a-z),
	// numbers (0-9), and underscores (\_).
	Name string `json:"name"`
	// The size of the asset content in bytes.
	SizeBytes int64 `json:"size_bytes"`
	// The URL where the asset content is fetched from.
	URL  string               `json:"url" format:"uri"`
	JSON assetNewResponseJSON `json:"-"`
}

// assetNewResponseJSON contains the JSON metadata for the struct
// [AssetNewResponse]
type assetNewResponseJSON struct {
	Description apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	SizeBytes   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetNewResponseJSON) RawJSON() string {
	return r.raw
}

type AssetUpdateResponse struct {
	// A short description of the custom asset.
	Description string    `json:"description"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The unique name of the custom asset. Can only contain letters (A-Z, a-z),
	// numbers (0-9), and underscores (\_).
	Name string `json:"name"`
	// The size of the asset content in bytes.
	SizeBytes int64 `json:"size_bytes"`
	// The URL where the asset content is fetched from.
	URL  string                  `json:"url" format:"uri"`
	JSON assetUpdateResponseJSON `json:"-"`
}

// assetUpdateResponseJSON contains the JSON metadata for the struct
// [AssetUpdateResponse]
type assetUpdateResponseJSON struct {
	Description apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	SizeBytes   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AssetListResponse struct {
	// A short description of the custom asset.
	Description string    `json:"description"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The unique name of the custom asset. Can only contain letters (A-Z, a-z),
	// numbers (0-9), and underscores (\_).
	Name string `json:"name"`
	// The size of the asset content in bytes.
	SizeBytes int64 `json:"size_bytes"`
	// The URL where the asset content is fetched from.
	URL  string                `json:"url" format:"uri"`
	JSON assetListResponseJSON `json:"-"`
}

// assetListResponseJSON contains the JSON metadata for the struct
// [AssetListResponse]
type assetListResponseJSON struct {
	Description apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	SizeBytes   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetListResponseJSON) RawJSON() string {
	return r.raw
}

type AssetGetResponse struct {
	// A short description of the custom asset.
	Description string    `json:"description"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The unique name of the custom asset. Can only contain letters (A-Z, a-z),
	// numbers (0-9), and underscores (\_).
	Name string `json:"name"`
	// The size of the asset content in bytes.
	SizeBytes int64 `json:"size_bytes"`
	// The URL where the asset content is fetched from.
	URL  string               `json:"url" format:"uri"`
	JSON assetGetResponseJSON `json:"-"`
}

// assetGetResponseJSON contains the JSON metadata for the struct
// [AssetGetResponse]
type assetGetResponseJSON struct {
	Description apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	SizeBytes   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetGetResponseJSON) RawJSON() string {
	return r.raw
}

type AssetNewParams struct {
	// A short description of the custom asset.
	Description param.Field[string] `json:"description" api:"required"`
	// The unique name of the custom asset. Can only contain letters (A-Z, a-z),
	// numbers (0-9), and underscores (\_).
	Name param.Field[string] `json:"name" api:"required"`
	// The URL where the asset content is fetched from.
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r AssetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AssetNewResponseEnvelope struct {
	Errors   []AssetNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AssetNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AssetNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AssetNewResponse                `json:"result"`
	JSON    assetNewResponseEnvelopeJSON    `json:"-"`
}

// assetNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AssetNewResponseEnvelope]
type assetNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AssetNewResponseEnvelopeErrors struct {
	Code             int64                                `json:"code" api:"required"`
	Message          string                               `json:"message" api:"required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           AssetNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             assetNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// assetNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AssetNewResponseEnvelopeErrors]
type assetNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AssetNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AssetNewResponseEnvelopeErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    assetNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// assetNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [AssetNewResponseEnvelopeErrorsSource]
type assetNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AssetNewResponseEnvelopeMessages struct {
	Code             int64                                  `json:"code" api:"required"`
	Message          string                                 `json:"message" api:"required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           AssetNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             assetNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// assetNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AssetNewResponseEnvelopeMessages]
type assetNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AssetNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AssetNewResponseEnvelopeMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    assetNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// assetNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [AssetNewResponseEnvelopeMessagesSource]
type assetNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AssetNewResponseEnvelopeSuccess bool

const (
	AssetNewResponseEnvelopeSuccessTrue AssetNewResponseEnvelopeSuccess = true
)

func (r AssetNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AssetNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AssetUpdateParams struct {
	// A short description of the custom asset.
	Description param.Field[string] `json:"description" api:"required"`
	// The URL where the asset content is fetched from.
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r AssetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AssetUpdateResponseEnvelope struct {
	Errors   []AssetUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AssetUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AssetUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AssetUpdateResponse                `json:"result"`
	JSON    assetUpdateResponseEnvelopeJSON    `json:"-"`
}

// assetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AssetUpdateResponseEnvelope]
type assetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AssetUpdateResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code" api:"required"`
	Message          string                                  `json:"message" api:"required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           AssetUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             assetUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// assetUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AssetUpdateResponseEnvelopeErrors]
type assetUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AssetUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AssetUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    assetUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// assetUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [AssetUpdateResponseEnvelopeErrorsSource]
type assetUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AssetUpdateResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code" api:"required"`
	Message          string                                    `json:"message" api:"required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           AssetUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             assetUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// assetUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AssetUpdateResponseEnvelopeMessages]
type assetUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AssetUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AssetUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    assetUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// assetUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [AssetUpdateResponseEnvelopeMessagesSource]
type assetUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AssetUpdateResponseEnvelopeSuccess bool

const (
	AssetUpdateResponseEnvelopeSuccessTrue AssetUpdateResponseEnvelopeSuccess = true
)

func (r AssetUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AssetUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AssetListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID  param.Field[string] `path:"zone_id"`
	Page    param.Field[int64]  `query:"page"`
	PerPage param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [AssetListParams]'s query parameters as `url.Values`.
func (r AssetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AssetDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AssetGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AssetGetResponseEnvelope struct {
	Errors   []AssetGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AssetGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AssetGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AssetGetResponse                `json:"result"`
	JSON    assetGetResponseEnvelopeJSON    `json:"-"`
}

// assetGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AssetGetResponseEnvelope]
type assetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AssetGetResponseEnvelopeErrors struct {
	Code             int64                                `json:"code" api:"required"`
	Message          string                               `json:"message" api:"required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           AssetGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             assetGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// assetGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AssetGetResponseEnvelopeErrors]
type assetGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AssetGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AssetGetResponseEnvelopeErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    assetGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// assetGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [AssetGetResponseEnvelopeErrorsSource]
type assetGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AssetGetResponseEnvelopeMessages struct {
	Code             int64                                  `json:"code" api:"required"`
	Message          string                                 `json:"message" api:"required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           AssetGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             assetGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// assetGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AssetGetResponseEnvelopeMessages]
type assetGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AssetGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AssetGetResponseEnvelopeMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    assetGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// assetGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [AssetGetResponseEnvelopeMessagesSource]
type assetGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AssetGetResponseEnvelopeSuccess bool

const (
	AssetGetResponseEnvelopeSuccessTrue AssetGetResponseEnvelopeSuccess = true
)

func (r AssetGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AssetGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

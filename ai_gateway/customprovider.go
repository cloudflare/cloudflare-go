// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// CustomProviderService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomProviderService] method instead.
type CustomProviderService struct {
	Options []option.RequestOption
}

// NewCustomProviderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomProviderService(opts ...option.RequestOption) (r *CustomProviderService) {
	r = &CustomProviderService{}
	r.Options = opts
	return
}

// Creates a new AI Gateway.
func (r *CustomProviderService) New(ctx context.Context, params CustomProviderNewParams, opts ...option.RequestOption) (res *CustomProviderNewResponse, err error) {
	var env CustomProviderNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/custom-providers", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all AI Gateway evaluator types configured for the account.
func (r *CustomProviderService) List(ctx context.Context, params CustomProviderListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[CustomProviderListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/custom-providers", params.AccountID)
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

// Lists all AI Gateway evaluator types configured for the account.
func (r *CustomProviderService) ListAutoPaging(ctx context.Context, params CustomProviderListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[CustomProviderListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an AI Gateway dataset.
func (r *CustomProviderService) Delete(ctx context.Context, id string, body CustomProviderDeleteParams, opts ...option.RequestOption) (res *CustomProviderDeleteResponse, err error) {
	var env CustomProviderDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/custom-providers/%s", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves details for a specific AI Gateway dataset.
func (r *CustomProviderService) Get(ctx context.Context, id string, query CustomProviderGetParams, opts ...option.RequestOption) (res *CustomProviderGetResponse, err error) {
	var env CustomProviderGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/custom-providers/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type CustomProviderNewResponse struct {
	ID          string                        `json:"id" api:"required" format:"uuid"`
	BaseURL     string                        `json:"base_url" api:"required" format:"uri"`
	CreatedAt   time.Time                     `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt  time.Time                     `json:"modified_at" api:"required" format:"date-time"`
	Name        string                        `json:"name" api:"required"`
	Slug        string                        `json:"slug" api:"required"`
	Beta        bool                          `json:"beta"`
	CurlExample string                        `json:"curl_example"`
	Description string                        `json:"description"`
	Enable      bool                          `json:"enable"`
	Headers     string                        `json:"headers"`
	JSExample   string                        `json:"js_example"`
	Link        string                        `json:"link"`
	Logo        string                        `json:"logo"`
	Position    int64                         `json:"position"`
	JSON        customProviderNewResponseJSON `json:"-"`
}

// customProviderNewResponseJSON contains the JSON metadata for the struct
// [CustomProviderNewResponse]
type customProviderNewResponseJSON struct {
	ID          apijson.Field
	BaseURL     apijson.Field
	CreatedAt   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Beta        apijson.Field
	CurlExample apijson.Field
	Description apijson.Field
	Enable      apijson.Field
	Headers     apijson.Field
	JSExample   apijson.Field
	Link        apijson.Field
	Logo        apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomProviderNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customProviderNewResponseJSON) RawJSON() string {
	return r.raw
}

type CustomProviderListResponse struct {
	ID          string                         `json:"id" api:"required" format:"uuid"`
	BaseURL     string                         `json:"base_url" api:"required" format:"uri"`
	CreatedAt   time.Time                      `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt  time.Time                      `json:"modified_at" api:"required" format:"date-time"`
	Name        string                         `json:"name" api:"required"`
	Slug        string                         `json:"slug" api:"required"`
	Beta        bool                           `json:"beta"`
	CurlExample string                         `json:"curl_example"`
	Description string                         `json:"description"`
	Enable      bool                           `json:"enable"`
	Headers     string                         `json:"headers"`
	JSExample   string                         `json:"js_example"`
	Link        string                         `json:"link"`
	Logo        string                         `json:"logo"`
	Position    int64                          `json:"position"`
	JSON        customProviderListResponseJSON `json:"-"`
}

// customProviderListResponseJSON contains the JSON metadata for the struct
// [CustomProviderListResponse]
type customProviderListResponseJSON struct {
	ID          apijson.Field
	BaseURL     apijson.Field
	CreatedAt   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Beta        apijson.Field
	CurlExample apijson.Field
	Description apijson.Field
	Enable      apijson.Field
	Headers     apijson.Field
	JSExample   apijson.Field
	Link        apijson.Field
	Logo        apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomProviderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customProviderListResponseJSON) RawJSON() string {
	return r.raw
}

type CustomProviderDeleteResponse struct {
	ID          string                           `json:"id" api:"required" format:"uuid"`
	BaseURL     string                           `json:"base_url" api:"required" format:"uri"`
	CreatedAt   time.Time                        `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt  time.Time                        `json:"modified_at" api:"required" format:"date-time"`
	Name        string                           `json:"name" api:"required"`
	Slug        string                           `json:"slug" api:"required"`
	Beta        bool                             `json:"beta"`
	CurlExample string                           `json:"curl_example"`
	Description string                           `json:"description"`
	Enable      bool                             `json:"enable"`
	Headers     string                           `json:"headers"`
	JSExample   string                           `json:"js_example"`
	Link        string                           `json:"link"`
	Logo        string                           `json:"logo"`
	Position    int64                            `json:"position"`
	JSON        customProviderDeleteResponseJSON `json:"-"`
}

// customProviderDeleteResponseJSON contains the JSON metadata for the struct
// [CustomProviderDeleteResponse]
type customProviderDeleteResponseJSON struct {
	ID          apijson.Field
	BaseURL     apijson.Field
	CreatedAt   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Beta        apijson.Field
	CurlExample apijson.Field
	Description apijson.Field
	Enable      apijson.Field
	Headers     apijson.Field
	JSExample   apijson.Field
	Link        apijson.Field
	Logo        apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomProviderDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customProviderDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CustomProviderGetResponse struct {
	ID          string                        `json:"id" api:"required" format:"uuid"`
	BaseURL     string                        `json:"base_url" api:"required" format:"uri"`
	CreatedAt   time.Time                     `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt  time.Time                     `json:"modified_at" api:"required" format:"date-time"`
	Name        string                        `json:"name" api:"required"`
	Slug        string                        `json:"slug" api:"required"`
	Beta        bool                          `json:"beta"`
	CurlExample string                        `json:"curl_example"`
	Description string                        `json:"description"`
	Enable      bool                          `json:"enable"`
	Headers     string                        `json:"headers"`
	JSExample   string                        `json:"js_example"`
	Link        string                        `json:"link"`
	Logo        string                        `json:"logo"`
	Position    int64                         `json:"position"`
	JSON        customProviderGetResponseJSON `json:"-"`
}

// customProviderGetResponseJSON contains the JSON metadata for the struct
// [CustomProviderGetResponse]
type customProviderGetResponseJSON struct {
	ID          apijson.Field
	BaseURL     apijson.Field
	CreatedAt   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Beta        apijson.Field
	CurlExample apijson.Field
	Description apijson.Field
	Enable      apijson.Field
	Headers     apijson.Field
	JSExample   apijson.Field
	Link        apijson.Field
	Logo        apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomProviderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customProviderGetResponseJSON) RawJSON() string {
	return r.raw
}

type CustomProviderNewParams struct {
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	BaseURL     param.Field[string] `json:"base_url" api:"required" format:"uri"`
	Name        param.Field[string] `json:"name" api:"required"`
	Slug        param.Field[string] `json:"slug" api:"required"`
	Beta        param.Field[bool]   `json:"beta"`
	CurlExample param.Field[string] `json:"curl_example"`
	Description param.Field[string] `json:"description"`
	Enable      param.Field[bool]   `json:"enable"`
	Headers     param.Field[string] `json:"headers"`
	JSExample   param.Field[string] `json:"js_example"`
	Link        param.Field[string] `json:"link"`
	Position    param.Field[int64]  `json:"position"`
}

func (r CustomProviderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomProviderNewResponseEnvelope struct {
	Result  CustomProviderNewResponse             `json:"result" api:"required"`
	Success bool                                  `json:"success" api:"required"`
	JSON    customProviderNewResponseEnvelopeJSON `json:"-"`
}

// customProviderNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomProviderNewResponseEnvelope]
type customProviderNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomProviderNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customProviderNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomProviderListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Beta      param.Field[bool]   `query:"beta"`
	Enable    param.Field[bool]   `query:"enable"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
	// Search by id, name, slug
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [CustomProviderListParams]'s query parameters as
// `url.Values`.
func (r CustomProviderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CustomProviderDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CustomProviderDeleteResponseEnvelope struct {
	Result  CustomProviderDeleteResponse             `json:"result" api:"required"`
	Success bool                                     `json:"success" api:"required"`
	JSON    customProviderDeleteResponseEnvelopeJSON `json:"-"`
}

// customProviderDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomProviderDeleteResponseEnvelope]
type customProviderDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomProviderDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customProviderDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomProviderGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CustomProviderGetResponseEnvelope struct {
	Result  CustomProviderGetResponse             `json:"result" api:"required"`
	Success bool                                  `json:"success" api:"required"`
	JSON    customProviderGetResponseEnvelopeJSON `json:"-"`
}

// customProviderGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomProviderGetResponseEnvelope]
type customProviderGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomProviderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customProviderGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search

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

// TokenService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenService] method instead.
type TokenService struct {
	Options []option.RequestOption
}

// NewTokenService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTokenService(opts ...option.RequestOption) (r *TokenService) {
	r = &TokenService{}
	r.Options = opts
	return
}

// Create new tokens.
func (r *TokenService) New(ctx context.Context, params TokenNewParams, opts ...option.RequestOption) (res *TokenNewResponse, err error) {
	var env TokenNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/tokens", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update tokens.
func (r *TokenService) Update(ctx context.Context, id string, params TokenUpdateParams, opts ...option.RequestOption) (res *TokenUpdateResponse, err error) {
	var env TokenUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/tokens/%s", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List tokens.
func (r *TokenService) List(ctx context.Context, params TokenListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[TokenListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/tokens", params.AccountID)
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

// List tokens.
func (r *TokenService) ListAutoPaging(ctx context.Context, params TokenListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[TokenListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete tokens.
func (r *TokenService) Delete(ctx context.Context, id string, body TokenDeleteParams, opts ...option.RequestOption) (res *TokenDeleteResponse, err error) {
	var env TokenDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/tokens/%s", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Read tokens.
func (r *TokenService) Read(ctx context.Context, id string, query TokenReadParams, opts ...option.RequestOption) (res *TokenReadResponse, err error) {
	var env TokenReadResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/tokens/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TokenNewResponse struct {
	ID         string               `json:"id,required" format:"uuid"`
	CfAPIID    string               `json:"cf_api_id,required"`
	CreatedAt  time.Time            `json:"created_at,required" format:"date-time"`
	ModifiedAt time.Time            `json:"modified_at,required" format:"date-time"`
	Name       string               `json:"name,required"`
	CreatedBy  string               `json:"created_by,nullable"`
	Enabled    bool                 `json:"enabled"`
	Legacy     bool                 `json:"legacy"`
	ModifiedBy string               `json:"modified_by,nullable"`
	JSON       tokenNewResponseJSON `json:"-"`
}

// tokenNewResponseJSON contains the JSON metadata for the struct
// [TokenNewResponse]
type tokenNewResponseJSON struct {
	ID          apijson.Field
	CfAPIID     apijson.Field
	CreatedAt   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	CreatedBy   apijson.Field
	Enabled     apijson.Field
	Legacy      apijson.Field
	ModifiedBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseJSON) RawJSON() string {
	return r.raw
}

type TokenUpdateResponse struct {
	ID         string                  `json:"id,required" format:"uuid"`
	CfAPIID    string                  `json:"cf_api_id,required"`
	CreatedAt  time.Time               `json:"created_at,required" format:"date-time"`
	ModifiedAt time.Time               `json:"modified_at,required" format:"date-time"`
	Name       string                  `json:"name,required"`
	CreatedBy  string                  `json:"created_by,nullable"`
	Enabled    bool                    `json:"enabled"`
	Legacy     bool                    `json:"legacy"`
	ModifiedBy string                  `json:"modified_by,nullable"`
	JSON       tokenUpdateResponseJSON `json:"-"`
}

// tokenUpdateResponseJSON contains the JSON metadata for the struct
// [TokenUpdateResponse]
type tokenUpdateResponseJSON struct {
	ID          apijson.Field
	CfAPIID     apijson.Field
	CreatedAt   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	CreatedBy   apijson.Field
	Enabled     apijson.Field
	Legacy      apijson.Field
	ModifiedBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type TokenListResponse struct {
	ID         string                `json:"id,required" format:"uuid"`
	CfAPIID    string                `json:"cf_api_id,required"`
	CreatedAt  time.Time             `json:"created_at,required" format:"date-time"`
	ModifiedAt time.Time             `json:"modified_at,required" format:"date-time"`
	Name       string                `json:"name,required"`
	CreatedBy  string                `json:"created_by,nullable"`
	Enabled    bool                  `json:"enabled"`
	Legacy     bool                  `json:"legacy"`
	ModifiedBy string                `json:"modified_by,nullable"`
	JSON       tokenListResponseJSON `json:"-"`
}

// tokenListResponseJSON contains the JSON metadata for the struct
// [TokenListResponse]
type tokenListResponseJSON struct {
	ID          apijson.Field
	CfAPIID     apijson.Field
	CreatedAt   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	CreatedBy   apijson.Field
	Enabled     apijson.Field
	Legacy      apijson.Field
	ModifiedBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenListResponseJSON) RawJSON() string {
	return r.raw
}

type TokenDeleteResponse struct {
	ID         string                  `json:"id,required" format:"uuid"`
	CfAPIID    string                  `json:"cf_api_id,required"`
	CreatedAt  time.Time               `json:"created_at,required" format:"date-time"`
	ModifiedAt time.Time               `json:"modified_at,required" format:"date-time"`
	Name       string                  `json:"name,required"`
	CreatedBy  string                  `json:"created_by,nullable"`
	Enabled    bool                    `json:"enabled"`
	Legacy     bool                    `json:"legacy"`
	ModifiedBy string                  `json:"modified_by,nullable"`
	JSON       tokenDeleteResponseJSON `json:"-"`
}

// tokenDeleteResponseJSON contains the JSON metadata for the struct
// [TokenDeleteResponse]
type tokenDeleteResponseJSON struct {
	ID          apijson.Field
	CfAPIID     apijson.Field
	CreatedAt   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	CreatedBy   apijson.Field
	Enabled     apijson.Field
	Legacy      apijson.Field
	ModifiedBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TokenReadResponse struct {
	ID         string                `json:"id,required" format:"uuid"`
	CfAPIID    string                `json:"cf_api_id,required"`
	CreatedAt  time.Time             `json:"created_at,required" format:"date-time"`
	ModifiedAt time.Time             `json:"modified_at,required" format:"date-time"`
	Name       string                `json:"name,required"`
	CreatedBy  string                `json:"created_by,nullable"`
	Enabled    bool                  `json:"enabled"`
	Legacy     bool                  `json:"legacy"`
	ModifiedBy string                `json:"modified_by,nullable"`
	JSON       tokenReadResponseJSON `json:"-"`
}

// tokenReadResponseJSON contains the JSON metadata for the struct
// [TokenReadResponse]
type tokenReadResponseJSON struct {
	ID          apijson.Field
	CfAPIID     apijson.Field
	CreatedAt   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	CreatedBy   apijson.Field
	Enabled     apijson.Field
	Legacy      apijson.Field
	ModifiedBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenReadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenReadResponseJSON) RawJSON() string {
	return r.raw
}

type TokenNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	CfAPIID   param.Field[string] `json:"cf_api_id,required"`
	CfAPIKey  param.Field[string] `json:"cf_api_key,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r TokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TokenNewResponseEnvelope struct {
	Result  TokenNewResponse             `json:"result,required"`
	Success bool                         `json:"success,required"`
	JSON    tokenNewResponseEnvelopeJSON `json:"-"`
}

// tokenNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenNewResponseEnvelope]
type tokenNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	CfAPIID   param.Field[string] `json:"cf_api_id,required"`
	CfAPIKey  param.Field[string] `json:"cf_api_key,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r TokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TokenUpdateResponseEnvelope struct {
	Result  TokenUpdateResponse             `json:"result,required"`
	Success bool                            `json:"success,required"`
	JSON    tokenUpdateResponseEnvelopeJSON `json:"-"`
}

// tokenUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenUpdateResponseEnvelope]
type tokenUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [TokenListParams]'s query parameters as `url.Values`.
func (r TokenListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type TokenDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type TokenDeleteResponseEnvelope struct {
	Result  TokenDeleteResponse             `json:"result,required"`
	Success bool                            `json:"success,required"`
	JSON    tokenDeleteResponseEnvelopeJSON `json:"-"`
}

// tokenDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenDeleteResponseEnvelope]
type tokenDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenReadParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type TokenReadResponseEnvelope struct {
	Result  TokenReadResponse             `json:"result,required"`
	Success bool                          `json:"success,required"`
	JSON    tokenReadResponseEnvelopeJSON `json:"-"`
}

// tokenReadResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenReadResponseEnvelope]
type tokenReadResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenReadResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenReadResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

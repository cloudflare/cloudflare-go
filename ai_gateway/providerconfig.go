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

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// ProviderConfigService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProviderConfigService] method instead.
type ProviderConfigService struct {
	Options []option.RequestOption
}

// NewProviderConfigService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProviderConfigService(opts ...option.RequestOption) (r *ProviderConfigService) {
	r = &ProviderConfigService{}
	r.Options = opts
	return
}

// Creates a new AI Gateway.
func (r *ProviderConfigService) New(ctx context.Context, gatewayID string, params ProviderConfigNewParams, opts ...option.RequestOption) (res *ProviderConfigNewResponse, err error) {
	var env ProviderConfigNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/provider_configs", params.AccountID, gatewayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all AI Gateway evaluator types configured for the account.
func (r *ProviderConfigService) List(ctx context.Context, gatewayID string, params ProviderConfigListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ProviderConfigListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/provider_configs", params.AccountID, gatewayID)
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
func (r *ProviderConfigService) ListAutoPaging(ctx context.Context, gatewayID string, params ProviderConfigListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ProviderConfigListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, gatewayID, params, opts...))
}

type ProviderConfigNewResponse struct {
	ID            string `json:"id" api:"required"`
	Alias         string `json:"alias" api:"required"`
	DefaultConfig bool   `json:"default_config" api:"required"`
	// gateway id
	GatewayID       string                        `json:"gateway_id" api:"required"`
	ModifiedAt      time.Time                     `json:"modified_at" api:"required" format:"date-time"`
	ProviderSlug    string                        `json:"provider_slug" api:"required"`
	SecretID        string                        `json:"secret_id" api:"required"`
	SecretPreview   string                        `json:"secret_preview" api:"required"`
	RateLimit       float64                       `json:"rate_limit"`
	RateLimitPeriod float64                       `json:"rate_limit_period"`
	JSON            providerConfigNewResponseJSON `json:"-"`
}

// providerConfigNewResponseJSON contains the JSON metadata for the struct
// [ProviderConfigNewResponse]
type providerConfigNewResponseJSON struct {
	ID              apijson.Field
	Alias           apijson.Field
	DefaultConfig   apijson.Field
	GatewayID       apijson.Field
	ModifiedAt      apijson.Field
	ProviderSlug    apijson.Field
	SecretID        apijson.Field
	SecretPreview   apijson.Field
	RateLimit       apijson.Field
	RateLimitPeriod apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProviderConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerConfigNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProviderConfigListResponse struct {
	ID            string `json:"id" api:"required"`
	Alias         string `json:"alias" api:"required"`
	DefaultConfig bool   `json:"default_config" api:"required"`
	// gateway id
	GatewayID       string                         `json:"gateway_id" api:"required"`
	ModifiedAt      time.Time                      `json:"modified_at" api:"required" format:"date-time"`
	ProviderSlug    string                         `json:"provider_slug" api:"required"`
	SecretID        string                         `json:"secret_id" api:"required"`
	SecretPreview   string                         `json:"secret_preview" api:"required"`
	RateLimit       float64                        `json:"rate_limit"`
	RateLimitPeriod float64                        `json:"rate_limit_period"`
	JSON            providerConfigListResponseJSON `json:"-"`
}

// providerConfigListResponseJSON contains the JSON metadata for the struct
// [ProviderConfigListResponse]
type providerConfigListResponseJSON struct {
	ID              apijson.Field
	Alias           apijson.Field
	DefaultConfig   apijson.Field
	GatewayID       apijson.Field
	ModifiedAt      apijson.Field
	ProviderSlug    apijson.Field
	SecretID        apijson.Field
	SecretPreview   apijson.Field
	RateLimit       apijson.Field
	RateLimitPeriod apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProviderConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerConfigListResponseJSON) RawJSON() string {
	return r.raw
}

type ProviderConfigNewParams struct {
	AccountID       param.Field[string]  `path:"account_id" api:"required"`
	Alias           param.Field[string]  `json:"alias" api:"required"`
	DefaultConfig   param.Field[bool]    `json:"default_config" api:"required"`
	ProviderSlug    param.Field[string]  `json:"provider_slug" api:"required"`
	Secret          param.Field[string]  `json:"secret" api:"required"`
	SecretID        param.Field[string]  `json:"secret_id" api:"required"`
	RateLimit       param.Field[float64] `json:"rate_limit"`
	RateLimitPeriod param.Field[float64] `json:"rate_limit_period"`
}

func (r ProviderConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProviderConfigNewResponseEnvelope struct {
	Result  ProviderConfigNewResponse             `json:"result" api:"required"`
	Success bool                                  `json:"success" api:"required"`
	JSON    providerConfigNewResponseEnvelopeJSON `json:"-"`
}

// providerConfigNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProviderConfigNewResponseEnvelope]
type providerConfigNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerConfigNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProviderConfigListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [ProviderConfigListParams]'s query parameters as
// `url.Values`.
func (r ProviderConfigListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AIGatewayService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIGatewayService] method instead.
type AIGatewayService struct {
	Options []option.RequestOption
	Logs    *LogService
}

// NewAIGatewayService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIGatewayService(opts ...option.RequestOption) (r *AIGatewayService) {
	r = &AIGatewayService{}
	r.Options = opts
	r.Logs = NewLogService(opts...)
	return
}

// Create a new Gateway
func (r *AIGatewayService) New(ctx context.Context, params AIGatewayNewParams, opts ...option.RequestOption) (res *AIGatewayNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AIGatewayNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Gateway
func (r *AIGatewayService) Update(ctx context.Context, id string, params AIGatewayUpdateParams, opts ...option.RequestOption) (res *AIGatewayUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AIGatewayUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Gateway's
func (r *AIGatewayService) List(ctx context.Context, params AIGatewayListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AIGatewayListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways", params.AccountID)
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

// List Gateway's
func (r *AIGatewayService) ListAutoPaging(ctx context.Context, params AIGatewayListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AIGatewayListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a Gateway
func (r *AIGatewayService) Delete(ctx context.Context, id string, body AIGatewayDeleteParams, opts ...option.RequestOption) (res *AIGatewayDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AIGatewayDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a Gateway
func (r *AIGatewayService) Get(ctx context.Context, id string, query AIGatewayGetParams, opts ...option.RequestOption) (res *AIGatewayGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AIGatewayGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIGatewayNewResponse struct {
	Task AIGatewayNewResponseTask `json:"task,required"`
	JSON aiGatewayNewResponseJSON `json:"-"`
}

// aiGatewayNewResponseJSON contains the JSON metadata for the struct
// [AIGatewayNewResponse]
type aiGatewayNewResponseJSON struct {
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayNewResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayNewResponseTask struct {
	// gateway id
	ID                      string                       `json:"id,required"`
	CacheInvalidateOnUpdate bool                         `json:"cache_invalidate_on_update,required"`
	CacheTTL                int64                        `json:"cache_ttl,required"`
	CollectLogs             bool                         `json:"collect_logs,required"`
	CreatedAt               time.Time                    `json:"created_at,required" format:"date-time"`
	ModifiedAt              time.Time                    `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                        `json:"rate_limiting_interval"`
	RateLimitingLimit       int64                        `json:"rate_limiting_limit"`
	RateLimitingTechnique   string                       `json:"rate_limiting_technique"`
	JSON                    aiGatewayNewResponseTaskJSON `json:"-"`
}

// aiGatewayNewResponseTaskJSON contains the JSON metadata for the struct
// [AIGatewayNewResponseTask]
type aiGatewayNewResponseTaskJSON struct {
	ID                      apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AIGatewayNewResponseTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayNewResponseTaskJSON) RawJSON() string {
	return r.raw
}

type AIGatewayUpdateResponse struct {
	// gateway id
	ID                      string                      `json:"id,required"`
	CacheInvalidateOnUpdate bool                        `json:"cache_invalidate_on_update,required"`
	CacheTTL                int64                       `json:"cache_ttl,required"`
	CollectLogs             bool                        `json:"collect_logs,required"`
	CreatedAt               time.Time                   `json:"created_at,required" format:"date-time"`
	ModifiedAt              time.Time                   `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                       `json:"rate_limiting_interval"`
	RateLimitingLimit       int64                       `json:"rate_limiting_limit"`
	RateLimitingTechnique   string                      `json:"rate_limiting_technique"`
	JSON                    aiGatewayUpdateResponseJSON `json:"-"`
}

// aiGatewayUpdateResponseJSON contains the JSON metadata for the struct
// [AIGatewayUpdateResponse]
type aiGatewayUpdateResponseJSON struct {
	ID                      apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AIGatewayUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayListResponse struct {
	// gateway id
	ID                      string                    `json:"id,required"`
	CacheInvalidateOnUpdate bool                      `json:"cache_invalidate_on_update,required"`
	CacheTTL                int64                     `json:"cache_ttl,required"`
	CollectLogs             bool                      `json:"collect_logs,required"`
	CreatedAt               time.Time                 `json:"created_at,required" format:"date-time"`
	ModifiedAt              time.Time                 `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                     `json:"rate_limiting_interval"`
	RateLimitingLimit       int64                     `json:"rate_limiting_limit"`
	RateLimitingTechnique   string                    `json:"rate_limiting_technique"`
	JSON                    aiGatewayListResponseJSON `json:"-"`
}

// aiGatewayListResponseJSON contains the JSON metadata for the struct
// [AIGatewayListResponse]
type aiGatewayListResponseJSON struct {
	ID                      apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AIGatewayListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayListResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayDeleteResponse struct {
	// gateway id
	ID                      string                      `json:"id,required"`
	CacheInvalidateOnUpdate bool                        `json:"cache_invalidate_on_update,required"`
	CacheTTL                int64                       `json:"cache_ttl,required"`
	CollectLogs             bool                        `json:"collect_logs,required"`
	CreatedAt               time.Time                   `json:"created_at,required" format:"date-time"`
	ModifiedAt              time.Time                   `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                       `json:"rate_limiting_interval"`
	RateLimitingLimit       int64                       `json:"rate_limiting_limit"`
	RateLimitingTechnique   string                      `json:"rate_limiting_technique"`
	JSON                    aiGatewayDeleteResponseJSON `json:"-"`
}

// aiGatewayDeleteResponseJSON contains the JSON metadata for the struct
// [AIGatewayDeleteResponse]
type aiGatewayDeleteResponseJSON struct {
	ID                      apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AIGatewayDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayGetResponse struct {
	// gateway id
	ID                      string                   `json:"id,required"`
	CacheInvalidateOnUpdate bool                     `json:"cache_invalidate_on_update,required"`
	CacheTTL                int64                    `json:"cache_ttl,required"`
	CollectLogs             bool                     `json:"collect_logs,required"`
	CreatedAt               time.Time                `json:"created_at,required" format:"date-time"`
	ModifiedAt              time.Time                `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                    `json:"rate_limiting_interval"`
	RateLimitingLimit       int64                    `json:"rate_limiting_limit"`
	RateLimitingTechnique   string                   `json:"rate_limiting_technique"`
	JSON                    aiGatewayGetResponseJSON `json:"-"`
}

// aiGatewayGetResponseJSON contains the JSON metadata for the struct
// [AIGatewayGetResponse]
type aiGatewayGetResponseJSON struct {
	ID                      apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AIGatewayGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayGetResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// gateway id
	ID                      param.Field[string] `json:"id,required"`
	CacheInvalidateOnUpdate param.Field[bool]   `json:"cache_invalidate_on_update,required"`
	CacheTTL                param.Field[int64]  `json:"cache_ttl,required"`
	CollectLogs             param.Field[bool]   `json:"collect_logs,required"`
	RateLimitingInterval    param.Field[int64]  `json:"rate_limiting_interval"`
	RateLimitingLimit       param.Field[int64]  `json:"rate_limiting_limit"`
	RateLimitingTechnique   param.Field[string] `json:"rate_limiting_technique"`
}

func (r AIGatewayNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIGatewayNewResponseEnvelope struct {
	Result  AIGatewayNewResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    aiGatewayNewResponseEnvelopeJSON `json:"-"`
}

// aiGatewayNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIGatewayNewResponseEnvelope]
type aiGatewayNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayUpdateParams struct {
	AccountID               param.Field[string] `path:"account_id,required"`
	CacheInvalidateOnUpdate param.Field[bool]   `json:"cache_invalidate_on_update,required"`
	CacheTTL                param.Field[int64]  `json:"cache_ttl,required"`
	CollectLogs             param.Field[bool]   `json:"collect_logs,required"`
	RateLimitingInterval    param.Field[int64]  `json:"rate_limiting_interval"`
	RateLimitingLimit       param.Field[int64]  `json:"rate_limiting_limit"`
	RateLimitingTechnique   param.Field[string] `json:"rate_limiting_technique"`
}

func (r AIGatewayUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIGatewayUpdateResponseEnvelope struct {
	Result  AIGatewayUpdateResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    aiGatewayUpdateResponseEnvelopeJSON `json:"-"`
}

// aiGatewayUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIGatewayUpdateResponseEnvelope]
type aiGatewayUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// gateway id
	ID param.Field[string] `query:"id"`
	// Order By Column Name
	OrderBy param.Field[string] `query:"order_by"`
	Page    param.Field[int64]  `query:"page"`
	PerPage param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [AIGatewayListParams]'s query parameters as `url.Values`.
func (r AIGatewayListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIGatewayDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type AIGatewayDeleteResponseEnvelope struct {
	Result  AIGatewayDeleteResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    aiGatewayDeleteResponseEnvelopeJSON `json:"-"`
}

// aiGatewayDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIGatewayDeleteResponseEnvelope]
type aiGatewayDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type AIGatewayGetResponseEnvelope struct {
	Result  AIGatewayGetResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    aiGatewayGetResponseEnvelopeJSON `json:"-"`
}

// aiGatewayGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIGatewayGetResponseEnvelope]
type aiGatewayGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

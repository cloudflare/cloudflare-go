// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/user"
)

// PoolService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPoolService] method instead.
type PoolService struct {
	Options    []option.RequestOption
	Health     *PoolHealthService
	References *PoolReferenceService
}

// NewPoolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPoolService(opts ...option.RequestOption) (r *PoolService) {
	r = &PoolService{}
	r.Options = opts
	r.Health = NewPoolHealthService(opts...)
	r.References = NewPoolReferenceService(opts...)
	return
}

// Create a new pool.
func (r *PoolService) New(ctx context.Context, params PoolNewParams, opts ...option.RequestOption) (res *user.Pool, err error) {
	opts = append(r.Options[:], opts...)
	var env PoolNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a configured pool.
func (r *PoolService) Update(ctx context.Context, poolID string, params PoolUpdateParams, opts ...option.RequestOption) (res *user.Pool, err error) {
	opts = append(r.Options[:], opts...)
	var env PoolUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", params.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured pools.
func (r *PoolService) List(ctx context.Context, params PoolListParams, opts ...option.RequestOption) (res *pagination.SinglePage[user.Pool], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", params.AccountID)
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

// List configured pools.
func (r *PoolService) ListAutoPaging(ctx context.Context, params PoolListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[user.Pool] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Delete a configured pool.
func (r *PoolService) Delete(ctx context.Context, poolID string, params PoolDeleteParams, opts ...option.RequestOption) (res *PoolDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PoolDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", params.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing pool, overwriting the supplied properties.
func (r *PoolService) Edit(ctx context.Context, poolID string, params PoolEditParams, opts ...option.RequestOption) (res *user.Pool, err error) {
	opts = append(r.Options[:], opts...)
	var env PoolEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", params.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured pool.
func (r *PoolService) Get(ctx context.Context, poolID string, query PoolGetParams, opts ...option.RequestOption) (res *user.Pool, err error) {
	opts = append(r.Options[:], opts...)
	var env PoolGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", query.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PoolDeleteResponse struct {
	ID   string                 `json:"id"`
	JSON poolDeleteResponseJSON `json:"-"`
}

// poolDeleteResponseJSON contains the JSON metadata for the struct
// [PoolDeleteResponse]
type poolDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PoolNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]OriginItemParam] `json:"origins,required"`
	// A human-readable description of the pool.
	Description param.Field[string] `json:"description"`
	// Whether to enable (the default) or disable this pool. Disabled pools will not
	// receive traffic and are excluded from health checks. Disabling a pool will cause
	// any load balancers using it to failover to the next pool (if any).
	Enabled param.Field[bool] `json:"enabled"`
	// The latitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, longitude must also be set.
	Latitude param.Field[float64] `json:"latitude"`
	// Configures load shedding policies and percentages for the pool.
	LoadShedding param.Field[LoadSheddingParam] `json:"load_shedding"`
	// The longitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, latitude must also be set.
	Longitude param.Field[float64] `json:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve
	// traffic. If the number of healthy origins falls below this number, the pool will
	// be marked unhealthy and will failover to the next available pool.
	MinimumOrigins param.Field[int64] `json:"minimum_origins"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[interface{}] `json:"monitor"`
	// This field is now deprecated. It has been moved to Cloudflare's Centralized
	// Notification service
	// https://developers.cloudflare.com/fundamentals/notifications/. The email address
	// to send health status notifications to. This can be an individual mailbox or a
	// mailing list. Multiple emails can be supplied as a comma delimited list.
	NotificationEmail param.Field[string] `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter param.Field[NotificationFilterParam] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[OriginSteeringParam] `json:"origin_steering"`
}

func (r PoolNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PoolNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   user.Pool                                                 `json:"result,required"`
	// Whether the API call was successful
	Success PoolNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    poolNewResponseEnvelopeJSON    `json:"-"`
}

// poolNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PoolNewResponseEnvelope]
type poolNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PoolNewResponseEnvelopeSuccess bool

const (
	PoolNewResponseEnvelopeSuccessTrue PoolNewResponseEnvelopeSuccess = true
)

func (r PoolNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PoolNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PoolUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]OriginItemParam] `json:"origins,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]PoolUpdateParamsCheckRegion] `json:"check_regions"`
	// A human-readable description of the pool.
	Description param.Field[string] `json:"description"`
	// Whether to enable (the default) or disable this pool. Disabled pools will not
	// receive traffic and are excluded from health checks. Disabling a pool will cause
	// any load balancers using it to failover to the next pool (if any).
	Enabled param.Field[bool] `json:"enabled"`
	// The latitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, longitude must also be set.
	Latitude param.Field[float64] `json:"latitude"`
	// Configures load shedding policies and percentages for the pool.
	LoadShedding param.Field[LoadSheddingParam] `json:"load_shedding"`
	// The longitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, latitude must also be set.
	Longitude param.Field[float64] `json:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve
	// traffic. If the number of healthy origins falls below this number, the pool will
	// be marked unhealthy and will failover to the next available pool.
	MinimumOrigins param.Field[int64] `json:"minimum_origins"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[interface{}] `json:"monitor"`
	// This field is now deprecated. It has been moved to Cloudflare's Centralized
	// Notification service
	// https://developers.cloudflare.com/fundamentals/notifications/. The email address
	// to send health status notifications to. This can be an individual mailbox or a
	// mailing list. Multiple emails can be supplied as a comma delimited list.
	NotificationEmail param.Field[string] `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter param.Field[NotificationFilterParam] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[OriginSteeringParam] `json:"origin_steering"`
}

func (r PoolUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type PoolUpdateParamsCheckRegion string

const (
	PoolUpdateParamsCheckRegionWnam       PoolUpdateParamsCheckRegion = "WNAM"
	PoolUpdateParamsCheckRegionEnam       PoolUpdateParamsCheckRegion = "ENAM"
	PoolUpdateParamsCheckRegionWeu        PoolUpdateParamsCheckRegion = "WEU"
	PoolUpdateParamsCheckRegionEeu        PoolUpdateParamsCheckRegion = "EEU"
	PoolUpdateParamsCheckRegionNsam       PoolUpdateParamsCheckRegion = "NSAM"
	PoolUpdateParamsCheckRegionSsam       PoolUpdateParamsCheckRegion = "SSAM"
	PoolUpdateParamsCheckRegionOc         PoolUpdateParamsCheckRegion = "OC"
	PoolUpdateParamsCheckRegionMe         PoolUpdateParamsCheckRegion = "ME"
	PoolUpdateParamsCheckRegionNaf        PoolUpdateParamsCheckRegion = "NAF"
	PoolUpdateParamsCheckRegionSaf        PoolUpdateParamsCheckRegion = "SAF"
	PoolUpdateParamsCheckRegionSas        PoolUpdateParamsCheckRegion = "SAS"
	PoolUpdateParamsCheckRegionSeas       PoolUpdateParamsCheckRegion = "SEAS"
	PoolUpdateParamsCheckRegionNeas       PoolUpdateParamsCheckRegion = "NEAS"
	PoolUpdateParamsCheckRegionAllRegions PoolUpdateParamsCheckRegion = "ALL_REGIONS"
)

func (r PoolUpdateParamsCheckRegion) IsKnown() bool {
	switch r {
	case PoolUpdateParamsCheckRegionWnam, PoolUpdateParamsCheckRegionEnam, PoolUpdateParamsCheckRegionWeu, PoolUpdateParamsCheckRegionEeu, PoolUpdateParamsCheckRegionNsam, PoolUpdateParamsCheckRegionSsam, PoolUpdateParamsCheckRegionOc, PoolUpdateParamsCheckRegionMe, PoolUpdateParamsCheckRegionNaf, PoolUpdateParamsCheckRegionSaf, PoolUpdateParamsCheckRegionSas, PoolUpdateParamsCheckRegionSeas, PoolUpdateParamsCheckRegionNeas, PoolUpdateParamsCheckRegionAllRegions:
		return true
	}
	return false
}

type PoolUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   user.Pool                                                 `json:"result,required"`
	// Whether the API call was successful
	Success PoolUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    poolUpdateResponseEnvelopeJSON    `json:"-"`
}

// poolUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PoolUpdateResponseEnvelope]
type poolUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PoolUpdateResponseEnvelopeSuccess bool

const (
	PoolUpdateResponseEnvelopeSuccessTrue PoolUpdateResponseEnvelopeSuccess = true
)

func (r PoolUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PoolUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PoolListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[interface{}] `query:"monitor"`
}

// URLQuery serializes [PoolListParams]'s query parameters as `url.Values`.
func (r PoolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PoolDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r PoolDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type PoolDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   PoolDeleteResponse                                        `json:"result,required"`
	// Whether the API call was successful
	Success PoolDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    poolDeleteResponseEnvelopeJSON    `json:"-"`
}

// poolDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [PoolDeleteResponseEnvelope]
type poolDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PoolDeleteResponseEnvelopeSuccess bool

const (
	PoolDeleteResponseEnvelopeSuccessTrue PoolDeleteResponseEnvelopeSuccess = true
)

func (r PoolDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PoolDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PoolEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]PoolEditParamsCheckRegion] `json:"check_regions"`
	// A human-readable description of the pool.
	Description param.Field[string] `json:"description"`
	// Whether to enable (the default) or disable this pool. Disabled pools will not
	// receive traffic and are excluded from health checks. Disabling a pool will cause
	// any load balancers using it to failover to the next pool (if any).
	Enabled param.Field[bool] `json:"enabled"`
	// The latitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, longitude must also be set.
	Latitude param.Field[float64] `json:"latitude"`
	// Configures load shedding policies and percentages for the pool.
	LoadShedding param.Field[LoadSheddingParam] `json:"load_shedding"`
	// The longitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, latitude must also be set.
	Longitude param.Field[float64] `json:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve
	// traffic. If the number of healthy origins falls below this number, the pool will
	// be marked unhealthy and will failover to the next available pool.
	MinimumOrigins param.Field[int64] `json:"minimum_origins"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[interface{}] `json:"monitor"`
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name"`
	// This field is now deprecated. It has been moved to Cloudflare's Centralized
	// Notification service
	// https://developers.cloudflare.com/fundamentals/notifications/. The email address
	// to send health status notifications to. This can be an individual mailbox or a
	// mailing list. Multiple emails can be supplied as a comma delimited list.
	NotificationEmail param.Field[string] `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter param.Field[NotificationFilterParam] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[OriginSteeringParam] `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]OriginItemParam] `json:"origins"`
}

func (r PoolEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type PoolEditParamsCheckRegion string

const (
	PoolEditParamsCheckRegionWnam       PoolEditParamsCheckRegion = "WNAM"
	PoolEditParamsCheckRegionEnam       PoolEditParamsCheckRegion = "ENAM"
	PoolEditParamsCheckRegionWeu        PoolEditParamsCheckRegion = "WEU"
	PoolEditParamsCheckRegionEeu        PoolEditParamsCheckRegion = "EEU"
	PoolEditParamsCheckRegionNsam       PoolEditParamsCheckRegion = "NSAM"
	PoolEditParamsCheckRegionSsam       PoolEditParamsCheckRegion = "SSAM"
	PoolEditParamsCheckRegionOc         PoolEditParamsCheckRegion = "OC"
	PoolEditParamsCheckRegionMe         PoolEditParamsCheckRegion = "ME"
	PoolEditParamsCheckRegionNaf        PoolEditParamsCheckRegion = "NAF"
	PoolEditParamsCheckRegionSaf        PoolEditParamsCheckRegion = "SAF"
	PoolEditParamsCheckRegionSas        PoolEditParamsCheckRegion = "SAS"
	PoolEditParamsCheckRegionSeas       PoolEditParamsCheckRegion = "SEAS"
	PoolEditParamsCheckRegionNeas       PoolEditParamsCheckRegion = "NEAS"
	PoolEditParamsCheckRegionAllRegions PoolEditParamsCheckRegion = "ALL_REGIONS"
)

func (r PoolEditParamsCheckRegion) IsKnown() bool {
	switch r {
	case PoolEditParamsCheckRegionWnam, PoolEditParamsCheckRegionEnam, PoolEditParamsCheckRegionWeu, PoolEditParamsCheckRegionEeu, PoolEditParamsCheckRegionNsam, PoolEditParamsCheckRegionSsam, PoolEditParamsCheckRegionOc, PoolEditParamsCheckRegionMe, PoolEditParamsCheckRegionNaf, PoolEditParamsCheckRegionSaf, PoolEditParamsCheckRegionSas, PoolEditParamsCheckRegionSeas, PoolEditParamsCheckRegionNeas, PoolEditParamsCheckRegionAllRegions:
		return true
	}
	return false
}

type PoolEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   user.Pool                                                 `json:"result,required"`
	// Whether the API call was successful
	Success PoolEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    poolEditResponseEnvelopeJSON    `json:"-"`
}

// poolEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [PoolEditResponseEnvelope]
type poolEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PoolEditResponseEnvelopeSuccess bool

const (
	PoolEditResponseEnvelopeSuccessTrue PoolEditResponseEnvelopeSuccess = true
)

func (r PoolEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PoolEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PoolGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PoolGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   user.Pool                                                 `json:"result,required"`
	// Whether the API call was successful
	Success PoolGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    poolGetResponseEnvelopeJSON    `json:"-"`
}

// poolGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PoolGetResponseEnvelope]
type poolGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PoolGetResponseEnvelopeSuccess bool

const (
	PoolGetResponseEnvelopeSuccessTrue PoolGetResponseEnvelopeSuccess = true
)

func (r PoolGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PoolGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

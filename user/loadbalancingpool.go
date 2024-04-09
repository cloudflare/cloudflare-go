// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/load_balancers"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// LoadBalancingPoolService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancingPoolService] method
// instead.
type LoadBalancingPoolService struct {
	Options []option.RequestOption
}

// NewLoadBalancingPoolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancingPoolService(opts ...option.RequestOption) (r *LoadBalancingPoolService) {
	r = &LoadBalancingPoolService{}
	r.Options = opts
	return
}

// Create a new pool.
func (r *LoadBalancingPoolService) New(ctx context.Context, body LoadBalancingPoolNewParams, opts ...option.RequestOption) (res *load_balancers.Pool, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingPoolNewResponseEnvelope
	path := "user/load_balancers/pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a configured pool.
func (r *LoadBalancingPoolService) Update(ctx context.Context, poolID string, body LoadBalancingPoolUpdateParams, opts ...option.RequestOption) (res *load_balancers.Pool, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingPoolUpdateResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured pools.
func (r *LoadBalancingPoolService) List(ctx context.Context, query LoadBalancingPoolListParams, opts ...option.RequestOption) (res *pagination.SinglePage[load_balancers.Pool], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/load_balancers/pools"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *LoadBalancingPoolService) ListAutoPaging(ctx context.Context, query LoadBalancingPoolListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[load_balancers.Pool] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a configured pool.
func (r *LoadBalancingPoolService) Delete(ctx context.Context, poolID string, body LoadBalancingPoolDeleteParams, opts ...option.RequestOption) (res *LoadBalancingPoolDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingPoolDeleteResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing pool, overwriting the supplied properties.
func (r *LoadBalancingPoolService) Edit(ctx context.Context, poolID string, body LoadBalancingPoolEditParams, opts ...option.RequestOption) (res *load_balancers.Pool, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingPoolEditResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured pool.
func (r *LoadBalancingPoolService) Get(ctx context.Context, poolID string, opts ...option.RequestOption) (res *load_balancers.Pool, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingPoolGetResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch the latest pool health status for a single pool.
func (r *LoadBalancingPoolService) Health(ctx context.Context, poolID string, opts ...option.RequestOption) (res *LoadBalancingPoolHealthResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingPoolHealthResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s/health", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Preview pool health using provided monitor details. The returned preview_id can
// be used in the preview endpoint to retrieve the results.
func (r *LoadBalancingPoolService) Preview(ctx context.Context, poolID string, body LoadBalancingPoolPreviewParams, opts ...option.RequestOption) (res *LoadBalancingPoolPreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingPoolPreviewResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s/preview", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the list of resources that reference the provided pool.
func (r *LoadBalancingPoolService) References(ctx context.Context, poolID string, opts ...option.RequestOption) (res *[]LoadBalancingPoolReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingPoolReferencesResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s/references", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancingPoolDeleteResponse struct {
	ID   string                              `json:"id"`
	JSON loadBalancingPoolDeleteResponseJSON `json:"-"`
}

// loadBalancingPoolDeleteResponseJSON contains the JSON metadata for the struct
// [LoadBalancingPoolDeleteResponse]
type loadBalancingPoolDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingPoolDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingPoolDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A list of regions from which to run health checks. Null means every Cloudflare
// data center.
//
// Union satisfied by [user.LoadBalancingPoolHealthResponseUnknown] or
// [shared.UnionString].
type LoadBalancingPoolHealthResponseUnion interface {
	ImplementsUserLoadBalancingPoolHealthResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LoadBalancingPoolHealthResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type LoadBalancingPoolPreviewResponse struct {
	// Monitored pool IDs mapped to their respective names.
	Pools     map[string]string                    `json:"pools"`
	PreviewID string                               `json:"preview_id"`
	JSON      loadBalancingPoolPreviewResponseJSON `json:"-"`
}

// loadBalancingPoolPreviewResponseJSON contains the JSON metadata for the struct
// [LoadBalancingPoolPreviewResponse]
type loadBalancingPoolPreviewResponseJSON struct {
	Pools       apijson.Field
	PreviewID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingPoolPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingPoolPreviewResponseJSON) RawJSON() string {
	return r.raw
}

type LoadBalancingPoolReferencesResponse struct {
	ReferenceType LoadBalancingPoolReferencesResponseReferenceType `json:"reference_type"`
	ResourceID    string                                           `json:"resource_id"`
	ResourceName  string                                           `json:"resource_name"`
	ResourceType  string                                           `json:"resource_type"`
	JSON          loadBalancingPoolReferencesResponseJSON          `json:"-"`
}

// loadBalancingPoolReferencesResponseJSON contains the JSON metadata for the
// struct [LoadBalancingPoolReferencesResponse]
type loadBalancingPoolReferencesResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancingPoolReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingPoolReferencesResponseJSON) RawJSON() string {
	return r.raw
}

type LoadBalancingPoolReferencesResponseReferenceType string

const (
	LoadBalancingPoolReferencesResponseReferenceTypeStar     LoadBalancingPoolReferencesResponseReferenceType = "*"
	LoadBalancingPoolReferencesResponseReferenceTypeReferral LoadBalancingPoolReferencesResponseReferenceType = "referral"
	LoadBalancingPoolReferencesResponseReferenceTypeReferrer LoadBalancingPoolReferencesResponseReferenceType = "referrer"
)

func (r LoadBalancingPoolReferencesResponseReferenceType) IsKnown() bool {
	switch r {
	case LoadBalancingPoolReferencesResponseReferenceTypeStar, LoadBalancingPoolReferencesResponseReferenceTypeReferral, LoadBalancingPoolReferencesResponseReferenceTypeReferrer:
		return true
	}
	return false
}

type LoadBalancingPoolNewParams struct {
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]load_balancers.OriginParam] `json:"origins,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]load_balancers.CheckRegion] `json:"check_regions"`
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
	LoadShedding param.Field[load_balancers.LoadSheddingParam] `json:"load_shedding"`
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
	NotificationFilter param.Field[load_balancers.NotificationFilterParam] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[load_balancers.OriginSteeringParam] `json:"origin_steering"`
}

func (r LoadBalancingPoolNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancingPoolNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   load_balancers.Pool   `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancingPoolNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancingPoolNewResponseEnvelopeJSON    `json:"-"`
}

// loadBalancingPoolNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancingPoolNewResponseEnvelope]
type loadBalancingPoolNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingPoolNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingPoolNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingPoolNewResponseEnvelopeSuccess bool

const (
	LoadBalancingPoolNewResponseEnvelopeSuccessTrue LoadBalancingPoolNewResponseEnvelopeSuccess = true
)

func (r LoadBalancingPoolNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingPoolNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingPoolUpdateParams struct {
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]load_balancers.OriginParam] `json:"origins,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]load_balancers.CheckRegion] `json:"check_regions"`
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
	LoadShedding param.Field[load_balancers.LoadSheddingParam] `json:"load_shedding"`
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
	NotificationFilter param.Field[load_balancers.NotificationFilterParam] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[load_balancers.OriginSteeringParam] `json:"origin_steering"`
}

func (r LoadBalancingPoolUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancingPoolUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   load_balancers.Pool   `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancingPoolUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancingPoolUpdateResponseEnvelopeJSON    `json:"-"`
}

// loadBalancingPoolUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancingPoolUpdateResponseEnvelope]
type loadBalancingPoolUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingPoolUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingPoolUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingPoolUpdateResponseEnvelopeSuccess bool

const (
	LoadBalancingPoolUpdateResponseEnvelopeSuccessTrue LoadBalancingPoolUpdateResponseEnvelopeSuccess = true
)

func (r LoadBalancingPoolUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingPoolUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingPoolListParams struct {
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[interface{}] `query:"monitor"`
}

// URLQuery serializes [LoadBalancingPoolListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancingPoolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LoadBalancingPoolDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r LoadBalancingPoolDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type LoadBalancingPoolDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   LoadBalancingPoolDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancingPoolDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancingPoolDeleteResponseEnvelopeJSON    `json:"-"`
}

// loadBalancingPoolDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancingPoolDeleteResponseEnvelope]
type loadBalancingPoolDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingPoolDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingPoolDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingPoolDeleteResponseEnvelopeSuccess bool

const (
	LoadBalancingPoolDeleteResponseEnvelopeSuccessTrue LoadBalancingPoolDeleteResponseEnvelopeSuccess = true
)

func (r LoadBalancingPoolDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingPoolDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingPoolEditParams struct {
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]load_balancers.CheckRegion] `json:"check_regions"`
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
	LoadShedding param.Field[load_balancers.LoadSheddingParam] `json:"load_shedding"`
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
	NotificationFilter param.Field[load_balancers.NotificationFilterParam] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[load_balancers.OriginSteeringParam] `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]load_balancers.OriginParam] `json:"origins"`
}

func (r LoadBalancingPoolEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancingPoolEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   load_balancers.Pool   `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancingPoolEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancingPoolEditResponseEnvelopeJSON    `json:"-"`
}

// loadBalancingPoolEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancingPoolEditResponseEnvelope]
type loadBalancingPoolEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingPoolEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingPoolEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingPoolEditResponseEnvelopeSuccess bool

const (
	LoadBalancingPoolEditResponseEnvelopeSuccessTrue LoadBalancingPoolEditResponseEnvelopeSuccess = true
)

func (r LoadBalancingPoolEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingPoolEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingPoolGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   load_balancers.Pool   `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancingPoolGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancingPoolGetResponseEnvelopeJSON    `json:"-"`
}

// loadBalancingPoolGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancingPoolGetResponseEnvelope]
type loadBalancingPoolGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingPoolGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingPoolGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingPoolGetResponseEnvelopeSuccess bool

const (
	LoadBalancingPoolGetResponseEnvelopeSuccessTrue LoadBalancingPoolGetResponseEnvelopeSuccess = true
)

func (r LoadBalancingPoolGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingPoolGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingPoolHealthResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	Result LoadBalancingPoolHealthResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancingPoolHealthResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancingPoolHealthResponseEnvelopeJSON    `json:"-"`
}

// loadBalancingPoolHealthResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancingPoolHealthResponseEnvelope]
type loadBalancingPoolHealthResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingPoolHealthResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingPoolHealthResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingPoolHealthResponseEnvelopeSuccess bool

const (
	LoadBalancingPoolHealthResponseEnvelopeSuccessTrue LoadBalancingPoolHealthResponseEnvelopeSuccess = true
)

func (r LoadBalancingPoolHealthResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingPoolHealthResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingPoolPreviewParams struct {
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes,required"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown param.Field[int64] `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp param.Field[int64] `json:"consecutive_up"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header param.Field[interface{}] `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method param.Field[string] `json:"method"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path param.Field[string] `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port param.Field[int64] `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone param.Field[string] `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type param.Field[LoadBalancingPoolPreviewParamsType] `json:"type"`
}

func (r LoadBalancingPoolPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancingPoolPreviewParamsType string

const (
	LoadBalancingPoolPreviewParamsTypeHTTP     LoadBalancingPoolPreviewParamsType = "http"
	LoadBalancingPoolPreviewParamsTypeHTTPS    LoadBalancingPoolPreviewParamsType = "https"
	LoadBalancingPoolPreviewParamsTypeTCP      LoadBalancingPoolPreviewParamsType = "tcp"
	LoadBalancingPoolPreviewParamsTypeUdpIcmp  LoadBalancingPoolPreviewParamsType = "udp_icmp"
	LoadBalancingPoolPreviewParamsTypeIcmpPing LoadBalancingPoolPreviewParamsType = "icmp_ping"
	LoadBalancingPoolPreviewParamsTypeSmtp     LoadBalancingPoolPreviewParamsType = "smtp"
)

func (r LoadBalancingPoolPreviewParamsType) IsKnown() bool {
	switch r {
	case LoadBalancingPoolPreviewParamsTypeHTTP, LoadBalancingPoolPreviewParamsTypeHTTPS, LoadBalancingPoolPreviewParamsTypeTCP, LoadBalancingPoolPreviewParamsTypeUdpIcmp, LoadBalancingPoolPreviewParamsTypeIcmpPing, LoadBalancingPoolPreviewParamsTypeSmtp:
		return true
	}
	return false
}

type LoadBalancingPoolPreviewResponseEnvelope struct {
	Errors   []shared.ResponseInfo            `json:"errors,required"`
	Messages []shared.ResponseInfo            `json:"messages,required"`
	Result   LoadBalancingPoolPreviewResponse `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancingPoolPreviewResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancingPoolPreviewResponseEnvelopeJSON    `json:"-"`
}

// loadBalancingPoolPreviewResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancingPoolPreviewResponseEnvelope]
type loadBalancingPoolPreviewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingPoolPreviewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingPoolPreviewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingPoolPreviewResponseEnvelopeSuccess bool

const (
	LoadBalancingPoolPreviewResponseEnvelopeSuccessTrue LoadBalancingPoolPreviewResponseEnvelopeSuccess = true
)

func (r LoadBalancingPoolPreviewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingPoolPreviewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingPoolReferencesResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// List of resources that reference a given pool.
	Result []LoadBalancingPoolReferencesResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    LoadBalancingPoolReferencesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo LoadBalancingPoolReferencesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       loadBalancingPoolReferencesResponseEnvelopeJSON       `json:"-"`
}

// loadBalancingPoolReferencesResponseEnvelopeJSON contains the JSON metadata for
// the struct [LoadBalancingPoolReferencesResponseEnvelope]
type loadBalancingPoolReferencesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingPoolReferencesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingPoolReferencesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingPoolReferencesResponseEnvelopeSuccess bool

const (
	LoadBalancingPoolReferencesResponseEnvelopeSuccessTrue LoadBalancingPoolReferencesResponseEnvelopeSuccess = true
)

func (r LoadBalancingPoolReferencesResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingPoolReferencesResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingPoolReferencesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                   `json:"total_count"`
	JSON       loadBalancingPoolReferencesResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancingPoolReferencesResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [LoadBalancingPoolReferencesResponseEnvelopeResultInfo]
type loadBalancingPoolReferencesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingPoolReferencesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingPoolReferencesResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

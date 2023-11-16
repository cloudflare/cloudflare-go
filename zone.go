// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewZoneService] method instead.
type ZoneService struct {
	Options       []option.RequestOption
	LoadBalancers *ZoneLoadBalancerService
	Dnssecs       *ZoneDnssecService
	RateLimits    *ZoneRateLimitService
	Settings      *ZoneSettingService
}

// NewZoneService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZoneService(opts ...option.RequestOption) (r *ZoneService) {
	r = &ZoneService{}
	r.Options = opts
	r.LoadBalancers = NewZoneLoadBalancerService(opts...)
	r.Dnssecs = NewZoneDnssecService(opts...)
	r.RateLimits = NewZoneRateLimitService(opts...)
	r.Settings = NewZoneSettingService(opts...)
	return
}

// Create Zone
func (r *ZoneService) New(ctx context.Context, body ZoneNewParams, opts ...option.RequestOption) (res *ZoneNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Zone Details
func (r *ZoneService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ZoneGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Editz a zone. Only one zone property can be changed at a time.
func (r *ZoneService) Update(ctx context.Context, identifier string, body ZoneUpdateParams, opts ...option.RequestOption) (res *ZoneUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists, searches, sorts, and filters your zones.
func (r *ZoneService) List(ctx context.Context, query ZoneListParams, opts ...option.RequestOption) (res *ZoneListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing zone.
func (r *ZoneService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ZoneAPIResponseSingleID, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// The session_affinity specifies the type of session affinity the load balancer
// should use unless specified as "none" or ""(default). The supported types are
// "cookie" and "ip_cookie". "cookie" - On the first request to a proxied load
// balancer, a cookie is generated, encoding information of which origin the
// request will be forwarded to. Subsequent requests, by the same client to the
// same load balancer, will be sent to the origin server the cookie encodes, for
// the duration of the cookie and as long as the origin server remains healthy. If
// the cookie has expired or the origin server is unhealthy then a new origin
// server is calculated and used. "ip_cookie" behaves the same as "cookie" except
// the initial origin selection is stable and based on the client’s ip address.
type SessionAffinity string

const (
	SessionAffinityNone     SessionAffinity = "none"
	SessionAffinityCookie   SessionAffinity = "cookie"
	SessionAffinityIPCookie SessionAffinity = "ip_cookie"
	SessionAffinityEmpty    SessionAffinity = "\"\""
)

type ZoneAPIResponseSingleID struct {
	Errors   []ZoneAPIResponseSingleIDError   `json:"errors"`
	Messages []ZoneAPIResponseSingleIDMessage `json:"messages"`
	Result   ZoneAPIResponseSingleIDResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success bool                        `json:"success"`
	JSON    zoneAPIResponseSingleIDJSON `json:"-"`
}

// zoneAPIResponseSingleIDJSON contains the JSON metadata for the struct
// [ZoneAPIResponseSingleID]
type zoneAPIResponseSingleIDJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIResponseSingleID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIResponseSingleIDError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    zoneAPIResponseSingleIDErrorJSON `json:"-"`
}

// zoneAPIResponseSingleIDErrorJSON contains the JSON metadata for the struct
// [ZoneAPIResponseSingleIDError]
type zoneAPIResponseSingleIDErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIResponseSingleIDError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIResponseSingleIDMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    zoneAPIResponseSingleIDMessageJSON `json:"-"`
}

// zoneAPIResponseSingleIDMessageJSON contains the JSON metadata for the struct
// [ZoneAPIResponseSingleIDMessage]
type zoneAPIResponseSingleIDMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIResponseSingleIDMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIResponseSingleIDResult struct {
	// Identifier
	ID   string                            `json:"id,required"`
	JSON zoneAPIResponseSingleIDResultJSON `json:"-"`
}

// zoneAPIResponseSingleIDResultJSON contains the JSON metadata for the struct
// [ZoneAPIResponseSingleIDResult]
type zoneAPIResponseSingleIDResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIResponseSingleIDResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneNewResponse struct {
	Errors   []ZoneNewResponseError   `json:"errors"`
	Messages []ZoneNewResponseMessage `json:"messages"`
	Result   ZoneNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                `json:"success"`
	JSON    zoneNewResponseJSON `json:"-"`
}

// zoneNewResponseJSON contains the JSON metadata for the struct [ZoneNewResponse]
type zoneNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneNewResponseError struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    zoneNewResponseErrorJSON `json:"-"`
}

// zoneNewResponseErrorJSON contains the JSON metadata for the struct
// [ZoneNewResponseError]
type zoneNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneNewResponseMessage struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    zoneNewResponseMessageJSON `json:"-"`
}

// zoneNewResponseMessageJSON contains the JSON metadata for the struct
// [ZoneNewResponseMessage]
type zoneNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneNewResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// The last time proof of ownership was detected and the zone was made active
	ActivatedOn time.Time `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode float64 `json:"development_mode,required"`
	// When the zone was last modified
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The domain name
	Name string `json:"name,required"`
	// DNS host at the time of switching to Cloudflare
	OriginalDnshost string `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare Notes: Is this only available
	// for full zones?
	OriginalNameServers []string `json:"original_name_servers,required,nullable" format:"hostname"`
	// Registrar for the domain at the time of switching to Cloudflare
	OriginalRegistrar string                    `json:"original_registrar,required,nullable"`
	JSON              zoneNewResponseResultJSON `json:"-"`
}

// zoneNewResponseResultJSON contains the JSON metadata for the struct
// [ZoneNewResponseResult]
type zoneNewResponseResultJSON struct {
	ID                  apijson.Field
	ActivatedOn         apijson.Field
	CreatedOn           apijson.Field
	DevelopmentMode     apijson.Field
	ModifiedOn          apijson.Field
	Name                apijson.Field
	OriginalDnshost     apijson.Field
	OriginalNameServers apijson.Field
	OriginalRegistrar   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGetResponse struct {
	Errors   []ZoneGetResponseError   `json:"errors"`
	Messages []ZoneGetResponseMessage `json:"messages"`
	Result   ZoneGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                `json:"success"`
	JSON    zoneGetResponseJSON `json:"-"`
}

// zoneGetResponseJSON contains the JSON metadata for the struct [ZoneGetResponse]
type zoneGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGetResponseError struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    zoneGetResponseErrorJSON `json:"-"`
}

// zoneGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneGetResponseError]
type zoneGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGetResponseMessage struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    zoneGetResponseMessageJSON `json:"-"`
}

// zoneGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneGetResponseMessage]
type zoneGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGetResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// The last time proof of ownership was detected and the zone was made active
	ActivatedOn time.Time `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode float64 `json:"development_mode,required"`
	// When the zone was last modified
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The domain name
	Name string `json:"name,required"`
	// DNS host at the time of switching to Cloudflare
	OriginalDnshost string `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare Notes: Is this only available
	// for full zones?
	OriginalNameServers []string `json:"original_name_servers,required,nullable" format:"hostname"`
	// Registrar for the domain at the time of switching to Cloudflare
	OriginalRegistrar string                    `json:"original_registrar,required,nullable"`
	JSON              zoneGetResponseResultJSON `json:"-"`
}

// zoneGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneGetResponseResult]
type zoneGetResponseResultJSON struct {
	ID                  apijson.Field
	ActivatedOn         apijson.Field
	CreatedOn           apijson.Field
	DevelopmentMode     apijson.Field
	ModifiedOn          apijson.Field
	Name                apijson.Field
	OriginalDnshost     apijson.Field
	OriginalNameServers apijson.Field
	OriginalRegistrar   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneUpdateResponse struct {
	Errors   []ZoneUpdateResponseError   `json:"errors"`
	Messages []ZoneUpdateResponseMessage `json:"messages"`
	Result   ZoneUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                   `json:"success"`
	JSON    zoneUpdateResponseJSON `json:"-"`
}

// zoneUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneUpdateResponse]
type zoneUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneUpdateResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    zoneUpdateResponseErrorJSON `json:"-"`
}

// zoneUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneUpdateResponseError]
type zoneUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneUpdateResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    zoneUpdateResponseMessageJSON `json:"-"`
}

// zoneUpdateResponseMessageJSON contains the JSON metadata for the struct
// [ZoneUpdateResponseMessage]
type zoneUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneUpdateResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// The last time proof of ownership was detected and the zone was made active
	ActivatedOn time.Time `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode float64 `json:"development_mode,required"`
	// When the zone was last modified
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The domain name
	Name string `json:"name,required"`
	// DNS host at the time of switching to Cloudflare
	OriginalDnshost string `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare Notes: Is this only available
	// for full zones?
	OriginalNameServers []string `json:"original_name_servers,required,nullable" format:"hostname"`
	// Registrar for the domain at the time of switching to Cloudflare
	OriginalRegistrar string                       `json:"original_registrar,required,nullable"`
	JSON              zoneUpdateResponseResultJSON `json:"-"`
}

// zoneUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZoneUpdateResponseResult]
type zoneUpdateResponseResultJSON struct {
	ID                  apijson.Field
	ActivatedOn         apijson.Field
	CreatedOn           apijson.Field
	DevelopmentMode     apijson.Field
	ModifiedOn          apijson.Field
	Name                apijson.Field
	OriginalDnshost     apijson.Field
	OriginalNameServers apijson.Field
	OriginalRegistrar   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneListResponse struct {
	Errors     []ZoneListResponseError    `json:"errors"`
	Messages   []ZoneListResponseMessage  `json:"messages"`
	Result     []ZoneListResponseResult   `json:"result"`
	ResultInfo ZoneListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success bool                 `json:"success"`
	JSON    zoneListResponseJSON `json:"-"`
}

// zoneListResponseJSON contains the JSON metadata for the struct
// [ZoneListResponse]
type zoneListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneListResponseError struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    zoneListResponseErrorJSON `json:"-"`
}

// zoneListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneListResponseError]
type zoneListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneListResponseMessage struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    zoneListResponseMessageJSON `json:"-"`
}

// zoneListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneListResponseMessage]
type zoneListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneListResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// The last time proof of ownership was detected and the zone was made active
	ActivatedOn time.Time `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode float64 `json:"development_mode,required"`
	// When the zone was last modified
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The domain name
	Name string `json:"name,required"`
	// DNS host at the time of switching to Cloudflare
	OriginalDnshost string `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare Notes: Is this only available
	// for full zones?
	OriginalNameServers []string `json:"original_name_servers,required,nullable" format:"hostname"`
	// Registrar for the domain at the time of switching to Cloudflare
	OriginalRegistrar string                     `json:"original_registrar,required,nullable"`
	JSON              zoneListResponseResultJSON `json:"-"`
}

// zoneListResponseResultJSON contains the JSON metadata for the struct
// [ZoneListResponseResult]
type zoneListResponseResultJSON struct {
	ID                  apijson.Field
	ActivatedOn         apijson.Field
	CreatedOn           apijson.Field
	DevelopmentMode     apijson.Field
	ModifiedOn          apijson.Field
	Name                apijson.Field
	OriginalDnshost     apijson.Field
	OriginalNameServers apijson.Field
	OriginalRegistrar   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                        `json:"total_count"`
	JSON       zoneListResponseResultInfoJSON `json:"-"`
}

// zoneListResponseResultInfoJSON contains the JSON metadata for the struct
// [ZoneListResponseResultInfo]
type zoneListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneNewParams struct {
	Account param.Field[ZoneNewParamsAccount] `json:"account,required"`
	// The domain name
	Name param.Field[string] `json:"name,required"`
	// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
	// typically a partner-hosted zone or a CNAME setup.
	Type param.Field[ZoneNewParamsType] `json:"type"`
}

func (r ZoneNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneNewParamsAccount struct {
	// Identifier
	ID param.Field[string] `json:"id"`
}

func (r ZoneNewParamsAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
// typically a partner-hosted zone or a CNAME setup.
type ZoneNewParamsType string

const (
	ZoneNewParamsTypeFull    ZoneNewParamsType = "full"
	ZoneNewParamsTypePartial ZoneNewParamsType = "partial"
)

type ZoneUpdateParams struct {
	// (Deprecated) Please use the `/zones/{identifier}/subscription` API to update a
	// zone's plan. Changing this value will create/cancel associated subscriptions. To
	// view available plans for this zone, see Zone Plans.
	Plan param.Field[ZoneUpdateParamsPlan] `json:"plan"`
	// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
	// typically a partner-hosted zone or a CNAME setup. This parameter is only
	// available to Enterprise customers or if it has been explicitly enabled on a
	// zone.
	Type param.Field[ZoneUpdateParamsType] `json:"type"`
	// An array of domains used for custom name servers. This is only available for
	// Business and Enterprise plans.
	VanityNameServers param.Field[[]string] `json:"vanity_name_servers" format:"hostname"`
}

func (r ZoneUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// (Deprecated) Please use the `/zones/{identifier}/subscription` API to update a
// zone's plan. Changing this value will create/cancel associated subscriptions. To
// view available plans for this zone, see Zone Plans.
type ZoneUpdateParamsPlan struct {
	// Identifier
	ID param.Field[string] `json:"id"`
}

func (r ZoneUpdateParamsPlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
// typically a partner-hosted zone or a CNAME setup. This parameter is only
// available to Enterprise customers or if it has been explicitly enabled on a
// zone.
type ZoneUpdateParamsType string

const (
	ZoneUpdateParamsTypeFull    ZoneUpdateParamsType = "full"
	ZoneUpdateParamsTypePartial ZoneUpdateParamsType = "partial"
)

type ZoneListParams struct {
	Account param.Field[ZoneListParamsAccount] `query:"account"`
	// Direction to order zones.
	Direction param.Field[ZoneListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any).
	Match param.Field[ZoneListParamsMatch] `query:"match"`
	// A domain name.
	Name param.Field[string] `query:"name" format:"hostname"`
	// Field to order zones by.
	Order param.Field[ZoneListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of zones per page.
	PerPage param.Field[float64] `query:"per_page"`
	Status  param.Field[string]  `query:"status"`
}

// URLQuery serializes [ZoneListParams]'s query parameters as `url.Values`.
func (r ZoneListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneListParamsAccount struct {
	ID   param.Field[string] `query:"id"`
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [ZoneListParamsAccount]'s query parameters as `url.Values`.
func (r ZoneListParamsAccount) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order zones.
type ZoneListParamsDirection string

const (
	ZoneListParamsDirectionAsc  ZoneListParamsDirection = "asc"
	ZoneListParamsDirectionDesc ZoneListParamsDirection = "desc"
)

// Whether to match all search requirements or at least one (any).
type ZoneListParamsMatch string

const (
	ZoneListParamsMatchAny ZoneListParamsMatch = "any"
	ZoneListParamsMatchAll ZoneListParamsMatch = "all"
)

// Field to order zones by.
type ZoneListParamsOrder string

const (
	ZoneListParamsOrderName        ZoneListParamsOrder = "name"
	ZoneListParamsOrderStatus      ZoneListParamsOrder = "status"
	ZoneListParamsOrderAccountID   ZoneListParamsOrder = "account.id"
	ZoneListParamsOrderAccountName ZoneListParamsOrder = "account.name"
)

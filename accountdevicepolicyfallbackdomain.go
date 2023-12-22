// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountDevicePolicyFallbackDomainService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountDevicePolicyFallbackDomainService] method instead.
type AccountDevicePolicyFallbackDomainService struct {
	Options []option.RequestOption
}

// NewAccountDevicePolicyFallbackDomainService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDevicePolicyFallbackDomainService(opts ...option.RequestOption) (r *AccountDevicePolicyFallbackDomainService) {
	r = &AccountDevicePolicyFallbackDomainService{}
	r.Options = opts
	return
}

// Get the list of domains to bypass Gateway for DNS resolution and instead use the
// specified server.
func (r *AccountDevicePolicyFallbackDomainService) DevicesGetLocalDomainFallbackList(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *FallbackDomainResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/fallback_domains", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the list of domains to bypass Gateway for DNS resolution and instead use the
// specified server.
func (r *AccountDevicePolicyFallbackDomainService) DevicesGetLocalDomainFallbackListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *FallbackDomainResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/fallback_domains", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set the list of domains to bypass Gateway for DNS resolution and instead use the
// specified server.
func (r *AccountDevicePolicyFallbackDomainService) DevicesSetLocalDomainFallbackList(ctx context.Context, identifier interface{}, body AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParams, opts ...option.RequestOption) (res *FallbackDomainResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/fallback_domains", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Set the list of domains to bypass Gateway for DNS resolution and instead use the
// specified server.
func (r *AccountDevicePolicyFallbackDomainService) DevicesSetLocalDomainFallbackListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, body AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParams, opts ...option.RequestOption) (res *FallbackDomainResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/fallback_domains", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type FallbackDomainResponseCollection struct {
	Errors     []FallbackDomainResponseCollectionError    `json:"errors"`
	Messages   []FallbackDomainResponseCollectionMessage  `json:"messages"`
	Result     []FallbackDomainResponseCollectionResult   `json:"result"`
	ResultInfo FallbackDomainResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success FallbackDomainResponseCollectionSuccess `json:"success"`
	JSON    fallbackDomainResponseCollectionJSON    `json:"-"`
}

// fallbackDomainResponseCollectionJSON contains the JSON metadata for the struct
// [FallbackDomainResponseCollection]
type fallbackDomainResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackDomainResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FallbackDomainResponseCollectionError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    fallbackDomainResponseCollectionErrorJSON `json:"-"`
}

// fallbackDomainResponseCollectionErrorJSON contains the JSON metadata for the
// struct [FallbackDomainResponseCollectionError]
type fallbackDomainResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackDomainResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FallbackDomainResponseCollectionMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    fallbackDomainResponseCollectionMessageJSON `json:"-"`
}

// fallbackDomainResponseCollectionMessageJSON contains the JSON metadata for the
// struct [FallbackDomainResponseCollectionMessage]
type fallbackDomainResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackDomainResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FallbackDomainResponseCollectionResult struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                              `json:"dns_server"`
	JSON      fallbackDomainResponseCollectionResultJSON `json:"-"`
}

// fallbackDomainResponseCollectionResultJSON contains the JSON metadata for the
// struct [FallbackDomainResponseCollectionResult]
type fallbackDomainResponseCollectionResultJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackDomainResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FallbackDomainResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       fallbackDomainResponseCollectionResultInfoJSON `json:"-"`
}

// fallbackDomainResponseCollectionResultInfoJSON contains the JSON metadata for
// the struct [FallbackDomainResponseCollectionResultInfo]
type fallbackDomainResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackDomainResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FallbackDomainResponseCollectionSuccess bool

const (
	FallbackDomainResponseCollectionSuccessTrue FallbackDomainResponseCollectionSuccess = true
)

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParams struct {
	Body param.Field[[]AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParamsBody] `json:"body,required"`
}

func (r AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParamsBody struct {
	// The domain suffix to match when resolving locally.
	Suffix param.Field[string] `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer param.Field[[]interface{}] `json:"dns_server"`
}

func (r AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParams struct {
	Body param.Field[[]AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParamsBody] `json:"body,required"`
}

func (r AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParamsBody struct {
	// The domain suffix to match when resolving locally.
	Suffix param.Field[string] `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer param.Field[[]interface{}] `json:"dns_server"`
}

func (r AccountDevicePolicyFallbackDomainDevicesSetLocalDomainFallbackListForADeviceSettingsPolicyParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

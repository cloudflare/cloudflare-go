// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// HostnameSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewHostnameSettingService] method
// instead.
type HostnameSettingService struct {
	Options []option.RequestOption
}

// NewHostnameSettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHostnameSettingService(opts ...option.RequestOption) (r *HostnameSettingService) {
	r = &HostnameSettingService{}
	r.Options = opts
	return
}

// List the requested TLS setting for the hostnames under this zone.
func (r *HostnameSettingService) List(ctx context.Context, zoneIdentifier string, tlsSetting HostnameSettingListParamsTlsSetting, opts ...option.RequestOption) (res *HostnameSettingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v", zoneIdentifier, tlsSetting)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the tls setting value for the hostname.
func (r *HostnameSettingService) Delete(ctx context.Context, zoneIdentifier string, tlsSetting HostnameSettingDeleteParamsTlsSetting, hostname string, opts ...option.RequestOption) (res *HostnameSettingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v/%s", zoneIdentifier, tlsSetting, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Update the tls setting value for the hostname.
func (r *HostnameSettingService) Add(ctx context.Context, zoneIdentifier string, tlsSetting HostnameSettingAddParamsTlsSetting, hostname string, body HostnameSettingAddParams, opts ...option.RequestOption) (res *HostnameSettingAddResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v/%s", zoneIdentifier, tlsSetting, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type HostnameSettingListResponse struct {
	Errors     []HostnameSettingListResponseError    `json:"errors"`
	Messages   []HostnameSettingListResponseMessage  `json:"messages"`
	Result     []HostnameSettingListResponseResult   `json:"result"`
	ResultInfo HostnameSettingListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success HostnameSettingListResponseSuccess `json:"success"`
	JSON    hostnameSettingListResponseJSON    `json:"-"`
}

// hostnameSettingListResponseJSON contains the JSON metadata for the struct
// [HostnameSettingListResponse]
type hostnameSettingListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingListResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    hostnameSettingListResponseErrorJSON `json:"-"`
}

// hostnameSettingListResponseErrorJSON contains the JSON metadata for the struct
// [HostnameSettingListResponseError]
type hostnameSettingListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingListResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    hostnameSettingListResponseMessageJSON `json:"-"`
}

// hostnameSettingListResponseMessageJSON contains the JSON metadata for the struct
// [HostnameSettingListResponseMessage]
type hostnameSettingListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingListResponseResult struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value HostnameSettingListResponseResultValue `json:"value"`
	JSON  hostnameSettingListResponseResultJSON  `json:"-"`
}

// hostnameSettingListResponseResultJSON contains the JSON metadata for the struct
// [HostnameSettingListResponseResult]
type hostnameSettingListResponseResultJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The tls setting value.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionString] or
// [HostnameSettingListResponseResultValueArray].
type HostnameSettingListResponseResultValue interface {
	ImplementsHostnameSettingListResponseResultValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HostnameSettingListResponseResultValue)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type HostnameSettingListResponseResultValueArray []string

func (r HostnameSettingListResponseResultValueArray) ImplementsHostnameSettingListResponseResultValue() {
}

type HostnameSettingListResponseResultInfo struct {
	Count      interface{} `json:"count"`
	Page       interface{} `json:"page"`
	PerPage    interface{} `json:"per_page"`
	TotalCount interface{} `json:"total_count"`
	// Total pages available of results
	TotalPages float64                                   `json:"total_pages"`
	JSON       hostnameSettingListResponseResultInfoJSON `json:"-"`
}

// hostnameSettingListResponseResultInfoJSON contains the JSON metadata for the
// struct [HostnameSettingListResponseResultInfo]
type hostnameSettingListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HostnameSettingListResponseSuccess bool

const (
	HostnameSettingListResponseSuccessTrue HostnameSettingListResponseSuccess = true
)

type HostnameSettingDeleteResponse struct {
	Errors   []HostnameSettingDeleteResponseError   `json:"errors"`
	Messages []HostnameSettingDeleteResponseMessage `json:"messages"`
	Result   HostnameSettingDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success HostnameSettingDeleteResponseSuccess `json:"success"`
	JSON    hostnameSettingDeleteResponseJSON    `json:"-"`
}

// hostnameSettingDeleteResponseJSON contains the JSON metadata for the struct
// [HostnameSettingDeleteResponse]
type hostnameSettingDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingDeleteResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    hostnameSettingDeleteResponseErrorJSON `json:"-"`
}

// hostnameSettingDeleteResponseErrorJSON contains the JSON metadata for the struct
// [HostnameSettingDeleteResponseError]
type hostnameSettingDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingDeleteResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    hostnameSettingDeleteResponseMessageJSON `json:"-"`
}

// hostnameSettingDeleteResponseMessageJSON contains the JSON metadata for the
// struct [HostnameSettingDeleteResponseMessage]
type hostnameSettingDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingDeleteResponseResult struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string      `json:"hostname"`
	Status   interface{} `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time                               `json:"updated_at" format:"date-time"`
	Value     interface{}                             `json:"value"`
	JSON      hostnameSettingDeleteResponseResultJSON `json:"-"`
}

// hostnameSettingDeleteResponseResultJSON contains the JSON metadata for the
// struct [HostnameSettingDeleteResponseResult]
type hostnameSettingDeleteResponseResultJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HostnameSettingDeleteResponseSuccess bool

const (
	HostnameSettingDeleteResponseSuccessTrue HostnameSettingDeleteResponseSuccess = true
)

type HostnameSettingAddResponse struct {
	Errors   []HostnameSettingAddResponseError   `json:"errors"`
	Messages []HostnameSettingAddResponseMessage `json:"messages"`
	Result   HostnameSettingAddResponseResult    `json:"result"`
	// Whether the API call was successful
	Success HostnameSettingAddResponseSuccess `json:"success"`
	JSON    hostnameSettingAddResponseJSON    `json:"-"`
}

// hostnameSettingAddResponseJSON contains the JSON metadata for the struct
// [HostnameSettingAddResponse]
type hostnameSettingAddResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingAddResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingAddResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    hostnameSettingAddResponseErrorJSON `json:"-"`
}

// hostnameSettingAddResponseErrorJSON contains the JSON metadata for the struct
// [HostnameSettingAddResponseError]
type hostnameSettingAddResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingAddResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingAddResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    hostnameSettingAddResponseMessageJSON `json:"-"`
}

// hostnameSettingAddResponseMessageJSON contains the JSON metadata for the struct
// [HostnameSettingAddResponseMessage]
type hostnameSettingAddResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingAddResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingAddResponseResult struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value HostnameSettingAddResponseResultValue `json:"value"`
	JSON  hostnameSettingAddResponseResultJSON  `json:"-"`
}

// hostnameSettingAddResponseResultJSON contains the JSON metadata for the struct
// [HostnameSettingAddResponseResult]
type hostnameSettingAddResponseResultJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingAddResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The tls setting value.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionString] or
// [HostnameSettingAddResponseResultValueArray].
type HostnameSettingAddResponseResultValue interface {
	ImplementsHostnameSettingAddResponseResultValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HostnameSettingAddResponseResultValue)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type HostnameSettingAddResponseResultValueArray []string

func (r HostnameSettingAddResponseResultValueArray) ImplementsHostnameSettingAddResponseResultValue() {
}

// Whether the API call was successful
type HostnameSettingAddResponseSuccess bool

const (
	HostnameSettingAddResponseSuccessTrue HostnameSettingAddResponseSuccess = true
)

// The TLS Setting name.
type HostnameSettingListParamsTlsSetting string

const (
	HostnameSettingListParamsTlsSettingCiphers       HostnameSettingListParamsTlsSetting = "ciphers"
	HostnameSettingListParamsTlsSettingMinTlsVersion HostnameSettingListParamsTlsSetting = "min_tls_version"
	HostnameSettingListParamsTlsSettingHttp2         HostnameSettingListParamsTlsSetting = "http2"
)

// The TLS Setting name.
type HostnameSettingDeleteParamsTlsSetting string

const (
	HostnameSettingDeleteParamsTlsSettingCiphers       HostnameSettingDeleteParamsTlsSetting = "ciphers"
	HostnameSettingDeleteParamsTlsSettingMinTlsVersion HostnameSettingDeleteParamsTlsSetting = "min_tls_version"
	HostnameSettingDeleteParamsTlsSettingHttp2         HostnameSettingDeleteParamsTlsSetting = "http2"
)

type HostnameSettingAddParams struct {
	// The tls setting value.
	Value param.Field[HostnameSettingAddParamsValue] `json:"value,required"`
}

func (r HostnameSettingAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TLS Setting name.
type HostnameSettingAddParamsTlsSetting string

const (
	HostnameSettingAddParamsTlsSettingCiphers       HostnameSettingAddParamsTlsSetting = "ciphers"
	HostnameSettingAddParamsTlsSettingMinTlsVersion HostnameSettingAddParamsTlsSetting = "min_tls_version"
	HostnameSettingAddParamsTlsSettingHttp2         HostnameSettingAddParamsTlsSetting = "http2"
)

// The tls setting value.
//
// Satisfied by [shared.UnionFloat], [shared.UnionString],
// [HostnameSettingAddParamsValueArray].
type HostnameSettingAddParamsValue interface {
	ImplementsHostnameSettingAddParamsValue()
}

type HostnameSettingAddParamsValueArray []string

func (r HostnameSettingAddParamsValueArray) ImplementsHostnameSettingAddParamsValue() {}

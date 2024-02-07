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

// HostnameSettingTLSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewHostnameSettingTLSService] method
// instead.
type HostnameSettingTLSService struct {
	Options []option.RequestOption
}

// NewHostnameSettingTLSService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHostnameSettingTLSService(opts ...option.RequestOption) (r *HostnameSettingTLSService) {
	r = &HostnameSettingTLSService{}
	r.Options = opts
	return
}

// List the requested TLS setting for the hostnames under this zone.
func (r *HostnameSettingTLSService) Get(ctx context.Context, zoneIdentifier string, tlsSetting HostnameSettingTLSGetParamsTLSSetting, opts ...option.RequestOption) (res *HostnameSettingTLSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v", zoneIdentifier, tlsSetting)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the tls setting value for the hostname.
func (r *HostnameSettingTLSService) Update(ctx context.Context, zoneIdentifier string, tlsSetting HostnameSettingTLSUpdateParamsTLSSetting, hostname string, body HostnameSettingTLSUpdateParams, opts ...option.RequestOption) (res *HostnameSettingTLSUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v/%s", zoneIdentifier, tlsSetting, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete the tls setting value for the hostname.
func (r *HostnameSettingTLSService) Delete(ctx context.Context, zoneIdentifier string, tlsSetting HostnameSettingTLSDeleteParamsTLSSetting, hostname string, opts ...option.RequestOption) (res *HostnameSettingTLSDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v/%s", zoneIdentifier, tlsSetting, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type HostnameSettingTLSGetResponse struct {
	Errors     []HostnameSettingTLSGetResponseError    `json:"errors"`
	Messages   []HostnameSettingTLSGetResponseMessage  `json:"messages"`
	Result     []HostnameSettingTLSGetResponseResult   `json:"result"`
	ResultInfo HostnameSettingTLSGetResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success HostnameSettingTLSGetResponseSuccess `json:"success"`
	JSON    hostnameSettingTLSGetResponseJSON    `json:"-"`
}

// hostnameSettingTLSGetResponseJSON contains the JSON metadata for the struct
// [HostnameSettingTLSGetResponse]
type hostnameSettingTLSGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingTLSGetResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    hostnameSettingTLSGetResponseErrorJSON `json:"-"`
}

// hostnameSettingTLSGetResponseErrorJSON contains the JSON metadata for the struct
// [HostnameSettingTLSGetResponseError]
type hostnameSettingTLSGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingTLSGetResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    hostnameSettingTLSGetResponseMessageJSON `json:"-"`
}

// hostnameSettingTLSGetResponseMessageJSON contains the JSON metadata for the
// struct [HostnameSettingTLSGetResponseMessage]
type hostnameSettingTLSGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingTLSGetResponseResult struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value HostnameSettingTLSGetResponseResultValue `json:"value"`
	JSON  hostnameSettingTLSGetResponseResultJSON  `json:"-"`
}

// hostnameSettingTLSGetResponseResultJSON contains the JSON metadata for the
// struct [HostnameSettingTLSGetResponseResult]
type hostnameSettingTLSGetResponseResultJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The tls setting value.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionString] or
// [HostnameSettingTLSGetResponseResultValueArray].
type HostnameSettingTLSGetResponseResultValue interface {
	ImplementsHostnameSettingTLSGetResponseResultValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HostnameSettingTLSGetResponseResultValue)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type HostnameSettingTLSGetResponseResultValueArray []string

func (r HostnameSettingTLSGetResponseResultValueArray) ImplementsHostnameSettingTLSGetResponseResultValue() {
}

type HostnameSettingTLSGetResponseResultInfo struct {
	Count      interface{} `json:"count"`
	Page       interface{} `json:"page"`
	PerPage    interface{} `json:"per_page"`
	TotalCount interface{} `json:"total_count"`
	// Total pages available of results
	TotalPages float64                                     `json:"total_pages"`
	JSON       hostnameSettingTLSGetResponseResultInfoJSON `json:"-"`
}

// hostnameSettingTLSGetResponseResultInfoJSON contains the JSON metadata for the
// struct [HostnameSettingTLSGetResponseResultInfo]
type hostnameSettingTLSGetResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSGetResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HostnameSettingTLSGetResponseSuccess bool

const (
	HostnameSettingTLSGetResponseSuccessTrue HostnameSettingTLSGetResponseSuccess = true
)

type HostnameSettingTLSUpdateResponse struct {
	Errors   []HostnameSettingTLSUpdateResponseError   `json:"errors"`
	Messages []HostnameSettingTLSUpdateResponseMessage `json:"messages"`
	Result   HostnameSettingTLSUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success HostnameSettingTLSUpdateResponseSuccess `json:"success"`
	JSON    hostnameSettingTLSUpdateResponseJSON    `json:"-"`
}

// hostnameSettingTLSUpdateResponseJSON contains the JSON metadata for the struct
// [HostnameSettingTLSUpdateResponse]
type hostnameSettingTLSUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingTLSUpdateResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    hostnameSettingTLSUpdateResponseErrorJSON `json:"-"`
}

// hostnameSettingTLSUpdateResponseErrorJSON contains the JSON metadata for the
// struct [HostnameSettingTLSUpdateResponseError]
type hostnameSettingTLSUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingTLSUpdateResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    hostnameSettingTLSUpdateResponseMessageJSON `json:"-"`
}

// hostnameSettingTLSUpdateResponseMessageJSON contains the JSON metadata for the
// struct [HostnameSettingTLSUpdateResponseMessage]
type hostnameSettingTLSUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingTLSUpdateResponseResult struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value HostnameSettingTLSUpdateResponseResultValue `json:"value"`
	JSON  hostnameSettingTLSUpdateResponseResultJSON  `json:"-"`
}

// hostnameSettingTLSUpdateResponseResultJSON contains the JSON metadata for the
// struct [HostnameSettingTLSUpdateResponseResult]
type hostnameSettingTLSUpdateResponseResultJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The tls setting value.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionString] or
// [HostnameSettingTLSUpdateResponseResultValueArray].
type HostnameSettingTLSUpdateResponseResultValue interface {
	ImplementsHostnameSettingTLSUpdateResponseResultValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HostnameSettingTLSUpdateResponseResultValue)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type HostnameSettingTLSUpdateResponseResultValueArray []string

func (r HostnameSettingTLSUpdateResponseResultValueArray) ImplementsHostnameSettingTLSUpdateResponseResultValue() {
}

// Whether the API call was successful
type HostnameSettingTLSUpdateResponseSuccess bool

const (
	HostnameSettingTLSUpdateResponseSuccessTrue HostnameSettingTLSUpdateResponseSuccess = true
)

type HostnameSettingTLSDeleteResponse struct {
	Errors   []HostnameSettingTLSDeleteResponseError   `json:"errors"`
	Messages []HostnameSettingTLSDeleteResponseMessage `json:"messages"`
	Result   HostnameSettingTLSDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success HostnameSettingTLSDeleteResponseSuccess `json:"success"`
	JSON    hostnameSettingTLSDeleteResponseJSON    `json:"-"`
}

// hostnameSettingTLSDeleteResponseJSON contains the JSON metadata for the struct
// [HostnameSettingTLSDeleteResponse]
type hostnameSettingTLSDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingTLSDeleteResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    hostnameSettingTLSDeleteResponseErrorJSON `json:"-"`
}

// hostnameSettingTLSDeleteResponseErrorJSON contains the JSON metadata for the
// struct [HostnameSettingTLSDeleteResponseError]
type hostnameSettingTLSDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingTLSDeleteResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    hostnameSettingTLSDeleteResponseMessageJSON `json:"-"`
}

// hostnameSettingTLSDeleteResponseMessageJSON contains the JSON metadata for the
// struct [HostnameSettingTLSDeleteResponseMessage]
type hostnameSettingTLSDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameSettingTLSDeleteResponseResult struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string      `json:"hostname"`
	Status   interface{} `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	Value     interface{}                                `json:"value"`
	JSON      hostnameSettingTLSDeleteResponseResultJSON `json:"-"`
}

// hostnameSettingTLSDeleteResponseResultJSON contains the JSON metadata for the
// struct [HostnameSettingTLSDeleteResponseResult]
type hostnameSettingTLSDeleteResponseResultJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HostnameSettingTLSDeleteResponseSuccess bool

const (
	HostnameSettingTLSDeleteResponseSuccessTrue HostnameSettingTLSDeleteResponseSuccess = true
)

// The TLS Setting name.
type HostnameSettingTLSGetParamsTLSSetting string

const (
	HostnameSettingTLSGetParamsTLSSettingCiphers       HostnameSettingTLSGetParamsTLSSetting = "ciphers"
	HostnameSettingTLSGetParamsTLSSettingMinTLSVersion HostnameSettingTLSGetParamsTLSSetting = "min_tls_version"
	HostnameSettingTLSGetParamsTLSSettingHTTP2         HostnameSettingTLSGetParamsTLSSetting = "http2"
)

type HostnameSettingTLSUpdateParams struct {
	// The tls setting value.
	Value param.Field[HostnameSettingTLSUpdateParamsValue] `json:"value,required"`
}

func (r HostnameSettingTLSUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TLS Setting name.
type HostnameSettingTLSUpdateParamsTLSSetting string

const (
	HostnameSettingTLSUpdateParamsTLSSettingCiphers       HostnameSettingTLSUpdateParamsTLSSetting = "ciphers"
	HostnameSettingTLSUpdateParamsTLSSettingMinTLSVersion HostnameSettingTLSUpdateParamsTLSSetting = "min_tls_version"
	HostnameSettingTLSUpdateParamsTLSSettingHTTP2         HostnameSettingTLSUpdateParamsTLSSetting = "http2"
)

// The tls setting value.
//
// Satisfied by [shared.UnionFloat], [shared.UnionString],
// [HostnameSettingTLSUpdateParamsValueArray].
type HostnameSettingTLSUpdateParamsValue interface {
	ImplementsHostnameSettingTLSUpdateParamsValue()
}

type HostnameSettingTLSUpdateParamsValueArray []string

func (r HostnameSettingTLSUpdateParamsValueArray) ImplementsHostnameSettingTLSUpdateParamsValue() {}

// The TLS Setting name.
type HostnameSettingTLSDeleteParamsTLSSetting string

const (
	HostnameSettingTLSDeleteParamsTLSSettingCiphers       HostnameSettingTLSDeleteParamsTLSSetting = "ciphers"
	HostnameSettingTLSDeleteParamsTLSSettingMinTLSVersion HostnameSettingTLSDeleteParamsTLSSetting = "min_tls_version"
	HostnameSettingTLSDeleteParamsTLSSettingHTTP2         HostnameSettingTLSDeleteParamsTLSSetting = "http2"
)

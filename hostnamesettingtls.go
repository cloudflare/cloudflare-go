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

// Update the tls setting value for the hostname.
func (r *HostnameSettingTLSService) Update(ctx context.Context, settingID HostnameSettingTLSUpdateParamsSettingID, hostname string, params HostnameSettingTLSUpdateParams, opts ...option.RequestOption) (res *HostnameSettingTLSUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameSettingTLSUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v/%s", params.ZoneID, settingID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete the tls setting value for the hostname.
func (r *HostnameSettingTLSService) Delete(ctx context.Context, settingID HostnameSettingTLSDeleteParamsSettingID, hostname string, body HostnameSettingTLSDeleteParams, opts ...option.RequestOption) (res *HostnameSettingTLSDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameSettingTLSDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v/%s", body.ZoneID, settingID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List the requested TLS setting for the hostnames under this zone.
func (r *HostnameSettingTLSService) Get(ctx context.Context, settingID HostnameSettingTLSGetParamsSettingID, query HostnameSettingTLSGetParams, opts ...option.RequestOption) (res *[]HostnameSettingTLSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameSettingTLSGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v", query.ZoneID, settingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HostnameSettingTLSUpdateResponse struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value HostnameSettingTLSUpdateResponseValue `json:"value"`
	JSON  hostnameSettingTLSUpdateResponseJSON  `json:"-"`
}

// hostnameSettingTLSUpdateResponseJSON contains the JSON metadata for the struct
// [HostnameSettingTLSUpdateResponse]
type hostnameSettingTLSUpdateResponseJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameSettingTLSUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The tls setting value.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionString] or
// [HostnameSettingTLSUpdateResponseValueArray].
type HostnameSettingTLSUpdateResponseValue interface {
	ImplementsHostnameSettingTLSUpdateResponseValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HostnameSettingTLSUpdateResponseValue)(nil)).Elem(),
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

type HostnameSettingTLSUpdateResponseValueArray []string

func (r HostnameSettingTLSUpdateResponseValueArray) ImplementsHostnameSettingTLSUpdateResponseValue() {
}

type HostnameSettingTLSDeleteResponse struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string      `json:"hostname"`
	Status   interface{} `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	Value     interface{}                          `json:"value"`
	JSON      hostnameSettingTLSDeleteResponseJSON `json:"-"`
}

// hostnameSettingTLSDeleteResponseJSON contains the JSON metadata for the struct
// [HostnameSettingTLSDeleteResponse]
type hostnameSettingTLSDeleteResponseJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameSettingTLSDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type HostnameSettingTLSGetResponse struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value HostnameSettingTLSGetResponseValue `json:"value"`
	JSON  hostnameSettingTLSGetResponseJSON  `json:"-"`
}

// hostnameSettingTLSGetResponseJSON contains the JSON metadata for the struct
// [HostnameSettingTLSGetResponse]
type hostnameSettingTLSGetResponseJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameSettingTLSGetResponseJSON) RawJSON() string {
	return r.raw
}

// The tls setting value.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionString] or
// [HostnameSettingTLSGetResponseValueArray].
type HostnameSettingTLSGetResponseValue interface {
	ImplementsHostnameSettingTLSGetResponseValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HostnameSettingTLSGetResponseValue)(nil)).Elem(),
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

type HostnameSettingTLSGetResponseValueArray []string

func (r HostnameSettingTLSGetResponseValueArray) ImplementsHostnameSettingTLSGetResponseValue() {}

type HostnameSettingTLSUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The tls setting value.
	Value param.Field[HostnameSettingTLSUpdateParamsValue] `json:"value,required"`
}

func (r HostnameSettingTLSUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TLS Setting name.
type HostnameSettingTLSUpdateParamsSettingID string

const (
	HostnameSettingTLSUpdateParamsSettingIDCiphers       HostnameSettingTLSUpdateParamsSettingID = "ciphers"
	HostnameSettingTLSUpdateParamsSettingIDMinTLSVersion HostnameSettingTLSUpdateParamsSettingID = "min_tls_version"
	HostnameSettingTLSUpdateParamsSettingIDHTTP2         HostnameSettingTLSUpdateParamsSettingID = "http2"
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

type HostnameSettingTLSUpdateResponseEnvelope struct {
	Errors   []HostnameSettingTLSUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameSettingTLSUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   HostnameSettingTLSUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HostnameSettingTLSUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameSettingTLSUpdateResponseEnvelopeJSON    `json:"-"`
}

// hostnameSettingTLSUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [HostnameSettingTLSUpdateResponseEnvelope]
type hostnameSettingTLSUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameSettingTLSUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameSettingTLSUpdateResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    hostnameSettingTLSUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameSettingTLSUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [HostnameSettingTLSUpdateResponseEnvelopeErrors]
type hostnameSettingTLSUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameSettingTLSUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameSettingTLSUpdateResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    hostnameSettingTLSUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameSettingTLSUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [HostnameSettingTLSUpdateResponseEnvelopeMessages]
type hostnameSettingTLSUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameSettingTLSUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameSettingTLSUpdateResponseEnvelopeSuccess bool

const (
	HostnameSettingTLSUpdateResponseEnvelopeSuccessTrue HostnameSettingTLSUpdateResponseEnvelopeSuccess = true
)

type HostnameSettingTLSDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

// The TLS Setting name.
type HostnameSettingTLSDeleteParamsSettingID string

const (
	HostnameSettingTLSDeleteParamsSettingIDCiphers       HostnameSettingTLSDeleteParamsSettingID = "ciphers"
	HostnameSettingTLSDeleteParamsSettingIDMinTLSVersion HostnameSettingTLSDeleteParamsSettingID = "min_tls_version"
	HostnameSettingTLSDeleteParamsSettingIDHTTP2         HostnameSettingTLSDeleteParamsSettingID = "http2"
)

type HostnameSettingTLSDeleteResponseEnvelope struct {
	Errors   []HostnameSettingTLSDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameSettingTLSDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   HostnameSettingTLSDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HostnameSettingTLSDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameSettingTLSDeleteResponseEnvelopeJSON    `json:"-"`
}

// hostnameSettingTLSDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [HostnameSettingTLSDeleteResponseEnvelope]
type hostnameSettingTLSDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameSettingTLSDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameSettingTLSDeleteResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    hostnameSettingTLSDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameSettingTLSDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [HostnameSettingTLSDeleteResponseEnvelopeErrors]
type hostnameSettingTLSDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameSettingTLSDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameSettingTLSDeleteResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    hostnameSettingTLSDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameSettingTLSDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [HostnameSettingTLSDeleteResponseEnvelopeMessages]
type hostnameSettingTLSDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameSettingTLSDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameSettingTLSDeleteResponseEnvelopeSuccess bool

const (
	HostnameSettingTLSDeleteResponseEnvelopeSuccessTrue HostnameSettingTLSDeleteResponseEnvelopeSuccess = true
)

type HostnameSettingTLSGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

// The TLS Setting name.
type HostnameSettingTLSGetParamsSettingID string

const (
	HostnameSettingTLSGetParamsSettingIDCiphers       HostnameSettingTLSGetParamsSettingID = "ciphers"
	HostnameSettingTLSGetParamsSettingIDMinTLSVersion HostnameSettingTLSGetParamsSettingID = "min_tls_version"
	HostnameSettingTLSGetParamsSettingIDHTTP2         HostnameSettingTLSGetParamsSettingID = "http2"
)

type HostnameSettingTLSGetResponseEnvelope struct {
	Errors   []HostnameSettingTLSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameSettingTLSGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []HostnameSettingTLSGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    HostnameSettingTLSGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo HostnameSettingTLSGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       hostnameSettingTLSGetResponseEnvelopeJSON       `json:"-"`
}

// hostnameSettingTLSGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HostnameSettingTLSGetResponseEnvelope]
type hostnameSettingTLSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameSettingTLSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameSettingTLSGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    hostnameSettingTLSGetResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameSettingTLSGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [HostnameSettingTLSGetResponseEnvelopeErrors]
type hostnameSettingTLSGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameSettingTLSGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameSettingTLSGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    hostnameSettingTLSGetResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameSettingTLSGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [HostnameSettingTLSGetResponseEnvelopeMessages]
type hostnameSettingTLSGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameSettingTLSGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameSettingTLSGetResponseEnvelopeSuccess bool

const (
	HostnameSettingTLSGetResponseEnvelopeSuccessTrue HostnameSettingTLSGetResponseEnvelopeSuccess = true
)

type HostnameSettingTLSGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64 `json:"total_count"`
	// Total pages available of results
	TotalPages float64                                             `json:"total_pages"`
	JSON       hostnameSettingTLSGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// hostnameSettingTLSGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [HostnameSettingTLSGetResponseEnvelopeResultInfo]
type hostnameSettingTLSGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingTLSGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameSettingTLSGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

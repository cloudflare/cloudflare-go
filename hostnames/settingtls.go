// File generated from our OpenAPI spec by Stainless.

package hostnames

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// SettingTLSService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSettingTLSService] method instead.
type SettingTLSService struct {
	Options []option.RequestOption
}

// NewSettingTLSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingTLSService(opts ...option.RequestOption) (r *SettingTLSService) {
	r = &SettingTLSService{}
	r.Options = opts
	return
}

// Update the tls setting value for the hostname.
func (r *SettingTLSService) Update(ctx context.Context, settingID SettingTLSUpdateParamsSettingID, hostname string, params SettingTLSUpdateParams, opts ...option.RequestOption) (res *TLSCertificatesAndHostnamesSettingObject, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTLSUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v/%s", params.ZoneID, settingID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete the tls setting value for the hostname.
func (r *SettingTLSService) Delete(ctx context.Context, settingID SettingTLSDeleteParamsSettingID, hostname string, body SettingTLSDeleteParams, opts ...option.RequestOption) (res *TLSCertificatesAndHostnamesSettingObjectDelete, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTLSDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v/%s", body.ZoneID, settingID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List the requested TLS setting for the hostnames under this zone.
func (r *SettingTLSService) Get(ctx context.Context, settingID SettingTLSGetParamsSettingID, query SettingTLSGetParams, opts ...option.RequestOption) (res *[]SettingTLSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTLSGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v", query.ZoneID, settingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TLSCertificatesAndHostnamesSettingObject struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value TLSCertificatesAndHostnamesSettingObjectValue `json:"value"`
	JSON  tlsCertificatesAndHostnamesSettingObjectJSON  `json:"-"`
}

// tlsCertificatesAndHostnamesSettingObjectJSON contains the JSON metadata for the
// struct [TLSCertificatesAndHostnamesSettingObject]
type tlsCertificatesAndHostnamesSettingObjectJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TLSCertificatesAndHostnamesSettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsCertificatesAndHostnamesSettingObjectJSON) RawJSON() string {
	return r.raw
}

// The tls setting value.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionString] or
// [hostnames.TLSCertificatesAndHostnamesSettingObjectValueArray].
type TLSCertificatesAndHostnamesSettingObjectValue interface {
	ImplementsHostnamesTLSCertificatesAndHostnamesSettingObjectValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TLSCertificatesAndHostnamesSettingObjectValue)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TLSCertificatesAndHostnamesSettingObjectValueArray{}),
		},
	)
}

type TLSCertificatesAndHostnamesSettingObjectValueArray []string

func (r TLSCertificatesAndHostnamesSettingObjectValueArray) ImplementsHostnamesTLSCertificatesAndHostnamesSettingObjectValue() {
}

type TLSCertificatesAndHostnamesSettingObjectDelete struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string      `json:"hostname"`
	Status   interface{} `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time                                          `json:"updated_at" format:"date-time"`
	Value     interface{}                                        `json:"value"`
	JSON      tlsCertificatesAndHostnamesSettingObjectDeleteJSON `json:"-"`
}

// tlsCertificatesAndHostnamesSettingObjectDeleteJSON contains the JSON metadata
// for the struct [TLSCertificatesAndHostnamesSettingObjectDelete]
type tlsCertificatesAndHostnamesSettingObjectDeleteJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TLSCertificatesAndHostnamesSettingObjectDelete) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsCertificatesAndHostnamesSettingObjectDeleteJSON) RawJSON() string {
	return r.raw
}

type SettingTLSGetResponse struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value SettingTLSGetResponseValue `json:"value"`
	JSON  settingTLSGetResponseJSON  `json:"-"`
}

// settingTLSGetResponseJSON contains the JSON metadata for the struct
// [SettingTLSGetResponse]
type settingTLSGetResponseJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSGetResponseJSON) RawJSON() string {
	return r.raw
}

// The tls setting value.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionString] or
// [hostnames.SettingTLSGetResponseValueArray].
type SettingTLSGetResponseValue interface {
	ImplementsHostnamesSettingTLSGetResponseValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingTLSGetResponseValue)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingTLSGetResponseValueArray{}),
		},
	)
}

type SettingTLSGetResponseValueArray []string

func (r SettingTLSGetResponseValueArray) ImplementsHostnamesSettingTLSGetResponseValue() {}

type SettingTLSUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The tls setting value.
	Value param.Field[SettingTLSUpdateParamsValue] `json:"value,required"`
}

func (r SettingTLSUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TLS Setting name.
type SettingTLSUpdateParamsSettingID string

const (
	SettingTLSUpdateParamsSettingIDCiphers       SettingTLSUpdateParamsSettingID = "ciphers"
	SettingTLSUpdateParamsSettingIDMinTLSVersion SettingTLSUpdateParamsSettingID = "min_tls_version"
	SettingTLSUpdateParamsSettingIDHTTP2         SettingTLSUpdateParamsSettingID = "http2"
)

// The tls setting value.
//
// Satisfied by [shared.UnionFloat], [shared.UnionString],
// [hostnames.SettingTLSUpdateParamsValueArray].
type SettingTLSUpdateParamsValue interface {
	ImplementsHostnamesSettingTLSUpdateParamsValue()
}

type SettingTLSUpdateParamsValueArray []string

func (r SettingTLSUpdateParamsValueArray) ImplementsHostnamesSettingTLSUpdateParamsValue() {}

type SettingTLSUpdateResponseEnvelope struct {
	Errors   []SettingTLSUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingTLSUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   TLSCertificatesAndHostnamesSettingObject   `json:"result,required"`
	// Whether the API call was successful
	Success SettingTLSUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    settingTLSUpdateResponseEnvelopeJSON    `json:"-"`
}

// settingTLSUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingTLSUpdateResponseEnvelope]
type settingTLSUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTLSUpdateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingTLSUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingTLSUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingTLSUpdateResponseEnvelopeErrors]
type settingTLSUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingTLSUpdateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingTLSUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingTLSUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingTLSUpdateResponseEnvelopeMessages]
type settingTLSUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingTLSUpdateResponseEnvelopeSuccess bool

const (
	SettingTLSUpdateResponseEnvelopeSuccessTrue SettingTLSUpdateResponseEnvelopeSuccess = true
)

type SettingTLSDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

// The TLS Setting name.
type SettingTLSDeleteParamsSettingID string

const (
	SettingTLSDeleteParamsSettingIDCiphers       SettingTLSDeleteParamsSettingID = "ciphers"
	SettingTLSDeleteParamsSettingIDMinTLSVersion SettingTLSDeleteParamsSettingID = "min_tls_version"
	SettingTLSDeleteParamsSettingIDHTTP2         SettingTLSDeleteParamsSettingID = "http2"
)

type SettingTLSDeleteResponseEnvelope struct {
	Errors   []SettingTLSDeleteResponseEnvelopeErrors       `json:"errors,required"`
	Messages []SettingTLSDeleteResponseEnvelopeMessages     `json:"messages,required"`
	Result   TLSCertificatesAndHostnamesSettingObjectDelete `json:"result,required"`
	// Whether the API call was successful
	Success SettingTLSDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    settingTLSDeleteResponseEnvelopeJSON    `json:"-"`
}

// settingTLSDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingTLSDeleteResponseEnvelope]
type settingTLSDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTLSDeleteResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingTLSDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// settingTLSDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingTLSDeleteResponseEnvelopeErrors]
type settingTLSDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingTLSDeleteResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingTLSDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// settingTLSDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingTLSDeleteResponseEnvelopeMessages]
type settingTLSDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingTLSDeleteResponseEnvelopeSuccess bool

const (
	SettingTLSDeleteResponseEnvelopeSuccessTrue SettingTLSDeleteResponseEnvelopeSuccess = true
)

type SettingTLSGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

// The TLS Setting name.
type SettingTLSGetParamsSettingID string

const (
	SettingTLSGetParamsSettingIDCiphers       SettingTLSGetParamsSettingID = "ciphers"
	SettingTLSGetParamsSettingIDMinTLSVersion SettingTLSGetParamsSettingID = "min_tls_version"
	SettingTLSGetParamsSettingIDHTTP2         SettingTLSGetParamsSettingID = "http2"
)

type SettingTLSGetResponseEnvelope struct {
	Errors   []SettingTLSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingTLSGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []SettingTLSGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    SettingTLSGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SettingTLSGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       settingTLSGetResponseEnvelopeJSON       `json:"-"`
}

// settingTLSGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingTLSGetResponseEnvelope]
type settingTLSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTLSGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    settingTLSGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingTLSGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingTLSGetResponseEnvelopeErrors]
type settingTLSGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingTLSGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    settingTLSGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingTLSGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingTLSGetResponseEnvelopeMessages]
type settingTLSGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingTLSGetResponseEnvelopeSuccess bool

const (
	SettingTLSGetResponseEnvelopeSuccessTrue SettingTLSGetResponseEnvelopeSuccess = true
)

type SettingTLSGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64 `json:"total_count"`
	// Total pages available of results
	TotalPages float64                                     `json:"total_pages"`
	JSON       settingTLSGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// settingTLSGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [SettingTLSGetResponseEnvelopeResultInfo]
type settingTLSGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

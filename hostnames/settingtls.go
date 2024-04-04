// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostnames

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *SettingTLSService) Update(ctx context.Context, settingID SettingTLSUpdateParamsSettingID, hostname string, params SettingTLSUpdateParams, opts ...option.RequestOption) (res *HostnameStting, err error) {
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
func (r *SettingTLSService) Delete(ctx context.Context, settingID SettingTLSDeleteParamsSettingID, hostname string, body SettingTLSDeleteParams, opts ...option.RequestOption) (res *HostnameSettingDelete, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTLSDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v/%s", body.ZoneID, settingID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HostnameSettingDelete struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	Status   string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	Value     string                    `json:"value"`
	JSON      hostnameSettingDeleteJSON `json:"-"`
}

// hostnameSettingDeleteJSON contains the JSON metadata for the struct
// [HostnameSettingDelete]
type hostnameSettingDeleteJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameSettingDelete) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameSettingDeleteJSON) RawJSON() string {
	return r.raw
}

type HostnameStting struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value HostnameSttingValueUnion `json:"value"`
	JSON  hostnameSttingJSON       `json:"-"`
}

// hostnameSttingJSON contains the JSON metadata for the struct [HostnameStting]
type hostnameSttingJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameStting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameSttingJSON) RawJSON() string {
	return r.raw
}

// The tls setting value.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionString] or
// [hostnames.HostnameSttingValueArray].
type HostnameSttingValueUnion interface {
	ImplementsHostnamesHostnameSttingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HostnameSttingValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(HostnameSttingValueArray{}),
		},
	)
}

type HostnameSttingValueArray []string

func (r HostnameSttingValueArray) ImplementsHostnamesHostnameSttingValueUnion() {}

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
	Value SettingTLSGetResponseValueUnion `json:"value"`
	JSON  settingTLSGetResponseJSON       `json:"-"`
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
type SettingTLSGetResponseValueUnion interface {
	ImplementsHostnamesSettingTLSGetResponseValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingTLSGetResponseValueUnion)(nil)).Elem(),
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

func (r SettingTLSGetResponseValueArray) ImplementsHostnamesSettingTLSGetResponseValueUnion() {}

type SettingTLSUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The tls setting value.
	Value param.Field[SettingTLSUpdateParamsValueUnion] `json:"value,required"`
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

func (r SettingTLSUpdateParamsSettingID) IsKnown() bool {
	switch r {
	case SettingTLSUpdateParamsSettingIDCiphers, SettingTLSUpdateParamsSettingIDMinTLSVersion, SettingTLSUpdateParamsSettingIDHTTP2:
		return true
	}
	return false
}

// The tls setting value.
//
// Satisfied by [shared.UnionFloat], [shared.UnionString],
// [hostnames.SettingTLSUpdateParamsValueArray].
type SettingTLSUpdateParamsValueUnion interface {
	ImplementsHostnamesSettingTLSUpdateParamsValueUnion()
}

type SettingTLSUpdateParamsValueArray []string

func (r SettingTLSUpdateParamsValueArray) ImplementsHostnamesSettingTLSUpdateParamsValueUnion() {}

type SettingTLSUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   HostnameStting                                            `json:"result,required"`
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

// Whether the API call was successful
type SettingTLSUpdateResponseEnvelopeSuccess bool

const (
	SettingTLSUpdateResponseEnvelopeSuccessTrue SettingTLSUpdateResponseEnvelopeSuccess = true
)

func (r SettingTLSUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingTLSUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r SettingTLSDeleteParamsSettingID) IsKnown() bool {
	switch r {
	case SettingTLSDeleteParamsSettingIDCiphers, SettingTLSDeleteParamsSettingIDMinTLSVersion, SettingTLSDeleteParamsSettingIDHTTP2:
		return true
	}
	return false
}

type SettingTLSDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   HostnameSettingDelete                                     `json:"result,required"`
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

// Whether the API call was successful
type SettingTLSDeleteResponseEnvelopeSuccess bool

const (
	SettingTLSDeleteResponseEnvelopeSuccessTrue SettingTLSDeleteResponseEnvelopeSuccess = true
)

func (r SettingTLSDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingTLSDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r SettingTLSGetParamsSettingID) IsKnown() bool {
	switch r {
	case SettingTLSGetParamsSettingIDCiphers, SettingTLSGetParamsSettingIDMinTLSVersion, SettingTLSGetParamsSettingIDHTTP2:
		return true
	}
	return false
}

type SettingTLSGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []SettingTLSGetResponse                                   `json:"result,required,nullable"`
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

// Whether the API call was successful
type SettingTLSGetResponseEnvelopeSuccess bool

const (
	SettingTLSGetResponseEnvelopeSuccessTrue SettingTLSGetResponseEnvelopeSuccess = true
)

func (r SettingTLSGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingTLSGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

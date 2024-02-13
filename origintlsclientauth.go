// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// OriginTLSClientAuthService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOriginTLSClientAuthService]
// method instead.
type OriginTLSClientAuthService struct {
	Options   []option.RequestOption
	Hostnames *OriginTLSClientAuthHostnameService
	Settings  *OriginTLSClientAuthSettingService
}

// NewOriginTLSClientAuthService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOriginTLSClientAuthService(opts ...option.RequestOption) (r *OriginTLSClientAuthService) {
	r = &OriginTLSClientAuthService{}
	r.Options = opts
	r.Hostnames = NewOriginTLSClientAuthHostnameService(opts...)
	r.Settings = NewOriginTLSClientAuthSettingService(opts...)
	return
}

// Upload your own certificate you want Cloudflare to use for edge-to-origin
// communication to override the shared certificate. Please note that it is
// important to keep only one certificate active. Also, make sure to enable
// zone-level authenticated origin pulls by making a PUT call to settings endpoint
// to see the uploaded certificate in use.
func (r *OriginTLSClientAuthService) New(ctx context.Context, zoneID string, body OriginTLSClientAuthNewParams, opts ...option.RequestOption) (res *OriginTLSClientAuthNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Certificates
func (r *OriginTLSClientAuthService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *[]OriginTLSClientAuthListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthListResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Certificate
func (r *OriginTLSClientAuthService) Delete(ctx context.Context, zoneID string, certificateID string, opts ...option.RequestOption) (res *OriginTLSClientAuthDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/%s", zoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Certificate Details
func (r *OriginTLSClientAuthService) Get(ctx context.Context, zoneID string, certificateID string, opts ...option.RequestOption) (res *OriginTLSClientAuthGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/%s", zoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [OriginTLSClientAuthNewResponseUnknown] or
// [shared.UnionString].
type OriginTLSClientAuthNewResponse interface {
	ImplementsOriginTLSClientAuthNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginTLSClientAuthNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type OriginTLSClientAuthListResponse struct {
	// Identifier
	ID string `json:"id"`
	// The zone's leaf certificate.
	Certificate string `json:"certificate"`
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool `json:"enabled"`
	// The zone's private key.
	PrivateKey string                              `json:"private_key"`
	JSON       originTLSClientAuthListResponseJSON `json:"-"`
}

// originTLSClientAuthListResponseJSON contains the JSON metadata for the struct
// [OriginTLSClientAuthListResponse]
type originTLSClientAuthListResponseJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	PrivateKey  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [OriginTLSClientAuthDeleteResponseUnknown] or
// [shared.UnionString].
type OriginTLSClientAuthDeleteResponse interface {
	ImplementsOriginTLSClientAuthDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginTLSClientAuthDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [OriginTLSClientAuthGetResponseUnknown] or
// [shared.UnionString].
type OriginTLSClientAuthGetResponse interface {
	ImplementsOriginTLSClientAuthGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginTLSClientAuthGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type OriginTLSClientAuthNewParams struct {
	// The zone's leaf certificate.
	Certificate param.Field[string] `json:"certificate,required"`
	// The zone's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
}

func (r OriginTLSClientAuthNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OriginTLSClientAuthNewResponseEnvelope struct {
	Errors   []OriginTLSClientAuthNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthNewResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientAuthNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    originTLSClientAuthNewResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthNewResponseEnvelope]
type originTLSClientAuthNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthNewResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    originTLSClientAuthNewResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [OriginTLSClientAuthNewResponseEnvelopeErrors]
type originTLSClientAuthNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthNewResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    originTLSClientAuthNewResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthNewResponseEnvelopeMessages]
type originTLSClientAuthNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OriginTLSClientAuthNewResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthNewResponseEnvelopeSuccessTrue OriginTLSClientAuthNewResponseEnvelopeSuccess = true
)

type OriginTLSClientAuthListResponseEnvelope struct {
	Errors   []OriginTLSClientAuthListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthListResponseEnvelopeMessages `json:"messages,required"`
	Result   []OriginTLSClientAuthListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    OriginTLSClientAuthListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo OriginTLSClientAuthListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       originTLSClientAuthListResponseEnvelopeJSON       `json:"-"`
}

// originTLSClientAuthListResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthListResponseEnvelope]
type originTLSClientAuthListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    originTLSClientAuthListResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [OriginTLSClientAuthListResponseEnvelopeErrors]
type originTLSClientAuthListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    originTLSClientAuthListResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthListResponseEnvelopeMessages]
type originTLSClientAuthListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OriginTLSClientAuthListResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthListResponseEnvelopeSuccessTrue OriginTLSClientAuthListResponseEnvelopeSuccess = true
)

type OriginTLSClientAuthListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       originTLSClientAuthListResponseEnvelopeResultInfoJSON `json:"-"`
}

// originTLSClientAuthListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthListResponseEnvelopeResultInfo]
type originTLSClientAuthListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthDeleteResponseEnvelope struct {
	Errors   []OriginTLSClientAuthDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientAuthDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    originTLSClientAuthDeleteResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthDeleteResponseEnvelope]
type originTLSClientAuthDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthDeleteResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    originTLSClientAuthDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthDeleteResponseEnvelopeErrors]
type originTLSClientAuthDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthDeleteResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    originTLSClientAuthDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthDeleteResponseEnvelopeMessages]
type originTLSClientAuthDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OriginTLSClientAuthDeleteResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthDeleteResponseEnvelopeSuccessTrue OriginTLSClientAuthDeleteResponseEnvelopeSuccess = true
)

type OriginTLSClientAuthGetResponseEnvelope struct {
	Errors   []OriginTLSClientAuthGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthGetResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientAuthGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    originTLSClientAuthGetResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthGetResponseEnvelope]
type originTLSClientAuthGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    originTLSClientAuthGetResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [OriginTLSClientAuthGetResponseEnvelopeErrors]
type originTLSClientAuthGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    originTLSClientAuthGetResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthGetResponseEnvelopeMessages]
type originTLSClientAuthGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OriginTLSClientAuthGetResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthGetResponseEnvelopeSuccessTrue OriginTLSClientAuthGetResponseEnvelopeSuccess = true
)

// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccessUserActiveSessionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccessUserActiveSessionService] method instead.
type AccessUserActiveSessionService struct {
	Options []option.RequestOption
}

// NewAccessUserActiveSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessUserActiveSessionService(opts ...option.RequestOption) (r *AccessUserActiveSessionService) {
	r = &AccessUserActiveSessionService{}
	r.Options = opts
	return
}

// Get an active session for a single user.
func (r *AccessUserActiveSessionService) Get(ctx context.Context, identifier string, id string, nonce string, opts ...option.RequestOption) (res *AccessUserActiveSessionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/users/%s/active_sessions/%s", identifier, id, nonce)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get active sessions for a single user.
func (r *AccessUserActiveSessionService) List(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *AccessUserActiveSessionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/users/%s/active_sessions", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessUserActiveSessionGetResponse struct {
	Errors   []AccessUserActiveSessionGetResponseError   `json:"errors"`
	Messages []AccessUserActiveSessionGetResponseMessage `json:"messages"`
	Result   AccessUserActiveSessionGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessUserActiveSessionGetResponseSuccess `json:"success"`
	JSON    accessUserActiveSessionGetResponseJSON    `json:"-"`
}

// accessUserActiveSessionGetResponseJSON contains the JSON metadata for the struct
// [AccessUserActiveSessionGetResponse]
type accessUserActiveSessionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserActiveSessionGetResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessUserActiveSessionGetResponseErrorJSON `json:"-"`
}

// accessUserActiveSessionGetResponseErrorJSON contains the JSON metadata for the
// struct [AccessUserActiveSessionGetResponseError]
type accessUserActiveSessionGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserActiveSessionGetResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accessUserActiveSessionGetResponseMessageJSON `json:"-"`
}

// accessUserActiveSessionGetResponseMessageJSON contains the JSON metadata for the
// struct [AccessUserActiveSessionGetResponseMessage]
type accessUserActiveSessionGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserActiveSessionGetResponseResult struct {
	AccountID          string                                           `json:"account_id"`
	AuthStatus         string                                           `json:"auth_status"`
	CommonName         string                                           `json:"common_name"`
	DeviceID           string                                           `json:"device_id"`
	DeviceSessions     interface{}                                      `json:"device_sessions"`
	DevicePosture      interface{}                                      `json:"devicePosture"`
	Email              string                                           `json:"email"`
	Geo                AccessUserActiveSessionGetResponseResultGeo      `json:"geo"`
	Iat                float64                                          `json:"iat"`
	Idp                AccessUserActiveSessionGetResponseResultIdp      `json:"idp"`
	IP                 string                                           `json:"ip"`
	IsGateway          bool                                             `json:"is_gateway"`
	IsWarp             bool                                             `json:"is_warp"`
	IsActive           bool                                             `json:"isActive"`
	MtlsAuth           AccessUserActiveSessionGetResponseResultMtlsAuth `json:"mtls_auth"`
	ServiceTokenID     string                                           `json:"service_token_id"`
	ServiceTokenStatus bool                                             `json:"service_token_status"`
	UserUuid           string                                           `json:"user_uuid"`
	Version            float64                                          `json:"version"`
	JSON               accessUserActiveSessionGetResponseResultJSON     `json:"-"`
}

// accessUserActiveSessionGetResponseResultJSON contains the JSON metadata for the
// struct [AccessUserActiveSessionGetResponseResult]
type accessUserActiveSessionGetResponseResultJSON struct {
	AccountID          apijson.Field
	AuthStatus         apijson.Field
	CommonName         apijson.Field
	DeviceID           apijson.Field
	DeviceSessions     apijson.Field
	DevicePosture      apijson.Field
	Email              apijson.Field
	Geo                apijson.Field
	Iat                apijson.Field
	Idp                apijson.Field
	IP                 apijson.Field
	IsGateway          apijson.Field
	IsWarp             apijson.Field
	IsActive           apijson.Field
	MtlsAuth           apijson.Field
	ServiceTokenID     apijson.Field
	ServiceTokenStatus apijson.Field
	UserUuid           apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserActiveSessionGetResponseResultGeo struct {
	Country string                                          `json:"country"`
	JSON    accessUserActiveSessionGetResponseResultGeoJSON `json:"-"`
}

// accessUserActiveSessionGetResponseResultGeoJSON contains the JSON metadata for
// the struct [AccessUserActiveSessionGetResponseResultGeo]
type accessUserActiveSessionGetResponseResultGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseResultGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserActiveSessionGetResponseResultIdp struct {
	ID   string                                          `json:"id"`
	Type string                                          `json:"type"`
	JSON accessUserActiveSessionGetResponseResultIdpJSON `json:"-"`
}

// accessUserActiveSessionGetResponseResultIdpJSON contains the JSON metadata for
// the struct [AccessUserActiveSessionGetResponseResultIdp]
type accessUserActiveSessionGetResponseResultIdpJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseResultIdp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserActiveSessionGetResponseResultMtlsAuth struct {
	AuthStatus    string                                               `json:"auth_status"`
	CertIssuerDn  string                                               `json:"cert_issuer_dn"`
	CertIssuerSki string                                               `json:"cert_issuer_ski"`
	CertPresented bool                                                 `json:"cert_presented"`
	CertSerial    string                                               `json:"cert_serial"`
	JSON          accessUserActiveSessionGetResponseResultMtlsAuthJSON `json:"-"`
}

// accessUserActiveSessionGetResponseResultMtlsAuthJSON contains the JSON metadata
// for the struct [AccessUserActiveSessionGetResponseResultMtlsAuth]
type accessUserActiveSessionGetResponseResultMtlsAuthJSON struct {
	AuthStatus    apijson.Field
	CertIssuerDn  apijson.Field
	CertIssuerSki apijson.Field
	CertPresented apijson.Field
	CertSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseResultMtlsAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessUserActiveSessionGetResponseSuccess bool

const (
	AccessUserActiveSessionGetResponseSuccessTrue AccessUserActiveSessionGetResponseSuccess = true
)

type AccessUserActiveSessionListResponse struct {
	Errors     []AccessUserActiveSessionListResponseError    `json:"errors"`
	Messages   []AccessUserActiveSessionListResponseMessage  `json:"messages"`
	Result     []AccessUserActiveSessionListResponseResult   `json:"result"`
	ResultInfo AccessUserActiveSessionListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessUserActiveSessionListResponseSuccess `json:"success"`
	JSON    accessUserActiveSessionListResponseJSON    `json:"-"`
}

// accessUserActiveSessionListResponseJSON contains the JSON metadata for the
// struct [AccessUserActiveSessionListResponse]
type accessUserActiveSessionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserActiveSessionListResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accessUserActiveSessionListResponseErrorJSON `json:"-"`
}

// accessUserActiveSessionListResponseErrorJSON contains the JSON metadata for the
// struct [AccessUserActiveSessionListResponseError]
type accessUserActiveSessionListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserActiveSessionListResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accessUserActiveSessionListResponseMessageJSON `json:"-"`
}

// accessUserActiveSessionListResponseMessageJSON contains the JSON metadata for
// the struct [AccessUserActiveSessionListResponseMessage]
type accessUserActiveSessionListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserActiveSessionListResponseResult struct {
	Expiration int64                                             `json:"expiration"`
	Metadata   AccessUserActiveSessionListResponseResultMetadata `json:"metadata"`
	Name       string                                            `json:"name"`
	JSON       accessUserActiveSessionListResponseResultJSON     `json:"-"`
}

// accessUserActiveSessionListResponseResultJSON contains the JSON metadata for the
// struct [AccessUserActiveSessionListResponseResult]
type accessUserActiveSessionListResponseResultJSON struct {
	Expiration  apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserActiveSessionListResponseResultMetadata struct {
	Apps    interface{}                                           `json:"apps"`
	Expires int64                                                 `json:"expires"`
	Iat     int64                                                 `json:"iat"`
	Nonce   string                                                `json:"nonce"`
	TTL     int64                                                 `json:"ttl"`
	JSON    accessUserActiveSessionListResponseResultMetadataJSON `json:"-"`
}

// accessUserActiveSessionListResponseResultMetadataJSON contains the JSON metadata
// for the struct [AccessUserActiveSessionListResponseResultMetadata]
type accessUserActiveSessionListResponseResultMetadataJSON struct {
	Apps        apijson.Field
	Expires     apijson.Field
	Iat         apijson.Field
	Nonce       apijson.Field
	TTL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponseResultMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserActiveSessionListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       accessUserActiveSessionListResponseResultInfoJSON `json:"-"`
}

// accessUserActiveSessionListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccessUserActiveSessionListResponseResultInfo]
type accessUserActiveSessionListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessUserActiveSessionListResponseSuccess bool

const (
	AccessUserActiveSessionListResponseSuccessTrue AccessUserActiveSessionListResponseSuccess = true
)

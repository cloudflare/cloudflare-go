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

// AccessUserLastSeenIdentityService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccessUserLastSeenIdentityService] method instead.
type AccessUserLastSeenIdentityService struct {
	Options []option.RequestOption
}

// NewAccessUserLastSeenIdentityService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessUserLastSeenIdentityService(opts ...option.RequestOption) (r *AccessUserLastSeenIdentityService) {
	r = &AccessUserLastSeenIdentityService{}
	r.Options = opts
	return
}

// Get last seen identity for a single user.
func (r *AccessUserLastSeenIdentityService) Get(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *AccessUserLastSeenIdentityGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/users/%s/last_seen_identity", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessUserLastSeenIdentityGetResponse struct {
	Errors   []AccessUserLastSeenIdentityGetResponseError   `json:"errors"`
	Messages []AccessUserLastSeenIdentityGetResponseMessage `json:"messages"`
	Result   AccessUserLastSeenIdentityGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessUserLastSeenIdentityGetResponseSuccess `json:"success"`
	JSON    accessUserLastSeenIdentityGetResponseJSON    `json:"-"`
}

// accessUserLastSeenIdentityGetResponseJSON contains the JSON metadata for the
// struct [AccessUserLastSeenIdentityGetResponse]
type accessUserLastSeenIdentityGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityGetResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accessUserLastSeenIdentityGetResponseErrorJSON `json:"-"`
}

// accessUserLastSeenIdentityGetResponseErrorJSON contains the JSON metadata for
// the struct [AccessUserLastSeenIdentityGetResponseError]
type accessUserLastSeenIdentityGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityGetResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessUserLastSeenIdentityGetResponseMessageJSON `json:"-"`
}

// accessUserLastSeenIdentityGetResponseMessageJSON contains the JSON metadata for
// the struct [AccessUserLastSeenIdentityGetResponseMessage]
type accessUserLastSeenIdentityGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityGetResponseResult struct {
	AccountID          string                                              `json:"account_id"`
	AuthStatus         string                                              `json:"auth_status"`
	CommonName         string                                              `json:"common_name"`
	DeviceID           string                                              `json:"device_id"`
	DeviceSessions     interface{}                                         `json:"device_sessions"`
	DevicePosture      interface{}                                         `json:"devicePosture"`
	Email              string                                              `json:"email"`
	Geo                AccessUserLastSeenIdentityGetResponseResultGeo      `json:"geo"`
	Iat                float64                                             `json:"iat"`
	Idp                AccessUserLastSeenIdentityGetResponseResultIdp      `json:"idp"`
	IP                 string                                              `json:"ip"`
	IsGateway          bool                                                `json:"is_gateway"`
	IsWarp             bool                                                `json:"is_warp"`
	MtlsAuth           AccessUserLastSeenIdentityGetResponseResultMtlsAuth `json:"mtls_auth"`
	ServiceTokenID     string                                              `json:"service_token_id"`
	ServiceTokenStatus bool                                                `json:"service_token_status"`
	UserUuid           string                                              `json:"user_uuid"`
	Version            float64                                             `json:"version"`
	JSON               accessUserLastSeenIdentityGetResponseResultJSON     `json:"-"`
}

// accessUserLastSeenIdentityGetResponseResultJSON contains the JSON metadata for
// the struct [AccessUserLastSeenIdentityGetResponseResult]
type accessUserLastSeenIdentityGetResponseResultJSON struct {
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
	MtlsAuth           apijson.Field
	ServiceTokenID     apijson.Field
	ServiceTokenStatus apijson.Field
	UserUuid           apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityGetResponseResultGeo struct {
	Country string                                             `json:"country"`
	JSON    accessUserLastSeenIdentityGetResponseResultGeoJSON `json:"-"`
}

// accessUserLastSeenIdentityGetResponseResultGeoJSON contains the JSON metadata
// for the struct [AccessUserLastSeenIdentityGetResponseResultGeo]
type accessUserLastSeenIdentityGetResponseResultGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseResultGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityGetResponseResultIdp struct {
	ID   string                                             `json:"id"`
	Type string                                             `json:"type"`
	JSON accessUserLastSeenIdentityGetResponseResultIdpJSON `json:"-"`
}

// accessUserLastSeenIdentityGetResponseResultIdpJSON contains the JSON metadata
// for the struct [AccessUserLastSeenIdentityGetResponseResultIdp]
type accessUserLastSeenIdentityGetResponseResultIdpJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseResultIdp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityGetResponseResultMtlsAuth struct {
	AuthStatus    string                                                  `json:"auth_status"`
	CertIssuerDn  string                                                  `json:"cert_issuer_dn"`
	CertIssuerSki string                                                  `json:"cert_issuer_ski"`
	CertPresented bool                                                    `json:"cert_presented"`
	CertSerial    string                                                  `json:"cert_serial"`
	JSON          accessUserLastSeenIdentityGetResponseResultMtlsAuthJSON `json:"-"`
}

// accessUserLastSeenIdentityGetResponseResultMtlsAuthJSON contains the JSON
// metadata for the struct [AccessUserLastSeenIdentityGetResponseResultMtlsAuth]
type accessUserLastSeenIdentityGetResponseResultMtlsAuthJSON struct {
	AuthStatus    apijson.Field
	CertIssuerDn  apijson.Field
	CertIssuerSki apijson.Field
	CertPresented apijson.Field
	CertSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseResultMtlsAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessUserLastSeenIdentityGetResponseSuccess bool

const (
	AccessUserLastSeenIdentityGetResponseSuccessTrue AccessUserLastSeenIdentityGetResponseSuccess = true
)

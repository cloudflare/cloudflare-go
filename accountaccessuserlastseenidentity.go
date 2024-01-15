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

// AccountAccessUserLastSeenIdentityService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAccessUserLastSeenIdentityService] method instead.
type AccountAccessUserLastSeenIdentityService struct {
	Options []option.RequestOption
}

// NewAccountAccessUserLastSeenIdentityService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessUserLastSeenIdentityService(opts ...option.RequestOption) (r *AccountAccessUserLastSeenIdentityService) {
	r = &AccountAccessUserLastSeenIdentityService{}
	r.Options = opts
	return
}

// Get last seen identity for a single user.
func (r *AccountAccessUserLastSeenIdentityService) Get(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *AccountAccessUserLastSeenIdentityGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/users/%s/last_seen_identity", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAccessUserLastSeenIdentityGetResponse struct {
	Errors   []AccountAccessUserLastSeenIdentityGetResponseError   `json:"errors"`
	Messages []AccountAccessUserLastSeenIdentityGetResponseMessage `json:"messages"`
	Result   AccountAccessUserLastSeenIdentityGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessUserLastSeenIdentityGetResponseSuccess `json:"success"`
	JSON    accountAccessUserLastSeenIdentityGetResponseJSON    `json:"-"`
}

// accountAccessUserLastSeenIdentityGetResponseJSON contains the JSON metadata for
// the struct [AccountAccessUserLastSeenIdentityGetResponse]
type accountAccessUserLastSeenIdentityGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserLastSeenIdentityGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserLastSeenIdentityGetResponseError struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountAccessUserLastSeenIdentityGetResponseErrorJSON `json:"-"`
}

// accountAccessUserLastSeenIdentityGetResponseErrorJSON contains the JSON metadata
// for the struct [AccountAccessUserLastSeenIdentityGetResponseError]
type accountAccessUserLastSeenIdentityGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserLastSeenIdentityGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserLastSeenIdentityGetResponseMessage struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accountAccessUserLastSeenIdentityGetResponseMessageJSON `json:"-"`
}

// accountAccessUserLastSeenIdentityGetResponseMessageJSON contains the JSON
// metadata for the struct [AccountAccessUserLastSeenIdentityGetResponseMessage]
type accountAccessUserLastSeenIdentityGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserLastSeenIdentityGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserLastSeenIdentityGetResponseResult struct {
	AccountID          string                                                     `json:"account_id"`
	AuthStatus         string                                                     `json:"auth_status"`
	CommonName         string                                                     `json:"common_name"`
	DeviceID           string                                                     `json:"device_id"`
	DeviceSessions     interface{}                                                `json:"device_sessions"`
	DevicePosture      interface{}                                                `json:"devicePosture"`
	Email              string                                                     `json:"email"`
	Geo                AccountAccessUserLastSeenIdentityGetResponseResultGeo      `json:"geo"`
	Iat                float64                                                    `json:"iat"`
	Idp                AccountAccessUserLastSeenIdentityGetResponseResultIdp      `json:"idp"`
	IP                 string                                                     `json:"ip"`
	IsGateway          bool                                                       `json:"is_gateway"`
	IsWarp             bool                                                       `json:"is_warp"`
	MtlsAuth           AccountAccessUserLastSeenIdentityGetResponseResultMtlsAuth `json:"mtls_auth"`
	ServiceTokenID     string                                                     `json:"service_token_id"`
	ServiceTokenStatus bool                                                       `json:"service_token_status"`
	UserUuid           string                                                     `json:"user_uuid"`
	Version            float64                                                    `json:"version"`
	JSON               accountAccessUserLastSeenIdentityGetResponseResultJSON     `json:"-"`
}

// accountAccessUserLastSeenIdentityGetResponseResultJSON contains the JSON
// metadata for the struct [AccountAccessUserLastSeenIdentityGetResponseResult]
type accountAccessUserLastSeenIdentityGetResponseResultJSON struct {
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

func (r *AccountAccessUserLastSeenIdentityGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserLastSeenIdentityGetResponseResultGeo struct {
	Country string                                                    `json:"country"`
	JSON    accountAccessUserLastSeenIdentityGetResponseResultGeoJSON `json:"-"`
}

// accountAccessUserLastSeenIdentityGetResponseResultGeoJSON contains the JSON
// metadata for the struct [AccountAccessUserLastSeenIdentityGetResponseResultGeo]
type accountAccessUserLastSeenIdentityGetResponseResultGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserLastSeenIdentityGetResponseResultGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserLastSeenIdentityGetResponseResultIdp struct {
	ID   string                                                    `json:"id"`
	Type string                                                    `json:"type"`
	JSON accountAccessUserLastSeenIdentityGetResponseResultIdpJSON `json:"-"`
}

// accountAccessUserLastSeenIdentityGetResponseResultIdpJSON contains the JSON
// metadata for the struct [AccountAccessUserLastSeenIdentityGetResponseResultIdp]
type accountAccessUserLastSeenIdentityGetResponseResultIdpJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserLastSeenIdentityGetResponseResultIdp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserLastSeenIdentityGetResponseResultMtlsAuth struct {
	AuthStatus    string                                                         `json:"auth_status"`
	CertIssuerDn  string                                                         `json:"cert_issuer_dn"`
	CertIssuerSki string                                                         `json:"cert_issuer_ski"`
	CertPresented bool                                                           `json:"cert_presented"`
	CertSerial    string                                                         `json:"cert_serial"`
	JSON          accountAccessUserLastSeenIdentityGetResponseResultMtlsAuthJSON `json:"-"`
}

// accountAccessUserLastSeenIdentityGetResponseResultMtlsAuthJSON contains the JSON
// metadata for the struct
// [AccountAccessUserLastSeenIdentityGetResponseResultMtlsAuth]
type accountAccessUserLastSeenIdentityGetResponseResultMtlsAuthJSON struct {
	AuthStatus    apijson.Field
	CertIssuerDn  apijson.Field
	CertIssuerSki apijson.Field
	CertPresented apijson.Field
	CertSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessUserLastSeenIdentityGetResponseResultMtlsAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessUserLastSeenIdentityGetResponseSuccess bool

const (
	AccountAccessUserLastSeenIdentityGetResponseSuccessTrue AccountAccessUserLastSeenIdentityGetResponseSuccess = true
)

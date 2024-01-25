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

// AccountAccessUserActiveSessionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAccessUserActiveSessionService] method instead.
type AccountAccessUserActiveSessionService struct {
	Options []option.RequestOption
}

// NewAccountAccessUserActiveSessionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessUserActiveSessionService(opts ...option.RequestOption) (r *AccountAccessUserActiveSessionService) {
	r = &AccountAccessUserActiveSessionService{}
	r.Options = opts
	return
}

// Get active sessions for a single user.
func (r *AccountAccessUserActiveSessionService) List(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *AccountAccessUserActiveSessionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/users/%s/active_sessions", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get an active session for a single user.
func (r *AccountAccessUserActiveSessionService) GetActiveSession(ctx context.Context, identifier string, id string, nonce string, opts ...option.RequestOption) (res *AccountAccessUserActiveSessionGetActiveSessionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/users/%s/active_sessions/%s", identifier, id, nonce)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAccessUserActiveSessionListResponse struct {
	Errors     []AccountAccessUserActiveSessionListResponseError    `json:"errors"`
	Messages   []AccountAccessUserActiveSessionListResponseMessage  `json:"messages"`
	Result     []AccountAccessUserActiveSessionListResponseResult   `json:"result"`
	ResultInfo AccountAccessUserActiveSessionListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAccessUserActiveSessionListResponseSuccess `json:"success"`
	JSON    accountAccessUserActiveSessionListResponseJSON    `json:"-"`
}

// accountAccessUserActiveSessionListResponseJSON contains the JSON metadata for
// the struct [AccountAccessUserActiveSessionListResponse]
type accountAccessUserActiveSessionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserActiveSessionListResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountAccessUserActiveSessionListResponseErrorJSON `json:"-"`
}

// accountAccessUserActiveSessionListResponseErrorJSON contains the JSON metadata
// for the struct [AccountAccessUserActiveSessionListResponseError]
type accountAccessUserActiveSessionListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserActiveSessionListResponseMessage struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountAccessUserActiveSessionListResponseMessageJSON `json:"-"`
}

// accountAccessUserActiveSessionListResponseMessageJSON contains the JSON metadata
// for the struct [AccountAccessUserActiveSessionListResponseMessage]
type accountAccessUserActiveSessionListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserActiveSessionListResponseResult struct {
	Expiration int64                                                    `json:"expiration"`
	Metadata   AccountAccessUserActiveSessionListResponseResultMetadata `json:"metadata"`
	Name       string                                                   `json:"name"`
	JSON       accountAccessUserActiveSessionListResponseResultJSON     `json:"-"`
}

// accountAccessUserActiveSessionListResponseResultJSON contains the JSON metadata
// for the struct [AccountAccessUserActiveSessionListResponseResult]
type accountAccessUserActiveSessionListResponseResultJSON struct {
	Expiration  apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserActiveSessionListResponseResultMetadata struct {
	Apps    interface{}                                                  `json:"apps"`
	Expires int64                                                        `json:"expires"`
	Iat     int64                                                        `json:"iat"`
	Nonce   string                                                       `json:"nonce"`
	Ttl     int64                                                        `json:"ttl"`
	JSON    accountAccessUserActiveSessionListResponseResultMetadataJSON `json:"-"`
}

// accountAccessUserActiveSessionListResponseResultMetadataJSON contains the JSON
// metadata for the struct
// [AccountAccessUserActiveSessionListResponseResultMetadata]
type accountAccessUserActiveSessionListResponseResultMetadataJSON struct {
	Apps        apijson.Field
	Expires     apijson.Field
	Iat         apijson.Field
	Nonce       apijson.Field
	Ttl         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionListResponseResultMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserActiveSessionListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       accountAccessUserActiveSessionListResponseResultInfoJSON `json:"-"`
}

// accountAccessUserActiveSessionListResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountAccessUserActiveSessionListResponseResultInfo]
type accountAccessUserActiveSessionListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessUserActiveSessionListResponseSuccess bool

const (
	AccountAccessUserActiveSessionListResponseSuccessTrue AccountAccessUserActiveSessionListResponseSuccess = true
)

type AccountAccessUserActiveSessionGetActiveSessionResponse struct {
	Errors   []AccountAccessUserActiveSessionGetActiveSessionResponseError   `json:"errors"`
	Messages []AccountAccessUserActiveSessionGetActiveSessionResponseMessage `json:"messages"`
	Result   AccountAccessUserActiveSessionGetActiveSessionResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessUserActiveSessionGetActiveSessionResponseSuccess `json:"success"`
	JSON    accountAccessUserActiveSessionGetActiveSessionResponseJSON    `json:"-"`
}

// accountAccessUserActiveSessionGetActiveSessionResponseJSON contains the JSON
// metadata for the struct [AccountAccessUserActiveSessionGetActiveSessionResponse]
type accountAccessUserActiveSessionGetActiveSessionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionGetActiveSessionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserActiveSessionGetActiveSessionResponseError struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    accountAccessUserActiveSessionGetActiveSessionResponseErrorJSON `json:"-"`
}

// accountAccessUserActiveSessionGetActiveSessionResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountAccessUserActiveSessionGetActiveSessionResponseError]
type accountAccessUserActiveSessionGetActiveSessionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionGetActiveSessionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserActiveSessionGetActiveSessionResponseMessage struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accountAccessUserActiveSessionGetActiveSessionResponseMessageJSON `json:"-"`
}

// accountAccessUserActiveSessionGetActiveSessionResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountAccessUserActiveSessionGetActiveSessionResponseMessage]
type accountAccessUserActiveSessionGetActiveSessionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionGetActiveSessionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserActiveSessionGetActiveSessionResponseResult struct {
	AccountID          string                                                               `json:"account_id"`
	AuthStatus         string                                                               `json:"auth_status"`
	CommonName         string                                                               `json:"common_name"`
	DeviceID           string                                                               `json:"device_id"`
	DeviceSessions     interface{}                                                          `json:"device_sessions"`
	DevicePosture      interface{}                                                          `json:"devicePosture"`
	Email              string                                                               `json:"email"`
	Geo                AccountAccessUserActiveSessionGetActiveSessionResponseResultGeo      `json:"geo"`
	Iat                float64                                                              `json:"iat"`
	Idp                AccountAccessUserActiveSessionGetActiveSessionResponseResultIdp      `json:"idp"`
	IP                 string                                                               `json:"ip"`
	IsGateway          bool                                                                 `json:"is_gateway"`
	IsWarp             bool                                                                 `json:"is_warp"`
	IsActive           bool                                                                 `json:"isActive"`
	MtlsAuth           AccountAccessUserActiveSessionGetActiveSessionResponseResultMtlsAuth `json:"mtls_auth"`
	ServiceTokenID     string                                                               `json:"service_token_id"`
	ServiceTokenStatus bool                                                                 `json:"service_token_status"`
	UserUuid           string                                                               `json:"user_uuid"`
	Version            float64                                                              `json:"version"`
	JSON               accountAccessUserActiveSessionGetActiveSessionResponseResultJSON     `json:"-"`
}

// accountAccessUserActiveSessionGetActiveSessionResponseResultJSON contains the
// JSON metadata for the struct
// [AccountAccessUserActiveSessionGetActiveSessionResponseResult]
type accountAccessUserActiveSessionGetActiveSessionResponseResultJSON struct {
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

func (r *AccountAccessUserActiveSessionGetActiveSessionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserActiveSessionGetActiveSessionResponseResultGeo struct {
	Country string                                                              `json:"country"`
	JSON    accountAccessUserActiveSessionGetActiveSessionResponseResultGeoJSON `json:"-"`
}

// accountAccessUserActiveSessionGetActiveSessionResponseResultGeoJSON contains the
// JSON metadata for the struct
// [AccountAccessUserActiveSessionGetActiveSessionResponseResultGeo]
type accountAccessUserActiveSessionGetActiveSessionResponseResultGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionGetActiveSessionResponseResultGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserActiveSessionGetActiveSessionResponseResultIdp struct {
	ID   string                                                              `json:"id"`
	Type string                                                              `json:"type"`
	JSON accountAccessUserActiveSessionGetActiveSessionResponseResultIdpJSON `json:"-"`
}

// accountAccessUserActiveSessionGetActiveSessionResponseResultIdpJSON contains the
// JSON metadata for the struct
// [AccountAccessUserActiveSessionGetActiveSessionResponseResultIdp]
type accountAccessUserActiveSessionGetActiveSessionResponseResultIdpJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionGetActiveSessionResponseResultIdp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserActiveSessionGetActiveSessionResponseResultMtlsAuth struct {
	AuthStatus    string                                                                   `json:"auth_status"`
	CertIssuerDn  string                                                                   `json:"cert_issuer_dn"`
	CertIssuerSki string                                                                   `json:"cert_issuer_ski"`
	CertPresented bool                                                                     `json:"cert_presented"`
	CertSerial    string                                                                   `json:"cert_serial"`
	JSON          accountAccessUserActiveSessionGetActiveSessionResponseResultMtlsAuthJSON `json:"-"`
}

// accountAccessUserActiveSessionGetActiveSessionResponseResultMtlsAuthJSON
// contains the JSON metadata for the struct
// [AccountAccessUserActiveSessionGetActiveSessionResponseResultMtlsAuth]
type accountAccessUserActiveSessionGetActiveSessionResponseResultMtlsAuthJSON struct {
	AuthStatus    apijson.Field
	CertIssuerDn  apijson.Field
	CertIssuerSki apijson.Field
	CertPresented apijson.Field
	CertSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionGetActiveSessionResponseResultMtlsAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessUserActiveSessionGetActiveSessionResponseSuccess bool

const (
	AccountAccessUserActiveSessionGetActiveSessionResponseSuccessTrue AccountAccessUserActiveSessionGetActiveSessionResponseSuccess = true
)

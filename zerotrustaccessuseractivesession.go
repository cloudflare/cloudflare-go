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

// ZeroTrustAccessUserActiveSessionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustAccessUserActiveSessionService] method instead.
type ZeroTrustAccessUserActiveSessionService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessUserActiveSessionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustAccessUserActiveSessionService(opts ...option.RequestOption) (r *ZeroTrustAccessUserActiveSessionService) {
	r = &ZeroTrustAccessUserActiveSessionService{}
	r.Options = opts
	return
}

// Get active sessions for a single user.
func (r *ZeroTrustAccessUserActiveSessionService) List(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *[]ZeroTrustAccessUserActiveSessionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessUserActiveSessionListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/users/%s/active_sessions", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get an active session for a single user.
func (r *ZeroTrustAccessUserActiveSessionService) Get(ctx context.Context, identifier string, id string, nonce string, opts ...option.RequestOption) (res *ZeroTrustAccessUserActiveSessionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessUserActiveSessionGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/users/%s/active_sessions/%s", identifier, id, nonce)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustAccessUserActiveSessionListResponse struct {
	Expiration int64                                                `json:"expiration"`
	Metadata   ZeroTrustAccessUserActiveSessionListResponseMetadata `json:"metadata"`
	Name       string                                               `json:"name"`
	JSON       zeroTrustAccessUserActiveSessionListResponseJSON     `json:"-"`
}

// zeroTrustAccessUserActiveSessionListResponseJSON contains the JSON metadata for
// the struct [ZeroTrustAccessUserActiveSessionListResponse]
type zeroTrustAccessUserActiveSessionListResponseJSON struct {
	Expiration  apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionListResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserActiveSessionListResponseMetadata struct {
	Apps    map[string]ZeroTrustAccessUserActiveSessionListResponseMetadataApp `json:"apps"`
	Expires int64                                                              `json:"expires"`
	Iat     int64                                                              `json:"iat"`
	Nonce   string                                                             `json:"nonce"`
	TTL     int64                                                              `json:"ttl"`
	JSON    zeroTrustAccessUserActiveSessionListResponseMetadataJSON           `json:"-"`
}

// zeroTrustAccessUserActiveSessionListResponseMetadataJSON contains the JSON
// metadata for the struct [ZeroTrustAccessUserActiveSessionListResponseMetadata]
type zeroTrustAccessUserActiveSessionListResponseMetadataJSON struct {
	Apps        apijson.Field
	Expires     apijson.Field
	Iat         apijson.Field
	Nonce       apijson.Field
	TTL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionListResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionListResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserActiveSessionListResponseMetadataApp struct {
	Hostname string                                                      `json:"hostname"`
	Name     string                                                      `json:"name"`
	Type     string                                                      `json:"type"`
	Uid      string                                                      `json:"uid"`
	JSON     zeroTrustAccessUserActiveSessionListResponseMetadataAppJSON `json:"-"`
}

// zeroTrustAccessUserActiveSessionListResponseMetadataAppJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessUserActiveSessionListResponseMetadataApp]
type zeroTrustAccessUserActiveSessionListResponseMetadataAppJSON struct {
	Hostname    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionListResponseMetadataApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionListResponseMetadataAppJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserActiveSessionGetResponse struct {
	AccountID          string                                                              `json:"account_id"`
	AuthStatus         string                                                              `json:"auth_status"`
	CommonName         string                                                              `json:"common_name"`
	DeviceID           string                                                              `json:"device_id"`
	DeviceSessions     map[string]ZeroTrustAccessUserActiveSessionGetResponseDeviceSession `json:"device_sessions"`
	DevicePosture      map[string]ZeroTrustAccessUserActiveSessionGetResponseDevicePosture `json:"devicePosture"`
	Email              string                                                              `json:"email"`
	Geo                ZeroTrustAccessUserActiveSessionGetResponseGeo                      `json:"geo"`
	Iat                float64                                                             `json:"iat"`
	Idp                ZeroTrustAccessUserActiveSessionGetResponseIdp                      `json:"idp"`
	IP                 string                                                              `json:"ip"`
	IsGateway          bool                                                                `json:"is_gateway"`
	IsWARP             bool                                                                `json:"is_warp"`
	IsActive           bool                                                                `json:"isActive"`
	MTLSAuth           ZeroTrustAccessUserActiveSessionGetResponseMTLSAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                                                              `json:"service_token_id"`
	ServiceTokenStatus bool                                                                `json:"service_token_status"`
	UserUUID           string                                                              `json:"user_uuid"`
	Version            float64                                                             `json:"version"`
	JSON               zeroTrustAccessUserActiveSessionGetResponseJSON                     `json:"-"`
}

// zeroTrustAccessUserActiveSessionGetResponseJSON contains the JSON metadata for
// the struct [ZeroTrustAccessUserActiveSessionGetResponse]
type zeroTrustAccessUserActiveSessionGetResponseJSON struct {
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
	IsWARP             apijson.Field
	IsActive           apijson.Field
	MTLSAuth           apijson.Field
	ServiceTokenID     apijson.Field
	ServiceTokenStatus apijson.Field
	UserUUID           apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserActiveSessionGetResponseDeviceSession struct {
	LastAuthenticated float64                                                      `json:"last_authenticated"`
	JSON              zeroTrustAccessUserActiveSessionGetResponseDeviceSessionJSON `json:"-"`
}

// zeroTrustAccessUserActiveSessionGetResponseDeviceSessionJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessUserActiveSessionGetResponseDeviceSession]
type zeroTrustAccessUserActiveSessionGetResponseDeviceSessionJSON struct {
	LastAuthenticated apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionGetResponseDeviceSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionGetResponseDeviceSessionJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserActiveSessionGetResponseDevicePosture struct {
	ID          string                                                        `json:"id"`
	Check       ZeroTrustAccessUserActiveSessionGetResponseDevicePostureCheck `json:"check"`
	Data        interface{}                                                   `json:"data"`
	Description string                                                        `json:"description"`
	Error       string                                                        `json:"error"`
	RuleName    string                                                        `json:"rule_name"`
	Success     bool                                                          `json:"success"`
	Timestamp   string                                                        `json:"timestamp"`
	Type        string                                                        `json:"type"`
	JSON        zeroTrustAccessUserActiveSessionGetResponseDevicePostureJSON  `json:"-"`
}

// zeroTrustAccessUserActiveSessionGetResponseDevicePostureJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessUserActiveSessionGetResponseDevicePosture]
type zeroTrustAccessUserActiveSessionGetResponseDevicePostureJSON struct {
	ID          apijson.Field
	Check       apijson.Field
	Data        apijson.Field
	Description apijson.Field
	Error       apijson.Field
	RuleName    apijson.Field
	Success     apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionGetResponseDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionGetResponseDevicePostureJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserActiveSessionGetResponseDevicePostureCheck struct {
	Exists bool                                                              `json:"exists"`
	Path   string                                                            `json:"path"`
	JSON   zeroTrustAccessUserActiveSessionGetResponseDevicePostureCheckJSON `json:"-"`
}

// zeroTrustAccessUserActiveSessionGetResponseDevicePostureCheckJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessUserActiveSessionGetResponseDevicePostureCheck]
type zeroTrustAccessUserActiveSessionGetResponseDevicePostureCheckJSON struct {
	Exists      apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionGetResponseDevicePostureCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionGetResponseDevicePostureCheckJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserActiveSessionGetResponseGeo struct {
	Country string                                             `json:"country"`
	JSON    zeroTrustAccessUserActiveSessionGetResponseGeoJSON `json:"-"`
}

// zeroTrustAccessUserActiveSessionGetResponseGeoJSON contains the JSON metadata
// for the struct [ZeroTrustAccessUserActiveSessionGetResponseGeo]
type zeroTrustAccessUserActiveSessionGetResponseGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionGetResponseGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionGetResponseGeoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserActiveSessionGetResponseIdp struct {
	ID   string                                             `json:"id"`
	Type string                                             `json:"type"`
	JSON zeroTrustAccessUserActiveSessionGetResponseIdpJSON `json:"-"`
}

// zeroTrustAccessUserActiveSessionGetResponseIdpJSON contains the JSON metadata
// for the struct [ZeroTrustAccessUserActiveSessionGetResponseIdp]
type zeroTrustAccessUserActiveSessionGetResponseIdpJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionGetResponseIdp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionGetResponseIdpJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserActiveSessionGetResponseMTLSAuth struct {
	AuthStatus    string                                                  `json:"auth_status"`
	CertIssuerDn  string                                                  `json:"cert_issuer_dn"`
	CertIssuerSki string                                                  `json:"cert_issuer_ski"`
	CertPresented bool                                                    `json:"cert_presented"`
	CertSerial    string                                                  `json:"cert_serial"`
	JSON          zeroTrustAccessUserActiveSessionGetResponseMTLSAuthJSON `json:"-"`
}

// zeroTrustAccessUserActiveSessionGetResponseMTLSAuthJSON contains the JSON
// metadata for the struct [ZeroTrustAccessUserActiveSessionGetResponseMTLSAuth]
type zeroTrustAccessUserActiveSessionGetResponseMTLSAuthJSON struct {
	AuthStatus    apijson.Field
	CertIssuerDn  apijson.Field
	CertIssuerSki apijson.Field
	CertPresented apijson.Field
	CertSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionGetResponseMTLSAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionGetResponseMTLSAuthJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserActiveSessionListResponseEnvelope struct {
	Errors   []ZeroTrustAccessUserActiveSessionListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessUserActiveSessionListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustAccessUserActiveSessionListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessUserActiveSessionListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessUserActiveSessionListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessUserActiveSessionListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessUserActiveSessionListResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessUserActiveSessionListResponseEnvelope]
type zeroTrustAccessUserActiveSessionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserActiveSessionListResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustAccessUserActiveSessionListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessUserActiveSessionListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessUserActiveSessionListResponseEnvelopeErrors]
type zeroTrustAccessUserActiveSessionListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserActiveSessionListResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zeroTrustAccessUserActiveSessionListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessUserActiveSessionListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessUserActiveSessionListResponseEnvelopeMessages]
type zeroTrustAccessUserActiveSessionListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessUserActiveSessionListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessUserActiveSessionListResponseEnvelopeSuccessTrue ZeroTrustAccessUserActiveSessionListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessUserActiveSessionListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                            `json:"total_count"`
	JSON       zeroTrustAccessUserActiveSessionListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessUserActiveSessionListResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessUserActiveSessionListResponseEnvelopeResultInfo]
type zeroTrustAccessUserActiveSessionListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserActiveSessionGetResponseEnvelope struct {
	Errors   []ZeroTrustAccessUserActiveSessionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessUserActiveSessionGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessUserActiveSessionGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessUserActiveSessionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessUserActiveSessionGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessUserActiveSessionGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessUserActiveSessionGetResponseEnvelope]
type zeroTrustAccessUserActiveSessionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserActiveSessionGetResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zeroTrustAccessUserActiveSessionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessUserActiveSessionGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessUserActiveSessionGetResponseEnvelopeErrors]
type zeroTrustAccessUserActiveSessionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserActiveSessionGetResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustAccessUserActiveSessionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessUserActiveSessionGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessUserActiveSessionGetResponseEnvelopeMessages]
type zeroTrustAccessUserActiveSessionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserActiveSessionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserActiveSessionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessUserActiveSessionGetResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessUserActiveSessionGetResponseEnvelopeSuccessTrue ZeroTrustAccessUserActiveSessionGetResponseEnvelopeSuccess = true
)

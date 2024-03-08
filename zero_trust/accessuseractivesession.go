// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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

// Get active sessions for a single user.
func (r *AccessUserActiveSessionService) List(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *[]AccessUserActiveSessionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessUserActiveSessionListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/users/%s/active_sessions", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get an active session for a single user.
func (r *AccessUserActiveSessionService) Get(ctx context.Context, identifier string, id string, nonce string, opts ...option.RequestOption) (res *AccessUserActiveSessionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessUserActiveSessionGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/users/%s/active_sessions/%s", identifier, id, nonce)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessUserActiveSessionListResponse struct {
	Expiration int64                                       `json:"expiration"`
	Metadata   AccessUserActiveSessionListResponseMetadata `json:"metadata"`
	Name       string                                      `json:"name"`
	JSON       accessUserActiveSessionListResponseJSON     `json:"-"`
}

// accessUserActiveSessionListResponseJSON contains the JSON metadata for the
// struct [AccessUserActiveSessionListResponse]
type accessUserActiveSessionListResponseJSON struct {
	Expiration  apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionListResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionListResponseMetadata struct {
	Apps    map[string]AccessUserActiveSessionListResponseMetadataApp `json:"apps"`
	Expires int64                                                     `json:"expires"`
	Iat     int64                                                     `json:"iat"`
	Nonce   string                                                    `json:"nonce"`
	TTL     int64                                                     `json:"ttl"`
	JSON    accessUserActiveSessionListResponseMetadataJSON           `json:"-"`
}

// accessUserActiveSessionListResponseMetadataJSON contains the JSON metadata for
// the struct [AccessUserActiveSessionListResponseMetadata]
type accessUserActiveSessionListResponseMetadataJSON struct {
	Apps        apijson.Field
	Expires     apijson.Field
	Iat         apijson.Field
	Nonce       apijson.Field
	TTL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionListResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionListResponseMetadataApp struct {
	Hostname string                                             `json:"hostname"`
	Name     string                                             `json:"name"`
	Type     string                                             `json:"type"`
	Uid      string                                             `json:"uid"`
	JSON     accessUserActiveSessionListResponseMetadataAppJSON `json:"-"`
}

// accessUserActiveSessionListResponseMetadataAppJSON contains the JSON metadata
// for the struct [AccessUserActiveSessionListResponseMetadataApp]
type accessUserActiveSessionListResponseMetadataAppJSON struct {
	Hostname    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponseMetadataApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionListResponseMetadataAppJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponse struct {
	AccountID          string                                                     `json:"account_id"`
	AuthStatus         string                                                     `json:"auth_status"`
	CommonName         string                                                     `json:"common_name"`
	DeviceID           string                                                     `json:"device_id"`
	DeviceSessions     map[string]AccessUserActiveSessionGetResponseDeviceSession `json:"device_sessions"`
	DevicePosture      map[string]AccessUserActiveSessionGetResponseDevicePosture `json:"devicePosture"`
	Email              string                                                     `json:"email"`
	Geo                AccessUserActiveSessionGetResponseGeo                      `json:"geo"`
	Iat                float64                                                    `json:"iat"`
	Idp                AccessUserActiveSessionGetResponseIdp                      `json:"idp"`
	IP                 string                                                     `json:"ip"`
	IsGateway          bool                                                       `json:"is_gateway"`
	IsWARP             bool                                                       `json:"is_warp"`
	IsActive           bool                                                       `json:"isActive"`
	MTLSAuth           AccessUserActiveSessionGetResponseMTLSAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                                                     `json:"service_token_id"`
	ServiceTokenStatus bool                                                       `json:"service_token_status"`
	UserUUID           string                                                     `json:"user_uuid"`
	Version            float64                                                    `json:"version"`
	JSON               accessUserActiveSessionGetResponseJSON                     `json:"-"`
}

// accessUserActiveSessionGetResponseJSON contains the JSON metadata for the struct
// [AccessUserActiveSessionGetResponse]
type accessUserActiveSessionGetResponseJSON struct {
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

func (r *AccessUserActiveSessionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseDeviceSession struct {
	LastAuthenticated float64                                             `json:"last_authenticated"`
	JSON              accessUserActiveSessionGetResponseDeviceSessionJSON `json:"-"`
}

// accessUserActiveSessionGetResponseDeviceSessionJSON contains the JSON metadata
// for the struct [AccessUserActiveSessionGetResponseDeviceSession]
type accessUserActiveSessionGetResponseDeviceSessionJSON struct {
	LastAuthenticated apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseDeviceSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseDeviceSessionJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseDevicePosture struct {
	ID          string                                               `json:"id"`
	Check       AccessUserActiveSessionGetResponseDevicePostureCheck `json:"check"`
	Data        interface{}                                          `json:"data"`
	Description string                                               `json:"description"`
	Error       string                                               `json:"error"`
	RuleName    string                                               `json:"rule_name"`
	Success     bool                                                 `json:"success"`
	Timestamp   string                                               `json:"timestamp"`
	Type        string                                               `json:"type"`
	JSON        accessUserActiveSessionGetResponseDevicePostureJSON  `json:"-"`
}

// accessUserActiveSessionGetResponseDevicePostureJSON contains the JSON metadata
// for the struct [AccessUserActiveSessionGetResponseDevicePosture]
type accessUserActiveSessionGetResponseDevicePostureJSON struct {
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

func (r *AccessUserActiveSessionGetResponseDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseDevicePostureJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseDevicePostureCheck struct {
	Exists bool                                                     `json:"exists"`
	Path   string                                                   `json:"path"`
	JSON   accessUserActiveSessionGetResponseDevicePostureCheckJSON `json:"-"`
}

// accessUserActiveSessionGetResponseDevicePostureCheckJSON contains the JSON
// metadata for the struct [AccessUserActiveSessionGetResponseDevicePostureCheck]
type accessUserActiveSessionGetResponseDevicePostureCheckJSON struct {
	Exists      apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseDevicePostureCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseDevicePostureCheckJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseGeo struct {
	Country string                                    `json:"country"`
	JSON    accessUserActiveSessionGetResponseGeoJSON `json:"-"`
}

// accessUserActiveSessionGetResponseGeoJSON contains the JSON metadata for the
// struct [AccessUserActiveSessionGetResponseGeo]
type accessUserActiveSessionGetResponseGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseGeoJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseIdp struct {
	ID   string                                    `json:"id"`
	Type string                                    `json:"type"`
	JSON accessUserActiveSessionGetResponseIdpJSON `json:"-"`
}

// accessUserActiveSessionGetResponseIdpJSON contains the JSON metadata for the
// struct [AccessUserActiveSessionGetResponseIdp]
type accessUserActiveSessionGetResponseIdpJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseIdp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseIdpJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseMTLSAuth struct {
	AuthStatus    string                                         `json:"auth_status"`
	CertIssuerDn  string                                         `json:"cert_issuer_dn"`
	CertIssuerSki string                                         `json:"cert_issuer_ski"`
	CertPresented bool                                           `json:"cert_presented"`
	CertSerial    string                                         `json:"cert_serial"`
	JSON          accessUserActiveSessionGetResponseMTLSAuthJSON `json:"-"`
}

// accessUserActiveSessionGetResponseMTLSAuthJSON contains the JSON metadata for
// the struct [AccessUserActiveSessionGetResponseMTLSAuth]
type accessUserActiveSessionGetResponseMTLSAuthJSON struct {
	AuthStatus    apijson.Field
	CertIssuerDn  apijson.Field
	CertIssuerSki apijson.Field
	CertPresented apijson.Field
	CertSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseMTLSAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseMTLSAuthJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionListResponseEnvelope struct {
	Errors   []AccessUserActiveSessionListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessUserActiveSessionListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessUserActiveSessionListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessUserActiveSessionListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessUserActiveSessionListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessUserActiveSessionListResponseEnvelopeJSON       `json:"-"`
}

// accessUserActiveSessionListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessUserActiveSessionListResponseEnvelope]
type accessUserActiveSessionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionListResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accessUserActiveSessionListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessUserActiveSessionListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessUserActiveSessionListResponseEnvelopeErrors]
type accessUserActiveSessionListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionListResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accessUserActiveSessionListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessUserActiveSessionListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessUserActiveSessionListResponseEnvelopeMessages]
type accessUserActiveSessionListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessUserActiveSessionListResponseEnvelopeSuccess bool

const (
	AccessUserActiveSessionListResponseEnvelopeSuccessTrue AccessUserActiveSessionListResponseEnvelopeSuccess = true
)

type AccessUserActiveSessionListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                   `json:"total_count"`
	JSON       accessUserActiveSessionListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessUserActiveSessionListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AccessUserActiveSessionListResponseEnvelopeResultInfo]
type accessUserActiveSessionListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseEnvelope struct {
	Errors   []AccessUserActiveSessionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessUserActiveSessionGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessUserActiveSessionGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessUserActiveSessionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessUserActiveSessionGetResponseEnvelopeJSON    `json:"-"`
}

// accessUserActiveSessionGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessUserActiveSessionGetResponseEnvelope]
type accessUserActiveSessionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accessUserActiveSessionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessUserActiveSessionGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessUserActiveSessionGetResponseEnvelopeErrors]
type accessUserActiveSessionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accessUserActiveSessionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessUserActiveSessionGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessUserActiveSessionGetResponseEnvelopeMessages]
type accessUserActiveSessionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessUserActiveSessionGetResponseEnvelopeSuccess bool

const (
	AccessUserActiveSessionGetResponseEnvelopeSuccessTrue AccessUserActiveSessionGetResponseEnvelopeSuccess = true
)

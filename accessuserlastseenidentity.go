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
func (r *AccessUserLastSeenIdentityService) List(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *AccessUserLastSeenIdentityListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessUserLastSeenIdentityListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/users/%s/last_seen_identity", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessUserLastSeenIdentityListResponse struct {
	AccountID          string                                                         `json:"account_id"`
	AuthStatus         string                                                         `json:"auth_status"`
	CommonName         string                                                         `json:"common_name"`
	DeviceID           string                                                         `json:"device_id"`
	DeviceSessions     map[string]AccessUserLastSeenIdentityListResponseDeviceSession `json:"device_sessions"`
	DevicePosture      map[string]AccessUserLastSeenIdentityListResponseDevicePosture `json:"devicePosture"`
	Email              string                                                         `json:"email"`
	Geo                AccessUserLastSeenIdentityListResponseGeo                      `json:"geo"`
	Iat                float64                                                        `json:"iat"`
	Idp                AccessUserLastSeenIdentityListResponseIdp                      `json:"idp"`
	IP                 string                                                         `json:"ip"`
	IsGateway          bool                                                           `json:"is_gateway"`
	IsWARP             bool                                                           `json:"is_warp"`
	MTLSAuth           AccessUserLastSeenIdentityListResponseMTLSAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                                                         `json:"service_token_id"`
	ServiceTokenStatus bool                                                           `json:"service_token_status"`
	UserUUID           string                                                         `json:"user_uuid"`
	Version            float64                                                        `json:"version"`
	JSON               accessUserLastSeenIdentityListResponseJSON                     `json:"-"`
}

// accessUserLastSeenIdentityListResponseJSON contains the JSON metadata for the
// struct [AccessUserLastSeenIdentityListResponse]
type accessUserLastSeenIdentityListResponseJSON struct {
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
	MTLSAuth           apijson.Field
	ServiceTokenID     apijson.Field
	ServiceTokenStatus apijson.Field
	UserUUID           apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityListResponseDeviceSession struct {
	LastAuthenticated float64                                                 `json:"last_authenticated"`
	JSON              accessUserLastSeenIdentityListResponseDeviceSessionJSON `json:"-"`
}

// accessUserLastSeenIdentityListResponseDeviceSessionJSON contains the JSON
// metadata for the struct [AccessUserLastSeenIdentityListResponseDeviceSession]
type accessUserLastSeenIdentityListResponseDeviceSessionJSON struct {
	LastAuthenticated apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityListResponseDeviceSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityListResponseDevicePosture struct {
	ID          string                                                   `json:"id"`
	Check       AccessUserLastSeenIdentityListResponseDevicePostureCheck `json:"check"`
	Data        interface{}                                              `json:"data"`
	Description string                                                   `json:"description"`
	Error       string                                                   `json:"error"`
	RuleName    string                                                   `json:"rule_name"`
	Success     bool                                                     `json:"success"`
	Timestamp   string                                                   `json:"timestamp"`
	Type        string                                                   `json:"type"`
	JSON        accessUserLastSeenIdentityListResponseDevicePostureJSON  `json:"-"`
}

// accessUserLastSeenIdentityListResponseDevicePostureJSON contains the JSON
// metadata for the struct [AccessUserLastSeenIdentityListResponseDevicePosture]
type accessUserLastSeenIdentityListResponseDevicePostureJSON struct {
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

func (r *AccessUserLastSeenIdentityListResponseDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityListResponseDevicePostureCheck struct {
	Exists bool                                                         `json:"exists"`
	Path   string                                                       `json:"path"`
	JSON   accessUserLastSeenIdentityListResponseDevicePostureCheckJSON `json:"-"`
}

// accessUserLastSeenIdentityListResponseDevicePostureCheckJSON contains the JSON
// metadata for the struct
// [AccessUserLastSeenIdentityListResponseDevicePostureCheck]
type accessUserLastSeenIdentityListResponseDevicePostureCheckJSON struct {
	Exists      apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityListResponseDevicePostureCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityListResponseGeo struct {
	Country string                                        `json:"country"`
	JSON    accessUserLastSeenIdentityListResponseGeoJSON `json:"-"`
}

// accessUserLastSeenIdentityListResponseGeoJSON contains the JSON metadata for the
// struct [AccessUserLastSeenIdentityListResponseGeo]
type accessUserLastSeenIdentityListResponseGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityListResponseGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityListResponseIdp struct {
	ID   string                                        `json:"id"`
	Type string                                        `json:"type"`
	JSON accessUserLastSeenIdentityListResponseIdpJSON `json:"-"`
}

// accessUserLastSeenIdentityListResponseIdpJSON contains the JSON metadata for the
// struct [AccessUserLastSeenIdentityListResponseIdp]
type accessUserLastSeenIdentityListResponseIdpJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityListResponseIdp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityListResponseMTLSAuth struct {
	AuthStatus    string                                             `json:"auth_status"`
	CertIssuerDn  string                                             `json:"cert_issuer_dn"`
	CertIssuerSki string                                             `json:"cert_issuer_ski"`
	CertPresented bool                                               `json:"cert_presented"`
	CertSerial    string                                             `json:"cert_serial"`
	JSON          accessUserLastSeenIdentityListResponseMTLSAuthJSON `json:"-"`
}

// accessUserLastSeenIdentityListResponseMTLSAuthJSON contains the JSON metadata
// for the struct [AccessUserLastSeenIdentityListResponseMTLSAuth]
type accessUserLastSeenIdentityListResponseMTLSAuthJSON struct {
	AuthStatus    apijson.Field
	CertIssuerDn  apijson.Field
	CertIssuerSki apijson.Field
	CertPresented apijson.Field
	CertSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityListResponseMTLSAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityListResponseEnvelope struct {
	Errors   []AccessUserLastSeenIdentityListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessUserLastSeenIdentityListResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessUserLastSeenIdentityListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessUserLastSeenIdentityListResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessUserLastSeenIdentityListResponseEnvelopeJSON    `json:"-"`
}

// accessUserLastSeenIdentityListResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessUserLastSeenIdentityListResponseEnvelope]
type accessUserLastSeenIdentityListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityListResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accessUserLastSeenIdentityListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessUserLastSeenIdentityListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessUserLastSeenIdentityListResponseEnvelopeErrors]
type accessUserLastSeenIdentityListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityListResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accessUserLastSeenIdentityListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessUserLastSeenIdentityListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessUserLastSeenIdentityListResponseEnvelopeMessages]
type accessUserLastSeenIdentityListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessUserLastSeenIdentityListResponseEnvelopeSuccess bool

const (
	AccessUserLastSeenIdentityListResponseEnvelopeSuccessTrue AccessUserLastSeenIdentityListResponseEnvelopeSuccess = true
)

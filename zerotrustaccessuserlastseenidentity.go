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

// ZeroTrustAccessUserLastSeenIdentityService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustAccessUserLastSeenIdentityService] method instead.
type ZeroTrustAccessUserLastSeenIdentityService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessUserLastSeenIdentityService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZeroTrustAccessUserLastSeenIdentityService(opts ...option.RequestOption) (r *ZeroTrustAccessUserLastSeenIdentityService) {
	r = &ZeroTrustAccessUserLastSeenIdentityService{}
	r.Options = opts
	return
}

// Get last seen identity for a single user.
func (r *ZeroTrustAccessUserLastSeenIdentityService) Get(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *ZeroTrustAccessUserLastSeenIdentityGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/users/%s/last_seen_identity", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustAccessUserLastSeenIdentityGetResponse struct {
	AccountID          string                                                                 `json:"account_id"`
	AuthStatus         string                                                                 `json:"auth_status"`
	CommonName         string                                                                 `json:"common_name"`
	DeviceID           string                                                                 `json:"device_id"`
	DeviceSessions     map[string]ZeroTrustAccessUserLastSeenIdentityGetResponseDeviceSession `json:"device_sessions"`
	DevicePosture      map[string]ZeroTrustAccessUserLastSeenIdentityGetResponseDevicePosture `json:"devicePosture"`
	Email              string                                                                 `json:"email"`
	Geo                ZeroTrustAccessUserLastSeenIdentityGetResponseGeo                      `json:"geo"`
	Iat                float64                                                                `json:"iat"`
	Idp                ZeroTrustAccessUserLastSeenIdentityGetResponseIdp                      `json:"idp"`
	IP                 string                                                                 `json:"ip"`
	IsGateway          bool                                                                   `json:"is_gateway"`
	IsWARP             bool                                                                   `json:"is_warp"`
	MTLSAuth           ZeroTrustAccessUserLastSeenIdentityGetResponseMTLSAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                                                                 `json:"service_token_id"`
	ServiceTokenStatus bool                                                                   `json:"service_token_status"`
	UserUUID           string                                                                 `json:"user_uuid"`
	Version            float64                                                                `json:"version"`
	JSON               zeroTrustAccessUserLastSeenIdentityGetResponseJSON                     `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityGetResponseJSON contains the JSON metadata
// for the struct [ZeroTrustAccessUserLastSeenIdentityGetResponse]
type zeroTrustAccessUserLastSeenIdentityGetResponseJSON struct {
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

func (r *ZeroTrustAccessUserLastSeenIdentityGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserLastSeenIdentityGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserLastSeenIdentityGetResponseDeviceSession struct {
	LastAuthenticated float64                                                         `json:"last_authenticated"`
	JSON              zeroTrustAccessUserLastSeenIdentityGetResponseDeviceSessionJSON `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityGetResponseDeviceSessionJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessUserLastSeenIdentityGetResponseDeviceSession]
type zeroTrustAccessUserLastSeenIdentityGetResponseDeviceSessionJSON struct {
	LastAuthenticated apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityGetResponseDeviceSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserLastSeenIdentityGetResponseDeviceSessionJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserLastSeenIdentityGetResponseDevicePosture struct {
	ID          string                                                           `json:"id"`
	Check       ZeroTrustAccessUserLastSeenIdentityGetResponseDevicePostureCheck `json:"check"`
	Data        interface{}                                                      `json:"data"`
	Description string                                                           `json:"description"`
	Error       string                                                           `json:"error"`
	RuleName    string                                                           `json:"rule_name"`
	Success     bool                                                             `json:"success"`
	Timestamp   string                                                           `json:"timestamp"`
	Type        string                                                           `json:"type"`
	JSON        zeroTrustAccessUserLastSeenIdentityGetResponseDevicePostureJSON  `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityGetResponseDevicePostureJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessUserLastSeenIdentityGetResponseDevicePosture]
type zeroTrustAccessUserLastSeenIdentityGetResponseDevicePostureJSON struct {
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

func (r *ZeroTrustAccessUserLastSeenIdentityGetResponseDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserLastSeenIdentityGetResponseDevicePostureJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserLastSeenIdentityGetResponseDevicePostureCheck struct {
	Exists bool                                                                 `json:"exists"`
	Path   string                                                               `json:"path"`
	JSON   zeroTrustAccessUserLastSeenIdentityGetResponseDevicePostureCheckJSON `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityGetResponseDevicePostureCheckJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessUserLastSeenIdentityGetResponseDevicePostureCheck]
type zeroTrustAccessUserLastSeenIdentityGetResponseDevicePostureCheckJSON struct {
	Exists      apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityGetResponseDevicePostureCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserLastSeenIdentityGetResponseDevicePostureCheckJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserLastSeenIdentityGetResponseGeo struct {
	Country string                                                `json:"country"`
	JSON    zeroTrustAccessUserLastSeenIdentityGetResponseGeoJSON `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityGetResponseGeoJSON contains the JSON metadata
// for the struct [ZeroTrustAccessUserLastSeenIdentityGetResponseGeo]
type zeroTrustAccessUserLastSeenIdentityGetResponseGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityGetResponseGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserLastSeenIdentityGetResponseGeoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserLastSeenIdentityGetResponseIdp struct {
	ID   string                                                `json:"id"`
	Type string                                                `json:"type"`
	JSON zeroTrustAccessUserLastSeenIdentityGetResponseIdpJSON `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityGetResponseIdpJSON contains the JSON metadata
// for the struct [ZeroTrustAccessUserLastSeenIdentityGetResponseIdp]
type zeroTrustAccessUserLastSeenIdentityGetResponseIdpJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityGetResponseIdp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserLastSeenIdentityGetResponseIdpJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserLastSeenIdentityGetResponseMTLSAuth struct {
	AuthStatus    string                                                     `json:"auth_status"`
	CertIssuerDn  string                                                     `json:"cert_issuer_dn"`
	CertIssuerSki string                                                     `json:"cert_issuer_ski"`
	CertPresented bool                                                       `json:"cert_presented"`
	CertSerial    string                                                     `json:"cert_serial"`
	JSON          zeroTrustAccessUserLastSeenIdentityGetResponseMTLSAuthJSON `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityGetResponseMTLSAuthJSON contains the JSON
// metadata for the struct [ZeroTrustAccessUserLastSeenIdentityGetResponseMTLSAuth]
type zeroTrustAccessUserLastSeenIdentityGetResponseMTLSAuthJSON struct {
	AuthStatus    apijson.Field
	CertIssuerDn  apijson.Field
	CertIssuerSki apijson.Field
	CertPresented apijson.Field
	CertSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityGetResponseMTLSAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserLastSeenIdentityGetResponseMTLSAuthJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelope struct {
	Errors   []ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessUserLastSeenIdentityGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelope]
type zeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeErrors]
type zeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    zeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeMessages]
type zeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeSuccessTrue ZeroTrustAccessUserLastSeenIdentityGetResponseEnvelopeSuccess = true
)

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
func (r *ZeroTrustAccessUserLastSeenIdentityService) List(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *ZeroTrustAccessUserLastSeenIdentityListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessUserLastSeenIdentityListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/users/%s/last_seen_identity", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustAccessUserLastSeenIdentityListResponse struct {
	AccountID          string                                                                  `json:"account_id"`
	AuthStatus         string                                                                  `json:"auth_status"`
	CommonName         string                                                                  `json:"common_name"`
	DeviceID           string                                                                  `json:"device_id"`
	DeviceSessions     map[string]ZeroTrustAccessUserLastSeenIdentityListResponseDeviceSession `json:"device_sessions"`
	DevicePosture      map[string]ZeroTrustAccessUserLastSeenIdentityListResponseDevicePosture `json:"devicePosture"`
	Email              string                                                                  `json:"email"`
	Geo                ZeroTrustAccessUserLastSeenIdentityListResponseGeo                      `json:"geo"`
	Iat                float64                                                                 `json:"iat"`
	Idp                ZeroTrustAccessUserLastSeenIdentityListResponseIdp                      `json:"idp"`
	IP                 string                                                                  `json:"ip"`
	IsGateway          bool                                                                    `json:"is_gateway"`
	IsWARP             bool                                                                    `json:"is_warp"`
	MTLSAuth           ZeroTrustAccessUserLastSeenIdentityListResponseMTLSAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                                                                  `json:"service_token_id"`
	ServiceTokenStatus bool                                                                    `json:"service_token_status"`
	UserUUID           string                                                                  `json:"user_uuid"`
	Version            float64                                                                 `json:"version"`
	JSON               zeroTrustAccessUserLastSeenIdentityListResponseJSON                     `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityListResponseJSON contains the JSON metadata
// for the struct [ZeroTrustAccessUserLastSeenIdentityListResponse]
type zeroTrustAccessUserLastSeenIdentityListResponseJSON struct {
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

func (r *ZeroTrustAccessUserLastSeenIdentityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessUserLastSeenIdentityListResponseDeviceSession struct {
	LastAuthenticated float64                                                          `json:"last_authenticated"`
	JSON              zeroTrustAccessUserLastSeenIdentityListResponseDeviceSessionJSON `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityListResponseDeviceSessionJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessUserLastSeenIdentityListResponseDeviceSession]
type zeroTrustAccessUserLastSeenIdentityListResponseDeviceSessionJSON struct {
	LastAuthenticated apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityListResponseDeviceSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessUserLastSeenIdentityListResponseDevicePosture struct {
	ID          string                                                            `json:"id"`
	Check       ZeroTrustAccessUserLastSeenIdentityListResponseDevicePostureCheck `json:"check"`
	Data        interface{}                                                       `json:"data"`
	Description string                                                            `json:"description"`
	Error       string                                                            `json:"error"`
	RuleName    string                                                            `json:"rule_name"`
	Success     bool                                                              `json:"success"`
	Timestamp   string                                                            `json:"timestamp"`
	Type        string                                                            `json:"type"`
	JSON        zeroTrustAccessUserLastSeenIdentityListResponseDevicePostureJSON  `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityListResponseDevicePostureJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessUserLastSeenIdentityListResponseDevicePosture]
type zeroTrustAccessUserLastSeenIdentityListResponseDevicePostureJSON struct {
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

func (r *ZeroTrustAccessUserLastSeenIdentityListResponseDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessUserLastSeenIdentityListResponseDevicePostureCheck struct {
	Exists bool                                                                  `json:"exists"`
	Path   string                                                                `json:"path"`
	JSON   zeroTrustAccessUserLastSeenIdentityListResponseDevicePostureCheckJSON `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityListResponseDevicePostureCheckJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessUserLastSeenIdentityListResponseDevicePostureCheck]
type zeroTrustAccessUserLastSeenIdentityListResponseDevicePostureCheckJSON struct {
	Exists      apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityListResponseDevicePostureCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessUserLastSeenIdentityListResponseGeo struct {
	Country string                                                 `json:"country"`
	JSON    zeroTrustAccessUserLastSeenIdentityListResponseGeoJSON `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityListResponseGeoJSON contains the JSON
// metadata for the struct [ZeroTrustAccessUserLastSeenIdentityListResponseGeo]
type zeroTrustAccessUserLastSeenIdentityListResponseGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityListResponseGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessUserLastSeenIdentityListResponseIdp struct {
	ID   string                                                 `json:"id"`
	Type string                                                 `json:"type"`
	JSON zeroTrustAccessUserLastSeenIdentityListResponseIdpJSON `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityListResponseIdpJSON contains the JSON
// metadata for the struct [ZeroTrustAccessUserLastSeenIdentityListResponseIdp]
type zeroTrustAccessUserLastSeenIdentityListResponseIdpJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityListResponseIdp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessUserLastSeenIdentityListResponseMTLSAuth struct {
	AuthStatus    string                                                      `json:"auth_status"`
	CertIssuerDn  string                                                      `json:"cert_issuer_dn"`
	CertIssuerSki string                                                      `json:"cert_issuer_ski"`
	CertPresented bool                                                        `json:"cert_presented"`
	CertSerial    string                                                      `json:"cert_serial"`
	JSON          zeroTrustAccessUserLastSeenIdentityListResponseMTLSAuthJSON `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityListResponseMTLSAuthJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessUserLastSeenIdentityListResponseMTLSAuth]
type zeroTrustAccessUserLastSeenIdentityListResponseMTLSAuthJSON struct {
	AuthStatus    apijson.Field
	CertIssuerDn  apijson.Field
	CertIssuerSki apijson.Field
	CertPresented apijson.Field
	CertSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityListResponseMTLSAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessUserLastSeenIdentityListResponseEnvelope struct {
	Errors   []ZeroTrustAccessUserLastSeenIdentityListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessUserLastSeenIdentityListResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessUserLastSeenIdentityListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessUserLastSeenIdentityListResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessUserLastSeenIdentityListResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityListResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessUserLastSeenIdentityListResponseEnvelope]
type zeroTrustAccessUserLastSeenIdentityListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessUserLastSeenIdentityListResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zeroTrustAccessUserLastSeenIdentityListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityListResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessUserLastSeenIdentityListResponseEnvelopeErrors]
type zeroTrustAccessUserLastSeenIdentityListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessUserLastSeenIdentityListResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zeroTrustAccessUserLastSeenIdentityListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessUserLastSeenIdentityListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessUserLastSeenIdentityListResponseEnvelopeMessages]
type zeroTrustAccessUserLastSeenIdentityListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserLastSeenIdentityListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessUserLastSeenIdentityListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessUserLastSeenIdentityListResponseEnvelopeSuccessTrue ZeroTrustAccessUserLastSeenIdentityListResponseEnvelopeSuccess = true
)

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
	var env AccessUserLastSeenIdentityGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/users/%s/last_seen_identity", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessUserLastSeenIdentityGetResponse struct {
	AccountID          string                                                        `json:"account_id"`
	AuthStatus         string                                                        `json:"auth_status"`
	CommonName         string                                                        `json:"common_name"`
	DeviceID           string                                                        `json:"device_id"`
	DeviceSessions     map[string]AccessUserLastSeenIdentityGetResponseDeviceSession `json:"device_sessions"`
	DevicePosture      map[string]AccessUserLastSeenIdentityGetResponseDevicePosture `json:"devicePosture"`
	Email              string                                                        `json:"email"`
	Geo                AccessUserLastSeenIdentityGetResponseGeo                      `json:"geo"`
	Iat                float64                                                       `json:"iat"`
	Idp                AccessUserLastSeenIdentityGetResponseIdp                      `json:"idp"`
	IP                 string                                                        `json:"ip"`
	IsGateway          bool                                                          `json:"is_gateway"`
	IsWarp             bool                                                          `json:"is_warp"`
	MtlsAuth           AccessUserLastSeenIdentityGetResponseMtlsAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                                                        `json:"service_token_id"`
	ServiceTokenStatus bool                                                          `json:"service_token_status"`
	UserUuid           string                                                        `json:"user_uuid"`
	Version            float64                                                       `json:"version"`
	JSON               accessUserLastSeenIdentityGetResponseJSON                     `json:"-"`
}

// accessUserLastSeenIdentityGetResponseJSON contains the JSON metadata for the
// struct [AccessUserLastSeenIdentityGetResponse]
type accessUserLastSeenIdentityGetResponseJSON struct {
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

func (r *AccessUserLastSeenIdentityGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityGetResponseDeviceSession struct {
	LastAuthenticated float64                                                `json:"last_authenticated"`
	JSON              accessUserLastSeenIdentityGetResponseDeviceSessionJSON `json:"-"`
}

// accessUserLastSeenIdentityGetResponseDeviceSessionJSON contains the JSON
// metadata for the struct [AccessUserLastSeenIdentityGetResponseDeviceSession]
type accessUserLastSeenIdentityGetResponseDeviceSessionJSON struct {
	LastAuthenticated apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseDeviceSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityGetResponseDevicePosture struct {
	ID          string                                                  `json:"id"`
	Check       AccessUserLastSeenIdentityGetResponseDevicePostureCheck `json:"check"`
	Data        interface{}                                             `json:"data"`
	Description string                                                  `json:"description"`
	Error       string                                                  `json:"error"`
	RuleName    string                                                  `json:"rule_name"`
	Success     bool                                                    `json:"success"`
	Timestamp   string                                                  `json:"timestamp"`
	Type        string                                                  `json:"type"`
	JSON        accessUserLastSeenIdentityGetResponseDevicePostureJSON  `json:"-"`
}

// accessUserLastSeenIdentityGetResponseDevicePostureJSON contains the JSON
// metadata for the struct [AccessUserLastSeenIdentityGetResponseDevicePosture]
type accessUserLastSeenIdentityGetResponseDevicePostureJSON struct {
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

func (r *AccessUserLastSeenIdentityGetResponseDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityGetResponseDevicePostureCheck struct {
	Exists bool                                                        `json:"exists"`
	Path   string                                                      `json:"path"`
	JSON   accessUserLastSeenIdentityGetResponseDevicePostureCheckJSON `json:"-"`
}

// accessUserLastSeenIdentityGetResponseDevicePostureCheckJSON contains the JSON
// metadata for the struct
// [AccessUserLastSeenIdentityGetResponseDevicePostureCheck]
type accessUserLastSeenIdentityGetResponseDevicePostureCheckJSON struct {
	Exists      apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseDevicePostureCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityGetResponseGeo struct {
	Country string                                       `json:"country"`
	JSON    accessUserLastSeenIdentityGetResponseGeoJSON `json:"-"`
}

// accessUserLastSeenIdentityGetResponseGeoJSON contains the JSON metadata for the
// struct [AccessUserLastSeenIdentityGetResponseGeo]
type accessUserLastSeenIdentityGetResponseGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityGetResponseIdp struct {
	ID   string                                       `json:"id"`
	Type string                                       `json:"type"`
	JSON accessUserLastSeenIdentityGetResponseIdpJSON `json:"-"`
}

// accessUserLastSeenIdentityGetResponseIdpJSON contains the JSON metadata for the
// struct [AccessUserLastSeenIdentityGetResponseIdp]
type accessUserLastSeenIdentityGetResponseIdpJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseIdp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityGetResponseMtlsAuth struct {
	AuthStatus    string                                            `json:"auth_status"`
	CertIssuerDn  string                                            `json:"cert_issuer_dn"`
	CertIssuerSki string                                            `json:"cert_issuer_ski"`
	CertPresented bool                                              `json:"cert_presented"`
	CertSerial    string                                            `json:"cert_serial"`
	JSON          accessUserLastSeenIdentityGetResponseMtlsAuthJSON `json:"-"`
}

// accessUserLastSeenIdentityGetResponseMtlsAuthJSON contains the JSON metadata for
// the struct [AccessUserLastSeenIdentityGetResponseMtlsAuth]
type accessUserLastSeenIdentityGetResponseMtlsAuthJSON struct {
	AuthStatus    apijson.Field
	CertIssuerDn  apijson.Field
	CertIssuerSki apijson.Field
	CertPresented apijson.Field
	CertSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseMtlsAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityGetResponseEnvelope struct {
	Errors   []AccessUserLastSeenIdentityGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessUserLastSeenIdentityGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessUserLastSeenIdentityGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessUserLastSeenIdentityGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessUserLastSeenIdentityGetResponseEnvelopeJSON    `json:"-"`
}

// accessUserLastSeenIdentityGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessUserLastSeenIdentityGetResponseEnvelope]
type accessUserLastSeenIdentityGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accessUserLastSeenIdentityGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessUserLastSeenIdentityGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessUserLastSeenIdentityGetResponseEnvelopeErrors]
type accessUserLastSeenIdentityGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserLastSeenIdentityGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    accessUserLastSeenIdentityGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessUserLastSeenIdentityGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessUserLastSeenIdentityGetResponseEnvelopeMessages]
type accessUserLastSeenIdentityGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessUserLastSeenIdentityGetResponseEnvelopeSuccess bool

const (
	AccessUserLastSeenIdentityGetResponseEnvelopeSuccessTrue AccessUserLastSeenIdentityGetResponseEnvelopeSuccess = true
)

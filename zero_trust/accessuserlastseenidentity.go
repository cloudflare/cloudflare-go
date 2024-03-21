// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *AccessUserLastSeenIdentityService) Get(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *AccessIdentity, err error) {
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

type AccessIdentity struct {
	AccountID          string                                 `json:"account_id"`
	AuthStatus         string                                 `json:"auth_status"`
	CommonName         string                                 `json:"common_name"`
	DeviceID           string                                 `json:"device_id"`
	DeviceSessions     map[string]AccessIdentityDeviceSession `json:"device_sessions"`
	DevicePosture      map[string]AccessIdentityDevicePosture `json:"devicePosture"`
	Email              string                                 `json:"email"`
	Geo                AccessIdentityGeo                      `json:"geo"`
	Iat                float64                                `json:"iat"`
	IDP                AccessIdentityIDP                      `json:"idp"`
	IP                 string                                 `json:"ip"`
	IsGateway          bool                                   `json:"is_gateway"`
	IsWARP             bool                                   `json:"is_warp"`
	MTLSAuth           AccessIdentityMTLSAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                                 `json:"service_token_id"`
	ServiceTokenStatus bool                                   `json:"service_token_status"`
	UserUUID           string                                 `json:"user_uuid"`
	Version            float64                                `json:"version"`
	JSON               accessIdentityJSON                     `json:"-"`
}

// accessIdentityJSON contains the JSON metadata for the struct [AccessIdentity]
type accessIdentityJSON struct {
	AccountID          apijson.Field
	AuthStatus         apijson.Field
	CommonName         apijson.Field
	DeviceID           apijson.Field
	DeviceSessions     apijson.Field
	DevicePosture      apijson.Field
	Email              apijson.Field
	Geo                apijson.Field
	Iat                apijson.Field
	IDP                apijson.Field
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

func (r *AccessIdentity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdentityJSON) RawJSON() string {
	return r.raw
}

type AccessIdentityDeviceSession struct {
	LastAuthenticated float64                         `json:"last_authenticated"`
	JSON              accessIdentityDeviceSessionJSON `json:"-"`
}

// accessIdentityDeviceSessionJSON contains the JSON metadata for the struct
// [AccessIdentityDeviceSession]
type accessIdentityDeviceSessionJSON struct {
	LastAuthenticated apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessIdentityDeviceSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdentityDeviceSessionJSON) RawJSON() string {
	return r.raw
}

type AccessIdentityDevicePosture struct {
	ID          string                           `json:"id"`
	Check       AccessIdentityDevicePostureCheck `json:"check"`
	Data        interface{}                      `json:"data"`
	Description string                           `json:"description"`
	Error       string                           `json:"error"`
	RuleName    string                           `json:"rule_name"`
	Success     bool                             `json:"success"`
	Timestamp   string                           `json:"timestamp"`
	Type        string                           `json:"type"`
	JSON        accessIdentityDevicePostureJSON  `json:"-"`
}

// accessIdentityDevicePostureJSON contains the JSON metadata for the struct
// [AccessIdentityDevicePosture]
type accessIdentityDevicePostureJSON struct {
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

func (r *AccessIdentityDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdentityDevicePostureJSON) RawJSON() string {
	return r.raw
}

type AccessIdentityDevicePostureCheck struct {
	Exists bool                                 `json:"exists"`
	Path   string                               `json:"path"`
	JSON   accessIdentityDevicePostureCheckJSON `json:"-"`
}

// accessIdentityDevicePostureCheckJSON contains the JSON metadata for the struct
// [AccessIdentityDevicePostureCheck]
type accessIdentityDevicePostureCheckJSON struct {
	Exists      apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityDevicePostureCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdentityDevicePostureCheckJSON) RawJSON() string {
	return r.raw
}

type AccessIdentityGeo struct {
	Country string                `json:"country"`
	JSON    accessIdentityGeoJSON `json:"-"`
}

// accessIdentityGeoJSON contains the JSON metadata for the struct
// [AccessIdentityGeo]
type accessIdentityGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdentityGeoJSON) RawJSON() string {
	return r.raw
}

type AccessIdentityIDP struct {
	ID   string                `json:"id"`
	Type string                `json:"type"`
	JSON accessIdentityIDPJSON `json:"-"`
}

// accessIdentityIDPJSON contains the JSON metadata for the struct
// [AccessIdentityIDP]
type accessIdentityIDPJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityIDP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdentityIDPJSON) RawJSON() string {
	return r.raw
}

type AccessIdentityMTLSAuth struct {
	AuthStatus    string                     `json:"auth_status"`
	CERTIssuerDn  string                     `json:"cert_issuer_dn"`
	CERTIssuerSki string                     `json:"cert_issuer_ski"`
	CERTPresented bool                       `json:"cert_presented"`
	CERTSerial    string                     `json:"cert_serial"`
	JSON          accessIdentityMTLSAuthJSON `json:"-"`
}

// accessIdentityMTLSAuthJSON contains the JSON metadata for the struct
// [AccessIdentityMTLSAuth]
type accessIdentityMTLSAuthJSON struct {
	AuthStatus    apijson.Field
	CERTIssuerDn  apijson.Field
	CERTIssuerSki apijson.Field
	CERTPresented apijson.Field
	CERTSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessIdentityMTLSAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdentityMTLSAuthJSON) RawJSON() string {
	return r.raw
}

type AccessUserLastSeenIdentityGetResponseEnvelope struct {
	Errors   []AccessUserLastSeenIdentityGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessUserLastSeenIdentityGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessIdentity                                          `json:"result,required"`
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

func (r accessUserLastSeenIdentityGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accessUserLastSeenIdentityGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accessUserLastSeenIdentityGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessUserLastSeenIdentityGetResponseEnvelopeSuccess bool

const (
	AccessUserLastSeenIdentityGetResponseEnvelopeSuccessTrue AccessUserLastSeenIdentityGetResponseEnvelopeSuccess = true
)

func (r AccessUserLastSeenIdentityGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserLastSeenIdentityGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

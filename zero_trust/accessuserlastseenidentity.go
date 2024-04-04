// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
func (r *AccessUserLastSeenIdentityService) Get(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *ZeroTrustIdentity, err error) {
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

type ZeroTrustIdentity struct {
	AccountID          string                                    `json:"account_id"`
	AuthStatus         string                                    `json:"auth_status"`
	CommonName         string                                    `json:"common_name"`
	DeviceID           string                                    `json:"device_id"`
	DeviceSessions     map[string]ZeroTrustIdentityDeviceSession `json:"device_sessions"`
	DevicePosture      map[string]ZeroTrustIdentityDevicePosture `json:"devicePosture"`
	Email              string                                    `json:"email"`
	Geo                ZeroTrustIdentityGeo                      `json:"geo"`
	Iat                float64                                   `json:"iat"`
	IDP                ZeroTrustIdentityIDP                      `json:"idp"`
	IP                 string                                    `json:"ip"`
	IsGateway          bool                                      `json:"is_gateway"`
	IsWARP             bool                                      `json:"is_warp"`
	MTLSAuth           ZeroTrustIdentityMTLSAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                                    `json:"service_token_id"`
	ServiceTokenStatus bool                                      `json:"service_token_status"`
	UserUUID           string                                    `json:"user_uuid"`
	Version            float64                                   `json:"version"`
	JSON               zeroTrustIdentityJSON                     `json:"-"`
}

// zeroTrustIdentityJSON contains the JSON metadata for the struct
// [ZeroTrustIdentity]
type zeroTrustIdentityJSON struct {
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

func (r *ZeroTrustIdentity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityDeviceSession struct {
	LastAuthenticated float64                            `json:"last_authenticated"`
	JSON              zeroTrustIdentityDeviceSessionJSON `json:"-"`
}

// zeroTrustIdentityDeviceSessionJSON contains the JSON metadata for the struct
// [ZeroTrustIdentityDeviceSession]
type zeroTrustIdentityDeviceSessionJSON struct {
	LastAuthenticated apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZeroTrustIdentityDeviceSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityDeviceSessionJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityDevicePosture struct {
	ID          string                              `json:"id"`
	Check       ZeroTrustIdentityDevicePostureCheck `json:"check"`
	Data        interface{}                         `json:"data"`
	Description string                              `json:"description"`
	Error       string                              `json:"error"`
	RuleName    string                              `json:"rule_name"`
	Success     bool                                `json:"success"`
	Timestamp   string                              `json:"timestamp"`
	Type        string                              `json:"type"`
	JSON        zeroTrustIdentityDevicePostureJSON  `json:"-"`
}

// zeroTrustIdentityDevicePostureJSON contains the JSON metadata for the struct
// [ZeroTrustIdentityDevicePosture]
type zeroTrustIdentityDevicePostureJSON struct {
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

func (r *ZeroTrustIdentityDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityDevicePostureJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityDevicePostureCheck struct {
	Exists bool                                    `json:"exists"`
	Path   string                                  `json:"path"`
	JSON   zeroTrustIdentityDevicePostureCheckJSON `json:"-"`
}

// zeroTrustIdentityDevicePostureCheckJSON contains the JSON metadata for the
// struct [ZeroTrustIdentityDevicePostureCheck]
type zeroTrustIdentityDevicePostureCheckJSON struct {
	Exists      apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityDevicePostureCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityDevicePostureCheckJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityGeo struct {
	Country string                   `json:"country"`
	JSON    zeroTrustIdentityGeoJSON `json:"-"`
}

// zeroTrustIdentityGeoJSON contains the JSON metadata for the struct
// [ZeroTrustIdentityGeo]
type zeroTrustIdentityGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityGeoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityIDP struct {
	ID   string                   `json:"id"`
	Type string                   `json:"type"`
	JSON zeroTrustIdentityIDPJSON `json:"-"`
}

// zeroTrustIdentityIDPJSON contains the JSON metadata for the struct
// [ZeroTrustIdentityIDP]
type zeroTrustIdentityIDPJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityIDP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityIDPJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityMTLSAuth struct {
	AuthStatus    string                        `json:"auth_status"`
	CERTIssuerDn  string                        `json:"cert_issuer_dn"`
	CERTIssuerSki string                        `json:"cert_issuer_ski"`
	CERTPresented bool                          `json:"cert_presented"`
	CERTSerial    string                        `json:"cert_serial"`
	JSON          zeroTrustIdentityMTLSAuthJSON `json:"-"`
}

// zeroTrustIdentityMTLSAuthJSON contains the JSON metadata for the struct
// [ZeroTrustIdentityMTLSAuth]
type zeroTrustIdentityMTLSAuthJSON struct {
	AuthStatus    apijson.Field
	CERTIssuerDn  apijson.Field
	CERTIssuerSki apijson.Field
	CERTPresented apijson.Field
	CERTSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustIdentityMTLSAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityMTLSAuthJSON) RawJSON() string {
	return r.raw
}

type AccessUserLastSeenIdentityGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ZeroTrustIdentity     `json:"result,required"`
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

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

// AccessAppUserPolicyCheckService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccessAppUserPolicyCheckService] method instead.
type AccessAppUserPolicyCheckService struct {
	Options []option.RequestOption
}

// NewAccessAppUserPolicyCheckService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessAppUserPolicyCheckService(opts ...option.RequestOption) (r *AccessAppUserPolicyCheckService) {
	r = &AccessAppUserPolicyCheckService{}
	r.Options = opts
	return
}

// Tests if a specific user has permission to access an application.
func (r *AccessAppUserPolicyCheckService) AccessApplicationsTestAccessPolicies(ctx context.Context, accountOrZone string, accountOrZoneID string, appID AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesParamsAppID, opts ...option.RequestOption) (res *AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%v/user_policy_checks", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponse struct {
	AppState     AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseAppState     `json:"app_state"`
	UserIdentity AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseUserIdentity `json:"user_identity"`
	JSON         accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseJSON         `json:"-"`
}

// accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseJSON
// contains the JSON metadata for the struct
// [AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponse]
type accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseJSON struct {
	AppState     apijson.Field
	UserIdentity apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseAppState struct {
	// UUID
	AppUid   string                                                                           `json:"app_uid"`
	Aud      string                                                                           `json:"aud"`
	Hostname string                                                                           `json:"hostname"`
	Name     string                                                                           `json:"name"`
	Policies []interface{}                                                                    `json:"policies"`
	Status   string                                                                           `json:"status"`
	JSON     accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseAppStateJSON `json:"-"`
}

// accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseAppStateJSON
// contains the JSON metadata for the struct
// [AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseAppState]
type accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseAppStateJSON struct {
	AppUid      apijson.Field
	Aud         apijson.Field
	Hostname    apijson.Field
	Name        apijson.Field
	Policies    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseAppState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseUserIdentity struct {
	ID             string                                                                              `json:"id"`
	AccountID      string                                                                              `json:"account_id"`
	DeviceSessions interface{}                                                                         `json:"device_sessions"`
	Email          string                                                                              `json:"email"`
	Geo            AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseUserIdentityGeo `json:"geo"`
	Iat            int64                                                                               `json:"iat"`
	IsGateway      bool                                                                                `json:"is_gateway"`
	IsWarp         bool                                                                                `json:"is_warp"`
	Name           string                                                                              `json:"name"`
	// UUID
	UserUuid string                                                                               `json:"user_uuid"`
	Version  int64                                                                                `json:"version"`
	JSON     accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseUserIdentityJSON `json:"-"`
}

// accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseUserIdentityJSON
// contains the JSON metadata for the struct
// [AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseUserIdentity]
type accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseUserIdentityJSON struct {
	ID             apijson.Field
	AccountID      apijson.Field
	DeviceSessions apijson.Field
	Email          apijson.Field
	Geo            apijson.Field
	Iat            apijson.Field
	IsGateway      apijson.Field
	IsWarp         apijson.Field
	Name           apijson.Field
	UserUuid       apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseUserIdentity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseUserIdentityGeo struct {
	Country string                                                                                  `json:"country"`
	JSON    accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseUserIdentityGeoJSON `json:"-"`
}

// accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseUserIdentityGeoJSON
// contains the JSON metadata for the struct
// [AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseUserIdentityGeo]
type accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseUserIdentityGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseUserIdentityGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesParamsAppID interface {
	ImplementsAccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesParamsAppID()
}

type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelope struct {
	Errors   []AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeMessages `json:"messages"`
	Result   AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeSuccess `json:"success"`
	JSON    accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeJSON    `json:"-"`
}

// accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelope]
type accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeErrors struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeErrors]
type accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeMessages struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeMessages]
type accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeSuccess bool

const (
	AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeSuccessTrue AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseEnvelopeSuccess = true
)

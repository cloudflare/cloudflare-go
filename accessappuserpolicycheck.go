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
	path := fmt.Sprintf("%s/%s/access/apps/%v/user_policy_checks", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponse struct {
	Errors   []AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseError   `json:"errors"`
	Messages []AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseMessage `json:"messages"`
	Result   AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseSuccess `json:"success"`
	JSON    accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseJSON    `json:"-"`
}

// accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseJSON
// contains the JSON metadata for the struct
// [AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponse]
type accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseError struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseErrorJSON `json:"-"`
}

// accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseError]
type accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseMessage struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseMessageJSON `json:"-"`
}

// accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseMessage]
type accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResult struct {
	AppState     AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultAppState     `json:"app_state"`
	UserIdentity AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultUserIdentity `json:"user_identity"`
	JSON         accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultJSON         `json:"-"`
}

// accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultJSON
// contains the JSON metadata for the struct
// [AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResult]
type accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultJSON struct {
	AppState     apijson.Field
	UserIdentity apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultAppState struct {
	// UUID
	AppUid   string                                                                                 `json:"app_uid"`
	Aud      string                                                                                 `json:"aud"`
	Hostname string                                                                                 `json:"hostname"`
	Name     string                                                                                 `json:"name"`
	Policies []interface{}                                                                          `json:"policies"`
	Status   string                                                                                 `json:"status"`
	JSON     accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultAppStateJSON `json:"-"`
}

// accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultAppStateJSON
// contains the JSON metadata for the struct
// [AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultAppState]
type accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultAppStateJSON struct {
	AppUid      apijson.Field
	Aud         apijson.Field
	Hostname    apijson.Field
	Name        apijson.Field
	Policies    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultAppState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultUserIdentity struct {
	ID             string                                                                                    `json:"id"`
	AccountID      string                                                                                    `json:"account_id"`
	DeviceSessions interface{}                                                                               `json:"device_sessions"`
	Email          string                                                                                    `json:"email"`
	Geo            AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultUserIdentityGeo `json:"geo"`
	Iat            int64                                                                                     `json:"iat"`
	IsGateway      bool                                                                                      `json:"is_gateway"`
	IsWarp         bool                                                                                      `json:"is_warp"`
	Name           string                                                                                    `json:"name"`
	// UUID
	UserUuid string                                                                                     `json:"user_uuid"`
	Version  int64                                                                                      `json:"version"`
	JSON     accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultUserIdentityJSON `json:"-"`
}

// accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultUserIdentityJSON
// contains the JSON metadata for the struct
// [AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultUserIdentity]
type accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultUserIdentityJSON struct {
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

func (r *AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultUserIdentity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultUserIdentityGeo struct {
	Country string                                                                                        `json:"country"`
	JSON    accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultUserIdentityGeoJSON `json:"-"`
}

// accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultUserIdentityGeoJSON
// contains the JSON metadata for the struct
// [AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultUserIdentityGeo]
type accessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultUserIdentityGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseResultUserIdentityGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseSuccess bool

const (
	AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseSuccessTrue AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesResponseSuccess = true
)

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesParamsAppID interface {
	ImplementsAccessAppUserPolicyCheckAccessApplicationsTestAccessPoliciesParamsAppID()
}

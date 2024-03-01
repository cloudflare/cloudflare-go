// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustAccessApplicationUserPolicyCheckService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZeroTrustAccessApplicationUserPolicyCheckService] method instead.
type ZeroTrustAccessApplicationUserPolicyCheckService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessApplicationUserPolicyCheckService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZeroTrustAccessApplicationUserPolicyCheckService(opts ...option.RequestOption) (r *ZeroTrustAccessApplicationUserPolicyCheckService) {
	r = &ZeroTrustAccessApplicationUserPolicyCheckService{}
	r.Options = opts
	return
}

// Tests if a specific user has permission to access an application.
func (r *ZeroTrustAccessApplicationUserPolicyCheckService) List(ctx context.Context, appID ZeroTrustAccessApplicationUserPolicyCheckListParamsAppID, query ZeroTrustAccessApplicationUserPolicyCheckListParams, opts ...option.RequestOption) (res *ZeroTrustAccessApplicationUserPolicyCheckListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps/%v/user_policy_checks", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustAccessApplicationUserPolicyCheckListResponse struct {
	AppState     ZeroTrustAccessApplicationUserPolicyCheckListResponseAppState     `json:"app_state"`
	UserIdentity ZeroTrustAccessApplicationUserPolicyCheckListResponseUserIdentity `json:"user_identity"`
	JSON         zeroTrustAccessApplicationUserPolicyCheckListResponseJSON         `json:"-"`
}

// zeroTrustAccessApplicationUserPolicyCheckListResponseJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationUserPolicyCheckListResponse]
type zeroTrustAccessApplicationUserPolicyCheckListResponseJSON struct {
	AppState     apijson.Field
	UserIdentity apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationUserPolicyCheckListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationUserPolicyCheckListResponseAppState struct {
	// UUID
	AppUid   string                                                            `json:"app_uid"`
	Aud      string                                                            `json:"aud"`
	Hostname string                                                            `json:"hostname"`
	Name     string                                                            `json:"name"`
	Policies []interface{}                                                     `json:"policies"`
	Status   string                                                            `json:"status"`
	JSON     zeroTrustAccessApplicationUserPolicyCheckListResponseAppStateJSON `json:"-"`
}

// zeroTrustAccessApplicationUserPolicyCheckListResponseAppStateJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationUserPolicyCheckListResponseAppState]
type zeroTrustAccessApplicationUserPolicyCheckListResponseAppStateJSON struct {
	AppUid      apijson.Field
	Aud         apijson.Field
	Hostname    apijson.Field
	Name        apijson.Field
	Policies    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationUserPolicyCheckListResponseAppState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationUserPolicyCheckListResponseUserIdentity struct {
	ID             string                                                               `json:"id"`
	AccountID      string                                                               `json:"account_id"`
	DeviceSessions interface{}                                                          `json:"device_sessions"`
	Email          string                                                               `json:"email"`
	Geo            ZeroTrustAccessApplicationUserPolicyCheckListResponseUserIdentityGeo `json:"geo"`
	Iat            int64                                                                `json:"iat"`
	IsGateway      bool                                                                 `json:"is_gateway"`
	IsWARP         bool                                                                 `json:"is_warp"`
	Name           string                                                               `json:"name"`
	// UUID
	UserUUID string                                                                `json:"user_uuid"`
	Version  int64                                                                 `json:"version"`
	JSON     zeroTrustAccessApplicationUserPolicyCheckListResponseUserIdentityJSON `json:"-"`
}

// zeroTrustAccessApplicationUserPolicyCheckListResponseUserIdentityJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationUserPolicyCheckListResponseUserIdentity]
type zeroTrustAccessApplicationUserPolicyCheckListResponseUserIdentityJSON struct {
	ID             apijson.Field
	AccountID      apijson.Field
	DeviceSessions apijson.Field
	Email          apijson.Field
	Geo            apijson.Field
	Iat            apijson.Field
	IsGateway      apijson.Field
	IsWARP         apijson.Field
	Name           apijson.Field
	UserUUID       apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationUserPolicyCheckListResponseUserIdentity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationUserPolicyCheckListResponseUserIdentityGeo struct {
	Country string                                                                   `json:"country"`
	JSON    zeroTrustAccessApplicationUserPolicyCheckListResponseUserIdentityGeoJSON `json:"-"`
}

// zeroTrustAccessApplicationUserPolicyCheckListResponseUserIdentityGeoJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationUserPolicyCheckListResponseUserIdentityGeo]
type zeroTrustAccessApplicationUserPolicyCheckListResponseUserIdentityGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationUserPolicyCheckListResponseUserIdentityGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationUserPolicyCheckListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type ZeroTrustAccessApplicationUserPolicyCheckListParamsAppID interface {
	ImplementsZeroTrustAccessApplicationUserPolicyCheckListParamsAppID()
}

type ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelope struct {
	Errors   []ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessApplicationUserPolicyCheckListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelope]
type zeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeErrors struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    zeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeErrors]
type zeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeMessages struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    zeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeMessages]
type zeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationUserPolicyCheckListResponseEnvelopeSuccess = true
)

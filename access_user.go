package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

/*---------------------*/
// GetAccessUserActiveSessions
//
/*---------------------*/

type AccessUserActiveSessionsResponse struct {
	Success    bool                      `json:"success"`
	Errors     []string                  `json:"errors"`
	Messages   []string                  `json:"messages"`
	Result     []AccessUserActiveSession `json:"result"`
	ResultInfo ResultInfo                `json:"result_info"`
}

type AccessUserActiveSession struct {
	Expiration int64  `json:"expiration"`
	Metadata   string `json:"metadata"`
	Name       string `json:"name"`
}

type AccessUserActiveSessionMetadata struct {
	Apps    map[string]AccessUserActiveSessionMetadataApp `json:"apps"`
	Expires int64                                         `json:"expires"`
	Iat     int64                                         `json:"iat"`
	Nonce   string                                        `json:"nonce"`
	Ttl     int64                                         `json:"ttl"`
}

type AccessUserActiveSessionMetadataApp struct {
	Hostname string `json:"hostname"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Uid      string `json:"uid"`
}

// GetAccessUserActiveSessions returns a list of active sessions for a user.
//
// API reference: https://developers.cloudflare.com/api/operations/zero-trust-users-get-active-sessions
func (api *API) GetAccessUserActiveSessions(ctx context.Context, rc *ResourceContainer, userID string) ([]AccessUserActiveSession, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/users/%s/active_sessions",
		rc.Level,
		rc.Identifier,
		userID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessUserActiveSession{}, err
	}

	var accessUserActiveSessionsResponse AccessUserActiveSessionsResponse
	err = json.Unmarshal(res, &accessUserActiveSessionsResponse)
	if err != nil {
		return []AccessUserActiveSession{}, err
	}
	return accessUserActiveSessionsResponse.Result, nil
}

/*---------------------*/
// GetAccessUserActiveSingleSessions
//
/*---------------------*/

type AccessUserSingleActiveSessionResponse struct {
	Success    bool                          `json:"success"`
	Errors     []string                      `json:"errors"`
	Messages   []string                      `json:"messages"`
	Result     AccessUserSingleActiveSession `json:"result"`
	ResultInfo ResultInfo                    `json:"result_info"`
}

type AccessUserSingleActiveSession struct {
	AccountID          string                                                `json:"account_id"`
	AuthStatus         string                                                `json:"auth_status"`
	CommonName         string                                                `json:"common_name"`
	DevicePosture      map[string]AccessUserSingleActiveSessionDevicePosture `json:"devicePosture"`
	DeviceID           string                                                `json:"device_id"`
	DeviceSessions     map[string]AccessUserSingleActiveSessionDeviceSession `json:"device_sessions"`
	Email              string                                                `json:"email"`
	Geo                AccessUserSingleActiveSessionGeo                      `json:"geo"`
	Iat                int64                                                 `json:"iat"`
	IDP                AccessUserSingleActiveSessionIDP                      `json:"idp"`
	IP                 string                                                `json:"ip"`
	IsGateway          bool                                                  `json:"is_gateway"`
	IsWarp             bool                                                  `json:"is_warp"`
	MTLSAuth           AccessUserSingleActiveSessionMTLSAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                                                `json:"service_token_id"`
	ServiceTokenStatus bool                                                  `json:"service_token_status"`
	UserUUID           string                                                `json:"user_uuid"`
	Version            int                                                   `json:"version"`
	IsActive           bool                                                  `json:"isActive"`
}

type AccessUserSingleActiveSessionDevicePosture struct {
	Check       AccessUserSingleActiveSessionDevicePostureCheck `json:"check"`
	Data        map[string]interface{}                          `json:"data"`
	Description string                                          `json:"description"`
	Error       string                                          `json:"error"`
	ID          string                                          `json:"id"`
	RuleName    string                                          `json:"rule_name"`
	Success     bool                                            `json:"success"`
	Timestamp   string                                          `json:"timestamp"`
	Type        string                                          `json:"type"`
}

type AccessUserSingleActiveSessionDevicePostureCheck struct {
	Exists bool   `json:"exists"`
	Path   string `json:"path"`
}

type AccessUserSingleActiveSessionDeviceSession struct {
	LastAuthenticated int64 `json:"last_authenticated"`
}

type AccessUserSingleActiveSessionGeo struct {
	Country string `json:"country"`
}

type AccessUserSingleActiveSessionIDP struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type AccessUserSingleActiveSessionMTLSAuth struct {
	AuthStatus    string `json:"auth_status"`
	CertIssuerDN  string `json:"cert_issuer_dn"`
	CertIssuerSKI string `json:"cert_issuer_ski"`
	CertPresented bool   `json:"cert_presented"`
	CertSerial    string `json:"cert_serial"`
}

// GetAccessUserSingleActiveSession returns a single active session for a user.
//
// API reference: https://developers.cloudflare.com/api/operations/zero-trust-users-get-active-session
func (api *API) GetAccessUserSingleActiveSession(ctx context.Context, rc *ResourceContainer, userID string, sessionID string) (AccessUserActiveSession, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/users/%s/active_sessions/%s",
		rc.Level,
		rc.Identifier,
		userID,
		sessionID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccessUserActiveSession{}, err
	}

	var accessUserActiveSessionsResponse AccessUserActiveSessionsResponse
	err = json.Unmarshal(res, &accessUserActiveSessionsResponse)
	if err != nil {
		return AccessUserActiveSession{}, err
	}
	return accessUserActiveSessionsResponse.Result[0], nil
}

/*---------------------*/
// GetAccessUserFailedLogins
//
/*---------------------*/

type AccessUserFailedLoginsResponse struct {
	Success    bool                    `json:"success"`
	Errors     []string                `json:"errors"`
	Messages   []string                `json:"messages"`
	Result     []AccessUserFailedLogin `json:"result"`
	ResultInfo ResultInfo              `json:"result_info"`
}

type AccessUserFailedLogin struct {
	Expiration int64  `json:"expiration"`
	Metadata   string `json:"metadata"`
}

type AccessUserFailedLoginMetadata struct {
	AppName   string `json:"app_name"`
	Aud       string `json:"aud"`
	Datetime  string `json:"datetime"`
	RayID     string `json:"ray_id"`
	UserEmail string `json:"user_email"`
	UserUUID  string `json:"user_uuid"`
}

// GetAccessUserFailedLogins returns a list of failed logins for a user.
//
// API reference: https://developers.cloudflare.com/api/operations/zero-trust-users-get-failed-logins
func (api *API) GetAccessUserFailedLogins(ctx context.Context, rc *ResourceContainer, userID string) ([]AccessUserFailedLogin, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/users/%s/failed_logins",
		rc.Level,
		rc.Identifier,
		userID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessUserFailedLogin{}, err
	}

	var accessUserFailedLoginsResponse AccessUserFailedLoginsResponse
	err = json.Unmarshal(res, &accessUserFailedLoginsResponse)
	if err != nil {
		return []AccessUserFailedLogin{}, err
	}
	return accessUserFailedLoginsResponse.Result, nil
}

/*---------------------*/
// GetAccessUserLastSeenIdentity
//
/*---------------------*/

type AccessUserLastSeenIdentityResponse struct {
	Success    bool                       `json:"success"`
	Errors     []string                   `json:"errors"`
	Messages   []string                   `json:"messages"`
	Result     AccessUserLastSeenIdentity `json:"result"`
	ResultInfo ResultInfo                 `json:"result_info"`
}

type AccessUserLastSeenIdentity struct {
	AccountID          string                                             `json:"account_id"`
	AuthStatus         string                                             `json:"auth_status"`
	CommonName         string                                             `json:"common_name"`
	DevicePosture      map[string]AccessUserLastSeenIdentityDevicePosture `json:"devicePosture"`
	DeviceID           string                                             `json:"device_id"`
	DeviceSessions     map[string]AccessUserLastSeenIdentityDeviceSession `json:"device_sessions"`
	Email              string                                             `json:"email"`
	Geo                AccessUserLastSeenIdentityGeo                      `json:"geo"`
	Iat                int64                                              `json:"iat"`
	IDP                AccessUserLastSeenIdentityIDP                      `json:"idp"`
	IP                 string                                             `json:"ip"`
	IsGateway          bool                                               `json:"is_gateway"`
	IsWarp             bool                                               `json:"is_warp"`
	MTLSAuth           AccessUserLastSeenIdentityMTLSAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                                             `json:"service_token_id"`
	ServiceTokenStatus bool                                               `json:"service_token_status"`
	UserUUID           string                                             `json:"user_uuid"`
	Version            int                                                `json:"version"`
	IsActive           bool                                               `json:"isActive"`
}

type AccessUserLastSeenIdentityDevicePosture struct {
	Check       AccessUserLastSeenIdentityDevicePostureCheck `json:"check"`
	Data        map[string]interface{}                       `json:"data"`
	Description string                                       `json:"description"`
	Error       string                                       `json:"error"`
	ID          string                                       `json:"id"`
	RuleName    string                                       `json:"rule_name"`
	Success     bool                                         `json:"success"`
	Timestamp   string                                       `json:"timestamp"`
	Type        string                                       `json:"type"`
}

type AccessUserLastSeenIdentityDevicePostureCheck struct {
	Exists bool   `json:"exists"`
	Path   string `json:"path"`
}

type AccessUserLastSeenIdentityDeviceSession struct {
	LastAuthenticated int64 `json:"last_authenticated"`
}

type AccessUserLastSeenIdentityGeo struct {
	Country string `json:"country"`
}

type AccessUserLastSeenIdentityIDP struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type AccessUserLastSeenIdentityMTLSAuth struct {
	AuthStatus    string `json:"auth_status"`
	CertIssuerDN  string `json:"cert_issuer_dn"`
	CertIssuerSKI string `json:"cert_issuer_ski"`
	CertPresented bool   `json:"cert_presented"`
	CertSerial    string `json:"cert_serial"`
}

// GetAccessUserLastSeenIdentity returns the last seen identity for a user.
//
// API reference: https://developers.cloudflare.com/api/operations/zero-trust-users-get-last-seen-identity
func (api *API) GetAccessUserLastSeenIdentity(ctx context.Context, rc *ResourceContainer, userID string) (AccessUserLastSeenIdentity, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/users/%s/last_seen_identity",
		rc.Level,
		rc.Identifier,
		userID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccessUserLastSeenIdentity{}, err
	}

	var accessUserLastSeenIdentityResponse AccessUserLastSeenIdentityResponse
	err = json.Unmarshal(res, &accessUserLastSeenIdentityResponse)
	if err != nil {
		return AccessUserLastSeenIdentity{}, err
	}
	return accessUserLastSeenIdentityResponse.Result, nil
}

type AccessUserListResponse struct {
	Success    bool         `json:"success"`
	Errors     []string     `json:"errors"`
	Messages   []string     `json:"messages"`
	Result     []AccessUser `json:"result"`
	ResultInfo ResultInfo   `json:"result_info"`
}

type AccessUser struct {
	ID                  string `json:"id"`
	AccessSeat          bool   `json:"access_seat"`
	ActiveDeviceCount   int    `json:"active_device_count"`
	CreatedAt           string `json:"created_at"`
	Email               string `json:"email"`
	GatewaySeat         bool   `json:"gateway_seat"`
	LastSuccessfulLogin string `json:"last_successful_login"`
	Name                string `json:"name"`
	SeatUID             string `json:"seat_uid"`
	UID                 string `json:"uid"`
	UpdatedAt           string `json:"updated_at"`
}

type AccessUserParams struct {
	ResultInfo
}

// GetAccessUsers returns a list of users.
//
// API reference: https://developers.cloudflare.com/api/operations/zero-trust-users-get-users
func (api *API) ListAccessUsers(ctx context.Context, rc *ResourceContainer, params ListAccessGroupsParams) ([]AccessUser, *ResultInfo, error) {
	baseURL := fmt.Sprintf("/%s/%s/access/users", rc.Level, rc.Identifier)

	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}

	if params.PerPage < 1 {
		params.PerPage = 25
	}

	if params.Page < 1 {
		params.Page = 1
	}

	var accessUsers []AccessUser
	var r AccessUserListResponse

	for {
		uri := buildURI(baseURL, params)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []AccessUser{}, &ResultInfo{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
		}

		err = json.Unmarshal(res, &r)
		if err != nil {
			return []AccessUser{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		accessUsers = append(accessUsers, r.Result...)
		params.ResultInfo = r.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return accessUsers, &r.ResultInfo, nil
}

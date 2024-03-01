// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustAccessOrganizationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustAccessOrganizationService] method instead.
type ZeroTrustAccessOrganizationService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessOrganizationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustAccessOrganizationService(opts ...option.RequestOption) (r *ZeroTrustAccessOrganizationService) {
	r = &ZeroTrustAccessOrganizationService{}
	r.Options = opts
	return
}

// Sets up a Zero Trust organization for your account or zone.
func (r *ZeroTrustAccessOrganizationService) New(ctx context.Context, params ZeroTrustAccessOrganizationNewParams, opts ...option.RequestOption) (res *ZeroTrustAccessOrganizationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessOrganizationNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the configuration for your Zero Trust organization.
func (r *ZeroTrustAccessOrganizationService) Update(ctx context.Context, params ZeroTrustAccessOrganizationUpdateParams, opts ...option.RequestOption) (res *ZeroTrustAccessOrganizationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessOrganizationUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the configuration for your Zero Trust organization.
func (r *ZeroTrustAccessOrganizationService) List(ctx context.Context, query ZeroTrustAccessOrganizationListParams, opts ...option.RequestOption) (res *ZeroTrustAccessOrganizationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessOrganizationListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Revokes a user's access across all applications.
func (r *ZeroTrustAccessOrganizationService) RevokeUsers(ctx context.Context, params ZeroTrustAccessOrganizationRevokeUsersParams, opts ...option.RequestOption) (res *ZeroTrustAccessOrganizationRevokeUsersResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessOrganizationRevokeUsersResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/organizations/revoke_user", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustAccessOrganizationNewResponse struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWARP bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                              `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                         `json:"created_at" format:"date-time"`
	CustomPages            ZeroTrustAccessOrganizationNewResponseCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                              `json:"is_ui_read_only"`
	LoginDesign  ZeroTrustAccessOrganizationNewResponseLoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration string `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WARPAuthSessionDuration string                                     `json:"warp_auth_session_duration"`
	JSON                    zeroTrustAccessOrganizationNewResponseJSON `json:"-"`
}

// zeroTrustAccessOrganizationNewResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessOrganizationNewResponse]
type zeroTrustAccessOrganizationNewResponseJSON struct {
	AllowAuthenticateViaWARP       apijson.Field
	AuthDomain                     apijson.Field
	AutoRedirectToIdentity         apijson.Field
	CreatedAt                      apijson.Field
	CustomPages                    apijson.Field
	IsUiReadOnly                   apijson.Field
	LoginDesign                    apijson.Field
	Name                           apijson.Field
	SessionDuration                apijson.Field
	UiReadOnlyToggleReason         apijson.Field
	UpdatedAt                      apijson.Field
	UserSeatExpirationInactiveTime apijson.Field
	WARPAuthSessionDuration        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationNewResponseCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                                `json:"identity_denied"`
	JSON           zeroTrustAccessOrganizationNewResponseCustomPagesJSON `json:"-"`
}

// zeroTrustAccessOrganizationNewResponseCustomPagesJSON contains the JSON metadata
// for the struct [ZeroTrustAccessOrganizationNewResponseCustomPages]
type zeroTrustAccessOrganizationNewResponseCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationNewResponseCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationNewResponseLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                                `json:"text_color"`
	JSON      zeroTrustAccessOrganizationNewResponseLoginDesignJSON `json:"-"`
}

// zeroTrustAccessOrganizationNewResponseLoginDesignJSON contains the JSON metadata
// for the struct [ZeroTrustAccessOrganizationNewResponseLoginDesign]
type zeroTrustAccessOrganizationNewResponseLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationNewResponseLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationUpdateResponse struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWARP bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                                 `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                            `json:"created_at" format:"date-time"`
	CustomPages            ZeroTrustAccessOrganizationUpdateResponseCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                                 `json:"is_ui_read_only"`
	LoginDesign  ZeroTrustAccessOrganizationUpdateResponseLoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration string `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WARPAuthSessionDuration string                                        `json:"warp_auth_session_duration"`
	JSON                    zeroTrustAccessOrganizationUpdateResponseJSON `json:"-"`
}

// zeroTrustAccessOrganizationUpdateResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessOrganizationUpdateResponse]
type zeroTrustAccessOrganizationUpdateResponseJSON struct {
	AllowAuthenticateViaWARP       apijson.Field
	AuthDomain                     apijson.Field
	AutoRedirectToIdentity         apijson.Field
	CreatedAt                      apijson.Field
	CustomPages                    apijson.Field
	IsUiReadOnly                   apijson.Field
	LoginDesign                    apijson.Field
	Name                           apijson.Field
	SessionDuration                apijson.Field
	UiReadOnlyToggleReason         apijson.Field
	UpdatedAt                      apijson.Field
	UserSeatExpirationInactiveTime apijson.Field
	WARPAuthSessionDuration        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationUpdateResponseCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                                   `json:"identity_denied"`
	JSON           zeroTrustAccessOrganizationUpdateResponseCustomPagesJSON `json:"-"`
}

// zeroTrustAccessOrganizationUpdateResponseCustomPagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessOrganizationUpdateResponseCustomPages]
type zeroTrustAccessOrganizationUpdateResponseCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationUpdateResponseCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationUpdateResponseLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                                   `json:"text_color"`
	JSON      zeroTrustAccessOrganizationUpdateResponseLoginDesignJSON `json:"-"`
}

// zeroTrustAccessOrganizationUpdateResponseLoginDesignJSON contains the JSON
// metadata for the struct [ZeroTrustAccessOrganizationUpdateResponseLoginDesign]
type zeroTrustAccessOrganizationUpdateResponseLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationUpdateResponseLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationListResponse struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWARP bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                               `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                          `json:"created_at" format:"date-time"`
	CustomPages            ZeroTrustAccessOrganizationListResponseCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                               `json:"is_ui_read_only"`
	LoginDesign  ZeroTrustAccessOrganizationListResponseLoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration string `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WARPAuthSessionDuration string                                      `json:"warp_auth_session_duration"`
	JSON                    zeroTrustAccessOrganizationListResponseJSON `json:"-"`
}

// zeroTrustAccessOrganizationListResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessOrganizationListResponse]
type zeroTrustAccessOrganizationListResponseJSON struct {
	AllowAuthenticateViaWARP       apijson.Field
	AuthDomain                     apijson.Field
	AutoRedirectToIdentity         apijson.Field
	CreatedAt                      apijson.Field
	CustomPages                    apijson.Field
	IsUiReadOnly                   apijson.Field
	LoginDesign                    apijson.Field
	Name                           apijson.Field
	SessionDuration                apijson.Field
	UiReadOnlyToggleReason         apijson.Field
	UpdatedAt                      apijson.Field
	UserSeatExpirationInactiveTime apijson.Field
	WARPAuthSessionDuration        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationListResponseCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                                 `json:"identity_denied"`
	JSON           zeroTrustAccessOrganizationListResponseCustomPagesJSON `json:"-"`
}

// zeroTrustAccessOrganizationListResponseCustomPagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessOrganizationListResponseCustomPages]
type zeroTrustAccessOrganizationListResponseCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationListResponseCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationListResponseLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                                 `json:"text_color"`
	JSON      zeroTrustAccessOrganizationListResponseLoginDesignJSON `json:"-"`
}

// zeroTrustAccessOrganizationListResponseLoginDesignJSON contains the JSON
// metadata for the struct [ZeroTrustAccessOrganizationListResponseLoginDesign]
type zeroTrustAccessOrganizationListResponseLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationListResponseLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationRevokeUsersResponse bool

const (
	ZeroTrustAccessOrganizationRevokeUsersResponseTrue  ZeroTrustAccessOrganizationRevokeUsersResponse = true
	ZeroTrustAccessOrganizationRevokeUsersResponseFalse ZeroTrustAccessOrganizationRevokeUsersResponse = false
)

type ZeroTrustAccessOrganizationNewParams struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain,required"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]                                            `json:"is_ui_read_only"`
	LoginDesign  param.Field[ZeroTrustAccessOrganizationNewParamsLoginDesign] `json:"login_design"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason param.Field[string] `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime param.Field[string] `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WARPAuthSessionDuration param.Field[string] `json:"warp_auth_session_duration"`
}

func (r ZeroTrustAccessOrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessOrganizationNewParamsLoginDesign struct {
	// The background color on your login page.
	BackgroundColor param.Field[string] `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText param.Field[string] `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText param.Field[string] `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath param.Field[string] `json:"logo_path"`
	// The text color on your login page.
	TextColor param.Field[string] `json:"text_color"`
}

func (r ZeroTrustAccessOrganizationNewParamsLoginDesign) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessOrganizationNewResponseEnvelope struct {
	Errors   []ZeroTrustAccessOrganizationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessOrganizationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessOrganizationNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessOrganizationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessOrganizationNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessOrganizationNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessOrganizationNewResponseEnvelope]
type zeroTrustAccessOrganizationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationNewResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustAccessOrganizationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessOrganizationNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessOrganizationNewResponseEnvelopeErrors]
type zeroTrustAccessOrganizationNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationNewResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustAccessOrganizationNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessOrganizationNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessOrganizationNewResponseEnvelopeMessages]
type zeroTrustAccessOrganizationNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessOrganizationNewResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessOrganizationNewResponseEnvelopeSuccessTrue ZeroTrustAccessOrganizationNewResponseEnvelopeSuccess = true
)

type ZeroTrustAccessOrganizationUpdateParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity param.Field[bool]                                               `json:"auto_redirect_to_identity"`
	CustomPages            param.Field[ZeroTrustAccessOrganizationUpdateParamsCustomPages] `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]                                               `json:"is_ui_read_only"`
	LoginDesign  param.Field[ZeroTrustAccessOrganizationUpdateParamsLoginDesign] `json:"login_design"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason param.Field[string] `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime param.Field[string] `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WARPAuthSessionDuration param.Field[string] `json:"warp_auth_session_duration"`
}

func (r ZeroTrustAccessOrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessOrganizationUpdateParamsCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden param.Field[string] `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied param.Field[string] `json:"identity_denied"`
}

func (r ZeroTrustAccessOrganizationUpdateParamsCustomPages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessOrganizationUpdateParamsLoginDesign struct {
	// The background color on your login page.
	BackgroundColor param.Field[string] `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText param.Field[string] `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText param.Field[string] `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath param.Field[string] `json:"logo_path"`
	// The text color on your login page.
	TextColor param.Field[string] `json:"text_color"`
}

func (r ZeroTrustAccessOrganizationUpdateParamsLoginDesign) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessOrganizationUpdateResponseEnvelope struct {
	Errors   []ZeroTrustAccessOrganizationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessOrganizationUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessOrganizationUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessOrganizationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessOrganizationUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessOrganizationUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessOrganizationUpdateResponseEnvelope]
type zeroTrustAccessOrganizationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationUpdateResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustAccessOrganizationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessOrganizationUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessOrganizationUpdateResponseEnvelopeErrors]
type zeroTrustAccessOrganizationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationUpdateResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zeroTrustAccessOrganizationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessOrganizationUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessOrganizationUpdateResponseEnvelopeMessages]
type zeroTrustAccessOrganizationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessOrganizationUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessOrganizationUpdateResponseEnvelopeSuccessTrue ZeroTrustAccessOrganizationUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustAccessOrganizationListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessOrganizationListResponseEnvelope struct {
	Errors   []ZeroTrustAccessOrganizationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessOrganizationListResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessOrganizationListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessOrganizationListResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessOrganizationListResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessOrganizationListResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessOrganizationListResponseEnvelope]
type zeroTrustAccessOrganizationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationListResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustAccessOrganizationListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessOrganizationListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessOrganizationListResponseEnvelopeErrors]
type zeroTrustAccessOrganizationListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationListResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustAccessOrganizationListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessOrganizationListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessOrganizationListResponseEnvelopeMessages]
type zeroTrustAccessOrganizationListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessOrganizationListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessOrganizationListResponseEnvelopeSuccessTrue ZeroTrustAccessOrganizationListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessOrganizationRevokeUsersParams struct {
	// The email of the user to revoke.
	Email param.Field[string] `json:"email,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r ZeroTrustAccessOrganizationRevokeUsersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessOrganizationRevokeUsersResponseEnvelope struct {
	Result  ZeroTrustAccessOrganizationRevokeUsersResponse                `json:"result"`
	Success ZeroTrustAccessOrganizationRevokeUsersResponseEnvelopeSuccess `json:"success"`
	JSON    zeroTrustAccessOrganizationRevokeUsersResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessOrganizationRevokeUsersResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessOrganizationRevokeUsersResponseEnvelope]
type zeroTrustAccessOrganizationRevokeUsersResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessOrganizationRevokeUsersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessOrganizationRevokeUsersResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessOrganizationRevokeUsersResponseEnvelopeSuccessTrue  ZeroTrustAccessOrganizationRevokeUsersResponseEnvelopeSuccess = true
	ZeroTrustAccessOrganizationRevokeUsersResponseEnvelopeSuccessFalse ZeroTrustAccessOrganizationRevokeUsersResponseEnvelopeSuccess = false
)

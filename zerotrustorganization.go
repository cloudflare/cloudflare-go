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

// ZeroTrustOrganizationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustOrganizationService]
// method instead.
type ZeroTrustOrganizationService struct {
	Options []option.RequestOption
}

// NewZeroTrustOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustOrganizationService(opts ...option.RequestOption) (r *ZeroTrustOrganizationService) {
	r = &ZeroTrustOrganizationService{}
	r.Options = opts
	return
}

// Sets up a Zero Trust organization for your account or zone.
func (r *ZeroTrustOrganizationService) New(ctx context.Context, params ZeroTrustOrganizationNewParams, opts ...option.RequestOption) (res *ZeroTrustOrganizationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustOrganizationNewResponseEnvelope
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
func (r *ZeroTrustOrganizationService) Update(ctx context.Context, params ZeroTrustOrganizationUpdateParams, opts ...option.RequestOption) (res *ZeroTrustOrganizationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustOrganizationUpdateResponseEnvelope
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
func (r *ZeroTrustOrganizationService) List(ctx context.Context, query ZeroTrustOrganizationListParams, opts ...option.RequestOption) (res *ZeroTrustOrganizationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustOrganizationListResponseEnvelope
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
func (r *ZeroTrustOrganizationService) RevokeUsers(ctx context.Context, params ZeroTrustOrganizationRevokeUsersParams, opts ...option.RequestOption) (res *ZeroTrustOrganizationRevokeUsersResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustOrganizationRevokeUsersResponseEnvelope
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

type ZeroTrustOrganizationNewResponse struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWARP bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                        `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                   `json:"created_at" format:"date-time"`
	CustomPages            ZeroTrustOrganizationNewResponseCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                        `json:"is_ui_read_only"`
	LoginDesign  ZeroTrustOrganizationNewResponseLoginDesign `json:"login_design"`
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
	WARPAuthSessionDuration string                               `json:"warp_auth_session_duration"`
	JSON                    zeroTrustOrganizationNewResponseJSON `json:"-"`
}

// zeroTrustOrganizationNewResponseJSON contains the JSON metadata for the struct
// [ZeroTrustOrganizationNewResponse]
type zeroTrustOrganizationNewResponseJSON struct {
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

func (r *ZeroTrustOrganizationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationNewResponseCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                          `json:"identity_denied"`
	JSON           zeroTrustOrganizationNewResponseCustomPagesJSON `json:"-"`
}

// zeroTrustOrganizationNewResponseCustomPagesJSON contains the JSON metadata for
// the struct [ZeroTrustOrganizationNewResponseCustomPages]
type zeroTrustOrganizationNewResponseCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustOrganizationNewResponseCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationNewResponseCustomPagesJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationNewResponseLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                          `json:"text_color"`
	JSON      zeroTrustOrganizationNewResponseLoginDesignJSON `json:"-"`
}

// zeroTrustOrganizationNewResponseLoginDesignJSON contains the JSON metadata for
// the struct [ZeroTrustOrganizationNewResponseLoginDesign]
type zeroTrustOrganizationNewResponseLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustOrganizationNewResponseLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationNewResponseLoginDesignJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationUpdateResponse struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWARP bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                           `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                      `json:"created_at" format:"date-time"`
	CustomPages            ZeroTrustOrganizationUpdateResponseCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                           `json:"is_ui_read_only"`
	LoginDesign  ZeroTrustOrganizationUpdateResponseLoginDesign `json:"login_design"`
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
	WARPAuthSessionDuration string                                  `json:"warp_auth_session_duration"`
	JSON                    zeroTrustOrganizationUpdateResponseJSON `json:"-"`
}

// zeroTrustOrganizationUpdateResponseJSON contains the JSON metadata for the
// struct [ZeroTrustOrganizationUpdateResponse]
type zeroTrustOrganizationUpdateResponseJSON struct {
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

func (r *ZeroTrustOrganizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationUpdateResponseCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                             `json:"identity_denied"`
	JSON           zeroTrustOrganizationUpdateResponseCustomPagesJSON `json:"-"`
}

// zeroTrustOrganizationUpdateResponseCustomPagesJSON contains the JSON metadata
// for the struct [ZeroTrustOrganizationUpdateResponseCustomPages]
type zeroTrustOrganizationUpdateResponseCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustOrganizationUpdateResponseCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationUpdateResponseCustomPagesJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationUpdateResponseLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                             `json:"text_color"`
	JSON      zeroTrustOrganizationUpdateResponseLoginDesignJSON `json:"-"`
}

// zeroTrustOrganizationUpdateResponseLoginDesignJSON contains the JSON metadata
// for the struct [ZeroTrustOrganizationUpdateResponseLoginDesign]
type zeroTrustOrganizationUpdateResponseLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustOrganizationUpdateResponseLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationUpdateResponseLoginDesignJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationListResponse struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWARP bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                         `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                    `json:"created_at" format:"date-time"`
	CustomPages            ZeroTrustOrganizationListResponseCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                         `json:"is_ui_read_only"`
	LoginDesign  ZeroTrustOrganizationListResponseLoginDesign `json:"login_design"`
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
	WARPAuthSessionDuration string                                `json:"warp_auth_session_duration"`
	JSON                    zeroTrustOrganizationListResponseJSON `json:"-"`
}

// zeroTrustOrganizationListResponseJSON contains the JSON metadata for the struct
// [ZeroTrustOrganizationListResponse]
type zeroTrustOrganizationListResponseJSON struct {
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

func (r *ZeroTrustOrganizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationListResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationListResponseCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                           `json:"identity_denied"`
	JSON           zeroTrustOrganizationListResponseCustomPagesJSON `json:"-"`
}

// zeroTrustOrganizationListResponseCustomPagesJSON contains the JSON metadata for
// the struct [ZeroTrustOrganizationListResponseCustomPages]
type zeroTrustOrganizationListResponseCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustOrganizationListResponseCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationListResponseCustomPagesJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationListResponseLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                           `json:"text_color"`
	JSON      zeroTrustOrganizationListResponseLoginDesignJSON `json:"-"`
}

// zeroTrustOrganizationListResponseLoginDesignJSON contains the JSON metadata for
// the struct [ZeroTrustOrganizationListResponseLoginDesign]
type zeroTrustOrganizationListResponseLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustOrganizationListResponseLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationListResponseLoginDesignJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationRevokeUsersResponse bool

const (
	ZeroTrustOrganizationRevokeUsersResponseTrue  ZeroTrustOrganizationRevokeUsersResponse = true
	ZeroTrustOrganizationRevokeUsersResponseFalse ZeroTrustOrganizationRevokeUsersResponse = false
)

type ZeroTrustOrganizationNewParams struct {
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
	IsUiReadOnly param.Field[bool]                                      `json:"is_ui_read_only"`
	LoginDesign  param.Field[ZeroTrustOrganizationNewParamsLoginDesign] `json:"login_design"`
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

func (r ZeroTrustOrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustOrganizationNewParamsLoginDesign struct {
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

func (r ZeroTrustOrganizationNewParamsLoginDesign) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustOrganizationNewResponseEnvelope struct {
	Errors   []ZeroTrustOrganizationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustOrganizationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustOrganizationNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustOrganizationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustOrganizationNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustOrganizationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustOrganizationNewResponseEnvelope]
type zeroTrustOrganizationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustOrganizationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationNewResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustOrganizationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustOrganizationNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustOrganizationNewResponseEnvelopeErrors]
type zeroTrustOrganizationNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustOrganizationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationNewResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustOrganizationNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustOrganizationNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustOrganizationNewResponseEnvelopeMessages]
type zeroTrustOrganizationNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustOrganizationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustOrganizationNewResponseEnvelopeSuccess bool

const (
	ZeroTrustOrganizationNewResponseEnvelopeSuccessTrue ZeroTrustOrganizationNewResponseEnvelopeSuccess = true
)

type ZeroTrustOrganizationUpdateParams struct {
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
	AutoRedirectToIdentity param.Field[bool]                                         `json:"auto_redirect_to_identity"`
	CustomPages            param.Field[ZeroTrustOrganizationUpdateParamsCustomPages] `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]                                         `json:"is_ui_read_only"`
	LoginDesign  param.Field[ZeroTrustOrganizationUpdateParamsLoginDesign] `json:"login_design"`
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

func (r ZeroTrustOrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustOrganizationUpdateParamsCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden param.Field[string] `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied param.Field[string] `json:"identity_denied"`
}

func (r ZeroTrustOrganizationUpdateParamsCustomPages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustOrganizationUpdateParamsLoginDesign struct {
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

func (r ZeroTrustOrganizationUpdateParamsLoginDesign) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustOrganizationUpdateResponseEnvelope struct {
	Errors   []ZeroTrustOrganizationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustOrganizationUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustOrganizationUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustOrganizationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustOrganizationUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustOrganizationUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustOrganizationUpdateResponseEnvelope]
type zeroTrustOrganizationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustOrganizationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationUpdateResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustOrganizationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustOrganizationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustOrganizationUpdateResponseEnvelopeErrors]
type zeroTrustOrganizationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustOrganizationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationUpdateResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustOrganizationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustOrganizationUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustOrganizationUpdateResponseEnvelopeMessages]
type zeroTrustOrganizationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustOrganizationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustOrganizationUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustOrganizationUpdateResponseEnvelopeSuccessTrue ZeroTrustOrganizationUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustOrganizationListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustOrganizationListResponseEnvelope struct {
	Errors   []ZeroTrustOrganizationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustOrganizationListResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustOrganizationListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustOrganizationListResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustOrganizationListResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustOrganizationListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustOrganizationListResponseEnvelope]
type zeroTrustOrganizationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustOrganizationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationListResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustOrganizationListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustOrganizationListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustOrganizationListResponseEnvelopeErrors]
type zeroTrustOrganizationListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustOrganizationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationListResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustOrganizationListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustOrganizationListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustOrganizationListResponseEnvelopeMessages]
type zeroTrustOrganizationListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustOrganizationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustOrganizationListResponseEnvelopeSuccess bool

const (
	ZeroTrustOrganizationListResponseEnvelopeSuccessTrue ZeroTrustOrganizationListResponseEnvelopeSuccess = true
)

type ZeroTrustOrganizationRevokeUsersParams struct {
	// The email of the user to revoke.
	Email param.Field[string] `json:"email,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r ZeroTrustOrganizationRevokeUsersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustOrganizationRevokeUsersResponseEnvelope struct {
	Result  ZeroTrustOrganizationRevokeUsersResponse                `json:"result"`
	Success ZeroTrustOrganizationRevokeUsersResponseEnvelopeSuccess `json:"success"`
	JSON    zeroTrustOrganizationRevokeUsersResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustOrganizationRevokeUsersResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustOrganizationRevokeUsersResponseEnvelope]
type zeroTrustOrganizationRevokeUsersResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustOrganizationRevokeUsersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustOrganizationRevokeUsersResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustOrganizationRevokeUsersResponseEnvelopeSuccess bool

const (
	ZeroTrustOrganizationRevokeUsersResponseEnvelopeSuccessTrue  ZeroTrustOrganizationRevokeUsersResponseEnvelopeSuccess = true
	ZeroTrustOrganizationRevokeUsersResponseEnvelopeSuccessFalse ZeroTrustOrganizationRevokeUsersResponseEnvelopeSuccess = false
)

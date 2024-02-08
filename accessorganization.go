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

// AccessOrganizationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessOrganizationService] method
// instead.
type AccessOrganizationService struct {
	Options []option.RequestOption
}

// NewAccessOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessOrganizationService(opts ...option.RequestOption) (r *AccessOrganizationService) {
	r = &AccessOrganizationService{}
	r.Options = opts
	return
}

// Sets up a Zero Trust organization for your account or zone.
func (r *AccessOrganizationService) ZeroTrustOrganizationNewYourZeroTrustOrganization(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParams, opts ...option.RequestOption) (res *AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the configuration for your Zero Trust organization.
func (r *AccessOrganizationService) ZeroTrustOrganizationGetYourZeroTrustOrganization(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the configuration for your Zero Trust organization.
func (r *AccessOrganizationService) ZeroTrustOrganizationUpdateYourZeroTrustOrganization(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams, opts ...option.RequestOption) (res *AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                                                                   `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                                                              `json:"created_at" format:"date-time"`
	CustomPages            AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                                                                   `json:"is_ui_read_only"`
	LoginDesign  AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseLoginDesign `json:"login_design"`
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
	WarpAuthSessionDuration string                                                                          `json:"warp_auth_session_duration"`
	JSON                    accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse]
type accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseJSON struct {
	AllowAuthenticateViaWarp       apijson.Field
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
	WarpAuthSessionDuration        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                                                                     `json:"identity_denied"`
	JSON           accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseCustomPagesJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseCustomPagesJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseCustomPages]
type accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                                                                     `json:"text_color"`
	JSON      accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseLoginDesignJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseLoginDesignJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseLoginDesign]
type accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                                                                   `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                                                              `json:"created_at" format:"date-time"`
	CustomPages            AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                                                                   `json:"is_ui_read_only"`
	LoginDesign  AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseLoginDesign `json:"login_design"`
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
	WarpAuthSessionDuration string                                                                          `json:"warp_auth_session_duration"`
	JSON                    accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse]
type accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseJSON struct {
	AllowAuthenticateViaWarp       apijson.Field
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
	WarpAuthSessionDuration        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                                                                     `json:"identity_denied"`
	JSON           accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseCustomPagesJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseCustomPagesJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseCustomPages]
type accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                                                                     `json:"text_color"`
	JSON      accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseLoginDesignJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseLoginDesignJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseLoginDesign]
type accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                                                                      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                                                                 `json:"created_at" format:"date-time"`
	CustomPages            AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                                                                      `json:"is_ui_read_only"`
	LoginDesign  AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseLoginDesign `json:"login_design"`
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
	WarpAuthSessionDuration string                                                                             `json:"warp_auth_session_duration"`
	JSON                    accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse]
type accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseJSON struct {
	AllowAuthenticateViaWarp       apijson.Field
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
	WarpAuthSessionDuration        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                                                                        `json:"identity_denied"`
	JSON           accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseCustomPagesJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseCustomPagesJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseCustomPages]
type accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                                                                        `json:"text_color"`
	JSON      accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseLoginDesignJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseLoginDesignJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseLoginDesign]
type accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParams struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain,required"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name,required"`
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp param.Field[bool] `json:"allow_authenticate_via_warp"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]                                                                                 `json:"is_ui_read_only"`
	LoginDesign  param.Field[AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParamsLoginDesign] `json:"login_design"`
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
	WarpAuthSessionDuration param.Field[string] `json:"warp_auth_session_duration"`
}

func (r AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParamsLoginDesign struct {
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

func (r AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParamsLoginDesign) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelope struct {
	Errors   []AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeJSON    `json:"-"`
}

// accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelope]
type accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeErrors struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeErrorsJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeErrors]
type accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeMessages struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeMessagesJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeMessages]
type accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeSuccess bool

const (
	AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeSuccessTrue AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseEnvelopeSuccess = true
)

type AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelope struct {
	Errors   []AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeJSON    `json:"-"`
}

// accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelope]
type accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeErrors struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeErrorsJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeErrors]
type accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeMessages struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeMessagesJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeMessages]
type accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeSuccess bool

const (
	AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeSuccessTrue AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseEnvelopeSuccess = true
)

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity param.Field[bool]                                                                                    `json:"auto_redirect_to_identity"`
	CustomPages            param.Field[AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsCustomPages] `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]                                                                                    `json:"is_ui_read_only"`
	LoginDesign  param.Field[AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign] `json:"login_design"`
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
	WarpAuthSessionDuration param.Field[string] `json:"warp_auth_session_duration"`
}

func (r AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden param.Field[string] `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied param.Field[string] `json:"identity_denied"`
}

func (r AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsCustomPages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign struct {
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

func (r AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelope struct {
	Errors   []AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeJSON    `json:"-"`
}

// accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelope]
type accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeErrors struct {
	Code    int64                                                                                            `json:"code,required"`
	Message string                                                                                           `json:"message,required"`
	JSON    accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeErrorsJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeErrors]
type accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeMessages struct {
	Code    int64                                                                                              `json:"code,required"`
	Message string                                                                                             `json:"message,required"`
	JSON    accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeMessagesJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeMessages]
type accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeSuccess bool

const (
	AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeSuccessTrue AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseEnvelopeSuccess = true
)

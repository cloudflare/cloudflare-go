// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// OrganizationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
	DOH     *OrganizationDOHService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	r.DOH = NewOrganizationDOHService(opts...)
	return
}

// Sets up a Zero Trust organization for your account or zone.
func (r *OrganizationService) New(ctx context.Context, params OrganizationNewParams, opts ...option.RequestOption) (res *Organization, err error) {
	var env OrganizationNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates the configuration for your Zero Trust organization.
func (r *OrganizationService) Update(ctx context.Context, params OrganizationUpdateParams, opts ...option.RequestOption) (res *Organization, err error) {
	var env OrganizationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the configuration for your Zero Trust organization.
func (r *OrganizationService) List(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) (res *Organization, err error) {
	var env OrganizationListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Revokes a user's access across all applications.
func (r *OrganizationService) RevokeUsers(ctx context.Context, params OrganizationRevokeUsersParams, opts ...option.RequestOption) (res *OrganizationRevokeUsersResponse, err error) {
	var env OrganizationRevokeUsersResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/organizations/revoke_user", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type LoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string          `json:"text_color"`
	JSON      loginDesignJSON `json:"-"`
}

// loginDesignJSON contains the JSON metadata for the struct [LoginDesign]
type loginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loginDesignJSON) RawJSON() string {
	return r.raw
}

type LoginDesignParam struct {
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

func (r LoginDesignParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Organization struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWARP bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                    `json:"auto_redirect_to_identity"`
	CustomPages            OrganizationCustomPages `json:"custom_pages"`
	// Determines whether to deny all requests to Cloudflare-protected resources that
	// lack an associated Access application. If enabled, you must explicitly configure
	// an Access application and policy to allow traffic to your Cloudflare-protected
	// resources. For domains you want to be public across all subdomains, add the
	// domain to the `deny_unmatched_requests_exempted_zone_names` array.
	DenyUnmatchedRequests bool `json:"deny_unmatched_requests"`
	// Contains zone names to exempt from the `deny_unmatched_requests` feature.
	// Requests to a subdomain in an exempted zone will block unauthenticated traffic
	// by default if there is a configured Access application and policy that matches
	// the request.
	DenyUnmatchedRequestsExemptedZoneNames []string `json:"deny_unmatched_requests_exempted_zone_names"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUIReadOnly bool        `json:"is_ui_read_only"`
	LoginDesign  LoginDesign `json:"login_design"`
	// Configures multi-factor authentication (MFA) settings for an organization.
	MfaConfig OrganizationMfaConfig `json:"mfa_config"`
	// Configures PIV key requirements for MFA using hardware security keys.
	MfaPivKeyRequirements OrganizationMfaPivKeyRequirements `json:"mfa_piv_key_requirements"`
	// Determines whether global MFA settings apply to applications by default. The
	// organization must have MFA enabled with at least one authentication method and a
	// session duration configured. Note: 'allowed_authenticators' cannot only contain
	// 'piv_key' if the organization has any non-infrastructure applications because
	// PIV keys are only compatible with infrastructure apps.
	MfaRequiredForAllApps bool `json:"mfa_required_for_all_apps"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration string `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UIReadOnlyToggleReason string `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Minimum value for this setting is 1
	// month (730h). Must be in the format `300ms` or `2h45m`. Valid time units are:
	// `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WARPAuthSessionDuration string           `json:"warp_auth_session_duration"`
	JSON                    organizationJSON `json:"-"`
}

// organizationJSON contains the JSON metadata for the struct [Organization]
type organizationJSON struct {
	AllowAuthenticateViaWARP               apijson.Field
	AuthDomain                             apijson.Field
	AutoRedirectToIdentity                 apijson.Field
	CustomPages                            apijson.Field
	DenyUnmatchedRequests                  apijson.Field
	DenyUnmatchedRequestsExemptedZoneNames apijson.Field
	IsUIReadOnly                           apijson.Field
	LoginDesign                            apijson.Field
	MfaConfig                              apijson.Field
	MfaPivKeyRequirements                  apijson.Field
	MfaRequiredForAllApps                  apijson.Field
	Name                                   apijson.Field
	SessionDuration                        apijson.Field
	UIReadOnlyToggleReason                 apijson.Field
	UserSeatExpirationInactiveTime         apijson.Field
	WARPAuthSessionDuration                apijson.Field
	raw                                    string
	ExtraFields                            map[string]apijson.Field
}

func (r *Organization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationJSON) RawJSON() string {
	return r.raw
}

type OrganizationCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                      `json:"identity_denied"`
	JSON           organizationCustomPagesJSON `json:"-"`
}

// organizationCustomPagesJSON contains the JSON metadata for the struct
// [OrganizationCustomPages]
type organizationCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OrganizationCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationCustomPagesJSON) RawJSON() string {
	return r.raw
}

// Configures multi-factor authentication (MFA) settings for an organization.
type OrganizationMfaConfig struct {
	// Lists the MFA methods that users can authenticate with.
	AllowedAuthenticators []OrganizationMfaConfigAllowedAuthenticator `json:"allowed_authenticators"`
	// Allows a user to skip MFA via Authentication Method Reference (AMR) matching
	// when the AMR claim provided by the IdP the user used to authenticate contains
	// "mfa". Must be in minutes (m) or hours (h). Minimum: 0m. Maximum: 720h (30
	// days).
	AmrMatchingSessionDuration string `json:"amr_matching_session_duration"`
	// Specifies a Cloudflare List of required FIDO2 authenticator device AAGUIDs.
	RequiredAaguids string `json:"required_aaguids" format:"uuid"`
	// Defines the duration of an MFA session. Must be in minutes (m) or hours (h).
	// Minimum: 0m. Maximum: 720h (30 days). Examples:`5m` or `24h`.
	SessionDuration string                    `json:"session_duration"`
	JSON            organizationMfaConfigJSON `json:"-"`
}

// organizationMfaConfigJSON contains the JSON metadata for the struct
// [OrganizationMfaConfig]
type organizationMfaConfigJSON struct {
	AllowedAuthenticators      apijson.Field
	AmrMatchingSessionDuration apijson.Field
	RequiredAaguids            apijson.Field
	SessionDuration            apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *OrganizationMfaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMfaConfigJSON) RawJSON() string {
	return r.raw
}

type OrganizationMfaConfigAllowedAuthenticator string

const (
	OrganizationMfaConfigAllowedAuthenticatorTotp        OrganizationMfaConfigAllowedAuthenticator = "totp"
	OrganizationMfaConfigAllowedAuthenticatorBiometrics  OrganizationMfaConfigAllowedAuthenticator = "biometrics"
	OrganizationMfaConfigAllowedAuthenticatorSecurityKey OrganizationMfaConfigAllowedAuthenticator = "security_key"
	OrganizationMfaConfigAllowedAuthenticatorPivKey      OrganizationMfaConfigAllowedAuthenticator = "piv_key"
)

func (r OrganizationMfaConfigAllowedAuthenticator) IsKnown() bool {
	switch r {
	case OrganizationMfaConfigAllowedAuthenticatorTotp, OrganizationMfaConfigAllowedAuthenticatorBiometrics, OrganizationMfaConfigAllowedAuthenticatorSecurityKey, OrganizationMfaConfigAllowedAuthenticatorPivKey:
		return true
	}
	return false
}

// Configures PIV key requirements for MFA using hardware security keys.
type OrganizationMfaPivKeyRequirements struct {
	// Defines when a PIN is required to use the SSH key. Valid values: `never` (no PIN
	// required), `once` (PIN required once per session), `always` (PIN required for
	// each use).
	PinPolicy OrganizationMfaPivKeyRequirementsPinPolicy `json:"pin_policy"`
	// Requires the PIV key to be stored on a FIPS 140-2 Level 1 or higher validated
	// device.
	RequireFipsDevice bool `json:"require_fips_device"`
	// Specifies the allowed SSH key sizes in bits. Valid sizes depend on key type.
	// Ed25519 has a fixed key size and does not accept this parameter.
	SSHKeySize []OrganizationMfaPivKeyRequirementsSSHKeySize `json:"ssh_key_size"`
	// Specifies the allowed SSH key types. Valid values are `ecdsa`, `ed25519`, and
	// `rsa`.
	SSHKeyType []OrganizationMfaPivKeyRequirementsSSHKeyType `json:"ssh_key_type"`
	// Defines when physical touch is required to use the SSH key. Valid values:
	// `never` (no touch required), `always` (touch required for each use), `cached`
	// (touch cached for 15 seconds).
	TouchPolicy OrganizationMfaPivKeyRequirementsTouchPolicy `json:"touch_policy"`
	JSON        organizationMfaPivKeyRequirementsJSON        `json:"-"`
}

// organizationMfaPivKeyRequirementsJSON contains the JSON metadata for the struct
// [OrganizationMfaPivKeyRequirements]
type organizationMfaPivKeyRequirementsJSON struct {
	PinPolicy         apijson.Field
	RequireFipsDevice apijson.Field
	SSHKeySize        apijson.Field
	SSHKeyType        apijson.Field
	TouchPolicy       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrganizationMfaPivKeyRequirements) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMfaPivKeyRequirementsJSON) RawJSON() string {
	return r.raw
}

// Defines when a PIN is required to use the SSH key. Valid values: `never` (no PIN
// required), `once` (PIN required once per session), `always` (PIN required for
// each use).
type OrganizationMfaPivKeyRequirementsPinPolicy string

const (
	OrganizationMfaPivKeyRequirementsPinPolicyNever  OrganizationMfaPivKeyRequirementsPinPolicy = "never"
	OrganizationMfaPivKeyRequirementsPinPolicyOnce   OrganizationMfaPivKeyRequirementsPinPolicy = "once"
	OrganizationMfaPivKeyRequirementsPinPolicyAlways OrganizationMfaPivKeyRequirementsPinPolicy = "always"
)

func (r OrganizationMfaPivKeyRequirementsPinPolicy) IsKnown() bool {
	switch r {
	case OrganizationMfaPivKeyRequirementsPinPolicyNever, OrganizationMfaPivKeyRequirementsPinPolicyOnce, OrganizationMfaPivKeyRequirementsPinPolicyAlways:
		return true
	}
	return false
}

type OrganizationMfaPivKeyRequirementsSSHKeySize int64

const (
	OrganizationMfaPivKeyRequirementsSSHKeySize256  OrganizationMfaPivKeyRequirementsSSHKeySize = 256
	OrganizationMfaPivKeyRequirementsSSHKeySize384  OrganizationMfaPivKeyRequirementsSSHKeySize = 384
	OrganizationMfaPivKeyRequirementsSSHKeySize521  OrganizationMfaPivKeyRequirementsSSHKeySize = 521
	OrganizationMfaPivKeyRequirementsSSHKeySize2048 OrganizationMfaPivKeyRequirementsSSHKeySize = 2048
	OrganizationMfaPivKeyRequirementsSSHKeySize3072 OrganizationMfaPivKeyRequirementsSSHKeySize = 3072
	OrganizationMfaPivKeyRequirementsSSHKeySize4096 OrganizationMfaPivKeyRequirementsSSHKeySize = 4096
)

func (r OrganizationMfaPivKeyRequirementsSSHKeySize) IsKnown() bool {
	switch r {
	case OrganizationMfaPivKeyRequirementsSSHKeySize256, OrganizationMfaPivKeyRequirementsSSHKeySize384, OrganizationMfaPivKeyRequirementsSSHKeySize521, OrganizationMfaPivKeyRequirementsSSHKeySize2048, OrganizationMfaPivKeyRequirementsSSHKeySize3072, OrganizationMfaPivKeyRequirementsSSHKeySize4096:
		return true
	}
	return false
}

type OrganizationMfaPivKeyRequirementsSSHKeyType string

const (
	OrganizationMfaPivKeyRequirementsSSHKeyTypeEcdsa   OrganizationMfaPivKeyRequirementsSSHKeyType = "ecdsa"
	OrganizationMfaPivKeyRequirementsSSHKeyTypeEd25519 OrganizationMfaPivKeyRequirementsSSHKeyType = "ed25519"
	OrganizationMfaPivKeyRequirementsSSHKeyTypeRSA     OrganizationMfaPivKeyRequirementsSSHKeyType = "rsa"
)

func (r OrganizationMfaPivKeyRequirementsSSHKeyType) IsKnown() bool {
	switch r {
	case OrganizationMfaPivKeyRequirementsSSHKeyTypeEcdsa, OrganizationMfaPivKeyRequirementsSSHKeyTypeEd25519, OrganizationMfaPivKeyRequirementsSSHKeyTypeRSA:
		return true
	}
	return false
}

// Defines when physical touch is required to use the SSH key. Valid values:
// `never` (no touch required), `always` (touch required for each use), `cached`
// (touch cached for 15 seconds).
type OrganizationMfaPivKeyRequirementsTouchPolicy string

const (
	OrganizationMfaPivKeyRequirementsTouchPolicyNever  OrganizationMfaPivKeyRequirementsTouchPolicy = "never"
	OrganizationMfaPivKeyRequirementsTouchPolicyAlways OrganizationMfaPivKeyRequirementsTouchPolicy = "always"
	OrganizationMfaPivKeyRequirementsTouchPolicyCached OrganizationMfaPivKeyRequirementsTouchPolicy = "cached"
)

func (r OrganizationMfaPivKeyRequirementsTouchPolicy) IsKnown() bool {
	switch r {
	case OrganizationMfaPivKeyRequirementsTouchPolicyNever, OrganizationMfaPivKeyRequirementsTouchPolicyAlways, OrganizationMfaPivKeyRequirementsTouchPolicyCached:
		return true
	}
	return false
}

type OrganizationRevokeUsersResponse bool

const (
	OrganizationRevokeUsersResponseTrue  OrganizationRevokeUsersResponse = true
	OrganizationRevokeUsersResponseFalse OrganizationRevokeUsersResponse = false
)

func (r OrganizationRevokeUsersResponse) IsKnown() bool {
	switch r {
	case OrganizationRevokeUsersResponseTrue, OrganizationRevokeUsersResponseFalse:
		return true
	}
	return false
}

type OrganizationNewParams struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain" api:"required"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name" api:"required"`
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
	// Determines whether to deny all requests to Cloudflare-protected resources that
	// lack an associated Access application. If enabled, you must explicitly configure
	// an Access application and policy to allow traffic to your Cloudflare-protected
	// resources. For domains you want to be public across all subdomains, add the
	// domain to the `deny_unmatched_requests_exempted_zone_names` array.
	DenyUnmatchedRequests param.Field[bool] `json:"deny_unmatched_requests"`
	// Contains zone names to exempt from the `deny_unmatched_requests` feature.
	// Requests to a subdomain in an exempted zone will block unauthenticated traffic
	// by default if there is a configured Access application and policy that matches
	// the request.
	DenyUnmatchedRequestsExemptedZoneNames param.Field[[]string] `json:"deny_unmatched_requests_exempted_zone_names"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUIReadOnly param.Field[bool]             `json:"is_ui_read_only"`
	LoginDesign  param.Field[LoginDesignParam] `json:"login_design"`
	// Configures multi-factor authentication (MFA) settings for an organization.
	MfaConfig param.Field[OrganizationNewParamsMfaConfig] `json:"mfa_config"`
	// Configures PIV key requirements for MFA using hardware security keys.
	MfaPivKeyRequirements param.Field[OrganizationNewParamsMfaPivKeyRequirements] `json:"mfa_piv_key_requirements"`
	// Determines whether global MFA settings apply to applications by default. The
	// organization must have MFA enabled with at least one authentication method and a
	// session duration configured. Note: 'allowed_authenticators' cannot only contain
	// 'piv_key' if the organization has any non-infrastructure applications because
	// PIV keys are only compatible with infrastructure apps.
	MfaRequiredForAllApps param.Field[bool] `json:"mfa_required_for_all_apps"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UIReadOnlyToggleReason param.Field[string] `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Minimum value for this setting is 1
	// month (730h). Must be in the format `300ms` or `2h45m`. Valid time units are:
	// `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime param.Field[string] `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WARPAuthSessionDuration param.Field[string] `json:"warp_auth_session_duration"`
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures multi-factor authentication (MFA) settings for an organization.
type OrganizationNewParamsMfaConfig struct {
	// Lists the MFA methods that users can authenticate with.
	AllowedAuthenticators param.Field[[]OrganizationNewParamsMfaConfigAllowedAuthenticator] `json:"allowed_authenticators"`
	// Allows a user to skip MFA via Authentication Method Reference (AMR) matching
	// when the AMR claim provided by the IdP the user used to authenticate contains
	// "mfa". Must be in minutes (m) or hours (h). Minimum: 0m. Maximum: 720h (30
	// days).
	AmrMatchingSessionDuration param.Field[string] `json:"amr_matching_session_duration"`
	// Specifies a Cloudflare List of required FIDO2 authenticator device AAGUIDs.
	RequiredAaguids param.Field[string] `json:"required_aaguids" format:"uuid"`
	// Defines the duration of an MFA session. Must be in minutes (m) or hours (h).
	// Minimum: 0m. Maximum: 720h (30 days). Examples:`5m` or `24h`.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r OrganizationNewParamsMfaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationNewParamsMfaConfigAllowedAuthenticator string

const (
	OrganizationNewParamsMfaConfigAllowedAuthenticatorTotp        OrganizationNewParamsMfaConfigAllowedAuthenticator = "totp"
	OrganizationNewParamsMfaConfigAllowedAuthenticatorBiometrics  OrganizationNewParamsMfaConfigAllowedAuthenticator = "biometrics"
	OrganizationNewParamsMfaConfigAllowedAuthenticatorSecurityKey OrganizationNewParamsMfaConfigAllowedAuthenticator = "security_key"
	OrganizationNewParamsMfaConfigAllowedAuthenticatorPivKey      OrganizationNewParamsMfaConfigAllowedAuthenticator = "piv_key"
)

func (r OrganizationNewParamsMfaConfigAllowedAuthenticator) IsKnown() bool {
	switch r {
	case OrganizationNewParamsMfaConfigAllowedAuthenticatorTotp, OrganizationNewParamsMfaConfigAllowedAuthenticatorBiometrics, OrganizationNewParamsMfaConfigAllowedAuthenticatorSecurityKey, OrganizationNewParamsMfaConfigAllowedAuthenticatorPivKey:
		return true
	}
	return false
}

// Configures PIV key requirements for MFA using hardware security keys.
type OrganizationNewParamsMfaPivKeyRequirements struct {
	// Defines when a PIN is required to use the SSH key. Valid values: `never` (no PIN
	// required), `once` (PIN required once per session), `always` (PIN required for
	// each use).
	PinPolicy param.Field[OrganizationNewParamsMfaPivKeyRequirementsPinPolicy] `json:"pin_policy"`
	// Requires the PIV key to be stored on a FIPS 140-2 Level 1 or higher validated
	// device.
	RequireFipsDevice param.Field[bool] `json:"require_fips_device"`
	// Specifies the allowed SSH key sizes in bits. Valid sizes depend on key type.
	// Ed25519 has a fixed key size and does not accept this parameter.
	SSHKeySize param.Field[[]OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize] `json:"ssh_key_size"`
	// Specifies the allowed SSH key types. Valid values are `ecdsa`, `ed25519`, and
	// `rsa`.
	SSHKeyType param.Field[[]OrganizationNewParamsMfaPivKeyRequirementsSSHKeyType] `json:"ssh_key_type"`
	// Defines when physical touch is required to use the SSH key. Valid values:
	// `never` (no touch required), `always` (touch required for each use), `cached`
	// (touch cached for 15 seconds).
	TouchPolicy param.Field[OrganizationNewParamsMfaPivKeyRequirementsTouchPolicy] `json:"touch_policy"`
}

func (r OrganizationNewParamsMfaPivKeyRequirements) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines when a PIN is required to use the SSH key. Valid values: `never` (no PIN
// required), `once` (PIN required once per session), `always` (PIN required for
// each use).
type OrganizationNewParamsMfaPivKeyRequirementsPinPolicy string

const (
	OrganizationNewParamsMfaPivKeyRequirementsPinPolicyNever  OrganizationNewParamsMfaPivKeyRequirementsPinPolicy = "never"
	OrganizationNewParamsMfaPivKeyRequirementsPinPolicyOnce   OrganizationNewParamsMfaPivKeyRequirementsPinPolicy = "once"
	OrganizationNewParamsMfaPivKeyRequirementsPinPolicyAlways OrganizationNewParamsMfaPivKeyRequirementsPinPolicy = "always"
)

func (r OrganizationNewParamsMfaPivKeyRequirementsPinPolicy) IsKnown() bool {
	switch r {
	case OrganizationNewParamsMfaPivKeyRequirementsPinPolicyNever, OrganizationNewParamsMfaPivKeyRequirementsPinPolicyOnce, OrganizationNewParamsMfaPivKeyRequirementsPinPolicyAlways:
		return true
	}
	return false
}

type OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize int64

const (
	OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize256  OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize = 256
	OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize384  OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize = 384
	OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize521  OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize = 521
	OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize2048 OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize = 2048
	OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize3072 OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize = 3072
	OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize4096 OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize = 4096
)

func (r OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize) IsKnown() bool {
	switch r {
	case OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize256, OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize384, OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize521, OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize2048, OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize3072, OrganizationNewParamsMfaPivKeyRequirementsSSHKeySize4096:
		return true
	}
	return false
}

type OrganizationNewParamsMfaPivKeyRequirementsSSHKeyType string

const (
	OrganizationNewParamsMfaPivKeyRequirementsSSHKeyTypeEcdsa   OrganizationNewParamsMfaPivKeyRequirementsSSHKeyType = "ecdsa"
	OrganizationNewParamsMfaPivKeyRequirementsSSHKeyTypeEd25519 OrganizationNewParamsMfaPivKeyRequirementsSSHKeyType = "ed25519"
	OrganizationNewParamsMfaPivKeyRequirementsSSHKeyTypeRSA     OrganizationNewParamsMfaPivKeyRequirementsSSHKeyType = "rsa"
)

func (r OrganizationNewParamsMfaPivKeyRequirementsSSHKeyType) IsKnown() bool {
	switch r {
	case OrganizationNewParamsMfaPivKeyRequirementsSSHKeyTypeEcdsa, OrganizationNewParamsMfaPivKeyRequirementsSSHKeyTypeEd25519, OrganizationNewParamsMfaPivKeyRequirementsSSHKeyTypeRSA:
		return true
	}
	return false
}

// Defines when physical touch is required to use the SSH key. Valid values:
// `never` (no touch required), `always` (touch required for each use), `cached`
// (touch cached for 15 seconds).
type OrganizationNewParamsMfaPivKeyRequirementsTouchPolicy string

const (
	OrganizationNewParamsMfaPivKeyRequirementsTouchPolicyNever  OrganizationNewParamsMfaPivKeyRequirementsTouchPolicy = "never"
	OrganizationNewParamsMfaPivKeyRequirementsTouchPolicyAlways OrganizationNewParamsMfaPivKeyRequirementsTouchPolicy = "always"
	OrganizationNewParamsMfaPivKeyRequirementsTouchPolicyCached OrganizationNewParamsMfaPivKeyRequirementsTouchPolicy = "cached"
)

func (r OrganizationNewParamsMfaPivKeyRequirementsTouchPolicy) IsKnown() bool {
	switch r {
	case OrganizationNewParamsMfaPivKeyRequirementsTouchPolicyNever, OrganizationNewParamsMfaPivKeyRequirementsTouchPolicyAlways, OrganizationNewParamsMfaPivKeyRequirementsTouchPolicyCached:
		return true
	}
	return false
}

type OrganizationNewResponseEnvelope struct {
	Errors   []OrganizationNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []OrganizationNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OrganizationNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  Organization                           `json:"result"`
	JSON    organizationNewResponseEnvelopeJSON    `json:"-"`
}

// organizationNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [OrganizationNewResponseEnvelope]
type organizationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           OrganizationNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             organizationNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// organizationNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OrganizationNewResponseEnvelopeErrors]
type organizationNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    organizationNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// organizationNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [OrganizationNewResponseEnvelopeErrorsSource]
type organizationNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           OrganizationNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             organizationNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// organizationNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [OrganizationNewResponseEnvelopeMessages]
type organizationNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    organizationNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// organizationNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [OrganizationNewResponseEnvelopeMessagesSource]
type organizationNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OrganizationNewResponseEnvelopeSuccess bool

const (
	OrganizationNewResponseEnvelopeSuccessTrue OrganizationNewResponseEnvelopeSuccess = true
)

func (r OrganizationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OrganizationUpdateParams struct {
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
	AutoRedirectToIdentity param.Field[bool]                                `json:"auto_redirect_to_identity"`
	CustomPages            param.Field[OrganizationUpdateParamsCustomPages] `json:"custom_pages"`
	// Determines whether to deny all requests to Cloudflare-protected resources that
	// lack an associated Access application. If enabled, you must explicitly configure
	// an Access application and policy to allow traffic to your Cloudflare-protected
	// resources. For domains you want to be public across all subdomains, add the
	// domain to the `deny_unmatched_requests_exempted_zone_names` array.
	DenyUnmatchedRequests param.Field[bool] `json:"deny_unmatched_requests"`
	// Contains zone names to exempt from the `deny_unmatched_requests` feature.
	// Requests to a subdomain in an exempted zone will block unauthenticated traffic
	// by default if there is a configured Access application and policy that matches
	// the request.
	DenyUnmatchedRequestsExemptedZoneNames param.Field[[]string] `json:"deny_unmatched_requests_exempted_zone_names"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUIReadOnly param.Field[bool]             `json:"is_ui_read_only"`
	LoginDesign  param.Field[LoginDesignParam] `json:"login_design"`
	// Configures multi-factor authentication (MFA) settings for an organization.
	MfaConfig param.Field[OrganizationUpdateParamsMfaConfig] `json:"mfa_config"`
	// Configures PIV key requirements for MFA using hardware security keys.
	MfaPivKeyRequirements param.Field[OrganizationUpdateParamsMfaPivKeyRequirements] `json:"mfa_piv_key_requirements"`
	// Determines whether global MFA settings apply to applications by default. The
	// organization must have MFA enabled with at least one authentication method and a
	// session duration configured. Note: 'allowed_authenticators' cannot only contain
	// 'piv_key' if the organization has any non-infrastructure applications because
	// PIV keys are only compatible with infrastructure apps.
	MfaRequiredForAllApps param.Field[bool] `json:"mfa_required_for_all_apps"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UIReadOnlyToggleReason param.Field[string] `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Minimum value for this setting is 1
	// month (730h). Must be in the format `300ms` or `2h45m`. Valid time units are:
	// `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime param.Field[string] `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WARPAuthSessionDuration param.Field[string] `json:"warp_auth_session_duration"`
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationUpdateParamsCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden param.Field[string] `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied param.Field[string] `json:"identity_denied"`
}

func (r OrganizationUpdateParamsCustomPages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures multi-factor authentication (MFA) settings for an organization.
type OrganizationUpdateParamsMfaConfig struct {
	// Lists the MFA methods that users can authenticate with.
	AllowedAuthenticators param.Field[[]OrganizationUpdateParamsMfaConfigAllowedAuthenticator] `json:"allowed_authenticators"`
	// Allows a user to skip MFA via Authentication Method Reference (AMR) matching
	// when the AMR claim provided by the IdP the user used to authenticate contains
	// "mfa". Must be in minutes (m) or hours (h). Minimum: 0m. Maximum: 720h (30
	// days).
	AmrMatchingSessionDuration param.Field[string] `json:"amr_matching_session_duration"`
	// Specifies a Cloudflare List of required FIDO2 authenticator device AAGUIDs.
	RequiredAaguids param.Field[string] `json:"required_aaguids" format:"uuid"`
	// Defines the duration of an MFA session. Must be in minutes (m) or hours (h).
	// Minimum: 0m. Maximum: 720h (30 days). Examples:`5m` or `24h`.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r OrganizationUpdateParamsMfaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationUpdateParamsMfaConfigAllowedAuthenticator string

const (
	OrganizationUpdateParamsMfaConfigAllowedAuthenticatorTotp        OrganizationUpdateParamsMfaConfigAllowedAuthenticator = "totp"
	OrganizationUpdateParamsMfaConfigAllowedAuthenticatorBiometrics  OrganizationUpdateParamsMfaConfigAllowedAuthenticator = "biometrics"
	OrganizationUpdateParamsMfaConfigAllowedAuthenticatorSecurityKey OrganizationUpdateParamsMfaConfigAllowedAuthenticator = "security_key"
	OrganizationUpdateParamsMfaConfigAllowedAuthenticatorPivKey      OrganizationUpdateParamsMfaConfigAllowedAuthenticator = "piv_key"
)

func (r OrganizationUpdateParamsMfaConfigAllowedAuthenticator) IsKnown() bool {
	switch r {
	case OrganizationUpdateParamsMfaConfigAllowedAuthenticatorTotp, OrganizationUpdateParamsMfaConfigAllowedAuthenticatorBiometrics, OrganizationUpdateParamsMfaConfigAllowedAuthenticatorSecurityKey, OrganizationUpdateParamsMfaConfigAllowedAuthenticatorPivKey:
		return true
	}
	return false
}

// Configures PIV key requirements for MFA using hardware security keys.
type OrganizationUpdateParamsMfaPivKeyRequirements struct {
	// Defines when a PIN is required to use the SSH key. Valid values: `never` (no PIN
	// required), `once` (PIN required once per session), `always` (PIN required for
	// each use).
	PinPolicy param.Field[OrganizationUpdateParamsMfaPivKeyRequirementsPinPolicy] `json:"pin_policy"`
	// Requires the PIV key to be stored on a FIPS 140-2 Level 1 or higher validated
	// device.
	RequireFipsDevice param.Field[bool] `json:"require_fips_device"`
	// Specifies the allowed SSH key sizes in bits. Valid sizes depend on key type.
	// Ed25519 has a fixed key size and does not accept this parameter.
	SSHKeySize param.Field[[]OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize] `json:"ssh_key_size"`
	// Specifies the allowed SSH key types. Valid values are `ecdsa`, `ed25519`, and
	// `rsa`.
	SSHKeyType param.Field[[]OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeyType] `json:"ssh_key_type"`
	// Defines when physical touch is required to use the SSH key. Valid values:
	// `never` (no touch required), `always` (touch required for each use), `cached`
	// (touch cached for 15 seconds).
	TouchPolicy param.Field[OrganizationUpdateParamsMfaPivKeyRequirementsTouchPolicy] `json:"touch_policy"`
}

func (r OrganizationUpdateParamsMfaPivKeyRequirements) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines when a PIN is required to use the SSH key. Valid values: `never` (no PIN
// required), `once` (PIN required once per session), `always` (PIN required for
// each use).
type OrganizationUpdateParamsMfaPivKeyRequirementsPinPolicy string

const (
	OrganizationUpdateParamsMfaPivKeyRequirementsPinPolicyNever  OrganizationUpdateParamsMfaPivKeyRequirementsPinPolicy = "never"
	OrganizationUpdateParamsMfaPivKeyRequirementsPinPolicyOnce   OrganizationUpdateParamsMfaPivKeyRequirementsPinPolicy = "once"
	OrganizationUpdateParamsMfaPivKeyRequirementsPinPolicyAlways OrganizationUpdateParamsMfaPivKeyRequirementsPinPolicy = "always"
)

func (r OrganizationUpdateParamsMfaPivKeyRequirementsPinPolicy) IsKnown() bool {
	switch r {
	case OrganizationUpdateParamsMfaPivKeyRequirementsPinPolicyNever, OrganizationUpdateParamsMfaPivKeyRequirementsPinPolicyOnce, OrganizationUpdateParamsMfaPivKeyRequirementsPinPolicyAlways:
		return true
	}
	return false
}

type OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize int64

const (
	OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize256  OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize = 256
	OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize384  OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize = 384
	OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize521  OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize = 521
	OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize2048 OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize = 2048
	OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize3072 OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize = 3072
	OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize4096 OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize = 4096
)

func (r OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize) IsKnown() bool {
	switch r {
	case OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize256, OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize384, OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize521, OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize2048, OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize3072, OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeySize4096:
		return true
	}
	return false
}

type OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeyType string

const (
	OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeyTypeEcdsa   OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeyType = "ecdsa"
	OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeyTypeEd25519 OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeyType = "ed25519"
	OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeyTypeRSA     OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeyType = "rsa"
)

func (r OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeyType) IsKnown() bool {
	switch r {
	case OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeyTypeEcdsa, OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeyTypeEd25519, OrganizationUpdateParamsMfaPivKeyRequirementsSSHKeyTypeRSA:
		return true
	}
	return false
}

// Defines when physical touch is required to use the SSH key. Valid values:
// `never` (no touch required), `always` (touch required for each use), `cached`
// (touch cached for 15 seconds).
type OrganizationUpdateParamsMfaPivKeyRequirementsTouchPolicy string

const (
	OrganizationUpdateParamsMfaPivKeyRequirementsTouchPolicyNever  OrganizationUpdateParamsMfaPivKeyRequirementsTouchPolicy = "never"
	OrganizationUpdateParamsMfaPivKeyRequirementsTouchPolicyAlways OrganizationUpdateParamsMfaPivKeyRequirementsTouchPolicy = "always"
	OrganizationUpdateParamsMfaPivKeyRequirementsTouchPolicyCached OrganizationUpdateParamsMfaPivKeyRequirementsTouchPolicy = "cached"
)

func (r OrganizationUpdateParamsMfaPivKeyRequirementsTouchPolicy) IsKnown() bool {
	switch r {
	case OrganizationUpdateParamsMfaPivKeyRequirementsTouchPolicyNever, OrganizationUpdateParamsMfaPivKeyRequirementsTouchPolicyAlways, OrganizationUpdateParamsMfaPivKeyRequirementsTouchPolicyCached:
		return true
	}
	return false
}

type OrganizationUpdateResponseEnvelope struct {
	Errors   []OrganizationUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []OrganizationUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OrganizationUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  Organization                              `json:"result"`
	JSON    organizationUpdateResponseEnvelopeJSON    `json:"-"`
}

// organizationUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [OrganizationUpdateResponseEnvelope]
type organizationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           OrganizationUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             organizationUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// organizationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OrganizationUpdateResponseEnvelopeErrors]
type organizationUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    organizationUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// organizationUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [OrganizationUpdateResponseEnvelopeErrorsSource]
type organizationUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code" api:"required"`
	Message          string                                           `json:"message" api:"required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           OrganizationUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             organizationUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// organizationUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [OrganizationUpdateResponseEnvelopeMessages]
type organizationUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    organizationUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// organizationUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [OrganizationUpdateResponseEnvelopeMessagesSource]
type organizationUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OrganizationUpdateResponseEnvelopeSuccess bool

const (
	OrganizationUpdateResponseEnvelopeSuccessTrue OrganizationUpdateResponseEnvelopeSuccess = true
)

func (r OrganizationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OrganizationListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type OrganizationListResponseEnvelope struct {
	Errors   []OrganizationListResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []OrganizationListResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OrganizationListResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  Organization                            `json:"result"`
	JSON    organizationListResponseEnvelopeJSON    `json:"-"`
}

// organizationListResponseEnvelopeJSON contains the JSON metadata for the struct
// [OrganizationListResponseEnvelope]
type organizationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           OrganizationListResponseEnvelopeErrorsSource `json:"source"`
	JSON             organizationListResponseEnvelopeErrorsJSON   `json:"-"`
}

// organizationListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OrganizationListResponseEnvelopeErrors]
type organizationListResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    organizationListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// organizationListResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [OrganizationListResponseEnvelopeErrorsSource]
type organizationListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           OrganizationListResponseEnvelopeMessagesSource `json:"source"`
	JSON             organizationListResponseEnvelopeMessagesJSON   `json:"-"`
}

// organizationListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [OrganizationListResponseEnvelopeMessages]
type organizationListResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    organizationListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// organizationListResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [OrganizationListResponseEnvelopeMessagesSource]
type organizationListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OrganizationListResponseEnvelopeSuccess bool

const (
	OrganizationListResponseEnvelopeSuccessTrue OrganizationListResponseEnvelopeSuccess = true
)

func (r OrganizationListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OrganizationRevokeUsersParams struct {
	// The email of the user to revoke.
	Email param.Field[string] `json:"email" api:"required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// When set to `true`, all devices associated with the user will be revoked.
	QueryDevices param.Field[bool] `query:"devices"`
	// When set to `true`, all devices associated with the user will be revoked.
	BodyDevices param.Field[bool] `json:"devices"`
	// The uuid of the user to revoke.
	UserUID param.Field[string] `json:"user_uid"`
	// When set to `true`, the user will be required to re-authenticate to WARP for all
	// Gateway policies that enforce a WARP client session duration. When `false`, the
	// user’s WARP session will remain active
	WARPSessionReauth param.Field[bool] `json:"warp_session_reauth"`
}

func (r OrganizationRevokeUsersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [OrganizationRevokeUsersParams]'s query parameters as
// `url.Values`.
func (r OrganizationRevokeUsersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type OrganizationRevokeUsersResponseEnvelope struct {
	Result  OrganizationRevokeUsersResponse                `json:"result"`
	Success OrganizationRevokeUsersResponseEnvelopeSuccess `json:"success"`
	JSON    organizationRevokeUsersResponseEnvelopeJSON    `json:"-"`
}

// organizationRevokeUsersResponseEnvelopeJSON contains the JSON metadata for the
// struct [OrganizationRevokeUsersResponseEnvelope]
type organizationRevokeUsersResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRevokeUsersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationRevokeUsersResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationRevokeUsersResponseEnvelopeSuccess bool

const (
	OrganizationRevokeUsersResponseEnvelopeSuccessTrue  OrganizationRevokeUsersResponseEnvelopeSuccess = true
	OrganizationRevokeUsersResponseEnvelopeSuccessFalse OrganizationRevokeUsersResponseEnvelopeSuccess = false
)

func (r OrganizationRevokeUsersResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationRevokeUsersResponseEnvelopeSuccessTrue, OrganizationRevokeUsersResponseEnvelopeSuccessFalse:
		return true
	}
	return false
}

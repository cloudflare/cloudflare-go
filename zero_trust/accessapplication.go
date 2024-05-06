// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// AccessApplicationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessApplicationService] method
// instead.
type AccessApplicationService struct {
	Options          []option.RequestOption
	CAs              *AccessApplicationCAService
	UserPolicyChecks *AccessApplicationUserPolicyCheckService
	Policies         *AccessApplicationPolicyService
}

// NewAccessApplicationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessApplicationService(opts ...option.RequestOption) (r *AccessApplicationService) {
	r = &AccessApplicationService{}
	r.Options = opts
	r.CAs = NewAccessApplicationCAService(opts...)
	r.UserPolicyChecks = NewAccessApplicationUserPolicyCheckService(opts...)
	r.Policies = NewAccessApplicationPolicyService(opts...)
	return
}

// Adds a new application to Access.
func (r *AccessApplicationService) New(ctx context.Context, params AccessApplicationNewParams, opts ...option.RequestOption) (res *Application, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an Access application.
func (r *AccessApplicationService) Update(ctx context.Context, appID AppIDUnionParam, params AccessApplicationUpdateParams, opts ...option.RequestOption) (res *Application, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps/%v", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Access applications in an account or zone.
func (r *AccessApplicationService) List(ctx context.Context, query AccessApplicationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Application], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps", accountOrZone, accountOrZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all Access applications in an account or zone.
func (r *AccessApplicationService) ListAutoPaging(ctx context.Context, query AccessApplicationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Application] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes an application from Access.
func (r *AccessApplicationService) Delete(ctx context.Context, appID AppIDUnionParam, body AccessApplicationDeleteParams, opts ...option.RequestOption) (res *AccessApplicationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps/%v", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches information about an Access application.
func (r *AccessApplicationService) Get(ctx context.Context, appID AppIDUnionParam, query AccessApplicationGetParams, opts ...option.RequestOption) (res *Application, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps/%v", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Revokes all tokens issued for an application.
func (r *AccessApplicationService) RevokeTokens(ctx context.Context, appID AppIDUnionParam, body AccessApplicationRevokeTokensParams, opts ...option.RequestOption) (res *AccessApplicationRevokeTokensResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationRevokeTokensResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps/%v/revoke_tokens", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AllowedHeadersh = string

type AllowedHeadershParam = string

type AllowedIdpsh = string

type AllowedIdpshParam = string

type AllowedMethodsh string

const (
	AllowedMethodshGet     AllowedMethodsh = "GET"
	AllowedMethodshPost    AllowedMethodsh = "POST"
	AllowedMethodshHead    AllowedMethodsh = "HEAD"
	AllowedMethodshPut     AllowedMethodsh = "PUT"
	AllowedMethodshDelete  AllowedMethodsh = "DELETE"
	AllowedMethodshConnect AllowedMethodsh = "CONNECT"
	AllowedMethodshOptions AllowedMethodsh = "OPTIONS"
	AllowedMethodshTrace   AllowedMethodsh = "TRACE"
	AllowedMethodshPatch   AllowedMethodsh = "PATCH"
)

func (r AllowedMethodsh) IsKnown() bool {
	switch r {
	case AllowedMethodshGet, AllowedMethodshPost, AllowedMethodshHead, AllowedMethodshPut, AllowedMethodshDelete, AllowedMethodshConnect, AllowedMethodshOptions, AllowedMethodshTrace, AllowedMethodshPatch:
		return true
	}
	return false
}

type AllowedOriginsh = string

type AllowedOriginshParam = string

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AppIDUnionParam interface {
	ImplementsZeroTrustAppIDUnionParam()
}

type Application struct {
	// Audience tag.
	AUD       string    `json:"aud"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	ID        string    `json:"id"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP bool        `json:"allow_authenticate_via_warp"`
	AllowedIdPs              interface{} `json:"allowed_idps,required"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage string `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL string `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL string      `json:"custom_non_identity_deny_url"`
	CustomPages              interface{} `json:"custom_pages,required"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// Enables the binding cookie, which increases security against compromised
	// authorization tokens and CSRF attacks.
	EnableBindingCookie bool `json:"enable_binding_cookie"`
	// Enables the HttpOnly cookie attribute, which increases security against XSS
	// attacks.
	HTTPOnlyCookieAttribute bool `json:"http_only_cookie_attribute"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name string `json:"name"`
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string      `json:"same_site_cookie_attribute"`
	SelfHostedDomains       interface{} `json:"self_hosted_domains,required"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect bool `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial bool        `json:"skip_interstitial"`
	Tags             interface{} `json:"tags,required"`
	// The application type.
	Type    string          `json:"type"`
	SaaSApp interface{}     `json:"saas_app,required"`
	JSON    applicationJSON `json:"-"`
	union   ApplicationUnion
}

// applicationJSON contains the JSON metadata for the struct [Application]
type applicationJSON struct {
	AUD                      apijson.Field
	CreatedAt                apijson.Field
	ID                       apijson.Field
	UpdatedAt                apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	Domain                   apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SelfHostedDomains        apijson.Field
	ServiceAuth401Redirect   apijson.Field
	SessionDuration          apijson.Field
	SkipInterstitial         apijson.Field
	Tags                     apijson.Field
	Type                     apijson.Field
	SaaSApp                  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r applicationJSON) RawJSON() string {
	return r.raw
}

func (r *Application) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r Application) AsUnion() ApplicationUnion {
	return r.union
}

// Union satisfied by [zero_trust.ApplicationSelfHostedApplication],
// [zero_trust.ApplicationSaaSApplication],
// [zero_trust.ApplicationBrowserSSHApplication],
// [zero_trust.ApplicationBrowserVncApplication],
// [zero_trust.ApplicationAppLauncherApplication],
// [zero_trust.ApplicationDeviceEnrollmentPermissionsApplication],
// [zero_trust.ApplicationBrowserIsolationPermissionsApplication] or
// [zero_trust.ApplicationBookmarkApplication].
type ApplicationUnion interface {
	implementsZeroTrustApplication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ApplicationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationSelfHostedApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationSaaSApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationBrowserSSHApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationBrowserVncApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationAppLauncherApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationDeviceEnrollmentPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationBrowserIsolationPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationBookmarkApplication{}),
		},
	)
}

type ApplicationSelfHostedApplication struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain,required"`
	// The application type.
	Type string `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP bool `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdpsh `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
	CreatedAt              time.Time   `json:"created_at" format:"date-time"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage string `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL string `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL string `json:"custom_non_identity_deny_url"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages []CustomPagesh `json:"custom_pages"`
	// Enables the binding cookie, which increases security against compromised
	// authorization tokens and CSRF attacks.
	EnableBindingCookie bool `json:"enable_binding_cookie"`
	// Enables the HttpOnly cookie attribute, which increases security against XSS
	// attacks.
	HTTPOnlyCookieAttribute bool `json:"http_only_cookie_attribute"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name string `json:"name"`
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomainsh `json:"self_hosted_domains"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect bool `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial bool `json:"skip_interstitial"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags      []string                             `json:"tags"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      applicationSelfHostedApplicationJSON `json:"-"`
}

// applicationSelfHostedApplicationJSON contains the JSON metadata for the struct
// [ApplicationSelfHostedApplication]
type applicationSelfHostedApplicationJSON struct {
	Domain                   apijson.Field
	Type                     apijson.Field
	ID                       apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AUD                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SelfHostedDomains        apijson.Field
	ServiceAuth401Redirect   apijson.Field
	SessionDuration          apijson.Field
	SkipInterstitial         apijson.Field
	Tags                     apijson.Field
	UpdatedAt                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ApplicationSelfHostedApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationSelfHostedApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationSelfHostedApplication) implementsZeroTrustApplication() {}

type ApplicationSaaSApplication struct {
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdpsh `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages []CustomPagesh `json:"custom_pages"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name    string                            `json:"name"`
	SaaSApp ApplicationSaaSApplicationSaaSApp `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags []string `json:"tags"`
	// The application type.
	Type      string                         `json:"type"`
	UpdatedAt time.Time                      `json:"updated_at" format:"date-time"`
	JSON      applicationSaaSApplicationJSON `json:"-"`
}

// applicationSaaSApplicationJSON contains the JSON metadata for the struct
// [ApplicationSaaSApplication]
type applicationSaaSApplicationJSON struct {
	ID                     apijson.Field
	AllowedIdPs            apijson.Field
	AppLauncherVisible     apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	CustomPages            apijson.Field
	LogoURL                apijson.Field
	Name                   apijson.Field
	SaaSApp                apijson.Field
	Tags                   apijson.Field
	Type                   apijson.Field
	UpdatedAt              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ApplicationSaaSApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationSaaSApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationSaaSApplication) implementsZeroTrustApplication() {}

type ApplicationSaaSApplicationSaaSApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType ApplicationSaaSApplicationSaaSAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string      `json:"consumer_service_url"`
	CreatedAt          time.Time   `json:"created_at" format:"date-time"`
	CustomAttributes   interface{} `json:"custom_attributes,required"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdPEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat SaaSAppNameIDFormat `json:"name_id_format"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's
	// user identities into a NameID value for its SAML assertion. This expression
	// should evaluate to a singular string. The output of this expression can override
	// the `name_id_format` setting.
	NameIDTransformJsonata string `json:"name_id_transform_jsonata"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// A [JSONata] (https://jsonata.org/) expression that transforms an application's
	// user identities into attribute assertions in the SAML response. The expression
	// can transform id, email, name, and groups values. It can also transform fields
	// listed in the saml_attributes or oidc_fields of the identity provider used to
	// authenticate. The output of this expression must be a JSON object.
	SAMLAttributeTransformJsonata string `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SPEntityID string `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint string    `json:"sso_endpoint"`
	UpdatedAt   time.Time `json:"updated_at" format:"date-time"`
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string      `json:"client_secret"`
	CustomClaims interface{} `json:"custom_claims,required"`
	GrantTypes   interface{} `json:"grant_types,required"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string                                `json:"group_filter_regex"`
	RedirectURIs     interface{}                           `json:"redirect_uris,required"`
	Scopes           interface{}                           `json:"scopes,required"`
	JSON             applicationSaaSApplicationSaaSAppJSON `json:"-"`
	union            ApplicationSaaSApplicationSaaSAppUnion
}

// applicationSaaSApplicationSaaSAppJSON contains the JSON metadata for the struct
// [ApplicationSaaSApplicationSaaSApp]
type applicationSaaSApplicationSaaSAppJSON struct {
	AuthType                      apijson.Field
	ConsumerServiceURL            apijson.Field
	CreatedAt                     apijson.Field
	CustomAttributes              apijson.Field
	DefaultRelayState             apijson.Field
	IdPEntityID                   apijson.Field
	NameIDFormat                  apijson.Field
	NameIDTransformJsonata        apijson.Field
	PublicKey                     apijson.Field
	SAMLAttributeTransformJsonata apijson.Field
	SPEntityID                    apijson.Field
	SSOEndpoint                   apijson.Field
	UpdatedAt                     apijson.Field
	AppLauncherURL                apijson.Field
	ClientID                      apijson.Field
	ClientSecret                  apijson.Field
	CustomClaims                  apijson.Field
	GrantTypes                    apijson.Field
	GroupFilterRegex              apijson.Field
	RedirectURIs                  apijson.Field
	Scopes                        apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r applicationSaaSApplicationSaaSAppJSON) RawJSON() string {
	return r.raw
}

func (r *ApplicationSaaSApplicationSaaSApp) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r ApplicationSaaSApplicationSaaSApp) AsUnion() ApplicationSaaSApplicationSaaSAppUnion {
	return r.union
}

// Union satisfied by [zero_trust.SAMLSaaSApp] or
// [zero_trust.ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSApp].
type ApplicationSaaSApplicationSaaSAppUnion interface {
	implementsZeroTrustApplicationSaaSApplicationSaaSApp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ApplicationSaaSApplicationSaaSAppUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SAMLSaaSApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSApp{}),
		},
	)
}

type ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string                                                         `json:"client_secret"`
	CreatedAt    time.Time                                                      `json:"created_at" format:"date-time"`
	CustomClaims ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaims `json:"custom_claims"`
	// The OIDC flows supported by this application
	GrantTypes []ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppGrantType `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs []string `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes    []ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScope `json:"scopes"`
	UpdatedAt time.Time                                                 `json:"updated_at" format:"date-time"`
	JSON      applicationSaaSApplicationSaaSAppAccessOIDCSaaSAppJSON    `json:"-"`
}

// applicationSaaSApplicationSaaSAppAccessOIDCSaaSAppJSON contains the JSON
// metadata for the struct [ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSApp]
type applicationSaaSApplicationSaaSAppAccessOIDCSaaSAppJSON struct {
	AppLauncherURL   apijson.Field
	AuthType         apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	CreatedAt        apijson.Field
	CustomClaims     apijson.Field
	GrantTypes       apijson.Field
	GroupFilterRegex apijson.Field
	PublicKey        apijson.Field
	RedirectURIs     apijson.Field
	Scopes           apijson.Field
	UpdatedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationSaaSApplicationSaaSAppAccessOIDCSaaSAppJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSApp) implementsZeroTrustApplicationSaaSApplicationSaaSApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppAuthType string

const (
	ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppAuthTypeSAML ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppAuthType = "saml"
	ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppAuthTypeOIDC ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppAuthType = "oidc"
)

func (r ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppAuthType) IsKnown() bool {
	switch r {
	case ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppAuthTypeSAML, ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppAuthTypeOIDC:
		return true
	}
	return false
}

type ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaims struct {
	// The name of the claim.
	Name string `json:"name"`
	// A mapping from IdP ID to claim name.
	NameByIdP map[string]string `json:"name_by_idp"`
	// If the claim is required when building an OIDC token.
	Required bool `json:"required"`
	// The scope of the claim.
	Scope  ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScope  `json:"scope"`
	Source ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsSource `json:"source"`
	JSON   applicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsJSON   `json:"-"`
}

// applicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsJSON contains the
// JSON metadata for the struct
// [ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaims]
type applicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsJSON struct {
	Name        apijson.Field
	NameByIdP   apijson.Field
	Required    apijson.Field
	Scope       apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaims) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsJSON) RawJSON() string {
	return r.raw
}

// The scope of the claim.
type ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScope string

const (
	ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScopeGroups  ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScope = "groups"
	ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScopeProfile ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScope = "profile"
	ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScopeEmail   ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScope = "email"
	ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScopeOpenid  ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScope = "openid"
)

func (r ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScope) IsKnown() bool {
	switch r {
	case ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScopeGroups, ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScopeProfile, ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScopeEmail, ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScopeOpenid:
		return true
	}
	return false
}

type ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsSource struct {
	// The name of the IdP claim.
	Name string                                                                   `json:"name"`
	JSON applicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsSourceJSON `json:"-"`
}

// applicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsSourceJSON
// contains the JSON metadata for the struct
// [ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsSource]
type applicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsSourceJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsSourceJSON) RawJSON() string {
	return r.raw
}

type ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppGrantType string

const (
	ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppGrantTypeAuthorizationCode         ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppGrantType = "authorization_code"
	ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppGrantTypeAuthorizationCodeWithPkce ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppGrantType = "authorization_code_with_pkce"
)

func (r ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppGrantType) IsKnown() bool {
	switch r {
	case ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppGrantTypeAuthorizationCode, ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppGrantTypeAuthorizationCodeWithPkce:
		return true
	}
	return false
}

type ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScope string

const (
	ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScopeOpenid  ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScope = "openid"
	ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScopeGroups  ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScope = "groups"
	ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScopeEmail   ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScope = "email"
	ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScopeProfile ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScope = "profile"
)

func (r ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScope) IsKnown() bool {
	switch r {
	case ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScopeOpenid, ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScopeGroups, ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScopeEmail, ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScopeProfile:
		return true
	}
	return false
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type ApplicationSaaSApplicationSaaSAppAuthType string

const (
	ApplicationSaaSApplicationSaaSAppAuthTypeSAML ApplicationSaaSApplicationSaaSAppAuthType = "saml"
	ApplicationSaaSApplicationSaaSAppAuthTypeOIDC ApplicationSaaSApplicationSaaSAppAuthType = "oidc"
)

func (r ApplicationSaaSApplicationSaaSAppAuthType) IsKnown() bool {
	switch r {
	case ApplicationSaaSApplicationSaaSAppAuthTypeSAML, ApplicationSaaSApplicationSaaSAppAuthTypeOIDC:
		return true
	}
	return false
}

type ApplicationBrowserSSHApplication struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain,required"`
	// The application type.
	Type string `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP bool `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdpsh `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
	CreatedAt              time.Time   `json:"created_at" format:"date-time"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage string `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL string `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL string `json:"custom_non_identity_deny_url"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages []CustomPagesh `json:"custom_pages"`
	// Enables the binding cookie, which increases security against compromised
	// authorization tokens and CSRF attacks.
	EnableBindingCookie bool `json:"enable_binding_cookie"`
	// Enables the HttpOnly cookie attribute, which increases security against XSS
	// attacks.
	HTTPOnlyCookieAttribute bool `json:"http_only_cookie_attribute"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name string `json:"name"`
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomainsh `json:"self_hosted_domains"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect bool `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial bool `json:"skip_interstitial"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags      []string                             `json:"tags"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      applicationBrowserSSHApplicationJSON `json:"-"`
}

// applicationBrowserSSHApplicationJSON contains the JSON metadata for the struct
// [ApplicationBrowserSSHApplication]
type applicationBrowserSSHApplicationJSON struct {
	Domain                   apijson.Field
	Type                     apijson.Field
	ID                       apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AUD                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SelfHostedDomains        apijson.Field
	ServiceAuth401Redirect   apijson.Field
	SessionDuration          apijson.Field
	SkipInterstitial         apijson.Field
	Tags                     apijson.Field
	UpdatedAt                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ApplicationBrowserSSHApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationBrowserSSHApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationBrowserSSHApplication) implementsZeroTrustApplication() {}

type ApplicationBrowserVncApplication struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain,required"`
	// The application type.
	Type string `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP bool `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdpsh `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
	CreatedAt              time.Time   `json:"created_at" format:"date-time"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage string `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL string `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL string `json:"custom_non_identity_deny_url"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages []CustomPagesh `json:"custom_pages"`
	// Enables the binding cookie, which increases security against compromised
	// authorization tokens and CSRF attacks.
	EnableBindingCookie bool `json:"enable_binding_cookie"`
	// Enables the HttpOnly cookie attribute, which increases security against XSS
	// attacks.
	HTTPOnlyCookieAttribute bool `json:"http_only_cookie_attribute"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name string `json:"name"`
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomainsh `json:"self_hosted_domains"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect bool `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial bool `json:"skip_interstitial"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags      []string                             `json:"tags"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      applicationBrowserVncApplicationJSON `json:"-"`
}

// applicationBrowserVncApplicationJSON contains the JSON metadata for the struct
// [ApplicationBrowserVncApplication]
type applicationBrowserVncApplicationJSON struct {
	Domain                   apijson.Field
	Type                     apijson.Field
	ID                       apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AUD                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SelfHostedDomains        apijson.Field
	ServiceAuth401Redirect   apijson.Field
	SessionDuration          apijson.Field
	SkipInterstitial         apijson.Field
	Tags                     apijson.Field
	UpdatedAt                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ApplicationBrowserVncApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationBrowserVncApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationBrowserVncApplication) implementsZeroTrustApplication() {}

type ApplicationAppLauncherApplication struct {
	// The application type.
	Type ApplicationAppLauncherApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdpsh `json:"allowed_idps"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// The name of the application.
	Name string `json:"name"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string                                `json:"session_duration"`
	UpdatedAt       time.Time                             `json:"updated_at" format:"date-time"`
	JSON            applicationAppLauncherApplicationJSON `json:"-"`
}

// applicationAppLauncherApplicationJSON contains the JSON metadata for the struct
// [ApplicationAppLauncherApplication]
type applicationAppLauncherApplicationJSON struct {
	Type                   apijson.Field
	ID                     apijson.Field
	AllowedIdPs            apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	SessionDuration        apijson.Field
	UpdatedAt              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ApplicationAppLauncherApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationAppLauncherApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationAppLauncherApplication) implementsZeroTrustApplication() {}

// The application type.
type ApplicationAppLauncherApplicationType string

const (
	ApplicationAppLauncherApplicationTypeSelfHosted  ApplicationAppLauncherApplicationType = "self_hosted"
	ApplicationAppLauncherApplicationTypeSaaS        ApplicationAppLauncherApplicationType = "saas"
	ApplicationAppLauncherApplicationTypeSSH         ApplicationAppLauncherApplicationType = "ssh"
	ApplicationAppLauncherApplicationTypeVnc         ApplicationAppLauncherApplicationType = "vnc"
	ApplicationAppLauncherApplicationTypeAppLauncher ApplicationAppLauncherApplicationType = "app_launcher"
	ApplicationAppLauncherApplicationTypeWARP        ApplicationAppLauncherApplicationType = "warp"
	ApplicationAppLauncherApplicationTypeBiso        ApplicationAppLauncherApplicationType = "biso"
	ApplicationAppLauncherApplicationTypeBookmark    ApplicationAppLauncherApplicationType = "bookmark"
	ApplicationAppLauncherApplicationTypeDashSSO     ApplicationAppLauncherApplicationType = "dash_sso"
)

func (r ApplicationAppLauncherApplicationType) IsKnown() bool {
	switch r {
	case ApplicationAppLauncherApplicationTypeSelfHosted, ApplicationAppLauncherApplicationTypeSaaS, ApplicationAppLauncherApplicationTypeSSH, ApplicationAppLauncherApplicationTypeVnc, ApplicationAppLauncherApplicationTypeAppLauncher, ApplicationAppLauncherApplicationTypeWARP, ApplicationAppLauncherApplicationTypeBiso, ApplicationAppLauncherApplicationTypeBookmark, ApplicationAppLauncherApplicationTypeDashSSO:
		return true
	}
	return false
}

type ApplicationDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type ApplicationDeviceEnrollmentPermissionsApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdpsh `json:"allowed_idps"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// The name of the application.
	Name string `json:"name"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string                                                `json:"session_duration"`
	UpdatedAt       time.Time                                             `json:"updated_at" format:"date-time"`
	JSON            applicationDeviceEnrollmentPermissionsApplicationJSON `json:"-"`
}

// applicationDeviceEnrollmentPermissionsApplicationJSON contains the JSON metadata
// for the struct [ApplicationDeviceEnrollmentPermissionsApplication]
type applicationDeviceEnrollmentPermissionsApplicationJSON struct {
	Type                   apijson.Field
	ID                     apijson.Field
	AllowedIdPs            apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	SessionDuration        apijson.Field
	UpdatedAt              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ApplicationDeviceEnrollmentPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationDeviceEnrollmentPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationDeviceEnrollmentPermissionsApplication) implementsZeroTrustApplication() {}

// The application type.
type ApplicationDeviceEnrollmentPermissionsApplicationType string

const (
	ApplicationDeviceEnrollmentPermissionsApplicationTypeSelfHosted  ApplicationDeviceEnrollmentPermissionsApplicationType = "self_hosted"
	ApplicationDeviceEnrollmentPermissionsApplicationTypeSaaS        ApplicationDeviceEnrollmentPermissionsApplicationType = "saas"
	ApplicationDeviceEnrollmentPermissionsApplicationTypeSSH         ApplicationDeviceEnrollmentPermissionsApplicationType = "ssh"
	ApplicationDeviceEnrollmentPermissionsApplicationTypeVnc         ApplicationDeviceEnrollmentPermissionsApplicationType = "vnc"
	ApplicationDeviceEnrollmentPermissionsApplicationTypeAppLauncher ApplicationDeviceEnrollmentPermissionsApplicationType = "app_launcher"
	ApplicationDeviceEnrollmentPermissionsApplicationTypeWARP        ApplicationDeviceEnrollmentPermissionsApplicationType = "warp"
	ApplicationDeviceEnrollmentPermissionsApplicationTypeBiso        ApplicationDeviceEnrollmentPermissionsApplicationType = "biso"
	ApplicationDeviceEnrollmentPermissionsApplicationTypeBookmark    ApplicationDeviceEnrollmentPermissionsApplicationType = "bookmark"
	ApplicationDeviceEnrollmentPermissionsApplicationTypeDashSSO     ApplicationDeviceEnrollmentPermissionsApplicationType = "dash_sso"
)

func (r ApplicationDeviceEnrollmentPermissionsApplicationType) IsKnown() bool {
	switch r {
	case ApplicationDeviceEnrollmentPermissionsApplicationTypeSelfHosted, ApplicationDeviceEnrollmentPermissionsApplicationTypeSaaS, ApplicationDeviceEnrollmentPermissionsApplicationTypeSSH, ApplicationDeviceEnrollmentPermissionsApplicationTypeVnc, ApplicationDeviceEnrollmentPermissionsApplicationTypeAppLauncher, ApplicationDeviceEnrollmentPermissionsApplicationTypeWARP, ApplicationDeviceEnrollmentPermissionsApplicationTypeBiso, ApplicationDeviceEnrollmentPermissionsApplicationTypeBookmark, ApplicationDeviceEnrollmentPermissionsApplicationTypeDashSSO:
		return true
	}
	return false
}

type ApplicationBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type ApplicationBrowserIsolationPermissionsApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdpsh `json:"allowed_idps"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// The name of the application.
	Name string `json:"name"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string                                                `json:"session_duration"`
	UpdatedAt       time.Time                                             `json:"updated_at" format:"date-time"`
	JSON            applicationBrowserIsolationPermissionsApplicationJSON `json:"-"`
}

// applicationBrowserIsolationPermissionsApplicationJSON contains the JSON metadata
// for the struct [ApplicationBrowserIsolationPermissionsApplication]
type applicationBrowserIsolationPermissionsApplicationJSON struct {
	Type                   apijson.Field
	ID                     apijson.Field
	AllowedIdPs            apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	SessionDuration        apijson.Field
	UpdatedAt              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ApplicationBrowserIsolationPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationBrowserIsolationPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationBrowserIsolationPermissionsApplication) implementsZeroTrustApplication() {}

// The application type.
type ApplicationBrowserIsolationPermissionsApplicationType string

const (
	ApplicationBrowserIsolationPermissionsApplicationTypeSelfHosted  ApplicationBrowserIsolationPermissionsApplicationType = "self_hosted"
	ApplicationBrowserIsolationPermissionsApplicationTypeSaaS        ApplicationBrowserIsolationPermissionsApplicationType = "saas"
	ApplicationBrowserIsolationPermissionsApplicationTypeSSH         ApplicationBrowserIsolationPermissionsApplicationType = "ssh"
	ApplicationBrowserIsolationPermissionsApplicationTypeVnc         ApplicationBrowserIsolationPermissionsApplicationType = "vnc"
	ApplicationBrowserIsolationPermissionsApplicationTypeAppLauncher ApplicationBrowserIsolationPermissionsApplicationType = "app_launcher"
	ApplicationBrowserIsolationPermissionsApplicationTypeWARP        ApplicationBrowserIsolationPermissionsApplicationType = "warp"
	ApplicationBrowserIsolationPermissionsApplicationTypeBiso        ApplicationBrowserIsolationPermissionsApplicationType = "biso"
	ApplicationBrowserIsolationPermissionsApplicationTypeBookmark    ApplicationBrowserIsolationPermissionsApplicationType = "bookmark"
	ApplicationBrowserIsolationPermissionsApplicationTypeDashSSO     ApplicationBrowserIsolationPermissionsApplicationType = "dash_sso"
)

func (r ApplicationBrowserIsolationPermissionsApplicationType) IsKnown() bool {
	switch r {
	case ApplicationBrowserIsolationPermissionsApplicationTypeSelfHosted, ApplicationBrowserIsolationPermissionsApplicationTypeSaaS, ApplicationBrowserIsolationPermissionsApplicationTypeSSH, ApplicationBrowserIsolationPermissionsApplicationTypeVnc, ApplicationBrowserIsolationPermissionsApplicationTypeAppLauncher, ApplicationBrowserIsolationPermissionsApplicationTypeWARP, ApplicationBrowserIsolationPermissionsApplicationTypeBiso, ApplicationBrowserIsolationPermissionsApplicationTypeBookmark, ApplicationBrowserIsolationPermissionsApplicationTypeDashSSO:
		return true
	}
	return false
}

type ApplicationBookmarkApplication struct {
	// UUID
	ID string `json:"id"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD       string    `json:"aud"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The URL or domain of the bookmark.
	Domain string `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name string `json:"name"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags []string `json:"tags"`
	// The application type.
	Type      string                             `json:"type"`
	UpdatedAt time.Time                          `json:"updated_at" format:"date-time"`
	JSON      applicationBookmarkApplicationJSON `json:"-"`
}

// applicationBookmarkApplicationJSON contains the JSON metadata for the struct
// [ApplicationBookmarkApplication]
type applicationBookmarkApplicationJSON struct {
	ID                 apijson.Field
	AppLauncherVisible apijson.Field
	AUD                apijson.Field
	CreatedAt          apijson.Field
	Domain             apijson.Field
	LogoURL            apijson.Field
	Name               apijson.Field
	Tags               apijson.Field
	Type               apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ApplicationBookmarkApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationBookmarkApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationBookmarkApplication) implementsZeroTrustApplication() {}

type ApplicationParam struct {
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool]        `json:"allow_authenticate_via_warp"`
	AllowedIdPs              param.Field[interface{}] `json:"allowed_idps,required"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CORSHeaders            param.Field[CORSHeadersParam] `json:"cors_headers"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage param.Field[string] `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL param.Field[string] `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL param.Field[string]      `json:"custom_non_identity_deny_url"`
	CustomPages              param.Field[interface{}] `json:"custom_pages,required"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain"`
	// Enables the binding cookie, which increases security against compromised
	// authorization tokens and CSRF attacks.
	EnableBindingCookie param.Field[bool] `json:"enable_binding_cookie"`
	// Enables the HttpOnly cookie attribute, which increases security against XSS
	// attacks.
	HTTPOnlyCookieAttribute param.Field[bool] `json:"http_only_cookie_attribute"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name param.Field[string] `json:"name"`
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass param.Field[bool] `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string]      `json:"same_site_cookie_attribute"`
	SelfHostedDomains       param.Field[interface{}] `json:"self_hosted_domains,required"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect param.Field[bool] `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial param.Field[bool]        `json:"skip_interstitial"`
	Tags             param.Field[interface{}] `json:"tags,required"`
	// The application type.
	Type    param.Field[string]      `json:"type"`
	SaaSApp param.Field[interface{}] `json:"saas_app,required"`
}

func (r ApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationParam) implementsZeroTrustApplicationUnionParam() {}

// Satisfied by [zero_trust.ApplicationSelfHostedApplicationParam],
// [zero_trust.ApplicationSaaSApplicationParam],
// [zero_trust.ApplicationBrowserSSHApplicationParam],
// [zero_trust.ApplicationBrowserVncApplicationParam],
// [zero_trust.ApplicationAppLauncherApplicationParam],
// [zero_trust.ApplicationDeviceEnrollmentPermissionsApplicationParam],
// [zero_trust.ApplicationBrowserIsolationPermissionsApplicationParam],
// [zero_trust.ApplicationBookmarkApplicationParam], [ApplicationParam].
type ApplicationUnionParam interface {
	implementsZeroTrustApplicationUnionParam()
}

type ApplicationSelfHostedApplicationParam struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdpshParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CORSHeaders            param.Field[CORSHeadersParam] `json:"cors_headers"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage param.Field[string] `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL param.Field[string] `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL param.Field[string] `json:"custom_non_identity_deny_url"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages param.Field[[]CustomPageshParam] `json:"custom_pages"`
	// Enables the binding cookie, which increases security against compromised
	// authorization tokens and CSRF attacks.
	EnableBindingCookie param.Field[bool] `json:"enable_binding_cookie"`
	// Enables the HttpOnly cookie attribute, which increases security against XSS
	// attacks.
	HTTPOnlyCookieAttribute param.Field[bool] `json:"http_only_cookie_attribute"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name param.Field[string] `json:"name"`
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass param.Field[bool] `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]SelfHostedDomainshParam] `json:"self_hosted_domains"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect param.Field[bool] `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial param.Field[bool] `json:"skip_interstitial"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ApplicationSelfHostedApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationSelfHostedApplicationParam) implementsZeroTrustApplicationUnionParam() {}

type ApplicationSaaSApplicationParam struct {
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdpshParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages param.Field[[]CustomPageshParam] `json:"custom_pages"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name    param.Field[string]                                      `json:"name"`
	SaaSApp param.Field[ApplicationSaaSApplicationSaaSAppUnionParam] `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r ApplicationSaaSApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationSaaSApplicationParam) implementsZeroTrustApplicationUnionParam() {}

type ApplicationSaaSApplicationSaaSAppParam struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[ApplicationSaaSApplicationSaaSAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]      `json:"consumer_service_url"`
	CustomAttributes   param.Field[interface{}] `json:"custom_attributes,required"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdPEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[SaaSAppNameIDFormat] `json:"name_id_format"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's
	// user identities into a NameID value for its SAML assertion. This expression
	// should evaluate to a singular string. The output of this expression can override
	// the `name_id_format` setting.
	NameIDTransformJsonata param.Field[string] `json:"name_id_transform_jsonata"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// A [JSONata] (https://jsonata.org/) expression that transforms an application's
	// user identities into attribute assertions in the SAML response. The expression
	// can transform id, email, name, and groups values. It can also transform fields
	// listed in the saml_attributes or oidc_fields of the identity provider used to
	// authenticate. The output of this expression must be a JSON object.
	SAMLAttributeTransformJsonata param.Field[string] `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SPEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string] `json:"sso_endpoint"`
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string]      `json:"client_secret"`
	CustomClaims param.Field[interface{}] `json:"custom_claims,required"`
	GrantTypes   param.Field[interface{}] `json:"grant_types,required"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex param.Field[string]      `json:"group_filter_regex"`
	RedirectURIs     param.Field[interface{}] `json:"redirect_uris,required"`
	Scopes           param.Field[interface{}] `json:"scopes,required"`
}

func (r ApplicationSaaSApplicationSaaSAppParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationSaaSApplicationSaaSAppParam) implementsZeroTrustApplicationSaaSApplicationSaaSAppUnionParam() {
}

// Satisfied by [zero_trust.SAMLSaaSAppParam],
// [zero_trust.ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppParam],
// [ApplicationSaaSApplicationSaaSAppParam].
type ApplicationSaaSApplicationSaaSAppUnionParam interface {
	implementsZeroTrustApplicationSaaSApplicationSaaSAppUnionParam()
}

type ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppParam struct {
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType param.Field[ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppAuthType] `json:"auth_type"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string]                                                              `json:"client_secret"`
	CustomClaims param.Field[ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsParam] `json:"custom_claims"`
	// The OIDC flows supported by this application
	GrantTypes param.Field[[]ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppGrantType] `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex param.Field[string] `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs param.Field[[]string] `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes param.Field[[]ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppScope] `json:"scopes"`
}

func (r ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppParam) implementsZeroTrustApplicationSaaSApplicationSaaSAppUnionParam() {
}

type ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsParam struct {
	// The name of the claim.
	Name param.Field[string] `json:"name"`
	// A mapping from IdP ID to claim name.
	NameByIdP param.Field[map[string]string] `json:"name_by_idp"`
	// If the claim is required when building an OIDC token.
	Required param.Field[bool] `json:"required"`
	// The scope of the claim.
	Scope  param.Field[ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsScope]       `json:"scope"`
	Source param.Field[ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsSourceParam] `json:"source"`
}

func (r ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsSourceParam struct {
	// The name of the IdP claim.
	Name param.Field[string] `json:"name"`
}

func (r ApplicationSaaSApplicationSaaSAppAccessOIDCSaaSAppCustomClaimsSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ApplicationBrowserSSHApplicationParam struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdpshParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CORSHeaders            param.Field[CORSHeadersParam] `json:"cors_headers"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage param.Field[string] `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL param.Field[string] `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL param.Field[string] `json:"custom_non_identity_deny_url"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages param.Field[[]CustomPageshParam] `json:"custom_pages"`
	// Enables the binding cookie, which increases security against compromised
	// authorization tokens and CSRF attacks.
	EnableBindingCookie param.Field[bool] `json:"enable_binding_cookie"`
	// Enables the HttpOnly cookie attribute, which increases security against XSS
	// attacks.
	HTTPOnlyCookieAttribute param.Field[bool] `json:"http_only_cookie_attribute"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name param.Field[string] `json:"name"`
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass param.Field[bool] `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]SelfHostedDomainshParam] `json:"self_hosted_domains"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect param.Field[bool] `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial param.Field[bool] `json:"skip_interstitial"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ApplicationBrowserSSHApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationBrowserSSHApplicationParam) implementsZeroTrustApplicationUnionParam() {}

type ApplicationBrowserVncApplicationParam struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdpshParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CORSHeaders            param.Field[CORSHeadersParam] `json:"cors_headers"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage param.Field[string] `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL param.Field[string] `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL param.Field[string] `json:"custom_non_identity_deny_url"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages param.Field[[]CustomPageshParam] `json:"custom_pages"`
	// Enables the binding cookie, which increases security against compromised
	// authorization tokens and CSRF attacks.
	EnableBindingCookie param.Field[bool] `json:"enable_binding_cookie"`
	// Enables the HttpOnly cookie attribute, which increases security against XSS
	// attacks.
	HTTPOnlyCookieAttribute param.Field[bool] `json:"http_only_cookie_attribute"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name param.Field[string] `json:"name"`
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass param.Field[bool] `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]SelfHostedDomainshParam] `json:"self_hosted_domains"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect param.Field[bool] `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial param.Field[bool] `json:"skip_interstitial"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ApplicationBrowserVncApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationBrowserVncApplicationParam) implementsZeroTrustApplicationUnionParam() {}

type ApplicationAppLauncherApplicationParam struct {
	// The application type.
	Type param.Field[ApplicationAppLauncherApplicationType] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdpshParam] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r ApplicationAppLauncherApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationAppLauncherApplicationParam) implementsZeroTrustApplicationUnionParam() {}

type ApplicationDeviceEnrollmentPermissionsApplicationParam struct {
	// The application type.
	Type param.Field[ApplicationDeviceEnrollmentPermissionsApplicationType] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdpshParam] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r ApplicationDeviceEnrollmentPermissionsApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationDeviceEnrollmentPermissionsApplicationParam) implementsZeroTrustApplicationUnionParam() {
}

type ApplicationBrowserIsolationPermissionsApplicationParam struct {
	// The application type.
	Type param.Field[ApplicationBrowserIsolationPermissionsApplicationType] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdpshParam] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r ApplicationBrowserIsolationPermissionsApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationBrowserIsolationPermissionsApplicationParam) implementsZeroTrustApplicationUnionParam() {
}

type ApplicationBookmarkApplicationParam struct {
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// The URL or domain of the bookmark.
	Domain param.Field[string] `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name param.Field[string] `json:"name"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r ApplicationBookmarkApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApplicationBookmarkApplicationParam) implementsZeroTrustApplicationUnionParam() {}

type CORSHeaders struct {
	// Allows all HTTP request headers.
	AllowAllHeaders bool `json:"allow_all_headers"`
	// Allows all HTTP request methods.
	AllowAllMethods bool `json:"allow_all_methods"`
	// Allows all origins.
	AllowAllOrigins bool `json:"allow_all_origins"`
	// When set to `true`, includes credentials (cookies, authorization headers, or TLS
	// client certificates) with requests.
	AllowCredentials bool `json:"allow_credentials"`
	// Allowed HTTP request headers.
	AllowedHeaders []AllowedHeadersh `json:"allowed_headers"`
	// Allowed HTTP request methods.
	AllowedMethods []AllowedMethodsh `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []AllowedOriginsh `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64         `json:"max_age"`
	JSON   corsHeadersJSON `json:"-"`
}

// corsHeadersJSON contains the JSON metadata for the struct [CORSHeaders]
type corsHeadersJSON struct {
	AllowAllHeaders  apijson.Field
	AllowAllMethods  apijson.Field
	AllowAllOrigins  apijson.Field
	AllowCredentials apijson.Field
	AllowedHeaders   apijson.Field
	AllowedMethods   apijson.Field
	AllowedOrigins   apijson.Field
	MaxAge           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CORSHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corsHeadersJSON) RawJSON() string {
	return r.raw
}

type CORSHeadersParam struct {
	// Allows all HTTP request headers.
	AllowAllHeaders param.Field[bool] `json:"allow_all_headers"`
	// Allows all HTTP request methods.
	AllowAllMethods param.Field[bool] `json:"allow_all_methods"`
	// Allows all origins.
	AllowAllOrigins param.Field[bool] `json:"allow_all_origins"`
	// When set to `true`, includes credentials (cookies, authorization headers, or TLS
	// client certificates) with requests.
	AllowCredentials param.Field[bool] `json:"allow_credentials"`
	// Allowed HTTP request headers.
	AllowedHeaders param.Field[[]AllowedHeadershParam] `json:"allowed_headers"`
	// Allowed HTTP request methods.
	AllowedMethods param.Field[[]AllowedMethodsh] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]AllowedOriginshParam] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r CORSHeadersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomPagesh = string

type CustomPageshParam = string

// A globally unique name for an identity or service provider.
type SaaSAppNameFormat string

const (
	SaaSAppNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatUnspecified SaaSAppNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	SaaSAppNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatBasic       SaaSAppNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	SaaSAppNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatURI         SaaSAppNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

func (r SaaSAppNameFormat) IsKnown() bool {
	switch r {
	case SaaSAppNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatUnspecified, SaaSAppNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatBasic, SaaSAppNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatURI:
		return true
	}
	return false
}

// The format of the name identifier sent to the SaaS application.
type SaaSAppNameIDFormat string

const (
	SaaSAppNameIDFormatID    SaaSAppNameIDFormat = "id"
	SaaSAppNameIDFormatEmail SaaSAppNameIDFormat = "email"
)

func (r SaaSAppNameIDFormat) IsKnown() bool {
	switch r {
	case SaaSAppNameIDFormatID, SaaSAppNameIDFormatEmail:
		return true
	}
	return false
}

type SaaSAppSource struct {
	// The name of the IdP attribute.
	Name string `json:"name"`
	// A mapping from IdP ID to attribute name.
	NameByIdP map[string]string `json:"name_by_idp"`
	JSON      SaaSAppSourceJSON `json:"-"`
}

// SaaSAppSourceJSON contains the JSON metadata for the struct [SaaSAppSource]
type SaaSAppSourceJSON struct {
	Name        apijson.Field
	NameByIdP   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SaaSAppSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SaaSAppSourceJSON) RawJSON() string {
	return r.raw
}

type SaaSAppSourceParam struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
	// A mapping from IdP ID to attribute name.
	NameByIdP param.Field[map[string]string] `json:"name_by_idp"`
}

func (r SaaSAppSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SAMLSaaSApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType SAMLSaaSAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string                      `json:"consumer_service_url"`
	CreatedAt          time.Time                   `json:"created_at" format:"date-time"`
	CustomAttributes   SAMLSaaSAppCustomAttributes `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdPEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat SaaSAppNameIDFormat `json:"name_id_format"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's
	// user identities into a NameID value for its SAML assertion. This expression
	// should evaluate to a singular string. The output of this expression can override
	// the `name_id_format` setting.
	NameIDTransformJsonata string `json:"name_id_transform_jsonata"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// A [JSONata] (https://jsonata.org/) expression that transforms an application's
	// user identities into attribute assertions in the SAML response. The expression
	// can transform id, email, name, and groups values. It can also transform fields
	// listed in the saml_attributes or oidc_fields of the identity provider used to
	// authenticate. The output of this expression must be a JSON object.
	SAMLAttributeTransformJsonata string `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SPEntityID string `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint string          `json:"sso_endpoint"`
	UpdatedAt   time.Time       `json:"updated_at" format:"date-time"`
	JSON        samlSaaSAppJSON `json:"-"`
}

// samlSaaSAppJSON contains the JSON metadata for the struct [SAMLSaaSApp]
type samlSaaSAppJSON struct {
	AuthType                      apijson.Field
	ConsumerServiceURL            apijson.Field
	CreatedAt                     apijson.Field
	CustomAttributes              apijson.Field
	DefaultRelayState             apijson.Field
	IdPEntityID                   apijson.Field
	NameIDFormat                  apijson.Field
	NameIDTransformJsonata        apijson.Field
	PublicKey                     apijson.Field
	SAMLAttributeTransformJsonata apijson.Field
	SPEntityID                    apijson.Field
	SSOEndpoint                   apijson.Field
	UpdatedAt                     apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *SAMLSaaSApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlSaaSAppJSON) RawJSON() string {
	return r.raw
}

func (r SAMLSaaSApp) implementsZeroTrustApplicationSaaSApplicationSaaSApp() {}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type SAMLSaaSAppAuthType string

const (
	SAMLSaaSAppAuthTypeSAML SAMLSaaSAppAuthType = "saml"
	SAMLSaaSAppAuthTypeOIDC SAMLSaaSAppAuthType = "oidc"
)

func (r SAMLSaaSAppAuthType) IsKnown() bool {
	switch r {
	case SAMLSaaSAppAuthTypeSAML, SAMLSaaSAppAuthTypeOIDC:
		return true
	}
	return false
}

type SAMLSaaSAppCustomAttributes struct {
	// The SAML FriendlyName of the attribute.
	FriendlyName string `json:"friendly_name"`
	// The name of the attribute.
	Name string `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat SaaSAppNameFormat `json:"name_format"`
	// If the attribute is required when building a SAML assertion.
	Required bool                            `json:"required"`
	Source   SaaSAppSource                   `json:"source"`
	JSON     samlSaaSAppCustomAttributesJSON `json:"-"`
}

// samlSaaSAppCustomAttributesJSON contains the JSON metadata for the struct
// [SAMLSaaSAppCustomAttributes]
type samlSaaSAppCustomAttributesJSON struct {
	FriendlyName apijson.Field
	Name         apijson.Field
	NameFormat   apijson.Field
	Required     apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SAMLSaaSAppCustomAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlSaaSAppCustomAttributesJSON) RawJSON() string {
	return r.raw
}

type SAMLSaaSAppParam struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[SAMLSaaSAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                           `json:"consumer_service_url"`
	CustomAttributes   param.Field[SAMLSaaSAppCustomAttributesParam] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdPEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[SaaSAppNameIDFormat] `json:"name_id_format"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's
	// user identities into a NameID value for its SAML assertion. This expression
	// should evaluate to a singular string. The output of this expression can override
	// the `name_id_format` setting.
	NameIDTransformJsonata param.Field[string] `json:"name_id_transform_jsonata"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// A [JSONata] (https://jsonata.org/) expression that transforms an application's
	// user identities into attribute assertions in the SAML response. The expression
	// can transform id, email, name, and groups values. It can also transform fields
	// listed in the saml_attributes or oidc_fields of the identity provider used to
	// authenticate. The output of this expression must be a JSON object.
	SAMLAttributeTransformJsonata param.Field[string] `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SPEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string] `json:"sso_endpoint"`
}

func (r SAMLSaaSAppParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SAMLSaaSAppParam) implementsZeroTrustApplicationSaaSApplicationSaaSAppUnionParam() {}

type SAMLSaaSAppCustomAttributesParam struct {
	// The SAML FriendlyName of the attribute.
	FriendlyName param.Field[string] `json:"friendly_name"`
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[SaaSAppNameFormat] `json:"name_format"`
	// If the attribute is required when building a SAML assertion.
	Required param.Field[bool]               `json:"required"`
	Source   param.Field[SaaSAppSourceParam] `json:"source"`
}

func (r SAMLSaaSAppCustomAttributesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SelfHostedDomainsh = string

type SelfHostedDomainshParam = string

type AccessApplicationDeleteResponse struct {
	// UUID
	ID   string                              `json:"id"`
	JSON accessApplicationDeleteResponseJSON `json:"-"`
}

// accessApplicationDeleteResponseJSON contains the JSON metadata for the struct
// [AccessApplicationDeleteResponse]
type accessApplicationDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationRevokeTokensResponse = interface{}

type AccessApplicationNewParams struct {
	Application ApplicationUnionParam `json:"application,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r AccessApplicationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Application)
}

type AccessApplicationNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessApplicationNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Application                                 `json:"result"`
	JSON    accessApplicationNewResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationNewResponseEnvelope]
type accessApplicationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationNewResponseEnvelopeSuccess bool

const (
	AccessApplicationNewResponseEnvelopeSuccessTrue AccessApplicationNewResponseEnvelopeSuccess = true
)

func (r AccessApplicationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationUpdateParams struct {
	Application ApplicationUnionParam `json:"application,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r AccessApplicationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Application)
}

type AccessApplicationUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessApplicationUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Application                                    `json:"result"`
	JSON    accessApplicationUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationUpdateResponseEnvelope]
type accessApplicationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationUpdateResponseEnvelopeSuccess bool

const (
	AccessApplicationUpdateResponseEnvelopeSuccessTrue AccessApplicationUpdateResponseEnvelopeSuccess = true
)

func (r AccessApplicationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessApplicationDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationDeleteResponse                `json:"result"`
	JSON    accessApplicationDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationDeleteResponseEnvelope]
type accessApplicationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationDeleteResponseEnvelopeSuccess bool

const (
	AccessApplicationDeleteResponseEnvelopeSuccessTrue AccessApplicationDeleteResponseEnvelopeSuccess = true
)

func (r AccessApplicationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessApplicationGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Application                                 `json:"result"`
	JSON    accessApplicationGetResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationGetResponseEnvelope]
type accessApplicationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationGetResponseEnvelopeSuccess bool

const (
	AccessApplicationGetResponseEnvelopeSuccessTrue AccessApplicationGetResponseEnvelopeSuccess = true
)

func (r AccessApplicationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationRevokeTokensParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationRevokeTokensResponseEnvelope struct {
	Result  AccessApplicationRevokeTokensResponse                `json:"result,nullable"`
	Success AccessApplicationRevokeTokensResponseEnvelopeSuccess `json:"success"`
	JSON    accessApplicationRevokeTokensResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationRevokeTokensResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessApplicationRevokeTokensResponseEnvelope]
type accessApplicationRevokeTokensResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationRevokeTokensResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationRevokeTokensResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationRevokeTokensResponseEnvelopeSuccess bool

const (
	AccessApplicationRevokeTokensResponseEnvelopeSuccessTrue  AccessApplicationRevokeTokensResponseEnvelopeSuccess = true
	AccessApplicationRevokeTokensResponseEnvelopeSuccessFalse AccessApplicationRevokeTokensResponseEnvelopeSuccess = false
)

func (r AccessApplicationRevokeTokensResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationRevokeTokensResponseEnvelopeSuccessTrue, AccessApplicationRevokeTokensResponseEnvelopeSuccessFalse:
		return true
	}
	return false
}

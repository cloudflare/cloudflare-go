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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
	if params.getAccountID().Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.getAccountID()
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.getZoneID()
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
	if params.getAccountID().Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.getAccountID()
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.getZoneID()
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
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

type AllowedHeadersItem = string

type AllowedHeadersItemParam = string

type AllowedIDPsItem = string

type AllowedIDPsItemParam = string

type AllowedMethodsItem string

const (
	AllowedMethodsItemGet     AllowedMethodsItem = "GET"
	AllowedMethodsItemPost    AllowedMethodsItem = "POST"
	AllowedMethodsItemHead    AllowedMethodsItem = "HEAD"
	AllowedMethodsItemPut     AllowedMethodsItem = "PUT"
	AllowedMethodsItemDelete  AllowedMethodsItem = "DELETE"
	AllowedMethodsItemConnect AllowedMethodsItem = "CONNECT"
	AllowedMethodsItemOptions AllowedMethodsItem = "OPTIONS"
	AllowedMethodsItemTrace   AllowedMethodsItem = "TRACE"
	AllowedMethodsItemPatch   AllowedMethodsItem = "PATCH"
)

func (r AllowedMethodsItem) IsKnown() bool {
	switch r {
	case AllowedMethodsItemGet, AllowedMethodsItemPost, AllowedMethodsItemHead, AllowedMethodsItemPut, AllowedMethodsItemDelete, AllowedMethodsItemConnect, AllowedMethodsItemOptions, AllowedMethodsItemTrace, AllowedMethodsItemPatch:
		return true
	}
	return false
}

type AllowedOriginsItem = string

type AllowedOriginsItemParam = string

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AppIDUnionParam interface {
	ImplementsZeroTrustAppIDUnionParam()
}

type Application struct {
	// Audience tag.
	Aud       string    `json:"aud"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	ID        string    `json:"id"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP bool        `json:"allow_authenticate_via_warp"`
	AllowedIDPs              interface{} `json:"allowed_idps,required"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CorsHeaders            CorsHeaders `json:"cors_headers"`
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
	SaasApp interface{}     `json:"saas_app,required"`
	JSON    applicationJSON `json:"-"`
	union   ApplicationUnion
}

// applicationJSON contains the JSON metadata for the struct [Application]
type applicationJSON struct {
	Aud                      apijson.Field
	CreatedAt                apijson.Field
	ID                       apijson.Field
	UpdatedAt                apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIDPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CorsHeaders              apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	Domain                   apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	PathCookieAttribute      apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SelfHostedDomains        apijson.Field
	ServiceAuth401Redirect   apijson.Field
	SessionDuration          apijson.Field
	SkipInterstitial         apijson.Field
	Tags                     apijson.Field
	Type                     apijson.Field
	SaasApp                  apijson.Field
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
	AllowedIDPs []AllowedIDPsItem `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	Aud string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CorsHeaders            CorsHeaders `json:"cors_headers"`
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
	CustomPages []CustomPagesItem `json:"custom_pages"`
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
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomainsItem `json:"self_hosted_domains"`
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
	AllowedIDPs              apijson.Field
	AppLauncherVisible       apijson.Field
	Aud                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CorsHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
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
	AllowedIDPs []AllowedIDPsItem `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	Aud string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages []CustomPagesItem `json:"custom_pages"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name    string                            `json:"name"`
	SaasApp ApplicationSaaSApplicationSaasApp `json:"saas_app"`
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
	AllowedIDPs            apijson.Field
	AppLauncherVisible     apijson.Field
	Aud                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	CustomPages            apijson.Field
	LogoURL                apijson.Field
	Name                   apijson.Field
	SaasApp                apijson.Field
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

type ApplicationSaaSApplicationSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType ApplicationSaaSApplicationSaasAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string      `json:"consumer_service_url"`
	CreatedAt          time.Time   `json:"created_at" format:"date-time"`
	CustomAttributes   interface{} `json:"custom_attributes,required"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IDPEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat ApplicationSaaSApplicationSaasAppNameIDFormat `json:"name_id_format"`
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
	SamlAttributeTransformJsonata string `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SpEntityID string `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint string    `json:"sso_endpoint"`
	UpdatedAt   time.Time `json:"updated_at" format:"date-time"`
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string      `json:"client_secret"`
	GrantTypes   interface{} `json:"grant_types,required"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string                                `json:"group_filter_regex"`
	RedirectURIs     interface{}                           `json:"redirect_uris,required"`
	Scopes           interface{}                           `json:"scopes,required"`
	JSON             applicationSaaSApplicationSaasAppJSON `json:"-"`
	union            ApplicationSaaSApplicationSaasAppUnion
}

// applicationSaaSApplicationSaasAppJSON contains the JSON metadata for the struct
// [ApplicationSaaSApplicationSaasApp]
type applicationSaaSApplicationSaasAppJSON struct {
	AuthType                      apijson.Field
	ConsumerServiceURL            apijson.Field
	CreatedAt                     apijson.Field
	CustomAttributes              apijson.Field
	DefaultRelayState             apijson.Field
	IDPEntityID                   apijson.Field
	NameIDFormat                  apijson.Field
	NameIDTransformJsonata        apijson.Field
	PublicKey                     apijson.Field
	SamlAttributeTransformJsonata apijson.Field
	SpEntityID                    apijson.Field
	SSOEndpoint                   apijson.Field
	UpdatedAt                     apijson.Field
	AppLauncherURL                apijson.Field
	ClientID                      apijson.Field
	ClientSecret                  apijson.Field
	GrantTypes                    apijson.Field
	GroupFilterRegex              apijson.Field
	RedirectURIs                  apijson.Field
	Scopes                        apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r applicationSaaSApplicationSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r *ApplicationSaaSApplicationSaasApp) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r ApplicationSaaSApplicationSaasApp) AsUnion() ApplicationSaaSApplicationSaasAppUnion {
	return r.union
}

// Union satisfied by [zero_trust.SamlSaasApp] or
// [zero_trust.ApplicationSaaSApplicationSaasAppAccessOidcSaasApp].
type ApplicationSaaSApplicationSaasAppUnion interface {
	implementsZeroTrustApplicationSaaSApplicationSaasApp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ApplicationSaaSApplicationSaasAppUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SamlSaasApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationSaaSApplicationSaasAppAccessOidcSaasApp{}),
		},
	)
}

type ApplicationSaaSApplicationSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType ApplicationSaaSApplicationSaasAppAccessOidcSaasAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The OIDC flows supported by this application
	GrantTypes []ApplicationSaaSApplicationSaasAppAccessOidcSaasAppGrantType `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs []string `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes    []ApplicationSaaSApplicationSaasAppAccessOidcSaasAppScope `json:"scopes"`
	UpdatedAt time.Time                                                 `json:"updated_at" format:"date-time"`
	JSON      applicationSaaSApplicationSaasAppAccessOidcSaasAppJSON    `json:"-"`
}

// applicationSaaSApplicationSaasAppAccessOidcSaasAppJSON contains the JSON
// metadata for the struct [ApplicationSaaSApplicationSaasAppAccessOidcSaasApp]
type applicationSaaSApplicationSaasAppAccessOidcSaasAppJSON struct {
	AppLauncherURL   apijson.Field
	AuthType         apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	CreatedAt        apijson.Field
	GrantTypes       apijson.Field
	GroupFilterRegex apijson.Field
	PublicKey        apijson.Field
	RedirectURIs     apijson.Field
	Scopes           apijson.Field
	UpdatedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationSaaSApplicationSaasAppAccessOidcSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationSaaSApplicationSaasAppAccessOidcSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationSaaSApplicationSaasAppAccessOidcSaasApp) implementsZeroTrustApplicationSaaSApplicationSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type ApplicationSaaSApplicationSaasAppAccessOidcSaasAppAuthType string

const (
	ApplicationSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml ApplicationSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "saml"
	ApplicationSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc ApplicationSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "oidc"
)

func (r ApplicationSaaSApplicationSaasAppAccessOidcSaasAppAuthType) IsKnown() bool {
	switch r {
	case ApplicationSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml, ApplicationSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc:
		return true
	}
	return false
}

type ApplicationSaaSApplicationSaasAppAccessOidcSaasAppGrantType string

const (
	ApplicationSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         ApplicationSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	ApplicationSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce ApplicationSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

func (r ApplicationSaaSApplicationSaasAppAccessOidcSaasAppGrantType) IsKnown() bool {
	switch r {
	case ApplicationSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode, ApplicationSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce:
		return true
	}
	return false
}

type ApplicationSaaSApplicationSaasAppAccessOidcSaasAppScope string

const (
	ApplicationSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid  ApplicationSaaSApplicationSaasAppAccessOidcSaasAppScope = "openid"
	ApplicationSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups  ApplicationSaaSApplicationSaasAppAccessOidcSaasAppScope = "groups"
	ApplicationSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail   ApplicationSaaSApplicationSaasAppAccessOidcSaasAppScope = "email"
	ApplicationSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile ApplicationSaaSApplicationSaasAppAccessOidcSaasAppScope = "profile"
)

func (r ApplicationSaaSApplicationSaasAppAccessOidcSaasAppScope) IsKnown() bool {
	switch r {
	case ApplicationSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid, ApplicationSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups, ApplicationSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail, ApplicationSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile:
		return true
	}
	return false
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type ApplicationSaaSApplicationSaasAppAuthType string

const (
	ApplicationSaaSApplicationSaasAppAuthTypeSaml ApplicationSaaSApplicationSaasAppAuthType = "saml"
	ApplicationSaaSApplicationSaasAppAuthTypeOidc ApplicationSaaSApplicationSaasAppAuthType = "oidc"
)

func (r ApplicationSaaSApplicationSaasAppAuthType) IsKnown() bool {
	switch r {
	case ApplicationSaaSApplicationSaasAppAuthTypeSaml, ApplicationSaaSApplicationSaasAppAuthTypeOidc:
		return true
	}
	return false
}

// The format of the name identifier sent to the SaaS application.
type ApplicationSaaSApplicationSaasAppNameIDFormat string

const (
	ApplicationSaaSApplicationSaasAppNameIDFormatID    ApplicationSaaSApplicationSaasAppNameIDFormat = "id"
	ApplicationSaaSApplicationSaasAppNameIDFormatEmail ApplicationSaaSApplicationSaasAppNameIDFormat = "email"
)

func (r ApplicationSaaSApplicationSaasAppNameIDFormat) IsKnown() bool {
	switch r {
	case ApplicationSaaSApplicationSaasAppNameIDFormatID, ApplicationSaaSApplicationSaasAppNameIDFormatEmail:
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
	AllowedIDPs []AllowedIDPsItem `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	Aud string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CorsHeaders            CorsHeaders `json:"cors_headers"`
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
	CustomPages []CustomPagesItem `json:"custom_pages"`
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
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomainsItem `json:"self_hosted_domains"`
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
	AllowedIDPs              apijson.Field
	AppLauncherVisible       apijson.Field
	Aud                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CorsHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
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
	AllowedIDPs []AllowedIDPsItem `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	Aud string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CorsHeaders            CorsHeaders `json:"cors_headers"`
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
	CustomPages []CustomPagesItem `json:"custom_pages"`
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
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomainsItem `json:"self_hosted_domains"`
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
	AllowedIDPs              apijson.Field
	AppLauncherVisible       apijson.Field
	Aud                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CorsHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
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
	AllowedIDPs []AllowedIDPsItem `json:"allowed_idps"`
	// Audience tag.
	Aud string `json:"aud"`
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
	AllowedIDPs            apijson.Field
	Aud                    apijson.Field
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
	ApplicationAppLauncherApplicationTypeSaas        ApplicationAppLauncherApplicationType = "saas"
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
	case ApplicationAppLauncherApplicationTypeSelfHosted, ApplicationAppLauncherApplicationTypeSaas, ApplicationAppLauncherApplicationTypeSSH, ApplicationAppLauncherApplicationTypeVnc, ApplicationAppLauncherApplicationTypeAppLauncher, ApplicationAppLauncherApplicationTypeWARP, ApplicationAppLauncherApplicationTypeBiso, ApplicationAppLauncherApplicationTypeBookmark, ApplicationAppLauncherApplicationTypeDashSSO:
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
	AllowedIDPs []AllowedIDPsItem `json:"allowed_idps"`
	// Audience tag.
	Aud string `json:"aud"`
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
	AllowedIDPs            apijson.Field
	Aud                    apijson.Field
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
	ApplicationDeviceEnrollmentPermissionsApplicationTypeSaas        ApplicationDeviceEnrollmentPermissionsApplicationType = "saas"
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
	case ApplicationDeviceEnrollmentPermissionsApplicationTypeSelfHosted, ApplicationDeviceEnrollmentPermissionsApplicationTypeSaas, ApplicationDeviceEnrollmentPermissionsApplicationTypeSSH, ApplicationDeviceEnrollmentPermissionsApplicationTypeVnc, ApplicationDeviceEnrollmentPermissionsApplicationTypeAppLauncher, ApplicationDeviceEnrollmentPermissionsApplicationTypeWARP, ApplicationDeviceEnrollmentPermissionsApplicationTypeBiso, ApplicationDeviceEnrollmentPermissionsApplicationTypeBookmark, ApplicationDeviceEnrollmentPermissionsApplicationTypeDashSSO:
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
	AllowedIDPs []AllowedIDPsItem `json:"allowed_idps"`
	// Audience tag.
	Aud string `json:"aud"`
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
	AllowedIDPs            apijson.Field
	Aud                    apijson.Field
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
	ApplicationBrowserIsolationPermissionsApplicationTypeSaas        ApplicationBrowserIsolationPermissionsApplicationType = "saas"
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
	case ApplicationBrowserIsolationPermissionsApplicationTypeSelfHosted, ApplicationBrowserIsolationPermissionsApplicationTypeSaas, ApplicationBrowserIsolationPermissionsApplicationTypeSSH, ApplicationBrowserIsolationPermissionsApplicationTypeVnc, ApplicationBrowserIsolationPermissionsApplicationTypeAppLauncher, ApplicationBrowserIsolationPermissionsApplicationTypeWARP, ApplicationBrowserIsolationPermissionsApplicationTypeBiso, ApplicationBrowserIsolationPermissionsApplicationTypeBookmark, ApplicationBrowserIsolationPermissionsApplicationTypeDashSSO:
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
	Aud       string    `json:"aud"`
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
	Aud                apijson.Field
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

type CorsHeaders struct {
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
	AllowedHeaders []AllowedHeadersItem `json:"allowed_headers"`
	// Allowed HTTP request methods.
	AllowedMethods []AllowedMethodsItem `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []AllowedOriginsItem `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64         `json:"max_age"`
	JSON   corsHeadersJSON `json:"-"`
}

// corsHeadersJSON contains the JSON metadata for the struct [CorsHeaders]
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

func (r *CorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corsHeadersJSON) RawJSON() string {
	return r.raw
}

type CorsHeadersParam struct {
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
	AllowedHeaders param.Field[[]AllowedHeadersItemParam] `json:"allowed_headers"`
	// Allowed HTTP request methods.
	AllowedMethods param.Field[[]AllowedMethodsItem] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]AllowedOriginsItemParam] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r CorsHeadersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomPagesItem = string

type CustomPagesItemParam = string

type SamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType SamlSaasAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string                      `json:"consumer_service_url"`
	CreatedAt          time.Time                   `json:"created_at" format:"date-time"`
	CustomAttributes   SamlSaasAppCustomAttributes `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IDPEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat SamlSaasAppNameIDFormat `json:"name_id_format"`
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
	SamlAttributeTransformJsonata string `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SpEntityID string `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint string          `json:"sso_endpoint"`
	UpdatedAt   time.Time       `json:"updated_at" format:"date-time"`
	JSON        samlSaasAppJSON `json:"-"`
}

// samlSaasAppJSON contains the JSON metadata for the struct [SamlSaasApp]
type samlSaasAppJSON struct {
	AuthType                      apijson.Field
	ConsumerServiceURL            apijson.Field
	CreatedAt                     apijson.Field
	CustomAttributes              apijson.Field
	DefaultRelayState             apijson.Field
	IDPEntityID                   apijson.Field
	NameIDFormat                  apijson.Field
	NameIDTransformJsonata        apijson.Field
	PublicKey                     apijson.Field
	SamlAttributeTransformJsonata apijson.Field
	SpEntityID                    apijson.Field
	SSOEndpoint                   apijson.Field
	UpdatedAt                     apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *SamlSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r SamlSaasApp) implementsZeroTrustApplicationSaaSApplicationSaasApp() {}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type SamlSaasAppAuthType string

const (
	SamlSaasAppAuthTypeSaml SamlSaasAppAuthType = "saml"
	SamlSaasAppAuthTypeOidc SamlSaasAppAuthType = "oidc"
)

func (r SamlSaasAppAuthType) IsKnown() bool {
	switch r {
	case SamlSaasAppAuthTypeSaml, SamlSaasAppAuthTypeOidc:
		return true
	}
	return false
}

type SamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name string `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat SamlSaasAppCustomAttributesNameFormat `json:"name_format"`
	Source     SamlSaasAppCustomAttributesSource     `json:"source"`
	JSON       samlSaasAppCustomAttributesJSON       `json:"-"`
}

// samlSaasAppCustomAttributesJSON contains the JSON metadata for the struct
// [SamlSaasAppCustomAttributes]
type samlSaasAppCustomAttributesJSON struct {
	Name        apijson.Field
	NameFormat  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SamlSaasAppCustomAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlSaasAppCustomAttributesJSON) RawJSON() string {
	return r.raw
}

// A globally unique name for an identity or service provider.
type SamlSaasAppCustomAttributesNameFormat string

const (
	SamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified SamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	SamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       SamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	SamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         SamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

func (r SamlSaasAppCustomAttributesNameFormat) IsKnown() bool {
	switch r {
	case SamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified, SamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic, SamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI:
		return true
	}
	return false
}

type SamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name string                                `json:"name"`
	JSON samlSaasAppCustomAttributesSourceJSON `json:"-"`
}

// samlSaasAppCustomAttributesSourceJSON contains the JSON metadata for the struct
// [SamlSaasAppCustomAttributesSource]
type samlSaasAppCustomAttributesSourceJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SamlSaasAppCustomAttributesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlSaasAppCustomAttributesSourceJSON) RawJSON() string {
	return r.raw
}

// The format of the name identifier sent to the SaaS application.
type SamlSaasAppNameIDFormat string

const (
	SamlSaasAppNameIDFormatID    SamlSaasAppNameIDFormat = "id"
	SamlSaasAppNameIDFormatEmail SamlSaasAppNameIDFormat = "email"
)

func (r SamlSaasAppNameIDFormat) IsKnown() bool {
	switch r {
	case SamlSaasAppNameIDFormatID, SamlSaasAppNameIDFormatEmail:
		return true
	}
	return false
}

type SamlSaasAppParam struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[SamlSaasAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                           `json:"consumer_service_url"`
	CustomAttributes   param.Field[SamlSaasAppCustomAttributesParam] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IDPEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[SamlSaasAppNameIDFormat] `json:"name_id_format"`
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
	SamlAttributeTransformJsonata param.Field[string] `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SpEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string] `json:"sso_endpoint"`
}

func (r SamlSaasAppParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SamlSaasAppParam) implementsZeroTrustAccessApplicationNewParamsSaaSApplicationSaasAppUnion() {
}

func (r SamlSaasAppParam) implementsZeroTrustAccessApplicationUpdateParamsSaaSApplicationSaasAppUnion() {
}

type SamlSaasAppCustomAttributesParam struct {
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[SamlSaasAppCustomAttributesNameFormat]  `json:"name_format"`
	Source     param.Field[SamlSaasAppCustomAttributesSourceParam] `json:"source"`
}

func (r SamlSaasAppCustomAttributesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SamlSaasAppCustomAttributesSourceParam struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
}

func (r SamlSaasAppCustomAttributesSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SelfHostedDomainsItem = string

type SelfHostedDomainsItemParam = string

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

// This interface is a union satisfied by one of the following:
// [AccessApplicationNewParamsSelfHostedApplication],
// [AccessApplicationNewParamsSaaSApplication],
// [AccessApplicationNewParamsBrowserSSHApplication],
// [AccessApplicationNewParamsBrowserVncApplication],
// [AccessApplicationNewParamsAppLauncherApplication],
// [AccessApplicationNewParamsDeviceEnrollmentPermissionsApplication],
// [AccessApplicationNewParamsBrowserIsolationPermissionsApplication],
// [AccessApplicationNewParamsBookmarkApplication].
type AccessApplicationNewParams interface {
	ImplementsAccessApplicationNewParams()

	getAccountID() param.Field[string]

	getZoneID() param.Field[string]
}

type AccessApplicationNewParamsSelfHostedApplication struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]AllowedIDPsItemParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[CorsHeadersParam] `json:"cors_headers"`
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
	CustomPages param.Field[[]CustomPagesItemParam] `json:"custom_pages"`
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
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]SelfHostedDomainsItemParam] `json:"self_hosted_domains"`
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

func (r AccessApplicationNewParamsSelfHostedApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsSelfHostedApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationNewParamsSelfHostedApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationNewParamsSelfHostedApplication) ImplementsAccessApplicationNewParams() {

}

type AccessApplicationNewParamsSaaSApplication struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]AllowedIDPsItemParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages param.Field[[]CustomPagesItemParam] `json:"custom_pages"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name    param.Field[string]                                                `json:"name"`
	SaasApp param.Field[AccessApplicationNewParamsSaaSApplicationSaasAppUnion] `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessApplicationNewParamsSaaSApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsSaaSApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationNewParamsSaaSApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationNewParamsSaaSApplication) ImplementsAccessApplicationNewParams() {

}

type AccessApplicationNewParamsSaaSApplicationSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[AccessApplicationNewParamsSaaSApplicationSaasAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]      `json:"consumer_service_url"`
	CustomAttributes   param.Field[interface{}] `json:"custom_attributes,required"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IDPEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[AccessApplicationNewParamsSaaSApplicationSaasAppNameIDFormat] `json:"name_id_format"`
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
	SamlAttributeTransformJsonata param.Field[string] `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SpEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string] `json:"sso_endpoint"`
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string]      `json:"client_secret"`
	GrantTypes   param.Field[interface{}] `json:"grant_types,required"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex param.Field[string]      `json:"group_filter_regex"`
	RedirectURIs     param.Field[interface{}] `json:"redirect_uris,required"`
	Scopes           param.Field[interface{}] `json:"scopes,required"`
}

func (r AccessApplicationNewParamsSaaSApplicationSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsSaaSApplicationSaasApp) implementsZeroTrustAccessApplicationNewParamsSaaSApplicationSaasAppUnion() {
}

// Satisfied by [zero_trust.SamlSaasAppParam],
// [zero_trust.AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasApp],
// [AccessApplicationNewParamsSaaSApplicationSaasApp].
type AccessApplicationNewParamsSaaSApplicationSaasAppUnion interface {
	implementsZeroTrustAccessApplicationNewParamsSaaSApplicationSaasAppUnion()
}

type AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType param.Field[AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthType] `json:"auth_type"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string] `json:"client_secret"`
	// The OIDC flows supported by this application
	GrantTypes param.Field[[]AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantType] `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex param.Field[string] `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs param.Field[[]string] `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes param.Field[[]AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppScope] `json:"scopes"`
}

func (r AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationNewParamsSaaSApplicationSaasAppUnion() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthType string

const (
	AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "saml"
	AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "oidc"
)

func (r AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthType) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml, AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc:
		return true
	}
	return false
}

type AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantType string

const (
	AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

func (r AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantType) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode, AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce:
		return true
	}
	return false
}

type AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppScope string

const (
	AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid  AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppScope = "openid"
	AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups  AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppScope = "groups"
	AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail   AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppScope = "email"
	AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppScope = "profile"
)

func (r AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppScope) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid, AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups, AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail, AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile:
		return true
	}
	return false
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationNewParamsSaaSApplicationSaasAppAuthType string

const (
	AccessApplicationNewParamsSaaSApplicationSaasAppAuthTypeSaml AccessApplicationNewParamsSaaSApplicationSaasAppAuthType = "saml"
	AccessApplicationNewParamsSaaSApplicationSaasAppAuthTypeOidc AccessApplicationNewParamsSaaSApplicationSaasAppAuthType = "oidc"
)

func (r AccessApplicationNewParamsSaaSApplicationSaasAppAuthType) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsSaaSApplicationSaasAppAuthTypeSaml, AccessApplicationNewParamsSaaSApplicationSaasAppAuthTypeOidc:
		return true
	}
	return false
}

// The format of the name identifier sent to the SaaS application.
type AccessApplicationNewParamsSaaSApplicationSaasAppNameIDFormat string

const (
	AccessApplicationNewParamsSaaSApplicationSaasAppNameIDFormatID    AccessApplicationNewParamsSaaSApplicationSaasAppNameIDFormat = "id"
	AccessApplicationNewParamsSaaSApplicationSaasAppNameIDFormatEmail AccessApplicationNewParamsSaaSApplicationSaasAppNameIDFormat = "email"
)

func (r AccessApplicationNewParamsSaaSApplicationSaasAppNameIDFormat) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsSaaSApplicationSaasAppNameIDFormatID, AccessApplicationNewParamsSaaSApplicationSaasAppNameIDFormatEmail:
		return true
	}
	return false
}

type AccessApplicationNewParamsBrowserSSHApplication struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]AllowedIDPsItemParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[CorsHeadersParam] `json:"cors_headers"`
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
	CustomPages param.Field[[]CustomPagesItemParam] `json:"custom_pages"`
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
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]SelfHostedDomainsItemParam] `json:"self_hosted_domains"`
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

func (r AccessApplicationNewParamsBrowserSSHApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBrowserSSHApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationNewParamsBrowserSSHApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationNewParamsBrowserSSHApplication) ImplementsAccessApplicationNewParams() {

}

type AccessApplicationNewParamsBrowserVncApplication struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]AllowedIDPsItemParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[CorsHeadersParam] `json:"cors_headers"`
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
	CustomPages param.Field[[]CustomPagesItemParam] `json:"custom_pages"`
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
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]SelfHostedDomainsItemParam] `json:"self_hosted_domains"`
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

func (r AccessApplicationNewParamsBrowserVncApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBrowserVncApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationNewParamsBrowserVncApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationNewParamsBrowserVncApplication) ImplementsAccessApplicationNewParams() {

}

type AccessApplicationNewParamsAppLauncherApplication struct {
	// The application type.
	Type param.Field[AccessApplicationNewParamsAppLauncherApplicationType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]AllowedIDPsItemParam] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsAppLauncherApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsAppLauncherApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationNewParamsAppLauncherApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationNewParamsAppLauncherApplication) ImplementsAccessApplicationNewParams() {

}

// The application type.
type AccessApplicationNewParamsAppLauncherApplicationType string

const (
	AccessApplicationNewParamsAppLauncherApplicationTypeSelfHosted  AccessApplicationNewParamsAppLauncherApplicationType = "self_hosted"
	AccessApplicationNewParamsAppLauncherApplicationTypeSaas        AccessApplicationNewParamsAppLauncherApplicationType = "saas"
	AccessApplicationNewParamsAppLauncherApplicationTypeSSH         AccessApplicationNewParamsAppLauncherApplicationType = "ssh"
	AccessApplicationNewParamsAppLauncherApplicationTypeVnc         AccessApplicationNewParamsAppLauncherApplicationType = "vnc"
	AccessApplicationNewParamsAppLauncherApplicationTypeAppLauncher AccessApplicationNewParamsAppLauncherApplicationType = "app_launcher"
	AccessApplicationNewParamsAppLauncherApplicationTypeWARP        AccessApplicationNewParamsAppLauncherApplicationType = "warp"
	AccessApplicationNewParamsAppLauncherApplicationTypeBiso        AccessApplicationNewParamsAppLauncherApplicationType = "biso"
	AccessApplicationNewParamsAppLauncherApplicationTypeBookmark    AccessApplicationNewParamsAppLauncherApplicationType = "bookmark"
	AccessApplicationNewParamsAppLauncherApplicationTypeDashSSO     AccessApplicationNewParamsAppLauncherApplicationType = "dash_sso"
)

func (r AccessApplicationNewParamsAppLauncherApplicationType) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsAppLauncherApplicationTypeSelfHosted, AccessApplicationNewParamsAppLauncherApplicationTypeSaas, AccessApplicationNewParamsAppLauncherApplicationTypeSSH, AccessApplicationNewParamsAppLauncherApplicationTypeVnc, AccessApplicationNewParamsAppLauncherApplicationTypeAppLauncher, AccessApplicationNewParamsAppLauncherApplicationTypeWARP, AccessApplicationNewParamsAppLauncherApplicationTypeBiso, AccessApplicationNewParamsAppLauncherApplicationTypeBookmark, AccessApplicationNewParamsAppLauncherApplicationTypeDashSSO:
		return true
	}
	return false
}

type AccessApplicationNewParamsDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type param.Field[AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]AllowedIDPsItemParam] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsDeviceEnrollmentPermissionsApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsDeviceEnrollmentPermissionsApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationNewParamsDeviceEnrollmentPermissionsApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationNewParamsDeviceEnrollmentPermissionsApplication) ImplementsAccessApplicationNewParams() {

}

// The application type.
type AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationType string

const (
	AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeSelfHosted  AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationType = "self_hosted"
	AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeSaas        AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationType = "saas"
	AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeSSH         AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationType = "ssh"
	AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeVnc         AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationType = "vnc"
	AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeAppLauncher AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationType = "app_launcher"
	AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeWARP        AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationType = "warp"
	AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeBiso        AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationType = "biso"
	AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeBookmark    AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationType = "bookmark"
	AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeDashSSO     AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationType = "dash_sso"
)

func (r AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationType) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeSelfHosted, AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeSaas, AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeSSH, AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeVnc, AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeAppLauncher, AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeWARP, AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeBiso, AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeBookmark, AccessApplicationNewParamsDeviceEnrollmentPermissionsApplicationTypeDashSSO:
		return true
	}
	return false
}

type AccessApplicationNewParamsBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type param.Field[AccessApplicationNewParamsBrowserIsolationPermissionsApplicationType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]AllowedIDPsItemParam] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBrowserIsolationPermissionsApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBrowserIsolationPermissionsApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationNewParamsBrowserIsolationPermissionsApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationNewParamsBrowserIsolationPermissionsApplication) ImplementsAccessApplicationNewParams() {

}

// The application type.
type AccessApplicationNewParamsBrowserIsolationPermissionsApplicationType string

const (
	AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeSelfHosted  AccessApplicationNewParamsBrowserIsolationPermissionsApplicationType = "self_hosted"
	AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeSaas        AccessApplicationNewParamsBrowserIsolationPermissionsApplicationType = "saas"
	AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeSSH         AccessApplicationNewParamsBrowserIsolationPermissionsApplicationType = "ssh"
	AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeVnc         AccessApplicationNewParamsBrowserIsolationPermissionsApplicationType = "vnc"
	AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeAppLauncher AccessApplicationNewParamsBrowserIsolationPermissionsApplicationType = "app_launcher"
	AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeWARP        AccessApplicationNewParamsBrowserIsolationPermissionsApplicationType = "warp"
	AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeBiso        AccessApplicationNewParamsBrowserIsolationPermissionsApplicationType = "biso"
	AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeBookmark    AccessApplicationNewParamsBrowserIsolationPermissionsApplicationType = "bookmark"
	AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeDashSSO     AccessApplicationNewParamsBrowserIsolationPermissionsApplicationType = "dash_sso"
)

func (r AccessApplicationNewParamsBrowserIsolationPermissionsApplicationType) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeSelfHosted, AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeSaas, AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeSSH, AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeVnc, AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeAppLauncher, AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeWARP, AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeBiso, AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeBookmark, AccessApplicationNewParamsBrowserIsolationPermissionsApplicationTypeDashSSO:
		return true
	}
	return false
}

type AccessApplicationNewParamsBookmarkApplication struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
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

func (r AccessApplicationNewParamsBookmarkApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBookmarkApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationNewParamsBookmarkApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationNewParamsBookmarkApplication) ImplementsAccessApplicationNewParams() {

}

type AccessApplicationNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Application           `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationNewResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationNewResponseEnvelope]
type accessApplicationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// This interface is a union satisfied by one of the following:
// [AccessApplicationUpdateParamsSelfHostedApplication],
// [AccessApplicationUpdateParamsSaaSApplication],
// [AccessApplicationUpdateParamsBrowserSSHApplication],
// [AccessApplicationUpdateParamsBrowserVncApplication],
// [AccessApplicationUpdateParamsAppLauncherApplication],
// [AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplication],
// [AccessApplicationUpdateParamsBrowserIsolationPermissionsApplication],
// [AccessApplicationUpdateParamsBookmarkApplication].
type AccessApplicationUpdateParams interface {
	ImplementsAccessApplicationUpdateParams()

	getAccountID() param.Field[string]

	getZoneID() param.Field[string]
}

type AccessApplicationUpdateParamsSelfHostedApplication struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]AllowedIDPsItemParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[CorsHeadersParam] `json:"cors_headers"`
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
	CustomPages param.Field[[]CustomPagesItemParam] `json:"custom_pages"`
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
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]SelfHostedDomainsItemParam] `json:"self_hosted_domains"`
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

func (r AccessApplicationUpdateParamsSelfHostedApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsSelfHostedApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationUpdateParamsSelfHostedApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationUpdateParamsSelfHostedApplication) ImplementsAccessApplicationUpdateParams() {

}

type AccessApplicationUpdateParamsSaaSApplication struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]AllowedIDPsItemParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages param.Field[[]CustomPagesItemParam] `json:"custom_pages"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name    param.Field[string]                                                   `json:"name"`
	SaasApp param.Field[AccessApplicationUpdateParamsSaaSApplicationSaasAppUnion] `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessApplicationUpdateParamsSaaSApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsSaaSApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationUpdateParamsSaaSApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationUpdateParamsSaaSApplication) ImplementsAccessApplicationUpdateParams() {

}

type AccessApplicationUpdateParamsSaaSApplicationSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[AccessApplicationUpdateParamsSaaSApplicationSaasAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]      `json:"consumer_service_url"`
	CustomAttributes   param.Field[interface{}] `json:"custom_attributes,required"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IDPEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[AccessApplicationUpdateParamsSaaSApplicationSaasAppNameIDFormat] `json:"name_id_format"`
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
	SamlAttributeTransformJsonata param.Field[string] `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SpEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string] `json:"sso_endpoint"`
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string]      `json:"client_secret"`
	GrantTypes   param.Field[interface{}] `json:"grant_types,required"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex param.Field[string]      `json:"group_filter_regex"`
	RedirectURIs     param.Field[interface{}] `json:"redirect_uris,required"`
	Scopes           param.Field[interface{}] `json:"scopes,required"`
}

func (r AccessApplicationUpdateParamsSaaSApplicationSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsSaaSApplicationSaasApp) implementsZeroTrustAccessApplicationUpdateParamsSaaSApplicationSaasAppUnion() {
}

// Satisfied by [zero_trust.SamlSaasAppParam],
// [zero_trust.AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasApp],
// [AccessApplicationUpdateParamsSaaSApplicationSaasApp].
type AccessApplicationUpdateParamsSaaSApplicationSaasAppUnion interface {
	implementsZeroTrustAccessApplicationUpdateParamsSaaSApplicationSaasAppUnion()
}

type AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType param.Field[AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthType] `json:"auth_type"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string] `json:"client_secret"`
	// The OIDC flows supported by this application
	GrantTypes param.Field[[]AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantType] `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex param.Field[string] `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs param.Field[[]string] `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes param.Field[[]AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppScope] `json:"scopes"`
}

func (r AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationUpdateParamsSaaSApplicationSaasAppUnion() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthType string

const (
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "saml"
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "oidc"
)

func (r AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthType) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml, AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantType string

const (
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

func (r AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantType) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode, AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppScope string

const (
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid  AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppScope = "openid"
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups  AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppScope = "groups"
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail   AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppScope = "email"
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppScope = "profile"
)

func (r AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppScope) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid, AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups, AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail, AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile:
		return true
	}
	return false
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationUpdateParamsSaaSApplicationSaasAppAuthType string

const (
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAuthTypeSaml AccessApplicationUpdateParamsSaaSApplicationSaasAppAuthType = "saml"
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAuthTypeOidc AccessApplicationUpdateParamsSaaSApplicationSaasAppAuthType = "oidc"
)

func (r AccessApplicationUpdateParamsSaaSApplicationSaasAppAuthType) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsSaaSApplicationSaasAppAuthTypeSaml, AccessApplicationUpdateParamsSaaSApplicationSaasAppAuthTypeOidc:
		return true
	}
	return false
}

// The format of the name identifier sent to the SaaS application.
type AccessApplicationUpdateParamsSaaSApplicationSaasAppNameIDFormat string

const (
	AccessApplicationUpdateParamsSaaSApplicationSaasAppNameIDFormatID    AccessApplicationUpdateParamsSaaSApplicationSaasAppNameIDFormat = "id"
	AccessApplicationUpdateParamsSaaSApplicationSaasAppNameIDFormatEmail AccessApplicationUpdateParamsSaaSApplicationSaasAppNameIDFormat = "email"
)

func (r AccessApplicationUpdateParamsSaaSApplicationSaasAppNameIDFormat) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsSaaSApplicationSaasAppNameIDFormatID, AccessApplicationUpdateParamsSaaSApplicationSaasAppNameIDFormatEmail:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsBrowserSSHApplication struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]AllowedIDPsItemParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[CorsHeadersParam] `json:"cors_headers"`
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
	CustomPages param.Field[[]CustomPagesItemParam] `json:"custom_pages"`
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
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]SelfHostedDomainsItemParam] `json:"self_hosted_domains"`
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

func (r AccessApplicationUpdateParamsBrowserSSHApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBrowserSSHApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationUpdateParamsBrowserSSHApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationUpdateParamsBrowserSSHApplication) ImplementsAccessApplicationUpdateParams() {

}

type AccessApplicationUpdateParamsBrowserVncApplication struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]AllowedIDPsItemParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[CorsHeadersParam] `json:"cors_headers"`
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
	CustomPages param.Field[[]CustomPagesItemParam] `json:"custom_pages"`
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
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]SelfHostedDomainsItemParam] `json:"self_hosted_domains"`
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

func (r AccessApplicationUpdateParamsBrowserVncApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBrowserVncApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationUpdateParamsBrowserVncApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationUpdateParamsBrowserVncApplication) ImplementsAccessApplicationUpdateParams() {

}

type AccessApplicationUpdateParamsAppLauncherApplication struct {
	// The application type.
	Type param.Field[AccessApplicationUpdateParamsAppLauncherApplicationType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]AllowedIDPsItemParam] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsAppLauncherApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsAppLauncherApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationUpdateParamsAppLauncherApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationUpdateParamsAppLauncherApplication) ImplementsAccessApplicationUpdateParams() {

}

// The application type.
type AccessApplicationUpdateParamsAppLauncherApplicationType string

const (
	AccessApplicationUpdateParamsAppLauncherApplicationTypeSelfHosted  AccessApplicationUpdateParamsAppLauncherApplicationType = "self_hosted"
	AccessApplicationUpdateParamsAppLauncherApplicationTypeSaas        AccessApplicationUpdateParamsAppLauncherApplicationType = "saas"
	AccessApplicationUpdateParamsAppLauncherApplicationTypeSSH         AccessApplicationUpdateParamsAppLauncherApplicationType = "ssh"
	AccessApplicationUpdateParamsAppLauncherApplicationTypeVnc         AccessApplicationUpdateParamsAppLauncherApplicationType = "vnc"
	AccessApplicationUpdateParamsAppLauncherApplicationTypeAppLauncher AccessApplicationUpdateParamsAppLauncherApplicationType = "app_launcher"
	AccessApplicationUpdateParamsAppLauncherApplicationTypeWARP        AccessApplicationUpdateParamsAppLauncherApplicationType = "warp"
	AccessApplicationUpdateParamsAppLauncherApplicationTypeBiso        AccessApplicationUpdateParamsAppLauncherApplicationType = "biso"
	AccessApplicationUpdateParamsAppLauncherApplicationTypeBookmark    AccessApplicationUpdateParamsAppLauncherApplicationType = "bookmark"
	AccessApplicationUpdateParamsAppLauncherApplicationTypeDashSSO     AccessApplicationUpdateParamsAppLauncherApplicationType = "dash_sso"
)

func (r AccessApplicationUpdateParamsAppLauncherApplicationType) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsAppLauncherApplicationTypeSelfHosted, AccessApplicationUpdateParamsAppLauncherApplicationTypeSaas, AccessApplicationUpdateParamsAppLauncherApplicationTypeSSH, AccessApplicationUpdateParamsAppLauncherApplicationTypeVnc, AccessApplicationUpdateParamsAppLauncherApplicationTypeAppLauncher, AccessApplicationUpdateParamsAppLauncherApplicationTypeWARP, AccessApplicationUpdateParamsAppLauncherApplicationTypeBiso, AccessApplicationUpdateParamsAppLauncherApplicationTypeBookmark, AccessApplicationUpdateParamsAppLauncherApplicationTypeDashSSO:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type param.Field[AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]AllowedIDPsItemParam] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplication) ImplementsAccessApplicationUpdateParams() {

}

// The application type.
type AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationType string

const (
	AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeSelfHosted  AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationType = "self_hosted"
	AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeSaas        AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationType = "saas"
	AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeSSH         AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationType = "ssh"
	AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeVnc         AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationType = "vnc"
	AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeAppLauncher AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationType = "app_launcher"
	AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeWARP        AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationType = "warp"
	AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeBiso        AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationType = "biso"
	AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeBookmark    AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationType = "bookmark"
	AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeDashSSO     AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationType = "dash_sso"
)

func (r AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationType) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeSelfHosted, AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeSaas, AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeSSH, AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeVnc, AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeAppLauncher, AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeWARP, AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeBiso, AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeBookmark, AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationTypeDashSSO:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type param.Field[AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]AllowedIDPsItemParam] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBrowserIsolationPermissionsApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBrowserIsolationPermissionsApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationUpdateParamsBrowserIsolationPermissionsApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationUpdateParamsBrowserIsolationPermissionsApplication) ImplementsAccessApplicationUpdateParams() {

}

// The application type.
type AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationType string

const (
	AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeSelfHosted  AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationType = "self_hosted"
	AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeSaas        AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationType = "saas"
	AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeSSH         AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationType = "ssh"
	AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeVnc         AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationType = "vnc"
	AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeAppLauncher AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationType = "app_launcher"
	AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeWARP        AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationType = "warp"
	AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeBiso        AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationType = "biso"
	AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeBookmark    AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationType = "bookmark"
	AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeDashSSO     AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationType = "dash_sso"
)

func (r AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationType) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeSelfHosted, AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeSaas, AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeSSH, AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeVnc, AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeAppLauncher, AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeWARP, AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeBiso, AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeBookmark, AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationTypeDashSSO:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsBookmarkApplication struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
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

func (r AccessApplicationUpdateParamsBookmarkApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBookmarkApplication) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r AccessApplicationUpdateParamsBookmarkApplication) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (AccessApplicationUpdateParamsBookmarkApplication) ImplementsAccessApplicationUpdateParams() {

}

type AccessApplicationUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Application           `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationUpdateResponseEnvelope]
type accessApplicationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   AccessApplicationDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationDeleteResponseEnvelope]
type accessApplicationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
	Result   Application           `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationGetResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationGetResponseEnvelope]
type accessApplicationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

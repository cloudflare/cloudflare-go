// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustAccessApplicationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustAccessApplicationService] method instead.
type ZeroTrustAccessApplicationService struct {
	Options          []option.RequestOption
	CAs              *ZeroTrustAccessApplicationCAService
	UserPolicyChecks *ZeroTrustAccessApplicationUserPolicyCheckService
	Policies         *ZeroTrustAccessApplicationPolicyService
}

// NewZeroTrustAccessApplicationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustAccessApplicationService(opts ...option.RequestOption) (r *ZeroTrustAccessApplicationService) {
	r = &ZeroTrustAccessApplicationService{}
	r.Options = opts
	r.CAs = NewZeroTrustAccessApplicationCAService(opts...)
	r.UserPolicyChecks = NewZeroTrustAccessApplicationUserPolicyCheckService(opts...)
	r.Policies = NewZeroTrustAccessApplicationPolicyService(opts...)
	return
}

// Adds a new application to Access.
func (r *ZeroTrustAccessApplicationService) New(ctx context.Context, params ZeroTrustAccessApplicationNewParams, opts ...option.RequestOption) (res *AccessApps, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationNewResponseEnvelope
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
func (r *ZeroTrustAccessApplicationService) Update(ctx context.Context, appID ZeroTrustAccessApplicationUpdateParamsAppID, params ZeroTrustAccessApplicationUpdateParams, opts ...option.RequestOption) (res *AccessApps, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationUpdateResponseEnvelope
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
func (r *ZeroTrustAccessApplicationService) List(ctx context.Context, query ZeroTrustAccessApplicationListParams, opts ...option.RequestOption) (res *[]AccessApps, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationListResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an application from Access.
func (r *ZeroTrustAccessApplicationService) Delete(ctx context.Context, appID ZeroTrustAccessApplicationDeleteParamsAppID, body ZeroTrustAccessApplicationDeleteParams, opts ...option.RequestOption) (res *ZeroTrustAccessApplicationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationDeleteResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches information about an Access application.
func (r *ZeroTrustAccessApplicationService) Get(ctx context.Context, appID ZeroTrustAccessApplicationGetParamsAppID, query ZeroTrustAccessApplicationGetParams, opts ...option.RequestOption) (res *AccessApps, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationGetResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Revokes all tokens issued for an application.
func (r *ZeroTrustAccessApplicationService) RevokeTokens(ctx context.Context, appID ZeroTrustAccessApplicationRevokeTokensParamsAppID, body ZeroTrustAccessApplicationRevokeTokensParams, opts ...option.RequestOption) (res *ZeroTrustAccessApplicationRevokeTokensResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationRevokeTokensResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AccessAppsSelfHostedApplication],
// [AccessAppsSaaSApplication], [AccessAppsBrowserSSHApplication],
// [AccessAppsBrowserVncApplication], [AccessAppsAppLauncherApplication],
// [AccessAppsDeviceEnrollmentPermissionsApplication],
// [AccessAppsBrowserIsolationPermissionsApplication] or
// [AccessAppsBookmarkApplication].
type AccessApps interface {
	implementsAccessApps()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApps)(nil)).Elem(), "")
}

type AccessAppsSelfHostedApplication struct {
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
	AllowedIdps []string `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	Aud string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool                                       `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessAppsSelfHostedApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                  `json:"created_at" format:"date-time"`
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
	CustomPages []string `json:"custom_pages"`
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
	SelfHostedDomains []string `json:"self_hosted_domains"`
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
	Tags      []string                            `json:"tags"`
	UpdatedAt time.Time                           `json:"updated_at" format:"date-time"`
	JSON      accessAppsSelfHostedApplicationJSON `json:"-"`
}

// accessAppsSelfHostedApplicationJSON contains the JSON metadata for the struct
// [AccessAppsSelfHostedApplication]
type accessAppsSelfHostedApplicationJSON struct {
	Domain                   apijson.Field
	Type                     apijson.Field
	ID                       apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIdps              apijson.Field
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

func (r *AccessAppsSelfHostedApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppsSelfHostedApplication) implementsAccessApps() {}

type AccessAppsSelfHostedApplicationCorsHeaders struct {
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
	AllowedHeaders []interface{} `json:"allowed_headers"`
	// Allowed HTTP request methods.
	AllowedMethods []AccessAppsSelfHostedApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                        `json:"max_age"`
	JSON   accessAppsSelfHostedApplicationCorsHeadersJSON `json:"-"`
}

// accessAppsSelfHostedApplicationCorsHeadersJSON contains the JSON metadata for
// the struct [AccessAppsSelfHostedApplicationCorsHeaders]
type accessAppsSelfHostedApplicationCorsHeadersJSON struct {
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

func (r *AccessAppsSelfHostedApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppsSelfHostedApplicationCorsHeadersAllowedMethod string

const (
	AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodGet     AccessAppsSelfHostedApplicationCorsHeadersAllowedMethod = "GET"
	AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodPost    AccessAppsSelfHostedApplicationCorsHeadersAllowedMethod = "POST"
	AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodHead    AccessAppsSelfHostedApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodPut     AccessAppsSelfHostedApplicationCorsHeadersAllowedMethod = "PUT"
	AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodDelete  AccessAppsSelfHostedApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodConnect AccessAppsSelfHostedApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodOptions AccessAppsSelfHostedApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodTrace   AccessAppsSelfHostedApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodPatch   AccessAppsSelfHostedApplicationCorsHeadersAllowedMethod = "PATCH"
)

type AccessAppsSaaSApplication struct {
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps []string `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	Aud string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages []string `json:"custom_pages"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name    string                           `json:"name"`
	SaasApp AccessAppsSaaSApplicationSaasApp `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags []string `json:"tags"`
	// The application type.
	Type      string                        `json:"type"`
	UpdatedAt time.Time                     `json:"updated_at" format:"date-time"`
	JSON      accessAppsSaaSApplicationJSON `json:"-"`
}

// accessAppsSaaSApplicationJSON contains the JSON metadata for the struct
// [AccessAppsSaaSApplication]
type accessAppsSaaSApplicationJSON struct {
	ID                     apijson.Field
	AllowedIdps            apijson.Field
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

func (r *AccessAppsSaaSApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppsSaaSApplication) implementsAccessApps() {}

// Union satisfied by [AccessAppsSaaSApplicationSaasAppAccessSamlSaasApp] or
// [AccessAppsSaaSApplicationSaasAppAccessOidcSaasApp].
type AccessAppsSaaSApplicationSaasApp interface {
	implementsAccessAppsSaaSApplicationSaasApp()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppsSaaSApplicationSaasApp)(nil)).Elem(), "")
}

type AccessAppsSaaSApplicationSaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string                                                            `json:"consumer_service_url"`
	CreatedAt          time.Time                                                         `json:"created_at" format:"date-time"`
	CustomAttributes   AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat `json:"name_id_format"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's
	// user identities into a NameID value for its SAML assertion. This expression
	// should evaluate to a singular string. The output of this expression can override
	// the `name_id_format` setting.
	NameIDTransformJsonata string `json:"name_id_transform_jsonata"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// A globally unique name for an identity or service provider.
	SpEntityID string `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint string                                                `json:"sso_endpoint"`
	UpdatedAt   time.Time                                             `json:"updated_at" format:"date-time"`
	JSON        accessAppsSaaSApplicationSaasAppAccessSamlSaasAppJSON `json:"-"`
}

// accessAppsSaaSApplicationSaasAppAccessSamlSaasAppJSON contains the JSON metadata
// for the struct [AccessAppsSaaSApplicationSaasAppAccessSamlSaasApp]
type accessAppsSaaSApplicationSaasAppAccessSamlSaasAppJSON struct {
	AuthType               apijson.Field
	ConsumerServiceURL     apijson.Field
	CreatedAt              apijson.Field
	CustomAttributes       apijson.Field
	DefaultRelayState      apijson.Field
	IdpEntityID            apijson.Field
	NameIDFormat           apijson.Field
	NameIDTransformJsonata apijson.Field
	PublicKey              apijson.Field
	SpEntityID             apijson.Field
	SSOEndpoint            apijson.Field
	UpdatedAt              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessAppsSaaSApplicationSaasAppAccessSamlSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppsSaaSApplicationSaasAppAccessSamlSaasApp) implementsAccessAppsSaaSApplicationSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppAuthType string

const (
	AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeSaml AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "saml"
	AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeOidc AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "oidc"
)

type AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name string `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat `json:"name_format"`
	Source     AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource     `json:"source"`
	JSON       accessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON       `json:"-"`
}

// accessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON contains
// the JSON metadata for the struct
// [AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes]
type accessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON struct {
	Name        apijson.Field
	NameFormat  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A globally unique name for an identity or service provider.
type AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUri         AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name string                                                                      `json:"name"`
	JSON accessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON `json:"-"`
}

// accessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON
// contains the JSON metadata for the struct
// [AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource]
type accessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The format of the name identifier sent to the SaaS application.
type AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat string

const (
	AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatID    AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "id"
	AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatEmail AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type AccessAppsSaaSApplicationSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The OIDC flows supported by this application
	GrantTypes []AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppGrantType `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectUris []string `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes    []AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScope `json:"scopes"`
	UpdatedAt time.Time                                                `json:"updated_at" format:"date-time"`
	JSON      accessAppsSaaSApplicationSaasAppAccessOidcSaasAppJSON    `json:"-"`
}

// accessAppsSaaSApplicationSaasAppAccessOidcSaasAppJSON contains the JSON metadata
// for the struct [AccessAppsSaaSApplicationSaasAppAccessOidcSaasApp]
type accessAppsSaaSApplicationSaasAppAccessOidcSaasAppJSON struct {
	AppLauncherURL   apijson.Field
	AuthType         apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	CreatedAt        apijson.Field
	GrantTypes       apijson.Field
	GroupFilterRegex apijson.Field
	PublicKey        apijson.Field
	RedirectUris     apijson.Field
	Scopes           apijson.Field
	UpdatedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessAppsSaaSApplicationSaasAppAccessOidcSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppsSaaSApplicationSaasAppAccessOidcSaasApp) implementsAccessAppsSaaSApplicationSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppAuthType string

const (
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "saml"
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "oidc"
)

type AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppGrantType string

const (
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScope string

const (
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid  AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScope = "openid"
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups  AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScope = "groups"
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail   AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScope = "email"
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScope = "profile"
)

type AccessAppsBrowserSSHApplication struct {
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
	AllowedIdps []string `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	Aud string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool                                       `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessAppsBrowserSSHApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                  `json:"created_at" format:"date-time"`
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
	CustomPages []string `json:"custom_pages"`
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
	SelfHostedDomains []string `json:"self_hosted_domains"`
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
	Tags      []string                            `json:"tags"`
	UpdatedAt time.Time                           `json:"updated_at" format:"date-time"`
	JSON      accessAppsBrowserSSHApplicationJSON `json:"-"`
}

// accessAppsBrowserSSHApplicationJSON contains the JSON metadata for the struct
// [AccessAppsBrowserSSHApplication]
type accessAppsBrowserSSHApplicationJSON struct {
	Domain                   apijson.Field
	Type                     apijson.Field
	ID                       apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIdps              apijson.Field
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

func (r *AccessAppsBrowserSSHApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppsBrowserSSHApplication) implementsAccessApps() {}

type AccessAppsBrowserSSHApplicationCorsHeaders struct {
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
	AllowedHeaders []interface{} `json:"allowed_headers"`
	// Allowed HTTP request methods.
	AllowedMethods []AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                        `json:"max_age"`
	JSON   accessAppsBrowserSSHApplicationCorsHeadersJSON `json:"-"`
}

// accessAppsBrowserSSHApplicationCorsHeadersJSON contains the JSON metadata for
// the struct [AccessAppsBrowserSSHApplicationCorsHeaders]
type accessAppsBrowserSSHApplicationCorsHeadersJSON struct {
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

func (r *AccessAppsBrowserSSHApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethod string

const (
	AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodGet     AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethod = "GET"
	AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodPost    AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethod = "POST"
	AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodHead    AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodPut     AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethod = "PUT"
	AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodDelete  AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodConnect AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodOptions AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodTrace   AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodPatch   AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethod = "PATCH"
)

type AccessAppsBrowserVncApplication struct {
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
	AllowedIdps []string `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	Aud string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool                                       `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessAppsBrowserVncApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                  `json:"created_at" format:"date-time"`
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
	CustomPages []string `json:"custom_pages"`
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
	SelfHostedDomains []string `json:"self_hosted_domains"`
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
	Tags      []string                            `json:"tags"`
	UpdatedAt time.Time                           `json:"updated_at" format:"date-time"`
	JSON      accessAppsBrowserVncApplicationJSON `json:"-"`
}

// accessAppsBrowserVncApplicationJSON contains the JSON metadata for the struct
// [AccessAppsBrowserVncApplication]
type accessAppsBrowserVncApplicationJSON struct {
	Domain                   apijson.Field
	Type                     apijson.Field
	ID                       apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIdps              apijson.Field
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

func (r *AccessAppsBrowserVncApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppsBrowserVncApplication) implementsAccessApps() {}

type AccessAppsBrowserVncApplicationCorsHeaders struct {
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
	AllowedHeaders []interface{} `json:"allowed_headers"`
	// Allowed HTTP request methods.
	AllowedMethods []AccessAppsBrowserVncApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                        `json:"max_age"`
	JSON   accessAppsBrowserVncApplicationCorsHeadersJSON `json:"-"`
}

// accessAppsBrowserVncApplicationCorsHeadersJSON contains the JSON metadata for
// the struct [AccessAppsBrowserVncApplicationCorsHeaders]
type accessAppsBrowserVncApplicationCorsHeadersJSON struct {
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

func (r *AccessAppsBrowserVncApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppsBrowserVncApplicationCorsHeadersAllowedMethod string

const (
	AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodGet     AccessAppsBrowserVncApplicationCorsHeadersAllowedMethod = "GET"
	AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodPost    AccessAppsBrowserVncApplicationCorsHeadersAllowedMethod = "POST"
	AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodHead    AccessAppsBrowserVncApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodPut     AccessAppsBrowserVncApplicationCorsHeadersAllowedMethod = "PUT"
	AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodDelete  AccessAppsBrowserVncApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodConnect AccessAppsBrowserVncApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodOptions AccessAppsBrowserVncApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodTrace   AccessAppsBrowserVncApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodPatch   AccessAppsBrowserVncApplicationCorsHeadersAllowedMethod = "PATCH"
)

type AccessAppsAppLauncherApplication struct {
	// The application type.
	Type AccessAppsAppLauncherApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps []string `json:"allowed_idps"`
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
	SessionDuration string                               `json:"session_duration"`
	UpdatedAt       time.Time                            `json:"updated_at" format:"date-time"`
	JSON            accessAppsAppLauncherApplicationJSON `json:"-"`
}

// accessAppsAppLauncherApplicationJSON contains the JSON metadata for the struct
// [AccessAppsAppLauncherApplication]
type accessAppsAppLauncherApplicationJSON struct {
	Type                   apijson.Field
	ID                     apijson.Field
	AllowedIdps            apijson.Field
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

func (r *AccessAppsAppLauncherApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppsAppLauncherApplication) implementsAccessApps() {}

// The application type.
type AccessAppsAppLauncherApplicationType string

const (
	AccessAppsAppLauncherApplicationTypeSelfHosted  AccessAppsAppLauncherApplicationType = "self_hosted"
	AccessAppsAppLauncherApplicationTypeSaas        AccessAppsAppLauncherApplicationType = "saas"
	AccessAppsAppLauncherApplicationTypeSSH         AccessAppsAppLauncherApplicationType = "ssh"
	AccessAppsAppLauncherApplicationTypeVnc         AccessAppsAppLauncherApplicationType = "vnc"
	AccessAppsAppLauncherApplicationTypeAppLauncher AccessAppsAppLauncherApplicationType = "app_launcher"
	AccessAppsAppLauncherApplicationTypeWARP        AccessAppsAppLauncherApplicationType = "warp"
	AccessAppsAppLauncherApplicationTypeBiso        AccessAppsAppLauncherApplicationType = "biso"
	AccessAppsAppLauncherApplicationTypeBookmark    AccessAppsAppLauncherApplicationType = "bookmark"
	AccessAppsAppLauncherApplicationTypeDashSSO     AccessAppsAppLauncherApplicationType = "dash_sso"
)

type AccessAppsDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type AccessAppsDeviceEnrollmentPermissionsApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps []string `json:"allowed_idps"`
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
	SessionDuration string                                               `json:"session_duration"`
	UpdatedAt       time.Time                                            `json:"updated_at" format:"date-time"`
	JSON            accessAppsDeviceEnrollmentPermissionsApplicationJSON `json:"-"`
}

// accessAppsDeviceEnrollmentPermissionsApplicationJSON contains the JSON metadata
// for the struct [AccessAppsDeviceEnrollmentPermissionsApplication]
type accessAppsDeviceEnrollmentPermissionsApplicationJSON struct {
	Type                   apijson.Field
	ID                     apijson.Field
	AllowedIdps            apijson.Field
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

func (r *AccessAppsDeviceEnrollmentPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppsDeviceEnrollmentPermissionsApplication) implementsAccessApps() {}

// The application type.
type AccessAppsDeviceEnrollmentPermissionsApplicationType string

const (
	AccessAppsDeviceEnrollmentPermissionsApplicationTypeSelfHosted  AccessAppsDeviceEnrollmentPermissionsApplicationType = "self_hosted"
	AccessAppsDeviceEnrollmentPermissionsApplicationTypeSaas        AccessAppsDeviceEnrollmentPermissionsApplicationType = "saas"
	AccessAppsDeviceEnrollmentPermissionsApplicationTypeSSH         AccessAppsDeviceEnrollmentPermissionsApplicationType = "ssh"
	AccessAppsDeviceEnrollmentPermissionsApplicationTypeVnc         AccessAppsDeviceEnrollmentPermissionsApplicationType = "vnc"
	AccessAppsDeviceEnrollmentPermissionsApplicationTypeAppLauncher AccessAppsDeviceEnrollmentPermissionsApplicationType = "app_launcher"
	AccessAppsDeviceEnrollmentPermissionsApplicationTypeWARP        AccessAppsDeviceEnrollmentPermissionsApplicationType = "warp"
	AccessAppsDeviceEnrollmentPermissionsApplicationTypeBiso        AccessAppsDeviceEnrollmentPermissionsApplicationType = "biso"
	AccessAppsDeviceEnrollmentPermissionsApplicationTypeBookmark    AccessAppsDeviceEnrollmentPermissionsApplicationType = "bookmark"
	AccessAppsDeviceEnrollmentPermissionsApplicationTypeDashSSO     AccessAppsDeviceEnrollmentPermissionsApplicationType = "dash_sso"
)

type AccessAppsBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type AccessAppsBrowserIsolationPermissionsApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps []string `json:"allowed_idps"`
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
	SessionDuration string                                               `json:"session_duration"`
	UpdatedAt       time.Time                                            `json:"updated_at" format:"date-time"`
	JSON            accessAppsBrowserIsolationPermissionsApplicationJSON `json:"-"`
}

// accessAppsBrowserIsolationPermissionsApplicationJSON contains the JSON metadata
// for the struct [AccessAppsBrowserIsolationPermissionsApplication]
type accessAppsBrowserIsolationPermissionsApplicationJSON struct {
	Type                   apijson.Field
	ID                     apijson.Field
	AllowedIdps            apijson.Field
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

func (r *AccessAppsBrowserIsolationPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppsBrowserIsolationPermissionsApplication) implementsAccessApps() {}

// The application type.
type AccessAppsBrowserIsolationPermissionsApplicationType string

const (
	AccessAppsBrowserIsolationPermissionsApplicationTypeSelfHosted  AccessAppsBrowserIsolationPermissionsApplicationType = "self_hosted"
	AccessAppsBrowserIsolationPermissionsApplicationTypeSaas        AccessAppsBrowserIsolationPermissionsApplicationType = "saas"
	AccessAppsBrowserIsolationPermissionsApplicationTypeSSH         AccessAppsBrowserIsolationPermissionsApplicationType = "ssh"
	AccessAppsBrowserIsolationPermissionsApplicationTypeVnc         AccessAppsBrowserIsolationPermissionsApplicationType = "vnc"
	AccessAppsBrowserIsolationPermissionsApplicationTypeAppLauncher AccessAppsBrowserIsolationPermissionsApplicationType = "app_launcher"
	AccessAppsBrowserIsolationPermissionsApplicationTypeWARP        AccessAppsBrowserIsolationPermissionsApplicationType = "warp"
	AccessAppsBrowserIsolationPermissionsApplicationTypeBiso        AccessAppsBrowserIsolationPermissionsApplicationType = "biso"
	AccessAppsBrowserIsolationPermissionsApplicationTypeBookmark    AccessAppsBrowserIsolationPermissionsApplicationType = "bookmark"
	AccessAppsBrowserIsolationPermissionsApplicationTypeDashSSO     AccessAppsBrowserIsolationPermissionsApplicationType = "dash_sso"
)

type AccessAppsBookmarkApplication struct {
	// UUID
	ID                 string      `json:"id"`
	AppLauncherVisible interface{} `json:"app_launcher_visible"`
	// Audience tag.
	Aud       string    `json:"aud"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The URL or domain of the bookmark.
	Domain interface{} `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name string `json:"name"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags []string `json:"tags"`
	// The application type.
	Type      string                            `json:"type"`
	UpdatedAt time.Time                         `json:"updated_at" format:"date-time"`
	JSON      accessAppsBookmarkApplicationJSON `json:"-"`
}

// accessAppsBookmarkApplicationJSON contains the JSON metadata for the struct
// [AccessAppsBookmarkApplication]
type accessAppsBookmarkApplicationJSON struct {
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

func (r *AccessAppsBookmarkApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppsBookmarkApplication) implementsAccessApps() {}

type ZeroTrustAccessApplicationDeleteResponse struct {
	// UUID
	ID   string                                       `json:"id"`
	JSON zeroTrustAccessApplicationDeleteResponseJSON `json:"-"`
}

// zeroTrustAccessApplicationDeleteResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessApplicationDeleteResponse]
type zeroTrustAccessApplicationDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationRevokeTokensResponse = interface{}

type ZeroTrustAccessApplicationNewParams struct {
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
	AllowedIdps        param.Field[[]string]    `json:"allowed_idps"`
	AppLauncherVisible param.Field[interface{}] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                           `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[ZeroTrustAccessApplicationNewParamsCorsHeaders] `json:"cors_headers"`
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
	CustomPages param.Field[[]string] `json:"custom_pages"`
	// The URL or domain of the bookmark.
	Domain param.Field[interface{}] `json:"domain"`
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
	PathCookieAttribute param.Field[bool]                                       `json:"path_cookie_attribute"`
	SaasApp             param.Field[ZeroTrustAccessApplicationNewParamsSaasApp] `json:"saas_app"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]string] `json:"self_hosted_domains"`
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
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r ZeroTrustAccessApplicationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessApplicationNewParamsCorsHeaders struct {
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
	AllowedHeaders param.Field[[]interface{}] `json:"allowed_headers"`
	// Allowed HTTP request methods.
	AllowedMethods param.Field[[]ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r ZeroTrustAccessApplicationNewParamsCorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethod string

const (
	ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethodGet     ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethod = "GET"
	ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethodPost    ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethod = "POST"
	ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethodHead    ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethod = "HEAD"
	ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethodPut     ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethod = "PUT"
	ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethodDelete  ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethod = "DELETE"
	ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethodConnect ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethod = "CONNECT"
	ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethodOptions ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethod = "OPTIONS"
	ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethodTrace   ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethod = "TRACE"
	ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethodPatch   ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethod = "PATCH"
)

// Satisfied by [ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasApp],
// [ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasApp].
type ZeroTrustAccessApplicationNewParamsSaasApp interface {
	implementsZeroTrustAccessApplicationNewParamsSaasApp()
}

type ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                                                                      `json:"consumer_service_url"`
	CustomAttributes   param.Field[ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributes] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppNameIDFormat] `json:"name_id_format"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's
	// user identities into a NameID value for its SAML assertion. This expression
	// should evaluate to a singular string. The output of this expression can override
	// the `name_id_format` setting.
	NameIDTransformJsonata param.Field[string] `json:"name_id_transform_jsonata"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// A globally unique name for an identity or service provider.
	SpEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string] `json:"sso_endpoint"`
}

func (r ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasApp) implementsZeroTrustAccessApplicationNewParamsSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppAuthType string

const (
	ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppAuthTypeSaml ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppAuthType = "saml"
	ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppAuthTypeOidc ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppAuthType = "oidc"
)

type ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat] `json:"name_format"`
	Source     param.Field[ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesSource]     `json:"source"`
}

func (r ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A globally unique name for an identity or service provider.
type ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUri         ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
}

func (r ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the name identifier sent to the SaaS application.
type ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppNameIDFormat string

const (
	ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppNameIDFormatID    ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppNameIDFormat = "id"
	ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppNameIDFormatEmail ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType param.Field[ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppAuthType] `json:"auth_type"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string] `json:"client_secret"`
	// The OIDC flows supported by this application
	GrantTypes param.Field[[]ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppGrantType] `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex param.Field[string] `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectUris param.Field[[]string] `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes param.Field[[]ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppScope] `json:"scopes"`
}

func (r ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationNewParamsSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppAuthType string

const (
	ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppAuthTypeSaml ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppAuthType = "saml"
	ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppAuthTypeOidc ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppAuthType = "oidc"
)

type ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppGrantType string

const (
	ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppScope string

const (
	ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppScopeOpenid  ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppScope = "openid"
	ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppScopeGroups  ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppScope = "groups"
	ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppScopeEmail   ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppScope = "email"
	ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppScopeProfile ZeroTrustAccessApplicationNewParamsSaasAppAccessOidcSaasAppScope = "profile"
)

type ZeroTrustAccessApplicationNewResponseEnvelope struct {
	Errors   []ZeroTrustAccessApplicationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessApplicationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApps                                              `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessApplicationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessApplicationNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessApplicationNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustAccessApplicationNewResponseEnvelope]
type zeroTrustAccessApplicationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationNewResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustAccessApplicationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessApplicationNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationNewResponseEnvelopeErrors]
type zeroTrustAccessApplicationNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationNewResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustAccessApplicationNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessApplicationNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationNewResponseEnvelopeMessages]
type zeroTrustAccessApplicationNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessApplicationNewResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationNewResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationNewResponseEnvelopeSuccess = true
)

type ZeroTrustAccessApplicationUpdateParams struct {
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
	AllowedIdps        param.Field[[]string]    `json:"allowed_idps"`
	AppLauncherVisible param.Field[interface{}] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                              `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[ZeroTrustAccessApplicationUpdateParamsCorsHeaders] `json:"cors_headers"`
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
	CustomPages param.Field[[]string] `json:"custom_pages"`
	// The URL or domain of the bookmark.
	Domain param.Field[interface{}] `json:"domain"`
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
	PathCookieAttribute param.Field[bool]                                          `json:"path_cookie_attribute"`
	SaasApp             param.Field[ZeroTrustAccessApplicationUpdateParamsSaasApp] `json:"saas_app"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]string] `json:"self_hosted_domains"`
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
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r ZeroTrustAccessApplicationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type ZeroTrustAccessApplicationUpdateParamsAppID interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsAppID()
}

type ZeroTrustAccessApplicationUpdateParamsCorsHeaders struct {
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
	AllowedHeaders param.Field[[]interface{}] `json:"allowed_headers"`
	// Allowed HTTP request methods.
	AllowedMethods param.Field[[]ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r ZeroTrustAccessApplicationUpdateParamsCorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethod string

const (
	ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethodGet     ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethod = "GET"
	ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethodPost    ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethod = "POST"
	ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethodHead    ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethod = "HEAD"
	ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethodPut     ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethod = "PUT"
	ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethodDelete  ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethod = "DELETE"
	ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethodConnect ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethod = "CONNECT"
	ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethodOptions ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethod = "OPTIONS"
	ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethodTrace   ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethod = "TRACE"
	ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethodPatch   ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethod = "PATCH"
)

// Satisfied by [ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasApp],
// [ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasApp].
type ZeroTrustAccessApplicationUpdateParamsSaasApp interface {
	implementsZeroTrustAccessApplicationUpdateParamsSaasApp()
}

type ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                                                                         `json:"consumer_service_url"`
	CustomAttributes   param.Field[ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributes] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppNameIDFormat] `json:"name_id_format"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's
	// user identities into a NameID value for its SAML assertion. This expression
	// should evaluate to a singular string. The output of this expression can override
	// the `name_id_format` setting.
	NameIDTransformJsonata param.Field[string] `json:"name_id_transform_jsonata"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// A globally unique name for an identity or service provider.
	SpEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string] `json:"sso_endpoint"`
}

func (r ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasApp) implementsZeroTrustAccessApplicationUpdateParamsSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppAuthType string

const (
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppAuthTypeSaml ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppAuthType = "saml"
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppAuthTypeOidc ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppAuthType = "oidc"
)

type ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat] `json:"name_format"`
	Source     param.Field[ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesSource]     `json:"source"`
}

func (r ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A globally unique name for an identity or service provider.
type ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUri         ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
}

func (r ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the name identifier sent to the SaaS application.
type ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppNameIDFormat string

const (
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppNameIDFormatID    ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppNameIDFormat = "id"
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppNameIDFormatEmail ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType param.Field[ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppAuthType] `json:"auth_type"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string] `json:"client_secret"`
	// The OIDC flows supported by this application
	GrantTypes param.Field[[]ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppGrantType] `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex param.Field[string] `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectUris param.Field[[]string] `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes param.Field[[]ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScope] `json:"scopes"`
}

func (r ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationUpdateParamsSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppAuthType string

const (
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppAuthTypeSaml ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppAuthType = "saml"
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppAuthTypeOidc ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppAuthType = "oidc"
)

type ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppGrantType string

const (
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScope string

const (
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScopeOpenid  ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScope = "openid"
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScopeGroups  ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScope = "groups"
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScopeEmail   ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScope = "email"
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScopeProfile ZeroTrustAccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScope = "profile"
)

type ZeroTrustAccessApplicationUpdateResponseEnvelope struct {
	Errors   []ZeroTrustAccessApplicationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessApplicationUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApps                                                 `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessApplicationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessApplicationUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessApplicationUpdateResponseEnvelope]
type zeroTrustAccessApplicationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationUpdateResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustAccessApplicationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationUpdateResponseEnvelopeErrors]
type zeroTrustAccessApplicationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationUpdateResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustAccessApplicationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseEnvelopeMessages]
type zeroTrustAccessApplicationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessApplicationUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationUpdateResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustAccessApplicationListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessApplicationListResponseEnvelope struct {
	Errors   []ZeroTrustAccessApplicationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessApplicationListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessApps                                             `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessApplicationListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessApplicationListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessApplicationListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessApplicationListResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessApplicationListResponseEnvelope]
type zeroTrustAccessApplicationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationListResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustAccessApplicationListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationListResponseEnvelopeErrors]
type zeroTrustAccessApplicationListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationListResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustAccessApplicationListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationListResponseEnvelopeMessages]
type zeroTrustAccessApplicationListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessApplicationListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationListResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessApplicationListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                      `json:"total_count"`
	JSON       zeroTrustAccessApplicationListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationListResponseEnvelopeResultInfo]
type zeroTrustAccessApplicationListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type ZeroTrustAccessApplicationDeleteParamsAppID interface {
	ImplementsZeroTrustAccessApplicationDeleteParamsAppID()
}

type ZeroTrustAccessApplicationDeleteResponseEnvelope struct {
	Errors   []ZeroTrustAccessApplicationDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessApplicationDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessApplicationDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessApplicationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessApplicationDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessApplicationDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessApplicationDeleteResponseEnvelope]
type zeroTrustAccessApplicationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationDeleteResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustAccessApplicationDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessApplicationDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationDeleteResponseEnvelopeErrors]
type zeroTrustAccessApplicationDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationDeleteResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustAccessApplicationDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessApplicationDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationDeleteResponseEnvelopeMessages]
type zeroTrustAccessApplicationDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessApplicationDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationDeleteResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustAccessApplicationGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type ZeroTrustAccessApplicationGetParamsAppID interface {
	ImplementsZeroTrustAccessApplicationGetParamsAppID()
}

type ZeroTrustAccessApplicationGetResponseEnvelope struct {
	Errors   []ZeroTrustAccessApplicationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessApplicationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApps                                              `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessApplicationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessApplicationGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessApplicationGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustAccessApplicationGetResponseEnvelope]
type zeroTrustAccessApplicationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustAccessApplicationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessApplicationGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationGetResponseEnvelopeErrors]
type zeroTrustAccessApplicationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustAccessApplicationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessApplicationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationGetResponseEnvelopeMessages]
type zeroTrustAccessApplicationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessApplicationGetResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationGetResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationGetResponseEnvelopeSuccess = true
)

type ZeroTrustAccessApplicationRevokeTokensParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type ZeroTrustAccessApplicationRevokeTokensParamsAppID interface {
	ImplementsZeroTrustAccessApplicationRevokeTokensParamsAppID()
}

type ZeroTrustAccessApplicationRevokeTokensResponseEnvelope struct {
	Result  ZeroTrustAccessApplicationRevokeTokensResponse                `json:"result,nullable"`
	Success ZeroTrustAccessApplicationRevokeTokensResponseEnvelopeSuccess `json:"success"`
	JSON    zeroTrustAccessApplicationRevokeTokensResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessApplicationRevokeTokensResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationRevokeTokensResponseEnvelope]
type zeroTrustAccessApplicationRevokeTokensResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationRevokeTokensResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationRevokeTokensResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationRevokeTokensResponseEnvelopeSuccessTrue  ZeroTrustAccessApplicationRevokeTokensResponseEnvelopeSuccess = true
	ZeroTrustAccessApplicationRevokeTokensResponseEnvelopeSuccessFalse ZeroTrustAccessApplicationRevokeTokensResponseEnvelopeSuccess = false
)

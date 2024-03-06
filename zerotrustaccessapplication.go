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
func (r *ZeroTrustAccessApplicationService) New(ctx context.Context, params ZeroTrustAccessApplicationNewParams, opts ...option.RequestOption) (res *ZeroTrustAccessApplicationNewResponse, err error) {
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
func (r *ZeroTrustAccessApplicationService) Update(ctx context.Context, appID ZeroTrustAccessApplicationUpdateParamsAppID, params ZeroTrustAccessApplicationUpdateParams, opts ...option.RequestOption) (res *ZeroTrustAccessApplicationUpdateResponse, err error) {
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
func (r *ZeroTrustAccessApplicationService) List(ctx context.Context, query ZeroTrustAccessApplicationListParams, opts ...option.RequestOption) (res *[]ZeroTrustAccessApplicationListResponse, err error) {
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
func (r *ZeroTrustAccessApplicationService) Get(ctx context.Context, appID ZeroTrustAccessApplicationGetParamsAppID, query ZeroTrustAccessApplicationGetParams, opts ...option.RequestOption) (res *ZeroTrustAccessApplicationGetResponse, err error) {
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

// Union satisfied by [ZeroTrustAccessApplicationNewResponseSelfHostedApplication],
// [ZeroTrustAccessApplicationNewResponseSaaSApplication],
// [ZeroTrustAccessApplicationNewResponseBrowserSSHApplication],
// [ZeroTrustAccessApplicationNewResponseBrowserVncApplication],
// [ZeroTrustAccessApplicationNewResponseAppLauncherApplication],
// [ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplication],
// [ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplication] or
// [ZeroTrustAccessApplicationNewResponseBookmarkApplication].
type ZeroTrustAccessApplicationNewResponse interface {
	implementsZeroTrustAccessApplicationNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessApplicationNewResponse)(nil)).Elem(), "")
}

type ZeroTrustAccessApplicationNewResponseSelfHostedApplication struct {
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
	AutoRedirectToIdentity bool                                                                  `json:"auto_redirect_to_identity"`
	CorsHeaders            ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                             `json:"created_at" format:"date-time"`
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
	Tags      []string                                                       `json:"tags"`
	UpdatedAt time.Time                                                      `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationNewResponseSelfHostedApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationNewResponseSelfHostedApplicationJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationNewResponseSelfHostedApplication]
type zeroTrustAccessApplicationNewResponseSelfHostedApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationNewResponseSelfHostedApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationNewResponseSelfHostedApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationNewResponseSelfHostedApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

type ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeaders struct {
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
	AllowedMethods []ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                   `json:"max_age"`
	JSON   zeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersJSON `json:"-"`
}

// zeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeaders]
type zeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersJSON struct {
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

func (r *ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod string

const (
	ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodGet     ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "GET"
	ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodPost    ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "POST"
	ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodHead    ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "HEAD"
	ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodPut     ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PUT"
	ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodDelete  ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "DELETE"
	ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodConnect ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "CONNECT"
	ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodOptions ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "OPTIONS"
	ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodTrace   ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "TRACE"
	ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodPatch   ZeroTrustAccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PATCH"
)

type ZeroTrustAccessApplicationNewResponseSaaSApplication struct {
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
	Name    string                                                      `json:"name"`
	SaasApp ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasApp `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags []string `json:"tags"`
	// The application type.
	Type      string                                                   `json:"type"`
	UpdatedAt time.Time                                                `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationNewResponseSaaSApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationNewResponseSaaSApplicationJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationNewResponseSaaSApplication]
type zeroTrustAccessApplicationNewResponseSaaSApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationNewResponseSaaSApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationNewResponseSaaSApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationNewResponseSaaSApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

// Union satisfied by
// [ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasApp]
// or
// [ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasApp].
type ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasApp interface {
	implementsZeroTrustAccessApplicationNewResponseSaaSApplicationSaasApp()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasApp)(nil)).Elem(), "")
}

type ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string                                                                                       `json:"consumer_service_url"`
	CreatedAt          time.Time                                                                                    `json:"created_at" format:"date-time"`
	CustomAttributes   ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat `json:"name_id_format"`
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
	SSOEndpoint string                                                                           `json:"sso_endpoint"`
	UpdatedAt   time.Time                                                                        `json:"updated_at" format:"date-time"`
	JSON        zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON `json:"-"`
}

// zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasApp]
type zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON struct {
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

func (r *ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasApp) implementsZeroTrustAccessApplicationNewResponseSaaSApplicationSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType string

const (
	ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeSaml ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "saml"
	ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeOidc ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "oidc"
)

type ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name string `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat `json:"name_format"`
	Source     ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource     `json:"source"`
	JSON       zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON       `json:"-"`
}

// zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes]
type zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON struct {
	Name        apijson.Field
	NameFormat  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON) RawJSON() string {
	return r.raw
}

// A globally unique name for an identity or service provider.
type ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name string                                                                                                 `json:"name"`
	JSON zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON `json:"-"`
}

// zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource]
type zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON) RawJSON() string {
	return r.raw
}

// The format of the name identifier sent to the SaaS application.
type ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat string

const (
	ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatID    ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "id"
	ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatEmail ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The OIDC flows supported by this application
	GrantTypes []ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs []string `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes    []ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScope `json:"scopes"`
	UpdatedAt time.Time                                                                           `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON    `json:"-"`
}

// zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasApp]
type zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON struct {
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

func (r *ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationNewResponseSaaSApplicationSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType string

const (
	ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "saml"
	ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "oidc"
)

type ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType string

const (
	ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScope string

const (
	ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid  ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "openid"
	ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups  ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "groups"
	ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail   ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "email"
	ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile ZeroTrustAccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "profile"
)

type ZeroTrustAccessApplicationNewResponseBrowserSSHApplication struct {
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
	AutoRedirectToIdentity bool                                                                  `json:"auto_redirect_to_identity"`
	CorsHeaders            ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                             `json:"created_at" format:"date-time"`
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
	Tags      []string                                                       `json:"tags"`
	UpdatedAt time.Time                                                      `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationNewResponseBrowserSSHApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationNewResponseBrowserSSHApplicationJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationNewResponseBrowserSSHApplication]
type zeroTrustAccessApplicationNewResponseBrowserSSHApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationNewResponseBrowserSSHApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationNewResponseBrowserSSHApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationNewResponseBrowserSSHApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

type ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeaders struct {
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
	AllowedMethods []ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                   `json:"max_age"`
	JSON   zeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersJSON `json:"-"`
}

// zeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeaders]
type zeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersJSON struct {
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

func (r *ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod string

const (
	ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodGet     ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "GET"
	ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodPost    ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "POST"
	ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodHead    ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "HEAD"
	ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodPut     ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PUT"
	ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodDelete  ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "DELETE"
	ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodConnect ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "CONNECT"
	ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodOptions ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "OPTIONS"
	ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodTrace   ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "TRACE"
	ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodPatch   ZeroTrustAccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PATCH"
)

type ZeroTrustAccessApplicationNewResponseBrowserVncApplication struct {
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
	AutoRedirectToIdentity bool                                                                  `json:"auto_redirect_to_identity"`
	CorsHeaders            ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                             `json:"created_at" format:"date-time"`
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
	Tags      []string                                                       `json:"tags"`
	UpdatedAt time.Time                                                      `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationNewResponseBrowserVncApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationNewResponseBrowserVncApplicationJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationNewResponseBrowserVncApplication]
type zeroTrustAccessApplicationNewResponseBrowserVncApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationNewResponseBrowserVncApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationNewResponseBrowserVncApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationNewResponseBrowserVncApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

type ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeaders struct {
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
	AllowedMethods []ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                   `json:"max_age"`
	JSON   zeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersJSON `json:"-"`
}

// zeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeaders]
type zeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersJSON struct {
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

func (r *ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod string

const (
	ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodGet     ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "GET"
	ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodPost    ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "POST"
	ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodHead    ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "HEAD"
	ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodPut     ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PUT"
	ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodDelete  ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "DELETE"
	ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodConnect ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "CONNECT"
	ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodOptions ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "OPTIONS"
	ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodTrace   ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "TRACE"
	ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodPatch   ZeroTrustAccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PATCH"
)

type ZeroTrustAccessApplicationNewResponseAppLauncherApplication struct {
	// The application type.
	Type ZeroTrustAccessApplicationNewResponseAppLauncherApplicationType `json:"type,required"`
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
	SessionDuration string                                                          `json:"session_duration"`
	UpdatedAt       time.Time                                                       `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationNewResponseAppLauncherApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationNewResponseAppLauncherApplicationJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationNewResponseAppLauncherApplication]
type zeroTrustAccessApplicationNewResponseAppLauncherApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationNewResponseAppLauncherApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationNewResponseAppLauncherApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationNewResponseAppLauncherApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

// The application type.
type ZeroTrustAccessApplicationNewResponseAppLauncherApplicationType string

const (
	ZeroTrustAccessApplicationNewResponseAppLauncherApplicationTypeSelfHosted  ZeroTrustAccessApplicationNewResponseAppLauncherApplicationType = "self_hosted"
	ZeroTrustAccessApplicationNewResponseAppLauncherApplicationTypeSaas        ZeroTrustAccessApplicationNewResponseAppLauncherApplicationType = "saas"
	ZeroTrustAccessApplicationNewResponseAppLauncherApplicationTypeSSH         ZeroTrustAccessApplicationNewResponseAppLauncherApplicationType = "ssh"
	ZeroTrustAccessApplicationNewResponseAppLauncherApplicationTypeVnc         ZeroTrustAccessApplicationNewResponseAppLauncherApplicationType = "vnc"
	ZeroTrustAccessApplicationNewResponseAppLauncherApplicationTypeAppLauncher ZeroTrustAccessApplicationNewResponseAppLauncherApplicationType = "app_launcher"
	ZeroTrustAccessApplicationNewResponseAppLauncherApplicationTypeWARP        ZeroTrustAccessApplicationNewResponseAppLauncherApplicationType = "warp"
	ZeroTrustAccessApplicationNewResponseAppLauncherApplicationTypeBiso        ZeroTrustAccessApplicationNewResponseAppLauncherApplicationType = "biso"
	ZeroTrustAccessApplicationNewResponseAppLauncherApplicationTypeBookmark    ZeroTrustAccessApplicationNewResponseAppLauncherApplicationType = "bookmark"
	ZeroTrustAccessApplicationNewResponseAppLauncherApplicationTypeDashSSO     ZeroTrustAccessApplicationNewResponseAppLauncherApplicationType = "dash_sso"
)

type ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                          `json:"session_duration"`
	UpdatedAt       time.Time                                                                       `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplication]
type zeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

// The application type.
type ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType string

const (
	ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeSelfHosted  ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "self_hosted"
	ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeSaas        ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "saas"
	ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeSSH         ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "ssh"
	ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeVnc         ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "vnc"
	ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeAppLauncher ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "app_launcher"
	ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeWARP        ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "warp"
	ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeBiso        ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "biso"
	ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeBookmark    ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "bookmark"
	ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeDashSSO     ZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "dash_sso"
)

type ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                          `json:"session_duration"`
	UpdatedAt       time.Time                                                                       `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplication]
type zeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

// The application type.
type ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationType string

const (
	ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeSelfHosted  ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "self_hosted"
	ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeSaas        ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "saas"
	ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeSSH         ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "ssh"
	ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeVnc         ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "vnc"
	ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeAppLauncher ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "app_launcher"
	ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeWARP        ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "warp"
	ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeBiso        ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "biso"
	ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeBookmark    ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "bookmark"
	ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeDashSSO     ZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "dash_sso"
)

type ZeroTrustAccessApplicationNewResponseBookmarkApplication struct {
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
	Type      string                                                       `json:"type"`
	UpdatedAt time.Time                                                    `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationNewResponseBookmarkApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationNewResponseBookmarkApplicationJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationNewResponseBookmarkApplication]
type zeroTrustAccessApplicationNewResponseBookmarkApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationNewResponseBookmarkApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationNewResponseBookmarkApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationNewResponseBookmarkApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

// Union satisfied by
// [ZeroTrustAccessApplicationUpdateResponseSelfHostedApplication],
// [ZeroTrustAccessApplicationUpdateResponseSaaSApplication],
// [ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplication],
// [ZeroTrustAccessApplicationUpdateResponseBrowserVncApplication],
// [ZeroTrustAccessApplicationUpdateResponseAppLauncherApplication],
// [ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplication],
// [ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplication]
// or [ZeroTrustAccessApplicationUpdateResponseBookmarkApplication].
type ZeroTrustAccessApplicationUpdateResponse interface {
	implementsZeroTrustAccessApplicationUpdateResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessApplicationUpdateResponse)(nil)).Elem(), "")
}

type ZeroTrustAccessApplicationUpdateResponseSelfHostedApplication struct {
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
	AutoRedirectToIdentity bool                                                                     `json:"auto_redirect_to_identity"`
	CorsHeaders            ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                                `json:"created_at" format:"date-time"`
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
	Tags      []string                                                          `json:"tags"`
	UpdatedAt time.Time                                                         `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationUpdateResponseSelfHostedApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseSelfHostedApplicationJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseSelfHostedApplication]
type zeroTrustAccessApplicationUpdateResponseSelfHostedApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationUpdateResponseSelfHostedApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationUpdateResponseSelfHostedApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationUpdateResponseSelfHostedApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

type ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeaders struct {
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
	AllowedMethods []ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                      `json:"max_age"`
	JSON   zeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersJSON `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeaders]
type zeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersJSON struct {
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

func (r *ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod string

const (
	ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodGet     ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "GET"
	ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodPost    ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "POST"
	ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodHead    ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "HEAD"
	ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodPut     ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PUT"
	ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodDelete  ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "DELETE"
	ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodConnect ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "CONNECT"
	ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodOptions ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "OPTIONS"
	ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodTrace   ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "TRACE"
	ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodPatch   ZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PATCH"
)

type ZeroTrustAccessApplicationUpdateResponseSaaSApplication struct {
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
	Name    string                                                         `json:"name"`
	SaasApp ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasApp `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags []string `json:"tags"`
	// The application type.
	Type      string                                                      `json:"type"`
	UpdatedAt time.Time                                                   `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationUpdateResponseSaaSApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseSaaSApplicationJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseSaaSApplication]
type zeroTrustAccessApplicationUpdateResponseSaaSApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationUpdateResponseSaaSApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationUpdateResponseSaaSApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationUpdateResponseSaaSApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

// Union satisfied by
// [ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasApp]
// or
// [ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasApp].
type ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasApp interface {
	implementsZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasApp()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasApp)(nil)).Elem(), "")
}

type ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string                                                                                          `json:"consumer_service_url"`
	CreatedAt          time.Time                                                                                       `json:"created_at" format:"date-time"`
	CustomAttributes   ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat `json:"name_id_format"`
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
	SSOEndpoint string                                                                              `json:"sso_endpoint"`
	UpdatedAt   time.Time                                                                           `json:"updated_at" format:"date-time"`
	JSON        zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasApp]
type zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON struct {
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

func (r *ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasApp) implementsZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType string

const (
	ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeSaml ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "saml"
	ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeOidc ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "oidc"
)

type ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name string `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat `json:"name_format"`
	Source     ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource     `json:"source"`
	JSON       zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON       `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes]
type zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON struct {
	Name        apijson.Field
	NameFormat  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON) RawJSON() string {
	return r.raw
}

// A globally unique name for an identity or service provider.
type ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name string                                                                                                    `json:"name"`
	JSON zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource]
type zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON) RawJSON() string {
	return r.raw
}

// The format of the name identifier sent to the SaaS application.
type ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat string

const (
	ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatID    ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "id"
	ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatEmail ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The OIDC flows supported by this application
	GrantTypes []ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs []string `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes    []ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScope `json:"scopes"`
	UpdatedAt time.Time                                                                              `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON    `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasApp]
type zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON struct {
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

func (r *ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType string

const (
	ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "saml"
	ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "oidc"
)

type ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType string

const (
	ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScope string

const (
	ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid  ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "openid"
	ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups  ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "groups"
	ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail   ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "email"
	ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile ZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "profile"
)

type ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplication struct {
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
	AutoRedirectToIdentity bool                                                                     `json:"auto_redirect_to_identity"`
	CorsHeaders            ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                                `json:"created_at" format:"date-time"`
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
	Tags      []string                                                          `json:"tags"`
	UpdatedAt time.Time                                                         `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplication]
type zeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

type ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeaders struct {
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
	AllowedMethods []ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                      `json:"max_age"`
	JSON   zeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersJSON `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeaders]
type zeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersJSON struct {
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

func (r *ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod string

const (
	ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodGet     ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "GET"
	ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodPost    ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "POST"
	ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodHead    ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "HEAD"
	ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodPut     ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PUT"
	ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodDelete  ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "DELETE"
	ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodConnect ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "CONNECT"
	ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodOptions ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "OPTIONS"
	ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodTrace   ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "TRACE"
	ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodPatch   ZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PATCH"
)

type ZeroTrustAccessApplicationUpdateResponseBrowserVncApplication struct {
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
	AutoRedirectToIdentity bool                                                                     `json:"auto_redirect_to_identity"`
	CorsHeaders            ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                                `json:"created_at" format:"date-time"`
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
	Tags      []string                                                          `json:"tags"`
	UpdatedAt time.Time                                                         `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationUpdateResponseBrowserVncApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseBrowserVncApplicationJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseBrowserVncApplication]
type zeroTrustAccessApplicationUpdateResponseBrowserVncApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationUpdateResponseBrowserVncApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationUpdateResponseBrowserVncApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationUpdateResponseBrowserVncApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

type ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeaders struct {
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
	AllowedMethods []ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                      `json:"max_age"`
	JSON   zeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersJSON `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeaders]
type zeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersJSON struct {
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

func (r *ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod string

const (
	ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodGet     ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "GET"
	ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodPost    ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "POST"
	ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodHead    ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "HEAD"
	ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodPut     ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PUT"
	ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodDelete  ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "DELETE"
	ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodConnect ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "CONNECT"
	ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodOptions ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "OPTIONS"
	ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodTrace   ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "TRACE"
	ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodPatch   ZeroTrustAccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PATCH"
)

type ZeroTrustAccessApplicationUpdateResponseAppLauncherApplication struct {
	// The application type.
	Type ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationType `json:"type,required"`
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
	SessionDuration string                                                             `json:"session_duration"`
	UpdatedAt       time.Time                                                          `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationUpdateResponseAppLauncherApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseAppLauncherApplicationJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseAppLauncherApplication]
type zeroTrustAccessApplicationUpdateResponseAppLauncherApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationUpdateResponseAppLauncherApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationUpdateResponseAppLauncherApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationUpdateResponseAppLauncherApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

// The application type.
type ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationType string

const (
	ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationTypeSelfHosted  ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationType = "self_hosted"
	ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationTypeSaas        ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationType = "saas"
	ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationTypeSSH         ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationType = "ssh"
	ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationTypeVnc         ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationType = "vnc"
	ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationTypeAppLauncher ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationType = "app_launcher"
	ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationTypeWARP        ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationType = "warp"
	ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationTypeBiso        ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationType = "biso"
	ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationTypeBookmark    ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationType = "bookmark"
	ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationTypeDashSSO     ZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationType = "dash_sso"
)

type ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                             `json:"session_duration"`
	UpdatedAt       time.Time                                                                          `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplication]
type zeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

// The application type.
type ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType string

const (
	ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeSelfHosted  ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "self_hosted"
	ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeSaas        ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "saas"
	ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeSSH         ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "ssh"
	ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeVnc         ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "vnc"
	ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeAppLauncher ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "app_launcher"
	ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeWARP        ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "warp"
	ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeBiso        ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "biso"
	ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeBookmark    ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "bookmark"
	ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeDashSSO     ZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "dash_sso"
)

type ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                             `json:"session_duration"`
	UpdatedAt       time.Time                                                                          `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplication]
type zeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

// The application type.
type ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType string

const (
	ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeSelfHosted  ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "self_hosted"
	ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeSaas        ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "saas"
	ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeSSH         ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "ssh"
	ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeVnc         ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "vnc"
	ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeAppLauncher ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "app_launcher"
	ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeWARP        ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "warp"
	ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeBiso        ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "biso"
	ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeBookmark    ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "bookmark"
	ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeDashSSO     ZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "dash_sso"
)

type ZeroTrustAccessApplicationUpdateResponseBookmarkApplication struct {
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
	Type      string                                                          `json:"type"`
	UpdatedAt time.Time                                                       `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationUpdateResponseBookmarkApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationUpdateResponseBookmarkApplicationJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationUpdateResponseBookmarkApplication]
type zeroTrustAccessApplicationUpdateResponseBookmarkApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationUpdateResponseBookmarkApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationUpdateResponseBookmarkApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationUpdateResponseBookmarkApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

// Union satisfied by
// [ZeroTrustAccessApplicationListResponseSelfHostedApplication],
// [ZeroTrustAccessApplicationListResponseSaaSApplication],
// [ZeroTrustAccessApplicationListResponseBrowserSSHApplication],
// [ZeroTrustAccessApplicationListResponseBrowserVncApplication],
// [ZeroTrustAccessApplicationListResponseAppLauncherApplication],
// [ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplication],
// [ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplication]
// or [ZeroTrustAccessApplicationListResponseBookmarkApplication].
type ZeroTrustAccessApplicationListResponse interface {
	implementsZeroTrustAccessApplicationListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessApplicationListResponse)(nil)).Elem(), "")
}

type ZeroTrustAccessApplicationListResponseSelfHostedApplication struct {
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
	AutoRedirectToIdentity bool                                                                   `json:"auto_redirect_to_identity"`
	CorsHeaders            ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                              `json:"created_at" format:"date-time"`
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
	Tags      []string                                                        `json:"tags"`
	UpdatedAt time.Time                                                       `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationListResponseSelfHostedApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseSelfHostedApplicationJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationListResponseSelfHostedApplication]
type zeroTrustAccessApplicationListResponseSelfHostedApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationListResponseSelfHostedApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationListResponseSelfHostedApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationListResponseSelfHostedApplication) implementsZeroTrustAccessApplicationListResponse() {
}

type ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeaders struct {
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
	AllowedMethods []ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                    `json:"max_age"`
	JSON   zeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeaders]
type zeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersJSON struct {
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

func (r *ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod string

const (
	ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodGet     ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "GET"
	ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodPost    ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "POST"
	ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodHead    ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "HEAD"
	ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodPut     ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PUT"
	ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodDelete  ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "DELETE"
	ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodConnect ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "CONNECT"
	ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodOptions ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "OPTIONS"
	ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodTrace   ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "TRACE"
	ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodPatch   ZeroTrustAccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PATCH"
)

type ZeroTrustAccessApplicationListResponseSaaSApplication struct {
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
	Name    string                                                       `json:"name"`
	SaasApp ZeroTrustAccessApplicationListResponseSaaSApplicationSaasApp `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags []string `json:"tags"`
	// The application type.
	Type      string                                                    `json:"type"`
	UpdatedAt time.Time                                                 `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationListResponseSaaSApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseSaaSApplicationJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationListResponseSaaSApplication]
type zeroTrustAccessApplicationListResponseSaaSApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationListResponseSaaSApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationListResponseSaaSApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationListResponseSaaSApplication) implementsZeroTrustAccessApplicationListResponse() {
}

// Union satisfied by
// [ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasApp]
// or
// [ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasApp].
type ZeroTrustAccessApplicationListResponseSaaSApplicationSaasApp interface {
	implementsZeroTrustAccessApplicationListResponseSaaSApplicationSaasApp()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessApplicationListResponseSaaSApplicationSaasApp)(nil)).Elem(), "")
}

type ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string                                                                                        `json:"consumer_service_url"`
	CreatedAt          time.Time                                                                                     `json:"created_at" format:"date-time"`
	CustomAttributes   ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat `json:"name_id_format"`
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
	SSOEndpoint string                                                                            `json:"sso_endpoint"`
	UpdatedAt   time.Time                                                                         `json:"updated_at" format:"date-time"`
	JSON        zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasApp]
type zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON struct {
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

func (r *ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasApp) implementsZeroTrustAccessApplicationListResponseSaaSApplicationSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType string

const (
	ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeSaml ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "saml"
	ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeOidc ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "oidc"
)

type ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name string `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat `json:"name_format"`
	Source     ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource     `json:"source"`
	JSON       zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON       `json:"-"`
}

// zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes]
type zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON struct {
	Name        apijson.Field
	NameFormat  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON) RawJSON() string {
	return r.raw
}

// A globally unique name for an identity or service provider.
type ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name string                                                                                                  `json:"name"`
	JSON zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource]
type zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON) RawJSON() string {
	return r.raw
}

// The format of the name identifier sent to the SaaS application.
type ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat string

const (
	ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatID    ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "id"
	ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatEmail ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The OIDC flows supported by this application
	GrantTypes []ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs []string `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes    []ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScope `json:"scopes"`
	UpdatedAt time.Time                                                                            `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON    `json:"-"`
}

// zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasApp]
type zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON struct {
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

func (r *ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationListResponseSaaSApplicationSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType string

const (
	ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "saml"
	ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "oidc"
)

type ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType string

const (
	ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScope string

const (
	ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid  ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "openid"
	ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups  ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "groups"
	ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail   ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "email"
	ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile ZeroTrustAccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "profile"
)

type ZeroTrustAccessApplicationListResponseBrowserSSHApplication struct {
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
	AutoRedirectToIdentity bool                                                                   `json:"auto_redirect_to_identity"`
	CorsHeaders            ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                              `json:"created_at" format:"date-time"`
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
	Tags      []string                                                        `json:"tags"`
	UpdatedAt time.Time                                                       `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationListResponseBrowserSSHApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseBrowserSSHApplicationJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationListResponseBrowserSSHApplication]
type zeroTrustAccessApplicationListResponseBrowserSSHApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationListResponseBrowserSSHApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationListResponseBrowserSSHApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationListResponseBrowserSSHApplication) implementsZeroTrustAccessApplicationListResponse() {
}

type ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeaders struct {
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
	AllowedMethods []ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                    `json:"max_age"`
	JSON   zeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeaders]
type zeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersJSON struct {
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

func (r *ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod string

const (
	ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodGet     ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "GET"
	ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodPost    ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "POST"
	ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodHead    ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "HEAD"
	ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodPut     ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PUT"
	ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodDelete  ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "DELETE"
	ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodConnect ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "CONNECT"
	ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodOptions ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "OPTIONS"
	ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodTrace   ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "TRACE"
	ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodPatch   ZeroTrustAccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PATCH"
)

type ZeroTrustAccessApplicationListResponseBrowserVncApplication struct {
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
	AutoRedirectToIdentity bool                                                                   `json:"auto_redirect_to_identity"`
	CorsHeaders            ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                              `json:"created_at" format:"date-time"`
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
	Tags      []string                                                        `json:"tags"`
	UpdatedAt time.Time                                                       `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationListResponseBrowserVncApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseBrowserVncApplicationJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationListResponseBrowserVncApplication]
type zeroTrustAccessApplicationListResponseBrowserVncApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationListResponseBrowserVncApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationListResponseBrowserVncApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationListResponseBrowserVncApplication) implementsZeroTrustAccessApplicationListResponse() {
}

type ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeaders struct {
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
	AllowedMethods []ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                    `json:"max_age"`
	JSON   zeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeaders]
type zeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersJSON struct {
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

func (r *ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod string

const (
	ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodGet     ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "GET"
	ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodPost    ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "POST"
	ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodHead    ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "HEAD"
	ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodPut     ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PUT"
	ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodDelete  ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "DELETE"
	ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodConnect ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "CONNECT"
	ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodOptions ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "OPTIONS"
	ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodTrace   ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "TRACE"
	ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodPatch   ZeroTrustAccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PATCH"
)

type ZeroTrustAccessApplicationListResponseAppLauncherApplication struct {
	// The application type.
	Type ZeroTrustAccessApplicationListResponseAppLauncherApplicationType `json:"type,required"`
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
	SessionDuration string                                                           `json:"session_duration"`
	UpdatedAt       time.Time                                                        `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationListResponseAppLauncherApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseAppLauncherApplicationJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationListResponseAppLauncherApplication]
type zeroTrustAccessApplicationListResponseAppLauncherApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationListResponseAppLauncherApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationListResponseAppLauncherApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationListResponseAppLauncherApplication) implementsZeroTrustAccessApplicationListResponse() {
}

// The application type.
type ZeroTrustAccessApplicationListResponseAppLauncherApplicationType string

const (
	ZeroTrustAccessApplicationListResponseAppLauncherApplicationTypeSelfHosted  ZeroTrustAccessApplicationListResponseAppLauncherApplicationType = "self_hosted"
	ZeroTrustAccessApplicationListResponseAppLauncherApplicationTypeSaas        ZeroTrustAccessApplicationListResponseAppLauncherApplicationType = "saas"
	ZeroTrustAccessApplicationListResponseAppLauncherApplicationTypeSSH         ZeroTrustAccessApplicationListResponseAppLauncherApplicationType = "ssh"
	ZeroTrustAccessApplicationListResponseAppLauncherApplicationTypeVnc         ZeroTrustAccessApplicationListResponseAppLauncherApplicationType = "vnc"
	ZeroTrustAccessApplicationListResponseAppLauncherApplicationTypeAppLauncher ZeroTrustAccessApplicationListResponseAppLauncherApplicationType = "app_launcher"
	ZeroTrustAccessApplicationListResponseAppLauncherApplicationTypeWARP        ZeroTrustAccessApplicationListResponseAppLauncherApplicationType = "warp"
	ZeroTrustAccessApplicationListResponseAppLauncherApplicationTypeBiso        ZeroTrustAccessApplicationListResponseAppLauncherApplicationType = "biso"
	ZeroTrustAccessApplicationListResponseAppLauncherApplicationTypeBookmark    ZeroTrustAccessApplicationListResponseAppLauncherApplicationType = "bookmark"
	ZeroTrustAccessApplicationListResponseAppLauncherApplicationTypeDashSSO     ZeroTrustAccessApplicationListResponseAppLauncherApplicationType = "dash_sso"
)

type ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                           `json:"session_duration"`
	UpdatedAt       time.Time                                                                        `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplication]
type zeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplication) implementsZeroTrustAccessApplicationListResponse() {
}

// The application type.
type ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType string

const (
	ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeSelfHosted  ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "self_hosted"
	ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeSaas        ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "saas"
	ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeSSH         ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "ssh"
	ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeVnc         ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "vnc"
	ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeAppLauncher ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "app_launcher"
	ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeWARP        ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "warp"
	ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeBiso        ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "biso"
	ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeBookmark    ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "bookmark"
	ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeDashSSO     ZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "dash_sso"
)

type ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                           `json:"session_duration"`
	UpdatedAt       time.Time                                                                        `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplication]
type zeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplication) implementsZeroTrustAccessApplicationListResponse() {
}

// The application type.
type ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationType string

const (
	ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeSelfHosted  ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "self_hosted"
	ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeSaas        ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "saas"
	ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeSSH         ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "ssh"
	ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeVnc         ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "vnc"
	ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeAppLauncher ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "app_launcher"
	ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeWARP        ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "warp"
	ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeBiso        ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "biso"
	ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeBookmark    ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "bookmark"
	ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeDashSSO     ZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "dash_sso"
)

type ZeroTrustAccessApplicationListResponseBookmarkApplication struct {
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
	Type      string                                                        `json:"type"`
	UpdatedAt time.Time                                                     `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationListResponseBookmarkApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationListResponseBookmarkApplicationJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationListResponseBookmarkApplication]
type zeroTrustAccessApplicationListResponseBookmarkApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationListResponseBookmarkApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationListResponseBookmarkApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationListResponseBookmarkApplication) implementsZeroTrustAccessApplicationListResponse() {
}

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

func (r zeroTrustAccessApplicationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ZeroTrustAccessApplicationGetResponseSelfHostedApplication],
// [ZeroTrustAccessApplicationGetResponseSaaSApplication],
// [ZeroTrustAccessApplicationGetResponseBrowserSSHApplication],
// [ZeroTrustAccessApplicationGetResponseBrowserVncApplication],
// [ZeroTrustAccessApplicationGetResponseAppLauncherApplication],
// [ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplication],
// [ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplication] or
// [ZeroTrustAccessApplicationGetResponseBookmarkApplication].
type ZeroTrustAccessApplicationGetResponse interface {
	implementsZeroTrustAccessApplicationGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessApplicationGetResponse)(nil)).Elem(), "")
}

type ZeroTrustAccessApplicationGetResponseSelfHostedApplication struct {
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
	AutoRedirectToIdentity bool                                                                  `json:"auto_redirect_to_identity"`
	CorsHeaders            ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                             `json:"created_at" format:"date-time"`
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
	Tags      []string                                                       `json:"tags"`
	UpdatedAt time.Time                                                      `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationGetResponseSelfHostedApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationGetResponseSelfHostedApplicationJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationGetResponseSelfHostedApplication]
type zeroTrustAccessApplicationGetResponseSelfHostedApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationGetResponseSelfHostedApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationGetResponseSelfHostedApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationGetResponseSelfHostedApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

type ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeaders struct {
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
	AllowedMethods []ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                   `json:"max_age"`
	JSON   zeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersJSON `json:"-"`
}

// zeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeaders]
type zeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersJSON struct {
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

func (r *ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod string

const (
	ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodGet     ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "GET"
	ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodPost    ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "POST"
	ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodHead    ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "HEAD"
	ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodPut     ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PUT"
	ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodDelete  ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "DELETE"
	ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodConnect ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "CONNECT"
	ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodOptions ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "OPTIONS"
	ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodTrace   ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "TRACE"
	ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodPatch   ZeroTrustAccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PATCH"
)

type ZeroTrustAccessApplicationGetResponseSaaSApplication struct {
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
	Name    string                                                      `json:"name"`
	SaasApp ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasApp `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags []string `json:"tags"`
	// The application type.
	Type      string                                                   `json:"type"`
	UpdatedAt time.Time                                                `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationGetResponseSaaSApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationGetResponseSaaSApplicationJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationGetResponseSaaSApplication]
type zeroTrustAccessApplicationGetResponseSaaSApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationGetResponseSaaSApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationGetResponseSaaSApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationGetResponseSaaSApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

// Union satisfied by
// [ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasApp]
// or
// [ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasApp].
type ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasApp interface {
	implementsZeroTrustAccessApplicationGetResponseSaaSApplicationSaasApp()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasApp)(nil)).Elem(), "")
}

type ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string                                                                                       `json:"consumer_service_url"`
	CreatedAt          time.Time                                                                                    `json:"created_at" format:"date-time"`
	CustomAttributes   ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat `json:"name_id_format"`
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
	SSOEndpoint string                                                                           `json:"sso_endpoint"`
	UpdatedAt   time.Time                                                                        `json:"updated_at" format:"date-time"`
	JSON        zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON `json:"-"`
}

// zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasApp]
type zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON struct {
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

func (r *ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasApp) implementsZeroTrustAccessApplicationGetResponseSaaSApplicationSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType string

const (
	ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeSaml ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "saml"
	ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeOidc ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "oidc"
)

type ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name string `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat `json:"name_format"`
	Source     ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource     `json:"source"`
	JSON       zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON       `json:"-"`
}

// zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes]
type zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON struct {
	Name        apijson.Field
	NameFormat  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON) RawJSON() string {
	return r.raw
}

// A globally unique name for an identity or service provider.
type ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name string                                                                                                 `json:"name"`
	JSON zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON `json:"-"`
}

// zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource]
type zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON) RawJSON() string {
	return r.raw
}

// The format of the name identifier sent to the SaaS application.
type ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat string

const (
	ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatID    ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "id"
	ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatEmail ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The OIDC flows supported by this application
	GrantTypes []ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs []string `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes    []ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScope `json:"scopes"`
	UpdatedAt time.Time                                                                           `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON    `json:"-"`
}

// zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasApp]
type zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON struct {
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

func (r *ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationGetResponseSaaSApplicationSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType string

const (
	ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "saml"
	ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "oidc"
)

type ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType string

const (
	ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScope string

const (
	ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid  ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "openid"
	ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups  ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "groups"
	ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail   ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "email"
	ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile ZeroTrustAccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "profile"
)

type ZeroTrustAccessApplicationGetResponseBrowserSSHApplication struct {
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
	AutoRedirectToIdentity bool                                                                  `json:"auto_redirect_to_identity"`
	CorsHeaders            ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                             `json:"created_at" format:"date-time"`
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
	Tags      []string                                                       `json:"tags"`
	UpdatedAt time.Time                                                      `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationGetResponseBrowserSSHApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationGetResponseBrowserSSHApplicationJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationGetResponseBrowserSSHApplication]
type zeroTrustAccessApplicationGetResponseBrowserSSHApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationGetResponseBrowserSSHApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationGetResponseBrowserSSHApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationGetResponseBrowserSSHApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

type ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeaders struct {
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
	AllowedMethods []ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                   `json:"max_age"`
	JSON   zeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersJSON `json:"-"`
}

// zeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeaders]
type zeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersJSON struct {
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

func (r *ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod string

const (
	ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodGet     ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "GET"
	ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodPost    ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "POST"
	ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodHead    ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "HEAD"
	ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodPut     ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PUT"
	ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodDelete  ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "DELETE"
	ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodConnect ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "CONNECT"
	ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodOptions ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "OPTIONS"
	ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodTrace   ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "TRACE"
	ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodPatch   ZeroTrustAccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PATCH"
)

type ZeroTrustAccessApplicationGetResponseBrowserVncApplication struct {
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
	AutoRedirectToIdentity bool                                                                  `json:"auto_redirect_to_identity"`
	CorsHeaders            ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                             `json:"created_at" format:"date-time"`
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
	Tags      []string                                                       `json:"tags"`
	UpdatedAt time.Time                                                      `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationGetResponseBrowserVncApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationGetResponseBrowserVncApplicationJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationGetResponseBrowserVncApplication]
type zeroTrustAccessApplicationGetResponseBrowserVncApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationGetResponseBrowserVncApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationGetResponseBrowserVncApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationGetResponseBrowserVncApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

type ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeaders struct {
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
	AllowedMethods []ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                   `json:"max_age"`
	JSON   zeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersJSON `json:"-"`
}

// zeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeaders]
type zeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersJSON struct {
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

func (r *ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod string

const (
	ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodGet     ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "GET"
	ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodPost    ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "POST"
	ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodHead    ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "HEAD"
	ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodPut     ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PUT"
	ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodDelete  ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "DELETE"
	ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodConnect ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "CONNECT"
	ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodOptions ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "OPTIONS"
	ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodTrace   ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "TRACE"
	ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodPatch   ZeroTrustAccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PATCH"
)

type ZeroTrustAccessApplicationGetResponseAppLauncherApplication struct {
	// The application type.
	Type ZeroTrustAccessApplicationGetResponseAppLauncherApplicationType `json:"type,required"`
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
	SessionDuration string                                                          `json:"session_duration"`
	UpdatedAt       time.Time                                                       `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationGetResponseAppLauncherApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationGetResponseAppLauncherApplicationJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationGetResponseAppLauncherApplication]
type zeroTrustAccessApplicationGetResponseAppLauncherApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationGetResponseAppLauncherApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationGetResponseAppLauncherApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationGetResponseAppLauncherApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

// The application type.
type ZeroTrustAccessApplicationGetResponseAppLauncherApplicationType string

const (
	ZeroTrustAccessApplicationGetResponseAppLauncherApplicationTypeSelfHosted  ZeroTrustAccessApplicationGetResponseAppLauncherApplicationType = "self_hosted"
	ZeroTrustAccessApplicationGetResponseAppLauncherApplicationTypeSaas        ZeroTrustAccessApplicationGetResponseAppLauncherApplicationType = "saas"
	ZeroTrustAccessApplicationGetResponseAppLauncherApplicationTypeSSH         ZeroTrustAccessApplicationGetResponseAppLauncherApplicationType = "ssh"
	ZeroTrustAccessApplicationGetResponseAppLauncherApplicationTypeVnc         ZeroTrustAccessApplicationGetResponseAppLauncherApplicationType = "vnc"
	ZeroTrustAccessApplicationGetResponseAppLauncherApplicationTypeAppLauncher ZeroTrustAccessApplicationGetResponseAppLauncherApplicationType = "app_launcher"
	ZeroTrustAccessApplicationGetResponseAppLauncherApplicationTypeWARP        ZeroTrustAccessApplicationGetResponseAppLauncherApplicationType = "warp"
	ZeroTrustAccessApplicationGetResponseAppLauncherApplicationTypeBiso        ZeroTrustAccessApplicationGetResponseAppLauncherApplicationType = "biso"
	ZeroTrustAccessApplicationGetResponseAppLauncherApplicationTypeBookmark    ZeroTrustAccessApplicationGetResponseAppLauncherApplicationType = "bookmark"
	ZeroTrustAccessApplicationGetResponseAppLauncherApplicationTypeDashSSO     ZeroTrustAccessApplicationGetResponseAppLauncherApplicationType = "dash_sso"
)

type ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                          `json:"session_duration"`
	UpdatedAt       time.Time                                                                       `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplication]
type zeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

// The application type.
type ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType string

const (
	ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeSelfHosted  ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "self_hosted"
	ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeSaas        ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "saas"
	ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeSSH         ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "ssh"
	ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeVnc         ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "vnc"
	ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeAppLauncher ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "app_launcher"
	ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeWARP        ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "warp"
	ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeBiso        ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "biso"
	ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeBookmark    ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "bookmark"
	ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeDashSSO     ZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "dash_sso"
)

type ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                          `json:"session_duration"`
	UpdatedAt       time.Time                                                                       `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplication]
type zeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

// The application type.
type ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationType string

const (
	ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeSelfHosted  ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "self_hosted"
	ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeSaas        ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "saas"
	ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeSSH         ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "ssh"
	ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeVnc         ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "vnc"
	ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeAppLauncher ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "app_launcher"
	ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeWARP        ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "warp"
	ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeBiso        ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "biso"
	ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeBookmark    ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "bookmark"
	ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeDashSSO     ZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "dash_sso"
)

type ZeroTrustAccessApplicationGetResponseBookmarkApplication struct {
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
	Type      string                                                       `json:"type"`
	UpdatedAt time.Time                                                    `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessApplicationGetResponseBookmarkApplicationJSON `json:"-"`
}

// zeroTrustAccessApplicationGetResponseBookmarkApplicationJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationGetResponseBookmarkApplication]
type zeroTrustAccessApplicationGetResponseBookmarkApplicationJSON struct {
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

func (r *ZeroTrustAccessApplicationGetResponseBookmarkApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationGetResponseBookmarkApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationGetResponseBookmarkApplication) implementsZeroTrustAccessApplicationGetResponse() {
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
	ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
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
	RedirectURIs param.Field[[]string] `json:"redirect_uris"`
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
	Result   ZeroTrustAccessApplicationNewResponse                   `json:"result,required"`
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

func (r zeroTrustAccessApplicationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
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
	RedirectURIs param.Field[[]string] `json:"redirect_uris"`
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
	Result   ZeroTrustAccessApplicationUpdateResponse                   `json:"result,required"`
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

func (r zeroTrustAccessApplicationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   []ZeroTrustAccessApplicationListResponse                 `json:"result,required,nullable"`
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

func (r zeroTrustAccessApplicationListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   ZeroTrustAccessApplicationGetResponse                   `json:"result,required"`
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

func (r zeroTrustAccessApplicationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationRevokeTokensResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationRevokeTokensResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationRevokeTokensResponseEnvelopeSuccessTrue  ZeroTrustAccessApplicationRevokeTokensResponseEnvelopeSuccess = true
	ZeroTrustAccessApplicationRevokeTokensResponseEnvelopeSuccessFalse ZeroTrustAccessApplicationRevokeTokensResponseEnvelopeSuccess = false
)

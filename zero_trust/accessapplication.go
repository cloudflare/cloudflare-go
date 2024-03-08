// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *AccessApplicationService) New(ctx context.Context, params AccessApplicationNewParams, opts ...option.RequestOption) (res *AccessApplicationNewResponse, err error) {
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
func (r *AccessApplicationService) Update(ctx context.Context, appID AccessApplicationUpdateParamsAppID, params AccessApplicationUpdateParams, opts ...option.RequestOption) (res *AccessApplicationUpdateResponse, err error) {
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
func (r *AccessApplicationService) List(ctx context.Context, query AccessApplicationListParams, opts ...option.RequestOption) (res *[]AccessApplicationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationListResponseEnvelope
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
func (r *AccessApplicationService) Delete(ctx context.Context, appID AccessApplicationDeleteParamsAppID, body AccessApplicationDeleteParams, opts ...option.RequestOption) (res *AccessApplicationDeleteResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches information about an Access application.
func (r *AccessApplicationService) Get(ctx context.Context, appID AccessApplicationGetParamsAppID, query AccessApplicationGetParams, opts ...option.RequestOption) (res *AccessApplicationGetResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Revokes all tokens issued for an application.
func (r *AccessApplicationService) RevokeTokens(ctx context.Context, appID AccessApplicationRevokeTokensParamsAppID, body AccessApplicationRevokeTokensParams, opts ...option.RequestOption) (res *AccessApplicationRevokeTokensResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [zero_trust.AccessApplicationNewResponseSelfHostedApplication],
// [zero_trust.AccessApplicationNewResponseSaaSApplication],
// [zero_trust.AccessApplicationNewResponseBrowserSSHApplication],
// [zero_trust.AccessApplicationNewResponseBrowserVncApplication],
// [zero_trust.AccessApplicationNewResponseAppLauncherApplication],
// [zero_trust.AccessApplicationNewResponseDeviceEnrollmentPermissionsApplication],
// [zero_trust.AccessApplicationNewResponseBrowserIsolationPermissionsApplication]
// or [zero_trust.AccessApplicationNewResponseBookmarkApplication].
type AccessApplicationNewResponse interface {
	implementsZeroTrustAccessApplicationNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationNewResponseSelfHostedApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationNewResponseSaaSApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationNewResponseBrowserSSHApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationNewResponseBrowserVncApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationNewResponseAppLauncherApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationNewResponseDeviceEnrollmentPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationNewResponseBrowserIsolationPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationNewResponseBookmarkApplication{}),
		},
	)
}

type AccessApplicationNewResponseSelfHostedApplication struct {
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
	AutoRedirectToIdentity bool                                                         `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationNewResponseSelfHostedApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                    `json:"created_at" format:"date-time"`
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
	Tags      []string                                              `json:"tags"`
	UpdatedAt time.Time                                             `json:"updated_at" format:"date-time"`
	JSON      accessApplicationNewResponseSelfHostedApplicationJSON `json:"-"`
}

// accessApplicationNewResponseSelfHostedApplicationJSON contains the JSON metadata
// for the struct [AccessApplicationNewResponseSelfHostedApplication]
type accessApplicationNewResponseSelfHostedApplicationJSON struct {
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

func (r *AccessApplicationNewResponseSelfHostedApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseSelfHostedApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationNewResponseSelfHostedApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

type AccessApplicationNewResponseSelfHostedApplicationCorsHeaders struct {
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
	AllowedMethods []AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                          `json:"max_age"`
	JSON   accessApplicationNewResponseSelfHostedApplicationCorsHeadersJSON `json:"-"`
}

// accessApplicationNewResponseSelfHostedApplicationCorsHeadersJSON contains the
// JSON metadata for the struct
// [AccessApplicationNewResponseSelfHostedApplicationCorsHeaders]
type accessApplicationNewResponseSelfHostedApplicationCorsHeadersJSON struct {
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

func (r *AccessApplicationNewResponseSelfHostedApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseSelfHostedApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodGet     AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodPost    AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodHead    AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodPut     AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodDelete  AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodConnect AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodOptions AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodTrace   AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethodPatch   AccessApplicationNewResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationNewResponseSaaSApplication struct {
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
	Name    string                                             `json:"name"`
	SaasApp AccessApplicationNewResponseSaaSApplicationSaasApp `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags []string `json:"tags"`
	// The application type.
	Type      string                                          `json:"type"`
	UpdatedAt time.Time                                       `json:"updated_at" format:"date-time"`
	JSON      accessApplicationNewResponseSaaSApplicationJSON `json:"-"`
}

// accessApplicationNewResponseSaaSApplicationJSON contains the JSON metadata for
// the struct [AccessApplicationNewResponseSaaSApplication]
type accessApplicationNewResponseSaaSApplicationJSON struct {
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

func (r *AccessApplicationNewResponseSaaSApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseSaaSApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationNewResponseSaaSApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

// Union satisfied by
// [zero_trust.AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasApp]
// or
// [zero_trust.AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasApp].
type AccessApplicationNewResponseSaaSApplicationSaasApp interface {
	implementsZeroTrustAccessApplicationNewResponseSaaSApplicationSaasApp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationNewResponseSaaSApplicationSaasApp)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasApp{}),
		},
	)
}

type AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string                                                                              `json:"consumer_service_url"`
	CreatedAt          time.Time                                                                           `json:"created_at" format:"date-time"`
	CustomAttributes   AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat `json:"name_id_format"`
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
	SSOEndpoint string                                                                  `json:"sso_endpoint"`
	UpdatedAt   time.Time                                                               `json:"updated_at" format:"date-time"`
	JSON        accessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON `json:"-"`
}

// accessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON contains
// the JSON metadata for the struct
// [AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasApp]
type accessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON struct {
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

func (r *AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasApp) implementsZeroTrustAccessApplicationNewResponseSaaSApplicationSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType string

const (
	AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeSaml AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "saml"
	AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeOidc AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "oidc"
)

type AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name string `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat `json:"name_format"`
	Source     AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource     `json:"source"`
	JSON       accessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON       `json:"-"`
}

// accessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON
// contains the JSON metadata for the struct
// [AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes]
type accessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON struct {
	Name        apijson.Field
	NameFormat  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON) RawJSON() string {
	return r.raw
}

// A globally unique name for an identity or service provider.
type AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name string                                                                                        `json:"name"`
	JSON accessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON `json:"-"`
}

// accessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON
// contains the JSON metadata for the struct
// [AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource]
type accessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON) RawJSON() string {
	return r.raw
}

// The format of the name identifier sent to the SaaS application.
type AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat string

const (
	AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatID    AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "id"
	AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatEmail AccessApplicationNewResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The OIDC flows supported by this application
	GrantTypes []AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs []string `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes    []AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScope `json:"scopes"`
	UpdatedAt time.Time                                                                  `json:"updated_at" format:"date-time"`
	JSON      accessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON    `json:"-"`
}

// accessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON contains
// the JSON metadata for the struct
// [AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasApp]
type accessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON struct {
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

func (r *AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationNewResponseSaaSApplicationSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType string

const (
	AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "saml"
	AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "oidc"
)

type AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType string

const (
	AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScope string

const (
	AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid  AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "openid"
	AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups  AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "groups"
	AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail   AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "email"
	AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile AccessApplicationNewResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "profile"
)

type AccessApplicationNewResponseBrowserSSHApplication struct {
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
	AutoRedirectToIdentity bool                                                         `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationNewResponseBrowserSSHApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                    `json:"created_at" format:"date-time"`
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
	Tags      []string                                              `json:"tags"`
	UpdatedAt time.Time                                             `json:"updated_at" format:"date-time"`
	JSON      accessApplicationNewResponseBrowserSSHApplicationJSON `json:"-"`
}

// accessApplicationNewResponseBrowserSSHApplicationJSON contains the JSON metadata
// for the struct [AccessApplicationNewResponseBrowserSSHApplication]
type accessApplicationNewResponseBrowserSSHApplicationJSON struct {
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

func (r *AccessApplicationNewResponseBrowserSSHApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseBrowserSSHApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationNewResponseBrowserSSHApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

type AccessApplicationNewResponseBrowserSSHApplicationCorsHeaders struct {
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
	AllowedMethods []AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                          `json:"max_age"`
	JSON   accessApplicationNewResponseBrowserSSHApplicationCorsHeadersJSON `json:"-"`
}

// accessApplicationNewResponseBrowserSSHApplicationCorsHeadersJSON contains the
// JSON metadata for the struct
// [AccessApplicationNewResponseBrowserSSHApplicationCorsHeaders]
type accessApplicationNewResponseBrowserSSHApplicationCorsHeadersJSON struct {
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

func (r *AccessApplicationNewResponseBrowserSSHApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseBrowserSSHApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodGet     AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodPost    AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodHead    AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodPut     AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodDelete  AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodConnect AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodOptions AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodTrace   AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethodPatch   AccessApplicationNewResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationNewResponseBrowserVncApplication struct {
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
	AutoRedirectToIdentity bool                                                         `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationNewResponseBrowserVncApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                    `json:"created_at" format:"date-time"`
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
	Tags      []string                                              `json:"tags"`
	UpdatedAt time.Time                                             `json:"updated_at" format:"date-time"`
	JSON      accessApplicationNewResponseBrowserVncApplicationJSON `json:"-"`
}

// accessApplicationNewResponseBrowserVncApplicationJSON contains the JSON metadata
// for the struct [AccessApplicationNewResponseBrowserVncApplication]
type accessApplicationNewResponseBrowserVncApplicationJSON struct {
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

func (r *AccessApplicationNewResponseBrowserVncApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseBrowserVncApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationNewResponseBrowserVncApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

type AccessApplicationNewResponseBrowserVncApplicationCorsHeaders struct {
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
	AllowedMethods []AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                          `json:"max_age"`
	JSON   accessApplicationNewResponseBrowserVncApplicationCorsHeadersJSON `json:"-"`
}

// accessApplicationNewResponseBrowserVncApplicationCorsHeadersJSON contains the
// JSON metadata for the struct
// [AccessApplicationNewResponseBrowserVncApplicationCorsHeaders]
type accessApplicationNewResponseBrowserVncApplicationCorsHeadersJSON struct {
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

func (r *AccessApplicationNewResponseBrowserVncApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseBrowserVncApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodGet     AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodPost    AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodHead    AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodPut     AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodDelete  AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodConnect AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodOptions AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodTrace   AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethodPatch   AccessApplicationNewResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationNewResponseAppLauncherApplication struct {
	// The application type.
	Type AccessApplicationNewResponseAppLauncherApplicationType `json:"type,required"`
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
	SessionDuration string                                                 `json:"session_duration"`
	UpdatedAt       time.Time                                              `json:"updated_at" format:"date-time"`
	JSON            accessApplicationNewResponseAppLauncherApplicationJSON `json:"-"`
}

// accessApplicationNewResponseAppLauncherApplicationJSON contains the JSON
// metadata for the struct [AccessApplicationNewResponseAppLauncherApplication]
type accessApplicationNewResponseAppLauncherApplicationJSON struct {
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

func (r *AccessApplicationNewResponseAppLauncherApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseAppLauncherApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationNewResponseAppLauncherApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

// The application type.
type AccessApplicationNewResponseAppLauncherApplicationType string

const (
	AccessApplicationNewResponseAppLauncherApplicationTypeSelfHosted  AccessApplicationNewResponseAppLauncherApplicationType = "self_hosted"
	AccessApplicationNewResponseAppLauncherApplicationTypeSaas        AccessApplicationNewResponseAppLauncherApplicationType = "saas"
	AccessApplicationNewResponseAppLauncherApplicationTypeSSH         AccessApplicationNewResponseAppLauncherApplicationType = "ssh"
	AccessApplicationNewResponseAppLauncherApplicationTypeVnc         AccessApplicationNewResponseAppLauncherApplicationType = "vnc"
	AccessApplicationNewResponseAppLauncherApplicationTypeAppLauncher AccessApplicationNewResponseAppLauncherApplicationType = "app_launcher"
	AccessApplicationNewResponseAppLauncherApplicationTypeWARP        AccessApplicationNewResponseAppLauncherApplicationType = "warp"
	AccessApplicationNewResponseAppLauncherApplicationTypeBiso        AccessApplicationNewResponseAppLauncherApplicationType = "biso"
	AccessApplicationNewResponseAppLauncherApplicationTypeBookmark    AccessApplicationNewResponseAppLauncherApplicationType = "bookmark"
	AccessApplicationNewResponseAppLauncherApplicationTypeDashSSO     AccessApplicationNewResponseAppLauncherApplicationType = "dash_sso"
)

type AccessApplicationNewResponseDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                 `json:"session_duration"`
	UpdatedAt       time.Time                                                              `json:"updated_at" format:"date-time"`
	JSON            accessApplicationNewResponseDeviceEnrollmentPermissionsApplicationJSON `json:"-"`
}

// accessApplicationNewResponseDeviceEnrollmentPermissionsApplicationJSON contains
// the JSON metadata for the struct
// [AccessApplicationNewResponseDeviceEnrollmentPermissionsApplication]
type accessApplicationNewResponseDeviceEnrollmentPermissionsApplicationJSON struct {
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

func (r *AccessApplicationNewResponseDeviceEnrollmentPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseDeviceEnrollmentPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationNewResponseDeviceEnrollmentPermissionsApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

// The application type.
type AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType string

const (
	AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeSelfHosted  AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "self_hosted"
	AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeSaas        AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "saas"
	AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeSSH         AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "ssh"
	AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeVnc         AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "vnc"
	AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeAppLauncher AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "app_launcher"
	AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeWARP        AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "warp"
	AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeBiso        AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "biso"
	AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeBookmark    AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "bookmark"
	AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationTypeDashSSO     AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationType = "dash_sso"
)

type AccessApplicationNewResponseBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type AccessApplicationNewResponseBrowserIsolationPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                 `json:"session_duration"`
	UpdatedAt       time.Time                                                              `json:"updated_at" format:"date-time"`
	JSON            accessApplicationNewResponseBrowserIsolationPermissionsApplicationJSON `json:"-"`
}

// accessApplicationNewResponseBrowserIsolationPermissionsApplicationJSON contains
// the JSON metadata for the struct
// [AccessApplicationNewResponseBrowserIsolationPermissionsApplication]
type accessApplicationNewResponseBrowserIsolationPermissionsApplicationJSON struct {
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

func (r *AccessApplicationNewResponseBrowserIsolationPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseBrowserIsolationPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationNewResponseBrowserIsolationPermissionsApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

// The application type.
type AccessApplicationNewResponseBrowserIsolationPermissionsApplicationType string

const (
	AccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeSelfHosted  AccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "self_hosted"
	AccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeSaas        AccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "saas"
	AccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeSSH         AccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "ssh"
	AccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeVnc         AccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "vnc"
	AccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeAppLauncher AccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "app_launcher"
	AccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeWARP        AccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "warp"
	AccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeBiso        AccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "biso"
	AccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeBookmark    AccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "bookmark"
	AccessApplicationNewResponseBrowserIsolationPermissionsApplicationTypeDashSSO     AccessApplicationNewResponseBrowserIsolationPermissionsApplicationType = "dash_sso"
)

type AccessApplicationNewResponseBookmarkApplication struct {
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
	Type      string                                              `json:"type"`
	UpdatedAt time.Time                                           `json:"updated_at" format:"date-time"`
	JSON      accessApplicationNewResponseBookmarkApplicationJSON `json:"-"`
}

// accessApplicationNewResponseBookmarkApplicationJSON contains the JSON metadata
// for the struct [AccessApplicationNewResponseBookmarkApplication]
type accessApplicationNewResponseBookmarkApplicationJSON struct {
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

func (r *AccessApplicationNewResponseBookmarkApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseBookmarkApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationNewResponseBookmarkApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

// Union satisfied by
// [zero_trust.AccessApplicationUpdateResponseSelfHostedApplication],
// [zero_trust.AccessApplicationUpdateResponseSaaSApplication],
// [zero_trust.AccessApplicationUpdateResponseBrowserSSHApplication],
// [zero_trust.AccessApplicationUpdateResponseBrowserVncApplication],
// [zero_trust.AccessApplicationUpdateResponseAppLauncherApplication],
// [zero_trust.AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplication],
// [zero_trust.AccessApplicationUpdateResponseBrowserIsolationPermissionsApplication]
// or [zero_trust.AccessApplicationUpdateResponseBookmarkApplication].
type AccessApplicationUpdateResponse interface {
	implementsZeroTrustAccessApplicationUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationUpdateResponseSelfHostedApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationUpdateResponseSaaSApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationUpdateResponseBrowserSSHApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationUpdateResponseBrowserVncApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationUpdateResponseAppLauncherApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationUpdateResponseBrowserIsolationPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationUpdateResponseBookmarkApplication{}),
		},
	)
}

type AccessApplicationUpdateResponseSelfHostedApplication struct {
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
	AutoRedirectToIdentity bool                                                            `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationUpdateResponseSelfHostedApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                       `json:"created_at" format:"date-time"`
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
	Tags      []string                                                 `json:"tags"`
	UpdatedAt time.Time                                                `json:"updated_at" format:"date-time"`
	JSON      accessApplicationUpdateResponseSelfHostedApplicationJSON `json:"-"`
}

// accessApplicationUpdateResponseSelfHostedApplicationJSON contains the JSON
// metadata for the struct [AccessApplicationUpdateResponseSelfHostedApplication]
type accessApplicationUpdateResponseSelfHostedApplicationJSON struct {
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

func (r *AccessApplicationUpdateResponseSelfHostedApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseSelfHostedApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationUpdateResponseSelfHostedApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

type AccessApplicationUpdateResponseSelfHostedApplicationCorsHeaders struct {
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
	AllowedMethods []AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                             `json:"max_age"`
	JSON   accessApplicationUpdateResponseSelfHostedApplicationCorsHeadersJSON `json:"-"`
}

// accessApplicationUpdateResponseSelfHostedApplicationCorsHeadersJSON contains the
// JSON metadata for the struct
// [AccessApplicationUpdateResponseSelfHostedApplicationCorsHeaders]
type accessApplicationUpdateResponseSelfHostedApplicationCorsHeadersJSON struct {
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

func (r *AccessApplicationUpdateResponseSelfHostedApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseSelfHostedApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodGet     AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodPost    AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodHead    AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodPut     AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodDelete  AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodConnect AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodOptions AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodTrace   AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethodPatch   AccessApplicationUpdateResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationUpdateResponseSaaSApplication struct {
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
	Name    string                                                `json:"name"`
	SaasApp AccessApplicationUpdateResponseSaaSApplicationSaasApp `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags []string `json:"tags"`
	// The application type.
	Type      string                                             `json:"type"`
	UpdatedAt time.Time                                          `json:"updated_at" format:"date-time"`
	JSON      accessApplicationUpdateResponseSaaSApplicationJSON `json:"-"`
}

// accessApplicationUpdateResponseSaaSApplicationJSON contains the JSON metadata
// for the struct [AccessApplicationUpdateResponseSaaSApplication]
type accessApplicationUpdateResponseSaaSApplicationJSON struct {
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

func (r *AccessApplicationUpdateResponseSaaSApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseSaaSApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationUpdateResponseSaaSApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

// Union satisfied by
// [zero_trust.AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasApp]
// or
// [zero_trust.AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasApp].
type AccessApplicationUpdateResponseSaaSApplicationSaasApp interface {
	implementsZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasApp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationUpdateResponseSaaSApplicationSaasApp)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasApp{}),
		},
	)
}

type AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string                                                                                 `json:"consumer_service_url"`
	CreatedAt          time.Time                                                                              `json:"created_at" format:"date-time"`
	CustomAttributes   AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat `json:"name_id_format"`
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
	SSOEndpoint string                                                                     `json:"sso_endpoint"`
	UpdatedAt   time.Time                                                                  `json:"updated_at" format:"date-time"`
	JSON        accessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON `json:"-"`
}

// accessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasApp]
type accessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON struct {
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

func (r *AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasApp) implementsZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType string

const (
	AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeSaml AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "saml"
	AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeOidc AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "oidc"
)

type AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name string `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat `json:"name_format"`
	Source     AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource     `json:"source"`
	JSON       accessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON       `json:"-"`
}

// accessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes]
type accessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON struct {
	Name        apijson.Field
	NameFormat  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON) RawJSON() string {
	return r.raw
}

// A globally unique name for an identity or service provider.
type AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name string                                                                                           `json:"name"`
	JSON accessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON `json:"-"`
}

// accessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource]
type accessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON) RawJSON() string {
	return r.raw
}

// The format of the name identifier sent to the SaaS application.
type AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat string

const (
	AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatID    AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "id"
	AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatEmail AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The OIDC flows supported by this application
	GrantTypes []AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs []string `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes    []AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScope `json:"scopes"`
	UpdatedAt time.Time                                                                     `json:"updated_at" format:"date-time"`
	JSON      accessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON    `json:"-"`
}

// accessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasApp]
type accessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON struct {
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

func (r *AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType string

const (
	AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "saml"
	AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "oidc"
)

type AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType string

const (
	AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScope string

const (
	AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid  AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "openid"
	AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups  AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "groups"
	AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail   AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "email"
	AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile AccessApplicationUpdateResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "profile"
)

type AccessApplicationUpdateResponseBrowserSSHApplication struct {
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
	AutoRedirectToIdentity bool                                                            `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                       `json:"created_at" format:"date-time"`
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
	Tags      []string                                                 `json:"tags"`
	UpdatedAt time.Time                                                `json:"updated_at" format:"date-time"`
	JSON      accessApplicationUpdateResponseBrowserSSHApplicationJSON `json:"-"`
}

// accessApplicationUpdateResponseBrowserSSHApplicationJSON contains the JSON
// metadata for the struct [AccessApplicationUpdateResponseBrowserSSHApplication]
type accessApplicationUpdateResponseBrowserSSHApplicationJSON struct {
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

func (r *AccessApplicationUpdateResponseBrowserSSHApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseBrowserSSHApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationUpdateResponseBrowserSSHApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

type AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeaders struct {
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
	AllowedMethods []AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                             `json:"max_age"`
	JSON   accessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersJSON `json:"-"`
}

// accessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersJSON contains the
// JSON metadata for the struct
// [AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeaders]
type accessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersJSON struct {
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

func (r *AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodGet     AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodPost    AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodHead    AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodPut     AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodDelete  AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodConnect AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodOptions AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodTrace   AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethodPatch   AccessApplicationUpdateResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationUpdateResponseBrowserVncApplication struct {
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
	AutoRedirectToIdentity bool                                                            `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationUpdateResponseBrowserVncApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                       `json:"created_at" format:"date-time"`
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
	Tags      []string                                                 `json:"tags"`
	UpdatedAt time.Time                                                `json:"updated_at" format:"date-time"`
	JSON      accessApplicationUpdateResponseBrowserVncApplicationJSON `json:"-"`
}

// accessApplicationUpdateResponseBrowserVncApplicationJSON contains the JSON
// metadata for the struct [AccessApplicationUpdateResponseBrowserVncApplication]
type accessApplicationUpdateResponseBrowserVncApplicationJSON struct {
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

func (r *AccessApplicationUpdateResponseBrowserVncApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseBrowserVncApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationUpdateResponseBrowserVncApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

type AccessApplicationUpdateResponseBrowserVncApplicationCorsHeaders struct {
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
	AllowedMethods []AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                             `json:"max_age"`
	JSON   accessApplicationUpdateResponseBrowserVncApplicationCorsHeadersJSON `json:"-"`
}

// accessApplicationUpdateResponseBrowserVncApplicationCorsHeadersJSON contains the
// JSON metadata for the struct
// [AccessApplicationUpdateResponseBrowserVncApplicationCorsHeaders]
type accessApplicationUpdateResponseBrowserVncApplicationCorsHeadersJSON struct {
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

func (r *AccessApplicationUpdateResponseBrowserVncApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseBrowserVncApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodGet     AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodPost    AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodHead    AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodPut     AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodDelete  AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodConnect AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodOptions AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodTrace   AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethodPatch   AccessApplicationUpdateResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationUpdateResponseAppLauncherApplication struct {
	// The application type.
	Type AccessApplicationUpdateResponseAppLauncherApplicationType `json:"type,required"`
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
	SessionDuration string                                                    `json:"session_duration"`
	UpdatedAt       time.Time                                                 `json:"updated_at" format:"date-time"`
	JSON            accessApplicationUpdateResponseAppLauncherApplicationJSON `json:"-"`
}

// accessApplicationUpdateResponseAppLauncherApplicationJSON contains the JSON
// metadata for the struct [AccessApplicationUpdateResponseAppLauncherApplication]
type accessApplicationUpdateResponseAppLauncherApplicationJSON struct {
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

func (r *AccessApplicationUpdateResponseAppLauncherApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseAppLauncherApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationUpdateResponseAppLauncherApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

// The application type.
type AccessApplicationUpdateResponseAppLauncherApplicationType string

const (
	AccessApplicationUpdateResponseAppLauncherApplicationTypeSelfHosted  AccessApplicationUpdateResponseAppLauncherApplicationType = "self_hosted"
	AccessApplicationUpdateResponseAppLauncherApplicationTypeSaas        AccessApplicationUpdateResponseAppLauncherApplicationType = "saas"
	AccessApplicationUpdateResponseAppLauncherApplicationTypeSSH         AccessApplicationUpdateResponseAppLauncherApplicationType = "ssh"
	AccessApplicationUpdateResponseAppLauncherApplicationTypeVnc         AccessApplicationUpdateResponseAppLauncherApplicationType = "vnc"
	AccessApplicationUpdateResponseAppLauncherApplicationTypeAppLauncher AccessApplicationUpdateResponseAppLauncherApplicationType = "app_launcher"
	AccessApplicationUpdateResponseAppLauncherApplicationTypeWARP        AccessApplicationUpdateResponseAppLauncherApplicationType = "warp"
	AccessApplicationUpdateResponseAppLauncherApplicationTypeBiso        AccessApplicationUpdateResponseAppLauncherApplicationType = "biso"
	AccessApplicationUpdateResponseAppLauncherApplicationTypeBookmark    AccessApplicationUpdateResponseAppLauncherApplicationType = "bookmark"
	AccessApplicationUpdateResponseAppLauncherApplicationTypeDashSSO     AccessApplicationUpdateResponseAppLauncherApplicationType = "dash_sso"
)

type AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                    `json:"session_duration"`
	UpdatedAt       time.Time                                                                 `json:"updated_at" format:"date-time"`
	JSON            accessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationJSON `json:"-"`
}

// accessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplication]
type accessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationJSON struct {
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

func (r *AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

// The application type.
type AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType string

const (
	AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeSelfHosted  AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "self_hosted"
	AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeSaas        AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "saas"
	AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeSSH         AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "ssh"
	AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeVnc         AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "vnc"
	AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeAppLauncher AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "app_launcher"
	AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeWARP        AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "warp"
	AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeBiso        AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "biso"
	AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeBookmark    AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "bookmark"
	AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationTypeDashSSO     AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationType = "dash_sso"
)

type AccessApplicationUpdateResponseBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                    `json:"session_duration"`
	UpdatedAt       time.Time                                                                 `json:"updated_at" format:"date-time"`
	JSON            accessApplicationUpdateResponseBrowserIsolationPermissionsApplicationJSON `json:"-"`
}

// accessApplicationUpdateResponseBrowserIsolationPermissionsApplicationJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseBrowserIsolationPermissionsApplication]
type accessApplicationUpdateResponseBrowserIsolationPermissionsApplicationJSON struct {
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

func (r *AccessApplicationUpdateResponseBrowserIsolationPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseBrowserIsolationPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationUpdateResponseBrowserIsolationPermissionsApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

// The application type.
type AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType string

const (
	AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeSelfHosted  AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "self_hosted"
	AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeSaas        AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "saas"
	AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeSSH         AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "ssh"
	AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeVnc         AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "vnc"
	AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeAppLauncher AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "app_launcher"
	AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeWARP        AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "warp"
	AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeBiso        AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "biso"
	AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeBookmark    AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "bookmark"
	AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationTypeDashSSO     AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationType = "dash_sso"
)

type AccessApplicationUpdateResponseBookmarkApplication struct {
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
	Type      string                                                 `json:"type"`
	UpdatedAt time.Time                                              `json:"updated_at" format:"date-time"`
	JSON      accessApplicationUpdateResponseBookmarkApplicationJSON `json:"-"`
}

// accessApplicationUpdateResponseBookmarkApplicationJSON contains the JSON
// metadata for the struct [AccessApplicationUpdateResponseBookmarkApplication]
type accessApplicationUpdateResponseBookmarkApplicationJSON struct {
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

func (r *AccessApplicationUpdateResponseBookmarkApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseBookmarkApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationUpdateResponseBookmarkApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

// Union satisfied by
// [zero_trust.AccessApplicationListResponseSelfHostedApplication],
// [zero_trust.AccessApplicationListResponseSaaSApplication],
// [zero_trust.AccessApplicationListResponseBrowserSSHApplication],
// [zero_trust.AccessApplicationListResponseBrowserVncApplication],
// [zero_trust.AccessApplicationListResponseAppLauncherApplication],
// [zero_trust.AccessApplicationListResponseDeviceEnrollmentPermissionsApplication],
// [zero_trust.AccessApplicationListResponseBrowserIsolationPermissionsApplication]
// or [zero_trust.AccessApplicationListResponseBookmarkApplication].
type AccessApplicationListResponse interface {
	implementsZeroTrustAccessApplicationListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationListResponseSelfHostedApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationListResponseSaaSApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationListResponseBrowserSSHApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationListResponseBrowserVncApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationListResponseAppLauncherApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationListResponseDeviceEnrollmentPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationListResponseBrowserIsolationPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationListResponseBookmarkApplication{}),
		},
	)
}

type AccessApplicationListResponseSelfHostedApplication struct {
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
	AutoRedirectToIdentity bool                                                          `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationListResponseSelfHostedApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                     `json:"created_at" format:"date-time"`
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
	Tags      []string                                               `json:"tags"`
	UpdatedAt time.Time                                              `json:"updated_at" format:"date-time"`
	JSON      accessApplicationListResponseSelfHostedApplicationJSON `json:"-"`
}

// accessApplicationListResponseSelfHostedApplicationJSON contains the JSON
// metadata for the struct [AccessApplicationListResponseSelfHostedApplication]
type accessApplicationListResponseSelfHostedApplicationJSON struct {
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

func (r *AccessApplicationListResponseSelfHostedApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseSelfHostedApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationListResponseSelfHostedApplication) implementsZeroTrustAccessApplicationListResponse() {
}

type AccessApplicationListResponseSelfHostedApplicationCorsHeaders struct {
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
	AllowedMethods []AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                           `json:"max_age"`
	JSON   accessApplicationListResponseSelfHostedApplicationCorsHeadersJSON `json:"-"`
}

// accessApplicationListResponseSelfHostedApplicationCorsHeadersJSON contains the
// JSON metadata for the struct
// [AccessApplicationListResponseSelfHostedApplicationCorsHeaders]
type accessApplicationListResponseSelfHostedApplicationCorsHeadersJSON struct {
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

func (r *AccessApplicationListResponseSelfHostedApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseSelfHostedApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodGet     AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodPost    AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodHead    AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodPut     AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodDelete  AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodConnect AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodOptions AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodTrace   AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethodPatch   AccessApplicationListResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationListResponseSaaSApplication struct {
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
	Name    string                                              `json:"name"`
	SaasApp AccessApplicationListResponseSaaSApplicationSaasApp `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags []string `json:"tags"`
	// The application type.
	Type      string                                           `json:"type"`
	UpdatedAt time.Time                                        `json:"updated_at" format:"date-time"`
	JSON      accessApplicationListResponseSaaSApplicationJSON `json:"-"`
}

// accessApplicationListResponseSaaSApplicationJSON contains the JSON metadata for
// the struct [AccessApplicationListResponseSaaSApplication]
type accessApplicationListResponseSaaSApplicationJSON struct {
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

func (r *AccessApplicationListResponseSaaSApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseSaaSApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationListResponseSaaSApplication) implementsZeroTrustAccessApplicationListResponse() {
}

// Union satisfied by
// [zero_trust.AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasApp]
// or
// [zero_trust.AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasApp].
type AccessApplicationListResponseSaaSApplicationSaasApp interface {
	implementsZeroTrustAccessApplicationListResponseSaaSApplicationSaasApp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationListResponseSaaSApplicationSaasApp)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasApp{}),
		},
	)
}

type AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string                                                                               `json:"consumer_service_url"`
	CreatedAt          time.Time                                                                            `json:"created_at" format:"date-time"`
	CustomAttributes   AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat `json:"name_id_format"`
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
	SSOEndpoint string                                                                   `json:"sso_endpoint"`
	UpdatedAt   time.Time                                                                `json:"updated_at" format:"date-time"`
	JSON        accessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON `json:"-"`
}

// accessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON
// contains the JSON metadata for the struct
// [AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasApp]
type accessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON struct {
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

func (r *AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasApp) implementsZeroTrustAccessApplicationListResponseSaaSApplicationSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType string

const (
	AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeSaml AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "saml"
	AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeOidc AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "oidc"
)

type AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name string `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat `json:"name_format"`
	Source     AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource     `json:"source"`
	JSON       accessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON       `json:"-"`
}

// accessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON
// contains the JSON metadata for the struct
// [AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes]
type accessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON struct {
	Name        apijson.Field
	NameFormat  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON) RawJSON() string {
	return r.raw
}

// A globally unique name for an identity or service provider.
type AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name string                                                                                         `json:"name"`
	JSON accessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON `json:"-"`
}

// accessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON
// contains the JSON metadata for the struct
// [AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource]
type accessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON) RawJSON() string {
	return r.raw
}

// The format of the name identifier sent to the SaaS application.
type AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat string

const (
	AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatID    AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "id"
	AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatEmail AccessApplicationListResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The OIDC flows supported by this application
	GrantTypes []AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs []string `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes    []AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScope `json:"scopes"`
	UpdatedAt time.Time                                                                   `json:"updated_at" format:"date-time"`
	JSON      accessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON    `json:"-"`
}

// accessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON
// contains the JSON metadata for the struct
// [AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasApp]
type accessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON struct {
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

func (r *AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationListResponseSaaSApplicationSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType string

const (
	AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "saml"
	AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "oidc"
)

type AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType string

const (
	AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScope string

const (
	AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid  AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "openid"
	AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups  AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "groups"
	AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail   AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "email"
	AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile AccessApplicationListResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "profile"
)

type AccessApplicationListResponseBrowserSSHApplication struct {
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
	AutoRedirectToIdentity bool                                                          `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationListResponseBrowserSSHApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                     `json:"created_at" format:"date-time"`
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
	Tags      []string                                               `json:"tags"`
	UpdatedAt time.Time                                              `json:"updated_at" format:"date-time"`
	JSON      accessApplicationListResponseBrowserSSHApplicationJSON `json:"-"`
}

// accessApplicationListResponseBrowserSSHApplicationJSON contains the JSON
// metadata for the struct [AccessApplicationListResponseBrowserSSHApplication]
type accessApplicationListResponseBrowserSSHApplicationJSON struct {
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

func (r *AccessApplicationListResponseBrowserSSHApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseBrowserSSHApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationListResponseBrowserSSHApplication) implementsZeroTrustAccessApplicationListResponse() {
}

type AccessApplicationListResponseBrowserSSHApplicationCorsHeaders struct {
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
	AllowedMethods []AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                           `json:"max_age"`
	JSON   accessApplicationListResponseBrowserSSHApplicationCorsHeadersJSON `json:"-"`
}

// accessApplicationListResponseBrowserSSHApplicationCorsHeadersJSON contains the
// JSON metadata for the struct
// [AccessApplicationListResponseBrowserSSHApplicationCorsHeaders]
type accessApplicationListResponseBrowserSSHApplicationCorsHeadersJSON struct {
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

func (r *AccessApplicationListResponseBrowserSSHApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseBrowserSSHApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodGet     AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodPost    AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodHead    AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodPut     AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodDelete  AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodConnect AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodOptions AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodTrace   AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethodPatch   AccessApplicationListResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationListResponseBrowserVncApplication struct {
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
	AutoRedirectToIdentity bool                                                          `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationListResponseBrowserVncApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                     `json:"created_at" format:"date-time"`
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
	Tags      []string                                               `json:"tags"`
	UpdatedAt time.Time                                              `json:"updated_at" format:"date-time"`
	JSON      accessApplicationListResponseBrowserVncApplicationJSON `json:"-"`
}

// accessApplicationListResponseBrowserVncApplicationJSON contains the JSON
// metadata for the struct [AccessApplicationListResponseBrowserVncApplication]
type accessApplicationListResponseBrowserVncApplicationJSON struct {
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

func (r *AccessApplicationListResponseBrowserVncApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseBrowserVncApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationListResponseBrowserVncApplication) implementsZeroTrustAccessApplicationListResponse() {
}

type AccessApplicationListResponseBrowserVncApplicationCorsHeaders struct {
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
	AllowedMethods []AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                           `json:"max_age"`
	JSON   accessApplicationListResponseBrowserVncApplicationCorsHeadersJSON `json:"-"`
}

// accessApplicationListResponseBrowserVncApplicationCorsHeadersJSON contains the
// JSON metadata for the struct
// [AccessApplicationListResponseBrowserVncApplicationCorsHeaders]
type accessApplicationListResponseBrowserVncApplicationCorsHeadersJSON struct {
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

func (r *AccessApplicationListResponseBrowserVncApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseBrowserVncApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodGet     AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodPost    AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodHead    AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodPut     AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodDelete  AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodConnect AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodOptions AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodTrace   AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethodPatch   AccessApplicationListResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationListResponseAppLauncherApplication struct {
	// The application type.
	Type AccessApplicationListResponseAppLauncherApplicationType `json:"type,required"`
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
	SessionDuration string                                                  `json:"session_duration"`
	UpdatedAt       time.Time                                               `json:"updated_at" format:"date-time"`
	JSON            accessApplicationListResponseAppLauncherApplicationJSON `json:"-"`
}

// accessApplicationListResponseAppLauncherApplicationJSON contains the JSON
// metadata for the struct [AccessApplicationListResponseAppLauncherApplication]
type accessApplicationListResponseAppLauncherApplicationJSON struct {
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

func (r *AccessApplicationListResponseAppLauncherApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseAppLauncherApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationListResponseAppLauncherApplication) implementsZeroTrustAccessApplicationListResponse() {
}

// The application type.
type AccessApplicationListResponseAppLauncherApplicationType string

const (
	AccessApplicationListResponseAppLauncherApplicationTypeSelfHosted  AccessApplicationListResponseAppLauncherApplicationType = "self_hosted"
	AccessApplicationListResponseAppLauncherApplicationTypeSaas        AccessApplicationListResponseAppLauncherApplicationType = "saas"
	AccessApplicationListResponseAppLauncherApplicationTypeSSH         AccessApplicationListResponseAppLauncherApplicationType = "ssh"
	AccessApplicationListResponseAppLauncherApplicationTypeVnc         AccessApplicationListResponseAppLauncherApplicationType = "vnc"
	AccessApplicationListResponseAppLauncherApplicationTypeAppLauncher AccessApplicationListResponseAppLauncherApplicationType = "app_launcher"
	AccessApplicationListResponseAppLauncherApplicationTypeWARP        AccessApplicationListResponseAppLauncherApplicationType = "warp"
	AccessApplicationListResponseAppLauncherApplicationTypeBiso        AccessApplicationListResponseAppLauncherApplicationType = "biso"
	AccessApplicationListResponseAppLauncherApplicationTypeBookmark    AccessApplicationListResponseAppLauncherApplicationType = "bookmark"
	AccessApplicationListResponseAppLauncherApplicationTypeDashSSO     AccessApplicationListResponseAppLauncherApplicationType = "dash_sso"
)

type AccessApplicationListResponseDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                  `json:"session_duration"`
	UpdatedAt       time.Time                                                               `json:"updated_at" format:"date-time"`
	JSON            accessApplicationListResponseDeviceEnrollmentPermissionsApplicationJSON `json:"-"`
}

// accessApplicationListResponseDeviceEnrollmentPermissionsApplicationJSON contains
// the JSON metadata for the struct
// [AccessApplicationListResponseDeviceEnrollmentPermissionsApplication]
type accessApplicationListResponseDeviceEnrollmentPermissionsApplicationJSON struct {
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

func (r *AccessApplicationListResponseDeviceEnrollmentPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseDeviceEnrollmentPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationListResponseDeviceEnrollmentPermissionsApplication) implementsZeroTrustAccessApplicationListResponse() {
}

// The application type.
type AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType string

const (
	AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeSelfHosted  AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "self_hosted"
	AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeSaas        AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "saas"
	AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeSSH         AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "ssh"
	AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeVnc         AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "vnc"
	AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeAppLauncher AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "app_launcher"
	AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeWARP        AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "warp"
	AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeBiso        AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "biso"
	AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeBookmark    AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "bookmark"
	AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationTypeDashSSO     AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationType = "dash_sso"
)

type AccessApplicationListResponseBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type AccessApplicationListResponseBrowserIsolationPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                  `json:"session_duration"`
	UpdatedAt       time.Time                                                               `json:"updated_at" format:"date-time"`
	JSON            accessApplicationListResponseBrowserIsolationPermissionsApplicationJSON `json:"-"`
}

// accessApplicationListResponseBrowserIsolationPermissionsApplicationJSON contains
// the JSON metadata for the struct
// [AccessApplicationListResponseBrowserIsolationPermissionsApplication]
type accessApplicationListResponseBrowserIsolationPermissionsApplicationJSON struct {
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

func (r *AccessApplicationListResponseBrowserIsolationPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseBrowserIsolationPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationListResponseBrowserIsolationPermissionsApplication) implementsZeroTrustAccessApplicationListResponse() {
}

// The application type.
type AccessApplicationListResponseBrowserIsolationPermissionsApplicationType string

const (
	AccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeSelfHosted  AccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "self_hosted"
	AccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeSaas        AccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "saas"
	AccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeSSH         AccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "ssh"
	AccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeVnc         AccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "vnc"
	AccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeAppLauncher AccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "app_launcher"
	AccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeWARP        AccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "warp"
	AccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeBiso        AccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "biso"
	AccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeBookmark    AccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "bookmark"
	AccessApplicationListResponseBrowserIsolationPermissionsApplicationTypeDashSSO     AccessApplicationListResponseBrowserIsolationPermissionsApplicationType = "dash_sso"
)

type AccessApplicationListResponseBookmarkApplication struct {
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
	Type      string                                               `json:"type"`
	UpdatedAt time.Time                                            `json:"updated_at" format:"date-time"`
	JSON      accessApplicationListResponseBookmarkApplicationJSON `json:"-"`
}

// accessApplicationListResponseBookmarkApplicationJSON contains the JSON metadata
// for the struct [AccessApplicationListResponseBookmarkApplication]
type accessApplicationListResponseBookmarkApplicationJSON struct {
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

func (r *AccessApplicationListResponseBookmarkApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseBookmarkApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationListResponseBookmarkApplication) implementsZeroTrustAccessApplicationListResponse() {
}

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

// Union satisfied by
// [zero_trust.AccessApplicationGetResponseSelfHostedApplication],
// [zero_trust.AccessApplicationGetResponseSaaSApplication],
// [zero_trust.AccessApplicationGetResponseBrowserSSHApplication],
// [zero_trust.AccessApplicationGetResponseBrowserVncApplication],
// [zero_trust.AccessApplicationGetResponseAppLauncherApplication],
// [zero_trust.AccessApplicationGetResponseDeviceEnrollmentPermissionsApplication],
// [zero_trust.AccessApplicationGetResponseBrowserIsolationPermissionsApplication]
// or [zero_trust.AccessApplicationGetResponseBookmarkApplication].
type AccessApplicationGetResponse interface {
	implementsZeroTrustAccessApplicationGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationGetResponseSelfHostedApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationGetResponseSaaSApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationGetResponseBrowserSSHApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationGetResponseBrowserVncApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationGetResponseAppLauncherApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationGetResponseDeviceEnrollmentPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationGetResponseBrowserIsolationPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationGetResponseBookmarkApplication{}),
		},
	)
}

type AccessApplicationGetResponseSelfHostedApplication struct {
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
	AutoRedirectToIdentity bool                                                         `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationGetResponseSelfHostedApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                    `json:"created_at" format:"date-time"`
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
	Tags      []string                                              `json:"tags"`
	UpdatedAt time.Time                                             `json:"updated_at" format:"date-time"`
	JSON      accessApplicationGetResponseSelfHostedApplicationJSON `json:"-"`
}

// accessApplicationGetResponseSelfHostedApplicationJSON contains the JSON metadata
// for the struct [AccessApplicationGetResponseSelfHostedApplication]
type accessApplicationGetResponseSelfHostedApplicationJSON struct {
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

func (r *AccessApplicationGetResponseSelfHostedApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseSelfHostedApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationGetResponseSelfHostedApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

type AccessApplicationGetResponseSelfHostedApplicationCorsHeaders struct {
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
	AllowedMethods []AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                          `json:"max_age"`
	JSON   accessApplicationGetResponseSelfHostedApplicationCorsHeadersJSON `json:"-"`
}

// accessApplicationGetResponseSelfHostedApplicationCorsHeadersJSON contains the
// JSON metadata for the struct
// [AccessApplicationGetResponseSelfHostedApplicationCorsHeaders]
type accessApplicationGetResponseSelfHostedApplicationCorsHeadersJSON struct {
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

func (r *AccessApplicationGetResponseSelfHostedApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseSelfHostedApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodGet     AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodPost    AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodHead    AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodPut     AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodDelete  AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodConnect AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodOptions AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodTrace   AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethodPatch   AccessApplicationGetResponseSelfHostedApplicationCorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationGetResponseSaaSApplication struct {
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
	Name    string                                             `json:"name"`
	SaasApp AccessApplicationGetResponseSaaSApplicationSaasApp `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags []string `json:"tags"`
	// The application type.
	Type      string                                          `json:"type"`
	UpdatedAt time.Time                                       `json:"updated_at" format:"date-time"`
	JSON      accessApplicationGetResponseSaaSApplicationJSON `json:"-"`
}

// accessApplicationGetResponseSaaSApplicationJSON contains the JSON metadata for
// the struct [AccessApplicationGetResponseSaaSApplication]
type accessApplicationGetResponseSaaSApplicationJSON struct {
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

func (r *AccessApplicationGetResponseSaaSApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseSaaSApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationGetResponseSaaSApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

// Union satisfied by
// [zero_trust.AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasApp]
// or
// [zero_trust.AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasApp].
type AccessApplicationGetResponseSaaSApplicationSaasApp interface {
	implementsZeroTrustAccessApplicationGetResponseSaaSApplicationSaasApp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationGetResponseSaaSApplicationSaasApp)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasApp{}),
		},
	)
}

type AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string                                                                              `json:"consumer_service_url"`
	CreatedAt          time.Time                                                                           `json:"created_at" format:"date-time"`
	CustomAttributes   AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat `json:"name_id_format"`
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
	SSOEndpoint string                                                                  `json:"sso_endpoint"`
	UpdatedAt   time.Time                                                               `json:"updated_at" format:"date-time"`
	JSON        accessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON `json:"-"`
}

// accessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON contains
// the JSON metadata for the struct
// [AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasApp]
type accessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON struct {
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

func (r *AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasApp) implementsZeroTrustAccessApplicationGetResponseSaaSApplicationSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType string

const (
	AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeSaml AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "saml"
	AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeOidc AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "oidc"
)

type AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name string `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat `json:"name_format"`
	Source     AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource     `json:"source"`
	JSON       accessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON       `json:"-"`
}

// accessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON
// contains the JSON metadata for the struct
// [AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes]
type accessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON struct {
	Name        apijson.Field
	NameFormat  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON) RawJSON() string {
	return r.raw
}

// A globally unique name for an identity or service provider.
type AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name string                                                                                        `json:"name"`
	JSON accessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON `json:"-"`
}

// accessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON
// contains the JSON metadata for the struct
// [AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource]
type accessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON) RawJSON() string {
	return r.raw
}

// The format of the name identifier sent to the SaaS application.
type AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat string

const (
	AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatID    AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "id"
	AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatEmail AccessApplicationGetResponseSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The OIDC flows supported by this application
	GrantTypes []AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs []string `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes    []AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScope `json:"scopes"`
	UpdatedAt time.Time                                                                  `json:"updated_at" format:"date-time"`
	JSON      accessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON    `json:"-"`
}

// accessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON contains
// the JSON metadata for the struct
// [AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasApp]
type accessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON struct {
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

func (r *AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationGetResponseSaaSApplicationSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType string

const (
	AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "saml"
	AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "oidc"
)

type AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType string

const (
	AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScope string

const (
	AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid  AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "openid"
	AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups  AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "groups"
	AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail   AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "email"
	AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile AccessApplicationGetResponseSaaSApplicationSaasAppAccessOidcSaasAppScope = "profile"
)

type AccessApplicationGetResponseBrowserSSHApplication struct {
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
	AutoRedirectToIdentity bool                                                         `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationGetResponseBrowserSSHApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                    `json:"created_at" format:"date-time"`
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
	Tags      []string                                              `json:"tags"`
	UpdatedAt time.Time                                             `json:"updated_at" format:"date-time"`
	JSON      accessApplicationGetResponseBrowserSSHApplicationJSON `json:"-"`
}

// accessApplicationGetResponseBrowserSSHApplicationJSON contains the JSON metadata
// for the struct [AccessApplicationGetResponseBrowserSSHApplication]
type accessApplicationGetResponseBrowserSSHApplicationJSON struct {
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

func (r *AccessApplicationGetResponseBrowserSSHApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseBrowserSSHApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationGetResponseBrowserSSHApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

type AccessApplicationGetResponseBrowserSSHApplicationCorsHeaders struct {
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
	AllowedMethods []AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                          `json:"max_age"`
	JSON   accessApplicationGetResponseBrowserSSHApplicationCorsHeadersJSON `json:"-"`
}

// accessApplicationGetResponseBrowserSSHApplicationCorsHeadersJSON contains the
// JSON metadata for the struct
// [AccessApplicationGetResponseBrowserSSHApplicationCorsHeaders]
type accessApplicationGetResponseBrowserSSHApplicationCorsHeadersJSON struct {
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

func (r *AccessApplicationGetResponseBrowserSSHApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseBrowserSSHApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodGet     AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodPost    AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodHead    AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodPut     AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodDelete  AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodConnect AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodOptions AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodTrace   AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethodPatch   AccessApplicationGetResponseBrowserSSHApplicationCorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationGetResponseBrowserVncApplication struct {
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
	AutoRedirectToIdentity bool                                                         `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationGetResponseBrowserVncApplicationCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                    `json:"created_at" format:"date-time"`
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
	Tags      []string                                              `json:"tags"`
	UpdatedAt time.Time                                             `json:"updated_at" format:"date-time"`
	JSON      accessApplicationGetResponseBrowserVncApplicationJSON `json:"-"`
}

// accessApplicationGetResponseBrowserVncApplicationJSON contains the JSON metadata
// for the struct [AccessApplicationGetResponseBrowserVncApplication]
type accessApplicationGetResponseBrowserVncApplicationJSON struct {
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

func (r *AccessApplicationGetResponseBrowserVncApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseBrowserVncApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationGetResponseBrowserVncApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

type AccessApplicationGetResponseBrowserVncApplicationCorsHeaders struct {
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
	AllowedMethods []AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                          `json:"max_age"`
	JSON   accessApplicationGetResponseBrowserVncApplicationCorsHeadersJSON `json:"-"`
}

// accessApplicationGetResponseBrowserVncApplicationCorsHeadersJSON contains the
// JSON metadata for the struct
// [AccessApplicationGetResponseBrowserVncApplicationCorsHeaders]
type accessApplicationGetResponseBrowserVncApplicationCorsHeadersJSON struct {
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

func (r *AccessApplicationGetResponseBrowserVncApplicationCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseBrowserVncApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodGet     AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodPost    AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodHead    AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodPut     AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodDelete  AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodConnect AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodOptions AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodTrace   AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethodPatch   AccessApplicationGetResponseBrowserVncApplicationCorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationGetResponseAppLauncherApplication struct {
	// The application type.
	Type AccessApplicationGetResponseAppLauncherApplicationType `json:"type,required"`
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
	SessionDuration string                                                 `json:"session_duration"`
	UpdatedAt       time.Time                                              `json:"updated_at" format:"date-time"`
	JSON            accessApplicationGetResponseAppLauncherApplicationJSON `json:"-"`
}

// accessApplicationGetResponseAppLauncherApplicationJSON contains the JSON
// metadata for the struct [AccessApplicationGetResponseAppLauncherApplication]
type accessApplicationGetResponseAppLauncherApplicationJSON struct {
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

func (r *AccessApplicationGetResponseAppLauncherApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseAppLauncherApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationGetResponseAppLauncherApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

// The application type.
type AccessApplicationGetResponseAppLauncherApplicationType string

const (
	AccessApplicationGetResponseAppLauncherApplicationTypeSelfHosted  AccessApplicationGetResponseAppLauncherApplicationType = "self_hosted"
	AccessApplicationGetResponseAppLauncherApplicationTypeSaas        AccessApplicationGetResponseAppLauncherApplicationType = "saas"
	AccessApplicationGetResponseAppLauncherApplicationTypeSSH         AccessApplicationGetResponseAppLauncherApplicationType = "ssh"
	AccessApplicationGetResponseAppLauncherApplicationTypeVnc         AccessApplicationGetResponseAppLauncherApplicationType = "vnc"
	AccessApplicationGetResponseAppLauncherApplicationTypeAppLauncher AccessApplicationGetResponseAppLauncherApplicationType = "app_launcher"
	AccessApplicationGetResponseAppLauncherApplicationTypeWARP        AccessApplicationGetResponseAppLauncherApplicationType = "warp"
	AccessApplicationGetResponseAppLauncherApplicationTypeBiso        AccessApplicationGetResponseAppLauncherApplicationType = "biso"
	AccessApplicationGetResponseAppLauncherApplicationTypeBookmark    AccessApplicationGetResponseAppLauncherApplicationType = "bookmark"
	AccessApplicationGetResponseAppLauncherApplicationTypeDashSSO     AccessApplicationGetResponseAppLauncherApplicationType = "dash_sso"
)

type AccessApplicationGetResponseDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                 `json:"session_duration"`
	UpdatedAt       time.Time                                                              `json:"updated_at" format:"date-time"`
	JSON            accessApplicationGetResponseDeviceEnrollmentPermissionsApplicationJSON `json:"-"`
}

// accessApplicationGetResponseDeviceEnrollmentPermissionsApplicationJSON contains
// the JSON metadata for the struct
// [AccessApplicationGetResponseDeviceEnrollmentPermissionsApplication]
type accessApplicationGetResponseDeviceEnrollmentPermissionsApplicationJSON struct {
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

func (r *AccessApplicationGetResponseDeviceEnrollmentPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseDeviceEnrollmentPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationGetResponseDeviceEnrollmentPermissionsApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

// The application type.
type AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType string

const (
	AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeSelfHosted  AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "self_hosted"
	AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeSaas        AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "saas"
	AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeSSH         AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "ssh"
	AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeVnc         AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "vnc"
	AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeAppLauncher AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "app_launcher"
	AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeWARP        AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "warp"
	AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeBiso        AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "biso"
	AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeBookmark    AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "bookmark"
	AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationTypeDashSSO     AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationType = "dash_sso"
)

type AccessApplicationGetResponseBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type AccessApplicationGetResponseBrowserIsolationPermissionsApplicationType `json:"type,required"`
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
	SessionDuration string                                                                 `json:"session_duration"`
	UpdatedAt       time.Time                                                              `json:"updated_at" format:"date-time"`
	JSON            accessApplicationGetResponseBrowserIsolationPermissionsApplicationJSON `json:"-"`
}

// accessApplicationGetResponseBrowserIsolationPermissionsApplicationJSON contains
// the JSON metadata for the struct
// [AccessApplicationGetResponseBrowserIsolationPermissionsApplication]
type accessApplicationGetResponseBrowserIsolationPermissionsApplicationJSON struct {
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

func (r *AccessApplicationGetResponseBrowserIsolationPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseBrowserIsolationPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationGetResponseBrowserIsolationPermissionsApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

// The application type.
type AccessApplicationGetResponseBrowserIsolationPermissionsApplicationType string

const (
	AccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeSelfHosted  AccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "self_hosted"
	AccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeSaas        AccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "saas"
	AccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeSSH         AccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "ssh"
	AccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeVnc         AccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "vnc"
	AccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeAppLauncher AccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "app_launcher"
	AccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeWARP        AccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "warp"
	AccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeBiso        AccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "biso"
	AccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeBookmark    AccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "bookmark"
	AccessApplicationGetResponseBrowserIsolationPermissionsApplicationTypeDashSSO     AccessApplicationGetResponseBrowserIsolationPermissionsApplicationType = "dash_sso"
)

type AccessApplicationGetResponseBookmarkApplication struct {
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
	Type      string                                              `json:"type"`
	UpdatedAt time.Time                                           `json:"updated_at" format:"date-time"`
	JSON      accessApplicationGetResponseBookmarkApplicationJSON `json:"-"`
}

// accessApplicationGetResponseBookmarkApplicationJSON contains the JSON metadata
// for the struct [AccessApplicationGetResponseBookmarkApplication]
type accessApplicationGetResponseBookmarkApplicationJSON struct {
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

func (r *AccessApplicationGetResponseBookmarkApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseBookmarkApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationGetResponseBookmarkApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

type AccessApplicationRevokeTokensResponse = interface{}

type AccessApplicationNewParams struct {
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
	AutoRedirectToIdentity param.Field[bool]                                  `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationNewParamsCorsHeaders] `json:"cors_headers"`
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
	PathCookieAttribute param.Field[bool]                              `json:"path_cookie_attribute"`
	SaasApp             param.Field[AccessApplicationNewParamsSaasApp] `json:"saas_app"`
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

func (r AccessApplicationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationNewParamsCorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationNewParamsCorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationNewParamsCorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationNewParamsCorsHeadersAllowedMethod string

const (
	AccessApplicationNewParamsCorsHeadersAllowedMethodGet     AccessApplicationNewParamsCorsHeadersAllowedMethod = "GET"
	AccessApplicationNewParamsCorsHeadersAllowedMethodPost    AccessApplicationNewParamsCorsHeadersAllowedMethod = "POST"
	AccessApplicationNewParamsCorsHeadersAllowedMethodHead    AccessApplicationNewParamsCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationNewParamsCorsHeadersAllowedMethodPut     AccessApplicationNewParamsCorsHeadersAllowedMethod = "PUT"
	AccessApplicationNewParamsCorsHeadersAllowedMethodDelete  AccessApplicationNewParamsCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationNewParamsCorsHeadersAllowedMethodConnect AccessApplicationNewParamsCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationNewParamsCorsHeadersAllowedMethodOptions AccessApplicationNewParamsCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationNewParamsCorsHeadersAllowedMethodTrace   AccessApplicationNewParamsCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationNewParamsCorsHeadersAllowedMethodPatch   AccessApplicationNewParamsCorsHeadersAllowedMethod = "PATCH"
)

// Satisfied by [zero_trust.AccessApplicationNewParamsSaasAppAccessSamlSaasApp],
// [zero_trust.AccessApplicationNewParamsSaasAppAccessOidcSaasApp].
type AccessApplicationNewParamsSaasApp interface {
	implementsZeroTrustAccessApplicationNewParamsSaasApp()
}

type AccessApplicationNewParamsSaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[AccessApplicationNewParamsSaasAppAccessSamlSaasAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                                                             `json:"consumer_service_url"`
	CustomAttributes   param.Field[AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributes] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[AccessApplicationNewParamsSaasAppAccessSamlSaasAppNameIDFormat] `json:"name_id_format"`
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

func (r AccessApplicationNewParamsSaasAppAccessSamlSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsSaasAppAccessSamlSaasApp) implementsZeroTrustAccessApplicationNewParamsSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationNewParamsSaasAppAccessSamlSaasAppAuthType string

const (
	AccessApplicationNewParamsSaasAppAccessSamlSaasAppAuthTypeSaml AccessApplicationNewParamsSaasAppAccessSamlSaasAppAuthType = "saml"
	AccessApplicationNewParamsSaasAppAccessSamlSaasAppAuthTypeOidc AccessApplicationNewParamsSaasAppAccessSamlSaasAppAuthType = "oidc"
)

type AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat] `json:"name_format"`
	Source     param.Field[AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesSource]     `json:"source"`
}

func (r AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A globally unique name for an identity or service provider.
type AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
}

func (r AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the name identifier sent to the SaaS application.
type AccessApplicationNewParamsSaasAppAccessSamlSaasAppNameIDFormat string

const (
	AccessApplicationNewParamsSaasAppAccessSamlSaasAppNameIDFormatID    AccessApplicationNewParamsSaasAppAccessSamlSaasAppNameIDFormat = "id"
	AccessApplicationNewParamsSaasAppAccessSamlSaasAppNameIDFormatEmail AccessApplicationNewParamsSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type AccessApplicationNewParamsSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType param.Field[AccessApplicationNewParamsSaasAppAccessOidcSaasAppAuthType] `json:"auth_type"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string] `json:"client_secret"`
	// The OIDC flows supported by this application
	GrantTypes param.Field[[]AccessApplicationNewParamsSaasAppAccessOidcSaasAppGrantType] `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex param.Field[string] `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs param.Field[[]string] `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes param.Field[[]AccessApplicationNewParamsSaasAppAccessOidcSaasAppScope] `json:"scopes"`
}

func (r AccessApplicationNewParamsSaasAppAccessOidcSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationNewParamsSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type AccessApplicationNewParamsSaasAppAccessOidcSaasAppAuthType string

const (
	AccessApplicationNewParamsSaasAppAccessOidcSaasAppAuthTypeSaml AccessApplicationNewParamsSaasAppAccessOidcSaasAppAuthType = "saml"
	AccessApplicationNewParamsSaasAppAccessOidcSaasAppAuthTypeOidc AccessApplicationNewParamsSaasAppAccessOidcSaasAppAuthType = "oidc"
)

type AccessApplicationNewParamsSaasAppAccessOidcSaasAppGrantType string

const (
	AccessApplicationNewParamsSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         AccessApplicationNewParamsSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	AccessApplicationNewParamsSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce AccessApplicationNewParamsSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type AccessApplicationNewParamsSaasAppAccessOidcSaasAppScope string

const (
	AccessApplicationNewParamsSaasAppAccessOidcSaasAppScopeOpenid  AccessApplicationNewParamsSaasAppAccessOidcSaasAppScope = "openid"
	AccessApplicationNewParamsSaasAppAccessOidcSaasAppScopeGroups  AccessApplicationNewParamsSaasAppAccessOidcSaasAppScope = "groups"
	AccessApplicationNewParamsSaasAppAccessOidcSaasAppScopeEmail   AccessApplicationNewParamsSaasAppAccessOidcSaasAppScope = "email"
	AccessApplicationNewParamsSaasAppAccessOidcSaasAppScopeProfile AccessApplicationNewParamsSaasAppAccessOidcSaasAppScope = "profile"
)

type AccessApplicationNewResponseEnvelope struct {
	Errors   []AccessApplicationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationNewResponse                   `json:"result,required"`
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

type AccessApplicationNewResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accessApplicationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessApplicationNewResponseEnvelopeErrors]
type accessApplicationNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationNewResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessApplicationNewResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessApplicationNewResponseEnvelopeMessages]
type accessApplicationNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationNewResponseEnvelopeSuccess bool

const (
	AccessApplicationNewResponseEnvelopeSuccessTrue AccessApplicationNewResponseEnvelopeSuccess = true
)

type AccessApplicationUpdateParams struct {
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
	AutoRedirectToIdentity param.Field[bool]                                     `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationUpdateParamsCorsHeaders] `json:"cors_headers"`
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
	PathCookieAttribute param.Field[bool]                                 `json:"path_cookie_attribute"`
	SaasApp             param.Field[AccessApplicationUpdateParamsSaasApp] `json:"saas_app"`
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

func (r AccessApplicationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsAppID interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsAppID()
}

type AccessApplicationUpdateParamsCorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationUpdateParamsCorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationUpdateParamsCorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationUpdateParamsCorsHeadersAllowedMethod string

const (
	AccessApplicationUpdateParamsCorsHeadersAllowedMethodGet     AccessApplicationUpdateParamsCorsHeadersAllowedMethod = "GET"
	AccessApplicationUpdateParamsCorsHeadersAllowedMethodPost    AccessApplicationUpdateParamsCorsHeadersAllowedMethod = "POST"
	AccessApplicationUpdateParamsCorsHeadersAllowedMethodHead    AccessApplicationUpdateParamsCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationUpdateParamsCorsHeadersAllowedMethodPut     AccessApplicationUpdateParamsCorsHeadersAllowedMethod = "PUT"
	AccessApplicationUpdateParamsCorsHeadersAllowedMethodDelete  AccessApplicationUpdateParamsCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationUpdateParamsCorsHeadersAllowedMethodConnect AccessApplicationUpdateParamsCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationUpdateParamsCorsHeadersAllowedMethodOptions AccessApplicationUpdateParamsCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationUpdateParamsCorsHeadersAllowedMethodTrace   AccessApplicationUpdateParamsCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationUpdateParamsCorsHeadersAllowedMethodPatch   AccessApplicationUpdateParamsCorsHeadersAllowedMethod = "PATCH"
)

// Satisfied by [zero_trust.AccessApplicationUpdateParamsSaasAppAccessSamlSaasApp],
// [zero_trust.AccessApplicationUpdateParamsSaasAppAccessOidcSaasApp].
type AccessApplicationUpdateParamsSaasApp interface {
	implementsZeroTrustAccessApplicationUpdateParamsSaasApp()
}

type AccessApplicationUpdateParamsSaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                                                                `json:"consumer_service_url"`
	CustomAttributes   param.Field[AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributes] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppNameIDFormat] `json:"name_id_format"`
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

func (r AccessApplicationUpdateParamsSaasAppAccessSamlSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsSaasAppAccessSamlSaasApp) implementsZeroTrustAccessApplicationUpdateParamsSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppAuthType string

const (
	AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppAuthTypeSaml AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppAuthType = "saml"
	AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppAuthTypeOidc AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppAuthType = "oidc"
)

type AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat] `json:"name_format"`
	Source     param.Field[AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesSource]     `json:"source"`
}

func (r AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A globally unique name for an identity or service provider.
type AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
}

func (r AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the name identifier sent to the SaaS application.
type AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppNameIDFormat string

const (
	AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppNameIDFormatID    AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppNameIDFormat = "id"
	AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppNameIDFormatEmail AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type AccessApplicationUpdateParamsSaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType param.Field[AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppAuthType] `json:"auth_type"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string] `json:"client_secret"`
	// The OIDC flows supported by this application
	GrantTypes param.Field[[]AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppGrantType] `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex param.Field[string] `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs param.Field[[]string] `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes param.Field[[]AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScope] `json:"scopes"`
}

func (r AccessApplicationUpdateParamsSaasAppAccessOidcSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationUpdateParamsSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppAuthType string

const (
	AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppAuthTypeSaml AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppAuthType = "saml"
	AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppAuthTypeOidc AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppAuthType = "oidc"
)

type AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppGrantType string

const (
	AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScope string

const (
	AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScopeOpenid  AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScope = "openid"
	AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScopeGroups  AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScope = "groups"
	AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScopeEmail   AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScope = "email"
	AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScopeProfile AccessApplicationUpdateParamsSaasAppAccessOidcSaasAppScope = "profile"
)

type AccessApplicationUpdateResponseEnvelope struct {
	Errors   []AccessApplicationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationUpdateResponse                   `json:"result,required"`
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

type AccessApplicationUpdateResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accessApplicationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessApplicationUpdateResponseEnvelopeErrors]
type accessApplicationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationUpdateResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accessApplicationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessApplicationUpdateResponseEnvelopeMessages]
type accessApplicationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationUpdateResponseEnvelopeSuccess bool

const (
	AccessApplicationUpdateResponseEnvelopeSuccessTrue AccessApplicationUpdateResponseEnvelopeSuccess = true
)

type AccessApplicationListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationListResponseEnvelope struct {
	Errors   []AccessApplicationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessApplicationListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessApplicationListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessApplicationListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessApplicationListResponseEnvelopeJSON       `json:"-"`
}

// accessApplicationListResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationListResponseEnvelope]
type accessApplicationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationListResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accessApplicationListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessApplicationListResponseEnvelopeErrors]
type accessApplicationListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationListResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accessApplicationListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessApplicationListResponseEnvelopeMessages]
type accessApplicationListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationListResponseEnvelopeSuccess bool

const (
	AccessApplicationListResponseEnvelopeSuccessTrue AccessApplicationListResponseEnvelopeSuccess = true
)

type AccessApplicationListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       accessApplicationListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessApplicationListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AccessApplicationListResponseEnvelopeResultInfo]
type accessApplicationListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationDeleteParamsAppID interface {
	ImplementsZeroTrustAccessApplicationDeleteParamsAppID()
}

type AccessApplicationDeleteResponseEnvelope struct {
	Errors   []AccessApplicationDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationDeleteResponse                   `json:"result,required"`
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

type AccessApplicationDeleteResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accessApplicationDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessApplicationDeleteResponseEnvelopeErrors]
type accessApplicationDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationDeleteResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accessApplicationDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessApplicationDeleteResponseEnvelopeMessages]
type accessApplicationDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationDeleteResponseEnvelopeSuccess bool

const (
	AccessApplicationDeleteResponseEnvelopeSuccessTrue AccessApplicationDeleteResponseEnvelopeSuccess = true
)

type AccessApplicationGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationGetParamsAppID interface {
	ImplementsZeroTrustAccessApplicationGetParamsAppID()
}

type AccessApplicationGetResponseEnvelope struct {
	Errors   []AccessApplicationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationGetResponse                   `json:"result,required"`
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

type AccessApplicationGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accessApplicationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessApplicationGetResponseEnvelopeErrors]
type accessApplicationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessApplicationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessApplicationGetResponseEnvelopeMessages]
type accessApplicationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationGetResponseEnvelopeSuccess bool

const (
	AccessApplicationGetResponseEnvelopeSuccessTrue AccessApplicationGetResponseEnvelopeSuccess = true
)

type AccessApplicationRevokeTokensParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationRevokeTokensParamsAppID interface {
	ImplementsZeroTrustAccessApplicationRevokeTokensParamsAppID()
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

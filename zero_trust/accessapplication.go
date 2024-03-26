// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
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
func (r *AccessApplicationService) New(ctx context.Context, params AccessApplicationNewParams, opts ...option.RequestOption) (res *AccessApps, err error) {
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
func (r *AccessApplicationService) Update(ctx context.Context, appID AccessApplicationUpdateParamsSelfHostedApplicationAppID, params AccessApplicationUpdateParams, opts ...option.RequestOption) (res *AccessApps, err error) {
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
func (r *AccessApplicationService) List(ctx context.Context, query AccessApplicationListParams, opts ...option.RequestOption) (res *[]AccessApps, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches information about an Access application.
func (r *AccessApplicationService) Get(ctx context.Context, appID AccessApplicationGetParamsAppID, query AccessApplicationGetParams, opts ...option.RequestOption) (res *AccessApps, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [zero_trust.AccessAppsSelfHostedApplication],
// [zero_trust.AccessAppsSaaSApplication],
// [zero_trust.AccessAppsBrowserSSHApplication],
// [zero_trust.AccessAppsBrowserVncApplication],
// [zero_trust.AccessAppsAppLauncherApplication],
// [zero_trust.AccessAppsDeviceEnrollmentPermissionsApplication],
// [zero_trust.AccessAppsBrowserIsolationPermissionsApplication] or
// [zero_trust.AccessAppsBookmarkApplication].
type AccessApps interface {
	implementsZeroTrustAccessApps()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApps)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsSelfHostedApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsSaaSApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsBrowserSSHApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsBrowserVncApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsAppLauncherApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsDeviceEnrollmentPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsBrowserIsolationPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsBookmarkApplication{}),
		},
	)
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
	AllowedIDPs []string `json:"allowed_idps"`
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

func (r *AccessAppsSelfHostedApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsSelfHostedApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsSelfHostedApplication) implementsZeroTrustAccessApps() {}

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

func (r accessAppsSelfHostedApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
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

func (r AccessAppsSelfHostedApplicationCorsHeadersAllowedMethod) IsKnown() bool {
	switch r {
	case AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodGet, AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodPost, AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodHead, AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodPut, AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodDelete, AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodConnect, AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodOptions, AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodTrace, AccessAppsSelfHostedApplicationCorsHeadersAllowedMethodPatch:
		return true
	}
	return false
}

type AccessAppsSaaSApplication struct {
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs []string `json:"allowed_idps"`
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

func (r *AccessAppsSaaSApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsSaaSApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsSaaSApplication) implementsZeroTrustAccessApps() {}

// Union satisfied by
// [zero_trust.AccessAppsSaaSApplicationSaasAppAccessSamlSaasApp] or
// [zero_trust.AccessAppsSaaSApplicationSaasAppAccessOidcSaasApp].
type AccessAppsSaaSApplicationSaasApp interface {
	implementsZeroTrustAccessAppsSaaSApplicationSaasApp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessAppsSaaSApplicationSaasApp)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsSaaSApplicationSaasAppAccessSamlSaasApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsSaaSApplicationSaasAppAccessOidcSaasApp{}),
		},
	)
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
	IDPEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat `json:"name_id_format"`
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
	SSOEndpoint string                                                `json:"sso_endpoint"`
	UpdatedAt   time.Time                                             `json:"updated_at" format:"date-time"`
	JSON        accessAppsSaaSApplicationSaasAppAccessSamlSaasAppJSON `json:"-"`
}

// accessAppsSaaSApplicationSaasAppAccessSamlSaasAppJSON contains the JSON metadata
// for the struct [AccessAppsSaaSApplicationSaasAppAccessSamlSaasApp]
type accessAppsSaaSApplicationSaasAppAccessSamlSaasAppJSON struct {
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

func (r *AccessAppsSaaSApplicationSaasAppAccessSamlSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsSaaSApplicationSaasAppAccessSamlSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsSaaSApplicationSaasAppAccessSamlSaasApp) implementsZeroTrustAccessAppsSaaSApplicationSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppAuthType string

const (
	AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeSaml AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "saml"
	AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeOidc AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "oidc"
)

func (r AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppAuthType) IsKnown() bool {
	switch r {
	case AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeSaml, AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeOidc:
		return true
	}
	return false
}

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

func (r accessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesJSON) RawJSON() string {
	return r.raw
}

// A globally unique name for an identity or service provider.
type AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

func (r AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat) IsKnown() bool {
	switch r {
	case AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified, AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic, AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI:
		return true
	}
	return false
}

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

func (r accessAppsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSourceJSON) RawJSON() string {
	return r.raw
}

// The format of the name identifier sent to the SaaS application.
type AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat string

const (
	AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatID    AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "id"
	AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatEmail AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

func (r AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat) IsKnown() bool {
	switch r {
	case AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatID, AccessAppsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatEmail:
		return true
	}
	return false
}

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
	RedirectURIs []string `json:"redirect_uris"`
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
	RedirectURIs     apijson.Field
	Scopes           apijson.Field
	UpdatedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessAppsSaaSApplicationSaasAppAccessOidcSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsSaaSApplicationSaasAppAccessOidcSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsSaaSApplicationSaasAppAccessOidcSaasApp) implementsZeroTrustAccessAppsSaaSApplicationSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppAuthType string

const (
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "saml"
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppAuthType = "oidc"
)

func (r AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppAuthType) IsKnown() bool {
	switch r {
	case AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeSaml, AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppAuthTypeOidc:
		return true
	}
	return false
}

type AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppGrantType string

const (
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

func (r AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppGrantType) IsKnown() bool {
	switch r {
	case AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode, AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce:
		return true
	}
	return false
}

type AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScope string

const (
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid  AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScope = "openid"
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups  AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScope = "groups"
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail   AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScope = "email"
	AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScope = "profile"
)

func (r AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScope) IsKnown() bool {
	switch r {
	case AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScopeOpenid, AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScopeGroups, AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScopeEmail, AccessAppsSaaSApplicationSaasAppAccessOidcSaasAppScopeProfile:
		return true
	}
	return false
}

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
	AllowedIDPs []string `json:"allowed_idps"`
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

func (r *AccessAppsBrowserSSHApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsBrowserSSHApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsBrowserSSHApplication) implementsZeroTrustAccessApps() {}

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

func (r accessAppsBrowserSSHApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
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

func (r AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethod) IsKnown() bool {
	switch r {
	case AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodGet, AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodPost, AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodHead, AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodPut, AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodDelete, AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodConnect, AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodOptions, AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodTrace, AccessAppsBrowserSSHApplicationCorsHeadersAllowedMethodPatch:
		return true
	}
	return false
}

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
	AllowedIDPs []string `json:"allowed_idps"`
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

func (r *AccessAppsBrowserVncApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsBrowserVncApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsBrowserVncApplication) implementsZeroTrustAccessApps() {}

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

func (r accessAppsBrowserVncApplicationCorsHeadersJSON) RawJSON() string {
	return r.raw
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

func (r AccessAppsBrowserVncApplicationCorsHeadersAllowedMethod) IsKnown() bool {
	switch r {
	case AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodGet, AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodPost, AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodHead, AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodPut, AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodDelete, AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodConnect, AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodOptions, AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodTrace, AccessAppsBrowserVncApplicationCorsHeadersAllowedMethodPatch:
		return true
	}
	return false
}

type AccessAppsAppLauncherApplication struct {
	// The application type.
	Type AccessAppsAppLauncherApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs []string `json:"allowed_idps"`
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

func (r *AccessAppsAppLauncherApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsAppLauncherApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsAppLauncherApplication) implementsZeroTrustAccessApps() {}

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

func (r AccessAppsAppLauncherApplicationType) IsKnown() bool {
	switch r {
	case AccessAppsAppLauncherApplicationTypeSelfHosted, AccessAppsAppLauncherApplicationTypeSaas, AccessAppsAppLauncherApplicationTypeSSH, AccessAppsAppLauncherApplicationTypeVnc, AccessAppsAppLauncherApplicationTypeAppLauncher, AccessAppsAppLauncherApplicationTypeWARP, AccessAppsAppLauncherApplicationTypeBiso, AccessAppsAppLauncherApplicationTypeBookmark, AccessAppsAppLauncherApplicationTypeDashSSO:
		return true
	}
	return false
}

type AccessAppsDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type AccessAppsDeviceEnrollmentPermissionsApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs []string `json:"allowed_idps"`
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

func (r *AccessAppsDeviceEnrollmentPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsDeviceEnrollmentPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsDeviceEnrollmentPermissionsApplication) implementsZeroTrustAccessApps() {}

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

func (r AccessAppsDeviceEnrollmentPermissionsApplicationType) IsKnown() bool {
	switch r {
	case AccessAppsDeviceEnrollmentPermissionsApplicationTypeSelfHosted, AccessAppsDeviceEnrollmentPermissionsApplicationTypeSaas, AccessAppsDeviceEnrollmentPermissionsApplicationTypeSSH, AccessAppsDeviceEnrollmentPermissionsApplicationTypeVnc, AccessAppsDeviceEnrollmentPermissionsApplicationTypeAppLauncher, AccessAppsDeviceEnrollmentPermissionsApplicationTypeWARP, AccessAppsDeviceEnrollmentPermissionsApplicationTypeBiso, AccessAppsDeviceEnrollmentPermissionsApplicationTypeBookmark, AccessAppsDeviceEnrollmentPermissionsApplicationTypeDashSSO:
		return true
	}
	return false
}

type AccessAppsBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type AccessAppsBrowserIsolationPermissionsApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs []string `json:"allowed_idps"`
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

func (r *AccessAppsBrowserIsolationPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsBrowserIsolationPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsBrowserIsolationPermissionsApplication) implementsZeroTrustAccessApps() {}

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

func (r AccessAppsBrowserIsolationPermissionsApplicationType) IsKnown() bool {
	switch r {
	case AccessAppsBrowserIsolationPermissionsApplicationTypeSelfHosted, AccessAppsBrowserIsolationPermissionsApplicationTypeSaas, AccessAppsBrowserIsolationPermissionsApplicationTypeSSH, AccessAppsBrowserIsolationPermissionsApplicationTypeVnc, AccessAppsBrowserIsolationPermissionsApplicationTypeAppLauncher, AccessAppsBrowserIsolationPermissionsApplicationTypeWARP, AccessAppsBrowserIsolationPermissionsApplicationTypeBiso, AccessAppsBrowserIsolationPermissionsApplicationTypeBookmark, AccessAppsBrowserIsolationPermissionsApplicationTypeDashSSO:
		return true
	}
	return false
}

type AccessAppsBookmarkApplication struct {
	// UUID
	ID                 string      `json:"id"`
	AppLauncherVisible interface{} `json:"app_launcher_visible"`
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

func (r accessAppsBookmarkApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsBookmarkApplication) implementsZeroTrustAccessApps() {}

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
	AllowedIDPs param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                                       `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationNewParamsSelfHostedApplicationCorsHeaders] `json:"cors_headers"`
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

type AccessApplicationNewParamsSelfHostedApplicationCorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationNewParamsSelfHostedApplicationCorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodGet     AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodPost    AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodHead    AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodPut     AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodDelete  AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodConnect AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodOptions AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodTrace   AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodPatch   AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethod = "PATCH"
)

func (r AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethod) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodGet, AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodPost, AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodHead, AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodPut, AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodDelete, AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodConnect, AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodOptions, AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodTrace, AccessApplicationNewParamsSelfHostedApplicationCorsHeadersAllowedMethodPatch:
		return true
	}
	return false
}

type AccessApplicationNewParamsSaaSApplication struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages param.Field[[]string] `json:"custom_pages"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name    param.Field[string]                                           `json:"name"`
	SaasApp param.Field[AccessApplicationNewParamsSaaSApplicationSaasApp] `json:"saas_app"`
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

// Satisfied by
// [zero_trust.AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasApp],
// [zero_trust.AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasApp].
type AccessApplicationNewParamsSaaSApplicationSaasApp interface {
	implementsZeroTrustAccessApplicationNewParamsSaaSApplicationSaasApp()
}

type AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                                                                            `json:"consumer_service_url"`
	CustomAttributes   param.Field[AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IDPEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat] `json:"name_id_format"`
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

func (r AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasApp) implementsZeroTrustAccessApplicationNewParamsSaaSApplicationSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthType string

const (
	AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeSaml AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "saml"
	AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeOidc AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "oidc"
)

func (r AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthType) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeSaml, AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeOidc:
		return true
	}
	return false
}

type AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat] `json:"name_format"`
	Source     param.Field[AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource]     `json:"source"`
}

func (r AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A globally unique name for an identity or service provider.
type AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

func (r AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified, AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic, AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI:
		return true
	}
	return false
}

type AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
}

func (r AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the name identifier sent to the SaaS application.
type AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat string

const (
	AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatID    AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "id"
	AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatEmail AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

func (r AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatID, AccessApplicationNewParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatEmail:
		return true
	}
	return false
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

func (r AccessApplicationNewParamsSaaSApplicationSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationNewParamsSaaSApplicationSaasApp() {
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
	AllowedIDPs param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                                       `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationNewParamsBrowserSSHApplicationCorsHeaders] `json:"cors_headers"`
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

type AccessApplicationNewParamsBrowserSSHApplicationCorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationNewParamsBrowserSSHApplicationCorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodGet     AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodPost    AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodHead    AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodPut     AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodDelete  AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodConnect AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodOptions AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodTrace   AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodPatch   AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "PATCH"
)

func (r AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethod) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodGet, AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodPost, AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodHead, AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodPut, AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodDelete, AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodConnect, AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodOptions, AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodTrace, AccessApplicationNewParamsBrowserSSHApplicationCorsHeadersAllowedMethodPatch:
		return true
	}
	return false
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
	AllowedIDPs param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                                       `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationNewParamsBrowserVncApplicationCorsHeaders] `json:"cors_headers"`
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

type AccessApplicationNewParamsBrowserVncApplicationCorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationNewParamsBrowserVncApplicationCorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodGet     AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodPost    AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodHead    AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodPut     AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodDelete  AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodConnect AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodOptions AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodTrace   AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodPatch   AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethod = "PATCH"
)

func (r AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethod) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodGet, AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodPost, AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodHead, AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodPut, AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodDelete, AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodConnect, AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodOptions, AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodTrace, AccessApplicationNewParamsBrowserVncApplicationCorsHeadersAllowedMethodPatch:
		return true
	}
	return false
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
	AllowedIDPs param.Field[[]string] `json:"allowed_idps"`
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
	AllowedIDPs param.Field[[]string] `json:"allowed_idps"`
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
	AllowedIDPs param.Field[[]string] `json:"allowed_idps"`
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
	ZoneID             param.Field[string]      `path:"zone_id"`
	AppLauncherVisible param.Field[interface{}] `json:"app_launcher_visible"`
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
	Errors   []AccessApplicationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApps                                     `json:"result,required"`
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
	AllowedIDPs param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                                          `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationUpdateParamsSelfHostedApplicationCorsHeaders] `json:"cors_headers"`
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

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsSelfHostedApplicationAppID interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsSelfHostedApplicationAppID()
}

type AccessApplicationUpdateParamsSelfHostedApplicationCorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationUpdateParamsSelfHostedApplicationCorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodGet     AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodPost    AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodHead    AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodPut     AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodDelete  AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodConnect AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodOptions AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodTrace   AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodPatch   AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethod = "PATCH"
)

func (r AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethod) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodGet, AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodPost, AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodHead, AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodPut, AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodDelete, AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodConnect, AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodOptions, AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodTrace, AccessApplicationUpdateParamsSelfHostedApplicationCorsHeadersAllowedMethodPatch:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsSaaSApplication struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIDPs param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages param.Field[[]string] `json:"custom_pages"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name    param.Field[string]                                              `json:"name"`
	SaasApp param.Field[AccessApplicationUpdateParamsSaaSApplicationSaasApp] `json:"saas_app"`
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

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsSaaSApplicationAppID interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsSaaSApplicationAppID()
}

// Satisfied by
// [zero_trust.AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasApp],
// [zero_trust.AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasApp].
type AccessApplicationUpdateParamsSaaSApplicationSaasApp interface {
	implementsZeroTrustAccessApplicationUpdateParamsSaaSApplicationSaasApp()
}

type AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                                                                               `json:"consumer_service_url"`
	CustomAttributes   param.Field[AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IDPEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat] `json:"name_id_format"`
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

func (r AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasApp) implementsZeroTrustAccessApplicationUpdateParamsSaaSApplicationSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthType string

const (
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeSaml AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "saml"
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeOidc AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthType = "oidc"
)

func (r AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthType) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeSaml, AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppAuthTypeOidc:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat] `json:"name_format"`
	Source     param.Field[AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource]     `json:"source"`
}

func (r AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A globally unique name for an identity or service provider.
type AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI         AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

func (r AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormat) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified, AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic, AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatURI:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
}

func (r AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppCustomAttributesSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the name identifier sent to the SaaS application.
type AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat string

const (
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatID    AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "id"
	AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatEmail AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat = "email"
)

func (r AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormat) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatID, AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessSamlSaasAppNameIDFormatEmail:
		return true
	}
	return false
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

func (r AccessApplicationUpdateParamsSaaSApplicationSaasAppAccessOidcSaasApp) implementsZeroTrustAccessApplicationUpdateParamsSaaSApplicationSaasApp() {
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
	AllowedIDPs param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                                          `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeaders] `json:"cors_headers"`
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

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsBrowserSSHApplicationAppID interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsBrowserSSHApplicationAppID()
}

type AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodGet     AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodPost    AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodHead    AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodPut     AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodDelete  AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodConnect AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodOptions AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodTrace   AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodPatch   AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethod = "PATCH"
)

func (r AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethod) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodGet, AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodPost, AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodHead, AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodPut, AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodDelete, AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodConnect, AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodOptions, AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodTrace, AccessApplicationUpdateParamsBrowserSSHApplicationCorsHeadersAllowedMethodPatch:
		return true
	}
	return false
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
	AllowedIDPs param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                                          `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationUpdateParamsBrowserVncApplicationCorsHeaders] `json:"cors_headers"`
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

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsBrowserVncApplicationAppID interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsBrowserVncApplicationAppID()
}

type AccessApplicationUpdateParamsBrowserVncApplicationCorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationUpdateParamsBrowserVncApplicationCorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethod string

const (
	AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodGet     AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethod = "GET"
	AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodPost    AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethod = "POST"
	AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodHead    AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodPut     AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethod = "PUT"
	AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodDelete  AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodConnect AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodOptions AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodTrace   AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodPatch   AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethod = "PATCH"
)

func (r AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethod) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodGet, AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodPost, AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodHead, AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodPut, AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodDelete, AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodConnect, AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodOptions, AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodTrace, AccessApplicationUpdateParamsBrowserVncApplicationCorsHeadersAllowedMethodPatch:
		return true
	}
	return false
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
	AllowedIDPs param.Field[[]string] `json:"allowed_idps"`
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

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsAppLauncherApplicationAppID interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsAppLauncherApplicationAppID()
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
	AllowedIDPs param.Field[[]string] `json:"allowed_idps"`
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

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationAppID interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsDeviceEnrollmentPermissionsApplicationAppID()
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
	AllowedIDPs param.Field[[]string] `json:"allowed_idps"`
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

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationAppID interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsBrowserIsolationPermissionsApplicationAppID()
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
	ZoneID             param.Field[string]      `path:"zone_id"`
	AppLauncherVisible param.Field[interface{}] `json:"app_launcher_visible"`
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

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsBookmarkApplicationAppID interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsBookmarkApplicationAppID()
}

type AccessApplicationUpdateResponseEnvelope struct {
	Errors   []AccessApplicationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApps                                        `json:"result,required"`
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

type AccessApplicationListResponseEnvelope struct {
	Errors   []AccessApplicationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessApps                                    `json:"result,required,nullable"`
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

func (r AccessApplicationListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationGetParamsAppID interface {
	ImplementsZeroTrustAccessApplicationGetParamsAppID()
}

type AccessApplicationGetResponseEnvelope struct {
	Errors   []AccessApplicationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApps                                     `json:"result,required"`
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

func (r AccessApplicationRevokeTokensResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationRevokeTokensResponseEnvelopeSuccessTrue, AccessApplicationRevokeTokensResponseEnvelopeSuccessFalse:
		return true
	}
	return false
}

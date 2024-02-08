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

// AccessAppService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccessAppService] method instead.
type AccessAppService struct {
	Options          []option.RequestOption
	Cas              *AccessAppCaService
	RevokeTokens     *AccessAppRevokeTokenService
	UserPolicyChecks *AccessAppUserPolicyCheckService
	Policies         *AccessAppPolicyService
}

// NewAccessAppService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessAppService(opts ...option.RequestOption) (r *AccessAppService) {
	r = &AccessAppService{}
	r.Options = opts
	r.Cas = NewAccessAppCaService(opts...)
	r.RevokeTokens = NewAccessAppRevokeTokenService(opts...)
	r.UserPolicyChecks = NewAccessAppUserPolicyCheckService(opts...)
	r.Policies = NewAccessAppPolicyService(opts...)
	return
}

// Adds a new application to Access.
func (r *AccessAppService) New(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessAppNewParams, opts ...option.RequestOption) (res *AccessAppNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppNewResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches information about an Access application.
func (r *AccessAppService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, appID AccessAppGetParamsAppID, opts ...option.RequestOption) (res *AccessAppGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppGetResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%v", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an Access application.
func (r *AccessAppService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, appID AccessAppUpdateParamsVariant0AppID, body AccessAppUpdateParams, opts ...option.RequestOption) (res *AccessAppUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppUpdateResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%v", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Access applications in an account or zone.
func (r *AccessAppService) List(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *[]AccessAppListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppListResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an application from Access.
func (r *AccessAppService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, appID AccessAppDeleteParamsAppID, opts ...option.RequestOption) (res *AccessAppDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppDeleteResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%v", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AccessAppNewResponseObject], [AccessAppNewResponseObject],
// [AccessAppNewResponseObject], [AccessAppNewResponseObject],
// [AccessAppNewResponseObject], [AccessAppNewResponseObject],
// [AccessAppNewResponseObject] or [AccessAppNewResponseObject].
type AccessAppNewResponse interface {
	implementsAccessAppNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppNewResponse)(nil)).Elem(), "")
}

type AccessAppNewResponseObject struct {
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
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps []string `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	Aud string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool                                  `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessAppNewResponseObjectCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                             `json:"created_at" format:"date-time"`
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
	Tags      []string                       `json:"tags"`
	UpdatedAt time.Time                      `json:"updated_at" format:"date-time"`
	JSON      accessAppNewResponseObjectJSON `json:"-"`
}

// accessAppNewResponseObjectJSON contains the JSON metadata for the struct
// [AccessAppNewResponseObject]
type accessAppNewResponseObjectJSON struct {
	Domain                   apijson.Field
	Type                     apijson.Field
	ID                       apijson.Field
	AllowAuthenticateViaWarp apijson.Field
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

func (r *AccessAppNewResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppNewResponseObject) implementsAccessAppNewResponse() {}

type AccessAppNewResponseObjectCorsHeaders struct {
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
	AllowedMethods []AccessAppNewResponseObjectCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                   `json:"max_age"`
	JSON   accessAppNewResponseObjectCorsHeadersJSON `json:"-"`
}

// accessAppNewResponseObjectCorsHeadersJSON contains the JSON metadata for the
// struct [AccessAppNewResponseObjectCorsHeaders]
type accessAppNewResponseObjectCorsHeadersJSON struct {
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

func (r *AccessAppNewResponseObjectCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppNewResponseObjectCorsHeadersAllowedMethod string

const (
	AccessAppNewResponseObjectCorsHeadersAllowedMethodGet     AccessAppNewResponseObjectCorsHeadersAllowedMethod = "GET"
	AccessAppNewResponseObjectCorsHeadersAllowedMethodPost    AccessAppNewResponseObjectCorsHeadersAllowedMethod = "POST"
	AccessAppNewResponseObjectCorsHeadersAllowedMethodHead    AccessAppNewResponseObjectCorsHeadersAllowedMethod = "HEAD"
	AccessAppNewResponseObjectCorsHeadersAllowedMethodPut     AccessAppNewResponseObjectCorsHeadersAllowedMethod = "PUT"
	AccessAppNewResponseObjectCorsHeadersAllowedMethodDelete  AccessAppNewResponseObjectCorsHeadersAllowedMethod = "DELETE"
	AccessAppNewResponseObjectCorsHeadersAllowedMethodConnect AccessAppNewResponseObjectCorsHeadersAllowedMethod = "CONNECT"
	AccessAppNewResponseObjectCorsHeadersAllowedMethodOptions AccessAppNewResponseObjectCorsHeadersAllowedMethod = "OPTIONS"
	AccessAppNewResponseObjectCorsHeadersAllowedMethodTrace   AccessAppNewResponseObjectCorsHeadersAllowedMethod = "TRACE"
	AccessAppNewResponseObjectCorsHeadersAllowedMethodPatch   AccessAppNewResponseObjectCorsHeadersAllowedMethod = "PATCH"
)

// Union satisfied by [AccessAppGetResponseObject], [AccessAppGetResponseObject],
// [AccessAppGetResponseObject], [AccessAppGetResponseObject],
// [AccessAppGetResponseObject], [AccessAppGetResponseObject],
// [AccessAppGetResponseObject] or [AccessAppGetResponseObject].
type AccessAppGetResponse interface {
	implementsAccessAppGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppGetResponse)(nil)).Elem(), "")
}

type AccessAppGetResponseObject struct {
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
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps []string `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	Aud string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool                                  `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessAppGetResponseObjectCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                             `json:"created_at" format:"date-time"`
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
	Tags      []string                       `json:"tags"`
	UpdatedAt time.Time                      `json:"updated_at" format:"date-time"`
	JSON      accessAppGetResponseObjectJSON `json:"-"`
}

// accessAppGetResponseObjectJSON contains the JSON metadata for the struct
// [AccessAppGetResponseObject]
type accessAppGetResponseObjectJSON struct {
	Domain                   apijson.Field
	Type                     apijson.Field
	ID                       apijson.Field
	AllowAuthenticateViaWarp apijson.Field
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

func (r *AccessAppGetResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppGetResponseObject) implementsAccessAppGetResponse() {}

type AccessAppGetResponseObjectCorsHeaders struct {
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
	AllowedMethods []AccessAppGetResponseObjectCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                   `json:"max_age"`
	JSON   accessAppGetResponseObjectCorsHeadersJSON `json:"-"`
}

// accessAppGetResponseObjectCorsHeadersJSON contains the JSON metadata for the
// struct [AccessAppGetResponseObjectCorsHeaders]
type accessAppGetResponseObjectCorsHeadersJSON struct {
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

func (r *AccessAppGetResponseObjectCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppGetResponseObjectCorsHeadersAllowedMethod string

const (
	AccessAppGetResponseObjectCorsHeadersAllowedMethodGet     AccessAppGetResponseObjectCorsHeadersAllowedMethod = "GET"
	AccessAppGetResponseObjectCorsHeadersAllowedMethodPost    AccessAppGetResponseObjectCorsHeadersAllowedMethod = "POST"
	AccessAppGetResponseObjectCorsHeadersAllowedMethodHead    AccessAppGetResponseObjectCorsHeadersAllowedMethod = "HEAD"
	AccessAppGetResponseObjectCorsHeadersAllowedMethodPut     AccessAppGetResponseObjectCorsHeadersAllowedMethod = "PUT"
	AccessAppGetResponseObjectCorsHeadersAllowedMethodDelete  AccessAppGetResponseObjectCorsHeadersAllowedMethod = "DELETE"
	AccessAppGetResponseObjectCorsHeadersAllowedMethodConnect AccessAppGetResponseObjectCorsHeadersAllowedMethod = "CONNECT"
	AccessAppGetResponseObjectCorsHeadersAllowedMethodOptions AccessAppGetResponseObjectCorsHeadersAllowedMethod = "OPTIONS"
	AccessAppGetResponseObjectCorsHeadersAllowedMethodTrace   AccessAppGetResponseObjectCorsHeadersAllowedMethod = "TRACE"
	AccessAppGetResponseObjectCorsHeadersAllowedMethodPatch   AccessAppGetResponseObjectCorsHeadersAllowedMethod = "PATCH"
)

// Union satisfied by [AccessAppUpdateResponseObject],
// [AccessAppUpdateResponseObject], [AccessAppUpdateResponseObject],
// [AccessAppUpdateResponseObject], [AccessAppUpdateResponseObject],
// [AccessAppUpdateResponseObject], [AccessAppUpdateResponseObject] or
// [AccessAppUpdateResponseObject].
type AccessAppUpdateResponse interface {
	implementsAccessAppUpdateResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppUpdateResponse)(nil)).Elem(), "")
}

type AccessAppUpdateResponseObject struct {
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
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps []string `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	Aud string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool                                     `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessAppUpdateResponseObjectCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                `json:"created_at" format:"date-time"`
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
	Tags      []string                          `json:"tags"`
	UpdatedAt time.Time                         `json:"updated_at" format:"date-time"`
	JSON      accessAppUpdateResponseObjectJSON `json:"-"`
}

// accessAppUpdateResponseObjectJSON contains the JSON metadata for the struct
// [AccessAppUpdateResponseObject]
type accessAppUpdateResponseObjectJSON struct {
	Domain                   apijson.Field
	Type                     apijson.Field
	ID                       apijson.Field
	AllowAuthenticateViaWarp apijson.Field
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

func (r *AccessAppUpdateResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppUpdateResponseObject) implementsAccessAppUpdateResponse() {}

type AccessAppUpdateResponseObjectCorsHeaders struct {
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
	AllowedMethods []AccessAppUpdateResponseObjectCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                      `json:"max_age"`
	JSON   accessAppUpdateResponseObjectCorsHeadersJSON `json:"-"`
}

// accessAppUpdateResponseObjectCorsHeadersJSON contains the JSON metadata for the
// struct [AccessAppUpdateResponseObjectCorsHeaders]
type accessAppUpdateResponseObjectCorsHeadersJSON struct {
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

func (r *AccessAppUpdateResponseObjectCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUpdateResponseObjectCorsHeadersAllowedMethod string

const (
	AccessAppUpdateResponseObjectCorsHeadersAllowedMethodGet     AccessAppUpdateResponseObjectCorsHeadersAllowedMethod = "GET"
	AccessAppUpdateResponseObjectCorsHeadersAllowedMethodPost    AccessAppUpdateResponseObjectCorsHeadersAllowedMethod = "POST"
	AccessAppUpdateResponseObjectCorsHeadersAllowedMethodHead    AccessAppUpdateResponseObjectCorsHeadersAllowedMethod = "HEAD"
	AccessAppUpdateResponseObjectCorsHeadersAllowedMethodPut     AccessAppUpdateResponseObjectCorsHeadersAllowedMethod = "PUT"
	AccessAppUpdateResponseObjectCorsHeadersAllowedMethodDelete  AccessAppUpdateResponseObjectCorsHeadersAllowedMethod = "DELETE"
	AccessAppUpdateResponseObjectCorsHeadersAllowedMethodConnect AccessAppUpdateResponseObjectCorsHeadersAllowedMethod = "CONNECT"
	AccessAppUpdateResponseObjectCorsHeadersAllowedMethodOptions AccessAppUpdateResponseObjectCorsHeadersAllowedMethod = "OPTIONS"
	AccessAppUpdateResponseObjectCorsHeadersAllowedMethodTrace   AccessAppUpdateResponseObjectCorsHeadersAllowedMethod = "TRACE"
	AccessAppUpdateResponseObjectCorsHeadersAllowedMethodPatch   AccessAppUpdateResponseObjectCorsHeadersAllowedMethod = "PATCH"
)

// Union satisfied by [AccessAppListResponseObject], [AccessAppListResponseObject],
// [AccessAppListResponseObject], [AccessAppListResponseObject],
// [AccessAppListResponseObject], [AccessAppListResponseObject],
// [AccessAppListResponseObject] or [AccessAppListResponseObject].
type AccessAppListResponse interface {
	implementsAccessAppListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppListResponse)(nil)).Elem(), "")
}

type AccessAppListResponseObject struct {
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
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps []string `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	Aud string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool                                   `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessAppListResponseObjectCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                              `json:"created_at" format:"date-time"`
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
	Tags      []string                        `json:"tags"`
	UpdatedAt time.Time                       `json:"updated_at" format:"date-time"`
	JSON      accessAppListResponseObjectJSON `json:"-"`
}

// accessAppListResponseObjectJSON contains the JSON metadata for the struct
// [AccessAppListResponseObject]
type accessAppListResponseObjectJSON struct {
	Domain                   apijson.Field
	Type                     apijson.Field
	ID                       apijson.Field
	AllowAuthenticateViaWarp apijson.Field
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

func (r *AccessAppListResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppListResponseObject) implementsAccessAppListResponse() {}

type AccessAppListResponseObjectCorsHeaders struct {
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
	AllowedMethods []AccessAppListResponseObjectCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                    `json:"max_age"`
	JSON   accessAppListResponseObjectCorsHeadersJSON `json:"-"`
}

// accessAppListResponseObjectCorsHeadersJSON contains the JSON metadata for the
// struct [AccessAppListResponseObjectCorsHeaders]
type accessAppListResponseObjectCorsHeadersJSON struct {
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

func (r *AccessAppListResponseObjectCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppListResponseObjectCorsHeadersAllowedMethod string

const (
	AccessAppListResponseObjectCorsHeadersAllowedMethodGet     AccessAppListResponseObjectCorsHeadersAllowedMethod = "GET"
	AccessAppListResponseObjectCorsHeadersAllowedMethodPost    AccessAppListResponseObjectCorsHeadersAllowedMethod = "POST"
	AccessAppListResponseObjectCorsHeadersAllowedMethodHead    AccessAppListResponseObjectCorsHeadersAllowedMethod = "HEAD"
	AccessAppListResponseObjectCorsHeadersAllowedMethodPut     AccessAppListResponseObjectCorsHeadersAllowedMethod = "PUT"
	AccessAppListResponseObjectCorsHeadersAllowedMethodDelete  AccessAppListResponseObjectCorsHeadersAllowedMethod = "DELETE"
	AccessAppListResponseObjectCorsHeadersAllowedMethodConnect AccessAppListResponseObjectCorsHeadersAllowedMethod = "CONNECT"
	AccessAppListResponseObjectCorsHeadersAllowedMethodOptions AccessAppListResponseObjectCorsHeadersAllowedMethod = "OPTIONS"
	AccessAppListResponseObjectCorsHeadersAllowedMethodTrace   AccessAppListResponseObjectCorsHeadersAllowedMethod = "TRACE"
	AccessAppListResponseObjectCorsHeadersAllowedMethodPatch   AccessAppListResponseObjectCorsHeadersAllowedMethod = "PATCH"
)

type AccessAppDeleteResponse struct {
	// UUID
	ID   string                      `json:"id"`
	JSON accessAppDeleteResponseJSON `json:"-"`
}

// accessAppDeleteResponseJSON contains the JSON metadata for the struct
// [AccessAppDeleteResponse]
type accessAppDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This interface is a union satisfied by one of the following:
// [AccessAppNewParamsVariant0], [AccessAppNewParamsVariant1],
// [AccessAppNewParamsVariant2], [AccessAppNewParamsVariant3],
// [AccessAppNewParamsVariant4], [AccessAppNewParamsVariant5],
// [AccessAppNewParamsVariant6], [AccessAppNewParamsVariant7].
type AccessAppNewParams interface {
	ImplementsAccessAppNewParams()
}

type AccessAppNewParamsVariant0 struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWarp param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                  `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessAppNewParamsVariant0CorsHeaders] `json:"cors_headers"`
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

func (r AccessAppNewParamsVariant0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppNewParamsVariant0) ImplementsAccessAppNewParams() {

}

type AccessAppNewParamsVariant0CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessAppNewParamsVariant0CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessAppNewParamsVariant0CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppNewParamsVariant0CorsHeadersAllowedMethod string

const (
	AccessAppNewParamsVariant0CorsHeadersAllowedMethodGet     AccessAppNewParamsVariant0CorsHeadersAllowedMethod = "GET"
	AccessAppNewParamsVariant0CorsHeadersAllowedMethodPost    AccessAppNewParamsVariant0CorsHeadersAllowedMethod = "POST"
	AccessAppNewParamsVariant0CorsHeadersAllowedMethodHead    AccessAppNewParamsVariant0CorsHeadersAllowedMethod = "HEAD"
	AccessAppNewParamsVariant0CorsHeadersAllowedMethodPut     AccessAppNewParamsVariant0CorsHeadersAllowedMethod = "PUT"
	AccessAppNewParamsVariant0CorsHeadersAllowedMethodDelete  AccessAppNewParamsVariant0CorsHeadersAllowedMethod = "DELETE"
	AccessAppNewParamsVariant0CorsHeadersAllowedMethodConnect AccessAppNewParamsVariant0CorsHeadersAllowedMethod = "CONNECT"
	AccessAppNewParamsVariant0CorsHeadersAllowedMethodOptions AccessAppNewParamsVariant0CorsHeadersAllowedMethod = "OPTIONS"
	AccessAppNewParamsVariant0CorsHeadersAllowedMethodTrace   AccessAppNewParamsVariant0CorsHeadersAllowedMethod = "TRACE"
	AccessAppNewParamsVariant0CorsHeadersAllowedMethodPatch   AccessAppNewParamsVariant0CorsHeadersAllowedMethod = "PATCH"
)

type AccessAppNewParamsVariant1 struct {
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
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
	Name    param.Field[string]                            `json:"name"`
	SaasApp param.Field[AccessAppNewParamsVariant1SaasApp] `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessAppNewParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppNewParamsVariant1) ImplementsAccessAppNewParams() {

}

type AccessAppNewParamsVariant1SaasApp struct {
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                                            `json:"consumer_service_url"`
	CustomAttributes   param.Field[AccessAppNewParamsVariant1SaasAppCustomAttributes] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[AccessAppNewParamsVariant1SaasAppNameIDFormat] `json:"name_id_format"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// A globally unique name for an identity or service provider.
	SpEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string] `json:"sso_endpoint"`
}

func (r AccessAppNewParamsVariant1SaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppNewParamsVariant1SaasAppCustomAttributes struct {
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[AccessAppNewParamsVariant1SaasAppCustomAttributesNameFormat] `json:"name_format"`
	Source     param.Field[AccessAppNewParamsVariant1SaasAppCustomAttributesSource]     `json:"source"`
}

func (r AccessAppNewParamsVariant1SaasAppCustomAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A globally unique name for an identity or service provider.
type AccessAppNewParamsVariant1SaasAppCustomAttributesNameFormat string

const (
	AccessAppNewParamsVariant1SaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessAppNewParamsVariant1SaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessAppNewParamsVariant1SaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessAppNewParamsVariant1SaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessAppNewParamsVariant1SaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUri         AccessAppNewParamsVariant1SaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type AccessAppNewParamsVariant1SaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
}

func (r AccessAppNewParamsVariant1SaasAppCustomAttributesSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the name identifier sent to the SaaS application.
type AccessAppNewParamsVariant1SaasAppNameIDFormat string

const (
	AccessAppNewParamsVariant1SaasAppNameIDFormatID    AccessAppNewParamsVariant1SaasAppNameIDFormat = "id"
	AccessAppNewParamsVariant1SaasAppNameIDFormatEmail AccessAppNewParamsVariant1SaasAppNameIDFormat = "email"
)

type AccessAppNewParamsVariant2 struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWarp param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                  `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessAppNewParamsVariant2CorsHeaders] `json:"cors_headers"`
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

func (r AccessAppNewParamsVariant2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppNewParamsVariant2) ImplementsAccessAppNewParams() {

}

type AccessAppNewParamsVariant2CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessAppNewParamsVariant2CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessAppNewParamsVariant2CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppNewParamsVariant2CorsHeadersAllowedMethod string

const (
	AccessAppNewParamsVariant2CorsHeadersAllowedMethodGet     AccessAppNewParamsVariant2CorsHeadersAllowedMethod = "GET"
	AccessAppNewParamsVariant2CorsHeadersAllowedMethodPost    AccessAppNewParamsVariant2CorsHeadersAllowedMethod = "POST"
	AccessAppNewParamsVariant2CorsHeadersAllowedMethodHead    AccessAppNewParamsVariant2CorsHeadersAllowedMethod = "HEAD"
	AccessAppNewParamsVariant2CorsHeadersAllowedMethodPut     AccessAppNewParamsVariant2CorsHeadersAllowedMethod = "PUT"
	AccessAppNewParamsVariant2CorsHeadersAllowedMethodDelete  AccessAppNewParamsVariant2CorsHeadersAllowedMethod = "DELETE"
	AccessAppNewParamsVariant2CorsHeadersAllowedMethodConnect AccessAppNewParamsVariant2CorsHeadersAllowedMethod = "CONNECT"
	AccessAppNewParamsVariant2CorsHeadersAllowedMethodOptions AccessAppNewParamsVariant2CorsHeadersAllowedMethod = "OPTIONS"
	AccessAppNewParamsVariant2CorsHeadersAllowedMethodTrace   AccessAppNewParamsVariant2CorsHeadersAllowedMethod = "TRACE"
	AccessAppNewParamsVariant2CorsHeadersAllowedMethodPatch   AccessAppNewParamsVariant2CorsHeadersAllowedMethod = "PATCH"
)

type AccessAppNewParamsVariant3 struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWarp param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                  `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessAppNewParamsVariant3CorsHeaders] `json:"cors_headers"`
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

func (r AccessAppNewParamsVariant3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppNewParamsVariant3) ImplementsAccessAppNewParams() {

}

type AccessAppNewParamsVariant3CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessAppNewParamsVariant3CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessAppNewParamsVariant3CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppNewParamsVariant3CorsHeadersAllowedMethod string

const (
	AccessAppNewParamsVariant3CorsHeadersAllowedMethodGet     AccessAppNewParamsVariant3CorsHeadersAllowedMethod = "GET"
	AccessAppNewParamsVariant3CorsHeadersAllowedMethodPost    AccessAppNewParamsVariant3CorsHeadersAllowedMethod = "POST"
	AccessAppNewParamsVariant3CorsHeadersAllowedMethodHead    AccessAppNewParamsVariant3CorsHeadersAllowedMethod = "HEAD"
	AccessAppNewParamsVariant3CorsHeadersAllowedMethodPut     AccessAppNewParamsVariant3CorsHeadersAllowedMethod = "PUT"
	AccessAppNewParamsVariant3CorsHeadersAllowedMethodDelete  AccessAppNewParamsVariant3CorsHeadersAllowedMethod = "DELETE"
	AccessAppNewParamsVariant3CorsHeadersAllowedMethodConnect AccessAppNewParamsVariant3CorsHeadersAllowedMethod = "CONNECT"
	AccessAppNewParamsVariant3CorsHeadersAllowedMethodOptions AccessAppNewParamsVariant3CorsHeadersAllowedMethod = "OPTIONS"
	AccessAppNewParamsVariant3CorsHeadersAllowedMethodTrace   AccessAppNewParamsVariant3CorsHeadersAllowedMethod = "TRACE"
	AccessAppNewParamsVariant3CorsHeadersAllowedMethodPatch   AccessAppNewParamsVariant3CorsHeadersAllowedMethod = "PATCH"
)

type AccessAppNewParamsVariant4 struct {
	// The application type.
	Type param.Field[AccessAppNewParamsVariant4Type] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessAppNewParamsVariant4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppNewParamsVariant4) ImplementsAccessAppNewParams() {

}

// The application type.
type AccessAppNewParamsVariant4Type string

const (
	AccessAppNewParamsVariant4TypeSelfHosted  AccessAppNewParamsVariant4Type = "self_hosted"
	AccessAppNewParamsVariant4TypeSaas        AccessAppNewParamsVariant4Type = "saas"
	AccessAppNewParamsVariant4TypeSSH         AccessAppNewParamsVariant4Type = "ssh"
	AccessAppNewParamsVariant4TypeVnc         AccessAppNewParamsVariant4Type = "vnc"
	AccessAppNewParamsVariant4TypeAppLauncher AccessAppNewParamsVariant4Type = "app_launcher"
	AccessAppNewParamsVariant4TypeWarp        AccessAppNewParamsVariant4Type = "warp"
	AccessAppNewParamsVariant4TypeBiso        AccessAppNewParamsVariant4Type = "biso"
	AccessAppNewParamsVariant4TypeBookmark    AccessAppNewParamsVariant4Type = "bookmark"
	AccessAppNewParamsVariant4TypeDashSSO     AccessAppNewParamsVariant4Type = "dash_sso"
)

type AccessAppNewParamsVariant5 struct {
	// The application type.
	Type param.Field[AccessAppNewParamsVariant5Type] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessAppNewParamsVariant5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppNewParamsVariant5) ImplementsAccessAppNewParams() {

}

// The application type.
type AccessAppNewParamsVariant5Type string

const (
	AccessAppNewParamsVariant5TypeSelfHosted  AccessAppNewParamsVariant5Type = "self_hosted"
	AccessAppNewParamsVariant5TypeSaas        AccessAppNewParamsVariant5Type = "saas"
	AccessAppNewParamsVariant5TypeSSH         AccessAppNewParamsVariant5Type = "ssh"
	AccessAppNewParamsVariant5TypeVnc         AccessAppNewParamsVariant5Type = "vnc"
	AccessAppNewParamsVariant5TypeAppLauncher AccessAppNewParamsVariant5Type = "app_launcher"
	AccessAppNewParamsVariant5TypeWarp        AccessAppNewParamsVariant5Type = "warp"
	AccessAppNewParamsVariant5TypeBiso        AccessAppNewParamsVariant5Type = "biso"
	AccessAppNewParamsVariant5TypeBookmark    AccessAppNewParamsVariant5Type = "bookmark"
	AccessAppNewParamsVariant5TypeDashSSO     AccessAppNewParamsVariant5Type = "dash_sso"
)

type AccessAppNewParamsVariant6 struct {
	// The application type.
	Type param.Field[AccessAppNewParamsVariant6Type] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessAppNewParamsVariant6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppNewParamsVariant6) ImplementsAccessAppNewParams() {

}

// The application type.
type AccessAppNewParamsVariant6Type string

const (
	AccessAppNewParamsVariant6TypeSelfHosted  AccessAppNewParamsVariant6Type = "self_hosted"
	AccessAppNewParamsVariant6TypeSaas        AccessAppNewParamsVariant6Type = "saas"
	AccessAppNewParamsVariant6TypeSSH         AccessAppNewParamsVariant6Type = "ssh"
	AccessAppNewParamsVariant6TypeVnc         AccessAppNewParamsVariant6Type = "vnc"
	AccessAppNewParamsVariant6TypeAppLauncher AccessAppNewParamsVariant6Type = "app_launcher"
	AccessAppNewParamsVariant6TypeWarp        AccessAppNewParamsVariant6Type = "warp"
	AccessAppNewParamsVariant6TypeBiso        AccessAppNewParamsVariant6Type = "biso"
	AccessAppNewParamsVariant6TypeBookmark    AccessAppNewParamsVariant6Type = "bookmark"
	AccessAppNewParamsVariant6TypeDashSSO     AccessAppNewParamsVariant6Type = "dash_sso"
)

type AccessAppNewParamsVariant7 struct {
	AppLauncherVisible param.Field[interface{}] `json:"app_launcher_visible"`
	// The URL or domain of the bookmark.
	Domain param.Field[interface{}] `json:"domain"`
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

func (r AccessAppNewParamsVariant7) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppNewParamsVariant7) ImplementsAccessAppNewParams() {

}

type AccessAppNewResponseEnvelope struct {
	Errors   []AccessAppNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessAppNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessAppNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessAppNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessAppNewResponseEnvelopeJSON    `json:"-"`
}

// accessAppNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessAppNewResponseEnvelope]
type accessAppNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accessAppNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AccessAppNewResponseEnvelopeErrors]
type accessAppNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessAppNewResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessAppNewResponseEnvelopeMessages]
type accessAppNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppNewResponseEnvelopeSuccess bool

const (
	AccessAppNewResponseEnvelopeSuccessTrue AccessAppNewResponseEnvelopeSuccess = true
)

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppGetParamsAppID interface {
	ImplementsAccessAppGetParamsAppID()
}

type AccessAppGetResponseEnvelope struct {
	Errors   []AccessAppGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessAppGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessAppGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessAppGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessAppGetResponseEnvelopeJSON    `json:"-"`
}

// accessAppGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessAppGetResponseEnvelope]
type accessAppGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppGetResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accessAppGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AccessAppGetResponseEnvelopeErrors]
type accessAppGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppGetResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessAppGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessAppGetResponseEnvelopeMessages]
type accessAppGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppGetResponseEnvelopeSuccess bool

const (
	AccessAppGetResponseEnvelopeSuccessTrue AccessAppGetResponseEnvelopeSuccess = true
)

// This interface is a union satisfied by one of the following:
// [AccessAppUpdateParamsVariant0], [AccessAppUpdateParamsVariant1],
// [AccessAppUpdateParamsVariant2], [AccessAppUpdateParamsVariant3],
// [AccessAppUpdateParamsVariant4], [AccessAppUpdateParamsVariant5],
// [AccessAppUpdateParamsVariant6], [AccessAppUpdateParamsVariant7].
type AccessAppUpdateParams interface {
	ImplementsAccessAppUpdateParams()
}

type AccessAppUpdateParamsVariant0 struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWarp param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                     `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessAppUpdateParamsVariant0CorsHeaders] `json:"cors_headers"`
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

func (r AccessAppUpdateParamsVariant0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppUpdateParamsVariant0) ImplementsAccessAppUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppUpdateParamsVariant0AppID interface {
	ImplementsAccessAppUpdateParamsVariant0AppID()
}

type AccessAppUpdateParamsVariant0CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessAppUpdateParamsVariant0CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessAppUpdateParamsVariant0CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppUpdateParamsVariant0CorsHeadersAllowedMethod string

const (
	AccessAppUpdateParamsVariant0CorsHeadersAllowedMethodGet     AccessAppUpdateParamsVariant0CorsHeadersAllowedMethod = "GET"
	AccessAppUpdateParamsVariant0CorsHeadersAllowedMethodPost    AccessAppUpdateParamsVariant0CorsHeadersAllowedMethod = "POST"
	AccessAppUpdateParamsVariant0CorsHeadersAllowedMethodHead    AccessAppUpdateParamsVariant0CorsHeadersAllowedMethod = "HEAD"
	AccessAppUpdateParamsVariant0CorsHeadersAllowedMethodPut     AccessAppUpdateParamsVariant0CorsHeadersAllowedMethod = "PUT"
	AccessAppUpdateParamsVariant0CorsHeadersAllowedMethodDelete  AccessAppUpdateParamsVariant0CorsHeadersAllowedMethod = "DELETE"
	AccessAppUpdateParamsVariant0CorsHeadersAllowedMethodConnect AccessAppUpdateParamsVariant0CorsHeadersAllowedMethod = "CONNECT"
	AccessAppUpdateParamsVariant0CorsHeadersAllowedMethodOptions AccessAppUpdateParamsVariant0CorsHeadersAllowedMethod = "OPTIONS"
	AccessAppUpdateParamsVariant0CorsHeadersAllowedMethodTrace   AccessAppUpdateParamsVariant0CorsHeadersAllowedMethod = "TRACE"
	AccessAppUpdateParamsVariant0CorsHeadersAllowedMethodPatch   AccessAppUpdateParamsVariant0CorsHeadersAllowedMethod = "PATCH"
)

type AccessAppUpdateParamsVariant1 struct {
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
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
	Name    param.Field[string]                               `json:"name"`
	SaasApp param.Field[AccessAppUpdateParamsVariant1SaasApp] `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessAppUpdateParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppUpdateParamsVariant1) ImplementsAccessAppUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppUpdateParamsVariant1AppID interface {
	ImplementsAccessAppUpdateParamsVariant1AppID()
}

type AccessAppUpdateParamsVariant1SaasApp struct {
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                                               `json:"consumer_service_url"`
	CustomAttributes   param.Field[AccessAppUpdateParamsVariant1SaasAppCustomAttributes] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[AccessAppUpdateParamsVariant1SaasAppNameIDFormat] `json:"name_id_format"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// A globally unique name for an identity or service provider.
	SpEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string] `json:"sso_endpoint"`
}

func (r AccessAppUpdateParamsVariant1SaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppUpdateParamsVariant1SaasAppCustomAttributes struct {
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[AccessAppUpdateParamsVariant1SaasAppCustomAttributesNameFormat] `json:"name_format"`
	Source     param.Field[AccessAppUpdateParamsVariant1SaasAppCustomAttributesSource]     `json:"source"`
}

func (r AccessAppUpdateParamsVariant1SaasAppCustomAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A globally unique name for an identity or service provider.
type AccessAppUpdateParamsVariant1SaasAppCustomAttributesNameFormat string

const (
	AccessAppUpdateParamsVariant1SaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessAppUpdateParamsVariant1SaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessAppUpdateParamsVariant1SaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessAppUpdateParamsVariant1SaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessAppUpdateParamsVariant1SaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUri         AccessAppUpdateParamsVariant1SaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type AccessAppUpdateParamsVariant1SaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
}

func (r AccessAppUpdateParamsVariant1SaasAppCustomAttributesSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the name identifier sent to the SaaS application.
type AccessAppUpdateParamsVariant1SaasAppNameIDFormat string

const (
	AccessAppUpdateParamsVariant1SaasAppNameIDFormatID    AccessAppUpdateParamsVariant1SaasAppNameIDFormat = "id"
	AccessAppUpdateParamsVariant1SaasAppNameIDFormatEmail AccessAppUpdateParamsVariant1SaasAppNameIDFormat = "email"
)

type AccessAppUpdateParamsVariant2 struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWarp param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                     `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessAppUpdateParamsVariant2CorsHeaders] `json:"cors_headers"`
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

func (r AccessAppUpdateParamsVariant2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppUpdateParamsVariant2) ImplementsAccessAppUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppUpdateParamsVariant2AppID interface {
	ImplementsAccessAppUpdateParamsVariant2AppID()
}

type AccessAppUpdateParamsVariant2CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessAppUpdateParamsVariant2CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessAppUpdateParamsVariant2CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppUpdateParamsVariant2CorsHeadersAllowedMethod string

const (
	AccessAppUpdateParamsVariant2CorsHeadersAllowedMethodGet     AccessAppUpdateParamsVariant2CorsHeadersAllowedMethod = "GET"
	AccessAppUpdateParamsVariant2CorsHeadersAllowedMethodPost    AccessAppUpdateParamsVariant2CorsHeadersAllowedMethod = "POST"
	AccessAppUpdateParamsVariant2CorsHeadersAllowedMethodHead    AccessAppUpdateParamsVariant2CorsHeadersAllowedMethod = "HEAD"
	AccessAppUpdateParamsVariant2CorsHeadersAllowedMethodPut     AccessAppUpdateParamsVariant2CorsHeadersAllowedMethod = "PUT"
	AccessAppUpdateParamsVariant2CorsHeadersAllowedMethodDelete  AccessAppUpdateParamsVariant2CorsHeadersAllowedMethod = "DELETE"
	AccessAppUpdateParamsVariant2CorsHeadersAllowedMethodConnect AccessAppUpdateParamsVariant2CorsHeadersAllowedMethod = "CONNECT"
	AccessAppUpdateParamsVariant2CorsHeadersAllowedMethodOptions AccessAppUpdateParamsVariant2CorsHeadersAllowedMethod = "OPTIONS"
	AccessAppUpdateParamsVariant2CorsHeadersAllowedMethodTrace   AccessAppUpdateParamsVariant2CorsHeadersAllowedMethod = "TRACE"
	AccessAppUpdateParamsVariant2CorsHeadersAllowedMethodPatch   AccessAppUpdateParamsVariant2CorsHeadersAllowedMethod = "PATCH"
)

type AccessAppUpdateParamsVariant3 struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWarp param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                     `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessAppUpdateParamsVariant3CorsHeaders] `json:"cors_headers"`
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

func (r AccessAppUpdateParamsVariant3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppUpdateParamsVariant3) ImplementsAccessAppUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppUpdateParamsVariant3AppID interface {
	ImplementsAccessAppUpdateParamsVariant3AppID()
}

type AccessAppUpdateParamsVariant3CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessAppUpdateParamsVariant3CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessAppUpdateParamsVariant3CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppUpdateParamsVariant3CorsHeadersAllowedMethod string

const (
	AccessAppUpdateParamsVariant3CorsHeadersAllowedMethodGet     AccessAppUpdateParamsVariant3CorsHeadersAllowedMethod = "GET"
	AccessAppUpdateParamsVariant3CorsHeadersAllowedMethodPost    AccessAppUpdateParamsVariant3CorsHeadersAllowedMethod = "POST"
	AccessAppUpdateParamsVariant3CorsHeadersAllowedMethodHead    AccessAppUpdateParamsVariant3CorsHeadersAllowedMethod = "HEAD"
	AccessAppUpdateParamsVariant3CorsHeadersAllowedMethodPut     AccessAppUpdateParamsVariant3CorsHeadersAllowedMethod = "PUT"
	AccessAppUpdateParamsVariant3CorsHeadersAllowedMethodDelete  AccessAppUpdateParamsVariant3CorsHeadersAllowedMethod = "DELETE"
	AccessAppUpdateParamsVariant3CorsHeadersAllowedMethodConnect AccessAppUpdateParamsVariant3CorsHeadersAllowedMethod = "CONNECT"
	AccessAppUpdateParamsVariant3CorsHeadersAllowedMethodOptions AccessAppUpdateParamsVariant3CorsHeadersAllowedMethod = "OPTIONS"
	AccessAppUpdateParamsVariant3CorsHeadersAllowedMethodTrace   AccessAppUpdateParamsVariant3CorsHeadersAllowedMethod = "TRACE"
	AccessAppUpdateParamsVariant3CorsHeadersAllowedMethodPatch   AccessAppUpdateParamsVariant3CorsHeadersAllowedMethod = "PATCH"
)

type AccessAppUpdateParamsVariant4 struct {
	// The application type.
	Type param.Field[AccessAppUpdateParamsVariant4Type] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessAppUpdateParamsVariant4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppUpdateParamsVariant4) ImplementsAccessAppUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppUpdateParamsVariant4AppID interface {
	ImplementsAccessAppUpdateParamsVariant4AppID()
}

// The application type.
type AccessAppUpdateParamsVariant4Type string

const (
	AccessAppUpdateParamsVariant4TypeSelfHosted  AccessAppUpdateParamsVariant4Type = "self_hosted"
	AccessAppUpdateParamsVariant4TypeSaas        AccessAppUpdateParamsVariant4Type = "saas"
	AccessAppUpdateParamsVariant4TypeSSH         AccessAppUpdateParamsVariant4Type = "ssh"
	AccessAppUpdateParamsVariant4TypeVnc         AccessAppUpdateParamsVariant4Type = "vnc"
	AccessAppUpdateParamsVariant4TypeAppLauncher AccessAppUpdateParamsVariant4Type = "app_launcher"
	AccessAppUpdateParamsVariant4TypeWarp        AccessAppUpdateParamsVariant4Type = "warp"
	AccessAppUpdateParamsVariant4TypeBiso        AccessAppUpdateParamsVariant4Type = "biso"
	AccessAppUpdateParamsVariant4TypeBookmark    AccessAppUpdateParamsVariant4Type = "bookmark"
	AccessAppUpdateParamsVariant4TypeDashSSO     AccessAppUpdateParamsVariant4Type = "dash_sso"
)

type AccessAppUpdateParamsVariant5 struct {
	// The application type.
	Type param.Field[AccessAppUpdateParamsVariant5Type] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessAppUpdateParamsVariant5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppUpdateParamsVariant5) ImplementsAccessAppUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppUpdateParamsVariant5AppID interface {
	ImplementsAccessAppUpdateParamsVariant5AppID()
}

// The application type.
type AccessAppUpdateParamsVariant5Type string

const (
	AccessAppUpdateParamsVariant5TypeSelfHosted  AccessAppUpdateParamsVariant5Type = "self_hosted"
	AccessAppUpdateParamsVariant5TypeSaas        AccessAppUpdateParamsVariant5Type = "saas"
	AccessAppUpdateParamsVariant5TypeSSH         AccessAppUpdateParamsVariant5Type = "ssh"
	AccessAppUpdateParamsVariant5TypeVnc         AccessAppUpdateParamsVariant5Type = "vnc"
	AccessAppUpdateParamsVariant5TypeAppLauncher AccessAppUpdateParamsVariant5Type = "app_launcher"
	AccessAppUpdateParamsVariant5TypeWarp        AccessAppUpdateParamsVariant5Type = "warp"
	AccessAppUpdateParamsVariant5TypeBiso        AccessAppUpdateParamsVariant5Type = "biso"
	AccessAppUpdateParamsVariant5TypeBookmark    AccessAppUpdateParamsVariant5Type = "bookmark"
	AccessAppUpdateParamsVariant5TypeDashSSO     AccessAppUpdateParamsVariant5Type = "dash_sso"
)

type AccessAppUpdateParamsVariant6 struct {
	// The application type.
	Type param.Field[AccessAppUpdateParamsVariant6Type] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessAppUpdateParamsVariant6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppUpdateParamsVariant6) ImplementsAccessAppUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppUpdateParamsVariant6AppID interface {
	ImplementsAccessAppUpdateParamsVariant6AppID()
}

// The application type.
type AccessAppUpdateParamsVariant6Type string

const (
	AccessAppUpdateParamsVariant6TypeSelfHosted  AccessAppUpdateParamsVariant6Type = "self_hosted"
	AccessAppUpdateParamsVariant6TypeSaas        AccessAppUpdateParamsVariant6Type = "saas"
	AccessAppUpdateParamsVariant6TypeSSH         AccessAppUpdateParamsVariant6Type = "ssh"
	AccessAppUpdateParamsVariant6TypeVnc         AccessAppUpdateParamsVariant6Type = "vnc"
	AccessAppUpdateParamsVariant6TypeAppLauncher AccessAppUpdateParamsVariant6Type = "app_launcher"
	AccessAppUpdateParamsVariant6TypeWarp        AccessAppUpdateParamsVariant6Type = "warp"
	AccessAppUpdateParamsVariant6TypeBiso        AccessAppUpdateParamsVariant6Type = "biso"
	AccessAppUpdateParamsVariant6TypeBookmark    AccessAppUpdateParamsVariant6Type = "bookmark"
	AccessAppUpdateParamsVariant6TypeDashSSO     AccessAppUpdateParamsVariant6Type = "dash_sso"
)

type AccessAppUpdateParamsVariant7 struct {
	AppLauncherVisible param.Field[interface{}] `json:"app_launcher_visible"`
	// The URL or domain of the bookmark.
	Domain param.Field[interface{}] `json:"domain"`
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

func (r AccessAppUpdateParamsVariant7) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppUpdateParamsVariant7) ImplementsAccessAppUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppUpdateParamsVariant7AppID interface {
	ImplementsAccessAppUpdateParamsVariant7AppID()
}

type AccessAppUpdateResponseEnvelope struct {
	Errors   []AccessAppUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessAppUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessAppUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessAppUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessAppUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessAppUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessAppUpdateResponseEnvelope]
type accessAppUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUpdateResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessAppUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessAppUpdateResponseEnvelopeErrors]
type accessAppUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUpdateResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessAppUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessAppUpdateResponseEnvelopeMessages]
type accessAppUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppUpdateResponseEnvelopeSuccess bool

const (
	AccessAppUpdateResponseEnvelopeSuccessTrue AccessAppUpdateResponseEnvelopeSuccess = true
)

type AccessAppListResponseEnvelope struct {
	Errors   []AccessAppListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessAppListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessAppListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessAppListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessAppListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessAppListResponseEnvelopeJSON       `json:"-"`
}

// accessAppListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessAppListResponseEnvelope]
type accessAppListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppListResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accessAppListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessAppListResponseEnvelopeErrors]
type accessAppListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppListResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessAppListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessAppListResponseEnvelopeMessages]
type accessAppListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppListResponseEnvelopeSuccess bool

const (
	AccessAppListResponseEnvelopeSuccessTrue AccessAppListResponseEnvelopeSuccess = true
)

type AccessAppListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       accessAppListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessAppListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [AccessAppListResponseEnvelopeResultInfo]
type accessAppListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppDeleteParamsAppID interface {
	ImplementsAccessAppDeleteParamsAppID()
}

type AccessAppDeleteResponseEnvelope struct {
	Errors   []AccessAppDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessAppDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessAppDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessAppDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessAppDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessAppDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessAppDeleteResponseEnvelope]
type accessAppDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppDeleteResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessAppDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessAppDeleteResponseEnvelopeErrors]
type accessAppDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppDeleteResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessAppDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessAppDeleteResponseEnvelopeMessages]
type accessAppDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppDeleteResponseEnvelopeSuccess bool

const (
	AccessAppDeleteResponseEnvelopeSuccessTrue AccessAppDeleteResponseEnvelopeSuccess = true
)

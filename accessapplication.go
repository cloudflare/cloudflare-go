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

// AccessApplicationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessApplicationService] method
// instead.
type AccessApplicationService struct {
	Options          []option.RequestOption
	Cas              *AccessApplicationCaService
	UserPolicyChecks *AccessApplicationUserPolicyCheckService
	Policies         *AccessApplicationPolicyService
}

// NewAccessApplicationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessApplicationService(opts ...option.RequestOption) (r *AccessApplicationService) {
	r = &AccessApplicationService{}
	r.Options = opts
	r.Cas = NewAccessApplicationCaService(opts...)
	r.UserPolicyChecks = NewAccessApplicationUserPolicyCheckService(opts...)
	r.Policies = NewAccessApplicationPolicyService(opts...)
	return
}

// Adds a new application to Access.
func (r *AccessApplicationService) New(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessApplicationNewParams, opts ...option.RequestOption) (res *AccessApplicationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationNewResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an Access application.
func (r *AccessApplicationService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, appID AccessApplicationUpdateParamsVariant0AppID, body AccessApplicationUpdateParams, opts ...option.RequestOption) (res *AccessApplicationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationUpdateResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%v", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Access applications in an account or zone.
func (r *AccessApplicationService) List(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *[]AccessApplicationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationListResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an application from Access.
func (r *AccessApplicationService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, appID AccessApplicationDeleteParamsAppID, opts ...option.RequestOption) (res *AccessApplicationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationDeleteResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%v", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches information about an Access application.
func (r *AccessApplicationService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, appID AccessApplicationGetParamsAppID, opts ...option.RequestOption) (res *AccessApplicationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationGetResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%v", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Revokes all tokens issued for an application.
func (r *AccessApplicationService) RevokeTokens(ctx context.Context, accountOrZone string, accountOrZoneID string, appID AccessApplicationRevokeTokensParamsAppID, opts ...option.RequestOption) (res *AccessApplicationRevokeTokensResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationRevokeTokensResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%v/revoke_tokens", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AccessApplicationNewResponseObject],
// [AccessApplicationNewResponseObject], [AccessApplicationNewResponseObject],
// [AccessApplicationNewResponseObject], [AccessApplicationNewResponseObject],
// [AccessApplicationNewResponseObject], [AccessApplicationNewResponseObject] or
// [AccessApplicationNewResponseObject].
type AccessApplicationNewResponse interface {
	implementsAccessApplicationNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationNewResponse)(nil)).Elem(), "")
}

type AccessApplicationNewResponseObject struct {
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
	AutoRedirectToIdentity bool                                          `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationNewResponseObjectCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                     `json:"created_at" format:"date-time"`
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
	Tags      []string                               `json:"tags"`
	UpdatedAt time.Time                              `json:"updated_at" format:"date-time"`
	JSON      accessApplicationNewResponseObjectJSON `json:"-"`
}

// accessApplicationNewResponseObjectJSON contains the JSON metadata for the struct
// [AccessApplicationNewResponseObject]
type accessApplicationNewResponseObjectJSON struct {
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

func (r *AccessApplicationNewResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationNewResponseObject) implementsAccessApplicationNewResponse() {}

type AccessApplicationNewResponseObjectCorsHeaders struct {
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
	AllowedMethods []AccessApplicationNewResponseObjectCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                           `json:"max_age"`
	JSON   accessApplicationNewResponseObjectCorsHeadersJSON `json:"-"`
}

// accessApplicationNewResponseObjectCorsHeadersJSON contains the JSON metadata for
// the struct [AccessApplicationNewResponseObjectCorsHeaders]
type accessApplicationNewResponseObjectCorsHeadersJSON struct {
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

func (r *AccessApplicationNewResponseObjectCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationNewResponseObjectCorsHeadersAllowedMethod string

const (
	AccessApplicationNewResponseObjectCorsHeadersAllowedMethodGet     AccessApplicationNewResponseObjectCorsHeadersAllowedMethod = "GET"
	AccessApplicationNewResponseObjectCorsHeadersAllowedMethodPost    AccessApplicationNewResponseObjectCorsHeadersAllowedMethod = "POST"
	AccessApplicationNewResponseObjectCorsHeadersAllowedMethodHead    AccessApplicationNewResponseObjectCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationNewResponseObjectCorsHeadersAllowedMethodPut     AccessApplicationNewResponseObjectCorsHeadersAllowedMethod = "PUT"
	AccessApplicationNewResponseObjectCorsHeadersAllowedMethodDelete  AccessApplicationNewResponseObjectCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationNewResponseObjectCorsHeadersAllowedMethodConnect AccessApplicationNewResponseObjectCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationNewResponseObjectCorsHeadersAllowedMethodOptions AccessApplicationNewResponseObjectCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationNewResponseObjectCorsHeadersAllowedMethodTrace   AccessApplicationNewResponseObjectCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationNewResponseObjectCorsHeadersAllowedMethodPatch   AccessApplicationNewResponseObjectCorsHeadersAllowedMethod = "PATCH"
)

// Union satisfied by [AccessApplicationUpdateResponseObject],
// [AccessApplicationUpdateResponseObject],
// [AccessApplicationUpdateResponseObject],
// [AccessApplicationUpdateResponseObject],
// [AccessApplicationUpdateResponseObject],
// [AccessApplicationUpdateResponseObject], [AccessApplicationUpdateResponseObject]
// or [AccessApplicationUpdateResponseObject].
type AccessApplicationUpdateResponse interface {
	implementsAccessApplicationUpdateResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationUpdateResponse)(nil)).Elem(), "")
}

type AccessApplicationUpdateResponseObject struct {
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
	AutoRedirectToIdentity bool                                             `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationUpdateResponseObjectCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                        `json:"created_at" format:"date-time"`
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
	Tags      []string                                  `json:"tags"`
	UpdatedAt time.Time                                 `json:"updated_at" format:"date-time"`
	JSON      accessApplicationUpdateResponseObjectJSON `json:"-"`
}

// accessApplicationUpdateResponseObjectJSON contains the JSON metadata for the
// struct [AccessApplicationUpdateResponseObject]
type accessApplicationUpdateResponseObjectJSON struct {
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

func (r *AccessApplicationUpdateResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationUpdateResponseObject) implementsAccessApplicationUpdateResponse() {}

type AccessApplicationUpdateResponseObjectCorsHeaders struct {
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
	AllowedMethods []AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                              `json:"max_age"`
	JSON   accessApplicationUpdateResponseObjectCorsHeadersJSON `json:"-"`
}

// accessApplicationUpdateResponseObjectCorsHeadersJSON contains the JSON metadata
// for the struct [AccessApplicationUpdateResponseObjectCorsHeaders]
type accessApplicationUpdateResponseObjectCorsHeadersJSON struct {
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

func (r *AccessApplicationUpdateResponseObjectCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethod string

const (
	AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethodGet     AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethod = "GET"
	AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethodPost    AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethod = "POST"
	AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethodHead    AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethodPut     AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethod = "PUT"
	AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethodDelete  AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethodConnect AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethodOptions AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethodTrace   AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethodPatch   AccessApplicationUpdateResponseObjectCorsHeadersAllowedMethod = "PATCH"
)

// Union satisfied by [AccessApplicationListResponseObject],
// [AccessApplicationListResponseObject], [AccessApplicationListResponseObject],
// [AccessApplicationListResponseObject], [AccessApplicationListResponseObject],
// [AccessApplicationListResponseObject], [AccessApplicationListResponseObject] or
// [AccessApplicationListResponseObject].
type AccessApplicationListResponse interface {
	implementsAccessApplicationListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationListResponse)(nil)).Elem(), "")
}

type AccessApplicationListResponseObject struct {
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
	AutoRedirectToIdentity bool                                           `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationListResponseObjectCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                      `json:"created_at" format:"date-time"`
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
	Tags      []string                                `json:"tags"`
	UpdatedAt time.Time                               `json:"updated_at" format:"date-time"`
	JSON      accessApplicationListResponseObjectJSON `json:"-"`
}

// accessApplicationListResponseObjectJSON contains the JSON metadata for the
// struct [AccessApplicationListResponseObject]
type accessApplicationListResponseObjectJSON struct {
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

func (r *AccessApplicationListResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationListResponseObject) implementsAccessApplicationListResponse() {}

type AccessApplicationListResponseObjectCorsHeaders struct {
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
	AllowedMethods []AccessApplicationListResponseObjectCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                            `json:"max_age"`
	JSON   accessApplicationListResponseObjectCorsHeadersJSON `json:"-"`
}

// accessApplicationListResponseObjectCorsHeadersJSON contains the JSON metadata
// for the struct [AccessApplicationListResponseObjectCorsHeaders]
type accessApplicationListResponseObjectCorsHeadersJSON struct {
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

func (r *AccessApplicationListResponseObjectCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationListResponseObjectCorsHeadersAllowedMethod string

const (
	AccessApplicationListResponseObjectCorsHeadersAllowedMethodGet     AccessApplicationListResponseObjectCorsHeadersAllowedMethod = "GET"
	AccessApplicationListResponseObjectCorsHeadersAllowedMethodPost    AccessApplicationListResponseObjectCorsHeadersAllowedMethod = "POST"
	AccessApplicationListResponseObjectCorsHeadersAllowedMethodHead    AccessApplicationListResponseObjectCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationListResponseObjectCorsHeadersAllowedMethodPut     AccessApplicationListResponseObjectCorsHeadersAllowedMethod = "PUT"
	AccessApplicationListResponseObjectCorsHeadersAllowedMethodDelete  AccessApplicationListResponseObjectCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationListResponseObjectCorsHeadersAllowedMethodConnect AccessApplicationListResponseObjectCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationListResponseObjectCorsHeadersAllowedMethodOptions AccessApplicationListResponseObjectCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationListResponseObjectCorsHeadersAllowedMethodTrace   AccessApplicationListResponseObjectCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationListResponseObjectCorsHeadersAllowedMethodPatch   AccessApplicationListResponseObjectCorsHeadersAllowedMethod = "PATCH"
)

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

// Union satisfied by [AccessApplicationGetResponseObject],
// [AccessApplicationGetResponseObject], [AccessApplicationGetResponseObject],
// [AccessApplicationGetResponseObject], [AccessApplicationGetResponseObject],
// [AccessApplicationGetResponseObject], [AccessApplicationGetResponseObject] or
// [AccessApplicationGetResponseObject].
type AccessApplicationGetResponse interface {
	implementsAccessApplicationGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationGetResponse)(nil)).Elem(), "")
}

type AccessApplicationGetResponseObject struct {
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
	AutoRedirectToIdentity bool                                          `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationGetResponseObjectCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                     `json:"created_at" format:"date-time"`
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
	Tags      []string                               `json:"tags"`
	UpdatedAt time.Time                              `json:"updated_at" format:"date-time"`
	JSON      accessApplicationGetResponseObjectJSON `json:"-"`
}

// accessApplicationGetResponseObjectJSON contains the JSON metadata for the struct
// [AccessApplicationGetResponseObject]
type accessApplicationGetResponseObjectJSON struct {
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

func (r *AccessApplicationGetResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationGetResponseObject) implementsAccessApplicationGetResponse() {}

type AccessApplicationGetResponseObjectCorsHeaders struct {
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
	AllowedMethods []AccessApplicationGetResponseObjectCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                           `json:"max_age"`
	JSON   accessApplicationGetResponseObjectCorsHeadersJSON `json:"-"`
}

// accessApplicationGetResponseObjectCorsHeadersJSON contains the JSON metadata for
// the struct [AccessApplicationGetResponseObjectCorsHeaders]
type accessApplicationGetResponseObjectCorsHeadersJSON struct {
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

func (r *AccessApplicationGetResponseObjectCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationGetResponseObjectCorsHeadersAllowedMethod string

const (
	AccessApplicationGetResponseObjectCorsHeadersAllowedMethodGet     AccessApplicationGetResponseObjectCorsHeadersAllowedMethod = "GET"
	AccessApplicationGetResponseObjectCorsHeadersAllowedMethodPost    AccessApplicationGetResponseObjectCorsHeadersAllowedMethod = "POST"
	AccessApplicationGetResponseObjectCorsHeadersAllowedMethodHead    AccessApplicationGetResponseObjectCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationGetResponseObjectCorsHeadersAllowedMethodPut     AccessApplicationGetResponseObjectCorsHeadersAllowedMethod = "PUT"
	AccessApplicationGetResponseObjectCorsHeadersAllowedMethodDelete  AccessApplicationGetResponseObjectCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationGetResponseObjectCorsHeadersAllowedMethodConnect AccessApplicationGetResponseObjectCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationGetResponseObjectCorsHeadersAllowedMethodOptions AccessApplicationGetResponseObjectCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationGetResponseObjectCorsHeadersAllowedMethodTrace   AccessApplicationGetResponseObjectCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationGetResponseObjectCorsHeadersAllowedMethodPatch   AccessApplicationGetResponseObjectCorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationRevokeTokensResponse = interface{}

// This interface is a union satisfied by one of the following:
// [AccessApplicationNewParamsVariant0], [AccessApplicationNewParamsVariant1],
// [AccessApplicationNewParamsVariant2], [AccessApplicationNewParamsVariant3],
// [AccessApplicationNewParamsVariant4], [AccessApplicationNewParamsVariant5],
// [AccessApplicationNewParamsVariant6], [AccessApplicationNewParamsVariant7].
type AccessApplicationNewParams interface {
	ImplementsAccessApplicationNewParams()
}

type AccessApplicationNewParamsVariant0 struct {
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
	AutoRedirectToIdentity param.Field[bool]                                          `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationNewParamsVariant0CorsHeaders] `json:"cors_headers"`
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

func (r AccessApplicationNewParamsVariant0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationNewParamsVariant0) ImplementsAccessApplicationNewParams() {

}

type AccessApplicationNewParamsVariant0CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationNewParamsVariant0CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationNewParamsVariant0CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationNewParamsVariant0CorsHeadersAllowedMethod string

const (
	AccessApplicationNewParamsVariant0CorsHeadersAllowedMethodGet     AccessApplicationNewParamsVariant0CorsHeadersAllowedMethod = "GET"
	AccessApplicationNewParamsVariant0CorsHeadersAllowedMethodPost    AccessApplicationNewParamsVariant0CorsHeadersAllowedMethod = "POST"
	AccessApplicationNewParamsVariant0CorsHeadersAllowedMethodHead    AccessApplicationNewParamsVariant0CorsHeadersAllowedMethod = "HEAD"
	AccessApplicationNewParamsVariant0CorsHeadersAllowedMethodPut     AccessApplicationNewParamsVariant0CorsHeadersAllowedMethod = "PUT"
	AccessApplicationNewParamsVariant0CorsHeadersAllowedMethodDelete  AccessApplicationNewParamsVariant0CorsHeadersAllowedMethod = "DELETE"
	AccessApplicationNewParamsVariant0CorsHeadersAllowedMethodConnect AccessApplicationNewParamsVariant0CorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationNewParamsVariant0CorsHeadersAllowedMethodOptions AccessApplicationNewParamsVariant0CorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationNewParamsVariant0CorsHeadersAllowedMethodTrace   AccessApplicationNewParamsVariant0CorsHeadersAllowedMethod = "TRACE"
	AccessApplicationNewParamsVariant0CorsHeadersAllowedMethodPatch   AccessApplicationNewParamsVariant0CorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationNewParamsVariant1 struct {
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
	Name    param.Field[string]                                    `json:"name"`
	SaasApp param.Field[AccessApplicationNewParamsVariant1SaasApp] `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessApplicationNewParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationNewParamsVariant1) ImplementsAccessApplicationNewParams() {

}

// Satisfied by [AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasApp],
// [AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasApp].
type AccessApplicationNewParamsVariant1SaasApp interface {
	implementsAccessApplicationNewParamsVariant1SaasApp()
}

type AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                                                                     `json:"consumer_service_url"`
	CustomAttributes   param.Field[AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppCustomAttributes] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppNameIDFormat] `json:"name_id_format"`
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

func (r AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasApp) implementsAccessApplicationNewParamsVariant1SaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppAuthType string

const (
	AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppAuthTypeSaml AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppAuthType = "saml"
	AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppAuthTypeOidc AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppAuthType = "oidc"
)

type AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormat] `json:"name_format"`
	Source     param.Field[AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesSource]     `json:"source"`
}

func (r AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppCustomAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A globally unique name for an identity or service provider.
type AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUri         AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
}

func (r AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the name identifier sent to the SaaS application.
type AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppNameIDFormat string

const (
	AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppNameIDFormatID    AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppNameIDFormat = "id"
	AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppNameIDFormatEmail AccessApplicationNewParamsVariant1SaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType param.Field[AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppAuthType] `json:"auth_type"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string] `json:"client_secret"`
	// The OIDC flows supported by this application
	GrantTypes param.Field[[]AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppGrantType] `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex param.Field[string] `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectUris param.Field[[]string] `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes param.Field[[]AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppScope] `json:"scopes"`
}

func (r AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasApp) implementsAccessApplicationNewParamsVariant1SaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppAuthType string

const (
	AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppAuthTypeSaml AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppAuthType = "saml"
	AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppAuthTypeOidc AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppAuthType = "oidc"
)

type AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppGrantType string

const (
	AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppGrantType = "authorization_code"
	AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppScope string

const (
	AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppScopeOpenid  AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppScope = "openid"
	AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppScopeGroups  AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppScope = "groups"
	AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppScopeEmail   AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppScope = "email"
	AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppScopeProfile AccessApplicationNewParamsVariant1SaasAppAccessOidcSaasAppScope = "profile"
)

type AccessApplicationNewParamsVariant2 struct {
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
	AutoRedirectToIdentity param.Field[bool]                                          `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationNewParamsVariant2CorsHeaders] `json:"cors_headers"`
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

func (r AccessApplicationNewParamsVariant2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationNewParamsVariant2) ImplementsAccessApplicationNewParams() {

}

type AccessApplicationNewParamsVariant2CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationNewParamsVariant2CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationNewParamsVariant2CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationNewParamsVariant2CorsHeadersAllowedMethod string

const (
	AccessApplicationNewParamsVariant2CorsHeadersAllowedMethodGet     AccessApplicationNewParamsVariant2CorsHeadersAllowedMethod = "GET"
	AccessApplicationNewParamsVariant2CorsHeadersAllowedMethodPost    AccessApplicationNewParamsVariant2CorsHeadersAllowedMethod = "POST"
	AccessApplicationNewParamsVariant2CorsHeadersAllowedMethodHead    AccessApplicationNewParamsVariant2CorsHeadersAllowedMethod = "HEAD"
	AccessApplicationNewParamsVariant2CorsHeadersAllowedMethodPut     AccessApplicationNewParamsVariant2CorsHeadersAllowedMethod = "PUT"
	AccessApplicationNewParamsVariant2CorsHeadersAllowedMethodDelete  AccessApplicationNewParamsVariant2CorsHeadersAllowedMethod = "DELETE"
	AccessApplicationNewParamsVariant2CorsHeadersAllowedMethodConnect AccessApplicationNewParamsVariant2CorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationNewParamsVariant2CorsHeadersAllowedMethodOptions AccessApplicationNewParamsVariant2CorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationNewParamsVariant2CorsHeadersAllowedMethodTrace   AccessApplicationNewParamsVariant2CorsHeadersAllowedMethod = "TRACE"
	AccessApplicationNewParamsVariant2CorsHeadersAllowedMethodPatch   AccessApplicationNewParamsVariant2CorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationNewParamsVariant3 struct {
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
	AutoRedirectToIdentity param.Field[bool]                                          `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationNewParamsVariant3CorsHeaders] `json:"cors_headers"`
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

func (r AccessApplicationNewParamsVariant3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationNewParamsVariant3) ImplementsAccessApplicationNewParams() {

}

type AccessApplicationNewParamsVariant3CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationNewParamsVariant3CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationNewParamsVariant3CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationNewParamsVariant3CorsHeadersAllowedMethod string

const (
	AccessApplicationNewParamsVariant3CorsHeadersAllowedMethodGet     AccessApplicationNewParamsVariant3CorsHeadersAllowedMethod = "GET"
	AccessApplicationNewParamsVariant3CorsHeadersAllowedMethodPost    AccessApplicationNewParamsVariant3CorsHeadersAllowedMethod = "POST"
	AccessApplicationNewParamsVariant3CorsHeadersAllowedMethodHead    AccessApplicationNewParamsVariant3CorsHeadersAllowedMethod = "HEAD"
	AccessApplicationNewParamsVariant3CorsHeadersAllowedMethodPut     AccessApplicationNewParamsVariant3CorsHeadersAllowedMethod = "PUT"
	AccessApplicationNewParamsVariant3CorsHeadersAllowedMethodDelete  AccessApplicationNewParamsVariant3CorsHeadersAllowedMethod = "DELETE"
	AccessApplicationNewParamsVariant3CorsHeadersAllowedMethodConnect AccessApplicationNewParamsVariant3CorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationNewParamsVariant3CorsHeadersAllowedMethodOptions AccessApplicationNewParamsVariant3CorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationNewParamsVariant3CorsHeadersAllowedMethodTrace   AccessApplicationNewParamsVariant3CorsHeadersAllowedMethod = "TRACE"
	AccessApplicationNewParamsVariant3CorsHeadersAllowedMethodPatch   AccessApplicationNewParamsVariant3CorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationNewParamsVariant4 struct {
	// The application type.
	Type param.Field[AccessApplicationNewParamsVariant4Type] `json:"type,required"`
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

func (r AccessApplicationNewParamsVariant4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationNewParamsVariant4) ImplementsAccessApplicationNewParams() {

}

// The application type.
type AccessApplicationNewParamsVariant4Type string

const (
	AccessApplicationNewParamsVariant4TypeSelfHosted  AccessApplicationNewParamsVariant4Type = "self_hosted"
	AccessApplicationNewParamsVariant4TypeSaas        AccessApplicationNewParamsVariant4Type = "saas"
	AccessApplicationNewParamsVariant4TypeSSH         AccessApplicationNewParamsVariant4Type = "ssh"
	AccessApplicationNewParamsVariant4TypeVnc         AccessApplicationNewParamsVariant4Type = "vnc"
	AccessApplicationNewParamsVariant4TypeAppLauncher AccessApplicationNewParamsVariant4Type = "app_launcher"
	AccessApplicationNewParamsVariant4TypeWarp        AccessApplicationNewParamsVariant4Type = "warp"
	AccessApplicationNewParamsVariant4TypeBiso        AccessApplicationNewParamsVariant4Type = "biso"
	AccessApplicationNewParamsVariant4TypeBookmark    AccessApplicationNewParamsVariant4Type = "bookmark"
	AccessApplicationNewParamsVariant4TypeDashSSO     AccessApplicationNewParamsVariant4Type = "dash_sso"
)

type AccessApplicationNewParamsVariant5 struct {
	// The application type.
	Type param.Field[AccessApplicationNewParamsVariant5Type] `json:"type,required"`
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

func (r AccessApplicationNewParamsVariant5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationNewParamsVariant5) ImplementsAccessApplicationNewParams() {

}

// The application type.
type AccessApplicationNewParamsVariant5Type string

const (
	AccessApplicationNewParamsVariant5TypeSelfHosted  AccessApplicationNewParamsVariant5Type = "self_hosted"
	AccessApplicationNewParamsVariant5TypeSaas        AccessApplicationNewParamsVariant5Type = "saas"
	AccessApplicationNewParamsVariant5TypeSSH         AccessApplicationNewParamsVariant5Type = "ssh"
	AccessApplicationNewParamsVariant5TypeVnc         AccessApplicationNewParamsVariant5Type = "vnc"
	AccessApplicationNewParamsVariant5TypeAppLauncher AccessApplicationNewParamsVariant5Type = "app_launcher"
	AccessApplicationNewParamsVariant5TypeWarp        AccessApplicationNewParamsVariant5Type = "warp"
	AccessApplicationNewParamsVariant5TypeBiso        AccessApplicationNewParamsVariant5Type = "biso"
	AccessApplicationNewParamsVariant5TypeBookmark    AccessApplicationNewParamsVariant5Type = "bookmark"
	AccessApplicationNewParamsVariant5TypeDashSSO     AccessApplicationNewParamsVariant5Type = "dash_sso"
)

type AccessApplicationNewParamsVariant6 struct {
	// The application type.
	Type param.Field[AccessApplicationNewParamsVariant6Type] `json:"type,required"`
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

func (r AccessApplicationNewParamsVariant6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationNewParamsVariant6) ImplementsAccessApplicationNewParams() {

}

// The application type.
type AccessApplicationNewParamsVariant6Type string

const (
	AccessApplicationNewParamsVariant6TypeSelfHosted  AccessApplicationNewParamsVariant6Type = "self_hosted"
	AccessApplicationNewParamsVariant6TypeSaas        AccessApplicationNewParamsVariant6Type = "saas"
	AccessApplicationNewParamsVariant6TypeSSH         AccessApplicationNewParamsVariant6Type = "ssh"
	AccessApplicationNewParamsVariant6TypeVnc         AccessApplicationNewParamsVariant6Type = "vnc"
	AccessApplicationNewParamsVariant6TypeAppLauncher AccessApplicationNewParamsVariant6Type = "app_launcher"
	AccessApplicationNewParamsVariant6TypeWarp        AccessApplicationNewParamsVariant6Type = "warp"
	AccessApplicationNewParamsVariant6TypeBiso        AccessApplicationNewParamsVariant6Type = "biso"
	AccessApplicationNewParamsVariant6TypeBookmark    AccessApplicationNewParamsVariant6Type = "bookmark"
	AccessApplicationNewParamsVariant6TypeDashSSO     AccessApplicationNewParamsVariant6Type = "dash_sso"
)

type AccessApplicationNewParamsVariant7 struct {
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

func (r AccessApplicationNewParamsVariant7) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationNewParamsVariant7) ImplementsAccessApplicationNewParams() {

}

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

// Whether the API call was successful
type AccessApplicationNewResponseEnvelopeSuccess bool

const (
	AccessApplicationNewResponseEnvelopeSuccessTrue AccessApplicationNewResponseEnvelopeSuccess = true
)

// This interface is a union satisfied by one of the following:
// [AccessApplicationUpdateParamsVariant0],
// [AccessApplicationUpdateParamsVariant1],
// [AccessApplicationUpdateParamsVariant2],
// [AccessApplicationUpdateParamsVariant3],
// [AccessApplicationUpdateParamsVariant4],
// [AccessApplicationUpdateParamsVariant5],
// [AccessApplicationUpdateParamsVariant6],
// [AccessApplicationUpdateParamsVariant7].
type AccessApplicationUpdateParams interface {
	ImplementsAccessApplicationUpdateParams()
}

type AccessApplicationUpdateParamsVariant0 struct {
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
	AutoRedirectToIdentity param.Field[bool]                                             `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationUpdateParamsVariant0CorsHeaders] `json:"cors_headers"`
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

func (r AccessApplicationUpdateParamsVariant0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationUpdateParamsVariant0) ImplementsAccessApplicationUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsVariant0AppID interface {
	ImplementsAccessApplicationUpdateParamsVariant0AppID()
}

type AccessApplicationUpdateParamsVariant0CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationUpdateParamsVariant0CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethod string

const (
	AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethodGet     AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethod = "GET"
	AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethodPost    AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethod = "POST"
	AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethodHead    AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethod = "HEAD"
	AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethodPut     AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethod = "PUT"
	AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethodDelete  AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethod = "DELETE"
	AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethodConnect AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethodOptions AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethodTrace   AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethod = "TRACE"
	AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethodPatch   AccessApplicationUpdateParamsVariant0CorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationUpdateParamsVariant1 struct {
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
	Name    param.Field[string]                                       `json:"name"`
	SaasApp param.Field[AccessApplicationUpdateParamsVariant1SaasApp] `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessApplicationUpdateParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationUpdateParamsVariant1) ImplementsAccessApplicationUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsVariant1AppID interface {
	ImplementsAccessApplicationUpdateParamsVariant1AppID()
}

// Satisfied by [AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasApp],
// [AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasApp].
type AccessApplicationUpdateParamsVariant1SaasApp interface {
	implementsAccessApplicationUpdateParamsVariant1SaasApp()
}

type AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                                                                        `json:"consumer_service_url"`
	CustomAttributes   param.Field[AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppCustomAttributes] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppNameIDFormat] `json:"name_id_format"`
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

func (r AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasApp) implementsAccessApplicationUpdateParamsVariant1SaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppAuthType string

const (
	AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppAuthTypeSaml AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppAuthType = "saml"
	AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppAuthTypeOidc AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppAuthType = "oidc"
)

type AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormat] `json:"name_format"`
	Source     param.Field[AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesSource]     `json:"source"`
}

func (r AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppCustomAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A globally unique name for an identity or service provider.
type AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUri         AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
}

func (r AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the name identifier sent to the SaaS application.
type AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppNameIDFormat string

const (
	AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppNameIDFormatID    AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppNameIDFormat = "id"
	AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppNameIDFormatEmail AccessApplicationUpdateParamsVariant1SaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType param.Field[AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppAuthType] `json:"auth_type"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string] `json:"client_secret"`
	// The OIDC flows supported by this application
	GrantTypes param.Field[[]AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppGrantType] `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex param.Field[string] `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectUris param.Field[[]string] `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes param.Field[[]AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppScope] `json:"scopes"`
}

func (r AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasApp) implementsAccessApplicationUpdateParamsVariant1SaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppAuthType string

const (
	AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppAuthTypeSaml AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppAuthType = "saml"
	AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppAuthTypeOidc AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppAuthType = "oidc"
)

type AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppGrantType string

const (
	AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppGrantType = "authorization_code"
	AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppScope string

const (
	AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppScopeOpenid  AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppScope = "openid"
	AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppScopeGroups  AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppScope = "groups"
	AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppScopeEmail   AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppScope = "email"
	AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppScopeProfile AccessApplicationUpdateParamsVariant1SaasAppAccessOidcSaasAppScope = "profile"
)

type AccessApplicationUpdateParamsVariant2 struct {
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
	AutoRedirectToIdentity param.Field[bool]                                             `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationUpdateParamsVariant2CorsHeaders] `json:"cors_headers"`
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

func (r AccessApplicationUpdateParamsVariant2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationUpdateParamsVariant2) ImplementsAccessApplicationUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsVariant2AppID interface {
	ImplementsAccessApplicationUpdateParamsVariant2AppID()
}

type AccessApplicationUpdateParamsVariant2CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationUpdateParamsVariant2CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethod string

const (
	AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethodGet     AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethod = "GET"
	AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethodPost    AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethod = "POST"
	AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethodHead    AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethod = "HEAD"
	AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethodPut     AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethod = "PUT"
	AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethodDelete  AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethod = "DELETE"
	AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethodConnect AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethodOptions AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethodTrace   AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethod = "TRACE"
	AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethodPatch   AccessApplicationUpdateParamsVariant2CorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationUpdateParamsVariant3 struct {
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
	AutoRedirectToIdentity param.Field[bool]                                             `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationUpdateParamsVariant3CorsHeaders] `json:"cors_headers"`
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

func (r AccessApplicationUpdateParamsVariant3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationUpdateParamsVariant3) ImplementsAccessApplicationUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsVariant3AppID interface {
	ImplementsAccessApplicationUpdateParamsVariant3AppID()
}

type AccessApplicationUpdateParamsVariant3CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationUpdateParamsVariant3CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethod string

const (
	AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethodGet     AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethod = "GET"
	AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethodPost    AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethod = "POST"
	AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethodHead    AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethod = "HEAD"
	AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethodPut     AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethod = "PUT"
	AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethodDelete  AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethod = "DELETE"
	AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethodConnect AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethodOptions AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethodTrace   AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethod = "TRACE"
	AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethodPatch   AccessApplicationUpdateParamsVariant3CorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationUpdateParamsVariant4 struct {
	// The application type.
	Type param.Field[AccessApplicationUpdateParamsVariant4Type] `json:"type,required"`
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

func (r AccessApplicationUpdateParamsVariant4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationUpdateParamsVariant4) ImplementsAccessApplicationUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsVariant4AppID interface {
	ImplementsAccessApplicationUpdateParamsVariant4AppID()
}

// The application type.
type AccessApplicationUpdateParamsVariant4Type string

const (
	AccessApplicationUpdateParamsVariant4TypeSelfHosted  AccessApplicationUpdateParamsVariant4Type = "self_hosted"
	AccessApplicationUpdateParamsVariant4TypeSaas        AccessApplicationUpdateParamsVariant4Type = "saas"
	AccessApplicationUpdateParamsVariant4TypeSSH         AccessApplicationUpdateParamsVariant4Type = "ssh"
	AccessApplicationUpdateParamsVariant4TypeVnc         AccessApplicationUpdateParamsVariant4Type = "vnc"
	AccessApplicationUpdateParamsVariant4TypeAppLauncher AccessApplicationUpdateParamsVariant4Type = "app_launcher"
	AccessApplicationUpdateParamsVariant4TypeWarp        AccessApplicationUpdateParamsVariant4Type = "warp"
	AccessApplicationUpdateParamsVariant4TypeBiso        AccessApplicationUpdateParamsVariant4Type = "biso"
	AccessApplicationUpdateParamsVariant4TypeBookmark    AccessApplicationUpdateParamsVariant4Type = "bookmark"
	AccessApplicationUpdateParamsVariant4TypeDashSSO     AccessApplicationUpdateParamsVariant4Type = "dash_sso"
)

type AccessApplicationUpdateParamsVariant5 struct {
	// The application type.
	Type param.Field[AccessApplicationUpdateParamsVariant5Type] `json:"type,required"`
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

func (r AccessApplicationUpdateParamsVariant5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationUpdateParamsVariant5) ImplementsAccessApplicationUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsVariant5AppID interface {
	ImplementsAccessApplicationUpdateParamsVariant5AppID()
}

// The application type.
type AccessApplicationUpdateParamsVariant5Type string

const (
	AccessApplicationUpdateParamsVariant5TypeSelfHosted  AccessApplicationUpdateParamsVariant5Type = "self_hosted"
	AccessApplicationUpdateParamsVariant5TypeSaas        AccessApplicationUpdateParamsVariant5Type = "saas"
	AccessApplicationUpdateParamsVariant5TypeSSH         AccessApplicationUpdateParamsVariant5Type = "ssh"
	AccessApplicationUpdateParamsVariant5TypeVnc         AccessApplicationUpdateParamsVariant5Type = "vnc"
	AccessApplicationUpdateParamsVariant5TypeAppLauncher AccessApplicationUpdateParamsVariant5Type = "app_launcher"
	AccessApplicationUpdateParamsVariant5TypeWarp        AccessApplicationUpdateParamsVariant5Type = "warp"
	AccessApplicationUpdateParamsVariant5TypeBiso        AccessApplicationUpdateParamsVariant5Type = "biso"
	AccessApplicationUpdateParamsVariant5TypeBookmark    AccessApplicationUpdateParamsVariant5Type = "bookmark"
	AccessApplicationUpdateParamsVariant5TypeDashSSO     AccessApplicationUpdateParamsVariant5Type = "dash_sso"
)

type AccessApplicationUpdateParamsVariant6 struct {
	// The application type.
	Type param.Field[AccessApplicationUpdateParamsVariant6Type] `json:"type,required"`
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

func (r AccessApplicationUpdateParamsVariant6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationUpdateParamsVariant6) ImplementsAccessApplicationUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsVariant6AppID interface {
	ImplementsAccessApplicationUpdateParamsVariant6AppID()
}

// The application type.
type AccessApplicationUpdateParamsVariant6Type string

const (
	AccessApplicationUpdateParamsVariant6TypeSelfHosted  AccessApplicationUpdateParamsVariant6Type = "self_hosted"
	AccessApplicationUpdateParamsVariant6TypeSaas        AccessApplicationUpdateParamsVariant6Type = "saas"
	AccessApplicationUpdateParamsVariant6TypeSSH         AccessApplicationUpdateParamsVariant6Type = "ssh"
	AccessApplicationUpdateParamsVariant6TypeVnc         AccessApplicationUpdateParamsVariant6Type = "vnc"
	AccessApplicationUpdateParamsVariant6TypeAppLauncher AccessApplicationUpdateParamsVariant6Type = "app_launcher"
	AccessApplicationUpdateParamsVariant6TypeWarp        AccessApplicationUpdateParamsVariant6Type = "warp"
	AccessApplicationUpdateParamsVariant6TypeBiso        AccessApplicationUpdateParamsVariant6Type = "biso"
	AccessApplicationUpdateParamsVariant6TypeBookmark    AccessApplicationUpdateParamsVariant6Type = "bookmark"
	AccessApplicationUpdateParamsVariant6TypeDashSSO     AccessApplicationUpdateParamsVariant6Type = "dash_sso"
)

type AccessApplicationUpdateParamsVariant7 struct {
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

func (r AccessApplicationUpdateParamsVariant7) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationUpdateParamsVariant7) ImplementsAccessApplicationUpdateParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationUpdateParamsVariant7AppID interface {
	ImplementsAccessApplicationUpdateParamsVariant7AppID()
}

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

// Whether the API call was successful
type AccessApplicationUpdateResponseEnvelopeSuccess bool

const (
	AccessApplicationUpdateResponseEnvelopeSuccessTrue AccessApplicationUpdateResponseEnvelopeSuccess = true
)

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

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationDeleteParamsAppID interface {
	ImplementsAccessApplicationDeleteParamsAppID()
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

// Whether the API call was successful
type AccessApplicationDeleteResponseEnvelopeSuccess bool

const (
	AccessApplicationDeleteResponseEnvelopeSuccessTrue AccessApplicationDeleteResponseEnvelopeSuccess = true
)

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationGetParamsAppID interface {
	ImplementsAccessApplicationGetParamsAppID()
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

// Whether the API call was successful
type AccessApplicationGetResponseEnvelopeSuccess bool

const (
	AccessApplicationGetResponseEnvelopeSuccessTrue AccessApplicationGetResponseEnvelopeSuccess = true
)

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationRevokeTokensParamsAppID interface {
	ImplementsAccessApplicationRevokeTokensParamsAppID()
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

type AccessApplicationRevokeTokensResponseEnvelopeSuccess bool

const (
	AccessApplicationRevokeTokensResponseEnvelopeSuccessTrue  AccessApplicationRevokeTokensResponseEnvelopeSuccess = true
	AccessApplicationRevokeTokensResponseEnvelopeSuccessFalse AccessApplicationRevokeTokensResponseEnvelopeSuccess = false
)

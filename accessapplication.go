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

// Updates an Access application.
func (r *AccessApplicationService) Replace(ctx context.Context, accountOrZone string, accountOrZoneID string, appID AccessApplicationReplaceParamsVariant0AppID, body AccessApplicationReplaceParams, opts ...option.RequestOption) (res *AccessApplicationReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationReplaceResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%v", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
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

// Union satisfied by [AccessApplicationReplaceResponseObject],
// [AccessApplicationReplaceResponseObject],
// [AccessApplicationReplaceResponseObject],
// [AccessApplicationReplaceResponseObject],
// [AccessApplicationReplaceResponseObject],
// [AccessApplicationReplaceResponseObject],
// [AccessApplicationReplaceResponseObject] or
// [AccessApplicationReplaceResponseObject].
type AccessApplicationReplaceResponse interface {
	implementsAccessApplicationReplaceResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationReplaceResponse)(nil)).Elem(), "")
}

type AccessApplicationReplaceResponseObject struct {
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
	AutoRedirectToIdentity bool                                              `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessApplicationReplaceResponseObjectCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                         `json:"created_at" format:"date-time"`
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
	Tags      []string                                   `json:"tags"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      accessApplicationReplaceResponseObjectJSON `json:"-"`
}

// accessApplicationReplaceResponseObjectJSON contains the JSON metadata for the
// struct [AccessApplicationReplaceResponseObject]
type accessApplicationReplaceResponseObjectJSON struct {
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

func (r *AccessApplicationReplaceResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationReplaceResponseObject) implementsAccessApplicationReplaceResponse() {}

type AccessApplicationReplaceResponseObjectCorsHeaders struct {
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
	AllowedMethods []AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                               `json:"max_age"`
	JSON   accessApplicationReplaceResponseObjectCorsHeadersJSON `json:"-"`
}

// accessApplicationReplaceResponseObjectCorsHeadersJSON contains the JSON metadata
// for the struct [AccessApplicationReplaceResponseObjectCorsHeaders]
type accessApplicationReplaceResponseObjectCorsHeadersJSON struct {
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

func (r *AccessApplicationReplaceResponseObjectCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethod string

const (
	AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethodGet     AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethod = "GET"
	AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethodPost    AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethod = "POST"
	AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethodHead    AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethod = "HEAD"
	AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethodPut     AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethod = "PUT"
	AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethodDelete  AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethod = "DELETE"
	AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethodConnect AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethodOptions AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethodTrace   AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethod = "TRACE"
	AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethodPatch   AccessApplicationReplaceResponseObjectCorsHeadersAllowedMethod = "PATCH"
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

// This interface is a union satisfied by one of the following:
// [AccessApplicationReplaceParamsVariant0],
// [AccessApplicationReplaceParamsVariant1],
// [AccessApplicationReplaceParamsVariant2],
// [AccessApplicationReplaceParamsVariant3],
// [AccessApplicationReplaceParamsVariant4],
// [AccessApplicationReplaceParamsVariant5],
// [AccessApplicationReplaceParamsVariant6],
// [AccessApplicationReplaceParamsVariant7].
type AccessApplicationReplaceParams interface {
	ImplementsAccessApplicationReplaceParams()
}

type AccessApplicationReplaceParamsVariant0 struct {
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
	AutoRedirectToIdentity param.Field[bool]                                              `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationReplaceParamsVariant0CorsHeaders] `json:"cors_headers"`
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

func (r AccessApplicationReplaceParamsVariant0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationReplaceParamsVariant0) ImplementsAccessApplicationReplaceParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationReplaceParamsVariant0AppID interface {
	ImplementsAccessApplicationReplaceParamsVariant0AppID()
}

type AccessApplicationReplaceParamsVariant0CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationReplaceParamsVariant0CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethod string

const (
	AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethodGet     AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethod = "GET"
	AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethodPost    AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethod = "POST"
	AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethodHead    AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethod = "HEAD"
	AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethodPut     AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethod = "PUT"
	AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethodDelete  AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethod = "DELETE"
	AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethodConnect AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethodOptions AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethodTrace   AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethod = "TRACE"
	AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethodPatch   AccessApplicationReplaceParamsVariant0CorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationReplaceParamsVariant1 struct {
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
	Name    param.Field[string]                                        `json:"name"`
	SaasApp param.Field[AccessApplicationReplaceParamsVariant1SaasApp] `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessApplicationReplaceParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationReplaceParamsVariant1) ImplementsAccessApplicationReplaceParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationReplaceParamsVariant1AppID interface {
	ImplementsAccessApplicationReplaceParamsVariant1AppID()
}

// Satisfied by [AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasApp],
// [AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasApp].
type AccessApplicationReplaceParamsVariant1SaasApp interface {
	implementsAccessApplicationReplaceParamsVariant1SaasApp()
}

type AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                                                                         `json:"consumer_service_url"`
	CustomAttributes   param.Field[AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppCustomAttributes] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppNameIDFormat] `json:"name_id_format"`
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

func (r AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasApp) implementsAccessApplicationReplaceParamsVariant1SaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppAuthType string

const (
	AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppAuthTypeSaml AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppAuthType = "saml"
	AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppAuthTypeOidc AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppAuthType = "oidc"
)

type AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppCustomAttributes struct {
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormat] `json:"name_format"`
	Source     param.Field[AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesSource]     `json:"source"`
}

func (r AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppCustomAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A globally unique name for an identity or service provider.
type AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormat string

const (
	AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUri         AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
}

func (r AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppCustomAttributesSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the name identifier sent to the SaaS application.
type AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppNameIDFormat string

const (
	AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppNameIDFormatID    AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppNameIDFormat = "id"
	AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppNameIDFormatEmail AccessApplicationReplaceParamsVariant1SaasAppAccessSamlSaasAppNameIDFormat = "email"
)

type AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasApp struct {
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType param.Field[AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppAuthType] `json:"auth_type"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string] `json:"client_secret"`
	// The OIDC flows supported by this application
	GrantTypes param.Field[[]AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppGrantType] `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex param.Field[string] `json:"group_filter_regex"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectUris param.Field[[]string] `json:"redirect_uris"`
	// Define the user information shared with access
	Scopes param.Field[[]AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppScope] `json:"scopes"`
}

func (r AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasApp) implementsAccessApplicationReplaceParamsVariant1SaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppAuthType string

const (
	AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppAuthTypeSaml AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppAuthType = "saml"
	AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppAuthTypeOidc AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppAuthType = "oidc"
)

type AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppGrantType string

const (
	AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppGrantType = "authorization_code"
	AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
)

type AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppScope string

const (
	AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppScopeOpenid  AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppScope = "openid"
	AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppScopeGroups  AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppScope = "groups"
	AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppScopeEmail   AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppScope = "email"
	AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppScopeProfile AccessApplicationReplaceParamsVariant1SaasAppAccessOidcSaasAppScope = "profile"
)

type AccessApplicationReplaceParamsVariant2 struct {
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
	AutoRedirectToIdentity param.Field[bool]                                              `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationReplaceParamsVariant2CorsHeaders] `json:"cors_headers"`
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

func (r AccessApplicationReplaceParamsVariant2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationReplaceParamsVariant2) ImplementsAccessApplicationReplaceParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationReplaceParamsVariant2AppID interface {
	ImplementsAccessApplicationReplaceParamsVariant2AppID()
}

type AccessApplicationReplaceParamsVariant2CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationReplaceParamsVariant2CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethod string

const (
	AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethodGet     AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethod = "GET"
	AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethodPost    AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethod = "POST"
	AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethodHead    AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethod = "HEAD"
	AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethodPut     AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethod = "PUT"
	AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethodDelete  AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethod = "DELETE"
	AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethodConnect AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethodOptions AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethodTrace   AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethod = "TRACE"
	AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethodPatch   AccessApplicationReplaceParamsVariant2CorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationReplaceParamsVariant3 struct {
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
	AutoRedirectToIdentity param.Field[bool]                                              `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessApplicationReplaceParamsVariant3CorsHeaders] `json:"cors_headers"`
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

func (r AccessApplicationReplaceParamsVariant3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationReplaceParamsVariant3) ImplementsAccessApplicationReplaceParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationReplaceParamsVariant3AppID interface {
	ImplementsAccessApplicationReplaceParamsVariant3AppID()
}

type AccessApplicationReplaceParamsVariant3CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessApplicationReplaceParamsVariant3CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethod string

const (
	AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethodGet     AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethod = "GET"
	AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethodPost    AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethod = "POST"
	AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethodHead    AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethod = "HEAD"
	AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethodPut     AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethod = "PUT"
	AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethodDelete  AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethod = "DELETE"
	AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethodConnect AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethod = "CONNECT"
	AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethodOptions AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethod = "OPTIONS"
	AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethodTrace   AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethod = "TRACE"
	AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethodPatch   AccessApplicationReplaceParamsVariant3CorsHeadersAllowedMethod = "PATCH"
)

type AccessApplicationReplaceParamsVariant4 struct {
	// The application type.
	Type param.Field[AccessApplicationReplaceParamsVariant4Type] `json:"type,required"`
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

func (r AccessApplicationReplaceParamsVariant4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationReplaceParamsVariant4) ImplementsAccessApplicationReplaceParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationReplaceParamsVariant4AppID interface {
	ImplementsAccessApplicationReplaceParamsVariant4AppID()
}

// The application type.
type AccessApplicationReplaceParamsVariant4Type string

const (
	AccessApplicationReplaceParamsVariant4TypeSelfHosted  AccessApplicationReplaceParamsVariant4Type = "self_hosted"
	AccessApplicationReplaceParamsVariant4TypeSaas        AccessApplicationReplaceParamsVariant4Type = "saas"
	AccessApplicationReplaceParamsVariant4TypeSSH         AccessApplicationReplaceParamsVariant4Type = "ssh"
	AccessApplicationReplaceParamsVariant4TypeVnc         AccessApplicationReplaceParamsVariant4Type = "vnc"
	AccessApplicationReplaceParamsVariant4TypeAppLauncher AccessApplicationReplaceParamsVariant4Type = "app_launcher"
	AccessApplicationReplaceParamsVariant4TypeWarp        AccessApplicationReplaceParamsVariant4Type = "warp"
	AccessApplicationReplaceParamsVariant4TypeBiso        AccessApplicationReplaceParamsVariant4Type = "biso"
	AccessApplicationReplaceParamsVariant4TypeBookmark    AccessApplicationReplaceParamsVariant4Type = "bookmark"
	AccessApplicationReplaceParamsVariant4TypeDashSSO     AccessApplicationReplaceParamsVariant4Type = "dash_sso"
)

type AccessApplicationReplaceParamsVariant5 struct {
	// The application type.
	Type param.Field[AccessApplicationReplaceParamsVariant5Type] `json:"type,required"`
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

func (r AccessApplicationReplaceParamsVariant5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationReplaceParamsVariant5) ImplementsAccessApplicationReplaceParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationReplaceParamsVariant5AppID interface {
	ImplementsAccessApplicationReplaceParamsVariant5AppID()
}

// The application type.
type AccessApplicationReplaceParamsVariant5Type string

const (
	AccessApplicationReplaceParamsVariant5TypeSelfHosted  AccessApplicationReplaceParamsVariant5Type = "self_hosted"
	AccessApplicationReplaceParamsVariant5TypeSaas        AccessApplicationReplaceParamsVariant5Type = "saas"
	AccessApplicationReplaceParamsVariant5TypeSSH         AccessApplicationReplaceParamsVariant5Type = "ssh"
	AccessApplicationReplaceParamsVariant5TypeVnc         AccessApplicationReplaceParamsVariant5Type = "vnc"
	AccessApplicationReplaceParamsVariant5TypeAppLauncher AccessApplicationReplaceParamsVariant5Type = "app_launcher"
	AccessApplicationReplaceParamsVariant5TypeWarp        AccessApplicationReplaceParamsVariant5Type = "warp"
	AccessApplicationReplaceParamsVariant5TypeBiso        AccessApplicationReplaceParamsVariant5Type = "biso"
	AccessApplicationReplaceParamsVariant5TypeBookmark    AccessApplicationReplaceParamsVariant5Type = "bookmark"
	AccessApplicationReplaceParamsVariant5TypeDashSSO     AccessApplicationReplaceParamsVariant5Type = "dash_sso"
)

type AccessApplicationReplaceParamsVariant6 struct {
	// The application type.
	Type param.Field[AccessApplicationReplaceParamsVariant6Type] `json:"type,required"`
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

func (r AccessApplicationReplaceParamsVariant6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationReplaceParamsVariant6) ImplementsAccessApplicationReplaceParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationReplaceParamsVariant6AppID interface {
	ImplementsAccessApplicationReplaceParamsVariant6AppID()
}

// The application type.
type AccessApplicationReplaceParamsVariant6Type string

const (
	AccessApplicationReplaceParamsVariant6TypeSelfHosted  AccessApplicationReplaceParamsVariant6Type = "self_hosted"
	AccessApplicationReplaceParamsVariant6TypeSaas        AccessApplicationReplaceParamsVariant6Type = "saas"
	AccessApplicationReplaceParamsVariant6TypeSSH         AccessApplicationReplaceParamsVariant6Type = "ssh"
	AccessApplicationReplaceParamsVariant6TypeVnc         AccessApplicationReplaceParamsVariant6Type = "vnc"
	AccessApplicationReplaceParamsVariant6TypeAppLauncher AccessApplicationReplaceParamsVariant6Type = "app_launcher"
	AccessApplicationReplaceParamsVariant6TypeWarp        AccessApplicationReplaceParamsVariant6Type = "warp"
	AccessApplicationReplaceParamsVariant6TypeBiso        AccessApplicationReplaceParamsVariant6Type = "biso"
	AccessApplicationReplaceParamsVariant6TypeBookmark    AccessApplicationReplaceParamsVariant6Type = "bookmark"
	AccessApplicationReplaceParamsVariant6TypeDashSSO     AccessApplicationReplaceParamsVariant6Type = "dash_sso"
)

type AccessApplicationReplaceParamsVariant7 struct {
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

func (r AccessApplicationReplaceParamsVariant7) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessApplicationReplaceParamsVariant7) ImplementsAccessApplicationReplaceParams() {

}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessApplicationReplaceParamsVariant7AppID interface {
	ImplementsAccessApplicationReplaceParamsVariant7AppID()
}

type AccessApplicationReplaceResponseEnvelope struct {
	Errors   []AccessApplicationReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationReplaceResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationReplaceResponseEnvelope]
type accessApplicationReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationReplaceResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accessApplicationReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationReplaceResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessApplicationReplaceResponseEnvelopeErrors]
type accessApplicationReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationReplaceResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accessApplicationReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationReplaceResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessApplicationReplaceResponseEnvelopeMessages]
type accessApplicationReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessApplicationReplaceResponseEnvelopeSuccess bool

const (
	AccessApplicationReplaceResponseEnvelopeSuccessTrue AccessApplicationReplaceResponseEnvelopeSuccess = true
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

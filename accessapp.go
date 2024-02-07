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

// Adds a new application to Access.
func (r *AccessAppService) AccessApplicationsAddAnApplication(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessAppAccessApplicationsAddAnApplicationParams, opts ...option.RequestOption) (res *AccessAppAccessApplicationsAddAnApplicationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppAccessApplicationsAddAnApplicationResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Access applications in an account or zone.
func (r *AccessAppService) AccessApplicationsListAccessApplications(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *[]AccessAppAccessApplicationsListAccessApplicationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppAccessApplicationsListAccessApplicationsResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

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
	Tags []string `json:"tags"`
	// The application type.
	Type      string                         `json:"type"`
	UpdatedAt time.Time                      `json:"updated_at" format:"date-time"`
	JSON      accessAppGetResponseObjectJSON `json:"-"`
}

// accessAppGetResponseObjectJSON contains the JSON metadata for the struct
// [AccessAppGetResponseObject]
type accessAppGetResponseObjectJSON struct {
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
	Tags []string `json:"tags"`
	// The application type.
	Type      string                            `json:"type"`
	UpdatedAt time.Time                         `json:"updated_at" format:"date-time"`
	JSON      accessAppUpdateResponseObjectJSON `json:"-"`
}

// accessAppUpdateResponseObjectJSON contains the JSON metadata for the struct
// [AccessAppUpdateResponseObject]
type accessAppUpdateResponseObjectJSON struct {
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

// Union satisfied by [AccessAppAccessApplicationsAddAnApplicationResponseObject],
// [AccessAppAccessApplicationsAddAnApplicationResponseObject],
// [AccessAppAccessApplicationsAddAnApplicationResponseObject],
// [AccessAppAccessApplicationsAddAnApplicationResponseObject],
// [AccessAppAccessApplicationsAddAnApplicationResponseObject],
// [AccessAppAccessApplicationsAddAnApplicationResponseObject],
// [AccessAppAccessApplicationsAddAnApplicationResponseObject] or
// [AccessAppAccessApplicationsAddAnApplicationResponseObject].
type AccessAppAccessApplicationsAddAnApplicationResponse interface {
	implementsAccessAppAccessApplicationsAddAnApplicationResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppAccessApplicationsAddAnApplicationResponse)(nil)).Elem(), "")
}

type AccessAppAccessApplicationsAddAnApplicationResponseObject struct {
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
	AutoRedirectToIdentity bool                                                                 `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                            `json:"created_at" format:"date-time"`
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
	Tags []string `json:"tags"`
	// The application type.
	Type      string                                                        `json:"type"`
	UpdatedAt time.Time                                                     `json:"updated_at" format:"date-time"`
	JSON      accessAppAccessApplicationsAddAnApplicationResponseObjectJSON `json:"-"`
}

// accessAppAccessApplicationsAddAnApplicationResponseObjectJSON contains the JSON
// metadata for the struct
// [AccessAppAccessApplicationsAddAnApplicationResponseObject]
type accessAppAccessApplicationsAddAnApplicationResponseObjectJSON struct {
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
	UpdatedAt                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsAddAnApplicationResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppAccessApplicationsAddAnApplicationResponseObject) implementsAccessAppAccessApplicationsAddAnApplicationResponse() {
}

type AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeaders struct {
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
	AllowedMethods []AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                  `json:"max_age"`
	JSON   accessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersJSON `json:"-"`
}

// accessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersJSON
// contains the JSON metadata for the struct
// [AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeaders]
type accessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersJSON struct {
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

func (r *AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethod string

const (
	AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethodGet     AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethod = "GET"
	AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethodPost    AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethod = "POST"
	AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethodHead    AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethod = "HEAD"
	AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethodPut     AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethod = "PUT"
	AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethodDelete  AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethod = "DELETE"
	AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethodConnect AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethod = "CONNECT"
	AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethodOptions AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethod = "OPTIONS"
	AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethodTrace   AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethod = "TRACE"
	AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethodPatch   AccessAppAccessApplicationsAddAnApplicationResponseObjectCorsHeadersAllowedMethod = "PATCH"
)

// Union satisfied by
// [AccessAppAccessApplicationsListAccessApplicationsResponseObject],
// [AccessAppAccessApplicationsListAccessApplicationsResponseObject],
// [AccessAppAccessApplicationsListAccessApplicationsResponseObject],
// [AccessAppAccessApplicationsListAccessApplicationsResponseObject],
// [AccessAppAccessApplicationsListAccessApplicationsResponseObject],
// [AccessAppAccessApplicationsListAccessApplicationsResponseObject],
// [AccessAppAccessApplicationsListAccessApplicationsResponseObject] or
// [AccessAppAccessApplicationsListAccessApplicationsResponseObject].
type AccessAppAccessApplicationsListAccessApplicationsResponse interface {
	implementsAccessAppAccessApplicationsListAccessApplicationsResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppAccessApplicationsListAccessApplicationsResponse)(nil)).Elem(), "")
}

type AccessAppAccessApplicationsListAccessApplicationsResponseObject struct {
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
	AutoRedirectToIdentity bool                                                                       `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                                  `json:"created_at" format:"date-time"`
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
	Tags []string `json:"tags"`
	// The application type.
	Type      string                                                              `json:"type"`
	UpdatedAt time.Time                                                           `json:"updated_at" format:"date-time"`
	JSON      accessAppAccessApplicationsListAccessApplicationsResponseObjectJSON `json:"-"`
}

// accessAppAccessApplicationsListAccessApplicationsResponseObjectJSON contains the
// JSON metadata for the struct
// [AccessAppAccessApplicationsListAccessApplicationsResponseObject]
type accessAppAccessApplicationsListAccessApplicationsResponseObjectJSON struct {
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
	UpdatedAt                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsListAccessApplicationsResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppAccessApplicationsListAccessApplicationsResponseObject) implementsAccessAppAccessApplicationsListAccessApplicationsResponse() {
}

type AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeaders struct {
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
	AllowedMethods []AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                        `json:"max_age"`
	JSON   accessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersJSON `json:"-"`
}

// accessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersJSON
// contains the JSON metadata for the struct
// [AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeaders]
type accessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersJSON struct {
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

func (r *AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethod string

const (
	AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethodGet     AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethod = "GET"
	AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethodPost    AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethod = "POST"
	AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethodHead    AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethod = "HEAD"
	AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethodPut     AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethod = "PUT"
	AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethodDelete  AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethod = "DELETE"
	AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethodConnect AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethod = "CONNECT"
	AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethodOptions AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethod = "OPTIONS"
	AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethodTrace   AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethod = "TRACE"
	AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethodPatch   AccessAppAccessApplicationsListAccessApplicationsResponseObjectCorsHeadersAllowedMethod = "PATCH"
)

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppGetParamsAppID interface {
	ImplementsAccessAppGetParamsAppID()
}

type AccessAppGetResponseEnvelope struct {
	Errors   []AccessAppGetResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessAppGetResponseEnvelopeMessages `json:"messages"`
	Result   AccessAppGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessAppGetResponseEnvelopeSuccess `json:"success"`
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
	// The application type.
	Type param.Field[string] `json:"type"`
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
	// The application type.
	Type param.Field[string] `json:"type"`
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
	// The application type.
	Type param.Field[string] `json:"type"`
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
	// The application type.
	Type param.Field[string] `json:"type"`
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

type AccessAppUpdateParamsVariant5 struct {
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
	// The application type.
	Type param.Field[string] `json:"type"`
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

type AccessAppUpdateParamsVariant6 struct {
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
	// The application type.
	Type param.Field[string] `json:"type"`
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
	Errors   []AccessAppUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessAppUpdateResponseEnvelopeMessages `json:"messages"`
	Result   AccessAppUpdateResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessAppUpdateResponseEnvelopeSuccess `json:"success"`
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

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppDeleteParamsAppID interface {
	ImplementsAccessAppDeleteParamsAppID()
}

type AccessAppDeleteResponseEnvelope struct {
	Errors   []AccessAppDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessAppDeleteResponseEnvelopeMessages `json:"messages"`
	Result   AccessAppDeleteResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessAppDeleteResponseEnvelopeSuccess `json:"success"`
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

// This interface is a union satisfied by one of the following:
// [AccessAppAccessApplicationsAddAnApplicationParamsVariant0],
// [AccessAppAccessApplicationsAddAnApplicationParamsVariant1],
// [AccessAppAccessApplicationsAddAnApplicationParamsVariant2],
// [AccessAppAccessApplicationsAddAnApplicationParamsVariant3],
// [AccessAppAccessApplicationsAddAnApplicationParamsVariant4],
// [AccessAppAccessApplicationsAddAnApplicationParamsVariant5],
// [AccessAppAccessApplicationsAddAnApplicationParamsVariant6],
// [AccessAppAccessApplicationsAddAnApplicationParamsVariant7].
type AccessAppAccessApplicationsAddAnApplicationParams interface {
	ImplementsAccessAppAccessApplicationsAddAnApplicationParams()
}

type AccessAppAccessApplicationsAddAnApplicationParamsVariant0 struct {
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
	AutoRedirectToIdentity param.Field[bool]                                                                 `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeaders] `json:"cors_headers"`
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
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessAppAccessApplicationsAddAnApplicationParamsVariant0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppAccessApplicationsAddAnApplicationParamsVariant0) ImplementsAccessAppAccessApplicationsAddAnApplicationParams() {

}

type AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethod string

const (
	AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethodGet     AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethod = "GET"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethodPost    AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethod = "POST"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethodHead    AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethod = "HEAD"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethodPut     AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethod = "PUT"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethodDelete  AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethod = "DELETE"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethodConnect AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethod = "CONNECT"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethodOptions AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethod = "OPTIONS"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethodTrace   AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethod = "TRACE"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethodPatch   AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethod = "PATCH"
)

type AccessAppAccessApplicationsAddAnApplicationParamsVariant1 struct {
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
	Name    param.Field[string]                                                           `json:"name"`
	SaasApp param.Field[AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasApp] `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessAppAccessApplicationsAddAnApplicationParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppAccessApplicationsAddAnApplicationParamsVariant1) ImplementsAccessAppAccessApplicationsAddAnApplicationParams() {

}

type AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasApp struct {
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                                                                           `json:"consumer_service_url"`
	CustomAttributes   param.Field[AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppCustomAttributes] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppNameIDFormat] `json:"name_id_format"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// A globally unique name for an identity or service provider.
	SpEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string] `json:"sso_endpoint"`
}

func (r AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppCustomAttributes struct {
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppCustomAttributesNameFormat] `json:"name_format"`
	Source     param.Field[AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppCustomAttributesSource]     `json:"source"`
}

func (r AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppCustomAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A globally unique name for an identity or service provider.
type AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppCustomAttributesNameFormat string

const (
	AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUnspecified AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic       AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatUri         AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

type AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
}

func (r AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppCustomAttributesSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the name identifier sent to the SaaS application.
type AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppNameIDFormat string

const (
	AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppNameIDFormatID    AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppNameIDFormat = "id"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppNameIDFormatEmail AccessAppAccessApplicationsAddAnApplicationParamsVariant1SaasAppNameIDFormat = "email"
)

type AccessAppAccessApplicationsAddAnApplicationParamsVariant2 struct {
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
	AutoRedirectToIdentity param.Field[bool]                                                                 `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeaders] `json:"cors_headers"`
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
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessAppAccessApplicationsAddAnApplicationParamsVariant2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppAccessApplicationsAddAnApplicationParamsVariant2) ImplementsAccessAppAccessApplicationsAddAnApplicationParams() {

}

type AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethod string

const (
	AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethodGet     AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethod = "GET"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethodPost    AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethod = "POST"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethodHead    AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethod = "HEAD"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethodPut     AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethod = "PUT"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethodDelete  AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethod = "DELETE"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethodConnect AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethod = "CONNECT"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethodOptions AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethod = "OPTIONS"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethodTrace   AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethod = "TRACE"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethodPatch   AccessAppAccessApplicationsAddAnApplicationParamsVariant2CorsHeadersAllowedMethod = "PATCH"
)

type AccessAppAccessApplicationsAddAnApplicationParamsVariant3 struct {
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
	AutoRedirectToIdentity param.Field[bool]                                                                 `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeaders] `json:"cors_headers"`
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
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessAppAccessApplicationsAddAnApplicationParamsVariant3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppAccessApplicationsAddAnApplicationParamsVariant3) ImplementsAccessAppAccessApplicationsAddAnApplicationParams() {

}

type AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeaders struct {
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
	AllowedMethods param.Field[[]AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethod] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethod string

const (
	AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethodGet     AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethod = "GET"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethodPost    AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethod = "POST"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethodHead    AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethod = "HEAD"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethodPut     AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethod = "PUT"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethodDelete  AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethod = "DELETE"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethodConnect AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethod = "CONNECT"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethodOptions AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethod = "OPTIONS"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethodTrace   AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethod = "TRACE"
	AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethodPatch   AccessAppAccessApplicationsAddAnApplicationParamsVariant3CorsHeadersAllowedMethod = "PATCH"
)

type AccessAppAccessApplicationsAddAnApplicationParamsVariant4 struct {
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
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessAppAccessApplicationsAddAnApplicationParamsVariant4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppAccessApplicationsAddAnApplicationParamsVariant4) ImplementsAccessAppAccessApplicationsAddAnApplicationParams() {

}

type AccessAppAccessApplicationsAddAnApplicationParamsVariant5 struct {
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
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessAppAccessApplicationsAddAnApplicationParamsVariant5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppAccessApplicationsAddAnApplicationParamsVariant5) ImplementsAccessAppAccessApplicationsAddAnApplicationParams() {

}

type AccessAppAccessApplicationsAddAnApplicationParamsVariant6 struct {
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
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessAppAccessApplicationsAddAnApplicationParamsVariant6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppAccessApplicationsAddAnApplicationParamsVariant6) ImplementsAccessAppAccessApplicationsAddAnApplicationParams() {

}

type AccessAppAccessApplicationsAddAnApplicationParamsVariant7 struct {
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

func (r AccessAppAccessApplicationsAddAnApplicationParamsVariant7) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessAppAccessApplicationsAddAnApplicationParamsVariant7) ImplementsAccessAppAccessApplicationsAddAnApplicationParams() {

}

type AccessAppAccessApplicationsAddAnApplicationResponseEnvelope struct {
	Errors   []AccessAppAccessApplicationsAddAnApplicationResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessAppAccessApplicationsAddAnApplicationResponseEnvelopeMessages `json:"messages"`
	Result   AccessAppAccessApplicationsAddAnApplicationResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessAppAccessApplicationsAddAnApplicationResponseEnvelopeSuccess `json:"success"`
	JSON    accessAppAccessApplicationsAddAnApplicationResponseEnvelopeJSON    `json:"-"`
}

// accessAppAccessApplicationsAddAnApplicationResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [AccessAppAccessApplicationsAddAnApplicationResponseEnvelope]
type accessAppAccessApplicationsAddAnApplicationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsAddAnApplicationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppAccessApplicationsAddAnApplicationResponseEnvelopeErrors struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accessAppAccessApplicationsAddAnApplicationResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppAccessApplicationsAddAnApplicationResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [AccessAppAccessApplicationsAddAnApplicationResponseEnvelopeErrors]
type accessAppAccessApplicationsAddAnApplicationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsAddAnApplicationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppAccessApplicationsAddAnApplicationResponseEnvelopeMessages struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    accessAppAccessApplicationsAddAnApplicationResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppAccessApplicationsAddAnApplicationResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [AccessAppAccessApplicationsAddAnApplicationResponseEnvelopeMessages]
type accessAppAccessApplicationsAddAnApplicationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsAddAnApplicationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppAccessApplicationsAddAnApplicationResponseEnvelopeSuccess bool

const (
	AccessAppAccessApplicationsAddAnApplicationResponseEnvelopeSuccessTrue AccessAppAccessApplicationsAddAnApplicationResponseEnvelopeSuccess = true
)

type AccessAppAccessApplicationsListAccessApplicationsResponseEnvelope struct {
	Errors     []AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeErrors   `json:"errors"`
	Messages   []AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeMessages `json:"messages"`
	Result     []AccessAppAccessApplicationsListAccessApplicationsResponse                 `json:"result"`
	ResultInfo AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeSuccess `json:"success"`
	JSON    accessAppAccessApplicationsListAccessApplicationsResponseEnvelopeJSON    `json:"-"`
}

// accessAppAccessApplicationsListAccessApplicationsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [AccessAppAccessApplicationsListAccessApplicationsResponseEnvelope]
type accessAppAccessApplicationsListAccessApplicationsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsListAccessApplicationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeErrors struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accessAppAccessApplicationsListAccessApplicationsResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppAccessApplicationsListAccessApplicationsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeErrors]
type accessAppAccessApplicationsListAccessApplicationsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeMessages struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    accessAppAccessApplicationsListAccessApplicationsResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppAccessApplicationsListAccessApplicationsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeMessages]
type accessAppAccessApplicationsListAccessApplicationsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                         `json:"total_count"`
	JSON       accessAppAccessApplicationsListAccessApplicationsResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessAppAccessApplicationsListAccessApplicationsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeResultInfo]
type accessAppAccessApplicationsListAccessApplicationsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeSuccess bool

const (
	AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeSuccessTrue AccessAppAccessApplicationsListAccessApplicationsResponseEnvelopeSuccess = true
)

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
	path := fmt.Sprintf("%s/%s/access/apps/%v", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an Access application.
func (r *AccessAppService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, appID AccessAppUpdateParamsVariant0AppID, body AccessAppUpdateParams, opts ...option.RequestOption) (res *AccessAppUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/apps/%v", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an application from Access.
func (r *AccessAppService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, appID AccessAppDeleteParamsAppID, opts ...option.RequestOption) (res *AccessAppDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/apps/%v", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new application to Access.
func (r *AccessAppService) AccessApplicationsAddAnApplication(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessAppAccessApplicationsAddAnApplicationParams, opts ...option.RequestOption) (res *AccessAppAccessApplicationsAddAnApplicationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/apps", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all Access applications in an account or zone.
func (r *AccessAppService) AccessApplicationsListAccessApplications(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *AccessAppAccessApplicationsListAccessApplicationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/apps", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessAppGetResponse struct {
	Errors   []AccessAppGetResponseError   `json:"errors"`
	Messages []AccessAppGetResponseMessage `json:"messages"`
	Result   AccessAppGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessAppGetResponseSuccess `json:"success"`
	JSON    accessAppGetResponseJSON    `json:"-"`
}

// accessAppGetResponseJSON contains the JSON metadata for the struct
// [AccessAppGetResponse]
type accessAppGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppGetResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    accessAppGetResponseErrorJSON `json:"-"`
}

// accessAppGetResponseErrorJSON contains the JSON metadata for the struct
// [AccessAppGetResponseError]
type accessAppGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppGetResponseMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    accessAppGetResponseMessageJSON `json:"-"`
}

// accessAppGetResponseMessageJSON contains the JSON metadata for the struct
// [AccessAppGetResponseMessage]
type accessAppGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccessAppGetResponseResultObject],
// [AccessAppGetResponseResultObject], [AccessAppGetResponseResultObject],
// [AccessAppGetResponseResultObject], [AccessAppGetResponseResultObject],
// [AccessAppGetResponseResultObject], [AccessAppGetResponseResultObject] or
// [AccessAppGetResponseResultObject].
type AccessAppGetResponseResult interface {
	implementsAccessAppGetResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppGetResponseResult)(nil)).Elem(), "")
}

type AccessAppGetResponseResultObject struct {
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
	AutoRedirectToIdentity bool                                        `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessAppGetResponseResultObjectCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                   `json:"created_at" format:"date-time"`
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
	Type      string                               `json:"type"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      accessAppGetResponseResultObjectJSON `json:"-"`
}

// accessAppGetResponseResultObjectJSON contains the JSON metadata for the struct
// [AccessAppGetResponseResultObject]
type accessAppGetResponseResultObjectJSON struct {
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

func (r *AccessAppGetResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppGetResponseResultObject) implementsAccessAppGetResponseResult() {}

type AccessAppGetResponseResultObjectCorsHeaders struct {
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
	AllowedMethods []AccessAppGetResponseResultObjectCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                         `json:"max_age"`
	JSON   accessAppGetResponseResultObjectCorsHeadersJSON `json:"-"`
}

// accessAppGetResponseResultObjectCorsHeadersJSON contains the JSON metadata for
// the struct [AccessAppGetResponseResultObjectCorsHeaders]
type accessAppGetResponseResultObjectCorsHeadersJSON struct {
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

func (r *AccessAppGetResponseResultObjectCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppGetResponseResultObjectCorsHeadersAllowedMethod string

const (
	AccessAppGetResponseResultObjectCorsHeadersAllowedMethodGet     AccessAppGetResponseResultObjectCorsHeadersAllowedMethod = "GET"
	AccessAppGetResponseResultObjectCorsHeadersAllowedMethodPost    AccessAppGetResponseResultObjectCorsHeadersAllowedMethod = "POST"
	AccessAppGetResponseResultObjectCorsHeadersAllowedMethodHead    AccessAppGetResponseResultObjectCorsHeadersAllowedMethod = "HEAD"
	AccessAppGetResponseResultObjectCorsHeadersAllowedMethodPut     AccessAppGetResponseResultObjectCorsHeadersAllowedMethod = "PUT"
	AccessAppGetResponseResultObjectCorsHeadersAllowedMethodDelete  AccessAppGetResponseResultObjectCorsHeadersAllowedMethod = "DELETE"
	AccessAppGetResponseResultObjectCorsHeadersAllowedMethodConnect AccessAppGetResponseResultObjectCorsHeadersAllowedMethod = "CONNECT"
	AccessAppGetResponseResultObjectCorsHeadersAllowedMethodOptions AccessAppGetResponseResultObjectCorsHeadersAllowedMethod = "OPTIONS"
	AccessAppGetResponseResultObjectCorsHeadersAllowedMethodTrace   AccessAppGetResponseResultObjectCorsHeadersAllowedMethod = "TRACE"
	AccessAppGetResponseResultObjectCorsHeadersAllowedMethodPatch   AccessAppGetResponseResultObjectCorsHeadersAllowedMethod = "PATCH"
)

// Whether the API call was successful
type AccessAppGetResponseSuccess bool

const (
	AccessAppGetResponseSuccessTrue AccessAppGetResponseSuccess = true
)

type AccessAppUpdateResponse struct {
	Errors   []AccessAppUpdateResponseError   `json:"errors"`
	Messages []AccessAppUpdateResponseMessage `json:"messages"`
	Result   AccessAppUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessAppUpdateResponseSuccess `json:"success"`
	JSON    accessAppUpdateResponseJSON    `json:"-"`
}

// accessAppUpdateResponseJSON contains the JSON metadata for the struct
// [AccessAppUpdateResponse]
type accessAppUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUpdateResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    accessAppUpdateResponseErrorJSON `json:"-"`
}

// accessAppUpdateResponseErrorJSON contains the JSON metadata for the struct
// [AccessAppUpdateResponseError]
type accessAppUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUpdateResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    accessAppUpdateResponseMessageJSON `json:"-"`
}

// accessAppUpdateResponseMessageJSON contains the JSON metadata for the struct
// [AccessAppUpdateResponseMessage]
type accessAppUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccessAppUpdateResponseResultObject],
// [AccessAppUpdateResponseResultObject], [AccessAppUpdateResponseResultObject],
// [AccessAppUpdateResponseResultObject], [AccessAppUpdateResponseResultObject],
// [AccessAppUpdateResponseResultObject], [AccessAppUpdateResponseResultObject] or
// [AccessAppUpdateResponseResultObject].
type AccessAppUpdateResponseResult interface {
	implementsAccessAppUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppUpdateResponseResult)(nil)).Elem(), "")
}

type AccessAppUpdateResponseResultObject struct {
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
	CorsHeaders            AccessAppUpdateResponseResultObjectCorsHeaders `json:"cors_headers"`
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
	Type      string                                  `json:"type"`
	UpdatedAt time.Time                               `json:"updated_at" format:"date-time"`
	JSON      accessAppUpdateResponseResultObjectJSON `json:"-"`
}

// accessAppUpdateResponseResultObjectJSON contains the JSON metadata for the
// struct [AccessAppUpdateResponseResultObject]
type accessAppUpdateResponseResultObjectJSON struct {
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

func (r *AccessAppUpdateResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppUpdateResponseResultObject) implementsAccessAppUpdateResponseResult() {}

type AccessAppUpdateResponseResultObjectCorsHeaders struct {
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
	AllowedMethods []AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                            `json:"max_age"`
	JSON   accessAppUpdateResponseResultObjectCorsHeadersJSON `json:"-"`
}

// accessAppUpdateResponseResultObjectCorsHeadersJSON contains the JSON metadata
// for the struct [AccessAppUpdateResponseResultObjectCorsHeaders]
type accessAppUpdateResponseResultObjectCorsHeadersJSON struct {
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

func (r *AccessAppUpdateResponseResultObjectCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethod string

const (
	AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethodGet     AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethod = "GET"
	AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethodPost    AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethod = "POST"
	AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethodHead    AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethod = "HEAD"
	AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethodPut     AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethod = "PUT"
	AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethodDelete  AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethod = "DELETE"
	AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethodConnect AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethod = "CONNECT"
	AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethodOptions AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethod = "OPTIONS"
	AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethodTrace   AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethod = "TRACE"
	AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethodPatch   AccessAppUpdateResponseResultObjectCorsHeadersAllowedMethod = "PATCH"
)

// Whether the API call was successful
type AccessAppUpdateResponseSuccess bool

const (
	AccessAppUpdateResponseSuccessTrue AccessAppUpdateResponseSuccess = true
)

type AccessAppDeleteResponse struct {
	Errors   []AccessAppDeleteResponseError   `json:"errors"`
	Messages []AccessAppDeleteResponseMessage `json:"messages"`
	Result   AccessAppDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessAppDeleteResponseSuccess `json:"success"`
	JSON    accessAppDeleteResponseJSON    `json:"-"`
}

// accessAppDeleteResponseJSON contains the JSON metadata for the struct
// [AccessAppDeleteResponse]
type accessAppDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppDeleteResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    accessAppDeleteResponseErrorJSON `json:"-"`
}

// accessAppDeleteResponseErrorJSON contains the JSON metadata for the struct
// [AccessAppDeleteResponseError]
type accessAppDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppDeleteResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    accessAppDeleteResponseMessageJSON `json:"-"`
}

// accessAppDeleteResponseMessageJSON contains the JSON metadata for the struct
// [AccessAppDeleteResponseMessage]
type accessAppDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppDeleteResponseResult struct {
	// UUID
	ID   string                            `json:"id"`
	JSON accessAppDeleteResponseResultJSON `json:"-"`
}

// accessAppDeleteResponseResultJSON contains the JSON metadata for the struct
// [AccessAppDeleteResponseResult]
type accessAppDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppDeleteResponseSuccess bool

const (
	AccessAppDeleteResponseSuccessTrue AccessAppDeleteResponseSuccess = true
)

type AccessAppAccessApplicationsAddAnApplicationResponse struct {
	Errors   []AccessAppAccessApplicationsAddAnApplicationResponseError   `json:"errors"`
	Messages []AccessAppAccessApplicationsAddAnApplicationResponseMessage `json:"messages"`
	Result   AccessAppAccessApplicationsAddAnApplicationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessAppAccessApplicationsAddAnApplicationResponseSuccess `json:"success"`
	JSON    accessAppAccessApplicationsAddAnApplicationResponseJSON    `json:"-"`
}

// accessAppAccessApplicationsAddAnApplicationResponseJSON contains the JSON
// metadata for the struct [AccessAppAccessApplicationsAddAnApplicationResponse]
type accessAppAccessApplicationsAddAnApplicationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsAddAnApplicationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppAccessApplicationsAddAnApplicationResponseError struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    accessAppAccessApplicationsAddAnApplicationResponseErrorJSON `json:"-"`
}

// accessAppAccessApplicationsAddAnApplicationResponseErrorJSON contains the JSON
// metadata for the struct
// [AccessAppAccessApplicationsAddAnApplicationResponseError]
type accessAppAccessApplicationsAddAnApplicationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsAddAnApplicationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppAccessApplicationsAddAnApplicationResponseMessage struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    accessAppAccessApplicationsAddAnApplicationResponseMessageJSON `json:"-"`
}

// accessAppAccessApplicationsAddAnApplicationResponseMessageJSON contains the JSON
// metadata for the struct
// [AccessAppAccessApplicationsAddAnApplicationResponseMessage]
type accessAppAccessApplicationsAddAnApplicationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsAddAnApplicationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccessAppAccessApplicationsAddAnApplicationResponseResultObject],
// [AccessAppAccessApplicationsAddAnApplicationResponseResultObject],
// [AccessAppAccessApplicationsAddAnApplicationResponseResultObject],
// [AccessAppAccessApplicationsAddAnApplicationResponseResultObject],
// [AccessAppAccessApplicationsAddAnApplicationResponseResultObject],
// [AccessAppAccessApplicationsAddAnApplicationResponseResultObject],
// [AccessAppAccessApplicationsAddAnApplicationResponseResultObject] or
// [AccessAppAccessApplicationsAddAnApplicationResponseResultObject].
type AccessAppAccessApplicationsAddAnApplicationResponseResult interface {
	implementsAccessAppAccessApplicationsAddAnApplicationResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppAccessApplicationsAddAnApplicationResponseResult)(nil)).Elem(), "")
}

type AccessAppAccessApplicationsAddAnApplicationResponseResultObject struct {
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
	CorsHeaders            AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeaders `json:"cors_headers"`
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
	JSON      accessAppAccessApplicationsAddAnApplicationResponseResultObjectJSON `json:"-"`
}

// accessAppAccessApplicationsAddAnApplicationResponseResultObjectJSON contains the
// JSON metadata for the struct
// [AccessAppAccessApplicationsAddAnApplicationResponseResultObject]
type accessAppAccessApplicationsAddAnApplicationResponseResultObjectJSON struct {
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

func (r *AccessAppAccessApplicationsAddAnApplicationResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppAccessApplicationsAddAnApplicationResponseResultObject) implementsAccessAppAccessApplicationsAddAnApplicationResponseResult() {
}

type AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeaders struct {
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
	AllowedMethods []AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                        `json:"max_age"`
	JSON   accessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersJSON `json:"-"`
}

// accessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersJSON
// contains the JSON metadata for the struct
// [AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeaders]
type accessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersJSON struct {
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

func (r *AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethod string

const (
	AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethodGet     AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethod = "GET"
	AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethodPost    AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethod = "POST"
	AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethodHead    AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethod = "HEAD"
	AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethodPut     AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethod = "PUT"
	AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethodDelete  AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethod = "DELETE"
	AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethodConnect AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethod = "CONNECT"
	AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethodOptions AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethod = "OPTIONS"
	AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethodTrace   AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethod = "TRACE"
	AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethodPatch   AccessAppAccessApplicationsAddAnApplicationResponseResultObjectCorsHeadersAllowedMethod = "PATCH"
)

// Whether the API call was successful
type AccessAppAccessApplicationsAddAnApplicationResponseSuccess bool

const (
	AccessAppAccessApplicationsAddAnApplicationResponseSuccessTrue AccessAppAccessApplicationsAddAnApplicationResponseSuccess = true
)

type AccessAppAccessApplicationsListAccessApplicationsResponse struct {
	Errors     []AccessAppAccessApplicationsListAccessApplicationsResponseError    `json:"errors"`
	Messages   []AccessAppAccessApplicationsListAccessApplicationsResponseMessage  `json:"messages"`
	Result     []AccessAppAccessApplicationsListAccessApplicationsResponseResult   `json:"result"`
	ResultInfo AccessAppAccessApplicationsListAccessApplicationsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessAppAccessApplicationsListAccessApplicationsResponseSuccess `json:"success"`
	JSON    accessAppAccessApplicationsListAccessApplicationsResponseJSON    `json:"-"`
}

// accessAppAccessApplicationsListAccessApplicationsResponseJSON contains the JSON
// metadata for the struct
// [AccessAppAccessApplicationsListAccessApplicationsResponse]
type accessAppAccessApplicationsListAccessApplicationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsListAccessApplicationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppAccessApplicationsListAccessApplicationsResponseError struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    accessAppAccessApplicationsListAccessApplicationsResponseErrorJSON `json:"-"`
}

// accessAppAccessApplicationsListAccessApplicationsResponseErrorJSON contains the
// JSON metadata for the struct
// [AccessAppAccessApplicationsListAccessApplicationsResponseError]
type accessAppAccessApplicationsListAccessApplicationsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsListAccessApplicationsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppAccessApplicationsListAccessApplicationsResponseMessage struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    accessAppAccessApplicationsListAccessApplicationsResponseMessageJSON `json:"-"`
}

// accessAppAccessApplicationsListAccessApplicationsResponseMessageJSON contains
// the JSON metadata for the struct
// [AccessAppAccessApplicationsListAccessApplicationsResponseMessage]
type accessAppAccessApplicationsListAccessApplicationsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsListAccessApplicationsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccessAppAccessApplicationsListAccessApplicationsResponseResultObject],
// [AccessAppAccessApplicationsListAccessApplicationsResponseResultObject],
// [AccessAppAccessApplicationsListAccessApplicationsResponseResultObject],
// [AccessAppAccessApplicationsListAccessApplicationsResponseResultObject],
// [AccessAppAccessApplicationsListAccessApplicationsResponseResultObject],
// [AccessAppAccessApplicationsListAccessApplicationsResponseResultObject],
// [AccessAppAccessApplicationsListAccessApplicationsResponseResultObject] or
// [AccessAppAccessApplicationsListAccessApplicationsResponseResultObject].
type AccessAppAccessApplicationsListAccessApplicationsResponseResult interface {
	implementsAccessAppAccessApplicationsListAccessApplicationsResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppAccessApplicationsListAccessApplicationsResponseResult)(nil)).Elem(), "")
}

type AccessAppAccessApplicationsListAccessApplicationsResponseResultObject struct {
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
	AutoRedirectToIdentity bool                                                                             `json:"auto_redirect_to_identity"`
	CorsHeaders            AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeaders `json:"cors_headers"`
	CreatedAt              time.Time                                                                        `json:"created_at" format:"date-time"`
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
	Type      string                                                                    `json:"type"`
	UpdatedAt time.Time                                                                 `json:"updated_at" format:"date-time"`
	JSON      accessAppAccessApplicationsListAccessApplicationsResponseResultObjectJSON `json:"-"`
}

// accessAppAccessApplicationsListAccessApplicationsResponseResultObjectJSON
// contains the JSON metadata for the struct
// [AccessAppAccessApplicationsListAccessApplicationsResponseResultObject]
type accessAppAccessApplicationsListAccessApplicationsResponseResultObjectJSON struct {
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

func (r *AccessAppAccessApplicationsListAccessApplicationsResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppAccessApplicationsListAccessApplicationsResponseResultObject) implementsAccessAppAccessApplicationsListAccessApplicationsResponseResult() {
}

type AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeaders struct {
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
	AllowedMethods []AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethod `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                                                                              `json:"max_age"`
	JSON   accessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersJSON `json:"-"`
}

// accessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersJSON
// contains the JSON metadata for the struct
// [AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeaders]
type accessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersJSON struct {
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

func (r *AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethod string

const (
	AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethodGet     AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethod = "GET"
	AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethodPost    AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethod = "POST"
	AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethodHead    AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethod = "HEAD"
	AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethodPut     AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethod = "PUT"
	AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethodDelete  AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethod = "DELETE"
	AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethodConnect AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethod = "CONNECT"
	AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethodOptions AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethod = "OPTIONS"
	AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethodTrace   AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethod = "TRACE"
	AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethodPatch   AccessAppAccessApplicationsListAccessApplicationsResponseResultObjectCorsHeadersAllowedMethod = "PATCH"
)

type AccessAppAccessApplicationsListAccessApplicationsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                 `json:"total_count"`
	JSON       accessAppAccessApplicationsListAccessApplicationsResponseResultInfoJSON `json:"-"`
}

// accessAppAccessApplicationsListAccessApplicationsResponseResultInfoJSON contains
// the JSON metadata for the struct
// [AccessAppAccessApplicationsListAccessApplicationsResponseResultInfo]
type accessAppAccessApplicationsListAccessApplicationsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppAccessApplicationsListAccessApplicationsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppAccessApplicationsListAccessApplicationsResponseSuccess bool

const (
	AccessAppAccessApplicationsListAccessApplicationsResponseSuccessTrue AccessAppAccessApplicationsListAccessApplicationsResponseSuccess = true
)

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppGetParamsAppID interface {
	ImplementsAccessAppGetParamsAppID()
}

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

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppDeleteParamsAppID interface {
	ImplementsAccessAppDeleteParamsAppID()
}

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

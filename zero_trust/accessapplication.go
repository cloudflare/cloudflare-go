// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// AccessApplicationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessApplicationService] method instead.
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
	var env AccessApplicationNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
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
func (r *AccessApplicationService) Update(ctx context.Context, appID AppIDUnionParam, params AccessApplicationUpdateParams, opts ...option.RequestOption) (res *AccessApplicationUpdateResponse, err error) {
	var env AccessApplicationUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
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
func (r *AccessApplicationService) List(ctx context.Context, query AccessApplicationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AccessApplicationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps", accountOrZone, accountOrZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *AccessApplicationService) ListAutoPaging(ctx context.Context, query AccessApplicationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AccessApplicationListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes an application from Access.
func (r *AccessApplicationService) Delete(ctx context.Context, appID AppIDUnionParam, body AccessApplicationDeleteParams, opts ...option.RequestOption) (res *AccessApplicationDeleteResponse, err error) {
	var env AccessApplicationDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
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
func (r *AccessApplicationService) Get(ctx context.Context, appID AppIDUnionParam, query AccessApplicationGetParams, opts ...option.RequestOption) (res *AccessApplicationGetResponse, err error) {
	var env AccessApplicationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
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
	var env AccessApplicationRevokeTokensResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
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

type AllowedHeaders = string

type AllowedHeadersParam = string

type AllowedIdPs = string

type AllowedIdPsParam = string

type AllowedMethods string

const (
	AllowedMethodsGet     AllowedMethods = "GET"
	AllowedMethodsPost    AllowedMethods = "POST"
	AllowedMethodsHead    AllowedMethods = "HEAD"
	AllowedMethodsPut     AllowedMethods = "PUT"
	AllowedMethodsDelete  AllowedMethods = "DELETE"
	AllowedMethodsConnect AllowedMethods = "CONNECT"
	AllowedMethodsOptions AllowedMethods = "OPTIONS"
	AllowedMethodsTrace   AllowedMethods = "TRACE"
	AllowedMethodsPatch   AllowedMethods = "PATCH"
)

func (r AllowedMethods) IsKnown() bool {
	switch r {
	case AllowedMethodsGet, AllowedMethodsPost, AllowedMethodsHead, AllowedMethodsPut, AllowedMethodsDelete, AllowedMethodsConnect, AllowedMethodsOptions, AllowedMethodsTrace, AllowedMethodsPatch:
		return true
	}
	return false
}

type AllowedOrigins = string

type AllowedOriginsParam = string

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AppIDUnionParam interface {
	ImplementsZeroTrustAppIDUnionParam()
}

type ApplicationPolicy struct {
	// The UUID of the policy
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []ApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision Decision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired bool `json:"isolation_required"`
	// The name of the Access policy.
	Name string `json:"name"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt string `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired bool `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require []AccessRule `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                `json:"session_duration"`
	UpdatedAt       time.Time             `json:"updated_at" format:"date-time"`
	JSON            applicationPolicyJSON `json:"-"`
}

// applicationPolicyJSON contains the JSON metadata for the struct
// [ApplicationPolicy]
type applicationPolicyJSON struct {
	ID                           apijson.Field
	ApprovalGroups               apijson.Field
	ApprovalRequired             apijson.Field
	CreatedAt                    apijson.Field
	Decision                     apijson.Field
	Exclude                      apijson.Field
	Include                      apijson.Field
	IsolationRequired            apijson.Field
	Name                         apijson.Field
	PurposeJustificationPrompt   apijson.Field
	PurposeJustificationRequired apijson.Field
	Require                      apijson.Field
	SessionDuration              apijson.Field
	UpdatedAt                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ApplicationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationPolicyJSON) RawJSON() string {
	return r.raw
}

// The application type.
type ApplicationType string

const (
	ApplicationTypeSelfHosted  ApplicationType = "self_hosted"
	ApplicationTypeSaaS        ApplicationType = "saas"
	ApplicationTypeSSH         ApplicationType = "ssh"
	ApplicationTypeVNC         ApplicationType = "vnc"
	ApplicationTypeAppLauncher ApplicationType = "app_launcher"
	ApplicationTypeWARP        ApplicationType = "warp"
	ApplicationTypeBISO        ApplicationType = "biso"
	ApplicationTypeBookmark    ApplicationType = "bookmark"
	ApplicationTypeDashSSO     ApplicationType = "dash_sso"
)

func (r ApplicationType) IsKnown() bool {
	switch r {
	case ApplicationTypeSelfHosted, ApplicationTypeSaaS, ApplicationTypeSSH, ApplicationTypeVNC, ApplicationTypeAppLauncher, ApplicationTypeWARP, ApplicationTypeBISO, ApplicationTypeBookmark, ApplicationTypeDashSSO:
		return true
	}
	return false
}

type CORSHeaders struct {
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
	AllowedHeaders []AllowedHeaders `json:"allowed_headers"`
	// Allowed HTTP request methods.
	AllowedMethods []AllowedMethods `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []AllowedOrigins `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64         `json:"max_age"`
	JSON   corsHeadersJSON `json:"-"`
}

// corsHeadersJSON contains the JSON metadata for the struct [CORSHeaders]
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

func (r *CORSHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corsHeadersJSON) RawJSON() string {
	return r.raw
}

type CORSHeadersParam struct {
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
	AllowedHeaders param.Field[[]AllowedHeadersParam] `json:"allowed_headers"`
	// Allowed HTTP request methods.
	AllowedMethods param.Field[[]AllowedMethods] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]AllowedOriginsParam] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r CORSHeadersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action Access will take if a user matches this policy.
type Decision string

const (
	DecisionAllow       Decision = "allow"
	DecisionDeny        Decision = "deny"
	DecisionNonIdentity Decision = "non_identity"
	DecisionBypass      Decision = "bypass"
)

func (r Decision) IsKnown() bool {
	switch r {
	case DecisionAllow, DecisionDeny, DecisionNonIdentity, DecisionBypass:
		return true
	}
	return false
}

type OIDCSaaSApp struct {
	// The lifetime of the OIDC Access Token after creation. Valid units are m,h. Must
	// be greater than or equal to 1m and less than or equal to 24h.
	AccessTokenLifetime string `json:"access_token_lifetime"`
	// If client secret should be required on the token endpoint when
	// authorization_code_with_pkce grant is used.
	AllowPKCEWithoutClientSecret bool `json:"allow_pkce_without_client_secret"`
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType OIDCSaaSAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string                  `json:"client_secret"`
	CreatedAt    time.Time               `json:"created_at" format:"date-time"`
	CustomClaims OIDCSaaSAppCustomClaims `json:"custom_claims"`
	// The OIDC flows supported by this application
	GrantTypes []OIDCSaaSAppGrantType `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex         string                              `json:"group_filter_regex"`
	HybridAndImplicitOptions OIDCSaaSAppHybridAndImplicitOptions `json:"hybrid_and_implicit_options"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs        []string                       `json:"redirect_uris"`
	RefreshTokenOptions OIDCSaaSAppRefreshTokenOptions `json:"refresh_token_options"`
	// Define the user information shared with access, "offline_access" scope will be
	// automatically enabled if refresh tokens are enabled
	Scopes    []OIDCSaaSAppScope `json:"scopes"`
	UpdatedAt time.Time          `json:"updated_at" format:"date-time"`
	JSON      oidcSaaSAppJSON    `json:"-"`
}

// oidcSaaSAppJSON contains the JSON metadata for the struct [OIDCSaaSApp]
type oidcSaaSAppJSON struct {
	AccessTokenLifetime          apijson.Field
	AllowPKCEWithoutClientSecret apijson.Field
	AppLauncherURL               apijson.Field
	AuthType                     apijson.Field
	ClientID                     apijson.Field
	ClientSecret                 apijson.Field
	CreatedAt                    apijson.Field
	CustomClaims                 apijson.Field
	GrantTypes                   apijson.Field
	GroupFilterRegex             apijson.Field
	HybridAndImplicitOptions     apijson.Field
	PublicKey                    apijson.Field
	RedirectURIs                 apijson.Field
	RefreshTokenOptions          apijson.Field
	Scopes                       apijson.Field
	UpdatedAt                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *OIDCSaaSApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oidcSaaSAppJSON) RawJSON() string {
	return r.raw
}

func (r OIDCSaaSApp) implementsZeroTrustAccessApplicationNewResponseSaaSApplicationSaaSApp() {}

func (r OIDCSaaSApp) implementsZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaaSApp() {}

func (r OIDCSaaSApp) implementsZeroTrustAccessApplicationListResponseSaaSApplicationSaaSApp() {}

func (r OIDCSaaSApp) implementsZeroTrustAccessApplicationGetResponseSaaSApplicationSaaSApp() {}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type OIDCSaaSAppAuthType string

const (
	OIDCSaaSAppAuthTypeSAML OIDCSaaSAppAuthType = "saml"
	OIDCSaaSAppAuthTypeOIDC OIDCSaaSAppAuthType = "oidc"
)

func (r OIDCSaaSAppAuthType) IsKnown() bool {
	switch r {
	case OIDCSaaSAppAuthTypeSAML, OIDCSaaSAppAuthTypeOIDC:
		return true
	}
	return false
}

type OIDCSaaSAppCustomClaims struct {
	// The name of the claim.
	Name string `json:"name"`
	// If the claim is required when building an OIDC token.
	Required bool `json:"required"`
	// The scope of the claim.
	Scope  OIDCSaaSAppCustomClaimsScope  `json:"scope"`
	Source OIDCSaaSAppCustomClaimsSource `json:"source"`
	JSON   oidcSaaSAppCustomClaimsJSON   `json:"-"`
}

// oidcSaaSAppCustomClaimsJSON contains the JSON metadata for the struct
// [OIDCSaaSAppCustomClaims]
type oidcSaaSAppCustomClaimsJSON struct {
	Name        apijson.Field
	Required    apijson.Field
	Scope       apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OIDCSaaSAppCustomClaims) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oidcSaaSAppCustomClaimsJSON) RawJSON() string {
	return r.raw
}

// The scope of the claim.
type OIDCSaaSAppCustomClaimsScope string

const (
	OIDCSaaSAppCustomClaimsScopeGroups  OIDCSaaSAppCustomClaimsScope = "groups"
	OIDCSaaSAppCustomClaimsScopeProfile OIDCSaaSAppCustomClaimsScope = "profile"
	OIDCSaaSAppCustomClaimsScopeEmail   OIDCSaaSAppCustomClaimsScope = "email"
	OIDCSaaSAppCustomClaimsScopeOpenid  OIDCSaaSAppCustomClaimsScope = "openid"
)

func (r OIDCSaaSAppCustomClaimsScope) IsKnown() bool {
	switch r {
	case OIDCSaaSAppCustomClaimsScopeGroups, OIDCSaaSAppCustomClaimsScopeProfile, OIDCSaaSAppCustomClaimsScopeEmail, OIDCSaaSAppCustomClaimsScopeOpenid:
		return true
	}
	return false
}

type OIDCSaaSAppCustomClaimsSource struct {
	// The name of the IdP claim.
	Name string `json:"name"`
	// A mapping from IdP ID to claim name.
	NameByIdP map[string]string                 `json:"name_by_idp"`
	JSON      oidcSaaSAppCustomClaimsSourceJSON `json:"-"`
}

// oidcSaaSAppCustomClaimsSourceJSON contains the JSON metadata for the struct
// [OIDCSaaSAppCustomClaimsSource]
type oidcSaaSAppCustomClaimsSourceJSON struct {
	Name        apijson.Field
	NameByIdP   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OIDCSaaSAppCustomClaimsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oidcSaaSAppCustomClaimsSourceJSON) RawJSON() string {
	return r.raw
}

type OIDCSaaSAppGrantType string

const (
	OIDCSaaSAppGrantTypeAuthorizationCode         OIDCSaaSAppGrantType = "authorization_code"
	OIDCSaaSAppGrantTypeAuthorizationCodeWithPKCE OIDCSaaSAppGrantType = "authorization_code_with_pkce"
	OIDCSaaSAppGrantTypeRefreshTokens             OIDCSaaSAppGrantType = "refresh_tokens"
	OIDCSaaSAppGrantTypeHybrid                    OIDCSaaSAppGrantType = "hybrid"
	OIDCSaaSAppGrantTypeImplicit                  OIDCSaaSAppGrantType = "implicit"
)

func (r OIDCSaaSAppGrantType) IsKnown() bool {
	switch r {
	case OIDCSaaSAppGrantTypeAuthorizationCode, OIDCSaaSAppGrantTypeAuthorizationCodeWithPKCE, OIDCSaaSAppGrantTypeRefreshTokens, OIDCSaaSAppGrantTypeHybrid, OIDCSaaSAppGrantTypeImplicit:
		return true
	}
	return false
}

type OIDCSaaSAppHybridAndImplicitOptions struct {
	// If an Access Token should be returned from the OIDC Authorization endpoint
	ReturnAccessTokenFromAuthorizationEndpoint bool `json:"return_access_token_from_authorization_endpoint"`
	// If an ID Token should be returned from the OIDC Authorization endpoint
	ReturnIDTokenFromAuthorizationEndpoint bool                                    `json:"return_id_token_from_authorization_endpoint"`
	JSON                                   oidcSaaSAppHybridAndImplicitOptionsJSON `json:"-"`
}

// oidcSaaSAppHybridAndImplicitOptionsJSON contains the JSON metadata for the
// struct [OIDCSaaSAppHybridAndImplicitOptions]
type oidcSaaSAppHybridAndImplicitOptionsJSON struct {
	ReturnAccessTokenFromAuthorizationEndpoint apijson.Field
	ReturnIDTokenFromAuthorizationEndpoint     apijson.Field
	raw                                        string
	ExtraFields                                map[string]apijson.Field
}

func (r *OIDCSaaSAppHybridAndImplicitOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oidcSaaSAppHybridAndImplicitOptionsJSON) RawJSON() string {
	return r.raw
}

type OIDCSaaSAppRefreshTokenOptions struct {
	// How long a refresh token will be valid for after creation. Valid units are
	// m,h,d. Must be longer than 1m.
	Lifetime string                             `json:"lifetime"`
	JSON     oidcSaaSAppRefreshTokenOptionsJSON `json:"-"`
}

// oidcSaaSAppRefreshTokenOptionsJSON contains the JSON metadata for the struct
// [OIDCSaaSAppRefreshTokenOptions]
type oidcSaaSAppRefreshTokenOptionsJSON struct {
	Lifetime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OIDCSaaSAppRefreshTokenOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oidcSaaSAppRefreshTokenOptionsJSON) RawJSON() string {
	return r.raw
}

type OIDCSaaSAppScope string

const (
	OIDCSaaSAppScopeOpenid  OIDCSaaSAppScope = "openid"
	OIDCSaaSAppScopeGroups  OIDCSaaSAppScope = "groups"
	OIDCSaaSAppScopeEmail   OIDCSaaSAppScope = "email"
	OIDCSaaSAppScopeProfile OIDCSaaSAppScope = "profile"
)

func (r OIDCSaaSAppScope) IsKnown() bool {
	switch r {
	case OIDCSaaSAppScopeOpenid, OIDCSaaSAppScopeGroups, OIDCSaaSAppScopeEmail, OIDCSaaSAppScopeProfile:
		return true
	}
	return false
}

type OIDCSaaSAppParam struct {
	// The lifetime of the OIDC Access Token after creation. Valid units are m,h. Must
	// be greater than or equal to 1m and less than or equal to 24h.
	AccessTokenLifetime param.Field[string] `json:"access_token_lifetime"`
	// If client secret should be required on the token endpoint when
	// authorization_code_with_pkce grant is used.
	AllowPKCEWithoutClientSecret param.Field[bool] `json:"allow_pkce_without_client_secret"`
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType param.Field[OIDCSaaSAppAuthType] `json:"auth_type"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string]                       `json:"client_secret"`
	CreatedAt    param.Field[time.Time]                    `json:"created_at" format:"date-time"`
	CustomClaims param.Field[OIDCSaaSAppCustomClaimsParam] `json:"custom_claims"`
	// The OIDC flows supported by this application
	GrantTypes param.Field[[]OIDCSaaSAppGrantType] `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex         param.Field[string]                                   `json:"group_filter_regex"`
	HybridAndImplicitOptions param.Field[OIDCSaaSAppHybridAndImplicitOptionsParam] `json:"hybrid_and_implicit_options"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectURIs        param.Field[[]string]                            `json:"redirect_uris"`
	RefreshTokenOptions param.Field[OIDCSaaSAppRefreshTokenOptionsParam] `json:"refresh_token_options"`
	// Define the user information shared with access, "offline_access" scope will be
	// automatically enabled if refresh tokens are enabled
	Scopes    param.Field[[]OIDCSaaSAppScope] `json:"scopes"`
	UpdatedAt param.Field[time.Time]          `json:"updated_at" format:"date-time"`
}

func (r OIDCSaaSAppParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OIDCSaaSAppParam) implementsZeroTrustAccessApplicationNewParamsBodySaaSApplicationSaaSAppUnion() {
}

func (r OIDCSaaSAppParam) implementsZeroTrustAccessApplicationUpdateParamsBodySaaSApplicationSaaSAppUnion() {
}

type OIDCSaaSAppCustomClaimsParam struct {
	// The name of the claim.
	Name param.Field[string] `json:"name"`
	// If the claim is required when building an OIDC token.
	Required param.Field[bool] `json:"required"`
	// The scope of the claim.
	Scope  param.Field[OIDCSaaSAppCustomClaimsScope]       `json:"scope"`
	Source param.Field[OIDCSaaSAppCustomClaimsSourceParam] `json:"source"`
}

func (r OIDCSaaSAppCustomClaimsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OIDCSaaSAppCustomClaimsSourceParam struct {
	// The name of the IdP claim.
	Name param.Field[string] `json:"name"`
	// A mapping from IdP ID to claim name.
	NameByIdP param.Field[map[string]string] `json:"name_by_idp"`
}

func (r OIDCSaaSAppCustomClaimsSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OIDCSaaSAppHybridAndImplicitOptionsParam struct {
	// If an Access Token should be returned from the OIDC Authorization endpoint
	ReturnAccessTokenFromAuthorizationEndpoint param.Field[bool] `json:"return_access_token_from_authorization_endpoint"`
	// If an ID Token should be returned from the OIDC Authorization endpoint
	ReturnIDTokenFromAuthorizationEndpoint param.Field[bool] `json:"return_id_token_from_authorization_endpoint"`
}

func (r OIDCSaaSAppHybridAndImplicitOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OIDCSaaSAppRefreshTokenOptionsParam struct {
	// How long a refresh token will be valid for after creation. Valid units are
	// m,h,d. Must be longer than 1m.
	Lifetime param.Field[string] `json:"lifetime"`
}

func (r OIDCSaaSAppRefreshTokenOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A globally unique name for an identity or service provider.
type SaaSAppNameFormat string

const (
	SaaSAppNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatUnspecified SaaSAppNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	SaaSAppNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatBasic       SaaSAppNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	SaaSAppNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatURI         SaaSAppNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

func (r SaaSAppNameFormat) IsKnown() bool {
	switch r {
	case SaaSAppNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatUnspecified, SaaSAppNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatBasic, SaaSAppNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatURI:
		return true
	}
	return false
}

// The format of the name identifier sent to the SaaS application.
type SaaSAppNameIDFormat string

const (
	SaaSAppNameIDFormatID    SaaSAppNameIDFormat = "id"
	SaaSAppNameIDFormatEmail SaaSAppNameIDFormat = "email"
)

func (r SaaSAppNameIDFormat) IsKnown() bool {
	switch r {
	case SaaSAppNameIDFormatID, SaaSAppNameIDFormatEmail:
		return true
	}
	return false
}

type SaaSAppSource struct {
	// The name of the IdP attribute.
	Name string `json:"name"`
	// A mapping from IdP ID to attribute name.
	NameByIdP map[string]string `json:"name_by_idp"`
	JSON      SaaSAppSourceJSON `json:"-"`
}

// SaaSAppSourceJSON contains the JSON metadata for the struct [SaaSAppSource]
type SaaSAppSourceJSON struct {
	Name        apijson.Field
	NameByIdP   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SaaSAppSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SaaSAppSourceJSON) RawJSON() string {
	return r.raw
}

type SaaSAppSourceParam struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
	// A mapping from IdP ID to attribute name.
	NameByIdP param.Field[map[string]string] `json:"name_by_idp"`
}

func (r SaaSAppSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SAMLSaaSApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType SAMLSaaSAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string                      `json:"consumer_service_url"`
	CreatedAt          time.Time                   `json:"created_at" format:"date-time"`
	CustomAttributes   SAMLSaaSAppCustomAttributes `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdPEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat SaaSAppNameIDFormat `json:"name_id_format"`
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
	SAMLAttributeTransformJsonata string `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SPEntityID string `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint string          `json:"sso_endpoint"`
	UpdatedAt   time.Time       `json:"updated_at" format:"date-time"`
	JSON        samlSaaSAppJSON `json:"-"`
}

// samlSaaSAppJSON contains the JSON metadata for the struct [SAMLSaaSApp]
type samlSaaSAppJSON struct {
	AuthType                      apijson.Field
	ConsumerServiceURL            apijson.Field
	CreatedAt                     apijson.Field
	CustomAttributes              apijson.Field
	DefaultRelayState             apijson.Field
	IdPEntityID                   apijson.Field
	NameIDFormat                  apijson.Field
	NameIDTransformJsonata        apijson.Field
	PublicKey                     apijson.Field
	SAMLAttributeTransformJsonata apijson.Field
	SPEntityID                    apijson.Field
	SSOEndpoint                   apijson.Field
	UpdatedAt                     apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *SAMLSaaSApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlSaaSAppJSON) RawJSON() string {
	return r.raw
}

func (r SAMLSaaSApp) implementsZeroTrustAccessApplicationNewResponseSaaSApplicationSaaSApp() {}

func (r SAMLSaaSApp) implementsZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaaSApp() {}

func (r SAMLSaaSApp) implementsZeroTrustAccessApplicationListResponseSaaSApplicationSaaSApp() {}

func (r SAMLSaaSApp) implementsZeroTrustAccessApplicationGetResponseSaaSApplicationSaaSApp() {}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type SAMLSaaSAppAuthType string

const (
	SAMLSaaSAppAuthTypeSAML SAMLSaaSAppAuthType = "saml"
	SAMLSaaSAppAuthTypeOIDC SAMLSaaSAppAuthType = "oidc"
)

func (r SAMLSaaSAppAuthType) IsKnown() bool {
	switch r {
	case SAMLSaaSAppAuthTypeSAML, SAMLSaaSAppAuthTypeOIDC:
		return true
	}
	return false
}

type SAMLSaaSAppCustomAttributes struct {
	// The SAML FriendlyName of the attribute.
	FriendlyName string `json:"friendly_name"`
	// The name of the attribute.
	Name string `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat SaaSAppNameFormat `json:"name_format"`
	// If the attribute is required when building a SAML assertion.
	Required bool                            `json:"required"`
	Source   SaaSAppSource                   `json:"source"`
	JSON     samlSaaSAppCustomAttributesJSON `json:"-"`
}

// samlSaaSAppCustomAttributesJSON contains the JSON metadata for the struct
// [SAMLSaaSAppCustomAttributes]
type samlSaaSAppCustomAttributesJSON struct {
	FriendlyName apijson.Field
	Name         apijson.Field
	NameFormat   apijson.Field
	Required     apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SAMLSaaSAppCustomAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlSaaSAppCustomAttributesJSON) RawJSON() string {
	return r.raw
}

type SAMLSaaSAppParam struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[SAMLSaaSAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                           `json:"consumer_service_url"`
	CreatedAt          param.Field[time.Time]                        `json:"created_at" format:"date-time"`
	CustomAttributes   param.Field[SAMLSaaSAppCustomAttributesParam] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdPEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[SaaSAppNameIDFormat] `json:"name_id_format"`
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
	SAMLAttributeTransformJsonata param.Field[string] `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SPEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string]    `json:"sso_endpoint"`
	UpdatedAt   param.Field[time.Time] `json:"updated_at" format:"date-time"`
}

func (r SAMLSaaSAppParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SAMLSaaSAppParam) implementsZeroTrustAccessApplicationNewParamsBodySaaSApplicationSaaSAppUnion() {
}

func (r SAMLSaaSAppParam) implementsZeroTrustAccessApplicationUpdateParamsBodySaaSApplicationSaaSAppUnion() {
}

type SAMLSaaSAppCustomAttributesParam struct {
	// The SAML FriendlyName of the attribute.
	FriendlyName param.Field[string] `json:"friendly_name"`
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[SaaSAppNameFormat] `json:"name_format"`
	// If the attribute is required when building a SAML assertion.
	Required param.Field[bool]               `json:"required"`
	Source   param.Field[SaaSAppSourceParam] `json:"source"`
}

func (r SAMLSaaSAppCustomAttributesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type SCIMConfigAuthenticationHTTPBasic struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme SCIMConfigAuthenticationHTTPBasicScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string                                `json:"user,required"`
	JSON scimConfigAuthenticationHTTPBasicJSON `json:"-"`
}

// scimConfigAuthenticationHTTPBasicJSON contains the JSON metadata for the struct
// [SCIMConfigAuthenticationHTTPBasic]
type scimConfigAuthenticationHTTPBasicJSON struct {
	Password    apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SCIMConfigAuthenticationHTTPBasic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigAuthenticationHTTPBasicJSON) RawJSON() string {
	return r.raw
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationNewResponseSaaSApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationListResponseSaaSApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationListResponseBookmarkApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationGetResponseSaaSApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationHTTPBasic) implementsZeroTrustAccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthentication() {
}

// The authentication scheme to use when making SCIM requests to this application.
type SCIMConfigAuthenticationHTTPBasicScheme string

const (
	SCIMConfigAuthenticationHTTPBasicSchemeHttpbasic SCIMConfigAuthenticationHTTPBasicScheme = "httpbasic"
)

func (r SCIMConfigAuthenticationHTTPBasicScheme) IsKnown() bool {
	switch r {
	case SCIMConfigAuthenticationHTTPBasicSchemeHttpbasic:
		return true
	}
	return false
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type SCIMConfigAuthenticationHTTPBasicParam struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[SCIMConfigAuthenticationHTTPBasicScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user,required"`
}

func (r SCIMConfigAuthenticationHTTPBasicParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationHTTPBasicParam) implementsZeroTrustAccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring OAuth Bearer Token authentication scheme for SCIM
// provisioning to an application.
type SCIMConfigAuthenticationOAuthBearerToken struct {
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme SCIMConfigAuthenticationOAuthBearerTokenScheme `json:"scheme,required"`
	JSON   scimConfigAuthenticationOAuthBearerTokenJSON   `json:"-"`
}

// scimConfigAuthenticationOAuthBearerTokenJSON contains the JSON metadata for the
// struct [SCIMConfigAuthenticationOAuthBearerToken]
type scimConfigAuthenticationOAuthBearerTokenJSON struct {
	Token       apijson.Field
	Scheme      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SCIMConfigAuthenticationOAuthBearerToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigAuthenticationOAuthBearerTokenJSON) RawJSON() string {
	return r.raw
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationNewResponseSaaSApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationListResponseSaaSApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationListResponseBookmarkApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationGetResponseSaaSApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOAuthBearerToken) implementsZeroTrustAccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthentication() {
}

// The authentication scheme to use when making SCIM requests to this application.
type SCIMConfigAuthenticationOAuthBearerTokenScheme string

const (
	SCIMConfigAuthenticationOAuthBearerTokenSchemeOauthbearertoken SCIMConfigAuthenticationOAuthBearerTokenScheme = "oauthbearertoken"
)

func (r SCIMConfigAuthenticationOAuthBearerTokenScheme) IsKnown() bool {
	switch r {
	case SCIMConfigAuthenticationOAuthBearerTokenSchemeOauthbearertoken:
		return true
	}
	return false
}

// Attributes for configuring OAuth Bearer Token authentication scheme for SCIM
// provisioning to an application.
type SCIMConfigAuthenticationOAuthBearerTokenParam struct {
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[SCIMConfigAuthenticationOAuthBearerTokenScheme] `json:"scheme,required"`
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOAuthBearerTokenParam) implementsZeroTrustAccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring OAuth 2 authentication scheme for SCIM provisioning
// to an application.
type SCIMConfigAuthenticationOauth2 struct {
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url,required"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id,required"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme SCIMConfigAuthenticationOauth2Scheme `json:"scheme,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string `json:"token_url,required"`
	// The authorization scopes to request when generating the token used to
	// authenticate with the remove SCIM service.
	Scopes []string                           `json:"scopes"`
	JSON   scimConfigAuthenticationOauth2JSON `json:"-"`
}

// scimConfigAuthenticationOauth2JSON contains the JSON metadata for the struct
// [SCIMConfigAuthenticationOauth2]
type scimConfigAuthenticationOauth2JSON struct {
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scheme           apijson.Field
	TokenURL         apijson.Field
	Scopes           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SCIMConfigAuthenticationOauth2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigAuthenticationOauth2JSON) RawJSON() string {
	return r.raw
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationNewResponseSaaSApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationListResponseSaaSApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationListResponseBookmarkApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationGetResponseSaaSApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication() {
}

func (r SCIMConfigAuthenticationOauth2) implementsZeroTrustAccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthentication() {
}

// The authentication scheme to use when making SCIM requests to this application.
type SCIMConfigAuthenticationOauth2Scheme string

const (
	SCIMConfigAuthenticationOauth2SchemeOauth2 SCIMConfigAuthenticationOauth2Scheme = "oauth2"
)

func (r SCIMConfigAuthenticationOauth2Scheme) IsKnown() bool {
	switch r {
	case SCIMConfigAuthenticationOauth2SchemeOauth2:
		return true
	}
	return false
}

// Attributes for configuring OAuth 2 authentication scheme for SCIM provisioning
// to an application.
type SCIMConfigAuthenticationOauth2Param struct {
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url,required"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id,required"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[SCIMConfigAuthenticationOauth2Scheme] `json:"scheme,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url,required"`
	// The authorization scopes to request when generating the token used to
	// authenticate with the remove SCIM service.
	Scopes param.Field[[]string] `json:"scopes"`
}

func (r SCIMConfigAuthenticationOauth2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

func (r SCIMConfigAuthenticationOauth2Param) implementsZeroTrustAccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion() {
}

// Transformations and filters applied to resources before they are provisioned in
// the remote SCIM service.
type SCIMConfigMapping struct {
	// Which SCIM resource type this mapping applies to.
	Schema string `json:"schema,required"`
	// Whether or not this mapping is enabled.
	Enabled bool `json:"enabled"`
	// A
	// [SCIM filter expression](https://datatracker.ietf.org/doc/html/rfc7644#section-3.4.2.2)
	// that matches resources that should be provisioned to this application.
	Filter string `json:"filter"`
	// Whether or not this mapping applies to creates, updates, or deletes.
	Operations SCIMConfigMappingOperations `json:"operations"`
	// A [JSONata](https://jsonata.org/) expression that transforms the resource before
	// provisioning it in the application.
	TransformJsonata string                `json:"transform_jsonata"`
	JSON             scimConfigMappingJSON `json:"-"`
}

// scimConfigMappingJSON contains the JSON metadata for the struct
// [SCIMConfigMapping]
type scimConfigMappingJSON struct {
	Schema           apijson.Field
	Enabled          apijson.Field
	Filter           apijson.Field
	Operations       apijson.Field
	TransformJsonata apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SCIMConfigMapping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigMappingJSON) RawJSON() string {
	return r.raw
}

// Whether or not this mapping applies to creates, updates, or deletes.
type SCIMConfigMappingOperations struct {
	// Whether or not this mapping applies to create (POST) operations.
	Create bool `json:"create"`
	// Whether or not this mapping applies to DELETE operations.
	Delete bool `json:"delete"`
	// Whether or not this mapping applies to update (PATCH/PUT) operations.
	Update bool                            `json:"update"`
	JSON   scimConfigMappingOperationsJSON `json:"-"`
}

// scimConfigMappingOperationsJSON contains the JSON metadata for the struct
// [SCIMConfigMappingOperations]
type scimConfigMappingOperationsJSON struct {
	Create      apijson.Field
	Delete      apijson.Field
	Update      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SCIMConfigMappingOperations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigMappingOperationsJSON) RawJSON() string {
	return r.raw
}

// Transformations and filters applied to resources before they are provisioned in
// the remote SCIM service.
type SCIMConfigMappingParam struct {
	// Which SCIM resource type this mapping applies to.
	Schema param.Field[string] `json:"schema,required"`
	// Whether or not this mapping is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// A
	// [SCIM filter expression](https://datatracker.ietf.org/doc/html/rfc7644#section-3.4.2.2)
	// that matches resources that should be provisioned to this application.
	Filter param.Field[string] `json:"filter"`
	// Whether or not this mapping applies to creates, updates, or deletes.
	Operations param.Field[SCIMConfigMappingOperationsParam] `json:"operations"`
	// A [JSONata](https://jsonata.org/) expression that transforms the resource before
	// provisioning it in the application.
	TransformJsonata param.Field[string] `json:"transform_jsonata"`
}

func (r SCIMConfigMappingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not this mapping applies to creates, updates, or deletes.
type SCIMConfigMappingOperationsParam struct {
	// Whether or not this mapping applies to create (POST) operations.
	Create param.Field[bool] `json:"create"`
	// Whether or not this mapping applies to DELETE operations.
	Delete param.Field[bool] `json:"delete"`
	// Whether or not this mapping applies to update (PATCH/PUT) operations.
	Update param.Field[bool] `json:"update"`
}

func (r SCIMConfigMappingOperationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SelfHostedDomains = string

type SelfHostedDomainsParam = string

type AccessApplicationNewResponse struct {
	// Audience tag.
	AUD       string    `json:"aud"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	ID string `json:"id"`
	// This field can have the runtime type of
	// [AccessApplicationNewResponseSelfHostedApplicationSCIMConfig],
	// [AccessApplicationNewResponseSaaSApplicationSCIMConfig],
	// [AccessApplicationNewResponseBrowserSSHApplicationSCIMConfig],
	// [AccessApplicationNewResponseBrowserVNCApplicationSCIMConfig],
	// [AccessApplicationNewResponseAppLauncherApplicationSCIMConfig],
	// [AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfig],
	// [AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfig],
	// [AccessApplicationNewResponseBookmarkApplicationSCIMConfig].
	SCIMConfig interface{} `json:"scim_config,required"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP bool `json:"allow_authenticate_via_warp"`
	// This field can have the runtime type of [[]AllowedIdPs].
	AllowedIdPs interface{} `json:"allowed_idps,required"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage string `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL string `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL string `json:"custom_non_identity_deny_url"`
	// This field can have the runtime type of [[]string].
	CustomPages interface{} `json:"custom_pages,required"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// This field can have the runtime type of [[]SelfHostedDomains].
	SelfHostedDomains interface{} `json:"self_hosted_domains,required"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect bool `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial bool `json:"skip_interstitial"`
	// This field can have the runtime type of [[]string].
	Tags interface{} `json:"tags,required"`
	// The application type.
	Type string `json:"type"`
	// This field can have the runtime type of [[]ApplicationPolicy].
	Policies interface{} `json:"policies,required"`
	// This field can have the runtime type of
	// [AccessApplicationNewResponseSaaSApplicationSaaSApp].
	SaaSApp interface{}                      `json:"saas_app,required"`
	JSON    accessApplicationNewResponseJSON `json:"-"`
	union   AccessApplicationNewResponseUnion
}

// accessApplicationNewResponseJSON contains the JSON metadata for the struct
// [AccessApplicationNewResponse]
type accessApplicationNewResponseJSON struct {
	AUD                      apijson.Field
	CreatedAt                apijson.Field
	ID                       apijson.Field
	SCIMConfig               apijson.Field
	UpdatedAt                apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	Domain                   apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SelfHostedDomains        apijson.Field
	ServiceAuth401Redirect   apijson.Field
	SessionDuration          apijson.Field
	SkipInterstitial         apijson.Field
	Tags                     apijson.Field
	Type                     apijson.Field
	Policies                 apijson.Field
	SaaSApp                  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r accessApplicationNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessApplicationNewResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.AccessApplicationNewResponseSelfHostedApplication],
// [zero_trust.AccessApplicationNewResponseSaaSApplication],
// [zero_trust.AccessApplicationNewResponseBrowserSSHApplication],
// [zero_trust.AccessApplicationNewResponseBrowserVNCApplication],
// [zero_trust.AccessApplicationNewResponseAppLauncherApplication],
// [zero_trust.AccessApplicationNewResponseDeviceEnrollmentPermissionsApplication],
// [zero_trust.AccessApplicationNewResponseBrowserIsolationPermissionsApplication],
// [zero_trust.AccessApplicationNewResponseBookmarkApplication].
func (r AccessApplicationNewResponse) AsUnion() AccessApplicationNewResponseUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.AccessApplicationNewResponseSelfHostedApplication],
// [zero_trust.AccessApplicationNewResponseSaaSApplication],
// [zero_trust.AccessApplicationNewResponseBrowserSSHApplication],
// [zero_trust.AccessApplicationNewResponseBrowserVNCApplication],
// [zero_trust.AccessApplicationNewResponseAppLauncherApplication],
// [zero_trust.AccessApplicationNewResponseDeviceEnrollmentPermissionsApplication],
// [zero_trust.AccessApplicationNewResponseBrowserIsolationPermissionsApplication]
// or [zero_trust.AccessApplicationNewResponseBookmarkApplication].
type AccessApplicationNewResponseUnion interface {
	implementsZeroTrustAccessApplicationNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationNewResponseUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AccessApplicationNewResponseBrowserVNCApplication{}),
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
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool                `json:"path_cookie_attribute"`
	Policies            []ApplicationPolicy `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationNewResponseSelfHostedApplicationSCIMConfig `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomains `json:"self_hosted_domains"`
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
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AUD                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	Policies                 apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SCIMConfig               apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewResponseSelfHostedApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                             `json:"mappings"`
	JSON     accessApplicationNewResponseSelfHostedApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationNewResponseSelfHostedApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationNewResponseSelfHostedApplicationSCIMConfig]
type accessApplicationNewResponseSelfHostedApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationNewResponseSelfHostedApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseSelfHostedApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                        `json:"token_url"`
	JSON     accessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthentication]
type accessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationNewResponseSaaSApplication struct {
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages []string `json:"custom_pages"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name     string                                             `json:"name"`
	Policies []ApplicationPolicy                                `json:"policies"`
	SaaSApp  AccessApplicationNewResponseSaaSApplicationSaaSApp `json:"saas_app"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationNewResponseSaaSApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AppLauncherVisible     apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	CustomPages            apijson.Field
	LogoURL                apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SaaSApp                apijson.Field
	SCIMConfig             apijson.Field
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

type AccessApplicationNewResponseSaaSApplicationSaaSApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType AccessApplicationNewResponseSaaSApplicationSaaSAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string    `json:"consumer_service_url"`
	CreatedAt          time.Time `json:"created_at" format:"date-time"`
	// This field can have the runtime type of [SAMLSaaSAppCustomAttributes].
	CustomAttributes interface{} `json:"custom_attributes,required"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdPEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat SaaSAppNameIDFormat `json:"name_id_format"`
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
	SAMLAttributeTransformJsonata string `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SPEntityID string `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint string    `json:"sso_endpoint"`
	UpdatedAt   time.Time `json:"updated_at" format:"date-time"`
	// The lifetime of the OIDC Access Token after creation. Valid units are m,h. Must
	// be greater than or equal to 1m and less than or equal to 24h.
	AccessTokenLifetime string `json:"access_token_lifetime"`
	// If client secret should be required on the token endpoint when
	// authorization_code_with_pkce grant is used.
	AllowPKCEWithoutClientSecret bool `json:"allow_pkce_without_client_secret"`
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [OIDCSaaSAppCustomClaims].
	CustomClaims interface{} `json:"custom_claims,required"`
	// This field can have the runtime type of [[]OIDCSaaSAppGrantType].
	GrantTypes interface{} `json:"grant_types,required"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// This field can have the runtime type of [OIDCSaaSAppHybridAndImplicitOptions].
	HybridAndImplicitOptions interface{} `json:"hybrid_and_implicit_options,required"`
	// This field can have the runtime type of [[]string].
	RedirectURIs interface{} `json:"redirect_uris,required"`
	// This field can have the runtime type of [OIDCSaaSAppRefreshTokenOptions].
	RefreshTokenOptions interface{} `json:"refresh_token_options,required"`
	// This field can have the runtime type of [[]OIDCSaaSAppScope].
	Scopes interface{}                                            `json:"scopes,required"`
	JSON   accessApplicationNewResponseSaaSApplicationSaaSAppJSON `json:"-"`
	union  AccessApplicationNewResponseSaaSApplicationSaaSAppUnion
}

// accessApplicationNewResponseSaaSApplicationSaaSAppJSON contains the JSON
// metadata for the struct [AccessApplicationNewResponseSaaSApplicationSaaSApp]
type accessApplicationNewResponseSaaSApplicationSaaSAppJSON struct {
	AuthType                      apijson.Field
	ConsumerServiceURL            apijson.Field
	CreatedAt                     apijson.Field
	CustomAttributes              apijson.Field
	DefaultRelayState             apijson.Field
	IdPEntityID                   apijson.Field
	NameIDFormat                  apijson.Field
	NameIDTransformJsonata        apijson.Field
	PublicKey                     apijson.Field
	SAMLAttributeTransformJsonata apijson.Field
	SPEntityID                    apijson.Field
	SSOEndpoint                   apijson.Field
	UpdatedAt                     apijson.Field
	AccessTokenLifetime           apijson.Field
	AllowPKCEWithoutClientSecret  apijson.Field
	AppLauncherURL                apijson.Field
	ClientID                      apijson.Field
	ClientSecret                  apijson.Field
	CustomClaims                  apijson.Field
	GrantTypes                    apijson.Field
	GroupFilterRegex              apijson.Field
	HybridAndImplicitOptions      apijson.Field
	RedirectURIs                  apijson.Field
	RefreshTokenOptions           apijson.Field
	Scopes                        apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r accessApplicationNewResponseSaaSApplicationSaaSAppJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationNewResponseSaaSApplicationSaaSApp) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationNewResponseSaaSApplicationSaaSApp{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessApplicationNewResponseSaaSApplicationSaaSAppUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [zero_trust.SAMLSaaSApp],
// [zero_trust.OIDCSaaSApp].
func (r AccessApplicationNewResponseSaaSApplicationSaaSApp) AsUnion() AccessApplicationNewResponseSaaSApplicationSaaSAppUnion {
	return r.union
}

// Union satisfied by [zero_trust.SAMLSaaSApp] or [zero_trust.OIDCSaaSApp].
type AccessApplicationNewResponseSaaSApplicationSaaSAppUnion interface {
	implementsZeroTrustAccessApplicationNewResponseSaaSApplicationSaaSApp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationNewResponseSaaSApplicationSaaSAppUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SAMLSaaSApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OIDCSaaSApp{}),
		},
	)
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationNewResponseSaaSApplicationSaaSAppAuthType string

const (
	AccessApplicationNewResponseSaaSApplicationSaaSAppAuthTypeSAML AccessApplicationNewResponseSaaSApplicationSaaSAppAuthType = "saml"
	AccessApplicationNewResponseSaaSApplicationSaaSAppAuthTypeOIDC AccessApplicationNewResponseSaaSApplicationSaaSAppAuthType = "oidc"
)

func (r AccessApplicationNewResponseSaaSApplicationSaaSAppAuthType) IsKnown() bool {
	switch r {
	case AccessApplicationNewResponseSaaSApplicationSaaSAppAuthTypeSAML, AccessApplicationNewResponseSaaSApplicationSaaSAppAuthTypeOIDC:
		return true
	}
	return false
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewResponseSaaSApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                       `json:"mappings"`
	JSON     accessApplicationNewResponseSaaSApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationNewResponseSaaSApplicationSCIMConfigJSON contains the JSON
// metadata for the struct [AccessApplicationNewResponseSaaSApplicationSCIMConfig]
type accessApplicationNewResponseSaaSApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationNewResponseSaaSApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseSaaSApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                  `json:"token_url"`
	JSON     accessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationJSON contains
// the JSON metadata for the struct
// [AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthentication]
type accessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewResponseSaaSApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

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
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool                `json:"path_cookie_attribute"`
	Policies            []ApplicationPolicy `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationNewResponseBrowserSSHApplicationSCIMConfig `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomains `json:"self_hosted_domains"`
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
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AUD                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	Policies                 apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SCIMConfig               apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewResponseBrowserSSHApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                             `json:"mappings"`
	JSON     accessApplicationNewResponseBrowserSSHApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationNewResponseBrowserSSHApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationNewResponseBrowserSSHApplicationSCIMConfig]
type accessApplicationNewResponseBrowserSSHApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationNewResponseBrowserSSHApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseBrowserSSHApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                        `json:"token_url"`
	JSON     accessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthentication]
type accessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationNewResponseBrowserVNCApplication struct {
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
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool                `json:"path_cookie_attribute"`
	Policies            []ApplicationPolicy `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationNewResponseBrowserVNCApplicationSCIMConfig `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomains `json:"self_hosted_domains"`
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
	JSON      accessApplicationNewResponseBrowserVNCApplicationJSON `json:"-"`
}

// accessApplicationNewResponseBrowserVNCApplicationJSON contains the JSON metadata
// for the struct [AccessApplicationNewResponseBrowserVNCApplication]
type accessApplicationNewResponseBrowserVNCApplicationJSON struct {
	Domain                   apijson.Field
	Type                     apijson.Field
	ID                       apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AUD                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	Policies                 apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SCIMConfig               apijson.Field
	SelfHostedDomains        apijson.Field
	ServiceAuth401Redirect   apijson.Field
	SessionDuration          apijson.Field
	SkipInterstitial         apijson.Field
	Tags                     apijson.Field
	UpdatedAt                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccessApplicationNewResponseBrowserVNCApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseBrowserVNCApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationNewResponseBrowserVNCApplication) implementsZeroTrustAccessApplicationNewResponse() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewResponseBrowserVNCApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                             `json:"mappings"`
	JSON     accessApplicationNewResponseBrowserVNCApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationNewResponseBrowserVNCApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationNewResponseBrowserVNCApplicationSCIMConfig]
type accessApplicationNewResponseBrowserVNCApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationNewResponseBrowserVNCApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseBrowserVNCApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                        `json:"token_url"`
	JSON     accessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthentication]
type accessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationNewResponseAppLauncherApplication struct {
	// The application type.
	Type ApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// The name of the application.
	Name     string              `json:"name"`
	Policies []ApplicationPolicy `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationNewResponseAppLauncherApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SCIMConfig             apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewResponseAppLauncherApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                              `json:"mappings"`
	JSON     accessApplicationNewResponseAppLauncherApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationNewResponseAppLauncherApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationNewResponseAppLauncherApplicationSCIMConfig]
type accessApplicationNewResponseAppLauncherApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationNewResponseAppLauncherApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseAppLauncherApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                         `json:"token_url"`
	JSON     accessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthentication]
type accessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationNewResponseDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type ApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// The name of the application.
	Name     string              `json:"name"`
	Policies []ApplicationPolicy `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SCIMConfig             apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                                              `json:"mappings"`
	JSON     accessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON
// contains the JSON metadata for the struct
// [AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfig]
type accessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                                         `json:"token_url"`
	JSON     accessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication]
type accessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationNewResponseBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type ApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// The name of the application.
	Name     string              `json:"name"`
	Policies []ApplicationPolicy `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SCIMConfig             apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                                              `json:"mappings"`
	JSON     accessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON
// contains the JSON metadata for the struct
// [AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfig]
type accessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                                         `json:"token_url"`
	JSON     accessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication]
type accessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationNewResponseBookmarkApplication struct {
	// UUID
	ID string `json:"id"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD       string    `json:"aud"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The URL or domain of the bookmark.
	Domain string `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name string `json:"name"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationNewResponseBookmarkApplicationSCIMConfig `json:"scim_config"`
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
	AUD                apijson.Field
	CreatedAt          apijson.Field
	Domain             apijson.Field
	LogoURL            apijson.Field
	Name               apijson.Field
	SCIMConfig         apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewResponseBookmarkApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                           `json:"mappings"`
	JSON     accessApplicationNewResponseBookmarkApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationNewResponseBookmarkApplicationSCIMConfigJSON contains the JSON
// metadata for the struct
// [AccessApplicationNewResponseBookmarkApplicationSCIMConfig]
type accessApplicationNewResponseBookmarkApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationNewResponseBookmarkApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationNewResponseBookmarkApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                      `json:"token_url"`
	JSON     accessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthentication]
type accessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationUpdateResponse struct {
	// Audience tag.
	AUD       string    `json:"aud"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	ID string `json:"id"`
	// This field can have the runtime type of
	// [AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfig],
	// [AccessApplicationUpdateResponseSaaSApplicationSCIMConfig],
	// [AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfig],
	// [AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfig],
	// [AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfig],
	// [AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfig],
	// [AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfig],
	// [AccessApplicationUpdateResponseBookmarkApplicationSCIMConfig].
	SCIMConfig interface{} `json:"scim_config,required"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP bool `json:"allow_authenticate_via_warp"`
	// This field can have the runtime type of [[]AllowedIdPs].
	AllowedIdPs interface{} `json:"allowed_idps,required"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage string `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL string `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL string `json:"custom_non_identity_deny_url"`
	// This field can have the runtime type of [[]string].
	CustomPages interface{} `json:"custom_pages,required"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// This field can have the runtime type of [[]SelfHostedDomains].
	SelfHostedDomains interface{} `json:"self_hosted_domains,required"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect bool `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial bool `json:"skip_interstitial"`
	// This field can have the runtime type of [[]string].
	Tags interface{} `json:"tags,required"`
	// The application type.
	Type string `json:"type"`
	// This field can have the runtime type of [[]ApplicationPolicy].
	Policies interface{} `json:"policies,required"`
	// This field can have the runtime type of
	// [AccessApplicationUpdateResponseSaaSApplicationSaaSApp].
	SaaSApp interface{}                         `json:"saas_app,required"`
	JSON    accessApplicationUpdateResponseJSON `json:"-"`
	union   AccessApplicationUpdateResponseUnion
}

// accessApplicationUpdateResponseJSON contains the JSON metadata for the struct
// [AccessApplicationUpdateResponse]
type accessApplicationUpdateResponseJSON struct {
	AUD                      apijson.Field
	CreatedAt                apijson.Field
	ID                       apijson.Field
	SCIMConfig               apijson.Field
	UpdatedAt                apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	Domain                   apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SelfHostedDomains        apijson.Field
	ServiceAuth401Redirect   apijson.Field
	SessionDuration          apijson.Field
	SkipInterstitial         apijson.Field
	Tags                     apijson.Field
	Type                     apijson.Field
	Policies                 apijson.Field
	SaaSApp                  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r accessApplicationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessApplicationUpdateResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.AccessApplicationUpdateResponseSelfHostedApplication],
// [zero_trust.AccessApplicationUpdateResponseSaaSApplication],
// [zero_trust.AccessApplicationUpdateResponseBrowserSSHApplication],
// [zero_trust.AccessApplicationUpdateResponseBrowserVNCApplication],
// [zero_trust.AccessApplicationUpdateResponseAppLauncherApplication],
// [zero_trust.AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplication],
// [zero_trust.AccessApplicationUpdateResponseBrowserIsolationPermissionsApplication],
// [zero_trust.AccessApplicationUpdateResponseBookmarkApplication].
func (r AccessApplicationUpdateResponse) AsUnion() AccessApplicationUpdateResponseUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.AccessApplicationUpdateResponseSelfHostedApplication],
// [zero_trust.AccessApplicationUpdateResponseSaaSApplication],
// [zero_trust.AccessApplicationUpdateResponseBrowserSSHApplication],
// [zero_trust.AccessApplicationUpdateResponseBrowserVNCApplication],
// [zero_trust.AccessApplicationUpdateResponseAppLauncherApplication],
// [zero_trust.AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplication],
// [zero_trust.AccessApplicationUpdateResponseBrowserIsolationPermissionsApplication]
// or [zero_trust.AccessApplicationUpdateResponseBookmarkApplication].
type AccessApplicationUpdateResponseUnion interface {
	implementsZeroTrustAccessApplicationUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationUpdateResponseUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AccessApplicationUpdateResponseBrowserVNCApplication{}),
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
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool                `json:"path_cookie_attribute"`
	Policies            []ApplicationPolicy `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfig `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomains `json:"self_hosted_domains"`
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
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AUD                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	Policies                 apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SCIMConfig               apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                                `json:"mappings"`
	JSON     accessApplicationUpdateResponseSelfHostedApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationUpdateResponseSelfHostedApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfig]
type accessApplicationUpdateResponseSelfHostedApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseSelfHostedApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                           `json:"token_url"`
	JSON     accessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthentication]
type accessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationUpdateResponseSaaSApplication struct {
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages []string `json:"custom_pages"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name     string                                                `json:"name"`
	Policies []ApplicationPolicy                                   `json:"policies"`
	SaaSApp  AccessApplicationUpdateResponseSaaSApplicationSaaSApp `json:"saas_app"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationUpdateResponseSaaSApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AppLauncherVisible     apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	CustomPages            apijson.Field
	LogoURL                apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SaaSApp                apijson.Field
	SCIMConfig             apijson.Field
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

type AccessApplicationUpdateResponseSaaSApplicationSaaSApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType AccessApplicationUpdateResponseSaaSApplicationSaaSAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string    `json:"consumer_service_url"`
	CreatedAt          time.Time `json:"created_at" format:"date-time"`
	// This field can have the runtime type of [SAMLSaaSAppCustomAttributes].
	CustomAttributes interface{} `json:"custom_attributes,required"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdPEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat SaaSAppNameIDFormat `json:"name_id_format"`
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
	SAMLAttributeTransformJsonata string `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SPEntityID string `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint string    `json:"sso_endpoint"`
	UpdatedAt   time.Time `json:"updated_at" format:"date-time"`
	// The lifetime of the OIDC Access Token after creation. Valid units are m,h. Must
	// be greater than or equal to 1m and less than or equal to 24h.
	AccessTokenLifetime string `json:"access_token_lifetime"`
	// If client secret should be required on the token endpoint when
	// authorization_code_with_pkce grant is used.
	AllowPKCEWithoutClientSecret bool `json:"allow_pkce_without_client_secret"`
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [OIDCSaaSAppCustomClaims].
	CustomClaims interface{} `json:"custom_claims,required"`
	// This field can have the runtime type of [[]OIDCSaaSAppGrantType].
	GrantTypes interface{} `json:"grant_types,required"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// This field can have the runtime type of [OIDCSaaSAppHybridAndImplicitOptions].
	HybridAndImplicitOptions interface{} `json:"hybrid_and_implicit_options,required"`
	// This field can have the runtime type of [[]string].
	RedirectURIs interface{} `json:"redirect_uris,required"`
	// This field can have the runtime type of [OIDCSaaSAppRefreshTokenOptions].
	RefreshTokenOptions interface{} `json:"refresh_token_options,required"`
	// This field can have the runtime type of [[]OIDCSaaSAppScope].
	Scopes interface{}                                               `json:"scopes,required"`
	JSON   accessApplicationUpdateResponseSaaSApplicationSaaSAppJSON `json:"-"`
	union  AccessApplicationUpdateResponseSaaSApplicationSaaSAppUnion
}

// accessApplicationUpdateResponseSaaSApplicationSaaSAppJSON contains the JSON
// metadata for the struct [AccessApplicationUpdateResponseSaaSApplicationSaaSApp]
type accessApplicationUpdateResponseSaaSApplicationSaaSAppJSON struct {
	AuthType                      apijson.Field
	ConsumerServiceURL            apijson.Field
	CreatedAt                     apijson.Field
	CustomAttributes              apijson.Field
	DefaultRelayState             apijson.Field
	IdPEntityID                   apijson.Field
	NameIDFormat                  apijson.Field
	NameIDTransformJsonata        apijson.Field
	PublicKey                     apijson.Field
	SAMLAttributeTransformJsonata apijson.Field
	SPEntityID                    apijson.Field
	SSOEndpoint                   apijson.Field
	UpdatedAt                     apijson.Field
	AccessTokenLifetime           apijson.Field
	AllowPKCEWithoutClientSecret  apijson.Field
	AppLauncherURL                apijson.Field
	ClientID                      apijson.Field
	ClientSecret                  apijson.Field
	CustomClaims                  apijson.Field
	GrantTypes                    apijson.Field
	GroupFilterRegex              apijson.Field
	HybridAndImplicitOptions      apijson.Field
	RedirectURIs                  apijson.Field
	RefreshTokenOptions           apijson.Field
	Scopes                        apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r accessApplicationUpdateResponseSaaSApplicationSaaSAppJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationUpdateResponseSaaSApplicationSaaSApp) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationUpdateResponseSaaSApplicationSaaSApp{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessApplicationUpdateResponseSaaSApplicationSaaSAppUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [zero_trust.SAMLSaaSApp],
// [zero_trust.OIDCSaaSApp].
func (r AccessApplicationUpdateResponseSaaSApplicationSaaSApp) AsUnion() AccessApplicationUpdateResponseSaaSApplicationSaaSAppUnion {
	return r.union
}

// Union satisfied by [zero_trust.SAMLSaaSApp] or [zero_trust.OIDCSaaSApp].
type AccessApplicationUpdateResponseSaaSApplicationSaaSAppUnion interface {
	implementsZeroTrustAccessApplicationUpdateResponseSaaSApplicationSaaSApp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationUpdateResponseSaaSApplicationSaaSAppUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SAMLSaaSApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OIDCSaaSApp{}),
		},
	)
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationUpdateResponseSaaSApplicationSaaSAppAuthType string

const (
	AccessApplicationUpdateResponseSaaSApplicationSaaSAppAuthTypeSAML AccessApplicationUpdateResponseSaaSApplicationSaaSAppAuthType = "saml"
	AccessApplicationUpdateResponseSaaSApplicationSaaSAppAuthTypeOIDC AccessApplicationUpdateResponseSaaSApplicationSaaSAppAuthType = "oidc"
)

func (r AccessApplicationUpdateResponseSaaSApplicationSaaSAppAuthType) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateResponseSaaSApplicationSaaSAppAuthTypeSAML, AccessApplicationUpdateResponseSaaSApplicationSaaSAppAuthTypeOIDC:
		return true
	}
	return false
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateResponseSaaSApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                          `json:"mappings"`
	JSON     accessApplicationUpdateResponseSaaSApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationUpdateResponseSaaSApplicationSCIMConfigJSON contains the JSON
// metadata for the struct
// [AccessApplicationUpdateResponseSaaSApplicationSCIMConfig]
type accessApplicationUpdateResponseSaaSApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationUpdateResponseSaaSApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseSaaSApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                     `json:"token_url"`
	JSON     accessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthentication]
type accessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

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
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool                `json:"path_cookie_attribute"`
	Policies            []ApplicationPolicy `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfig `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomains `json:"self_hosted_domains"`
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
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AUD                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	Policies                 apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SCIMConfig               apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                                `json:"mappings"`
	JSON     accessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfig]
type accessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                           `json:"token_url"`
	JSON     accessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthentication]
type accessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationUpdateResponseBrowserVNCApplication struct {
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
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool                `json:"path_cookie_attribute"`
	Policies            []ApplicationPolicy `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfig `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomains `json:"self_hosted_domains"`
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
	JSON      accessApplicationUpdateResponseBrowserVNCApplicationJSON `json:"-"`
}

// accessApplicationUpdateResponseBrowserVNCApplicationJSON contains the JSON
// metadata for the struct [AccessApplicationUpdateResponseBrowserVNCApplication]
type accessApplicationUpdateResponseBrowserVNCApplicationJSON struct {
	Domain                   apijson.Field
	Type                     apijson.Field
	ID                       apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AUD                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	Policies                 apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SCIMConfig               apijson.Field
	SelfHostedDomains        apijson.Field
	ServiceAuth401Redirect   apijson.Field
	SessionDuration          apijson.Field
	SkipInterstitial         apijson.Field
	Tags                     apijson.Field
	UpdatedAt                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccessApplicationUpdateResponseBrowserVNCApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseBrowserVNCApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationUpdateResponseBrowserVNCApplication) implementsZeroTrustAccessApplicationUpdateResponse() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                                `json:"mappings"`
	JSON     accessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfig]
type accessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                           `json:"token_url"`
	JSON     accessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthentication]
type accessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationUpdateResponseAppLauncherApplication struct {
	// The application type.
	Type ApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// The name of the application.
	Name     string              `json:"name"`
	Policies []ApplicationPolicy `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SCIMConfig             apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                                 `json:"mappings"`
	JSON     accessApplicationUpdateResponseAppLauncherApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationUpdateResponseAppLauncherApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfig]
type accessApplicationUpdateResponseAppLauncherApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseAppLauncherApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                            `json:"token_url"`
	JSON     accessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthentication]
type accessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type ApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// The name of the application.
	Name     string              `json:"name"`
	Policies []ApplicationPolicy `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SCIMConfig             apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                                                 `json:"mappings"`
	JSON     accessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfig]
type accessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                                            `json:"token_url"`
	JSON     accessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication]
type accessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationUpdateResponseBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type ApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// The name of the application.
	Name     string              `json:"name"`
	Policies []ApplicationPolicy `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SCIMConfig             apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                                                 `json:"mappings"`
	JSON     accessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfig]
type accessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                                            `json:"token_url"`
	JSON     accessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication]
type accessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationUpdateResponseBookmarkApplication struct {
	// UUID
	ID string `json:"id"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD       string    `json:"aud"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The URL or domain of the bookmark.
	Domain string `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name string `json:"name"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationUpdateResponseBookmarkApplicationSCIMConfig `json:"scim_config"`
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
	AUD                apijson.Field
	CreatedAt          apijson.Field
	Domain             apijson.Field
	LogoURL            apijson.Field
	Name               apijson.Field
	SCIMConfig         apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateResponseBookmarkApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                              `json:"mappings"`
	JSON     accessApplicationUpdateResponseBookmarkApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationUpdateResponseBookmarkApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationUpdateResponseBookmarkApplicationSCIMConfig]
type accessApplicationUpdateResponseBookmarkApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationUpdateResponseBookmarkApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUpdateResponseBookmarkApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                         `json:"token_url"`
	JSON     accessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthentication]
type accessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationListResponse struct {
	// Audience tag.
	AUD       string    `json:"aud"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	ID string `json:"id"`
	// This field can have the runtime type of
	// [AccessApplicationListResponseSelfHostedApplicationSCIMConfig],
	// [AccessApplicationListResponseSaaSApplicationSCIMConfig],
	// [AccessApplicationListResponseBrowserSSHApplicationSCIMConfig],
	// [AccessApplicationListResponseBrowserVNCApplicationSCIMConfig],
	// [AccessApplicationListResponseAppLauncherApplicationSCIMConfig],
	// [AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfig],
	// [AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfig],
	// [AccessApplicationListResponseBookmarkApplicationSCIMConfig].
	SCIMConfig interface{} `json:"scim_config,required"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP bool `json:"allow_authenticate_via_warp"`
	// This field can have the runtime type of [[]AllowedIdPs].
	AllowedIdPs interface{} `json:"allowed_idps,required"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage string `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL string `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL string `json:"custom_non_identity_deny_url"`
	// This field can have the runtime type of [[]string].
	CustomPages interface{} `json:"custom_pages,required"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// This field can have the runtime type of [[]SelfHostedDomains].
	SelfHostedDomains interface{} `json:"self_hosted_domains,required"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect bool `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial bool `json:"skip_interstitial"`
	// This field can have the runtime type of [[]string].
	Tags interface{} `json:"tags,required"`
	// The application type.
	Type string `json:"type"`
	// This field can have the runtime type of [[]ApplicationPolicy].
	Policies interface{} `json:"policies,required"`
	// This field can have the runtime type of
	// [AccessApplicationListResponseSaaSApplicationSaaSApp].
	SaaSApp interface{}                       `json:"saas_app,required"`
	JSON    accessApplicationListResponseJSON `json:"-"`
	union   AccessApplicationListResponseUnion
}

// accessApplicationListResponseJSON contains the JSON metadata for the struct
// [AccessApplicationListResponse]
type accessApplicationListResponseJSON struct {
	AUD                      apijson.Field
	CreatedAt                apijson.Field
	ID                       apijson.Field
	SCIMConfig               apijson.Field
	UpdatedAt                apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	Domain                   apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SelfHostedDomains        apijson.Field
	ServiceAuth401Redirect   apijson.Field
	SessionDuration          apijson.Field
	SkipInterstitial         apijson.Field
	Tags                     apijson.Field
	Type                     apijson.Field
	Policies                 apijson.Field
	SaaSApp                  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r accessApplicationListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessApplicationListResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.AccessApplicationListResponseSelfHostedApplication],
// [zero_trust.AccessApplicationListResponseSaaSApplication],
// [zero_trust.AccessApplicationListResponseBrowserSSHApplication],
// [zero_trust.AccessApplicationListResponseBrowserVNCApplication],
// [zero_trust.AccessApplicationListResponseAppLauncherApplication],
// [zero_trust.AccessApplicationListResponseDeviceEnrollmentPermissionsApplication],
// [zero_trust.AccessApplicationListResponseBrowserIsolationPermissionsApplication],
// [zero_trust.AccessApplicationListResponseBookmarkApplication].
func (r AccessApplicationListResponse) AsUnion() AccessApplicationListResponseUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.AccessApplicationListResponseSelfHostedApplication],
// [zero_trust.AccessApplicationListResponseSaaSApplication],
// [zero_trust.AccessApplicationListResponseBrowserSSHApplication],
// [zero_trust.AccessApplicationListResponseBrowserVNCApplication],
// [zero_trust.AccessApplicationListResponseAppLauncherApplication],
// [zero_trust.AccessApplicationListResponseDeviceEnrollmentPermissionsApplication],
// [zero_trust.AccessApplicationListResponseBrowserIsolationPermissionsApplication]
// or [zero_trust.AccessApplicationListResponseBookmarkApplication].
type AccessApplicationListResponseUnion interface {
	implementsZeroTrustAccessApplicationListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationListResponseUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AccessApplicationListResponseBrowserVNCApplication{}),
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
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool                `json:"path_cookie_attribute"`
	Policies            []ApplicationPolicy `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationListResponseSelfHostedApplicationSCIMConfig `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomains `json:"self_hosted_domains"`
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
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AUD                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	Policies                 apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SCIMConfig               apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationListResponseSelfHostedApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                              `json:"mappings"`
	JSON     accessApplicationListResponseSelfHostedApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationListResponseSelfHostedApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationListResponseSelfHostedApplicationSCIMConfig]
type accessApplicationListResponseSelfHostedApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationListResponseSelfHostedApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseSelfHostedApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                         `json:"token_url"`
	JSON     accessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthentication]
type accessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationListResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationListResponseSaaSApplication struct {
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages []string `json:"custom_pages"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name     string                                              `json:"name"`
	Policies []ApplicationPolicy                                 `json:"policies"`
	SaaSApp  AccessApplicationListResponseSaaSApplicationSaaSApp `json:"saas_app"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationListResponseSaaSApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AppLauncherVisible     apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	CustomPages            apijson.Field
	LogoURL                apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SaaSApp                apijson.Field
	SCIMConfig             apijson.Field
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

type AccessApplicationListResponseSaaSApplicationSaaSApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType AccessApplicationListResponseSaaSApplicationSaaSAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string    `json:"consumer_service_url"`
	CreatedAt          time.Time `json:"created_at" format:"date-time"`
	// This field can have the runtime type of [SAMLSaaSAppCustomAttributes].
	CustomAttributes interface{} `json:"custom_attributes,required"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdPEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat SaaSAppNameIDFormat `json:"name_id_format"`
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
	SAMLAttributeTransformJsonata string `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SPEntityID string `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint string    `json:"sso_endpoint"`
	UpdatedAt   time.Time `json:"updated_at" format:"date-time"`
	// The lifetime of the OIDC Access Token after creation. Valid units are m,h. Must
	// be greater than or equal to 1m and less than or equal to 24h.
	AccessTokenLifetime string `json:"access_token_lifetime"`
	// If client secret should be required on the token endpoint when
	// authorization_code_with_pkce grant is used.
	AllowPKCEWithoutClientSecret bool `json:"allow_pkce_without_client_secret"`
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [OIDCSaaSAppCustomClaims].
	CustomClaims interface{} `json:"custom_claims,required"`
	// This field can have the runtime type of [[]OIDCSaaSAppGrantType].
	GrantTypes interface{} `json:"grant_types,required"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// This field can have the runtime type of [OIDCSaaSAppHybridAndImplicitOptions].
	HybridAndImplicitOptions interface{} `json:"hybrid_and_implicit_options,required"`
	// This field can have the runtime type of [[]string].
	RedirectURIs interface{} `json:"redirect_uris,required"`
	// This field can have the runtime type of [OIDCSaaSAppRefreshTokenOptions].
	RefreshTokenOptions interface{} `json:"refresh_token_options,required"`
	// This field can have the runtime type of [[]OIDCSaaSAppScope].
	Scopes interface{}                                             `json:"scopes,required"`
	JSON   accessApplicationListResponseSaaSApplicationSaaSAppJSON `json:"-"`
	union  AccessApplicationListResponseSaaSApplicationSaaSAppUnion
}

// accessApplicationListResponseSaaSApplicationSaaSAppJSON contains the JSON
// metadata for the struct [AccessApplicationListResponseSaaSApplicationSaaSApp]
type accessApplicationListResponseSaaSApplicationSaaSAppJSON struct {
	AuthType                      apijson.Field
	ConsumerServiceURL            apijson.Field
	CreatedAt                     apijson.Field
	CustomAttributes              apijson.Field
	DefaultRelayState             apijson.Field
	IdPEntityID                   apijson.Field
	NameIDFormat                  apijson.Field
	NameIDTransformJsonata        apijson.Field
	PublicKey                     apijson.Field
	SAMLAttributeTransformJsonata apijson.Field
	SPEntityID                    apijson.Field
	SSOEndpoint                   apijson.Field
	UpdatedAt                     apijson.Field
	AccessTokenLifetime           apijson.Field
	AllowPKCEWithoutClientSecret  apijson.Field
	AppLauncherURL                apijson.Field
	ClientID                      apijson.Field
	ClientSecret                  apijson.Field
	CustomClaims                  apijson.Field
	GrantTypes                    apijson.Field
	GroupFilterRegex              apijson.Field
	HybridAndImplicitOptions      apijson.Field
	RedirectURIs                  apijson.Field
	RefreshTokenOptions           apijson.Field
	Scopes                        apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r accessApplicationListResponseSaaSApplicationSaaSAppJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationListResponseSaaSApplicationSaaSApp) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationListResponseSaaSApplicationSaaSApp{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessApplicationListResponseSaaSApplicationSaaSAppUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [zero_trust.SAMLSaaSApp],
// [zero_trust.OIDCSaaSApp].
func (r AccessApplicationListResponseSaaSApplicationSaaSApp) AsUnion() AccessApplicationListResponseSaaSApplicationSaaSAppUnion {
	return r.union
}

// Union satisfied by [zero_trust.SAMLSaaSApp] or [zero_trust.OIDCSaaSApp].
type AccessApplicationListResponseSaaSApplicationSaaSAppUnion interface {
	implementsZeroTrustAccessApplicationListResponseSaaSApplicationSaaSApp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationListResponseSaaSApplicationSaaSAppUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SAMLSaaSApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OIDCSaaSApp{}),
		},
	)
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationListResponseSaaSApplicationSaaSAppAuthType string

const (
	AccessApplicationListResponseSaaSApplicationSaaSAppAuthTypeSAML AccessApplicationListResponseSaaSApplicationSaaSAppAuthType = "saml"
	AccessApplicationListResponseSaaSApplicationSaaSAppAuthTypeOIDC AccessApplicationListResponseSaaSApplicationSaaSAppAuthType = "oidc"
)

func (r AccessApplicationListResponseSaaSApplicationSaaSAppAuthType) IsKnown() bool {
	switch r {
	case AccessApplicationListResponseSaaSApplicationSaaSAppAuthTypeSAML, AccessApplicationListResponseSaaSApplicationSaaSAppAuthTypeOIDC:
		return true
	}
	return false
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationListResponseSaaSApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationListResponseSaaSApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                        `json:"mappings"`
	JSON     accessApplicationListResponseSaaSApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationListResponseSaaSApplicationSCIMConfigJSON contains the JSON
// metadata for the struct [AccessApplicationListResponseSaaSApplicationSCIMConfig]
type accessApplicationListResponseSaaSApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationListResponseSaaSApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseSaaSApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationListResponseSaaSApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                   `json:"token_url"`
	JSON     accessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationListResponseSaaSApplicationSCIMConfigAuthentication]
type accessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationListResponseSaaSApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationListResponseSaaSApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationListResponseSaaSApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationListResponseSaaSApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationListResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

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
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool                `json:"path_cookie_attribute"`
	Policies            []ApplicationPolicy `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationListResponseBrowserSSHApplicationSCIMConfig `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomains `json:"self_hosted_domains"`
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
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AUD                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	Policies                 apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SCIMConfig               apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationListResponseBrowserSSHApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                              `json:"mappings"`
	JSON     accessApplicationListResponseBrowserSSHApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationListResponseBrowserSSHApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationListResponseBrowserSSHApplicationSCIMConfig]
type accessApplicationListResponseBrowserSSHApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationListResponseBrowserSSHApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseBrowserSSHApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                         `json:"token_url"`
	JSON     accessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthentication]
type accessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationListResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationListResponseBrowserVNCApplication struct {
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
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool                `json:"path_cookie_attribute"`
	Policies            []ApplicationPolicy `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationListResponseBrowserVNCApplicationSCIMConfig `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomains `json:"self_hosted_domains"`
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
	JSON      accessApplicationListResponseBrowserVNCApplicationJSON `json:"-"`
}

// accessApplicationListResponseBrowserVNCApplicationJSON contains the JSON
// metadata for the struct [AccessApplicationListResponseBrowserVNCApplication]
type accessApplicationListResponseBrowserVNCApplicationJSON struct {
	Domain                   apijson.Field
	Type                     apijson.Field
	ID                       apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AUD                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	Policies                 apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SCIMConfig               apijson.Field
	SelfHostedDomains        apijson.Field
	ServiceAuth401Redirect   apijson.Field
	SessionDuration          apijson.Field
	SkipInterstitial         apijson.Field
	Tags                     apijson.Field
	UpdatedAt                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccessApplicationListResponseBrowserVNCApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseBrowserVNCApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationListResponseBrowserVNCApplication) implementsZeroTrustAccessApplicationListResponse() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationListResponseBrowserVNCApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                              `json:"mappings"`
	JSON     accessApplicationListResponseBrowserVNCApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationListResponseBrowserVNCApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationListResponseBrowserVNCApplicationSCIMConfig]
type accessApplicationListResponseBrowserVNCApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationListResponseBrowserVNCApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseBrowserVNCApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                         `json:"token_url"`
	JSON     accessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthentication]
type accessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationListResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationListResponseAppLauncherApplication struct {
	// The application type.
	Type ApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// The name of the application.
	Name     string              `json:"name"`
	Policies []ApplicationPolicy `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationListResponseAppLauncherApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SCIMConfig             apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationListResponseAppLauncherApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                               `json:"mappings"`
	JSON     accessApplicationListResponseAppLauncherApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationListResponseAppLauncherApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationListResponseAppLauncherApplicationSCIMConfig]
type accessApplicationListResponseAppLauncherApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationListResponseAppLauncherApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseAppLauncherApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                          `json:"token_url"`
	JSON     accessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthentication]
type accessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationListResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationListResponseDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type ApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// The name of the application.
	Name     string              `json:"name"`
	Policies []ApplicationPolicy `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SCIMConfig             apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                                               `json:"mappings"`
	JSON     accessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON
// contains the JSON metadata for the struct
// [AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfig]
type accessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                                          `json:"token_url"`
	JSON     accessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication]
type accessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationListResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationListResponseBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type ApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// The name of the application.
	Name     string              `json:"name"`
	Policies []ApplicationPolicy `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SCIMConfig             apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                                               `json:"mappings"`
	JSON     accessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON
// contains the JSON metadata for the struct
// [AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfig]
type accessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                                          `json:"token_url"`
	JSON     accessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication]
type accessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationListResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationListResponseBookmarkApplication struct {
	// UUID
	ID string `json:"id"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD       string    `json:"aud"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The URL or domain of the bookmark.
	Domain string `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name string `json:"name"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationListResponseBookmarkApplicationSCIMConfig `json:"scim_config"`
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
	AUD                apijson.Field
	CreatedAt          apijson.Field
	Domain             apijson.Field
	LogoURL            apijson.Field
	Name               apijson.Field
	SCIMConfig         apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationListResponseBookmarkApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                            `json:"mappings"`
	JSON     accessApplicationListResponseBookmarkApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationListResponseBookmarkApplicationSCIMConfigJSON contains the JSON
// metadata for the struct
// [AccessApplicationListResponseBookmarkApplicationSCIMConfig]
type accessApplicationListResponseBookmarkApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationListResponseBookmarkApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationListResponseBookmarkApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                       `json:"token_url"`
	JSON     accessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthentication]
type accessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationListResponseBookmarkApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationListResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
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

type AccessApplicationGetResponse struct {
	// Audience tag.
	AUD       string    `json:"aud"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	ID string `json:"id"`
	// This field can have the runtime type of
	// [AccessApplicationGetResponseSelfHostedApplicationSCIMConfig],
	// [AccessApplicationGetResponseSaaSApplicationSCIMConfig],
	// [AccessApplicationGetResponseBrowserSSHApplicationSCIMConfig],
	// [AccessApplicationGetResponseBrowserVNCApplicationSCIMConfig],
	// [AccessApplicationGetResponseAppLauncherApplicationSCIMConfig],
	// [AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfig],
	// [AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfig],
	// [AccessApplicationGetResponseBookmarkApplicationSCIMConfig].
	SCIMConfig interface{} `json:"scim_config,required"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP bool `json:"allow_authenticate_via_warp"`
	// This field can have the runtime type of [[]AllowedIdPs].
	AllowedIdPs interface{} `json:"allowed_idps,required"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage string `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL string `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL string `json:"custom_non_identity_deny_url"`
	// This field can have the runtime type of [[]string].
	CustomPages interface{} `json:"custom_pages,required"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// This field can have the runtime type of [[]SelfHostedDomains].
	SelfHostedDomains interface{} `json:"self_hosted_domains,required"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect bool `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial bool `json:"skip_interstitial"`
	// This field can have the runtime type of [[]string].
	Tags interface{} `json:"tags,required"`
	// The application type.
	Type string `json:"type"`
	// This field can have the runtime type of [[]ApplicationPolicy].
	Policies interface{} `json:"policies,required"`
	// This field can have the runtime type of
	// [AccessApplicationGetResponseSaaSApplicationSaaSApp].
	SaaSApp interface{}                      `json:"saas_app,required"`
	JSON    accessApplicationGetResponseJSON `json:"-"`
	union   AccessApplicationGetResponseUnion
}

// accessApplicationGetResponseJSON contains the JSON metadata for the struct
// [AccessApplicationGetResponse]
type accessApplicationGetResponseJSON struct {
	AUD                      apijson.Field
	CreatedAt                apijson.Field
	ID                       apijson.Field
	SCIMConfig               apijson.Field
	UpdatedAt                apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	Domain                   apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SelfHostedDomains        apijson.Field
	ServiceAuth401Redirect   apijson.Field
	SessionDuration          apijson.Field
	SkipInterstitial         apijson.Field
	Tags                     apijson.Field
	Type                     apijson.Field
	Policies                 apijson.Field
	SaaSApp                  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r accessApplicationGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessApplicationGetResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.AccessApplicationGetResponseSelfHostedApplication],
// [zero_trust.AccessApplicationGetResponseSaaSApplication],
// [zero_trust.AccessApplicationGetResponseBrowserSSHApplication],
// [zero_trust.AccessApplicationGetResponseBrowserVNCApplication],
// [zero_trust.AccessApplicationGetResponseAppLauncherApplication],
// [zero_trust.AccessApplicationGetResponseDeviceEnrollmentPermissionsApplication],
// [zero_trust.AccessApplicationGetResponseBrowserIsolationPermissionsApplication],
// [zero_trust.AccessApplicationGetResponseBookmarkApplication].
func (r AccessApplicationGetResponse) AsUnion() AccessApplicationGetResponseUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.AccessApplicationGetResponseSelfHostedApplication],
// [zero_trust.AccessApplicationGetResponseSaaSApplication],
// [zero_trust.AccessApplicationGetResponseBrowserSSHApplication],
// [zero_trust.AccessApplicationGetResponseBrowserVNCApplication],
// [zero_trust.AccessApplicationGetResponseAppLauncherApplication],
// [zero_trust.AccessApplicationGetResponseDeviceEnrollmentPermissionsApplication],
// [zero_trust.AccessApplicationGetResponseBrowserIsolationPermissionsApplication]
// or [zero_trust.AccessApplicationGetResponseBookmarkApplication].
type AccessApplicationGetResponseUnion interface {
	implementsZeroTrustAccessApplicationGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationGetResponseUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AccessApplicationGetResponseBrowserVNCApplication{}),
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
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool                `json:"path_cookie_attribute"`
	Policies            []ApplicationPolicy `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationGetResponseSelfHostedApplicationSCIMConfig `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomains `json:"self_hosted_domains"`
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
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AUD                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	Policies                 apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SCIMConfig               apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationGetResponseSelfHostedApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                             `json:"mappings"`
	JSON     accessApplicationGetResponseSelfHostedApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationGetResponseSelfHostedApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationGetResponseSelfHostedApplicationSCIMConfig]
type accessApplicationGetResponseSelfHostedApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationGetResponseSelfHostedApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseSelfHostedApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                        `json:"token_url"`
	JSON     accessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthentication]
type accessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationGetResponseSelfHostedApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationGetResponseSaaSApplication struct {
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages []string `json:"custom_pages"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name     string                                             `json:"name"`
	Policies []ApplicationPolicy                                `json:"policies"`
	SaaSApp  AccessApplicationGetResponseSaaSApplicationSaaSApp `json:"saas_app"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationGetResponseSaaSApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AppLauncherVisible     apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	CustomPages            apijson.Field
	LogoURL                apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SaaSApp                apijson.Field
	SCIMConfig             apijson.Field
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

type AccessApplicationGetResponseSaaSApplicationSaaSApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType AccessApplicationGetResponseSaaSApplicationSaaSAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string    `json:"consumer_service_url"`
	CreatedAt          time.Time `json:"created_at" format:"date-time"`
	// This field can have the runtime type of [SAMLSaaSAppCustomAttributes].
	CustomAttributes interface{} `json:"custom_attributes,required"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdPEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat SaaSAppNameIDFormat `json:"name_id_format"`
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
	SAMLAttributeTransformJsonata string `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SPEntityID string `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint string    `json:"sso_endpoint"`
	UpdatedAt   time.Time `json:"updated_at" format:"date-time"`
	// The lifetime of the OIDC Access Token after creation. Valid units are m,h. Must
	// be greater than or equal to 1m and less than or equal to 24h.
	AccessTokenLifetime string `json:"access_token_lifetime"`
	// If client secret should be required on the token endpoint when
	// authorization_code_with_pkce grant is used.
	AllowPKCEWithoutClientSecret bool `json:"allow_pkce_without_client_secret"`
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [OIDCSaaSAppCustomClaims].
	CustomClaims interface{} `json:"custom_claims,required"`
	// This field can have the runtime type of [[]OIDCSaaSAppGrantType].
	GrantTypes interface{} `json:"grant_types,required"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// This field can have the runtime type of [OIDCSaaSAppHybridAndImplicitOptions].
	HybridAndImplicitOptions interface{} `json:"hybrid_and_implicit_options,required"`
	// This field can have the runtime type of [[]string].
	RedirectURIs interface{} `json:"redirect_uris,required"`
	// This field can have the runtime type of [OIDCSaaSAppRefreshTokenOptions].
	RefreshTokenOptions interface{} `json:"refresh_token_options,required"`
	// This field can have the runtime type of [[]OIDCSaaSAppScope].
	Scopes interface{}                                            `json:"scopes,required"`
	JSON   accessApplicationGetResponseSaaSApplicationSaaSAppJSON `json:"-"`
	union  AccessApplicationGetResponseSaaSApplicationSaaSAppUnion
}

// accessApplicationGetResponseSaaSApplicationSaaSAppJSON contains the JSON
// metadata for the struct [AccessApplicationGetResponseSaaSApplicationSaaSApp]
type accessApplicationGetResponseSaaSApplicationSaaSAppJSON struct {
	AuthType                      apijson.Field
	ConsumerServiceURL            apijson.Field
	CreatedAt                     apijson.Field
	CustomAttributes              apijson.Field
	DefaultRelayState             apijson.Field
	IdPEntityID                   apijson.Field
	NameIDFormat                  apijson.Field
	NameIDTransformJsonata        apijson.Field
	PublicKey                     apijson.Field
	SAMLAttributeTransformJsonata apijson.Field
	SPEntityID                    apijson.Field
	SSOEndpoint                   apijson.Field
	UpdatedAt                     apijson.Field
	AccessTokenLifetime           apijson.Field
	AllowPKCEWithoutClientSecret  apijson.Field
	AppLauncherURL                apijson.Field
	ClientID                      apijson.Field
	ClientSecret                  apijson.Field
	CustomClaims                  apijson.Field
	GrantTypes                    apijson.Field
	GroupFilterRegex              apijson.Field
	HybridAndImplicitOptions      apijson.Field
	RedirectURIs                  apijson.Field
	RefreshTokenOptions           apijson.Field
	Scopes                        apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r accessApplicationGetResponseSaaSApplicationSaaSAppJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationGetResponseSaaSApplicationSaaSApp) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationGetResponseSaaSApplicationSaaSApp{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessApplicationGetResponseSaaSApplicationSaaSAppUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [zero_trust.SAMLSaaSApp],
// [zero_trust.OIDCSaaSApp].
func (r AccessApplicationGetResponseSaaSApplicationSaaSApp) AsUnion() AccessApplicationGetResponseSaaSApplicationSaaSAppUnion {
	return r.union
}

// Union satisfied by [zero_trust.SAMLSaaSApp] or [zero_trust.OIDCSaaSApp].
type AccessApplicationGetResponseSaaSApplicationSaaSAppUnion interface {
	implementsZeroTrustAccessApplicationGetResponseSaaSApplicationSaaSApp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationGetResponseSaaSApplicationSaaSAppUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SAMLSaaSApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OIDCSaaSApp{}),
		},
	)
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationGetResponseSaaSApplicationSaaSAppAuthType string

const (
	AccessApplicationGetResponseSaaSApplicationSaaSAppAuthTypeSAML AccessApplicationGetResponseSaaSApplicationSaaSAppAuthType = "saml"
	AccessApplicationGetResponseSaaSApplicationSaaSAppAuthTypeOIDC AccessApplicationGetResponseSaaSApplicationSaaSAppAuthType = "oidc"
)

func (r AccessApplicationGetResponseSaaSApplicationSaaSAppAuthType) IsKnown() bool {
	switch r {
	case AccessApplicationGetResponseSaaSApplicationSaaSAppAuthTypeSAML, AccessApplicationGetResponseSaaSApplicationSaaSAppAuthTypeOIDC:
		return true
	}
	return false
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationGetResponseSaaSApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                       `json:"mappings"`
	JSON     accessApplicationGetResponseSaaSApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationGetResponseSaaSApplicationSCIMConfigJSON contains the JSON
// metadata for the struct [AccessApplicationGetResponseSaaSApplicationSCIMConfig]
type accessApplicationGetResponseSaaSApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationGetResponseSaaSApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseSaaSApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                  `json:"token_url"`
	JSON     accessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationJSON contains
// the JSON metadata for the struct
// [AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthentication]
type accessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationGetResponseSaaSApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationGetResponseSaaSApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

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
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool                `json:"path_cookie_attribute"`
	Policies            []ApplicationPolicy `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationGetResponseBrowserSSHApplicationSCIMConfig `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomains `json:"self_hosted_domains"`
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
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AUD                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	Policies                 apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SCIMConfig               apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationGetResponseBrowserSSHApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                             `json:"mappings"`
	JSON     accessApplicationGetResponseBrowserSSHApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationGetResponseBrowserSSHApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationGetResponseBrowserSSHApplicationSCIMConfig]
type accessApplicationGetResponseBrowserSSHApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationGetResponseBrowserSSHApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseBrowserSSHApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                        `json:"token_url"`
	JSON     accessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthentication]
type accessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationGetResponseBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationGetResponseBrowserVNCApplication struct {
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
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CORSHeaders            CORSHeaders `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool                `json:"path_cookie_attribute"`
	Policies            []ApplicationPolicy `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationGetResponseBrowserVNCApplicationSCIMConfig `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains []SelfHostedDomains `json:"self_hosted_domains"`
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
	JSON      accessApplicationGetResponseBrowserVNCApplicationJSON `json:"-"`
}

// accessApplicationGetResponseBrowserVNCApplicationJSON contains the JSON metadata
// for the struct [AccessApplicationGetResponseBrowserVNCApplication]
type accessApplicationGetResponseBrowserVNCApplicationJSON struct {
	Domain                   apijson.Field
	Type                     apijson.Field
	ID                       apijson.Field
	AllowAuthenticateViaWARP apijson.Field
	AllowedIdPs              apijson.Field
	AppLauncherVisible       apijson.Field
	AUD                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CORSHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	Policies                 apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SCIMConfig               apijson.Field
	SelfHostedDomains        apijson.Field
	ServiceAuth401Redirect   apijson.Field
	SessionDuration          apijson.Field
	SkipInterstitial         apijson.Field
	Tags                     apijson.Field
	UpdatedAt                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccessApplicationGetResponseBrowserVNCApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseBrowserVNCApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationGetResponseBrowserVNCApplication) implementsZeroTrustAccessApplicationGetResponse() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationGetResponseBrowserVNCApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                             `json:"mappings"`
	JSON     accessApplicationGetResponseBrowserVNCApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationGetResponseBrowserVNCApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationGetResponseBrowserVNCApplicationSCIMConfig]
type accessApplicationGetResponseBrowserVNCApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationGetResponseBrowserVNCApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseBrowserVNCApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                        `json:"token_url"`
	JSON     accessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthentication]
type accessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationGetResponseBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationGetResponseAppLauncherApplication struct {
	// The application type.
	Type ApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// The name of the application.
	Name     string              `json:"name"`
	Policies []ApplicationPolicy `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationGetResponseAppLauncherApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SCIMConfig             apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationGetResponseAppLauncherApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                              `json:"mappings"`
	JSON     accessApplicationGetResponseAppLauncherApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationGetResponseAppLauncherApplicationSCIMConfigJSON contains the
// JSON metadata for the struct
// [AccessApplicationGetResponseAppLauncherApplicationSCIMConfig]
type accessApplicationGetResponseAppLauncherApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationGetResponseAppLauncherApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseAppLauncherApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                         `json:"token_url"`
	JSON     accessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthentication]
type accessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationGetResponseAppLauncherApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationGetResponseDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type ApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// The name of the application.
	Name     string              `json:"name"`
	Policies []ApplicationPolicy `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SCIMConfig             apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                                              `json:"mappings"`
	JSON     accessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON
// contains the JSON metadata for the struct
// [AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfig]
type accessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                                         `json:"token_url"`
	JSON     accessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication]
type accessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationGetResponseDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationGetResponseBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type ApplicationType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs []AllowedIdPs `json:"allowed_idps"`
	// Audience tag.
	AUD string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time `json:"created_at" format:"date-time"`
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain string `json:"domain"`
	// The name of the application.
	Name     string              `json:"name"`
	Policies []ApplicationPolicy `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfig `json:"scim_config"`
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
	AllowedIdPs            apijson.Field
	AUD                    apijson.Field
	AutoRedirectToIdentity apijson.Field
	CreatedAt              apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	Policies               apijson.Field
	SCIMConfig             apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                                              `json:"mappings"`
	JSON     accessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON
// contains the JSON metadata for the struct
// [AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfig]
type accessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                                         `json:"token_url"`
	JSON     accessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication]
type accessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationGetResponseBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationGetResponseBookmarkApplication struct {
	// UUID
	ID string `json:"id"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	AUD       string    `json:"aud"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The URL or domain of the bookmark.
	Domain string `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name string `json:"name"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig AccessApplicationGetResponseBookmarkApplicationSCIMConfig `json:"scim_config"`
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
	AUD                apijson.Field
	CreatedAt          apijson.Field
	Domain             apijson.Field
	LogoURL            apijson.Field
	Name               apijson.Field
	SCIMConfig         apijson.Field
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

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationGetResponseBookmarkApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthentication `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []SCIMConfigMapping                                           `json:"mappings"`
	JSON     accessApplicationGetResponseBookmarkApplicationSCIMConfigJSON `json:"-"`
}

// accessApplicationGetResponseBookmarkApplicationSCIMConfigJSON contains the JSON
// metadata for the struct
// [AccessApplicationGetResponseBookmarkApplicationSCIMConfig]
type accessApplicationGetResponseBookmarkApplicationSCIMConfigJSON struct {
	IdPUID             apijson.Field
	RemoteURI          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationGetResponseBookmarkApplicationSCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationGetResponseBookmarkApplicationSCIMConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string                                                                      `json:"token_url"`
	JSON     accessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationJSON `json:"-"`
	union    AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationUnion
}

// accessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationJSON
// contains the JSON metadata for the struct
// [AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthentication]
type accessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationJSON struct {
	Password         apijson.Field
	Scheme           apijson.Field
	User             apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r accessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken],
// [zero_trust.SCIMConfigAuthenticationOauth2].
func (r AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthentication) AsUnion() AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasic],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerToken] or
// [zero_trust.SCIMConfigAuthenticationOauth2].
type AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SCIMConfigAuthenticationOauth2{}),
		},
	)
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationGetResponseBookmarkApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationRevokeTokensResponse = interface{}

type AccessApplicationNewParams struct {
	Body AccessApplicationNewParamsBodyUnion `json:"body,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r AccessApplicationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccessApplicationNewParamsBody struct {
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool]        `json:"allow_authenticate_via_warp"`
	AllowedIdPs              param.Field[interface{}] `json:"allowed_idps,required"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CORSHeaders            param.Field[CORSHeadersParam] `json:"cors_headers"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage param.Field[string] `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL param.Field[string] `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL param.Field[string]      `json:"custom_non_identity_deny_url"`
	CustomPages              param.Field[interface{}] `json:"custom_pages,required"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass param.Field[bool] `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string]      `json:"same_site_cookie_attribute"`
	SelfHostedDomains       param.Field[interface{}] `json:"self_hosted_domains,required"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect param.Field[bool] `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial param.Field[bool]        `json:"skip_interstitial"`
	Tags             param.Field[interface{}] `json:"tags,required"`
	// The application type.
	Type       param.Field[string]      `json:"type"`
	Policies   param.Field[interface{}] `json:"policies,required"`
	SCIMConfig param.Field[interface{}] `json:"scim_config,required"`
	SaaSApp    param.Field[interface{}] `json:"saas_app,required"`
}

func (r AccessApplicationNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBody) implementsZeroTrustAccessApplicationNewParamsBodyUnion() {}

// Satisfied by [zero_trust.AccessApplicationNewParamsBodySelfHostedApplication],
// [zero_trust.AccessApplicationNewParamsBodySaaSApplication],
// [zero_trust.AccessApplicationNewParamsBodyBrowserSSHApplication],
// [zero_trust.AccessApplicationNewParamsBodyBrowserVNCApplication],
// [zero_trust.AccessApplicationNewParamsBodyAppLauncherApplication],
// [zero_trust.AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplication],
// [zero_trust.AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplication],
// [zero_trust.AccessApplicationNewParamsBodyBookmarkApplication],
// [AccessApplicationNewParamsBody].
type AccessApplicationNewParamsBodyUnion interface {
	implementsZeroTrustAccessApplicationNewParamsBodyUnion()
}

type AccessApplicationNewParamsBodySelfHostedApplication struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdPsParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CORSHeaders            param.Field[CORSHeadersParam] `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass param.Field[bool] `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// The policies that will apply to the application, in ascending order of
	// precedence. Items can reference existing policies or create new policies
	// exclusive to the application.
	Policies param.Field[[]AccessApplicationNewParamsBodySelfHostedApplicationPolicyUnion] `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfig] `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]SelfHostedDomainsParam] `json:"self_hosted_domains"`
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

func (r AccessApplicationNewParamsBodySelfHostedApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodySelfHostedApplication) implementsZeroTrustAccessApplicationNewParamsBodyUnion() {
}

// A JSON that links a reusable policy to an application.
type AccessApplicationNewParamsBodySelfHostedApplicationPolicy struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence     param.Field[int64]       `json:"precedence"`
	ApprovalGroups param.Field[interface{}] `json:"approval_groups,required"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude,required"`
	Include  param.Field[interface{}] `json:"include,required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require,required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodySelfHostedApplicationPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodySelfHostedApplicationPolicy) ImplementsZeroTrustAccessApplicationNewParamsBodySelfHostedApplicationPolicyUnion() {
}

// A JSON that links a reusable policy to an application.
//
// Satisfied by
// [zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink],
// [shared.UnionString],
// [zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationPoliciesObject],
// [AccessApplicationNewParamsBodySelfHostedApplicationPolicy].
type AccessApplicationNewParamsBodySelfHostedApplicationPolicyUnion interface {
	ImplementsZeroTrustAccessApplicationNewParamsBodySelfHostedApplicationPolicyUnion()
}

// A JSON that links a reusable policy to an application.
type AccessApplicationNewParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
}

func (r AccessApplicationNewParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink) ImplementsZeroTrustAccessApplicationNewParamsBodySelfHostedApplicationPolicyUnion() {
}

type AccessApplicationNewParamsBodySelfHostedApplicationPoliciesObject struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodySelfHostedApplicationPoliciesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodySelfHostedApplicationPoliciesObject) ImplementsZeroTrustAccessApplicationNewParamsBodySelfHostedApplicationPolicyUnion() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthentication].
type AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationNewParamsBodySaaSApplication struct {
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdPsParam] `json:"allowed_idps"`
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
	Name param.Field[string] `json:"name"`
	// The policies that will apply to the application, in ascending order of
	// precedence. Items can reference existing policies or create new policies
	// exclusive to the application.
	Policies param.Field[[]AccessApplicationNewParamsBodySaaSApplicationPolicyUnion] `json:"policies"`
	SaaSApp  param.Field[AccessApplicationNewParamsBodySaaSApplicationSaaSAppUnion]  `json:"saas_app"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationNewParamsBodySaaSApplicationSCIMConfig] `json:"scim_config"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessApplicationNewParamsBodySaaSApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodySaaSApplication) implementsZeroTrustAccessApplicationNewParamsBodyUnion() {
}

// A JSON that links a reusable policy to an application.
type AccessApplicationNewParamsBodySaaSApplicationPolicy struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence     param.Field[int64]       `json:"precedence"`
	ApprovalGroups param.Field[interface{}] `json:"approval_groups,required"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude,required"`
	Include  param.Field[interface{}] `json:"include,required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require,required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodySaaSApplicationPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodySaaSApplicationPolicy) ImplementsZeroTrustAccessApplicationNewParamsBodySaaSApplicationPolicyUnion() {
}

// A JSON that links a reusable policy to an application.
//
// Satisfied by
// [zero_trust.AccessApplicationNewParamsBodySaaSApplicationPoliciesAccessAppPolicyLink],
// [shared.UnionString],
// [zero_trust.AccessApplicationNewParamsBodySaaSApplicationPoliciesObject],
// [AccessApplicationNewParamsBodySaaSApplicationPolicy].
type AccessApplicationNewParamsBodySaaSApplicationPolicyUnion interface {
	ImplementsZeroTrustAccessApplicationNewParamsBodySaaSApplicationPolicyUnion()
}

// A JSON that links a reusable policy to an application.
type AccessApplicationNewParamsBodySaaSApplicationPoliciesAccessAppPolicyLink struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
}

func (r AccessApplicationNewParamsBodySaaSApplicationPoliciesAccessAppPolicyLink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodySaaSApplicationPoliciesAccessAppPolicyLink) ImplementsZeroTrustAccessApplicationNewParamsBodySaaSApplicationPolicyUnion() {
}

type AccessApplicationNewParamsBodySaaSApplicationPoliciesObject struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodySaaSApplicationPoliciesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodySaaSApplicationPoliciesObject) ImplementsZeroTrustAccessApplicationNewParamsBodySaaSApplicationPolicyUnion() {
}

type AccessApplicationNewParamsBodySaaSApplicationSaaSApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[AccessApplicationNewParamsBodySaaSApplicationSaaSAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]      `json:"consumer_service_url"`
	CreatedAt          param.Field[time.Time]   `json:"created_at" format:"date-time"`
	CustomAttributes   param.Field[interface{}] `json:"custom_attributes,required"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdPEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[SaaSAppNameIDFormat] `json:"name_id_format"`
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
	SAMLAttributeTransformJsonata param.Field[string] `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SPEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string]    `json:"sso_endpoint"`
	UpdatedAt   param.Field[time.Time] `json:"updated_at" format:"date-time"`
	// The lifetime of the OIDC Access Token after creation. Valid units are m,h. Must
	// be greater than or equal to 1m and less than or equal to 24h.
	AccessTokenLifetime param.Field[string] `json:"access_token_lifetime"`
	// If client secret should be required on the token endpoint when
	// authorization_code_with_pkce grant is used.
	AllowPKCEWithoutClientSecret param.Field[bool] `json:"allow_pkce_without_client_secret"`
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string]      `json:"client_secret"`
	CustomClaims param.Field[interface{}] `json:"custom_claims,required"`
	GrantTypes   param.Field[interface{}] `json:"grant_types,required"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex         param.Field[string]      `json:"group_filter_regex"`
	HybridAndImplicitOptions param.Field[interface{}] `json:"hybrid_and_implicit_options,required"`
	RedirectURIs             param.Field[interface{}] `json:"redirect_uris,required"`
	RefreshTokenOptions      param.Field[interface{}] `json:"refresh_token_options,required"`
	Scopes                   param.Field[interface{}] `json:"scopes,required"`
}

func (r AccessApplicationNewParamsBodySaaSApplicationSaaSApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodySaaSApplicationSaaSApp) implementsZeroTrustAccessApplicationNewParamsBodySaaSApplicationSaaSAppUnion() {
}

// Satisfied by [zero_trust.SAMLSaaSAppParam], [zero_trust.OIDCSaaSAppParam],
// [AccessApplicationNewParamsBodySaaSApplicationSaaSApp].
type AccessApplicationNewParamsBodySaaSApplicationSaaSAppUnion interface {
	implementsZeroTrustAccessApplicationNewParamsBodySaaSApplicationSaaSAppUnion()
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationNewParamsBodySaaSApplicationSaaSAppAuthType string

const (
	AccessApplicationNewParamsBodySaaSApplicationSaaSAppAuthTypeSAML AccessApplicationNewParamsBodySaaSApplicationSaaSAppAuthType = "saml"
	AccessApplicationNewParamsBodySaaSApplicationSaaSAppAuthTypeOIDC AccessApplicationNewParamsBodySaaSApplicationSaaSAppAuthType = "oidc"
)

func (r AccessApplicationNewParamsBodySaaSApplicationSaaSAppAuthType) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsBodySaaSApplicationSaaSAppAuthTypeSAML, AccessApplicationNewParamsBodySaaSApplicationSaaSAppAuthTypeOIDC:
		return true
	}
	return false
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewParamsBodySaaSApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationNewParamsBodySaaSApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthentication].
type AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewParamsBodySaaSApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationNewParamsBodyBrowserSSHApplication struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdPsParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CORSHeaders            param.Field[CORSHeadersParam] `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass param.Field[bool] `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// The policies that will apply to the application, in ascending order of
	// precedence. Items can reference existing policies or create new policies
	// exclusive to the application.
	Policies param.Field[[]AccessApplicationNewParamsBodyBrowserSSHApplicationPolicyUnion] `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfig] `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]SelfHostedDomainsParam] `json:"self_hosted_domains"`
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

func (r AccessApplicationNewParamsBodyBrowserSSHApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBrowserSSHApplication) implementsZeroTrustAccessApplicationNewParamsBodyUnion() {
}

// A JSON that links a reusable policy to an application.
type AccessApplicationNewParamsBodyBrowserSSHApplicationPolicy struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence     param.Field[int64]       `json:"precedence"`
	ApprovalGroups param.Field[interface{}] `json:"approval_groups,required"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude,required"`
	Include  param.Field[interface{}] `json:"include,required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require,required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodyBrowserSSHApplicationPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBrowserSSHApplicationPolicy) ImplementsZeroTrustAccessApplicationNewParamsBodyBrowserSSHApplicationPolicyUnion() {
}

// A JSON that links a reusable policy to an application.
//
// Satisfied by
// [zero_trust.AccessApplicationNewParamsBodyBrowserSSHApplicationPoliciesAccessAppPolicyLink],
// [shared.UnionString],
// [zero_trust.AccessApplicationNewParamsBodyBrowserSSHApplicationPoliciesObject],
// [AccessApplicationNewParamsBodyBrowserSSHApplicationPolicy].
type AccessApplicationNewParamsBodyBrowserSSHApplicationPolicyUnion interface {
	ImplementsZeroTrustAccessApplicationNewParamsBodyBrowserSSHApplicationPolicyUnion()
}

// A JSON that links a reusable policy to an application.
type AccessApplicationNewParamsBodyBrowserSSHApplicationPoliciesAccessAppPolicyLink struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
}

func (r AccessApplicationNewParamsBodyBrowserSSHApplicationPoliciesAccessAppPolicyLink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBrowserSSHApplicationPoliciesAccessAppPolicyLink) ImplementsZeroTrustAccessApplicationNewParamsBodyBrowserSSHApplicationPolicyUnion() {
}

type AccessApplicationNewParamsBodyBrowserSSHApplicationPoliciesObject struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodyBrowserSSHApplicationPoliciesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBrowserSSHApplicationPoliciesObject) ImplementsZeroTrustAccessApplicationNewParamsBodyBrowserSSHApplicationPolicyUnion() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthentication].
type AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationNewParamsBodyBrowserVNCApplication struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdPsParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CORSHeaders            param.Field[CORSHeadersParam] `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass param.Field[bool] `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// The policies that will apply to the application, in ascending order of
	// precedence. Items can reference existing policies or create new policies
	// exclusive to the application.
	Policies param.Field[[]AccessApplicationNewParamsBodyBrowserVNCApplicationPolicyUnion] `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfig] `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]SelfHostedDomainsParam] `json:"self_hosted_domains"`
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

func (r AccessApplicationNewParamsBodyBrowserVNCApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBrowserVNCApplication) implementsZeroTrustAccessApplicationNewParamsBodyUnion() {
}

// A JSON that links a reusable policy to an application.
type AccessApplicationNewParamsBodyBrowserVNCApplicationPolicy struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence     param.Field[int64]       `json:"precedence"`
	ApprovalGroups param.Field[interface{}] `json:"approval_groups,required"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude,required"`
	Include  param.Field[interface{}] `json:"include,required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require,required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodyBrowserVNCApplicationPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBrowserVNCApplicationPolicy) ImplementsZeroTrustAccessApplicationNewParamsBodyBrowserVNCApplicationPolicyUnion() {
}

// A JSON that links a reusable policy to an application.
//
// Satisfied by
// [zero_trust.AccessApplicationNewParamsBodyBrowserVNCApplicationPoliciesAccessAppPolicyLink],
// [shared.UnionString],
// [zero_trust.AccessApplicationNewParamsBodyBrowserVNCApplicationPoliciesObject],
// [AccessApplicationNewParamsBodyBrowserVNCApplicationPolicy].
type AccessApplicationNewParamsBodyBrowserVNCApplicationPolicyUnion interface {
	ImplementsZeroTrustAccessApplicationNewParamsBodyBrowserVNCApplicationPolicyUnion()
}

// A JSON that links a reusable policy to an application.
type AccessApplicationNewParamsBodyBrowserVNCApplicationPoliciesAccessAppPolicyLink struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
}

func (r AccessApplicationNewParamsBodyBrowserVNCApplicationPoliciesAccessAppPolicyLink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBrowserVNCApplicationPoliciesAccessAppPolicyLink) ImplementsZeroTrustAccessApplicationNewParamsBodyBrowserVNCApplicationPolicyUnion() {
}

type AccessApplicationNewParamsBodyBrowserVNCApplicationPoliciesObject struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodyBrowserVNCApplicationPoliciesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBrowserVNCApplicationPoliciesObject) ImplementsZeroTrustAccessApplicationNewParamsBodyBrowserVNCApplicationPolicyUnion() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthentication].
type AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationNewParamsBodyAppLauncherApplication struct {
	// The application type.
	Type param.Field[ApplicationType] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdPsParam] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The policies that will apply to the application, in ascending order of
	// precedence. Items can reference existing policies or create new policies
	// exclusive to the application.
	Policies param.Field[[]AccessApplicationNewParamsBodyAppLauncherApplicationPolicyUnion] `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfig] `json:"scim_config"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodyAppLauncherApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyAppLauncherApplication) implementsZeroTrustAccessApplicationNewParamsBodyUnion() {
}

// A JSON that links a reusable policy to an application.
type AccessApplicationNewParamsBodyAppLauncherApplicationPolicy struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence     param.Field[int64]       `json:"precedence"`
	ApprovalGroups param.Field[interface{}] `json:"approval_groups,required"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude,required"`
	Include  param.Field[interface{}] `json:"include,required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require,required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodyAppLauncherApplicationPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyAppLauncherApplicationPolicy) ImplementsZeroTrustAccessApplicationNewParamsBodyAppLauncherApplicationPolicyUnion() {
}

// A JSON that links a reusable policy to an application.
//
// Satisfied by
// [zero_trust.AccessApplicationNewParamsBodyAppLauncherApplicationPoliciesAccessAppPolicyLink],
// [shared.UnionString],
// [zero_trust.AccessApplicationNewParamsBodyAppLauncherApplicationPoliciesObject],
// [AccessApplicationNewParamsBodyAppLauncherApplicationPolicy].
type AccessApplicationNewParamsBodyAppLauncherApplicationPolicyUnion interface {
	ImplementsZeroTrustAccessApplicationNewParamsBodyAppLauncherApplicationPolicyUnion()
}

// A JSON that links a reusable policy to an application.
type AccessApplicationNewParamsBodyAppLauncherApplicationPoliciesAccessAppPolicyLink struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
}

func (r AccessApplicationNewParamsBodyAppLauncherApplicationPoliciesAccessAppPolicyLink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyAppLauncherApplicationPoliciesAccessAppPolicyLink) ImplementsZeroTrustAccessApplicationNewParamsBodyAppLauncherApplicationPolicyUnion() {
}

type AccessApplicationNewParamsBodyAppLauncherApplicationPoliciesObject struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodyAppLauncherApplicationPoliciesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyAppLauncherApplicationPoliciesObject) ImplementsZeroTrustAccessApplicationNewParamsBodyAppLauncherApplicationPolicyUnion() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthentication].
type AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewParamsBodyAppLauncherApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type param.Field[ApplicationType] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdPsParam] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The policies that will apply to the application, in ascending order of
	// precedence. Items can reference existing policies or create new policies
	// exclusive to the application.
	Policies param.Field[[]AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion] `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfig] `json:"scim_config"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplication) implementsZeroTrustAccessApplicationNewParamsBodyUnion() {
}

// A JSON that links a reusable policy to an application.
type AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPolicy struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence     param.Field[int64]       `json:"precedence"`
	ApprovalGroups param.Field[interface{}] `json:"approval_groups,required"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude,required"`
	Include  param.Field[interface{}] `json:"include,required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require,required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPolicy) ImplementsZeroTrustAccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion() {
}

// A JSON that links a reusable policy to an application.
//
// Satisfied by
// [zero_trust.AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesAccessAppPolicyLink],
// [shared.UnionString],
// [zero_trust.AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesObject],
// [AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPolicy].
type AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion interface {
	ImplementsZeroTrustAccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion()
}

// A JSON that links a reusable policy to an application.
type AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesAccessAppPolicyLink struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
}

func (r AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesAccessAppPolicyLink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesAccessAppPolicyLink) ImplementsZeroTrustAccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion() {
}

type AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesObject struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesObject) ImplementsZeroTrustAccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication].
type AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type param.Field[ApplicationType] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdPsParam] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The policies that will apply to the application, in ascending order of
	// precedence. Items can reference existing policies or create new policies
	// exclusive to the application.
	Policies param.Field[[]AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion] `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfig] `json:"scim_config"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplication) implementsZeroTrustAccessApplicationNewParamsBodyUnion() {
}

// A JSON that links a reusable policy to an application.
type AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPolicy struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence     param.Field[int64]       `json:"precedence"`
	ApprovalGroups param.Field[interface{}] `json:"approval_groups,required"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude,required"`
	Include  param.Field[interface{}] `json:"include,required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require,required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPolicy) ImplementsZeroTrustAccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion() {
}

// A JSON that links a reusable policy to an application.
//
// Satisfied by
// [zero_trust.AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPoliciesAccessAppPolicyLink],
// [shared.UnionString],
// [zero_trust.AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPoliciesObject],
// [AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPolicy].
type AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion interface {
	ImplementsZeroTrustAccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion()
}

// A JSON that links a reusable policy to an application.
type AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPoliciesAccessAppPolicyLink struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
}

func (r AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPoliciesAccessAppPolicyLink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPoliciesAccessAppPolicyLink) ImplementsZeroTrustAccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion() {
}

type AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPoliciesObject struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPoliciesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPoliciesObject) ImplementsZeroTrustAccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthentication].
type AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationNewParamsBodyBookmarkApplication struct {
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// The URL or domain of the bookmark.
	Domain param.Field[string] `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name param.Field[string] `json:"name"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfig] `json:"scim_config"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessApplicationNewParamsBodyBookmarkApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBookmarkApplication) implementsZeroTrustAccessApplicationNewParamsBodyUnion() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthentication].
type AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessApplicationNewResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationNewResponse                `json:"result"`
	JSON    accessApplicationNewResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationNewResponseEnvelope]
type accessApplicationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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

type AccessApplicationUpdateParams struct {
	Body AccessApplicationUpdateParamsBodyUnion `json:"body,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r AccessApplicationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccessApplicationUpdateParamsBody struct {
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool]        `json:"allow_authenticate_via_warp"`
	AllowedIdPs              param.Field[interface{}] `json:"allowed_idps,required"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CORSHeaders            param.Field[CORSHeadersParam] `json:"cors_headers"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage param.Field[string] `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL param.Field[string] `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL param.Field[string]      `json:"custom_non_identity_deny_url"`
	CustomPages              param.Field[interface{}] `json:"custom_pages,required"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass param.Field[bool] `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string]      `json:"same_site_cookie_attribute"`
	SelfHostedDomains       param.Field[interface{}] `json:"self_hosted_domains,required"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect param.Field[bool] `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial param.Field[bool]        `json:"skip_interstitial"`
	Tags             param.Field[interface{}] `json:"tags,required"`
	// The application type.
	Type       param.Field[string]      `json:"type"`
	Policies   param.Field[interface{}] `json:"policies,required"`
	SCIMConfig param.Field[interface{}] `json:"scim_config,required"`
	SaaSApp    param.Field[interface{}] `json:"saas_app,required"`
}

func (r AccessApplicationUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBody) implementsZeroTrustAccessApplicationUpdateParamsBodyUnion() {
}

// Satisfied by
// [zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplication],
// [zero_trust.AccessApplicationUpdateParamsBodySaaSApplication],
// [zero_trust.AccessApplicationUpdateParamsBodyBrowserSSHApplication],
// [zero_trust.AccessApplicationUpdateParamsBodyBrowserVNCApplication],
// [zero_trust.AccessApplicationUpdateParamsBodyAppLauncherApplication],
// [zero_trust.AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplication],
// [zero_trust.AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplication],
// [zero_trust.AccessApplicationUpdateParamsBodyBookmarkApplication],
// [AccessApplicationUpdateParamsBody].
type AccessApplicationUpdateParamsBodyUnion interface {
	implementsZeroTrustAccessApplicationUpdateParamsBodyUnion()
}

type AccessApplicationUpdateParamsBodySelfHostedApplication struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdPsParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CORSHeaders            param.Field[CORSHeadersParam] `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass param.Field[bool] `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// The policies that will apply to the application, in ascending order of
	// precedence. Items can reference existing policies or create new policies
	// exclusive to the application.
	Policies param.Field[[]AccessApplicationUpdateParamsBodySelfHostedApplicationPolicyUnion] `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfig] `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]SelfHostedDomainsParam] `json:"self_hosted_domains"`
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

func (r AccessApplicationUpdateParamsBodySelfHostedApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodySelfHostedApplication) implementsZeroTrustAccessApplicationUpdateParamsBodyUnion() {
}

// A JSON that links a reusable policy to an application.
type AccessApplicationUpdateParamsBodySelfHostedApplicationPolicy struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence     param.Field[int64]       `json:"precedence"`
	ApprovalGroups param.Field[interface{}] `json:"approval_groups,required"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude,required"`
	Include  param.Field[interface{}] `json:"include,required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require,required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodySelfHostedApplicationPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodySelfHostedApplicationPolicy) ImplementsZeroTrustAccessApplicationUpdateParamsBodySelfHostedApplicationPolicyUnion() {
}

// A JSON that links a reusable policy to an application.
//
// Satisfied by
// [zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink],
// [shared.UnionString],
// [zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationPoliciesObject],
// [AccessApplicationUpdateParamsBodySelfHostedApplicationPolicy].
type AccessApplicationUpdateParamsBodySelfHostedApplicationPolicyUnion interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsBodySelfHostedApplicationPolicyUnion()
}

// A JSON that links a reusable policy to an application.
type AccessApplicationUpdateParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
}

func (r AccessApplicationUpdateParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink) ImplementsZeroTrustAccessApplicationUpdateParamsBodySelfHostedApplicationPolicyUnion() {
}

type AccessApplicationUpdateParamsBodySelfHostedApplicationPoliciesObject struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodySelfHostedApplicationPoliciesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodySelfHostedApplicationPoliciesObject) ImplementsZeroTrustAccessApplicationUpdateParamsBodySelfHostedApplicationPolicyUnion() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthentication].
type AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsBodySaaSApplication struct {
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdPsParam] `json:"allowed_idps"`
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
	Name param.Field[string] `json:"name"`
	// The policies that will apply to the application, in ascending order of
	// precedence. Items can reference existing policies or create new policies
	// exclusive to the application.
	Policies param.Field[[]AccessApplicationUpdateParamsBodySaaSApplicationPolicyUnion] `json:"policies"`
	SaaSApp  param.Field[AccessApplicationUpdateParamsBodySaaSApplicationSaaSAppUnion]  `json:"saas_app"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfig] `json:"scim_config"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessApplicationUpdateParamsBodySaaSApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodySaaSApplication) implementsZeroTrustAccessApplicationUpdateParamsBodyUnion() {
}

// A JSON that links a reusable policy to an application.
type AccessApplicationUpdateParamsBodySaaSApplicationPolicy struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence     param.Field[int64]       `json:"precedence"`
	ApprovalGroups param.Field[interface{}] `json:"approval_groups,required"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude,required"`
	Include  param.Field[interface{}] `json:"include,required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require,required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodySaaSApplicationPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodySaaSApplicationPolicy) ImplementsZeroTrustAccessApplicationUpdateParamsBodySaaSApplicationPolicyUnion() {
}

// A JSON that links a reusable policy to an application.
//
// Satisfied by
// [zero_trust.AccessApplicationUpdateParamsBodySaaSApplicationPoliciesAccessAppPolicyLink],
// [shared.UnionString],
// [zero_trust.AccessApplicationUpdateParamsBodySaaSApplicationPoliciesObject],
// [AccessApplicationUpdateParamsBodySaaSApplicationPolicy].
type AccessApplicationUpdateParamsBodySaaSApplicationPolicyUnion interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsBodySaaSApplicationPolicyUnion()
}

// A JSON that links a reusable policy to an application.
type AccessApplicationUpdateParamsBodySaaSApplicationPoliciesAccessAppPolicyLink struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
}

func (r AccessApplicationUpdateParamsBodySaaSApplicationPoliciesAccessAppPolicyLink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodySaaSApplicationPoliciesAccessAppPolicyLink) ImplementsZeroTrustAccessApplicationUpdateParamsBodySaaSApplicationPolicyUnion() {
}

type AccessApplicationUpdateParamsBodySaaSApplicationPoliciesObject struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodySaaSApplicationPoliciesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodySaaSApplicationPoliciesObject) ImplementsZeroTrustAccessApplicationUpdateParamsBodySaaSApplicationPolicyUnion() {
}

type AccessApplicationUpdateParamsBodySaaSApplicationSaaSApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[AccessApplicationUpdateParamsBodySaaSApplicationSaaSAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]      `json:"consumer_service_url"`
	CreatedAt          param.Field[time.Time]   `json:"created_at" format:"date-time"`
	CustomAttributes   param.Field[interface{}] `json:"custom_attributes,required"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdPEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[SaaSAppNameIDFormat] `json:"name_id_format"`
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
	SAMLAttributeTransformJsonata param.Field[string] `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SPEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string]    `json:"sso_endpoint"`
	UpdatedAt   param.Field[time.Time] `json:"updated_at" format:"date-time"`
	// The lifetime of the OIDC Access Token after creation. Valid units are m,h. Must
	// be greater than or equal to 1m and less than or equal to 24h.
	AccessTokenLifetime param.Field[string] `json:"access_token_lifetime"`
	// If client secret should be required on the token endpoint when
	// authorization_code_with_pkce grant is used.
	AllowPKCEWithoutClientSecret param.Field[bool] `json:"allow_pkce_without_client_secret"`
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string]      `json:"client_secret"`
	CustomClaims param.Field[interface{}] `json:"custom_claims,required"`
	GrantTypes   param.Field[interface{}] `json:"grant_types,required"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex         param.Field[string]      `json:"group_filter_regex"`
	HybridAndImplicitOptions param.Field[interface{}] `json:"hybrid_and_implicit_options,required"`
	RedirectURIs             param.Field[interface{}] `json:"redirect_uris,required"`
	RefreshTokenOptions      param.Field[interface{}] `json:"refresh_token_options,required"`
	Scopes                   param.Field[interface{}] `json:"scopes,required"`
}

func (r AccessApplicationUpdateParamsBodySaaSApplicationSaaSApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodySaaSApplicationSaaSApp) implementsZeroTrustAccessApplicationUpdateParamsBodySaaSApplicationSaaSAppUnion() {
}

// Satisfied by [zero_trust.SAMLSaaSAppParam], [zero_trust.OIDCSaaSAppParam],
// [AccessApplicationUpdateParamsBodySaaSApplicationSaaSApp].
type AccessApplicationUpdateParamsBodySaaSApplicationSaaSAppUnion interface {
	implementsZeroTrustAccessApplicationUpdateParamsBodySaaSApplicationSaaSAppUnion()
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessApplicationUpdateParamsBodySaaSApplicationSaaSAppAuthType string

const (
	AccessApplicationUpdateParamsBodySaaSApplicationSaaSAppAuthTypeSAML AccessApplicationUpdateParamsBodySaaSApplicationSaaSAppAuthType = "saml"
	AccessApplicationUpdateParamsBodySaaSApplicationSaaSAppAuthTypeOIDC AccessApplicationUpdateParamsBodySaaSApplicationSaaSAppAuthType = "oidc"
)

func (r AccessApplicationUpdateParamsBodySaaSApplicationSaaSAppAuthType) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsBodySaaSApplicationSaaSAppAuthTypeSAML, AccessApplicationUpdateParamsBodySaaSApplicationSaaSAppAuthTypeOIDC:
		return true
	}
	return false
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthentication].
type AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateParamsBodySaaSApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsBodyBrowserSSHApplication struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdPsParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CORSHeaders            param.Field[CORSHeadersParam] `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass param.Field[bool] `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// The policies that will apply to the application, in ascending order of
	// precedence. Items can reference existing policies or create new policies
	// exclusive to the application.
	Policies param.Field[[]AccessApplicationUpdateParamsBodyBrowserSSHApplicationPolicyUnion] `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfig] `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]SelfHostedDomainsParam] `json:"self_hosted_domains"`
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

func (r AccessApplicationUpdateParamsBodyBrowserSSHApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBrowserSSHApplication) implementsZeroTrustAccessApplicationUpdateParamsBodyUnion() {
}

// A JSON that links a reusable policy to an application.
type AccessApplicationUpdateParamsBodyBrowserSSHApplicationPolicy struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence     param.Field[int64]       `json:"precedence"`
	ApprovalGroups param.Field[interface{}] `json:"approval_groups,required"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude,required"`
	Include  param.Field[interface{}] `json:"include,required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require,required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodyBrowserSSHApplicationPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBrowserSSHApplicationPolicy) ImplementsZeroTrustAccessApplicationUpdateParamsBodyBrowserSSHApplicationPolicyUnion() {
}

// A JSON that links a reusable policy to an application.
//
// Satisfied by
// [zero_trust.AccessApplicationUpdateParamsBodyBrowserSSHApplicationPoliciesAccessAppPolicyLink],
// [shared.UnionString],
// [zero_trust.AccessApplicationUpdateParamsBodyBrowserSSHApplicationPoliciesObject],
// [AccessApplicationUpdateParamsBodyBrowserSSHApplicationPolicy].
type AccessApplicationUpdateParamsBodyBrowserSSHApplicationPolicyUnion interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsBodyBrowserSSHApplicationPolicyUnion()
}

// A JSON that links a reusable policy to an application.
type AccessApplicationUpdateParamsBodyBrowserSSHApplicationPoliciesAccessAppPolicyLink struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
}

func (r AccessApplicationUpdateParamsBodyBrowserSSHApplicationPoliciesAccessAppPolicyLink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBrowserSSHApplicationPoliciesAccessAppPolicyLink) ImplementsZeroTrustAccessApplicationUpdateParamsBodyBrowserSSHApplicationPolicyUnion() {
}

type AccessApplicationUpdateParamsBodyBrowserSSHApplicationPoliciesObject struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodyBrowserSSHApplicationPoliciesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBrowserSSHApplicationPoliciesObject) ImplementsZeroTrustAccessApplicationUpdateParamsBodyBrowserSSHApplicationPolicyUnion() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthentication].
type AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateParamsBodyBrowserSSHApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsBodyBrowserVNCApplication struct {
	// The primary hostname and path that Access will secure. If the app is visible in
	// the App Launcher dashboard, this is the domain that will be displayed.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdPsParam] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CORSHeaders            param.Field[CORSHeadersParam] `json:"cors_headers"`
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
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass param.Field[bool] `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// The policies that will apply to the application, in ascending order of
	// precedence. Items can reference existing policies or create new policies
	// exclusive to the application.
	Policies param.Field[[]AccessApplicationUpdateParamsBodyBrowserVNCApplicationPolicyUnion] `json:"policies"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfig] `json:"scim_config"`
	// List of domains that Access will secure.
	SelfHostedDomains param.Field[[]SelfHostedDomainsParam] `json:"self_hosted_domains"`
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

func (r AccessApplicationUpdateParamsBodyBrowserVNCApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBrowserVNCApplication) implementsZeroTrustAccessApplicationUpdateParamsBodyUnion() {
}

// A JSON that links a reusable policy to an application.
type AccessApplicationUpdateParamsBodyBrowserVNCApplicationPolicy struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence     param.Field[int64]       `json:"precedence"`
	ApprovalGroups param.Field[interface{}] `json:"approval_groups,required"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude,required"`
	Include  param.Field[interface{}] `json:"include,required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require,required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodyBrowserVNCApplicationPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBrowserVNCApplicationPolicy) ImplementsZeroTrustAccessApplicationUpdateParamsBodyBrowserVNCApplicationPolicyUnion() {
}

// A JSON that links a reusable policy to an application.
//
// Satisfied by
// [zero_trust.AccessApplicationUpdateParamsBodyBrowserVNCApplicationPoliciesAccessAppPolicyLink],
// [shared.UnionString],
// [zero_trust.AccessApplicationUpdateParamsBodyBrowserVNCApplicationPoliciesObject],
// [AccessApplicationUpdateParamsBodyBrowserVNCApplicationPolicy].
type AccessApplicationUpdateParamsBodyBrowserVNCApplicationPolicyUnion interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsBodyBrowserVNCApplicationPolicyUnion()
}

// A JSON that links a reusable policy to an application.
type AccessApplicationUpdateParamsBodyBrowserVNCApplicationPoliciesAccessAppPolicyLink struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
}

func (r AccessApplicationUpdateParamsBodyBrowserVNCApplicationPoliciesAccessAppPolicyLink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBrowserVNCApplicationPoliciesAccessAppPolicyLink) ImplementsZeroTrustAccessApplicationUpdateParamsBodyBrowserVNCApplicationPolicyUnion() {
}

type AccessApplicationUpdateParamsBodyBrowserVNCApplicationPoliciesObject struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodyBrowserVNCApplicationPoliciesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBrowserVNCApplicationPoliciesObject) ImplementsZeroTrustAccessApplicationUpdateParamsBodyBrowserVNCApplicationPolicyUnion() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthentication].
type AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateParamsBodyBrowserVNCApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsBodyAppLauncherApplication struct {
	// The application type.
	Type param.Field[ApplicationType] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdPsParam] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The policies that will apply to the application, in ascending order of
	// precedence. Items can reference existing policies or create new policies
	// exclusive to the application.
	Policies param.Field[[]AccessApplicationUpdateParamsBodyAppLauncherApplicationPolicyUnion] `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfig] `json:"scim_config"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodyAppLauncherApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyAppLauncherApplication) implementsZeroTrustAccessApplicationUpdateParamsBodyUnion() {
}

// A JSON that links a reusable policy to an application.
type AccessApplicationUpdateParamsBodyAppLauncherApplicationPolicy struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence     param.Field[int64]       `json:"precedence"`
	ApprovalGroups param.Field[interface{}] `json:"approval_groups,required"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude,required"`
	Include  param.Field[interface{}] `json:"include,required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require,required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodyAppLauncherApplicationPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyAppLauncherApplicationPolicy) ImplementsZeroTrustAccessApplicationUpdateParamsBodyAppLauncherApplicationPolicyUnion() {
}

// A JSON that links a reusable policy to an application.
//
// Satisfied by
// [zero_trust.AccessApplicationUpdateParamsBodyAppLauncherApplicationPoliciesAccessAppPolicyLink],
// [shared.UnionString],
// [zero_trust.AccessApplicationUpdateParamsBodyAppLauncherApplicationPoliciesObject],
// [AccessApplicationUpdateParamsBodyAppLauncherApplicationPolicy].
type AccessApplicationUpdateParamsBodyAppLauncherApplicationPolicyUnion interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsBodyAppLauncherApplicationPolicyUnion()
}

// A JSON that links a reusable policy to an application.
type AccessApplicationUpdateParamsBodyAppLauncherApplicationPoliciesAccessAppPolicyLink struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
}

func (r AccessApplicationUpdateParamsBodyAppLauncherApplicationPoliciesAccessAppPolicyLink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyAppLauncherApplicationPoliciesAccessAppPolicyLink) ImplementsZeroTrustAccessApplicationUpdateParamsBodyAppLauncherApplicationPolicyUnion() {
}

type AccessApplicationUpdateParamsBodyAppLauncherApplicationPoliciesObject struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodyAppLauncherApplicationPoliciesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyAppLauncherApplicationPoliciesObject) ImplementsZeroTrustAccessApplicationUpdateParamsBodyAppLauncherApplicationPolicyUnion() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthentication].
type AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateParamsBodyAppLauncherApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplication struct {
	// The application type.
	Type param.Field[ApplicationType] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdPsParam] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The policies that will apply to the application, in ascending order of
	// precedence. Items can reference existing policies or create new policies
	// exclusive to the application.
	Policies param.Field[[]AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion] `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfig] `json:"scim_config"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplication) implementsZeroTrustAccessApplicationUpdateParamsBodyUnion() {
}

// A JSON that links a reusable policy to an application.
type AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPolicy struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence     param.Field[int64]       `json:"precedence"`
	ApprovalGroups param.Field[interface{}] `json:"approval_groups,required"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude,required"`
	Include  param.Field[interface{}] `json:"include,required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require,required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPolicy) ImplementsZeroTrustAccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion() {
}

// A JSON that links a reusable policy to an application.
//
// Satisfied by
// [zero_trust.AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesAccessAppPolicyLink],
// [shared.UnionString],
// [zero_trust.AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesObject],
// [AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPolicy].
type AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion()
}

// A JSON that links a reusable policy to an application.
type AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesAccessAppPolicyLink struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
}

func (r AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesAccessAppPolicyLink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesAccessAppPolicyLink) ImplementsZeroTrustAccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion() {
}

type AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesObject struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPoliciesObject) ImplementsZeroTrustAccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthentication].
type AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplication struct {
	// The application type.
	Type param.Field[ApplicationType] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdPs param.Field[[]AllowedIdPsParam] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The policies that will apply to the application, in ascending order of
	// precedence. Items can reference existing policies or create new policies
	// exclusive to the application.
	Policies param.Field[[]AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion] `json:"policies"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfig] `json:"scim_config"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplication) implementsZeroTrustAccessApplicationUpdateParamsBodyUnion() {
}

// A JSON that links a reusable policy to an application.
type AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPolicy struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence     param.Field[int64]       `json:"precedence"`
	ApprovalGroups param.Field[interface{}] `json:"approval_groups,required"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude,required"`
	Include  param.Field[interface{}] `json:"include,required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require,required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPolicy) ImplementsZeroTrustAccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion() {
}

// A JSON that links a reusable policy to an application.
//
// Satisfied by
// [zero_trust.AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPoliciesAccessAppPolicyLink],
// [shared.UnionString],
// [zero_trust.AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPoliciesObject],
// [AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPolicy].
type AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion interface {
	ImplementsZeroTrustAccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion()
}

// A JSON that links a reusable policy to an application.
type AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPoliciesAccessAppPolicyLink struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
}

func (r AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPoliciesAccessAppPolicyLink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPoliciesAccessAppPolicyLink) ImplementsZeroTrustAccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion() {
}

type AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPoliciesObject struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPoliciesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPoliciesObject) ImplementsZeroTrustAccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthentication].
type AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationUpdateParamsBodyBookmarkApplication struct {
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// The URL or domain of the bookmark.
	Domain param.Field[string] `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name param.Field[string] `json:"name"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	SCIMConfig param.Field[AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfig] `json:"scim_config"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r AccessApplicationUpdateParamsBodyBookmarkApplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBookmarkApplication) implementsZeroTrustAccessApplicationUpdateParamsBodyUnion() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdPUID param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteURI param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]SCIMConfigMappingParam] `json:"mappings"`
}

func (r AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthentication struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string]      `json:"client_secret"`
	Scopes       param.Field[interface{}] `json:"scopes,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthentication) implementsZeroTrustAccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [zero_trust.SCIMConfigAuthenticationHTTPBasicParam],
// [zero_trust.SCIMConfigAuthenticationOAuthBearerTokenParam],
// [zero_trust.SCIMConfigAuthenticationOauth2Param],
// [AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthentication].
type AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion interface {
	implementsZeroTrustAccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion()
}

// The authentication scheme to use when making SCIM requests to this application.
type AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationScheme string

const (
	AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationSchemeHttpbasic        AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationScheme = "httpbasic"
	AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationSchemeOauthbearertoken AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationScheme = "oauthbearertoken"
	AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationSchemeOauth2           AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationScheme = "oauth2"
)

func (r AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationSchemeHttpbasic, AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationSchemeOauthbearertoken, AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationSchemeOauth2:
		return true
	}
	return false
}

type AccessApplicationUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessApplicationUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationUpdateResponse                `json:"result"`
	JSON    accessApplicationUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationUpdateResponseEnvelope]
type accessApplicationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessApplicationDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationDeleteResponse                `json:"result"`
	JSON    accessApplicationDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationDeleteResponseEnvelope]
type accessApplicationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	// Whether the API call was successful
	Success AccessApplicationGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationGetResponse                `json:"result"`
	JSON    accessApplicationGetResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationGetResponseEnvelope]
type accessApplicationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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

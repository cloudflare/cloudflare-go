// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package iam

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// OAuthClientService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOAuthClientService] method instead.
type OAuthClientService struct {
	Options []option.RequestOption
}

// NewOAuthClientService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOAuthClientService(opts ...option.RequestOption) (r *OAuthClientService) {
	r = &OAuthClientService{}
	r.Options = opts
	return
}

// Create a new OAuth client for an account.
func (r *OAuthClientService) New(ctx context.Context, params OAuthClientNewParams, opts ...option.RequestOption) (res *OAuthClientNewResponse, err error) {
	var env OAuthClientNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/oauth_clients", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update an existing OAuth client. Only include fields you want to update.
func (r *OAuthClientService) Update(ctx context.Context, oauthClientID string, params OAuthClientUpdateParams, opts ...option.RequestOption) (res *OAuthClientUpdateResponse, err error) {
	var env OAuthClientUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if oauthClientID == "" {
		err = errors.New("missing required oauth_client_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/oauth_clients/%s", params.AccountID, oauthClientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List all OAuth clients for an account.
func (r *OAuthClientService) List(ctx context.Context, query OAuthClientListParams, opts ...option.RequestOption) (res *pagination.SinglePage[OAuthClientListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/oauth_clients", query.AccountID)
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

// List all OAuth clients for an account.
func (r *OAuthClientService) ListAutoPaging(ctx context.Context, query OAuthClientListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[OAuthClientListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete an OAuth client.
func (r *OAuthClientService) Delete(ctx context.Context, oauthClientID string, body OAuthClientDeleteParams, opts ...option.RequestOption) (res *OAuthClientDeleteResponse, err error) {
	var env OAuthClientDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if oauthClientID == "" {
		err = errors.New("missing required oauth_client_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/oauth_clients/%s", body.AccountID, oauthClientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Removes the old client secret after a rotation, keeping only the new one. Use
// this after you have updated your client configuration to use the new secret. The
// `has_rotated_secret` field on the client indicates whether there is an old
// secret to delete.
func (r *OAuthClientService) DeleteRotatedSecret(ctx context.Context, oauthClientID string, body OAuthClientDeleteRotatedSecretParams, opts ...option.RequestOption) (res *OAuthClientDeleteRotatedSecretResponse, err error) {
	var env OAuthClientDeleteRotatedSecretResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if oauthClientID == "" {
		err = errors.New("missing required oauth_client_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/oauth_clients/%s/rotate_secret", body.AccountID, oauthClientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get details of a specific OAuth client.
func (r *OAuthClientService) Get(ctx context.Context, oauthClientID string, query OAuthClientGetParams, opts ...option.RequestOption) (res *OAuthClientGetResponse, err error) {
	var env OAuthClientGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if oauthClientID == "" {
		err = errors.New("missing required oauth_client_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/oauth_clients/%s", query.AccountID, oauthClientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Creates a second client secret so you can update your client configuration
// before deleting the old one. The `has_rotated_secret` field on the client will
// be set to `true`.
func (r *OAuthClientService) RotateSecret(ctx context.Context, oauthClientID string, body OAuthClientRotateSecretParams, opts ...option.RequestOption) (res *OAuthClientRotateSecretResponse, err error) {
	var env OAuthClientRotateSecretResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if oauthClientID == "" {
		err = errors.New("missing required oauth_client_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/oauth_clients/%s/rotate_secret", body.AccountID, oauthClientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fields shared by OAuth client responses and create/update requests.
type OAuthClientNewResponse struct {
	// The unique identifier for an OAuth client.
	ClientID string `json:"client_id" api:"required"`
	// Visibility of the OAuth client.
	Visibility OAuthClientNewResponseVisibility `json:"visibility" api:"required"`
	// Array of allowed CORS origins.
	AllowedCORSOrigins []string `json:"allowed_cors_origins"`
	// Human-readable name of the OAuth client.
	ClientName string `json:"client_name"`
	// The client secret. This is the only time the secret is returned in a response.
	ClientSecret string `json:"client_secret"`
	// URL of the home page of the client.
	ClientURI string `json:"client_uri"`
	// Client URI domain control verification state.
	ClientURIVerification OAuthClientNewResponseClientURIVerification `json:"client_uri_verification"`
	// Timestamp when the OAuth client was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Array of OAuth grant types the client is allowed to use. `authorization_code` is
	// required; `refresh_token` may be included optionally.
	GrantTypes []OAuthClientNewResponseGrantType `json:"grant_types"`
	// Indicates whether the client has a rotated secret that has not yet been deleted.
	HasRotatedSecret bool `json:"has_rotated_secret"`
	// URL of the client's logo.
	LogoURI string `json:"logo_uri"`
	// URL that points to a privacy policy document.
	PolicyURI string `json:"policy_uri"`
	// Array of allowed post-logout redirect URIs.
	PostLogoutRedirectURIs []string `json:"post_logout_redirect_uris"`
	// Timestamp when the OAuth client was promoted to public visibility.
	PromotedAt time.Time `json:"promoted_at" format:"date-time"`
	// Array of allowed redirect URIs for the client.
	RedirectURIs []string `json:"redirect_uris"`
	// Array of OAuth response types the client is allowed to use.
	ResponseTypes []OAuthClientNewResponseResponseType `json:"response_types"`
	// Array of OAuth scopes the client is allowed to request. Colon-delimited scopes
	// are not accepted. Dot-delimited scopes are validated against available OAuth API
	// scopes; simple identity scopes are allowed. Protocol scopes `offline_access` and
	// `openid` are added or removed automatically based on `grant_types` and
	// `response_types`.
	Scopes []string `json:"scopes"`
	// The authentication method the client uses at the token endpoint.
	TokenEndpointAuthMethod OAuthClientNewResponseTokenEndpointAuthMethod `json:"token_endpoint_auth_method"`
	// URL that points to a terms of service document.
	TosURI string `json:"tos_uri"`
	// Timestamp when the OAuth client was last updated.
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      oauthClientNewResponseJSON `json:"-"`
}

// oauthClientNewResponseJSON contains the JSON metadata for the struct
// [OAuthClientNewResponse]
type oauthClientNewResponseJSON struct {
	ClientID                apijson.Field
	Visibility              apijson.Field
	AllowedCORSOrigins      apijson.Field
	ClientName              apijson.Field
	ClientSecret            apijson.Field
	ClientURI               apijson.Field
	ClientURIVerification   apijson.Field
	CreatedAt               apijson.Field
	GrantTypes              apijson.Field
	HasRotatedSecret        apijson.Field
	LogoURI                 apijson.Field
	PolicyURI               apijson.Field
	PostLogoutRedirectURIs  apijson.Field
	PromotedAt              apijson.Field
	RedirectURIs            apijson.Field
	ResponseTypes           apijson.Field
	Scopes                  apijson.Field
	TokenEndpointAuthMethod apijson.Field
	TosURI                  apijson.Field
	UpdatedAt               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *OAuthClientNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientNewResponseJSON) RawJSON() string {
	return r.raw
}

// Visibility of the OAuth client.
type OAuthClientNewResponseVisibility string

const (
	OAuthClientNewResponseVisibilityPublic  OAuthClientNewResponseVisibility = "public"
	OAuthClientNewResponseVisibilityPrivate OAuthClientNewResponseVisibility = "private"
)

func (r OAuthClientNewResponseVisibility) IsKnown() bool {
	switch r {
	case OAuthClientNewResponseVisibilityPublic, OAuthClientNewResponseVisibilityPrivate:
		return true
	}
	return false
}

// Client URI domain control verification state.
type OAuthClientNewResponseClientURIVerification struct {
	// Current verification status for the client URI host.
	Status OAuthClientNewResponseClientURIVerificationStatus `json:"status"`
	// Exact TXT record value that must be added to DNS to prove ownership of the
	// client URI host.
	Text string                                          `json:"text"`
	JSON oauthClientNewResponseClientURIVerificationJSON `json:"-"`
}

// oauthClientNewResponseClientURIVerificationJSON contains the JSON metadata for
// the struct [OAuthClientNewResponseClientURIVerification]
type oauthClientNewResponseClientURIVerificationJSON struct {
	Status      apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientNewResponseClientURIVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientNewResponseClientURIVerificationJSON) RawJSON() string {
	return r.raw
}

// Current verification status for the client URI host.
type OAuthClientNewResponseClientURIVerificationStatus string

const (
	OAuthClientNewResponseClientURIVerificationStatusPending    OAuthClientNewResponseClientURIVerificationStatus = "pending"
	OAuthClientNewResponseClientURIVerificationStatusInProgress OAuthClientNewResponseClientURIVerificationStatus = "in_progress"
	OAuthClientNewResponseClientURIVerificationStatusVerified   OAuthClientNewResponseClientURIVerificationStatus = "verified"
	OAuthClientNewResponseClientURIVerificationStatusFailed     OAuthClientNewResponseClientURIVerificationStatus = "failed"
)

func (r OAuthClientNewResponseClientURIVerificationStatus) IsKnown() bool {
	switch r {
	case OAuthClientNewResponseClientURIVerificationStatusPending, OAuthClientNewResponseClientURIVerificationStatusInProgress, OAuthClientNewResponseClientURIVerificationStatusVerified, OAuthClientNewResponseClientURIVerificationStatusFailed:
		return true
	}
	return false
}

type OAuthClientNewResponseGrantType string

const (
	OAuthClientNewResponseGrantTypeAuthorizationCode OAuthClientNewResponseGrantType = "authorization_code"
	OAuthClientNewResponseGrantTypeRefreshToken      OAuthClientNewResponseGrantType = "refresh_token"
)

func (r OAuthClientNewResponseGrantType) IsKnown() bool {
	switch r {
	case OAuthClientNewResponseGrantTypeAuthorizationCode, OAuthClientNewResponseGrantTypeRefreshToken:
		return true
	}
	return false
}

type OAuthClientNewResponseResponseType string

const (
	OAuthClientNewResponseResponseTypeToken   OAuthClientNewResponseResponseType = "token"
	OAuthClientNewResponseResponseTypeIDToken OAuthClientNewResponseResponseType = "id_token"
	OAuthClientNewResponseResponseTypeCode    OAuthClientNewResponseResponseType = "code"
)

func (r OAuthClientNewResponseResponseType) IsKnown() bool {
	switch r {
	case OAuthClientNewResponseResponseTypeToken, OAuthClientNewResponseResponseTypeIDToken, OAuthClientNewResponseResponseTypeCode:
		return true
	}
	return false
}

// The authentication method the client uses at the token endpoint.
type OAuthClientNewResponseTokenEndpointAuthMethod string

const (
	OAuthClientNewResponseTokenEndpointAuthMethodNone              OAuthClientNewResponseTokenEndpointAuthMethod = "none"
	OAuthClientNewResponseTokenEndpointAuthMethodClientSecretBasic OAuthClientNewResponseTokenEndpointAuthMethod = "client_secret_basic"
	OAuthClientNewResponseTokenEndpointAuthMethodClientSecretPost  OAuthClientNewResponseTokenEndpointAuthMethod = "client_secret_post"
)

func (r OAuthClientNewResponseTokenEndpointAuthMethod) IsKnown() bool {
	switch r {
	case OAuthClientNewResponseTokenEndpointAuthMethodNone, OAuthClientNewResponseTokenEndpointAuthMethodClientSecretBasic, OAuthClientNewResponseTokenEndpointAuthMethodClientSecretPost:
		return true
	}
	return false
}

// Fields shared by OAuth client responses and create/update requests.
type OAuthClientUpdateResponse struct {
	// The unique identifier for an OAuth client.
	ClientID string `json:"client_id" api:"required"`
	// Visibility of the OAuth client.
	Visibility OAuthClientUpdateResponseVisibility `json:"visibility" api:"required"`
	// Array of allowed CORS origins.
	AllowedCORSOrigins []string `json:"allowed_cors_origins"`
	// Human-readable name of the OAuth client.
	ClientName string `json:"client_name"`
	// URL of the home page of the client.
	ClientURI string `json:"client_uri"`
	// Client URI domain control verification state.
	ClientURIVerification OAuthClientUpdateResponseClientURIVerification `json:"client_uri_verification"`
	// Timestamp when the OAuth client was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Array of OAuth grant types the client is allowed to use. `authorization_code` is
	// required; `refresh_token` may be included optionally.
	GrantTypes []OAuthClientUpdateResponseGrantType `json:"grant_types"`
	// Indicates whether the client has a rotated secret that has not yet been deleted.
	HasRotatedSecret bool `json:"has_rotated_secret"`
	// URL of the client's logo.
	LogoURI string `json:"logo_uri"`
	// URL that points to a privacy policy document.
	PolicyURI string `json:"policy_uri"`
	// Array of allowed post-logout redirect URIs.
	PostLogoutRedirectURIs []string `json:"post_logout_redirect_uris"`
	// Timestamp when the OAuth client was promoted to public visibility.
	PromotedAt time.Time `json:"promoted_at" format:"date-time"`
	// Array of allowed redirect URIs for the client.
	RedirectURIs []string `json:"redirect_uris"`
	// Array of OAuth response types the client is allowed to use.
	ResponseTypes []OAuthClientUpdateResponseResponseType `json:"response_types"`
	// Array of OAuth scopes the client is allowed to request. Colon-delimited scopes
	// are not accepted. Dot-delimited scopes are validated against available OAuth API
	// scopes; simple identity scopes are allowed. Protocol scopes `offline_access` and
	// `openid` are added or removed automatically based on `grant_types` and
	// `response_types`.
	Scopes []string `json:"scopes"`
	// The authentication method the client uses at the token endpoint.
	TokenEndpointAuthMethod OAuthClientUpdateResponseTokenEndpointAuthMethod `json:"token_endpoint_auth_method"`
	// URL that points to a terms of service document.
	TosURI string `json:"tos_uri"`
	// Timestamp when the OAuth client was last updated.
	UpdatedAt time.Time                     `json:"updated_at" format:"date-time"`
	JSON      oauthClientUpdateResponseJSON `json:"-"`
}

// oauthClientUpdateResponseJSON contains the JSON metadata for the struct
// [OAuthClientUpdateResponse]
type oauthClientUpdateResponseJSON struct {
	ClientID                apijson.Field
	Visibility              apijson.Field
	AllowedCORSOrigins      apijson.Field
	ClientName              apijson.Field
	ClientURI               apijson.Field
	ClientURIVerification   apijson.Field
	CreatedAt               apijson.Field
	GrantTypes              apijson.Field
	HasRotatedSecret        apijson.Field
	LogoURI                 apijson.Field
	PolicyURI               apijson.Field
	PostLogoutRedirectURIs  apijson.Field
	PromotedAt              apijson.Field
	RedirectURIs            apijson.Field
	ResponseTypes           apijson.Field
	Scopes                  apijson.Field
	TokenEndpointAuthMethod apijson.Field
	TosURI                  apijson.Field
	UpdatedAt               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *OAuthClientUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Visibility of the OAuth client.
type OAuthClientUpdateResponseVisibility string

const (
	OAuthClientUpdateResponseVisibilityPublic  OAuthClientUpdateResponseVisibility = "public"
	OAuthClientUpdateResponseVisibilityPrivate OAuthClientUpdateResponseVisibility = "private"
)

func (r OAuthClientUpdateResponseVisibility) IsKnown() bool {
	switch r {
	case OAuthClientUpdateResponseVisibilityPublic, OAuthClientUpdateResponseVisibilityPrivate:
		return true
	}
	return false
}

// Client URI domain control verification state.
type OAuthClientUpdateResponseClientURIVerification struct {
	// Current verification status for the client URI host.
	Status OAuthClientUpdateResponseClientURIVerificationStatus `json:"status"`
	// Exact TXT record value that must be added to DNS to prove ownership of the
	// client URI host.
	Text string                                             `json:"text"`
	JSON oauthClientUpdateResponseClientURIVerificationJSON `json:"-"`
}

// oauthClientUpdateResponseClientURIVerificationJSON contains the JSON metadata
// for the struct [OAuthClientUpdateResponseClientURIVerification]
type oauthClientUpdateResponseClientURIVerificationJSON struct {
	Status      apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientUpdateResponseClientURIVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientUpdateResponseClientURIVerificationJSON) RawJSON() string {
	return r.raw
}

// Current verification status for the client URI host.
type OAuthClientUpdateResponseClientURIVerificationStatus string

const (
	OAuthClientUpdateResponseClientURIVerificationStatusPending    OAuthClientUpdateResponseClientURIVerificationStatus = "pending"
	OAuthClientUpdateResponseClientURIVerificationStatusInProgress OAuthClientUpdateResponseClientURIVerificationStatus = "in_progress"
	OAuthClientUpdateResponseClientURIVerificationStatusVerified   OAuthClientUpdateResponseClientURIVerificationStatus = "verified"
	OAuthClientUpdateResponseClientURIVerificationStatusFailed     OAuthClientUpdateResponseClientURIVerificationStatus = "failed"
)

func (r OAuthClientUpdateResponseClientURIVerificationStatus) IsKnown() bool {
	switch r {
	case OAuthClientUpdateResponseClientURIVerificationStatusPending, OAuthClientUpdateResponseClientURIVerificationStatusInProgress, OAuthClientUpdateResponseClientURIVerificationStatusVerified, OAuthClientUpdateResponseClientURIVerificationStatusFailed:
		return true
	}
	return false
}

type OAuthClientUpdateResponseGrantType string

const (
	OAuthClientUpdateResponseGrantTypeAuthorizationCode OAuthClientUpdateResponseGrantType = "authorization_code"
	OAuthClientUpdateResponseGrantTypeRefreshToken      OAuthClientUpdateResponseGrantType = "refresh_token"
)

func (r OAuthClientUpdateResponseGrantType) IsKnown() bool {
	switch r {
	case OAuthClientUpdateResponseGrantTypeAuthorizationCode, OAuthClientUpdateResponseGrantTypeRefreshToken:
		return true
	}
	return false
}

type OAuthClientUpdateResponseResponseType string

const (
	OAuthClientUpdateResponseResponseTypeToken   OAuthClientUpdateResponseResponseType = "token"
	OAuthClientUpdateResponseResponseTypeIDToken OAuthClientUpdateResponseResponseType = "id_token"
	OAuthClientUpdateResponseResponseTypeCode    OAuthClientUpdateResponseResponseType = "code"
)

func (r OAuthClientUpdateResponseResponseType) IsKnown() bool {
	switch r {
	case OAuthClientUpdateResponseResponseTypeToken, OAuthClientUpdateResponseResponseTypeIDToken, OAuthClientUpdateResponseResponseTypeCode:
		return true
	}
	return false
}

// The authentication method the client uses at the token endpoint.
type OAuthClientUpdateResponseTokenEndpointAuthMethod string

const (
	OAuthClientUpdateResponseTokenEndpointAuthMethodNone              OAuthClientUpdateResponseTokenEndpointAuthMethod = "none"
	OAuthClientUpdateResponseTokenEndpointAuthMethodClientSecretBasic OAuthClientUpdateResponseTokenEndpointAuthMethod = "client_secret_basic"
	OAuthClientUpdateResponseTokenEndpointAuthMethodClientSecretPost  OAuthClientUpdateResponseTokenEndpointAuthMethod = "client_secret_post"
)

func (r OAuthClientUpdateResponseTokenEndpointAuthMethod) IsKnown() bool {
	switch r {
	case OAuthClientUpdateResponseTokenEndpointAuthMethodNone, OAuthClientUpdateResponseTokenEndpointAuthMethodClientSecretBasic, OAuthClientUpdateResponseTokenEndpointAuthMethodClientSecretPost:
		return true
	}
	return false
}

// Fields shared by OAuth client responses and create/update requests.
type OAuthClientListResponse struct {
	// The unique identifier for an OAuth client.
	ClientID string `json:"client_id" api:"required"`
	// Visibility of the OAuth client.
	Visibility OAuthClientListResponseVisibility `json:"visibility" api:"required"`
	// Array of allowed CORS origins.
	AllowedCORSOrigins []string `json:"allowed_cors_origins"`
	// Human-readable name of the OAuth client.
	ClientName string `json:"client_name"`
	// URL of the home page of the client.
	ClientURI string `json:"client_uri"`
	// Client URI domain control verification state.
	ClientURIVerification OAuthClientListResponseClientURIVerification `json:"client_uri_verification"`
	// Timestamp when the OAuth client was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Array of OAuth grant types the client is allowed to use. `authorization_code` is
	// required; `refresh_token` may be included optionally.
	GrantTypes []OAuthClientListResponseGrantType `json:"grant_types"`
	// Indicates whether the client has a rotated secret that has not yet been deleted.
	HasRotatedSecret bool `json:"has_rotated_secret"`
	// URL of the client's logo.
	LogoURI string `json:"logo_uri"`
	// URL that points to a privacy policy document.
	PolicyURI string `json:"policy_uri"`
	// Array of allowed post-logout redirect URIs.
	PostLogoutRedirectURIs []string `json:"post_logout_redirect_uris"`
	// Timestamp when the OAuth client was promoted to public visibility.
	PromotedAt time.Time `json:"promoted_at" format:"date-time"`
	// Array of allowed redirect URIs for the client.
	RedirectURIs []string `json:"redirect_uris"`
	// Array of OAuth response types the client is allowed to use.
	ResponseTypes []OAuthClientListResponseResponseType `json:"response_types"`
	// Array of OAuth scopes the client is allowed to request. Colon-delimited scopes
	// are not accepted. Dot-delimited scopes are validated against available OAuth API
	// scopes; simple identity scopes are allowed. Protocol scopes `offline_access` and
	// `openid` are added or removed automatically based on `grant_types` and
	// `response_types`.
	Scopes []string `json:"scopes"`
	// The authentication method the client uses at the token endpoint.
	TokenEndpointAuthMethod OAuthClientListResponseTokenEndpointAuthMethod `json:"token_endpoint_auth_method"`
	// URL that points to a terms of service document.
	TosURI string `json:"tos_uri"`
	// Timestamp when the OAuth client was last updated.
	UpdatedAt time.Time                   `json:"updated_at" format:"date-time"`
	JSON      oauthClientListResponseJSON `json:"-"`
}

// oauthClientListResponseJSON contains the JSON metadata for the struct
// [OAuthClientListResponse]
type oauthClientListResponseJSON struct {
	ClientID                apijson.Field
	Visibility              apijson.Field
	AllowedCORSOrigins      apijson.Field
	ClientName              apijson.Field
	ClientURI               apijson.Field
	ClientURIVerification   apijson.Field
	CreatedAt               apijson.Field
	GrantTypes              apijson.Field
	HasRotatedSecret        apijson.Field
	LogoURI                 apijson.Field
	PolicyURI               apijson.Field
	PostLogoutRedirectURIs  apijson.Field
	PromotedAt              apijson.Field
	RedirectURIs            apijson.Field
	ResponseTypes           apijson.Field
	Scopes                  apijson.Field
	TokenEndpointAuthMethod apijson.Field
	TosURI                  apijson.Field
	UpdatedAt               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *OAuthClientListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientListResponseJSON) RawJSON() string {
	return r.raw
}

// Visibility of the OAuth client.
type OAuthClientListResponseVisibility string

const (
	OAuthClientListResponseVisibilityPublic  OAuthClientListResponseVisibility = "public"
	OAuthClientListResponseVisibilityPrivate OAuthClientListResponseVisibility = "private"
)

func (r OAuthClientListResponseVisibility) IsKnown() bool {
	switch r {
	case OAuthClientListResponseVisibilityPublic, OAuthClientListResponseVisibilityPrivate:
		return true
	}
	return false
}

// Client URI domain control verification state.
type OAuthClientListResponseClientURIVerification struct {
	// Current verification status for the client URI host.
	Status OAuthClientListResponseClientURIVerificationStatus `json:"status"`
	// Exact TXT record value that must be added to DNS to prove ownership of the
	// client URI host.
	Text string                                           `json:"text"`
	JSON oauthClientListResponseClientURIVerificationJSON `json:"-"`
}

// oauthClientListResponseClientURIVerificationJSON contains the JSON metadata for
// the struct [OAuthClientListResponseClientURIVerification]
type oauthClientListResponseClientURIVerificationJSON struct {
	Status      apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientListResponseClientURIVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientListResponseClientURIVerificationJSON) RawJSON() string {
	return r.raw
}

// Current verification status for the client URI host.
type OAuthClientListResponseClientURIVerificationStatus string

const (
	OAuthClientListResponseClientURIVerificationStatusPending    OAuthClientListResponseClientURIVerificationStatus = "pending"
	OAuthClientListResponseClientURIVerificationStatusInProgress OAuthClientListResponseClientURIVerificationStatus = "in_progress"
	OAuthClientListResponseClientURIVerificationStatusVerified   OAuthClientListResponseClientURIVerificationStatus = "verified"
	OAuthClientListResponseClientURIVerificationStatusFailed     OAuthClientListResponseClientURIVerificationStatus = "failed"
)

func (r OAuthClientListResponseClientURIVerificationStatus) IsKnown() bool {
	switch r {
	case OAuthClientListResponseClientURIVerificationStatusPending, OAuthClientListResponseClientURIVerificationStatusInProgress, OAuthClientListResponseClientURIVerificationStatusVerified, OAuthClientListResponseClientURIVerificationStatusFailed:
		return true
	}
	return false
}

type OAuthClientListResponseGrantType string

const (
	OAuthClientListResponseGrantTypeAuthorizationCode OAuthClientListResponseGrantType = "authorization_code"
	OAuthClientListResponseGrantTypeRefreshToken      OAuthClientListResponseGrantType = "refresh_token"
)

func (r OAuthClientListResponseGrantType) IsKnown() bool {
	switch r {
	case OAuthClientListResponseGrantTypeAuthorizationCode, OAuthClientListResponseGrantTypeRefreshToken:
		return true
	}
	return false
}

type OAuthClientListResponseResponseType string

const (
	OAuthClientListResponseResponseTypeToken   OAuthClientListResponseResponseType = "token"
	OAuthClientListResponseResponseTypeIDToken OAuthClientListResponseResponseType = "id_token"
	OAuthClientListResponseResponseTypeCode    OAuthClientListResponseResponseType = "code"
)

func (r OAuthClientListResponseResponseType) IsKnown() bool {
	switch r {
	case OAuthClientListResponseResponseTypeToken, OAuthClientListResponseResponseTypeIDToken, OAuthClientListResponseResponseTypeCode:
		return true
	}
	return false
}

// The authentication method the client uses at the token endpoint.
type OAuthClientListResponseTokenEndpointAuthMethod string

const (
	OAuthClientListResponseTokenEndpointAuthMethodNone              OAuthClientListResponseTokenEndpointAuthMethod = "none"
	OAuthClientListResponseTokenEndpointAuthMethodClientSecretBasic OAuthClientListResponseTokenEndpointAuthMethod = "client_secret_basic"
	OAuthClientListResponseTokenEndpointAuthMethodClientSecretPost  OAuthClientListResponseTokenEndpointAuthMethod = "client_secret_post"
)

func (r OAuthClientListResponseTokenEndpointAuthMethod) IsKnown() bool {
	switch r {
	case OAuthClientListResponseTokenEndpointAuthMethodNone, OAuthClientListResponseTokenEndpointAuthMethodClientSecretBasic, OAuthClientListResponseTokenEndpointAuthMethodClientSecretPost:
		return true
	}
	return false
}

type OAuthClientDeleteResponse struct {
	// Identifier
	ID   string                        `json:"id" api:"required"`
	JSON oauthClientDeleteResponseJSON `json:"-"`
}

// oauthClientDeleteResponseJSON contains the JSON metadata for the struct
// [OAuthClientDeleteResponse]
type oauthClientDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OAuthClientDeleteRotatedSecretResponse struct {
	// Identifier
	ID   string                                     `json:"id" api:"required"`
	JSON oauthClientDeleteRotatedSecretResponseJSON `json:"-"`
}

// oauthClientDeleteRotatedSecretResponseJSON contains the JSON metadata for the
// struct [OAuthClientDeleteRotatedSecretResponse]
type oauthClientDeleteRotatedSecretResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientDeleteRotatedSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientDeleteRotatedSecretResponseJSON) RawJSON() string {
	return r.raw
}

// Fields shared by OAuth client responses and create/update requests.
type OAuthClientGetResponse struct {
	// The unique identifier for an OAuth client.
	ClientID string `json:"client_id" api:"required"`
	// Visibility of the OAuth client.
	Visibility OAuthClientGetResponseVisibility `json:"visibility" api:"required"`
	// Array of allowed CORS origins.
	AllowedCORSOrigins []string `json:"allowed_cors_origins"`
	// Human-readable name of the OAuth client.
	ClientName string `json:"client_name"`
	// URL of the home page of the client.
	ClientURI string `json:"client_uri"`
	// Client URI domain control verification state.
	ClientURIVerification OAuthClientGetResponseClientURIVerification `json:"client_uri_verification"`
	// Timestamp when the OAuth client was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Array of OAuth grant types the client is allowed to use. `authorization_code` is
	// required; `refresh_token` may be included optionally.
	GrantTypes []OAuthClientGetResponseGrantType `json:"grant_types"`
	// Indicates whether the client has a rotated secret that has not yet been deleted.
	HasRotatedSecret bool `json:"has_rotated_secret"`
	// URL of the client's logo.
	LogoURI string `json:"logo_uri"`
	// URL that points to a privacy policy document.
	PolicyURI string `json:"policy_uri"`
	// Array of allowed post-logout redirect URIs.
	PostLogoutRedirectURIs []string `json:"post_logout_redirect_uris"`
	// Timestamp when the OAuth client was promoted to public visibility.
	PromotedAt time.Time `json:"promoted_at" format:"date-time"`
	// Array of allowed redirect URIs for the client.
	RedirectURIs []string `json:"redirect_uris"`
	// Array of OAuth response types the client is allowed to use.
	ResponseTypes []OAuthClientGetResponseResponseType `json:"response_types"`
	// Array of OAuth scopes the client is allowed to request. Colon-delimited scopes
	// are not accepted. Dot-delimited scopes are validated against available OAuth API
	// scopes; simple identity scopes are allowed. Protocol scopes `offline_access` and
	// `openid` are added or removed automatically based on `grant_types` and
	// `response_types`.
	Scopes []string `json:"scopes"`
	// The authentication method the client uses at the token endpoint.
	TokenEndpointAuthMethod OAuthClientGetResponseTokenEndpointAuthMethod `json:"token_endpoint_auth_method"`
	// URL that points to a terms of service document.
	TosURI string `json:"tos_uri"`
	// Timestamp when the OAuth client was last updated.
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      oauthClientGetResponseJSON `json:"-"`
}

// oauthClientGetResponseJSON contains the JSON metadata for the struct
// [OAuthClientGetResponse]
type oauthClientGetResponseJSON struct {
	ClientID                apijson.Field
	Visibility              apijson.Field
	AllowedCORSOrigins      apijson.Field
	ClientName              apijson.Field
	ClientURI               apijson.Field
	ClientURIVerification   apijson.Field
	CreatedAt               apijson.Field
	GrantTypes              apijson.Field
	HasRotatedSecret        apijson.Field
	LogoURI                 apijson.Field
	PolicyURI               apijson.Field
	PostLogoutRedirectURIs  apijson.Field
	PromotedAt              apijson.Field
	RedirectURIs            apijson.Field
	ResponseTypes           apijson.Field
	Scopes                  apijson.Field
	TokenEndpointAuthMethod apijson.Field
	TosURI                  apijson.Field
	UpdatedAt               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *OAuthClientGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientGetResponseJSON) RawJSON() string {
	return r.raw
}

// Visibility of the OAuth client.
type OAuthClientGetResponseVisibility string

const (
	OAuthClientGetResponseVisibilityPublic  OAuthClientGetResponseVisibility = "public"
	OAuthClientGetResponseVisibilityPrivate OAuthClientGetResponseVisibility = "private"
)

func (r OAuthClientGetResponseVisibility) IsKnown() bool {
	switch r {
	case OAuthClientGetResponseVisibilityPublic, OAuthClientGetResponseVisibilityPrivate:
		return true
	}
	return false
}

// Client URI domain control verification state.
type OAuthClientGetResponseClientURIVerification struct {
	// Current verification status for the client URI host.
	Status OAuthClientGetResponseClientURIVerificationStatus `json:"status"`
	// Exact TXT record value that must be added to DNS to prove ownership of the
	// client URI host.
	Text string                                          `json:"text"`
	JSON oauthClientGetResponseClientURIVerificationJSON `json:"-"`
}

// oauthClientGetResponseClientURIVerificationJSON contains the JSON metadata for
// the struct [OAuthClientGetResponseClientURIVerification]
type oauthClientGetResponseClientURIVerificationJSON struct {
	Status      apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientGetResponseClientURIVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientGetResponseClientURIVerificationJSON) RawJSON() string {
	return r.raw
}

// Current verification status for the client URI host.
type OAuthClientGetResponseClientURIVerificationStatus string

const (
	OAuthClientGetResponseClientURIVerificationStatusPending    OAuthClientGetResponseClientURIVerificationStatus = "pending"
	OAuthClientGetResponseClientURIVerificationStatusInProgress OAuthClientGetResponseClientURIVerificationStatus = "in_progress"
	OAuthClientGetResponseClientURIVerificationStatusVerified   OAuthClientGetResponseClientURIVerificationStatus = "verified"
	OAuthClientGetResponseClientURIVerificationStatusFailed     OAuthClientGetResponseClientURIVerificationStatus = "failed"
)

func (r OAuthClientGetResponseClientURIVerificationStatus) IsKnown() bool {
	switch r {
	case OAuthClientGetResponseClientURIVerificationStatusPending, OAuthClientGetResponseClientURIVerificationStatusInProgress, OAuthClientGetResponseClientURIVerificationStatusVerified, OAuthClientGetResponseClientURIVerificationStatusFailed:
		return true
	}
	return false
}

type OAuthClientGetResponseGrantType string

const (
	OAuthClientGetResponseGrantTypeAuthorizationCode OAuthClientGetResponseGrantType = "authorization_code"
	OAuthClientGetResponseGrantTypeRefreshToken      OAuthClientGetResponseGrantType = "refresh_token"
)

func (r OAuthClientGetResponseGrantType) IsKnown() bool {
	switch r {
	case OAuthClientGetResponseGrantTypeAuthorizationCode, OAuthClientGetResponseGrantTypeRefreshToken:
		return true
	}
	return false
}

type OAuthClientGetResponseResponseType string

const (
	OAuthClientGetResponseResponseTypeToken   OAuthClientGetResponseResponseType = "token"
	OAuthClientGetResponseResponseTypeIDToken OAuthClientGetResponseResponseType = "id_token"
	OAuthClientGetResponseResponseTypeCode    OAuthClientGetResponseResponseType = "code"
)

func (r OAuthClientGetResponseResponseType) IsKnown() bool {
	switch r {
	case OAuthClientGetResponseResponseTypeToken, OAuthClientGetResponseResponseTypeIDToken, OAuthClientGetResponseResponseTypeCode:
		return true
	}
	return false
}

// The authentication method the client uses at the token endpoint.
type OAuthClientGetResponseTokenEndpointAuthMethod string

const (
	OAuthClientGetResponseTokenEndpointAuthMethodNone              OAuthClientGetResponseTokenEndpointAuthMethod = "none"
	OAuthClientGetResponseTokenEndpointAuthMethodClientSecretBasic OAuthClientGetResponseTokenEndpointAuthMethod = "client_secret_basic"
	OAuthClientGetResponseTokenEndpointAuthMethodClientSecretPost  OAuthClientGetResponseTokenEndpointAuthMethod = "client_secret_post"
)

func (r OAuthClientGetResponseTokenEndpointAuthMethod) IsKnown() bool {
	switch r {
	case OAuthClientGetResponseTokenEndpointAuthMethodNone, OAuthClientGetResponseTokenEndpointAuthMethodClientSecretBasic, OAuthClientGetResponseTokenEndpointAuthMethodClientSecretPost:
		return true
	}
	return false
}

type OAuthClientRotateSecretResponse struct {
	// The new client secret.
	ClientSecret string                              `json:"client_secret"`
	JSON         oauthClientRotateSecretResponseJSON `json:"-"`
}

// oauthClientRotateSecretResponseJSON contains the JSON metadata for the struct
// [OAuthClientRotateSecretResponse]
type oauthClientRotateSecretResponseJSON struct {
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OAuthClientRotateSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientRotateSecretResponseJSON) RawJSON() string {
	return r.raw
}

type OAuthClientNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Human-readable name of the OAuth client.
	ClientName param.Field[string] `json:"client_name" api:"required"`
	// Array of OAuth grant types the client is allowed to use. `authorization_code` is
	// required; `refresh_token` may be included optionally.
	GrantTypes param.Field[[]OAuthClientNewParamsGrantType] `json:"grant_types" api:"required"`
	// Array of allowed redirect URIs for the client.
	RedirectURIs param.Field[[]string] `json:"redirect_uris" api:"required"`
	// Array of OAuth response types the client is allowed to use.
	ResponseTypes param.Field[[]OAuthClientNewParamsResponseType] `json:"response_types" api:"required"`
	// Array of OAuth scopes the client is allowed to request. Colon-delimited scopes
	// are not accepted. Dot-delimited scopes are validated against available OAuth API
	// scopes; simple identity scopes are allowed. Protocol scopes `offline_access` and
	// `openid` are added or removed automatically based on `grant_types` and
	// `response_types`.
	Scopes param.Field[[]string] `json:"scopes" api:"required"`
	// The authentication method the client uses at the token endpoint.
	TokenEndpointAuthMethod param.Field[OAuthClientNewParamsTokenEndpointAuthMethod] `json:"token_endpoint_auth_method" api:"required"`
	// Array of allowed CORS origins.
	AllowedCORSOrigins param.Field[[]string] `json:"allowed_cors_origins"`
	// URL of the home page of the client.
	ClientURI param.Field[string] `json:"client_uri"`
	// URL of the client's logo.
	LogoURI param.Field[string] `json:"logo_uri"`
	// URL that points to a privacy policy document.
	PolicyURI param.Field[string] `json:"policy_uri"`
	// Array of allowed post-logout redirect URIs.
	PostLogoutRedirectURIs param.Field[[]string] `json:"post_logout_redirect_uris"`
	// URL that points to a terms of service document.
	TosURI param.Field[string] `json:"tos_uri"`
}

func (r OAuthClientNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OAuthClientNewParamsGrantType string

const (
	OAuthClientNewParamsGrantTypeAuthorizationCode OAuthClientNewParamsGrantType = "authorization_code"
	OAuthClientNewParamsGrantTypeRefreshToken      OAuthClientNewParamsGrantType = "refresh_token"
)

func (r OAuthClientNewParamsGrantType) IsKnown() bool {
	switch r {
	case OAuthClientNewParamsGrantTypeAuthorizationCode, OAuthClientNewParamsGrantTypeRefreshToken:
		return true
	}
	return false
}

type OAuthClientNewParamsResponseType string

const (
	OAuthClientNewParamsResponseTypeToken   OAuthClientNewParamsResponseType = "token"
	OAuthClientNewParamsResponseTypeIDToken OAuthClientNewParamsResponseType = "id_token"
	OAuthClientNewParamsResponseTypeCode    OAuthClientNewParamsResponseType = "code"
)

func (r OAuthClientNewParamsResponseType) IsKnown() bool {
	switch r {
	case OAuthClientNewParamsResponseTypeToken, OAuthClientNewParamsResponseTypeIDToken, OAuthClientNewParamsResponseTypeCode:
		return true
	}
	return false
}

// The authentication method the client uses at the token endpoint.
type OAuthClientNewParamsTokenEndpointAuthMethod string

const (
	OAuthClientNewParamsTokenEndpointAuthMethodNone              OAuthClientNewParamsTokenEndpointAuthMethod = "none"
	OAuthClientNewParamsTokenEndpointAuthMethodClientSecretBasic OAuthClientNewParamsTokenEndpointAuthMethod = "client_secret_basic"
	OAuthClientNewParamsTokenEndpointAuthMethodClientSecretPost  OAuthClientNewParamsTokenEndpointAuthMethod = "client_secret_post"
)

func (r OAuthClientNewParamsTokenEndpointAuthMethod) IsKnown() bool {
	switch r {
	case OAuthClientNewParamsTokenEndpointAuthMethodNone, OAuthClientNewParamsTokenEndpointAuthMethodClientSecretBasic, OAuthClientNewParamsTokenEndpointAuthMethodClientSecretPost:
		return true
	}
	return false
}

type OAuthClientNewResponseEnvelope struct {
	Errors   []OAuthClientNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []OAuthClientNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OAuthClientNewResponseEnvelopeSuccess `json:"success" api:"required"`
	// Fields shared by OAuth client responses and create/update requests.
	Result OAuthClientNewResponse             `json:"result"`
	JSON   oauthClientNewResponseEnvelopeJSON `json:"-"`
}

// oauthClientNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [OAuthClientNewResponseEnvelope]
type oauthClientNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OAuthClientNewResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           OAuthClientNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             oauthClientNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// oauthClientNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OAuthClientNewResponseEnvelopeErrors]
type oauthClientNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OAuthClientNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OAuthClientNewResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    oauthClientNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// oauthClientNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [OAuthClientNewResponseEnvelopeErrorsSource]
type oauthClientNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OAuthClientNewResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           OAuthClientNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             oauthClientNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// oauthClientNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [OAuthClientNewResponseEnvelopeMessages]
type oauthClientNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OAuthClientNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OAuthClientNewResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    oauthClientNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// oauthClientNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [OAuthClientNewResponseEnvelopeMessagesSource]
type oauthClientNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OAuthClientNewResponseEnvelopeSuccess bool

const (
	OAuthClientNewResponseEnvelopeSuccessTrue OAuthClientNewResponseEnvelopeSuccess = true
)

func (r OAuthClientNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OAuthClientNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OAuthClientUpdateParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Array of allowed CORS origins.
	AllowedCORSOrigins param.Field[[]string] `json:"allowed_cors_origins"`
	// Human-readable name of the OAuth client.
	ClientName param.Field[string] `json:"client_name"`
	// URL of the home page of the client.
	ClientURI param.Field[string] `json:"client_uri"`
	// Array of OAuth grant types the client is allowed to use. `authorization_code` is
	// required; `refresh_token` may be included optionally.
	GrantTypes param.Field[[]OAuthClientUpdateParamsGrantType] `json:"grant_types"`
	// URL of the client's logo.
	LogoURI param.Field[string] `json:"logo_uri"`
	// URL that points to a privacy policy document.
	PolicyURI param.Field[string] `json:"policy_uri"`
	// Array of allowed post-logout redirect URIs.
	PostLogoutRedirectURIs param.Field[[]string] `json:"post_logout_redirect_uris"`
	// Array of allowed redirect URIs for the client.
	RedirectURIs param.Field[[]string] `json:"redirect_uris"`
	// Array of OAuth response types the client is allowed to use.
	ResponseTypes param.Field[[]OAuthClientUpdateParamsResponseType] `json:"response_types"`
	// Array of OAuth scopes the client is allowed to request. Colon-delimited scopes
	// are not accepted. Dot-delimited scopes are validated against available OAuth API
	// scopes; simple identity scopes are allowed. Protocol scopes `offline_access` and
	// `openid` are added or removed automatically based on `grant_types` and
	// `response_types`.
	Scopes param.Field[[]string] `json:"scopes"`
	// The authentication method the client uses at the token endpoint.
	TokenEndpointAuthMethod param.Field[OAuthClientUpdateParamsTokenEndpointAuthMethod] `json:"token_endpoint_auth_method"`
	// URL that points to a terms of service document.
	TosURI param.Field[string] `json:"tos_uri"`
	// Promote the OAuth client from private to public visibility. Only `public` is
	// accepted; demotion to `private` is not supported. Promotion requires a non-empty
	// client name, logo URI, verified client URI host, and at least one non-identity
	// scope.
	Visibility param.Field[OAuthClientUpdateParamsVisibility] `json:"visibility"`
}

func (r OAuthClientUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OAuthClientUpdateParamsGrantType string

const (
	OAuthClientUpdateParamsGrantTypeAuthorizationCode OAuthClientUpdateParamsGrantType = "authorization_code"
	OAuthClientUpdateParamsGrantTypeRefreshToken      OAuthClientUpdateParamsGrantType = "refresh_token"
)

func (r OAuthClientUpdateParamsGrantType) IsKnown() bool {
	switch r {
	case OAuthClientUpdateParamsGrantTypeAuthorizationCode, OAuthClientUpdateParamsGrantTypeRefreshToken:
		return true
	}
	return false
}

type OAuthClientUpdateParamsResponseType string

const (
	OAuthClientUpdateParamsResponseTypeToken   OAuthClientUpdateParamsResponseType = "token"
	OAuthClientUpdateParamsResponseTypeIDToken OAuthClientUpdateParamsResponseType = "id_token"
	OAuthClientUpdateParamsResponseTypeCode    OAuthClientUpdateParamsResponseType = "code"
)

func (r OAuthClientUpdateParamsResponseType) IsKnown() bool {
	switch r {
	case OAuthClientUpdateParamsResponseTypeToken, OAuthClientUpdateParamsResponseTypeIDToken, OAuthClientUpdateParamsResponseTypeCode:
		return true
	}
	return false
}

// The authentication method the client uses at the token endpoint.
type OAuthClientUpdateParamsTokenEndpointAuthMethod string

const (
	OAuthClientUpdateParamsTokenEndpointAuthMethodNone              OAuthClientUpdateParamsTokenEndpointAuthMethod = "none"
	OAuthClientUpdateParamsTokenEndpointAuthMethodClientSecretBasic OAuthClientUpdateParamsTokenEndpointAuthMethod = "client_secret_basic"
	OAuthClientUpdateParamsTokenEndpointAuthMethodClientSecretPost  OAuthClientUpdateParamsTokenEndpointAuthMethod = "client_secret_post"
)

func (r OAuthClientUpdateParamsTokenEndpointAuthMethod) IsKnown() bool {
	switch r {
	case OAuthClientUpdateParamsTokenEndpointAuthMethodNone, OAuthClientUpdateParamsTokenEndpointAuthMethodClientSecretBasic, OAuthClientUpdateParamsTokenEndpointAuthMethodClientSecretPost:
		return true
	}
	return false
}

// Promote the OAuth client from private to public visibility. Only `public` is
// accepted; demotion to `private` is not supported. Promotion requires a non-empty
// client name, logo URI, verified client URI host, and at least one non-identity
// scope.
type OAuthClientUpdateParamsVisibility string

const (
	OAuthClientUpdateParamsVisibilityPublic OAuthClientUpdateParamsVisibility = "public"
)

func (r OAuthClientUpdateParamsVisibility) IsKnown() bool {
	switch r {
	case OAuthClientUpdateParamsVisibilityPublic:
		return true
	}
	return false
}

type OAuthClientUpdateResponseEnvelope struct {
	Errors   []OAuthClientUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []OAuthClientUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OAuthClientUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	// Fields shared by OAuth client responses and create/update requests.
	Result OAuthClientUpdateResponse             `json:"result"`
	JSON   oauthClientUpdateResponseEnvelopeJSON `json:"-"`
}

// oauthClientUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [OAuthClientUpdateResponseEnvelope]
type oauthClientUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OAuthClientUpdateResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           OAuthClientUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             oauthClientUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// oauthClientUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OAuthClientUpdateResponseEnvelopeErrors]
type oauthClientUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OAuthClientUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OAuthClientUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    oauthClientUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// oauthClientUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [OAuthClientUpdateResponseEnvelopeErrorsSource]
type oauthClientUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OAuthClientUpdateResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           OAuthClientUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             oauthClientUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// oauthClientUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [OAuthClientUpdateResponseEnvelopeMessages]
type oauthClientUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OAuthClientUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OAuthClientUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    oauthClientUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// oauthClientUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [OAuthClientUpdateResponseEnvelopeMessagesSource]
type oauthClientUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OAuthClientUpdateResponseEnvelopeSuccess bool

const (
	OAuthClientUpdateResponseEnvelopeSuccessTrue OAuthClientUpdateResponseEnvelopeSuccess = true
)

func (r OAuthClientUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OAuthClientUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OAuthClientListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type OAuthClientDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type OAuthClientDeleteResponseEnvelope struct {
	Errors   []OAuthClientDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []OAuthClientDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OAuthClientDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  OAuthClientDeleteResponse                `json:"result" api:"nullable"`
	JSON    oauthClientDeleteResponseEnvelopeJSON    `json:"-"`
}

// oauthClientDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [OAuthClientDeleteResponseEnvelope]
type oauthClientDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OAuthClientDeleteResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           OAuthClientDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             oauthClientDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// oauthClientDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OAuthClientDeleteResponseEnvelopeErrors]
type oauthClientDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OAuthClientDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OAuthClientDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    oauthClientDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// oauthClientDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [OAuthClientDeleteResponseEnvelopeErrorsSource]
type oauthClientDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OAuthClientDeleteResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           OAuthClientDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             oauthClientDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// oauthClientDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [OAuthClientDeleteResponseEnvelopeMessages]
type oauthClientDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OAuthClientDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OAuthClientDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    oauthClientDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// oauthClientDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [OAuthClientDeleteResponseEnvelopeMessagesSource]
type oauthClientDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OAuthClientDeleteResponseEnvelopeSuccess bool

const (
	OAuthClientDeleteResponseEnvelopeSuccessTrue OAuthClientDeleteResponseEnvelopeSuccess = true
)

func (r OAuthClientDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OAuthClientDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OAuthClientDeleteRotatedSecretParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type OAuthClientDeleteRotatedSecretResponseEnvelope struct {
	Errors   []OAuthClientDeleteRotatedSecretResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []OAuthClientDeleteRotatedSecretResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OAuthClientDeleteRotatedSecretResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  OAuthClientDeleteRotatedSecretResponse                `json:"result" api:"nullable"`
	JSON    oauthClientDeleteRotatedSecretResponseEnvelopeJSON    `json:"-"`
}

// oauthClientDeleteRotatedSecretResponseEnvelopeJSON contains the JSON metadata
// for the struct [OAuthClientDeleteRotatedSecretResponseEnvelope]
type oauthClientDeleteRotatedSecretResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientDeleteRotatedSecretResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientDeleteRotatedSecretResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OAuthClientDeleteRotatedSecretResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           OAuthClientDeleteRotatedSecretResponseEnvelopeErrorsSource `json:"source"`
	JSON             oauthClientDeleteRotatedSecretResponseEnvelopeErrorsJSON   `json:"-"`
}

// oauthClientDeleteRotatedSecretResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [OAuthClientDeleteRotatedSecretResponseEnvelopeErrors]
type oauthClientDeleteRotatedSecretResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OAuthClientDeleteRotatedSecretResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientDeleteRotatedSecretResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OAuthClientDeleteRotatedSecretResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    oauthClientDeleteRotatedSecretResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// oauthClientDeleteRotatedSecretResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [OAuthClientDeleteRotatedSecretResponseEnvelopeErrorsSource]
type oauthClientDeleteRotatedSecretResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientDeleteRotatedSecretResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientDeleteRotatedSecretResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OAuthClientDeleteRotatedSecretResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           OAuthClientDeleteRotatedSecretResponseEnvelopeMessagesSource `json:"source"`
	JSON             oauthClientDeleteRotatedSecretResponseEnvelopeMessagesJSON   `json:"-"`
}

// oauthClientDeleteRotatedSecretResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [OAuthClientDeleteRotatedSecretResponseEnvelopeMessages]
type oauthClientDeleteRotatedSecretResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OAuthClientDeleteRotatedSecretResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientDeleteRotatedSecretResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OAuthClientDeleteRotatedSecretResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    oauthClientDeleteRotatedSecretResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// oauthClientDeleteRotatedSecretResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [OAuthClientDeleteRotatedSecretResponseEnvelopeMessagesSource]
type oauthClientDeleteRotatedSecretResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientDeleteRotatedSecretResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientDeleteRotatedSecretResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OAuthClientDeleteRotatedSecretResponseEnvelopeSuccess bool

const (
	OAuthClientDeleteRotatedSecretResponseEnvelopeSuccessTrue OAuthClientDeleteRotatedSecretResponseEnvelopeSuccess = true
)

func (r OAuthClientDeleteRotatedSecretResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OAuthClientDeleteRotatedSecretResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OAuthClientGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type OAuthClientGetResponseEnvelope struct {
	Errors   []OAuthClientGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []OAuthClientGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OAuthClientGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// Fields shared by OAuth client responses and create/update requests.
	Result OAuthClientGetResponse             `json:"result"`
	JSON   oauthClientGetResponseEnvelopeJSON `json:"-"`
}

// oauthClientGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [OAuthClientGetResponseEnvelope]
type oauthClientGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OAuthClientGetResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           OAuthClientGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             oauthClientGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// oauthClientGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OAuthClientGetResponseEnvelopeErrors]
type oauthClientGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OAuthClientGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OAuthClientGetResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    oauthClientGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// oauthClientGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [OAuthClientGetResponseEnvelopeErrorsSource]
type oauthClientGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OAuthClientGetResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           OAuthClientGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             oauthClientGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// oauthClientGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [OAuthClientGetResponseEnvelopeMessages]
type oauthClientGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OAuthClientGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OAuthClientGetResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    oauthClientGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// oauthClientGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [OAuthClientGetResponseEnvelopeMessagesSource]
type oauthClientGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OAuthClientGetResponseEnvelopeSuccess bool

const (
	OAuthClientGetResponseEnvelopeSuccessTrue OAuthClientGetResponseEnvelopeSuccess = true
)

func (r OAuthClientGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OAuthClientGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OAuthClientRotateSecretParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type OAuthClientRotateSecretResponseEnvelope struct {
	Errors   []OAuthClientRotateSecretResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []OAuthClientRotateSecretResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OAuthClientRotateSecretResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  OAuthClientRotateSecretResponse                `json:"result"`
	JSON    oauthClientRotateSecretResponseEnvelopeJSON    `json:"-"`
}

// oauthClientRotateSecretResponseEnvelopeJSON contains the JSON metadata for the
// struct [OAuthClientRotateSecretResponseEnvelope]
type oauthClientRotateSecretResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientRotateSecretResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientRotateSecretResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OAuthClientRotateSecretResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           OAuthClientRotateSecretResponseEnvelopeErrorsSource `json:"source"`
	JSON             oauthClientRotateSecretResponseEnvelopeErrorsJSON   `json:"-"`
}

// oauthClientRotateSecretResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [OAuthClientRotateSecretResponseEnvelopeErrors]
type oauthClientRotateSecretResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OAuthClientRotateSecretResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientRotateSecretResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OAuthClientRotateSecretResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    oauthClientRotateSecretResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// oauthClientRotateSecretResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [OAuthClientRotateSecretResponseEnvelopeErrorsSource]
type oauthClientRotateSecretResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientRotateSecretResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientRotateSecretResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OAuthClientRotateSecretResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           OAuthClientRotateSecretResponseEnvelopeMessagesSource `json:"source"`
	JSON             oauthClientRotateSecretResponseEnvelopeMessagesJSON   `json:"-"`
}

// oauthClientRotateSecretResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OAuthClientRotateSecretResponseEnvelopeMessages]
type oauthClientRotateSecretResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OAuthClientRotateSecretResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientRotateSecretResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OAuthClientRotateSecretResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    oauthClientRotateSecretResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// oauthClientRotateSecretResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [OAuthClientRotateSecretResponseEnvelopeMessagesSource]
type oauthClientRotateSecretResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthClientRotateSecretResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthClientRotateSecretResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OAuthClientRotateSecretResponseEnvelopeSuccess bool

const (
	OAuthClientRotateSecretResponseEnvelopeSuccessTrue OAuthClientRotateSecretResponseEnvelopeSuccess = true
)

func (r OAuthClientRotateSecretResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OAuthClientRotateSecretResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

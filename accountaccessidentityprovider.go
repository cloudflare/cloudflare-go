// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAccessIdentityProviderService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAccessIdentityProviderService] method instead.
type AccountAccessIdentityProviderService struct {
	Options []option.RequestOption
}

// NewAccountAccessIdentityProviderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessIdentityProviderService(opts ...option.RequestOption) (r *AccountAccessIdentityProviderService) {
	r = &AccountAccessIdentityProviderService{}
	r.Options = opts
	return
}

// Fetches a configured identity provider.
func (r *AccountAccessIdentityProviderService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *SchemasSingleResponseEiGa0KyX, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/identity_providers/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured identity provider.
func (r *AccountAccessIdentityProviderService) Update(ctx context.Context, identifier interface{}, uuid string, body AccountAccessIdentityProviderUpdateParams, opts ...option.RequestOption) (res *SchemasSingleResponseEiGa0KyX, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/identity_providers/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an identity provider from Access.
func (r *AccountAccessIdentityProviderService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *APIResponseSingleIDY7yIk9Qy, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/identity_providers/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new identity provider to Access.
func (r *AccountAccessIdentityProviderService) AccessIdentityProvidersAddAnAccessIdentityProvider(ctx context.Context, identifier interface{}, body AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams, opts ...option.RequestOption) (res *SchemasSingleResponseEiGa0KyX, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/identity_providers", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all configured identity providers.
func (r *AccountAccessIdentityProviderService) AccessIdentityProvidersListAccessIdentityProviders(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *ResponseCollectionUy60SYxZ, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/identity_providers", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIResponseSingleIDY7yIk9Qy struct {
	Errors   []APIResponseSingleIDY7yIk9QyError   `json:"errors"`
	Messages []APIResponseSingleIDY7yIk9QyMessage `json:"messages"`
	Result   APIResponseSingleIDY7yIk9QyResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success APIResponseSingleIDY7yIk9QySuccess `json:"success"`
	JSON    apiResponseSingleIdy7yIk9QyJSON    `json:"-"`
}

// apiResponseSingleIdy7yIk9QyJSON contains the JSON metadata for the struct
// [APIResponseSingleIDY7yIk9Qy]
type apiResponseSingleIdy7yIk9QyJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIDY7yIk9Qy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIResponseSingleIDY7yIk9QyError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    apiResponseSingleIdy7yIk9QyErrorJSON `json:"-"`
}

// apiResponseSingleIdy7yIk9QyErrorJSON contains the JSON metadata for the struct
// [APIResponseSingleIDY7yIk9QyError]
type apiResponseSingleIdy7yIk9QyErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIDY7yIk9QyError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIResponseSingleIDY7yIk9QyMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    apiResponseSingleIdy7yIk9QyMessageJSON `json:"-"`
}

// apiResponseSingleIdy7yIk9QyMessageJSON contains the JSON metadata for the struct
// [APIResponseSingleIDY7yIk9QyMessage]
type apiResponseSingleIdy7yIk9QyMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIDY7yIk9QyMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIResponseSingleIDY7yIk9QyResult struct {
	// Identifier
	ID   string                                `json:"id,required"`
	JSON apiResponseSingleIdy7yIk9QyResultJSON `json:"-"`
}

// apiResponseSingleIdy7yIk9QyResultJSON contains the JSON metadata for the struct
// [APIResponseSingleIDY7yIk9QyResult]
type apiResponseSingleIdy7yIk9QyResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIDY7yIk9QyResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type APIResponseSingleIDY7yIk9QySuccess bool

const (
	APIResponseSingleIDY7yIk9QySuccessTrue APIResponseSingleIDY7yIk9QySuccess = true
)

type ResponseCollectionUy60SYxZ struct {
	Errors     []ResponseCollectionUy60SYxZError    `json:"errors"`
	Messages   []ResponseCollectionUy60SYxZMessage  `json:"messages"`
	Result     []ResponseCollectionUy60SYxZResult   `json:"result"`
	ResultInfo ResponseCollectionUy60SYxZResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ResponseCollectionUy60SYxZSuccess `json:"success"`
	JSON    responseCollectionUy60SYxZJSON    `json:"-"`
}

// responseCollectionUy60SYxZJSON contains the JSON metadata for the struct
// [ResponseCollectionUy60SYxZ]
type responseCollectionUy60SYxZJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZ) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    responseCollectionUy60SYxZErrorJSON `json:"-"`
}

// responseCollectionUy60SYxZErrorJSON contains the JSON metadata for the struct
// [ResponseCollectionUy60SYxZError]
type responseCollectionUy60SYxZErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    responseCollectionUy60SYxZMessageJSON `json:"-"`
}

// responseCollectionUy60SYxZMessageJSON contains the JSON metadata for the struct
// [ResponseCollectionUy60SYxZMessage]
type responseCollectionUy60SYxZMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ResponseCollectionUy60SYxZResultAzureAd],
// [ResponseCollectionUy60SYxZResultCentrify],
// [ResponseCollectionUy60SYxZResultFacebook],
// [ResponseCollectionUy60SYxZResultGitHub],
// [ResponseCollectionUy60SYxZResultGoogle],
// [ResponseCollectionUy60SYxZResultGoogleApps],
// [ResponseCollectionUy60SYxZResultLinkedin],
// [ResponseCollectionUy60SYxZResultOidc], [ResponseCollectionUy60SYxZResultOkta],
// [ResponseCollectionUy60SYxZResultOnelogin],
// [ResponseCollectionUy60SYxZResultPingone],
// [ResponseCollectionUy60SYxZResultSaml] or
// [ResponseCollectionUy60SYxZResultYandex].
type ResponseCollectionUy60SYxZResult interface {
	implementsResponseCollectionUy60SYxZResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ResponseCollectionUy60SYxZResult)(nil)).Elem(), "")
}

type ResponseCollectionUy60SYxZResultAzureAd struct {
	// UUID
	ID     string                                        `json:"id"`
	Config ResponseCollectionUy60SYxZResultAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ResponseCollectionUy60SYxZResultAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                      `json:"type"`
	JSON responseCollectionUy60SYxZResultAzureAdJSON `json:"-"`
}

// responseCollectionUy60SYxZResultAzureAdJSON contains the JSON metadata for the
// struct [ResponseCollectionUy60SYxZResultAzureAd]
type responseCollectionUy60SYxZResultAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ResponseCollectionUy60SYxZResultAzureAd) implementsResponseCollectionUy60SYxZResult() {}

type ResponseCollectionUy60SYxZResultAzureAdConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your Azure directory uuid
	DirectoryID string `json:"directory_id"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool                                              `json:"support_groups"`
	JSON          responseCollectionUy60SYxZResultAzureAdConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultAzureAdConfigJSON contains the JSON metadata for
// the struct [ResponseCollectionUy60SYxZResultAzureAdConfig]
type responseCollectionUy60SYxZResultAzureAdConfigJSON struct {
	ClientID      apijson.Field
	ClientSecret  apijson.Field
	DirectoryID   apijson.Field
	SupportGroups apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ResponseCollectionUy60SYxZResultAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                  `json:"user_deprovision"`
	JSON            responseCollectionUy60SYxZResultAzureAdScimConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultAzureAdScimConfigJSON contains the JSON metadata
// for the struct [ResponseCollectionUy60SYxZResultAzureAdScimConfig]
type responseCollectionUy60SYxZResultAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZResultCentrify struct {
	// UUID
	ID     string                                         `json:"id"`
	Config ResponseCollectionUy60SYxZResultCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ResponseCollectionUy60SYxZResultCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                       `json:"type"`
	JSON responseCollectionUy60SYxZResultCentrifyJSON `json:"-"`
}

// responseCollectionUy60SYxZResultCentrifyJSON contains the JSON metadata for the
// struct [ResponseCollectionUy60SYxZResultCentrify]
type responseCollectionUy60SYxZResultCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ResponseCollectionUy60SYxZResultCentrify) implementsResponseCollectionUy60SYxZResult() {}

type ResponseCollectionUy60SYxZResultCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount string `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID string `json:"centrify_app_id"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                             `json:"client_secret"`
	JSON         responseCollectionUy60SYxZResultCentrifyConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultCentrifyConfigJSON contains the JSON metadata
// for the struct [ResponseCollectionUy60SYxZResultCentrifyConfig]
type responseCollectionUy60SYxZResultCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ResponseCollectionUy60SYxZResultCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                   `json:"user_deprovision"`
	JSON            responseCollectionUy60SYxZResultCentrifyScimConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultCentrifyScimConfigJSON contains the JSON
// metadata for the struct [ResponseCollectionUy60SYxZResultCentrifyScimConfig]
type responseCollectionUy60SYxZResultCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZResultFacebook struct {
	// UUID
	ID     string                                         `json:"id"`
	Config ResponseCollectionUy60SYxZResultFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ResponseCollectionUy60SYxZResultFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                       `json:"type"`
	JSON responseCollectionUy60SYxZResultFacebookJSON `json:"-"`
}

// responseCollectionUy60SYxZResultFacebookJSON contains the JSON metadata for the
// struct [ResponseCollectionUy60SYxZResultFacebook]
type responseCollectionUy60SYxZResultFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ResponseCollectionUy60SYxZResultFacebook) implementsResponseCollectionUy60SYxZResult() {}

type ResponseCollectionUy60SYxZResultFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                             `json:"client_secret"`
	JSON         responseCollectionUy60SYxZResultFacebookConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultFacebookConfigJSON contains the JSON metadata
// for the struct [ResponseCollectionUy60SYxZResultFacebookConfig]
type responseCollectionUy60SYxZResultFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ResponseCollectionUy60SYxZResultFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                   `json:"user_deprovision"`
	JSON            responseCollectionUy60SYxZResultFacebookScimConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultFacebookScimConfigJSON contains the JSON
// metadata for the struct [ResponseCollectionUy60SYxZResultFacebookScimConfig]
type responseCollectionUy60SYxZResultFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZResultGitHub struct {
	// UUID
	ID     string                                       `json:"id"`
	Config ResponseCollectionUy60SYxZResultGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ResponseCollectionUy60SYxZResultGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                     `json:"type"`
	JSON responseCollectionUy60SYxZResultGitHubJSON `json:"-"`
}

// responseCollectionUy60SYxZResultGitHubJSON contains the JSON metadata for the
// struct [ResponseCollectionUy60SYxZResultGitHub]
type responseCollectionUy60SYxZResultGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ResponseCollectionUy60SYxZResultGitHub) implementsResponseCollectionUy60SYxZResult() {}

type ResponseCollectionUy60SYxZResultGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                           `json:"client_secret"`
	JSON         responseCollectionUy60SYxZResultGitHubConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultGitHubConfigJSON contains the JSON metadata for
// the struct [ResponseCollectionUy60SYxZResultGitHubConfig]
type responseCollectionUy60SYxZResultGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ResponseCollectionUy60SYxZResultGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                 `json:"user_deprovision"`
	JSON            responseCollectionUy60SYxZResultGitHubScimConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultGitHubScimConfigJSON contains the JSON metadata
// for the struct [ResponseCollectionUy60SYxZResultGitHubScimConfig]
type responseCollectionUy60SYxZResultGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZResultGoogle struct {
	// UUID
	ID     string                                       `json:"id"`
	Config ResponseCollectionUy60SYxZResultGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ResponseCollectionUy60SYxZResultGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                     `json:"type"`
	JSON responseCollectionUy60SYxZResultGoogleJSON `json:"-"`
}

// responseCollectionUy60SYxZResultGoogleJSON contains the JSON metadata for the
// struct [ResponseCollectionUy60SYxZResultGoogle]
type responseCollectionUy60SYxZResultGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ResponseCollectionUy60SYxZResultGoogle) implementsResponseCollectionUy60SYxZResult() {}

type ResponseCollectionUy60SYxZResultGoogleConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                           `json:"client_secret"`
	JSON         responseCollectionUy60SYxZResultGoogleConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultGoogleConfigJSON contains the JSON metadata for
// the struct [ResponseCollectionUy60SYxZResultGoogleConfig]
type responseCollectionUy60SYxZResultGoogleConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ResponseCollectionUy60SYxZResultGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                 `json:"user_deprovision"`
	JSON            responseCollectionUy60SYxZResultGoogleScimConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultGoogleScimConfigJSON contains the JSON metadata
// for the struct [ResponseCollectionUy60SYxZResultGoogleScimConfig]
type responseCollectionUy60SYxZResultGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZResultGoogleApps struct {
	// UUID
	ID     string                                           `json:"id"`
	Config ResponseCollectionUy60SYxZResultGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ResponseCollectionUy60SYxZResultGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                         `json:"type"`
	JSON responseCollectionUy60SYxZResultGoogleAppsJSON `json:"-"`
}

// responseCollectionUy60SYxZResultGoogleAppsJSON contains the JSON metadata for
// the struct [ResponseCollectionUy60SYxZResultGoogleApps]
type responseCollectionUy60SYxZResultGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ResponseCollectionUy60SYxZResultGoogleApps) implementsResponseCollectionUy60SYxZResult() {}

type ResponseCollectionUy60SYxZResultGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                               `json:"client_secret"`
	JSON         responseCollectionUy60SYxZResultGoogleAppsConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultGoogleAppsConfigJSON contains the JSON metadata
// for the struct [ResponseCollectionUy60SYxZResultGoogleAppsConfig]
type responseCollectionUy60SYxZResultGoogleAppsConfigJSON struct {
	AppsDomain   apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ResponseCollectionUy60SYxZResultGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                     `json:"user_deprovision"`
	JSON            responseCollectionUy60SYxZResultGoogleAppsScimConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultGoogleAppsScimConfigJSON contains the JSON
// metadata for the struct [ResponseCollectionUy60SYxZResultGoogleAppsScimConfig]
type responseCollectionUy60SYxZResultGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZResultLinkedin struct {
	// UUID
	ID     string                                         `json:"id"`
	Config ResponseCollectionUy60SYxZResultLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ResponseCollectionUy60SYxZResultLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                       `json:"type"`
	JSON responseCollectionUy60SYxZResultLinkedinJSON `json:"-"`
}

// responseCollectionUy60SYxZResultLinkedinJSON contains the JSON metadata for the
// struct [ResponseCollectionUy60SYxZResultLinkedin]
type responseCollectionUy60SYxZResultLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ResponseCollectionUy60SYxZResultLinkedin) implementsResponseCollectionUy60SYxZResult() {}

type ResponseCollectionUy60SYxZResultLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                             `json:"client_secret"`
	JSON         responseCollectionUy60SYxZResultLinkedinConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultLinkedinConfigJSON contains the JSON metadata
// for the struct [ResponseCollectionUy60SYxZResultLinkedinConfig]
type responseCollectionUy60SYxZResultLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ResponseCollectionUy60SYxZResultLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                   `json:"user_deprovision"`
	JSON            responseCollectionUy60SYxZResultLinkedinScimConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultLinkedinScimConfigJSON contains the JSON
// metadata for the struct [ResponseCollectionUy60SYxZResultLinkedinScimConfig]
type responseCollectionUy60SYxZResultLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZResultOidc struct {
	// UUID
	ID     string                                     `json:"id"`
	Config ResponseCollectionUy60SYxZResultOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ResponseCollectionUy60SYxZResultOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                   `json:"type"`
	JSON responseCollectionUy60SYxZResultOidcJSON `json:"-"`
}

// responseCollectionUy60SYxZResultOidcJSON contains the JSON metadata for the
// struct [ResponseCollectionUy60SYxZResultOidc]
type responseCollectionUy60SYxZResultOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ResponseCollectionUy60SYxZResultOidc) implementsResponseCollectionUy60SYxZResult() {}

type ResponseCollectionUy60SYxZResultOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL string `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL string `json:"certs_url"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// OAuth scopes
	Scopes []string `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL string                                         `json:"token_url"`
	JSON     responseCollectionUy60SYxZResultOidcConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultOidcConfigJSON contains the JSON metadata for
// the struct [ResponseCollectionUy60SYxZResultOidcConfig]
type responseCollectionUy60SYxZResultOidcConfigJSON struct {
	AuthURL      apijson.Field
	CertsURL     apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	Scopes       apijson.Field
	TokenURL     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ResponseCollectionUy60SYxZResultOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                               `json:"user_deprovision"`
	JSON            responseCollectionUy60SYxZResultOidcScimConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultOidcScimConfigJSON contains the JSON metadata
// for the struct [ResponseCollectionUy60SYxZResultOidcScimConfig]
type responseCollectionUy60SYxZResultOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZResultOkta struct {
	// UUID
	ID     string                                     `json:"id"`
	Config ResponseCollectionUy60SYxZResultOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ResponseCollectionUy60SYxZResultOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                   `json:"type"`
	JSON responseCollectionUy60SYxZResultOktaJSON `json:"-"`
}

// responseCollectionUy60SYxZResultOktaJSON contains the JSON metadata for the
// struct [ResponseCollectionUy60SYxZResultOkta]
type responseCollectionUy60SYxZResultOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ResponseCollectionUy60SYxZResultOkta) implementsResponseCollectionUy60SYxZResult() {}

type ResponseCollectionUy60SYxZResultOktaConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your okta account url
	OktaAccount string                                         `json:"okta_account"`
	JSON        responseCollectionUy60SYxZResultOktaConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultOktaConfigJSON contains the JSON metadata for
// the struct [ResponseCollectionUy60SYxZResultOktaConfig]
type responseCollectionUy60SYxZResultOktaConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	OktaAccount  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ResponseCollectionUy60SYxZResultOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                               `json:"user_deprovision"`
	JSON            responseCollectionUy60SYxZResultOktaScimConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultOktaScimConfigJSON contains the JSON metadata
// for the struct [ResponseCollectionUy60SYxZResultOktaScimConfig]
type responseCollectionUy60SYxZResultOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZResultOnelogin struct {
	// UUID
	ID     string                                         `json:"id"`
	Config ResponseCollectionUy60SYxZResultOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ResponseCollectionUy60SYxZResultOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                       `json:"type"`
	JSON responseCollectionUy60SYxZResultOneloginJSON `json:"-"`
}

// responseCollectionUy60SYxZResultOneloginJSON contains the JSON metadata for the
// struct [ResponseCollectionUy60SYxZResultOnelogin]
type responseCollectionUy60SYxZResultOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ResponseCollectionUy60SYxZResultOnelogin) implementsResponseCollectionUy60SYxZResult() {}

type ResponseCollectionUy60SYxZResultOneloginConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your OneLogin account url
	OneloginAccount string                                             `json:"onelogin_account"`
	JSON            responseCollectionUy60SYxZResultOneloginConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultOneloginConfigJSON contains the JSON metadata
// for the struct [ResponseCollectionUy60SYxZResultOneloginConfig]
type responseCollectionUy60SYxZResultOneloginConfigJSON struct {
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ResponseCollectionUy60SYxZResultOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                   `json:"user_deprovision"`
	JSON            responseCollectionUy60SYxZResultOneloginScimConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultOneloginScimConfigJSON contains the JSON
// metadata for the struct [ResponseCollectionUy60SYxZResultOneloginScimConfig]
type responseCollectionUy60SYxZResultOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZResultPingone struct {
	// UUID
	ID     string                                        `json:"id"`
	Config ResponseCollectionUy60SYxZResultPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ResponseCollectionUy60SYxZResultPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                      `json:"type"`
	JSON responseCollectionUy60SYxZResultPingoneJSON `json:"-"`
}

// responseCollectionUy60SYxZResultPingoneJSON contains the JSON metadata for the
// struct [ResponseCollectionUy60SYxZResultPingone]
type responseCollectionUy60SYxZResultPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ResponseCollectionUy60SYxZResultPingone) implementsResponseCollectionUy60SYxZResult() {}

type ResponseCollectionUy60SYxZResultPingoneConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your PingOne environment identifier
	PingEnvID string                                            `json:"ping_env_id"`
	JSON      responseCollectionUy60SYxZResultPingoneConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultPingoneConfigJSON contains the JSON metadata for
// the struct [ResponseCollectionUy60SYxZResultPingoneConfig]
type responseCollectionUy60SYxZResultPingoneConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	PingEnvID    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ResponseCollectionUy60SYxZResultPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                  `json:"user_deprovision"`
	JSON            responseCollectionUy60SYxZResultPingoneScimConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultPingoneScimConfigJSON contains the JSON metadata
// for the struct [ResponseCollectionUy60SYxZResultPingoneScimConfig]
type responseCollectionUy60SYxZResultPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZResultSaml struct {
	// UUID
	ID     string                                     `json:"id"`
	Config ResponseCollectionUy60SYxZResultSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ResponseCollectionUy60SYxZResultSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                   `json:"type"`
	JSON responseCollectionUy60SYxZResultSamlJSON `json:"-"`
}

// responseCollectionUy60SYxZResultSamlJSON contains the JSON metadata for the
// struct [ResponseCollectionUy60SYxZResultSaml]
type responseCollectionUy60SYxZResultSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ResponseCollectionUy60SYxZResultSaml) implementsResponseCollectionUy60SYxZResult() {}

type ResponseCollectionUy60SYxZResultSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []ResponseCollectionUy60SYxZResultSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                         `json:"sso_target_url"`
	JSON         responseCollectionUy60SYxZResultSamlConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultSamlConfigJSON contains the JSON metadata for
// the struct [ResponseCollectionUy60SYxZResultSamlConfig]
type responseCollectionUy60SYxZResultSamlConfigJSON struct {
	Attributes         apijson.Field
	EmailAttributeName apijson.Field
	HeaderAttributes   apijson.Field
	IdpPublicCerts     apijson.Field
	IssuerURL          apijson.Field
	SignRequest        apijson.Field
	SSOTargetURL       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZResultSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                        `json:"header_name"`
	JSON       responseCollectionUy60SYxZResultSamlConfigHeaderAttributeJSON `json:"-"`
}

// responseCollectionUy60SYxZResultSamlConfigHeaderAttributeJSON contains the JSON
// metadata for the struct
// [ResponseCollectionUy60SYxZResultSamlConfigHeaderAttribute]
type responseCollectionUy60SYxZResultSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ResponseCollectionUy60SYxZResultSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                               `json:"user_deprovision"`
	JSON            responseCollectionUy60SYxZResultSamlScimConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultSamlScimConfigJSON contains the JSON metadata
// for the struct [ResponseCollectionUy60SYxZResultSamlScimConfig]
type responseCollectionUy60SYxZResultSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZResultYandex struct {
	// UUID
	ID     string                                       `json:"id"`
	Config ResponseCollectionUy60SYxZResultYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ResponseCollectionUy60SYxZResultYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                     `json:"type"`
	JSON responseCollectionUy60SYxZResultYandexJSON `json:"-"`
}

// responseCollectionUy60SYxZResultYandexJSON contains the JSON metadata for the
// struct [ResponseCollectionUy60SYxZResultYandex]
type responseCollectionUy60SYxZResultYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ResponseCollectionUy60SYxZResultYandex) implementsResponseCollectionUy60SYxZResult() {}

type ResponseCollectionUy60SYxZResultYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                           `json:"client_secret"`
	JSON         responseCollectionUy60SYxZResultYandexConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultYandexConfigJSON contains the JSON metadata for
// the struct [ResponseCollectionUy60SYxZResultYandexConfig]
type responseCollectionUy60SYxZResultYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ResponseCollectionUy60SYxZResultYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                 `json:"user_deprovision"`
	JSON            responseCollectionUy60SYxZResultYandexScimConfigJSON `json:"-"`
}

// responseCollectionUy60SYxZResultYandexScimConfigJSON contains the JSON metadata
// for the struct [ResponseCollectionUy60SYxZResultYandexScimConfig]
type responseCollectionUy60SYxZResultYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUy60SYxZResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       responseCollectionUy60SYxZResultInfoJSON `json:"-"`
}

// responseCollectionUy60SYxZResultInfoJSON contains the JSON metadata for the
// struct [ResponseCollectionUy60SYxZResultInfo]
type responseCollectionUy60SYxZResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUy60SYxZResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ResponseCollectionUy60SYxZSuccess bool

const (
	ResponseCollectionUy60SYxZSuccessTrue ResponseCollectionUy60SYxZSuccess = true
)

type SchemasSingleResponseEiGa0KyX struct {
	Errors   []SchemasSingleResponseEiGa0KyXError   `json:"errors"`
	Messages []SchemasSingleResponseEiGa0KyXMessage `json:"messages"`
	Result   SchemasSingleResponseEiGa0KyXResult    `json:"result"`
	// Whether the API call was successful
	Success SchemasSingleResponseEiGa0KyXSuccess `json:"success"`
	JSON    schemasSingleResponseEiGa0KyXJSON    `json:"-"`
}

// schemasSingleResponseEiGa0KyXJSON contains the JSON metadata for the struct
// [SchemasSingleResponseEiGa0KyX]
type schemasSingleResponseEiGa0KyXJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyX) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseEiGa0KyXError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    schemasSingleResponseEiGa0KyXErrorJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXErrorJSON contains the JSON metadata for the struct
// [SchemasSingleResponseEiGa0KyXError]
type schemasSingleResponseEiGa0KyXErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseEiGa0KyXMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    schemasSingleResponseEiGa0KyXMessageJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXMessageJSON contains the JSON metadata for the
// struct [SchemasSingleResponseEiGa0KyXMessage]
type schemasSingleResponseEiGa0KyXMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [SchemasSingleResponseEiGa0KyXResultAzureAd],
// [SchemasSingleResponseEiGa0KyXResultCentrify],
// [SchemasSingleResponseEiGa0KyXResultFacebook],
// [SchemasSingleResponseEiGa0KyXResultGitHub],
// [SchemasSingleResponseEiGa0KyXResultGoogle],
// [SchemasSingleResponseEiGa0KyXResultGoogleApps],
// [SchemasSingleResponseEiGa0KyXResultLinkedin],
// [SchemasSingleResponseEiGa0KyXResultOidc],
// [SchemasSingleResponseEiGa0KyXResultOkta],
// [SchemasSingleResponseEiGa0KyXResultOnelogin],
// [SchemasSingleResponseEiGa0KyXResultPingone],
// [SchemasSingleResponseEiGa0KyXResultSaml] or
// [SchemasSingleResponseEiGa0KyXResultYandex].
type SchemasSingleResponseEiGa0KyXResult interface {
	implementsSchemasSingleResponseEiGa0KyXResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*SchemasSingleResponseEiGa0KyXResult)(nil)).Elem(), "")
}

type SchemasSingleResponseEiGa0KyXResultAzureAd struct {
	// UUID
	ID     string                                           `json:"id"`
	Config SchemasSingleResponseEiGa0KyXResultAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig SchemasSingleResponseEiGa0KyXResultAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                         `json:"type"`
	JSON schemasSingleResponseEiGa0KyXResultAzureAdJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultAzureAdJSON contains the JSON metadata for
// the struct [SchemasSingleResponseEiGa0KyXResultAzureAd]
type schemasSingleResponseEiGa0KyXResultAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasSingleResponseEiGa0KyXResultAzureAd) implementsSchemasSingleResponseEiGa0KyXResult() {}

type SchemasSingleResponseEiGa0KyXResultAzureAdConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your Azure directory uuid
	DirectoryID string `json:"directory_id"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool                                                 `json:"support_groups"`
	JSON          schemasSingleResponseEiGa0KyXResultAzureAdConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultAzureAdConfigJSON contains the JSON metadata
// for the struct [SchemasSingleResponseEiGa0KyXResultAzureAdConfig]
type schemasSingleResponseEiGa0KyXResultAzureAdConfigJSON struct {
	ClientID      apijson.Field
	ClientSecret  apijson.Field
	DirectoryID   apijson.Field
	SupportGroups apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type SchemasSingleResponseEiGa0KyXResultAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                     `json:"user_deprovision"`
	JSON            schemasSingleResponseEiGa0KyXResultAzureAdScimConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultAzureAdScimConfigJSON contains the JSON
// metadata for the struct [SchemasSingleResponseEiGa0KyXResultAzureAdScimConfig]
type schemasSingleResponseEiGa0KyXResultAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseEiGa0KyXResultCentrify struct {
	// UUID
	ID     string                                            `json:"id"`
	Config SchemasSingleResponseEiGa0KyXResultCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig SchemasSingleResponseEiGa0KyXResultCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                          `json:"type"`
	JSON schemasSingleResponseEiGa0KyXResultCentrifyJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultCentrifyJSON contains the JSON metadata for
// the struct [SchemasSingleResponseEiGa0KyXResultCentrify]
type schemasSingleResponseEiGa0KyXResultCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasSingleResponseEiGa0KyXResultCentrify) implementsSchemasSingleResponseEiGa0KyXResult() {
}

type SchemasSingleResponseEiGa0KyXResultCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount string `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID string `json:"centrify_app_id"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                `json:"client_secret"`
	JSON         schemasSingleResponseEiGa0KyXResultCentrifyConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultCentrifyConfigJSON contains the JSON metadata
// for the struct [SchemasSingleResponseEiGa0KyXResultCentrifyConfig]
type schemasSingleResponseEiGa0KyXResultCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type SchemasSingleResponseEiGa0KyXResultCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                      `json:"user_deprovision"`
	JSON            schemasSingleResponseEiGa0KyXResultCentrifyScimConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultCentrifyScimConfigJSON contains the JSON
// metadata for the struct [SchemasSingleResponseEiGa0KyXResultCentrifyScimConfig]
type schemasSingleResponseEiGa0KyXResultCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseEiGa0KyXResultFacebook struct {
	// UUID
	ID     string                                            `json:"id"`
	Config SchemasSingleResponseEiGa0KyXResultFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig SchemasSingleResponseEiGa0KyXResultFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                          `json:"type"`
	JSON schemasSingleResponseEiGa0KyXResultFacebookJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultFacebookJSON contains the JSON metadata for
// the struct [SchemasSingleResponseEiGa0KyXResultFacebook]
type schemasSingleResponseEiGa0KyXResultFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasSingleResponseEiGa0KyXResultFacebook) implementsSchemasSingleResponseEiGa0KyXResult() {
}

type SchemasSingleResponseEiGa0KyXResultFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                `json:"client_secret"`
	JSON         schemasSingleResponseEiGa0KyXResultFacebookConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultFacebookConfigJSON contains the JSON metadata
// for the struct [SchemasSingleResponseEiGa0KyXResultFacebookConfig]
type schemasSingleResponseEiGa0KyXResultFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type SchemasSingleResponseEiGa0KyXResultFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                      `json:"user_deprovision"`
	JSON            schemasSingleResponseEiGa0KyXResultFacebookScimConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultFacebookScimConfigJSON contains the JSON
// metadata for the struct [SchemasSingleResponseEiGa0KyXResultFacebookScimConfig]
type schemasSingleResponseEiGa0KyXResultFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseEiGa0KyXResultGitHub struct {
	// UUID
	ID     string                                          `json:"id"`
	Config SchemasSingleResponseEiGa0KyXResultGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig SchemasSingleResponseEiGa0KyXResultGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                        `json:"type"`
	JSON schemasSingleResponseEiGa0KyXResultGitHubJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultGitHubJSON contains the JSON metadata for the
// struct [SchemasSingleResponseEiGa0KyXResultGitHub]
type schemasSingleResponseEiGa0KyXResultGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasSingleResponseEiGa0KyXResultGitHub) implementsSchemasSingleResponseEiGa0KyXResult() {}

type SchemasSingleResponseEiGa0KyXResultGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                              `json:"client_secret"`
	JSON         schemasSingleResponseEiGa0KyXResultGitHubConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultGitHubConfigJSON contains the JSON metadata
// for the struct [SchemasSingleResponseEiGa0KyXResultGitHubConfig]
type schemasSingleResponseEiGa0KyXResultGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type SchemasSingleResponseEiGa0KyXResultGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                    `json:"user_deprovision"`
	JSON            schemasSingleResponseEiGa0KyXResultGitHubScimConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultGitHubScimConfigJSON contains the JSON
// metadata for the struct [SchemasSingleResponseEiGa0KyXResultGitHubScimConfig]
type schemasSingleResponseEiGa0KyXResultGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseEiGa0KyXResultGoogle struct {
	// UUID
	ID     string                                          `json:"id"`
	Config SchemasSingleResponseEiGa0KyXResultGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig SchemasSingleResponseEiGa0KyXResultGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                        `json:"type"`
	JSON schemasSingleResponseEiGa0KyXResultGoogleJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultGoogleJSON contains the JSON metadata for the
// struct [SchemasSingleResponseEiGa0KyXResultGoogle]
type schemasSingleResponseEiGa0KyXResultGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasSingleResponseEiGa0KyXResultGoogle) implementsSchemasSingleResponseEiGa0KyXResult() {}

type SchemasSingleResponseEiGa0KyXResultGoogleConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                              `json:"client_secret"`
	JSON         schemasSingleResponseEiGa0KyXResultGoogleConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultGoogleConfigJSON contains the JSON metadata
// for the struct [SchemasSingleResponseEiGa0KyXResultGoogleConfig]
type schemasSingleResponseEiGa0KyXResultGoogleConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type SchemasSingleResponseEiGa0KyXResultGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                    `json:"user_deprovision"`
	JSON            schemasSingleResponseEiGa0KyXResultGoogleScimConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultGoogleScimConfigJSON contains the JSON
// metadata for the struct [SchemasSingleResponseEiGa0KyXResultGoogleScimConfig]
type schemasSingleResponseEiGa0KyXResultGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseEiGa0KyXResultGoogleApps struct {
	// UUID
	ID     string                                              `json:"id"`
	Config SchemasSingleResponseEiGa0KyXResultGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig SchemasSingleResponseEiGa0KyXResultGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                            `json:"type"`
	JSON schemasSingleResponseEiGa0KyXResultGoogleAppsJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultGoogleAppsJSON contains the JSON metadata for
// the struct [SchemasSingleResponseEiGa0KyXResultGoogleApps]
type schemasSingleResponseEiGa0KyXResultGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasSingleResponseEiGa0KyXResultGoogleApps) implementsSchemasSingleResponseEiGa0KyXResult() {
}

type SchemasSingleResponseEiGa0KyXResultGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                  `json:"client_secret"`
	JSON         schemasSingleResponseEiGa0KyXResultGoogleAppsConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultGoogleAppsConfigJSON contains the JSON
// metadata for the struct [SchemasSingleResponseEiGa0KyXResultGoogleAppsConfig]
type schemasSingleResponseEiGa0KyXResultGoogleAppsConfigJSON struct {
	AppsDomain   apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type SchemasSingleResponseEiGa0KyXResultGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                        `json:"user_deprovision"`
	JSON            schemasSingleResponseEiGa0KyXResultGoogleAppsScimConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultGoogleAppsScimConfigJSON contains the JSON
// metadata for the struct
// [SchemasSingleResponseEiGa0KyXResultGoogleAppsScimConfig]
type schemasSingleResponseEiGa0KyXResultGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseEiGa0KyXResultLinkedin struct {
	// UUID
	ID     string                                            `json:"id"`
	Config SchemasSingleResponseEiGa0KyXResultLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig SchemasSingleResponseEiGa0KyXResultLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                          `json:"type"`
	JSON schemasSingleResponseEiGa0KyXResultLinkedinJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultLinkedinJSON contains the JSON metadata for
// the struct [SchemasSingleResponseEiGa0KyXResultLinkedin]
type schemasSingleResponseEiGa0KyXResultLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasSingleResponseEiGa0KyXResultLinkedin) implementsSchemasSingleResponseEiGa0KyXResult() {
}

type SchemasSingleResponseEiGa0KyXResultLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                `json:"client_secret"`
	JSON         schemasSingleResponseEiGa0KyXResultLinkedinConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultLinkedinConfigJSON contains the JSON metadata
// for the struct [SchemasSingleResponseEiGa0KyXResultLinkedinConfig]
type schemasSingleResponseEiGa0KyXResultLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type SchemasSingleResponseEiGa0KyXResultLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                      `json:"user_deprovision"`
	JSON            schemasSingleResponseEiGa0KyXResultLinkedinScimConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultLinkedinScimConfigJSON contains the JSON
// metadata for the struct [SchemasSingleResponseEiGa0KyXResultLinkedinScimConfig]
type schemasSingleResponseEiGa0KyXResultLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseEiGa0KyXResultOidc struct {
	// UUID
	ID     string                                        `json:"id"`
	Config SchemasSingleResponseEiGa0KyXResultOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig SchemasSingleResponseEiGa0KyXResultOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                      `json:"type"`
	JSON schemasSingleResponseEiGa0KyXResultOidcJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultOidcJSON contains the JSON metadata for the
// struct [SchemasSingleResponseEiGa0KyXResultOidc]
type schemasSingleResponseEiGa0KyXResultOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasSingleResponseEiGa0KyXResultOidc) implementsSchemasSingleResponseEiGa0KyXResult() {}

type SchemasSingleResponseEiGa0KyXResultOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL string `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL string `json:"certs_url"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// OAuth scopes
	Scopes []string `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL string                                            `json:"token_url"`
	JSON     schemasSingleResponseEiGa0KyXResultOidcConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultOidcConfigJSON contains the JSON metadata for
// the struct [SchemasSingleResponseEiGa0KyXResultOidcConfig]
type schemasSingleResponseEiGa0KyXResultOidcConfigJSON struct {
	AuthURL      apijson.Field
	CertsURL     apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	Scopes       apijson.Field
	TokenURL     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type SchemasSingleResponseEiGa0KyXResultOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                  `json:"user_deprovision"`
	JSON            schemasSingleResponseEiGa0KyXResultOidcScimConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultOidcScimConfigJSON contains the JSON metadata
// for the struct [SchemasSingleResponseEiGa0KyXResultOidcScimConfig]
type schemasSingleResponseEiGa0KyXResultOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseEiGa0KyXResultOkta struct {
	// UUID
	ID     string                                        `json:"id"`
	Config SchemasSingleResponseEiGa0KyXResultOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig SchemasSingleResponseEiGa0KyXResultOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                      `json:"type"`
	JSON schemasSingleResponseEiGa0KyXResultOktaJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultOktaJSON contains the JSON metadata for the
// struct [SchemasSingleResponseEiGa0KyXResultOkta]
type schemasSingleResponseEiGa0KyXResultOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasSingleResponseEiGa0KyXResultOkta) implementsSchemasSingleResponseEiGa0KyXResult() {}

type SchemasSingleResponseEiGa0KyXResultOktaConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your okta account url
	OktaAccount string                                            `json:"okta_account"`
	JSON        schemasSingleResponseEiGa0KyXResultOktaConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultOktaConfigJSON contains the JSON metadata for
// the struct [SchemasSingleResponseEiGa0KyXResultOktaConfig]
type schemasSingleResponseEiGa0KyXResultOktaConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	OktaAccount  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type SchemasSingleResponseEiGa0KyXResultOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                  `json:"user_deprovision"`
	JSON            schemasSingleResponseEiGa0KyXResultOktaScimConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultOktaScimConfigJSON contains the JSON metadata
// for the struct [SchemasSingleResponseEiGa0KyXResultOktaScimConfig]
type schemasSingleResponseEiGa0KyXResultOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseEiGa0KyXResultOnelogin struct {
	// UUID
	ID     string                                            `json:"id"`
	Config SchemasSingleResponseEiGa0KyXResultOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig SchemasSingleResponseEiGa0KyXResultOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                          `json:"type"`
	JSON schemasSingleResponseEiGa0KyXResultOneloginJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultOneloginJSON contains the JSON metadata for
// the struct [SchemasSingleResponseEiGa0KyXResultOnelogin]
type schemasSingleResponseEiGa0KyXResultOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasSingleResponseEiGa0KyXResultOnelogin) implementsSchemasSingleResponseEiGa0KyXResult() {
}

type SchemasSingleResponseEiGa0KyXResultOneloginConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your OneLogin account url
	OneloginAccount string                                                `json:"onelogin_account"`
	JSON            schemasSingleResponseEiGa0KyXResultOneloginConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultOneloginConfigJSON contains the JSON metadata
// for the struct [SchemasSingleResponseEiGa0KyXResultOneloginConfig]
type schemasSingleResponseEiGa0KyXResultOneloginConfigJSON struct {
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type SchemasSingleResponseEiGa0KyXResultOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                      `json:"user_deprovision"`
	JSON            schemasSingleResponseEiGa0KyXResultOneloginScimConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultOneloginScimConfigJSON contains the JSON
// metadata for the struct [SchemasSingleResponseEiGa0KyXResultOneloginScimConfig]
type schemasSingleResponseEiGa0KyXResultOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseEiGa0KyXResultPingone struct {
	// UUID
	ID     string                                           `json:"id"`
	Config SchemasSingleResponseEiGa0KyXResultPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig SchemasSingleResponseEiGa0KyXResultPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                         `json:"type"`
	JSON schemasSingleResponseEiGa0KyXResultPingoneJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultPingoneJSON contains the JSON metadata for
// the struct [SchemasSingleResponseEiGa0KyXResultPingone]
type schemasSingleResponseEiGa0KyXResultPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasSingleResponseEiGa0KyXResultPingone) implementsSchemasSingleResponseEiGa0KyXResult() {}

type SchemasSingleResponseEiGa0KyXResultPingoneConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your PingOne environment identifier
	PingEnvID string                                               `json:"ping_env_id"`
	JSON      schemasSingleResponseEiGa0KyXResultPingoneConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultPingoneConfigJSON contains the JSON metadata
// for the struct [SchemasSingleResponseEiGa0KyXResultPingoneConfig]
type schemasSingleResponseEiGa0KyXResultPingoneConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	PingEnvID    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type SchemasSingleResponseEiGa0KyXResultPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                     `json:"user_deprovision"`
	JSON            schemasSingleResponseEiGa0KyXResultPingoneScimConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultPingoneScimConfigJSON contains the JSON
// metadata for the struct [SchemasSingleResponseEiGa0KyXResultPingoneScimConfig]
type schemasSingleResponseEiGa0KyXResultPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseEiGa0KyXResultSaml struct {
	// UUID
	ID     string                                        `json:"id"`
	Config SchemasSingleResponseEiGa0KyXResultSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig SchemasSingleResponseEiGa0KyXResultSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                      `json:"type"`
	JSON schemasSingleResponseEiGa0KyXResultSamlJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultSamlJSON contains the JSON metadata for the
// struct [SchemasSingleResponseEiGa0KyXResultSaml]
type schemasSingleResponseEiGa0KyXResultSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasSingleResponseEiGa0KyXResultSaml) implementsSchemasSingleResponseEiGa0KyXResult() {}

type SchemasSingleResponseEiGa0KyXResultSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []SchemasSingleResponseEiGa0KyXResultSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                            `json:"sso_target_url"`
	JSON         schemasSingleResponseEiGa0KyXResultSamlConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultSamlConfigJSON contains the JSON metadata for
// the struct [SchemasSingleResponseEiGa0KyXResultSamlConfig]
type schemasSingleResponseEiGa0KyXResultSamlConfigJSON struct {
	Attributes         apijson.Field
	EmailAttributeName apijson.Field
	HeaderAttributes   apijson.Field
	IdpPublicCerts     apijson.Field
	IssuerURL          apijson.Field
	SignRequest        apijson.Field
	SSOTargetURL       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseEiGa0KyXResultSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                           `json:"header_name"`
	JSON       schemasSingleResponseEiGa0KyXResultSamlConfigHeaderAttributeJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultSamlConfigHeaderAttributeJSON contains the
// JSON metadata for the struct
// [SchemasSingleResponseEiGa0KyXResultSamlConfigHeaderAttribute]
type schemasSingleResponseEiGa0KyXResultSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type SchemasSingleResponseEiGa0KyXResultSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                  `json:"user_deprovision"`
	JSON            schemasSingleResponseEiGa0KyXResultSamlScimConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultSamlScimConfigJSON contains the JSON metadata
// for the struct [SchemasSingleResponseEiGa0KyXResultSamlScimConfig]
type schemasSingleResponseEiGa0KyXResultSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseEiGa0KyXResultYandex struct {
	// UUID
	ID     string                                          `json:"id"`
	Config SchemasSingleResponseEiGa0KyXResultYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig SchemasSingleResponseEiGa0KyXResultYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                        `json:"type"`
	JSON schemasSingleResponseEiGa0KyXResultYandexJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultYandexJSON contains the JSON metadata for the
// struct [SchemasSingleResponseEiGa0KyXResultYandex]
type schemasSingleResponseEiGa0KyXResultYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasSingleResponseEiGa0KyXResultYandex) implementsSchemasSingleResponseEiGa0KyXResult() {}

type SchemasSingleResponseEiGa0KyXResultYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                              `json:"client_secret"`
	JSON         schemasSingleResponseEiGa0KyXResultYandexConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultYandexConfigJSON contains the JSON metadata
// for the struct [SchemasSingleResponseEiGa0KyXResultYandexConfig]
type schemasSingleResponseEiGa0KyXResultYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type SchemasSingleResponseEiGa0KyXResultYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                    `json:"user_deprovision"`
	JSON            schemasSingleResponseEiGa0KyXResultYandexScimConfigJSON `json:"-"`
}

// schemasSingleResponseEiGa0KyXResultYandexScimConfigJSON contains the JSON
// metadata for the struct [SchemasSingleResponseEiGa0KyXResultYandexScimConfig]
type schemasSingleResponseEiGa0KyXResultYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SchemasSingleResponseEiGa0KyXResultYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasSingleResponseEiGa0KyXSuccess bool

const (
	SchemasSingleResponseEiGa0KyXSuccessTrue SchemasSingleResponseEiGa0KyXSuccess = true
)

// This interface is a union satisfied by one of the following:
// [AccountAccessIdentityProviderUpdateParamsAzureAd],
// [AccountAccessIdentityProviderUpdateParamsCentrify],
// [AccountAccessIdentityProviderUpdateParamsFacebook],
// [AccountAccessIdentityProviderUpdateParamsGitHub],
// [AccountAccessIdentityProviderUpdateParamsGoogle],
// [AccountAccessIdentityProviderUpdateParamsGoogleApps],
// [AccountAccessIdentityProviderUpdateParamsLinkedin],
// [AccountAccessIdentityProviderUpdateParamsOidc],
// [AccountAccessIdentityProviderUpdateParamsOkta],
// [AccountAccessIdentityProviderUpdateParamsOnelogin],
// [AccountAccessIdentityProviderUpdateParamsPingone],
// [AccountAccessIdentityProviderUpdateParamsSaml],
// [AccountAccessIdentityProviderUpdateParamsYandex].
type AccountAccessIdentityProviderUpdateParams interface {
	ImplementsAccountAccessIdentityProviderUpdateParams()
}

type AccountAccessIdentityProviderUpdateParamsAzureAd struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsAzureAdConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsAzureAdScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsAzureAd) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsAzureAdConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your Azure directory uuid
	DirectoryID param.Field[string] `json:"directory_id"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
}

func (r AccountAccessIdentityProviderUpdateParamsAzureAdConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderUpdateParamsAzureAdScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderUpdateParamsCentrify struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsCentrifyConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsCentrifyScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsCentrify) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount param.Field[string] `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID param.Field[string] `json:"centrify_app_id"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderUpdateParamsCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderUpdateParamsCentrifyScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderUpdateParamsFacebook struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsFacebookConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsFacebookScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsFacebook) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsFacebookConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderUpdateParamsFacebookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderUpdateParamsFacebookScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderUpdateParamsGitHub struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsGitHubConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsGitHubScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsGitHub) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsGitHubConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderUpdateParamsGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderUpdateParamsGitHubScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderUpdateParamsGoogle struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsGoogleConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsGoogleScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsGoogle) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsGoogleConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderUpdateParamsGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderUpdateParamsGoogleScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderUpdateParamsGoogleApps struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsGoogleAppsConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsGoogleAppsScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsGoogleApps) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain param.Field[string] `json:"apps_domain"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderUpdateParamsGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderUpdateParamsGoogleAppsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderUpdateParamsLinkedin struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsLinkedinConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsLinkedinScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsLinkedin) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderUpdateParamsLinkedinConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderUpdateParamsLinkedinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderUpdateParamsOidc struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsOidcConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsOidcScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsOidc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsOidc) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL param.Field[string] `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL param.Field[string] `json:"certs_url"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// OAuth scopes
	Scopes param.Field[[]string] `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccountAccessIdentityProviderUpdateParamsOidcConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderUpdateParamsOidcScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderUpdateParamsOkta struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsOktaConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsOktaScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsOkta) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsOktaConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your okta account url
	OktaAccount param.Field[string] `json:"okta_account"`
}

func (r AccountAccessIdentityProviderUpdateParamsOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderUpdateParamsOktaScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderUpdateParamsOnelogin struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsOneloginConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsOneloginScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsOnelogin) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsOneloginConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your OneLogin account url
	OneloginAccount param.Field[string] `json:"onelogin_account"`
}

func (r AccountAccessIdentityProviderUpdateParamsOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderUpdateParamsOneloginScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderUpdateParamsPingone struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsPingoneConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsPingoneScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsPingone) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsPingoneConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your PingOne environment identifier
	PingEnvID param.Field[string] `json:"ping_env_id"`
}

func (r AccountAccessIdentityProviderUpdateParamsPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderUpdateParamsPingoneScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderUpdateParamsSaml struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsSamlConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsSamlScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsSaml) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]AccountAccessIdentityProviderUpdateParamsSamlConfigHeaderAttribute] `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts param.Field[[]string] `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL param.Field[string] `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest param.Field[bool] `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL param.Field[string] `json:"sso_target_url"`
}

func (r AccountAccessIdentityProviderUpdateParamsSamlConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderUpdateParamsSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r AccountAccessIdentityProviderUpdateParamsSamlConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderUpdateParamsSamlScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderUpdateParamsYandex struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsYandexConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsYandexScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsYandex) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsYandexConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderUpdateParamsYandexConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderUpdateParamsYandexScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This interface is a union satisfied by one of the following:
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAzureAd],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsCentrify],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsFacebook],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGitHub],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogle],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleApps],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsLinkedin],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOidc],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOkta],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOnelogin],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPingone],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsSaml],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsYandex].
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams interface {
	ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams()
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAzureAd struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAzureAdConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAzureAdScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAzureAd) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAzureAdConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your Azure directory uuid
	DirectoryID param.Field[string] `json:"directory_id"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAzureAdConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAzureAdScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsCentrify struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsCentrifyConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsCentrifyScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsCentrify) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount param.Field[string] `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID param.Field[string] `json:"centrify_app_id"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsCentrifyScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsFacebook struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsFacebookConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsFacebookScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsFacebook) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsFacebookConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsFacebookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsFacebookScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGitHub struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGitHubConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGitHubScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGitHub) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGitHubConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGitHubScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogle struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogle) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleApps struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleAppsConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleAppsScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleApps) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain param.Field[string] `json:"apps_domain"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsGoogleAppsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsLinkedin struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsLinkedinConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsLinkedinScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsLinkedin) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsLinkedinConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsLinkedinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOidc struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOidcConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOidcScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOidc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOidc) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL param.Field[string] `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL param.Field[string] `json:"certs_url"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// OAuth scopes
	Scopes param.Field[[]string] `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOidcConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOidcScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOkta struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOktaConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOktaScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOkta) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOktaConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your okta account url
	OktaAccount param.Field[string] `json:"okta_account"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOktaScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOnelogin struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOneloginConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOneloginScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOnelogin) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOneloginConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your OneLogin account url
	OneloginAccount param.Field[string] `json:"onelogin_account"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsOneloginScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPingone struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPingoneConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPingoneScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPingone) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPingoneConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your PingOne environment identifier
	PingEnvID param.Field[string] `json:"ping_env_id"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPingoneScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsSaml struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsSamlConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsSamlScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsSaml) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsSamlConfigHeaderAttribute] `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts param.Field[[]string] `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL param.Field[string] `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest param.Field[bool] `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL param.Field[string] `json:"sso_target_url"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsSamlConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsSamlConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsSamlScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsYandex struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsYandexConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsYandexScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsYandex) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsYandexConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsYandexConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision param.Field[bool] `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret param.Field[string] `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsYandexScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

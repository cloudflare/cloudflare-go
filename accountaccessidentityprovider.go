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
func (r *AccountAccessIdentityProviderService) Get(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccountAccessIdentityProviderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured identity provider.
func (r *AccountAccessIdentityProviderService) Update(ctx context.Context, identifier string, uuid string, body AccountAccessIdentityProviderUpdateParams, opts ...option.RequestOption) (res *AccountAccessIdentityProviderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an identity provider from Access.
func (r *AccountAccessIdentityProviderService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccountAccessIdentityProviderDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new identity provider to Access.
func (r *AccountAccessIdentityProviderService) AccessIdentityProvidersAddAnAccessIdentityProvider(ctx context.Context, identifier string, body AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams, opts ...option.RequestOption) (res *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/identity_providers", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all configured identity providers.
func (r *AccountAccessIdentityProviderService) AccessIdentityProvidersListAccessIdentityProviders(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/identity_providers", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAccessIdentityProviderGetResponse struct {
	Errors   []AccountAccessIdentityProviderGetResponseError   `json:"errors"`
	Messages []AccountAccessIdentityProviderGetResponseMessage `json:"messages"`
	Result   AccountAccessIdentityProviderGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessIdentityProviderGetResponseSuccess `json:"success"`
	JSON    accountAccessIdentityProviderGetResponseJSON    `json:"-"`
}

// accountAccessIdentityProviderGetResponseJSON contains the JSON metadata for the
// struct [AccountAccessIdentityProviderGetResponse]
type accountAccessIdentityProviderGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderGetResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountAccessIdentityProviderGetResponseErrorJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseErrorJSON contains the JSON metadata for
// the struct [AccountAccessIdentityProviderGetResponseError]
type accountAccessIdentityProviderGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderGetResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountAccessIdentityProviderGetResponseMessageJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseMessageJSON contains the JSON metadata
// for the struct [AccountAccessIdentityProviderGetResponseMessage]
type accountAccessIdentityProviderGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAd],
// [AccountAccessIdentityProviderGetResponseResultPajwohLqCentrify],
// [AccountAccessIdentityProviderGetResponseResultPajwohLqFacebook],
// [AccountAccessIdentityProviderGetResponseResultPajwohLqGitHub],
// [AccountAccessIdentityProviderGetResponseResultPajwohLqGoogle],
// [AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleApps],
// [AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedin],
// [AccountAccessIdentityProviderGetResponseResultPajwohLqOidc],
// [AccountAccessIdentityProviderGetResponseResultPajwohLqOkta],
// [AccountAccessIdentityProviderGetResponseResultPajwohLqOnelogin],
// [AccountAccessIdentityProviderGetResponseResultPajwohLqPingone],
// [AccountAccessIdentityProviderGetResponseResultPajwohLqSaml],
// [AccountAccessIdentityProviderGetResponseResultPajwohLqYandex] or
// [AccountAccessIdentityProviderGetResponseResultPajwohLqOnetimepin].
type AccountAccessIdentityProviderGetResponseResult interface {
	implementsAccountAccessIdentityProviderGetResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessIdentityProviderGetResponseResult)(nil)).Elem(), "")
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAd struct {
	// UUID
	ID     string                                                              `json:"id"`
	Config AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType `json:"type"`
	JSON accountAccessIdentityProviderGetResponseResultPajwohLqAzureAdJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqAzureAdJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAd]
type accountAccessIdentityProviderGetResponseResultPajwohLqAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAd) implementsAccountAccessIdentityProviderGetResponseResult() {
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Should Cloudflare try to load authentication contexts from your account
	ConditionalAccessEnabled bool `json:"conditional_access_enabled"`
	// Your Azure directory uuid
	DirectoryID string `json:"directory_id"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool                                                                    `json:"support_groups"`
	JSON          accountAccessIdentityProviderGetResponseResultPajwohLqAzureAdConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqAzureAdConfigJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqAzureAdConfigJSON struct {
	Claims                   apijson.Field
	ClientID                 apijson.Field
	ClientSecret             apijson.Field
	ConditionalAccessEnabled apijson.Field
	DirectoryID              apijson.Field
	EmailClaimName           apijson.Field
	SupportGroups            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdScimConfig struct {
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
	UserDeprovision bool                                                                        `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderGetResponseResultPajwohLqAzureAdScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqAzureAdScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdScimConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType string

const (
	AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdTypeOnetimepin AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType = "onetimepin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdTypeAzureAd    AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType = "azureAD"
	AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdTypeSaml       AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType = "saml"
	AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdTypeCentrify   AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType = "centrify"
	AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdTypeFacebook   AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType = "facebook"
	AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdTypeGitHub     AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType = "github"
	AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdTypeGoogleApps AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType = "google-apps"
	AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdTypeGoogle     AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType = "google"
	AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdTypeLinkedin   AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType = "linkedin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdTypeOidc       AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType = "oidc"
	AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdTypeOkta       AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType = "okta"
	AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdTypeOnelogin   AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType = "onelogin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdTypePingone    AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType = "pingone"
	AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdTypeYandex     AccountAccessIdentityProviderGetResponseResultPajwohLqAzureAdType = "yandex"
)

type AccountAccessIdentityProviderGetResponseResultPajwohLqCentrify struct {
	// UUID
	ID     string                                                               `json:"id"`
	Config AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType `json:"type"`
	JSON accountAccessIdentityProviderGetResponseResultPajwohLqCentrifyJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqCentrifyJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqCentrify]
type accountAccessIdentityProviderGetResponseResultPajwohLqCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderGetResponseResultPajwohLqCentrify) implementsAccountAccessIdentityProviderGetResponseResult() {
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount string `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID string `json:"centrify_app_id"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                   `json:"email_claim_name"`
	JSON           accountAccessIdentityProviderGetResponseResultPajwohLqCentrifyConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqCentrifyConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyScimConfig struct {
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
	UserDeprovision bool                                                                         `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderGetResponseResultPajwohLqCentrifyScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqCentrifyScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyScimConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType string

const (
	AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyTypeOnetimepin AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType = "onetimepin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyTypeAzureAd    AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType = "azureAD"
	AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyTypeSaml       AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType = "saml"
	AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyTypeCentrify   AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType = "centrify"
	AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyTypeFacebook   AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType = "facebook"
	AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyTypeGitHub     AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType = "github"
	AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyTypeGoogleApps AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType = "google-apps"
	AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyTypeGoogle     AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType = "google"
	AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyTypeLinkedin   AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType = "linkedin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyTypeOidc       AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType = "oidc"
	AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyTypeOkta       AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType = "okta"
	AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyTypeOnelogin   AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType = "onelogin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyTypePingone    AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType = "pingone"
	AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyTypeYandex     AccountAccessIdentityProviderGetResponseResultPajwohLqCentrifyType = "yandex"
)

type AccountAccessIdentityProviderGetResponseResultPajwohLqFacebook struct {
	// UUID
	ID     string                                                               `json:"id"`
	Config AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType `json:"type"`
	JSON accountAccessIdentityProviderGetResponseResultPajwohLqFacebookJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqFacebookJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqFacebook]
type accountAccessIdentityProviderGetResponseResultPajwohLqFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderGetResponseResultPajwohLqFacebook) implementsAccountAccessIdentityProviderGetResponseResult() {
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                   `json:"client_secret"`
	JSON         accountAccessIdentityProviderGetResponseResultPajwohLqFacebookConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqFacebookConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookScimConfig struct {
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
	UserDeprovision bool                                                                         `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderGetResponseResultPajwohLqFacebookScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqFacebookScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookScimConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType string

const (
	AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookTypeOnetimepin AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType = "onetimepin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookTypeAzureAd    AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType = "azureAD"
	AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookTypeSaml       AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType = "saml"
	AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookTypeCentrify   AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType = "centrify"
	AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookTypeFacebook   AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType = "facebook"
	AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookTypeGitHub     AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType = "github"
	AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookTypeGoogleApps AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType = "google-apps"
	AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookTypeGoogle     AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType = "google"
	AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookTypeLinkedin   AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType = "linkedin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookTypeOidc       AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType = "oidc"
	AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookTypeOkta       AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType = "okta"
	AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookTypeOnelogin   AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType = "onelogin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookTypePingone    AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType = "pingone"
	AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookTypeYandex     AccountAccessIdentityProviderGetResponseResultPajwohLqFacebookType = "yandex"
)

type AccountAccessIdentityProviderGetResponseResultPajwohLqGitHub struct {
	// UUID
	ID     string                                                             `json:"id"`
	Config AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType `json:"type"`
	JSON accountAccessIdentityProviderGetResponseResultPajwohLqGitHubJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqGitHubJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqGitHub]
type accountAccessIdentityProviderGetResponseResultPajwohLqGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderGetResponseResultPajwohLqGitHub) implementsAccountAccessIdentityProviderGetResponseResult() {
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                 `json:"client_secret"`
	JSON         accountAccessIdentityProviderGetResponseResultPajwohLqGitHubConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqGitHubConfigJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubScimConfig struct {
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
	UserDeprovision bool                                                                       `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderGetResponseResultPajwohLqGitHubScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqGitHubScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubScimConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType string

const (
	AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubTypeOnetimepin AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType = "onetimepin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubTypeAzureAd    AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType = "azureAD"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubTypeSaml       AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType = "saml"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubTypeCentrify   AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType = "centrify"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubTypeFacebook   AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType = "facebook"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubTypeGitHub     AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType = "github"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubTypeGoogleApps AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType = "google-apps"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubTypeGoogle     AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType = "google"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubTypeLinkedin   AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType = "linkedin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubTypeOidc       AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType = "oidc"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubTypeOkta       AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType = "okta"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubTypeOnelogin   AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType = "onelogin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubTypePingone    AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType = "pingone"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubTypeYandex     AccountAccessIdentityProviderGetResponseResultPajwohLqGitHubType = "yandex"
)

type AccountAccessIdentityProviderGetResponseResultPajwohLqGoogle struct {
	// UUID
	ID     string                                                             `json:"id"`
	Config AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType `json:"type"`
	JSON accountAccessIdentityProviderGetResponseResultPajwohLqGoogleJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqGoogleJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqGoogle]
type accountAccessIdentityProviderGetResponseResultPajwohLqGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderGetResponseResultPajwohLqGoogle) implementsAccountAccessIdentityProviderGetResponseResult() {
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                 `json:"email_claim_name"`
	JSON           accountAccessIdentityProviderGetResponseResultPajwohLqGoogleConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqGoogleConfigJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleScimConfig struct {
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
	UserDeprovision bool                                                                       `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderGetResponseResultPajwohLqGoogleScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqGoogleScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleScimConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType string

const (
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleTypeOnetimepin AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType = "onetimepin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleTypeAzureAd    AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType = "azureAD"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleTypeSaml       AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType = "saml"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleTypeCentrify   AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType = "centrify"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleTypeFacebook   AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType = "facebook"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleTypeGitHub     AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType = "github"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleTypeGoogleApps AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType = "google-apps"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleTypeGoogle     AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType = "google"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleTypeLinkedin   AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType = "linkedin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleTypeOidc       AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType = "oidc"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleTypeOkta       AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType = "okta"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleTypeOnelogin   AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType = "onelogin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleTypePingone    AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType = "pingone"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleTypeYandex     AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleType = "yandex"
)

type AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleApps struct {
	// UUID
	ID     string                                                                 `json:"id"`
	Config AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType `json:"type"`
	JSON accountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleApps]
type accountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleApps) implementsAccountAccessIdentityProviderGetResponseResult() {
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                     `json:"email_claim_name"`
	JSON           accountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsScimConfig struct {
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
	UserDeprovision bool                                                                           `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsScimConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType string

const (
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsTypeOnetimepin AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType = "onetimepin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsTypeAzureAd    AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType = "azureAD"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsTypeSaml       AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType = "saml"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsTypeCentrify   AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType = "centrify"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsTypeFacebook   AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType = "facebook"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsTypeGitHub     AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType = "github"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsTypeGoogleApps AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType = "google-apps"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsTypeGoogle     AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType = "google"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsTypeLinkedin   AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType = "linkedin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsTypeOidc       AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType = "oidc"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsTypeOkta       AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType = "okta"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsTypeOnelogin   AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType = "onelogin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsTypePingone    AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType = "pingone"
	AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsTypeYandex     AccountAccessIdentityProviderGetResponseResultPajwohLqGoogleAppsType = "yandex"
)

type AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedin struct {
	// UUID
	ID     string                                                               `json:"id"`
	Config AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType `json:"type"`
	JSON accountAccessIdentityProviderGetResponseResultPajwohLqLinkedinJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqLinkedinJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedin]
type accountAccessIdentityProviderGetResponseResultPajwohLqLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedin) implementsAccountAccessIdentityProviderGetResponseResult() {
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                   `json:"client_secret"`
	JSON         accountAccessIdentityProviderGetResponseResultPajwohLqLinkedinConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqLinkedinConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinScimConfig struct {
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
	UserDeprovision bool                                                                         `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderGetResponseResultPajwohLqLinkedinScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqLinkedinScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinScimConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType string

const (
	AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinTypeOnetimepin AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType = "onetimepin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinTypeAzureAd    AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType = "azureAD"
	AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinTypeSaml       AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType = "saml"
	AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinTypeCentrify   AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType = "centrify"
	AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinTypeFacebook   AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType = "facebook"
	AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinTypeGitHub     AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType = "github"
	AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinTypeGoogleApps AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType = "google-apps"
	AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinTypeGoogle     AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType = "google"
	AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinTypeLinkedin   AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType = "linkedin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinTypeOidc       AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType = "oidc"
	AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinTypeOkta       AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType = "okta"
	AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinTypeOnelogin   AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType = "onelogin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinTypePingone    AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType = "pingone"
	AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinTypeYandex     AccountAccessIdentityProviderGetResponseResultPajwohLqLinkedinType = "yandex"
)

type AccountAccessIdentityProviderGetResponseResultPajwohLqOidc struct {
	// UUID
	ID     string                                                           `json:"id"`
	Config AccountAccessIdentityProviderGetResponseResultPajwohLqOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderGetResponseResultPajwohLqOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType `json:"type"`
	JSON accountAccessIdentityProviderGetResponseResultPajwohLqOidcJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqOidcJSON contains the JSON
// metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqOidc]
type accountAccessIdentityProviderGetResponseResultPajwohLqOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderGetResponseResultPajwohLqOidc) implementsAccountAccessIdentityProviderGetResponseResult() {
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL string `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL string `json:"certs_url"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// OAuth scopes
	Scopes []string `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL string                                                               `json:"token_url"`
	JSON     accountAccessIdentityProviderGetResponseResultPajwohLqOidcConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqOidcConfigJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqOidcConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqOidcConfigJSON struct {
	AuthURL        apijson.Field
	CertsURL       apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	Scopes         apijson.Field
	TokenURL       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderGetResponseResultPajwohLqOidcScimConfig struct {
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
	UserDeprovision bool                                                                     `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderGetResponseResultPajwohLqOidcScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqOidcScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqOidcScimConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType string

const (
	AccountAccessIdentityProviderGetResponseResultPajwohLqOidcTypeOnetimepin AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType = "onetimepin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOidcTypeAzureAd    AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType = "azureAD"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOidcTypeSaml       AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType = "saml"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOidcTypeCentrify   AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType = "centrify"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOidcTypeFacebook   AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType = "facebook"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOidcTypeGitHub     AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType = "github"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOidcTypeGoogleApps AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType = "google-apps"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOidcTypeGoogle     AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType = "google"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOidcTypeLinkedin   AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType = "linkedin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOidcTypeOidc       AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType = "oidc"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOidcTypeOkta       AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType = "okta"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOidcTypeOnelogin   AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType = "onelogin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOidcTypePingone    AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType = "pingone"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOidcTypeYandex     AccountAccessIdentityProviderGetResponseResultPajwohLqOidcType = "yandex"
)

type AccountAccessIdentityProviderGetResponseResultPajwohLqOkta struct {
	// UUID
	ID     string                                                           `json:"id"`
	Config AccountAccessIdentityProviderGetResponseResultPajwohLqOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderGetResponseResultPajwohLqOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType `json:"type"`
	JSON accountAccessIdentityProviderGetResponseResultPajwohLqOktaJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqOktaJSON contains the JSON
// metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqOkta]
type accountAccessIdentityProviderGetResponseResultPajwohLqOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderGetResponseResultPajwohLqOkta) implementsAccountAccessIdentityProviderGetResponseResult() {
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqOktaConfig struct {
	// Your okta authorization server id
	AuthorizationServerID string `json:"authorization_server_id"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your okta account url
	OktaAccount string                                                               `json:"okta_account"`
	JSON        accountAccessIdentityProviderGetResponseResultPajwohLqOktaConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqOktaConfigJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqOktaConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderGetResponseResultPajwohLqOktaScimConfig struct {
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
	UserDeprovision bool                                                                     `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderGetResponseResultPajwohLqOktaScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqOktaScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqOktaScimConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType string

const (
	AccountAccessIdentityProviderGetResponseResultPajwohLqOktaTypeOnetimepin AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType = "onetimepin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOktaTypeAzureAd    AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType = "azureAD"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOktaTypeSaml       AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType = "saml"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOktaTypeCentrify   AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType = "centrify"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOktaTypeFacebook   AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType = "facebook"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOktaTypeGitHub     AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType = "github"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOktaTypeGoogleApps AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType = "google-apps"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOktaTypeGoogle     AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType = "google"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOktaTypeLinkedin   AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType = "linkedin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOktaTypeOidc       AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType = "oidc"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOktaTypeOkta       AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType = "okta"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOktaTypeOnelogin   AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType = "onelogin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOktaTypePingone    AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType = "pingone"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOktaTypeYandex     AccountAccessIdentityProviderGetResponseResultPajwohLqOktaType = "yandex"
)

type AccountAccessIdentityProviderGetResponseResultPajwohLqOnelogin struct {
	// UUID
	ID     string                                                               `json:"id"`
	Config AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType `json:"type"`
	JSON accountAccessIdentityProviderGetResponseResultPajwohLqOneloginJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqOneloginJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqOnelogin]
type accountAccessIdentityProviderGetResponseResultPajwohLqOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderGetResponseResultPajwohLqOnelogin) implementsAccountAccessIdentityProviderGetResponseResult() {
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                                   `json:"onelogin_account"`
	JSON            accountAccessIdentityProviderGetResponseResultPajwohLqOneloginConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqOneloginConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginScimConfig struct {
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
	UserDeprovision bool                                                                         `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderGetResponseResultPajwohLqOneloginScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqOneloginScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginScimConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType string

const (
	AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginTypeOnetimepin AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType = "onetimepin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginTypeAzureAd    AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType = "azureAD"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginTypeSaml       AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType = "saml"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginTypeCentrify   AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType = "centrify"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginTypeFacebook   AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType = "facebook"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginTypeGitHub     AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType = "github"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginTypeGoogleApps AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType = "google-apps"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginTypeGoogle     AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType = "google"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginTypeLinkedin   AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType = "linkedin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginTypeOidc       AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType = "oidc"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginTypeOkta       AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType = "okta"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginTypeOnelogin   AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType = "onelogin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginTypePingone    AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType = "pingone"
	AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginTypeYandex     AccountAccessIdentityProviderGetResponseResultPajwohLqOneloginType = "yandex"
)

type AccountAccessIdentityProviderGetResponseResultPajwohLqPingone struct {
	// UUID
	ID     string                                                              `json:"id"`
	Config AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType `json:"type"`
	JSON accountAccessIdentityProviderGetResponseResultPajwohLqPingoneJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqPingoneJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqPingone]
type accountAccessIdentityProviderGetResponseResultPajwohLqPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderGetResponseResultPajwohLqPingone) implementsAccountAccessIdentityProviderGetResponseResult() {
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                                  `json:"ping_env_id"`
	JSON      accountAccessIdentityProviderGetResponseResultPajwohLqPingoneConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqPingoneConfigJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneScimConfig struct {
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
	UserDeprovision bool                                                                        `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderGetResponseResultPajwohLqPingoneScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqPingoneScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneScimConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType string

const (
	AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneTypeOnetimepin AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType = "onetimepin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneTypeAzureAd    AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType = "azureAD"
	AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneTypeSaml       AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType = "saml"
	AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneTypeCentrify   AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType = "centrify"
	AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneTypeFacebook   AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType = "facebook"
	AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneTypeGitHub     AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType = "github"
	AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneTypeGoogleApps AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType = "google-apps"
	AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneTypeGoogle     AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType = "google"
	AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneTypeLinkedin   AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType = "linkedin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneTypeOidc       AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType = "oidc"
	AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneTypeOkta       AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType = "okta"
	AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneTypeOnelogin   AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType = "onelogin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneTypePingone    AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType = "pingone"
	AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneTypeYandex     AccountAccessIdentityProviderGetResponseResultPajwohLqPingoneType = "yandex"
)

type AccountAccessIdentityProviderGetResponseResultPajwohLqSaml struct {
	// UUID
	ID     string                                                           `json:"id"`
	Config AccountAccessIdentityProviderGetResponseResultPajwohLqSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderGetResponseResultPajwohLqSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType `json:"type"`
	JSON accountAccessIdentityProviderGetResponseResultPajwohLqSamlJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqSamlJSON contains the JSON
// metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqSaml]
type accountAccessIdentityProviderGetResponseResultPajwohLqSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderGetResponseResultPajwohLqSaml) implementsAccountAccessIdentityProviderGetResponseResult() {
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccountAccessIdentityProviderGetResponseResultPajwohLqSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                               `json:"sso_target_url"`
	JSON         accountAccessIdentityProviderGetResponseResultPajwohLqSamlConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqSamlConfigJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqSamlConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqSamlConfigJSON struct {
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

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                              `json:"header_name"`
	JSON       accountAccessIdentityProviderGetResponseResultPajwohLqSamlConfigHeaderAttributeJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqSamlConfigHeaderAttribute]
type accountAccessIdentityProviderGetResponseResultPajwohLqSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderGetResponseResultPajwohLqSamlScimConfig struct {
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
	UserDeprovision bool                                                                     `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderGetResponseResultPajwohLqSamlScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqSamlScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqSamlScimConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType string

const (
	AccountAccessIdentityProviderGetResponseResultPajwohLqSamlTypeOnetimepin AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType = "onetimepin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqSamlTypeAzureAd    AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType = "azureAD"
	AccountAccessIdentityProviderGetResponseResultPajwohLqSamlTypeSaml       AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType = "saml"
	AccountAccessIdentityProviderGetResponseResultPajwohLqSamlTypeCentrify   AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType = "centrify"
	AccountAccessIdentityProviderGetResponseResultPajwohLqSamlTypeFacebook   AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType = "facebook"
	AccountAccessIdentityProviderGetResponseResultPajwohLqSamlTypeGitHub     AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType = "github"
	AccountAccessIdentityProviderGetResponseResultPajwohLqSamlTypeGoogleApps AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType = "google-apps"
	AccountAccessIdentityProviderGetResponseResultPajwohLqSamlTypeGoogle     AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType = "google"
	AccountAccessIdentityProviderGetResponseResultPajwohLqSamlTypeLinkedin   AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType = "linkedin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqSamlTypeOidc       AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType = "oidc"
	AccountAccessIdentityProviderGetResponseResultPajwohLqSamlTypeOkta       AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType = "okta"
	AccountAccessIdentityProviderGetResponseResultPajwohLqSamlTypeOnelogin   AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType = "onelogin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqSamlTypePingone    AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType = "pingone"
	AccountAccessIdentityProviderGetResponseResultPajwohLqSamlTypeYandex     AccountAccessIdentityProviderGetResponseResultPajwohLqSamlType = "yandex"
)

type AccountAccessIdentityProviderGetResponseResultPajwohLqYandex struct {
	// UUID
	ID     string                                                             `json:"id"`
	Config AccountAccessIdentityProviderGetResponseResultPajwohLqYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderGetResponseResultPajwohLqYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType `json:"type"`
	JSON accountAccessIdentityProviderGetResponseResultPajwohLqYandexJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqYandexJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqYandex]
type accountAccessIdentityProviderGetResponseResultPajwohLqYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderGetResponseResultPajwohLqYandex) implementsAccountAccessIdentityProviderGetResponseResult() {
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                 `json:"client_secret"`
	JSON         accountAccessIdentityProviderGetResponseResultPajwohLqYandexConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqYandexConfigJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqYandexConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderGetResponseResultPajwohLqYandexScimConfig struct {
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
	UserDeprovision bool                                                                       `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderGetResponseResultPajwohLqYandexScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqYandexScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqYandexScimConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType string

const (
	AccountAccessIdentityProviderGetResponseResultPajwohLqYandexTypeOnetimepin AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType = "onetimepin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqYandexTypeAzureAd    AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType = "azureAD"
	AccountAccessIdentityProviderGetResponseResultPajwohLqYandexTypeSaml       AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType = "saml"
	AccountAccessIdentityProviderGetResponseResultPajwohLqYandexTypeCentrify   AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType = "centrify"
	AccountAccessIdentityProviderGetResponseResultPajwohLqYandexTypeFacebook   AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType = "facebook"
	AccountAccessIdentityProviderGetResponseResultPajwohLqYandexTypeGitHub     AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType = "github"
	AccountAccessIdentityProviderGetResponseResultPajwohLqYandexTypeGoogleApps AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType = "google-apps"
	AccountAccessIdentityProviderGetResponseResultPajwohLqYandexTypeGoogle     AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType = "google"
	AccountAccessIdentityProviderGetResponseResultPajwohLqYandexTypeLinkedin   AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType = "linkedin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqYandexTypeOidc       AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType = "oidc"
	AccountAccessIdentityProviderGetResponseResultPajwohLqYandexTypeOkta       AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType = "okta"
	AccountAccessIdentityProviderGetResponseResultPajwohLqYandexTypeOnelogin   AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType = "onelogin"
	AccountAccessIdentityProviderGetResponseResultPajwohLqYandexTypePingone    AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType = "pingone"
	AccountAccessIdentityProviderGetResponseResultPajwohLqYandexTypeYandex     AccountAccessIdentityProviderGetResponseResultPajwohLqYandexType = "yandex"
)

type AccountAccessIdentityProviderGetResponseResultPajwohLqOnetimepin struct {
	// UUID
	ID     string      `json:"id"`
	Config interface{} `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderGetResponseResultPajwohLqOnetimepinScimConfig `json:"scim_config"`
	Type       AccountAccessIdentityProviderGetResponseResultPajwohLqOnetimepinType       `json:"type"`
	JSON       accountAccessIdentityProviderGetResponseResultPajwohLqOnetimepinJSON       `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqOnetimepinJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqOnetimepin]
type accountAccessIdentityProviderGetResponseResultPajwohLqOnetimepinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderGetResponseResultPajwohLqOnetimepin) implementsAccountAccessIdentityProviderGetResponseResult() {
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderGetResponseResultPajwohLqOnetimepinScimConfig struct {
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
	UserDeprovision bool                                                                           `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderGetResponseResultPajwohLqOnetimepinScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderGetResponseResultPajwohLqOnetimepinScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderGetResponseResultPajwohLqOnetimepinScimConfig]
type accountAccessIdentityProviderGetResponseResultPajwohLqOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderGetResponseResultPajwohLqOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderGetResponseResultPajwohLqOnetimepinType string

const (
	AccountAccessIdentityProviderGetResponseResultPajwohLqOnetimepinTypeOnetimepin AccountAccessIdentityProviderGetResponseResultPajwohLqOnetimepinType = "onetimepin"
)

// Whether the API call was successful
type AccountAccessIdentityProviderGetResponseSuccess bool

const (
	AccountAccessIdentityProviderGetResponseSuccessTrue AccountAccessIdentityProviderGetResponseSuccess = true
)

type AccountAccessIdentityProviderUpdateResponse struct {
	Errors   []AccountAccessIdentityProviderUpdateResponseError   `json:"errors"`
	Messages []AccountAccessIdentityProviderUpdateResponseMessage `json:"messages"`
	Result   AccountAccessIdentityProviderUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessIdentityProviderUpdateResponseSuccess `json:"success"`
	JSON    accountAccessIdentityProviderUpdateResponseJSON    `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseJSON contains the JSON metadata for
// the struct [AccountAccessIdentityProviderUpdateResponse]
type accountAccessIdentityProviderUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderUpdateResponseError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accountAccessIdentityProviderUpdateResponseErrorJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseErrorJSON contains the JSON metadata
// for the struct [AccountAccessIdentityProviderUpdateResponseError]
type accountAccessIdentityProviderUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderUpdateResponseMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountAccessIdentityProviderUpdateResponseMessageJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseMessageJSON contains the JSON
// metadata for the struct [AccountAccessIdentityProviderUpdateResponseMessage]
type accountAccessIdentityProviderUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAd],
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrify],
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebook],
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHub],
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogle],
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleApps],
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedin],
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidc],
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqOkta],
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnelogin],
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingone],
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqSaml],
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandex] or
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepin].
type AccountAccessIdentityProviderUpdateResponseResult interface {
	implementsAccountAccessIdentityProviderUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessIdentityProviderUpdateResponseResult)(nil)).Elem(), "")
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAd struct {
	// UUID
	ID     string                                                                 `json:"id"`
	Config AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType `json:"type"`
	JSON accountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAd]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAd) implementsAccountAccessIdentityProviderUpdateResponseResult() {
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Should Cloudflare try to load authentication contexts from your account
	ConditionalAccessEnabled bool `json:"conditional_access_enabled"`
	// Your Azure directory uuid
	DirectoryID string `json:"directory_id"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool                                                                       `json:"support_groups"`
	JSON          accountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdConfigJSON struct {
	Claims                   apijson.Field
	ClientID                 apijson.Field
	ClientSecret             apijson.Field
	ConditionalAccessEnabled apijson.Field
	DirectoryID              apijson.Field
	EmailClaimName           apijson.Field
	SupportGroups            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdScimConfig struct {
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
	UserDeprovision bool                                                                           `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdScimConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType string

const (
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdTypeOnetimepin AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType = "onetimepin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdTypeAzureAd    AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType = "azureAD"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdTypeSaml       AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType = "saml"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdTypeCentrify   AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType = "centrify"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdTypeFacebook   AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType = "facebook"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdTypeGitHub     AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType = "github"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdTypeGoogleApps AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType = "google-apps"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdTypeGoogle     AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType = "google"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdTypeLinkedin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType = "linkedin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdTypeOidc       AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType = "oidc"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdTypeOkta       AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType = "okta"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdTypeOnelogin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType = "onelogin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdTypePingone    AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType = "pingone"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdTypeYandex     AccountAccessIdentityProviderUpdateResponseResultPajwohLqAzureAdType = "yandex"
)

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrify struct {
	// UUID
	ID     string                                                                  `json:"id"`
	Config AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType `json:"type"`
	JSON accountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrify]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrify) implementsAccountAccessIdentityProviderUpdateResponseResult() {
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount string `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID string `json:"centrify_app_id"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                      `json:"email_claim_name"`
	JSON           accountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyScimConfig struct {
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
	UserDeprovision bool                                                                            `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyScimConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType string

const (
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyTypeOnetimepin AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType = "onetimepin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyTypeAzureAd    AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType = "azureAD"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyTypeSaml       AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType = "saml"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyTypeCentrify   AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType = "centrify"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyTypeFacebook   AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType = "facebook"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyTypeGitHub     AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType = "github"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyTypeGoogleApps AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType = "google-apps"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyTypeGoogle     AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType = "google"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyTypeLinkedin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType = "linkedin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyTypeOidc       AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType = "oidc"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyTypeOkta       AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType = "okta"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyTypeOnelogin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType = "onelogin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyTypePingone    AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType = "pingone"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyTypeYandex     AccountAccessIdentityProviderUpdateResponseResultPajwohLqCentrifyType = "yandex"
)

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebook struct {
	// UUID
	ID     string                                                                  `json:"id"`
	Config AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType `json:"type"`
	JSON accountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebook]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebook) implementsAccountAccessIdentityProviderUpdateResponseResult() {
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                      `json:"client_secret"`
	JSON         accountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookScimConfig struct {
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
	UserDeprovision bool                                                                            `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookScimConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType string

const (
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookTypeOnetimepin AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType = "onetimepin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookTypeAzureAd    AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType = "azureAD"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookTypeSaml       AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType = "saml"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookTypeCentrify   AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType = "centrify"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookTypeFacebook   AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType = "facebook"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookTypeGitHub     AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType = "github"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookTypeGoogleApps AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType = "google-apps"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookTypeGoogle     AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType = "google"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookTypeLinkedin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType = "linkedin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookTypeOidc       AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType = "oidc"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookTypeOkta       AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType = "okta"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookTypeOnelogin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType = "onelogin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookTypePingone    AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType = "pingone"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookTypeYandex     AccountAccessIdentityProviderUpdateResponseResultPajwohLqFacebookType = "yandex"
)

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHub struct {
	// UUID
	ID     string                                                                `json:"id"`
	Config AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType `json:"type"`
	JSON accountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHub]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHub) implementsAccountAccessIdentityProviderUpdateResponseResult() {
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                    `json:"client_secret"`
	JSON         accountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubScimConfig struct {
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
	UserDeprovision bool                                                                          `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubScimConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType string

const (
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubTypeOnetimepin AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType = "onetimepin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubTypeAzureAd    AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType = "azureAD"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubTypeSaml       AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType = "saml"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubTypeCentrify   AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType = "centrify"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubTypeFacebook   AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType = "facebook"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubTypeGitHub     AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType = "github"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubTypeGoogleApps AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType = "google-apps"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubTypeGoogle     AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType = "google"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubTypeLinkedin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType = "linkedin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubTypeOidc       AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType = "oidc"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubTypeOkta       AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType = "okta"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubTypeOnelogin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType = "onelogin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubTypePingone    AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType = "pingone"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubTypeYandex     AccountAccessIdentityProviderUpdateResponseResultPajwohLqGitHubType = "yandex"
)

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogle struct {
	// UUID
	ID     string                                                                `json:"id"`
	Config AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType `json:"type"`
	JSON accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogle]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogle) implementsAccountAccessIdentityProviderUpdateResponseResult() {
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                    `json:"email_claim_name"`
	JSON           accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleScimConfig struct {
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
	UserDeprovision bool                                                                          `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleScimConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType string

const (
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleTypeOnetimepin AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType = "onetimepin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleTypeAzureAd    AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType = "azureAD"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleTypeSaml       AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType = "saml"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleTypeCentrify   AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType = "centrify"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleTypeFacebook   AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType = "facebook"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleTypeGitHub     AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType = "github"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleTypeGoogleApps AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType = "google-apps"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleTypeGoogle     AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType = "google"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleTypeLinkedin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType = "linkedin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleTypeOidc       AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType = "oidc"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleTypeOkta       AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType = "okta"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleTypeOnelogin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType = "onelogin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleTypePingone    AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType = "pingone"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleTypeYandex     AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleType = "yandex"
)

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleApps struct {
	// UUID
	ID     string                                                                    `json:"id"`
	Config AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType `json:"type"`
	JSON accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleApps]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleApps) implementsAccountAccessIdentityProviderUpdateResponseResult() {
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                        `json:"email_claim_name"`
	JSON           accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsScimConfig struct {
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
	UserDeprovision bool                                                                              `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsScimConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType string

const (
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsTypeOnetimepin AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType = "onetimepin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsTypeAzureAd    AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType = "azureAD"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsTypeSaml       AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType = "saml"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsTypeCentrify   AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType = "centrify"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsTypeFacebook   AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType = "facebook"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsTypeGitHub     AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType = "github"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsTypeGoogleApps AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType = "google-apps"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsTypeGoogle     AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType = "google"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsTypeLinkedin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType = "linkedin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsTypeOidc       AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType = "oidc"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsTypeOkta       AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType = "okta"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsTypeOnelogin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType = "onelogin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsTypePingone    AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType = "pingone"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsTypeYandex     AccountAccessIdentityProviderUpdateResponseResultPajwohLqGoogleAppsType = "yandex"
)

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedin struct {
	// UUID
	ID     string                                                                  `json:"id"`
	Config AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType `json:"type"`
	JSON accountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedin]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedin) implementsAccountAccessIdentityProviderUpdateResponseResult() {
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                      `json:"client_secret"`
	JSON         accountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinScimConfig struct {
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
	UserDeprovision bool                                                                            `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinScimConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType string

const (
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinTypeOnetimepin AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType = "onetimepin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinTypeAzureAd    AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType = "azureAD"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinTypeSaml       AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType = "saml"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinTypeCentrify   AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType = "centrify"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinTypeFacebook   AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType = "facebook"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinTypeGitHub     AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType = "github"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinTypeGoogleApps AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType = "google-apps"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinTypeGoogle     AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType = "google"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinTypeLinkedin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType = "linkedin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinTypeOidc       AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType = "oidc"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinTypeOkta       AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType = "okta"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinTypeOnelogin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType = "onelogin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinTypePingone    AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType = "pingone"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinTypeYandex     AccountAccessIdentityProviderUpdateResponseResultPajwohLqLinkedinType = "yandex"
)

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidc struct {
	// UUID
	ID     string                                                              `json:"id"`
	Config AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType `json:"type"`
	JSON accountAccessIdentityProviderUpdateResponseResultPajwohLqOidcJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqOidcJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidc]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidc) implementsAccountAccessIdentityProviderUpdateResponseResult() {
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL string `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL string `json:"certs_url"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// OAuth scopes
	Scopes []string `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL string                                                                  `json:"token_url"`
	JSON     accountAccessIdentityProviderUpdateResponseResultPajwohLqOidcConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqOidcConfigJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqOidcConfigJSON struct {
	AuthURL        apijson.Field
	CertsURL       apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	Scopes         apijson.Field
	TokenURL       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcScimConfig struct {
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
	UserDeprovision bool                                                                        `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderUpdateResponseResultPajwohLqOidcScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqOidcScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcScimConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType string

const (
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcTypeOnetimepin AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType = "onetimepin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcTypeAzureAd    AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType = "azureAD"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcTypeSaml       AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType = "saml"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcTypeCentrify   AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType = "centrify"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcTypeFacebook   AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType = "facebook"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcTypeGitHub     AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType = "github"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcTypeGoogleApps AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType = "google-apps"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcTypeGoogle     AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType = "google"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcTypeLinkedin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType = "linkedin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcTypeOidc       AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType = "oidc"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcTypeOkta       AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType = "okta"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcTypeOnelogin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType = "onelogin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcTypePingone    AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType = "pingone"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcTypeYandex     AccountAccessIdentityProviderUpdateResponseResultPajwohLqOidcType = "yandex"
)

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOkta struct {
	// UUID
	ID     string                                                              `json:"id"`
	Config AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType `json:"type"`
	JSON accountAccessIdentityProviderUpdateResponseResultPajwohLqOktaJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqOktaJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqOkta]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderUpdateResponseResultPajwohLqOkta) implementsAccountAccessIdentityProviderUpdateResponseResult() {
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaConfig struct {
	// Your okta authorization server id
	AuthorizationServerID string `json:"authorization_server_id"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your okta account url
	OktaAccount string                                                                  `json:"okta_account"`
	JSON        accountAccessIdentityProviderUpdateResponseResultPajwohLqOktaConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqOktaConfigJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaScimConfig struct {
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
	UserDeprovision bool                                                                        `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderUpdateResponseResultPajwohLqOktaScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqOktaScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaScimConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType string

const (
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaTypeOnetimepin AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType = "onetimepin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaTypeAzureAd    AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType = "azureAD"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaTypeSaml       AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType = "saml"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaTypeCentrify   AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType = "centrify"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaTypeFacebook   AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType = "facebook"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaTypeGitHub     AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType = "github"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaTypeGoogleApps AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType = "google-apps"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaTypeGoogle     AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType = "google"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaTypeLinkedin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType = "linkedin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaTypeOidc       AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType = "oidc"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaTypeOkta       AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType = "okta"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaTypeOnelogin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType = "onelogin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaTypePingone    AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType = "pingone"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaTypeYandex     AccountAccessIdentityProviderUpdateResponseResultPajwohLqOktaType = "yandex"
)

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnelogin struct {
	// UUID
	ID     string                                                                  `json:"id"`
	Config AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType `json:"type"`
	JSON accountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnelogin]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnelogin) implementsAccountAccessIdentityProviderUpdateResponseResult() {
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                                      `json:"onelogin_account"`
	JSON            accountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginScimConfig struct {
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
	UserDeprovision bool                                                                            `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginScimConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType string

const (
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginTypeOnetimepin AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType = "onetimepin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginTypeAzureAd    AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType = "azureAD"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginTypeSaml       AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType = "saml"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginTypeCentrify   AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType = "centrify"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginTypeFacebook   AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType = "facebook"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginTypeGitHub     AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType = "github"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginTypeGoogleApps AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType = "google-apps"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginTypeGoogle     AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType = "google"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginTypeLinkedin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType = "linkedin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginTypeOidc       AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType = "oidc"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginTypeOkta       AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType = "okta"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginTypeOnelogin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType = "onelogin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginTypePingone    AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType = "pingone"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginTypeYandex     AccountAccessIdentityProviderUpdateResponseResultPajwohLqOneloginType = "yandex"
)

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingone struct {
	// UUID
	ID     string                                                                 `json:"id"`
	Config AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType `json:"type"`
	JSON accountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingone]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingone) implementsAccountAccessIdentityProviderUpdateResponseResult() {
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                                     `json:"ping_env_id"`
	JSON      accountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneScimConfig struct {
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
	UserDeprovision bool                                                                           `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneScimConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType string

const (
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneTypeOnetimepin AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType = "onetimepin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneTypeAzureAd    AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType = "azureAD"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneTypeSaml       AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType = "saml"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneTypeCentrify   AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType = "centrify"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneTypeFacebook   AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType = "facebook"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneTypeGitHub     AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType = "github"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneTypeGoogleApps AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType = "google-apps"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneTypeGoogle     AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType = "google"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneTypeLinkedin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType = "linkedin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneTypeOidc       AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType = "oidc"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneTypeOkta       AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType = "okta"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneTypeOnelogin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType = "onelogin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneTypePingone    AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType = "pingone"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneTypeYandex     AccountAccessIdentityProviderUpdateResponseResultPajwohLqPingoneType = "yandex"
)

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqSaml struct {
	// UUID
	ID     string                                                              `json:"id"`
	Config AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType `json:"type"`
	JSON accountAccessIdentityProviderUpdateResponseResultPajwohLqSamlJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqSamlJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqSaml]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderUpdateResponseResultPajwohLqSaml) implementsAccountAccessIdentityProviderUpdateResponseResult() {
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                                  `json:"sso_target_url"`
	JSON         accountAccessIdentityProviderUpdateResponseResultPajwohLqSamlConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqSamlConfigJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqSamlConfigJSON struct {
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

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                                 `json:"header_name"`
	JSON       accountAccessIdentityProviderUpdateResponseResultPajwohLqSamlConfigHeaderAttributeJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlConfigHeaderAttribute]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlScimConfig struct {
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
	UserDeprovision bool                                                                        `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderUpdateResponseResultPajwohLqSamlScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqSamlScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlScimConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType string

const (
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlTypeOnetimepin AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType = "onetimepin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlTypeAzureAd    AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType = "azureAD"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlTypeSaml       AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType = "saml"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlTypeCentrify   AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType = "centrify"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlTypeFacebook   AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType = "facebook"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlTypeGitHub     AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType = "github"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlTypeGoogleApps AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType = "google-apps"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlTypeGoogle     AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType = "google"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlTypeLinkedin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType = "linkedin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlTypeOidc       AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType = "oidc"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlTypeOkta       AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType = "okta"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlTypeOnelogin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType = "onelogin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlTypePingone    AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType = "pingone"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlTypeYandex     AccountAccessIdentityProviderUpdateResponseResultPajwohLqSamlType = "yandex"
)

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandex struct {
	// UUID
	ID     string                                                                `json:"id"`
	Config AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType `json:"type"`
	JSON accountAccessIdentityProviderUpdateResponseResultPajwohLqYandexJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqYandexJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandex]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandex) implementsAccountAccessIdentityProviderUpdateResponseResult() {
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                    `json:"client_secret"`
	JSON         accountAccessIdentityProviderUpdateResponseResultPajwohLqYandexConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqYandexConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexScimConfig struct {
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
	UserDeprovision bool                                                                          `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderUpdateResponseResultPajwohLqYandexScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqYandexScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexScimConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType string

const (
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexTypeOnetimepin AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType = "onetimepin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexTypeAzureAd    AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType = "azureAD"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexTypeSaml       AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType = "saml"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexTypeCentrify   AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType = "centrify"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexTypeFacebook   AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType = "facebook"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexTypeGitHub     AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType = "github"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexTypeGoogleApps AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType = "google-apps"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexTypeGoogle     AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType = "google"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexTypeLinkedin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType = "linkedin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexTypeOidc       AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType = "oidc"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexTypeOkta       AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType = "okta"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexTypeOnelogin   AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType = "onelogin"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexTypePingone    AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType = "pingone"
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexTypeYandex     AccountAccessIdentityProviderUpdateResponseResultPajwohLqYandexType = "yandex"
)

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepin struct {
	// UUID
	ID     string      `json:"id"`
	Config interface{} `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepinScimConfig `json:"scim_config"`
	Type       AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepinType       `json:"type"`
	JSON       accountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepinJSON       `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepinJSON contains
// the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepin]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepin) implementsAccountAccessIdentityProviderUpdateResponseResult() {
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepinScimConfig struct {
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
	UserDeprovision bool                                                                              `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepinScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepinScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepinScimConfig]
type accountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepinType string

const (
	AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepinTypeOnetimepin AccountAccessIdentityProviderUpdateResponseResultPajwohLqOnetimepinType = "onetimepin"
)

// Whether the API call was successful
type AccountAccessIdentityProviderUpdateResponseSuccess bool

const (
	AccountAccessIdentityProviderUpdateResponseSuccessTrue AccountAccessIdentityProviderUpdateResponseSuccess = true
)

type AccountAccessIdentityProviderDeleteResponse struct {
	Errors   []AccountAccessIdentityProviderDeleteResponseError   `json:"errors"`
	Messages []AccountAccessIdentityProviderDeleteResponseMessage `json:"messages"`
	Result   AccountAccessIdentityProviderDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessIdentityProviderDeleteResponseSuccess `json:"success"`
	JSON    accountAccessIdentityProviderDeleteResponseJSON    `json:"-"`
}

// accountAccessIdentityProviderDeleteResponseJSON contains the JSON metadata for
// the struct [AccountAccessIdentityProviderDeleteResponse]
type accountAccessIdentityProviderDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderDeleteResponseError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accountAccessIdentityProviderDeleteResponseErrorJSON `json:"-"`
}

// accountAccessIdentityProviderDeleteResponseErrorJSON contains the JSON metadata
// for the struct [AccountAccessIdentityProviderDeleteResponseError]
type accountAccessIdentityProviderDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderDeleteResponseMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountAccessIdentityProviderDeleteResponseMessageJSON `json:"-"`
}

// accountAccessIdentityProviderDeleteResponseMessageJSON contains the JSON
// metadata for the struct [AccountAccessIdentityProviderDeleteResponseMessage]
type accountAccessIdentityProviderDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderDeleteResponseResult struct {
	// UUID
	ID   string                                                `json:"id"`
	JSON accountAccessIdentityProviderDeleteResponseResultJSON `json:"-"`
}

// accountAccessIdentityProviderDeleteResponseResultJSON contains the JSON metadata
// for the struct [AccountAccessIdentityProviderDeleteResponseResult]
type accountAccessIdentityProviderDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessIdentityProviderDeleteResponseSuccess bool

const (
	AccountAccessIdentityProviderDeleteResponseSuccessTrue AccountAccessIdentityProviderDeleteResponseSuccess = true
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse struct {
	Errors   []AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseError   `json:"errors"`
	Messages []AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseMessage `json:"messages"`
	Result   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseSuccess `json:"success"`
	JSON    accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseJSON    `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseError struct {
	Code    int64                                                                                            `json:"code,required"`
	Message string                                                                                           `json:"message,required"`
	JSON    accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseErrorJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseError]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseMessage struct {
	Code    int64                                                                                              `json:"code,required"`
	Message string                                                                                             `json:"message,required"`
	JSON    accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseMessageJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseMessage]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAd],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrify],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebook],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHub],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogle],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleApps],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedin],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidc],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOkta],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnelogin],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingone],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSaml],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandex]
// or
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepin].
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult interface {
	implementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult)(nil)).Elem(), "")
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAd struct {
	// UUID
	ID     string                                                                                                             `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAd]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAd) implementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Should Cloudflare try to load authentication contexts from your account
	ConditionalAccessEnabled bool `json:"conditional_access_enabled"`
	// Your Azure directory uuid
	DirectoryID string `json:"directory_id"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool                                                                                                                   `json:"support_groups"`
	JSON          accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdConfigJSON struct {
	Claims                   apijson.Field
	ClientID                 apijson.Field
	ClientSecret             apijson.Field
	ConditionalAccessEnabled apijson.Field
	DirectoryID              apijson.Field
	EmailClaimName           apijson.Field
	SupportGroups            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdScimConfig struct {
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
	UserDeprovision bool                                                                                                                       `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqAzureAdType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrify struct {
	// UUID
	ID     string                                                                                                              `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrify]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrify) implementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount string `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID string `json:"centrify_app_id"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                                                                  `json:"email_claim_name"`
	JSON           accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyScimConfig struct {
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
	UserDeprovision bool                                                                                                                        `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqCentrifyType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebook struct {
	// UUID
	ID     string                                                                                                              `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebook]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebook) implementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                                  `json:"client_secret"`
	JSON         accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookScimConfig struct {
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
	UserDeprovision bool                                                                                                                        `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqFacebookType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHub struct {
	// UUID
	ID     string                                                                                                            `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHub]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHub) implementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                                `json:"client_secret"`
	JSON         accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubScimConfig struct {
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
	UserDeprovision bool                                                                                                                      `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGitHubType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogle struct {
	// UUID
	ID     string                                                                                                            `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogle]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogle) implementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                                                                `json:"email_claim_name"`
	JSON           accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleScimConfig struct {
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
	UserDeprovision bool                                                                                                                      `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleApps struct {
	// UUID
	ID     string                                                                                                                `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleApps]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleApps) implementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                                                                    `json:"email_claim_name"`
	JSON           accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsScimConfig struct {
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
	UserDeprovision bool                                                                                                                          `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqGoogleAppsType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedin struct {
	// UUID
	ID     string                                                                                                              `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedin]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedin) implementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                                  `json:"client_secret"`
	JSON         accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinScimConfig struct {
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
	UserDeprovision bool                                                                                                                        `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqLinkedinType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidc struct {
	// UUID
	ID     string                                                                                                          `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidc]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidc) implementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL string `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL string `json:"certs_url"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// OAuth scopes
	Scopes []string `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL string                                                                                                              `json:"token_url"`
	JSON     accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcConfigJSON struct {
	AuthURL        apijson.Field
	CertsURL       apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	Scopes         apijson.Field
	TokenURL       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcScimConfig struct {
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
	UserDeprovision bool                                                                                                                    `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOidcType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOkta struct {
	// UUID
	ID     string                                                                                                          `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOkta]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOkta) implementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaConfig struct {
	// Your okta authorization server id
	AuthorizationServerID string `json:"authorization_server_id"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your okta account url
	OktaAccount string                                                                                                              `json:"okta_account"`
	JSON        accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaScimConfig struct {
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
	UserDeprovision bool                                                                                                                    `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOktaType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnelogin struct {
	// UUID
	ID     string                                                                                                              `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnelogin]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnelogin) implementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                                                                                  `json:"onelogin_account"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginScimConfig struct {
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
	UserDeprovision bool                                                                                                                        `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOneloginType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingone struct {
	// UUID
	ID     string                                                                                                             `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingone]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingone) implementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                                                                                 `json:"ping_env_id"`
	JSON      accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneScimConfig struct {
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
	UserDeprovision bool                                                                                                                       `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqPingoneType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSaml struct {
	// UUID
	ID     string                                                                                                          `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSaml]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSaml) implementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                                                                              `json:"sso_target_url"`
	JSON         accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlConfigJSON struct {
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

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                                                                             `json:"header_name"`
	JSON       accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlConfigHeaderAttributeJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlConfigHeaderAttribute]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlScimConfig struct {
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
	UserDeprovision bool                                                                                                                    `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqSamlType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandex struct {
	// UUID
	ID     string                                                                                                            `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandex]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandex) implementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                                `json:"client_secret"`
	JSON         accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexScimConfig struct {
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
	UserDeprovision bool                                                                                                                      `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqYandexType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepin struct {
	// UUID
	ID     string      `json:"id"`
	Config interface{} `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepinScimConfig `json:"scim_config"`
	Type       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepinType       `json:"type"`
	JSON       accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepinJSON       `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepinJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepin]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepin) implementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepinScimConfig struct {
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
	UserDeprovision bool                                                                                                                          `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepinScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepinScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepinScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepinType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepinTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultPajwohLqOnetimepinType = "onetimepin"
)

// Whether the API call was successful
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseSuccess bool

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseSuccessTrue AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseSuccess = true
)

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse struct {
	Errors     []AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseError    `json:"errors"`
	Messages   []AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseMessage  `json:"messages"`
	Result     []AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult   `json:"result"`
	ResultInfo AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseSuccess `json:"success"`
	JSON    accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseJSON    `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseError struct {
	Code    int64                                                                                            `json:"code,required"`
	Message string                                                                                           `json:"message,required"`
	JSON    accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseErrorJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseError]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseMessage struct {
	Code    int64                                                                                              `json:"code,required"`
	Message string                                                                                             `json:"message,required"`
	JSON    accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseMessageJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseMessage]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAd],
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrify],
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebook],
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHub],
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogle],
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleApps],
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedin],
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidc],
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOkta],
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOnelogin],
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingone],
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSaml]
// or
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandex].
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult interface {
	implementsAccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult)(nil)).Elem(), "")
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAd struct {
	// UUID
	ID     string                                                                                                             `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAd]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAd) implementsAccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Should Cloudflare try to load authentication contexts from your account
	ConditionalAccessEnabled bool `json:"conditional_access_enabled"`
	// Your Azure directory uuid
	DirectoryID string `json:"directory_id"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool                                                                                                                   `json:"support_groups"`
	JSON          accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdConfigJSON struct {
	Claims                   apijson.Field
	ClientID                 apijson.Field
	ClientSecret             apijson.Field
	ConditionalAccessEnabled apijson.Field
	DirectoryID              apijson.Field
	EmailClaimName           apijson.Field
	SupportGroups            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdScimConfig struct {
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
	UserDeprovision bool                                                                                                                       `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqAzureAdType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrify struct {
	// UUID
	ID     string                                                                                                              `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrify]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrify) implementsAccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount string `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID string `json:"centrify_app_id"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                                                                  `json:"email_claim_name"`
	JSON           accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyScimConfig struct {
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
	UserDeprovision bool                                                                                                                        `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqCentrifyType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebook struct {
	// UUID
	ID     string                                                                                                              `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebook]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebook) implementsAccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                                  `json:"client_secret"`
	JSON         accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookScimConfig struct {
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
	UserDeprovision bool                                                                                                                        `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqFacebookType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHub struct {
	// UUID
	ID     string                                                                                                            `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHub]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHub) implementsAccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                                `json:"client_secret"`
	JSON         accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubScimConfig struct {
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
	UserDeprovision bool                                                                                                                      `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGitHubType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogle struct {
	// UUID
	ID     string                                                                                                            `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogle]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogle) implementsAccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                                                                `json:"email_claim_name"`
	JSON           accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleScimConfig struct {
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
	UserDeprovision bool                                                                                                                      `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleApps struct {
	// UUID
	ID     string                                                                                                                `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleApps]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleApps) implementsAccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                                                                    `json:"email_claim_name"`
	JSON           accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsScimConfig struct {
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
	UserDeprovision bool                                                                                                                          `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqGoogleAppsType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedin struct {
	// UUID
	ID     string                                                                                                              `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedin]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedin) implementsAccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                                  `json:"client_secret"`
	JSON         accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinScimConfig struct {
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
	UserDeprovision bool                                                                                                                        `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqLinkedinType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidc struct {
	// UUID
	ID     string                                                                                                          `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidc]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidc) implementsAccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL string `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL string `json:"certs_url"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// OAuth scopes
	Scopes []string `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL string                                                                                                              `json:"token_url"`
	JSON     accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcConfigJSON struct {
	AuthURL        apijson.Field
	CertsURL       apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	Scopes         apijson.Field
	TokenURL       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcScimConfig struct {
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
	UserDeprovision bool                                                                                                                    `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOidcType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOkta struct {
	// UUID
	ID     string                                                                                                          `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOkta]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOkta) implementsAccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaConfig struct {
	// Your okta authorization server id
	AuthorizationServerID string `json:"authorization_server_id"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your okta account url
	OktaAccount string                                                                                                              `json:"okta_account"`
	JSON        accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaScimConfig struct {
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
	UserDeprovision bool                                                                                                                    `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOktaType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOnelogin struct {
	// UUID
	ID     string                                                                                                              `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOnelogin]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOnelogin) implementsAccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                                                                                  `json:"onelogin_account"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginScimConfig struct {
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
	UserDeprovision bool                                                                                                                        `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqOneloginType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingone struct {
	// UUID
	ID     string                                                                                                             `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingone]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingone) implementsAccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                                                                                 `json:"ping_env_id"`
	JSON      accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneScimConfig struct {
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
	UserDeprovision bool                                                                                                                       `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqPingoneType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSaml struct {
	// UUID
	ID     string                                                                                                          `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSaml]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSaml) implementsAccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                                                                              `json:"sso_target_url"`
	JSON         accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlConfigJSON struct {
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

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                                                                             `json:"header_name"`
	JSON       accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlConfigHeaderAttributeJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlConfigHeaderAttribute]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlScimConfig struct {
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
	UserDeprovision bool                                                                                                                    `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqSamlType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandex struct {
	// UUID
	ID     string                                                                                                            `json:"id"`
	Config AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType `json:"type"`
	JSON accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandex]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandex) implementsAccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                                `json:"client_secret"`
	JSON         accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexScimConfig struct {
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
	UserDeprovision bool                                                                                                                      `json:"user_deprovision"`
	JSON            accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexScimConfigJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexScimConfigJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexScimConfig]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultPajwohLqYandexType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                               `json:"total_count"`
	JSON       accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultInfoJSON `json:"-"`
}

// accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultInfo]
type accountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseSuccess bool

const (
	AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseSuccessTrue AccountAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseSuccess = true
)

// This interface is a union satisfied by one of the following:
// [AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAd],
// [AccountAccessIdentityProviderUpdateParamsPajwohLqCentrify],
// [AccountAccessIdentityProviderUpdateParamsPajwohLqFacebook],
// [AccountAccessIdentityProviderUpdateParamsPajwohLqGitHub],
// [AccountAccessIdentityProviderUpdateParamsPajwohLqGoogle],
// [AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleApps],
// [AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedin],
// [AccountAccessIdentityProviderUpdateParamsPajwohLqOidc],
// [AccountAccessIdentityProviderUpdateParamsPajwohLqOkta],
// [AccountAccessIdentityProviderUpdateParamsPajwohLqOnelogin],
// [AccountAccessIdentityProviderUpdateParamsPajwohLqPingone],
// [AccountAccessIdentityProviderUpdateParamsPajwohLqSaml],
// [AccountAccessIdentityProviderUpdateParamsPajwohLqYandex],
// [AccountAccessIdentityProviderUpdateParamsPajwohLqOnetimepin].
type AccountAccessIdentityProviderUpdateParams interface {
	ImplementsAccountAccessIdentityProviderUpdateParams()
}

type AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAd struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAd) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Should Cloudflare try to load authentication contexts from your account
	ConditionalAccessEnabled param.Field[bool] `json:"conditional_access_enabled"`
	// Your Azure directory uuid
	DirectoryID param.Field[string] `json:"directory_id"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdScimConfig struct {
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

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType string

const (
	AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdTypeOnetimepin AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType = "onetimepin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdTypeAzureAd    AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType = "azureAD"
	AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdTypeSaml       AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType = "saml"
	AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdTypeCentrify   AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType = "centrify"
	AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdTypeFacebook   AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType = "facebook"
	AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdTypeGitHub     AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType = "github"
	AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdTypeGoogleApps AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType = "google-apps"
	AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdTypeGoogle     AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType = "google"
	AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdTypeLinkedin   AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType = "linkedin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdTypeOidc       AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType = "oidc"
	AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdTypeOkta       AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType = "okta"
	AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdTypeOnelogin   AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType = "onelogin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdTypePingone    AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType = "pingone"
	AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdTypeYandex     AccountAccessIdentityProviderUpdateParamsPajwohLqAzureAdType = "yandex"
)

type AccountAccessIdentityProviderUpdateParamsPajwohLqCentrify struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsPajwohLqCentrify) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount param.Field[string] `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID param.Field[string] `json:"centrify_app_id"`
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyScimConfig struct {
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

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType string

const (
	AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyTypeOnetimepin AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType = "onetimepin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyTypeAzureAd    AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType = "azureAD"
	AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyTypeSaml       AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType = "saml"
	AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyTypeCentrify   AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType = "centrify"
	AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyTypeFacebook   AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType = "facebook"
	AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyTypeGitHub     AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType = "github"
	AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyTypeGoogleApps AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType = "google-apps"
	AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyTypeGoogle     AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType = "google"
	AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyTypeLinkedin   AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType = "linkedin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyTypeOidc       AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType = "oidc"
	AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyTypeOkta       AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType = "okta"
	AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyTypeOnelogin   AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType = "onelogin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyTypePingone    AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType = "pingone"
	AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyTypeYandex     AccountAccessIdentityProviderUpdateParamsPajwohLqCentrifyType = "yandex"
)

type AccountAccessIdentityProviderUpdateParamsPajwohLqFacebook struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsPajwohLqFacebook) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookScimConfig struct {
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

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType string

const (
	AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookTypeOnetimepin AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType = "onetimepin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookTypeAzureAd    AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType = "azureAD"
	AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookTypeSaml       AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType = "saml"
	AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookTypeCentrify   AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType = "centrify"
	AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookTypeFacebook   AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType = "facebook"
	AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookTypeGitHub     AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType = "github"
	AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookTypeGoogleApps AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType = "google-apps"
	AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookTypeGoogle     AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType = "google"
	AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookTypeLinkedin   AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType = "linkedin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookTypeOidc       AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType = "oidc"
	AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookTypeOkta       AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType = "okta"
	AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookTypeOnelogin   AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType = "onelogin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookTypePingone    AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType = "pingone"
	AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookTypeYandex     AccountAccessIdentityProviderUpdateParamsPajwohLqFacebookType = "yandex"
)

type AccountAccessIdentityProviderUpdateParamsPajwohLqGitHub struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsPajwohLqGitHub) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubScimConfig struct {
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

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType string

const (
	AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubTypeOnetimepin AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType = "onetimepin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubTypeAzureAd    AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType = "azureAD"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubTypeSaml       AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType = "saml"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubTypeCentrify   AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType = "centrify"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubTypeFacebook   AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType = "facebook"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubTypeGitHub     AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType = "github"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubTypeGoogleApps AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType = "google-apps"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubTypeGoogle     AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType = "google"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubTypeLinkedin   AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType = "linkedin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubTypeOidc       AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType = "oidc"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubTypeOkta       AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType = "okta"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubTypeOnelogin   AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType = "onelogin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubTypePingone    AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType = "pingone"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubTypeYandex     AccountAccessIdentityProviderUpdateParamsPajwohLqGitHubType = "yandex"
)

type AccountAccessIdentityProviderUpdateParamsPajwohLqGoogle struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsPajwohLqGoogle) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleScimConfig struct {
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

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType string

const (
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleTypeOnetimepin AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType = "onetimepin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleTypeAzureAd    AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType = "azureAD"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleTypeSaml       AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType = "saml"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleTypeCentrify   AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType = "centrify"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleTypeFacebook   AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType = "facebook"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleTypeGitHub     AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType = "github"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleTypeGoogleApps AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType = "google-apps"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleTypeGoogle     AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType = "google"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleTypeLinkedin   AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType = "linkedin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleTypeOidc       AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType = "oidc"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleTypeOkta       AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType = "okta"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleTypeOnelogin   AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType = "onelogin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleTypePingone    AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType = "pingone"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleTypeYandex     AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleType = "yandex"
)

type AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleApps struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleApps) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain param.Field[string] `json:"apps_domain"`
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsScimConfig struct {
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

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType string

const (
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsTypeOnetimepin AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType = "onetimepin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsTypeAzureAd    AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType = "azureAD"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsTypeSaml       AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType = "saml"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsTypeCentrify   AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType = "centrify"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsTypeFacebook   AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType = "facebook"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsTypeGitHub     AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType = "github"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsTypeGoogleApps AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType = "google-apps"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsTypeGoogle     AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType = "google"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsTypeLinkedin   AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType = "linkedin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsTypeOidc       AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType = "oidc"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsTypeOkta       AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType = "okta"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsTypeOnelogin   AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType = "onelogin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsTypePingone    AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType = "pingone"
	AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsTypeYandex     AccountAccessIdentityProviderUpdateParamsPajwohLqGoogleAppsType = "yandex"
)

type AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedin struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedin) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinScimConfig struct {
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

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType string

const (
	AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinTypeOnetimepin AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType = "onetimepin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinTypeAzureAd    AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType = "azureAD"
	AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinTypeSaml       AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType = "saml"
	AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinTypeCentrify   AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType = "centrify"
	AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinTypeFacebook   AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType = "facebook"
	AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinTypeGitHub     AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType = "github"
	AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinTypeGoogleApps AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType = "google-apps"
	AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinTypeGoogle     AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType = "google"
	AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinTypeLinkedin   AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType = "linkedin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinTypeOidc       AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType = "oidc"
	AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinTypeOkta       AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType = "okta"
	AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinTypeOnelogin   AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType = "onelogin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinTypePingone    AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType = "pingone"
	AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinTypeYandex     AccountAccessIdentityProviderUpdateParamsPajwohLqLinkedinType = "yandex"
)

type AccountAccessIdentityProviderUpdateParamsPajwohLqOidc struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqOidcConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqOidcScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqOidc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsPajwohLqOidc) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsPajwohLqOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL param.Field[string] `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL param.Field[string] `json:"certs_url"`
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
	// OAuth scopes
	Scopes param.Field[[]string] `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqOidcConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsPajwohLqOidcScimConfig struct {
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

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqOidcScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType string

const (
	AccountAccessIdentityProviderUpdateParamsPajwohLqOidcTypeOnetimepin AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType = "onetimepin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOidcTypeAzureAd    AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType = "azureAD"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOidcTypeSaml       AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType = "saml"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOidcTypeCentrify   AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType = "centrify"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOidcTypeFacebook   AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType = "facebook"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOidcTypeGitHub     AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType = "github"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOidcTypeGoogleApps AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType = "google-apps"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOidcTypeGoogle     AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType = "google"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOidcTypeLinkedin   AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType = "linkedin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOidcTypeOidc       AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType = "oidc"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOidcTypeOkta       AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType = "okta"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOidcTypeOnelogin   AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType = "onelogin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOidcTypePingone    AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType = "pingone"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOidcTypeYandex     AccountAccessIdentityProviderUpdateParamsPajwohLqOidcType = "yandex"
)

type AccountAccessIdentityProviderUpdateParamsPajwohLqOkta struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqOktaConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqOktaScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsPajwohLqOkta) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsPajwohLqOktaConfig struct {
	// Your okta authorization server id
	AuthorizationServerID param.Field[string] `json:"authorization_server_id"`
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
	// Your okta account url
	OktaAccount param.Field[string] `json:"okta_account"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsPajwohLqOktaScimConfig struct {
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

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqOktaScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType string

const (
	AccountAccessIdentityProviderUpdateParamsPajwohLqOktaTypeOnetimepin AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType = "onetimepin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOktaTypeAzureAd    AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType = "azureAD"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOktaTypeSaml       AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType = "saml"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOktaTypeCentrify   AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType = "centrify"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOktaTypeFacebook   AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType = "facebook"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOktaTypeGitHub     AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType = "github"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOktaTypeGoogleApps AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType = "google-apps"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOktaTypeGoogle     AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType = "google"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOktaTypeLinkedin   AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType = "linkedin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOktaTypeOidc       AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType = "oidc"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOktaTypeOkta       AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType = "okta"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOktaTypeOnelogin   AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType = "onelogin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOktaTypePingone    AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType = "pingone"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOktaTypeYandex     AccountAccessIdentityProviderUpdateParamsPajwohLqOktaType = "yandex"
)

type AccountAccessIdentityProviderUpdateParamsPajwohLqOnelogin struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsPajwohLqOnelogin) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount param.Field[string] `json:"onelogin_account"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginScimConfig struct {
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

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType string

const (
	AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginTypeOnetimepin AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType = "onetimepin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginTypeAzureAd    AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType = "azureAD"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginTypeSaml       AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType = "saml"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginTypeCentrify   AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType = "centrify"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginTypeFacebook   AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType = "facebook"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginTypeGitHub     AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType = "github"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginTypeGoogleApps AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType = "google-apps"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginTypeGoogle     AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType = "google"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginTypeLinkedin   AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType = "linkedin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginTypeOidc       AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType = "oidc"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginTypeOkta       AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType = "okta"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginTypeOnelogin   AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType = "onelogin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginTypePingone    AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType = "pingone"
	AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginTypeYandex     AccountAccessIdentityProviderUpdateParamsPajwohLqOneloginType = "yandex"
)

type AccountAccessIdentityProviderUpdateParamsPajwohLqPingone struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsPajwohLqPingone) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID param.Field[string] `json:"ping_env_id"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneScimConfig struct {
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

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType string

const (
	AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneTypeOnetimepin AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType = "onetimepin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneTypeAzureAd    AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType = "azureAD"
	AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneTypeSaml       AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType = "saml"
	AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneTypeCentrify   AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType = "centrify"
	AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneTypeFacebook   AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType = "facebook"
	AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneTypeGitHub     AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType = "github"
	AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneTypeGoogleApps AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType = "google-apps"
	AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneTypeGoogle     AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType = "google"
	AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneTypeLinkedin   AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType = "linkedin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneTypeOidc       AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType = "oidc"
	AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneTypeOkta       AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType = "okta"
	AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneTypeOnelogin   AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType = "onelogin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneTypePingone    AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType = "pingone"
	AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneTypeYandex     AccountAccessIdentityProviderUpdateParamsPajwohLqPingoneType = "yandex"
)

type AccountAccessIdentityProviderUpdateParamsPajwohLqSaml struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqSamlConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqSamlScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsPajwohLqSaml) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsPajwohLqSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]AccountAccessIdentityProviderUpdateParamsPajwohLqSamlConfigHeaderAttribute] `json:"header_attributes"`
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

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqSamlConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderUpdateParamsPajwohLqSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqSamlConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsPajwohLqSamlScimConfig struct {
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

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqSamlScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType string

const (
	AccountAccessIdentityProviderUpdateParamsPajwohLqSamlTypeOnetimepin AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType = "onetimepin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqSamlTypeAzureAd    AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType = "azureAD"
	AccountAccessIdentityProviderUpdateParamsPajwohLqSamlTypeSaml       AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType = "saml"
	AccountAccessIdentityProviderUpdateParamsPajwohLqSamlTypeCentrify   AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType = "centrify"
	AccountAccessIdentityProviderUpdateParamsPajwohLqSamlTypeFacebook   AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType = "facebook"
	AccountAccessIdentityProviderUpdateParamsPajwohLqSamlTypeGitHub     AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType = "github"
	AccountAccessIdentityProviderUpdateParamsPajwohLqSamlTypeGoogleApps AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType = "google-apps"
	AccountAccessIdentityProviderUpdateParamsPajwohLqSamlTypeGoogle     AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType = "google"
	AccountAccessIdentityProviderUpdateParamsPajwohLqSamlTypeLinkedin   AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType = "linkedin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqSamlTypeOidc       AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType = "oidc"
	AccountAccessIdentityProviderUpdateParamsPajwohLqSamlTypeOkta       AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType = "okta"
	AccountAccessIdentityProviderUpdateParamsPajwohLqSamlTypeOnelogin   AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType = "onelogin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqSamlTypePingone    AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType = "pingone"
	AccountAccessIdentityProviderUpdateParamsPajwohLqSamlTypeYandex     AccountAccessIdentityProviderUpdateParamsPajwohLqSamlType = "yandex"
)

type AccountAccessIdentityProviderUpdateParamsPajwohLqYandex struct {
	Config param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqYandexConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqYandexScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType] `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsPajwohLqYandex) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

type AccountAccessIdentityProviderUpdateParamsPajwohLqYandexConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqYandexConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsPajwohLqYandexScimConfig struct {
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

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqYandexScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType string

const (
	AccountAccessIdentityProviderUpdateParamsPajwohLqYandexTypeOnetimepin AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType = "onetimepin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqYandexTypeAzureAd    AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType = "azureAD"
	AccountAccessIdentityProviderUpdateParamsPajwohLqYandexTypeSaml       AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType = "saml"
	AccountAccessIdentityProviderUpdateParamsPajwohLqYandexTypeCentrify   AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType = "centrify"
	AccountAccessIdentityProviderUpdateParamsPajwohLqYandexTypeFacebook   AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType = "facebook"
	AccountAccessIdentityProviderUpdateParamsPajwohLqYandexTypeGitHub     AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType = "github"
	AccountAccessIdentityProviderUpdateParamsPajwohLqYandexTypeGoogleApps AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType = "google-apps"
	AccountAccessIdentityProviderUpdateParamsPajwohLqYandexTypeGoogle     AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType = "google"
	AccountAccessIdentityProviderUpdateParamsPajwohLqYandexTypeLinkedin   AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType = "linkedin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqYandexTypeOidc       AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType = "oidc"
	AccountAccessIdentityProviderUpdateParamsPajwohLqYandexTypeOkta       AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType = "okta"
	AccountAccessIdentityProviderUpdateParamsPajwohLqYandexTypeOnelogin   AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType = "onelogin"
	AccountAccessIdentityProviderUpdateParamsPajwohLqYandexTypePingone    AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType = "pingone"
	AccountAccessIdentityProviderUpdateParamsPajwohLqYandexTypeYandex     AccountAccessIdentityProviderUpdateParamsPajwohLqYandexType = "yandex"
)

type AccountAccessIdentityProviderUpdateParamsPajwohLqOnetimepin struct {
	Config param.Field[interface{}] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqOnetimepinScimConfig] `json:"scim_config"`
	Type       param.Field[AccountAccessIdentityProviderUpdateParamsPajwohLqOnetimepinType]       `json:"type"`
}

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqOnetimepin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderUpdateParamsPajwohLqOnetimepin) ImplementsAccountAccessIdentityProviderUpdateParams() {

}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderUpdateParamsPajwohLqOnetimepinScimConfig struct {
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

func (r AccountAccessIdentityProviderUpdateParamsPajwohLqOnetimepinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderUpdateParamsPajwohLqOnetimepinType string

const (
	AccountAccessIdentityProviderUpdateParamsPajwohLqOnetimepinTypeOnetimepin AccountAccessIdentityProviderUpdateParamsPajwohLqOnetimepinType = "onetimepin"
)

// This interface is a union satisfied by one of the following:
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAd],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrify],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebook],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHub],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogle],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleApps],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedin],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidc],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOkta],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOnelogin],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingone],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSaml],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandex],
// [AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOnetimepin].
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams interface {
	ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams()
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAd struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAd) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Should Cloudflare try to load authentication contexts from your account
	ConditionalAccessEnabled param.Field[bool] `json:"conditional_access_enabled"`
	// Your Azure directory uuid
	DirectoryID param.Field[string] `json:"directory_id"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdScimConfig struct {
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

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqAzureAdType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrify struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrify) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount param.Field[string] `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID param.Field[string] `json:"centrify_app_id"`
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyScimConfig struct {
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

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqCentrifyType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebook struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebook) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookScimConfig struct {
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

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqFacebookType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHub struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHub) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubScimConfig struct {
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

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGitHubType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogle struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogle) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleScimConfig struct {
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

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleApps struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleApps) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain param.Field[string] `json:"apps_domain"`
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsScimConfig struct {
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

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqGoogleAppsType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedin struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedin) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinScimConfig struct {
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

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqLinkedinType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidc struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidc) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL param.Field[string] `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL param.Field[string] `json:"certs_url"`
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
	// OAuth scopes
	Scopes param.Field[[]string] `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL param.Field[string] `json:"token_url"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcScimConfig struct {
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

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOidcType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOkta struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOkta) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaConfig struct {
	// Your okta authorization server id
	AuthorizationServerID param.Field[string] `json:"authorization_server_id"`
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
	// Your okta account url
	OktaAccount param.Field[string] `json:"okta_account"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaScimConfig struct {
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

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOktaType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOnelogin struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOnelogin) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount param.Field[string] `json:"onelogin_account"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginScimConfig struct {
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

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOneloginType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingone struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingone) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID param.Field[string] `json:"ping_env_id"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneScimConfig struct {
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

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqPingoneType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSaml struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSaml) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlConfigHeaderAttribute] `json:"header_attributes"`
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

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlScimConfig struct {
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

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqSamlType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandex struct {
	Config param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType] `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandex) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexScimConfig struct {
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

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType = "onetimepin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexTypeAzureAd    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType = "azureAD"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexTypeSaml       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType = "saml"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexTypeCentrify   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType = "centrify"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexTypeFacebook   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType = "facebook"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexTypeGitHub     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType = "github"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexTypeGoogleApps AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType = "google-apps"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexTypeGoogle     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType = "google"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexTypeLinkedin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType = "linkedin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexTypeOidc       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType = "oidc"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexTypeOkta       AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType = "okta"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexTypeOnelogin   AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType = "onelogin"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexTypePingone    AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType = "pingone"
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexTypeYandex     AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqYandexType = "yandex"
)

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOnetimepin struct {
	Config param.Field[interface{}] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOnetimepinScimConfig] `json:"scim_config"`
	Type       param.Field[AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOnetimepinType]       `json:"type"`
}

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOnetimepin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOnetimepin) ImplementsAccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOnetimepinScimConfig struct {
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

func (r AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOnetimepinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOnetimepinType string

const (
	AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOnetimepinTypeOnetimepin AccountAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsPajwohLqOnetimepinType = "onetimepin"
)

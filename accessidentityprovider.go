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

// AccessIdentityProviderService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessIdentityProviderService]
// method instead.
type AccessIdentityProviderService struct {
	Options []option.RequestOption
}

// NewAccessIdentityProviderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessIdentityProviderService(opts ...option.RequestOption) (r *AccessIdentityProviderService) {
	r = &AccessIdentityProviderService{}
	r.Options = opts
	return
}

// Fetches a configured identity provider.
func (r *AccessIdentityProviderService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessIdentityProviderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/identity_providers/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured identity provider.
func (r *AccessIdentityProviderService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, body AccessIdentityProviderUpdateParams, opts ...option.RequestOption) (res *AccessIdentityProviderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/identity_providers/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an identity provider from Access.
func (r *AccessIdentityProviderService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessIdentityProviderDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/identity_providers/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new identity provider to Access.
func (r *AccessIdentityProviderService) AccessIdentityProvidersAddAnAccessIdentityProvider(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams, opts ...option.RequestOption) (res *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/identity_providers", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all configured identity providers.
func (r *AccessIdentityProviderService) AccessIdentityProvidersListAccessIdentityProviders(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/identity_providers", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessIdentityProviderGetResponse struct {
	Errors   []AccessIdentityProviderGetResponseError   `json:"errors"`
	Messages []AccessIdentityProviderGetResponseMessage `json:"messages"`
	Result   AccessIdentityProviderGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessIdentityProviderGetResponseSuccess `json:"success"`
	JSON    accessIdentityProviderGetResponseJSON    `json:"-"`
}

// accessIdentityProviderGetResponseJSON contains the JSON metadata for the struct
// [AccessIdentityProviderGetResponse]
type accessIdentityProviderGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderGetResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accessIdentityProviderGetResponseErrorJSON `json:"-"`
}

// accessIdentityProviderGetResponseErrorJSON contains the JSON metadata for the
// struct [AccessIdentityProviderGetResponseError]
type accessIdentityProviderGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderGetResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accessIdentityProviderGetResponseMessageJSON `json:"-"`
}

// accessIdentityProviderGetResponseMessageJSON contains the JSON metadata for the
// struct [AccessIdentityProviderGetResponseMessage]
type accessIdentityProviderGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccessIdentityProviderGetResponseResultAccessAzureAd],
// [AccessIdentityProviderGetResponseResultAccessCentrify],
// [AccessIdentityProviderGetResponseResultAccessFacebook],
// [AccessIdentityProviderGetResponseResultAccessGitHub],
// [AccessIdentityProviderGetResponseResultAccessGoogle],
// [AccessIdentityProviderGetResponseResultAccessGoogleApps],
// [AccessIdentityProviderGetResponseResultAccessLinkedin],
// [AccessIdentityProviderGetResponseResultAccessOidc],
// [AccessIdentityProviderGetResponseResultAccessOkta],
// [AccessIdentityProviderGetResponseResultAccessOnelogin],
// [AccessIdentityProviderGetResponseResultAccessPingone],
// [AccessIdentityProviderGetResponseResultAccessSaml],
// [AccessIdentityProviderGetResponseResultAccessYandex] or
// [AccessIdentityProviderGetResponseResultAccessOnetimepin].
type AccessIdentityProviderGetResponseResult interface {
	implementsAccessIdentityProviderGetResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessIdentityProviderGetResponseResult)(nil)).Elem(), "")
}

type AccessIdentityProviderGetResponseResultAccessAzureAd struct {
	// UUID
	ID     string                                                     `json:"id"`
	Config AccessIdentityProviderGetResponseResultAccessAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseResultAccessAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseResultAccessAzureAdType `json:"type"`
	JSON accessIdentityProviderGetResponseResultAccessAzureAdJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessAzureAdJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseResultAccessAzureAd]
type accessIdentityProviderGetResponseResultAccessAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseResultAccessAzureAd) implementsAccessIdentityProviderGetResponseResult() {
}

type AccessIdentityProviderGetResponseResultAccessAzureAdConfig struct {
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
	SupportGroups bool                                                           `json:"support_groups"`
	JSON          accessIdentityProviderGetResponseResultAccessAzureAdConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessAzureAdConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessAzureAdConfig]
type accessIdentityProviderGetResponseResultAccessAzureAdConfigJSON struct {
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

func (r *AccessIdentityProviderGetResponseResultAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseResultAccessAzureAdScimConfig struct {
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
	UserDeprovision bool                                                               `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseResultAccessAzureAdScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessAzureAdScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessAzureAdScimConfig]
type accessIdentityProviderGetResponseResultAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseResultAccessAzureAdType string

const (
	AccessIdentityProviderGetResponseResultAccessAzureAdTypeOnetimepin AccessIdentityProviderGetResponseResultAccessAzureAdType = "onetimepin"
	AccessIdentityProviderGetResponseResultAccessAzureAdTypeAzureAd    AccessIdentityProviderGetResponseResultAccessAzureAdType = "azureAD"
	AccessIdentityProviderGetResponseResultAccessAzureAdTypeSaml       AccessIdentityProviderGetResponseResultAccessAzureAdType = "saml"
	AccessIdentityProviderGetResponseResultAccessAzureAdTypeCentrify   AccessIdentityProviderGetResponseResultAccessAzureAdType = "centrify"
	AccessIdentityProviderGetResponseResultAccessAzureAdTypeFacebook   AccessIdentityProviderGetResponseResultAccessAzureAdType = "facebook"
	AccessIdentityProviderGetResponseResultAccessAzureAdTypeGitHub     AccessIdentityProviderGetResponseResultAccessAzureAdType = "github"
	AccessIdentityProviderGetResponseResultAccessAzureAdTypeGoogleApps AccessIdentityProviderGetResponseResultAccessAzureAdType = "google-apps"
	AccessIdentityProviderGetResponseResultAccessAzureAdTypeGoogle     AccessIdentityProviderGetResponseResultAccessAzureAdType = "google"
	AccessIdentityProviderGetResponseResultAccessAzureAdTypeLinkedin   AccessIdentityProviderGetResponseResultAccessAzureAdType = "linkedin"
	AccessIdentityProviderGetResponseResultAccessAzureAdTypeOidc       AccessIdentityProviderGetResponseResultAccessAzureAdType = "oidc"
	AccessIdentityProviderGetResponseResultAccessAzureAdTypeOkta       AccessIdentityProviderGetResponseResultAccessAzureAdType = "okta"
	AccessIdentityProviderGetResponseResultAccessAzureAdTypeOnelogin   AccessIdentityProviderGetResponseResultAccessAzureAdType = "onelogin"
	AccessIdentityProviderGetResponseResultAccessAzureAdTypePingone    AccessIdentityProviderGetResponseResultAccessAzureAdType = "pingone"
	AccessIdentityProviderGetResponseResultAccessAzureAdTypeYandex     AccessIdentityProviderGetResponseResultAccessAzureAdType = "yandex"
)

type AccessIdentityProviderGetResponseResultAccessCentrify struct {
	// UUID
	ID     string                                                      `json:"id"`
	Config AccessIdentityProviderGetResponseResultAccessCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseResultAccessCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseResultAccessCentrifyType `json:"type"`
	JSON accessIdentityProviderGetResponseResultAccessCentrifyJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessCentrifyJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseResultAccessCentrify]
type accessIdentityProviderGetResponseResultAccessCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseResultAccessCentrify) implementsAccessIdentityProviderGetResponseResult() {
}

type AccessIdentityProviderGetResponseResultAccessCentrifyConfig struct {
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
	EmailClaimName string                                                          `json:"email_claim_name"`
	JSON           accessIdentityProviderGetResponseResultAccessCentrifyConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessCentrifyConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessCentrifyConfig]
type accessIdentityProviderGetResponseResultAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseResultAccessCentrifyScimConfig struct {
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
	UserDeprovision bool                                                                `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseResultAccessCentrifyScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessCentrifyScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessCentrifyScimConfig]
type accessIdentityProviderGetResponseResultAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseResultAccessCentrifyType string

const (
	AccessIdentityProviderGetResponseResultAccessCentrifyTypeOnetimepin AccessIdentityProviderGetResponseResultAccessCentrifyType = "onetimepin"
	AccessIdentityProviderGetResponseResultAccessCentrifyTypeAzureAd    AccessIdentityProviderGetResponseResultAccessCentrifyType = "azureAD"
	AccessIdentityProviderGetResponseResultAccessCentrifyTypeSaml       AccessIdentityProviderGetResponseResultAccessCentrifyType = "saml"
	AccessIdentityProviderGetResponseResultAccessCentrifyTypeCentrify   AccessIdentityProviderGetResponseResultAccessCentrifyType = "centrify"
	AccessIdentityProviderGetResponseResultAccessCentrifyTypeFacebook   AccessIdentityProviderGetResponseResultAccessCentrifyType = "facebook"
	AccessIdentityProviderGetResponseResultAccessCentrifyTypeGitHub     AccessIdentityProviderGetResponseResultAccessCentrifyType = "github"
	AccessIdentityProviderGetResponseResultAccessCentrifyTypeGoogleApps AccessIdentityProviderGetResponseResultAccessCentrifyType = "google-apps"
	AccessIdentityProviderGetResponseResultAccessCentrifyTypeGoogle     AccessIdentityProviderGetResponseResultAccessCentrifyType = "google"
	AccessIdentityProviderGetResponseResultAccessCentrifyTypeLinkedin   AccessIdentityProviderGetResponseResultAccessCentrifyType = "linkedin"
	AccessIdentityProviderGetResponseResultAccessCentrifyTypeOidc       AccessIdentityProviderGetResponseResultAccessCentrifyType = "oidc"
	AccessIdentityProviderGetResponseResultAccessCentrifyTypeOkta       AccessIdentityProviderGetResponseResultAccessCentrifyType = "okta"
	AccessIdentityProviderGetResponseResultAccessCentrifyTypeOnelogin   AccessIdentityProviderGetResponseResultAccessCentrifyType = "onelogin"
	AccessIdentityProviderGetResponseResultAccessCentrifyTypePingone    AccessIdentityProviderGetResponseResultAccessCentrifyType = "pingone"
	AccessIdentityProviderGetResponseResultAccessCentrifyTypeYandex     AccessIdentityProviderGetResponseResultAccessCentrifyType = "yandex"
)

type AccessIdentityProviderGetResponseResultAccessFacebook struct {
	// UUID
	ID     string                                                      `json:"id"`
	Config AccessIdentityProviderGetResponseResultAccessFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseResultAccessFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseResultAccessFacebookType `json:"type"`
	JSON accessIdentityProviderGetResponseResultAccessFacebookJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessFacebookJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseResultAccessFacebook]
type accessIdentityProviderGetResponseResultAccessFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseResultAccessFacebook) implementsAccessIdentityProviderGetResponseResult() {
}

type AccessIdentityProviderGetResponseResultAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                          `json:"client_secret"`
	JSON         accessIdentityProviderGetResponseResultAccessFacebookConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessFacebookConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessFacebookConfig]
type accessIdentityProviderGetResponseResultAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseResultAccessFacebookScimConfig struct {
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
	UserDeprovision bool                                                                `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseResultAccessFacebookScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessFacebookScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessFacebookScimConfig]
type accessIdentityProviderGetResponseResultAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseResultAccessFacebookType string

const (
	AccessIdentityProviderGetResponseResultAccessFacebookTypeOnetimepin AccessIdentityProviderGetResponseResultAccessFacebookType = "onetimepin"
	AccessIdentityProviderGetResponseResultAccessFacebookTypeAzureAd    AccessIdentityProviderGetResponseResultAccessFacebookType = "azureAD"
	AccessIdentityProviderGetResponseResultAccessFacebookTypeSaml       AccessIdentityProviderGetResponseResultAccessFacebookType = "saml"
	AccessIdentityProviderGetResponseResultAccessFacebookTypeCentrify   AccessIdentityProviderGetResponseResultAccessFacebookType = "centrify"
	AccessIdentityProviderGetResponseResultAccessFacebookTypeFacebook   AccessIdentityProviderGetResponseResultAccessFacebookType = "facebook"
	AccessIdentityProviderGetResponseResultAccessFacebookTypeGitHub     AccessIdentityProviderGetResponseResultAccessFacebookType = "github"
	AccessIdentityProviderGetResponseResultAccessFacebookTypeGoogleApps AccessIdentityProviderGetResponseResultAccessFacebookType = "google-apps"
	AccessIdentityProviderGetResponseResultAccessFacebookTypeGoogle     AccessIdentityProviderGetResponseResultAccessFacebookType = "google"
	AccessIdentityProviderGetResponseResultAccessFacebookTypeLinkedin   AccessIdentityProviderGetResponseResultAccessFacebookType = "linkedin"
	AccessIdentityProviderGetResponseResultAccessFacebookTypeOidc       AccessIdentityProviderGetResponseResultAccessFacebookType = "oidc"
	AccessIdentityProviderGetResponseResultAccessFacebookTypeOkta       AccessIdentityProviderGetResponseResultAccessFacebookType = "okta"
	AccessIdentityProviderGetResponseResultAccessFacebookTypeOnelogin   AccessIdentityProviderGetResponseResultAccessFacebookType = "onelogin"
	AccessIdentityProviderGetResponseResultAccessFacebookTypePingone    AccessIdentityProviderGetResponseResultAccessFacebookType = "pingone"
	AccessIdentityProviderGetResponseResultAccessFacebookTypeYandex     AccessIdentityProviderGetResponseResultAccessFacebookType = "yandex"
)

type AccessIdentityProviderGetResponseResultAccessGitHub struct {
	// UUID
	ID     string                                                    `json:"id"`
	Config AccessIdentityProviderGetResponseResultAccessGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseResultAccessGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseResultAccessGitHubType `json:"type"`
	JSON accessIdentityProviderGetResponseResultAccessGitHubJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessGitHubJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseResultAccessGitHub]
type accessIdentityProviderGetResponseResultAccessGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseResultAccessGitHub) implementsAccessIdentityProviderGetResponseResult() {
}

type AccessIdentityProviderGetResponseResultAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                        `json:"client_secret"`
	JSON         accessIdentityProviderGetResponseResultAccessGitHubConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessGitHubConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessGitHubConfig]
type accessIdentityProviderGetResponseResultAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseResultAccessGitHubScimConfig struct {
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
	UserDeprovision bool                                                              `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseResultAccessGitHubScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessGitHubScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessGitHubScimConfig]
type accessIdentityProviderGetResponseResultAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseResultAccessGitHubType string

const (
	AccessIdentityProviderGetResponseResultAccessGitHubTypeOnetimepin AccessIdentityProviderGetResponseResultAccessGitHubType = "onetimepin"
	AccessIdentityProviderGetResponseResultAccessGitHubTypeAzureAd    AccessIdentityProviderGetResponseResultAccessGitHubType = "azureAD"
	AccessIdentityProviderGetResponseResultAccessGitHubTypeSaml       AccessIdentityProviderGetResponseResultAccessGitHubType = "saml"
	AccessIdentityProviderGetResponseResultAccessGitHubTypeCentrify   AccessIdentityProviderGetResponseResultAccessGitHubType = "centrify"
	AccessIdentityProviderGetResponseResultAccessGitHubTypeFacebook   AccessIdentityProviderGetResponseResultAccessGitHubType = "facebook"
	AccessIdentityProviderGetResponseResultAccessGitHubTypeGitHub     AccessIdentityProviderGetResponseResultAccessGitHubType = "github"
	AccessIdentityProviderGetResponseResultAccessGitHubTypeGoogleApps AccessIdentityProviderGetResponseResultAccessGitHubType = "google-apps"
	AccessIdentityProviderGetResponseResultAccessGitHubTypeGoogle     AccessIdentityProviderGetResponseResultAccessGitHubType = "google"
	AccessIdentityProviderGetResponseResultAccessGitHubTypeLinkedin   AccessIdentityProviderGetResponseResultAccessGitHubType = "linkedin"
	AccessIdentityProviderGetResponseResultAccessGitHubTypeOidc       AccessIdentityProviderGetResponseResultAccessGitHubType = "oidc"
	AccessIdentityProviderGetResponseResultAccessGitHubTypeOkta       AccessIdentityProviderGetResponseResultAccessGitHubType = "okta"
	AccessIdentityProviderGetResponseResultAccessGitHubTypeOnelogin   AccessIdentityProviderGetResponseResultAccessGitHubType = "onelogin"
	AccessIdentityProviderGetResponseResultAccessGitHubTypePingone    AccessIdentityProviderGetResponseResultAccessGitHubType = "pingone"
	AccessIdentityProviderGetResponseResultAccessGitHubTypeYandex     AccessIdentityProviderGetResponseResultAccessGitHubType = "yandex"
)

type AccessIdentityProviderGetResponseResultAccessGoogle struct {
	// UUID
	ID     string                                                    `json:"id"`
	Config AccessIdentityProviderGetResponseResultAccessGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseResultAccessGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseResultAccessGoogleType `json:"type"`
	JSON accessIdentityProviderGetResponseResultAccessGoogleJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessGoogleJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseResultAccessGoogle]
type accessIdentityProviderGetResponseResultAccessGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseResultAccessGoogle) implementsAccessIdentityProviderGetResponseResult() {
}

type AccessIdentityProviderGetResponseResultAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                        `json:"email_claim_name"`
	JSON           accessIdentityProviderGetResponseResultAccessGoogleConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessGoogleConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessGoogleConfig]
type accessIdentityProviderGetResponseResultAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseResultAccessGoogleScimConfig struct {
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
	UserDeprovision bool                                                              `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseResultAccessGoogleScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessGoogleScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessGoogleScimConfig]
type accessIdentityProviderGetResponseResultAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseResultAccessGoogleType string

const (
	AccessIdentityProviderGetResponseResultAccessGoogleTypeOnetimepin AccessIdentityProviderGetResponseResultAccessGoogleType = "onetimepin"
	AccessIdentityProviderGetResponseResultAccessGoogleTypeAzureAd    AccessIdentityProviderGetResponseResultAccessGoogleType = "azureAD"
	AccessIdentityProviderGetResponseResultAccessGoogleTypeSaml       AccessIdentityProviderGetResponseResultAccessGoogleType = "saml"
	AccessIdentityProviderGetResponseResultAccessGoogleTypeCentrify   AccessIdentityProviderGetResponseResultAccessGoogleType = "centrify"
	AccessIdentityProviderGetResponseResultAccessGoogleTypeFacebook   AccessIdentityProviderGetResponseResultAccessGoogleType = "facebook"
	AccessIdentityProviderGetResponseResultAccessGoogleTypeGitHub     AccessIdentityProviderGetResponseResultAccessGoogleType = "github"
	AccessIdentityProviderGetResponseResultAccessGoogleTypeGoogleApps AccessIdentityProviderGetResponseResultAccessGoogleType = "google-apps"
	AccessIdentityProviderGetResponseResultAccessGoogleTypeGoogle     AccessIdentityProviderGetResponseResultAccessGoogleType = "google"
	AccessIdentityProviderGetResponseResultAccessGoogleTypeLinkedin   AccessIdentityProviderGetResponseResultAccessGoogleType = "linkedin"
	AccessIdentityProviderGetResponseResultAccessGoogleTypeOidc       AccessIdentityProviderGetResponseResultAccessGoogleType = "oidc"
	AccessIdentityProviderGetResponseResultAccessGoogleTypeOkta       AccessIdentityProviderGetResponseResultAccessGoogleType = "okta"
	AccessIdentityProviderGetResponseResultAccessGoogleTypeOnelogin   AccessIdentityProviderGetResponseResultAccessGoogleType = "onelogin"
	AccessIdentityProviderGetResponseResultAccessGoogleTypePingone    AccessIdentityProviderGetResponseResultAccessGoogleType = "pingone"
	AccessIdentityProviderGetResponseResultAccessGoogleTypeYandex     AccessIdentityProviderGetResponseResultAccessGoogleType = "yandex"
)

type AccessIdentityProviderGetResponseResultAccessGoogleApps struct {
	// UUID
	ID     string                                                        `json:"id"`
	Config AccessIdentityProviderGetResponseResultAccessGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseResultAccessGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseResultAccessGoogleAppsType `json:"type"`
	JSON accessIdentityProviderGetResponseResultAccessGoogleAppsJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessGoogleAppsJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessGoogleApps]
type accessIdentityProviderGetResponseResultAccessGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseResultAccessGoogleApps) implementsAccessIdentityProviderGetResponseResult() {
}

type AccessIdentityProviderGetResponseResultAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                            `json:"email_claim_name"`
	JSON           accessIdentityProviderGetResponseResultAccessGoogleAppsConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessGoogleAppsConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessGoogleAppsConfig]
type accessIdentityProviderGetResponseResultAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseResultAccessGoogleAppsScimConfig struct {
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
	UserDeprovision bool                                                                  `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseResultAccessGoogleAppsScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessGoogleAppsScimConfigJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessGoogleAppsScimConfig]
type accessIdentityProviderGetResponseResultAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseResultAccessGoogleAppsType string

const (
	AccessIdentityProviderGetResponseResultAccessGoogleAppsTypeOnetimepin AccessIdentityProviderGetResponseResultAccessGoogleAppsType = "onetimepin"
	AccessIdentityProviderGetResponseResultAccessGoogleAppsTypeAzureAd    AccessIdentityProviderGetResponseResultAccessGoogleAppsType = "azureAD"
	AccessIdentityProviderGetResponseResultAccessGoogleAppsTypeSaml       AccessIdentityProviderGetResponseResultAccessGoogleAppsType = "saml"
	AccessIdentityProviderGetResponseResultAccessGoogleAppsTypeCentrify   AccessIdentityProviderGetResponseResultAccessGoogleAppsType = "centrify"
	AccessIdentityProviderGetResponseResultAccessGoogleAppsTypeFacebook   AccessIdentityProviderGetResponseResultAccessGoogleAppsType = "facebook"
	AccessIdentityProviderGetResponseResultAccessGoogleAppsTypeGitHub     AccessIdentityProviderGetResponseResultAccessGoogleAppsType = "github"
	AccessIdentityProviderGetResponseResultAccessGoogleAppsTypeGoogleApps AccessIdentityProviderGetResponseResultAccessGoogleAppsType = "google-apps"
	AccessIdentityProviderGetResponseResultAccessGoogleAppsTypeGoogle     AccessIdentityProviderGetResponseResultAccessGoogleAppsType = "google"
	AccessIdentityProviderGetResponseResultAccessGoogleAppsTypeLinkedin   AccessIdentityProviderGetResponseResultAccessGoogleAppsType = "linkedin"
	AccessIdentityProviderGetResponseResultAccessGoogleAppsTypeOidc       AccessIdentityProviderGetResponseResultAccessGoogleAppsType = "oidc"
	AccessIdentityProviderGetResponseResultAccessGoogleAppsTypeOkta       AccessIdentityProviderGetResponseResultAccessGoogleAppsType = "okta"
	AccessIdentityProviderGetResponseResultAccessGoogleAppsTypeOnelogin   AccessIdentityProviderGetResponseResultAccessGoogleAppsType = "onelogin"
	AccessIdentityProviderGetResponseResultAccessGoogleAppsTypePingone    AccessIdentityProviderGetResponseResultAccessGoogleAppsType = "pingone"
	AccessIdentityProviderGetResponseResultAccessGoogleAppsTypeYandex     AccessIdentityProviderGetResponseResultAccessGoogleAppsType = "yandex"
)

type AccessIdentityProviderGetResponseResultAccessLinkedin struct {
	// UUID
	ID     string                                                      `json:"id"`
	Config AccessIdentityProviderGetResponseResultAccessLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseResultAccessLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseResultAccessLinkedinType `json:"type"`
	JSON accessIdentityProviderGetResponseResultAccessLinkedinJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessLinkedinJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseResultAccessLinkedin]
type accessIdentityProviderGetResponseResultAccessLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseResultAccessLinkedin) implementsAccessIdentityProviderGetResponseResult() {
}

type AccessIdentityProviderGetResponseResultAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                          `json:"client_secret"`
	JSON         accessIdentityProviderGetResponseResultAccessLinkedinConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessLinkedinConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessLinkedinConfig]
type accessIdentityProviderGetResponseResultAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseResultAccessLinkedinScimConfig struct {
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
	UserDeprovision bool                                                                `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseResultAccessLinkedinScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessLinkedinScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessLinkedinScimConfig]
type accessIdentityProviderGetResponseResultAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseResultAccessLinkedinType string

const (
	AccessIdentityProviderGetResponseResultAccessLinkedinTypeOnetimepin AccessIdentityProviderGetResponseResultAccessLinkedinType = "onetimepin"
	AccessIdentityProviderGetResponseResultAccessLinkedinTypeAzureAd    AccessIdentityProviderGetResponseResultAccessLinkedinType = "azureAD"
	AccessIdentityProviderGetResponseResultAccessLinkedinTypeSaml       AccessIdentityProviderGetResponseResultAccessLinkedinType = "saml"
	AccessIdentityProviderGetResponseResultAccessLinkedinTypeCentrify   AccessIdentityProviderGetResponseResultAccessLinkedinType = "centrify"
	AccessIdentityProviderGetResponseResultAccessLinkedinTypeFacebook   AccessIdentityProviderGetResponseResultAccessLinkedinType = "facebook"
	AccessIdentityProviderGetResponseResultAccessLinkedinTypeGitHub     AccessIdentityProviderGetResponseResultAccessLinkedinType = "github"
	AccessIdentityProviderGetResponseResultAccessLinkedinTypeGoogleApps AccessIdentityProviderGetResponseResultAccessLinkedinType = "google-apps"
	AccessIdentityProviderGetResponseResultAccessLinkedinTypeGoogle     AccessIdentityProviderGetResponseResultAccessLinkedinType = "google"
	AccessIdentityProviderGetResponseResultAccessLinkedinTypeLinkedin   AccessIdentityProviderGetResponseResultAccessLinkedinType = "linkedin"
	AccessIdentityProviderGetResponseResultAccessLinkedinTypeOidc       AccessIdentityProviderGetResponseResultAccessLinkedinType = "oidc"
	AccessIdentityProviderGetResponseResultAccessLinkedinTypeOkta       AccessIdentityProviderGetResponseResultAccessLinkedinType = "okta"
	AccessIdentityProviderGetResponseResultAccessLinkedinTypeOnelogin   AccessIdentityProviderGetResponseResultAccessLinkedinType = "onelogin"
	AccessIdentityProviderGetResponseResultAccessLinkedinTypePingone    AccessIdentityProviderGetResponseResultAccessLinkedinType = "pingone"
	AccessIdentityProviderGetResponseResultAccessLinkedinTypeYandex     AccessIdentityProviderGetResponseResultAccessLinkedinType = "yandex"
)

type AccessIdentityProviderGetResponseResultAccessOidc struct {
	// UUID
	ID     string                                                  `json:"id"`
	Config AccessIdentityProviderGetResponseResultAccessOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseResultAccessOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseResultAccessOidcType `json:"type"`
	JSON accessIdentityProviderGetResponseResultAccessOidcJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessOidcJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseResultAccessOidc]
type accessIdentityProviderGetResponseResultAccessOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseResultAccessOidc) implementsAccessIdentityProviderGetResponseResult() {
}

type AccessIdentityProviderGetResponseResultAccessOidcConfig struct {
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
	TokenURL string                                                      `json:"token_url"`
	JSON     accessIdentityProviderGetResponseResultAccessOidcConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessOidcConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessOidcConfig]
type accessIdentityProviderGetResponseResultAccessOidcConfigJSON struct {
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

func (r *AccessIdentityProviderGetResponseResultAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseResultAccessOidcScimConfig struct {
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
	UserDeprovision bool                                                            `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseResultAccessOidcScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessOidcScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessOidcScimConfig]
type accessIdentityProviderGetResponseResultAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseResultAccessOidcType string

const (
	AccessIdentityProviderGetResponseResultAccessOidcTypeOnetimepin AccessIdentityProviderGetResponseResultAccessOidcType = "onetimepin"
	AccessIdentityProviderGetResponseResultAccessOidcTypeAzureAd    AccessIdentityProviderGetResponseResultAccessOidcType = "azureAD"
	AccessIdentityProviderGetResponseResultAccessOidcTypeSaml       AccessIdentityProviderGetResponseResultAccessOidcType = "saml"
	AccessIdentityProviderGetResponseResultAccessOidcTypeCentrify   AccessIdentityProviderGetResponseResultAccessOidcType = "centrify"
	AccessIdentityProviderGetResponseResultAccessOidcTypeFacebook   AccessIdentityProviderGetResponseResultAccessOidcType = "facebook"
	AccessIdentityProviderGetResponseResultAccessOidcTypeGitHub     AccessIdentityProviderGetResponseResultAccessOidcType = "github"
	AccessIdentityProviderGetResponseResultAccessOidcTypeGoogleApps AccessIdentityProviderGetResponseResultAccessOidcType = "google-apps"
	AccessIdentityProviderGetResponseResultAccessOidcTypeGoogle     AccessIdentityProviderGetResponseResultAccessOidcType = "google"
	AccessIdentityProviderGetResponseResultAccessOidcTypeLinkedin   AccessIdentityProviderGetResponseResultAccessOidcType = "linkedin"
	AccessIdentityProviderGetResponseResultAccessOidcTypeOidc       AccessIdentityProviderGetResponseResultAccessOidcType = "oidc"
	AccessIdentityProviderGetResponseResultAccessOidcTypeOkta       AccessIdentityProviderGetResponseResultAccessOidcType = "okta"
	AccessIdentityProviderGetResponseResultAccessOidcTypeOnelogin   AccessIdentityProviderGetResponseResultAccessOidcType = "onelogin"
	AccessIdentityProviderGetResponseResultAccessOidcTypePingone    AccessIdentityProviderGetResponseResultAccessOidcType = "pingone"
	AccessIdentityProviderGetResponseResultAccessOidcTypeYandex     AccessIdentityProviderGetResponseResultAccessOidcType = "yandex"
)

type AccessIdentityProviderGetResponseResultAccessOkta struct {
	// UUID
	ID     string                                                  `json:"id"`
	Config AccessIdentityProviderGetResponseResultAccessOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseResultAccessOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseResultAccessOktaType `json:"type"`
	JSON accessIdentityProviderGetResponseResultAccessOktaJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessOktaJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseResultAccessOkta]
type accessIdentityProviderGetResponseResultAccessOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseResultAccessOkta) implementsAccessIdentityProviderGetResponseResult() {
}

type AccessIdentityProviderGetResponseResultAccessOktaConfig struct {
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
	OktaAccount string                                                      `json:"okta_account"`
	JSON        accessIdentityProviderGetResponseResultAccessOktaConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessOktaConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessOktaConfig]
type accessIdentityProviderGetResponseResultAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseResultAccessOktaScimConfig struct {
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
	UserDeprovision bool                                                            `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseResultAccessOktaScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessOktaScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessOktaScimConfig]
type accessIdentityProviderGetResponseResultAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseResultAccessOktaType string

const (
	AccessIdentityProviderGetResponseResultAccessOktaTypeOnetimepin AccessIdentityProviderGetResponseResultAccessOktaType = "onetimepin"
	AccessIdentityProviderGetResponseResultAccessOktaTypeAzureAd    AccessIdentityProviderGetResponseResultAccessOktaType = "azureAD"
	AccessIdentityProviderGetResponseResultAccessOktaTypeSaml       AccessIdentityProviderGetResponseResultAccessOktaType = "saml"
	AccessIdentityProviderGetResponseResultAccessOktaTypeCentrify   AccessIdentityProviderGetResponseResultAccessOktaType = "centrify"
	AccessIdentityProviderGetResponseResultAccessOktaTypeFacebook   AccessIdentityProviderGetResponseResultAccessOktaType = "facebook"
	AccessIdentityProviderGetResponseResultAccessOktaTypeGitHub     AccessIdentityProviderGetResponseResultAccessOktaType = "github"
	AccessIdentityProviderGetResponseResultAccessOktaTypeGoogleApps AccessIdentityProviderGetResponseResultAccessOktaType = "google-apps"
	AccessIdentityProviderGetResponseResultAccessOktaTypeGoogle     AccessIdentityProviderGetResponseResultAccessOktaType = "google"
	AccessIdentityProviderGetResponseResultAccessOktaTypeLinkedin   AccessIdentityProviderGetResponseResultAccessOktaType = "linkedin"
	AccessIdentityProviderGetResponseResultAccessOktaTypeOidc       AccessIdentityProviderGetResponseResultAccessOktaType = "oidc"
	AccessIdentityProviderGetResponseResultAccessOktaTypeOkta       AccessIdentityProviderGetResponseResultAccessOktaType = "okta"
	AccessIdentityProviderGetResponseResultAccessOktaTypeOnelogin   AccessIdentityProviderGetResponseResultAccessOktaType = "onelogin"
	AccessIdentityProviderGetResponseResultAccessOktaTypePingone    AccessIdentityProviderGetResponseResultAccessOktaType = "pingone"
	AccessIdentityProviderGetResponseResultAccessOktaTypeYandex     AccessIdentityProviderGetResponseResultAccessOktaType = "yandex"
)

type AccessIdentityProviderGetResponseResultAccessOnelogin struct {
	// UUID
	ID     string                                                      `json:"id"`
	Config AccessIdentityProviderGetResponseResultAccessOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseResultAccessOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseResultAccessOneloginType `json:"type"`
	JSON accessIdentityProviderGetResponseResultAccessOneloginJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessOneloginJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseResultAccessOnelogin]
type accessIdentityProviderGetResponseResultAccessOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseResultAccessOnelogin) implementsAccessIdentityProviderGetResponseResult() {
}

type AccessIdentityProviderGetResponseResultAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                          `json:"onelogin_account"`
	JSON            accessIdentityProviderGetResponseResultAccessOneloginConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessOneloginConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessOneloginConfig]
type accessIdentityProviderGetResponseResultAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseResultAccessOneloginScimConfig struct {
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
	UserDeprovision bool                                                                `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseResultAccessOneloginScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessOneloginScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessOneloginScimConfig]
type accessIdentityProviderGetResponseResultAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseResultAccessOneloginType string

const (
	AccessIdentityProviderGetResponseResultAccessOneloginTypeOnetimepin AccessIdentityProviderGetResponseResultAccessOneloginType = "onetimepin"
	AccessIdentityProviderGetResponseResultAccessOneloginTypeAzureAd    AccessIdentityProviderGetResponseResultAccessOneloginType = "azureAD"
	AccessIdentityProviderGetResponseResultAccessOneloginTypeSaml       AccessIdentityProviderGetResponseResultAccessOneloginType = "saml"
	AccessIdentityProviderGetResponseResultAccessOneloginTypeCentrify   AccessIdentityProviderGetResponseResultAccessOneloginType = "centrify"
	AccessIdentityProviderGetResponseResultAccessOneloginTypeFacebook   AccessIdentityProviderGetResponseResultAccessOneloginType = "facebook"
	AccessIdentityProviderGetResponseResultAccessOneloginTypeGitHub     AccessIdentityProviderGetResponseResultAccessOneloginType = "github"
	AccessIdentityProviderGetResponseResultAccessOneloginTypeGoogleApps AccessIdentityProviderGetResponseResultAccessOneloginType = "google-apps"
	AccessIdentityProviderGetResponseResultAccessOneloginTypeGoogle     AccessIdentityProviderGetResponseResultAccessOneloginType = "google"
	AccessIdentityProviderGetResponseResultAccessOneloginTypeLinkedin   AccessIdentityProviderGetResponseResultAccessOneloginType = "linkedin"
	AccessIdentityProviderGetResponseResultAccessOneloginTypeOidc       AccessIdentityProviderGetResponseResultAccessOneloginType = "oidc"
	AccessIdentityProviderGetResponseResultAccessOneloginTypeOkta       AccessIdentityProviderGetResponseResultAccessOneloginType = "okta"
	AccessIdentityProviderGetResponseResultAccessOneloginTypeOnelogin   AccessIdentityProviderGetResponseResultAccessOneloginType = "onelogin"
	AccessIdentityProviderGetResponseResultAccessOneloginTypePingone    AccessIdentityProviderGetResponseResultAccessOneloginType = "pingone"
	AccessIdentityProviderGetResponseResultAccessOneloginTypeYandex     AccessIdentityProviderGetResponseResultAccessOneloginType = "yandex"
)

type AccessIdentityProviderGetResponseResultAccessPingone struct {
	// UUID
	ID     string                                                     `json:"id"`
	Config AccessIdentityProviderGetResponseResultAccessPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseResultAccessPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseResultAccessPingoneType `json:"type"`
	JSON accessIdentityProviderGetResponseResultAccessPingoneJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessPingoneJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseResultAccessPingone]
type accessIdentityProviderGetResponseResultAccessPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseResultAccessPingone) implementsAccessIdentityProviderGetResponseResult() {
}

type AccessIdentityProviderGetResponseResultAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                         `json:"ping_env_id"`
	JSON      accessIdentityProviderGetResponseResultAccessPingoneConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessPingoneConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessPingoneConfig]
type accessIdentityProviderGetResponseResultAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseResultAccessPingoneScimConfig struct {
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
	UserDeprovision bool                                                               `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseResultAccessPingoneScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessPingoneScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessPingoneScimConfig]
type accessIdentityProviderGetResponseResultAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseResultAccessPingoneType string

const (
	AccessIdentityProviderGetResponseResultAccessPingoneTypeOnetimepin AccessIdentityProviderGetResponseResultAccessPingoneType = "onetimepin"
	AccessIdentityProviderGetResponseResultAccessPingoneTypeAzureAd    AccessIdentityProviderGetResponseResultAccessPingoneType = "azureAD"
	AccessIdentityProviderGetResponseResultAccessPingoneTypeSaml       AccessIdentityProviderGetResponseResultAccessPingoneType = "saml"
	AccessIdentityProviderGetResponseResultAccessPingoneTypeCentrify   AccessIdentityProviderGetResponseResultAccessPingoneType = "centrify"
	AccessIdentityProviderGetResponseResultAccessPingoneTypeFacebook   AccessIdentityProviderGetResponseResultAccessPingoneType = "facebook"
	AccessIdentityProviderGetResponseResultAccessPingoneTypeGitHub     AccessIdentityProviderGetResponseResultAccessPingoneType = "github"
	AccessIdentityProviderGetResponseResultAccessPingoneTypeGoogleApps AccessIdentityProviderGetResponseResultAccessPingoneType = "google-apps"
	AccessIdentityProviderGetResponseResultAccessPingoneTypeGoogle     AccessIdentityProviderGetResponseResultAccessPingoneType = "google"
	AccessIdentityProviderGetResponseResultAccessPingoneTypeLinkedin   AccessIdentityProviderGetResponseResultAccessPingoneType = "linkedin"
	AccessIdentityProviderGetResponseResultAccessPingoneTypeOidc       AccessIdentityProviderGetResponseResultAccessPingoneType = "oidc"
	AccessIdentityProviderGetResponseResultAccessPingoneTypeOkta       AccessIdentityProviderGetResponseResultAccessPingoneType = "okta"
	AccessIdentityProviderGetResponseResultAccessPingoneTypeOnelogin   AccessIdentityProviderGetResponseResultAccessPingoneType = "onelogin"
	AccessIdentityProviderGetResponseResultAccessPingoneTypePingone    AccessIdentityProviderGetResponseResultAccessPingoneType = "pingone"
	AccessIdentityProviderGetResponseResultAccessPingoneTypeYandex     AccessIdentityProviderGetResponseResultAccessPingoneType = "yandex"
)

type AccessIdentityProviderGetResponseResultAccessSaml struct {
	// UUID
	ID     string                                                  `json:"id"`
	Config AccessIdentityProviderGetResponseResultAccessSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseResultAccessSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseResultAccessSamlType `json:"type"`
	JSON accessIdentityProviderGetResponseResultAccessSamlJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessSamlJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseResultAccessSaml]
type accessIdentityProviderGetResponseResultAccessSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseResultAccessSaml) implementsAccessIdentityProviderGetResponseResult() {
}

type AccessIdentityProviderGetResponseResultAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccessIdentityProviderGetResponseResultAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                      `json:"sso_target_url"`
	JSON         accessIdentityProviderGetResponseResultAccessSamlConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessSamlConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessSamlConfig]
type accessIdentityProviderGetResponseResultAccessSamlConfigJSON struct {
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

func (r *AccessIdentityProviderGetResponseResultAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderGetResponseResultAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                     `json:"header_name"`
	JSON       accessIdentityProviderGetResponseResultAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessSamlConfigHeaderAttribute]
type accessIdentityProviderGetResponseResultAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseResultAccessSamlScimConfig struct {
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
	UserDeprovision bool                                                            `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseResultAccessSamlScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessSamlScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessSamlScimConfig]
type accessIdentityProviderGetResponseResultAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseResultAccessSamlType string

const (
	AccessIdentityProviderGetResponseResultAccessSamlTypeOnetimepin AccessIdentityProviderGetResponseResultAccessSamlType = "onetimepin"
	AccessIdentityProviderGetResponseResultAccessSamlTypeAzureAd    AccessIdentityProviderGetResponseResultAccessSamlType = "azureAD"
	AccessIdentityProviderGetResponseResultAccessSamlTypeSaml       AccessIdentityProviderGetResponseResultAccessSamlType = "saml"
	AccessIdentityProviderGetResponseResultAccessSamlTypeCentrify   AccessIdentityProviderGetResponseResultAccessSamlType = "centrify"
	AccessIdentityProviderGetResponseResultAccessSamlTypeFacebook   AccessIdentityProviderGetResponseResultAccessSamlType = "facebook"
	AccessIdentityProviderGetResponseResultAccessSamlTypeGitHub     AccessIdentityProviderGetResponseResultAccessSamlType = "github"
	AccessIdentityProviderGetResponseResultAccessSamlTypeGoogleApps AccessIdentityProviderGetResponseResultAccessSamlType = "google-apps"
	AccessIdentityProviderGetResponseResultAccessSamlTypeGoogle     AccessIdentityProviderGetResponseResultAccessSamlType = "google"
	AccessIdentityProviderGetResponseResultAccessSamlTypeLinkedin   AccessIdentityProviderGetResponseResultAccessSamlType = "linkedin"
	AccessIdentityProviderGetResponseResultAccessSamlTypeOidc       AccessIdentityProviderGetResponseResultAccessSamlType = "oidc"
	AccessIdentityProviderGetResponseResultAccessSamlTypeOkta       AccessIdentityProviderGetResponseResultAccessSamlType = "okta"
	AccessIdentityProviderGetResponseResultAccessSamlTypeOnelogin   AccessIdentityProviderGetResponseResultAccessSamlType = "onelogin"
	AccessIdentityProviderGetResponseResultAccessSamlTypePingone    AccessIdentityProviderGetResponseResultAccessSamlType = "pingone"
	AccessIdentityProviderGetResponseResultAccessSamlTypeYandex     AccessIdentityProviderGetResponseResultAccessSamlType = "yandex"
)

type AccessIdentityProviderGetResponseResultAccessYandex struct {
	// UUID
	ID     string                                                    `json:"id"`
	Config AccessIdentityProviderGetResponseResultAccessYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseResultAccessYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseResultAccessYandexType `json:"type"`
	JSON accessIdentityProviderGetResponseResultAccessYandexJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessYandexJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseResultAccessYandex]
type accessIdentityProviderGetResponseResultAccessYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseResultAccessYandex) implementsAccessIdentityProviderGetResponseResult() {
}

type AccessIdentityProviderGetResponseResultAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                        `json:"client_secret"`
	JSON         accessIdentityProviderGetResponseResultAccessYandexConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessYandexConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessYandexConfig]
type accessIdentityProviderGetResponseResultAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseResultAccessYandexScimConfig struct {
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
	UserDeprovision bool                                                              `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseResultAccessYandexScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessYandexScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessYandexScimConfig]
type accessIdentityProviderGetResponseResultAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseResultAccessYandexType string

const (
	AccessIdentityProviderGetResponseResultAccessYandexTypeOnetimepin AccessIdentityProviderGetResponseResultAccessYandexType = "onetimepin"
	AccessIdentityProviderGetResponseResultAccessYandexTypeAzureAd    AccessIdentityProviderGetResponseResultAccessYandexType = "azureAD"
	AccessIdentityProviderGetResponseResultAccessYandexTypeSaml       AccessIdentityProviderGetResponseResultAccessYandexType = "saml"
	AccessIdentityProviderGetResponseResultAccessYandexTypeCentrify   AccessIdentityProviderGetResponseResultAccessYandexType = "centrify"
	AccessIdentityProviderGetResponseResultAccessYandexTypeFacebook   AccessIdentityProviderGetResponseResultAccessYandexType = "facebook"
	AccessIdentityProviderGetResponseResultAccessYandexTypeGitHub     AccessIdentityProviderGetResponseResultAccessYandexType = "github"
	AccessIdentityProviderGetResponseResultAccessYandexTypeGoogleApps AccessIdentityProviderGetResponseResultAccessYandexType = "google-apps"
	AccessIdentityProviderGetResponseResultAccessYandexTypeGoogle     AccessIdentityProviderGetResponseResultAccessYandexType = "google"
	AccessIdentityProviderGetResponseResultAccessYandexTypeLinkedin   AccessIdentityProviderGetResponseResultAccessYandexType = "linkedin"
	AccessIdentityProviderGetResponseResultAccessYandexTypeOidc       AccessIdentityProviderGetResponseResultAccessYandexType = "oidc"
	AccessIdentityProviderGetResponseResultAccessYandexTypeOkta       AccessIdentityProviderGetResponseResultAccessYandexType = "okta"
	AccessIdentityProviderGetResponseResultAccessYandexTypeOnelogin   AccessIdentityProviderGetResponseResultAccessYandexType = "onelogin"
	AccessIdentityProviderGetResponseResultAccessYandexTypePingone    AccessIdentityProviderGetResponseResultAccessYandexType = "pingone"
	AccessIdentityProviderGetResponseResultAccessYandexTypeYandex     AccessIdentityProviderGetResponseResultAccessYandexType = "yandex"
)

type AccessIdentityProviderGetResponseResultAccessOnetimepin struct {
	// UUID
	ID     string      `json:"id"`
	Config interface{} `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseResultAccessOnetimepinScimConfig `json:"scim_config"`
	Type       AccessIdentityProviderGetResponseResultAccessOnetimepinType       `json:"type"`
	JSON       accessIdentityProviderGetResponseResultAccessOnetimepinJSON       `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessOnetimepinJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessOnetimepin]
type accessIdentityProviderGetResponseResultAccessOnetimepinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseResultAccessOnetimepin) implementsAccessIdentityProviderGetResponseResult() {
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseResultAccessOnetimepinScimConfig struct {
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
	UserDeprovision bool                                                                  `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseResultAccessOnetimepinScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseResultAccessOnetimepinScimConfigJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderGetResponseResultAccessOnetimepinScimConfig]
type accessIdentityProviderGetResponseResultAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseResultAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderGetResponseResultAccessOnetimepinType string

const (
	AccessIdentityProviderGetResponseResultAccessOnetimepinTypeOnetimepin AccessIdentityProviderGetResponseResultAccessOnetimepinType = "onetimepin"
)

// Whether the API call was successful
type AccessIdentityProviderGetResponseSuccess bool

const (
	AccessIdentityProviderGetResponseSuccessTrue AccessIdentityProviderGetResponseSuccess = true
)

type AccessIdentityProviderUpdateResponse struct {
	Errors   []AccessIdentityProviderUpdateResponseError   `json:"errors"`
	Messages []AccessIdentityProviderUpdateResponseMessage `json:"messages"`
	Result   AccessIdentityProviderUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessIdentityProviderUpdateResponseSuccess `json:"success"`
	JSON    accessIdentityProviderUpdateResponseJSON    `json:"-"`
}

// accessIdentityProviderUpdateResponseJSON contains the JSON metadata for the
// struct [AccessIdentityProviderUpdateResponse]
type accessIdentityProviderUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderUpdateResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accessIdentityProviderUpdateResponseErrorJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccessIdentityProviderUpdateResponseError]
type accessIdentityProviderUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderUpdateResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accessIdentityProviderUpdateResponseMessageJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccessIdentityProviderUpdateResponseMessage]
type accessIdentityProviderUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccessIdentityProviderUpdateResponseResultAccessAzureAd],
// [AccessIdentityProviderUpdateResponseResultAccessCentrify],
// [AccessIdentityProviderUpdateResponseResultAccessFacebook],
// [AccessIdentityProviderUpdateResponseResultAccessGitHub],
// [AccessIdentityProviderUpdateResponseResultAccessGoogle],
// [AccessIdentityProviderUpdateResponseResultAccessGoogleApps],
// [AccessIdentityProviderUpdateResponseResultAccessLinkedin],
// [AccessIdentityProviderUpdateResponseResultAccessOidc],
// [AccessIdentityProviderUpdateResponseResultAccessOkta],
// [AccessIdentityProviderUpdateResponseResultAccessOnelogin],
// [AccessIdentityProviderUpdateResponseResultAccessPingone],
// [AccessIdentityProviderUpdateResponseResultAccessSaml],
// [AccessIdentityProviderUpdateResponseResultAccessYandex] or
// [AccessIdentityProviderUpdateResponseResultAccessOnetimepin].
type AccessIdentityProviderUpdateResponseResult interface {
	implementsAccessIdentityProviderUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessIdentityProviderUpdateResponseResult)(nil)).Elem(), "")
}

type AccessIdentityProviderUpdateResponseResultAccessAzureAd struct {
	// UUID
	ID     string                                                        `json:"id"`
	Config AccessIdentityProviderUpdateResponseResultAccessAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseResultAccessAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseResultAccessAzureAdType `json:"type"`
	JSON accessIdentityProviderUpdateResponseResultAccessAzureAdJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessAzureAdJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessAzureAd]
type accessIdentityProviderUpdateResponseResultAccessAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseResultAccessAzureAd) implementsAccessIdentityProviderUpdateResponseResult() {
}

type AccessIdentityProviderUpdateResponseResultAccessAzureAdConfig struct {
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
	SupportGroups bool                                                              `json:"support_groups"`
	JSON          accessIdentityProviderUpdateResponseResultAccessAzureAdConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessAzureAdConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessAzureAdConfig]
type accessIdentityProviderUpdateResponseResultAccessAzureAdConfigJSON struct {
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

func (r *AccessIdentityProviderUpdateResponseResultAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseResultAccessAzureAdScimConfig struct {
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
	UserDeprovision bool                                                                  `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseResultAccessAzureAdScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessAzureAdScimConfigJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessAzureAdScimConfig]
type accessIdentityProviderUpdateResponseResultAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseResultAccessAzureAdType string

const (
	AccessIdentityProviderUpdateResponseResultAccessAzureAdTypeOnetimepin AccessIdentityProviderUpdateResponseResultAccessAzureAdType = "onetimepin"
	AccessIdentityProviderUpdateResponseResultAccessAzureAdTypeAzureAd    AccessIdentityProviderUpdateResponseResultAccessAzureAdType = "azureAD"
	AccessIdentityProviderUpdateResponseResultAccessAzureAdTypeSaml       AccessIdentityProviderUpdateResponseResultAccessAzureAdType = "saml"
	AccessIdentityProviderUpdateResponseResultAccessAzureAdTypeCentrify   AccessIdentityProviderUpdateResponseResultAccessAzureAdType = "centrify"
	AccessIdentityProviderUpdateResponseResultAccessAzureAdTypeFacebook   AccessIdentityProviderUpdateResponseResultAccessAzureAdType = "facebook"
	AccessIdentityProviderUpdateResponseResultAccessAzureAdTypeGitHub     AccessIdentityProviderUpdateResponseResultAccessAzureAdType = "github"
	AccessIdentityProviderUpdateResponseResultAccessAzureAdTypeGoogleApps AccessIdentityProviderUpdateResponseResultAccessAzureAdType = "google-apps"
	AccessIdentityProviderUpdateResponseResultAccessAzureAdTypeGoogle     AccessIdentityProviderUpdateResponseResultAccessAzureAdType = "google"
	AccessIdentityProviderUpdateResponseResultAccessAzureAdTypeLinkedin   AccessIdentityProviderUpdateResponseResultAccessAzureAdType = "linkedin"
	AccessIdentityProviderUpdateResponseResultAccessAzureAdTypeOidc       AccessIdentityProviderUpdateResponseResultAccessAzureAdType = "oidc"
	AccessIdentityProviderUpdateResponseResultAccessAzureAdTypeOkta       AccessIdentityProviderUpdateResponseResultAccessAzureAdType = "okta"
	AccessIdentityProviderUpdateResponseResultAccessAzureAdTypeOnelogin   AccessIdentityProviderUpdateResponseResultAccessAzureAdType = "onelogin"
	AccessIdentityProviderUpdateResponseResultAccessAzureAdTypePingone    AccessIdentityProviderUpdateResponseResultAccessAzureAdType = "pingone"
	AccessIdentityProviderUpdateResponseResultAccessAzureAdTypeYandex     AccessIdentityProviderUpdateResponseResultAccessAzureAdType = "yandex"
)

type AccessIdentityProviderUpdateResponseResultAccessCentrify struct {
	// UUID
	ID     string                                                         `json:"id"`
	Config AccessIdentityProviderUpdateResponseResultAccessCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseResultAccessCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseResultAccessCentrifyType `json:"type"`
	JSON accessIdentityProviderUpdateResponseResultAccessCentrifyJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessCentrifyJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessCentrify]
type accessIdentityProviderUpdateResponseResultAccessCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseResultAccessCentrify) implementsAccessIdentityProviderUpdateResponseResult() {
}

type AccessIdentityProviderUpdateResponseResultAccessCentrifyConfig struct {
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
	EmailClaimName string                                                             `json:"email_claim_name"`
	JSON           accessIdentityProviderUpdateResponseResultAccessCentrifyConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessCentrifyConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessCentrifyConfig]
type accessIdentityProviderUpdateResponseResultAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseResultAccessCentrifyScimConfig struct {
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
	UserDeprovision bool                                                                   `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseResultAccessCentrifyScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessCentrifyScimConfigJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessCentrifyScimConfig]
type accessIdentityProviderUpdateResponseResultAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseResultAccessCentrifyType string

const (
	AccessIdentityProviderUpdateResponseResultAccessCentrifyTypeOnetimepin AccessIdentityProviderUpdateResponseResultAccessCentrifyType = "onetimepin"
	AccessIdentityProviderUpdateResponseResultAccessCentrifyTypeAzureAd    AccessIdentityProviderUpdateResponseResultAccessCentrifyType = "azureAD"
	AccessIdentityProviderUpdateResponseResultAccessCentrifyTypeSaml       AccessIdentityProviderUpdateResponseResultAccessCentrifyType = "saml"
	AccessIdentityProviderUpdateResponseResultAccessCentrifyTypeCentrify   AccessIdentityProviderUpdateResponseResultAccessCentrifyType = "centrify"
	AccessIdentityProviderUpdateResponseResultAccessCentrifyTypeFacebook   AccessIdentityProviderUpdateResponseResultAccessCentrifyType = "facebook"
	AccessIdentityProviderUpdateResponseResultAccessCentrifyTypeGitHub     AccessIdentityProviderUpdateResponseResultAccessCentrifyType = "github"
	AccessIdentityProviderUpdateResponseResultAccessCentrifyTypeGoogleApps AccessIdentityProviderUpdateResponseResultAccessCentrifyType = "google-apps"
	AccessIdentityProviderUpdateResponseResultAccessCentrifyTypeGoogle     AccessIdentityProviderUpdateResponseResultAccessCentrifyType = "google"
	AccessIdentityProviderUpdateResponseResultAccessCentrifyTypeLinkedin   AccessIdentityProviderUpdateResponseResultAccessCentrifyType = "linkedin"
	AccessIdentityProviderUpdateResponseResultAccessCentrifyTypeOidc       AccessIdentityProviderUpdateResponseResultAccessCentrifyType = "oidc"
	AccessIdentityProviderUpdateResponseResultAccessCentrifyTypeOkta       AccessIdentityProviderUpdateResponseResultAccessCentrifyType = "okta"
	AccessIdentityProviderUpdateResponseResultAccessCentrifyTypeOnelogin   AccessIdentityProviderUpdateResponseResultAccessCentrifyType = "onelogin"
	AccessIdentityProviderUpdateResponseResultAccessCentrifyTypePingone    AccessIdentityProviderUpdateResponseResultAccessCentrifyType = "pingone"
	AccessIdentityProviderUpdateResponseResultAccessCentrifyTypeYandex     AccessIdentityProviderUpdateResponseResultAccessCentrifyType = "yandex"
)

type AccessIdentityProviderUpdateResponseResultAccessFacebook struct {
	// UUID
	ID     string                                                         `json:"id"`
	Config AccessIdentityProviderUpdateResponseResultAccessFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseResultAccessFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseResultAccessFacebookType `json:"type"`
	JSON accessIdentityProviderUpdateResponseResultAccessFacebookJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessFacebookJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessFacebook]
type accessIdentityProviderUpdateResponseResultAccessFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseResultAccessFacebook) implementsAccessIdentityProviderUpdateResponseResult() {
}

type AccessIdentityProviderUpdateResponseResultAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                             `json:"client_secret"`
	JSON         accessIdentityProviderUpdateResponseResultAccessFacebookConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessFacebookConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessFacebookConfig]
type accessIdentityProviderUpdateResponseResultAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseResultAccessFacebookScimConfig struct {
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
	UserDeprovision bool                                                                   `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseResultAccessFacebookScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessFacebookScimConfigJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessFacebookScimConfig]
type accessIdentityProviderUpdateResponseResultAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseResultAccessFacebookType string

const (
	AccessIdentityProviderUpdateResponseResultAccessFacebookTypeOnetimepin AccessIdentityProviderUpdateResponseResultAccessFacebookType = "onetimepin"
	AccessIdentityProviderUpdateResponseResultAccessFacebookTypeAzureAd    AccessIdentityProviderUpdateResponseResultAccessFacebookType = "azureAD"
	AccessIdentityProviderUpdateResponseResultAccessFacebookTypeSaml       AccessIdentityProviderUpdateResponseResultAccessFacebookType = "saml"
	AccessIdentityProviderUpdateResponseResultAccessFacebookTypeCentrify   AccessIdentityProviderUpdateResponseResultAccessFacebookType = "centrify"
	AccessIdentityProviderUpdateResponseResultAccessFacebookTypeFacebook   AccessIdentityProviderUpdateResponseResultAccessFacebookType = "facebook"
	AccessIdentityProviderUpdateResponseResultAccessFacebookTypeGitHub     AccessIdentityProviderUpdateResponseResultAccessFacebookType = "github"
	AccessIdentityProviderUpdateResponseResultAccessFacebookTypeGoogleApps AccessIdentityProviderUpdateResponseResultAccessFacebookType = "google-apps"
	AccessIdentityProviderUpdateResponseResultAccessFacebookTypeGoogle     AccessIdentityProviderUpdateResponseResultAccessFacebookType = "google"
	AccessIdentityProviderUpdateResponseResultAccessFacebookTypeLinkedin   AccessIdentityProviderUpdateResponseResultAccessFacebookType = "linkedin"
	AccessIdentityProviderUpdateResponseResultAccessFacebookTypeOidc       AccessIdentityProviderUpdateResponseResultAccessFacebookType = "oidc"
	AccessIdentityProviderUpdateResponseResultAccessFacebookTypeOkta       AccessIdentityProviderUpdateResponseResultAccessFacebookType = "okta"
	AccessIdentityProviderUpdateResponseResultAccessFacebookTypeOnelogin   AccessIdentityProviderUpdateResponseResultAccessFacebookType = "onelogin"
	AccessIdentityProviderUpdateResponseResultAccessFacebookTypePingone    AccessIdentityProviderUpdateResponseResultAccessFacebookType = "pingone"
	AccessIdentityProviderUpdateResponseResultAccessFacebookTypeYandex     AccessIdentityProviderUpdateResponseResultAccessFacebookType = "yandex"
)

type AccessIdentityProviderUpdateResponseResultAccessGitHub struct {
	// UUID
	ID     string                                                       `json:"id"`
	Config AccessIdentityProviderUpdateResponseResultAccessGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseResultAccessGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseResultAccessGitHubType `json:"type"`
	JSON accessIdentityProviderUpdateResponseResultAccessGitHubJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessGitHubJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseResultAccessGitHub]
type accessIdentityProviderUpdateResponseResultAccessGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseResultAccessGitHub) implementsAccessIdentityProviderUpdateResponseResult() {
}

type AccessIdentityProviderUpdateResponseResultAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                           `json:"client_secret"`
	JSON         accessIdentityProviderUpdateResponseResultAccessGitHubConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessGitHubConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessGitHubConfig]
type accessIdentityProviderUpdateResponseResultAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseResultAccessGitHubScimConfig struct {
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
	UserDeprovision bool                                                                 `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseResultAccessGitHubScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessGitHubScimConfigJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessGitHubScimConfig]
type accessIdentityProviderUpdateResponseResultAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseResultAccessGitHubType string

const (
	AccessIdentityProviderUpdateResponseResultAccessGitHubTypeOnetimepin AccessIdentityProviderUpdateResponseResultAccessGitHubType = "onetimepin"
	AccessIdentityProviderUpdateResponseResultAccessGitHubTypeAzureAd    AccessIdentityProviderUpdateResponseResultAccessGitHubType = "azureAD"
	AccessIdentityProviderUpdateResponseResultAccessGitHubTypeSaml       AccessIdentityProviderUpdateResponseResultAccessGitHubType = "saml"
	AccessIdentityProviderUpdateResponseResultAccessGitHubTypeCentrify   AccessIdentityProviderUpdateResponseResultAccessGitHubType = "centrify"
	AccessIdentityProviderUpdateResponseResultAccessGitHubTypeFacebook   AccessIdentityProviderUpdateResponseResultAccessGitHubType = "facebook"
	AccessIdentityProviderUpdateResponseResultAccessGitHubTypeGitHub     AccessIdentityProviderUpdateResponseResultAccessGitHubType = "github"
	AccessIdentityProviderUpdateResponseResultAccessGitHubTypeGoogleApps AccessIdentityProviderUpdateResponseResultAccessGitHubType = "google-apps"
	AccessIdentityProviderUpdateResponseResultAccessGitHubTypeGoogle     AccessIdentityProviderUpdateResponseResultAccessGitHubType = "google"
	AccessIdentityProviderUpdateResponseResultAccessGitHubTypeLinkedin   AccessIdentityProviderUpdateResponseResultAccessGitHubType = "linkedin"
	AccessIdentityProviderUpdateResponseResultAccessGitHubTypeOidc       AccessIdentityProviderUpdateResponseResultAccessGitHubType = "oidc"
	AccessIdentityProviderUpdateResponseResultAccessGitHubTypeOkta       AccessIdentityProviderUpdateResponseResultAccessGitHubType = "okta"
	AccessIdentityProviderUpdateResponseResultAccessGitHubTypeOnelogin   AccessIdentityProviderUpdateResponseResultAccessGitHubType = "onelogin"
	AccessIdentityProviderUpdateResponseResultAccessGitHubTypePingone    AccessIdentityProviderUpdateResponseResultAccessGitHubType = "pingone"
	AccessIdentityProviderUpdateResponseResultAccessGitHubTypeYandex     AccessIdentityProviderUpdateResponseResultAccessGitHubType = "yandex"
)

type AccessIdentityProviderUpdateResponseResultAccessGoogle struct {
	// UUID
	ID     string                                                       `json:"id"`
	Config AccessIdentityProviderUpdateResponseResultAccessGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseResultAccessGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseResultAccessGoogleType `json:"type"`
	JSON accessIdentityProviderUpdateResponseResultAccessGoogleJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessGoogleJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseResultAccessGoogle]
type accessIdentityProviderUpdateResponseResultAccessGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseResultAccessGoogle) implementsAccessIdentityProviderUpdateResponseResult() {
}

type AccessIdentityProviderUpdateResponseResultAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                           `json:"email_claim_name"`
	JSON           accessIdentityProviderUpdateResponseResultAccessGoogleConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessGoogleConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessGoogleConfig]
type accessIdentityProviderUpdateResponseResultAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseResultAccessGoogleScimConfig struct {
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
	UserDeprovision bool                                                                 `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseResultAccessGoogleScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessGoogleScimConfigJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessGoogleScimConfig]
type accessIdentityProviderUpdateResponseResultAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseResultAccessGoogleType string

const (
	AccessIdentityProviderUpdateResponseResultAccessGoogleTypeOnetimepin AccessIdentityProviderUpdateResponseResultAccessGoogleType = "onetimepin"
	AccessIdentityProviderUpdateResponseResultAccessGoogleTypeAzureAd    AccessIdentityProviderUpdateResponseResultAccessGoogleType = "azureAD"
	AccessIdentityProviderUpdateResponseResultAccessGoogleTypeSaml       AccessIdentityProviderUpdateResponseResultAccessGoogleType = "saml"
	AccessIdentityProviderUpdateResponseResultAccessGoogleTypeCentrify   AccessIdentityProviderUpdateResponseResultAccessGoogleType = "centrify"
	AccessIdentityProviderUpdateResponseResultAccessGoogleTypeFacebook   AccessIdentityProviderUpdateResponseResultAccessGoogleType = "facebook"
	AccessIdentityProviderUpdateResponseResultAccessGoogleTypeGitHub     AccessIdentityProviderUpdateResponseResultAccessGoogleType = "github"
	AccessIdentityProviderUpdateResponseResultAccessGoogleTypeGoogleApps AccessIdentityProviderUpdateResponseResultAccessGoogleType = "google-apps"
	AccessIdentityProviderUpdateResponseResultAccessGoogleTypeGoogle     AccessIdentityProviderUpdateResponseResultAccessGoogleType = "google"
	AccessIdentityProviderUpdateResponseResultAccessGoogleTypeLinkedin   AccessIdentityProviderUpdateResponseResultAccessGoogleType = "linkedin"
	AccessIdentityProviderUpdateResponseResultAccessGoogleTypeOidc       AccessIdentityProviderUpdateResponseResultAccessGoogleType = "oidc"
	AccessIdentityProviderUpdateResponseResultAccessGoogleTypeOkta       AccessIdentityProviderUpdateResponseResultAccessGoogleType = "okta"
	AccessIdentityProviderUpdateResponseResultAccessGoogleTypeOnelogin   AccessIdentityProviderUpdateResponseResultAccessGoogleType = "onelogin"
	AccessIdentityProviderUpdateResponseResultAccessGoogleTypePingone    AccessIdentityProviderUpdateResponseResultAccessGoogleType = "pingone"
	AccessIdentityProviderUpdateResponseResultAccessGoogleTypeYandex     AccessIdentityProviderUpdateResponseResultAccessGoogleType = "yandex"
)

type AccessIdentityProviderUpdateResponseResultAccessGoogleApps struct {
	// UUID
	ID     string                                                           `json:"id"`
	Config AccessIdentityProviderUpdateResponseResultAccessGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseResultAccessGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType `json:"type"`
	JSON accessIdentityProviderUpdateResponseResultAccessGoogleAppsJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessGoogleAppsJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessGoogleApps]
type accessIdentityProviderUpdateResponseResultAccessGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseResultAccessGoogleApps) implementsAccessIdentityProviderUpdateResponseResult() {
}

type AccessIdentityProviderUpdateResponseResultAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                               `json:"email_claim_name"`
	JSON           accessIdentityProviderUpdateResponseResultAccessGoogleAppsConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessGoogleAppsConfigJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessGoogleAppsConfig]
type accessIdentityProviderUpdateResponseResultAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseResultAccessGoogleAppsScimConfig struct {
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
	JSON            accessIdentityProviderUpdateResponseResultAccessGoogleAppsScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessGoogleAppsScimConfig]
type accessIdentityProviderUpdateResponseResultAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType string

const (
	AccessIdentityProviderUpdateResponseResultAccessGoogleAppsTypeOnetimepin AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType = "onetimepin"
	AccessIdentityProviderUpdateResponseResultAccessGoogleAppsTypeAzureAd    AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType = "azureAD"
	AccessIdentityProviderUpdateResponseResultAccessGoogleAppsTypeSaml       AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType = "saml"
	AccessIdentityProviderUpdateResponseResultAccessGoogleAppsTypeCentrify   AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType = "centrify"
	AccessIdentityProviderUpdateResponseResultAccessGoogleAppsTypeFacebook   AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType = "facebook"
	AccessIdentityProviderUpdateResponseResultAccessGoogleAppsTypeGitHub     AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType = "github"
	AccessIdentityProviderUpdateResponseResultAccessGoogleAppsTypeGoogleApps AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType = "google-apps"
	AccessIdentityProviderUpdateResponseResultAccessGoogleAppsTypeGoogle     AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType = "google"
	AccessIdentityProviderUpdateResponseResultAccessGoogleAppsTypeLinkedin   AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType = "linkedin"
	AccessIdentityProviderUpdateResponseResultAccessGoogleAppsTypeOidc       AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType = "oidc"
	AccessIdentityProviderUpdateResponseResultAccessGoogleAppsTypeOkta       AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType = "okta"
	AccessIdentityProviderUpdateResponseResultAccessGoogleAppsTypeOnelogin   AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType = "onelogin"
	AccessIdentityProviderUpdateResponseResultAccessGoogleAppsTypePingone    AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType = "pingone"
	AccessIdentityProviderUpdateResponseResultAccessGoogleAppsTypeYandex     AccessIdentityProviderUpdateResponseResultAccessGoogleAppsType = "yandex"
)

type AccessIdentityProviderUpdateResponseResultAccessLinkedin struct {
	// UUID
	ID     string                                                         `json:"id"`
	Config AccessIdentityProviderUpdateResponseResultAccessLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseResultAccessLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseResultAccessLinkedinType `json:"type"`
	JSON accessIdentityProviderUpdateResponseResultAccessLinkedinJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessLinkedinJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessLinkedin]
type accessIdentityProviderUpdateResponseResultAccessLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseResultAccessLinkedin) implementsAccessIdentityProviderUpdateResponseResult() {
}

type AccessIdentityProviderUpdateResponseResultAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                             `json:"client_secret"`
	JSON         accessIdentityProviderUpdateResponseResultAccessLinkedinConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessLinkedinConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessLinkedinConfig]
type accessIdentityProviderUpdateResponseResultAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseResultAccessLinkedinScimConfig struct {
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
	UserDeprovision bool                                                                   `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseResultAccessLinkedinScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessLinkedinScimConfigJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessLinkedinScimConfig]
type accessIdentityProviderUpdateResponseResultAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseResultAccessLinkedinType string

const (
	AccessIdentityProviderUpdateResponseResultAccessLinkedinTypeOnetimepin AccessIdentityProviderUpdateResponseResultAccessLinkedinType = "onetimepin"
	AccessIdentityProviderUpdateResponseResultAccessLinkedinTypeAzureAd    AccessIdentityProviderUpdateResponseResultAccessLinkedinType = "azureAD"
	AccessIdentityProviderUpdateResponseResultAccessLinkedinTypeSaml       AccessIdentityProviderUpdateResponseResultAccessLinkedinType = "saml"
	AccessIdentityProviderUpdateResponseResultAccessLinkedinTypeCentrify   AccessIdentityProviderUpdateResponseResultAccessLinkedinType = "centrify"
	AccessIdentityProviderUpdateResponseResultAccessLinkedinTypeFacebook   AccessIdentityProviderUpdateResponseResultAccessLinkedinType = "facebook"
	AccessIdentityProviderUpdateResponseResultAccessLinkedinTypeGitHub     AccessIdentityProviderUpdateResponseResultAccessLinkedinType = "github"
	AccessIdentityProviderUpdateResponseResultAccessLinkedinTypeGoogleApps AccessIdentityProviderUpdateResponseResultAccessLinkedinType = "google-apps"
	AccessIdentityProviderUpdateResponseResultAccessLinkedinTypeGoogle     AccessIdentityProviderUpdateResponseResultAccessLinkedinType = "google"
	AccessIdentityProviderUpdateResponseResultAccessLinkedinTypeLinkedin   AccessIdentityProviderUpdateResponseResultAccessLinkedinType = "linkedin"
	AccessIdentityProviderUpdateResponseResultAccessLinkedinTypeOidc       AccessIdentityProviderUpdateResponseResultAccessLinkedinType = "oidc"
	AccessIdentityProviderUpdateResponseResultAccessLinkedinTypeOkta       AccessIdentityProviderUpdateResponseResultAccessLinkedinType = "okta"
	AccessIdentityProviderUpdateResponseResultAccessLinkedinTypeOnelogin   AccessIdentityProviderUpdateResponseResultAccessLinkedinType = "onelogin"
	AccessIdentityProviderUpdateResponseResultAccessLinkedinTypePingone    AccessIdentityProviderUpdateResponseResultAccessLinkedinType = "pingone"
	AccessIdentityProviderUpdateResponseResultAccessLinkedinTypeYandex     AccessIdentityProviderUpdateResponseResultAccessLinkedinType = "yandex"
)

type AccessIdentityProviderUpdateResponseResultAccessOidc struct {
	// UUID
	ID     string                                                     `json:"id"`
	Config AccessIdentityProviderUpdateResponseResultAccessOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseResultAccessOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseResultAccessOidcType `json:"type"`
	JSON accessIdentityProviderUpdateResponseResultAccessOidcJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessOidcJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseResultAccessOidc]
type accessIdentityProviderUpdateResponseResultAccessOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseResultAccessOidc) implementsAccessIdentityProviderUpdateResponseResult() {
}

type AccessIdentityProviderUpdateResponseResultAccessOidcConfig struct {
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
	TokenURL string                                                         `json:"token_url"`
	JSON     accessIdentityProviderUpdateResponseResultAccessOidcConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessOidcConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessOidcConfig]
type accessIdentityProviderUpdateResponseResultAccessOidcConfigJSON struct {
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

func (r *AccessIdentityProviderUpdateResponseResultAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseResultAccessOidcScimConfig struct {
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
	UserDeprovision bool                                                               `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseResultAccessOidcScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessOidcScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessOidcScimConfig]
type accessIdentityProviderUpdateResponseResultAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseResultAccessOidcType string

const (
	AccessIdentityProviderUpdateResponseResultAccessOidcTypeOnetimepin AccessIdentityProviderUpdateResponseResultAccessOidcType = "onetimepin"
	AccessIdentityProviderUpdateResponseResultAccessOidcTypeAzureAd    AccessIdentityProviderUpdateResponseResultAccessOidcType = "azureAD"
	AccessIdentityProviderUpdateResponseResultAccessOidcTypeSaml       AccessIdentityProviderUpdateResponseResultAccessOidcType = "saml"
	AccessIdentityProviderUpdateResponseResultAccessOidcTypeCentrify   AccessIdentityProviderUpdateResponseResultAccessOidcType = "centrify"
	AccessIdentityProviderUpdateResponseResultAccessOidcTypeFacebook   AccessIdentityProviderUpdateResponseResultAccessOidcType = "facebook"
	AccessIdentityProviderUpdateResponseResultAccessOidcTypeGitHub     AccessIdentityProviderUpdateResponseResultAccessOidcType = "github"
	AccessIdentityProviderUpdateResponseResultAccessOidcTypeGoogleApps AccessIdentityProviderUpdateResponseResultAccessOidcType = "google-apps"
	AccessIdentityProviderUpdateResponseResultAccessOidcTypeGoogle     AccessIdentityProviderUpdateResponseResultAccessOidcType = "google"
	AccessIdentityProviderUpdateResponseResultAccessOidcTypeLinkedin   AccessIdentityProviderUpdateResponseResultAccessOidcType = "linkedin"
	AccessIdentityProviderUpdateResponseResultAccessOidcTypeOidc       AccessIdentityProviderUpdateResponseResultAccessOidcType = "oidc"
	AccessIdentityProviderUpdateResponseResultAccessOidcTypeOkta       AccessIdentityProviderUpdateResponseResultAccessOidcType = "okta"
	AccessIdentityProviderUpdateResponseResultAccessOidcTypeOnelogin   AccessIdentityProviderUpdateResponseResultAccessOidcType = "onelogin"
	AccessIdentityProviderUpdateResponseResultAccessOidcTypePingone    AccessIdentityProviderUpdateResponseResultAccessOidcType = "pingone"
	AccessIdentityProviderUpdateResponseResultAccessOidcTypeYandex     AccessIdentityProviderUpdateResponseResultAccessOidcType = "yandex"
)

type AccessIdentityProviderUpdateResponseResultAccessOkta struct {
	// UUID
	ID     string                                                     `json:"id"`
	Config AccessIdentityProviderUpdateResponseResultAccessOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseResultAccessOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseResultAccessOktaType `json:"type"`
	JSON accessIdentityProviderUpdateResponseResultAccessOktaJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessOktaJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseResultAccessOkta]
type accessIdentityProviderUpdateResponseResultAccessOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseResultAccessOkta) implementsAccessIdentityProviderUpdateResponseResult() {
}

type AccessIdentityProviderUpdateResponseResultAccessOktaConfig struct {
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
	OktaAccount string                                                         `json:"okta_account"`
	JSON        accessIdentityProviderUpdateResponseResultAccessOktaConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessOktaConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessOktaConfig]
type accessIdentityProviderUpdateResponseResultAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseResultAccessOktaScimConfig struct {
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
	UserDeprovision bool                                                               `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseResultAccessOktaScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessOktaScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessOktaScimConfig]
type accessIdentityProviderUpdateResponseResultAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseResultAccessOktaType string

const (
	AccessIdentityProviderUpdateResponseResultAccessOktaTypeOnetimepin AccessIdentityProviderUpdateResponseResultAccessOktaType = "onetimepin"
	AccessIdentityProviderUpdateResponseResultAccessOktaTypeAzureAd    AccessIdentityProviderUpdateResponseResultAccessOktaType = "azureAD"
	AccessIdentityProviderUpdateResponseResultAccessOktaTypeSaml       AccessIdentityProviderUpdateResponseResultAccessOktaType = "saml"
	AccessIdentityProviderUpdateResponseResultAccessOktaTypeCentrify   AccessIdentityProviderUpdateResponseResultAccessOktaType = "centrify"
	AccessIdentityProviderUpdateResponseResultAccessOktaTypeFacebook   AccessIdentityProviderUpdateResponseResultAccessOktaType = "facebook"
	AccessIdentityProviderUpdateResponseResultAccessOktaTypeGitHub     AccessIdentityProviderUpdateResponseResultAccessOktaType = "github"
	AccessIdentityProviderUpdateResponseResultAccessOktaTypeGoogleApps AccessIdentityProviderUpdateResponseResultAccessOktaType = "google-apps"
	AccessIdentityProviderUpdateResponseResultAccessOktaTypeGoogle     AccessIdentityProviderUpdateResponseResultAccessOktaType = "google"
	AccessIdentityProviderUpdateResponseResultAccessOktaTypeLinkedin   AccessIdentityProviderUpdateResponseResultAccessOktaType = "linkedin"
	AccessIdentityProviderUpdateResponseResultAccessOktaTypeOidc       AccessIdentityProviderUpdateResponseResultAccessOktaType = "oidc"
	AccessIdentityProviderUpdateResponseResultAccessOktaTypeOkta       AccessIdentityProviderUpdateResponseResultAccessOktaType = "okta"
	AccessIdentityProviderUpdateResponseResultAccessOktaTypeOnelogin   AccessIdentityProviderUpdateResponseResultAccessOktaType = "onelogin"
	AccessIdentityProviderUpdateResponseResultAccessOktaTypePingone    AccessIdentityProviderUpdateResponseResultAccessOktaType = "pingone"
	AccessIdentityProviderUpdateResponseResultAccessOktaTypeYandex     AccessIdentityProviderUpdateResponseResultAccessOktaType = "yandex"
)

type AccessIdentityProviderUpdateResponseResultAccessOnelogin struct {
	// UUID
	ID     string                                                         `json:"id"`
	Config AccessIdentityProviderUpdateResponseResultAccessOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseResultAccessOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseResultAccessOneloginType `json:"type"`
	JSON accessIdentityProviderUpdateResponseResultAccessOneloginJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessOneloginJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessOnelogin]
type accessIdentityProviderUpdateResponseResultAccessOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseResultAccessOnelogin) implementsAccessIdentityProviderUpdateResponseResult() {
}

type AccessIdentityProviderUpdateResponseResultAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                             `json:"onelogin_account"`
	JSON            accessIdentityProviderUpdateResponseResultAccessOneloginConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessOneloginConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessOneloginConfig]
type accessIdentityProviderUpdateResponseResultAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseResultAccessOneloginScimConfig struct {
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
	UserDeprovision bool                                                                   `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseResultAccessOneloginScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessOneloginScimConfigJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessOneloginScimConfig]
type accessIdentityProviderUpdateResponseResultAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseResultAccessOneloginType string

const (
	AccessIdentityProviderUpdateResponseResultAccessOneloginTypeOnetimepin AccessIdentityProviderUpdateResponseResultAccessOneloginType = "onetimepin"
	AccessIdentityProviderUpdateResponseResultAccessOneloginTypeAzureAd    AccessIdentityProviderUpdateResponseResultAccessOneloginType = "azureAD"
	AccessIdentityProviderUpdateResponseResultAccessOneloginTypeSaml       AccessIdentityProviderUpdateResponseResultAccessOneloginType = "saml"
	AccessIdentityProviderUpdateResponseResultAccessOneloginTypeCentrify   AccessIdentityProviderUpdateResponseResultAccessOneloginType = "centrify"
	AccessIdentityProviderUpdateResponseResultAccessOneloginTypeFacebook   AccessIdentityProviderUpdateResponseResultAccessOneloginType = "facebook"
	AccessIdentityProviderUpdateResponseResultAccessOneloginTypeGitHub     AccessIdentityProviderUpdateResponseResultAccessOneloginType = "github"
	AccessIdentityProviderUpdateResponseResultAccessOneloginTypeGoogleApps AccessIdentityProviderUpdateResponseResultAccessOneloginType = "google-apps"
	AccessIdentityProviderUpdateResponseResultAccessOneloginTypeGoogle     AccessIdentityProviderUpdateResponseResultAccessOneloginType = "google"
	AccessIdentityProviderUpdateResponseResultAccessOneloginTypeLinkedin   AccessIdentityProviderUpdateResponseResultAccessOneloginType = "linkedin"
	AccessIdentityProviderUpdateResponseResultAccessOneloginTypeOidc       AccessIdentityProviderUpdateResponseResultAccessOneloginType = "oidc"
	AccessIdentityProviderUpdateResponseResultAccessOneloginTypeOkta       AccessIdentityProviderUpdateResponseResultAccessOneloginType = "okta"
	AccessIdentityProviderUpdateResponseResultAccessOneloginTypeOnelogin   AccessIdentityProviderUpdateResponseResultAccessOneloginType = "onelogin"
	AccessIdentityProviderUpdateResponseResultAccessOneloginTypePingone    AccessIdentityProviderUpdateResponseResultAccessOneloginType = "pingone"
	AccessIdentityProviderUpdateResponseResultAccessOneloginTypeYandex     AccessIdentityProviderUpdateResponseResultAccessOneloginType = "yandex"
)

type AccessIdentityProviderUpdateResponseResultAccessPingone struct {
	// UUID
	ID     string                                                        `json:"id"`
	Config AccessIdentityProviderUpdateResponseResultAccessPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseResultAccessPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseResultAccessPingoneType `json:"type"`
	JSON accessIdentityProviderUpdateResponseResultAccessPingoneJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessPingoneJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessPingone]
type accessIdentityProviderUpdateResponseResultAccessPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseResultAccessPingone) implementsAccessIdentityProviderUpdateResponseResult() {
}

type AccessIdentityProviderUpdateResponseResultAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                            `json:"ping_env_id"`
	JSON      accessIdentityProviderUpdateResponseResultAccessPingoneConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessPingoneConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessPingoneConfig]
type accessIdentityProviderUpdateResponseResultAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseResultAccessPingoneScimConfig struct {
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
	UserDeprovision bool                                                                  `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseResultAccessPingoneScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessPingoneScimConfigJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessPingoneScimConfig]
type accessIdentityProviderUpdateResponseResultAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseResultAccessPingoneType string

const (
	AccessIdentityProviderUpdateResponseResultAccessPingoneTypeOnetimepin AccessIdentityProviderUpdateResponseResultAccessPingoneType = "onetimepin"
	AccessIdentityProviderUpdateResponseResultAccessPingoneTypeAzureAd    AccessIdentityProviderUpdateResponseResultAccessPingoneType = "azureAD"
	AccessIdentityProviderUpdateResponseResultAccessPingoneTypeSaml       AccessIdentityProviderUpdateResponseResultAccessPingoneType = "saml"
	AccessIdentityProviderUpdateResponseResultAccessPingoneTypeCentrify   AccessIdentityProviderUpdateResponseResultAccessPingoneType = "centrify"
	AccessIdentityProviderUpdateResponseResultAccessPingoneTypeFacebook   AccessIdentityProviderUpdateResponseResultAccessPingoneType = "facebook"
	AccessIdentityProviderUpdateResponseResultAccessPingoneTypeGitHub     AccessIdentityProviderUpdateResponseResultAccessPingoneType = "github"
	AccessIdentityProviderUpdateResponseResultAccessPingoneTypeGoogleApps AccessIdentityProviderUpdateResponseResultAccessPingoneType = "google-apps"
	AccessIdentityProviderUpdateResponseResultAccessPingoneTypeGoogle     AccessIdentityProviderUpdateResponseResultAccessPingoneType = "google"
	AccessIdentityProviderUpdateResponseResultAccessPingoneTypeLinkedin   AccessIdentityProviderUpdateResponseResultAccessPingoneType = "linkedin"
	AccessIdentityProviderUpdateResponseResultAccessPingoneTypeOidc       AccessIdentityProviderUpdateResponseResultAccessPingoneType = "oidc"
	AccessIdentityProviderUpdateResponseResultAccessPingoneTypeOkta       AccessIdentityProviderUpdateResponseResultAccessPingoneType = "okta"
	AccessIdentityProviderUpdateResponseResultAccessPingoneTypeOnelogin   AccessIdentityProviderUpdateResponseResultAccessPingoneType = "onelogin"
	AccessIdentityProviderUpdateResponseResultAccessPingoneTypePingone    AccessIdentityProviderUpdateResponseResultAccessPingoneType = "pingone"
	AccessIdentityProviderUpdateResponseResultAccessPingoneTypeYandex     AccessIdentityProviderUpdateResponseResultAccessPingoneType = "yandex"
)

type AccessIdentityProviderUpdateResponseResultAccessSaml struct {
	// UUID
	ID     string                                                     `json:"id"`
	Config AccessIdentityProviderUpdateResponseResultAccessSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseResultAccessSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseResultAccessSamlType `json:"type"`
	JSON accessIdentityProviderUpdateResponseResultAccessSamlJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessSamlJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseResultAccessSaml]
type accessIdentityProviderUpdateResponseResultAccessSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseResultAccessSaml) implementsAccessIdentityProviderUpdateResponseResult() {
}

type AccessIdentityProviderUpdateResponseResultAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccessIdentityProviderUpdateResponseResultAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                         `json:"sso_target_url"`
	JSON         accessIdentityProviderUpdateResponseResultAccessSamlConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessSamlConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessSamlConfig]
type accessIdentityProviderUpdateResponseResultAccessSamlConfigJSON struct {
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

func (r *AccessIdentityProviderUpdateResponseResultAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderUpdateResponseResultAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                        `json:"header_name"`
	JSON       accessIdentityProviderUpdateResponseResultAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessSamlConfigHeaderAttribute]
type accessIdentityProviderUpdateResponseResultAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseResultAccessSamlScimConfig struct {
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
	UserDeprovision bool                                                               `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseResultAccessSamlScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessSamlScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessSamlScimConfig]
type accessIdentityProviderUpdateResponseResultAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseResultAccessSamlType string

const (
	AccessIdentityProviderUpdateResponseResultAccessSamlTypeOnetimepin AccessIdentityProviderUpdateResponseResultAccessSamlType = "onetimepin"
	AccessIdentityProviderUpdateResponseResultAccessSamlTypeAzureAd    AccessIdentityProviderUpdateResponseResultAccessSamlType = "azureAD"
	AccessIdentityProviderUpdateResponseResultAccessSamlTypeSaml       AccessIdentityProviderUpdateResponseResultAccessSamlType = "saml"
	AccessIdentityProviderUpdateResponseResultAccessSamlTypeCentrify   AccessIdentityProviderUpdateResponseResultAccessSamlType = "centrify"
	AccessIdentityProviderUpdateResponseResultAccessSamlTypeFacebook   AccessIdentityProviderUpdateResponseResultAccessSamlType = "facebook"
	AccessIdentityProviderUpdateResponseResultAccessSamlTypeGitHub     AccessIdentityProviderUpdateResponseResultAccessSamlType = "github"
	AccessIdentityProviderUpdateResponseResultAccessSamlTypeGoogleApps AccessIdentityProviderUpdateResponseResultAccessSamlType = "google-apps"
	AccessIdentityProviderUpdateResponseResultAccessSamlTypeGoogle     AccessIdentityProviderUpdateResponseResultAccessSamlType = "google"
	AccessIdentityProviderUpdateResponseResultAccessSamlTypeLinkedin   AccessIdentityProviderUpdateResponseResultAccessSamlType = "linkedin"
	AccessIdentityProviderUpdateResponseResultAccessSamlTypeOidc       AccessIdentityProviderUpdateResponseResultAccessSamlType = "oidc"
	AccessIdentityProviderUpdateResponseResultAccessSamlTypeOkta       AccessIdentityProviderUpdateResponseResultAccessSamlType = "okta"
	AccessIdentityProviderUpdateResponseResultAccessSamlTypeOnelogin   AccessIdentityProviderUpdateResponseResultAccessSamlType = "onelogin"
	AccessIdentityProviderUpdateResponseResultAccessSamlTypePingone    AccessIdentityProviderUpdateResponseResultAccessSamlType = "pingone"
	AccessIdentityProviderUpdateResponseResultAccessSamlTypeYandex     AccessIdentityProviderUpdateResponseResultAccessSamlType = "yandex"
)

type AccessIdentityProviderUpdateResponseResultAccessYandex struct {
	// UUID
	ID     string                                                       `json:"id"`
	Config AccessIdentityProviderUpdateResponseResultAccessYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseResultAccessYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseResultAccessYandexType `json:"type"`
	JSON accessIdentityProviderUpdateResponseResultAccessYandexJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessYandexJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseResultAccessYandex]
type accessIdentityProviderUpdateResponseResultAccessYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseResultAccessYandex) implementsAccessIdentityProviderUpdateResponseResult() {
}

type AccessIdentityProviderUpdateResponseResultAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                           `json:"client_secret"`
	JSON         accessIdentityProviderUpdateResponseResultAccessYandexConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessYandexConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessYandexConfig]
type accessIdentityProviderUpdateResponseResultAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseResultAccessYandexScimConfig struct {
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
	UserDeprovision bool                                                                 `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseResultAccessYandexScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessYandexScimConfigJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessYandexScimConfig]
type accessIdentityProviderUpdateResponseResultAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseResultAccessYandexType string

const (
	AccessIdentityProviderUpdateResponseResultAccessYandexTypeOnetimepin AccessIdentityProviderUpdateResponseResultAccessYandexType = "onetimepin"
	AccessIdentityProviderUpdateResponseResultAccessYandexTypeAzureAd    AccessIdentityProviderUpdateResponseResultAccessYandexType = "azureAD"
	AccessIdentityProviderUpdateResponseResultAccessYandexTypeSaml       AccessIdentityProviderUpdateResponseResultAccessYandexType = "saml"
	AccessIdentityProviderUpdateResponseResultAccessYandexTypeCentrify   AccessIdentityProviderUpdateResponseResultAccessYandexType = "centrify"
	AccessIdentityProviderUpdateResponseResultAccessYandexTypeFacebook   AccessIdentityProviderUpdateResponseResultAccessYandexType = "facebook"
	AccessIdentityProviderUpdateResponseResultAccessYandexTypeGitHub     AccessIdentityProviderUpdateResponseResultAccessYandexType = "github"
	AccessIdentityProviderUpdateResponseResultAccessYandexTypeGoogleApps AccessIdentityProviderUpdateResponseResultAccessYandexType = "google-apps"
	AccessIdentityProviderUpdateResponseResultAccessYandexTypeGoogle     AccessIdentityProviderUpdateResponseResultAccessYandexType = "google"
	AccessIdentityProviderUpdateResponseResultAccessYandexTypeLinkedin   AccessIdentityProviderUpdateResponseResultAccessYandexType = "linkedin"
	AccessIdentityProviderUpdateResponseResultAccessYandexTypeOidc       AccessIdentityProviderUpdateResponseResultAccessYandexType = "oidc"
	AccessIdentityProviderUpdateResponseResultAccessYandexTypeOkta       AccessIdentityProviderUpdateResponseResultAccessYandexType = "okta"
	AccessIdentityProviderUpdateResponseResultAccessYandexTypeOnelogin   AccessIdentityProviderUpdateResponseResultAccessYandexType = "onelogin"
	AccessIdentityProviderUpdateResponseResultAccessYandexTypePingone    AccessIdentityProviderUpdateResponseResultAccessYandexType = "pingone"
	AccessIdentityProviderUpdateResponseResultAccessYandexTypeYandex     AccessIdentityProviderUpdateResponseResultAccessYandexType = "yandex"
)

type AccessIdentityProviderUpdateResponseResultAccessOnetimepin struct {
	// UUID
	ID     string      `json:"id"`
	Config interface{} `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseResultAccessOnetimepinScimConfig `json:"scim_config"`
	Type       AccessIdentityProviderUpdateResponseResultAccessOnetimepinType       `json:"type"`
	JSON       accessIdentityProviderUpdateResponseResultAccessOnetimepinJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessOnetimepinJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessOnetimepin]
type accessIdentityProviderUpdateResponseResultAccessOnetimepinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseResultAccessOnetimepin) implementsAccessIdentityProviderUpdateResponseResult() {
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseResultAccessOnetimepinScimConfig struct {
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
	JSON            accessIdentityProviderUpdateResponseResultAccessOnetimepinScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseResultAccessOnetimepinScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseResultAccessOnetimepinScimConfig]
type accessIdentityProviderUpdateResponseResultAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseResultAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderUpdateResponseResultAccessOnetimepinType string

const (
	AccessIdentityProviderUpdateResponseResultAccessOnetimepinTypeOnetimepin AccessIdentityProviderUpdateResponseResultAccessOnetimepinType = "onetimepin"
)

// Whether the API call was successful
type AccessIdentityProviderUpdateResponseSuccess bool

const (
	AccessIdentityProviderUpdateResponseSuccessTrue AccessIdentityProviderUpdateResponseSuccess = true
)

type AccessIdentityProviderDeleteResponse struct {
	Errors   []AccessIdentityProviderDeleteResponseError   `json:"errors"`
	Messages []AccessIdentityProviderDeleteResponseMessage `json:"messages"`
	Result   AccessIdentityProviderDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessIdentityProviderDeleteResponseSuccess `json:"success"`
	JSON    accessIdentityProviderDeleteResponseJSON    `json:"-"`
}

// accessIdentityProviderDeleteResponseJSON contains the JSON metadata for the
// struct [AccessIdentityProviderDeleteResponse]
type accessIdentityProviderDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderDeleteResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accessIdentityProviderDeleteResponseErrorJSON `json:"-"`
}

// accessIdentityProviderDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccessIdentityProviderDeleteResponseError]
type accessIdentityProviderDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderDeleteResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accessIdentityProviderDeleteResponseMessageJSON `json:"-"`
}

// accessIdentityProviderDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccessIdentityProviderDeleteResponseMessage]
type accessIdentityProviderDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderDeleteResponseResult struct {
	// UUID
	ID   string                                         `json:"id"`
	JSON accessIdentityProviderDeleteResponseResultJSON `json:"-"`
}

// accessIdentityProviderDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccessIdentityProviderDeleteResponseResult]
type accessIdentityProviderDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessIdentityProviderDeleteResponseSuccess bool

const (
	AccessIdentityProviderDeleteResponseSuccessTrue AccessIdentityProviderDeleteResponseSuccess = true
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse struct {
	Errors   []AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseError   `json:"errors"`
	Messages []AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseMessage `json:"messages"`
	Result   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseSuccess `json:"success"`
	JSON    accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseJSON    `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseError struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseErrorJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseError]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseMessage struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseMessageJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseMessage]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAd],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrify],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebook],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHub],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogle],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleApps],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedin],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidc],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOkta],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnelogin],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingone],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSaml],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandex]
// or
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepin].
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult interface {
	implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult)(nil)).Elem(), "")
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAd struct {
	// UUID
	ID     string                                                                                                    `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAd]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAd) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdConfig struct {
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
	SupportGroups bool                                                                                                          `json:"support_groups"`
	JSON          accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdConfigJSON struct {
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

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdScimConfig struct {
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
	UserDeprovision bool                                                                                                              `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessAzureAdType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrify struct {
	// UUID
	ID     string                                                                                                     `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrify]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrify) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyConfig struct {
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
	EmailClaimName string                                                                                                         `json:"email_claim_name"`
	JSON           accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyScimConfig struct {
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
	UserDeprovision bool                                                                                                               `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessCentrifyType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebook struct {
	// UUID
	ID     string                                                                                                     `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebook]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebook) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                         `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookScimConfig struct {
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
	UserDeprovision bool                                                                                                               `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessFacebookType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHub struct {
	// UUID
	ID     string                                                                                                   `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHub]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHub) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                       `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubScimConfig struct {
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
	UserDeprovision bool                                                                                                             `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGitHubType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogle struct {
	// UUID
	ID     string                                                                                                   `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogle]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogle) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                                                       `json:"email_claim_name"`
	JSON           accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleScimConfig struct {
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
	UserDeprovision bool                                                                                                             `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleApps struct {
	// UUID
	ID     string                                                                                                       `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleApps]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleApps) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                                                           `json:"email_claim_name"`
	JSON           accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsScimConfig struct {
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
	UserDeprovision bool                                                                                                                 `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessGoogleAppsType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedin struct {
	// UUID
	ID     string                                                                                                     `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedin]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedin) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                         `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinScimConfig struct {
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
	UserDeprovision bool                                                                                                               `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessLinkedinType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidc struct {
	// UUID
	ID     string                                                                                                 `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidc]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidc) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcConfig struct {
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
	TokenURL string                                                                                                     `json:"token_url"`
	JSON     accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcConfigJSON struct {
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

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcScimConfig struct {
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
	UserDeprovision bool                                                                                                           `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOidcType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOkta struct {
	// UUID
	ID     string                                                                                                 `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOkta]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOkta) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaConfig struct {
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
	OktaAccount string                                                                                                     `json:"okta_account"`
	JSON        accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaScimConfig struct {
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
	UserDeprovision bool                                                                                                           `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOktaType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnelogin struct {
	// UUID
	ID     string                                                                                                     `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnelogin]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnelogin) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                                                                         `json:"onelogin_account"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginScimConfig struct {
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
	UserDeprovision bool                                                                                                               `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOneloginType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingone struct {
	// UUID
	ID     string                                                                                                    `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingone]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingone) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                                                                        `json:"ping_env_id"`
	JSON      accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneScimConfig struct {
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
	UserDeprovision bool                                                                                                              `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessPingoneType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSaml struct {
	// UUID
	ID     string                                                                                                 `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSaml]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSaml) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                                                                     `json:"sso_target_url"`
	JSON         accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlConfigJSON struct {
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

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                                                                    `json:"header_name"`
	JSON       accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlConfigHeaderAttribute]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlScimConfig struct {
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
	UserDeprovision bool                                                                                                           `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessSamlType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandex struct {
	// UUID
	ID     string                                                                                                   `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandex]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandex) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                       `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexScimConfig struct {
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
	UserDeprovision bool                                                                                                             `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessYandexType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepin struct {
	// UUID
	ID     string      `json:"id"`
	Config interface{} `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepinScimConfig `json:"scim_config"`
	Type       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepinType       `json:"type"`
	JSON       accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepinJSON       `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepinJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepin]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepin) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResult() {
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepinScimConfig struct {
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
	UserDeprovision bool                                                                                                                 `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepinScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepinScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepinScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepinType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepinTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseResultAccessOnetimepinType = "onetimepin"
)

// Whether the API call was successful
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseSuccess bool

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseSuccessTrue AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseSuccess = true
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse struct {
	Errors     []AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseError    `json:"errors"`
	Messages   []AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseMessage  `json:"messages"`
	Result     []AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult   `json:"result"`
	ResultInfo AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseSuccess `json:"success"`
	JSON    accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseJSON    `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseError struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseErrorJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseError]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseMessage struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseMessageJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseMessage]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAd],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrify],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebook],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHub],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogle],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleApps],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedin],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidc],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOkta],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOnelogin],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingone],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSaml]
// or
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandex].
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult interface {
	implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult)(nil)).Elem(), "")
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAd struct {
	// UUID
	ID     string                                                                                                    `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAd]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAd) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdConfig struct {
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
	SupportGroups bool                                                                                                          `json:"support_groups"`
	JSON          accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdConfigJSON struct {
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

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdScimConfig struct {
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
	UserDeprovision bool                                                                                                              `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessAzureAdType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrify struct {
	// UUID
	ID     string                                                                                                     `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrify]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrify) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyConfig struct {
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
	EmailClaimName string                                                                                                         `json:"email_claim_name"`
	JSON           accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyScimConfig struct {
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
	UserDeprovision bool                                                                                                               `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessCentrifyType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebook struct {
	// UUID
	ID     string                                                                                                     `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebook]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebook) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                         `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookScimConfig struct {
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
	UserDeprovision bool                                                                                                               `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessFacebookType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHub struct {
	// UUID
	ID     string                                                                                                   `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHub]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHub) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                       `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubScimConfig struct {
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
	UserDeprovision bool                                                                                                             `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGitHubType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogle struct {
	// UUID
	ID     string                                                                                                   `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogle]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogle) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                                                       `json:"email_claim_name"`
	JSON           accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleScimConfig struct {
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
	UserDeprovision bool                                                                                                             `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleApps struct {
	// UUID
	ID     string                                                                                                       `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleApps]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleApps) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                                                           `json:"email_claim_name"`
	JSON           accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsScimConfig struct {
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
	UserDeprovision bool                                                                                                                 `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessGoogleAppsType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedin struct {
	// UUID
	ID     string                                                                                                     `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedin]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedin) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                         `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinScimConfig struct {
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
	UserDeprovision bool                                                                                                               `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessLinkedinType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidc struct {
	// UUID
	ID     string                                                                                                 `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidc]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidc) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcConfig struct {
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
	TokenURL string                                                                                                     `json:"token_url"`
	JSON     accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcConfigJSON struct {
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

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcScimConfig struct {
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
	UserDeprovision bool                                                                                                           `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOidcType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOkta struct {
	// UUID
	ID     string                                                                                                 `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOkta]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOkta) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaConfig struct {
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
	OktaAccount string                                                                                                     `json:"okta_account"`
	JSON        accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaScimConfig struct {
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
	UserDeprovision bool                                                                                                           `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOktaType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOnelogin struct {
	// UUID
	ID     string                                                                                                     `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOnelogin]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOnelogin) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                                                                         `json:"onelogin_account"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginScimConfig struct {
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
	UserDeprovision bool                                                                                                               `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessOneloginType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingone struct {
	// UUID
	ID     string                                                                                                    `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingone]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingone) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                                                                        `json:"ping_env_id"`
	JSON      accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneScimConfig struct {
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
	UserDeprovision bool                                                                                                              `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessPingoneType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSaml struct {
	// UUID
	ID     string                                                                                                 `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSaml]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSaml) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                                                                     `json:"sso_target_url"`
	JSON         accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlConfigJSON struct {
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

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                                                                    `json:"header_name"`
	JSON       accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlConfigHeaderAttribute]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlScimConfig struct {
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
	UserDeprovision bool                                                                                                           `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessSamlType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandex struct {
	// UUID
	ID     string                                                                                                   `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandex]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandex) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResult() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                       `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexScimConfig struct {
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
	UserDeprovision bool                                                                                                             `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultAccessYandexType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                        `json:"total_count"`
	JSON       accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultInfoJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultInfo]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseSuccess bool

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseSuccessTrue AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseSuccess = true
)

// This interface is a union satisfied by one of the following:
// [AccessIdentityProviderUpdateParamsAccessAzureAd],
// [AccessIdentityProviderUpdateParamsAccessCentrify],
// [AccessIdentityProviderUpdateParamsAccessFacebook],
// [AccessIdentityProviderUpdateParamsAccessGitHub],
// [AccessIdentityProviderUpdateParamsAccessGoogle],
// [AccessIdentityProviderUpdateParamsAccessGoogleApps],
// [AccessIdentityProviderUpdateParamsAccessLinkedin],
// [AccessIdentityProviderUpdateParamsAccessOidc],
// [AccessIdentityProviderUpdateParamsAccessOkta],
// [AccessIdentityProviderUpdateParamsAccessOnelogin],
// [AccessIdentityProviderUpdateParamsAccessPingone],
// [AccessIdentityProviderUpdateParamsAccessSaml],
// [AccessIdentityProviderUpdateParamsAccessYandex],
// [AccessIdentityProviderUpdateParamsAccessOnetimepin].
type AccessIdentityProviderUpdateParams interface {
	ImplementsAccessIdentityProviderUpdateParams()
}

type AccessIdentityProviderUpdateParamsAccessAzureAd struct {
	Config param.Field[AccessIdentityProviderUpdateParamsAccessAzureAdConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessAzureAdScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessAzureAdType] `json:"type"`
}

func (r AccessIdentityProviderUpdateParamsAccessAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessAzureAd) ImplementsAccessIdentityProviderUpdateParams() {

}

type AccessIdentityProviderUpdateParamsAccessAzureAdConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessAzureAdConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateParamsAccessAzureAdScimConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessAzureAdScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessAzureAdType string

const (
	AccessIdentityProviderUpdateParamsAccessAzureAdTypeOnetimepin AccessIdentityProviderUpdateParamsAccessAzureAdType = "onetimepin"
	AccessIdentityProviderUpdateParamsAccessAzureAdTypeAzureAd    AccessIdentityProviderUpdateParamsAccessAzureAdType = "azureAD"
	AccessIdentityProviderUpdateParamsAccessAzureAdTypeSaml       AccessIdentityProviderUpdateParamsAccessAzureAdType = "saml"
	AccessIdentityProviderUpdateParamsAccessAzureAdTypeCentrify   AccessIdentityProviderUpdateParamsAccessAzureAdType = "centrify"
	AccessIdentityProviderUpdateParamsAccessAzureAdTypeFacebook   AccessIdentityProviderUpdateParamsAccessAzureAdType = "facebook"
	AccessIdentityProviderUpdateParamsAccessAzureAdTypeGitHub     AccessIdentityProviderUpdateParamsAccessAzureAdType = "github"
	AccessIdentityProviderUpdateParamsAccessAzureAdTypeGoogleApps AccessIdentityProviderUpdateParamsAccessAzureAdType = "google-apps"
	AccessIdentityProviderUpdateParamsAccessAzureAdTypeGoogle     AccessIdentityProviderUpdateParamsAccessAzureAdType = "google"
	AccessIdentityProviderUpdateParamsAccessAzureAdTypeLinkedin   AccessIdentityProviderUpdateParamsAccessAzureAdType = "linkedin"
	AccessIdentityProviderUpdateParamsAccessAzureAdTypeOidc       AccessIdentityProviderUpdateParamsAccessAzureAdType = "oidc"
	AccessIdentityProviderUpdateParamsAccessAzureAdTypeOkta       AccessIdentityProviderUpdateParamsAccessAzureAdType = "okta"
	AccessIdentityProviderUpdateParamsAccessAzureAdTypeOnelogin   AccessIdentityProviderUpdateParamsAccessAzureAdType = "onelogin"
	AccessIdentityProviderUpdateParamsAccessAzureAdTypePingone    AccessIdentityProviderUpdateParamsAccessAzureAdType = "pingone"
	AccessIdentityProviderUpdateParamsAccessAzureAdTypeYandex     AccessIdentityProviderUpdateParamsAccessAzureAdType = "yandex"
)

type AccessIdentityProviderUpdateParamsAccessCentrify struct {
	Config param.Field[AccessIdentityProviderUpdateParamsAccessCentrifyConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessCentrifyScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessCentrifyType] `json:"type"`
}

func (r AccessIdentityProviderUpdateParamsAccessCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessCentrify) ImplementsAccessIdentityProviderUpdateParams() {

}

type AccessIdentityProviderUpdateParamsAccessCentrifyConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateParamsAccessCentrifyScimConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessCentrifyScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessCentrifyType string

const (
	AccessIdentityProviderUpdateParamsAccessCentrifyTypeOnetimepin AccessIdentityProviderUpdateParamsAccessCentrifyType = "onetimepin"
	AccessIdentityProviderUpdateParamsAccessCentrifyTypeAzureAd    AccessIdentityProviderUpdateParamsAccessCentrifyType = "azureAD"
	AccessIdentityProviderUpdateParamsAccessCentrifyTypeSaml       AccessIdentityProviderUpdateParamsAccessCentrifyType = "saml"
	AccessIdentityProviderUpdateParamsAccessCentrifyTypeCentrify   AccessIdentityProviderUpdateParamsAccessCentrifyType = "centrify"
	AccessIdentityProviderUpdateParamsAccessCentrifyTypeFacebook   AccessIdentityProviderUpdateParamsAccessCentrifyType = "facebook"
	AccessIdentityProviderUpdateParamsAccessCentrifyTypeGitHub     AccessIdentityProviderUpdateParamsAccessCentrifyType = "github"
	AccessIdentityProviderUpdateParamsAccessCentrifyTypeGoogleApps AccessIdentityProviderUpdateParamsAccessCentrifyType = "google-apps"
	AccessIdentityProviderUpdateParamsAccessCentrifyTypeGoogle     AccessIdentityProviderUpdateParamsAccessCentrifyType = "google"
	AccessIdentityProviderUpdateParamsAccessCentrifyTypeLinkedin   AccessIdentityProviderUpdateParamsAccessCentrifyType = "linkedin"
	AccessIdentityProviderUpdateParamsAccessCentrifyTypeOidc       AccessIdentityProviderUpdateParamsAccessCentrifyType = "oidc"
	AccessIdentityProviderUpdateParamsAccessCentrifyTypeOkta       AccessIdentityProviderUpdateParamsAccessCentrifyType = "okta"
	AccessIdentityProviderUpdateParamsAccessCentrifyTypeOnelogin   AccessIdentityProviderUpdateParamsAccessCentrifyType = "onelogin"
	AccessIdentityProviderUpdateParamsAccessCentrifyTypePingone    AccessIdentityProviderUpdateParamsAccessCentrifyType = "pingone"
	AccessIdentityProviderUpdateParamsAccessCentrifyTypeYandex     AccessIdentityProviderUpdateParamsAccessCentrifyType = "yandex"
)

type AccessIdentityProviderUpdateParamsAccessFacebook struct {
	Config param.Field[AccessIdentityProviderUpdateParamsAccessFacebookConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessFacebookScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessFacebookType] `json:"type"`
}

func (r AccessIdentityProviderUpdateParamsAccessFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessFacebook) ImplementsAccessIdentityProviderUpdateParams() {

}

type AccessIdentityProviderUpdateParamsAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderUpdateParamsAccessFacebookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateParamsAccessFacebookScimConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessFacebookScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessFacebookType string

const (
	AccessIdentityProviderUpdateParamsAccessFacebookTypeOnetimepin AccessIdentityProviderUpdateParamsAccessFacebookType = "onetimepin"
	AccessIdentityProviderUpdateParamsAccessFacebookTypeAzureAd    AccessIdentityProviderUpdateParamsAccessFacebookType = "azureAD"
	AccessIdentityProviderUpdateParamsAccessFacebookTypeSaml       AccessIdentityProviderUpdateParamsAccessFacebookType = "saml"
	AccessIdentityProviderUpdateParamsAccessFacebookTypeCentrify   AccessIdentityProviderUpdateParamsAccessFacebookType = "centrify"
	AccessIdentityProviderUpdateParamsAccessFacebookTypeFacebook   AccessIdentityProviderUpdateParamsAccessFacebookType = "facebook"
	AccessIdentityProviderUpdateParamsAccessFacebookTypeGitHub     AccessIdentityProviderUpdateParamsAccessFacebookType = "github"
	AccessIdentityProviderUpdateParamsAccessFacebookTypeGoogleApps AccessIdentityProviderUpdateParamsAccessFacebookType = "google-apps"
	AccessIdentityProviderUpdateParamsAccessFacebookTypeGoogle     AccessIdentityProviderUpdateParamsAccessFacebookType = "google"
	AccessIdentityProviderUpdateParamsAccessFacebookTypeLinkedin   AccessIdentityProviderUpdateParamsAccessFacebookType = "linkedin"
	AccessIdentityProviderUpdateParamsAccessFacebookTypeOidc       AccessIdentityProviderUpdateParamsAccessFacebookType = "oidc"
	AccessIdentityProviderUpdateParamsAccessFacebookTypeOkta       AccessIdentityProviderUpdateParamsAccessFacebookType = "okta"
	AccessIdentityProviderUpdateParamsAccessFacebookTypeOnelogin   AccessIdentityProviderUpdateParamsAccessFacebookType = "onelogin"
	AccessIdentityProviderUpdateParamsAccessFacebookTypePingone    AccessIdentityProviderUpdateParamsAccessFacebookType = "pingone"
	AccessIdentityProviderUpdateParamsAccessFacebookTypeYandex     AccessIdentityProviderUpdateParamsAccessFacebookType = "yandex"
)

type AccessIdentityProviderUpdateParamsAccessGitHub struct {
	Config param.Field[AccessIdentityProviderUpdateParamsAccessGitHubConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessGitHubScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessGitHubType] `json:"type"`
}

func (r AccessIdentityProviderUpdateParamsAccessGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessGitHub) ImplementsAccessIdentityProviderUpdateParams() {

}

type AccessIdentityProviderUpdateParamsAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderUpdateParamsAccessGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateParamsAccessGitHubScimConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessGitHubScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessGitHubType string

const (
	AccessIdentityProviderUpdateParamsAccessGitHubTypeOnetimepin AccessIdentityProviderUpdateParamsAccessGitHubType = "onetimepin"
	AccessIdentityProviderUpdateParamsAccessGitHubTypeAzureAd    AccessIdentityProviderUpdateParamsAccessGitHubType = "azureAD"
	AccessIdentityProviderUpdateParamsAccessGitHubTypeSaml       AccessIdentityProviderUpdateParamsAccessGitHubType = "saml"
	AccessIdentityProviderUpdateParamsAccessGitHubTypeCentrify   AccessIdentityProviderUpdateParamsAccessGitHubType = "centrify"
	AccessIdentityProviderUpdateParamsAccessGitHubTypeFacebook   AccessIdentityProviderUpdateParamsAccessGitHubType = "facebook"
	AccessIdentityProviderUpdateParamsAccessGitHubTypeGitHub     AccessIdentityProviderUpdateParamsAccessGitHubType = "github"
	AccessIdentityProviderUpdateParamsAccessGitHubTypeGoogleApps AccessIdentityProviderUpdateParamsAccessGitHubType = "google-apps"
	AccessIdentityProviderUpdateParamsAccessGitHubTypeGoogle     AccessIdentityProviderUpdateParamsAccessGitHubType = "google"
	AccessIdentityProviderUpdateParamsAccessGitHubTypeLinkedin   AccessIdentityProviderUpdateParamsAccessGitHubType = "linkedin"
	AccessIdentityProviderUpdateParamsAccessGitHubTypeOidc       AccessIdentityProviderUpdateParamsAccessGitHubType = "oidc"
	AccessIdentityProviderUpdateParamsAccessGitHubTypeOkta       AccessIdentityProviderUpdateParamsAccessGitHubType = "okta"
	AccessIdentityProviderUpdateParamsAccessGitHubTypeOnelogin   AccessIdentityProviderUpdateParamsAccessGitHubType = "onelogin"
	AccessIdentityProviderUpdateParamsAccessGitHubTypePingone    AccessIdentityProviderUpdateParamsAccessGitHubType = "pingone"
	AccessIdentityProviderUpdateParamsAccessGitHubTypeYandex     AccessIdentityProviderUpdateParamsAccessGitHubType = "yandex"
)

type AccessIdentityProviderUpdateParamsAccessGoogle struct {
	Config param.Field[AccessIdentityProviderUpdateParamsAccessGoogleConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessGoogleScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessGoogleType] `json:"type"`
}

func (r AccessIdentityProviderUpdateParamsAccessGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessGoogle) ImplementsAccessIdentityProviderUpdateParams() {

}

type AccessIdentityProviderUpdateParamsAccessGoogleConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r AccessIdentityProviderUpdateParamsAccessGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateParamsAccessGoogleScimConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessGoogleScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessGoogleType string

const (
	AccessIdentityProviderUpdateParamsAccessGoogleTypeOnetimepin AccessIdentityProviderUpdateParamsAccessGoogleType = "onetimepin"
	AccessIdentityProviderUpdateParamsAccessGoogleTypeAzureAd    AccessIdentityProviderUpdateParamsAccessGoogleType = "azureAD"
	AccessIdentityProviderUpdateParamsAccessGoogleTypeSaml       AccessIdentityProviderUpdateParamsAccessGoogleType = "saml"
	AccessIdentityProviderUpdateParamsAccessGoogleTypeCentrify   AccessIdentityProviderUpdateParamsAccessGoogleType = "centrify"
	AccessIdentityProviderUpdateParamsAccessGoogleTypeFacebook   AccessIdentityProviderUpdateParamsAccessGoogleType = "facebook"
	AccessIdentityProviderUpdateParamsAccessGoogleTypeGitHub     AccessIdentityProviderUpdateParamsAccessGoogleType = "github"
	AccessIdentityProviderUpdateParamsAccessGoogleTypeGoogleApps AccessIdentityProviderUpdateParamsAccessGoogleType = "google-apps"
	AccessIdentityProviderUpdateParamsAccessGoogleTypeGoogle     AccessIdentityProviderUpdateParamsAccessGoogleType = "google"
	AccessIdentityProviderUpdateParamsAccessGoogleTypeLinkedin   AccessIdentityProviderUpdateParamsAccessGoogleType = "linkedin"
	AccessIdentityProviderUpdateParamsAccessGoogleTypeOidc       AccessIdentityProviderUpdateParamsAccessGoogleType = "oidc"
	AccessIdentityProviderUpdateParamsAccessGoogleTypeOkta       AccessIdentityProviderUpdateParamsAccessGoogleType = "okta"
	AccessIdentityProviderUpdateParamsAccessGoogleTypeOnelogin   AccessIdentityProviderUpdateParamsAccessGoogleType = "onelogin"
	AccessIdentityProviderUpdateParamsAccessGoogleTypePingone    AccessIdentityProviderUpdateParamsAccessGoogleType = "pingone"
	AccessIdentityProviderUpdateParamsAccessGoogleTypeYandex     AccessIdentityProviderUpdateParamsAccessGoogleType = "yandex"
)

type AccessIdentityProviderUpdateParamsAccessGoogleApps struct {
	Config param.Field[AccessIdentityProviderUpdateParamsAccessGoogleAppsConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessGoogleAppsScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessGoogleAppsType] `json:"type"`
}

func (r AccessIdentityProviderUpdateParamsAccessGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessGoogleApps) ImplementsAccessIdentityProviderUpdateParams() {

}

type AccessIdentityProviderUpdateParamsAccessGoogleAppsConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateParamsAccessGoogleAppsScimConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessGoogleAppsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessGoogleAppsType string

const (
	AccessIdentityProviderUpdateParamsAccessGoogleAppsTypeOnetimepin AccessIdentityProviderUpdateParamsAccessGoogleAppsType = "onetimepin"
	AccessIdentityProviderUpdateParamsAccessGoogleAppsTypeAzureAd    AccessIdentityProviderUpdateParamsAccessGoogleAppsType = "azureAD"
	AccessIdentityProviderUpdateParamsAccessGoogleAppsTypeSaml       AccessIdentityProviderUpdateParamsAccessGoogleAppsType = "saml"
	AccessIdentityProviderUpdateParamsAccessGoogleAppsTypeCentrify   AccessIdentityProviderUpdateParamsAccessGoogleAppsType = "centrify"
	AccessIdentityProviderUpdateParamsAccessGoogleAppsTypeFacebook   AccessIdentityProviderUpdateParamsAccessGoogleAppsType = "facebook"
	AccessIdentityProviderUpdateParamsAccessGoogleAppsTypeGitHub     AccessIdentityProviderUpdateParamsAccessGoogleAppsType = "github"
	AccessIdentityProviderUpdateParamsAccessGoogleAppsTypeGoogleApps AccessIdentityProviderUpdateParamsAccessGoogleAppsType = "google-apps"
	AccessIdentityProviderUpdateParamsAccessGoogleAppsTypeGoogle     AccessIdentityProviderUpdateParamsAccessGoogleAppsType = "google"
	AccessIdentityProviderUpdateParamsAccessGoogleAppsTypeLinkedin   AccessIdentityProviderUpdateParamsAccessGoogleAppsType = "linkedin"
	AccessIdentityProviderUpdateParamsAccessGoogleAppsTypeOidc       AccessIdentityProviderUpdateParamsAccessGoogleAppsType = "oidc"
	AccessIdentityProviderUpdateParamsAccessGoogleAppsTypeOkta       AccessIdentityProviderUpdateParamsAccessGoogleAppsType = "okta"
	AccessIdentityProviderUpdateParamsAccessGoogleAppsTypeOnelogin   AccessIdentityProviderUpdateParamsAccessGoogleAppsType = "onelogin"
	AccessIdentityProviderUpdateParamsAccessGoogleAppsTypePingone    AccessIdentityProviderUpdateParamsAccessGoogleAppsType = "pingone"
	AccessIdentityProviderUpdateParamsAccessGoogleAppsTypeYandex     AccessIdentityProviderUpdateParamsAccessGoogleAppsType = "yandex"
)

type AccessIdentityProviderUpdateParamsAccessLinkedin struct {
	Config param.Field[AccessIdentityProviderUpdateParamsAccessLinkedinConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessLinkedinScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessLinkedinType] `json:"type"`
}

func (r AccessIdentityProviderUpdateParamsAccessLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessLinkedin) ImplementsAccessIdentityProviderUpdateParams() {

}

type AccessIdentityProviderUpdateParamsAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderUpdateParamsAccessLinkedinConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateParamsAccessLinkedinScimConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessLinkedinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessLinkedinType string

const (
	AccessIdentityProviderUpdateParamsAccessLinkedinTypeOnetimepin AccessIdentityProviderUpdateParamsAccessLinkedinType = "onetimepin"
	AccessIdentityProviderUpdateParamsAccessLinkedinTypeAzureAd    AccessIdentityProviderUpdateParamsAccessLinkedinType = "azureAD"
	AccessIdentityProviderUpdateParamsAccessLinkedinTypeSaml       AccessIdentityProviderUpdateParamsAccessLinkedinType = "saml"
	AccessIdentityProviderUpdateParamsAccessLinkedinTypeCentrify   AccessIdentityProviderUpdateParamsAccessLinkedinType = "centrify"
	AccessIdentityProviderUpdateParamsAccessLinkedinTypeFacebook   AccessIdentityProviderUpdateParamsAccessLinkedinType = "facebook"
	AccessIdentityProviderUpdateParamsAccessLinkedinTypeGitHub     AccessIdentityProviderUpdateParamsAccessLinkedinType = "github"
	AccessIdentityProviderUpdateParamsAccessLinkedinTypeGoogleApps AccessIdentityProviderUpdateParamsAccessLinkedinType = "google-apps"
	AccessIdentityProviderUpdateParamsAccessLinkedinTypeGoogle     AccessIdentityProviderUpdateParamsAccessLinkedinType = "google"
	AccessIdentityProviderUpdateParamsAccessLinkedinTypeLinkedin   AccessIdentityProviderUpdateParamsAccessLinkedinType = "linkedin"
	AccessIdentityProviderUpdateParamsAccessLinkedinTypeOidc       AccessIdentityProviderUpdateParamsAccessLinkedinType = "oidc"
	AccessIdentityProviderUpdateParamsAccessLinkedinTypeOkta       AccessIdentityProviderUpdateParamsAccessLinkedinType = "okta"
	AccessIdentityProviderUpdateParamsAccessLinkedinTypeOnelogin   AccessIdentityProviderUpdateParamsAccessLinkedinType = "onelogin"
	AccessIdentityProviderUpdateParamsAccessLinkedinTypePingone    AccessIdentityProviderUpdateParamsAccessLinkedinType = "pingone"
	AccessIdentityProviderUpdateParamsAccessLinkedinTypeYandex     AccessIdentityProviderUpdateParamsAccessLinkedinType = "yandex"
)

type AccessIdentityProviderUpdateParamsAccessOidc struct {
	Config param.Field[AccessIdentityProviderUpdateParamsAccessOidcConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessOidcScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessOidcType] `json:"type"`
}

func (r AccessIdentityProviderUpdateParamsAccessOidc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessOidc) ImplementsAccessIdentityProviderUpdateParams() {

}

type AccessIdentityProviderUpdateParamsAccessOidcConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessOidcConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateParamsAccessOidcScimConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessOidcScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessOidcType string

const (
	AccessIdentityProviderUpdateParamsAccessOidcTypeOnetimepin AccessIdentityProviderUpdateParamsAccessOidcType = "onetimepin"
	AccessIdentityProviderUpdateParamsAccessOidcTypeAzureAd    AccessIdentityProviderUpdateParamsAccessOidcType = "azureAD"
	AccessIdentityProviderUpdateParamsAccessOidcTypeSaml       AccessIdentityProviderUpdateParamsAccessOidcType = "saml"
	AccessIdentityProviderUpdateParamsAccessOidcTypeCentrify   AccessIdentityProviderUpdateParamsAccessOidcType = "centrify"
	AccessIdentityProviderUpdateParamsAccessOidcTypeFacebook   AccessIdentityProviderUpdateParamsAccessOidcType = "facebook"
	AccessIdentityProviderUpdateParamsAccessOidcTypeGitHub     AccessIdentityProviderUpdateParamsAccessOidcType = "github"
	AccessIdentityProviderUpdateParamsAccessOidcTypeGoogleApps AccessIdentityProviderUpdateParamsAccessOidcType = "google-apps"
	AccessIdentityProviderUpdateParamsAccessOidcTypeGoogle     AccessIdentityProviderUpdateParamsAccessOidcType = "google"
	AccessIdentityProviderUpdateParamsAccessOidcTypeLinkedin   AccessIdentityProviderUpdateParamsAccessOidcType = "linkedin"
	AccessIdentityProviderUpdateParamsAccessOidcTypeOidc       AccessIdentityProviderUpdateParamsAccessOidcType = "oidc"
	AccessIdentityProviderUpdateParamsAccessOidcTypeOkta       AccessIdentityProviderUpdateParamsAccessOidcType = "okta"
	AccessIdentityProviderUpdateParamsAccessOidcTypeOnelogin   AccessIdentityProviderUpdateParamsAccessOidcType = "onelogin"
	AccessIdentityProviderUpdateParamsAccessOidcTypePingone    AccessIdentityProviderUpdateParamsAccessOidcType = "pingone"
	AccessIdentityProviderUpdateParamsAccessOidcTypeYandex     AccessIdentityProviderUpdateParamsAccessOidcType = "yandex"
)

type AccessIdentityProviderUpdateParamsAccessOkta struct {
	Config param.Field[AccessIdentityProviderUpdateParamsAccessOktaConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessOktaScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessOktaType] `json:"type"`
}

func (r AccessIdentityProviderUpdateParamsAccessOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessOkta) ImplementsAccessIdentityProviderUpdateParams() {

}

type AccessIdentityProviderUpdateParamsAccessOktaConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateParamsAccessOktaScimConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessOktaScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessOktaType string

const (
	AccessIdentityProviderUpdateParamsAccessOktaTypeOnetimepin AccessIdentityProviderUpdateParamsAccessOktaType = "onetimepin"
	AccessIdentityProviderUpdateParamsAccessOktaTypeAzureAd    AccessIdentityProviderUpdateParamsAccessOktaType = "azureAD"
	AccessIdentityProviderUpdateParamsAccessOktaTypeSaml       AccessIdentityProviderUpdateParamsAccessOktaType = "saml"
	AccessIdentityProviderUpdateParamsAccessOktaTypeCentrify   AccessIdentityProviderUpdateParamsAccessOktaType = "centrify"
	AccessIdentityProviderUpdateParamsAccessOktaTypeFacebook   AccessIdentityProviderUpdateParamsAccessOktaType = "facebook"
	AccessIdentityProviderUpdateParamsAccessOktaTypeGitHub     AccessIdentityProviderUpdateParamsAccessOktaType = "github"
	AccessIdentityProviderUpdateParamsAccessOktaTypeGoogleApps AccessIdentityProviderUpdateParamsAccessOktaType = "google-apps"
	AccessIdentityProviderUpdateParamsAccessOktaTypeGoogle     AccessIdentityProviderUpdateParamsAccessOktaType = "google"
	AccessIdentityProviderUpdateParamsAccessOktaTypeLinkedin   AccessIdentityProviderUpdateParamsAccessOktaType = "linkedin"
	AccessIdentityProviderUpdateParamsAccessOktaTypeOidc       AccessIdentityProviderUpdateParamsAccessOktaType = "oidc"
	AccessIdentityProviderUpdateParamsAccessOktaTypeOkta       AccessIdentityProviderUpdateParamsAccessOktaType = "okta"
	AccessIdentityProviderUpdateParamsAccessOktaTypeOnelogin   AccessIdentityProviderUpdateParamsAccessOktaType = "onelogin"
	AccessIdentityProviderUpdateParamsAccessOktaTypePingone    AccessIdentityProviderUpdateParamsAccessOktaType = "pingone"
	AccessIdentityProviderUpdateParamsAccessOktaTypeYandex     AccessIdentityProviderUpdateParamsAccessOktaType = "yandex"
)

type AccessIdentityProviderUpdateParamsAccessOnelogin struct {
	Config param.Field[AccessIdentityProviderUpdateParamsAccessOneloginConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessOneloginScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessOneloginType] `json:"type"`
}

func (r AccessIdentityProviderUpdateParamsAccessOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessOnelogin) ImplementsAccessIdentityProviderUpdateParams() {

}

type AccessIdentityProviderUpdateParamsAccessOneloginConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateParamsAccessOneloginScimConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessOneloginScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessOneloginType string

const (
	AccessIdentityProviderUpdateParamsAccessOneloginTypeOnetimepin AccessIdentityProviderUpdateParamsAccessOneloginType = "onetimepin"
	AccessIdentityProviderUpdateParamsAccessOneloginTypeAzureAd    AccessIdentityProviderUpdateParamsAccessOneloginType = "azureAD"
	AccessIdentityProviderUpdateParamsAccessOneloginTypeSaml       AccessIdentityProviderUpdateParamsAccessOneloginType = "saml"
	AccessIdentityProviderUpdateParamsAccessOneloginTypeCentrify   AccessIdentityProviderUpdateParamsAccessOneloginType = "centrify"
	AccessIdentityProviderUpdateParamsAccessOneloginTypeFacebook   AccessIdentityProviderUpdateParamsAccessOneloginType = "facebook"
	AccessIdentityProviderUpdateParamsAccessOneloginTypeGitHub     AccessIdentityProviderUpdateParamsAccessOneloginType = "github"
	AccessIdentityProviderUpdateParamsAccessOneloginTypeGoogleApps AccessIdentityProviderUpdateParamsAccessOneloginType = "google-apps"
	AccessIdentityProviderUpdateParamsAccessOneloginTypeGoogle     AccessIdentityProviderUpdateParamsAccessOneloginType = "google"
	AccessIdentityProviderUpdateParamsAccessOneloginTypeLinkedin   AccessIdentityProviderUpdateParamsAccessOneloginType = "linkedin"
	AccessIdentityProviderUpdateParamsAccessOneloginTypeOidc       AccessIdentityProviderUpdateParamsAccessOneloginType = "oidc"
	AccessIdentityProviderUpdateParamsAccessOneloginTypeOkta       AccessIdentityProviderUpdateParamsAccessOneloginType = "okta"
	AccessIdentityProviderUpdateParamsAccessOneloginTypeOnelogin   AccessIdentityProviderUpdateParamsAccessOneloginType = "onelogin"
	AccessIdentityProviderUpdateParamsAccessOneloginTypePingone    AccessIdentityProviderUpdateParamsAccessOneloginType = "pingone"
	AccessIdentityProviderUpdateParamsAccessOneloginTypeYandex     AccessIdentityProviderUpdateParamsAccessOneloginType = "yandex"
)

type AccessIdentityProviderUpdateParamsAccessPingone struct {
	Config param.Field[AccessIdentityProviderUpdateParamsAccessPingoneConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessPingoneScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessPingoneType] `json:"type"`
}

func (r AccessIdentityProviderUpdateParamsAccessPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessPingone) ImplementsAccessIdentityProviderUpdateParams() {

}

type AccessIdentityProviderUpdateParamsAccessPingoneConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateParamsAccessPingoneScimConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessPingoneScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessPingoneType string

const (
	AccessIdentityProviderUpdateParamsAccessPingoneTypeOnetimepin AccessIdentityProviderUpdateParamsAccessPingoneType = "onetimepin"
	AccessIdentityProviderUpdateParamsAccessPingoneTypeAzureAd    AccessIdentityProviderUpdateParamsAccessPingoneType = "azureAD"
	AccessIdentityProviderUpdateParamsAccessPingoneTypeSaml       AccessIdentityProviderUpdateParamsAccessPingoneType = "saml"
	AccessIdentityProviderUpdateParamsAccessPingoneTypeCentrify   AccessIdentityProviderUpdateParamsAccessPingoneType = "centrify"
	AccessIdentityProviderUpdateParamsAccessPingoneTypeFacebook   AccessIdentityProviderUpdateParamsAccessPingoneType = "facebook"
	AccessIdentityProviderUpdateParamsAccessPingoneTypeGitHub     AccessIdentityProviderUpdateParamsAccessPingoneType = "github"
	AccessIdentityProviderUpdateParamsAccessPingoneTypeGoogleApps AccessIdentityProviderUpdateParamsAccessPingoneType = "google-apps"
	AccessIdentityProviderUpdateParamsAccessPingoneTypeGoogle     AccessIdentityProviderUpdateParamsAccessPingoneType = "google"
	AccessIdentityProviderUpdateParamsAccessPingoneTypeLinkedin   AccessIdentityProviderUpdateParamsAccessPingoneType = "linkedin"
	AccessIdentityProviderUpdateParamsAccessPingoneTypeOidc       AccessIdentityProviderUpdateParamsAccessPingoneType = "oidc"
	AccessIdentityProviderUpdateParamsAccessPingoneTypeOkta       AccessIdentityProviderUpdateParamsAccessPingoneType = "okta"
	AccessIdentityProviderUpdateParamsAccessPingoneTypeOnelogin   AccessIdentityProviderUpdateParamsAccessPingoneType = "onelogin"
	AccessIdentityProviderUpdateParamsAccessPingoneTypePingone    AccessIdentityProviderUpdateParamsAccessPingoneType = "pingone"
	AccessIdentityProviderUpdateParamsAccessPingoneTypeYandex     AccessIdentityProviderUpdateParamsAccessPingoneType = "yandex"
)

type AccessIdentityProviderUpdateParamsAccessSaml struct {
	Config param.Field[AccessIdentityProviderUpdateParamsAccessSamlConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessSamlScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessSamlType] `json:"type"`
}

func (r AccessIdentityProviderUpdateParamsAccessSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessSaml) ImplementsAccessIdentityProviderUpdateParams() {

}

type AccessIdentityProviderUpdateParamsAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]AccessIdentityProviderUpdateParamsAccessSamlConfigHeaderAttribute] `json:"header_attributes"`
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

func (r AccessIdentityProviderUpdateParamsAccessSamlConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderUpdateParamsAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r AccessIdentityProviderUpdateParamsAccessSamlConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateParamsAccessSamlScimConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessSamlScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessSamlType string

const (
	AccessIdentityProviderUpdateParamsAccessSamlTypeOnetimepin AccessIdentityProviderUpdateParamsAccessSamlType = "onetimepin"
	AccessIdentityProviderUpdateParamsAccessSamlTypeAzureAd    AccessIdentityProviderUpdateParamsAccessSamlType = "azureAD"
	AccessIdentityProviderUpdateParamsAccessSamlTypeSaml       AccessIdentityProviderUpdateParamsAccessSamlType = "saml"
	AccessIdentityProviderUpdateParamsAccessSamlTypeCentrify   AccessIdentityProviderUpdateParamsAccessSamlType = "centrify"
	AccessIdentityProviderUpdateParamsAccessSamlTypeFacebook   AccessIdentityProviderUpdateParamsAccessSamlType = "facebook"
	AccessIdentityProviderUpdateParamsAccessSamlTypeGitHub     AccessIdentityProviderUpdateParamsAccessSamlType = "github"
	AccessIdentityProviderUpdateParamsAccessSamlTypeGoogleApps AccessIdentityProviderUpdateParamsAccessSamlType = "google-apps"
	AccessIdentityProviderUpdateParamsAccessSamlTypeGoogle     AccessIdentityProviderUpdateParamsAccessSamlType = "google"
	AccessIdentityProviderUpdateParamsAccessSamlTypeLinkedin   AccessIdentityProviderUpdateParamsAccessSamlType = "linkedin"
	AccessIdentityProviderUpdateParamsAccessSamlTypeOidc       AccessIdentityProviderUpdateParamsAccessSamlType = "oidc"
	AccessIdentityProviderUpdateParamsAccessSamlTypeOkta       AccessIdentityProviderUpdateParamsAccessSamlType = "okta"
	AccessIdentityProviderUpdateParamsAccessSamlTypeOnelogin   AccessIdentityProviderUpdateParamsAccessSamlType = "onelogin"
	AccessIdentityProviderUpdateParamsAccessSamlTypePingone    AccessIdentityProviderUpdateParamsAccessSamlType = "pingone"
	AccessIdentityProviderUpdateParamsAccessSamlTypeYandex     AccessIdentityProviderUpdateParamsAccessSamlType = "yandex"
)

type AccessIdentityProviderUpdateParamsAccessYandex struct {
	Config param.Field[AccessIdentityProviderUpdateParamsAccessYandexConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessYandexScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessYandexType] `json:"type"`
}

func (r AccessIdentityProviderUpdateParamsAccessYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessYandex) ImplementsAccessIdentityProviderUpdateParams() {

}

type AccessIdentityProviderUpdateParamsAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderUpdateParamsAccessYandexConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateParamsAccessYandexScimConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessYandexScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessYandexType string

const (
	AccessIdentityProviderUpdateParamsAccessYandexTypeOnetimepin AccessIdentityProviderUpdateParamsAccessYandexType = "onetimepin"
	AccessIdentityProviderUpdateParamsAccessYandexTypeAzureAd    AccessIdentityProviderUpdateParamsAccessYandexType = "azureAD"
	AccessIdentityProviderUpdateParamsAccessYandexTypeSaml       AccessIdentityProviderUpdateParamsAccessYandexType = "saml"
	AccessIdentityProviderUpdateParamsAccessYandexTypeCentrify   AccessIdentityProviderUpdateParamsAccessYandexType = "centrify"
	AccessIdentityProviderUpdateParamsAccessYandexTypeFacebook   AccessIdentityProviderUpdateParamsAccessYandexType = "facebook"
	AccessIdentityProviderUpdateParamsAccessYandexTypeGitHub     AccessIdentityProviderUpdateParamsAccessYandexType = "github"
	AccessIdentityProviderUpdateParamsAccessYandexTypeGoogleApps AccessIdentityProviderUpdateParamsAccessYandexType = "google-apps"
	AccessIdentityProviderUpdateParamsAccessYandexTypeGoogle     AccessIdentityProviderUpdateParamsAccessYandexType = "google"
	AccessIdentityProviderUpdateParamsAccessYandexTypeLinkedin   AccessIdentityProviderUpdateParamsAccessYandexType = "linkedin"
	AccessIdentityProviderUpdateParamsAccessYandexTypeOidc       AccessIdentityProviderUpdateParamsAccessYandexType = "oidc"
	AccessIdentityProviderUpdateParamsAccessYandexTypeOkta       AccessIdentityProviderUpdateParamsAccessYandexType = "okta"
	AccessIdentityProviderUpdateParamsAccessYandexTypeOnelogin   AccessIdentityProviderUpdateParamsAccessYandexType = "onelogin"
	AccessIdentityProviderUpdateParamsAccessYandexTypePingone    AccessIdentityProviderUpdateParamsAccessYandexType = "pingone"
	AccessIdentityProviderUpdateParamsAccessYandexTypeYandex     AccessIdentityProviderUpdateParamsAccessYandexType = "yandex"
)

type AccessIdentityProviderUpdateParamsAccessOnetimepin struct {
	Config param.Field[interface{}] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessOnetimepinScimConfig] `json:"scim_config"`
	Type       param.Field[AccessIdentityProviderUpdateParamsAccessOnetimepinType]       `json:"type"`
}

func (r AccessIdentityProviderUpdateParamsAccessOnetimepin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessOnetimepin) ImplementsAccessIdentityProviderUpdateParams() {

}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateParamsAccessOnetimepinScimConfig struct {
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

func (r AccessIdentityProviderUpdateParamsAccessOnetimepinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderUpdateParamsAccessOnetimepinType string

const (
	AccessIdentityProviderUpdateParamsAccessOnetimepinTypeOnetimepin AccessIdentityProviderUpdateParamsAccessOnetimepinType = "onetimepin"
)

// This interface is a union satisfied by one of the following:
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAd],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrify],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebook],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHub],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogle],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleApps],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedin],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidc],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOkta],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOnelogin],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingone],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSaml],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandex],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOnetimepin].
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams interface {
	ImplementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams()
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAd struct {
	Config param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType] `json:"type"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAd) ImplementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdScimConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessAzureAdType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrify struct {
	Config param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType] `json:"type"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrify) ImplementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyScimConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessCentrifyType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebook struct {
	Config param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType] `json:"type"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebook) ImplementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookScimConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessFacebookType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHub struct {
	Config param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType] `json:"type"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHub) ImplementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubScimConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGitHubType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogle struct {
	Config param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType] `json:"type"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogle) ImplementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleScimConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleApps struct {
	Config param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType] `json:"type"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleApps) ImplementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsScimConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessGoogleAppsType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedin struct {
	Config param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType] `json:"type"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedin) ImplementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinScimConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessLinkedinType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidc struct {
	Config param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType] `json:"type"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidc) ImplementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcScimConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOidcType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOkta struct {
	Config param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType] `json:"type"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOkta) ImplementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaScimConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOktaType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOnelogin struct {
	Config param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType] `json:"type"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOnelogin) ImplementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginScimConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOneloginType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingone struct {
	Config param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType] `json:"type"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingone) ImplementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneScimConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessPingoneType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSaml struct {
	Config param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType] `json:"type"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSaml) ImplementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlConfigHeaderAttribute] `json:"header_attributes"`
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlScimConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessSamlType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandex struct {
	Config param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType] `json:"type"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandex) ImplementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexScimConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessYandexType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOnetimepin struct {
	Config param.Field[interface{}] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOnetimepinScimConfig] `json:"scim_config"`
	Type       param.Field[AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOnetimepinType]       `json:"type"`
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOnetimepin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOnetimepin) ImplementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams() {

}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOnetimepinScimConfig struct {
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

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOnetimepinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOnetimepinType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOnetimepinTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParamsAccessOnetimepinType = "onetimepin"
)

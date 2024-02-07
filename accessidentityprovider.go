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
	var env AccessIdentityProviderGetResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/identity_providers/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured identity provider.
func (r *AccessIdentityProviderService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, body AccessIdentityProviderUpdateParams, opts ...option.RequestOption) (res *AccessIdentityProviderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessIdentityProviderUpdateResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/identity_providers/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an identity provider from Access.
func (r *AccessIdentityProviderService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessIdentityProviderDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessIdentityProviderDeleteResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/identity_providers/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Adds a new identity provider to Access.
func (r *AccessIdentityProviderService) AccessIdentityProvidersAddAnAccessIdentityProvider(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderParams, opts ...option.RequestOption) (res *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/identity_providers", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all configured identity providers.
func (r *AccessIdentityProviderService) AccessIdentityProvidersListAccessIdentityProviders(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *[]AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/identity_providers", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AccessIdentityProviderGetResponseAccessAzureAd],
// [AccessIdentityProviderGetResponseAccessCentrify],
// [AccessIdentityProviderGetResponseAccessFacebook],
// [AccessIdentityProviderGetResponseAccessGitHub],
// [AccessIdentityProviderGetResponseAccessGoogle],
// [AccessIdentityProviderGetResponseAccessGoogleApps],
// [AccessIdentityProviderGetResponseAccessLinkedin],
// [AccessIdentityProviderGetResponseAccessOidc],
// [AccessIdentityProviderGetResponseAccessOkta],
// [AccessIdentityProviderGetResponseAccessOnelogin],
// [AccessIdentityProviderGetResponseAccessPingone],
// [AccessIdentityProviderGetResponseAccessSaml],
// [AccessIdentityProviderGetResponseAccessYandex] or
// [AccessIdentityProviderGetResponseAccessOnetimepin].
type AccessIdentityProviderGetResponse interface {
	implementsAccessIdentityProviderGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessIdentityProviderGetResponse)(nil)).Elem(), "")
}

type AccessIdentityProviderGetResponseAccessAzureAd struct {
	// UUID
	ID     string                                               `json:"id"`
	Config AccessIdentityProviderGetResponseAccessAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessAzureAdType `json:"type"`
	JSON accessIdentityProviderGetResponseAccessAzureAdJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessAzureAdJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessAzureAd]
type accessIdentityProviderGetResponseAccessAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessAzureAd) implementsAccessIdentityProviderGetResponse() {
}

type AccessIdentityProviderGetResponseAccessAzureAdConfig struct {
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
	SupportGroups bool                                                     `json:"support_groups"`
	JSON          accessIdentityProviderGetResponseAccessAzureAdConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessAzureAdConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseAccessAzureAdConfig]
type accessIdentityProviderGetResponseAccessAzureAdConfigJSON struct {
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

func (r *AccessIdentityProviderGetResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseAccessAzureAdScimConfig struct {
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
	UserDeprovision bool                                                         `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessAzureAdScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseAccessAzureAdScimConfig]
type accessIdentityProviderGetResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseAccessAzureAdType string

const (
	AccessIdentityProviderGetResponseAccessAzureAdTypeOnetimepin AccessIdentityProviderGetResponseAccessAzureAdType = "onetimepin"
	AccessIdentityProviderGetResponseAccessAzureAdTypeAzureAd    AccessIdentityProviderGetResponseAccessAzureAdType = "azureAD"
	AccessIdentityProviderGetResponseAccessAzureAdTypeSaml       AccessIdentityProviderGetResponseAccessAzureAdType = "saml"
	AccessIdentityProviderGetResponseAccessAzureAdTypeCentrify   AccessIdentityProviderGetResponseAccessAzureAdType = "centrify"
	AccessIdentityProviderGetResponseAccessAzureAdTypeFacebook   AccessIdentityProviderGetResponseAccessAzureAdType = "facebook"
	AccessIdentityProviderGetResponseAccessAzureAdTypeGitHub     AccessIdentityProviderGetResponseAccessAzureAdType = "github"
	AccessIdentityProviderGetResponseAccessAzureAdTypeGoogleApps AccessIdentityProviderGetResponseAccessAzureAdType = "google-apps"
	AccessIdentityProviderGetResponseAccessAzureAdTypeGoogle     AccessIdentityProviderGetResponseAccessAzureAdType = "google"
	AccessIdentityProviderGetResponseAccessAzureAdTypeLinkedin   AccessIdentityProviderGetResponseAccessAzureAdType = "linkedin"
	AccessIdentityProviderGetResponseAccessAzureAdTypeOidc       AccessIdentityProviderGetResponseAccessAzureAdType = "oidc"
	AccessIdentityProviderGetResponseAccessAzureAdTypeOkta       AccessIdentityProviderGetResponseAccessAzureAdType = "okta"
	AccessIdentityProviderGetResponseAccessAzureAdTypeOnelogin   AccessIdentityProviderGetResponseAccessAzureAdType = "onelogin"
	AccessIdentityProviderGetResponseAccessAzureAdTypePingone    AccessIdentityProviderGetResponseAccessAzureAdType = "pingone"
	AccessIdentityProviderGetResponseAccessAzureAdTypeYandex     AccessIdentityProviderGetResponseAccessAzureAdType = "yandex"
)

type AccessIdentityProviderGetResponseAccessCentrify struct {
	// UUID
	ID     string                                                `json:"id"`
	Config AccessIdentityProviderGetResponseAccessCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessCentrifyType `json:"type"`
	JSON accessIdentityProviderGetResponseAccessCentrifyJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessCentrifyJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessCentrify]
type accessIdentityProviderGetResponseAccessCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessCentrify) implementsAccessIdentityProviderGetResponse() {
}

type AccessIdentityProviderGetResponseAccessCentrifyConfig struct {
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
	EmailClaimName string                                                    `json:"email_claim_name"`
	JSON           accessIdentityProviderGetResponseAccessCentrifyConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessCentrifyConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseAccessCentrifyConfig]
type accessIdentityProviderGetResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseAccessCentrifyScimConfig struct {
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
	UserDeprovision bool                                                          `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessCentrifyScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseAccessCentrifyScimConfig]
type accessIdentityProviderGetResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseAccessCentrifyType string

const (
	AccessIdentityProviderGetResponseAccessCentrifyTypeOnetimepin AccessIdentityProviderGetResponseAccessCentrifyType = "onetimepin"
	AccessIdentityProviderGetResponseAccessCentrifyTypeAzureAd    AccessIdentityProviderGetResponseAccessCentrifyType = "azureAD"
	AccessIdentityProviderGetResponseAccessCentrifyTypeSaml       AccessIdentityProviderGetResponseAccessCentrifyType = "saml"
	AccessIdentityProviderGetResponseAccessCentrifyTypeCentrify   AccessIdentityProviderGetResponseAccessCentrifyType = "centrify"
	AccessIdentityProviderGetResponseAccessCentrifyTypeFacebook   AccessIdentityProviderGetResponseAccessCentrifyType = "facebook"
	AccessIdentityProviderGetResponseAccessCentrifyTypeGitHub     AccessIdentityProviderGetResponseAccessCentrifyType = "github"
	AccessIdentityProviderGetResponseAccessCentrifyTypeGoogleApps AccessIdentityProviderGetResponseAccessCentrifyType = "google-apps"
	AccessIdentityProviderGetResponseAccessCentrifyTypeGoogle     AccessIdentityProviderGetResponseAccessCentrifyType = "google"
	AccessIdentityProviderGetResponseAccessCentrifyTypeLinkedin   AccessIdentityProviderGetResponseAccessCentrifyType = "linkedin"
	AccessIdentityProviderGetResponseAccessCentrifyTypeOidc       AccessIdentityProviderGetResponseAccessCentrifyType = "oidc"
	AccessIdentityProviderGetResponseAccessCentrifyTypeOkta       AccessIdentityProviderGetResponseAccessCentrifyType = "okta"
	AccessIdentityProviderGetResponseAccessCentrifyTypeOnelogin   AccessIdentityProviderGetResponseAccessCentrifyType = "onelogin"
	AccessIdentityProviderGetResponseAccessCentrifyTypePingone    AccessIdentityProviderGetResponseAccessCentrifyType = "pingone"
	AccessIdentityProviderGetResponseAccessCentrifyTypeYandex     AccessIdentityProviderGetResponseAccessCentrifyType = "yandex"
)

type AccessIdentityProviderGetResponseAccessFacebook struct {
	// UUID
	ID     string                                                `json:"id"`
	Config AccessIdentityProviderGetResponseAccessFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessFacebookType `json:"type"`
	JSON accessIdentityProviderGetResponseAccessFacebookJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessFacebookJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessFacebook]
type accessIdentityProviderGetResponseAccessFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessFacebook) implementsAccessIdentityProviderGetResponse() {
}

type AccessIdentityProviderGetResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                    `json:"client_secret"`
	JSON         accessIdentityProviderGetResponseAccessFacebookConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessFacebookConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseAccessFacebookConfig]
type accessIdentityProviderGetResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseAccessFacebookScimConfig struct {
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
	UserDeprovision bool                                                          `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseAccessFacebookScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessFacebookScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseAccessFacebookScimConfig]
type accessIdentityProviderGetResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseAccessFacebookType string

const (
	AccessIdentityProviderGetResponseAccessFacebookTypeOnetimepin AccessIdentityProviderGetResponseAccessFacebookType = "onetimepin"
	AccessIdentityProviderGetResponseAccessFacebookTypeAzureAd    AccessIdentityProviderGetResponseAccessFacebookType = "azureAD"
	AccessIdentityProviderGetResponseAccessFacebookTypeSaml       AccessIdentityProviderGetResponseAccessFacebookType = "saml"
	AccessIdentityProviderGetResponseAccessFacebookTypeCentrify   AccessIdentityProviderGetResponseAccessFacebookType = "centrify"
	AccessIdentityProviderGetResponseAccessFacebookTypeFacebook   AccessIdentityProviderGetResponseAccessFacebookType = "facebook"
	AccessIdentityProviderGetResponseAccessFacebookTypeGitHub     AccessIdentityProviderGetResponseAccessFacebookType = "github"
	AccessIdentityProviderGetResponseAccessFacebookTypeGoogleApps AccessIdentityProviderGetResponseAccessFacebookType = "google-apps"
	AccessIdentityProviderGetResponseAccessFacebookTypeGoogle     AccessIdentityProviderGetResponseAccessFacebookType = "google"
	AccessIdentityProviderGetResponseAccessFacebookTypeLinkedin   AccessIdentityProviderGetResponseAccessFacebookType = "linkedin"
	AccessIdentityProviderGetResponseAccessFacebookTypeOidc       AccessIdentityProviderGetResponseAccessFacebookType = "oidc"
	AccessIdentityProviderGetResponseAccessFacebookTypeOkta       AccessIdentityProviderGetResponseAccessFacebookType = "okta"
	AccessIdentityProviderGetResponseAccessFacebookTypeOnelogin   AccessIdentityProviderGetResponseAccessFacebookType = "onelogin"
	AccessIdentityProviderGetResponseAccessFacebookTypePingone    AccessIdentityProviderGetResponseAccessFacebookType = "pingone"
	AccessIdentityProviderGetResponseAccessFacebookTypeYandex     AccessIdentityProviderGetResponseAccessFacebookType = "yandex"
)

type AccessIdentityProviderGetResponseAccessGitHub struct {
	// UUID
	ID     string                                              `json:"id"`
	Config AccessIdentityProviderGetResponseAccessGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessGitHubType `json:"type"`
	JSON accessIdentityProviderGetResponseAccessGitHubJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessGitHubJSON contains the JSON metadata for
// the struct [AccessIdentityProviderGetResponseAccessGitHub]
type accessIdentityProviderGetResponseAccessGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessGitHub) implementsAccessIdentityProviderGetResponse() {
}

type AccessIdentityProviderGetResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                  `json:"client_secret"`
	JSON         accessIdentityProviderGetResponseAccessGitHubConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessGitHubConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseAccessGitHubConfig]
type accessIdentityProviderGetResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseAccessGitHubScimConfig struct {
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
	JSON            accessIdentityProviderGetResponseAccessGitHubScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessGitHubScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseAccessGitHubScimConfig]
type accessIdentityProviderGetResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseAccessGitHubType string

const (
	AccessIdentityProviderGetResponseAccessGitHubTypeOnetimepin AccessIdentityProviderGetResponseAccessGitHubType = "onetimepin"
	AccessIdentityProviderGetResponseAccessGitHubTypeAzureAd    AccessIdentityProviderGetResponseAccessGitHubType = "azureAD"
	AccessIdentityProviderGetResponseAccessGitHubTypeSaml       AccessIdentityProviderGetResponseAccessGitHubType = "saml"
	AccessIdentityProviderGetResponseAccessGitHubTypeCentrify   AccessIdentityProviderGetResponseAccessGitHubType = "centrify"
	AccessIdentityProviderGetResponseAccessGitHubTypeFacebook   AccessIdentityProviderGetResponseAccessGitHubType = "facebook"
	AccessIdentityProviderGetResponseAccessGitHubTypeGitHub     AccessIdentityProviderGetResponseAccessGitHubType = "github"
	AccessIdentityProviderGetResponseAccessGitHubTypeGoogleApps AccessIdentityProviderGetResponseAccessGitHubType = "google-apps"
	AccessIdentityProviderGetResponseAccessGitHubTypeGoogle     AccessIdentityProviderGetResponseAccessGitHubType = "google"
	AccessIdentityProviderGetResponseAccessGitHubTypeLinkedin   AccessIdentityProviderGetResponseAccessGitHubType = "linkedin"
	AccessIdentityProviderGetResponseAccessGitHubTypeOidc       AccessIdentityProviderGetResponseAccessGitHubType = "oidc"
	AccessIdentityProviderGetResponseAccessGitHubTypeOkta       AccessIdentityProviderGetResponseAccessGitHubType = "okta"
	AccessIdentityProviderGetResponseAccessGitHubTypeOnelogin   AccessIdentityProviderGetResponseAccessGitHubType = "onelogin"
	AccessIdentityProviderGetResponseAccessGitHubTypePingone    AccessIdentityProviderGetResponseAccessGitHubType = "pingone"
	AccessIdentityProviderGetResponseAccessGitHubTypeYandex     AccessIdentityProviderGetResponseAccessGitHubType = "yandex"
)

type AccessIdentityProviderGetResponseAccessGoogle struct {
	// UUID
	ID     string                                              `json:"id"`
	Config AccessIdentityProviderGetResponseAccessGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessGoogleType `json:"type"`
	JSON accessIdentityProviderGetResponseAccessGoogleJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessGoogleJSON contains the JSON metadata for
// the struct [AccessIdentityProviderGetResponseAccessGoogle]
type accessIdentityProviderGetResponseAccessGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessGoogle) implementsAccessIdentityProviderGetResponse() {
}

type AccessIdentityProviderGetResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                  `json:"email_claim_name"`
	JSON           accessIdentityProviderGetResponseAccessGoogleConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessGoogleConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseAccessGoogleConfig]
type accessIdentityProviderGetResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseAccessGoogleScimConfig struct {
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
	JSON            accessIdentityProviderGetResponseAccessGoogleScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessGoogleScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseAccessGoogleScimConfig]
type accessIdentityProviderGetResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseAccessGoogleType string

const (
	AccessIdentityProviderGetResponseAccessGoogleTypeOnetimepin AccessIdentityProviderGetResponseAccessGoogleType = "onetimepin"
	AccessIdentityProviderGetResponseAccessGoogleTypeAzureAd    AccessIdentityProviderGetResponseAccessGoogleType = "azureAD"
	AccessIdentityProviderGetResponseAccessGoogleTypeSaml       AccessIdentityProviderGetResponseAccessGoogleType = "saml"
	AccessIdentityProviderGetResponseAccessGoogleTypeCentrify   AccessIdentityProviderGetResponseAccessGoogleType = "centrify"
	AccessIdentityProviderGetResponseAccessGoogleTypeFacebook   AccessIdentityProviderGetResponseAccessGoogleType = "facebook"
	AccessIdentityProviderGetResponseAccessGoogleTypeGitHub     AccessIdentityProviderGetResponseAccessGoogleType = "github"
	AccessIdentityProviderGetResponseAccessGoogleTypeGoogleApps AccessIdentityProviderGetResponseAccessGoogleType = "google-apps"
	AccessIdentityProviderGetResponseAccessGoogleTypeGoogle     AccessIdentityProviderGetResponseAccessGoogleType = "google"
	AccessIdentityProviderGetResponseAccessGoogleTypeLinkedin   AccessIdentityProviderGetResponseAccessGoogleType = "linkedin"
	AccessIdentityProviderGetResponseAccessGoogleTypeOidc       AccessIdentityProviderGetResponseAccessGoogleType = "oidc"
	AccessIdentityProviderGetResponseAccessGoogleTypeOkta       AccessIdentityProviderGetResponseAccessGoogleType = "okta"
	AccessIdentityProviderGetResponseAccessGoogleTypeOnelogin   AccessIdentityProviderGetResponseAccessGoogleType = "onelogin"
	AccessIdentityProviderGetResponseAccessGoogleTypePingone    AccessIdentityProviderGetResponseAccessGoogleType = "pingone"
	AccessIdentityProviderGetResponseAccessGoogleTypeYandex     AccessIdentityProviderGetResponseAccessGoogleType = "yandex"
)

type AccessIdentityProviderGetResponseAccessGoogleApps struct {
	// UUID
	ID     string                                                  `json:"id"`
	Config AccessIdentityProviderGetResponseAccessGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessGoogleAppsType `json:"type"`
	JSON accessIdentityProviderGetResponseAccessGoogleAppsJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessGoogleAppsJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessGoogleApps]
type accessIdentityProviderGetResponseAccessGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessGoogleApps) implementsAccessIdentityProviderGetResponse() {
}

type AccessIdentityProviderGetResponseAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                      `json:"email_claim_name"`
	JSON           accessIdentityProviderGetResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessGoogleAppsConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseAccessGoogleAppsConfig]
type accessIdentityProviderGetResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseAccessGoogleAppsScimConfig struct {
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
	JSON            accessIdentityProviderGetResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessGoogleAppsScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseAccessGoogleAppsScimConfig]
type accessIdentityProviderGetResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseAccessGoogleAppsType string

const (
	AccessIdentityProviderGetResponseAccessGoogleAppsTypeOnetimepin AccessIdentityProviderGetResponseAccessGoogleAppsType = "onetimepin"
	AccessIdentityProviderGetResponseAccessGoogleAppsTypeAzureAd    AccessIdentityProviderGetResponseAccessGoogleAppsType = "azureAD"
	AccessIdentityProviderGetResponseAccessGoogleAppsTypeSaml       AccessIdentityProviderGetResponseAccessGoogleAppsType = "saml"
	AccessIdentityProviderGetResponseAccessGoogleAppsTypeCentrify   AccessIdentityProviderGetResponseAccessGoogleAppsType = "centrify"
	AccessIdentityProviderGetResponseAccessGoogleAppsTypeFacebook   AccessIdentityProviderGetResponseAccessGoogleAppsType = "facebook"
	AccessIdentityProviderGetResponseAccessGoogleAppsTypeGitHub     AccessIdentityProviderGetResponseAccessGoogleAppsType = "github"
	AccessIdentityProviderGetResponseAccessGoogleAppsTypeGoogleApps AccessIdentityProviderGetResponseAccessGoogleAppsType = "google-apps"
	AccessIdentityProviderGetResponseAccessGoogleAppsTypeGoogle     AccessIdentityProviderGetResponseAccessGoogleAppsType = "google"
	AccessIdentityProviderGetResponseAccessGoogleAppsTypeLinkedin   AccessIdentityProviderGetResponseAccessGoogleAppsType = "linkedin"
	AccessIdentityProviderGetResponseAccessGoogleAppsTypeOidc       AccessIdentityProviderGetResponseAccessGoogleAppsType = "oidc"
	AccessIdentityProviderGetResponseAccessGoogleAppsTypeOkta       AccessIdentityProviderGetResponseAccessGoogleAppsType = "okta"
	AccessIdentityProviderGetResponseAccessGoogleAppsTypeOnelogin   AccessIdentityProviderGetResponseAccessGoogleAppsType = "onelogin"
	AccessIdentityProviderGetResponseAccessGoogleAppsTypePingone    AccessIdentityProviderGetResponseAccessGoogleAppsType = "pingone"
	AccessIdentityProviderGetResponseAccessGoogleAppsTypeYandex     AccessIdentityProviderGetResponseAccessGoogleAppsType = "yandex"
)

type AccessIdentityProviderGetResponseAccessLinkedin struct {
	// UUID
	ID     string                                                `json:"id"`
	Config AccessIdentityProviderGetResponseAccessLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessLinkedinType `json:"type"`
	JSON accessIdentityProviderGetResponseAccessLinkedinJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessLinkedinJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessLinkedin]
type accessIdentityProviderGetResponseAccessLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessLinkedin) implementsAccessIdentityProviderGetResponse() {
}

type AccessIdentityProviderGetResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                    `json:"client_secret"`
	JSON         accessIdentityProviderGetResponseAccessLinkedinConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessLinkedinConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseAccessLinkedinConfig]
type accessIdentityProviderGetResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseAccessLinkedinScimConfig struct {
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
	UserDeprovision bool                                                          `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessLinkedinScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseAccessLinkedinScimConfig]
type accessIdentityProviderGetResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseAccessLinkedinType string

const (
	AccessIdentityProviderGetResponseAccessLinkedinTypeOnetimepin AccessIdentityProviderGetResponseAccessLinkedinType = "onetimepin"
	AccessIdentityProviderGetResponseAccessLinkedinTypeAzureAd    AccessIdentityProviderGetResponseAccessLinkedinType = "azureAD"
	AccessIdentityProviderGetResponseAccessLinkedinTypeSaml       AccessIdentityProviderGetResponseAccessLinkedinType = "saml"
	AccessIdentityProviderGetResponseAccessLinkedinTypeCentrify   AccessIdentityProviderGetResponseAccessLinkedinType = "centrify"
	AccessIdentityProviderGetResponseAccessLinkedinTypeFacebook   AccessIdentityProviderGetResponseAccessLinkedinType = "facebook"
	AccessIdentityProviderGetResponseAccessLinkedinTypeGitHub     AccessIdentityProviderGetResponseAccessLinkedinType = "github"
	AccessIdentityProviderGetResponseAccessLinkedinTypeGoogleApps AccessIdentityProviderGetResponseAccessLinkedinType = "google-apps"
	AccessIdentityProviderGetResponseAccessLinkedinTypeGoogle     AccessIdentityProviderGetResponseAccessLinkedinType = "google"
	AccessIdentityProviderGetResponseAccessLinkedinTypeLinkedin   AccessIdentityProviderGetResponseAccessLinkedinType = "linkedin"
	AccessIdentityProviderGetResponseAccessLinkedinTypeOidc       AccessIdentityProviderGetResponseAccessLinkedinType = "oidc"
	AccessIdentityProviderGetResponseAccessLinkedinTypeOkta       AccessIdentityProviderGetResponseAccessLinkedinType = "okta"
	AccessIdentityProviderGetResponseAccessLinkedinTypeOnelogin   AccessIdentityProviderGetResponseAccessLinkedinType = "onelogin"
	AccessIdentityProviderGetResponseAccessLinkedinTypePingone    AccessIdentityProviderGetResponseAccessLinkedinType = "pingone"
	AccessIdentityProviderGetResponseAccessLinkedinTypeYandex     AccessIdentityProviderGetResponseAccessLinkedinType = "yandex"
)

type AccessIdentityProviderGetResponseAccessOidc struct {
	// UUID
	ID     string                                            `json:"id"`
	Config AccessIdentityProviderGetResponseAccessOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessOidcType `json:"type"`
	JSON accessIdentityProviderGetResponseAccessOidcJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessOidcJSON contains the JSON metadata for
// the struct [AccessIdentityProviderGetResponseAccessOidc]
type accessIdentityProviderGetResponseAccessOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessOidc) implementsAccessIdentityProviderGetResponse() {}

type AccessIdentityProviderGetResponseAccessOidcConfig struct {
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
	TokenURL string                                                `json:"token_url"`
	JSON     accessIdentityProviderGetResponseAccessOidcConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessOidcConfigJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessOidcConfig]
type accessIdentityProviderGetResponseAccessOidcConfigJSON struct {
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

func (r *AccessIdentityProviderGetResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseAccessOidcScimConfig struct {
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
	JSON            accessIdentityProviderGetResponseAccessOidcScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessOidcScimConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseAccessOidcScimConfig]
type accessIdentityProviderGetResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseAccessOidcType string

const (
	AccessIdentityProviderGetResponseAccessOidcTypeOnetimepin AccessIdentityProviderGetResponseAccessOidcType = "onetimepin"
	AccessIdentityProviderGetResponseAccessOidcTypeAzureAd    AccessIdentityProviderGetResponseAccessOidcType = "azureAD"
	AccessIdentityProviderGetResponseAccessOidcTypeSaml       AccessIdentityProviderGetResponseAccessOidcType = "saml"
	AccessIdentityProviderGetResponseAccessOidcTypeCentrify   AccessIdentityProviderGetResponseAccessOidcType = "centrify"
	AccessIdentityProviderGetResponseAccessOidcTypeFacebook   AccessIdentityProviderGetResponseAccessOidcType = "facebook"
	AccessIdentityProviderGetResponseAccessOidcTypeGitHub     AccessIdentityProviderGetResponseAccessOidcType = "github"
	AccessIdentityProviderGetResponseAccessOidcTypeGoogleApps AccessIdentityProviderGetResponseAccessOidcType = "google-apps"
	AccessIdentityProviderGetResponseAccessOidcTypeGoogle     AccessIdentityProviderGetResponseAccessOidcType = "google"
	AccessIdentityProviderGetResponseAccessOidcTypeLinkedin   AccessIdentityProviderGetResponseAccessOidcType = "linkedin"
	AccessIdentityProviderGetResponseAccessOidcTypeOidc       AccessIdentityProviderGetResponseAccessOidcType = "oidc"
	AccessIdentityProviderGetResponseAccessOidcTypeOkta       AccessIdentityProviderGetResponseAccessOidcType = "okta"
	AccessIdentityProviderGetResponseAccessOidcTypeOnelogin   AccessIdentityProviderGetResponseAccessOidcType = "onelogin"
	AccessIdentityProviderGetResponseAccessOidcTypePingone    AccessIdentityProviderGetResponseAccessOidcType = "pingone"
	AccessIdentityProviderGetResponseAccessOidcTypeYandex     AccessIdentityProviderGetResponseAccessOidcType = "yandex"
)

type AccessIdentityProviderGetResponseAccessOkta struct {
	// UUID
	ID     string                                            `json:"id"`
	Config AccessIdentityProviderGetResponseAccessOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessOktaType `json:"type"`
	JSON accessIdentityProviderGetResponseAccessOktaJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessOktaJSON contains the JSON metadata for
// the struct [AccessIdentityProviderGetResponseAccessOkta]
type accessIdentityProviderGetResponseAccessOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessOkta) implementsAccessIdentityProviderGetResponse() {}

type AccessIdentityProviderGetResponseAccessOktaConfig struct {
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
	OktaAccount string                                                `json:"okta_account"`
	JSON        accessIdentityProviderGetResponseAccessOktaConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessOktaConfigJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessOktaConfig]
type accessIdentityProviderGetResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseAccessOktaScimConfig struct {
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
	JSON            accessIdentityProviderGetResponseAccessOktaScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessOktaScimConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseAccessOktaScimConfig]
type accessIdentityProviderGetResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseAccessOktaType string

const (
	AccessIdentityProviderGetResponseAccessOktaTypeOnetimepin AccessIdentityProviderGetResponseAccessOktaType = "onetimepin"
	AccessIdentityProviderGetResponseAccessOktaTypeAzureAd    AccessIdentityProviderGetResponseAccessOktaType = "azureAD"
	AccessIdentityProviderGetResponseAccessOktaTypeSaml       AccessIdentityProviderGetResponseAccessOktaType = "saml"
	AccessIdentityProviderGetResponseAccessOktaTypeCentrify   AccessIdentityProviderGetResponseAccessOktaType = "centrify"
	AccessIdentityProviderGetResponseAccessOktaTypeFacebook   AccessIdentityProviderGetResponseAccessOktaType = "facebook"
	AccessIdentityProviderGetResponseAccessOktaTypeGitHub     AccessIdentityProviderGetResponseAccessOktaType = "github"
	AccessIdentityProviderGetResponseAccessOktaTypeGoogleApps AccessIdentityProviderGetResponseAccessOktaType = "google-apps"
	AccessIdentityProviderGetResponseAccessOktaTypeGoogle     AccessIdentityProviderGetResponseAccessOktaType = "google"
	AccessIdentityProviderGetResponseAccessOktaTypeLinkedin   AccessIdentityProviderGetResponseAccessOktaType = "linkedin"
	AccessIdentityProviderGetResponseAccessOktaTypeOidc       AccessIdentityProviderGetResponseAccessOktaType = "oidc"
	AccessIdentityProviderGetResponseAccessOktaTypeOkta       AccessIdentityProviderGetResponseAccessOktaType = "okta"
	AccessIdentityProviderGetResponseAccessOktaTypeOnelogin   AccessIdentityProviderGetResponseAccessOktaType = "onelogin"
	AccessIdentityProviderGetResponseAccessOktaTypePingone    AccessIdentityProviderGetResponseAccessOktaType = "pingone"
	AccessIdentityProviderGetResponseAccessOktaTypeYandex     AccessIdentityProviderGetResponseAccessOktaType = "yandex"
)

type AccessIdentityProviderGetResponseAccessOnelogin struct {
	// UUID
	ID     string                                                `json:"id"`
	Config AccessIdentityProviderGetResponseAccessOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessOneloginType `json:"type"`
	JSON accessIdentityProviderGetResponseAccessOneloginJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessOneloginJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessOnelogin]
type accessIdentityProviderGetResponseAccessOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessOnelogin) implementsAccessIdentityProviderGetResponse() {
}

type AccessIdentityProviderGetResponseAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                    `json:"onelogin_account"`
	JSON            accessIdentityProviderGetResponseAccessOneloginConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessOneloginConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseAccessOneloginConfig]
type accessIdentityProviderGetResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseAccessOneloginScimConfig struct {
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
	UserDeprovision bool                                                          `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseAccessOneloginScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessOneloginScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseAccessOneloginScimConfig]
type accessIdentityProviderGetResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseAccessOneloginType string

const (
	AccessIdentityProviderGetResponseAccessOneloginTypeOnetimepin AccessIdentityProviderGetResponseAccessOneloginType = "onetimepin"
	AccessIdentityProviderGetResponseAccessOneloginTypeAzureAd    AccessIdentityProviderGetResponseAccessOneloginType = "azureAD"
	AccessIdentityProviderGetResponseAccessOneloginTypeSaml       AccessIdentityProviderGetResponseAccessOneloginType = "saml"
	AccessIdentityProviderGetResponseAccessOneloginTypeCentrify   AccessIdentityProviderGetResponseAccessOneloginType = "centrify"
	AccessIdentityProviderGetResponseAccessOneloginTypeFacebook   AccessIdentityProviderGetResponseAccessOneloginType = "facebook"
	AccessIdentityProviderGetResponseAccessOneloginTypeGitHub     AccessIdentityProviderGetResponseAccessOneloginType = "github"
	AccessIdentityProviderGetResponseAccessOneloginTypeGoogleApps AccessIdentityProviderGetResponseAccessOneloginType = "google-apps"
	AccessIdentityProviderGetResponseAccessOneloginTypeGoogle     AccessIdentityProviderGetResponseAccessOneloginType = "google"
	AccessIdentityProviderGetResponseAccessOneloginTypeLinkedin   AccessIdentityProviderGetResponseAccessOneloginType = "linkedin"
	AccessIdentityProviderGetResponseAccessOneloginTypeOidc       AccessIdentityProviderGetResponseAccessOneloginType = "oidc"
	AccessIdentityProviderGetResponseAccessOneloginTypeOkta       AccessIdentityProviderGetResponseAccessOneloginType = "okta"
	AccessIdentityProviderGetResponseAccessOneloginTypeOnelogin   AccessIdentityProviderGetResponseAccessOneloginType = "onelogin"
	AccessIdentityProviderGetResponseAccessOneloginTypePingone    AccessIdentityProviderGetResponseAccessOneloginType = "pingone"
	AccessIdentityProviderGetResponseAccessOneloginTypeYandex     AccessIdentityProviderGetResponseAccessOneloginType = "yandex"
)

type AccessIdentityProviderGetResponseAccessPingone struct {
	// UUID
	ID     string                                               `json:"id"`
	Config AccessIdentityProviderGetResponseAccessPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessPingoneType `json:"type"`
	JSON accessIdentityProviderGetResponseAccessPingoneJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessPingoneJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessPingone]
type accessIdentityProviderGetResponseAccessPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessPingone) implementsAccessIdentityProviderGetResponse() {
}

type AccessIdentityProviderGetResponseAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                   `json:"ping_env_id"`
	JSON      accessIdentityProviderGetResponseAccessPingoneConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessPingoneConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseAccessPingoneConfig]
type accessIdentityProviderGetResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseAccessPingoneScimConfig struct {
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
	UserDeprovision bool                                                         `json:"user_deprovision"`
	JSON            accessIdentityProviderGetResponseAccessPingoneScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessPingoneScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseAccessPingoneScimConfig]
type accessIdentityProviderGetResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseAccessPingoneType string

const (
	AccessIdentityProviderGetResponseAccessPingoneTypeOnetimepin AccessIdentityProviderGetResponseAccessPingoneType = "onetimepin"
	AccessIdentityProviderGetResponseAccessPingoneTypeAzureAd    AccessIdentityProviderGetResponseAccessPingoneType = "azureAD"
	AccessIdentityProviderGetResponseAccessPingoneTypeSaml       AccessIdentityProviderGetResponseAccessPingoneType = "saml"
	AccessIdentityProviderGetResponseAccessPingoneTypeCentrify   AccessIdentityProviderGetResponseAccessPingoneType = "centrify"
	AccessIdentityProviderGetResponseAccessPingoneTypeFacebook   AccessIdentityProviderGetResponseAccessPingoneType = "facebook"
	AccessIdentityProviderGetResponseAccessPingoneTypeGitHub     AccessIdentityProviderGetResponseAccessPingoneType = "github"
	AccessIdentityProviderGetResponseAccessPingoneTypeGoogleApps AccessIdentityProviderGetResponseAccessPingoneType = "google-apps"
	AccessIdentityProviderGetResponseAccessPingoneTypeGoogle     AccessIdentityProviderGetResponseAccessPingoneType = "google"
	AccessIdentityProviderGetResponseAccessPingoneTypeLinkedin   AccessIdentityProviderGetResponseAccessPingoneType = "linkedin"
	AccessIdentityProviderGetResponseAccessPingoneTypeOidc       AccessIdentityProviderGetResponseAccessPingoneType = "oidc"
	AccessIdentityProviderGetResponseAccessPingoneTypeOkta       AccessIdentityProviderGetResponseAccessPingoneType = "okta"
	AccessIdentityProviderGetResponseAccessPingoneTypeOnelogin   AccessIdentityProviderGetResponseAccessPingoneType = "onelogin"
	AccessIdentityProviderGetResponseAccessPingoneTypePingone    AccessIdentityProviderGetResponseAccessPingoneType = "pingone"
	AccessIdentityProviderGetResponseAccessPingoneTypeYandex     AccessIdentityProviderGetResponseAccessPingoneType = "yandex"
)

type AccessIdentityProviderGetResponseAccessSaml struct {
	// UUID
	ID     string                                            `json:"id"`
	Config AccessIdentityProviderGetResponseAccessSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessSamlType `json:"type"`
	JSON accessIdentityProviderGetResponseAccessSamlJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessSamlJSON contains the JSON metadata for
// the struct [AccessIdentityProviderGetResponseAccessSaml]
type accessIdentityProviderGetResponseAccessSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessSaml) implementsAccessIdentityProviderGetResponse() {}

type AccessIdentityProviderGetResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccessIdentityProviderGetResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                `json:"sso_target_url"`
	JSON         accessIdentityProviderGetResponseAccessSamlConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessSamlConfigJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessSamlConfig]
type accessIdentityProviderGetResponseAccessSamlConfigJSON struct {
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

func (r *AccessIdentityProviderGetResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderGetResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                               `json:"header_name"`
	JSON       accessIdentityProviderGetResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessSamlConfigHeaderAttributeJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderGetResponseAccessSamlConfigHeaderAttribute]
type accessIdentityProviderGetResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseAccessSamlScimConfig struct {
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
	JSON            accessIdentityProviderGetResponseAccessSamlScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessSamlScimConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseAccessSamlScimConfig]
type accessIdentityProviderGetResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseAccessSamlType string

const (
	AccessIdentityProviderGetResponseAccessSamlTypeOnetimepin AccessIdentityProviderGetResponseAccessSamlType = "onetimepin"
	AccessIdentityProviderGetResponseAccessSamlTypeAzureAd    AccessIdentityProviderGetResponseAccessSamlType = "azureAD"
	AccessIdentityProviderGetResponseAccessSamlTypeSaml       AccessIdentityProviderGetResponseAccessSamlType = "saml"
	AccessIdentityProviderGetResponseAccessSamlTypeCentrify   AccessIdentityProviderGetResponseAccessSamlType = "centrify"
	AccessIdentityProviderGetResponseAccessSamlTypeFacebook   AccessIdentityProviderGetResponseAccessSamlType = "facebook"
	AccessIdentityProviderGetResponseAccessSamlTypeGitHub     AccessIdentityProviderGetResponseAccessSamlType = "github"
	AccessIdentityProviderGetResponseAccessSamlTypeGoogleApps AccessIdentityProviderGetResponseAccessSamlType = "google-apps"
	AccessIdentityProviderGetResponseAccessSamlTypeGoogle     AccessIdentityProviderGetResponseAccessSamlType = "google"
	AccessIdentityProviderGetResponseAccessSamlTypeLinkedin   AccessIdentityProviderGetResponseAccessSamlType = "linkedin"
	AccessIdentityProviderGetResponseAccessSamlTypeOidc       AccessIdentityProviderGetResponseAccessSamlType = "oidc"
	AccessIdentityProviderGetResponseAccessSamlTypeOkta       AccessIdentityProviderGetResponseAccessSamlType = "okta"
	AccessIdentityProviderGetResponseAccessSamlTypeOnelogin   AccessIdentityProviderGetResponseAccessSamlType = "onelogin"
	AccessIdentityProviderGetResponseAccessSamlTypePingone    AccessIdentityProviderGetResponseAccessSamlType = "pingone"
	AccessIdentityProviderGetResponseAccessSamlTypeYandex     AccessIdentityProviderGetResponseAccessSamlType = "yandex"
)

type AccessIdentityProviderGetResponseAccessYandex struct {
	// UUID
	ID     string                                              `json:"id"`
	Config AccessIdentityProviderGetResponseAccessYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessYandexType `json:"type"`
	JSON accessIdentityProviderGetResponseAccessYandexJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessYandexJSON contains the JSON metadata for
// the struct [AccessIdentityProviderGetResponseAccessYandex]
type accessIdentityProviderGetResponseAccessYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessYandex) implementsAccessIdentityProviderGetResponse() {
}

type AccessIdentityProviderGetResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                  `json:"client_secret"`
	JSON         accessIdentityProviderGetResponseAccessYandexConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessYandexConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderGetResponseAccessYandexConfig]
type accessIdentityProviderGetResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseAccessYandexScimConfig struct {
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
	JSON            accessIdentityProviderGetResponseAccessYandexScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessYandexScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderGetResponseAccessYandexScimConfig]
type accessIdentityProviderGetResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseAccessYandexType string

const (
	AccessIdentityProviderGetResponseAccessYandexTypeOnetimepin AccessIdentityProviderGetResponseAccessYandexType = "onetimepin"
	AccessIdentityProviderGetResponseAccessYandexTypeAzureAd    AccessIdentityProviderGetResponseAccessYandexType = "azureAD"
	AccessIdentityProviderGetResponseAccessYandexTypeSaml       AccessIdentityProviderGetResponseAccessYandexType = "saml"
	AccessIdentityProviderGetResponseAccessYandexTypeCentrify   AccessIdentityProviderGetResponseAccessYandexType = "centrify"
	AccessIdentityProviderGetResponseAccessYandexTypeFacebook   AccessIdentityProviderGetResponseAccessYandexType = "facebook"
	AccessIdentityProviderGetResponseAccessYandexTypeGitHub     AccessIdentityProviderGetResponseAccessYandexType = "github"
	AccessIdentityProviderGetResponseAccessYandexTypeGoogleApps AccessIdentityProviderGetResponseAccessYandexType = "google-apps"
	AccessIdentityProviderGetResponseAccessYandexTypeGoogle     AccessIdentityProviderGetResponseAccessYandexType = "google"
	AccessIdentityProviderGetResponseAccessYandexTypeLinkedin   AccessIdentityProviderGetResponseAccessYandexType = "linkedin"
	AccessIdentityProviderGetResponseAccessYandexTypeOidc       AccessIdentityProviderGetResponseAccessYandexType = "oidc"
	AccessIdentityProviderGetResponseAccessYandexTypeOkta       AccessIdentityProviderGetResponseAccessYandexType = "okta"
	AccessIdentityProviderGetResponseAccessYandexTypeOnelogin   AccessIdentityProviderGetResponseAccessYandexType = "onelogin"
	AccessIdentityProviderGetResponseAccessYandexTypePingone    AccessIdentityProviderGetResponseAccessYandexType = "pingone"
	AccessIdentityProviderGetResponseAccessYandexTypeYandex     AccessIdentityProviderGetResponseAccessYandexType = "yandex"
)

type AccessIdentityProviderGetResponseAccessOnetimepin struct {
	// UUID
	ID     string      `json:"id"`
	Config interface{} `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessOnetimepinScimConfig `json:"scim_config"`
	Type       AccessIdentityProviderGetResponseAccessOnetimepinType       `json:"type"`
	JSON       accessIdentityProviderGetResponseAccessOnetimepinJSON       `json:"-"`
}

// accessIdentityProviderGetResponseAccessOnetimepinJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessOnetimepin]
type accessIdentityProviderGetResponseAccessOnetimepinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessOnetimepin) implementsAccessIdentityProviderGetResponse() {
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderGetResponseAccessOnetimepinScimConfig struct {
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
	JSON            accessIdentityProviderGetResponseAccessOnetimepinScimConfigJSON `json:"-"`
}

// accessIdentityProviderGetResponseAccessOnetimepinScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderGetResponseAccessOnetimepinScimConfig]
type accessIdentityProviderGetResponseAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderGetResponseAccessOnetimepinType string

const (
	AccessIdentityProviderGetResponseAccessOnetimepinTypeOnetimepin AccessIdentityProviderGetResponseAccessOnetimepinType = "onetimepin"
)

// Union satisfied by [AccessIdentityProviderUpdateResponseAccessAzureAd],
// [AccessIdentityProviderUpdateResponseAccessCentrify],
// [AccessIdentityProviderUpdateResponseAccessFacebook],
// [AccessIdentityProviderUpdateResponseAccessGitHub],
// [AccessIdentityProviderUpdateResponseAccessGoogle],
// [AccessIdentityProviderUpdateResponseAccessGoogleApps],
// [AccessIdentityProviderUpdateResponseAccessLinkedin],
// [AccessIdentityProviderUpdateResponseAccessOidc],
// [AccessIdentityProviderUpdateResponseAccessOkta],
// [AccessIdentityProviderUpdateResponseAccessOnelogin],
// [AccessIdentityProviderUpdateResponseAccessPingone],
// [AccessIdentityProviderUpdateResponseAccessSaml],
// [AccessIdentityProviderUpdateResponseAccessYandex] or
// [AccessIdentityProviderUpdateResponseAccessOnetimepin].
type AccessIdentityProviderUpdateResponse interface {
	implementsAccessIdentityProviderUpdateResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessIdentityProviderUpdateResponse)(nil)).Elem(), "")
}

type AccessIdentityProviderUpdateResponseAccessAzureAd struct {
	// UUID
	ID     string                                                  `json:"id"`
	Config AccessIdentityProviderUpdateResponseAccessAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessAzureAdType `json:"type"`
	JSON accessIdentityProviderUpdateResponseAccessAzureAdJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessAzureAdJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessAzureAd]
type accessIdentityProviderUpdateResponseAccessAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessAzureAd) implementsAccessIdentityProviderUpdateResponse() {
}

type AccessIdentityProviderUpdateResponseAccessAzureAdConfig struct {
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
	SupportGroups bool                                                        `json:"support_groups"`
	JSON          accessIdentityProviderUpdateResponseAccessAzureAdConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessAzureAdConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessAzureAdConfig]
type accessIdentityProviderUpdateResponseAccessAzureAdConfigJSON struct {
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

func (r *AccessIdentityProviderUpdateResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseAccessAzureAdScimConfig struct {
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
	JSON            accessIdentityProviderUpdateResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessAzureAdScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessAzureAdScimConfig]
type accessIdentityProviderUpdateResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseAccessAzureAdType string

const (
	AccessIdentityProviderUpdateResponseAccessAzureAdTypeOnetimepin AccessIdentityProviderUpdateResponseAccessAzureAdType = "onetimepin"
	AccessIdentityProviderUpdateResponseAccessAzureAdTypeAzureAd    AccessIdentityProviderUpdateResponseAccessAzureAdType = "azureAD"
	AccessIdentityProviderUpdateResponseAccessAzureAdTypeSaml       AccessIdentityProviderUpdateResponseAccessAzureAdType = "saml"
	AccessIdentityProviderUpdateResponseAccessAzureAdTypeCentrify   AccessIdentityProviderUpdateResponseAccessAzureAdType = "centrify"
	AccessIdentityProviderUpdateResponseAccessAzureAdTypeFacebook   AccessIdentityProviderUpdateResponseAccessAzureAdType = "facebook"
	AccessIdentityProviderUpdateResponseAccessAzureAdTypeGitHub     AccessIdentityProviderUpdateResponseAccessAzureAdType = "github"
	AccessIdentityProviderUpdateResponseAccessAzureAdTypeGoogleApps AccessIdentityProviderUpdateResponseAccessAzureAdType = "google-apps"
	AccessIdentityProviderUpdateResponseAccessAzureAdTypeGoogle     AccessIdentityProviderUpdateResponseAccessAzureAdType = "google"
	AccessIdentityProviderUpdateResponseAccessAzureAdTypeLinkedin   AccessIdentityProviderUpdateResponseAccessAzureAdType = "linkedin"
	AccessIdentityProviderUpdateResponseAccessAzureAdTypeOidc       AccessIdentityProviderUpdateResponseAccessAzureAdType = "oidc"
	AccessIdentityProviderUpdateResponseAccessAzureAdTypeOkta       AccessIdentityProviderUpdateResponseAccessAzureAdType = "okta"
	AccessIdentityProviderUpdateResponseAccessAzureAdTypeOnelogin   AccessIdentityProviderUpdateResponseAccessAzureAdType = "onelogin"
	AccessIdentityProviderUpdateResponseAccessAzureAdTypePingone    AccessIdentityProviderUpdateResponseAccessAzureAdType = "pingone"
	AccessIdentityProviderUpdateResponseAccessAzureAdTypeYandex     AccessIdentityProviderUpdateResponseAccessAzureAdType = "yandex"
)

type AccessIdentityProviderUpdateResponseAccessCentrify struct {
	// UUID
	ID     string                                                   `json:"id"`
	Config AccessIdentityProviderUpdateResponseAccessCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessCentrifyType `json:"type"`
	JSON accessIdentityProviderUpdateResponseAccessCentrifyJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessCentrifyJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessCentrify]
type accessIdentityProviderUpdateResponseAccessCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessCentrify) implementsAccessIdentityProviderUpdateResponse() {
}

type AccessIdentityProviderUpdateResponseAccessCentrifyConfig struct {
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
	EmailClaimName string                                                       `json:"email_claim_name"`
	JSON           accessIdentityProviderUpdateResponseAccessCentrifyConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessCentrifyConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessCentrifyConfig]
type accessIdentityProviderUpdateResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseAccessCentrifyScimConfig struct {
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
	UserDeprovision bool                                                             `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessCentrifyScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessCentrifyScimConfig]
type accessIdentityProviderUpdateResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseAccessCentrifyType string

const (
	AccessIdentityProviderUpdateResponseAccessCentrifyTypeOnetimepin AccessIdentityProviderUpdateResponseAccessCentrifyType = "onetimepin"
	AccessIdentityProviderUpdateResponseAccessCentrifyTypeAzureAd    AccessIdentityProviderUpdateResponseAccessCentrifyType = "azureAD"
	AccessIdentityProviderUpdateResponseAccessCentrifyTypeSaml       AccessIdentityProviderUpdateResponseAccessCentrifyType = "saml"
	AccessIdentityProviderUpdateResponseAccessCentrifyTypeCentrify   AccessIdentityProviderUpdateResponseAccessCentrifyType = "centrify"
	AccessIdentityProviderUpdateResponseAccessCentrifyTypeFacebook   AccessIdentityProviderUpdateResponseAccessCentrifyType = "facebook"
	AccessIdentityProviderUpdateResponseAccessCentrifyTypeGitHub     AccessIdentityProviderUpdateResponseAccessCentrifyType = "github"
	AccessIdentityProviderUpdateResponseAccessCentrifyTypeGoogleApps AccessIdentityProviderUpdateResponseAccessCentrifyType = "google-apps"
	AccessIdentityProviderUpdateResponseAccessCentrifyTypeGoogle     AccessIdentityProviderUpdateResponseAccessCentrifyType = "google"
	AccessIdentityProviderUpdateResponseAccessCentrifyTypeLinkedin   AccessIdentityProviderUpdateResponseAccessCentrifyType = "linkedin"
	AccessIdentityProviderUpdateResponseAccessCentrifyTypeOidc       AccessIdentityProviderUpdateResponseAccessCentrifyType = "oidc"
	AccessIdentityProviderUpdateResponseAccessCentrifyTypeOkta       AccessIdentityProviderUpdateResponseAccessCentrifyType = "okta"
	AccessIdentityProviderUpdateResponseAccessCentrifyTypeOnelogin   AccessIdentityProviderUpdateResponseAccessCentrifyType = "onelogin"
	AccessIdentityProviderUpdateResponseAccessCentrifyTypePingone    AccessIdentityProviderUpdateResponseAccessCentrifyType = "pingone"
	AccessIdentityProviderUpdateResponseAccessCentrifyTypeYandex     AccessIdentityProviderUpdateResponseAccessCentrifyType = "yandex"
)

type AccessIdentityProviderUpdateResponseAccessFacebook struct {
	// UUID
	ID     string                                                   `json:"id"`
	Config AccessIdentityProviderUpdateResponseAccessFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessFacebookType `json:"type"`
	JSON accessIdentityProviderUpdateResponseAccessFacebookJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessFacebookJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessFacebook]
type accessIdentityProviderUpdateResponseAccessFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessFacebook) implementsAccessIdentityProviderUpdateResponse() {
}

type AccessIdentityProviderUpdateResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                       `json:"client_secret"`
	JSON         accessIdentityProviderUpdateResponseAccessFacebookConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessFacebookConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessFacebookConfig]
type accessIdentityProviderUpdateResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseAccessFacebookScimConfig struct {
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
	UserDeprovision bool                                                             `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseAccessFacebookScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessFacebookScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessFacebookScimConfig]
type accessIdentityProviderUpdateResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseAccessFacebookType string

const (
	AccessIdentityProviderUpdateResponseAccessFacebookTypeOnetimepin AccessIdentityProviderUpdateResponseAccessFacebookType = "onetimepin"
	AccessIdentityProviderUpdateResponseAccessFacebookTypeAzureAd    AccessIdentityProviderUpdateResponseAccessFacebookType = "azureAD"
	AccessIdentityProviderUpdateResponseAccessFacebookTypeSaml       AccessIdentityProviderUpdateResponseAccessFacebookType = "saml"
	AccessIdentityProviderUpdateResponseAccessFacebookTypeCentrify   AccessIdentityProviderUpdateResponseAccessFacebookType = "centrify"
	AccessIdentityProviderUpdateResponseAccessFacebookTypeFacebook   AccessIdentityProviderUpdateResponseAccessFacebookType = "facebook"
	AccessIdentityProviderUpdateResponseAccessFacebookTypeGitHub     AccessIdentityProviderUpdateResponseAccessFacebookType = "github"
	AccessIdentityProviderUpdateResponseAccessFacebookTypeGoogleApps AccessIdentityProviderUpdateResponseAccessFacebookType = "google-apps"
	AccessIdentityProviderUpdateResponseAccessFacebookTypeGoogle     AccessIdentityProviderUpdateResponseAccessFacebookType = "google"
	AccessIdentityProviderUpdateResponseAccessFacebookTypeLinkedin   AccessIdentityProviderUpdateResponseAccessFacebookType = "linkedin"
	AccessIdentityProviderUpdateResponseAccessFacebookTypeOidc       AccessIdentityProviderUpdateResponseAccessFacebookType = "oidc"
	AccessIdentityProviderUpdateResponseAccessFacebookTypeOkta       AccessIdentityProviderUpdateResponseAccessFacebookType = "okta"
	AccessIdentityProviderUpdateResponseAccessFacebookTypeOnelogin   AccessIdentityProviderUpdateResponseAccessFacebookType = "onelogin"
	AccessIdentityProviderUpdateResponseAccessFacebookTypePingone    AccessIdentityProviderUpdateResponseAccessFacebookType = "pingone"
	AccessIdentityProviderUpdateResponseAccessFacebookTypeYandex     AccessIdentityProviderUpdateResponseAccessFacebookType = "yandex"
)

type AccessIdentityProviderUpdateResponseAccessGitHub struct {
	// UUID
	ID     string                                                 `json:"id"`
	Config AccessIdentityProviderUpdateResponseAccessGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessGitHubType `json:"type"`
	JSON accessIdentityProviderUpdateResponseAccessGitHubJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessGitHubJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessGitHub]
type accessIdentityProviderUpdateResponseAccessGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessGitHub) implementsAccessIdentityProviderUpdateResponse() {
}

type AccessIdentityProviderUpdateResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                     `json:"client_secret"`
	JSON         accessIdentityProviderUpdateResponseAccessGitHubConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessGitHubConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessGitHubConfig]
type accessIdentityProviderUpdateResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseAccessGitHubScimConfig struct {
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
	UserDeprovision bool                                                           `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseAccessGitHubScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessGitHubScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessGitHubScimConfig]
type accessIdentityProviderUpdateResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseAccessGitHubType string

const (
	AccessIdentityProviderUpdateResponseAccessGitHubTypeOnetimepin AccessIdentityProviderUpdateResponseAccessGitHubType = "onetimepin"
	AccessIdentityProviderUpdateResponseAccessGitHubTypeAzureAd    AccessIdentityProviderUpdateResponseAccessGitHubType = "azureAD"
	AccessIdentityProviderUpdateResponseAccessGitHubTypeSaml       AccessIdentityProviderUpdateResponseAccessGitHubType = "saml"
	AccessIdentityProviderUpdateResponseAccessGitHubTypeCentrify   AccessIdentityProviderUpdateResponseAccessGitHubType = "centrify"
	AccessIdentityProviderUpdateResponseAccessGitHubTypeFacebook   AccessIdentityProviderUpdateResponseAccessGitHubType = "facebook"
	AccessIdentityProviderUpdateResponseAccessGitHubTypeGitHub     AccessIdentityProviderUpdateResponseAccessGitHubType = "github"
	AccessIdentityProviderUpdateResponseAccessGitHubTypeGoogleApps AccessIdentityProviderUpdateResponseAccessGitHubType = "google-apps"
	AccessIdentityProviderUpdateResponseAccessGitHubTypeGoogle     AccessIdentityProviderUpdateResponseAccessGitHubType = "google"
	AccessIdentityProviderUpdateResponseAccessGitHubTypeLinkedin   AccessIdentityProviderUpdateResponseAccessGitHubType = "linkedin"
	AccessIdentityProviderUpdateResponseAccessGitHubTypeOidc       AccessIdentityProviderUpdateResponseAccessGitHubType = "oidc"
	AccessIdentityProviderUpdateResponseAccessGitHubTypeOkta       AccessIdentityProviderUpdateResponseAccessGitHubType = "okta"
	AccessIdentityProviderUpdateResponseAccessGitHubTypeOnelogin   AccessIdentityProviderUpdateResponseAccessGitHubType = "onelogin"
	AccessIdentityProviderUpdateResponseAccessGitHubTypePingone    AccessIdentityProviderUpdateResponseAccessGitHubType = "pingone"
	AccessIdentityProviderUpdateResponseAccessGitHubTypeYandex     AccessIdentityProviderUpdateResponseAccessGitHubType = "yandex"
)

type AccessIdentityProviderUpdateResponseAccessGoogle struct {
	// UUID
	ID     string                                                 `json:"id"`
	Config AccessIdentityProviderUpdateResponseAccessGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessGoogleType `json:"type"`
	JSON accessIdentityProviderUpdateResponseAccessGoogleJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessGoogleJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessGoogle]
type accessIdentityProviderUpdateResponseAccessGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessGoogle) implementsAccessIdentityProviderUpdateResponse() {
}

type AccessIdentityProviderUpdateResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                     `json:"email_claim_name"`
	JSON           accessIdentityProviderUpdateResponseAccessGoogleConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessGoogleConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessGoogleConfig]
type accessIdentityProviderUpdateResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseAccessGoogleScimConfig struct {
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
	UserDeprovision bool                                                           `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseAccessGoogleScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessGoogleScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessGoogleScimConfig]
type accessIdentityProviderUpdateResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseAccessGoogleType string

const (
	AccessIdentityProviderUpdateResponseAccessGoogleTypeOnetimepin AccessIdentityProviderUpdateResponseAccessGoogleType = "onetimepin"
	AccessIdentityProviderUpdateResponseAccessGoogleTypeAzureAd    AccessIdentityProviderUpdateResponseAccessGoogleType = "azureAD"
	AccessIdentityProviderUpdateResponseAccessGoogleTypeSaml       AccessIdentityProviderUpdateResponseAccessGoogleType = "saml"
	AccessIdentityProviderUpdateResponseAccessGoogleTypeCentrify   AccessIdentityProviderUpdateResponseAccessGoogleType = "centrify"
	AccessIdentityProviderUpdateResponseAccessGoogleTypeFacebook   AccessIdentityProviderUpdateResponseAccessGoogleType = "facebook"
	AccessIdentityProviderUpdateResponseAccessGoogleTypeGitHub     AccessIdentityProviderUpdateResponseAccessGoogleType = "github"
	AccessIdentityProviderUpdateResponseAccessGoogleTypeGoogleApps AccessIdentityProviderUpdateResponseAccessGoogleType = "google-apps"
	AccessIdentityProviderUpdateResponseAccessGoogleTypeGoogle     AccessIdentityProviderUpdateResponseAccessGoogleType = "google"
	AccessIdentityProviderUpdateResponseAccessGoogleTypeLinkedin   AccessIdentityProviderUpdateResponseAccessGoogleType = "linkedin"
	AccessIdentityProviderUpdateResponseAccessGoogleTypeOidc       AccessIdentityProviderUpdateResponseAccessGoogleType = "oidc"
	AccessIdentityProviderUpdateResponseAccessGoogleTypeOkta       AccessIdentityProviderUpdateResponseAccessGoogleType = "okta"
	AccessIdentityProviderUpdateResponseAccessGoogleTypeOnelogin   AccessIdentityProviderUpdateResponseAccessGoogleType = "onelogin"
	AccessIdentityProviderUpdateResponseAccessGoogleTypePingone    AccessIdentityProviderUpdateResponseAccessGoogleType = "pingone"
	AccessIdentityProviderUpdateResponseAccessGoogleTypeYandex     AccessIdentityProviderUpdateResponseAccessGoogleType = "yandex"
)

type AccessIdentityProviderUpdateResponseAccessGoogleApps struct {
	// UUID
	ID     string                                                     `json:"id"`
	Config AccessIdentityProviderUpdateResponseAccessGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessGoogleAppsType `json:"type"`
	JSON accessIdentityProviderUpdateResponseAccessGoogleAppsJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessGoogleAppsJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessGoogleApps]
type accessIdentityProviderUpdateResponseAccessGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessGoogleApps) implementsAccessIdentityProviderUpdateResponse() {
}

type AccessIdentityProviderUpdateResponseAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                         `json:"email_claim_name"`
	JSON           accessIdentityProviderUpdateResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessGoogleAppsConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessGoogleAppsConfig]
type accessIdentityProviderUpdateResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseAccessGoogleAppsScimConfig struct {
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
	JSON            accessIdentityProviderUpdateResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessGoogleAppsScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessGoogleAppsScimConfig]
type accessIdentityProviderUpdateResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseAccessGoogleAppsType string

const (
	AccessIdentityProviderUpdateResponseAccessGoogleAppsTypeOnetimepin AccessIdentityProviderUpdateResponseAccessGoogleAppsType = "onetimepin"
	AccessIdentityProviderUpdateResponseAccessGoogleAppsTypeAzureAd    AccessIdentityProviderUpdateResponseAccessGoogleAppsType = "azureAD"
	AccessIdentityProviderUpdateResponseAccessGoogleAppsTypeSaml       AccessIdentityProviderUpdateResponseAccessGoogleAppsType = "saml"
	AccessIdentityProviderUpdateResponseAccessGoogleAppsTypeCentrify   AccessIdentityProviderUpdateResponseAccessGoogleAppsType = "centrify"
	AccessIdentityProviderUpdateResponseAccessGoogleAppsTypeFacebook   AccessIdentityProviderUpdateResponseAccessGoogleAppsType = "facebook"
	AccessIdentityProviderUpdateResponseAccessGoogleAppsTypeGitHub     AccessIdentityProviderUpdateResponseAccessGoogleAppsType = "github"
	AccessIdentityProviderUpdateResponseAccessGoogleAppsTypeGoogleApps AccessIdentityProviderUpdateResponseAccessGoogleAppsType = "google-apps"
	AccessIdentityProviderUpdateResponseAccessGoogleAppsTypeGoogle     AccessIdentityProviderUpdateResponseAccessGoogleAppsType = "google"
	AccessIdentityProviderUpdateResponseAccessGoogleAppsTypeLinkedin   AccessIdentityProviderUpdateResponseAccessGoogleAppsType = "linkedin"
	AccessIdentityProviderUpdateResponseAccessGoogleAppsTypeOidc       AccessIdentityProviderUpdateResponseAccessGoogleAppsType = "oidc"
	AccessIdentityProviderUpdateResponseAccessGoogleAppsTypeOkta       AccessIdentityProviderUpdateResponseAccessGoogleAppsType = "okta"
	AccessIdentityProviderUpdateResponseAccessGoogleAppsTypeOnelogin   AccessIdentityProviderUpdateResponseAccessGoogleAppsType = "onelogin"
	AccessIdentityProviderUpdateResponseAccessGoogleAppsTypePingone    AccessIdentityProviderUpdateResponseAccessGoogleAppsType = "pingone"
	AccessIdentityProviderUpdateResponseAccessGoogleAppsTypeYandex     AccessIdentityProviderUpdateResponseAccessGoogleAppsType = "yandex"
)

type AccessIdentityProviderUpdateResponseAccessLinkedin struct {
	// UUID
	ID     string                                                   `json:"id"`
	Config AccessIdentityProviderUpdateResponseAccessLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessLinkedinType `json:"type"`
	JSON accessIdentityProviderUpdateResponseAccessLinkedinJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessLinkedinJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessLinkedin]
type accessIdentityProviderUpdateResponseAccessLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessLinkedin) implementsAccessIdentityProviderUpdateResponse() {
}

type AccessIdentityProviderUpdateResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                       `json:"client_secret"`
	JSON         accessIdentityProviderUpdateResponseAccessLinkedinConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessLinkedinConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessLinkedinConfig]
type accessIdentityProviderUpdateResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseAccessLinkedinScimConfig struct {
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
	UserDeprovision bool                                                             `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessLinkedinScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessLinkedinScimConfig]
type accessIdentityProviderUpdateResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseAccessLinkedinType string

const (
	AccessIdentityProviderUpdateResponseAccessLinkedinTypeOnetimepin AccessIdentityProviderUpdateResponseAccessLinkedinType = "onetimepin"
	AccessIdentityProviderUpdateResponseAccessLinkedinTypeAzureAd    AccessIdentityProviderUpdateResponseAccessLinkedinType = "azureAD"
	AccessIdentityProviderUpdateResponseAccessLinkedinTypeSaml       AccessIdentityProviderUpdateResponseAccessLinkedinType = "saml"
	AccessIdentityProviderUpdateResponseAccessLinkedinTypeCentrify   AccessIdentityProviderUpdateResponseAccessLinkedinType = "centrify"
	AccessIdentityProviderUpdateResponseAccessLinkedinTypeFacebook   AccessIdentityProviderUpdateResponseAccessLinkedinType = "facebook"
	AccessIdentityProviderUpdateResponseAccessLinkedinTypeGitHub     AccessIdentityProviderUpdateResponseAccessLinkedinType = "github"
	AccessIdentityProviderUpdateResponseAccessLinkedinTypeGoogleApps AccessIdentityProviderUpdateResponseAccessLinkedinType = "google-apps"
	AccessIdentityProviderUpdateResponseAccessLinkedinTypeGoogle     AccessIdentityProviderUpdateResponseAccessLinkedinType = "google"
	AccessIdentityProviderUpdateResponseAccessLinkedinTypeLinkedin   AccessIdentityProviderUpdateResponseAccessLinkedinType = "linkedin"
	AccessIdentityProviderUpdateResponseAccessLinkedinTypeOidc       AccessIdentityProviderUpdateResponseAccessLinkedinType = "oidc"
	AccessIdentityProviderUpdateResponseAccessLinkedinTypeOkta       AccessIdentityProviderUpdateResponseAccessLinkedinType = "okta"
	AccessIdentityProviderUpdateResponseAccessLinkedinTypeOnelogin   AccessIdentityProviderUpdateResponseAccessLinkedinType = "onelogin"
	AccessIdentityProviderUpdateResponseAccessLinkedinTypePingone    AccessIdentityProviderUpdateResponseAccessLinkedinType = "pingone"
	AccessIdentityProviderUpdateResponseAccessLinkedinTypeYandex     AccessIdentityProviderUpdateResponseAccessLinkedinType = "yandex"
)

type AccessIdentityProviderUpdateResponseAccessOidc struct {
	// UUID
	ID     string                                               `json:"id"`
	Config AccessIdentityProviderUpdateResponseAccessOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessOidcType `json:"type"`
	JSON accessIdentityProviderUpdateResponseAccessOidcJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessOidcJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessOidc]
type accessIdentityProviderUpdateResponseAccessOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessOidc) implementsAccessIdentityProviderUpdateResponse() {
}

type AccessIdentityProviderUpdateResponseAccessOidcConfig struct {
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
	TokenURL string                                                   `json:"token_url"`
	JSON     accessIdentityProviderUpdateResponseAccessOidcConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessOidcConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessOidcConfig]
type accessIdentityProviderUpdateResponseAccessOidcConfigJSON struct {
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

func (r *AccessIdentityProviderUpdateResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseAccessOidcScimConfig struct {
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
	UserDeprovision bool                                                         `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseAccessOidcScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessOidcScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessOidcScimConfig]
type accessIdentityProviderUpdateResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseAccessOidcType string

const (
	AccessIdentityProviderUpdateResponseAccessOidcTypeOnetimepin AccessIdentityProviderUpdateResponseAccessOidcType = "onetimepin"
	AccessIdentityProviderUpdateResponseAccessOidcTypeAzureAd    AccessIdentityProviderUpdateResponseAccessOidcType = "azureAD"
	AccessIdentityProviderUpdateResponseAccessOidcTypeSaml       AccessIdentityProviderUpdateResponseAccessOidcType = "saml"
	AccessIdentityProviderUpdateResponseAccessOidcTypeCentrify   AccessIdentityProviderUpdateResponseAccessOidcType = "centrify"
	AccessIdentityProviderUpdateResponseAccessOidcTypeFacebook   AccessIdentityProviderUpdateResponseAccessOidcType = "facebook"
	AccessIdentityProviderUpdateResponseAccessOidcTypeGitHub     AccessIdentityProviderUpdateResponseAccessOidcType = "github"
	AccessIdentityProviderUpdateResponseAccessOidcTypeGoogleApps AccessIdentityProviderUpdateResponseAccessOidcType = "google-apps"
	AccessIdentityProviderUpdateResponseAccessOidcTypeGoogle     AccessIdentityProviderUpdateResponseAccessOidcType = "google"
	AccessIdentityProviderUpdateResponseAccessOidcTypeLinkedin   AccessIdentityProviderUpdateResponseAccessOidcType = "linkedin"
	AccessIdentityProviderUpdateResponseAccessOidcTypeOidc       AccessIdentityProviderUpdateResponseAccessOidcType = "oidc"
	AccessIdentityProviderUpdateResponseAccessOidcTypeOkta       AccessIdentityProviderUpdateResponseAccessOidcType = "okta"
	AccessIdentityProviderUpdateResponseAccessOidcTypeOnelogin   AccessIdentityProviderUpdateResponseAccessOidcType = "onelogin"
	AccessIdentityProviderUpdateResponseAccessOidcTypePingone    AccessIdentityProviderUpdateResponseAccessOidcType = "pingone"
	AccessIdentityProviderUpdateResponseAccessOidcTypeYandex     AccessIdentityProviderUpdateResponseAccessOidcType = "yandex"
)

type AccessIdentityProviderUpdateResponseAccessOkta struct {
	// UUID
	ID     string                                               `json:"id"`
	Config AccessIdentityProviderUpdateResponseAccessOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessOktaType `json:"type"`
	JSON accessIdentityProviderUpdateResponseAccessOktaJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessOktaJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessOkta]
type accessIdentityProviderUpdateResponseAccessOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessOkta) implementsAccessIdentityProviderUpdateResponse() {
}

type AccessIdentityProviderUpdateResponseAccessOktaConfig struct {
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
	OktaAccount string                                                   `json:"okta_account"`
	JSON        accessIdentityProviderUpdateResponseAccessOktaConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessOktaConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessOktaConfig]
type accessIdentityProviderUpdateResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseAccessOktaScimConfig struct {
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
	UserDeprovision bool                                                         `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseAccessOktaScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessOktaScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessOktaScimConfig]
type accessIdentityProviderUpdateResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseAccessOktaType string

const (
	AccessIdentityProviderUpdateResponseAccessOktaTypeOnetimepin AccessIdentityProviderUpdateResponseAccessOktaType = "onetimepin"
	AccessIdentityProviderUpdateResponseAccessOktaTypeAzureAd    AccessIdentityProviderUpdateResponseAccessOktaType = "azureAD"
	AccessIdentityProviderUpdateResponseAccessOktaTypeSaml       AccessIdentityProviderUpdateResponseAccessOktaType = "saml"
	AccessIdentityProviderUpdateResponseAccessOktaTypeCentrify   AccessIdentityProviderUpdateResponseAccessOktaType = "centrify"
	AccessIdentityProviderUpdateResponseAccessOktaTypeFacebook   AccessIdentityProviderUpdateResponseAccessOktaType = "facebook"
	AccessIdentityProviderUpdateResponseAccessOktaTypeGitHub     AccessIdentityProviderUpdateResponseAccessOktaType = "github"
	AccessIdentityProviderUpdateResponseAccessOktaTypeGoogleApps AccessIdentityProviderUpdateResponseAccessOktaType = "google-apps"
	AccessIdentityProviderUpdateResponseAccessOktaTypeGoogle     AccessIdentityProviderUpdateResponseAccessOktaType = "google"
	AccessIdentityProviderUpdateResponseAccessOktaTypeLinkedin   AccessIdentityProviderUpdateResponseAccessOktaType = "linkedin"
	AccessIdentityProviderUpdateResponseAccessOktaTypeOidc       AccessIdentityProviderUpdateResponseAccessOktaType = "oidc"
	AccessIdentityProviderUpdateResponseAccessOktaTypeOkta       AccessIdentityProviderUpdateResponseAccessOktaType = "okta"
	AccessIdentityProviderUpdateResponseAccessOktaTypeOnelogin   AccessIdentityProviderUpdateResponseAccessOktaType = "onelogin"
	AccessIdentityProviderUpdateResponseAccessOktaTypePingone    AccessIdentityProviderUpdateResponseAccessOktaType = "pingone"
	AccessIdentityProviderUpdateResponseAccessOktaTypeYandex     AccessIdentityProviderUpdateResponseAccessOktaType = "yandex"
)

type AccessIdentityProviderUpdateResponseAccessOnelogin struct {
	// UUID
	ID     string                                                   `json:"id"`
	Config AccessIdentityProviderUpdateResponseAccessOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessOneloginType `json:"type"`
	JSON accessIdentityProviderUpdateResponseAccessOneloginJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessOneloginJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessOnelogin]
type accessIdentityProviderUpdateResponseAccessOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessOnelogin) implementsAccessIdentityProviderUpdateResponse() {
}

type AccessIdentityProviderUpdateResponseAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                       `json:"onelogin_account"`
	JSON            accessIdentityProviderUpdateResponseAccessOneloginConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessOneloginConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessOneloginConfig]
type accessIdentityProviderUpdateResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseAccessOneloginScimConfig struct {
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
	UserDeprovision bool                                                             `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseAccessOneloginScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessOneloginScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessOneloginScimConfig]
type accessIdentityProviderUpdateResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseAccessOneloginType string

const (
	AccessIdentityProviderUpdateResponseAccessOneloginTypeOnetimepin AccessIdentityProviderUpdateResponseAccessOneloginType = "onetimepin"
	AccessIdentityProviderUpdateResponseAccessOneloginTypeAzureAd    AccessIdentityProviderUpdateResponseAccessOneloginType = "azureAD"
	AccessIdentityProviderUpdateResponseAccessOneloginTypeSaml       AccessIdentityProviderUpdateResponseAccessOneloginType = "saml"
	AccessIdentityProviderUpdateResponseAccessOneloginTypeCentrify   AccessIdentityProviderUpdateResponseAccessOneloginType = "centrify"
	AccessIdentityProviderUpdateResponseAccessOneloginTypeFacebook   AccessIdentityProviderUpdateResponseAccessOneloginType = "facebook"
	AccessIdentityProviderUpdateResponseAccessOneloginTypeGitHub     AccessIdentityProviderUpdateResponseAccessOneloginType = "github"
	AccessIdentityProviderUpdateResponseAccessOneloginTypeGoogleApps AccessIdentityProviderUpdateResponseAccessOneloginType = "google-apps"
	AccessIdentityProviderUpdateResponseAccessOneloginTypeGoogle     AccessIdentityProviderUpdateResponseAccessOneloginType = "google"
	AccessIdentityProviderUpdateResponseAccessOneloginTypeLinkedin   AccessIdentityProviderUpdateResponseAccessOneloginType = "linkedin"
	AccessIdentityProviderUpdateResponseAccessOneloginTypeOidc       AccessIdentityProviderUpdateResponseAccessOneloginType = "oidc"
	AccessIdentityProviderUpdateResponseAccessOneloginTypeOkta       AccessIdentityProviderUpdateResponseAccessOneloginType = "okta"
	AccessIdentityProviderUpdateResponseAccessOneloginTypeOnelogin   AccessIdentityProviderUpdateResponseAccessOneloginType = "onelogin"
	AccessIdentityProviderUpdateResponseAccessOneloginTypePingone    AccessIdentityProviderUpdateResponseAccessOneloginType = "pingone"
	AccessIdentityProviderUpdateResponseAccessOneloginTypeYandex     AccessIdentityProviderUpdateResponseAccessOneloginType = "yandex"
)

type AccessIdentityProviderUpdateResponseAccessPingone struct {
	// UUID
	ID     string                                                  `json:"id"`
	Config AccessIdentityProviderUpdateResponseAccessPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessPingoneType `json:"type"`
	JSON accessIdentityProviderUpdateResponseAccessPingoneJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessPingoneJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessPingone]
type accessIdentityProviderUpdateResponseAccessPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessPingone) implementsAccessIdentityProviderUpdateResponse() {
}

type AccessIdentityProviderUpdateResponseAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                      `json:"ping_env_id"`
	JSON      accessIdentityProviderUpdateResponseAccessPingoneConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessPingoneConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessPingoneConfig]
type accessIdentityProviderUpdateResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseAccessPingoneScimConfig struct {
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
	JSON            accessIdentityProviderUpdateResponseAccessPingoneScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessPingoneScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessPingoneScimConfig]
type accessIdentityProviderUpdateResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseAccessPingoneType string

const (
	AccessIdentityProviderUpdateResponseAccessPingoneTypeOnetimepin AccessIdentityProviderUpdateResponseAccessPingoneType = "onetimepin"
	AccessIdentityProviderUpdateResponseAccessPingoneTypeAzureAd    AccessIdentityProviderUpdateResponseAccessPingoneType = "azureAD"
	AccessIdentityProviderUpdateResponseAccessPingoneTypeSaml       AccessIdentityProviderUpdateResponseAccessPingoneType = "saml"
	AccessIdentityProviderUpdateResponseAccessPingoneTypeCentrify   AccessIdentityProviderUpdateResponseAccessPingoneType = "centrify"
	AccessIdentityProviderUpdateResponseAccessPingoneTypeFacebook   AccessIdentityProviderUpdateResponseAccessPingoneType = "facebook"
	AccessIdentityProviderUpdateResponseAccessPingoneTypeGitHub     AccessIdentityProviderUpdateResponseAccessPingoneType = "github"
	AccessIdentityProviderUpdateResponseAccessPingoneTypeGoogleApps AccessIdentityProviderUpdateResponseAccessPingoneType = "google-apps"
	AccessIdentityProviderUpdateResponseAccessPingoneTypeGoogle     AccessIdentityProviderUpdateResponseAccessPingoneType = "google"
	AccessIdentityProviderUpdateResponseAccessPingoneTypeLinkedin   AccessIdentityProviderUpdateResponseAccessPingoneType = "linkedin"
	AccessIdentityProviderUpdateResponseAccessPingoneTypeOidc       AccessIdentityProviderUpdateResponseAccessPingoneType = "oidc"
	AccessIdentityProviderUpdateResponseAccessPingoneTypeOkta       AccessIdentityProviderUpdateResponseAccessPingoneType = "okta"
	AccessIdentityProviderUpdateResponseAccessPingoneTypeOnelogin   AccessIdentityProviderUpdateResponseAccessPingoneType = "onelogin"
	AccessIdentityProviderUpdateResponseAccessPingoneTypePingone    AccessIdentityProviderUpdateResponseAccessPingoneType = "pingone"
	AccessIdentityProviderUpdateResponseAccessPingoneTypeYandex     AccessIdentityProviderUpdateResponseAccessPingoneType = "yandex"
)

type AccessIdentityProviderUpdateResponseAccessSaml struct {
	// UUID
	ID     string                                               `json:"id"`
	Config AccessIdentityProviderUpdateResponseAccessSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessSamlType `json:"type"`
	JSON accessIdentityProviderUpdateResponseAccessSamlJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessSamlJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessSaml]
type accessIdentityProviderUpdateResponseAccessSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessSaml) implementsAccessIdentityProviderUpdateResponse() {
}

type AccessIdentityProviderUpdateResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccessIdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                   `json:"sso_target_url"`
	JSON         accessIdentityProviderUpdateResponseAccessSamlConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessSamlConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessSamlConfig]
type accessIdentityProviderUpdateResponseAccessSamlConfigJSON struct {
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

func (r *AccessIdentityProviderUpdateResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                  `json:"header_name"`
	JSON       accessIdentityProviderUpdateResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessSamlConfigHeaderAttributeJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute]
type accessIdentityProviderUpdateResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseAccessSamlScimConfig struct {
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
	UserDeprovision bool                                                         `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseAccessSamlScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessSamlScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessSamlScimConfig]
type accessIdentityProviderUpdateResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseAccessSamlType string

const (
	AccessIdentityProviderUpdateResponseAccessSamlTypeOnetimepin AccessIdentityProviderUpdateResponseAccessSamlType = "onetimepin"
	AccessIdentityProviderUpdateResponseAccessSamlTypeAzureAd    AccessIdentityProviderUpdateResponseAccessSamlType = "azureAD"
	AccessIdentityProviderUpdateResponseAccessSamlTypeSaml       AccessIdentityProviderUpdateResponseAccessSamlType = "saml"
	AccessIdentityProviderUpdateResponseAccessSamlTypeCentrify   AccessIdentityProviderUpdateResponseAccessSamlType = "centrify"
	AccessIdentityProviderUpdateResponseAccessSamlTypeFacebook   AccessIdentityProviderUpdateResponseAccessSamlType = "facebook"
	AccessIdentityProviderUpdateResponseAccessSamlTypeGitHub     AccessIdentityProviderUpdateResponseAccessSamlType = "github"
	AccessIdentityProviderUpdateResponseAccessSamlTypeGoogleApps AccessIdentityProviderUpdateResponseAccessSamlType = "google-apps"
	AccessIdentityProviderUpdateResponseAccessSamlTypeGoogle     AccessIdentityProviderUpdateResponseAccessSamlType = "google"
	AccessIdentityProviderUpdateResponseAccessSamlTypeLinkedin   AccessIdentityProviderUpdateResponseAccessSamlType = "linkedin"
	AccessIdentityProviderUpdateResponseAccessSamlTypeOidc       AccessIdentityProviderUpdateResponseAccessSamlType = "oidc"
	AccessIdentityProviderUpdateResponseAccessSamlTypeOkta       AccessIdentityProviderUpdateResponseAccessSamlType = "okta"
	AccessIdentityProviderUpdateResponseAccessSamlTypeOnelogin   AccessIdentityProviderUpdateResponseAccessSamlType = "onelogin"
	AccessIdentityProviderUpdateResponseAccessSamlTypePingone    AccessIdentityProviderUpdateResponseAccessSamlType = "pingone"
	AccessIdentityProviderUpdateResponseAccessSamlTypeYandex     AccessIdentityProviderUpdateResponseAccessSamlType = "yandex"
)

type AccessIdentityProviderUpdateResponseAccessYandex struct {
	// UUID
	ID     string                                                 `json:"id"`
	Config AccessIdentityProviderUpdateResponseAccessYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessYandexType `json:"type"`
	JSON accessIdentityProviderUpdateResponseAccessYandexJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessYandexJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessYandex]
type accessIdentityProviderUpdateResponseAccessYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessYandex) implementsAccessIdentityProviderUpdateResponse() {
}

type AccessIdentityProviderUpdateResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                     `json:"client_secret"`
	JSON         accessIdentityProviderUpdateResponseAccessYandexConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessYandexConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessYandexConfig]
type accessIdentityProviderUpdateResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseAccessYandexScimConfig struct {
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
	UserDeprovision bool                                                           `json:"user_deprovision"`
	JSON            accessIdentityProviderUpdateResponseAccessYandexScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessYandexScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessYandexScimConfig]
type accessIdentityProviderUpdateResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseAccessYandexType string

const (
	AccessIdentityProviderUpdateResponseAccessYandexTypeOnetimepin AccessIdentityProviderUpdateResponseAccessYandexType = "onetimepin"
	AccessIdentityProviderUpdateResponseAccessYandexTypeAzureAd    AccessIdentityProviderUpdateResponseAccessYandexType = "azureAD"
	AccessIdentityProviderUpdateResponseAccessYandexTypeSaml       AccessIdentityProviderUpdateResponseAccessYandexType = "saml"
	AccessIdentityProviderUpdateResponseAccessYandexTypeCentrify   AccessIdentityProviderUpdateResponseAccessYandexType = "centrify"
	AccessIdentityProviderUpdateResponseAccessYandexTypeFacebook   AccessIdentityProviderUpdateResponseAccessYandexType = "facebook"
	AccessIdentityProviderUpdateResponseAccessYandexTypeGitHub     AccessIdentityProviderUpdateResponseAccessYandexType = "github"
	AccessIdentityProviderUpdateResponseAccessYandexTypeGoogleApps AccessIdentityProviderUpdateResponseAccessYandexType = "google-apps"
	AccessIdentityProviderUpdateResponseAccessYandexTypeGoogle     AccessIdentityProviderUpdateResponseAccessYandexType = "google"
	AccessIdentityProviderUpdateResponseAccessYandexTypeLinkedin   AccessIdentityProviderUpdateResponseAccessYandexType = "linkedin"
	AccessIdentityProviderUpdateResponseAccessYandexTypeOidc       AccessIdentityProviderUpdateResponseAccessYandexType = "oidc"
	AccessIdentityProviderUpdateResponseAccessYandexTypeOkta       AccessIdentityProviderUpdateResponseAccessYandexType = "okta"
	AccessIdentityProviderUpdateResponseAccessYandexTypeOnelogin   AccessIdentityProviderUpdateResponseAccessYandexType = "onelogin"
	AccessIdentityProviderUpdateResponseAccessYandexTypePingone    AccessIdentityProviderUpdateResponseAccessYandexType = "pingone"
	AccessIdentityProviderUpdateResponseAccessYandexTypeYandex     AccessIdentityProviderUpdateResponseAccessYandexType = "yandex"
)

type AccessIdentityProviderUpdateResponseAccessOnetimepin struct {
	// UUID
	ID     string      `json:"id"`
	Config interface{} `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessOnetimepinScimConfig `json:"scim_config"`
	Type       AccessIdentityProviderUpdateResponseAccessOnetimepinType       `json:"type"`
	JSON       accessIdentityProviderUpdateResponseAccessOnetimepinJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessOnetimepinJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessOnetimepin]
type accessIdentityProviderUpdateResponseAccessOnetimepinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessOnetimepin) implementsAccessIdentityProviderUpdateResponse() {
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderUpdateResponseAccessOnetimepinScimConfig struct {
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
	JSON            accessIdentityProviderUpdateResponseAccessOnetimepinScimConfigJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessOnetimepinScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderUpdateResponseAccessOnetimepinScimConfig]
type accessIdentityProviderUpdateResponseAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderUpdateResponseAccessOnetimepinType string

const (
	AccessIdentityProviderUpdateResponseAccessOnetimepinTypeOnetimepin AccessIdentityProviderUpdateResponseAccessOnetimepinType = "onetimepin"
)

type AccessIdentityProviderDeleteResponse struct {
	// UUID
	ID   string                                   `json:"id"`
	JSON accessIdentityProviderDeleteResponseJSON `json:"-"`
}

// accessIdentityProviderDeleteResponseJSON contains the JSON metadata for the
// struct [AccessIdentityProviderDeleteResponse]
type accessIdentityProviderDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAd],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrify],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebook],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHub],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogle],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleApps],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedin],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidc],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOkta],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnelogin],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingone],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSaml],
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandex]
// or
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepin].
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse interface {
	implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse)(nil)).Elem(), "")
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAd struct {
	// UUID
	ID     string                                                                                              `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAd]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAd) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdConfig struct {
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
	SupportGroups bool                                                                                                    `json:"support_groups"`
	JSON          accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdConfigJSON struct {
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

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdScimConfig struct {
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
	UserDeprovision bool                                                                                                        `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessAzureAdType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrify struct {
	// UUID
	ID     string                                                                                               `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrify]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrify) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyConfig struct {
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
	EmailClaimName string                                                                                                   `json:"email_claim_name"`
	JSON           accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyScimConfig struct {
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
	UserDeprovision bool                                                                                                         `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessCentrifyType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebook struct {
	// UUID
	ID     string                                                                                               `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebook]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebook) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                   `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookScimConfig struct {
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
	UserDeprovision bool                                                                                                         `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessFacebookType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHub struct {
	// UUID
	ID     string                                                                                             `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHub]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHub) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                 `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubScimConfig struct {
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
	UserDeprovision bool                                                                                                       `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGitHubType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogle struct {
	// UUID
	ID     string                                                                                             `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogle]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogle) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                                                 `json:"email_claim_name"`
	JSON           accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleScimConfig struct {
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
	UserDeprovision bool                                                                                                       `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleApps struct {
	// UUID
	ID     string                                                                                                 `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleApps]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleApps) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                                                     `json:"email_claim_name"`
	JSON           accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsScimConfig struct {
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
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessGoogleAppsType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedin struct {
	// UUID
	ID     string                                                                                               `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedin]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedin) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                   `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinScimConfig struct {
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
	UserDeprovision bool                                                                                                         `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessLinkedinType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidc struct {
	// UUID
	ID     string                                                                                           `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidc]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidc) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcConfig struct {
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
	TokenURL string                                                                                               `json:"token_url"`
	JSON     accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcConfigJSON struct {
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

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcScimConfig struct {
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
	UserDeprovision bool                                                                                                     `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOidcType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOkta struct {
	// UUID
	ID     string                                                                                           `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOkta]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOkta) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaConfig struct {
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
	OktaAccount string                                                                                               `json:"okta_account"`
	JSON        accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaScimConfig struct {
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
	UserDeprovision bool                                                                                                     `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOktaType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnelogin struct {
	// UUID
	ID     string                                                                                               `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnelogin]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnelogin) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                                                                   `json:"onelogin_account"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginScimConfig struct {
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
	UserDeprovision bool                                                                                                         `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOneloginType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingone struct {
	// UUID
	ID     string                                                                                              `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingone]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingone) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                                                                  `json:"ping_env_id"`
	JSON      accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneScimConfig struct {
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
	UserDeprovision bool                                                                                                        `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessPingoneType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSaml struct {
	// UUID
	ID     string                                                                                           `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSaml]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSaml) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                                                               `json:"sso_target_url"`
	JSON         accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlConfigJSON struct {
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

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                                                              `json:"header_name"`
	JSON       accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlConfigHeaderAttribute]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlScimConfig struct {
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
	UserDeprovision bool                                                                                                     `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessSamlType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandex struct {
	// UUID
	ID     string                                                                                             `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandex]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandex) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                 `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexScimConfig struct {
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
	UserDeprovision bool                                                                                                       `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexTypeSaml       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType = "saml"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexTypeCentrify   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexTypeFacebook   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexTypeGitHub     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType = "github"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexTypeGoogle     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType = "google"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexTypeOidc       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexTypeOkta       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType = "okta"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexTypePingone    AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexTypeYandex     AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessYandexType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepin struct {
	// UUID
	ID     string      `json:"id"`
	Config interface{} `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepinScimConfig `json:"scim_config"`
	Type       AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepinType       `json:"type"`
	JSON       accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepinJSON       `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepinJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepin]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepin) implementsAccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse() {
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepinScimConfig struct {
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
	JSON            accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepinScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepinScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepinScimConfig]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepinType string

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepinTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseAccessOnetimepinType = "onetimepin"
)

// Union satisfied by
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAd],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrify],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebook],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHub],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogle],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleApps],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedin],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidc],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOkta],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOnelogin],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingone],
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSaml]
// or
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandex].
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse interface {
	implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse)(nil)).Elem(), "")
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAd struct {
	// UUID
	ID     string                                                                                              `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAd]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAd) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdConfig struct {
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
	SupportGroups bool                                                                                                    `json:"support_groups"`
	JSON          accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdConfigJSON struct {
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

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdScimConfig struct {
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
	UserDeprovision bool                                                                                                        `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessAzureAdType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrify struct {
	// UUID
	ID     string                                                                                               `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrify]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrify) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyConfig struct {
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
	EmailClaimName string                                                                                                   `json:"email_claim_name"`
	JSON           accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyScimConfig struct {
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
	UserDeprovision bool                                                                                                         `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessCentrifyType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebook struct {
	// UUID
	ID     string                                                                                               `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebook]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebook) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                   `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookScimConfig struct {
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
	UserDeprovision bool                                                                                                         `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessFacebookType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHub struct {
	// UUID
	ID     string                                                                                             `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHub]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHub) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                 `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubScimConfig struct {
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
	UserDeprovision bool                                                                                                       `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGitHubType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogle struct {
	// UUID
	ID     string                                                                                             `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogle]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogle) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                                                 `json:"email_claim_name"`
	JSON           accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleScimConfig struct {
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
	UserDeprovision bool                                                                                                       `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleApps struct {
	// UUID
	ID     string                                                                                                 `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleApps]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleApps) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                                                     `json:"email_claim_name"`
	JSON           accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsScimConfig struct {
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
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessGoogleAppsType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedin struct {
	// UUID
	ID     string                                                                                               `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedin]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedin) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                   `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinScimConfig struct {
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
	UserDeprovision bool                                                                                                         `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessLinkedinType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidc struct {
	// UUID
	ID     string                                                                                           `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidc]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidc) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcConfig struct {
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
	TokenURL string                                                                                               `json:"token_url"`
	JSON     accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcConfigJSON struct {
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

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcScimConfig struct {
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
	UserDeprovision bool                                                                                                     `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOidcType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOkta struct {
	// UUID
	ID     string                                                                                           `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOkta]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOkta) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaConfig struct {
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
	OktaAccount string                                                                                               `json:"okta_account"`
	JSON        accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaScimConfig struct {
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
	UserDeprovision bool                                                                                                     `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOktaType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOnelogin struct {
	// UUID
	ID     string                                                                                               `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOnelogin]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOnelogin) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                                                                   `json:"onelogin_account"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginScimConfig struct {
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
	UserDeprovision bool                                                                                                         `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessOneloginType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingone struct {
	// UUID
	ID     string                                                                                              `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingone]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingone) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                                                                  `json:"ping_env_id"`
	JSON      accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneScimConfig struct {
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
	UserDeprovision bool                                                                                                        `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessPingoneType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSaml struct {
	// UUID
	ID     string                                                                                           `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSaml]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSaml) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                                                               `json:"sso_target_url"`
	JSON         accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlConfigJSON struct {
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

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                                                              `json:"header_name"`
	JSON       accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlConfigHeaderAttribute]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlScimConfig struct {
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
	UserDeprovision bool                                                                                                     `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessSamlType = "yandex"
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandex struct {
	// UUID
	ID     string                                                                                             `json:"id"`
	Config AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType `json:"type"`
	JSON accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandex]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandex) implementsAccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse() {
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                                                 `json:"client_secret"`
	JSON         accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexScimConfig struct {
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
	UserDeprovision bool                                                                                                       `json:"user_deprovision"`
	JSON            accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexScimConfigJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexScimConfigJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexScimConfig]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType string

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexTypeOnetimepin AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType = "onetimepin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexTypeAzureAd    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType = "azureAD"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexTypeSaml       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType = "saml"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexTypeCentrify   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType = "centrify"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexTypeFacebook   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType = "facebook"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexTypeGitHub     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType = "github"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexTypeGoogleApps AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType = "google-apps"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexTypeGoogle     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType = "google"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexTypeLinkedin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType = "linkedin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexTypeOidc       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType = "oidc"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexTypeOkta       AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType = "okta"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexTypeOnelogin   AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType = "onelogin"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexTypePingone    AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType = "pingone"
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexTypeYandex     AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseAccessYandexType = "yandex"
)

type AccessIdentityProviderGetResponseEnvelope struct {
	Errors   []AccessIdentityProviderGetResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessIdentityProviderGetResponseEnvelopeMessages `json:"messages"`
	Result   AccessIdentityProviderGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessIdentityProviderGetResponseEnvelopeSuccess `json:"success"`
	JSON    accessIdentityProviderGetResponseEnvelopeJSON    `json:"-"`
}

// accessIdentityProviderGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessIdentityProviderGetResponseEnvelope]
type accessIdentityProviderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accessIdentityProviderGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessIdentityProviderGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseEnvelopeErrors]
type accessIdentityProviderGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accessIdentityProviderGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessIdentityProviderGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseEnvelopeMessages]
type accessIdentityProviderGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessIdentityProviderGetResponseEnvelopeSuccess bool

const (
	AccessIdentityProviderGetResponseEnvelopeSuccessTrue AccessIdentityProviderGetResponseEnvelopeSuccess = true
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

type AccessIdentityProviderUpdateResponseEnvelope struct {
	Errors   []AccessIdentityProviderUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessIdentityProviderUpdateResponseEnvelopeMessages `json:"messages"`
	Result   AccessIdentityProviderUpdateResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessIdentityProviderUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    accessIdentityProviderUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessIdentityProviderUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessIdentityProviderUpdateResponseEnvelope]
type accessIdentityProviderUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderUpdateResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accessIdentityProviderUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseEnvelopeErrors]
type accessIdentityProviderUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderUpdateResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accessIdentityProviderUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessIdentityProviderUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseEnvelopeMessages]
type accessIdentityProviderUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessIdentityProviderUpdateResponseEnvelopeSuccess bool

const (
	AccessIdentityProviderUpdateResponseEnvelopeSuccessTrue AccessIdentityProviderUpdateResponseEnvelopeSuccess = true
)

type AccessIdentityProviderDeleteResponseEnvelope struct {
	Errors   []AccessIdentityProviderDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessIdentityProviderDeleteResponseEnvelopeMessages `json:"messages"`
	Result   AccessIdentityProviderDeleteResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessIdentityProviderDeleteResponseEnvelopeSuccess `json:"success"`
	JSON    accessIdentityProviderDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessIdentityProviderDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessIdentityProviderDeleteResponseEnvelope]
type accessIdentityProviderDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderDeleteResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accessIdentityProviderDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessIdentityProviderDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessIdentityProviderDeleteResponseEnvelopeErrors]
type accessIdentityProviderDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderDeleteResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accessIdentityProviderDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessIdentityProviderDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessIdentityProviderDeleteResponseEnvelopeMessages]
type accessIdentityProviderDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessIdentityProviderDeleteResponseEnvelopeSuccess bool

const (
	AccessIdentityProviderDeleteResponseEnvelopeSuccessTrue AccessIdentityProviderDeleteResponseEnvelopeSuccess = true
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

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelope struct {
	Errors   []AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeMessages `json:"messages"`
	Result   AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeSuccess `json:"success"`
	JSON    accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeJSON    `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelope]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeErrors struct {
	Code    int64                                                                                              `json:"code,required"`
	Message string                                                                                             `json:"message,required"`
	JSON    accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeErrorsJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeErrors]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeMessages struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeMessagesJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeMessages]
type accessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeSuccess bool

const (
	AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeSuccessTrue AccessIdentityProviderAccessIdentityProvidersAddAnAccessIdentityProviderResponseEnvelopeSuccess = true
)

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelope struct {
	Errors     []AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeErrors   `json:"errors"`
	Messages   []AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeMessages `json:"messages"`
	Result     []AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponse                 `json:"result"`
	ResultInfo AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeSuccess `json:"success"`
	JSON    accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeJSON    `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelope]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeErrors struct {
	Code    int64                                                                                              `json:"code,required"`
	Message string                                                                                             `json:"message,required"`
	JSON    accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeErrorsJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeErrors]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeMessages struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeMessagesJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeMessages]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                `json:"total_count"`
	JSON       accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeResultInfo]
type accessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeSuccess bool

const (
	AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeSuccessTrue AccessIdentityProviderAccessIdentityProvidersListAccessIdentityProvidersResponseEnvelopeSuccess = true
)

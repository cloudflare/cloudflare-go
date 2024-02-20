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

// Adds a new identity provider to Access.
func (r *AccessIdentityProviderService) New(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessIdentityProviderNewParams, opts ...option.RequestOption) (res *AccessIdentityProviderNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessIdentityProviderNewResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/identity_providers", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all configured identity providers.
func (r *AccessIdentityProviderService) List(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *[]AccessIdentityProviderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessIdentityProviderListResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/identity_providers", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
func (r *AccessIdentityProviderService) Replace(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, body AccessIdentityProviderReplaceParams, opts ...option.RequestOption) (res *AccessIdentityProviderReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessIdentityProviderReplaceResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/identity_providers/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AccessIdentityProviderNewResponseAccessAzureAd],
// [AccessIdentityProviderNewResponseAccessCentrify],
// [AccessIdentityProviderNewResponseAccessFacebook],
// [AccessIdentityProviderNewResponseAccessGitHub],
// [AccessIdentityProviderNewResponseAccessGoogle],
// [AccessIdentityProviderNewResponseAccessGoogleApps],
// [AccessIdentityProviderNewResponseAccessLinkedin],
// [AccessIdentityProviderNewResponseAccessOidc],
// [AccessIdentityProviderNewResponseAccessOkta],
// [AccessIdentityProviderNewResponseAccessOnelogin],
// [AccessIdentityProviderNewResponseAccessPingone],
// [AccessIdentityProviderNewResponseAccessSaml],
// [AccessIdentityProviderNewResponseAccessYandex] or
// [AccessIdentityProviderNewResponseAccessOnetimepin].
type AccessIdentityProviderNewResponse interface {
	implementsAccessIdentityProviderNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessIdentityProviderNewResponse)(nil)).Elem(), "")
}

type AccessIdentityProviderNewResponseAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderNewResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderNewResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderNewResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderNewResponseAccessAzureAdJSON       `json:"-"`
}

// accessIdentityProviderNewResponseAccessAzureAdJSON contains the JSON metadata
// for the struct [AccessIdentityProviderNewResponseAccessAzureAd]
type accessIdentityProviderNewResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderNewResponseAccessAzureAd) implementsAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessAzureAdConfig struct {
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
	JSON          accessIdentityProviderNewResponseAccessAzureAdConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessAzureAdConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderNewResponseAccessAzureAdConfig]
type accessIdentityProviderNewResponseAccessAzureAdConfigJSON struct {
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

func (r *AccessIdentityProviderNewResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessAzureAdType string

const (
	AccessIdentityProviderNewResponseAccessAzureAdTypeOnetimepin AccessIdentityProviderNewResponseAccessAzureAdType = "onetimepin"
	AccessIdentityProviderNewResponseAccessAzureAdTypeAzureAd    AccessIdentityProviderNewResponseAccessAzureAdType = "azureAD"
	AccessIdentityProviderNewResponseAccessAzureAdTypeSaml       AccessIdentityProviderNewResponseAccessAzureAdType = "saml"
	AccessIdentityProviderNewResponseAccessAzureAdTypeCentrify   AccessIdentityProviderNewResponseAccessAzureAdType = "centrify"
	AccessIdentityProviderNewResponseAccessAzureAdTypeFacebook   AccessIdentityProviderNewResponseAccessAzureAdType = "facebook"
	AccessIdentityProviderNewResponseAccessAzureAdTypeGitHub     AccessIdentityProviderNewResponseAccessAzureAdType = "github"
	AccessIdentityProviderNewResponseAccessAzureAdTypeGoogleApps AccessIdentityProviderNewResponseAccessAzureAdType = "google-apps"
	AccessIdentityProviderNewResponseAccessAzureAdTypeGoogle     AccessIdentityProviderNewResponseAccessAzureAdType = "google"
	AccessIdentityProviderNewResponseAccessAzureAdTypeLinkedin   AccessIdentityProviderNewResponseAccessAzureAdType = "linkedin"
	AccessIdentityProviderNewResponseAccessAzureAdTypeOidc       AccessIdentityProviderNewResponseAccessAzureAdType = "oidc"
	AccessIdentityProviderNewResponseAccessAzureAdTypeOkta       AccessIdentityProviderNewResponseAccessAzureAdType = "okta"
	AccessIdentityProviderNewResponseAccessAzureAdTypeOnelogin   AccessIdentityProviderNewResponseAccessAzureAdType = "onelogin"
	AccessIdentityProviderNewResponseAccessAzureAdTypePingone    AccessIdentityProviderNewResponseAccessAzureAdType = "pingone"
	AccessIdentityProviderNewResponseAccessAzureAdTypeYandex     AccessIdentityProviderNewResponseAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewResponseAccessAzureAdScimConfig struct {
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
	JSON            accessIdentityProviderNewResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessAzureAdScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderNewResponseAccessAzureAdScimConfig]
type accessIdentityProviderNewResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderNewResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderNewResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderNewResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderNewResponseAccessCentrifyJSON       `json:"-"`
}

// accessIdentityProviderNewResponseAccessCentrifyJSON contains the JSON metadata
// for the struct [AccessIdentityProviderNewResponseAccessCentrify]
type accessIdentityProviderNewResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderNewResponseAccessCentrify) implementsAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessCentrifyConfig struct {
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
	JSON           accessIdentityProviderNewResponseAccessCentrifyConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessCentrifyConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderNewResponseAccessCentrifyConfig]
type accessIdentityProviderNewResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessCentrifyType string

const (
	AccessIdentityProviderNewResponseAccessCentrifyTypeOnetimepin AccessIdentityProviderNewResponseAccessCentrifyType = "onetimepin"
	AccessIdentityProviderNewResponseAccessCentrifyTypeAzureAd    AccessIdentityProviderNewResponseAccessCentrifyType = "azureAD"
	AccessIdentityProviderNewResponseAccessCentrifyTypeSaml       AccessIdentityProviderNewResponseAccessCentrifyType = "saml"
	AccessIdentityProviderNewResponseAccessCentrifyTypeCentrify   AccessIdentityProviderNewResponseAccessCentrifyType = "centrify"
	AccessIdentityProviderNewResponseAccessCentrifyTypeFacebook   AccessIdentityProviderNewResponseAccessCentrifyType = "facebook"
	AccessIdentityProviderNewResponseAccessCentrifyTypeGitHub     AccessIdentityProviderNewResponseAccessCentrifyType = "github"
	AccessIdentityProviderNewResponseAccessCentrifyTypeGoogleApps AccessIdentityProviderNewResponseAccessCentrifyType = "google-apps"
	AccessIdentityProviderNewResponseAccessCentrifyTypeGoogle     AccessIdentityProviderNewResponseAccessCentrifyType = "google"
	AccessIdentityProviderNewResponseAccessCentrifyTypeLinkedin   AccessIdentityProviderNewResponseAccessCentrifyType = "linkedin"
	AccessIdentityProviderNewResponseAccessCentrifyTypeOidc       AccessIdentityProviderNewResponseAccessCentrifyType = "oidc"
	AccessIdentityProviderNewResponseAccessCentrifyTypeOkta       AccessIdentityProviderNewResponseAccessCentrifyType = "okta"
	AccessIdentityProviderNewResponseAccessCentrifyTypeOnelogin   AccessIdentityProviderNewResponseAccessCentrifyType = "onelogin"
	AccessIdentityProviderNewResponseAccessCentrifyTypePingone    AccessIdentityProviderNewResponseAccessCentrifyType = "pingone"
	AccessIdentityProviderNewResponseAccessCentrifyTypeYandex     AccessIdentityProviderNewResponseAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewResponseAccessCentrifyScimConfig struct {
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
	JSON            accessIdentityProviderNewResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessCentrifyScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderNewResponseAccessCentrifyScimConfig]
type accessIdentityProviderNewResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderNewResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderNewResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderNewResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderNewResponseAccessFacebookJSON       `json:"-"`
}

// accessIdentityProviderNewResponseAccessFacebookJSON contains the JSON metadata
// for the struct [AccessIdentityProviderNewResponseAccessFacebook]
type accessIdentityProviderNewResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderNewResponseAccessFacebook) implementsAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                    `json:"client_secret"`
	JSON         accessIdentityProviderNewResponseAccessFacebookConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessFacebookConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderNewResponseAccessFacebookConfig]
type accessIdentityProviderNewResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessFacebookType string

const (
	AccessIdentityProviderNewResponseAccessFacebookTypeOnetimepin AccessIdentityProviderNewResponseAccessFacebookType = "onetimepin"
	AccessIdentityProviderNewResponseAccessFacebookTypeAzureAd    AccessIdentityProviderNewResponseAccessFacebookType = "azureAD"
	AccessIdentityProviderNewResponseAccessFacebookTypeSaml       AccessIdentityProviderNewResponseAccessFacebookType = "saml"
	AccessIdentityProviderNewResponseAccessFacebookTypeCentrify   AccessIdentityProviderNewResponseAccessFacebookType = "centrify"
	AccessIdentityProviderNewResponseAccessFacebookTypeFacebook   AccessIdentityProviderNewResponseAccessFacebookType = "facebook"
	AccessIdentityProviderNewResponseAccessFacebookTypeGitHub     AccessIdentityProviderNewResponseAccessFacebookType = "github"
	AccessIdentityProviderNewResponseAccessFacebookTypeGoogleApps AccessIdentityProviderNewResponseAccessFacebookType = "google-apps"
	AccessIdentityProviderNewResponseAccessFacebookTypeGoogle     AccessIdentityProviderNewResponseAccessFacebookType = "google"
	AccessIdentityProviderNewResponseAccessFacebookTypeLinkedin   AccessIdentityProviderNewResponseAccessFacebookType = "linkedin"
	AccessIdentityProviderNewResponseAccessFacebookTypeOidc       AccessIdentityProviderNewResponseAccessFacebookType = "oidc"
	AccessIdentityProviderNewResponseAccessFacebookTypeOkta       AccessIdentityProviderNewResponseAccessFacebookType = "okta"
	AccessIdentityProviderNewResponseAccessFacebookTypeOnelogin   AccessIdentityProviderNewResponseAccessFacebookType = "onelogin"
	AccessIdentityProviderNewResponseAccessFacebookTypePingone    AccessIdentityProviderNewResponseAccessFacebookType = "pingone"
	AccessIdentityProviderNewResponseAccessFacebookTypeYandex     AccessIdentityProviderNewResponseAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewResponseAccessFacebookScimConfig struct {
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
	JSON            accessIdentityProviderNewResponseAccessFacebookScimConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessFacebookScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderNewResponseAccessFacebookScimConfig]
type accessIdentityProviderNewResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderNewResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderNewResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderNewResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderNewResponseAccessGitHubJSON       `json:"-"`
}

// accessIdentityProviderNewResponseAccessGitHubJSON contains the JSON metadata for
// the struct [AccessIdentityProviderNewResponseAccessGitHub]
type accessIdentityProviderNewResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderNewResponseAccessGitHub) implementsAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                  `json:"client_secret"`
	JSON         accessIdentityProviderNewResponseAccessGitHubConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessGitHubConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderNewResponseAccessGitHubConfig]
type accessIdentityProviderNewResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessGitHubType string

const (
	AccessIdentityProviderNewResponseAccessGitHubTypeOnetimepin AccessIdentityProviderNewResponseAccessGitHubType = "onetimepin"
	AccessIdentityProviderNewResponseAccessGitHubTypeAzureAd    AccessIdentityProviderNewResponseAccessGitHubType = "azureAD"
	AccessIdentityProviderNewResponseAccessGitHubTypeSaml       AccessIdentityProviderNewResponseAccessGitHubType = "saml"
	AccessIdentityProviderNewResponseAccessGitHubTypeCentrify   AccessIdentityProviderNewResponseAccessGitHubType = "centrify"
	AccessIdentityProviderNewResponseAccessGitHubTypeFacebook   AccessIdentityProviderNewResponseAccessGitHubType = "facebook"
	AccessIdentityProviderNewResponseAccessGitHubTypeGitHub     AccessIdentityProviderNewResponseAccessGitHubType = "github"
	AccessIdentityProviderNewResponseAccessGitHubTypeGoogleApps AccessIdentityProviderNewResponseAccessGitHubType = "google-apps"
	AccessIdentityProviderNewResponseAccessGitHubTypeGoogle     AccessIdentityProviderNewResponseAccessGitHubType = "google"
	AccessIdentityProviderNewResponseAccessGitHubTypeLinkedin   AccessIdentityProviderNewResponseAccessGitHubType = "linkedin"
	AccessIdentityProviderNewResponseAccessGitHubTypeOidc       AccessIdentityProviderNewResponseAccessGitHubType = "oidc"
	AccessIdentityProviderNewResponseAccessGitHubTypeOkta       AccessIdentityProviderNewResponseAccessGitHubType = "okta"
	AccessIdentityProviderNewResponseAccessGitHubTypeOnelogin   AccessIdentityProviderNewResponseAccessGitHubType = "onelogin"
	AccessIdentityProviderNewResponseAccessGitHubTypePingone    AccessIdentityProviderNewResponseAccessGitHubType = "pingone"
	AccessIdentityProviderNewResponseAccessGitHubTypeYandex     AccessIdentityProviderNewResponseAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewResponseAccessGitHubScimConfig struct {
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
	JSON            accessIdentityProviderNewResponseAccessGitHubScimConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessGitHubScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderNewResponseAccessGitHubScimConfig]
type accessIdentityProviderNewResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderNewResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderNewResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderNewResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderNewResponseAccessGoogleJSON       `json:"-"`
}

// accessIdentityProviderNewResponseAccessGoogleJSON contains the JSON metadata for
// the struct [AccessIdentityProviderNewResponseAccessGoogle]
type accessIdentityProviderNewResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderNewResponseAccessGoogle) implementsAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                  `json:"email_claim_name"`
	JSON           accessIdentityProviderNewResponseAccessGoogleConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessGoogleConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderNewResponseAccessGoogleConfig]
type accessIdentityProviderNewResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessGoogleType string

const (
	AccessIdentityProviderNewResponseAccessGoogleTypeOnetimepin AccessIdentityProviderNewResponseAccessGoogleType = "onetimepin"
	AccessIdentityProviderNewResponseAccessGoogleTypeAzureAd    AccessIdentityProviderNewResponseAccessGoogleType = "azureAD"
	AccessIdentityProviderNewResponseAccessGoogleTypeSaml       AccessIdentityProviderNewResponseAccessGoogleType = "saml"
	AccessIdentityProviderNewResponseAccessGoogleTypeCentrify   AccessIdentityProviderNewResponseAccessGoogleType = "centrify"
	AccessIdentityProviderNewResponseAccessGoogleTypeFacebook   AccessIdentityProviderNewResponseAccessGoogleType = "facebook"
	AccessIdentityProviderNewResponseAccessGoogleTypeGitHub     AccessIdentityProviderNewResponseAccessGoogleType = "github"
	AccessIdentityProviderNewResponseAccessGoogleTypeGoogleApps AccessIdentityProviderNewResponseAccessGoogleType = "google-apps"
	AccessIdentityProviderNewResponseAccessGoogleTypeGoogle     AccessIdentityProviderNewResponseAccessGoogleType = "google"
	AccessIdentityProviderNewResponseAccessGoogleTypeLinkedin   AccessIdentityProviderNewResponseAccessGoogleType = "linkedin"
	AccessIdentityProviderNewResponseAccessGoogleTypeOidc       AccessIdentityProviderNewResponseAccessGoogleType = "oidc"
	AccessIdentityProviderNewResponseAccessGoogleTypeOkta       AccessIdentityProviderNewResponseAccessGoogleType = "okta"
	AccessIdentityProviderNewResponseAccessGoogleTypeOnelogin   AccessIdentityProviderNewResponseAccessGoogleType = "onelogin"
	AccessIdentityProviderNewResponseAccessGoogleTypePingone    AccessIdentityProviderNewResponseAccessGoogleType = "pingone"
	AccessIdentityProviderNewResponseAccessGoogleTypeYandex     AccessIdentityProviderNewResponseAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewResponseAccessGoogleScimConfig struct {
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
	JSON            accessIdentityProviderNewResponseAccessGoogleScimConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessGoogleScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderNewResponseAccessGoogleScimConfig]
type accessIdentityProviderNewResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderNewResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderNewResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderNewResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderNewResponseAccessGoogleAppsJSON       `json:"-"`
}

// accessIdentityProviderNewResponseAccessGoogleAppsJSON contains the JSON metadata
// for the struct [AccessIdentityProviderNewResponseAccessGoogleApps]
type accessIdentityProviderNewResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderNewResponseAccessGoogleApps) implementsAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessGoogleAppsConfig struct {
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
	JSON           accessIdentityProviderNewResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessGoogleAppsConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderNewResponseAccessGoogleAppsConfig]
type accessIdentityProviderNewResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessGoogleAppsType string

const (
	AccessIdentityProviderNewResponseAccessGoogleAppsTypeOnetimepin AccessIdentityProviderNewResponseAccessGoogleAppsType = "onetimepin"
	AccessIdentityProviderNewResponseAccessGoogleAppsTypeAzureAd    AccessIdentityProviderNewResponseAccessGoogleAppsType = "azureAD"
	AccessIdentityProviderNewResponseAccessGoogleAppsTypeSaml       AccessIdentityProviderNewResponseAccessGoogleAppsType = "saml"
	AccessIdentityProviderNewResponseAccessGoogleAppsTypeCentrify   AccessIdentityProviderNewResponseAccessGoogleAppsType = "centrify"
	AccessIdentityProviderNewResponseAccessGoogleAppsTypeFacebook   AccessIdentityProviderNewResponseAccessGoogleAppsType = "facebook"
	AccessIdentityProviderNewResponseAccessGoogleAppsTypeGitHub     AccessIdentityProviderNewResponseAccessGoogleAppsType = "github"
	AccessIdentityProviderNewResponseAccessGoogleAppsTypeGoogleApps AccessIdentityProviderNewResponseAccessGoogleAppsType = "google-apps"
	AccessIdentityProviderNewResponseAccessGoogleAppsTypeGoogle     AccessIdentityProviderNewResponseAccessGoogleAppsType = "google"
	AccessIdentityProviderNewResponseAccessGoogleAppsTypeLinkedin   AccessIdentityProviderNewResponseAccessGoogleAppsType = "linkedin"
	AccessIdentityProviderNewResponseAccessGoogleAppsTypeOidc       AccessIdentityProviderNewResponseAccessGoogleAppsType = "oidc"
	AccessIdentityProviderNewResponseAccessGoogleAppsTypeOkta       AccessIdentityProviderNewResponseAccessGoogleAppsType = "okta"
	AccessIdentityProviderNewResponseAccessGoogleAppsTypeOnelogin   AccessIdentityProviderNewResponseAccessGoogleAppsType = "onelogin"
	AccessIdentityProviderNewResponseAccessGoogleAppsTypePingone    AccessIdentityProviderNewResponseAccessGoogleAppsType = "pingone"
	AccessIdentityProviderNewResponseAccessGoogleAppsTypeYandex     AccessIdentityProviderNewResponseAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewResponseAccessGoogleAppsScimConfig struct {
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
	JSON            accessIdentityProviderNewResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessGoogleAppsScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderNewResponseAccessGoogleAppsScimConfig]
type accessIdentityProviderNewResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderNewResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderNewResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderNewResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderNewResponseAccessLinkedinJSON       `json:"-"`
}

// accessIdentityProviderNewResponseAccessLinkedinJSON contains the JSON metadata
// for the struct [AccessIdentityProviderNewResponseAccessLinkedin]
type accessIdentityProviderNewResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderNewResponseAccessLinkedin) implementsAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                    `json:"client_secret"`
	JSON         accessIdentityProviderNewResponseAccessLinkedinConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessLinkedinConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderNewResponseAccessLinkedinConfig]
type accessIdentityProviderNewResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessLinkedinType string

const (
	AccessIdentityProviderNewResponseAccessLinkedinTypeOnetimepin AccessIdentityProviderNewResponseAccessLinkedinType = "onetimepin"
	AccessIdentityProviderNewResponseAccessLinkedinTypeAzureAd    AccessIdentityProviderNewResponseAccessLinkedinType = "azureAD"
	AccessIdentityProviderNewResponseAccessLinkedinTypeSaml       AccessIdentityProviderNewResponseAccessLinkedinType = "saml"
	AccessIdentityProviderNewResponseAccessLinkedinTypeCentrify   AccessIdentityProviderNewResponseAccessLinkedinType = "centrify"
	AccessIdentityProviderNewResponseAccessLinkedinTypeFacebook   AccessIdentityProviderNewResponseAccessLinkedinType = "facebook"
	AccessIdentityProviderNewResponseAccessLinkedinTypeGitHub     AccessIdentityProviderNewResponseAccessLinkedinType = "github"
	AccessIdentityProviderNewResponseAccessLinkedinTypeGoogleApps AccessIdentityProviderNewResponseAccessLinkedinType = "google-apps"
	AccessIdentityProviderNewResponseAccessLinkedinTypeGoogle     AccessIdentityProviderNewResponseAccessLinkedinType = "google"
	AccessIdentityProviderNewResponseAccessLinkedinTypeLinkedin   AccessIdentityProviderNewResponseAccessLinkedinType = "linkedin"
	AccessIdentityProviderNewResponseAccessLinkedinTypeOidc       AccessIdentityProviderNewResponseAccessLinkedinType = "oidc"
	AccessIdentityProviderNewResponseAccessLinkedinTypeOkta       AccessIdentityProviderNewResponseAccessLinkedinType = "okta"
	AccessIdentityProviderNewResponseAccessLinkedinTypeOnelogin   AccessIdentityProviderNewResponseAccessLinkedinType = "onelogin"
	AccessIdentityProviderNewResponseAccessLinkedinTypePingone    AccessIdentityProviderNewResponseAccessLinkedinType = "pingone"
	AccessIdentityProviderNewResponseAccessLinkedinTypeYandex     AccessIdentityProviderNewResponseAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewResponseAccessLinkedinScimConfig struct {
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
	JSON            accessIdentityProviderNewResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessLinkedinScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderNewResponseAccessLinkedinScimConfig]
type accessIdentityProviderNewResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderNewResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderNewResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderNewResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderNewResponseAccessOidcJSON       `json:"-"`
}

// accessIdentityProviderNewResponseAccessOidcJSON contains the JSON metadata for
// the struct [AccessIdentityProviderNewResponseAccessOidc]
type accessIdentityProviderNewResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderNewResponseAccessOidc) implementsAccessIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessOidcConfig struct {
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
	JSON     accessIdentityProviderNewResponseAccessOidcConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessOidcConfigJSON contains the JSON metadata
// for the struct [AccessIdentityProviderNewResponseAccessOidcConfig]
type accessIdentityProviderNewResponseAccessOidcConfigJSON struct {
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

func (r *AccessIdentityProviderNewResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessOidcType string

const (
	AccessIdentityProviderNewResponseAccessOidcTypeOnetimepin AccessIdentityProviderNewResponseAccessOidcType = "onetimepin"
	AccessIdentityProviderNewResponseAccessOidcTypeAzureAd    AccessIdentityProviderNewResponseAccessOidcType = "azureAD"
	AccessIdentityProviderNewResponseAccessOidcTypeSaml       AccessIdentityProviderNewResponseAccessOidcType = "saml"
	AccessIdentityProviderNewResponseAccessOidcTypeCentrify   AccessIdentityProviderNewResponseAccessOidcType = "centrify"
	AccessIdentityProviderNewResponseAccessOidcTypeFacebook   AccessIdentityProviderNewResponseAccessOidcType = "facebook"
	AccessIdentityProviderNewResponseAccessOidcTypeGitHub     AccessIdentityProviderNewResponseAccessOidcType = "github"
	AccessIdentityProviderNewResponseAccessOidcTypeGoogleApps AccessIdentityProviderNewResponseAccessOidcType = "google-apps"
	AccessIdentityProviderNewResponseAccessOidcTypeGoogle     AccessIdentityProviderNewResponseAccessOidcType = "google"
	AccessIdentityProviderNewResponseAccessOidcTypeLinkedin   AccessIdentityProviderNewResponseAccessOidcType = "linkedin"
	AccessIdentityProviderNewResponseAccessOidcTypeOidc       AccessIdentityProviderNewResponseAccessOidcType = "oidc"
	AccessIdentityProviderNewResponseAccessOidcTypeOkta       AccessIdentityProviderNewResponseAccessOidcType = "okta"
	AccessIdentityProviderNewResponseAccessOidcTypeOnelogin   AccessIdentityProviderNewResponseAccessOidcType = "onelogin"
	AccessIdentityProviderNewResponseAccessOidcTypePingone    AccessIdentityProviderNewResponseAccessOidcType = "pingone"
	AccessIdentityProviderNewResponseAccessOidcTypeYandex     AccessIdentityProviderNewResponseAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewResponseAccessOidcScimConfig struct {
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
	JSON            accessIdentityProviderNewResponseAccessOidcScimConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessOidcScimConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderNewResponseAccessOidcScimConfig]
type accessIdentityProviderNewResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderNewResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderNewResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderNewResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderNewResponseAccessOktaJSON       `json:"-"`
}

// accessIdentityProviderNewResponseAccessOktaJSON contains the JSON metadata for
// the struct [AccessIdentityProviderNewResponseAccessOkta]
type accessIdentityProviderNewResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderNewResponseAccessOkta) implementsAccessIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessOktaConfig struct {
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
	JSON        accessIdentityProviderNewResponseAccessOktaConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessOktaConfigJSON contains the JSON metadata
// for the struct [AccessIdentityProviderNewResponseAccessOktaConfig]
type accessIdentityProviderNewResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessOktaType string

const (
	AccessIdentityProviderNewResponseAccessOktaTypeOnetimepin AccessIdentityProviderNewResponseAccessOktaType = "onetimepin"
	AccessIdentityProviderNewResponseAccessOktaTypeAzureAd    AccessIdentityProviderNewResponseAccessOktaType = "azureAD"
	AccessIdentityProviderNewResponseAccessOktaTypeSaml       AccessIdentityProviderNewResponseAccessOktaType = "saml"
	AccessIdentityProviderNewResponseAccessOktaTypeCentrify   AccessIdentityProviderNewResponseAccessOktaType = "centrify"
	AccessIdentityProviderNewResponseAccessOktaTypeFacebook   AccessIdentityProviderNewResponseAccessOktaType = "facebook"
	AccessIdentityProviderNewResponseAccessOktaTypeGitHub     AccessIdentityProviderNewResponseAccessOktaType = "github"
	AccessIdentityProviderNewResponseAccessOktaTypeGoogleApps AccessIdentityProviderNewResponseAccessOktaType = "google-apps"
	AccessIdentityProviderNewResponseAccessOktaTypeGoogle     AccessIdentityProviderNewResponseAccessOktaType = "google"
	AccessIdentityProviderNewResponseAccessOktaTypeLinkedin   AccessIdentityProviderNewResponseAccessOktaType = "linkedin"
	AccessIdentityProviderNewResponseAccessOktaTypeOidc       AccessIdentityProviderNewResponseAccessOktaType = "oidc"
	AccessIdentityProviderNewResponseAccessOktaTypeOkta       AccessIdentityProviderNewResponseAccessOktaType = "okta"
	AccessIdentityProviderNewResponseAccessOktaTypeOnelogin   AccessIdentityProviderNewResponseAccessOktaType = "onelogin"
	AccessIdentityProviderNewResponseAccessOktaTypePingone    AccessIdentityProviderNewResponseAccessOktaType = "pingone"
	AccessIdentityProviderNewResponseAccessOktaTypeYandex     AccessIdentityProviderNewResponseAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewResponseAccessOktaScimConfig struct {
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
	JSON            accessIdentityProviderNewResponseAccessOktaScimConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessOktaScimConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderNewResponseAccessOktaScimConfig]
type accessIdentityProviderNewResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderNewResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderNewResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderNewResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderNewResponseAccessOneloginJSON       `json:"-"`
}

// accessIdentityProviderNewResponseAccessOneloginJSON contains the JSON metadata
// for the struct [AccessIdentityProviderNewResponseAccessOnelogin]
type accessIdentityProviderNewResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderNewResponseAccessOnelogin) implementsAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessOneloginConfig struct {
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
	JSON            accessIdentityProviderNewResponseAccessOneloginConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessOneloginConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderNewResponseAccessOneloginConfig]
type accessIdentityProviderNewResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessOneloginType string

const (
	AccessIdentityProviderNewResponseAccessOneloginTypeOnetimepin AccessIdentityProviderNewResponseAccessOneloginType = "onetimepin"
	AccessIdentityProviderNewResponseAccessOneloginTypeAzureAd    AccessIdentityProviderNewResponseAccessOneloginType = "azureAD"
	AccessIdentityProviderNewResponseAccessOneloginTypeSaml       AccessIdentityProviderNewResponseAccessOneloginType = "saml"
	AccessIdentityProviderNewResponseAccessOneloginTypeCentrify   AccessIdentityProviderNewResponseAccessOneloginType = "centrify"
	AccessIdentityProviderNewResponseAccessOneloginTypeFacebook   AccessIdentityProviderNewResponseAccessOneloginType = "facebook"
	AccessIdentityProviderNewResponseAccessOneloginTypeGitHub     AccessIdentityProviderNewResponseAccessOneloginType = "github"
	AccessIdentityProviderNewResponseAccessOneloginTypeGoogleApps AccessIdentityProviderNewResponseAccessOneloginType = "google-apps"
	AccessIdentityProviderNewResponseAccessOneloginTypeGoogle     AccessIdentityProviderNewResponseAccessOneloginType = "google"
	AccessIdentityProviderNewResponseAccessOneloginTypeLinkedin   AccessIdentityProviderNewResponseAccessOneloginType = "linkedin"
	AccessIdentityProviderNewResponseAccessOneloginTypeOidc       AccessIdentityProviderNewResponseAccessOneloginType = "oidc"
	AccessIdentityProviderNewResponseAccessOneloginTypeOkta       AccessIdentityProviderNewResponseAccessOneloginType = "okta"
	AccessIdentityProviderNewResponseAccessOneloginTypeOnelogin   AccessIdentityProviderNewResponseAccessOneloginType = "onelogin"
	AccessIdentityProviderNewResponseAccessOneloginTypePingone    AccessIdentityProviderNewResponseAccessOneloginType = "pingone"
	AccessIdentityProviderNewResponseAccessOneloginTypeYandex     AccessIdentityProviderNewResponseAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewResponseAccessOneloginScimConfig struct {
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
	JSON            accessIdentityProviderNewResponseAccessOneloginScimConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessOneloginScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderNewResponseAccessOneloginScimConfig]
type accessIdentityProviderNewResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderNewResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderNewResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderNewResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderNewResponseAccessPingoneJSON       `json:"-"`
}

// accessIdentityProviderNewResponseAccessPingoneJSON contains the JSON metadata
// for the struct [AccessIdentityProviderNewResponseAccessPingone]
type accessIdentityProviderNewResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderNewResponseAccessPingone) implementsAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessPingoneConfig struct {
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
	JSON      accessIdentityProviderNewResponseAccessPingoneConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessPingoneConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderNewResponseAccessPingoneConfig]
type accessIdentityProviderNewResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessPingoneType string

const (
	AccessIdentityProviderNewResponseAccessPingoneTypeOnetimepin AccessIdentityProviderNewResponseAccessPingoneType = "onetimepin"
	AccessIdentityProviderNewResponseAccessPingoneTypeAzureAd    AccessIdentityProviderNewResponseAccessPingoneType = "azureAD"
	AccessIdentityProviderNewResponseAccessPingoneTypeSaml       AccessIdentityProviderNewResponseAccessPingoneType = "saml"
	AccessIdentityProviderNewResponseAccessPingoneTypeCentrify   AccessIdentityProviderNewResponseAccessPingoneType = "centrify"
	AccessIdentityProviderNewResponseAccessPingoneTypeFacebook   AccessIdentityProviderNewResponseAccessPingoneType = "facebook"
	AccessIdentityProviderNewResponseAccessPingoneTypeGitHub     AccessIdentityProviderNewResponseAccessPingoneType = "github"
	AccessIdentityProviderNewResponseAccessPingoneTypeGoogleApps AccessIdentityProviderNewResponseAccessPingoneType = "google-apps"
	AccessIdentityProviderNewResponseAccessPingoneTypeGoogle     AccessIdentityProviderNewResponseAccessPingoneType = "google"
	AccessIdentityProviderNewResponseAccessPingoneTypeLinkedin   AccessIdentityProviderNewResponseAccessPingoneType = "linkedin"
	AccessIdentityProviderNewResponseAccessPingoneTypeOidc       AccessIdentityProviderNewResponseAccessPingoneType = "oidc"
	AccessIdentityProviderNewResponseAccessPingoneTypeOkta       AccessIdentityProviderNewResponseAccessPingoneType = "okta"
	AccessIdentityProviderNewResponseAccessPingoneTypeOnelogin   AccessIdentityProviderNewResponseAccessPingoneType = "onelogin"
	AccessIdentityProviderNewResponseAccessPingoneTypePingone    AccessIdentityProviderNewResponseAccessPingoneType = "pingone"
	AccessIdentityProviderNewResponseAccessPingoneTypeYandex     AccessIdentityProviderNewResponseAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewResponseAccessPingoneScimConfig struct {
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
	JSON            accessIdentityProviderNewResponseAccessPingoneScimConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessPingoneScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderNewResponseAccessPingoneScimConfig]
type accessIdentityProviderNewResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderNewResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderNewResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderNewResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderNewResponseAccessSamlJSON       `json:"-"`
}

// accessIdentityProviderNewResponseAccessSamlJSON contains the JSON metadata for
// the struct [AccessIdentityProviderNewResponseAccessSaml]
type accessIdentityProviderNewResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderNewResponseAccessSaml) implementsAccessIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccessIdentityProviderNewResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                `json:"sso_target_url"`
	JSON         accessIdentityProviderNewResponseAccessSamlConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessSamlConfigJSON contains the JSON metadata
// for the struct [AccessIdentityProviderNewResponseAccessSamlConfig]
type accessIdentityProviderNewResponseAccessSamlConfigJSON struct {
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

func (r *AccessIdentityProviderNewResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                               `json:"header_name"`
	JSON       accessIdentityProviderNewResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessSamlConfigHeaderAttributeJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderNewResponseAccessSamlConfigHeaderAttribute]
type accessIdentityProviderNewResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessSamlType string

const (
	AccessIdentityProviderNewResponseAccessSamlTypeOnetimepin AccessIdentityProviderNewResponseAccessSamlType = "onetimepin"
	AccessIdentityProviderNewResponseAccessSamlTypeAzureAd    AccessIdentityProviderNewResponseAccessSamlType = "azureAD"
	AccessIdentityProviderNewResponseAccessSamlTypeSaml       AccessIdentityProviderNewResponseAccessSamlType = "saml"
	AccessIdentityProviderNewResponseAccessSamlTypeCentrify   AccessIdentityProviderNewResponseAccessSamlType = "centrify"
	AccessIdentityProviderNewResponseAccessSamlTypeFacebook   AccessIdentityProviderNewResponseAccessSamlType = "facebook"
	AccessIdentityProviderNewResponseAccessSamlTypeGitHub     AccessIdentityProviderNewResponseAccessSamlType = "github"
	AccessIdentityProviderNewResponseAccessSamlTypeGoogleApps AccessIdentityProviderNewResponseAccessSamlType = "google-apps"
	AccessIdentityProviderNewResponseAccessSamlTypeGoogle     AccessIdentityProviderNewResponseAccessSamlType = "google"
	AccessIdentityProviderNewResponseAccessSamlTypeLinkedin   AccessIdentityProviderNewResponseAccessSamlType = "linkedin"
	AccessIdentityProviderNewResponseAccessSamlTypeOidc       AccessIdentityProviderNewResponseAccessSamlType = "oidc"
	AccessIdentityProviderNewResponseAccessSamlTypeOkta       AccessIdentityProviderNewResponseAccessSamlType = "okta"
	AccessIdentityProviderNewResponseAccessSamlTypeOnelogin   AccessIdentityProviderNewResponseAccessSamlType = "onelogin"
	AccessIdentityProviderNewResponseAccessSamlTypePingone    AccessIdentityProviderNewResponseAccessSamlType = "pingone"
	AccessIdentityProviderNewResponseAccessSamlTypeYandex     AccessIdentityProviderNewResponseAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewResponseAccessSamlScimConfig struct {
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
	JSON            accessIdentityProviderNewResponseAccessSamlScimConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessSamlScimConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderNewResponseAccessSamlScimConfig]
type accessIdentityProviderNewResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderNewResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderNewResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderNewResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderNewResponseAccessYandexJSON       `json:"-"`
}

// accessIdentityProviderNewResponseAccessYandexJSON contains the JSON metadata for
// the struct [AccessIdentityProviderNewResponseAccessYandex]
type accessIdentityProviderNewResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderNewResponseAccessYandex) implementsAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                  `json:"client_secret"`
	JSON         accessIdentityProviderNewResponseAccessYandexConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessYandexConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderNewResponseAccessYandexConfig]
type accessIdentityProviderNewResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessYandexType string

const (
	AccessIdentityProviderNewResponseAccessYandexTypeOnetimepin AccessIdentityProviderNewResponseAccessYandexType = "onetimepin"
	AccessIdentityProviderNewResponseAccessYandexTypeAzureAd    AccessIdentityProviderNewResponseAccessYandexType = "azureAD"
	AccessIdentityProviderNewResponseAccessYandexTypeSaml       AccessIdentityProviderNewResponseAccessYandexType = "saml"
	AccessIdentityProviderNewResponseAccessYandexTypeCentrify   AccessIdentityProviderNewResponseAccessYandexType = "centrify"
	AccessIdentityProviderNewResponseAccessYandexTypeFacebook   AccessIdentityProviderNewResponseAccessYandexType = "facebook"
	AccessIdentityProviderNewResponseAccessYandexTypeGitHub     AccessIdentityProviderNewResponseAccessYandexType = "github"
	AccessIdentityProviderNewResponseAccessYandexTypeGoogleApps AccessIdentityProviderNewResponseAccessYandexType = "google-apps"
	AccessIdentityProviderNewResponseAccessYandexTypeGoogle     AccessIdentityProviderNewResponseAccessYandexType = "google"
	AccessIdentityProviderNewResponseAccessYandexTypeLinkedin   AccessIdentityProviderNewResponseAccessYandexType = "linkedin"
	AccessIdentityProviderNewResponseAccessYandexTypeOidc       AccessIdentityProviderNewResponseAccessYandexType = "oidc"
	AccessIdentityProviderNewResponseAccessYandexTypeOkta       AccessIdentityProviderNewResponseAccessYandexType = "okta"
	AccessIdentityProviderNewResponseAccessYandexTypeOnelogin   AccessIdentityProviderNewResponseAccessYandexType = "onelogin"
	AccessIdentityProviderNewResponseAccessYandexTypePingone    AccessIdentityProviderNewResponseAccessYandexType = "pingone"
	AccessIdentityProviderNewResponseAccessYandexTypeYandex     AccessIdentityProviderNewResponseAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewResponseAccessYandexScimConfig struct {
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
	JSON            accessIdentityProviderNewResponseAccessYandexScimConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessYandexScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderNewResponseAccessYandexScimConfig]
type accessIdentityProviderNewResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderNewResponseAccessOnetimepinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderNewResponseAccessOnetimepinScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderNewResponseAccessOnetimepinJSON       `json:"-"`
}

// accessIdentityProviderNewResponseAccessOnetimepinJSON contains the JSON metadata
// for the struct [AccessIdentityProviderNewResponseAccessOnetimepin]
type accessIdentityProviderNewResponseAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderNewResponseAccessOnetimepin) implementsAccessIdentityProviderNewResponse() {
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewResponseAccessOnetimepinType string

const (
	AccessIdentityProviderNewResponseAccessOnetimepinTypeOnetimepin AccessIdentityProviderNewResponseAccessOnetimepinType = "onetimepin"
	AccessIdentityProviderNewResponseAccessOnetimepinTypeAzureAd    AccessIdentityProviderNewResponseAccessOnetimepinType = "azureAD"
	AccessIdentityProviderNewResponseAccessOnetimepinTypeSaml       AccessIdentityProviderNewResponseAccessOnetimepinType = "saml"
	AccessIdentityProviderNewResponseAccessOnetimepinTypeCentrify   AccessIdentityProviderNewResponseAccessOnetimepinType = "centrify"
	AccessIdentityProviderNewResponseAccessOnetimepinTypeFacebook   AccessIdentityProviderNewResponseAccessOnetimepinType = "facebook"
	AccessIdentityProviderNewResponseAccessOnetimepinTypeGitHub     AccessIdentityProviderNewResponseAccessOnetimepinType = "github"
	AccessIdentityProviderNewResponseAccessOnetimepinTypeGoogleApps AccessIdentityProviderNewResponseAccessOnetimepinType = "google-apps"
	AccessIdentityProviderNewResponseAccessOnetimepinTypeGoogle     AccessIdentityProviderNewResponseAccessOnetimepinType = "google"
	AccessIdentityProviderNewResponseAccessOnetimepinTypeLinkedin   AccessIdentityProviderNewResponseAccessOnetimepinType = "linkedin"
	AccessIdentityProviderNewResponseAccessOnetimepinTypeOidc       AccessIdentityProviderNewResponseAccessOnetimepinType = "oidc"
	AccessIdentityProviderNewResponseAccessOnetimepinTypeOkta       AccessIdentityProviderNewResponseAccessOnetimepinType = "okta"
	AccessIdentityProviderNewResponseAccessOnetimepinTypeOnelogin   AccessIdentityProviderNewResponseAccessOnetimepinType = "onelogin"
	AccessIdentityProviderNewResponseAccessOnetimepinTypePingone    AccessIdentityProviderNewResponseAccessOnetimepinType = "pingone"
	AccessIdentityProviderNewResponseAccessOnetimepinTypeYandex     AccessIdentityProviderNewResponseAccessOnetimepinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewResponseAccessOnetimepinScimConfig struct {
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
	JSON            accessIdentityProviderNewResponseAccessOnetimepinScimConfigJSON `json:"-"`
}

// accessIdentityProviderNewResponseAccessOnetimepinScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderNewResponseAccessOnetimepinScimConfig]
type accessIdentityProviderNewResponseAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccessIdentityProviderListResponseAccessAzureAd],
// [AccessIdentityProviderListResponseAccessCentrify],
// [AccessIdentityProviderListResponseAccessFacebook],
// [AccessIdentityProviderListResponseAccessGitHub],
// [AccessIdentityProviderListResponseAccessGoogle],
// [AccessIdentityProviderListResponseAccessGoogleApps],
// [AccessIdentityProviderListResponseAccessLinkedin],
// [AccessIdentityProviderListResponseAccessOidc],
// [AccessIdentityProviderListResponseAccessOkta],
// [AccessIdentityProviderListResponseAccessOnelogin],
// [AccessIdentityProviderListResponseAccessPingone],
// [AccessIdentityProviderListResponseAccessSaml] or
// [AccessIdentityProviderListResponseAccessYandex].
type AccessIdentityProviderListResponse interface {
	implementsAccessIdentityProviderListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessIdentityProviderListResponse)(nil)).Elem(), "")
}

type AccessIdentityProviderListResponseAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderListResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderListResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderListResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderListResponseAccessAzureAdJSON       `json:"-"`
}

// accessIdentityProviderListResponseAccessAzureAdJSON contains the JSON metadata
// for the struct [AccessIdentityProviderListResponseAccessAzureAd]
type accessIdentityProviderListResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderListResponseAccessAzureAd) implementsAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessAzureAdConfig struct {
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
	SupportGroups bool                                                      `json:"support_groups"`
	JSON          accessIdentityProviderListResponseAccessAzureAdConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessAzureAdConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessAzureAdConfig]
type accessIdentityProviderListResponseAccessAzureAdConfigJSON struct {
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

func (r *AccessIdentityProviderListResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessAzureAdType string

const (
	AccessIdentityProviderListResponseAccessAzureAdTypeOnetimepin AccessIdentityProviderListResponseAccessAzureAdType = "onetimepin"
	AccessIdentityProviderListResponseAccessAzureAdTypeAzureAd    AccessIdentityProviderListResponseAccessAzureAdType = "azureAD"
	AccessIdentityProviderListResponseAccessAzureAdTypeSaml       AccessIdentityProviderListResponseAccessAzureAdType = "saml"
	AccessIdentityProviderListResponseAccessAzureAdTypeCentrify   AccessIdentityProviderListResponseAccessAzureAdType = "centrify"
	AccessIdentityProviderListResponseAccessAzureAdTypeFacebook   AccessIdentityProviderListResponseAccessAzureAdType = "facebook"
	AccessIdentityProviderListResponseAccessAzureAdTypeGitHub     AccessIdentityProviderListResponseAccessAzureAdType = "github"
	AccessIdentityProviderListResponseAccessAzureAdTypeGoogleApps AccessIdentityProviderListResponseAccessAzureAdType = "google-apps"
	AccessIdentityProviderListResponseAccessAzureAdTypeGoogle     AccessIdentityProviderListResponseAccessAzureAdType = "google"
	AccessIdentityProviderListResponseAccessAzureAdTypeLinkedin   AccessIdentityProviderListResponseAccessAzureAdType = "linkedin"
	AccessIdentityProviderListResponseAccessAzureAdTypeOidc       AccessIdentityProviderListResponseAccessAzureAdType = "oidc"
	AccessIdentityProviderListResponseAccessAzureAdTypeOkta       AccessIdentityProviderListResponseAccessAzureAdType = "okta"
	AccessIdentityProviderListResponseAccessAzureAdTypeOnelogin   AccessIdentityProviderListResponseAccessAzureAdType = "onelogin"
	AccessIdentityProviderListResponseAccessAzureAdTypePingone    AccessIdentityProviderListResponseAccessAzureAdType = "pingone"
	AccessIdentityProviderListResponseAccessAzureAdTypeYandex     AccessIdentityProviderListResponseAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderListResponseAccessAzureAdScimConfig struct {
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
	JSON            accessIdentityProviderListResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessAzureAdScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderListResponseAccessAzureAdScimConfig]
type accessIdentityProviderListResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderListResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderListResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderListResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderListResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderListResponseAccessCentrifyJSON       `json:"-"`
}

// accessIdentityProviderListResponseAccessCentrifyJSON contains the JSON metadata
// for the struct [AccessIdentityProviderListResponseAccessCentrify]
type accessIdentityProviderListResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderListResponseAccessCentrify) implementsAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessCentrifyConfig struct {
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
	EmailClaimName string                                                     `json:"email_claim_name"`
	JSON           accessIdentityProviderListResponseAccessCentrifyConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessCentrifyConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessCentrifyConfig]
type accessIdentityProviderListResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessCentrifyType string

const (
	AccessIdentityProviderListResponseAccessCentrifyTypeOnetimepin AccessIdentityProviderListResponseAccessCentrifyType = "onetimepin"
	AccessIdentityProviderListResponseAccessCentrifyTypeAzureAd    AccessIdentityProviderListResponseAccessCentrifyType = "azureAD"
	AccessIdentityProviderListResponseAccessCentrifyTypeSaml       AccessIdentityProviderListResponseAccessCentrifyType = "saml"
	AccessIdentityProviderListResponseAccessCentrifyTypeCentrify   AccessIdentityProviderListResponseAccessCentrifyType = "centrify"
	AccessIdentityProviderListResponseAccessCentrifyTypeFacebook   AccessIdentityProviderListResponseAccessCentrifyType = "facebook"
	AccessIdentityProviderListResponseAccessCentrifyTypeGitHub     AccessIdentityProviderListResponseAccessCentrifyType = "github"
	AccessIdentityProviderListResponseAccessCentrifyTypeGoogleApps AccessIdentityProviderListResponseAccessCentrifyType = "google-apps"
	AccessIdentityProviderListResponseAccessCentrifyTypeGoogle     AccessIdentityProviderListResponseAccessCentrifyType = "google"
	AccessIdentityProviderListResponseAccessCentrifyTypeLinkedin   AccessIdentityProviderListResponseAccessCentrifyType = "linkedin"
	AccessIdentityProviderListResponseAccessCentrifyTypeOidc       AccessIdentityProviderListResponseAccessCentrifyType = "oidc"
	AccessIdentityProviderListResponseAccessCentrifyTypeOkta       AccessIdentityProviderListResponseAccessCentrifyType = "okta"
	AccessIdentityProviderListResponseAccessCentrifyTypeOnelogin   AccessIdentityProviderListResponseAccessCentrifyType = "onelogin"
	AccessIdentityProviderListResponseAccessCentrifyTypePingone    AccessIdentityProviderListResponseAccessCentrifyType = "pingone"
	AccessIdentityProviderListResponseAccessCentrifyTypeYandex     AccessIdentityProviderListResponseAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderListResponseAccessCentrifyScimConfig struct {
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
	JSON            accessIdentityProviderListResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessCentrifyScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderListResponseAccessCentrifyScimConfig]
type accessIdentityProviderListResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderListResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderListResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderListResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderListResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderListResponseAccessFacebookJSON       `json:"-"`
}

// accessIdentityProviderListResponseAccessFacebookJSON contains the JSON metadata
// for the struct [AccessIdentityProviderListResponseAccessFacebook]
type accessIdentityProviderListResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderListResponseAccessFacebook) implementsAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                     `json:"client_secret"`
	JSON         accessIdentityProviderListResponseAccessFacebookConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessFacebookConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessFacebookConfig]
type accessIdentityProviderListResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessFacebookType string

const (
	AccessIdentityProviderListResponseAccessFacebookTypeOnetimepin AccessIdentityProviderListResponseAccessFacebookType = "onetimepin"
	AccessIdentityProviderListResponseAccessFacebookTypeAzureAd    AccessIdentityProviderListResponseAccessFacebookType = "azureAD"
	AccessIdentityProviderListResponseAccessFacebookTypeSaml       AccessIdentityProviderListResponseAccessFacebookType = "saml"
	AccessIdentityProviderListResponseAccessFacebookTypeCentrify   AccessIdentityProviderListResponseAccessFacebookType = "centrify"
	AccessIdentityProviderListResponseAccessFacebookTypeFacebook   AccessIdentityProviderListResponseAccessFacebookType = "facebook"
	AccessIdentityProviderListResponseAccessFacebookTypeGitHub     AccessIdentityProviderListResponseAccessFacebookType = "github"
	AccessIdentityProviderListResponseAccessFacebookTypeGoogleApps AccessIdentityProviderListResponseAccessFacebookType = "google-apps"
	AccessIdentityProviderListResponseAccessFacebookTypeGoogle     AccessIdentityProviderListResponseAccessFacebookType = "google"
	AccessIdentityProviderListResponseAccessFacebookTypeLinkedin   AccessIdentityProviderListResponseAccessFacebookType = "linkedin"
	AccessIdentityProviderListResponseAccessFacebookTypeOidc       AccessIdentityProviderListResponseAccessFacebookType = "oidc"
	AccessIdentityProviderListResponseAccessFacebookTypeOkta       AccessIdentityProviderListResponseAccessFacebookType = "okta"
	AccessIdentityProviderListResponseAccessFacebookTypeOnelogin   AccessIdentityProviderListResponseAccessFacebookType = "onelogin"
	AccessIdentityProviderListResponseAccessFacebookTypePingone    AccessIdentityProviderListResponseAccessFacebookType = "pingone"
	AccessIdentityProviderListResponseAccessFacebookTypeYandex     AccessIdentityProviderListResponseAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderListResponseAccessFacebookScimConfig struct {
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
	JSON            accessIdentityProviderListResponseAccessFacebookScimConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessFacebookScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderListResponseAccessFacebookScimConfig]
type accessIdentityProviderListResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderListResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderListResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderListResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderListResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderListResponseAccessGitHubJSON       `json:"-"`
}

// accessIdentityProviderListResponseAccessGitHubJSON contains the JSON metadata
// for the struct [AccessIdentityProviderListResponseAccessGitHub]
type accessIdentityProviderListResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderListResponseAccessGitHub) implementsAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                   `json:"client_secret"`
	JSON         accessIdentityProviderListResponseAccessGitHubConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessGitHubConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessGitHubConfig]
type accessIdentityProviderListResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessGitHubType string

const (
	AccessIdentityProviderListResponseAccessGitHubTypeOnetimepin AccessIdentityProviderListResponseAccessGitHubType = "onetimepin"
	AccessIdentityProviderListResponseAccessGitHubTypeAzureAd    AccessIdentityProviderListResponseAccessGitHubType = "azureAD"
	AccessIdentityProviderListResponseAccessGitHubTypeSaml       AccessIdentityProviderListResponseAccessGitHubType = "saml"
	AccessIdentityProviderListResponseAccessGitHubTypeCentrify   AccessIdentityProviderListResponseAccessGitHubType = "centrify"
	AccessIdentityProviderListResponseAccessGitHubTypeFacebook   AccessIdentityProviderListResponseAccessGitHubType = "facebook"
	AccessIdentityProviderListResponseAccessGitHubTypeGitHub     AccessIdentityProviderListResponseAccessGitHubType = "github"
	AccessIdentityProviderListResponseAccessGitHubTypeGoogleApps AccessIdentityProviderListResponseAccessGitHubType = "google-apps"
	AccessIdentityProviderListResponseAccessGitHubTypeGoogle     AccessIdentityProviderListResponseAccessGitHubType = "google"
	AccessIdentityProviderListResponseAccessGitHubTypeLinkedin   AccessIdentityProviderListResponseAccessGitHubType = "linkedin"
	AccessIdentityProviderListResponseAccessGitHubTypeOidc       AccessIdentityProviderListResponseAccessGitHubType = "oidc"
	AccessIdentityProviderListResponseAccessGitHubTypeOkta       AccessIdentityProviderListResponseAccessGitHubType = "okta"
	AccessIdentityProviderListResponseAccessGitHubTypeOnelogin   AccessIdentityProviderListResponseAccessGitHubType = "onelogin"
	AccessIdentityProviderListResponseAccessGitHubTypePingone    AccessIdentityProviderListResponseAccessGitHubType = "pingone"
	AccessIdentityProviderListResponseAccessGitHubTypeYandex     AccessIdentityProviderListResponseAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderListResponseAccessGitHubScimConfig struct {
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
	JSON            accessIdentityProviderListResponseAccessGitHubScimConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessGitHubScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderListResponseAccessGitHubScimConfig]
type accessIdentityProviderListResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderListResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderListResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderListResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderListResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderListResponseAccessGoogleJSON       `json:"-"`
}

// accessIdentityProviderListResponseAccessGoogleJSON contains the JSON metadata
// for the struct [AccessIdentityProviderListResponseAccessGoogle]
type accessIdentityProviderListResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderListResponseAccessGoogle) implementsAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                   `json:"email_claim_name"`
	JSON           accessIdentityProviderListResponseAccessGoogleConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessGoogleConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessGoogleConfig]
type accessIdentityProviderListResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessGoogleType string

const (
	AccessIdentityProviderListResponseAccessGoogleTypeOnetimepin AccessIdentityProviderListResponseAccessGoogleType = "onetimepin"
	AccessIdentityProviderListResponseAccessGoogleTypeAzureAd    AccessIdentityProviderListResponseAccessGoogleType = "azureAD"
	AccessIdentityProviderListResponseAccessGoogleTypeSaml       AccessIdentityProviderListResponseAccessGoogleType = "saml"
	AccessIdentityProviderListResponseAccessGoogleTypeCentrify   AccessIdentityProviderListResponseAccessGoogleType = "centrify"
	AccessIdentityProviderListResponseAccessGoogleTypeFacebook   AccessIdentityProviderListResponseAccessGoogleType = "facebook"
	AccessIdentityProviderListResponseAccessGoogleTypeGitHub     AccessIdentityProviderListResponseAccessGoogleType = "github"
	AccessIdentityProviderListResponseAccessGoogleTypeGoogleApps AccessIdentityProviderListResponseAccessGoogleType = "google-apps"
	AccessIdentityProviderListResponseAccessGoogleTypeGoogle     AccessIdentityProviderListResponseAccessGoogleType = "google"
	AccessIdentityProviderListResponseAccessGoogleTypeLinkedin   AccessIdentityProviderListResponseAccessGoogleType = "linkedin"
	AccessIdentityProviderListResponseAccessGoogleTypeOidc       AccessIdentityProviderListResponseAccessGoogleType = "oidc"
	AccessIdentityProviderListResponseAccessGoogleTypeOkta       AccessIdentityProviderListResponseAccessGoogleType = "okta"
	AccessIdentityProviderListResponseAccessGoogleTypeOnelogin   AccessIdentityProviderListResponseAccessGoogleType = "onelogin"
	AccessIdentityProviderListResponseAccessGoogleTypePingone    AccessIdentityProviderListResponseAccessGoogleType = "pingone"
	AccessIdentityProviderListResponseAccessGoogleTypeYandex     AccessIdentityProviderListResponseAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderListResponseAccessGoogleScimConfig struct {
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
	JSON            accessIdentityProviderListResponseAccessGoogleScimConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessGoogleScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderListResponseAccessGoogleScimConfig]
type accessIdentityProviderListResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderListResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderListResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderListResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderListResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderListResponseAccessGoogleAppsJSON       `json:"-"`
}

// accessIdentityProviderListResponseAccessGoogleAppsJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessGoogleApps]
type accessIdentityProviderListResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderListResponseAccessGoogleApps) implementsAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                       `json:"email_claim_name"`
	JSON           accessIdentityProviderListResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessGoogleAppsConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderListResponseAccessGoogleAppsConfig]
type accessIdentityProviderListResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessGoogleAppsType string

const (
	AccessIdentityProviderListResponseAccessGoogleAppsTypeOnetimepin AccessIdentityProviderListResponseAccessGoogleAppsType = "onetimepin"
	AccessIdentityProviderListResponseAccessGoogleAppsTypeAzureAd    AccessIdentityProviderListResponseAccessGoogleAppsType = "azureAD"
	AccessIdentityProviderListResponseAccessGoogleAppsTypeSaml       AccessIdentityProviderListResponseAccessGoogleAppsType = "saml"
	AccessIdentityProviderListResponseAccessGoogleAppsTypeCentrify   AccessIdentityProviderListResponseAccessGoogleAppsType = "centrify"
	AccessIdentityProviderListResponseAccessGoogleAppsTypeFacebook   AccessIdentityProviderListResponseAccessGoogleAppsType = "facebook"
	AccessIdentityProviderListResponseAccessGoogleAppsTypeGitHub     AccessIdentityProviderListResponseAccessGoogleAppsType = "github"
	AccessIdentityProviderListResponseAccessGoogleAppsTypeGoogleApps AccessIdentityProviderListResponseAccessGoogleAppsType = "google-apps"
	AccessIdentityProviderListResponseAccessGoogleAppsTypeGoogle     AccessIdentityProviderListResponseAccessGoogleAppsType = "google"
	AccessIdentityProviderListResponseAccessGoogleAppsTypeLinkedin   AccessIdentityProviderListResponseAccessGoogleAppsType = "linkedin"
	AccessIdentityProviderListResponseAccessGoogleAppsTypeOidc       AccessIdentityProviderListResponseAccessGoogleAppsType = "oidc"
	AccessIdentityProviderListResponseAccessGoogleAppsTypeOkta       AccessIdentityProviderListResponseAccessGoogleAppsType = "okta"
	AccessIdentityProviderListResponseAccessGoogleAppsTypeOnelogin   AccessIdentityProviderListResponseAccessGoogleAppsType = "onelogin"
	AccessIdentityProviderListResponseAccessGoogleAppsTypePingone    AccessIdentityProviderListResponseAccessGoogleAppsType = "pingone"
	AccessIdentityProviderListResponseAccessGoogleAppsTypeYandex     AccessIdentityProviderListResponseAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderListResponseAccessGoogleAppsScimConfig struct {
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
	JSON            accessIdentityProviderListResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessGoogleAppsScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderListResponseAccessGoogleAppsScimConfig]
type accessIdentityProviderListResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderListResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderListResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderListResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderListResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderListResponseAccessLinkedinJSON       `json:"-"`
}

// accessIdentityProviderListResponseAccessLinkedinJSON contains the JSON metadata
// for the struct [AccessIdentityProviderListResponseAccessLinkedin]
type accessIdentityProviderListResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderListResponseAccessLinkedin) implementsAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                     `json:"client_secret"`
	JSON         accessIdentityProviderListResponseAccessLinkedinConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessLinkedinConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessLinkedinConfig]
type accessIdentityProviderListResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessLinkedinType string

const (
	AccessIdentityProviderListResponseAccessLinkedinTypeOnetimepin AccessIdentityProviderListResponseAccessLinkedinType = "onetimepin"
	AccessIdentityProviderListResponseAccessLinkedinTypeAzureAd    AccessIdentityProviderListResponseAccessLinkedinType = "azureAD"
	AccessIdentityProviderListResponseAccessLinkedinTypeSaml       AccessIdentityProviderListResponseAccessLinkedinType = "saml"
	AccessIdentityProviderListResponseAccessLinkedinTypeCentrify   AccessIdentityProviderListResponseAccessLinkedinType = "centrify"
	AccessIdentityProviderListResponseAccessLinkedinTypeFacebook   AccessIdentityProviderListResponseAccessLinkedinType = "facebook"
	AccessIdentityProviderListResponseAccessLinkedinTypeGitHub     AccessIdentityProviderListResponseAccessLinkedinType = "github"
	AccessIdentityProviderListResponseAccessLinkedinTypeGoogleApps AccessIdentityProviderListResponseAccessLinkedinType = "google-apps"
	AccessIdentityProviderListResponseAccessLinkedinTypeGoogle     AccessIdentityProviderListResponseAccessLinkedinType = "google"
	AccessIdentityProviderListResponseAccessLinkedinTypeLinkedin   AccessIdentityProviderListResponseAccessLinkedinType = "linkedin"
	AccessIdentityProviderListResponseAccessLinkedinTypeOidc       AccessIdentityProviderListResponseAccessLinkedinType = "oidc"
	AccessIdentityProviderListResponseAccessLinkedinTypeOkta       AccessIdentityProviderListResponseAccessLinkedinType = "okta"
	AccessIdentityProviderListResponseAccessLinkedinTypeOnelogin   AccessIdentityProviderListResponseAccessLinkedinType = "onelogin"
	AccessIdentityProviderListResponseAccessLinkedinTypePingone    AccessIdentityProviderListResponseAccessLinkedinType = "pingone"
	AccessIdentityProviderListResponseAccessLinkedinTypeYandex     AccessIdentityProviderListResponseAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderListResponseAccessLinkedinScimConfig struct {
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
	JSON            accessIdentityProviderListResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessLinkedinScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderListResponseAccessLinkedinScimConfig]
type accessIdentityProviderListResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderListResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderListResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderListResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderListResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderListResponseAccessOidcJSON       `json:"-"`
}

// accessIdentityProviderListResponseAccessOidcJSON contains the JSON metadata for
// the struct [AccessIdentityProviderListResponseAccessOidc]
type accessIdentityProviderListResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderListResponseAccessOidc) implementsAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessOidcConfig struct {
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
	TokenURL string                                                 `json:"token_url"`
	JSON     accessIdentityProviderListResponseAccessOidcConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessOidcConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessOidcConfig]
type accessIdentityProviderListResponseAccessOidcConfigJSON struct {
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

func (r *AccessIdentityProviderListResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessOidcType string

const (
	AccessIdentityProviderListResponseAccessOidcTypeOnetimepin AccessIdentityProviderListResponseAccessOidcType = "onetimepin"
	AccessIdentityProviderListResponseAccessOidcTypeAzureAd    AccessIdentityProviderListResponseAccessOidcType = "azureAD"
	AccessIdentityProviderListResponseAccessOidcTypeSaml       AccessIdentityProviderListResponseAccessOidcType = "saml"
	AccessIdentityProviderListResponseAccessOidcTypeCentrify   AccessIdentityProviderListResponseAccessOidcType = "centrify"
	AccessIdentityProviderListResponseAccessOidcTypeFacebook   AccessIdentityProviderListResponseAccessOidcType = "facebook"
	AccessIdentityProviderListResponseAccessOidcTypeGitHub     AccessIdentityProviderListResponseAccessOidcType = "github"
	AccessIdentityProviderListResponseAccessOidcTypeGoogleApps AccessIdentityProviderListResponseAccessOidcType = "google-apps"
	AccessIdentityProviderListResponseAccessOidcTypeGoogle     AccessIdentityProviderListResponseAccessOidcType = "google"
	AccessIdentityProviderListResponseAccessOidcTypeLinkedin   AccessIdentityProviderListResponseAccessOidcType = "linkedin"
	AccessIdentityProviderListResponseAccessOidcTypeOidc       AccessIdentityProviderListResponseAccessOidcType = "oidc"
	AccessIdentityProviderListResponseAccessOidcTypeOkta       AccessIdentityProviderListResponseAccessOidcType = "okta"
	AccessIdentityProviderListResponseAccessOidcTypeOnelogin   AccessIdentityProviderListResponseAccessOidcType = "onelogin"
	AccessIdentityProviderListResponseAccessOidcTypePingone    AccessIdentityProviderListResponseAccessOidcType = "pingone"
	AccessIdentityProviderListResponseAccessOidcTypeYandex     AccessIdentityProviderListResponseAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderListResponseAccessOidcScimConfig struct {
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
	UserDeprovision bool                                                       `json:"user_deprovision"`
	JSON            accessIdentityProviderListResponseAccessOidcScimConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessOidcScimConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessOidcScimConfig]
type accessIdentityProviderListResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderListResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderListResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderListResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderListResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderListResponseAccessOktaJSON       `json:"-"`
}

// accessIdentityProviderListResponseAccessOktaJSON contains the JSON metadata for
// the struct [AccessIdentityProviderListResponseAccessOkta]
type accessIdentityProviderListResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderListResponseAccessOkta) implementsAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessOktaConfig struct {
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
	OktaAccount string                                                 `json:"okta_account"`
	JSON        accessIdentityProviderListResponseAccessOktaConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessOktaConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessOktaConfig]
type accessIdentityProviderListResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessOktaType string

const (
	AccessIdentityProviderListResponseAccessOktaTypeOnetimepin AccessIdentityProviderListResponseAccessOktaType = "onetimepin"
	AccessIdentityProviderListResponseAccessOktaTypeAzureAd    AccessIdentityProviderListResponseAccessOktaType = "azureAD"
	AccessIdentityProviderListResponseAccessOktaTypeSaml       AccessIdentityProviderListResponseAccessOktaType = "saml"
	AccessIdentityProviderListResponseAccessOktaTypeCentrify   AccessIdentityProviderListResponseAccessOktaType = "centrify"
	AccessIdentityProviderListResponseAccessOktaTypeFacebook   AccessIdentityProviderListResponseAccessOktaType = "facebook"
	AccessIdentityProviderListResponseAccessOktaTypeGitHub     AccessIdentityProviderListResponseAccessOktaType = "github"
	AccessIdentityProviderListResponseAccessOktaTypeGoogleApps AccessIdentityProviderListResponseAccessOktaType = "google-apps"
	AccessIdentityProviderListResponseAccessOktaTypeGoogle     AccessIdentityProviderListResponseAccessOktaType = "google"
	AccessIdentityProviderListResponseAccessOktaTypeLinkedin   AccessIdentityProviderListResponseAccessOktaType = "linkedin"
	AccessIdentityProviderListResponseAccessOktaTypeOidc       AccessIdentityProviderListResponseAccessOktaType = "oidc"
	AccessIdentityProviderListResponseAccessOktaTypeOkta       AccessIdentityProviderListResponseAccessOktaType = "okta"
	AccessIdentityProviderListResponseAccessOktaTypeOnelogin   AccessIdentityProviderListResponseAccessOktaType = "onelogin"
	AccessIdentityProviderListResponseAccessOktaTypePingone    AccessIdentityProviderListResponseAccessOktaType = "pingone"
	AccessIdentityProviderListResponseAccessOktaTypeYandex     AccessIdentityProviderListResponseAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderListResponseAccessOktaScimConfig struct {
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
	UserDeprovision bool                                                       `json:"user_deprovision"`
	JSON            accessIdentityProviderListResponseAccessOktaScimConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessOktaScimConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessOktaScimConfig]
type accessIdentityProviderListResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderListResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderListResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderListResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderListResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderListResponseAccessOneloginJSON       `json:"-"`
}

// accessIdentityProviderListResponseAccessOneloginJSON contains the JSON metadata
// for the struct [AccessIdentityProviderListResponseAccessOnelogin]
type accessIdentityProviderListResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderListResponseAccessOnelogin) implementsAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                     `json:"onelogin_account"`
	JSON            accessIdentityProviderListResponseAccessOneloginConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessOneloginConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessOneloginConfig]
type accessIdentityProviderListResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessOneloginType string

const (
	AccessIdentityProviderListResponseAccessOneloginTypeOnetimepin AccessIdentityProviderListResponseAccessOneloginType = "onetimepin"
	AccessIdentityProviderListResponseAccessOneloginTypeAzureAd    AccessIdentityProviderListResponseAccessOneloginType = "azureAD"
	AccessIdentityProviderListResponseAccessOneloginTypeSaml       AccessIdentityProviderListResponseAccessOneloginType = "saml"
	AccessIdentityProviderListResponseAccessOneloginTypeCentrify   AccessIdentityProviderListResponseAccessOneloginType = "centrify"
	AccessIdentityProviderListResponseAccessOneloginTypeFacebook   AccessIdentityProviderListResponseAccessOneloginType = "facebook"
	AccessIdentityProviderListResponseAccessOneloginTypeGitHub     AccessIdentityProviderListResponseAccessOneloginType = "github"
	AccessIdentityProviderListResponseAccessOneloginTypeGoogleApps AccessIdentityProviderListResponseAccessOneloginType = "google-apps"
	AccessIdentityProviderListResponseAccessOneloginTypeGoogle     AccessIdentityProviderListResponseAccessOneloginType = "google"
	AccessIdentityProviderListResponseAccessOneloginTypeLinkedin   AccessIdentityProviderListResponseAccessOneloginType = "linkedin"
	AccessIdentityProviderListResponseAccessOneloginTypeOidc       AccessIdentityProviderListResponseAccessOneloginType = "oidc"
	AccessIdentityProviderListResponseAccessOneloginTypeOkta       AccessIdentityProviderListResponseAccessOneloginType = "okta"
	AccessIdentityProviderListResponseAccessOneloginTypeOnelogin   AccessIdentityProviderListResponseAccessOneloginType = "onelogin"
	AccessIdentityProviderListResponseAccessOneloginTypePingone    AccessIdentityProviderListResponseAccessOneloginType = "pingone"
	AccessIdentityProviderListResponseAccessOneloginTypeYandex     AccessIdentityProviderListResponseAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderListResponseAccessOneloginScimConfig struct {
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
	JSON            accessIdentityProviderListResponseAccessOneloginScimConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessOneloginScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderListResponseAccessOneloginScimConfig]
type accessIdentityProviderListResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderListResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderListResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderListResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderListResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderListResponseAccessPingoneJSON       `json:"-"`
}

// accessIdentityProviderListResponseAccessPingoneJSON contains the JSON metadata
// for the struct [AccessIdentityProviderListResponseAccessPingone]
type accessIdentityProviderListResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderListResponseAccessPingone) implementsAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                    `json:"ping_env_id"`
	JSON      accessIdentityProviderListResponseAccessPingoneConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessPingoneConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessPingoneConfig]
type accessIdentityProviderListResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessPingoneType string

const (
	AccessIdentityProviderListResponseAccessPingoneTypeOnetimepin AccessIdentityProviderListResponseAccessPingoneType = "onetimepin"
	AccessIdentityProviderListResponseAccessPingoneTypeAzureAd    AccessIdentityProviderListResponseAccessPingoneType = "azureAD"
	AccessIdentityProviderListResponseAccessPingoneTypeSaml       AccessIdentityProviderListResponseAccessPingoneType = "saml"
	AccessIdentityProviderListResponseAccessPingoneTypeCentrify   AccessIdentityProviderListResponseAccessPingoneType = "centrify"
	AccessIdentityProviderListResponseAccessPingoneTypeFacebook   AccessIdentityProviderListResponseAccessPingoneType = "facebook"
	AccessIdentityProviderListResponseAccessPingoneTypeGitHub     AccessIdentityProviderListResponseAccessPingoneType = "github"
	AccessIdentityProviderListResponseAccessPingoneTypeGoogleApps AccessIdentityProviderListResponseAccessPingoneType = "google-apps"
	AccessIdentityProviderListResponseAccessPingoneTypeGoogle     AccessIdentityProviderListResponseAccessPingoneType = "google"
	AccessIdentityProviderListResponseAccessPingoneTypeLinkedin   AccessIdentityProviderListResponseAccessPingoneType = "linkedin"
	AccessIdentityProviderListResponseAccessPingoneTypeOidc       AccessIdentityProviderListResponseAccessPingoneType = "oidc"
	AccessIdentityProviderListResponseAccessPingoneTypeOkta       AccessIdentityProviderListResponseAccessPingoneType = "okta"
	AccessIdentityProviderListResponseAccessPingoneTypeOnelogin   AccessIdentityProviderListResponseAccessPingoneType = "onelogin"
	AccessIdentityProviderListResponseAccessPingoneTypePingone    AccessIdentityProviderListResponseAccessPingoneType = "pingone"
	AccessIdentityProviderListResponseAccessPingoneTypeYandex     AccessIdentityProviderListResponseAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderListResponseAccessPingoneScimConfig struct {
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
	JSON            accessIdentityProviderListResponseAccessPingoneScimConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessPingoneScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderListResponseAccessPingoneScimConfig]
type accessIdentityProviderListResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderListResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderListResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderListResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderListResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderListResponseAccessSamlJSON       `json:"-"`
}

// accessIdentityProviderListResponseAccessSamlJSON contains the JSON metadata for
// the struct [AccessIdentityProviderListResponseAccessSaml]
type accessIdentityProviderListResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderListResponseAccessSaml) implementsAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccessIdentityProviderListResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                 `json:"sso_target_url"`
	JSON         accessIdentityProviderListResponseAccessSamlConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessSamlConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessSamlConfig]
type accessIdentityProviderListResponseAccessSamlConfigJSON struct {
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

func (r *AccessIdentityProviderListResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderListResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                `json:"header_name"`
	JSON       accessIdentityProviderListResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessSamlConfigHeaderAttributeJSON contains
// the JSON metadata for the struct
// [AccessIdentityProviderListResponseAccessSamlConfigHeaderAttribute]
type accessIdentityProviderListResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessSamlType string

const (
	AccessIdentityProviderListResponseAccessSamlTypeOnetimepin AccessIdentityProviderListResponseAccessSamlType = "onetimepin"
	AccessIdentityProviderListResponseAccessSamlTypeAzureAd    AccessIdentityProviderListResponseAccessSamlType = "azureAD"
	AccessIdentityProviderListResponseAccessSamlTypeSaml       AccessIdentityProviderListResponseAccessSamlType = "saml"
	AccessIdentityProviderListResponseAccessSamlTypeCentrify   AccessIdentityProviderListResponseAccessSamlType = "centrify"
	AccessIdentityProviderListResponseAccessSamlTypeFacebook   AccessIdentityProviderListResponseAccessSamlType = "facebook"
	AccessIdentityProviderListResponseAccessSamlTypeGitHub     AccessIdentityProviderListResponseAccessSamlType = "github"
	AccessIdentityProviderListResponseAccessSamlTypeGoogleApps AccessIdentityProviderListResponseAccessSamlType = "google-apps"
	AccessIdentityProviderListResponseAccessSamlTypeGoogle     AccessIdentityProviderListResponseAccessSamlType = "google"
	AccessIdentityProviderListResponseAccessSamlTypeLinkedin   AccessIdentityProviderListResponseAccessSamlType = "linkedin"
	AccessIdentityProviderListResponseAccessSamlTypeOidc       AccessIdentityProviderListResponseAccessSamlType = "oidc"
	AccessIdentityProviderListResponseAccessSamlTypeOkta       AccessIdentityProviderListResponseAccessSamlType = "okta"
	AccessIdentityProviderListResponseAccessSamlTypeOnelogin   AccessIdentityProviderListResponseAccessSamlType = "onelogin"
	AccessIdentityProviderListResponseAccessSamlTypePingone    AccessIdentityProviderListResponseAccessSamlType = "pingone"
	AccessIdentityProviderListResponseAccessSamlTypeYandex     AccessIdentityProviderListResponseAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderListResponseAccessSamlScimConfig struct {
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
	UserDeprovision bool                                                       `json:"user_deprovision"`
	JSON            accessIdentityProviderListResponseAccessSamlScimConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessSamlScimConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessSamlScimConfig]
type accessIdentityProviderListResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderListResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderListResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderListResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderListResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderListResponseAccessYandexJSON       `json:"-"`
}

// accessIdentityProviderListResponseAccessYandexJSON contains the JSON metadata
// for the struct [AccessIdentityProviderListResponseAccessYandex]
type accessIdentityProviderListResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderListResponseAccessYandex) implementsAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                   `json:"client_secret"`
	JSON         accessIdentityProviderListResponseAccessYandexConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessYandexConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseAccessYandexConfig]
type accessIdentityProviderListResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderListResponseAccessYandexType string

const (
	AccessIdentityProviderListResponseAccessYandexTypeOnetimepin AccessIdentityProviderListResponseAccessYandexType = "onetimepin"
	AccessIdentityProviderListResponseAccessYandexTypeAzureAd    AccessIdentityProviderListResponseAccessYandexType = "azureAD"
	AccessIdentityProviderListResponseAccessYandexTypeSaml       AccessIdentityProviderListResponseAccessYandexType = "saml"
	AccessIdentityProviderListResponseAccessYandexTypeCentrify   AccessIdentityProviderListResponseAccessYandexType = "centrify"
	AccessIdentityProviderListResponseAccessYandexTypeFacebook   AccessIdentityProviderListResponseAccessYandexType = "facebook"
	AccessIdentityProviderListResponseAccessYandexTypeGitHub     AccessIdentityProviderListResponseAccessYandexType = "github"
	AccessIdentityProviderListResponseAccessYandexTypeGoogleApps AccessIdentityProviderListResponseAccessYandexType = "google-apps"
	AccessIdentityProviderListResponseAccessYandexTypeGoogle     AccessIdentityProviderListResponseAccessYandexType = "google"
	AccessIdentityProviderListResponseAccessYandexTypeLinkedin   AccessIdentityProviderListResponseAccessYandexType = "linkedin"
	AccessIdentityProviderListResponseAccessYandexTypeOidc       AccessIdentityProviderListResponseAccessYandexType = "oidc"
	AccessIdentityProviderListResponseAccessYandexTypeOkta       AccessIdentityProviderListResponseAccessYandexType = "okta"
	AccessIdentityProviderListResponseAccessYandexTypeOnelogin   AccessIdentityProviderListResponseAccessYandexType = "onelogin"
	AccessIdentityProviderListResponseAccessYandexTypePingone    AccessIdentityProviderListResponseAccessYandexType = "pingone"
	AccessIdentityProviderListResponseAccessYandexTypeYandex     AccessIdentityProviderListResponseAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderListResponseAccessYandexScimConfig struct {
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
	JSON            accessIdentityProviderListResponseAccessYandexScimConfigJSON `json:"-"`
}

// accessIdentityProviderListResponseAccessYandexScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderListResponseAccessYandexScimConfig]
type accessIdentityProviderListResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderGetResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderGetResponseAccessAzureAdJSON       `json:"-"`
}

// accessIdentityProviderGetResponseAccessAzureAdJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessAzureAd]
type accessIdentityProviderGetResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessAzureAd) implementsAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderGetResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderGetResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderGetResponseAccessCentrifyJSON       `json:"-"`
}

// accessIdentityProviderGetResponseAccessCentrifyJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessCentrify]
type accessIdentityProviderGetResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessCentrify) implementsAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderGetResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderGetResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderGetResponseAccessFacebookJSON       `json:"-"`
}

// accessIdentityProviderGetResponseAccessFacebookJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessFacebook]
type accessIdentityProviderGetResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessFacebook) implementsAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderGetResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderGetResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderGetResponseAccessGitHubJSON       `json:"-"`
}

// accessIdentityProviderGetResponseAccessGitHubJSON contains the JSON metadata for
// the struct [AccessIdentityProviderGetResponseAccessGitHub]
type accessIdentityProviderGetResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessGitHub) implementsAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderGetResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderGetResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderGetResponseAccessGoogleJSON       `json:"-"`
}

// accessIdentityProviderGetResponseAccessGoogleJSON contains the JSON metadata for
// the struct [AccessIdentityProviderGetResponseAccessGoogle]
type accessIdentityProviderGetResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessGoogle) implementsAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderGetResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderGetResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderGetResponseAccessGoogleAppsJSON       `json:"-"`
}

// accessIdentityProviderGetResponseAccessGoogleAppsJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessGoogleApps]
type accessIdentityProviderGetResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessGoogleApps) implementsAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderGetResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderGetResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderGetResponseAccessLinkedinJSON       `json:"-"`
}

// accessIdentityProviderGetResponseAccessLinkedinJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessLinkedin]
type accessIdentityProviderGetResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessLinkedin) implementsAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderGetResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderGetResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderGetResponseAccessOidcJSON       `json:"-"`
}

// accessIdentityProviderGetResponseAccessOidcJSON contains the JSON metadata for
// the struct [AccessIdentityProviderGetResponseAccessOidc]
type accessIdentityProviderGetResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessOidc) implementsAccessIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderGetResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderGetResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderGetResponseAccessOktaJSON       `json:"-"`
}

// accessIdentityProviderGetResponseAccessOktaJSON contains the JSON metadata for
// the struct [AccessIdentityProviderGetResponseAccessOkta]
type accessIdentityProviderGetResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessOkta) implementsAccessIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderGetResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderGetResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderGetResponseAccessOneloginJSON       `json:"-"`
}

// accessIdentityProviderGetResponseAccessOneloginJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessOnelogin]
type accessIdentityProviderGetResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessOnelogin) implementsAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderGetResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderGetResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderGetResponseAccessPingoneJSON       `json:"-"`
}

// accessIdentityProviderGetResponseAccessPingoneJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessPingone]
type accessIdentityProviderGetResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessPingone) implementsAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderGetResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderGetResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderGetResponseAccessSamlJSON       `json:"-"`
}

// accessIdentityProviderGetResponseAccessSamlJSON contains the JSON metadata for
// the struct [AccessIdentityProviderGetResponseAccessSaml]
type accessIdentityProviderGetResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessSaml) implementsAccessIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderGetResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderGetResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderGetResponseAccessYandexJSON       `json:"-"`
}

// accessIdentityProviderGetResponseAccessYandexJSON contains the JSON metadata for
// the struct [AccessIdentityProviderGetResponseAccessYandex]
type accessIdentityProviderGetResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessYandex) implementsAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderGetResponseAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderGetResponseAccessOnetimepinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderGetResponseAccessOnetimepinScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderGetResponseAccessOnetimepinJSON       `json:"-"`
}

// accessIdentityProviderGetResponseAccessOnetimepinJSON contains the JSON metadata
// for the struct [AccessIdentityProviderGetResponseAccessOnetimepin]
type accessIdentityProviderGetResponseAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderGetResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderGetResponseAccessOnetimepin) implementsAccessIdentityProviderGetResponse() {
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderGetResponseAccessOnetimepinType string

const (
	AccessIdentityProviderGetResponseAccessOnetimepinTypeOnetimepin AccessIdentityProviderGetResponseAccessOnetimepinType = "onetimepin"
	AccessIdentityProviderGetResponseAccessOnetimepinTypeAzureAd    AccessIdentityProviderGetResponseAccessOnetimepinType = "azureAD"
	AccessIdentityProviderGetResponseAccessOnetimepinTypeSaml       AccessIdentityProviderGetResponseAccessOnetimepinType = "saml"
	AccessIdentityProviderGetResponseAccessOnetimepinTypeCentrify   AccessIdentityProviderGetResponseAccessOnetimepinType = "centrify"
	AccessIdentityProviderGetResponseAccessOnetimepinTypeFacebook   AccessIdentityProviderGetResponseAccessOnetimepinType = "facebook"
	AccessIdentityProviderGetResponseAccessOnetimepinTypeGitHub     AccessIdentityProviderGetResponseAccessOnetimepinType = "github"
	AccessIdentityProviderGetResponseAccessOnetimepinTypeGoogleApps AccessIdentityProviderGetResponseAccessOnetimepinType = "google-apps"
	AccessIdentityProviderGetResponseAccessOnetimepinTypeGoogle     AccessIdentityProviderGetResponseAccessOnetimepinType = "google"
	AccessIdentityProviderGetResponseAccessOnetimepinTypeLinkedin   AccessIdentityProviderGetResponseAccessOnetimepinType = "linkedin"
	AccessIdentityProviderGetResponseAccessOnetimepinTypeOidc       AccessIdentityProviderGetResponseAccessOnetimepinType = "oidc"
	AccessIdentityProviderGetResponseAccessOnetimepinTypeOkta       AccessIdentityProviderGetResponseAccessOnetimepinType = "okta"
	AccessIdentityProviderGetResponseAccessOnetimepinTypeOnelogin   AccessIdentityProviderGetResponseAccessOnetimepinType = "onelogin"
	AccessIdentityProviderGetResponseAccessOnetimepinTypePingone    AccessIdentityProviderGetResponseAccessOnetimepinType = "pingone"
	AccessIdentityProviderGetResponseAccessOnetimepinTypeYandex     AccessIdentityProviderGetResponseAccessOnetimepinType = "yandex"
)

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

// Union satisfied by [AccessIdentityProviderReplaceResponseAccessAzureAd],
// [AccessIdentityProviderReplaceResponseAccessCentrify],
// [AccessIdentityProviderReplaceResponseAccessFacebook],
// [AccessIdentityProviderReplaceResponseAccessGitHub],
// [AccessIdentityProviderReplaceResponseAccessGoogle],
// [AccessIdentityProviderReplaceResponseAccessGoogleApps],
// [AccessIdentityProviderReplaceResponseAccessLinkedin],
// [AccessIdentityProviderReplaceResponseAccessOidc],
// [AccessIdentityProviderReplaceResponseAccessOkta],
// [AccessIdentityProviderReplaceResponseAccessOnelogin],
// [AccessIdentityProviderReplaceResponseAccessPingone],
// [AccessIdentityProviderReplaceResponseAccessSaml],
// [AccessIdentityProviderReplaceResponseAccessYandex] or
// [AccessIdentityProviderReplaceResponseAccessOnetimepin].
type AccessIdentityProviderReplaceResponse interface {
	implementsAccessIdentityProviderReplaceResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessIdentityProviderReplaceResponse)(nil)).Elem(), "")
}

type AccessIdentityProviderReplaceResponseAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderReplaceResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderReplaceResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderReplaceResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderReplaceResponseAccessAzureAdJSON       `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessAzureAdJSON contains the JSON
// metadata for the struct [AccessIdentityProviderReplaceResponseAccessAzureAd]
type accessIdentityProviderReplaceResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderReplaceResponseAccessAzureAd) implementsAccessIdentityProviderReplaceResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessAzureAdConfig struct {
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
	SupportGroups bool                                                         `json:"support_groups"`
	JSON          accessIdentityProviderReplaceResponseAccessAzureAdConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessAzureAdConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessAzureAdConfig]
type accessIdentityProviderReplaceResponseAccessAzureAdConfigJSON struct {
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

func (r *AccessIdentityProviderReplaceResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessAzureAdType string

const (
	AccessIdentityProviderReplaceResponseAccessAzureAdTypeOnetimepin AccessIdentityProviderReplaceResponseAccessAzureAdType = "onetimepin"
	AccessIdentityProviderReplaceResponseAccessAzureAdTypeAzureAd    AccessIdentityProviderReplaceResponseAccessAzureAdType = "azureAD"
	AccessIdentityProviderReplaceResponseAccessAzureAdTypeSaml       AccessIdentityProviderReplaceResponseAccessAzureAdType = "saml"
	AccessIdentityProviderReplaceResponseAccessAzureAdTypeCentrify   AccessIdentityProviderReplaceResponseAccessAzureAdType = "centrify"
	AccessIdentityProviderReplaceResponseAccessAzureAdTypeFacebook   AccessIdentityProviderReplaceResponseAccessAzureAdType = "facebook"
	AccessIdentityProviderReplaceResponseAccessAzureAdTypeGitHub     AccessIdentityProviderReplaceResponseAccessAzureAdType = "github"
	AccessIdentityProviderReplaceResponseAccessAzureAdTypeGoogleApps AccessIdentityProviderReplaceResponseAccessAzureAdType = "google-apps"
	AccessIdentityProviderReplaceResponseAccessAzureAdTypeGoogle     AccessIdentityProviderReplaceResponseAccessAzureAdType = "google"
	AccessIdentityProviderReplaceResponseAccessAzureAdTypeLinkedin   AccessIdentityProviderReplaceResponseAccessAzureAdType = "linkedin"
	AccessIdentityProviderReplaceResponseAccessAzureAdTypeOidc       AccessIdentityProviderReplaceResponseAccessAzureAdType = "oidc"
	AccessIdentityProviderReplaceResponseAccessAzureAdTypeOkta       AccessIdentityProviderReplaceResponseAccessAzureAdType = "okta"
	AccessIdentityProviderReplaceResponseAccessAzureAdTypeOnelogin   AccessIdentityProviderReplaceResponseAccessAzureAdType = "onelogin"
	AccessIdentityProviderReplaceResponseAccessAzureAdTypePingone    AccessIdentityProviderReplaceResponseAccessAzureAdType = "pingone"
	AccessIdentityProviderReplaceResponseAccessAzureAdTypeYandex     AccessIdentityProviderReplaceResponseAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceResponseAccessAzureAdScimConfig struct {
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
	JSON            accessIdentityProviderReplaceResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessAzureAdScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessAzureAdScimConfig]
type accessIdentityProviderReplaceResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderReplaceResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderReplaceResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderReplaceResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderReplaceResponseAccessCentrifyJSON       `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessCentrifyJSON contains the JSON
// metadata for the struct [AccessIdentityProviderReplaceResponseAccessCentrify]
type accessIdentityProviderReplaceResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderReplaceResponseAccessCentrify) implementsAccessIdentityProviderReplaceResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessCentrifyConfig struct {
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
	EmailClaimName string                                                        `json:"email_claim_name"`
	JSON           accessIdentityProviderReplaceResponseAccessCentrifyConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessCentrifyConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessCentrifyConfig]
type accessIdentityProviderReplaceResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessCentrifyType string

const (
	AccessIdentityProviderReplaceResponseAccessCentrifyTypeOnetimepin AccessIdentityProviderReplaceResponseAccessCentrifyType = "onetimepin"
	AccessIdentityProviderReplaceResponseAccessCentrifyTypeAzureAd    AccessIdentityProviderReplaceResponseAccessCentrifyType = "azureAD"
	AccessIdentityProviderReplaceResponseAccessCentrifyTypeSaml       AccessIdentityProviderReplaceResponseAccessCentrifyType = "saml"
	AccessIdentityProviderReplaceResponseAccessCentrifyTypeCentrify   AccessIdentityProviderReplaceResponseAccessCentrifyType = "centrify"
	AccessIdentityProviderReplaceResponseAccessCentrifyTypeFacebook   AccessIdentityProviderReplaceResponseAccessCentrifyType = "facebook"
	AccessIdentityProviderReplaceResponseAccessCentrifyTypeGitHub     AccessIdentityProviderReplaceResponseAccessCentrifyType = "github"
	AccessIdentityProviderReplaceResponseAccessCentrifyTypeGoogleApps AccessIdentityProviderReplaceResponseAccessCentrifyType = "google-apps"
	AccessIdentityProviderReplaceResponseAccessCentrifyTypeGoogle     AccessIdentityProviderReplaceResponseAccessCentrifyType = "google"
	AccessIdentityProviderReplaceResponseAccessCentrifyTypeLinkedin   AccessIdentityProviderReplaceResponseAccessCentrifyType = "linkedin"
	AccessIdentityProviderReplaceResponseAccessCentrifyTypeOidc       AccessIdentityProviderReplaceResponseAccessCentrifyType = "oidc"
	AccessIdentityProviderReplaceResponseAccessCentrifyTypeOkta       AccessIdentityProviderReplaceResponseAccessCentrifyType = "okta"
	AccessIdentityProviderReplaceResponseAccessCentrifyTypeOnelogin   AccessIdentityProviderReplaceResponseAccessCentrifyType = "onelogin"
	AccessIdentityProviderReplaceResponseAccessCentrifyTypePingone    AccessIdentityProviderReplaceResponseAccessCentrifyType = "pingone"
	AccessIdentityProviderReplaceResponseAccessCentrifyTypeYandex     AccessIdentityProviderReplaceResponseAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceResponseAccessCentrifyScimConfig struct {
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
	JSON            accessIdentityProviderReplaceResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessCentrifyScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessCentrifyScimConfig]
type accessIdentityProviderReplaceResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderReplaceResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderReplaceResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderReplaceResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderReplaceResponseAccessFacebookJSON       `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessFacebookJSON contains the JSON
// metadata for the struct [AccessIdentityProviderReplaceResponseAccessFacebook]
type accessIdentityProviderReplaceResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderReplaceResponseAccessFacebook) implementsAccessIdentityProviderReplaceResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                        `json:"client_secret"`
	JSON         accessIdentityProviderReplaceResponseAccessFacebookConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessFacebookConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessFacebookConfig]
type accessIdentityProviderReplaceResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessFacebookType string

const (
	AccessIdentityProviderReplaceResponseAccessFacebookTypeOnetimepin AccessIdentityProviderReplaceResponseAccessFacebookType = "onetimepin"
	AccessIdentityProviderReplaceResponseAccessFacebookTypeAzureAd    AccessIdentityProviderReplaceResponseAccessFacebookType = "azureAD"
	AccessIdentityProviderReplaceResponseAccessFacebookTypeSaml       AccessIdentityProviderReplaceResponseAccessFacebookType = "saml"
	AccessIdentityProviderReplaceResponseAccessFacebookTypeCentrify   AccessIdentityProviderReplaceResponseAccessFacebookType = "centrify"
	AccessIdentityProviderReplaceResponseAccessFacebookTypeFacebook   AccessIdentityProviderReplaceResponseAccessFacebookType = "facebook"
	AccessIdentityProviderReplaceResponseAccessFacebookTypeGitHub     AccessIdentityProviderReplaceResponseAccessFacebookType = "github"
	AccessIdentityProviderReplaceResponseAccessFacebookTypeGoogleApps AccessIdentityProviderReplaceResponseAccessFacebookType = "google-apps"
	AccessIdentityProviderReplaceResponseAccessFacebookTypeGoogle     AccessIdentityProviderReplaceResponseAccessFacebookType = "google"
	AccessIdentityProviderReplaceResponseAccessFacebookTypeLinkedin   AccessIdentityProviderReplaceResponseAccessFacebookType = "linkedin"
	AccessIdentityProviderReplaceResponseAccessFacebookTypeOidc       AccessIdentityProviderReplaceResponseAccessFacebookType = "oidc"
	AccessIdentityProviderReplaceResponseAccessFacebookTypeOkta       AccessIdentityProviderReplaceResponseAccessFacebookType = "okta"
	AccessIdentityProviderReplaceResponseAccessFacebookTypeOnelogin   AccessIdentityProviderReplaceResponseAccessFacebookType = "onelogin"
	AccessIdentityProviderReplaceResponseAccessFacebookTypePingone    AccessIdentityProviderReplaceResponseAccessFacebookType = "pingone"
	AccessIdentityProviderReplaceResponseAccessFacebookTypeYandex     AccessIdentityProviderReplaceResponseAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceResponseAccessFacebookScimConfig struct {
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
	JSON            accessIdentityProviderReplaceResponseAccessFacebookScimConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessFacebookScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessFacebookScimConfig]
type accessIdentityProviderReplaceResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderReplaceResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderReplaceResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderReplaceResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderReplaceResponseAccessGitHubJSON       `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessGitHubJSON contains the JSON metadata
// for the struct [AccessIdentityProviderReplaceResponseAccessGitHub]
type accessIdentityProviderReplaceResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderReplaceResponseAccessGitHub) implementsAccessIdentityProviderReplaceResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                      `json:"client_secret"`
	JSON         accessIdentityProviderReplaceResponseAccessGitHubConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessGitHubConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessGitHubConfig]
type accessIdentityProviderReplaceResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessGitHubType string

const (
	AccessIdentityProviderReplaceResponseAccessGitHubTypeOnetimepin AccessIdentityProviderReplaceResponseAccessGitHubType = "onetimepin"
	AccessIdentityProviderReplaceResponseAccessGitHubTypeAzureAd    AccessIdentityProviderReplaceResponseAccessGitHubType = "azureAD"
	AccessIdentityProviderReplaceResponseAccessGitHubTypeSaml       AccessIdentityProviderReplaceResponseAccessGitHubType = "saml"
	AccessIdentityProviderReplaceResponseAccessGitHubTypeCentrify   AccessIdentityProviderReplaceResponseAccessGitHubType = "centrify"
	AccessIdentityProviderReplaceResponseAccessGitHubTypeFacebook   AccessIdentityProviderReplaceResponseAccessGitHubType = "facebook"
	AccessIdentityProviderReplaceResponseAccessGitHubTypeGitHub     AccessIdentityProviderReplaceResponseAccessGitHubType = "github"
	AccessIdentityProviderReplaceResponseAccessGitHubTypeGoogleApps AccessIdentityProviderReplaceResponseAccessGitHubType = "google-apps"
	AccessIdentityProviderReplaceResponseAccessGitHubTypeGoogle     AccessIdentityProviderReplaceResponseAccessGitHubType = "google"
	AccessIdentityProviderReplaceResponseAccessGitHubTypeLinkedin   AccessIdentityProviderReplaceResponseAccessGitHubType = "linkedin"
	AccessIdentityProviderReplaceResponseAccessGitHubTypeOidc       AccessIdentityProviderReplaceResponseAccessGitHubType = "oidc"
	AccessIdentityProviderReplaceResponseAccessGitHubTypeOkta       AccessIdentityProviderReplaceResponseAccessGitHubType = "okta"
	AccessIdentityProviderReplaceResponseAccessGitHubTypeOnelogin   AccessIdentityProviderReplaceResponseAccessGitHubType = "onelogin"
	AccessIdentityProviderReplaceResponseAccessGitHubTypePingone    AccessIdentityProviderReplaceResponseAccessGitHubType = "pingone"
	AccessIdentityProviderReplaceResponseAccessGitHubTypeYandex     AccessIdentityProviderReplaceResponseAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceResponseAccessGitHubScimConfig struct {
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
	JSON            accessIdentityProviderReplaceResponseAccessGitHubScimConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessGitHubScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessGitHubScimConfig]
type accessIdentityProviderReplaceResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderReplaceResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderReplaceResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderReplaceResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderReplaceResponseAccessGoogleJSON       `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessGoogleJSON contains the JSON metadata
// for the struct [AccessIdentityProviderReplaceResponseAccessGoogle]
type accessIdentityProviderReplaceResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderReplaceResponseAccessGoogle) implementsAccessIdentityProviderReplaceResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                      `json:"email_claim_name"`
	JSON           accessIdentityProviderReplaceResponseAccessGoogleConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessGoogleConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessGoogleConfig]
type accessIdentityProviderReplaceResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessGoogleType string

const (
	AccessIdentityProviderReplaceResponseAccessGoogleTypeOnetimepin AccessIdentityProviderReplaceResponseAccessGoogleType = "onetimepin"
	AccessIdentityProviderReplaceResponseAccessGoogleTypeAzureAd    AccessIdentityProviderReplaceResponseAccessGoogleType = "azureAD"
	AccessIdentityProviderReplaceResponseAccessGoogleTypeSaml       AccessIdentityProviderReplaceResponseAccessGoogleType = "saml"
	AccessIdentityProviderReplaceResponseAccessGoogleTypeCentrify   AccessIdentityProviderReplaceResponseAccessGoogleType = "centrify"
	AccessIdentityProviderReplaceResponseAccessGoogleTypeFacebook   AccessIdentityProviderReplaceResponseAccessGoogleType = "facebook"
	AccessIdentityProviderReplaceResponseAccessGoogleTypeGitHub     AccessIdentityProviderReplaceResponseAccessGoogleType = "github"
	AccessIdentityProviderReplaceResponseAccessGoogleTypeGoogleApps AccessIdentityProviderReplaceResponseAccessGoogleType = "google-apps"
	AccessIdentityProviderReplaceResponseAccessGoogleTypeGoogle     AccessIdentityProviderReplaceResponseAccessGoogleType = "google"
	AccessIdentityProviderReplaceResponseAccessGoogleTypeLinkedin   AccessIdentityProviderReplaceResponseAccessGoogleType = "linkedin"
	AccessIdentityProviderReplaceResponseAccessGoogleTypeOidc       AccessIdentityProviderReplaceResponseAccessGoogleType = "oidc"
	AccessIdentityProviderReplaceResponseAccessGoogleTypeOkta       AccessIdentityProviderReplaceResponseAccessGoogleType = "okta"
	AccessIdentityProviderReplaceResponseAccessGoogleTypeOnelogin   AccessIdentityProviderReplaceResponseAccessGoogleType = "onelogin"
	AccessIdentityProviderReplaceResponseAccessGoogleTypePingone    AccessIdentityProviderReplaceResponseAccessGoogleType = "pingone"
	AccessIdentityProviderReplaceResponseAccessGoogleTypeYandex     AccessIdentityProviderReplaceResponseAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceResponseAccessGoogleScimConfig struct {
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
	JSON            accessIdentityProviderReplaceResponseAccessGoogleScimConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessGoogleScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessGoogleScimConfig]
type accessIdentityProviderReplaceResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderReplaceResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderReplaceResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderReplaceResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderReplaceResponseAccessGoogleAppsJSON       `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessGoogleAppsJSON contains the JSON
// metadata for the struct [AccessIdentityProviderReplaceResponseAccessGoogleApps]
type accessIdentityProviderReplaceResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderReplaceResponseAccessGoogleApps) implementsAccessIdentityProviderReplaceResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                          `json:"email_claim_name"`
	JSON           accessIdentityProviderReplaceResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessGoogleAppsConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessGoogleAppsConfig]
type accessIdentityProviderReplaceResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessGoogleAppsType string

const (
	AccessIdentityProviderReplaceResponseAccessGoogleAppsTypeOnetimepin AccessIdentityProviderReplaceResponseAccessGoogleAppsType = "onetimepin"
	AccessIdentityProviderReplaceResponseAccessGoogleAppsTypeAzureAd    AccessIdentityProviderReplaceResponseAccessGoogleAppsType = "azureAD"
	AccessIdentityProviderReplaceResponseAccessGoogleAppsTypeSaml       AccessIdentityProviderReplaceResponseAccessGoogleAppsType = "saml"
	AccessIdentityProviderReplaceResponseAccessGoogleAppsTypeCentrify   AccessIdentityProviderReplaceResponseAccessGoogleAppsType = "centrify"
	AccessIdentityProviderReplaceResponseAccessGoogleAppsTypeFacebook   AccessIdentityProviderReplaceResponseAccessGoogleAppsType = "facebook"
	AccessIdentityProviderReplaceResponseAccessGoogleAppsTypeGitHub     AccessIdentityProviderReplaceResponseAccessGoogleAppsType = "github"
	AccessIdentityProviderReplaceResponseAccessGoogleAppsTypeGoogleApps AccessIdentityProviderReplaceResponseAccessGoogleAppsType = "google-apps"
	AccessIdentityProviderReplaceResponseAccessGoogleAppsTypeGoogle     AccessIdentityProviderReplaceResponseAccessGoogleAppsType = "google"
	AccessIdentityProviderReplaceResponseAccessGoogleAppsTypeLinkedin   AccessIdentityProviderReplaceResponseAccessGoogleAppsType = "linkedin"
	AccessIdentityProviderReplaceResponseAccessGoogleAppsTypeOidc       AccessIdentityProviderReplaceResponseAccessGoogleAppsType = "oidc"
	AccessIdentityProviderReplaceResponseAccessGoogleAppsTypeOkta       AccessIdentityProviderReplaceResponseAccessGoogleAppsType = "okta"
	AccessIdentityProviderReplaceResponseAccessGoogleAppsTypeOnelogin   AccessIdentityProviderReplaceResponseAccessGoogleAppsType = "onelogin"
	AccessIdentityProviderReplaceResponseAccessGoogleAppsTypePingone    AccessIdentityProviderReplaceResponseAccessGoogleAppsType = "pingone"
	AccessIdentityProviderReplaceResponseAccessGoogleAppsTypeYandex     AccessIdentityProviderReplaceResponseAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceResponseAccessGoogleAppsScimConfig struct {
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
	JSON            accessIdentityProviderReplaceResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessGoogleAppsScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessGoogleAppsScimConfig]
type accessIdentityProviderReplaceResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderReplaceResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderReplaceResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderReplaceResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderReplaceResponseAccessLinkedinJSON       `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessLinkedinJSON contains the JSON
// metadata for the struct [AccessIdentityProviderReplaceResponseAccessLinkedin]
type accessIdentityProviderReplaceResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderReplaceResponseAccessLinkedin) implementsAccessIdentityProviderReplaceResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                        `json:"client_secret"`
	JSON         accessIdentityProviderReplaceResponseAccessLinkedinConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessLinkedinConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessLinkedinConfig]
type accessIdentityProviderReplaceResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessLinkedinType string

const (
	AccessIdentityProviderReplaceResponseAccessLinkedinTypeOnetimepin AccessIdentityProviderReplaceResponseAccessLinkedinType = "onetimepin"
	AccessIdentityProviderReplaceResponseAccessLinkedinTypeAzureAd    AccessIdentityProviderReplaceResponseAccessLinkedinType = "azureAD"
	AccessIdentityProviderReplaceResponseAccessLinkedinTypeSaml       AccessIdentityProviderReplaceResponseAccessLinkedinType = "saml"
	AccessIdentityProviderReplaceResponseAccessLinkedinTypeCentrify   AccessIdentityProviderReplaceResponseAccessLinkedinType = "centrify"
	AccessIdentityProviderReplaceResponseAccessLinkedinTypeFacebook   AccessIdentityProviderReplaceResponseAccessLinkedinType = "facebook"
	AccessIdentityProviderReplaceResponseAccessLinkedinTypeGitHub     AccessIdentityProviderReplaceResponseAccessLinkedinType = "github"
	AccessIdentityProviderReplaceResponseAccessLinkedinTypeGoogleApps AccessIdentityProviderReplaceResponseAccessLinkedinType = "google-apps"
	AccessIdentityProviderReplaceResponseAccessLinkedinTypeGoogle     AccessIdentityProviderReplaceResponseAccessLinkedinType = "google"
	AccessIdentityProviderReplaceResponseAccessLinkedinTypeLinkedin   AccessIdentityProviderReplaceResponseAccessLinkedinType = "linkedin"
	AccessIdentityProviderReplaceResponseAccessLinkedinTypeOidc       AccessIdentityProviderReplaceResponseAccessLinkedinType = "oidc"
	AccessIdentityProviderReplaceResponseAccessLinkedinTypeOkta       AccessIdentityProviderReplaceResponseAccessLinkedinType = "okta"
	AccessIdentityProviderReplaceResponseAccessLinkedinTypeOnelogin   AccessIdentityProviderReplaceResponseAccessLinkedinType = "onelogin"
	AccessIdentityProviderReplaceResponseAccessLinkedinTypePingone    AccessIdentityProviderReplaceResponseAccessLinkedinType = "pingone"
	AccessIdentityProviderReplaceResponseAccessLinkedinTypeYandex     AccessIdentityProviderReplaceResponseAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceResponseAccessLinkedinScimConfig struct {
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
	JSON            accessIdentityProviderReplaceResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessLinkedinScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessLinkedinScimConfig]
type accessIdentityProviderReplaceResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderReplaceResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderReplaceResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderReplaceResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderReplaceResponseAccessOidcJSON       `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessOidcJSON contains the JSON metadata
// for the struct [AccessIdentityProviderReplaceResponseAccessOidc]
type accessIdentityProviderReplaceResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderReplaceResponseAccessOidc) implementsAccessIdentityProviderReplaceResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessOidcConfig struct {
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
	TokenURL string                                                    `json:"token_url"`
	JSON     accessIdentityProviderReplaceResponseAccessOidcConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessOidcConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderReplaceResponseAccessOidcConfig]
type accessIdentityProviderReplaceResponseAccessOidcConfigJSON struct {
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

func (r *AccessIdentityProviderReplaceResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessOidcType string

const (
	AccessIdentityProviderReplaceResponseAccessOidcTypeOnetimepin AccessIdentityProviderReplaceResponseAccessOidcType = "onetimepin"
	AccessIdentityProviderReplaceResponseAccessOidcTypeAzureAd    AccessIdentityProviderReplaceResponseAccessOidcType = "azureAD"
	AccessIdentityProviderReplaceResponseAccessOidcTypeSaml       AccessIdentityProviderReplaceResponseAccessOidcType = "saml"
	AccessIdentityProviderReplaceResponseAccessOidcTypeCentrify   AccessIdentityProviderReplaceResponseAccessOidcType = "centrify"
	AccessIdentityProviderReplaceResponseAccessOidcTypeFacebook   AccessIdentityProviderReplaceResponseAccessOidcType = "facebook"
	AccessIdentityProviderReplaceResponseAccessOidcTypeGitHub     AccessIdentityProviderReplaceResponseAccessOidcType = "github"
	AccessIdentityProviderReplaceResponseAccessOidcTypeGoogleApps AccessIdentityProviderReplaceResponseAccessOidcType = "google-apps"
	AccessIdentityProviderReplaceResponseAccessOidcTypeGoogle     AccessIdentityProviderReplaceResponseAccessOidcType = "google"
	AccessIdentityProviderReplaceResponseAccessOidcTypeLinkedin   AccessIdentityProviderReplaceResponseAccessOidcType = "linkedin"
	AccessIdentityProviderReplaceResponseAccessOidcTypeOidc       AccessIdentityProviderReplaceResponseAccessOidcType = "oidc"
	AccessIdentityProviderReplaceResponseAccessOidcTypeOkta       AccessIdentityProviderReplaceResponseAccessOidcType = "okta"
	AccessIdentityProviderReplaceResponseAccessOidcTypeOnelogin   AccessIdentityProviderReplaceResponseAccessOidcType = "onelogin"
	AccessIdentityProviderReplaceResponseAccessOidcTypePingone    AccessIdentityProviderReplaceResponseAccessOidcType = "pingone"
	AccessIdentityProviderReplaceResponseAccessOidcTypeYandex     AccessIdentityProviderReplaceResponseAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceResponseAccessOidcScimConfig struct {
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
	JSON            accessIdentityProviderReplaceResponseAccessOidcScimConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessOidcScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessOidcScimConfig]
type accessIdentityProviderReplaceResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderReplaceResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderReplaceResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderReplaceResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderReplaceResponseAccessOktaJSON       `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessOktaJSON contains the JSON metadata
// for the struct [AccessIdentityProviderReplaceResponseAccessOkta]
type accessIdentityProviderReplaceResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderReplaceResponseAccessOkta) implementsAccessIdentityProviderReplaceResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessOktaConfig struct {
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
	OktaAccount string                                                    `json:"okta_account"`
	JSON        accessIdentityProviderReplaceResponseAccessOktaConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessOktaConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderReplaceResponseAccessOktaConfig]
type accessIdentityProviderReplaceResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessOktaType string

const (
	AccessIdentityProviderReplaceResponseAccessOktaTypeOnetimepin AccessIdentityProviderReplaceResponseAccessOktaType = "onetimepin"
	AccessIdentityProviderReplaceResponseAccessOktaTypeAzureAd    AccessIdentityProviderReplaceResponseAccessOktaType = "azureAD"
	AccessIdentityProviderReplaceResponseAccessOktaTypeSaml       AccessIdentityProviderReplaceResponseAccessOktaType = "saml"
	AccessIdentityProviderReplaceResponseAccessOktaTypeCentrify   AccessIdentityProviderReplaceResponseAccessOktaType = "centrify"
	AccessIdentityProviderReplaceResponseAccessOktaTypeFacebook   AccessIdentityProviderReplaceResponseAccessOktaType = "facebook"
	AccessIdentityProviderReplaceResponseAccessOktaTypeGitHub     AccessIdentityProviderReplaceResponseAccessOktaType = "github"
	AccessIdentityProviderReplaceResponseAccessOktaTypeGoogleApps AccessIdentityProviderReplaceResponseAccessOktaType = "google-apps"
	AccessIdentityProviderReplaceResponseAccessOktaTypeGoogle     AccessIdentityProviderReplaceResponseAccessOktaType = "google"
	AccessIdentityProviderReplaceResponseAccessOktaTypeLinkedin   AccessIdentityProviderReplaceResponseAccessOktaType = "linkedin"
	AccessIdentityProviderReplaceResponseAccessOktaTypeOidc       AccessIdentityProviderReplaceResponseAccessOktaType = "oidc"
	AccessIdentityProviderReplaceResponseAccessOktaTypeOkta       AccessIdentityProviderReplaceResponseAccessOktaType = "okta"
	AccessIdentityProviderReplaceResponseAccessOktaTypeOnelogin   AccessIdentityProviderReplaceResponseAccessOktaType = "onelogin"
	AccessIdentityProviderReplaceResponseAccessOktaTypePingone    AccessIdentityProviderReplaceResponseAccessOktaType = "pingone"
	AccessIdentityProviderReplaceResponseAccessOktaTypeYandex     AccessIdentityProviderReplaceResponseAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceResponseAccessOktaScimConfig struct {
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
	JSON            accessIdentityProviderReplaceResponseAccessOktaScimConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessOktaScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessOktaScimConfig]
type accessIdentityProviderReplaceResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderReplaceResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderReplaceResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderReplaceResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderReplaceResponseAccessOneloginJSON       `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessOneloginJSON contains the JSON
// metadata for the struct [AccessIdentityProviderReplaceResponseAccessOnelogin]
type accessIdentityProviderReplaceResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderReplaceResponseAccessOnelogin) implementsAccessIdentityProviderReplaceResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                        `json:"onelogin_account"`
	JSON            accessIdentityProviderReplaceResponseAccessOneloginConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessOneloginConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessOneloginConfig]
type accessIdentityProviderReplaceResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessOneloginType string

const (
	AccessIdentityProviderReplaceResponseAccessOneloginTypeOnetimepin AccessIdentityProviderReplaceResponseAccessOneloginType = "onetimepin"
	AccessIdentityProviderReplaceResponseAccessOneloginTypeAzureAd    AccessIdentityProviderReplaceResponseAccessOneloginType = "azureAD"
	AccessIdentityProviderReplaceResponseAccessOneloginTypeSaml       AccessIdentityProviderReplaceResponseAccessOneloginType = "saml"
	AccessIdentityProviderReplaceResponseAccessOneloginTypeCentrify   AccessIdentityProviderReplaceResponseAccessOneloginType = "centrify"
	AccessIdentityProviderReplaceResponseAccessOneloginTypeFacebook   AccessIdentityProviderReplaceResponseAccessOneloginType = "facebook"
	AccessIdentityProviderReplaceResponseAccessOneloginTypeGitHub     AccessIdentityProviderReplaceResponseAccessOneloginType = "github"
	AccessIdentityProviderReplaceResponseAccessOneloginTypeGoogleApps AccessIdentityProviderReplaceResponseAccessOneloginType = "google-apps"
	AccessIdentityProviderReplaceResponseAccessOneloginTypeGoogle     AccessIdentityProviderReplaceResponseAccessOneloginType = "google"
	AccessIdentityProviderReplaceResponseAccessOneloginTypeLinkedin   AccessIdentityProviderReplaceResponseAccessOneloginType = "linkedin"
	AccessIdentityProviderReplaceResponseAccessOneloginTypeOidc       AccessIdentityProviderReplaceResponseAccessOneloginType = "oidc"
	AccessIdentityProviderReplaceResponseAccessOneloginTypeOkta       AccessIdentityProviderReplaceResponseAccessOneloginType = "okta"
	AccessIdentityProviderReplaceResponseAccessOneloginTypeOnelogin   AccessIdentityProviderReplaceResponseAccessOneloginType = "onelogin"
	AccessIdentityProviderReplaceResponseAccessOneloginTypePingone    AccessIdentityProviderReplaceResponseAccessOneloginType = "pingone"
	AccessIdentityProviderReplaceResponseAccessOneloginTypeYandex     AccessIdentityProviderReplaceResponseAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceResponseAccessOneloginScimConfig struct {
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
	JSON            accessIdentityProviderReplaceResponseAccessOneloginScimConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessOneloginScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessOneloginScimConfig]
type accessIdentityProviderReplaceResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderReplaceResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderReplaceResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderReplaceResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderReplaceResponseAccessPingoneJSON       `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessPingoneJSON contains the JSON
// metadata for the struct [AccessIdentityProviderReplaceResponseAccessPingone]
type accessIdentityProviderReplaceResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderReplaceResponseAccessPingone) implementsAccessIdentityProviderReplaceResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                       `json:"ping_env_id"`
	JSON      accessIdentityProviderReplaceResponseAccessPingoneConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessPingoneConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessPingoneConfig]
type accessIdentityProviderReplaceResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessPingoneType string

const (
	AccessIdentityProviderReplaceResponseAccessPingoneTypeOnetimepin AccessIdentityProviderReplaceResponseAccessPingoneType = "onetimepin"
	AccessIdentityProviderReplaceResponseAccessPingoneTypeAzureAd    AccessIdentityProviderReplaceResponseAccessPingoneType = "azureAD"
	AccessIdentityProviderReplaceResponseAccessPingoneTypeSaml       AccessIdentityProviderReplaceResponseAccessPingoneType = "saml"
	AccessIdentityProviderReplaceResponseAccessPingoneTypeCentrify   AccessIdentityProviderReplaceResponseAccessPingoneType = "centrify"
	AccessIdentityProviderReplaceResponseAccessPingoneTypeFacebook   AccessIdentityProviderReplaceResponseAccessPingoneType = "facebook"
	AccessIdentityProviderReplaceResponseAccessPingoneTypeGitHub     AccessIdentityProviderReplaceResponseAccessPingoneType = "github"
	AccessIdentityProviderReplaceResponseAccessPingoneTypeGoogleApps AccessIdentityProviderReplaceResponseAccessPingoneType = "google-apps"
	AccessIdentityProviderReplaceResponseAccessPingoneTypeGoogle     AccessIdentityProviderReplaceResponseAccessPingoneType = "google"
	AccessIdentityProviderReplaceResponseAccessPingoneTypeLinkedin   AccessIdentityProviderReplaceResponseAccessPingoneType = "linkedin"
	AccessIdentityProviderReplaceResponseAccessPingoneTypeOidc       AccessIdentityProviderReplaceResponseAccessPingoneType = "oidc"
	AccessIdentityProviderReplaceResponseAccessPingoneTypeOkta       AccessIdentityProviderReplaceResponseAccessPingoneType = "okta"
	AccessIdentityProviderReplaceResponseAccessPingoneTypeOnelogin   AccessIdentityProviderReplaceResponseAccessPingoneType = "onelogin"
	AccessIdentityProviderReplaceResponseAccessPingoneTypePingone    AccessIdentityProviderReplaceResponseAccessPingoneType = "pingone"
	AccessIdentityProviderReplaceResponseAccessPingoneTypeYandex     AccessIdentityProviderReplaceResponseAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceResponseAccessPingoneScimConfig struct {
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
	JSON            accessIdentityProviderReplaceResponseAccessPingoneScimConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessPingoneScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessPingoneScimConfig]
type accessIdentityProviderReplaceResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderReplaceResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderReplaceResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderReplaceResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderReplaceResponseAccessSamlJSON       `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessSamlJSON contains the JSON metadata
// for the struct [AccessIdentityProviderReplaceResponseAccessSaml]
type accessIdentityProviderReplaceResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderReplaceResponseAccessSaml) implementsAccessIdentityProviderReplaceResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccessIdentityProviderReplaceResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                    `json:"sso_target_url"`
	JSON         accessIdentityProviderReplaceResponseAccessSamlConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessSamlConfigJSON contains the JSON
// metadata for the struct [AccessIdentityProviderReplaceResponseAccessSamlConfig]
type accessIdentityProviderReplaceResponseAccessSamlConfigJSON struct {
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

func (r *AccessIdentityProviderReplaceResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                   `json:"header_name"`
	JSON       accessIdentityProviderReplaceResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessSamlConfigHeaderAttribute]
type accessIdentityProviderReplaceResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessSamlType string

const (
	AccessIdentityProviderReplaceResponseAccessSamlTypeOnetimepin AccessIdentityProviderReplaceResponseAccessSamlType = "onetimepin"
	AccessIdentityProviderReplaceResponseAccessSamlTypeAzureAd    AccessIdentityProviderReplaceResponseAccessSamlType = "azureAD"
	AccessIdentityProviderReplaceResponseAccessSamlTypeSaml       AccessIdentityProviderReplaceResponseAccessSamlType = "saml"
	AccessIdentityProviderReplaceResponseAccessSamlTypeCentrify   AccessIdentityProviderReplaceResponseAccessSamlType = "centrify"
	AccessIdentityProviderReplaceResponseAccessSamlTypeFacebook   AccessIdentityProviderReplaceResponseAccessSamlType = "facebook"
	AccessIdentityProviderReplaceResponseAccessSamlTypeGitHub     AccessIdentityProviderReplaceResponseAccessSamlType = "github"
	AccessIdentityProviderReplaceResponseAccessSamlTypeGoogleApps AccessIdentityProviderReplaceResponseAccessSamlType = "google-apps"
	AccessIdentityProviderReplaceResponseAccessSamlTypeGoogle     AccessIdentityProviderReplaceResponseAccessSamlType = "google"
	AccessIdentityProviderReplaceResponseAccessSamlTypeLinkedin   AccessIdentityProviderReplaceResponseAccessSamlType = "linkedin"
	AccessIdentityProviderReplaceResponseAccessSamlTypeOidc       AccessIdentityProviderReplaceResponseAccessSamlType = "oidc"
	AccessIdentityProviderReplaceResponseAccessSamlTypeOkta       AccessIdentityProviderReplaceResponseAccessSamlType = "okta"
	AccessIdentityProviderReplaceResponseAccessSamlTypeOnelogin   AccessIdentityProviderReplaceResponseAccessSamlType = "onelogin"
	AccessIdentityProviderReplaceResponseAccessSamlTypePingone    AccessIdentityProviderReplaceResponseAccessSamlType = "pingone"
	AccessIdentityProviderReplaceResponseAccessSamlTypeYandex     AccessIdentityProviderReplaceResponseAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceResponseAccessSamlScimConfig struct {
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
	JSON            accessIdentityProviderReplaceResponseAccessSamlScimConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessSamlScimConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessSamlScimConfig]
type accessIdentityProviderReplaceResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderReplaceResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderReplaceResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderReplaceResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderReplaceResponseAccessYandexJSON       `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessYandexJSON contains the JSON metadata
// for the struct [AccessIdentityProviderReplaceResponseAccessYandex]
type accessIdentityProviderReplaceResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderReplaceResponseAccessYandex) implementsAccessIdentityProviderReplaceResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                      `json:"client_secret"`
	JSON         accessIdentityProviderReplaceResponseAccessYandexConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessYandexConfigJSON contains the JSON
// metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessYandexConfig]
type accessIdentityProviderReplaceResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessYandexType string

const (
	AccessIdentityProviderReplaceResponseAccessYandexTypeOnetimepin AccessIdentityProviderReplaceResponseAccessYandexType = "onetimepin"
	AccessIdentityProviderReplaceResponseAccessYandexTypeAzureAd    AccessIdentityProviderReplaceResponseAccessYandexType = "azureAD"
	AccessIdentityProviderReplaceResponseAccessYandexTypeSaml       AccessIdentityProviderReplaceResponseAccessYandexType = "saml"
	AccessIdentityProviderReplaceResponseAccessYandexTypeCentrify   AccessIdentityProviderReplaceResponseAccessYandexType = "centrify"
	AccessIdentityProviderReplaceResponseAccessYandexTypeFacebook   AccessIdentityProviderReplaceResponseAccessYandexType = "facebook"
	AccessIdentityProviderReplaceResponseAccessYandexTypeGitHub     AccessIdentityProviderReplaceResponseAccessYandexType = "github"
	AccessIdentityProviderReplaceResponseAccessYandexTypeGoogleApps AccessIdentityProviderReplaceResponseAccessYandexType = "google-apps"
	AccessIdentityProviderReplaceResponseAccessYandexTypeGoogle     AccessIdentityProviderReplaceResponseAccessYandexType = "google"
	AccessIdentityProviderReplaceResponseAccessYandexTypeLinkedin   AccessIdentityProviderReplaceResponseAccessYandexType = "linkedin"
	AccessIdentityProviderReplaceResponseAccessYandexTypeOidc       AccessIdentityProviderReplaceResponseAccessYandexType = "oidc"
	AccessIdentityProviderReplaceResponseAccessYandexTypeOkta       AccessIdentityProviderReplaceResponseAccessYandexType = "okta"
	AccessIdentityProviderReplaceResponseAccessYandexTypeOnelogin   AccessIdentityProviderReplaceResponseAccessYandexType = "onelogin"
	AccessIdentityProviderReplaceResponseAccessYandexTypePingone    AccessIdentityProviderReplaceResponseAccessYandexType = "pingone"
	AccessIdentityProviderReplaceResponseAccessYandexTypeYandex     AccessIdentityProviderReplaceResponseAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceResponseAccessYandexScimConfig struct {
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
	JSON            accessIdentityProviderReplaceResponseAccessYandexScimConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessYandexScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessYandexScimConfig]
type accessIdentityProviderReplaceResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderReplaceResponseAccessOnetimepinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderReplaceResponseAccessOnetimepinScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderReplaceResponseAccessOnetimepinJSON       `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessOnetimepinJSON contains the JSON
// metadata for the struct [AccessIdentityProviderReplaceResponseAccessOnetimepin]
type accessIdentityProviderReplaceResponseAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderReplaceResponseAccessOnetimepin) implementsAccessIdentityProviderReplaceResponse() {
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceResponseAccessOnetimepinType string

const (
	AccessIdentityProviderReplaceResponseAccessOnetimepinTypeOnetimepin AccessIdentityProviderReplaceResponseAccessOnetimepinType = "onetimepin"
	AccessIdentityProviderReplaceResponseAccessOnetimepinTypeAzureAd    AccessIdentityProviderReplaceResponseAccessOnetimepinType = "azureAD"
	AccessIdentityProviderReplaceResponseAccessOnetimepinTypeSaml       AccessIdentityProviderReplaceResponseAccessOnetimepinType = "saml"
	AccessIdentityProviderReplaceResponseAccessOnetimepinTypeCentrify   AccessIdentityProviderReplaceResponseAccessOnetimepinType = "centrify"
	AccessIdentityProviderReplaceResponseAccessOnetimepinTypeFacebook   AccessIdentityProviderReplaceResponseAccessOnetimepinType = "facebook"
	AccessIdentityProviderReplaceResponseAccessOnetimepinTypeGitHub     AccessIdentityProviderReplaceResponseAccessOnetimepinType = "github"
	AccessIdentityProviderReplaceResponseAccessOnetimepinTypeGoogleApps AccessIdentityProviderReplaceResponseAccessOnetimepinType = "google-apps"
	AccessIdentityProviderReplaceResponseAccessOnetimepinTypeGoogle     AccessIdentityProviderReplaceResponseAccessOnetimepinType = "google"
	AccessIdentityProviderReplaceResponseAccessOnetimepinTypeLinkedin   AccessIdentityProviderReplaceResponseAccessOnetimepinType = "linkedin"
	AccessIdentityProviderReplaceResponseAccessOnetimepinTypeOidc       AccessIdentityProviderReplaceResponseAccessOnetimepinType = "oidc"
	AccessIdentityProviderReplaceResponseAccessOnetimepinTypeOkta       AccessIdentityProviderReplaceResponseAccessOnetimepinType = "okta"
	AccessIdentityProviderReplaceResponseAccessOnetimepinTypeOnelogin   AccessIdentityProviderReplaceResponseAccessOnetimepinType = "onelogin"
	AccessIdentityProviderReplaceResponseAccessOnetimepinTypePingone    AccessIdentityProviderReplaceResponseAccessOnetimepinType = "pingone"
	AccessIdentityProviderReplaceResponseAccessOnetimepinTypeYandex     AccessIdentityProviderReplaceResponseAccessOnetimepinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceResponseAccessOnetimepinScimConfig struct {
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
	JSON            accessIdentityProviderReplaceResponseAccessOnetimepinScimConfigJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseAccessOnetimepinScimConfigJSON contains the
// JSON metadata for the struct
// [AccessIdentityProviderReplaceResponseAccessOnetimepinScimConfig]
type accessIdentityProviderReplaceResponseAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This interface is a union satisfied by one of the following:
// [AccessIdentityProviderNewParamsAccessAzureAd],
// [AccessIdentityProviderNewParamsAccessCentrify],
// [AccessIdentityProviderNewParamsAccessFacebook],
// [AccessIdentityProviderNewParamsAccessGitHub],
// [AccessIdentityProviderNewParamsAccessGoogle],
// [AccessIdentityProviderNewParamsAccessGoogleApps],
// [AccessIdentityProviderNewParamsAccessLinkedin],
// [AccessIdentityProviderNewParamsAccessOidc],
// [AccessIdentityProviderNewParamsAccessOkta],
// [AccessIdentityProviderNewParamsAccessOnelogin],
// [AccessIdentityProviderNewParamsAccessPingone],
// [AccessIdentityProviderNewParamsAccessSaml],
// [AccessIdentityProviderNewParamsAccessYandex],
// [AccessIdentityProviderNewParamsAccessOnetimepin].
type AccessIdentityProviderNewParams interface {
	ImplementsAccessIdentityProviderNewParams()
}

type AccessIdentityProviderNewParamsAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderNewParamsAccessAzureAdConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderNewParamsAccessAzureAdType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderNewParamsAccessAzureAdScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderNewParamsAccessAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderNewParamsAccessAzureAd) ImplementsAccessIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessAzureAdConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessAzureAdConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessAzureAdType string

const (
	AccessIdentityProviderNewParamsAccessAzureAdTypeOnetimepin AccessIdentityProviderNewParamsAccessAzureAdType = "onetimepin"
	AccessIdentityProviderNewParamsAccessAzureAdTypeAzureAd    AccessIdentityProviderNewParamsAccessAzureAdType = "azureAD"
	AccessIdentityProviderNewParamsAccessAzureAdTypeSaml       AccessIdentityProviderNewParamsAccessAzureAdType = "saml"
	AccessIdentityProviderNewParamsAccessAzureAdTypeCentrify   AccessIdentityProviderNewParamsAccessAzureAdType = "centrify"
	AccessIdentityProviderNewParamsAccessAzureAdTypeFacebook   AccessIdentityProviderNewParamsAccessAzureAdType = "facebook"
	AccessIdentityProviderNewParamsAccessAzureAdTypeGitHub     AccessIdentityProviderNewParamsAccessAzureAdType = "github"
	AccessIdentityProviderNewParamsAccessAzureAdTypeGoogleApps AccessIdentityProviderNewParamsAccessAzureAdType = "google-apps"
	AccessIdentityProviderNewParamsAccessAzureAdTypeGoogle     AccessIdentityProviderNewParamsAccessAzureAdType = "google"
	AccessIdentityProviderNewParamsAccessAzureAdTypeLinkedin   AccessIdentityProviderNewParamsAccessAzureAdType = "linkedin"
	AccessIdentityProviderNewParamsAccessAzureAdTypeOidc       AccessIdentityProviderNewParamsAccessAzureAdType = "oidc"
	AccessIdentityProviderNewParamsAccessAzureAdTypeOkta       AccessIdentityProviderNewParamsAccessAzureAdType = "okta"
	AccessIdentityProviderNewParamsAccessAzureAdTypeOnelogin   AccessIdentityProviderNewParamsAccessAzureAdType = "onelogin"
	AccessIdentityProviderNewParamsAccessAzureAdTypePingone    AccessIdentityProviderNewParamsAccessAzureAdType = "pingone"
	AccessIdentityProviderNewParamsAccessAzureAdTypeYandex     AccessIdentityProviderNewParamsAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewParamsAccessAzureAdScimConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessAzureAdScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderNewParamsAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderNewParamsAccessCentrifyConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderNewParamsAccessCentrifyType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderNewParamsAccessCentrifyScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderNewParamsAccessCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderNewParamsAccessCentrify) ImplementsAccessIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessCentrifyConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessCentrifyType string

const (
	AccessIdentityProviderNewParamsAccessCentrifyTypeOnetimepin AccessIdentityProviderNewParamsAccessCentrifyType = "onetimepin"
	AccessIdentityProviderNewParamsAccessCentrifyTypeAzureAd    AccessIdentityProviderNewParamsAccessCentrifyType = "azureAD"
	AccessIdentityProviderNewParamsAccessCentrifyTypeSaml       AccessIdentityProviderNewParamsAccessCentrifyType = "saml"
	AccessIdentityProviderNewParamsAccessCentrifyTypeCentrify   AccessIdentityProviderNewParamsAccessCentrifyType = "centrify"
	AccessIdentityProviderNewParamsAccessCentrifyTypeFacebook   AccessIdentityProviderNewParamsAccessCentrifyType = "facebook"
	AccessIdentityProviderNewParamsAccessCentrifyTypeGitHub     AccessIdentityProviderNewParamsAccessCentrifyType = "github"
	AccessIdentityProviderNewParamsAccessCentrifyTypeGoogleApps AccessIdentityProviderNewParamsAccessCentrifyType = "google-apps"
	AccessIdentityProviderNewParamsAccessCentrifyTypeGoogle     AccessIdentityProviderNewParamsAccessCentrifyType = "google"
	AccessIdentityProviderNewParamsAccessCentrifyTypeLinkedin   AccessIdentityProviderNewParamsAccessCentrifyType = "linkedin"
	AccessIdentityProviderNewParamsAccessCentrifyTypeOidc       AccessIdentityProviderNewParamsAccessCentrifyType = "oidc"
	AccessIdentityProviderNewParamsAccessCentrifyTypeOkta       AccessIdentityProviderNewParamsAccessCentrifyType = "okta"
	AccessIdentityProviderNewParamsAccessCentrifyTypeOnelogin   AccessIdentityProviderNewParamsAccessCentrifyType = "onelogin"
	AccessIdentityProviderNewParamsAccessCentrifyTypePingone    AccessIdentityProviderNewParamsAccessCentrifyType = "pingone"
	AccessIdentityProviderNewParamsAccessCentrifyTypeYandex     AccessIdentityProviderNewParamsAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewParamsAccessCentrifyScimConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessCentrifyScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderNewParamsAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderNewParamsAccessFacebookConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderNewParamsAccessFacebookType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderNewParamsAccessFacebookScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderNewParamsAccessFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderNewParamsAccessFacebook) ImplementsAccessIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderNewParamsAccessFacebookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessFacebookType string

const (
	AccessIdentityProviderNewParamsAccessFacebookTypeOnetimepin AccessIdentityProviderNewParamsAccessFacebookType = "onetimepin"
	AccessIdentityProviderNewParamsAccessFacebookTypeAzureAd    AccessIdentityProviderNewParamsAccessFacebookType = "azureAD"
	AccessIdentityProviderNewParamsAccessFacebookTypeSaml       AccessIdentityProviderNewParamsAccessFacebookType = "saml"
	AccessIdentityProviderNewParamsAccessFacebookTypeCentrify   AccessIdentityProviderNewParamsAccessFacebookType = "centrify"
	AccessIdentityProviderNewParamsAccessFacebookTypeFacebook   AccessIdentityProviderNewParamsAccessFacebookType = "facebook"
	AccessIdentityProviderNewParamsAccessFacebookTypeGitHub     AccessIdentityProviderNewParamsAccessFacebookType = "github"
	AccessIdentityProviderNewParamsAccessFacebookTypeGoogleApps AccessIdentityProviderNewParamsAccessFacebookType = "google-apps"
	AccessIdentityProviderNewParamsAccessFacebookTypeGoogle     AccessIdentityProviderNewParamsAccessFacebookType = "google"
	AccessIdentityProviderNewParamsAccessFacebookTypeLinkedin   AccessIdentityProviderNewParamsAccessFacebookType = "linkedin"
	AccessIdentityProviderNewParamsAccessFacebookTypeOidc       AccessIdentityProviderNewParamsAccessFacebookType = "oidc"
	AccessIdentityProviderNewParamsAccessFacebookTypeOkta       AccessIdentityProviderNewParamsAccessFacebookType = "okta"
	AccessIdentityProviderNewParamsAccessFacebookTypeOnelogin   AccessIdentityProviderNewParamsAccessFacebookType = "onelogin"
	AccessIdentityProviderNewParamsAccessFacebookTypePingone    AccessIdentityProviderNewParamsAccessFacebookType = "pingone"
	AccessIdentityProviderNewParamsAccessFacebookTypeYandex     AccessIdentityProviderNewParamsAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewParamsAccessFacebookScimConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessFacebookScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderNewParamsAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderNewParamsAccessGitHubConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderNewParamsAccessGitHubType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderNewParamsAccessGitHubScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderNewParamsAccessGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderNewParamsAccessGitHub) ImplementsAccessIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderNewParamsAccessGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessGitHubType string

const (
	AccessIdentityProviderNewParamsAccessGitHubTypeOnetimepin AccessIdentityProviderNewParamsAccessGitHubType = "onetimepin"
	AccessIdentityProviderNewParamsAccessGitHubTypeAzureAd    AccessIdentityProviderNewParamsAccessGitHubType = "azureAD"
	AccessIdentityProviderNewParamsAccessGitHubTypeSaml       AccessIdentityProviderNewParamsAccessGitHubType = "saml"
	AccessIdentityProviderNewParamsAccessGitHubTypeCentrify   AccessIdentityProviderNewParamsAccessGitHubType = "centrify"
	AccessIdentityProviderNewParamsAccessGitHubTypeFacebook   AccessIdentityProviderNewParamsAccessGitHubType = "facebook"
	AccessIdentityProviderNewParamsAccessGitHubTypeGitHub     AccessIdentityProviderNewParamsAccessGitHubType = "github"
	AccessIdentityProviderNewParamsAccessGitHubTypeGoogleApps AccessIdentityProviderNewParamsAccessGitHubType = "google-apps"
	AccessIdentityProviderNewParamsAccessGitHubTypeGoogle     AccessIdentityProviderNewParamsAccessGitHubType = "google"
	AccessIdentityProviderNewParamsAccessGitHubTypeLinkedin   AccessIdentityProviderNewParamsAccessGitHubType = "linkedin"
	AccessIdentityProviderNewParamsAccessGitHubTypeOidc       AccessIdentityProviderNewParamsAccessGitHubType = "oidc"
	AccessIdentityProviderNewParamsAccessGitHubTypeOkta       AccessIdentityProviderNewParamsAccessGitHubType = "okta"
	AccessIdentityProviderNewParamsAccessGitHubTypeOnelogin   AccessIdentityProviderNewParamsAccessGitHubType = "onelogin"
	AccessIdentityProviderNewParamsAccessGitHubTypePingone    AccessIdentityProviderNewParamsAccessGitHubType = "pingone"
	AccessIdentityProviderNewParamsAccessGitHubTypeYandex     AccessIdentityProviderNewParamsAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewParamsAccessGitHubScimConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessGitHubScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderNewParamsAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderNewParamsAccessGoogleConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderNewParamsAccessGoogleType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderNewParamsAccessGoogleScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderNewParamsAccessGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderNewParamsAccessGoogle) ImplementsAccessIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessGoogleConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r AccessIdentityProviderNewParamsAccessGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessGoogleType string

const (
	AccessIdentityProviderNewParamsAccessGoogleTypeOnetimepin AccessIdentityProviderNewParamsAccessGoogleType = "onetimepin"
	AccessIdentityProviderNewParamsAccessGoogleTypeAzureAd    AccessIdentityProviderNewParamsAccessGoogleType = "azureAD"
	AccessIdentityProviderNewParamsAccessGoogleTypeSaml       AccessIdentityProviderNewParamsAccessGoogleType = "saml"
	AccessIdentityProviderNewParamsAccessGoogleTypeCentrify   AccessIdentityProviderNewParamsAccessGoogleType = "centrify"
	AccessIdentityProviderNewParamsAccessGoogleTypeFacebook   AccessIdentityProviderNewParamsAccessGoogleType = "facebook"
	AccessIdentityProviderNewParamsAccessGoogleTypeGitHub     AccessIdentityProviderNewParamsAccessGoogleType = "github"
	AccessIdentityProviderNewParamsAccessGoogleTypeGoogleApps AccessIdentityProviderNewParamsAccessGoogleType = "google-apps"
	AccessIdentityProviderNewParamsAccessGoogleTypeGoogle     AccessIdentityProviderNewParamsAccessGoogleType = "google"
	AccessIdentityProviderNewParamsAccessGoogleTypeLinkedin   AccessIdentityProviderNewParamsAccessGoogleType = "linkedin"
	AccessIdentityProviderNewParamsAccessGoogleTypeOidc       AccessIdentityProviderNewParamsAccessGoogleType = "oidc"
	AccessIdentityProviderNewParamsAccessGoogleTypeOkta       AccessIdentityProviderNewParamsAccessGoogleType = "okta"
	AccessIdentityProviderNewParamsAccessGoogleTypeOnelogin   AccessIdentityProviderNewParamsAccessGoogleType = "onelogin"
	AccessIdentityProviderNewParamsAccessGoogleTypePingone    AccessIdentityProviderNewParamsAccessGoogleType = "pingone"
	AccessIdentityProviderNewParamsAccessGoogleTypeYandex     AccessIdentityProviderNewParamsAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewParamsAccessGoogleScimConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessGoogleScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderNewParamsAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderNewParamsAccessGoogleAppsConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderNewParamsAccessGoogleAppsType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderNewParamsAccessGoogleAppsScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderNewParamsAccessGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderNewParamsAccessGoogleApps) ImplementsAccessIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessGoogleAppsConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessGoogleAppsType string

const (
	AccessIdentityProviderNewParamsAccessGoogleAppsTypeOnetimepin AccessIdentityProviderNewParamsAccessGoogleAppsType = "onetimepin"
	AccessIdentityProviderNewParamsAccessGoogleAppsTypeAzureAd    AccessIdentityProviderNewParamsAccessGoogleAppsType = "azureAD"
	AccessIdentityProviderNewParamsAccessGoogleAppsTypeSaml       AccessIdentityProviderNewParamsAccessGoogleAppsType = "saml"
	AccessIdentityProviderNewParamsAccessGoogleAppsTypeCentrify   AccessIdentityProviderNewParamsAccessGoogleAppsType = "centrify"
	AccessIdentityProviderNewParamsAccessGoogleAppsTypeFacebook   AccessIdentityProviderNewParamsAccessGoogleAppsType = "facebook"
	AccessIdentityProviderNewParamsAccessGoogleAppsTypeGitHub     AccessIdentityProviderNewParamsAccessGoogleAppsType = "github"
	AccessIdentityProviderNewParamsAccessGoogleAppsTypeGoogleApps AccessIdentityProviderNewParamsAccessGoogleAppsType = "google-apps"
	AccessIdentityProviderNewParamsAccessGoogleAppsTypeGoogle     AccessIdentityProviderNewParamsAccessGoogleAppsType = "google"
	AccessIdentityProviderNewParamsAccessGoogleAppsTypeLinkedin   AccessIdentityProviderNewParamsAccessGoogleAppsType = "linkedin"
	AccessIdentityProviderNewParamsAccessGoogleAppsTypeOidc       AccessIdentityProviderNewParamsAccessGoogleAppsType = "oidc"
	AccessIdentityProviderNewParamsAccessGoogleAppsTypeOkta       AccessIdentityProviderNewParamsAccessGoogleAppsType = "okta"
	AccessIdentityProviderNewParamsAccessGoogleAppsTypeOnelogin   AccessIdentityProviderNewParamsAccessGoogleAppsType = "onelogin"
	AccessIdentityProviderNewParamsAccessGoogleAppsTypePingone    AccessIdentityProviderNewParamsAccessGoogleAppsType = "pingone"
	AccessIdentityProviderNewParamsAccessGoogleAppsTypeYandex     AccessIdentityProviderNewParamsAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewParamsAccessGoogleAppsScimConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessGoogleAppsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderNewParamsAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderNewParamsAccessLinkedinConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderNewParamsAccessLinkedinType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderNewParamsAccessLinkedinScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderNewParamsAccessLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderNewParamsAccessLinkedin) ImplementsAccessIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderNewParamsAccessLinkedinConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessLinkedinType string

const (
	AccessIdentityProviderNewParamsAccessLinkedinTypeOnetimepin AccessIdentityProviderNewParamsAccessLinkedinType = "onetimepin"
	AccessIdentityProviderNewParamsAccessLinkedinTypeAzureAd    AccessIdentityProviderNewParamsAccessLinkedinType = "azureAD"
	AccessIdentityProviderNewParamsAccessLinkedinTypeSaml       AccessIdentityProviderNewParamsAccessLinkedinType = "saml"
	AccessIdentityProviderNewParamsAccessLinkedinTypeCentrify   AccessIdentityProviderNewParamsAccessLinkedinType = "centrify"
	AccessIdentityProviderNewParamsAccessLinkedinTypeFacebook   AccessIdentityProviderNewParamsAccessLinkedinType = "facebook"
	AccessIdentityProviderNewParamsAccessLinkedinTypeGitHub     AccessIdentityProviderNewParamsAccessLinkedinType = "github"
	AccessIdentityProviderNewParamsAccessLinkedinTypeGoogleApps AccessIdentityProviderNewParamsAccessLinkedinType = "google-apps"
	AccessIdentityProviderNewParamsAccessLinkedinTypeGoogle     AccessIdentityProviderNewParamsAccessLinkedinType = "google"
	AccessIdentityProviderNewParamsAccessLinkedinTypeLinkedin   AccessIdentityProviderNewParamsAccessLinkedinType = "linkedin"
	AccessIdentityProviderNewParamsAccessLinkedinTypeOidc       AccessIdentityProviderNewParamsAccessLinkedinType = "oidc"
	AccessIdentityProviderNewParamsAccessLinkedinTypeOkta       AccessIdentityProviderNewParamsAccessLinkedinType = "okta"
	AccessIdentityProviderNewParamsAccessLinkedinTypeOnelogin   AccessIdentityProviderNewParamsAccessLinkedinType = "onelogin"
	AccessIdentityProviderNewParamsAccessLinkedinTypePingone    AccessIdentityProviderNewParamsAccessLinkedinType = "pingone"
	AccessIdentityProviderNewParamsAccessLinkedinTypeYandex     AccessIdentityProviderNewParamsAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewParamsAccessLinkedinScimConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessLinkedinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderNewParamsAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderNewParamsAccessOidcConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderNewParamsAccessOidcType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderNewParamsAccessOidcScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderNewParamsAccessOidc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderNewParamsAccessOidc) ImplementsAccessIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessOidcConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessOidcConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessOidcType string

const (
	AccessIdentityProviderNewParamsAccessOidcTypeOnetimepin AccessIdentityProviderNewParamsAccessOidcType = "onetimepin"
	AccessIdentityProviderNewParamsAccessOidcTypeAzureAd    AccessIdentityProviderNewParamsAccessOidcType = "azureAD"
	AccessIdentityProviderNewParamsAccessOidcTypeSaml       AccessIdentityProviderNewParamsAccessOidcType = "saml"
	AccessIdentityProviderNewParamsAccessOidcTypeCentrify   AccessIdentityProviderNewParamsAccessOidcType = "centrify"
	AccessIdentityProviderNewParamsAccessOidcTypeFacebook   AccessIdentityProviderNewParamsAccessOidcType = "facebook"
	AccessIdentityProviderNewParamsAccessOidcTypeGitHub     AccessIdentityProviderNewParamsAccessOidcType = "github"
	AccessIdentityProviderNewParamsAccessOidcTypeGoogleApps AccessIdentityProviderNewParamsAccessOidcType = "google-apps"
	AccessIdentityProviderNewParamsAccessOidcTypeGoogle     AccessIdentityProviderNewParamsAccessOidcType = "google"
	AccessIdentityProviderNewParamsAccessOidcTypeLinkedin   AccessIdentityProviderNewParamsAccessOidcType = "linkedin"
	AccessIdentityProviderNewParamsAccessOidcTypeOidc       AccessIdentityProviderNewParamsAccessOidcType = "oidc"
	AccessIdentityProviderNewParamsAccessOidcTypeOkta       AccessIdentityProviderNewParamsAccessOidcType = "okta"
	AccessIdentityProviderNewParamsAccessOidcTypeOnelogin   AccessIdentityProviderNewParamsAccessOidcType = "onelogin"
	AccessIdentityProviderNewParamsAccessOidcTypePingone    AccessIdentityProviderNewParamsAccessOidcType = "pingone"
	AccessIdentityProviderNewParamsAccessOidcTypeYandex     AccessIdentityProviderNewParamsAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewParamsAccessOidcScimConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessOidcScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderNewParamsAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderNewParamsAccessOktaConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderNewParamsAccessOktaType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderNewParamsAccessOktaScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderNewParamsAccessOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderNewParamsAccessOkta) ImplementsAccessIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessOktaConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessOktaType string

const (
	AccessIdentityProviderNewParamsAccessOktaTypeOnetimepin AccessIdentityProviderNewParamsAccessOktaType = "onetimepin"
	AccessIdentityProviderNewParamsAccessOktaTypeAzureAd    AccessIdentityProviderNewParamsAccessOktaType = "azureAD"
	AccessIdentityProviderNewParamsAccessOktaTypeSaml       AccessIdentityProviderNewParamsAccessOktaType = "saml"
	AccessIdentityProviderNewParamsAccessOktaTypeCentrify   AccessIdentityProviderNewParamsAccessOktaType = "centrify"
	AccessIdentityProviderNewParamsAccessOktaTypeFacebook   AccessIdentityProviderNewParamsAccessOktaType = "facebook"
	AccessIdentityProviderNewParamsAccessOktaTypeGitHub     AccessIdentityProviderNewParamsAccessOktaType = "github"
	AccessIdentityProviderNewParamsAccessOktaTypeGoogleApps AccessIdentityProviderNewParamsAccessOktaType = "google-apps"
	AccessIdentityProviderNewParamsAccessOktaTypeGoogle     AccessIdentityProviderNewParamsAccessOktaType = "google"
	AccessIdentityProviderNewParamsAccessOktaTypeLinkedin   AccessIdentityProviderNewParamsAccessOktaType = "linkedin"
	AccessIdentityProviderNewParamsAccessOktaTypeOidc       AccessIdentityProviderNewParamsAccessOktaType = "oidc"
	AccessIdentityProviderNewParamsAccessOktaTypeOkta       AccessIdentityProviderNewParamsAccessOktaType = "okta"
	AccessIdentityProviderNewParamsAccessOktaTypeOnelogin   AccessIdentityProviderNewParamsAccessOktaType = "onelogin"
	AccessIdentityProviderNewParamsAccessOktaTypePingone    AccessIdentityProviderNewParamsAccessOktaType = "pingone"
	AccessIdentityProviderNewParamsAccessOktaTypeYandex     AccessIdentityProviderNewParamsAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewParamsAccessOktaScimConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessOktaScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderNewParamsAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderNewParamsAccessOneloginConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderNewParamsAccessOneloginType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderNewParamsAccessOneloginScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderNewParamsAccessOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderNewParamsAccessOnelogin) ImplementsAccessIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessOneloginConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessOneloginType string

const (
	AccessIdentityProviderNewParamsAccessOneloginTypeOnetimepin AccessIdentityProviderNewParamsAccessOneloginType = "onetimepin"
	AccessIdentityProviderNewParamsAccessOneloginTypeAzureAd    AccessIdentityProviderNewParamsAccessOneloginType = "azureAD"
	AccessIdentityProviderNewParamsAccessOneloginTypeSaml       AccessIdentityProviderNewParamsAccessOneloginType = "saml"
	AccessIdentityProviderNewParamsAccessOneloginTypeCentrify   AccessIdentityProviderNewParamsAccessOneloginType = "centrify"
	AccessIdentityProviderNewParamsAccessOneloginTypeFacebook   AccessIdentityProviderNewParamsAccessOneloginType = "facebook"
	AccessIdentityProviderNewParamsAccessOneloginTypeGitHub     AccessIdentityProviderNewParamsAccessOneloginType = "github"
	AccessIdentityProviderNewParamsAccessOneloginTypeGoogleApps AccessIdentityProviderNewParamsAccessOneloginType = "google-apps"
	AccessIdentityProviderNewParamsAccessOneloginTypeGoogle     AccessIdentityProviderNewParamsAccessOneloginType = "google"
	AccessIdentityProviderNewParamsAccessOneloginTypeLinkedin   AccessIdentityProviderNewParamsAccessOneloginType = "linkedin"
	AccessIdentityProviderNewParamsAccessOneloginTypeOidc       AccessIdentityProviderNewParamsAccessOneloginType = "oidc"
	AccessIdentityProviderNewParamsAccessOneloginTypeOkta       AccessIdentityProviderNewParamsAccessOneloginType = "okta"
	AccessIdentityProviderNewParamsAccessOneloginTypeOnelogin   AccessIdentityProviderNewParamsAccessOneloginType = "onelogin"
	AccessIdentityProviderNewParamsAccessOneloginTypePingone    AccessIdentityProviderNewParamsAccessOneloginType = "pingone"
	AccessIdentityProviderNewParamsAccessOneloginTypeYandex     AccessIdentityProviderNewParamsAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewParamsAccessOneloginScimConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessOneloginScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderNewParamsAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderNewParamsAccessPingoneConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderNewParamsAccessPingoneType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderNewParamsAccessPingoneScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderNewParamsAccessPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderNewParamsAccessPingone) ImplementsAccessIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessPingoneConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessPingoneType string

const (
	AccessIdentityProviderNewParamsAccessPingoneTypeOnetimepin AccessIdentityProviderNewParamsAccessPingoneType = "onetimepin"
	AccessIdentityProviderNewParamsAccessPingoneTypeAzureAd    AccessIdentityProviderNewParamsAccessPingoneType = "azureAD"
	AccessIdentityProviderNewParamsAccessPingoneTypeSaml       AccessIdentityProviderNewParamsAccessPingoneType = "saml"
	AccessIdentityProviderNewParamsAccessPingoneTypeCentrify   AccessIdentityProviderNewParamsAccessPingoneType = "centrify"
	AccessIdentityProviderNewParamsAccessPingoneTypeFacebook   AccessIdentityProviderNewParamsAccessPingoneType = "facebook"
	AccessIdentityProviderNewParamsAccessPingoneTypeGitHub     AccessIdentityProviderNewParamsAccessPingoneType = "github"
	AccessIdentityProviderNewParamsAccessPingoneTypeGoogleApps AccessIdentityProviderNewParamsAccessPingoneType = "google-apps"
	AccessIdentityProviderNewParamsAccessPingoneTypeGoogle     AccessIdentityProviderNewParamsAccessPingoneType = "google"
	AccessIdentityProviderNewParamsAccessPingoneTypeLinkedin   AccessIdentityProviderNewParamsAccessPingoneType = "linkedin"
	AccessIdentityProviderNewParamsAccessPingoneTypeOidc       AccessIdentityProviderNewParamsAccessPingoneType = "oidc"
	AccessIdentityProviderNewParamsAccessPingoneTypeOkta       AccessIdentityProviderNewParamsAccessPingoneType = "okta"
	AccessIdentityProviderNewParamsAccessPingoneTypeOnelogin   AccessIdentityProviderNewParamsAccessPingoneType = "onelogin"
	AccessIdentityProviderNewParamsAccessPingoneTypePingone    AccessIdentityProviderNewParamsAccessPingoneType = "pingone"
	AccessIdentityProviderNewParamsAccessPingoneTypeYandex     AccessIdentityProviderNewParamsAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewParamsAccessPingoneScimConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessPingoneScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderNewParamsAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderNewParamsAccessSamlConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderNewParamsAccessSamlType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderNewParamsAccessSamlScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderNewParamsAccessSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderNewParamsAccessSaml) ImplementsAccessIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]AccessIdentityProviderNewParamsAccessSamlConfigHeaderAttribute] `json:"header_attributes"`
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

func (r AccessIdentityProviderNewParamsAccessSamlConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderNewParamsAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r AccessIdentityProviderNewParamsAccessSamlConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessSamlType string

const (
	AccessIdentityProviderNewParamsAccessSamlTypeOnetimepin AccessIdentityProviderNewParamsAccessSamlType = "onetimepin"
	AccessIdentityProviderNewParamsAccessSamlTypeAzureAd    AccessIdentityProviderNewParamsAccessSamlType = "azureAD"
	AccessIdentityProviderNewParamsAccessSamlTypeSaml       AccessIdentityProviderNewParamsAccessSamlType = "saml"
	AccessIdentityProviderNewParamsAccessSamlTypeCentrify   AccessIdentityProviderNewParamsAccessSamlType = "centrify"
	AccessIdentityProviderNewParamsAccessSamlTypeFacebook   AccessIdentityProviderNewParamsAccessSamlType = "facebook"
	AccessIdentityProviderNewParamsAccessSamlTypeGitHub     AccessIdentityProviderNewParamsAccessSamlType = "github"
	AccessIdentityProviderNewParamsAccessSamlTypeGoogleApps AccessIdentityProviderNewParamsAccessSamlType = "google-apps"
	AccessIdentityProviderNewParamsAccessSamlTypeGoogle     AccessIdentityProviderNewParamsAccessSamlType = "google"
	AccessIdentityProviderNewParamsAccessSamlTypeLinkedin   AccessIdentityProviderNewParamsAccessSamlType = "linkedin"
	AccessIdentityProviderNewParamsAccessSamlTypeOidc       AccessIdentityProviderNewParamsAccessSamlType = "oidc"
	AccessIdentityProviderNewParamsAccessSamlTypeOkta       AccessIdentityProviderNewParamsAccessSamlType = "okta"
	AccessIdentityProviderNewParamsAccessSamlTypeOnelogin   AccessIdentityProviderNewParamsAccessSamlType = "onelogin"
	AccessIdentityProviderNewParamsAccessSamlTypePingone    AccessIdentityProviderNewParamsAccessSamlType = "pingone"
	AccessIdentityProviderNewParamsAccessSamlTypeYandex     AccessIdentityProviderNewParamsAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewParamsAccessSamlScimConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessSamlScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderNewParamsAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderNewParamsAccessYandexConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderNewParamsAccessYandexType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderNewParamsAccessYandexScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderNewParamsAccessYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderNewParamsAccessYandex) ImplementsAccessIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderNewParamsAccessYandexConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessYandexType string

const (
	AccessIdentityProviderNewParamsAccessYandexTypeOnetimepin AccessIdentityProviderNewParamsAccessYandexType = "onetimepin"
	AccessIdentityProviderNewParamsAccessYandexTypeAzureAd    AccessIdentityProviderNewParamsAccessYandexType = "azureAD"
	AccessIdentityProviderNewParamsAccessYandexTypeSaml       AccessIdentityProviderNewParamsAccessYandexType = "saml"
	AccessIdentityProviderNewParamsAccessYandexTypeCentrify   AccessIdentityProviderNewParamsAccessYandexType = "centrify"
	AccessIdentityProviderNewParamsAccessYandexTypeFacebook   AccessIdentityProviderNewParamsAccessYandexType = "facebook"
	AccessIdentityProviderNewParamsAccessYandexTypeGitHub     AccessIdentityProviderNewParamsAccessYandexType = "github"
	AccessIdentityProviderNewParamsAccessYandexTypeGoogleApps AccessIdentityProviderNewParamsAccessYandexType = "google-apps"
	AccessIdentityProviderNewParamsAccessYandexTypeGoogle     AccessIdentityProviderNewParamsAccessYandexType = "google"
	AccessIdentityProviderNewParamsAccessYandexTypeLinkedin   AccessIdentityProviderNewParamsAccessYandexType = "linkedin"
	AccessIdentityProviderNewParamsAccessYandexTypeOidc       AccessIdentityProviderNewParamsAccessYandexType = "oidc"
	AccessIdentityProviderNewParamsAccessYandexTypeOkta       AccessIdentityProviderNewParamsAccessYandexType = "okta"
	AccessIdentityProviderNewParamsAccessYandexTypeOnelogin   AccessIdentityProviderNewParamsAccessYandexType = "onelogin"
	AccessIdentityProviderNewParamsAccessYandexTypePingone    AccessIdentityProviderNewParamsAccessYandexType = "pingone"
	AccessIdentityProviderNewParamsAccessYandexTypeYandex     AccessIdentityProviderNewParamsAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewParamsAccessYandexScimConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessYandexScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderNewParamsAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[interface{}] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderNewParamsAccessOnetimepinType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderNewParamsAccessOnetimepinScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderNewParamsAccessOnetimepin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderNewParamsAccessOnetimepin) ImplementsAccessIdentityProviderNewParams() {

}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderNewParamsAccessOnetimepinType string

const (
	AccessIdentityProviderNewParamsAccessOnetimepinTypeOnetimepin AccessIdentityProviderNewParamsAccessOnetimepinType = "onetimepin"
	AccessIdentityProviderNewParamsAccessOnetimepinTypeAzureAd    AccessIdentityProviderNewParamsAccessOnetimepinType = "azureAD"
	AccessIdentityProviderNewParamsAccessOnetimepinTypeSaml       AccessIdentityProviderNewParamsAccessOnetimepinType = "saml"
	AccessIdentityProviderNewParamsAccessOnetimepinTypeCentrify   AccessIdentityProviderNewParamsAccessOnetimepinType = "centrify"
	AccessIdentityProviderNewParamsAccessOnetimepinTypeFacebook   AccessIdentityProviderNewParamsAccessOnetimepinType = "facebook"
	AccessIdentityProviderNewParamsAccessOnetimepinTypeGitHub     AccessIdentityProviderNewParamsAccessOnetimepinType = "github"
	AccessIdentityProviderNewParamsAccessOnetimepinTypeGoogleApps AccessIdentityProviderNewParamsAccessOnetimepinType = "google-apps"
	AccessIdentityProviderNewParamsAccessOnetimepinTypeGoogle     AccessIdentityProviderNewParamsAccessOnetimepinType = "google"
	AccessIdentityProviderNewParamsAccessOnetimepinTypeLinkedin   AccessIdentityProviderNewParamsAccessOnetimepinType = "linkedin"
	AccessIdentityProviderNewParamsAccessOnetimepinTypeOidc       AccessIdentityProviderNewParamsAccessOnetimepinType = "oidc"
	AccessIdentityProviderNewParamsAccessOnetimepinTypeOkta       AccessIdentityProviderNewParamsAccessOnetimepinType = "okta"
	AccessIdentityProviderNewParamsAccessOnetimepinTypeOnelogin   AccessIdentityProviderNewParamsAccessOnetimepinType = "onelogin"
	AccessIdentityProviderNewParamsAccessOnetimepinTypePingone    AccessIdentityProviderNewParamsAccessOnetimepinType = "pingone"
	AccessIdentityProviderNewParamsAccessOnetimepinTypeYandex     AccessIdentityProviderNewParamsAccessOnetimepinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderNewParamsAccessOnetimepinScimConfig struct {
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

func (r AccessIdentityProviderNewParamsAccessOnetimepinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderNewResponseEnvelope struct {
	Errors   []AccessIdentityProviderNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessIdentityProviderNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessIdentityProviderNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessIdentityProviderNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessIdentityProviderNewResponseEnvelopeJSON    `json:"-"`
}

// accessIdentityProviderNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessIdentityProviderNewResponseEnvelope]
type accessIdentityProviderNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accessIdentityProviderNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessIdentityProviderNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessIdentityProviderNewResponseEnvelopeErrors]
type accessIdentityProviderNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderNewResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accessIdentityProviderNewResponseEnvelopeMessagesJSON `json:"-"`
}

// accessIdentityProviderNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessIdentityProviderNewResponseEnvelopeMessages]
type accessIdentityProviderNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessIdentityProviderNewResponseEnvelopeSuccess bool

const (
	AccessIdentityProviderNewResponseEnvelopeSuccessTrue AccessIdentityProviderNewResponseEnvelopeSuccess = true
)

type AccessIdentityProviderListResponseEnvelope struct {
	Errors   []AccessIdentityProviderListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessIdentityProviderListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessIdentityProviderListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessIdentityProviderListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessIdentityProviderListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessIdentityProviderListResponseEnvelopeJSON       `json:"-"`
}

// accessIdentityProviderListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessIdentityProviderListResponseEnvelope]
type accessIdentityProviderListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderListResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accessIdentityProviderListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessIdentityProviderListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessIdentityProviderListResponseEnvelopeErrors]
type accessIdentityProviderListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderListResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accessIdentityProviderListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessIdentityProviderListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseEnvelopeMessages]
type accessIdentityProviderListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessIdentityProviderListResponseEnvelopeSuccess bool

const (
	AccessIdentityProviderListResponseEnvelopeSuccessTrue AccessIdentityProviderListResponseEnvelopeSuccess = true
)

type AccessIdentityProviderListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       accessIdentityProviderListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessIdentityProviderListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AccessIdentityProviderListResponseEnvelopeResultInfo]
type accessIdentityProviderListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderDeleteResponseEnvelope struct {
	Errors   []AccessIdentityProviderDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessIdentityProviderDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessIdentityProviderDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessIdentityProviderDeleteResponseEnvelopeSuccess `json:"success,required"`
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

type AccessIdentityProviderGetResponseEnvelope struct {
	Errors   []AccessIdentityProviderGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessIdentityProviderGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessIdentityProviderGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessIdentityProviderGetResponseEnvelopeSuccess `json:"success,required"`
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
// [AccessIdentityProviderReplaceParamsAccessAzureAd],
// [AccessIdentityProviderReplaceParamsAccessCentrify],
// [AccessIdentityProviderReplaceParamsAccessFacebook],
// [AccessIdentityProviderReplaceParamsAccessGitHub],
// [AccessIdentityProviderReplaceParamsAccessGoogle],
// [AccessIdentityProviderReplaceParamsAccessGoogleApps],
// [AccessIdentityProviderReplaceParamsAccessLinkedin],
// [AccessIdentityProviderReplaceParamsAccessOidc],
// [AccessIdentityProviderReplaceParamsAccessOkta],
// [AccessIdentityProviderReplaceParamsAccessOnelogin],
// [AccessIdentityProviderReplaceParamsAccessPingone],
// [AccessIdentityProviderReplaceParamsAccessSaml],
// [AccessIdentityProviderReplaceParamsAccessYandex],
// [AccessIdentityProviderReplaceParamsAccessOnetimepin].
type AccessIdentityProviderReplaceParams interface {
	ImplementsAccessIdentityProviderReplaceParams()
}

type AccessIdentityProviderReplaceParamsAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderReplaceParamsAccessAzureAdConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderReplaceParamsAccessAzureAdType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderReplaceParamsAccessAzureAdScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderReplaceParamsAccessAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderReplaceParamsAccessAzureAd) ImplementsAccessIdentityProviderReplaceParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessAzureAdConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessAzureAdConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessAzureAdType string

const (
	AccessIdentityProviderReplaceParamsAccessAzureAdTypeOnetimepin AccessIdentityProviderReplaceParamsAccessAzureAdType = "onetimepin"
	AccessIdentityProviderReplaceParamsAccessAzureAdTypeAzureAd    AccessIdentityProviderReplaceParamsAccessAzureAdType = "azureAD"
	AccessIdentityProviderReplaceParamsAccessAzureAdTypeSaml       AccessIdentityProviderReplaceParamsAccessAzureAdType = "saml"
	AccessIdentityProviderReplaceParamsAccessAzureAdTypeCentrify   AccessIdentityProviderReplaceParamsAccessAzureAdType = "centrify"
	AccessIdentityProviderReplaceParamsAccessAzureAdTypeFacebook   AccessIdentityProviderReplaceParamsAccessAzureAdType = "facebook"
	AccessIdentityProviderReplaceParamsAccessAzureAdTypeGitHub     AccessIdentityProviderReplaceParamsAccessAzureAdType = "github"
	AccessIdentityProviderReplaceParamsAccessAzureAdTypeGoogleApps AccessIdentityProviderReplaceParamsAccessAzureAdType = "google-apps"
	AccessIdentityProviderReplaceParamsAccessAzureAdTypeGoogle     AccessIdentityProviderReplaceParamsAccessAzureAdType = "google"
	AccessIdentityProviderReplaceParamsAccessAzureAdTypeLinkedin   AccessIdentityProviderReplaceParamsAccessAzureAdType = "linkedin"
	AccessIdentityProviderReplaceParamsAccessAzureAdTypeOidc       AccessIdentityProviderReplaceParamsAccessAzureAdType = "oidc"
	AccessIdentityProviderReplaceParamsAccessAzureAdTypeOkta       AccessIdentityProviderReplaceParamsAccessAzureAdType = "okta"
	AccessIdentityProviderReplaceParamsAccessAzureAdTypeOnelogin   AccessIdentityProviderReplaceParamsAccessAzureAdType = "onelogin"
	AccessIdentityProviderReplaceParamsAccessAzureAdTypePingone    AccessIdentityProviderReplaceParamsAccessAzureAdType = "pingone"
	AccessIdentityProviderReplaceParamsAccessAzureAdTypeYandex     AccessIdentityProviderReplaceParamsAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceParamsAccessAzureAdScimConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessAzureAdScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderReplaceParamsAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderReplaceParamsAccessCentrifyConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderReplaceParamsAccessCentrifyType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderReplaceParamsAccessCentrifyScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderReplaceParamsAccessCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderReplaceParamsAccessCentrify) ImplementsAccessIdentityProviderReplaceParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessCentrifyConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessCentrifyType string

const (
	AccessIdentityProviderReplaceParamsAccessCentrifyTypeOnetimepin AccessIdentityProviderReplaceParamsAccessCentrifyType = "onetimepin"
	AccessIdentityProviderReplaceParamsAccessCentrifyTypeAzureAd    AccessIdentityProviderReplaceParamsAccessCentrifyType = "azureAD"
	AccessIdentityProviderReplaceParamsAccessCentrifyTypeSaml       AccessIdentityProviderReplaceParamsAccessCentrifyType = "saml"
	AccessIdentityProviderReplaceParamsAccessCentrifyTypeCentrify   AccessIdentityProviderReplaceParamsAccessCentrifyType = "centrify"
	AccessIdentityProviderReplaceParamsAccessCentrifyTypeFacebook   AccessIdentityProviderReplaceParamsAccessCentrifyType = "facebook"
	AccessIdentityProviderReplaceParamsAccessCentrifyTypeGitHub     AccessIdentityProviderReplaceParamsAccessCentrifyType = "github"
	AccessIdentityProviderReplaceParamsAccessCentrifyTypeGoogleApps AccessIdentityProviderReplaceParamsAccessCentrifyType = "google-apps"
	AccessIdentityProviderReplaceParamsAccessCentrifyTypeGoogle     AccessIdentityProviderReplaceParamsAccessCentrifyType = "google"
	AccessIdentityProviderReplaceParamsAccessCentrifyTypeLinkedin   AccessIdentityProviderReplaceParamsAccessCentrifyType = "linkedin"
	AccessIdentityProviderReplaceParamsAccessCentrifyTypeOidc       AccessIdentityProviderReplaceParamsAccessCentrifyType = "oidc"
	AccessIdentityProviderReplaceParamsAccessCentrifyTypeOkta       AccessIdentityProviderReplaceParamsAccessCentrifyType = "okta"
	AccessIdentityProviderReplaceParamsAccessCentrifyTypeOnelogin   AccessIdentityProviderReplaceParamsAccessCentrifyType = "onelogin"
	AccessIdentityProviderReplaceParamsAccessCentrifyTypePingone    AccessIdentityProviderReplaceParamsAccessCentrifyType = "pingone"
	AccessIdentityProviderReplaceParamsAccessCentrifyTypeYandex     AccessIdentityProviderReplaceParamsAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceParamsAccessCentrifyScimConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessCentrifyScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderReplaceParamsAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderReplaceParamsAccessFacebookConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderReplaceParamsAccessFacebookType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderReplaceParamsAccessFacebookScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderReplaceParamsAccessFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderReplaceParamsAccessFacebook) ImplementsAccessIdentityProviderReplaceParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderReplaceParamsAccessFacebookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessFacebookType string

const (
	AccessIdentityProviderReplaceParamsAccessFacebookTypeOnetimepin AccessIdentityProviderReplaceParamsAccessFacebookType = "onetimepin"
	AccessIdentityProviderReplaceParamsAccessFacebookTypeAzureAd    AccessIdentityProviderReplaceParamsAccessFacebookType = "azureAD"
	AccessIdentityProviderReplaceParamsAccessFacebookTypeSaml       AccessIdentityProviderReplaceParamsAccessFacebookType = "saml"
	AccessIdentityProviderReplaceParamsAccessFacebookTypeCentrify   AccessIdentityProviderReplaceParamsAccessFacebookType = "centrify"
	AccessIdentityProviderReplaceParamsAccessFacebookTypeFacebook   AccessIdentityProviderReplaceParamsAccessFacebookType = "facebook"
	AccessIdentityProviderReplaceParamsAccessFacebookTypeGitHub     AccessIdentityProviderReplaceParamsAccessFacebookType = "github"
	AccessIdentityProviderReplaceParamsAccessFacebookTypeGoogleApps AccessIdentityProviderReplaceParamsAccessFacebookType = "google-apps"
	AccessIdentityProviderReplaceParamsAccessFacebookTypeGoogle     AccessIdentityProviderReplaceParamsAccessFacebookType = "google"
	AccessIdentityProviderReplaceParamsAccessFacebookTypeLinkedin   AccessIdentityProviderReplaceParamsAccessFacebookType = "linkedin"
	AccessIdentityProviderReplaceParamsAccessFacebookTypeOidc       AccessIdentityProviderReplaceParamsAccessFacebookType = "oidc"
	AccessIdentityProviderReplaceParamsAccessFacebookTypeOkta       AccessIdentityProviderReplaceParamsAccessFacebookType = "okta"
	AccessIdentityProviderReplaceParamsAccessFacebookTypeOnelogin   AccessIdentityProviderReplaceParamsAccessFacebookType = "onelogin"
	AccessIdentityProviderReplaceParamsAccessFacebookTypePingone    AccessIdentityProviderReplaceParamsAccessFacebookType = "pingone"
	AccessIdentityProviderReplaceParamsAccessFacebookTypeYandex     AccessIdentityProviderReplaceParamsAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceParamsAccessFacebookScimConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessFacebookScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderReplaceParamsAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderReplaceParamsAccessGitHubConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderReplaceParamsAccessGitHubType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderReplaceParamsAccessGitHubScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderReplaceParamsAccessGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderReplaceParamsAccessGitHub) ImplementsAccessIdentityProviderReplaceParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderReplaceParamsAccessGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessGitHubType string

const (
	AccessIdentityProviderReplaceParamsAccessGitHubTypeOnetimepin AccessIdentityProviderReplaceParamsAccessGitHubType = "onetimepin"
	AccessIdentityProviderReplaceParamsAccessGitHubTypeAzureAd    AccessIdentityProviderReplaceParamsAccessGitHubType = "azureAD"
	AccessIdentityProviderReplaceParamsAccessGitHubTypeSaml       AccessIdentityProviderReplaceParamsAccessGitHubType = "saml"
	AccessIdentityProviderReplaceParamsAccessGitHubTypeCentrify   AccessIdentityProviderReplaceParamsAccessGitHubType = "centrify"
	AccessIdentityProviderReplaceParamsAccessGitHubTypeFacebook   AccessIdentityProviderReplaceParamsAccessGitHubType = "facebook"
	AccessIdentityProviderReplaceParamsAccessGitHubTypeGitHub     AccessIdentityProviderReplaceParamsAccessGitHubType = "github"
	AccessIdentityProviderReplaceParamsAccessGitHubTypeGoogleApps AccessIdentityProviderReplaceParamsAccessGitHubType = "google-apps"
	AccessIdentityProviderReplaceParamsAccessGitHubTypeGoogle     AccessIdentityProviderReplaceParamsAccessGitHubType = "google"
	AccessIdentityProviderReplaceParamsAccessGitHubTypeLinkedin   AccessIdentityProviderReplaceParamsAccessGitHubType = "linkedin"
	AccessIdentityProviderReplaceParamsAccessGitHubTypeOidc       AccessIdentityProviderReplaceParamsAccessGitHubType = "oidc"
	AccessIdentityProviderReplaceParamsAccessGitHubTypeOkta       AccessIdentityProviderReplaceParamsAccessGitHubType = "okta"
	AccessIdentityProviderReplaceParamsAccessGitHubTypeOnelogin   AccessIdentityProviderReplaceParamsAccessGitHubType = "onelogin"
	AccessIdentityProviderReplaceParamsAccessGitHubTypePingone    AccessIdentityProviderReplaceParamsAccessGitHubType = "pingone"
	AccessIdentityProviderReplaceParamsAccessGitHubTypeYandex     AccessIdentityProviderReplaceParamsAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceParamsAccessGitHubScimConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessGitHubScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderReplaceParamsAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderReplaceParamsAccessGoogleConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderReplaceParamsAccessGoogleType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderReplaceParamsAccessGoogleScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderReplaceParamsAccessGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderReplaceParamsAccessGoogle) ImplementsAccessIdentityProviderReplaceParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessGoogleConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r AccessIdentityProviderReplaceParamsAccessGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessGoogleType string

const (
	AccessIdentityProviderReplaceParamsAccessGoogleTypeOnetimepin AccessIdentityProviderReplaceParamsAccessGoogleType = "onetimepin"
	AccessIdentityProviderReplaceParamsAccessGoogleTypeAzureAd    AccessIdentityProviderReplaceParamsAccessGoogleType = "azureAD"
	AccessIdentityProviderReplaceParamsAccessGoogleTypeSaml       AccessIdentityProviderReplaceParamsAccessGoogleType = "saml"
	AccessIdentityProviderReplaceParamsAccessGoogleTypeCentrify   AccessIdentityProviderReplaceParamsAccessGoogleType = "centrify"
	AccessIdentityProviderReplaceParamsAccessGoogleTypeFacebook   AccessIdentityProviderReplaceParamsAccessGoogleType = "facebook"
	AccessIdentityProviderReplaceParamsAccessGoogleTypeGitHub     AccessIdentityProviderReplaceParamsAccessGoogleType = "github"
	AccessIdentityProviderReplaceParamsAccessGoogleTypeGoogleApps AccessIdentityProviderReplaceParamsAccessGoogleType = "google-apps"
	AccessIdentityProviderReplaceParamsAccessGoogleTypeGoogle     AccessIdentityProviderReplaceParamsAccessGoogleType = "google"
	AccessIdentityProviderReplaceParamsAccessGoogleTypeLinkedin   AccessIdentityProviderReplaceParamsAccessGoogleType = "linkedin"
	AccessIdentityProviderReplaceParamsAccessGoogleTypeOidc       AccessIdentityProviderReplaceParamsAccessGoogleType = "oidc"
	AccessIdentityProviderReplaceParamsAccessGoogleTypeOkta       AccessIdentityProviderReplaceParamsAccessGoogleType = "okta"
	AccessIdentityProviderReplaceParamsAccessGoogleTypeOnelogin   AccessIdentityProviderReplaceParamsAccessGoogleType = "onelogin"
	AccessIdentityProviderReplaceParamsAccessGoogleTypePingone    AccessIdentityProviderReplaceParamsAccessGoogleType = "pingone"
	AccessIdentityProviderReplaceParamsAccessGoogleTypeYandex     AccessIdentityProviderReplaceParamsAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceParamsAccessGoogleScimConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessGoogleScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderReplaceParamsAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderReplaceParamsAccessGoogleAppsConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderReplaceParamsAccessGoogleAppsType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderReplaceParamsAccessGoogleAppsScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderReplaceParamsAccessGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderReplaceParamsAccessGoogleApps) ImplementsAccessIdentityProviderReplaceParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessGoogleAppsConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessGoogleAppsType string

const (
	AccessIdentityProviderReplaceParamsAccessGoogleAppsTypeOnetimepin AccessIdentityProviderReplaceParamsAccessGoogleAppsType = "onetimepin"
	AccessIdentityProviderReplaceParamsAccessGoogleAppsTypeAzureAd    AccessIdentityProviderReplaceParamsAccessGoogleAppsType = "azureAD"
	AccessIdentityProviderReplaceParamsAccessGoogleAppsTypeSaml       AccessIdentityProviderReplaceParamsAccessGoogleAppsType = "saml"
	AccessIdentityProviderReplaceParamsAccessGoogleAppsTypeCentrify   AccessIdentityProviderReplaceParamsAccessGoogleAppsType = "centrify"
	AccessIdentityProviderReplaceParamsAccessGoogleAppsTypeFacebook   AccessIdentityProviderReplaceParamsAccessGoogleAppsType = "facebook"
	AccessIdentityProviderReplaceParamsAccessGoogleAppsTypeGitHub     AccessIdentityProviderReplaceParamsAccessGoogleAppsType = "github"
	AccessIdentityProviderReplaceParamsAccessGoogleAppsTypeGoogleApps AccessIdentityProviderReplaceParamsAccessGoogleAppsType = "google-apps"
	AccessIdentityProviderReplaceParamsAccessGoogleAppsTypeGoogle     AccessIdentityProviderReplaceParamsAccessGoogleAppsType = "google"
	AccessIdentityProviderReplaceParamsAccessGoogleAppsTypeLinkedin   AccessIdentityProviderReplaceParamsAccessGoogleAppsType = "linkedin"
	AccessIdentityProviderReplaceParamsAccessGoogleAppsTypeOidc       AccessIdentityProviderReplaceParamsAccessGoogleAppsType = "oidc"
	AccessIdentityProviderReplaceParamsAccessGoogleAppsTypeOkta       AccessIdentityProviderReplaceParamsAccessGoogleAppsType = "okta"
	AccessIdentityProviderReplaceParamsAccessGoogleAppsTypeOnelogin   AccessIdentityProviderReplaceParamsAccessGoogleAppsType = "onelogin"
	AccessIdentityProviderReplaceParamsAccessGoogleAppsTypePingone    AccessIdentityProviderReplaceParamsAccessGoogleAppsType = "pingone"
	AccessIdentityProviderReplaceParamsAccessGoogleAppsTypeYandex     AccessIdentityProviderReplaceParamsAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceParamsAccessGoogleAppsScimConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessGoogleAppsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderReplaceParamsAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderReplaceParamsAccessLinkedinConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderReplaceParamsAccessLinkedinType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderReplaceParamsAccessLinkedinScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderReplaceParamsAccessLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderReplaceParamsAccessLinkedin) ImplementsAccessIdentityProviderReplaceParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderReplaceParamsAccessLinkedinConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessLinkedinType string

const (
	AccessIdentityProviderReplaceParamsAccessLinkedinTypeOnetimepin AccessIdentityProviderReplaceParamsAccessLinkedinType = "onetimepin"
	AccessIdentityProviderReplaceParamsAccessLinkedinTypeAzureAd    AccessIdentityProviderReplaceParamsAccessLinkedinType = "azureAD"
	AccessIdentityProviderReplaceParamsAccessLinkedinTypeSaml       AccessIdentityProviderReplaceParamsAccessLinkedinType = "saml"
	AccessIdentityProviderReplaceParamsAccessLinkedinTypeCentrify   AccessIdentityProviderReplaceParamsAccessLinkedinType = "centrify"
	AccessIdentityProviderReplaceParamsAccessLinkedinTypeFacebook   AccessIdentityProviderReplaceParamsAccessLinkedinType = "facebook"
	AccessIdentityProviderReplaceParamsAccessLinkedinTypeGitHub     AccessIdentityProviderReplaceParamsAccessLinkedinType = "github"
	AccessIdentityProviderReplaceParamsAccessLinkedinTypeGoogleApps AccessIdentityProviderReplaceParamsAccessLinkedinType = "google-apps"
	AccessIdentityProviderReplaceParamsAccessLinkedinTypeGoogle     AccessIdentityProviderReplaceParamsAccessLinkedinType = "google"
	AccessIdentityProviderReplaceParamsAccessLinkedinTypeLinkedin   AccessIdentityProviderReplaceParamsAccessLinkedinType = "linkedin"
	AccessIdentityProviderReplaceParamsAccessLinkedinTypeOidc       AccessIdentityProviderReplaceParamsAccessLinkedinType = "oidc"
	AccessIdentityProviderReplaceParamsAccessLinkedinTypeOkta       AccessIdentityProviderReplaceParamsAccessLinkedinType = "okta"
	AccessIdentityProviderReplaceParamsAccessLinkedinTypeOnelogin   AccessIdentityProviderReplaceParamsAccessLinkedinType = "onelogin"
	AccessIdentityProviderReplaceParamsAccessLinkedinTypePingone    AccessIdentityProviderReplaceParamsAccessLinkedinType = "pingone"
	AccessIdentityProviderReplaceParamsAccessLinkedinTypeYandex     AccessIdentityProviderReplaceParamsAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceParamsAccessLinkedinScimConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessLinkedinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderReplaceParamsAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderReplaceParamsAccessOidcConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderReplaceParamsAccessOidcType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderReplaceParamsAccessOidcScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderReplaceParamsAccessOidc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderReplaceParamsAccessOidc) ImplementsAccessIdentityProviderReplaceParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessOidcConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessOidcConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessOidcType string

const (
	AccessIdentityProviderReplaceParamsAccessOidcTypeOnetimepin AccessIdentityProviderReplaceParamsAccessOidcType = "onetimepin"
	AccessIdentityProviderReplaceParamsAccessOidcTypeAzureAd    AccessIdentityProviderReplaceParamsAccessOidcType = "azureAD"
	AccessIdentityProviderReplaceParamsAccessOidcTypeSaml       AccessIdentityProviderReplaceParamsAccessOidcType = "saml"
	AccessIdentityProviderReplaceParamsAccessOidcTypeCentrify   AccessIdentityProviderReplaceParamsAccessOidcType = "centrify"
	AccessIdentityProviderReplaceParamsAccessOidcTypeFacebook   AccessIdentityProviderReplaceParamsAccessOidcType = "facebook"
	AccessIdentityProviderReplaceParamsAccessOidcTypeGitHub     AccessIdentityProviderReplaceParamsAccessOidcType = "github"
	AccessIdentityProviderReplaceParamsAccessOidcTypeGoogleApps AccessIdentityProviderReplaceParamsAccessOidcType = "google-apps"
	AccessIdentityProviderReplaceParamsAccessOidcTypeGoogle     AccessIdentityProviderReplaceParamsAccessOidcType = "google"
	AccessIdentityProviderReplaceParamsAccessOidcTypeLinkedin   AccessIdentityProviderReplaceParamsAccessOidcType = "linkedin"
	AccessIdentityProviderReplaceParamsAccessOidcTypeOidc       AccessIdentityProviderReplaceParamsAccessOidcType = "oidc"
	AccessIdentityProviderReplaceParamsAccessOidcTypeOkta       AccessIdentityProviderReplaceParamsAccessOidcType = "okta"
	AccessIdentityProviderReplaceParamsAccessOidcTypeOnelogin   AccessIdentityProviderReplaceParamsAccessOidcType = "onelogin"
	AccessIdentityProviderReplaceParamsAccessOidcTypePingone    AccessIdentityProviderReplaceParamsAccessOidcType = "pingone"
	AccessIdentityProviderReplaceParamsAccessOidcTypeYandex     AccessIdentityProviderReplaceParamsAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceParamsAccessOidcScimConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessOidcScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderReplaceParamsAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderReplaceParamsAccessOktaConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderReplaceParamsAccessOktaType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderReplaceParamsAccessOktaScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderReplaceParamsAccessOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderReplaceParamsAccessOkta) ImplementsAccessIdentityProviderReplaceParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessOktaConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessOktaType string

const (
	AccessIdentityProviderReplaceParamsAccessOktaTypeOnetimepin AccessIdentityProviderReplaceParamsAccessOktaType = "onetimepin"
	AccessIdentityProviderReplaceParamsAccessOktaTypeAzureAd    AccessIdentityProviderReplaceParamsAccessOktaType = "azureAD"
	AccessIdentityProviderReplaceParamsAccessOktaTypeSaml       AccessIdentityProviderReplaceParamsAccessOktaType = "saml"
	AccessIdentityProviderReplaceParamsAccessOktaTypeCentrify   AccessIdentityProviderReplaceParamsAccessOktaType = "centrify"
	AccessIdentityProviderReplaceParamsAccessOktaTypeFacebook   AccessIdentityProviderReplaceParamsAccessOktaType = "facebook"
	AccessIdentityProviderReplaceParamsAccessOktaTypeGitHub     AccessIdentityProviderReplaceParamsAccessOktaType = "github"
	AccessIdentityProviderReplaceParamsAccessOktaTypeGoogleApps AccessIdentityProviderReplaceParamsAccessOktaType = "google-apps"
	AccessIdentityProviderReplaceParamsAccessOktaTypeGoogle     AccessIdentityProviderReplaceParamsAccessOktaType = "google"
	AccessIdentityProviderReplaceParamsAccessOktaTypeLinkedin   AccessIdentityProviderReplaceParamsAccessOktaType = "linkedin"
	AccessIdentityProviderReplaceParamsAccessOktaTypeOidc       AccessIdentityProviderReplaceParamsAccessOktaType = "oidc"
	AccessIdentityProviderReplaceParamsAccessOktaTypeOkta       AccessIdentityProviderReplaceParamsAccessOktaType = "okta"
	AccessIdentityProviderReplaceParamsAccessOktaTypeOnelogin   AccessIdentityProviderReplaceParamsAccessOktaType = "onelogin"
	AccessIdentityProviderReplaceParamsAccessOktaTypePingone    AccessIdentityProviderReplaceParamsAccessOktaType = "pingone"
	AccessIdentityProviderReplaceParamsAccessOktaTypeYandex     AccessIdentityProviderReplaceParamsAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceParamsAccessOktaScimConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessOktaScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderReplaceParamsAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderReplaceParamsAccessOneloginConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderReplaceParamsAccessOneloginType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderReplaceParamsAccessOneloginScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderReplaceParamsAccessOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderReplaceParamsAccessOnelogin) ImplementsAccessIdentityProviderReplaceParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessOneloginConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessOneloginType string

const (
	AccessIdentityProviderReplaceParamsAccessOneloginTypeOnetimepin AccessIdentityProviderReplaceParamsAccessOneloginType = "onetimepin"
	AccessIdentityProviderReplaceParamsAccessOneloginTypeAzureAd    AccessIdentityProviderReplaceParamsAccessOneloginType = "azureAD"
	AccessIdentityProviderReplaceParamsAccessOneloginTypeSaml       AccessIdentityProviderReplaceParamsAccessOneloginType = "saml"
	AccessIdentityProviderReplaceParamsAccessOneloginTypeCentrify   AccessIdentityProviderReplaceParamsAccessOneloginType = "centrify"
	AccessIdentityProviderReplaceParamsAccessOneloginTypeFacebook   AccessIdentityProviderReplaceParamsAccessOneloginType = "facebook"
	AccessIdentityProviderReplaceParamsAccessOneloginTypeGitHub     AccessIdentityProviderReplaceParamsAccessOneloginType = "github"
	AccessIdentityProviderReplaceParamsAccessOneloginTypeGoogleApps AccessIdentityProviderReplaceParamsAccessOneloginType = "google-apps"
	AccessIdentityProviderReplaceParamsAccessOneloginTypeGoogle     AccessIdentityProviderReplaceParamsAccessOneloginType = "google"
	AccessIdentityProviderReplaceParamsAccessOneloginTypeLinkedin   AccessIdentityProviderReplaceParamsAccessOneloginType = "linkedin"
	AccessIdentityProviderReplaceParamsAccessOneloginTypeOidc       AccessIdentityProviderReplaceParamsAccessOneloginType = "oidc"
	AccessIdentityProviderReplaceParamsAccessOneloginTypeOkta       AccessIdentityProviderReplaceParamsAccessOneloginType = "okta"
	AccessIdentityProviderReplaceParamsAccessOneloginTypeOnelogin   AccessIdentityProviderReplaceParamsAccessOneloginType = "onelogin"
	AccessIdentityProviderReplaceParamsAccessOneloginTypePingone    AccessIdentityProviderReplaceParamsAccessOneloginType = "pingone"
	AccessIdentityProviderReplaceParamsAccessOneloginTypeYandex     AccessIdentityProviderReplaceParamsAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceParamsAccessOneloginScimConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessOneloginScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderReplaceParamsAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderReplaceParamsAccessPingoneConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderReplaceParamsAccessPingoneType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderReplaceParamsAccessPingoneScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderReplaceParamsAccessPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderReplaceParamsAccessPingone) ImplementsAccessIdentityProviderReplaceParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessPingoneConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessPingoneType string

const (
	AccessIdentityProviderReplaceParamsAccessPingoneTypeOnetimepin AccessIdentityProviderReplaceParamsAccessPingoneType = "onetimepin"
	AccessIdentityProviderReplaceParamsAccessPingoneTypeAzureAd    AccessIdentityProviderReplaceParamsAccessPingoneType = "azureAD"
	AccessIdentityProviderReplaceParamsAccessPingoneTypeSaml       AccessIdentityProviderReplaceParamsAccessPingoneType = "saml"
	AccessIdentityProviderReplaceParamsAccessPingoneTypeCentrify   AccessIdentityProviderReplaceParamsAccessPingoneType = "centrify"
	AccessIdentityProviderReplaceParamsAccessPingoneTypeFacebook   AccessIdentityProviderReplaceParamsAccessPingoneType = "facebook"
	AccessIdentityProviderReplaceParamsAccessPingoneTypeGitHub     AccessIdentityProviderReplaceParamsAccessPingoneType = "github"
	AccessIdentityProviderReplaceParamsAccessPingoneTypeGoogleApps AccessIdentityProviderReplaceParamsAccessPingoneType = "google-apps"
	AccessIdentityProviderReplaceParamsAccessPingoneTypeGoogle     AccessIdentityProviderReplaceParamsAccessPingoneType = "google"
	AccessIdentityProviderReplaceParamsAccessPingoneTypeLinkedin   AccessIdentityProviderReplaceParamsAccessPingoneType = "linkedin"
	AccessIdentityProviderReplaceParamsAccessPingoneTypeOidc       AccessIdentityProviderReplaceParamsAccessPingoneType = "oidc"
	AccessIdentityProviderReplaceParamsAccessPingoneTypeOkta       AccessIdentityProviderReplaceParamsAccessPingoneType = "okta"
	AccessIdentityProviderReplaceParamsAccessPingoneTypeOnelogin   AccessIdentityProviderReplaceParamsAccessPingoneType = "onelogin"
	AccessIdentityProviderReplaceParamsAccessPingoneTypePingone    AccessIdentityProviderReplaceParamsAccessPingoneType = "pingone"
	AccessIdentityProviderReplaceParamsAccessPingoneTypeYandex     AccessIdentityProviderReplaceParamsAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceParamsAccessPingoneScimConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessPingoneScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderReplaceParamsAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderReplaceParamsAccessSamlConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderReplaceParamsAccessSamlType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderReplaceParamsAccessSamlScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderReplaceParamsAccessSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderReplaceParamsAccessSaml) ImplementsAccessIdentityProviderReplaceParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]AccessIdentityProviderReplaceParamsAccessSamlConfigHeaderAttribute] `json:"header_attributes"`
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

func (r AccessIdentityProviderReplaceParamsAccessSamlConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderReplaceParamsAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r AccessIdentityProviderReplaceParamsAccessSamlConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessSamlType string

const (
	AccessIdentityProviderReplaceParamsAccessSamlTypeOnetimepin AccessIdentityProviderReplaceParamsAccessSamlType = "onetimepin"
	AccessIdentityProviderReplaceParamsAccessSamlTypeAzureAd    AccessIdentityProviderReplaceParamsAccessSamlType = "azureAD"
	AccessIdentityProviderReplaceParamsAccessSamlTypeSaml       AccessIdentityProviderReplaceParamsAccessSamlType = "saml"
	AccessIdentityProviderReplaceParamsAccessSamlTypeCentrify   AccessIdentityProviderReplaceParamsAccessSamlType = "centrify"
	AccessIdentityProviderReplaceParamsAccessSamlTypeFacebook   AccessIdentityProviderReplaceParamsAccessSamlType = "facebook"
	AccessIdentityProviderReplaceParamsAccessSamlTypeGitHub     AccessIdentityProviderReplaceParamsAccessSamlType = "github"
	AccessIdentityProviderReplaceParamsAccessSamlTypeGoogleApps AccessIdentityProviderReplaceParamsAccessSamlType = "google-apps"
	AccessIdentityProviderReplaceParamsAccessSamlTypeGoogle     AccessIdentityProviderReplaceParamsAccessSamlType = "google"
	AccessIdentityProviderReplaceParamsAccessSamlTypeLinkedin   AccessIdentityProviderReplaceParamsAccessSamlType = "linkedin"
	AccessIdentityProviderReplaceParamsAccessSamlTypeOidc       AccessIdentityProviderReplaceParamsAccessSamlType = "oidc"
	AccessIdentityProviderReplaceParamsAccessSamlTypeOkta       AccessIdentityProviderReplaceParamsAccessSamlType = "okta"
	AccessIdentityProviderReplaceParamsAccessSamlTypeOnelogin   AccessIdentityProviderReplaceParamsAccessSamlType = "onelogin"
	AccessIdentityProviderReplaceParamsAccessSamlTypePingone    AccessIdentityProviderReplaceParamsAccessSamlType = "pingone"
	AccessIdentityProviderReplaceParamsAccessSamlTypeYandex     AccessIdentityProviderReplaceParamsAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceParamsAccessSamlScimConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessSamlScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderReplaceParamsAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderReplaceParamsAccessYandexConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderReplaceParamsAccessYandexType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderReplaceParamsAccessYandexScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderReplaceParamsAccessYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderReplaceParamsAccessYandex) ImplementsAccessIdentityProviderReplaceParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderReplaceParamsAccessYandexConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessYandexType string

const (
	AccessIdentityProviderReplaceParamsAccessYandexTypeOnetimepin AccessIdentityProviderReplaceParamsAccessYandexType = "onetimepin"
	AccessIdentityProviderReplaceParamsAccessYandexTypeAzureAd    AccessIdentityProviderReplaceParamsAccessYandexType = "azureAD"
	AccessIdentityProviderReplaceParamsAccessYandexTypeSaml       AccessIdentityProviderReplaceParamsAccessYandexType = "saml"
	AccessIdentityProviderReplaceParamsAccessYandexTypeCentrify   AccessIdentityProviderReplaceParamsAccessYandexType = "centrify"
	AccessIdentityProviderReplaceParamsAccessYandexTypeFacebook   AccessIdentityProviderReplaceParamsAccessYandexType = "facebook"
	AccessIdentityProviderReplaceParamsAccessYandexTypeGitHub     AccessIdentityProviderReplaceParamsAccessYandexType = "github"
	AccessIdentityProviderReplaceParamsAccessYandexTypeGoogleApps AccessIdentityProviderReplaceParamsAccessYandexType = "google-apps"
	AccessIdentityProviderReplaceParamsAccessYandexTypeGoogle     AccessIdentityProviderReplaceParamsAccessYandexType = "google"
	AccessIdentityProviderReplaceParamsAccessYandexTypeLinkedin   AccessIdentityProviderReplaceParamsAccessYandexType = "linkedin"
	AccessIdentityProviderReplaceParamsAccessYandexTypeOidc       AccessIdentityProviderReplaceParamsAccessYandexType = "oidc"
	AccessIdentityProviderReplaceParamsAccessYandexTypeOkta       AccessIdentityProviderReplaceParamsAccessYandexType = "okta"
	AccessIdentityProviderReplaceParamsAccessYandexTypeOnelogin   AccessIdentityProviderReplaceParamsAccessYandexType = "onelogin"
	AccessIdentityProviderReplaceParamsAccessYandexTypePingone    AccessIdentityProviderReplaceParamsAccessYandexType = "pingone"
	AccessIdentityProviderReplaceParamsAccessYandexTypeYandex     AccessIdentityProviderReplaceParamsAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceParamsAccessYandexScimConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessYandexScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderReplaceParamsAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[interface{}] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderReplaceParamsAccessOnetimepinType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderReplaceParamsAccessOnetimepinScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderReplaceParamsAccessOnetimepin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderReplaceParamsAccessOnetimepin) ImplementsAccessIdentityProviderReplaceParams() {

}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderReplaceParamsAccessOnetimepinType string

const (
	AccessIdentityProviderReplaceParamsAccessOnetimepinTypeOnetimepin AccessIdentityProviderReplaceParamsAccessOnetimepinType = "onetimepin"
	AccessIdentityProviderReplaceParamsAccessOnetimepinTypeAzureAd    AccessIdentityProviderReplaceParamsAccessOnetimepinType = "azureAD"
	AccessIdentityProviderReplaceParamsAccessOnetimepinTypeSaml       AccessIdentityProviderReplaceParamsAccessOnetimepinType = "saml"
	AccessIdentityProviderReplaceParamsAccessOnetimepinTypeCentrify   AccessIdentityProviderReplaceParamsAccessOnetimepinType = "centrify"
	AccessIdentityProviderReplaceParamsAccessOnetimepinTypeFacebook   AccessIdentityProviderReplaceParamsAccessOnetimepinType = "facebook"
	AccessIdentityProviderReplaceParamsAccessOnetimepinTypeGitHub     AccessIdentityProviderReplaceParamsAccessOnetimepinType = "github"
	AccessIdentityProviderReplaceParamsAccessOnetimepinTypeGoogleApps AccessIdentityProviderReplaceParamsAccessOnetimepinType = "google-apps"
	AccessIdentityProviderReplaceParamsAccessOnetimepinTypeGoogle     AccessIdentityProviderReplaceParamsAccessOnetimepinType = "google"
	AccessIdentityProviderReplaceParamsAccessOnetimepinTypeLinkedin   AccessIdentityProviderReplaceParamsAccessOnetimepinType = "linkedin"
	AccessIdentityProviderReplaceParamsAccessOnetimepinTypeOidc       AccessIdentityProviderReplaceParamsAccessOnetimepinType = "oidc"
	AccessIdentityProviderReplaceParamsAccessOnetimepinTypeOkta       AccessIdentityProviderReplaceParamsAccessOnetimepinType = "okta"
	AccessIdentityProviderReplaceParamsAccessOnetimepinTypeOnelogin   AccessIdentityProviderReplaceParamsAccessOnetimepinType = "onelogin"
	AccessIdentityProviderReplaceParamsAccessOnetimepinTypePingone    AccessIdentityProviderReplaceParamsAccessOnetimepinType = "pingone"
	AccessIdentityProviderReplaceParamsAccessOnetimepinTypeYandex     AccessIdentityProviderReplaceParamsAccessOnetimepinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProviderReplaceParamsAccessOnetimepinScimConfig struct {
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

func (r AccessIdentityProviderReplaceParamsAccessOnetimepinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdentityProviderReplaceResponseEnvelope struct {
	Errors   []AccessIdentityProviderReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessIdentityProviderReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessIdentityProviderReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessIdentityProviderReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessIdentityProviderReplaceResponseEnvelopeJSON    `json:"-"`
}

// accessIdentityProviderReplaceResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessIdentityProviderReplaceResponseEnvelope]
type accessIdentityProviderReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accessIdentityProviderReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessIdentityProviderReplaceResponseEnvelopeErrors]
type accessIdentityProviderReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProviderReplaceResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    accessIdentityProviderReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// accessIdentityProviderReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessIdentityProviderReplaceResponseEnvelopeMessages]
type accessIdentityProviderReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessIdentityProviderReplaceResponseEnvelopeSuccess bool

const (
	AccessIdentityProviderReplaceResponseEnvelopeSuccessTrue AccessIdentityProviderReplaceResponseEnvelopeSuccess = true
)

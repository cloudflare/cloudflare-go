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
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderUpdateResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderUpdateResponseAccessAzureAdJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessAzureAdJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessAzureAd]
type accessIdentityProviderUpdateResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessAzureAd) implementsAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderUpdateResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderUpdateResponseAccessCentrifyJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessCentrifyJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessCentrify]
type accessIdentityProviderUpdateResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessCentrify) implementsAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderUpdateResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderUpdateResponseAccessFacebookJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessFacebookJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessFacebook]
type accessIdentityProviderUpdateResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessFacebook) implementsAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderUpdateResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderUpdateResponseAccessGitHubJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessGitHubJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessGitHub]
type accessIdentityProviderUpdateResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessGitHub) implementsAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderUpdateResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderUpdateResponseAccessGoogleJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessGoogleJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessGoogle]
type accessIdentityProviderUpdateResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessGoogle) implementsAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderUpdateResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderUpdateResponseAccessGoogleAppsJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessGoogleAppsJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessGoogleApps]
type accessIdentityProviderUpdateResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessGoogleApps) implementsAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderUpdateResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderUpdateResponseAccessLinkedinJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessLinkedinJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessLinkedin]
type accessIdentityProviderUpdateResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessLinkedin) implementsAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderUpdateResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderUpdateResponseAccessOidcJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessOidcJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessOidc]
type accessIdentityProviderUpdateResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessOidc) implementsAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderUpdateResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderUpdateResponseAccessOktaJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessOktaJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessOkta]
type accessIdentityProviderUpdateResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessOkta) implementsAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderUpdateResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderUpdateResponseAccessOneloginJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessOneloginJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessOnelogin]
type accessIdentityProviderUpdateResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessOnelogin) implementsAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderUpdateResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderUpdateResponseAccessPingoneJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessPingoneJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessPingone]
type accessIdentityProviderUpdateResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessPingone) implementsAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderUpdateResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderUpdateResponseAccessSamlJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessSamlJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessSaml]
type accessIdentityProviderUpdateResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessSaml) implementsAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProviderUpdateResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderUpdateResponseAccessYandexJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessYandexJSON contains the JSON metadata
// for the struct [AccessIdentityProviderUpdateResponseAccessYandex]
type accessIdentityProviderUpdateResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessYandex) implementsAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateResponseAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProviderUpdateResponseAccessOnetimepinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProviderUpdateResponseAccessOnetimepinScimConfig `json:"scim_config"`
	JSON       accessIdentityProviderUpdateResponseAccessOnetimepinJSON       `json:"-"`
}

// accessIdentityProviderUpdateResponseAccessOnetimepinJSON contains the JSON
// metadata for the struct [AccessIdentityProviderUpdateResponseAccessOnetimepin]
type accessIdentityProviderUpdateResponseAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProviderUpdateResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProviderUpdateResponseAccessOnetimepin) implementsAccessIdentityProviderUpdateResponse() {
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateResponseAccessOnetimepinType string

const (
	AccessIdentityProviderUpdateResponseAccessOnetimepinTypeOnetimepin AccessIdentityProviderUpdateResponseAccessOnetimepinType = "onetimepin"
	AccessIdentityProviderUpdateResponseAccessOnetimepinTypeAzureAd    AccessIdentityProviderUpdateResponseAccessOnetimepinType = "azureAD"
	AccessIdentityProviderUpdateResponseAccessOnetimepinTypeSaml       AccessIdentityProviderUpdateResponseAccessOnetimepinType = "saml"
	AccessIdentityProviderUpdateResponseAccessOnetimepinTypeCentrify   AccessIdentityProviderUpdateResponseAccessOnetimepinType = "centrify"
	AccessIdentityProviderUpdateResponseAccessOnetimepinTypeFacebook   AccessIdentityProviderUpdateResponseAccessOnetimepinType = "facebook"
	AccessIdentityProviderUpdateResponseAccessOnetimepinTypeGitHub     AccessIdentityProviderUpdateResponseAccessOnetimepinType = "github"
	AccessIdentityProviderUpdateResponseAccessOnetimepinTypeGoogleApps AccessIdentityProviderUpdateResponseAccessOnetimepinType = "google-apps"
	AccessIdentityProviderUpdateResponseAccessOnetimepinTypeGoogle     AccessIdentityProviderUpdateResponseAccessOnetimepinType = "google"
	AccessIdentityProviderUpdateResponseAccessOnetimepinTypeLinkedin   AccessIdentityProviderUpdateResponseAccessOnetimepinType = "linkedin"
	AccessIdentityProviderUpdateResponseAccessOnetimepinTypeOidc       AccessIdentityProviderUpdateResponseAccessOnetimepinType = "oidc"
	AccessIdentityProviderUpdateResponseAccessOnetimepinTypeOkta       AccessIdentityProviderUpdateResponseAccessOnetimepinType = "okta"
	AccessIdentityProviderUpdateResponseAccessOnetimepinTypeOnelogin   AccessIdentityProviderUpdateResponseAccessOnetimepinType = "onelogin"
	AccessIdentityProviderUpdateResponseAccessOnetimepinTypePingone    AccessIdentityProviderUpdateResponseAccessOnetimepinType = "pingone"
	AccessIdentityProviderUpdateResponseAccessOnetimepinTypeYandex     AccessIdentityProviderUpdateResponseAccessOnetimepinType = "yandex"
)

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
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderUpdateParamsAccessAzureAdConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessAzureAdType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessAzureAdScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderUpdateParamsAccessAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessAzureAd) ImplementsAccessIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateParamsAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderUpdateParamsAccessCentrifyConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessCentrifyType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessCentrifyScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderUpdateParamsAccessCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessCentrify) ImplementsAccessIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateParamsAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderUpdateParamsAccessFacebookConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessFacebookType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessFacebookScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderUpdateParamsAccessFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessFacebook) ImplementsAccessIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderUpdateParamsAccessFacebookConfig) MarshalJSON() (data []byte, err error) {
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

type AccessIdentityProviderUpdateParamsAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderUpdateParamsAccessGitHubConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessGitHubType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessGitHubScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderUpdateParamsAccessGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessGitHub) ImplementsAccessIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderUpdateParamsAccessGitHubConfig) MarshalJSON() (data []byte, err error) {
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

type AccessIdentityProviderUpdateParamsAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderUpdateParamsAccessGoogleConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessGoogleType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessGoogleScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderUpdateParamsAccessGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessGoogle) ImplementsAccessIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateParamsAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderUpdateParamsAccessGoogleAppsConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessGoogleAppsType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessGoogleAppsScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderUpdateParamsAccessGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessGoogleApps) ImplementsAccessIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateParamsAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderUpdateParamsAccessLinkedinConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessLinkedinType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessLinkedinScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderUpdateParamsAccessLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessLinkedin) ImplementsAccessIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderUpdateParamsAccessLinkedinConfig) MarshalJSON() (data []byte, err error) {
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

type AccessIdentityProviderUpdateParamsAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderUpdateParamsAccessOidcConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessOidcType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessOidcScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderUpdateParamsAccessOidc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessOidc) ImplementsAccessIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateParamsAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderUpdateParamsAccessOktaConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessOktaType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessOktaScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderUpdateParamsAccessOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessOkta) ImplementsAccessIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateParamsAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderUpdateParamsAccessOneloginConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessOneloginType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessOneloginScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderUpdateParamsAccessOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessOnelogin) ImplementsAccessIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateParamsAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderUpdateParamsAccessPingoneConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessPingoneType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessPingoneScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderUpdateParamsAccessPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessPingone) ImplementsAccessIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateParamsAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderUpdateParamsAccessSamlConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessSamlType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessSamlScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderUpdateParamsAccessSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessSaml) ImplementsAccessIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
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

type AccessIdentityProviderUpdateParamsAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AccessIdentityProviderUpdateParamsAccessYandexConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessYandexType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessYandexScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderUpdateParamsAccessYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessYandex) ImplementsAccessIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r AccessIdentityProviderUpdateParamsAccessYandexConfig) MarshalJSON() (data []byte, err error) {
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

type AccessIdentityProviderUpdateParamsAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[interface{}] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[AccessIdentityProviderUpdateParamsAccessOnetimepinType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[AccessIdentityProviderUpdateParamsAccessOnetimepinScimConfig] `json:"scim_config"`
}

func (r AccessIdentityProviderUpdateParamsAccessOnetimepin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccessIdentityProviderUpdateParamsAccessOnetimepin) ImplementsAccessIdentityProviderUpdateParams() {

}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProviderUpdateParamsAccessOnetimepinType string

const (
	AccessIdentityProviderUpdateParamsAccessOnetimepinTypeOnetimepin AccessIdentityProviderUpdateParamsAccessOnetimepinType = "onetimepin"
	AccessIdentityProviderUpdateParamsAccessOnetimepinTypeAzureAd    AccessIdentityProviderUpdateParamsAccessOnetimepinType = "azureAD"
	AccessIdentityProviderUpdateParamsAccessOnetimepinTypeSaml       AccessIdentityProviderUpdateParamsAccessOnetimepinType = "saml"
	AccessIdentityProviderUpdateParamsAccessOnetimepinTypeCentrify   AccessIdentityProviderUpdateParamsAccessOnetimepinType = "centrify"
	AccessIdentityProviderUpdateParamsAccessOnetimepinTypeFacebook   AccessIdentityProviderUpdateParamsAccessOnetimepinType = "facebook"
	AccessIdentityProviderUpdateParamsAccessOnetimepinTypeGitHub     AccessIdentityProviderUpdateParamsAccessOnetimepinType = "github"
	AccessIdentityProviderUpdateParamsAccessOnetimepinTypeGoogleApps AccessIdentityProviderUpdateParamsAccessOnetimepinType = "google-apps"
	AccessIdentityProviderUpdateParamsAccessOnetimepinTypeGoogle     AccessIdentityProviderUpdateParamsAccessOnetimepinType = "google"
	AccessIdentityProviderUpdateParamsAccessOnetimepinTypeLinkedin   AccessIdentityProviderUpdateParamsAccessOnetimepinType = "linkedin"
	AccessIdentityProviderUpdateParamsAccessOnetimepinTypeOidc       AccessIdentityProviderUpdateParamsAccessOnetimepinType = "oidc"
	AccessIdentityProviderUpdateParamsAccessOnetimepinTypeOkta       AccessIdentityProviderUpdateParamsAccessOnetimepinType = "okta"
	AccessIdentityProviderUpdateParamsAccessOnetimepinTypeOnelogin   AccessIdentityProviderUpdateParamsAccessOnetimepinType = "onelogin"
	AccessIdentityProviderUpdateParamsAccessOnetimepinTypePingone    AccessIdentityProviderUpdateParamsAccessOnetimepinType = "pingone"
	AccessIdentityProviderUpdateParamsAccessOnetimepinTypeYandex     AccessIdentityProviderUpdateParamsAccessOnetimepinType = "yandex"
)

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

type AccessIdentityProviderUpdateResponseEnvelope struct {
	Errors   []AccessIdentityProviderUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessIdentityProviderUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessIdentityProviderUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessIdentityProviderUpdateResponseEnvelopeSuccess `json:"success,required"`
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

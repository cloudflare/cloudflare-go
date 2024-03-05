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

// ZeroTrustIdentityProviderService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustIdentityProviderService] method instead.
type ZeroTrustIdentityProviderService struct {
	Options []option.RequestOption
}

// NewZeroTrustIdentityProviderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustIdentityProviderService(opts ...option.RequestOption) (r *ZeroTrustIdentityProviderService) {
	r = &ZeroTrustIdentityProviderService{}
	r.Options = opts
	return
}

// Adds a new identity provider to Access.
func (r *ZeroTrustIdentityProviderService) New(ctx context.Context, params ZeroTrustIdentityProviderNewParams, opts ...option.RequestOption) (res *AccessIdentityProviders, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustIdentityProviderNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/identity_providers", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured identity provider.
func (r *ZeroTrustIdentityProviderService) Update(ctx context.Context, uuid string, params ZeroTrustIdentityProviderUpdateParams, opts ...option.RequestOption) (res *AccessIdentityProviders, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustIdentityProviderUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/identity_providers/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all configured identity providers.
func (r *ZeroTrustIdentityProviderService) List(ctx context.Context, query ZeroTrustIdentityProviderListParams, opts ...option.RequestOption) (res *[]ZeroTrustIdentityProviderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustIdentityProviderListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/identity_providers", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an identity provider from Access.
func (r *ZeroTrustIdentityProviderService) Delete(ctx context.Context, uuid string, body ZeroTrustIdentityProviderDeleteParams, opts ...option.RequestOption) (res *ZeroTrustIdentityProviderDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustIdentityProviderDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/identity_providers/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a configured identity provider.
func (r *ZeroTrustIdentityProviderService) Get(ctx context.Context, uuid string, query ZeroTrustIdentityProviderGetParams, opts ...option.RequestOption) (res *AccessIdentityProviders, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustIdentityProviderGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/identity_providers/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AccessIdentityProvidersAccessAzureAd],
// [AccessIdentityProvidersAccessCentrify],
// [AccessIdentityProvidersAccessFacebook], [AccessIdentityProvidersAccessGitHub],
// [AccessIdentityProvidersAccessGoogle],
// [AccessIdentityProvidersAccessGoogleApps],
// [AccessIdentityProvidersAccessLinkedin], [AccessIdentityProvidersAccessOidc],
// [AccessIdentityProvidersAccessOkta], [AccessIdentityProvidersAccessOnelogin],
// [AccessIdentityProvidersAccessPingone], [AccessIdentityProvidersAccessSaml],
// [AccessIdentityProvidersAccessYandex] or
// [AccessIdentityProvidersAccessOnetimepin].
type AccessIdentityProviders interface {
	implementsAccessIdentityProviders()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessIdentityProviders)(nil)).Elem(), "")
}

type AccessIdentityProvidersAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProvidersAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProvidersAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProvidersAccessAzureAdScimConfig `json:"scim_config"`
	JSON       accessIdentityProvidersAccessAzureAdJSON       `json:"-"`
}

// accessIdentityProvidersAccessAzureAdJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessAzureAd]
type accessIdentityProvidersAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProvidersAccessAzureAd) implementsAccessIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessAzureAdConfig struct {
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
	SupportGroups bool                                           `json:"support_groups"`
	JSON          accessIdentityProvidersAccessAzureAdConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessAzureAdConfigJSON contains the JSON metadata for
// the struct [AccessIdentityProvidersAccessAzureAdConfig]
type accessIdentityProvidersAccessAzureAdConfigJSON struct {
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

func (r *AccessIdentityProvidersAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessAzureAdType string

const (
	AccessIdentityProvidersAccessAzureAdTypeOnetimepin AccessIdentityProvidersAccessAzureAdType = "onetimepin"
	AccessIdentityProvidersAccessAzureAdTypeAzureAd    AccessIdentityProvidersAccessAzureAdType = "azureAD"
	AccessIdentityProvidersAccessAzureAdTypeSaml       AccessIdentityProvidersAccessAzureAdType = "saml"
	AccessIdentityProvidersAccessAzureAdTypeCentrify   AccessIdentityProvidersAccessAzureAdType = "centrify"
	AccessIdentityProvidersAccessAzureAdTypeFacebook   AccessIdentityProvidersAccessAzureAdType = "facebook"
	AccessIdentityProvidersAccessAzureAdTypeGitHub     AccessIdentityProvidersAccessAzureAdType = "github"
	AccessIdentityProvidersAccessAzureAdTypeGoogleApps AccessIdentityProvidersAccessAzureAdType = "google-apps"
	AccessIdentityProvidersAccessAzureAdTypeGoogle     AccessIdentityProvidersAccessAzureAdType = "google"
	AccessIdentityProvidersAccessAzureAdTypeLinkedin   AccessIdentityProvidersAccessAzureAdType = "linkedin"
	AccessIdentityProvidersAccessAzureAdTypeOidc       AccessIdentityProvidersAccessAzureAdType = "oidc"
	AccessIdentityProvidersAccessAzureAdTypeOkta       AccessIdentityProvidersAccessAzureAdType = "okta"
	AccessIdentityProvidersAccessAzureAdTypeOnelogin   AccessIdentityProvidersAccessAzureAdType = "onelogin"
	AccessIdentityProvidersAccessAzureAdTypePingone    AccessIdentityProvidersAccessAzureAdType = "pingone"
	AccessIdentityProvidersAccessAzureAdTypeYandex     AccessIdentityProvidersAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProvidersAccessAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            accessIdentityProvidersAccessAzureAdScimConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessAzureAdScimConfigJSON contains the JSON metadata
// for the struct [AccessIdentityProvidersAccessAzureAdScimConfig]
type accessIdentityProvidersAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProvidersAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProvidersAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProvidersAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProvidersAccessCentrifyScimConfig `json:"scim_config"`
	JSON       accessIdentityProvidersAccessCentrifyJSON       `json:"-"`
}

// accessIdentityProvidersAccessCentrifyJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessCentrify]
type accessIdentityProvidersAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProvidersAccessCentrify) implementsAccessIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessCentrifyConfig struct {
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
	EmailClaimName string                                          `json:"email_claim_name"`
	JSON           accessIdentityProvidersAccessCentrifyConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessCentrifyConfigJSON contains the JSON metadata for
// the struct [AccessIdentityProvidersAccessCentrifyConfig]
type accessIdentityProvidersAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessCentrifyType string

const (
	AccessIdentityProvidersAccessCentrifyTypeOnetimepin AccessIdentityProvidersAccessCentrifyType = "onetimepin"
	AccessIdentityProvidersAccessCentrifyTypeAzureAd    AccessIdentityProvidersAccessCentrifyType = "azureAD"
	AccessIdentityProvidersAccessCentrifyTypeSaml       AccessIdentityProvidersAccessCentrifyType = "saml"
	AccessIdentityProvidersAccessCentrifyTypeCentrify   AccessIdentityProvidersAccessCentrifyType = "centrify"
	AccessIdentityProvidersAccessCentrifyTypeFacebook   AccessIdentityProvidersAccessCentrifyType = "facebook"
	AccessIdentityProvidersAccessCentrifyTypeGitHub     AccessIdentityProvidersAccessCentrifyType = "github"
	AccessIdentityProvidersAccessCentrifyTypeGoogleApps AccessIdentityProvidersAccessCentrifyType = "google-apps"
	AccessIdentityProvidersAccessCentrifyTypeGoogle     AccessIdentityProvidersAccessCentrifyType = "google"
	AccessIdentityProvidersAccessCentrifyTypeLinkedin   AccessIdentityProvidersAccessCentrifyType = "linkedin"
	AccessIdentityProvidersAccessCentrifyTypeOidc       AccessIdentityProvidersAccessCentrifyType = "oidc"
	AccessIdentityProvidersAccessCentrifyTypeOkta       AccessIdentityProvidersAccessCentrifyType = "okta"
	AccessIdentityProvidersAccessCentrifyTypeOnelogin   AccessIdentityProvidersAccessCentrifyType = "onelogin"
	AccessIdentityProvidersAccessCentrifyTypePingone    AccessIdentityProvidersAccessCentrifyType = "pingone"
	AccessIdentityProvidersAccessCentrifyTypeYandex     AccessIdentityProvidersAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProvidersAccessCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                `json:"user_deprovision"`
	JSON            accessIdentityProvidersAccessCentrifyScimConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessCentrifyScimConfigJSON contains the JSON metadata
// for the struct [AccessIdentityProvidersAccessCentrifyScimConfig]
type accessIdentityProvidersAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProvidersAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProvidersAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProvidersAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProvidersAccessFacebookScimConfig `json:"scim_config"`
	JSON       accessIdentityProvidersAccessFacebookJSON       `json:"-"`
}

// accessIdentityProvidersAccessFacebookJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessFacebook]
type accessIdentityProvidersAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProvidersAccessFacebook) implementsAccessIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                          `json:"client_secret"`
	JSON         accessIdentityProvidersAccessFacebookConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessFacebookConfigJSON contains the JSON metadata for
// the struct [AccessIdentityProvidersAccessFacebookConfig]
type accessIdentityProvidersAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessFacebookType string

const (
	AccessIdentityProvidersAccessFacebookTypeOnetimepin AccessIdentityProvidersAccessFacebookType = "onetimepin"
	AccessIdentityProvidersAccessFacebookTypeAzureAd    AccessIdentityProvidersAccessFacebookType = "azureAD"
	AccessIdentityProvidersAccessFacebookTypeSaml       AccessIdentityProvidersAccessFacebookType = "saml"
	AccessIdentityProvidersAccessFacebookTypeCentrify   AccessIdentityProvidersAccessFacebookType = "centrify"
	AccessIdentityProvidersAccessFacebookTypeFacebook   AccessIdentityProvidersAccessFacebookType = "facebook"
	AccessIdentityProvidersAccessFacebookTypeGitHub     AccessIdentityProvidersAccessFacebookType = "github"
	AccessIdentityProvidersAccessFacebookTypeGoogleApps AccessIdentityProvidersAccessFacebookType = "google-apps"
	AccessIdentityProvidersAccessFacebookTypeGoogle     AccessIdentityProvidersAccessFacebookType = "google"
	AccessIdentityProvidersAccessFacebookTypeLinkedin   AccessIdentityProvidersAccessFacebookType = "linkedin"
	AccessIdentityProvidersAccessFacebookTypeOidc       AccessIdentityProvidersAccessFacebookType = "oidc"
	AccessIdentityProvidersAccessFacebookTypeOkta       AccessIdentityProvidersAccessFacebookType = "okta"
	AccessIdentityProvidersAccessFacebookTypeOnelogin   AccessIdentityProvidersAccessFacebookType = "onelogin"
	AccessIdentityProvidersAccessFacebookTypePingone    AccessIdentityProvidersAccessFacebookType = "pingone"
	AccessIdentityProvidersAccessFacebookTypeYandex     AccessIdentityProvidersAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProvidersAccessFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                `json:"user_deprovision"`
	JSON            accessIdentityProvidersAccessFacebookScimConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessFacebookScimConfigJSON contains the JSON metadata
// for the struct [AccessIdentityProvidersAccessFacebookScimConfig]
type accessIdentityProvidersAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProvidersAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProvidersAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProvidersAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProvidersAccessGitHubScimConfig `json:"scim_config"`
	JSON       accessIdentityProvidersAccessGitHubJSON       `json:"-"`
}

// accessIdentityProvidersAccessGitHubJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessGitHub]
type accessIdentityProvidersAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProvidersAccessGitHub) implementsAccessIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                        `json:"client_secret"`
	JSON         accessIdentityProvidersAccessGitHubConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessGitHubConfigJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessGitHubConfig]
type accessIdentityProvidersAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessGitHubType string

const (
	AccessIdentityProvidersAccessGitHubTypeOnetimepin AccessIdentityProvidersAccessGitHubType = "onetimepin"
	AccessIdentityProvidersAccessGitHubTypeAzureAd    AccessIdentityProvidersAccessGitHubType = "azureAD"
	AccessIdentityProvidersAccessGitHubTypeSaml       AccessIdentityProvidersAccessGitHubType = "saml"
	AccessIdentityProvidersAccessGitHubTypeCentrify   AccessIdentityProvidersAccessGitHubType = "centrify"
	AccessIdentityProvidersAccessGitHubTypeFacebook   AccessIdentityProvidersAccessGitHubType = "facebook"
	AccessIdentityProvidersAccessGitHubTypeGitHub     AccessIdentityProvidersAccessGitHubType = "github"
	AccessIdentityProvidersAccessGitHubTypeGoogleApps AccessIdentityProvidersAccessGitHubType = "google-apps"
	AccessIdentityProvidersAccessGitHubTypeGoogle     AccessIdentityProvidersAccessGitHubType = "google"
	AccessIdentityProvidersAccessGitHubTypeLinkedin   AccessIdentityProvidersAccessGitHubType = "linkedin"
	AccessIdentityProvidersAccessGitHubTypeOidc       AccessIdentityProvidersAccessGitHubType = "oidc"
	AccessIdentityProvidersAccessGitHubTypeOkta       AccessIdentityProvidersAccessGitHubType = "okta"
	AccessIdentityProvidersAccessGitHubTypeOnelogin   AccessIdentityProvidersAccessGitHubType = "onelogin"
	AccessIdentityProvidersAccessGitHubTypePingone    AccessIdentityProvidersAccessGitHubType = "pingone"
	AccessIdentityProvidersAccessGitHubTypeYandex     AccessIdentityProvidersAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProvidersAccessGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                              `json:"user_deprovision"`
	JSON            accessIdentityProvidersAccessGitHubScimConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessGitHubScimConfigJSON contains the JSON metadata for
// the struct [AccessIdentityProvidersAccessGitHubScimConfig]
type accessIdentityProvidersAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProvidersAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProvidersAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProvidersAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProvidersAccessGoogleScimConfig `json:"scim_config"`
	JSON       accessIdentityProvidersAccessGoogleJSON       `json:"-"`
}

// accessIdentityProvidersAccessGoogleJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessGoogle]
type accessIdentityProvidersAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProvidersAccessGoogle) implementsAccessIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                        `json:"email_claim_name"`
	JSON           accessIdentityProvidersAccessGoogleConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessGoogleConfigJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessGoogleConfig]
type accessIdentityProvidersAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessGoogleType string

const (
	AccessIdentityProvidersAccessGoogleTypeOnetimepin AccessIdentityProvidersAccessGoogleType = "onetimepin"
	AccessIdentityProvidersAccessGoogleTypeAzureAd    AccessIdentityProvidersAccessGoogleType = "azureAD"
	AccessIdentityProvidersAccessGoogleTypeSaml       AccessIdentityProvidersAccessGoogleType = "saml"
	AccessIdentityProvidersAccessGoogleTypeCentrify   AccessIdentityProvidersAccessGoogleType = "centrify"
	AccessIdentityProvidersAccessGoogleTypeFacebook   AccessIdentityProvidersAccessGoogleType = "facebook"
	AccessIdentityProvidersAccessGoogleTypeGitHub     AccessIdentityProvidersAccessGoogleType = "github"
	AccessIdentityProvidersAccessGoogleTypeGoogleApps AccessIdentityProvidersAccessGoogleType = "google-apps"
	AccessIdentityProvidersAccessGoogleTypeGoogle     AccessIdentityProvidersAccessGoogleType = "google"
	AccessIdentityProvidersAccessGoogleTypeLinkedin   AccessIdentityProvidersAccessGoogleType = "linkedin"
	AccessIdentityProvidersAccessGoogleTypeOidc       AccessIdentityProvidersAccessGoogleType = "oidc"
	AccessIdentityProvidersAccessGoogleTypeOkta       AccessIdentityProvidersAccessGoogleType = "okta"
	AccessIdentityProvidersAccessGoogleTypeOnelogin   AccessIdentityProvidersAccessGoogleType = "onelogin"
	AccessIdentityProvidersAccessGoogleTypePingone    AccessIdentityProvidersAccessGoogleType = "pingone"
	AccessIdentityProvidersAccessGoogleTypeYandex     AccessIdentityProvidersAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProvidersAccessGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                              `json:"user_deprovision"`
	JSON            accessIdentityProvidersAccessGoogleScimConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessGoogleScimConfigJSON contains the JSON metadata for
// the struct [AccessIdentityProvidersAccessGoogleScimConfig]
type accessIdentityProvidersAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProvidersAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProvidersAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProvidersAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProvidersAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       accessIdentityProvidersAccessGoogleAppsJSON       `json:"-"`
}

// accessIdentityProvidersAccessGoogleAppsJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessGoogleApps]
type accessIdentityProvidersAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProvidersAccessGoogleApps) implementsAccessIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                            `json:"email_claim_name"`
	JSON           accessIdentityProvidersAccessGoogleAppsConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessGoogleAppsConfigJSON contains the JSON metadata for
// the struct [AccessIdentityProvidersAccessGoogleAppsConfig]
type accessIdentityProvidersAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessGoogleAppsType string

const (
	AccessIdentityProvidersAccessGoogleAppsTypeOnetimepin AccessIdentityProvidersAccessGoogleAppsType = "onetimepin"
	AccessIdentityProvidersAccessGoogleAppsTypeAzureAd    AccessIdentityProvidersAccessGoogleAppsType = "azureAD"
	AccessIdentityProvidersAccessGoogleAppsTypeSaml       AccessIdentityProvidersAccessGoogleAppsType = "saml"
	AccessIdentityProvidersAccessGoogleAppsTypeCentrify   AccessIdentityProvidersAccessGoogleAppsType = "centrify"
	AccessIdentityProvidersAccessGoogleAppsTypeFacebook   AccessIdentityProvidersAccessGoogleAppsType = "facebook"
	AccessIdentityProvidersAccessGoogleAppsTypeGitHub     AccessIdentityProvidersAccessGoogleAppsType = "github"
	AccessIdentityProvidersAccessGoogleAppsTypeGoogleApps AccessIdentityProvidersAccessGoogleAppsType = "google-apps"
	AccessIdentityProvidersAccessGoogleAppsTypeGoogle     AccessIdentityProvidersAccessGoogleAppsType = "google"
	AccessIdentityProvidersAccessGoogleAppsTypeLinkedin   AccessIdentityProvidersAccessGoogleAppsType = "linkedin"
	AccessIdentityProvidersAccessGoogleAppsTypeOidc       AccessIdentityProvidersAccessGoogleAppsType = "oidc"
	AccessIdentityProvidersAccessGoogleAppsTypeOkta       AccessIdentityProvidersAccessGoogleAppsType = "okta"
	AccessIdentityProvidersAccessGoogleAppsTypeOnelogin   AccessIdentityProvidersAccessGoogleAppsType = "onelogin"
	AccessIdentityProvidersAccessGoogleAppsTypePingone    AccessIdentityProvidersAccessGoogleAppsType = "pingone"
	AccessIdentityProvidersAccessGoogleAppsTypeYandex     AccessIdentityProvidersAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProvidersAccessGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            accessIdentityProvidersAccessGoogleAppsScimConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessGoogleAppsScimConfigJSON contains the JSON metadata
// for the struct [AccessIdentityProvidersAccessGoogleAppsScimConfig]
type accessIdentityProvidersAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProvidersAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProvidersAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProvidersAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProvidersAccessLinkedinScimConfig `json:"scim_config"`
	JSON       accessIdentityProvidersAccessLinkedinJSON       `json:"-"`
}

// accessIdentityProvidersAccessLinkedinJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessLinkedin]
type accessIdentityProvidersAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProvidersAccessLinkedin) implementsAccessIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                          `json:"client_secret"`
	JSON         accessIdentityProvidersAccessLinkedinConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessLinkedinConfigJSON contains the JSON metadata for
// the struct [AccessIdentityProvidersAccessLinkedinConfig]
type accessIdentityProvidersAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessLinkedinType string

const (
	AccessIdentityProvidersAccessLinkedinTypeOnetimepin AccessIdentityProvidersAccessLinkedinType = "onetimepin"
	AccessIdentityProvidersAccessLinkedinTypeAzureAd    AccessIdentityProvidersAccessLinkedinType = "azureAD"
	AccessIdentityProvidersAccessLinkedinTypeSaml       AccessIdentityProvidersAccessLinkedinType = "saml"
	AccessIdentityProvidersAccessLinkedinTypeCentrify   AccessIdentityProvidersAccessLinkedinType = "centrify"
	AccessIdentityProvidersAccessLinkedinTypeFacebook   AccessIdentityProvidersAccessLinkedinType = "facebook"
	AccessIdentityProvidersAccessLinkedinTypeGitHub     AccessIdentityProvidersAccessLinkedinType = "github"
	AccessIdentityProvidersAccessLinkedinTypeGoogleApps AccessIdentityProvidersAccessLinkedinType = "google-apps"
	AccessIdentityProvidersAccessLinkedinTypeGoogle     AccessIdentityProvidersAccessLinkedinType = "google"
	AccessIdentityProvidersAccessLinkedinTypeLinkedin   AccessIdentityProvidersAccessLinkedinType = "linkedin"
	AccessIdentityProvidersAccessLinkedinTypeOidc       AccessIdentityProvidersAccessLinkedinType = "oidc"
	AccessIdentityProvidersAccessLinkedinTypeOkta       AccessIdentityProvidersAccessLinkedinType = "okta"
	AccessIdentityProvidersAccessLinkedinTypeOnelogin   AccessIdentityProvidersAccessLinkedinType = "onelogin"
	AccessIdentityProvidersAccessLinkedinTypePingone    AccessIdentityProvidersAccessLinkedinType = "pingone"
	AccessIdentityProvidersAccessLinkedinTypeYandex     AccessIdentityProvidersAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProvidersAccessLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                `json:"user_deprovision"`
	JSON            accessIdentityProvidersAccessLinkedinScimConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessLinkedinScimConfigJSON contains the JSON metadata
// for the struct [AccessIdentityProvidersAccessLinkedinScimConfig]
type accessIdentityProvidersAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProvidersAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProvidersAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProvidersAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProvidersAccessOidcScimConfig `json:"scim_config"`
	JSON       accessIdentityProvidersAccessOidcJSON       `json:"-"`
}

// accessIdentityProvidersAccessOidcJSON contains the JSON metadata for the struct
// [AccessIdentityProvidersAccessOidc]
type accessIdentityProvidersAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProvidersAccessOidc) implementsAccessIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessOidcConfig struct {
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
	TokenURL string                                      `json:"token_url"`
	JSON     accessIdentityProvidersAccessOidcConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessOidcConfigJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessOidcConfig]
type accessIdentityProvidersAccessOidcConfigJSON struct {
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

func (r *AccessIdentityProvidersAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessOidcType string

const (
	AccessIdentityProvidersAccessOidcTypeOnetimepin AccessIdentityProvidersAccessOidcType = "onetimepin"
	AccessIdentityProvidersAccessOidcTypeAzureAd    AccessIdentityProvidersAccessOidcType = "azureAD"
	AccessIdentityProvidersAccessOidcTypeSaml       AccessIdentityProvidersAccessOidcType = "saml"
	AccessIdentityProvidersAccessOidcTypeCentrify   AccessIdentityProvidersAccessOidcType = "centrify"
	AccessIdentityProvidersAccessOidcTypeFacebook   AccessIdentityProvidersAccessOidcType = "facebook"
	AccessIdentityProvidersAccessOidcTypeGitHub     AccessIdentityProvidersAccessOidcType = "github"
	AccessIdentityProvidersAccessOidcTypeGoogleApps AccessIdentityProvidersAccessOidcType = "google-apps"
	AccessIdentityProvidersAccessOidcTypeGoogle     AccessIdentityProvidersAccessOidcType = "google"
	AccessIdentityProvidersAccessOidcTypeLinkedin   AccessIdentityProvidersAccessOidcType = "linkedin"
	AccessIdentityProvidersAccessOidcTypeOidc       AccessIdentityProvidersAccessOidcType = "oidc"
	AccessIdentityProvidersAccessOidcTypeOkta       AccessIdentityProvidersAccessOidcType = "okta"
	AccessIdentityProvidersAccessOidcTypeOnelogin   AccessIdentityProvidersAccessOidcType = "onelogin"
	AccessIdentityProvidersAccessOidcTypePingone    AccessIdentityProvidersAccessOidcType = "pingone"
	AccessIdentityProvidersAccessOidcTypeYandex     AccessIdentityProvidersAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProvidersAccessOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                            `json:"user_deprovision"`
	JSON            accessIdentityProvidersAccessOidcScimConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessOidcScimConfigJSON contains the JSON metadata for
// the struct [AccessIdentityProvidersAccessOidcScimConfig]
type accessIdentityProvidersAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProvidersAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProvidersAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProvidersAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProvidersAccessOktaScimConfig `json:"scim_config"`
	JSON       accessIdentityProvidersAccessOktaJSON       `json:"-"`
}

// accessIdentityProvidersAccessOktaJSON contains the JSON metadata for the struct
// [AccessIdentityProvidersAccessOkta]
type accessIdentityProvidersAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProvidersAccessOkta) implementsAccessIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessOktaConfig struct {
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
	OktaAccount string                                      `json:"okta_account"`
	JSON        accessIdentityProvidersAccessOktaConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessOktaConfigJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessOktaConfig]
type accessIdentityProvidersAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessOktaType string

const (
	AccessIdentityProvidersAccessOktaTypeOnetimepin AccessIdentityProvidersAccessOktaType = "onetimepin"
	AccessIdentityProvidersAccessOktaTypeAzureAd    AccessIdentityProvidersAccessOktaType = "azureAD"
	AccessIdentityProvidersAccessOktaTypeSaml       AccessIdentityProvidersAccessOktaType = "saml"
	AccessIdentityProvidersAccessOktaTypeCentrify   AccessIdentityProvidersAccessOktaType = "centrify"
	AccessIdentityProvidersAccessOktaTypeFacebook   AccessIdentityProvidersAccessOktaType = "facebook"
	AccessIdentityProvidersAccessOktaTypeGitHub     AccessIdentityProvidersAccessOktaType = "github"
	AccessIdentityProvidersAccessOktaTypeGoogleApps AccessIdentityProvidersAccessOktaType = "google-apps"
	AccessIdentityProvidersAccessOktaTypeGoogle     AccessIdentityProvidersAccessOktaType = "google"
	AccessIdentityProvidersAccessOktaTypeLinkedin   AccessIdentityProvidersAccessOktaType = "linkedin"
	AccessIdentityProvidersAccessOktaTypeOidc       AccessIdentityProvidersAccessOktaType = "oidc"
	AccessIdentityProvidersAccessOktaTypeOkta       AccessIdentityProvidersAccessOktaType = "okta"
	AccessIdentityProvidersAccessOktaTypeOnelogin   AccessIdentityProvidersAccessOktaType = "onelogin"
	AccessIdentityProvidersAccessOktaTypePingone    AccessIdentityProvidersAccessOktaType = "pingone"
	AccessIdentityProvidersAccessOktaTypeYandex     AccessIdentityProvidersAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProvidersAccessOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                            `json:"user_deprovision"`
	JSON            accessIdentityProvidersAccessOktaScimConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessOktaScimConfigJSON contains the JSON metadata for
// the struct [AccessIdentityProvidersAccessOktaScimConfig]
type accessIdentityProvidersAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProvidersAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProvidersAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProvidersAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProvidersAccessOneloginScimConfig `json:"scim_config"`
	JSON       accessIdentityProvidersAccessOneloginJSON       `json:"-"`
}

// accessIdentityProvidersAccessOneloginJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessOnelogin]
type accessIdentityProvidersAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProvidersAccessOnelogin) implementsAccessIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                          `json:"onelogin_account"`
	JSON            accessIdentityProvidersAccessOneloginConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessOneloginConfigJSON contains the JSON metadata for
// the struct [AccessIdentityProvidersAccessOneloginConfig]
type accessIdentityProvidersAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessOneloginType string

const (
	AccessIdentityProvidersAccessOneloginTypeOnetimepin AccessIdentityProvidersAccessOneloginType = "onetimepin"
	AccessIdentityProvidersAccessOneloginTypeAzureAd    AccessIdentityProvidersAccessOneloginType = "azureAD"
	AccessIdentityProvidersAccessOneloginTypeSaml       AccessIdentityProvidersAccessOneloginType = "saml"
	AccessIdentityProvidersAccessOneloginTypeCentrify   AccessIdentityProvidersAccessOneloginType = "centrify"
	AccessIdentityProvidersAccessOneloginTypeFacebook   AccessIdentityProvidersAccessOneloginType = "facebook"
	AccessIdentityProvidersAccessOneloginTypeGitHub     AccessIdentityProvidersAccessOneloginType = "github"
	AccessIdentityProvidersAccessOneloginTypeGoogleApps AccessIdentityProvidersAccessOneloginType = "google-apps"
	AccessIdentityProvidersAccessOneloginTypeGoogle     AccessIdentityProvidersAccessOneloginType = "google"
	AccessIdentityProvidersAccessOneloginTypeLinkedin   AccessIdentityProvidersAccessOneloginType = "linkedin"
	AccessIdentityProvidersAccessOneloginTypeOidc       AccessIdentityProvidersAccessOneloginType = "oidc"
	AccessIdentityProvidersAccessOneloginTypeOkta       AccessIdentityProvidersAccessOneloginType = "okta"
	AccessIdentityProvidersAccessOneloginTypeOnelogin   AccessIdentityProvidersAccessOneloginType = "onelogin"
	AccessIdentityProvidersAccessOneloginTypePingone    AccessIdentityProvidersAccessOneloginType = "pingone"
	AccessIdentityProvidersAccessOneloginTypeYandex     AccessIdentityProvidersAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProvidersAccessOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                `json:"user_deprovision"`
	JSON            accessIdentityProvidersAccessOneloginScimConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessOneloginScimConfigJSON contains the JSON metadata
// for the struct [AccessIdentityProvidersAccessOneloginScimConfig]
type accessIdentityProvidersAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProvidersAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProvidersAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProvidersAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProvidersAccessPingoneScimConfig `json:"scim_config"`
	JSON       accessIdentityProvidersAccessPingoneJSON       `json:"-"`
}

// accessIdentityProvidersAccessPingoneJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessPingone]
type accessIdentityProvidersAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProvidersAccessPingone) implementsAccessIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                         `json:"ping_env_id"`
	JSON      accessIdentityProvidersAccessPingoneConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessPingoneConfigJSON contains the JSON metadata for
// the struct [AccessIdentityProvidersAccessPingoneConfig]
type accessIdentityProvidersAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessPingoneType string

const (
	AccessIdentityProvidersAccessPingoneTypeOnetimepin AccessIdentityProvidersAccessPingoneType = "onetimepin"
	AccessIdentityProvidersAccessPingoneTypeAzureAd    AccessIdentityProvidersAccessPingoneType = "azureAD"
	AccessIdentityProvidersAccessPingoneTypeSaml       AccessIdentityProvidersAccessPingoneType = "saml"
	AccessIdentityProvidersAccessPingoneTypeCentrify   AccessIdentityProvidersAccessPingoneType = "centrify"
	AccessIdentityProvidersAccessPingoneTypeFacebook   AccessIdentityProvidersAccessPingoneType = "facebook"
	AccessIdentityProvidersAccessPingoneTypeGitHub     AccessIdentityProvidersAccessPingoneType = "github"
	AccessIdentityProvidersAccessPingoneTypeGoogleApps AccessIdentityProvidersAccessPingoneType = "google-apps"
	AccessIdentityProvidersAccessPingoneTypeGoogle     AccessIdentityProvidersAccessPingoneType = "google"
	AccessIdentityProvidersAccessPingoneTypeLinkedin   AccessIdentityProvidersAccessPingoneType = "linkedin"
	AccessIdentityProvidersAccessPingoneTypeOidc       AccessIdentityProvidersAccessPingoneType = "oidc"
	AccessIdentityProvidersAccessPingoneTypeOkta       AccessIdentityProvidersAccessPingoneType = "okta"
	AccessIdentityProvidersAccessPingoneTypeOnelogin   AccessIdentityProvidersAccessPingoneType = "onelogin"
	AccessIdentityProvidersAccessPingoneTypePingone    AccessIdentityProvidersAccessPingoneType = "pingone"
	AccessIdentityProvidersAccessPingoneTypeYandex     AccessIdentityProvidersAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProvidersAccessPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            accessIdentityProvidersAccessPingoneScimConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessPingoneScimConfigJSON contains the JSON metadata
// for the struct [AccessIdentityProvidersAccessPingoneScimConfig]
type accessIdentityProvidersAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProvidersAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProvidersAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProvidersAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProvidersAccessSamlScimConfig `json:"scim_config"`
	JSON       accessIdentityProvidersAccessSamlJSON       `json:"-"`
}

// accessIdentityProvidersAccessSamlJSON contains the JSON metadata for the struct
// [AccessIdentityProvidersAccessSaml]
type accessIdentityProvidersAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProvidersAccessSaml) implementsAccessIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []AccessIdentityProvidersAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                      `json:"sso_target_url"`
	JSON         accessIdentityProvidersAccessSamlConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessSamlConfigJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessSamlConfig]
type accessIdentityProvidersAccessSamlConfigJSON struct {
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

func (r *AccessIdentityProvidersAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProvidersAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                     `json:"header_name"`
	JSON       accessIdentityProvidersAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// accessIdentityProvidersAccessSamlConfigHeaderAttributeJSON contains the JSON
// metadata for the struct [AccessIdentityProvidersAccessSamlConfigHeaderAttribute]
type accessIdentityProvidersAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessSamlType string

const (
	AccessIdentityProvidersAccessSamlTypeOnetimepin AccessIdentityProvidersAccessSamlType = "onetimepin"
	AccessIdentityProvidersAccessSamlTypeAzureAd    AccessIdentityProvidersAccessSamlType = "azureAD"
	AccessIdentityProvidersAccessSamlTypeSaml       AccessIdentityProvidersAccessSamlType = "saml"
	AccessIdentityProvidersAccessSamlTypeCentrify   AccessIdentityProvidersAccessSamlType = "centrify"
	AccessIdentityProvidersAccessSamlTypeFacebook   AccessIdentityProvidersAccessSamlType = "facebook"
	AccessIdentityProvidersAccessSamlTypeGitHub     AccessIdentityProvidersAccessSamlType = "github"
	AccessIdentityProvidersAccessSamlTypeGoogleApps AccessIdentityProvidersAccessSamlType = "google-apps"
	AccessIdentityProvidersAccessSamlTypeGoogle     AccessIdentityProvidersAccessSamlType = "google"
	AccessIdentityProvidersAccessSamlTypeLinkedin   AccessIdentityProvidersAccessSamlType = "linkedin"
	AccessIdentityProvidersAccessSamlTypeOidc       AccessIdentityProvidersAccessSamlType = "oidc"
	AccessIdentityProvidersAccessSamlTypeOkta       AccessIdentityProvidersAccessSamlType = "okta"
	AccessIdentityProvidersAccessSamlTypeOnelogin   AccessIdentityProvidersAccessSamlType = "onelogin"
	AccessIdentityProvidersAccessSamlTypePingone    AccessIdentityProvidersAccessSamlType = "pingone"
	AccessIdentityProvidersAccessSamlTypeYandex     AccessIdentityProvidersAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProvidersAccessSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                            `json:"user_deprovision"`
	JSON            accessIdentityProvidersAccessSamlScimConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessSamlScimConfigJSON contains the JSON metadata for
// the struct [AccessIdentityProvidersAccessSamlScimConfig]
type accessIdentityProvidersAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProvidersAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AccessIdentityProvidersAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProvidersAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProvidersAccessYandexScimConfig `json:"scim_config"`
	JSON       accessIdentityProvidersAccessYandexJSON       `json:"-"`
}

// accessIdentityProvidersAccessYandexJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessYandex]
type accessIdentityProvidersAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProvidersAccessYandex) implementsAccessIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                        `json:"client_secret"`
	JSON         accessIdentityProvidersAccessYandexConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessYandexConfigJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessYandexConfig]
type accessIdentityProvidersAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessYandexType string

const (
	AccessIdentityProvidersAccessYandexTypeOnetimepin AccessIdentityProvidersAccessYandexType = "onetimepin"
	AccessIdentityProvidersAccessYandexTypeAzureAd    AccessIdentityProvidersAccessYandexType = "azureAD"
	AccessIdentityProvidersAccessYandexTypeSaml       AccessIdentityProvidersAccessYandexType = "saml"
	AccessIdentityProvidersAccessYandexTypeCentrify   AccessIdentityProvidersAccessYandexType = "centrify"
	AccessIdentityProvidersAccessYandexTypeFacebook   AccessIdentityProvidersAccessYandexType = "facebook"
	AccessIdentityProvidersAccessYandexTypeGitHub     AccessIdentityProvidersAccessYandexType = "github"
	AccessIdentityProvidersAccessYandexTypeGoogleApps AccessIdentityProvidersAccessYandexType = "google-apps"
	AccessIdentityProvidersAccessYandexTypeGoogle     AccessIdentityProvidersAccessYandexType = "google"
	AccessIdentityProvidersAccessYandexTypeLinkedin   AccessIdentityProvidersAccessYandexType = "linkedin"
	AccessIdentityProvidersAccessYandexTypeOidc       AccessIdentityProvidersAccessYandexType = "oidc"
	AccessIdentityProvidersAccessYandexTypeOkta       AccessIdentityProvidersAccessYandexType = "okta"
	AccessIdentityProvidersAccessYandexTypeOnelogin   AccessIdentityProvidersAccessYandexType = "onelogin"
	AccessIdentityProvidersAccessYandexTypePingone    AccessIdentityProvidersAccessYandexType = "pingone"
	AccessIdentityProvidersAccessYandexTypeYandex     AccessIdentityProvidersAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProvidersAccessYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                              `json:"user_deprovision"`
	JSON            accessIdentityProvidersAccessYandexScimConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessYandexScimConfigJSON contains the JSON metadata for
// the struct [AccessIdentityProvidersAccessYandexScimConfig]
type accessIdentityProvidersAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIdentityProvidersAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type AccessIdentityProvidersAccessOnetimepinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig AccessIdentityProvidersAccessOnetimepinScimConfig `json:"scim_config"`
	JSON       accessIdentityProvidersAccessOnetimepinJSON       `json:"-"`
}

// accessIdentityProvidersAccessOnetimepinJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessOnetimepin]
type accessIdentityProvidersAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessIdentityProvidersAccessOnetimepin) implementsAccessIdentityProviders() {}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessOnetimepinType string

const (
	AccessIdentityProvidersAccessOnetimepinTypeOnetimepin AccessIdentityProvidersAccessOnetimepinType = "onetimepin"
	AccessIdentityProvidersAccessOnetimepinTypeAzureAd    AccessIdentityProvidersAccessOnetimepinType = "azureAD"
	AccessIdentityProvidersAccessOnetimepinTypeSaml       AccessIdentityProvidersAccessOnetimepinType = "saml"
	AccessIdentityProvidersAccessOnetimepinTypeCentrify   AccessIdentityProvidersAccessOnetimepinType = "centrify"
	AccessIdentityProvidersAccessOnetimepinTypeFacebook   AccessIdentityProvidersAccessOnetimepinType = "facebook"
	AccessIdentityProvidersAccessOnetimepinTypeGitHub     AccessIdentityProvidersAccessOnetimepinType = "github"
	AccessIdentityProvidersAccessOnetimepinTypeGoogleApps AccessIdentityProvidersAccessOnetimepinType = "google-apps"
	AccessIdentityProvidersAccessOnetimepinTypeGoogle     AccessIdentityProvidersAccessOnetimepinType = "google"
	AccessIdentityProvidersAccessOnetimepinTypeLinkedin   AccessIdentityProvidersAccessOnetimepinType = "linkedin"
	AccessIdentityProvidersAccessOnetimepinTypeOidc       AccessIdentityProvidersAccessOnetimepinType = "oidc"
	AccessIdentityProvidersAccessOnetimepinTypeOkta       AccessIdentityProvidersAccessOnetimepinType = "okta"
	AccessIdentityProvidersAccessOnetimepinTypeOnelogin   AccessIdentityProvidersAccessOnetimepinType = "onelogin"
	AccessIdentityProvidersAccessOnetimepinTypePingone    AccessIdentityProvidersAccessOnetimepinType = "pingone"
	AccessIdentityProvidersAccessOnetimepinTypeYandex     AccessIdentityProvidersAccessOnetimepinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type AccessIdentityProvidersAccessOnetimepinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            accessIdentityProvidersAccessOnetimepinScimConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessOnetimepinScimConfigJSON contains the JSON metadata
// for the struct [AccessIdentityProvidersAccessOnetimepinScimConfig]
type accessIdentityProvidersAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZeroTrustIdentityProviderListResponseAccessAzureAd],
// [ZeroTrustIdentityProviderListResponseAccessCentrify],
// [ZeroTrustIdentityProviderListResponseAccessFacebook],
// [ZeroTrustIdentityProviderListResponseAccessGitHub],
// [ZeroTrustIdentityProviderListResponseAccessGoogle],
// [ZeroTrustIdentityProviderListResponseAccessGoogleApps],
// [ZeroTrustIdentityProviderListResponseAccessLinkedin],
// [ZeroTrustIdentityProviderListResponseAccessOidc],
// [ZeroTrustIdentityProviderListResponseAccessOkta],
// [ZeroTrustIdentityProviderListResponseAccessOnelogin],
// [ZeroTrustIdentityProviderListResponseAccessPingone],
// [ZeroTrustIdentityProviderListResponseAccessSaml] or
// [ZeroTrustIdentityProviderListResponseAccessYandex].
type ZeroTrustIdentityProviderListResponse interface {
	implementsZeroTrustIdentityProviderListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustIdentityProviderListResponse)(nil)).Elem(), "")
}

type ZeroTrustIdentityProviderListResponseAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderListResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderListResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderListResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderListResponseAccessAzureAdJSON       `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessAzureAdJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderListResponseAccessAzureAd]
type zeroTrustIdentityProviderListResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustIdentityProviderListResponseAccessAzureAd) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessAzureAdConfig struct {
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
	JSON          zeroTrustIdentityProviderListResponseAccessAzureAdConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessAzureAdConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessAzureAdConfig]
type zeroTrustIdentityProviderListResponseAccessAzureAdConfigJSON struct {
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

func (r *ZeroTrustIdentityProviderListResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessAzureAdType string

const (
	ZeroTrustIdentityProviderListResponseAccessAzureAdTypeOnetimepin ZeroTrustIdentityProviderListResponseAccessAzureAdType = "onetimepin"
	ZeroTrustIdentityProviderListResponseAccessAzureAdTypeAzureAd    ZeroTrustIdentityProviderListResponseAccessAzureAdType = "azureAD"
	ZeroTrustIdentityProviderListResponseAccessAzureAdTypeSaml       ZeroTrustIdentityProviderListResponseAccessAzureAdType = "saml"
	ZeroTrustIdentityProviderListResponseAccessAzureAdTypeCentrify   ZeroTrustIdentityProviderListResponseAccessAzureAdType = "centrify"
	ZeroTrustIdentityProviderListResponseAccessAzureAdTypeFacebook   ZeroTrustIdentityProviderListResponseAccessAzureAdType = "facebook"
	ZeroTrustIdentityProviderListResponseAccessAzureAdTypeGitHub     ZeroTrustIdentityProviderListResponseAccessAzureAdType = "github"
	ZeroTrustIdentityProviderListResponseAccessAzureAdTypeGoogleApps ZeroTrustIdentityProviderListResponseAccessAzureAdType = "google-apps"
	ZeroTrustIdentityProviderListResponseAccessAzureAdTypeGoogle     ZeroTrustIdentityProviderListResponseAccessAzureAdType = "google"
	ZeroTrustIdentityProviderListResponseAccessAzureAdTypeLinkedin   ZeroTrustIdentityProviderListResponseAccessAzureAdType = "linkedin"
	ZeroTrustIdentityProviderListResponseAccessAzureAdTypeOidc       ZeroTrustIdentityProviderListResponseAccessAzureAdType = "oidc"
	ZeroTrustIdentityProviderListResponseAccessAzureAdTypeOkta       ZeroTrustIdentityProviderListResponseAccessAzureAdType = "okta"
	ZeroTrustIdentityProviderListResponseAccessAzureAdTypeOnelogin   ZeroTrustIdentityProviderListResponseAccessAzureAdType = "onelogin"
	ZeroTrustIdentityProviderListResponseAccessAzureAdTypePingone    ZeroTrustIdentityProviderListResponseAccessAzureAdType = "pingone"
	ZeroTrustIdentityProviderListResponseAccessAzureAdTypeYandex     ZeroTrustIdentityProviderListResponseAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderListResponseAccessAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderListResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessAzureAdScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessAzureAdScimConfig]
type zeroTrustIdentityProviderListResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderListResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderListResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderListResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderListResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderListResponseAccessCentrifyJSON       `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessCentrifyJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderListResponseAccessCentrify]
type zeroTrustIdentityProviderListResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustIdentityProviderListResponseAccessCentrify) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessCentrifyConfig struct {
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
	JSON           zeroTrustIdentityProviderListResponseAccessCentrifyConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessCentrifyConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessCentrifyConfig]
type zeroTrustIdentityProviderListResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessCentrifyType string

const (
	ZeroTrustIdentityProviderListResponseAccessCentrifyTypeOnetimepin ZeroTrustIdentityProviderListResponseAccessCentrifyType = "onetimepin"
	ZeroTrustIdentityProviderListResponseAccessCentrifyTypeAzureAd    ZeroTrustIdentityProviderListResponseAccessCentrifyType = "azureAD"
	ZeroTrustIdentityProviderListResponseAccessCentrifyTypeSaml       ZeroTrustIdentityProviderListResponseAccessCentrifyType = "saml"
	ZeroTrustIdentityProviderListResponseAccessCentrifyTypeCentrify   ZeroTrustIdentityProviderListResponseAccessCentrifyType = "centrify"
	ZeroTrustIdentityProviderListResponseAccessCentrifyTypeFacebook   ZeroTrustIdentityProviderListResponseAccessCentrifyType = "facebook"
	ZeroTrustIdentityProviderListResponseAccessCentrifyTypeGitHub     ZeroTrustIdentityProviderListResponseAccessCentrifyType = "github"
	ZeroTrustIdentityProviderListResponseAccessCentrifyTypeGoogleApps ZeroTrustIdentityProviderListResponseAccessCentrifyType = "google-apps"
	ZeroTrustIdentityProviderListResponseAccessCentrifyTypeGoogle     ZeroTrustIdentityProviderListResponseAccessCentrifyType = "google"
	ZeroTrustIdentityProviderListResponseAccessCentrifyTypeLinkedin   ZeroTrustIdentityProviderListResponseAccessCentrifyType = "linkedin"
	ZeroTrustIdentityProviderListResponseAccessCentrifyTypeOidc       ZeroTrustIdentityProviderListResponseAccessCentrifyType = "oidc"
	ZeroTrustIdentityProviderListResponseAccessCentrifyTypeOkta       ZeroTrustIdentityProviderListResponseAccessCentrifyType = "okta"
	ZeroTrustIdentityProviderListResponseAccessCentrifyTypeOnelogin   ZeroTrustIdentityProviderListResponseAccessCentrifyType = "onelogin"
	ZeroTrustIdentityProviderListResponseAccessCentrifyTypePingone    ZeroTrustIdentityProviderListResponseAccessCentrifyType = "pingone"
	ZeroTrustIdentityProviderListResponseAccessCentrifyTypeYandex     ZeroTrustIdentityProviderListResponseAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderListResponseAccessCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderListResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessCentrifyScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessCentrifyScimConfig]
type zeroTrustIdentityProviderListResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderListResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderListResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderListResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderListResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderListResponseAccessFacebookJSON       `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessFacebookJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderListResponseAccessFacebook]
type zeroTrustIdentityProviderListResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustIdentityProviderListResponseAccessFacebook) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                        `json:"client_secret"`
	JSON         zeroTrustIdentityProviderListResponseAccessFacebookConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessFacebookConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessFacebookConfig]
type zeroTrustIdentityProviderListResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessFacebookType string

const (
	ZeroTrustIdentityProviderListResponseAccessFacebookTypeOnetimepin ZeroTrustIdentityProviderListResponseAccessFacebookType = "onetimepin"
	ZeroTrustIdentityProviderListResponseAccessFacebookTypeAzureAd    ZeroTrustIdentityProviderListResponseAccessFacebookType = "azureAD"
	ZeroTrustIdentityProviderListResponseAccessFacebookTypeSaml       ZeroTrustIdentityProviderListResponseAccessFacebookType = "saml"
	ZeroTrustIdentityProviderListResponseAccessFacebookTypeCentrify   ZeroTrustIdentityProviderListResponseAccessFacebookType = "centrify"
	ZeroTrustIdentityProviderListResponseAccessFacebookTypeFacebook   ZeroTrustIdentityProviderListResponseAccessFacebookType = "facebook"
	ZeroTrustIdentityProviderListResponseAccessFacebookTypeGitHub     ZeroTrustIdentityProviderListResponseAccessFacebookType = "github"
	ZeroTrustIdentityProviderListResponseAccessFacebookTypeGoogleApps ZeroTrustIdentityProviderListResponseAccessFacebookType = "google-apps"
	ZeroTrustIdentityProviderListResponseAccessFacebookTypeGoogle     ZeroTrustIdentityProviderListResponseAccessFacebookType = "google"
	ZeroTrustIdentityProviderListResponseAccessFacebookTypeLinkedin   ZeroTrustIdentityProviderListResponseAccessFacebookType = "linkedin"
	ZeroTrustIdentityProviderListResponseAccessFacebookTypeOidc       ZeroTrustIdentityProviderListResponseAccessFacebookType = "oidc"
	ZeroTrustIdentityProviderListResponseAccessFacebookTypeOkta       ZeroTrustIdentityProviderListResponseAccessFacebookType = "okta"
	ZeroTrustIdentityProviderListResponseAccessFacebookTypeOnelogin   ZeroTrustIdentityProviderListResponseAccessFacebookType = "onelogin"
	ZeroTrustIdentityProviderListResponseAccessFacebookTypePingone    ZeroTrustIdentityProviderListResponseAccessFacebookType = "pingone"
	ZeroTrustIdentityProviderListResponseAccessFacebookTypeYandex     ZeroTrustIdentityProviderListResponseAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderListResponseAccessFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderListResponseAccessFacebookScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessFacebookScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessFacebookScimConfig]
type zeroTrustIdentityProviderListResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderListResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderListResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderListResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderListResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderListResponseAccessGitHubJSON       `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessGitHubJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderListResponseAccessGitHub]
type zeroTrustIdentityProviderListResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustIdentityProviderListResponseAccessGitHub) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                      `json:"client_secret"`
	JSON         zeroTrustIdentityProviderListResponseAccessGitHubConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessGitHubConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessGitHubConfig]
type zeroTrustIdentityProviderListResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessGitHubType string

const (
	ZeroTrustIdentityProviderListResponseAccessGitHubTypeOnetimepin ZeroTrustIdentityProviderListResponseAccessGitHubType = "onetimepin"
	ZeroTrustIdentityProviderListResponseAccessGitHubTypeAzureAd    ZeroTrustIdentityProviderListResponseAccessGitHubType = "azureAD"
	ZeroTrustIdentityProviderListResponseAccessGitHubTypeSaml       ZeroTrustIdentityProviderListResponseAccessGitHubType = "saml"
	ZeroTrustIdentityProviderListResponseAccessGitHubTypeCentrify   ZeroTrustIdentityProviderListResponseAccessGitHubType = "centrify"
	ZeroTrustIdentityProviderListResponseAccessGitHubTypeFacebook   ZeroTrustIdentityProviderListResponseAccessGitHubType = "facebook"
	ZeroTrustIdentityProviderListResponseAccessGitHubTypeGitHub     ZeroTrustIdentityProviderListResponseAccessGitHubType = "github"
	ZeroTrustIdentityProviderListResponseAccessGitHubTypeGoogleApps ZeroTrustIdentityProviderListResponseAccessGitHubType = "google-apps"
	ZeroTrustIdentityProviderListResponseAccessGitHubTypeGoogle     ZeroTrustIdentityProviderListResponseAccessGitHubType = "google"
	ZeroTrustIdentityProviderListResponseAccessGitHubTypeLinkedin   ZeroTrustIdentityProviderListResponseAccessGitHubType = "linkedin"
	ZeroTrustIdentityProviderListResponseAccessGitHubTypeOidc       ZeroTrustIdentityProviderListResponseAccessGitHubType = "oidc"
	ZeroTrustIdentityProviderListResponseAccessGitHubTypeOkta       ZeroTrustIdentityProviderListResponseAccessGitHubType = "okta"
	ZeroTrustIdentityProviderListResponseAccessGitHubTypeOnelogin   ZeroTrustIdentityProviderListResponseAccessGitHubType = "onelogin"
	ZeroTrustIdentityProviderListResponseAccessGitHubTypePingone    ZeroTrustIdentityProviderListResponseAccessGitHubType = "pingone"
	ZeroTrustIdentityProviderListResponseAccessGitHubTypeYandex     ZeroTrustIdentityProviderListResponseAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderListResponseAccessGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderListResponseAccessGitHubScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessGitHubScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessGitHubScimConfig]
type zeroTrustIdentityProviderListResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderListResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderListResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderListResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderListResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderListResponseAccessGoogleJSON       `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessGoogleJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderListResponseAccessGoogle]
type zeroTrustIdentityProviderListResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustIdentityProviderListResponseAccessGoogle) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                      `json:"email_claim_name"`
	JSON           zeroTrustIdentityProviderListResponseAccessGoogleConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessGoogleConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessGoogleConfig]
type zeroTrustIdentityProviderListResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessGoogleType string

const (
	ZeroTrustIdentityProviderListResponseAccessGoogleTypeOnetimepin ZeroTrustIdentityProviderListResponseAccessGoogleType = "onetimepin"
	ZeroTrustIdentityProviderListResponseAccessGoogleTypeAzureAd    ZeroTrustIdentityProviderListResponseAccessGoogleType = "azureAD"
	ZeroTrustIdentityProviderListResponseAccessGoogleTypeSaml       ZeroTrustIdentityProviderListResponseAccessGoogleType = "saml"
	ZeroTrustIdentityProviderListResponseAccessGoogleTypeCentrify   ZeroTrustIdentityProviderListResponseAccessGoogleType = "centrify"
	ZeroTrustIdentityProviderListResponseAccessGoogleTypeFacebook   ZeroTrustIdentityProviderListResponseAccessGoogleType = "facebook"
	ZeroTrustIdentityProviderListResponseAccessGoogleTypeGitHub     ZeroTrustIdentityProviderListResponseAccessGoogleType = "github"
	ZeroTrustIdentityProviderListResponseAccessGoogleTypeGoogleApps ZeroTrustIdentityProviderListResponseAccessGoogleType = "google-apps"
	ZeroTrustIdentityProviderListResponseAccessGoogleTypeGoogle     ZeroTrustIdentityProviderListResponseAccessGoogleType = "google"
	ZeroTrustIdentityProviderListResponseAccessGoogleTypeLinkedin   ZeroTrustIdentityProviderListResponseAccessGoogleType = "linkedin"
	ZeroTrustIdentityProviderListResponseAccessGoogleTypeOidc       ZeroTrustIdentityProviderListResponseAccessGoogleType = "oidc"
	ZeroTrustIdentityProviderListResponseAccessGoogleTypeOkta       ZeroTrustIdentityProviderListResponseAccessGoogleType = "okta"
	ZeroTrustIdentityProviderListResponseAccessGoogleTypeOnelogin   ZeroTrustIdentityProviderListResponseAccessGoogleType = "onelogin"
	ZeroTrustIdentityProviderListResponseAccessGoogleTypePingone    ZeroTrustIdentityProviderListResponseAccessGoogleType = "pingone"
	ZeroTrustIdentityProviderListResponseAccessGoogleTypeYandex     ZeroTrustIdentityProviderListResponseAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderListResponseAccessGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderListResponseAccessGoogleScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessGoogleScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessGoogleScimConfig]
type zeroTrustIdentityProviderListResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderListResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderListResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderListResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderListResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderListResponseAccessGoogleAppsJSON       `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessGoogleAppsJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderListResponseAccessGoogleApps]
type zeroTrustIdentityProviderListResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustIdentityProviderListResponseAccessGoogleApps) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessGoogleAppsConfig struct {
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
	JSON           zeroTrustIdentityProviderListResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessGoogleAppsConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessGoogleAppsConfig]
type zeroTrustIdentityProviderListResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessGoogleAppsType string

const (
	ZeroTrustIdentityProviderListResponseAccessGoogleAppsTypeOnetimepin ZeroTrustIdentityProviderListResponseAccessGoogleAppsType = "onetimepin"
	ZeroTrustIdentityProviderListResponseAccessGoogleAppsTypeAzureAd    ZeroTrustIdentityProviderListResponseAccessGoogleAppsType = "azureAD"
	ZeroTrustIdentityProviderListResponseAccessGoogleAppsTypeSaml       ZeroTrustIdentityProviderListResponseAccessGoogleAppsType = "saml"
	ZeroTrustIdentityProviderListResponseAccessGoogleAppsTypeCentrify   ZeroTrustIdentityProviderListResponseAccessGoogleAppsType = "centrify"
	ZeroTrustIdentityProviderListResponseAccessGoogleAppsTypeFacebook   ZeroTrustIdentityProviderListResponseAccessGoogleAppsType = "facebook"
	ZeroTrustIdentityProviderListResponseAccessGoogleAppsTypeGitHub     ZeroTrustIdentityProviderListResponseAccessGoogleAppsType = "github"
	ZeroTrustIdentityProviderListResponseAccessGoogleAppsTypeGoogleApps ZeroTrustIdentityProviderListResponseAccessGoogleAppsType = "google-apps"
	ZeroTrustIdentityProviderListResponseAccessGoogleAppsTypeGoogle     ZeroTrustIdentityProviderListResponseAccessGoogleAppsType = "google"
	ZeroTrustIdentityProviderListResponseAccessGoogleAppsTypeLinkedin   ZeroTrustIdentityProviderListResponseAccessGoogleAppsType = "linkedin"
	ZeroTrustIdentityProviderListResponseAccessGoogleAppsTypeOidc       ZeroTrustIdentityProviderListResponseAccessGoogleAppsType = "oidc"
	ZeroTrustIdentityProviderListResponseAccessGoogleAppsTypeOkta       ZeroTrustIdentityProviderListResponseAccessGoogleAppsType = "okta"
	ZeroTrustIdentityProviderListResponseAccessGoogleAppsTypeOnelogin   ZeroTrustIdentityProviderListResponseAccessGoogleAppsType = "onelogin"
	ZeroTrustIdentityProviderListResponseAccessGoogleAppsTypePingone    ZeroTrustIdentityProviderListResponseAccessGoogleAppsType = "pingone"
	ZeroTrustIdentityProviderListResponseAccessGoogleAppsTypeYandex     ZeroTrustIdentityProviderListResponseAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderListResponseAccessGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderListResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessGoogleAppsScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessGoogleAppsScimConfig]
type zeroTrustIdentityProviderListResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderListResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderListResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderListResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderListResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderListResponseAccessLinkedinJSON       `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessLinkedinJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderListResponseAccessLinkedin]
type zeroTrustIdentityProviderListResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustIdentityProviderListResponseAccessLinkedin) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                        `json:"client_secret"`
	JSON         zeroTrustIdentityProviderListResponseAccessLinkedinConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessLinkedinConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessLinkedinConfig]
type zeroTrustIdentityProviderListResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessLinkedinType string

const (
	ZeroTrustIdentityProviderListResponseAccessLinkedinTypeOnetimepin ZeroTrustIdentityProviderListResponseAccessLinkedinType = "onetimepin"
	ZeroTrustIdentityProviderListResponseAccessLinkedinTypeAzureAd    ZeroTrustIdentityProviderListResponseAccessLinkedinType = "azureAD"
	ZeroTrustIdentityProviderListResponseAccessLinkedinTypeSaml       ZeroTrustIdentityProviderListResponseAccessLinkedinType = "saml"
	ZeroTrustIdentityProviderListResponseAccessLinkedinTypeCentrify   ZeroTrustIdentityProviderListResponseAccessLinkedinType = "centrify"
	ZeroTrustIdentityProviderListResponseAccessLinkedinTypeFacebook   ZeroTrustIdentityProviderListResponseAccessLinkedinType = "facebook"
	ZeroTrustIdentityProviderListResponseAccessLinkedinTypeGitHub     ZeroTrustIdentityProviderListResponseAccessLinkedinType = "github"
	ZeroTrustIdentityProviderListResponseAccessLinkedinTypeGoogleApps ZeroTrustIdentityProviderListResponseAccessLinkedinType = "google-apps"
	ZeroTrustIdentityProviderListResponseAccessLinkedinTypeGoogle     ZeroTrustIdentityProviderListResponseAccessLinkedinType = "google"
	ZeroTrustIdentityProviderListResponseAccessLinkedinTypeLinkedin   ZeroTrustIdentityProviderListResponseAccessLinkedinType = "linkedin"
	ZeroTrustIdentityProviderListResponseAccessLinkedinTypeOidc       ZeroTrustIdentityProviderListResponseAccessLinkedinType = "oidc"
	ZeroTrustIdentityProviderListResponseAccessLinkedinTypeOkta       ZeroTrustIdentityProviderListResponseAccessLinkedinType = "okta"
	ZeroTrustIdentityProviderListResponseAccessLinkedinTypeOnelogin   ZeroTrustIdentityProviderListResponseAccessLinkedinType = "onelogin"
	ZeroTrustIdentityProviderListResponseAccessLinkedinTypePingone    ZeroTrustIdentityProviderListResponseAccessLinkedinType = "pingone"
	ZeroTrustIdentityProviderListResponseAccessLinkedinTypeYandex     ZeroTrustIdentityProviderListResponseAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderListResponseAccessLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderListResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessLinkedinScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessLinkedinScimConfig]
type zeroTrustIdentityProviderListResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderListResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderListResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderListResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderListResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderListResponseAccessOidcJSON       `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessOidcJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderListResponseAccessOidc]
type zeroTrustIdentityProviderListResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustIdentityProviderListResponseAccessOidc) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessOidcConfig struct {
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
	JSON     zeroTrustIdentityProviderListResponseAccessOidcConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessOidcConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderListResponseAccessOidcConfig]
type zeroTrustIdentityProviderListResponseAccessOidcConfigJSON struct {
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

func (r *ZeroTrustIdentityProviderListResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessOidcType string

const (
	ZeroTrustIdentityProviderListResponseAccessOidcTypeOnetimepin ZeroTrustIdentityProviderListResponseAccessOidcType = "onetimepin"
	ZeroTrustIdentityProviderListResponseAccessOidcTypeAzureAd    ZeroTrustIdentityProviderListResponseAccessOidcType = "azureAD"
	ZeroTrustIdentityProviderListResponseAccessOidcTypeSaml       ZeroTrustIdentityProviderListResponseAccessOidcType = "saml"
	ZeroTrustIdentityProviderListResponseAccessOidcTypeCentrify   ZeroTrustIdentityProviderListResponseAccessOidcType = "centrify"
	ZeroTrustIdentityProviderListResponseAccessOidcTypeFacebook   ZeroTrustIdentityProviderListResponseAccessOidcType = "facebook"
	ZeroTrustIdentityProviderListResponseAccessOidcTypeGitHub     ZeroTrustIdentityProviderListResponseAccessOidcType = "github"
	ZeroTrustIdentityProviderListResponseAccessOidcTypeGoogleApps ZeroTrustIdentityProviderListResponseAccessOidcType = "google-apps"
	ZeroTrustIdentityProviderListResponseAccessOidcTypeGoogle     ZeroTrustIdentityProviderListResponseAccessOidcType = "google"
	ZeroTrustIdentityProviderListResponseAccessOidcTypeLinkedin   ZeroTrustIdentityProviderListResponseAccessOidcType = "linkedin"
	ZeroTrustIdentityProviderListResponseAccessOidcTypeOidc       ZeroTrustIdentityProviderListResponseAccessOidcType = "oidc"
	ZeroTrustIdentityProviderListResponseAccessOidcTypeOkta       ZeroTrustIdentityProviderListResponseAccessOidcType = "okta"
	ZeroTrustIdentityProviderListResponseAccessOidcTypeOnelogin   ZeroTrustIdentityProviderListResponseAccessOidcType = "onelogin"
	ZeroTrustIdentityProviderListResponseAccessOidcTypePingone    ZeroTrustIdentityProviderListResponseAccessOidcType = "pingone"
	ZeroTrustIdentityProviderListResponseAccessOidcTypeYandex     ZeroTrustIdentityProviderListResponseAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderListResponseAccessOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderListResponseAccessOidcScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessOidcScimConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessOidcScimConfig]
type zeroTrustIdentityProviderListResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderListResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderListResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderListResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderListResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderListResponseAccessOktaJSON       `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessOktaJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderListResponseAccessOkta]
type zeroTrustIdentityProviderListResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustIdentityProviderListResponseAccessOkta) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessOktaConfig struct {
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
	JSON        zeroTrustIdentityProviderListResponseAccessOktaConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessOktaConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderListResponseAccessOktaConfig]
type zeroTrustIdentityProviderListResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessOktaType string

const (
	ZeroTrustIdentityProviderListResponseAccessOktaTypeOnetimepin ZeroTrustIdentityProviderListResponseAccessOktaType = "onetimepin"
	ZeroTrustIdentityProviderListResponseAccessOktaTypeAzureAd    ZeroTrustIdentityProviderListResponseAccessOktaType = "azureAD"
	ZeroTrustIdentityProviderListResponseAccessOktaTypeSaml       ZeroTrustIdentityProviderListResponseAccessOktaType = "saml"
	ZeroTrustIdentityProviderListResponseAccessOktaTypeCentrify   ZeroTrustIdentityProviderListResponseAccessOktaType = "centrify"
	ZeroTrustIdentityProviderListResponseAccessOktaTypeFacebook   ZeroTrustIdentityProviderListResponseAccessOktaType = "facebook"
	ZeroTrustIdentityProviderListResponseAccessOktaTypeGitHub     ZeroTrustIdentityProviderListResponseAccessOktaType = "github"
	ZeroTrustIdentityProviderListResponseAccessOktaTypeGoogleApps ZeroTrustIdentityProviderListResponseAccessOktaType = "google-apps"
	ZeroTrustIdentityProviderListResponseAccessOktaTypeGoogle     ZeroTrustIdentityProviderListResponseAccessOktaType = "google"
	ZeroTrustIdentityProviderListResponseAccessOktaTypeLinkedin   ZeroTrustIdentityProviderListResponseAccessOktaType = "linkedin"
	ZeroTrustIdentityProviderListResponseAccessOktaTypeOidc       ZeroTrustIdentityProviderListResponseAccessOktaType = "oidc"
	ZeroTrustIdentityProviderListResponseAccessOktaTypeOkta       ZeroTrustIdentityProviderListResponseAccessOktaType = "okta"
	ZeroTrustIdentityProviderListResponseAccessOktaTypeOnelogin   ZeroTrustIdentityProviderListResponseAccessOktaType = "onelogin"
	ZeroTrustIdentityProviderListResponseAccessOktaTypePingone    ZeroTrustIdentityProviderListResponseAccessOktaType = "pingone"
	ZeroTrustIdentityProviderListResponseAccessOktaTypeYandex     ZeroTrustIdentityProviderListResponseAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderListResponseAccessOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderListResponseAccessOktaScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessOktaScimConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessOktaScimConfig]
type zeroTrustIdentityProviderListResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderListResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderListResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderListResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderListResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderListResponseAccessOneloginJSON       `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessOneloginJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderListResponseAccessOnelogin]
type zeroTrustIdentityProviderListResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustIdentityProviderListResponseAccessOnelogin) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessOneloginConfig struct {
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
	JSON            zeroTrustIdentityProviderListResponseAccessOneloginConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessOneloginConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessOneloginConfig]
type zeroTrustIdentityProviderListResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessOneloginType string

const (
	ZeroTrustIdentityProviderListResponseAccessOneloginTypeOnetimepin ZeroTrustIdentityProviderListResponseAccessOneloginType = "onetimepin"
	ZeroTrustIdentityProviderListResponseAccessOneloginTypeAzureAd    ZeroTrustIdentityProviderListResponseAccessOneloginType = "azureAD"
	ZeroTrustIdentityProviderListResponseAccessOneloginTypeSaml       ZeroTrustIdentityProviderListResponseAccessOneloginType = "saml"
	ZeroTrustIdentityProviderListResponseAccessOneloginTypeCentrify   ZeroTrustIdentityProviderListResponseAccessOneloginType = "centrify"
	ZeroTrustIdentityProviderListResponseAccessOneloginTypeFacebook   ZeroTrustIdentityProviderListResponseAccessOneloginType = "facebook"
	ZeroTrustIdentityProviderListResponseAccessOneloginTypeGitHub     ZeroTrustIdentityProviderListResponseAccessOneloginType = "github"
	ZeroTrustIdentityProviderListResponseAccessOneloginTypeGoogleApps ZeroTrustIdentityProviderListResponseAccessOneloginType = "google-apps"
	ZeroTrustIdentityProviderListResponseAccessOneloginTypeGoogle     ZeroTrustIdentityProviderListResponseAccessOneloginType = "google"
	ZeroTrustIdentityProviderListResponseAccessOneloginTypeLinkedin   ZeroTrustIdentityProviderListResponseAccessOneloginType = "linkedin"
	ZeroTrustIdentityProviderListResponseAccessOneloginTypeOidc       ZeroTrustIdentityProviderListResponseAccessOneloginType = "oidc"
	ZeroTrustIdentityProviderListResponseAccessOneloginTypeOkta       ZeroTrustIdentityProviderListResponseAccessOneloginType = "okta"
	ZeroTrustIdentityProviderListResponseAccessOneloginTypeOnelogin   ZeroTrustIdentityProviderListResponseAccessOneloginType = "onelogin"
	ZeroTrustIdentityProviderListResponseAccessOneloginTypePingone    ZeroTrustIdentityProviderListResponseAccessOneloginType = "pingone"
	ZeroTrustIdentityProviderListResponseAccessOneloginTypeYandex     ZeroTrustIdentityProviderListResponseAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderListResponseAccessOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderListResponseAccessOneloginScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessOneloginScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessOneloginScimConfig]
type zeroTrustIdentityProviderListResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderListResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderListResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderListResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderListResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderListResponseAccessPingoneJSON       `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessPingoneJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderListResponseAccessPingone]
type zeroTrustIdentityProviderListResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustIdentityProviderListResponseAccessPingone) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessPingoneConfig struct {
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
	JSON      zeroTrustIdentityProviderListResponseAccessPingoneConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessPingoneConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessPingoneConfig]
type zeroTrustIdentityProviderListResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessPingoneType string

const (
	ZeroTrustIdentityProviderListResponseAccessPingoneTypeOnetimepin ZeroTrustIdentityProviderListResponseAccessPingoneType = "onetimepin"
	ZeroTrustIdentityProviderListResponseAccessPingoneTypeAzureAd    ZeroTrustIdentityProviderListResponseAccessPingoneType = "azureAD"
	ZeroTrustIdentityProviderListResponseAccessPingoneTypeSaml       ZeroTrustIdentityProviderListResponseAccessPingoneType = "saml"
	ZeroTrustIdentityProviderListResponseAccessPingoneTypeCentrify   ZeroTrustIdentityProviderListResponseAccessPingoneType = "centrify"
	ZeroTrustIdentityProviderListResponseAccessPingoneTypeFacebook   ZeroTrustIdentityProviderListResponseAccessPingoneType = "facebook"
	ZeroTrustIdentityProviderListResponseAccessPingoneTypeGitHub     ZeroTrustIdentityProviderListResponseAccessPingoneType = "github"
	ZeroTrustIdentityProviderListResponseAccessPingoneTypeGoogleApps ZeroTrustIdentityProviderListResponseAccessPingoneType = "google-apps"
	ZeroTrustIdentityProviderListResponseAccessPingoneTypeGoogle     ZeroTrustIdentityProviderListResponseAccessPingoneType = "google"
	ZeroTrustIdentityProviderListResponseAccessPingoneTypeLinkedin   ZeroTrustIdentityProviderListResponseAccessPingoneType = "linkedin"
	ZeroTrustIdentityProviderListResponseAccessPingoneTypeOidc       ZeroTrustIdentityProviderListResponseAccessPingoneType = "oidc"
	ZeroTrustIdentityProviderListResponseAccessPingoneTypeOkta       ZeroTrustIdentityProviderListResponseAccessPingoneType = "okta"
	ZeroTrustIdentityProviderListResponseAccessPingoneTypeOnelogin   ZeroTrustIdentityProviderListResponseAccessPingoneType = "onelogin"
	ZeroTrustIdentityProviderListResponseAccessPingoneTypePingone    ZeroTrustIdentityProviderListResponseAccessPingoneType = "pingone"
	ZeroTrustIdentityProviderListResponseAccessPingoneTypeYandex     ZeroTrustIdentityProviderListResponseAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderListResponseAccessPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderListResponseAccessPingoneScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessPingoneScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessPingoneScimConfig]
type zeroTrustIdentityProviderListResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderListResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderListResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderListResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderListResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderListResponseAccessSamlJSON       `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessSamlJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderListResponseAccessSaml]
type zeroTrustIdentityProviderListResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustIdentityProviderListResponseAccessSaml) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []ZeroTrustIdentityProviderListResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                    `json:"sso_target_url"`
	JSON         zeroTrustIdentityProviderListResponseAccessSamlConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessSamlConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderListResponseAccessSamlConfig]
type zeroTrustIdentityProviderListResponseAccessSamlConfigJSON struct {
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

func (r *ZeroTrustIdentityProviderListResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderListResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                   `json:"header_name"`
	JSON       zeroTrustIdentityProviderListResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessSamlConfigHeaderAttribute]
type zeroTrustIdentityProviderListResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessSamlType string

const (
	ZeroTrustIdentityProviderListResponseAccessSamlTypeOnetimepin ZeroTrustIdentityProviderListResponseAccessSamlType = "onetimepin"
	ZeroTrustIdentityProviderListResponseAccessSamlTypeAzureAd    ZeroTrustIdentityProviderListResponseAccessSamlType = "azureAD"
	ZeroTrustIdentityProviderListResponseAccessSamlTypeSaml       ZeroTrustIdentityProviderListResponseAccessSamlType = "saml"
	ZeroTrustIdentityProviderListResponseAccessSamlTypeCentrify   ZeroTrustIdentityProviderListResponseAccessSamlType = "centrify"
	ZeroTrustIdentityProviderListResponseAccessSamlTypeFacebook   ZeroTrustIdentityProviderListResponseAccessSamlType = "facebook"
	ZeroTrustIdentityProviderListResponseAccessSamlTypeGitHub     ZeroTrustIdentityProviderListResponseAccessSamlType = "github"
	ZeroTrustIdentityProviderListResponseAccessSamlTypeGoogleApps ZeroTrustIdentityProviderListResponseAccessSamlType = "google-apps"
	ZeroTrustIdentityProviderListResponseAccessSamlTypeGoogle     ZeroTrustIdentityProviderListResponseAccessSamlType = "google"
	ZeroTrustIdentityProviderListResponseAccessSamlTypeLinkedin   ZeroTrustIdentityProviderListResponseAccessSamlType = "linkedin"
	ZeroTrustIdentityProviderListResponseAccessSamlTypeOidc       ZeroTrustIdentityProviderListResponseAccessSamlType = "oidc"
	ZeroTrustIdentityProviderListResponseAccessSamlTypeOkta       ZeroTrustIdentityProviderListResponseAccessSamlType = "okta"
	ZeroTrustIdentityProviderListResponseAccessSamlTypeOnelogin   ZeroTrustIdentityProviderListResponseAccessSamlType = "onelogin"
	ZeroTrustIdentityProviderListResponseAccessSamlTypePingone    ZeroTrustIdentityProviderListResponseAccessSamlType = "pingone"
	ZeroTrustIdentityProviderListResponseAccessSamlTypeYandex     ZeroTrustIdentityProviderListResponseAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderListResponseAccessSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderListResponseAccessSamlScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessSamlScimConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessSamlScimConfig]
type zeroTrustIdentityProviderListResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderListResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderListResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderListResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderListResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderListResponseAccessYandexJSON       `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessYandexJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderListResponseAccessYandex]
type zeroTrustIdentityProviderListResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustIdentityProviderListResponseAccessYandex) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                      `json:"client_secret"`
	JSON         zeroTrustIdentityProviderListResponseAccessYandexConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessYandexConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessYandexConfig]
type zeroTrustIdentityProviderListResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderListResponseAccessYandexType string

const (
	ZeroTrustIdentityProviderListResponseAccessYandexTypeOnetimepin ZeroTrustIdentityProviderListResponseAccessYandexType = "onetimepin"
	ZeroTrustIdentityProviderListResponseAccessYandexTypeAzureAd    ZeroTrustIdentityProviderListResponseAccessYandexType = "azureAD"
	ZeroTrustIdentityProviderListResponseAccessYandexTypeSaml       ZeroTrustIdentityProviderListResponseAccessYandexType = "saml"
	ZeroTrustIdentityProviderListResponseAccessYandexTypeCentrify   ZeroTrustIdentityProviderListResponseAccessYandexType = "centrify"
	ZeroTrustIdentityProviderListResponseAccessYandexTypeFacebook   ZeroTrustIdentityProviderListResponseAccessYandexType = "facebook"
	ZeroTrustIdentityProviderListResponseAccessYandexTypeGitHub     ZeroTrustIdentityProviderListResponseAccessYandexType = "github"
	ZeroTrustIdentityProviderListResponseAccessYandexTypeGoogleApps ZeroTrustIdentityProviderListResponseAccessYandexType = "google-apps"
	ZeroTrustIdentityProviderListResponseAccessYandexTypeGoogle     ZeroTrustIdentityProviderListResponseAccessYandexType = "google"
	ZeroTrustIdentityProviderListResponseAccessYandexTypeLinkedin   ZeroTrustIdentityProviderListResponseAccessYandexType = "linkedin"
	ZeroTrustIdentityProviderListResponseAccessYandexTypeOidc       ZeroTrustIdentityProviderListResponseAccessYandexType = "oidc"
	ZeroTrustIdentityProviderListResponseAccessYandexTypeOkta       ZeroTrustIdentityProviderListResponseAccessYandexType = "okta"
	ZeroTrustIdentityProviderListResponseAccessYandexTypeOnelogin   ZeroTrustIdentityProviderListResponseAccessYandexType = "onelogin"
	ZeroTrustIdentityProviderListResponseAccessYandexTypePingone    ZeroTrustIdentityProviderListResponseAccessYandexType = "pingone"
	ZeroTrustIdentityProviderListResponseAccessYandexTypeYandex     ZeroTrustIdentityProviderListResponseAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderListResponseAccessYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderListResponseAccessYandexScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseAccessYandexScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderListResponseAccessYandexScimConfig]
type zeroTrustIdentityProviderListResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderDeleteResponse struct {
	// UUID
	ID   string                                      `json:"id"`
	JSON zeroTrustIdentityProviderDeleteResponseJSON `json:"-"`
}

// zeroTrustIdentityProviderDeleteResponseJSON contains the JSON metadata for the
// struct [ZeroTrustIdentityProviderDeleteResponse]
type zeroTrustIdentityProviderDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderNewParams struct {
	Config param.Field[ZeroTrustIdentityProviderNewParamsConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZeroTrustIdentityProviderNewParamsType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID     param.Field[string]                                       `path:"zone_id"`
	ScimConfig param.Field[ZeroTrustIdentityProviderNewParamsScimConfig] `json:"scim_config"`
}

func (r ZeroTrustIdentityProviderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustIdentityProviderNewParamsConfig struct {
	// Your companies TLD
	AppsDomain param.Field[string] `json:"apps_domain"`
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The authorization_endpoint URL of your IdP
	AuthURL param.Field[string] `json:"auth_url"`
	// Your okta authorization server id
	AuthorizationServerID param.Field[string] `json:"authorization_server_id"`
	// Your centrify account url
	CentrifyAccount param.Field[string] `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID param.Field[string] `json:"centrify_app_id"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL param.Field[string] `json:"certs_url"`
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
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]ZeroTrustIdentityProviderNewParamsConfigHeaderAttribute] `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts param.Field[[]string] `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL param.Field[string] `json:"issuer_url"`
	// Your okta account url
	OktaAccount param.Field[string] `json:"okta_account"`
	// Your OneLogin account url
	OneloginAccount param.Field[string] `json:"onelogin_account"`
	// Your PingOne environment identifier
	PingEnvID param.Field[string] `json:"ping_env_id"`
	// OAuth scopes
	Scopes param.Field[[]string] `json:"scopes"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest param.Field[bool] `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL param.Field[string] `json:"sso_target_url"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
	// The token_endpoint URL of your IdP
	TokenURL param.Field[string] `json:"token_url"`
}

func (r ZeroTrustIdentityProviderNewParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustIdentityProviderNewParamsConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r ZeroTrustIdentityProviderNewParamsConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewParamsType string

const (
	ZeroTrustIdentityProviderNewParamsTypeOnetimepin ZeroTrustIdentityProviderNewParamsType = "onetimepin"
	ZeroTrustIdentityProviderNewParamsTypeAzureAd    ZeroTrustIdentityProviderNewParamsType = "azureAD"
	ZeroTrustIdentityProviderNewParamsTypeSaml       ZeroTrustIdentityProviderNewParamsType = "saml"
	ZeroTrustIdentityProviderNewParamsTypeCentrify   ZeroTrustIdentityProviderNewParamsType = "centrify"
	ZeroTrustIdentityProviderNewParamsTypeFacebook   ZeroTrustIdentityProviderNewParamsType = "facebook"
	ZeroTrustIdentityProviderNewParamsTypeGitHub     ZeroTrustIdentityProviderNewParamsType = "github"
	ZeroTrustIdentityProviderNewParamsTypeGoogleApps ZeroTrustIdentityProviderNewParamsType = "google-apps"
	ZeroTrustIdentityProviderNewParamsTypeGoogle     ZeroTrustIdentityProviderNewParamsType = "google"
	ZeroTrustIdentityProviderNewParamsTypeLinkedin   ZeroTrustIdentityProviderNewParamsType = "linkedin"
	ZeroTrustIdentityProviderNewParamsTypeOidc       ZeroTrustIdentityProviderNewParamsType = "oidc"
	ZeroTrustIdentityProviderNewParamsTypeOkta       ZeroTrustIdentityProviderNewParamsType = "okta"
	ZeroTrustIdentityProviderNewParamsTypeOnelogin   ZeroTrustIdentityProviderNewParamsType = "onelogin"
	ZeroTrustIdentityProviderNewParamsTypePingone    ZeroTrustIdentityProviderNewParamsType = "pingone"
	ZeroTrustIdentityProviderNewParamsTypeYandex     ZeroTrustIdentityProviderNewParamsType = "yandex"
)

type ZeroTrustIdentityProviderNewParamsScimConfig struct {
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

func (r ZeroTrustIdentityProviderNewParamsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustIdentityProviderNewResponseEnvelope struct {
	Errors   []ZeroTrustIdentityProviderNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustIdentityProviderNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessIdentityProviders                                `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustIdentityProviderNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustIdentityProviderNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustIdentityProviderNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustIdentityProviderNewResponseEnvelope]
type zeroTrustIdentityProviderNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderNewResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustIdentityProviderNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderNewResponseEnvelopeErrors]
type zeroTrustIdentityProviderNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderNewResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustIdentityProviderNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderNewResponseEnvelopeMessages]
type zeroTrustIdentityProviderNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustIdentityProviderNewResponseEnvelopeSuccess bool

const (
	ZeroTrustIdentityProviderNewResponseEnvelopeSuccessTrue ZeroTrustIdentityProviderNewResponseEnvelopeSuccess = true
)

type ZeroTrustIdentityProviderUpdateParams struct {
	Config param.Field[ZeroTrustIdentityProviderUpdateParamsConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZeroTrustIdentityProviderUpdateParamsType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID     param.Field[string]                                          `path:"zone_id"`
	ScimConfig param.Field[ZeroTrustIdentityProviderUpdateParamsScimConfig] `json:"scim_config"`
}

func (r ZeroTrustIdentityProviderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustIdentityProviderUpdateParamsConfig struct {
	// Your companies TLD
	AppsDomain param.Field[string] `json:"apps_domain"`
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The authorization_endpoint URL of your IdP
	AuthURL param.Field[string] `json:"auth_url"`
	// Your okta authorization server id
	AuthorizationServerID param.Field[string] `json:"authorization_server_id"`
	// Your centrify account url
	CentrifyAccount param.Field[string] `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID param.Field[string] `json:"centrify_app_id"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL param.Field[string] `json:"certs_url"`
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
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]ZeroTrustIdentityProviderUpdateParamsConfigHeaderAttribute] `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts param.Field[[]string] `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL param.Field[string] `json:"issuer_url"`
	// Your okta account url
	OktaAccount param.Field[string] `json:"okta_account"`
	// Your OneLogin account url
	OneloginAccount param.Field[string] `json:"onelogin_account"`
	// Your PingOne environment identifier
	PingEnvID param.Field[string] `json:"ping_env_id"`
	// OAuth scopes
	Scopes param.Field[[]string] `json:"scopes"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest param.Field[bool] `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL param.Field[string] `json:"sso_target_url"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
	// The token_endpoint URL of your IdP
	TokenURL param.Field[string] `json:"token_url"`
}

func (r ZeroTrustIdentityProviderUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustIdentityProviderUpdateParamsConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r ZeroTrustIdentityProviderUpdateParamsConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateParamsType string

const (
	ZeroTrustIdentityProviderUpdateParamsTypeOnetimepin ZeroTrustIdentityProviderUpdateParamsType = "onetimepin"
	ZeroTrustIdentityProviderUpdateParamsTypeAzureAd    ZeroTrustIdentityProviderUpdateParamsType = "azureAD"
	ZeroTrustIdentityProviderUpdateParamsTypeSaml       ZeroTrustIdentityProviderUpdateParamsType = "saml"
	ZeroTrustIdentityProviderUpdateParamsTypeCentrify   ZeroTrustIdentityProviderUpdateParamsType = "centrify"
	ZeroTrustIdentityProviderUpdateParamsTypeFacebook   ZeroTrustIdentityProviderUpdateParamsType = "facebook"
	ZeroTrustIdentityProviderUpdateParamsTypeGitHub     ZeroTrustIdentityProviderUpdateParamsType = "github"
	ZeroTrustIdentityProviderUpdateParamsTypeGoogleApps ZeroTrustIdentityProviderUpdateParamsType = "google-apps"
	ZeroTrustIdentityProviderUpdateParamsTypeGoogle     ZeroTrustIdentityProviderUpdateParamsType = "google"
	ZeroTrustIdentityProviderUpdateParamsTypeLinkedin   ZeroTrustIdentityProviderUpdateParamsType = "linkedin"
	ZeroTrustIdentityProviderUpdateParamsTypeOidc       ZeroTrustIdentityProviderUpdateParamsType = "oidc"
	ZeroTrustIdentityProviderUpdateParamsTypeOkta       ZeroTrustIdentityProviderUpdateParamsType = "okta"
	ZeroTrustIdentityProviderUpdateParamsTypeOnelogin   ZeroTrustIdentityProviderUpdateParamsType = "onelogin"
	ZeroTrustIdentityProviderUpdateParamsTypePingone    ZeroTrustIdentityProviderUpdateParamsType = "pingone"
	ZeroTrustIdentityProviderUpdateParamsTypeYandex     ZeroTrustIdentityProviderUpdateParamsType = "yandex"
)

type ZeroTrustIdentityProviderUpdateParamsScimConfig struct {
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

func (r ZeroTrustIdentityProviderUpdateParamsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustIdentityProviderUpdateResponseEnvelope struct {
	Errors   []ZeroTrustIdentityProviderUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustIdentityProviderUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessIdentityProviders                                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustIdentityProviderUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustIdentityProviderUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderUpdateResponseEnvelope]
type zeroTrustIdentityProviderUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderUpdateResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustIdentityProviderUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderUpdateResponseEnvelopeErrors]
type zeroTrustIdentityProviderUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderUpdateResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustIdentityProviderUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseEnvelopeMessages]
type zeroTrustIdentityProviderUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustIdentityProviderUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustIdentityProviderUpdateResponseEnvelopeSuccessTrue ZeroTrustIdentityProviderUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustIdentityProviderListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustIdentityProviderListResponseEnvelope struct {
	Errors   []ZeroTrustIdentityProviderListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustIdentityProviderListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustIdentityProviderListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustIdentityProviderListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustIdentityProviderListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustIdentityProviderListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustIdentityProviderListResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustIdentityProviderListResponseEnvelope]
type zeroTrustIdentityProviderListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderListResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustIdentityProviderListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderListResponseEnvelopeErrors]
type zeroTrustIdentityProviderListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderListResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustIdentityProviderListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderListResponseEnvelopeMessages]
type zeroTrustIdentityProviderListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustIdentityProviderListResponseEnvelopeSuccess bool

const (
	ZeroTrustIdentityProviderListResponseEnvelopeSuccessTrue ZeroTrustIdentityProviderListResponseEnvelopeSuccess = true
)

type ZeroTrustIdentityProviderListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       zeroTrustIdentityProviderListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustIdentityProviderListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderListResponseEnvelopeResultInfo]
type zeroTrustIdentityProviderListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustIdentityProviderDeleteResponseEnvelope struct {
	Errors   []ZeroTrustIdentityProviderDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustIdentityProviderDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustIdentityProviderDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustIdentityProviderDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustIdentityProviderDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustIdentityProviderDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderDeleteResponseEnvelope]
type zeroTrustIdentityProviderDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderDeleteResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustIdentityProviderDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustIdentityProviderDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderDeleteResponseEnvelopeErrors]
type zeroTrustIdentityProviderDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderDeleteResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustIdentityProviderDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustIdentityProviderDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderDeleteResponseEnvelopeMessages]
type zeroTrustIdentityProviderDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustIdentityProviderDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustIdentityProviderDeleteResponseEnvelopeSuccessTrue ZeroTrustIdentityProviderDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustIdentityProviderGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustIdentityProviderGetResponseEnvelope struct {
	Errors   []ZeroTrustIdentityProviderGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustIdentityProviderGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessIdentityProviders                                `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustIdentityProviderGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustIdentityProviderGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustIdentityProviderGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustIdentityProviderGetResponseEnvelope]
type zeroTrustIdentityProviderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustIdentityProviderGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderGetResponseEnvelopeErrors]
type zeroTrustIdentityProviderGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustIdentityProviderGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustIdentityProviderGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderGetResponseEnvelopeMessages]
type zeroTrustIdentityProviderGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustIdentityProviderGetResponseEnvelopeSuccess bool

const (
	ZeroTrustIdentityProviderGetResponseEnvelopeSuccessTrue ZeroTrustIdentityProviderGetResponseEnvelopeSuccess = true
)
